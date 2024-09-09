package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/riquemorozine/todo_list_go/cmd/app"
	"github.com/riquemorozine/todo_list_go/cmd/config"
	"github.com/riquemorozine/todo_list_go/cmd/entities"
	"github.com/riquemorozine/todo_list_go/cmd/infra/dependencies"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {
	r := gin.Default()
	l := log.New(os.Stdout, "[todo-list-go] ", log.LstdFlags)
	c, err := config.LoadConfig(".")

	if err != nil {
		panic(err)
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", "localhost", c.DBUSer, c.DBPass, c.DBName, "5432")
	db, err := gorm.Open(postgres.Open(dsn))
	db.AutoMigrate(&entities.Todo{})

	if err != nil {
		panic(err)
	}

	handlers := dependencies.Start(db)

	app.ConfigureMappings(r, handlers)

	l.Println("Starting server on port 8080")
	r.Run()
}
