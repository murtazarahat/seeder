package seeder

import (
	"testing"
	"time"
)

func TestGenerateName(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	generator := NewGenerator(seed)
	if generator == nil {
		t.Errorf("seeder is null %v", generator)
	}

	name := generator.GenerateName()
	if name == "" {
		t.Error("could not generate name")
	}
}
