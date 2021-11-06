package main

import (
	"context"
	"log"
	"os"
	"strconv"

	"github.com/jackc/pgx/v4"
)

const (
	MigrationsTable = "goose_db_version"
	SelectVersion   = "SELECT version_id FROM " + MigrationsTable + " ORDER BY version_id DESC LIMIT 1"
)

func main() {
	expectedVersion := getExpectedVersion()

	connection, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("failed to connect to postgres:", err)
	}
	err = connection.Ping(context.Background())
	if err != nil {
		log.Fatal("failed to ping postgres:", err)
	}

	row := connection.QueryRow(context.Background(), SelectVersion)
	var actualVersion int64
	err = row.Scan(&actualVersion)
	if err != nil {
		log.Fatalf("failed to select migrations version from %s: %s", MigrationsTable, err)
	}

	if actualVersion < int64(expectedVersion) {
		log.Fatalf("verification failed: actual migrations version %d is earlier than expected %d", actualVersion, expectedVersion)
	}

	log.Printf("verification succeeded: actual migrations version %d is later or equal than expected %d\n", actualVersion, expectedVersion)
}

func getExpectedVersion() int {
	versionString := os.Getenv("MIGRATION_VERSION")
	if versionString == "" {
		log.Fatal("env variable MIGRATION_VERSION is required")
	}
	version, err := strconv.Atoi(versionString)
	if err != nil {
		log.Fatal("MIGRATION_VERSION should be a valid integer:", err)
	}
	return version
}
