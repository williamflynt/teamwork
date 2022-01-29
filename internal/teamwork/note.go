package teamwork

import "time"

// Note documents any meaningful/important piece of knowledge. A Note is something we want to remember later.
// A Note is freeform text. It may be related to 0-N other entities - including Person, Group, Interaction, MetaTag.
// A Note may have a Category, like "personal goal" or "focus area".
type Note struct {
	Id          string    `json:"id"`
	CreatedUtc  time.Time `json:"created_utc"`
	ModifiedUtc time.Time `json:"modified_utc"`
	Text string `json:"text"`
	Category string `json:"category"`
}
