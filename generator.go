package seeder

import (
	"strings"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/goombaio/namegenerator"
	"github.com/iancoleman/strcase"
)

func NewGenerator(seed int64) *Generator {
	return &Generator{
		nameGenerator: namegenerator.NewNameGenerator(seed),
		faker:         *gofakeit.New(seed),
	}
}

type Generator struct {
	nameGenerator namegenerator.Generator
	faker         gofakeit.Faker
}

func (g Generator) GenerateName() string {
	name := g.nameGenerator.Generate()
	split := strings.Split(name, "-")
	split[0] = strcase.ToCamel(split[0])
	split[1] = strcase.ToCamel(split[1])
	return strings.Join(split, " ")
}

func (g Generator) GenerateEmail() string {
	return g.faker.Email()
}

func (g Generator) GeneratePhoneNumber() string {
	return g.faker.Phone()
}
