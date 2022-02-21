package models

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Shortener struct {
	ShortenerID uuid.UUID `json:"id" gorm:"type:uuid;primary_key;not null;"`
	From        string    `json:"from"`
	To          string    `json:"to"`
	TTL         time.Time `json:"ttl"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (shortener *Shortener) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()

	return scope.SetColumn("ShortenerID", uuid)
}
