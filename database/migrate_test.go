package database

import (
	"os"
	"path/filepath"
	"testing"
)

func TestMigrateUpDown(t *testing.T) {
	// Set up the database
	db, close := tearUp(t)
	defer close()

	// Set the source of the migrations
	db.Migrator.SetSource(filepath.Join(os.TempDir(), "migrations"))

	// Create a new migration
	err := db.Migrator.NewMigration("test1", "sql")
	if err != nil {
		t.Fatalf("Failed to create new migration: %v", err)
	}

	// Migrate up
	err = db.Migrator.Up()
	if err != nil {
		t.Fatalf("Failed to migrate up: %v", err)
	}

	// Get the status of the migrations
	err = db.Migrator.Status()
	if err != nil {
		t.Fatalf("Failed to get migration status: %v", err)
	}

	// Migrate down
	err = db.Migrator.Down()
	if err != nil {
		t.Fatalf("Failed to migrate down: %v", err)
	}

	// Reset the migrations
	err = db.Migrator.Reset()
	if err != nil {
		t.Fatalf("Failed to reset migrations: %v", err)
	}

	// Delete the migrations
	err = os.RemoveAll(db.Migrator.Source())
	if err != nil {
		t.Fatalf("Failed to delete migrations: %v", err)
	}
}
