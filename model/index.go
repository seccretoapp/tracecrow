package model

type Index struct {
	Id            string
	Name          string
	Type          string
	Fields        map[string]interface{}
	IndexMetadata IndexMetadata
}
