package model

type EditHistoric struct {
	Old        []byte
	New        []byte
	Timestamp  int64
	ModifiedBy []byte
}
