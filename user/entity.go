package user

import "time"

type User struct {
	ID             int
	Nama           string
	Pekerjaan      string
	Email          string
	PasswordHash   string
	AvatarFileName string
	Role           string
	Token          string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
