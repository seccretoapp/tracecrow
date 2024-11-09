package model

import (
	"sync"
)

type Operation string

//const (
//	Insert Operation = Operation("insert")
//	Update
//	Delete
//)

type Entry struct {
	Operation Operation
	Data      map[string]interface{}
	Metadata  Metadata
	Index     Index
	mu        sync.RWMutex
}
