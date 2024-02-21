package project

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SourceCode struct {
	gorm.Model
	ID  uuid.UUID `gorm:"primaryKey;default:uuid_generate_v4()"`
	URI string    `gorm:"unique;not null"`
}
