package model

import "github.com/google/uuid"

type Index struct {
	Id            uuid.UUID
	Name          string
	Type          string
	Fields        map[string]interface{}
	IndexMetadata IndexMetadata
}
