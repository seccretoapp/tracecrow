package logs

import (
	"fmt"
	"os"
	"testing"
)

func TestSegmentRotation(t *testing.T) {
	tempDir := "./test_logs"
	defer os.RemoveAll(tempDir)

	// Inicializa o logger com limite de 100 bytes por segmento
	logger, err := NewLogger(tempDir, 100, []byte("supersecretkey"))
	if err != nil {
		t.Fatalf("Failed to initialize logger: %v", err)
	}

	// Adiciona logs até o limite de tamanho do segmento
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

	// Verifica o número de segmentos
	segments, err := getAllSegments(tempDir)
	if err != nil {
		t.Fatalf("Failed to get segments: %v", err)
	}

	// Verifica se ao menos dois segmentos foram criados
	if len(segments) < 2 {
		t.Errorf("Expected at least 2 segments, but got %d", len(segments))
	}
}

func getAllSegments(dir string) ([]string, error) {
	var segments []string
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		segments = append(segments, file.Name())
	}
	return segments, nil
}
