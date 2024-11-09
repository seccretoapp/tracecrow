package logs

import (
	"archive/zip"
	"fmt"
	"os"
	"path/filepath"
)

// CompressSegment compacta um segmento de log
func CompressSegment(segment *Segment) error {
	compressedFilePath := filepath.Join(segment.segmentID, ".zip")
	outFile, err := os.Create(compressedFilePath)
	if err != nil {
		return fmt.Errorf("failed to create compressed file: %v", err)
	}
	defer outFile.Close()

	zipWriter := zip.NewWriter(outFile)
	defer zipWriter.Close()

	writer, err := zipWriter.Create(segment.segmentID)
	if err != nil {
		return fmt.Errorf("failed to create zip entry: %v", err)
	}

	_, err = writer.Write([]byte(fmt.Sprintf("Compressed log data from segment: %s", segment.segmentID)))
	if err != nil {
		return fmt.Errorf("failed to write compressed data: %v", err)
	}

	return nil
}
