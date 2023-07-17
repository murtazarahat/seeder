package seeder

import (
	"testing"
	"time"
)

func TestGenerateName(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	seeder := NewSeeder(seed)
	if seeder == nil {
		t.Errorf("seeder is null %v", seeder)
	}

	name := seeder.GenerateName()
	if name == "" {
		t.Error("could not generate name")
	}
}
