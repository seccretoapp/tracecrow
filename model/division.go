package model

import (
	"fmt"
	"github.com/google/uuid"
)

type Division struct {
	ID   uuid.UUID
	Name string
}

// NewDivision cria uma nova divisão com o nome fornecido.
func NewDivision(name string) Division {
	return Division{
		ID:   uuid.New(),
		Name: name,
	}
}

// Validate verifica se a divisão é válida.
func (d Division) Validate() error {
	if d.Name == "" {
		return fmt.Errorf("o nome da divisão não pode ser vazio")
	}
	return nil
}

// Equals compara duas divisões para verificar igualdade.
func (d Division) Equals(other Division) bool {
	return d.ID == other.ID && d.Name == other.Name
}

// UpdateName atualiza o nome da divisão.
func (d *Division) UpdateName(newName string) error {
	if newName == "" {
		return fmt.Errorf("o novo nome não pode ser vazio")
	}
	d.Name = newName
	return nil
}

// GetID retorna o ID da divisão.
func (d Division) GetID() uuid.UUID {
	return d.ID
}

// GetName retorna o nome da divisão.
func (d Division) GetName() string {
	return d.Name
}
