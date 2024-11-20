package model

import (
	"time"
)

type SSHKey struct {
	ID        uint32    `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	PublicKey string    `json:"public_key" db:"public_key"`
	Created   time.Time `json:"created" db:"created"`
	Updated   time.Time `json:"updated" db:"updated"`
}
