package logs

import (
	"fmt"
	"os"
	"testing"
)

func TestCompressSegment(t *testing.T) {
	tempDir := "./test_logs"
	defer os.RemoveAll(tempDir)

	// Inicializa o logger com limite de 100 bytes por segmento
	logger, err := NewLogger(tempDir, 100, []byte("supersecretkey"))
	if err != nil {
		t.Fatalf("Failed to initialize logger: %v", err)
	}

	// Adiciona alguns logs
	for i := 0; i < 5; i++ {
		logEntry := LogEntry{
			ID:        fmt.Sprintf("log-%03d", i),
			Timestamp: 1234567890,
			Level:     "INFO",
			Message:   fmt.Sprintf("Test log entry %d", i),
		}
		err := logger.AddLog(logEntry)
		if err != nil {
			t.Fatalf("Failed to add log: %v", err)
		}
	}

	// Compacta o segmento
	segment, err := NewSegment(tempDir, 100)
	if err != nil {
		t.Fatalf("Failed to create segment: %v", err)
	}
	err = CompressSegment(segment)
	if err != nil {
		t.Fatalf("Failed to compress segment: %v", err)
	}

	// Verifica se o arquivo zip foi gerado
	zipFile := fmt.Sprintf("%s.zip", segment.segmentID)
	if _, err := os.Stat(zipFile); os.IsNotExist(err) {
		t.Errorf("Expected zip file %s to be created, but it doesn't exist", zipFile)
	}
}
