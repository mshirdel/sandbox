package database

import (
	"context"
	"fmt"

	"github.com/mshirdel/sandbox/database/repo"
)

func Run() {
	database := NewSandboxDatabase()
	if err := database.Init(); err != nil {
		panic(err)
	}

	fmt.Println("connected to database")
	defer database.Close()

	repo := repo.NewUserRepo(database.DB)
	if err := repo.CreateUser(context.Background(), "ali"); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("ok")
}
