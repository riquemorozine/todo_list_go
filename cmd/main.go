package main

import (
	"github.com/gin-gonic/gin"
	"github.com/riquemorozine/todo_list_go/cmd/app"
	"log"
	"os"
)

func main() {
	r := gin.Default()
	l := log.New(os.Stdout, "todo_list_go", log.LstdFlags)

	router := gin.Default()

	app.ConfigureMappings(router)

	l.Println("Starting server on port 8080")
	r.Run()
}
