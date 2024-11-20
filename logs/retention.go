package logs

import (
	"fmt"
	"time"

	"github.com/seccretoapp/tracecrow/model"
)

// RetentionCriteria define os critérios para retenção de logs
type RetentionCriteria struct {
	RetentionPeriod time.Duration
	LogLevel        string // Representado como string para compatibilidade com o modelo
	Environment     string
}

// RetentionResult contém informações sobre os resultados da retenção
type RetentionResult struct {
	RetainedLogs     []Log
	DiscardedLogs    []Log
	TotalLogs        int
	RetainedCount    int
	DiscardedCount   int
	RetentionApplied bool
}

func RetainLogs(logs []Log, criteria RetentionCriteria) (RetentionResult, error) {
	var retainedLogs, discardedLogs []Log

	// Tempo atual em segundos (UNIX epoch)
	currentTime := time.Now().Unix()
	retentionThreshold := currentTime - int64(criteria.RetentionPeriod.Seconds())

	// Converte critérios de string para tipos do modelo
	logLevel, logLevelErr := model.FromStringToLogLevel(criteria.LogLevel)
	env, envErr := model.FromStringToEnvironment(criteria.Environment)

	// Validação de critérios
	if logLevelErr != nil && criteria.LogLevel != "" {
		return RetentionResult{}, fmt.Errorf("invalid log level: %v", logLevelErr)
	}
	if envErr != nil && criteria.Environment != "" {
		return RetentionResult{}, fmt.Errorf("invalid environment: %v", envErr)
	}

	// Itera sobre os logs e aplica os critérios de retenção
	for _, log := range logs {
		// Verifica se o log está dentro do período de retenção
		isWithinRetention := log.Header.Timestamp >= retentionThreshold

		// Verifica critérios adicionais (LogLevel e Environment)
		matchesLogLevel := logLevelErr == nil && (criteria.LogLevel == "" || log.Header.LogLevel == logLevel)
		matchesEnvironment := envErr == nil && (criteria.Environment == "" || log.Header.Environment == env)

		// Adiciona o log à lista apropriada
		if isWithinRetention && matchesLogLevel && matchesEnvironment {
			retainedLogs = append(retainedLogs, log)
		} else {
			discardedLogs = append(discardedLogs, log)
		}
	}

	// Cria o resultado da retenção
	result := RetentionResult{
		RetainedLogs:     retainedLogs,
		DiscardedLogs:    discardedLogs,
		TotalLogs:        len(logs),
		RetainedCount:    len(retainedLogs),
		DiscardedCount:   len(discardedLogs),
		RetentionApplied: true,
	}

	// Retorna um erro se nenhum log foi retido
	if result.RetainedCount == 0 {
		return result, fmt.Errorf("no logs retained with the given criteria")
	}

	return result, nil
}
