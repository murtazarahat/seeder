package main

import (
	"fmt"
	"time"

	"github.com/murtazarahat/seeder"
)

type Student struct {
	Name string `json:"name"`
}

func main() {
	seeder := seeder.NewSeeder(time.Now().UTC().UnixNano(), Student{Name: "Rahat Murtaza"})

	fmt.Println(seeder.GetData().(Student))
}
