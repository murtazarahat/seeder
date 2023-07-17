package seeder

type Seeder struct {
	generator *Generator
}

func NewSeeder(seed int64) *Seeder {
	generator := NewGenerator(seed)
	return &Seeder{
		generator: generator,
	}
}
