package logs

import (
	"testing"
	"time"
)

func TestRetainLogs(t *testing.T) {
	// Definindo logs de exemplo
	logs := []LogEntry{
		{ID: "log-001", Timestamp: time.Now().Add(-1 * time.Hour).Unix(), Level: "INFO", Message: "Old log"},
		{ID: "log-002", Timestamp: time.Now().Unix(), Level: "INFO", Message: "New log"},
	}

	retentionPeriod := time.Minute * 30 // Retenção de 30 minutos
	retainedLogs, err := RetainLogs(logs, retentionPeriod)
	if err != nil {
		t.Fatalf("Failed to retain logs: %v", err)
	}

	if len(retainedLogs) != 1 {
		t.Errorf("Expected 1 log to be retained, but got %d", len(retainedLogs))
	}
}
