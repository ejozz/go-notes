package main

import (
	"log"

	"github.com/ejozz/go-notes/db"
	"github.com/ejozz/go-notes/db/migrations"
	"github.com/ejozz/go-notes/lib/config"
	logger "github.com/ejozz/go-notes/lib/logger"
	"github.com/ejozz/go-notes/note"
	server "github.com/ejozz/go-notes/server/http"
)

func main() {

	config, err := config.Load(".")
	if err != nil {
		log.Fatalf("Could not load config. %v", err)
	}

	l := logger.New(config.LogLevel)

	sqldb, err := db.NewSQL("postgres", config.DBConnString, &l)
	if err != nil {
		l.Fatal().Err(err).Send()
	}

	err = migrations.MigrateDB(config.DBConnString, "file://db/migrations/", &l)
	if err != nil {
		l.Fatal().Err(err).Send()
	}

	s := note.NewService(sqldb)

	r, err := server.NewChiRouter(s, config.PASETOSecret, config.AccessTokenDuration, &l)
	if err != nil {
		l.Fatal().Err(err).Send()
	}

	httpServer, err := server.NewHTTP(r, config.HTTPServerAddress, &l)
	if err != nil {
		l.Fatal().Err(err).Send()
	}

	err = httpServer.Shutdown()
	if err != nil {
		l.Fatal().Err(err).Send()
	}
}
