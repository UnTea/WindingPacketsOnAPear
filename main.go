package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/UnTea/WindingPacketsOnAPear/api"
	db "github.com/UnTea/WindingPacketsOnAPear/db/sqlc"
	"github.com/UnTea/WindingPacketsOnAPear/util"
)

func main() {
	config, err := util.LoadConfig(".")
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to database:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
