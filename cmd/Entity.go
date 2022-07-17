package cmd

import "time"

type User struct {
	ID               int
	Name             string
	Occupation       string
	Email            string
	Password_Hash    string
	Avatar_file_name string
	Role             string
	Created_at       time.Time
	Updated_at       time.Time
}
