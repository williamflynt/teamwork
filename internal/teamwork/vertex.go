package teamwork

import (
	"encoding/json"
	"errors"
	"teamwork/internal/models"
	"time"
)

// Vertex is any entity in our graph representation for TeamWork.
type Vertex interface {
	models.Idable
	models.Typer
	Created() time.Time
	Modified() time.Time
}

// vertex implements Vertex.
type vertex struct {
	Id          string         `json:"id"`
	Type_       string         `json:"type"`
	CreatedUtc  time.Time      `json:"created_utc"`
	ModifiedUtc time.Time      `json:"modified_utc"`
	Attrs       map[string]any `json:"attrs"`
}

// NewGenericVertex returns a pointer to a new `vertex` struct, with `options` set at attrs.
// To make `vertex` useful, you may need to marshal into a known struct type (ex: Person).
func NewGenericVertex(id, type_ string, options ...Option) *vertex {
	attrs := resolveAttrs(options...)
	vtx := vertex{
		Id:          id,
		Type_:       type_,
		CreatedUtc:  time.Now(),
		ModifiedUtc: time.Now(),
		Attrs:       attrs,
	}
	if val, ok := attrs["created_utc"]; ok {
		t, isTime := val.(time.Time)
		if isTime {
			vtx.CreatedUtc = t
		}
	}
	if val, ok := attrs["modified_utc"]; ok {
		t, isTime := val.(time.Time)
		if isTime {
			vtx.CreatedUtc = t
		}
	}
	return &vtx
}

// --- IMPLEMENT VERTEX ---

func (v vertex) GetId() string {
	return v.Id
}

func (v vertex) Type() string {
	return v.Type_
}

func (v vertex) Created() time.Time {
	return v.CreatedUtc
}

func (v vertex) Modified() time.Time {
	return v.ModifiedUtc
}

// --- CUSTOM JSON MARSHAL ---

/*
 * Implement a custom MarshalJSON function to let us pass a generic vertex directly to the UI as needed,
 * without implementing error-prone resolution.
 */

func (v *vertex) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte{}, errors.New("got nil pointer for vertex in MarshalJSON")
	}
	attrs := v.Attrs
	attrs["id"] = v.Id
	attrs["type"] = v.Type_
	return json.Marshal(attrs)
}
