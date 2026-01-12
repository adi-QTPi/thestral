package model

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

const ListenerName = "thestral_channel"

type action string

const (
	EventCreate action = "CREATE"
	EventUpdate action = "UPDATE"
	EventDelete action = "DELETE"
)

type Route struct {
	Host     string         `gorm:"uniqueIndex;not null" json:"host"`
	Targets  pq.StringArray `gorm:"type:text[];not null" json:"targets"`
	IsActive *bool          `gorm:"default:true" json:"is_active"`

	gorm.Model
}

type EventPayload struct {
	Action action `json:"action"`
	ID     uint   `json:"id"`
}
