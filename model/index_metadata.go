package model

type IndexMetadata struct {
	InvertedIndex map[string]interface{}
	BTree         map[string]interface{}
	ShardInfo     map[string]interface{}
}
