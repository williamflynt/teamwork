package backends

import (
	"context"
	"github.com/stretchr/testify/assert"
	"teamwork/internal/models"
	"teamwork/internal/teamwork"
	"testing"
)

func TestDaggerDb_CreateVertex(t *testing.T) {
	d, err := NewDagger()
	if err != nil {
		t.Fatal("could not init Dagger backend")
	}
	ctx := context.Background()

	testId := "abc123"
	testType := "generic"
	vtx := genTestVertex(testId, testType)

	f, err := d.CreateVertex(ctx, vtx)
	assert.Nil(t, err)
	assert.Equal(t, testId, f.GetId())
	assert.Equal(t, testType, f.Type())

	v, err := d.GetVertex(ctx, models.NewFetchable(testId, testType))
	assert.Nil(t, err)
	assert.Equal(t, testId, v.GetId())
	assert.Equal(t, testType, v.Type())
	// Due to JSON and floats and time, we have ~1ms difference in exact times.
	// This check is enough.
	assert.Equal(t, vtx.Created().Unix(), v.Created().Unix())
	// assert.Equal(t, vtx.Created().UnixMilli(), v.Created().UnixMilli())
}

// --- HELPERS ---

func genTestVertex(id, type_ string, options ...teamwork.Option) teamwork.Vertex {
	vtx := teamwork.NewGenericVertex(id, type_, options...)
	return *vtx
}
