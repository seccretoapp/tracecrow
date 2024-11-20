package main

import (
	"fmt"
	"github.com/seccretoapp/tracecrow/logs"
	"github.com/seccretoapp/tracecrow/model"
	"log"
)

func main() {
	// Configurações do logger
	segmentDir := "./logs"
	segmentSize := int64(1024 * 1024) // 1 MB
	encryptionKey := []byte("encryption-key-placeholder")

	// Cria um logger
	logger, err := logs.NewLogger(segmentDir, segmentSize, encryptionKey)
	if err != nil {
		log.Fatalf("Erro ao criar logger: %v", err)
	}

	// Criação de canal e divisão para o segmento
	channel := model.NewChannel("Main Channel")
	division := model.NewDivision("Primary Division")

	// Validação de canal e divisão
	if err := channel.Validate(); err != nil {
		log.Fatalf("Erro ao validar canal: %v", err)
	}
	if err := division.Validate(); err != nil {
		log.Fatalf("Erro ao validar divisão: %v", err)
	}

	// Criação de um segmento
	segment := model.NewSegment(channel, division, 42)
	if err := segment.Validate(); err != nil {
		log.Fatalf("Erro ao validar segmento: %v", err)
	}

	// Cria um log
	logEntry := logs.CreateLog(
		model.LogLevelInfo,          // Nível de log
		model.EnvironmentProduction, // Ambiente
		"correlation-id-12345",      // Correlation ID
		segment,                     // Segmento
		&model.Metrics{
			ProcessingTime: 123,
			DataSize:       456,
		},
		&model.Alert{
			IsCritical:   true,
			AlertMessage: "Test Alert",
		},
		&model.Entry{
			Operation: "INSERT",
			Data: map[string]interface{}{
				"field1": "value1",
				"field2": "value2",
			},
		},
		&model.Index{
			Id:   "index-001", // Corrigido: Campo `Id`
			Name: "test-index",
			Type: "test-type",
			Fields: map[string]interface{}{ // Corrigido: Usar `map[string]interface{}`
				"field1": "hash1",
				"field2": "hash2",
			},
		},
	)

	// Adiciona o log ao logger
	err = logger.AddLog(logEntry)
	if err != nil {
		log.Fatalf("Erro ao adicionar log: %v", err)
	}
	fmt.Println("Log adicionado com sucesso.")

	// Lê os logs do segmento atual
	readLogs, err := logger.ReadLogs()
	if err != nil {
		log.Fatalf("Erro ao ler logs: %v", err)
	}

	// Exibe os logs lidos
	fmt.Println("Logs lidos do segmento atual:")
	for _, l := range readLogs {
		fmt.Printf("Log ID: %s, Timestamp: %d, LogLevel: %s, Channel: %s, Division: %s, Offset: %d\n",
			l.Header.ID, l.Header.Timestamp, l.Header.LogLevel, l.Segment.Channel, l.Segment.Division, l.Segment.Offset)
	}
}
