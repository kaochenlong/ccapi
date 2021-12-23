package main

import (
	"net/http"

	"cancanapi/internal/comment"
	"cancanapi/internal/database"
	transportHttp "cancanapi/internal/transport/http"

	log "github.com/sirupsen/logrus"
)

type App struct {
	Name    string
	Version string
}

func (app *App) Run() error {
	log.SetFormatter(&log.JSONFormatter{})
	log.WithFields(
		log.Fields{
			"AppName":    app.Name,
			"AppVersion": app.Version,
		},
	).Info("Setting up application")

	db, err := database.NewDatabase()
	if err != nil {
		return err
	}

	err = database.MigrateDB(db)
	if err != nil {
		return err
	}

	commentService := comment.NewService(db)

	handler := transportHttp.NewHandler(commentService)
	handler.SetupRoutes()

	if err := http.ListenAndServe(":3000", handler.Router); err != nil {
		log.Error("Fail to set up server!")
		return err
	}

	return nil
}

func main() {
	app := App{Name: "CanCanAPIService", Version: "1.0.0"}

	if err := app.Run(); err != nil {
		log.Error("Error!")
		log.Fatal(err)
	}
}
