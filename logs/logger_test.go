package logs

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestAddLog(t *testing.T) {
	// Definindo o diretório para os logs de teste
	tempDir := "./test_logs"
	defer os.RemoveAll(tempDir)

	// Inicializa o logger com tamanho de segmento de 1024 bytes
	logger, err := NewLogger(tempDir, 1024, []byte("supersecretkey"))
	if err != nil {
		t.Fatalf("Failed to initialize logger: %v", err)
	}

	// Definir um log de exemplo
	logEntry := LogEntry{
		ID:        "log-001",
		Timestamp: time.Now().Unix(),
		Level:     "INFO",
		Message:   "Test log entry",
	}

	// Adiciona o log
	err = logger.AddLog(logEntry)
	if err != nil {
		t.Fatalf("Failed to add log: %v", err)
	}

	// Ler os logs adicionados
	logs, err := logger.ReadLogs()
	if err != nil {
		t.Fatalf("Failed to read logs: %v", err)
	}

	// Verificar se o log foi corretamente adicionado
	if len(logs) == 0 || logs[0].ID != logEntry.ID {
		t.Errorf("Log not added correctly. Expected %v, got %v", logEntry.ID, logs[0].ID)
	}
}

func TestAddMultipleLogs(t *testing.T) {
	tempDir := "./test_logs"
	defer os.RemoveAll(tempDir)

	logger, err := NewLogger(tempDir, 1024, []byte("supersecretkey"))
	if err != nil {
		t.Fatalf("Failed to initialize logger: %v", err)
	}

	// Adiciona múltiplos logs
	for i := 0; i < 10; i++ {
		logEntry := LogEntry{
			ID:        fmt.Sprintf("log-%03d", i),
			Timestamp: time.Now().Unix(),
			Level:     "INFO",
			Message:   fmt.Sprintf("Test log entry %d", i),
		}
		err := logger.AddLog(logEntry)
		if err != nil {
			t.Fatalf("Failed to add log: %v", err)
		}
	}

	// Lê todos os logs adicionados
	logs, err := logger.ReadLogs()
	if err != nil {
		t.Fatalf("Failed to read logs: %v", err)
	}

	// Verifica se o número de logs lidos é correto
	if len(logs) != 10 {
		t.Errorf("Expected 10 logs, but got %d", len(logs))
	}
}
