package model

import (
	"fmt"
	"github.com/google/uuid"
)

type Channel struct {
	ID   uuid.UUID
	Name string
}

// NewChannel cria um novo canal com o nome fornecido.
func NewChannel(name string) Channel {
	return Channel{
		ID:   uuid.New(),
		Name: name,
	}
}

// Validate verifica se o canal é válido.
func (c Channel) Validate() error {
	if c.Name == "" {
		return fmt.Errorf("o nome do canal não pode ser vazio")
	}
	return nil
}

// Equals compara dois canais para verificar igualdade.
func (c Channel) Equals(other Channel) bool {
	return c.ID == other.ID && c.Name == other.Name
}

// UpdateName atualiza o nome do canal.
func (c *Channel) UpdateName(newName string) error {
	if newName == "" {
		return fmt.Errorf("o novo nome não pode ser vazio")
	}
	c.Name = newName
	return nil
}

// GetID retorna o ID do canal.
func (c Channel) GetID() uuid.UUID {
	return c.ID
}

// GetName retorna o nome do canal.
func (c Channel) GetName() string {
	return c.Name
}

// FilterChannelsByName filtra canais pelo nome.
func FilterChannelsByName(channels []Channel, name string) []Channel {
	var filtered []Channel
	for _, c := range channels {
		if c.Name == name {
			filtered = append(filtered, c)
		}
	}
	return filtered
}
