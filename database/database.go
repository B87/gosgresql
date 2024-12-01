package database

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DB struct {
	// Apply migrations to the database
	Migrator Migrator

	db     *sqlx.DB
	config Config
}

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	// goose migrations path
	MigrationsDir string
	// backup data path
	BackupsDir string
}

func (c Config) ConnectionString() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		c.Host, c.Port, c.User, c.Password, c.DBName)
}

func New(config Config) (*DB, error) {
	sqlxConn, err := sqlx.Connect("postgres", config.ConnectionString())
	if err != nil {
		return nil, err
	}
	return &DB{
		db:       sqlxConn,
		config:   config,
		Migrator: NewGooseMigrator(sqlxConn, config.MigrationsDir),
	}, nil
}

func NewDefault() (*DB, error) {
	port, err := strconv.Atoi(envOrDefault("POSTGRES_PORT", "5432"))
	if err != nil {
		return nil, err
	}
	config := Config{
		Host:          envOrDefault("POSTGRES_HOST", "localhost"),
		Port:          port,
		User:          envOrDefault("POSTGRES_USER", "postgres"),
		Password:      envOrDefault("POSTGRES_PASSWORD", "postgres"),
		DBName:        envOrDefault("POSTGRES_DB", "postgres"),
		MigrationsDir: envOrDefault("MIGRATIONS_DIR", "../tmp/migrations"),
		BackupsDir:    envOrDefault("DATA_DIR", "../tmp"),
	}
	return New(config)
}

// Close should be called when the application is shutting down.
func (d *DB) Close() error {
	return d.db.Close()
}

func envOrDefault(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}
