package postgres

import (
	"context"
	"fmt"

	"user-service/internal/postgres/database"
)

const (
	migrationsTable = "goose_db_version"
	selectVersion   = "SELECT version_id FROM " + migrationsTable + " ORDER BY version_id DESC LIMIT 1"
)

func GetMigrationsVersion(db database.DBTX) (int64, error) {
	row := db.QueryRow(context.Background(), selectVersion)
	var actualVersion int64
	err := row.Scan(&actualVersion)
	if err != nil {
		return 0, fmt.Errorf("failed to select migrations version from %s: %s", migrationsTable, err)
	}
	return actualVersion, nil
}
