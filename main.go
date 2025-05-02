package main

import (
	"database/sql"
	"log"

	"github.com/Ma3au88/goSimpleBank/api"
	db "github.com/Ma3au88/goSimpleBank/db/sqlc"
	"github.com/Ma3au88/goSimpleBank/util"
	_ "github.com/lib/pq"
)

const configPath = "."

func main() {
	config, err := util.LoadConfig(configPath)
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	if err = server.Start(config.ServerAddress); err != nil {
		log.Fatal("cannot start server:", err)
	}
}
