package main

import (
	"log"

	_ "github.com/lib/pq"

	"github.com/radish-miyazaki/go-clean-architecture/api"
	"github.com/radish-miyazaki/go-clean-architecture/config"
	"github.com/radish-miyazaki/go-clean-architecture/infrastructure/db"
	"github.com/radish-miyazaki/go-clean-architecture/infrastructure/server"
)

func main() {
	// DB接続
	conn, err := db.NewDB()
	if err != nil {
		log.Fatalf("Failed to initialize the database: %v", err)
	}
	defer conn.Close()

	// Server立ち上げ
	srv := server.NewServer(
		api.BuildRouter(conn),
	)
	log.Printf("Serving on localhost:%v\n", config.Config.ServerPort)
	log.Fatal(srv.ListenAndServe())
}
