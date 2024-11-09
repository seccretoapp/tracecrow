package model

import (
	"encoding/binary"
	"errors"
	"math"
)

type PrimitiveType string

const (
	TypeInt    PrimitiveType = "int"
	TypeFloat  PrimitiveType = "float"
	TypeString PrimitiveType = "string"
	TypeBool   PrimitiveType = "bool"
)

type StoredValue struct {
	Type  PrimitiveType
	Value []byte
}

func ToBytes(value interface{}) (StoredValue, error) {
	switch v := value.(type) {
	case int:
		b := make([]byte, 8)
		binary.BigEndian.PutUint64(b, uint64(v))
		return StoredValue{Type: TypeInt, Value: b}, nil
	case float64:
		b := make([]byte, 8)
		binary.BigEndian.PutUint64(b, math.Float64bits(v))
		return StoredValue{Type: TypeFloat, Value: b}, nil
	case string:
		return StoredValue{Type: TypeString, Value: []byte(v)}, nil
	case bool:
		b := []byte{0}
		if v {
			b[0] = 1
		}
		return StoredValue{Type: TypeBool, Value: b}, nil
	default:
		return StoredValue{}, errors.New("tipo não suportado")
	}
}

func (sv StoredValue) FromBytes() (interface{}, error) {
	switch sv.Type {
	case TypeInt:
		if len(sv.Value) != 8 {
			return nil, errors.New("tamanho inválido para int")
		}
		return int(binary.BigEndian.Uint64(sv.Value)), nil
	case TypeFloat:
		if len(sv.Value) != 8 {
			return nil, errors.New("tamanho inválido para float")
		}
		return math.Float64frombits(binary.BigEndian.Uint64(sv.Value)), nil
	case TypeString:
		return string(sv.Value), nil
	case TypeBool:
		if len(sv.Value) != 1 {
			return nil, errors.New("tamanho inválido para bool")
		}
		return sv.Value[0] == 1, nil
	default:
		return nil, errors.New("tipo não suportado")
	}
}
