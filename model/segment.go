package model

import (
	"fmt"
	"github.com/google/uuid"
)

type Segment struct {
	Channel  uuid.UUID
	Division uuid.UUID
	Offset   int64
}

// NewSegment cria um novo segmento com canal, divisão e offset.
func NewSegment(channel Channel, division Division, offset int64) Segment {
	return Segment{
		Channel:  channel.ID,
		Division: division.ID,
		Offset:   offset,
	}
}

// Validate verifica se o segmento é válido.
func (s Segment) Validate() error {
	if s.Offset < 0 {
		return fmt.Errorf("o offset não pode ser negativo")
	}
	return nil
}

// Equals compara dois segmentos para verificar igualdade.
func (s Segment) Equals(other Segment) bool {
	return s.Channel == other.Channel &&
		s.Division == other.Division &&
		s.Offset == other.Offset
}
