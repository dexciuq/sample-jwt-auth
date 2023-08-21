package app

import (
	"fmt"
	"github.com/dexciuq/sample-jwt-auth/config"
	"github.com/dexciuq/sample-jwt-auth/internal/data"
	"github.com/dexciuq/sample-jwt-auth/internal/db"
	"github.com/dexciuq/sample-jwt-auth/pkg/mailer"
	"github.com/gin-gonic/gin"
	"log"
)

type application struct {
	config config.Config
	models data.Models
	mailer mailer.Mailer
}

func Run(cfg *config.Config) {
	psql := cfg.DB.Postgres
	smtp := cfg.SMTP

	conn, err := db.PostgresConnect(psql.Port, psql.Host, psql.User, psql.Password, psql.DBName)
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
		return
	}

	app := &application{
		config: *cfg,
		models: data.NewModels(conn),
		mailer: mailer.New(smtp.Host, smtp.Port, smtp.Username, smtp.Password, smtp.Sender),
	}

	err = app.serve()
	if err != nil {
		log.Fatalf("failed to server app: %v", err)
		return
	}
}

func (app *application) serve() error {
	engine := gin.Default()

	engine.ForwardedByClientIP = true
	err := engine.SetTrustedProxies(nil)
	if err != nil {
		return err
	}

	// setup routes
	app.routes(&engine.RouterGroup)

	addr := fmt.Sprintf("%s:%d", app.config.Host, app.config.Port)
	err = engine.Run(addr)
	if err != nil {
		return err
	}

	return nil
}
