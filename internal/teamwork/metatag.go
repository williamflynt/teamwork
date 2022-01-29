package teamwork

import "time"

// MetaTag can be anything - a skill, a job title, an email address, a phone number, a project, a task, a vacation...
// The Category of a MetaTag encodes the "thing" that it is. The Value depends on the thing itself - for example, a
// "project" might have the project name as Value. An email address would have the actual email address.
// EarlyUtc and LaterUtc should be interpreted dynamically based on Category, and bound a thing with respect to time.
type MetaTag struct {
	Id          string    `json:"id"`
	CreatedUtc  time.Time `json:"created_utc"`
	ModifiedUtc time.Time `json:"modified_utc"`
	Category    string    `json:"category"`
	Value       string    `json:"value"`
	Description string    `json:"description"`
	EarlyUtc    time.Time `json:"early_utc"`
	LaterUtc    time.Time `json:"later_utc"`
}
