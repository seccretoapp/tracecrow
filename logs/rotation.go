package logs

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// Segment representa um segmento de log
type Segment struct {
	file        *os.File
	writer      *os.File
	segmentID   string
	sizeLimit   int64
	currentSize int64
}

// NewSegment cria um novo segmento
func NewSegment(segmentDir string, sizeLimit int64) (*Segment, error) {
	segmentID := fmt.Sprintf("%d", time.Now().Unix())
	filePath := filepath.Join(segmentDir, fmt.Sprintf("segment-%s.log", segmentID))

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to create segment file: %v", err)
	}

	segment := &Segment{
		file:        file,
		writer:      file,
		segmentID:   segmentID,
		sizeLimit:   sizeLimit,
		currentSize: 0,
	}

	return segment, nil
}

// Append adiciona dados ao segmento, criando um novo se o tamanho máximo for atingido
func (s *Segment) Append(logData []byte) error {
	//if s.currentSize+int64(len(logData)) > s.sizeLimit {
	//	return fmt.Errorf("segment size exceeded")
	//}

	_, err := s.writer.Write(logData)
	if err != nil {
		return fmt.Errorf("failed to write log to segment: %v", err)
	}
	s.currentSize += int64(len(logData))
	return nil
}

// Read lê todos os logs de um segmento
func (s *Segment) Read() ([][]byte, error) {
	var logs [][]byte
	scanner := bufio.NewScanner(s.file)
	for scanner.Scan() {
		logs = append(logs, scanner.Bytes())
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to read from segment: %v", err)
	}
	return logs, nil
}

// Close fecha o segmento de log
func (s *Segment) Close() error {
	return s.file.Close()
}
