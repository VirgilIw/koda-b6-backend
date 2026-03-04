package config

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func InitDB() (*pgx.Conn, error) {

	host := os.Getenv("PGHOST")
	port := os.Getenv("PGPORT")
	user := os.Getenv("PGUSER")
	password := os.Getenv("PGPASSWORD")
	dbName := os.Getenv("PGDATABASE")
	sslMode := os.Getenv("PGSSLMODE")

	if host == "" || port == "" || user == "" || password == "" || dbName == "" {
		return nil, fmt.Errorf("database environment variables not set")
	}

	if sslMode == "" {
		sslMode = "disable"
	}

	connString := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=%s",
		user,
		password,
		host,
		port,
		dbName,
		sslMode,
	)

	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
