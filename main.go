package main

import (
	"context"
	"fmt"
	"os"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/leavedtrait/go-ota/cmd/api"
	"github.com/leavedtrait/go-ota/internal/server"
)

func main() {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	router := server.NewRouter(conn)
	srv := api.NewServer(":3000", router)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err.Error())
	}
}