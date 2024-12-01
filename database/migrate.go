package database

import (
	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose/v3"
)

// Migrator is an interface that interacts with the database migrations
type Migrator interface {
	// Apply migrations to the database
	Up() error
	// Rollback migrations to the database
	Down() error
	// Reset the database to the initial state
	Reset() error
	// Get the status of the migrations
	Status() error
	// Create a new migration file
	NewMigration(name, migrationType string) error
	// Get the source of the migrations
	Source() string
	// Set the source of the migrations
	SetSource(source string)
}

// GooseMigrator is a concrete implementation of the Migrator interface
type GooseMigrator struct {
	db            *sqlx.DB
	migrationsDir string
}

func NewGooseMigrator(db *sqlx.DB, migrationsDir string) *GooseMigrator {
	return &GooseMigrator{db: db, migrationsDir: migrationsDir}
}

func (migrator *GooseMigrator) Up() error {
	return goose.Up(migrator.db.DB, migrator.migrationsDir)
}

func (migrator *GooseMigrator) Down() error {
	return goose.Down(migrator.db.DB, migrator.migrationsDir)
}

func (migrator *GooseMigrator) Reset() error {
	return goose.Reset(migrator.db.DB, migrator.migrationsDir)
}

func (migrator *GooseMigrator) Status() error {
	return goose.Status(migrator.db.DB, migrator.migrationsDir)
}

func (migrator *GooseMigrator) NewMigration(name, migrationType string) error {
	return goose.Create(migrator.db.DB, migrator.migrationsDir, name, migrationType)
}

func (migrator *GooseMigrator) Source() string {
	return migrator.migrationsDir
}

func (migrator *GooseMigrator) SetSource(source string) {
	migrator.migrationsDir = source
}
