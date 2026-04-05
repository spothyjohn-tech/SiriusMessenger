package models

import (
	"time"

	"gorm.io/gorm"
)

// User table
type User struct {
	gorm.Model
	Email      string `gorm:"uniqueIndex;not null"`
	Salt       []byte `gorm: "not null"`
	Username   string `gorm:"uniqueIndex"`
	Verifier   []byte `gorm: "not null"`
	PublicKey  []byte `gorm: "not null"`
	EncPrivKey []byte `gorm: "not null"`
}

// Messages table
type Messages struct {
	ID            uint   `gorm: "primarykey"`
	SenderID      uint   `gorm: "index"`
	ReceiverID    uint   `gorm: "index"`
	GroupID       uint   `gorm: "index"`
	EncryptedData []byte `gorm: "not null"`
	Nonce         []byte `gorm: "not null"`
	IsRead        bool   `gorm:"default:false"`
	CreatedAt     time.Time
}

// Group table
type Group struct {
	gorm.Model
	Name    string
	OwnerID uint
	Members []User `gorm:"many2many:user_groups"`
}

// Channel table
type Channel struct {
	gorm.Model
	Name        string
	Description string
	OwnerID     uint
	Members     []User `gorm:"many2many:channel_members"`
}
