package model

import (
	"fmt"
	"regexp"
	"sync"
)

type Metrics struct {
	ProcessingTime int64
	DataSize       int64
	Fields         map[string]interface{}
	mu             sync.RWMutex
}

func NewMetrics(processingTime int64, dataSize int64) *Metrics {
	return &Metrics{
		ProcessingTime: processingTime,
		DataSize:       dataSize,
		Fields:         make(map[string]interface{}),
	}
}

func (m *Metrics) AddField(key string, value interface{}) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	validKeyPattern := `^field\d+(\.[a-zA-Z0-9_]+)*$`
	matched, err := regexp.MatchString(validKeyPattern, key)
	if err != nil {
		return fmt.Errorf("erro ao validar a chave: %w", err)
	}
	if !matched {
		return fmt.Errorf("chave inválida: %s", key)
	}

	switch value.(type) {
	case string, int, float64, bool, map[string]interface{}:

	default:
		return fmt.Errorf("tipo de valor inválido para a chave: %s", key)
	}

	if m.Fields == nil {
		m.Fields = make(map[string]interface{})
	}
	m.Fields[key] = value
	return nil
}

func (m *Metrics) GetField(key string) (interface{}, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	value, exists := m.Fields[key]
	return value, exists
}

func (m *Metrics) RemoveField(key string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.Fields != nil {
		delete(m.Fields, key)
	}
}

func (m *Metrics) Merge(other *Metrics) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.Fields == nil {
		m.Fields = make(map[string]interface{})
	}
	for key, value := range other.Fields {
		m.Fields[key] = value
	}
}

func (m *Metrics) ClearFields() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.Fields = make(map[string]interface{})
}

func (m *Metrics) UpdateField(key string, newValue interface{}) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.Fields[key]; !exists {
		return fmt.Errorf("campo não encontrado: %s", key)
	}

	m.Fields[key] = newValue
	return nil
}

func (m *Metrics) MergeWithPriority(other *Metrics, overwrite bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.Fields == nil {
		m.Fields = make(map[string]interface{})
	}
	for key, value := range other.Fields {
		if _, exists := m.Fields[key]; !exists || overwrite {
			m.Fields[key] = value
		}
	}
}

//metrics.IterateFields(func(key string, value interface{}) {
//	fmt.Printf("Campo: %s, Valor: %v\n", key, value)
//})

func (m *Metrics) IterateFields(callback func(key string, value interface{})) {
	for key, value := range m.Fields {
		callback(key, value)
	}
}
