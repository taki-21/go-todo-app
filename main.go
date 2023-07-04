package main

import (
	"fmt"
	"log"

	"github.com/taki-21/go-todo-app/config"
)

func main() {
	fmt.Println(config.Config.Port)
	fmt.Println(config.Config.SQLDriver)
	fmt.Println(config.Config.DbName)
	fmt.Println(config.Config.LogFile)

	log.Println("test")
}
