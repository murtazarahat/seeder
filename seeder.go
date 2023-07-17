package seeder

import (
	"strings"

	"github.com/goombaio/namegenerator"
	"github.com/iancoleman/strcase"
)

type Seeder struct {
	nameGenerator namegenerator.Generator
}

func (s Seeder) GenerateName() string {
	name := s.nameGenerator.Generate()
	split := strings.Split(name, "-")
	split[0] = strcase.ToCamel(split[0])
	split[1] = strcase.ToCamel(split[1])
	return strings.Join(split, " ")
}

func NewSeeder(seed int64) *Seeder {
	return &Seeder{
		nameGenerator: namegenerator.NewNameGenerator(seed),
	}
}
