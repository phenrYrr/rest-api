package main

import (
	"context"
	"fmt"

	"github.com/yggdrasiI1/rest-api/internal/comment"
	"github.com/yggdrasiI1/rest-api/internal/db"
)

func Run() error {
	fmt.Println("starting up our application")

	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("failed to connect to the database")
		return err
	}

	if err := db.MigrateDB(); err != nil {
		fmt.Println("failed to migrate database")
		return err
	}

	cmtService := comment.NewService(db)
	fmt.Println(cmtService.GetComment(
		context.Background(),
		"9a31bf83-28dc-4b8d-bf70-7d347a24ff2e",
	))

	return nil
}

func main() {
	fmt.Println("Go start")

	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
