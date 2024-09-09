package main

import (
	"github.com/gin-gonic/gin"
	"github.com/riquemorozine/todo_list_go/cmd/app"
	"github.com/riquemorozine/todo_list_go/cmd/infra/dependencies"
	"log"
	"os"
)

func main() {
	r := gin.Default()
	l := log.New(os.Stdout, "[todo-list-go] ", log.LstdFlags)

	handlers := dependencies.Start()

	app.ConfigureMappings(r, handlers)
	dependencies.Start()

	l.Println("Starting server on port 8080")
	r.Run()
}
