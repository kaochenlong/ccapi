package main

import (
	"fmt"
	"net/http"

	"cancanapi/internal/comment"
	"cancanapi/internal/database"
	transportHttp "cancanapi/internal/transport/http"
)

type App struct {
}

func (app *App) Run() error {
	fmt.Println("Setup app up!")

	db, err := database.NewDatabase()
	if err != nil {
		return err
	}

	commentService := comment.NewService(db)

	handler := transportHttp.NewHandler(commentService)
	handler.SetupRoutes()

	if err := http.ListenAndServe(":3000", handler.Router); err != nil {
		fmt.Println("Fail to set up server!")
		return err
	}

	return nil
}

func main() {
	fmt.Println("hi")
	app := App{}

	if err := app.Run(); err != nil {
		fmt.Println("Error!")
		fmt.Println(err)
	}

}
