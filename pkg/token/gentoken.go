package token

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
)

//GenUUIDv4 is a
func GenUUIDv4() uuid.UUID {
	id := uuid.NewV4()
	return id
}

//GenToken is a
func GenToken() string {
	return GenUUIDv4String()
}

//GenUUIDv4String is a
func GenUUIDv4String() string {
	token := GenUUIDv4()
	return fmt.Sprintf("%s", token)
}
