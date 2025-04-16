package db

import (
	"context"
	"os"
	"github.com/jackc/pgx/v5"
)
const schema = " CREATE TABLE IF NOT EXISTS users (id BIGSERIAL PRIMARY KEY,  email TEXT UNIQUE NOT NULL, name TEXT NOT NULL, password TEXT NOT NULL);"
func ConnectToDB() (*pgx.Conn ,error){
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	conn.Exec(context.Background(), schema)
	return conn,err
}