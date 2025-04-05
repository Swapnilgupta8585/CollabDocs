// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package database

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Doc struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    uuid.UUID
	Content   string
}

type Link struct {
	Token      string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DocID      uuid.UUID
	Permission string
	ExpiresAt  time.Time
}

type RefreshToken struct {
	RefreshToken string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	UserID       uuid.UUID
	ExpiredAt    time.Time
	RevokedAt    sql.NullTime
}

type User struct {
	ID             uuid.UUID
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Email          string
	HashedPassword string
}
