package logs

import (
	"encoding/json"
	"fmt"
	"github.com/seccretoapp/tracecrow/cryptography"
	"os"
)

// LogEntry define a estrutura básica de um log
type LogEntry struct {
	ID        string `json:"id"`
	Timestamp int64  `json:"timestamp"`
	Level     string `json:"level"`
	Message   string `json:"message"`
	Signature string `json:"signature,omitempty"`
}

// Logger estrutura responsável pela manipulação de logs
type Logger struct {
	currentSegment *Segment
	segmentDir     string
	segmentSize    int64
	encryptionKey  []byte
}

// NewLogger cria um novo logger, configurado com chave de criptografia e diretório de segmentos
func NewLogger(segmentDir string, segmentSize int64, encryptionKey []byte) (*Logger, error) {
	if err := os.MkdirAll(segmentDir, os.ModePerm); err != nil {
		return nil, fmt.Errorf("failed to create log directory: %v", err)
	}

	logger := &Logger{
		segmentDir:    segmentDir,
		segmentSize:   segmentSize,
		encryptionKey: encryptionKey,
	}

	// Cria o primeiro segmento
	segment, err := NewSegment(segmentDir, segmentSize)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize first log segment: %v", err)
	}

	logger.currentSegment = segment
	return logger, nil
}

// AddLog adiciona um novo log no segmento atual, realizando a criptografia se necessário
func (l *Logger) AddLog(logEntry LogEntry) error {
	// Serializa o log
	logData, err := json.Marshal(logEntry)
	if err != nil {
		return fmt.Errorf("failed to serialize log: %v", err)
	}

	// Criptografa o log (se necessário)
	_, _, sig, err := cryptography.Sign(logData)
	if err != nil {
		return fmt.Errorf("failed to encrypt log: %v", err)
	}

	// Adiciona o log criptografado ao segmento
	err = l.currentSegment.Append(sig)
	if err != nil {
		return fmt.Errorf("failed to append log to segment: %v", err)
	}

	return nil
}

// ReadLogs lê todos os logs de um segmento
func (l *Logger) ReadLogs() ([]LogEntry, error) {
	var logs []LogEntry
	logsBytes, err := l.currentSegment.Read()
	if err != nil {
		return nil, fmt.Errorf("failed to read logs from segment: %v", err)
	}

	for _, logBytes := range logsBytes {
		var log LogEntry
		err := json.Unmarshal(logBytes, &log)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal log: %v", err)
		}
		logs = append(logs, log)
	}

	return logs, nil
}
