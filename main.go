package main

import (
	"fmt"

	"github.com/taki-21/go-todo-app/app/controllers"
	"github.com/taki-21/go-todo-app/app/models"
)

func main() {
	fmt.Println(models.Db)
	controllers.StartMainServer()
}
