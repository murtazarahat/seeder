package main

import (
	"fmt"
	"time"

	"github.com/murtazarahat/seeder"
)

func main() {
	seeder := seeder.NewSeeder(time.Now().UTC().UnixNano())

	fmt.Println(seeder.GenerateName())
}
