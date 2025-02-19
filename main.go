package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Println("Новый проект")
	name := flag.String("name", "Anton", "Имя пользователя")

	flag.Parse()

	fmt.Println(*name)
}
