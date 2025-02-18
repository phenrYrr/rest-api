package main

import (
	"fmt"

	"github.com/yggdrasiI1/rest-api/internal/comment"
	"github.com/yggdrasiI1/rest-api/internal/db"
	transportHttp "github.com/yggdrasiI1/rest-api/internal/transport/http"
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

	httpHandler := transportHttp.NewHandler(cmtService)

	if err := httpHandler.Serve(); err != nil {
		return err
	}

	return nil
}

func main() {
	fmt.Println("Go start")

	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
