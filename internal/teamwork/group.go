package teamwork

import "time"

// Group is a group, organization, or team. It can be associated with 0-N Person.
// It can also be associated with other Group.
type Group struct {
	Id          string    `json:"id"`
	CreatedUtc  time.Time `json:"created_utc"`
	ModifiedUtc time.Time `json:"modified_utc"`
	Name string `json:"name"`
	Description string `json:"description"`
}
