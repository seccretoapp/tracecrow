package logs

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type Segment struct {
	file        *os.File
	writer      *os.File
	segmentID   string
	sizeLimit   int64
	currentSize int64
}

func NewSegment(segmentDir string, sizeLimit int64) (*Segment, error) {
	// Gera um ID único para o segmento
	segmentID := fmt.Sprintf("%d", time.Now().Unix())
	filePath := filepath.Join(segmentDir, fmt.Sprintf("segment-%s.log", segmentID))

	// Abre ou cria o arquivo do segmento
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to create or open segment file: %v", err)
	}

	// Obtém o tamanho atual do arquivo
	info, err := file.Stat()
	if err != nil {
		file.Close()
		return nil, fmt.Errorf("failed to get segment file info: %v", err)
	}

	segment := &Segment{
		file:        file,
		writer:      file,
		segmentID:   segmentID,
		sizeLimit:   sizeLimit,
		currentSize: info.Size(),
	}

	return segment, nil
}

func (s *Segment) Append(logData []byte) error {
	// Verifica se o segmento atingiu o limite de tamanho
	if s.currentSize+int64(len(logData)) > s.sizeLimit {
		return fmt.Errorf("segment size exceeded")
	}

	// Escreve os dados no arquivo
	_, err := s.writer.Write(logData)
	if err != nil {
		return fmt.Errorf("failed to write log to segment: %v", err)
	}

	// Atualiza o tamanho atual
	s.currentSize += int64(len(logData))
	return nil
}

func (s *Segment) Read() ([][]byte, error) {
	// Reposiciona o ponteiro do arquivo para o início
	_, err := s.file.Seek(0, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to seek to the beginning of the segment: %v", err)
	}

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

func (s *Segment) Close() error {
	if err := s.writer.Close(); err != nil {
		return err
	}
	return s.file.Close()
}

type SegmentManager struct {
	segments    []*Segment
	segmentDir  string
	sizeLimit   int64
	currentFile *Segment
}

func NewSegmentManager(segmentDir string, sizeLimit int64) (*SegmentManager, error) {
	if _, err := os.Stat(segmentDir); os.IsNotExist(err) {
		if err := os.MkdirAll(segmentDir, os.ModePerm); err != nil {
			return nil, fmt.Errorf("failed to create segment directory: %v", err)
		}
	}

	manager := &SegmentManager{
		segmentDir: segmentDir,
		sizeLimit:  sizeLimit,
	}

	// Cria o primeiro segmento
	segment, err := NewSegment(segmentDir, sizeLimit)
	if err != nil {
		return nil, err
	}
	manager.currentFile = segment
	manager.segments = append(manager.segments, segment)

	return manager, nil
}

func (sm *SegmentManager) AddLog(logData []byte) error {
	err := sm.currentFile.Append(logData)
	if err != nil {
		if err.Error() == "segment size exceeded" {
			// Cria um novo segmento
			newSegment, err := NewSegment(sm.segmentDir, sm.sizeLimit)
			if err != nil {
				return fmt.Errorf("failed to create new segment: %v", err)
			}

			sm.currentFile = newSegment
			sm.segments = append(sm.segments, newSegment)
			return sm.currentFile.Append(logData)
		}
		return err
	}
	return nil
}
