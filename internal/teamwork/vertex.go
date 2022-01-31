package teamwork

import (
	"encoding/json"
	"errors"
	"time"
)

// Vertex is any entity in our graph representation for TeamWork.
type Vertex interface {
	Idable
	Typer
	Created() time.Time
	Modified() time.Time
}

// vertex implements Vertex.
type vertex struct {
	id          string
	type_       string
	createdUtc  time.Time
	modifiedUtc time.Time
	attrs       map[string]any
}

// NewGenericVertex returns a pointer to a new `vertex` struct, with `options` set at attrs.
// To make `vertex` useful, you may need to marshal into a known struct type (ex: Person).
func NewGenericVertex(id, type_ string, options ...Option) *vertex {
	attrs := resolveAttrs(options...)
	vtx := vertex{
		id:          id,
		type_:       type_,
		createdUtc:  time.Now(),
		modifiedUtc: time.Now(),
		attrs:       attrs,
	}
	if val, ok := attrs["created_utc"]; ok {
		t, isTime := val.(time.Time)
		if isTime {
			vtx.createdUtc = t
		}
	}
	if val, ok := attrs["modified_utc"]; ok {
		t, isTime := val.(time.Time)
		if isTime {
			vtx.createdUtc = t
		}
	}
	return &vtx
}

// --- IMPLEMENT VERTEX ---

func (v vertex) GetId() string {
	return v.id
}

func (v vertex) Type() string {
	return v.type_
}

func (v vertex) Created() time.Time {
	return v.createdUtc
}

func (v vertex) Modified() time.Time {
	return v.modifiedUtc
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
	attrs := v.attrs
	attrs["id"] = v.id
	attrs["type"] = v.type_
	return json.Marshal(attrs)
}
