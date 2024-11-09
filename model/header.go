package model

import "github.com/google/uuid"

type LogLevel int
type Environment int

//const (
//	Entry LogLevel = LogLevel("entry")
//	Info
//	Warn
//	Error
//	Fatal
//	Debug
//	Trace
//)

const (
	Production Environment = iota
	Staging
	Development
)

type Header struct {
	ID            uuid.UUID
	Timestamp     int64
	HMAC          string
	LogLevel      LogLevel
	Environment   Environment
	CorrelationID uuid.UUID
	Signature     []byte
	Nonce         []byte
	PublicKey     []byte
}
