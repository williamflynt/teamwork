package teamwork

import "time"

// Person is a type that represents an individual in the system, including users, contacts, bosses - anyone.
type Person struct {
	Id             string    `json:"id"`
	CreatedUtc     time.Time `json:"created_utc"`
	ModifiedUtc    time.Time `json:"modified_utc"`
	Name           string    `json:"name"`
	BirthdayOn     time.Time `json:"birthday_on"`
	AvatarPath     string    `json:"avatar_path"`
	CanLogin       bool      `json:"can_login"`
	HashedPassword string    `json:"hashed_password"`
}
