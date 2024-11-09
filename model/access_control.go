package model

type AccessControl struct {
	AllowedUsers []Users
	DeniedUsers  []Users
}
