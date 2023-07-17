package seeder

func NewSeeder(seed int64, data any) *Seeder[any] {
	generator := NewGenerator(seed)
	return &Seeder[any]{
		generator: generator,
		data:      data,
	}
}

type Seeder[T any] struct {
	generator *Generator
	data      T
}

func (s Seeder[T]) GetData() T {
	return s.data
}
