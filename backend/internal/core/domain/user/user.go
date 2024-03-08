package user

import (
	"time"

	"github.com/google/uuid"
)

type Language string

const (
	Eng Language = "english"
	Amh Language = "amharic"
)

type User struct {
	UserID      uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Email       string    `gorm:"unique;not null"`
	Avatar      string
	DisplayName string    `gorm:"unique;not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}

type UserSettings struct {
	UserID   uuid.UUID `gorm:"foreignKey:UserID"`
	Language Language  `gorm:"default:'english'"`
}

func NewUser(g *GithubAuthResp) *User {
	return &User{
		UserID:      uuid.New(),
		Email:       g.Email,
		Avatar:      g.AvatarURL,
		DisplayName: g.Name,
		CreatedAt:   time.Now(),
	}
}

func NewUserSettings(u *User, l Language) *UserSettings {
	return &UserSettings{
		UserID:   u.UserID,
		Language: l,
	}
}
