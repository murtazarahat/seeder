package main

import (
	"fmt"
	"time"

	"github.com/murtazarahat/seeder"
)

func main() {
	generator := seeder.NewGenerator(time.Now().UTC().UnixNano())

	fmt.Println(generator.GenerateName())
}
