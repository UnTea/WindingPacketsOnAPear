package main

import (
	"database/sql"
	"log"

	"github.com/UnTea/WindingPacketsOnAPear/api"
	db "github.com/UnTea/WindingPacketsOnAPear/db/sqlc"
	"github.com/UnTea/WindingPacketsOnAPear/util"
	_ "github.com/lib/pq"
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
