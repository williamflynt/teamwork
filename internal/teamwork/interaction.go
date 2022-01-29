package teamwork

import "time"

// Interaction is any interaction between 2-N Person and/or Group. It can have a Channel, like "chat".
type Interaction struct {
	Id              string    `json:"id"`
	CreatedUtc      time.Time `json:"created_utc"`
	ModifiedUtc     time.Time `json:"modified_utc"`
	Title           string    `json:"title"`
	Summary         string    `json:"summary"`
	Channel         string    `json:"channel"`
	OccurredUtc     string    `json:"occurred_utc"`
	DurationSeconds int       `json:"duration_seconds"`
}
