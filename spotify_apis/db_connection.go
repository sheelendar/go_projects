package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

const DATABASE_URL string = "postgres://postgres:password@localhost:5432/spotify_db"

// getDBIsntance create Db connection for application used.
func getDBIsntance() *pgxpool.Pool {
	dbPool, err := pgxpool.NewWithConfig(context.Background(), Config())

	connection, err := dbPool.Acquire(context.Background())
	if err != nil {
		log.Fatal("Error while acquiring connection from the database pool!!")
	}
	defer connection.Release()

	err = connection.Ping(context.Background())
	if err != nil {
		log.Fatal("Could not ping database")
	}

	fmt.Println("Connected to the database!!")
	return dbPool
}

func Config() *pgxpool.Config {
	const defaultMaxConns = int32(4)
	const defaultMinConns = int32(0)

	// Can set other param based on requirement

	dbConfig, err := pgxpool.ParseConfig(DATABASE_URL)
	if err != nil {
		log.Fatal("Failed to create a config, error: ", err)
	}
	dbConfig.MaxConns = defaultMaxConns
	dbConfig.MinConns = defaultMinConns
	return dbConfig
}
