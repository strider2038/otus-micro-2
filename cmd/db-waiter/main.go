package main

import (
	"context"
	"log"
	"os"
	"strconv"

	"user-service/internal/postgres"

	"github.com/jackc/pgx/v4"
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

	actualVersion, err := postgres.GetMigrationsVersion(connection)
	if err != nil {
		log.Fatal(err)
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
