package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/s0m380dy/simplebank/api"
	db "github.com/s0m380dy/simplebank/db/sqlc"
	"github.com/s0m380dy/simplebank/util"
)

func main() {
	config, err := util.LoadConfig(".") // "." потому что app.env в той же директории что и main.go
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	connection, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(connection)
	server, err := api.NewServer(config, store)

	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
