package user

import "time"

type User struct {
	ID           int
	Name         string
	Occupotion   string
	Email        string
	PasswordHash string
	Avatarfile   string
	Role         string
	Createdat    time.Time
	Updatedat    time.Time
}
