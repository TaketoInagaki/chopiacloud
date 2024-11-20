package model

import "time"

// Instance represents a model for the instances table
type Instance struct {
	ID      uint32    `json:"id" db:"id"`
	Name    string    `json:"name" db:"name"`
	Status  uint8     `json:"status" db:"status"`
	HostIP  string    `json:"host_ip" db:"host_ip"`
	GuestIP string    `json:"guest_ip" db:"guest_ip"`
	Created time.Time `json:"created" db:"created"`
	Updated time.Time `json:"updated" db:"updated"`
}
