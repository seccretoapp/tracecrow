package logs

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// GetFilePath retorna o caminho do arquivo do segmento
func (s *Segment) GetFilePath() string {
	if s == nil || s.segmentID == "" {
		return ""
	}
	return filepath.Join("/path/to/segments", fmt.Sprintf("segment-%s.log", s.segmentID))
}

// CompressSegment compacta um segmento de log
func CompressSegment(segment *Segment) error {
	// Caminho para o arquivo de segmento
	segmentPath := segment.GetFilePath()
	if segmentPath == "" {
		return fmt.Errorf("invalid segment path")
	}

	// Caminho para o arquivo compactado
	compressedFilePath := segmentPath + ".zip"

	// Verifica se o arquivo compactado já existe
	if _, err := os.Stat(compressedFilePath); err == nil {
		return fmt.Errorf("compressed file already exists: %s", compressedFilePath)
	}

	// Cria o arquivo compactado
	outFile, err := os.Create(compressedFilePath)
	if err != nil {
		return fmt.Errorf("failed to create compressed file: %v", err)
	}
	defer outFile.Close()

	// Cria um escritor ZIP
	zipWriter := zip.NewWriter(outFile)
	defer zipWriter.Close()

	// Adiciona o segmento ao ZIP
	writer, err := zipWriter.Create(filepath.Base(segmentPath))
	if err != nil {
		return fmt.Errorf("failed to create zip entry: %v", err)
	}

	// Abre o arquivo do segmento
	segmentFile, err := os.Open(segmentPath)
	if err != nil {
		return fmt.Errorf("failed to open segment file: %v", err)
	}
	defer segmentFile.Close()

	// Copia os dados do segmento para o ZIP
	_, err = io.Copy(writer, segmentFile)
	if err != nil {
		return fmt.Errorf("failed to write compressed data: %v", err)
	}

	return nil
}
