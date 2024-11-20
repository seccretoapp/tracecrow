package logs

import (
	"encoding/json"
	"fmt"
	"github.com/seccretoapp/tracecrow/model"
	"os"
	"time"

	"github.com/google/uuid"
)

type Log struct {
	Header  model.Header   `json:"header"`
	Segment model.Segment  `json:"segment"`
	Metrics *model.Metrics `json:"metrics,omitempty"`
	Alert   *model.Alert   `json:"alert,omitempty"`
	Entry   *model.Entry   `json:"entry,omitempty"`
	Index   *model.Index   `json:"index,omitempty"`
}

type Logger struct {
	currentSegment *Segment
	segmentDir     string
	segmentSize    int64
	encryptionKey  []byte
}

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

func (l *Logger) AddLog(log Log) error {
	if l.currentSegment == nil {
		return fmt.Errorf("current segment is not initialized")
	}

	// Serializa o log
	logData, err := json.Marshal(log)
	if err != nil {
		return fmt.Errorf("failed to serialize log: %v", err)
	}

	// Adiciona o log ao segmento
	err = l.currentSegment.Append(logData)
	if err != nil {
		return fmt.Errorf("failed to append log to segment: %v", err)
	}

	return nil
}

func (l *Logger) ReadLogs() ([]Log, error) {
	if l.currentSegment == nil {
		return nil, fmt.Errorf("current segment is not initialized")
	}

	var logs []Log
	logsBytes, err := l.currentSegment.Read()
	if err != nil {
		return nil, fmt.Errorf("failed to read logs from segment: %v", err)
	}

	for _, logBytes := range logsBytes {
		var log Log
		err := json.Unmarshal(logBytes, &log)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal log: %v", err)
		}
		logs = append(logs, log)
	}

	return logs, nil
}

func CreateLog(logLevel model.LogLevel, environment model.Environment, correlationID string, segment model.Segment,
	metrics *model.Metrics, alert *model.Alert, entry *model.Entry, index *model.Index) Log {
	return Log{
		Header: model.Header{
			ID:            uuid.New().String(),
			Timestamp:     time.Now().Unix(),
			LogLevel:      logLevel,
			Environment:   environment,
			CorrelationID: correlationID,
			Signature:     nil,
		},
		Segment: segment,
		Metrics: metrics,
		Alert:   alert,
		Entry:   entry,
		Index:   index,
	}
}
