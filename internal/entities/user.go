package entities

import "time"

type Gender int8

const (
	Male   = 1
	Female = 2
	Other  = 3
)

type User struct {
	ID        int
	Username  string
	Bio       string
	Gender    Gender
	Email     string
	Followers []Follower
	Following []Follower
	Password  string
}

type Follower struct {
	FollowerID int
	UserID     int
	CreatedAt  time.Time
}
