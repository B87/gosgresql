package database

import (
	"fmt"
	"os"
	"testing"

	embeddedpostgres "github.com/fergusstrange/embedded-postgres"
)

// TestMain is the entry point for the test suite of the package.
func TestMain(m *testing.M) {
	var code int
	var postgres *embeddedpostgres.EmbeddedPostgres

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from panic: %v\n", r)
			if postgres != nil {
				if err := postgres.Stop(); err != nil {
					fmt.Printf("Failed to stop embedded postgres after panic: %v\n", err)
				}
				fmt.Println("Embedded postgres stopped after panic")
			}
			os.Exit(1)
		}
		os.Exit(code)
	}()

	// Set up
	postgres = embeddedpostgres.NewDatabase(
		embeddedpostgres.DefaultConfig().
			Username("postgres").
			Password("postgres").
			Database("postgres").
			Port(5432).
			Version(embeddedpostgres.V16),
	)
	if err := postgres.Start(); err != nil {
		fmt.Printf("Failed to start embedded postgres: %v", err)
		os.Exit(1)
	}

	// Run tests
	code = m.Run()

	// Tear down
	if err := postgres.Stop(); err != nil {
		fmt.Printf("Failed to stop embedded postgres: %v", err)
		os.Exit(1)
	}
}

func tearUp(t *testing.T) (*DB, func()) {
	t.Helper()
	db, err := NewDefault()
	if err != nil {
		t.Fatalf("Failed to create database: %v", err)
	}
	return db, func() { db.Close() }
}
