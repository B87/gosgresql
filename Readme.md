# db-kit

Lib and utils to create apps using postgres as database

* Embedded postgres option with [embedded-postgres](https://github.com/fergusstrange/embedded-postgres/)
* SQLX / SQL database
* Goose migrations
* backup and restore: Requires postgresql-client
* e2e test utils

## Usage

``` go

db, err := NewDefault()
if err != nil {
    panic("Failed to create database: %v", err)
}

// use db and its sql/sqlx conections

db.Close()

```

## Database Configuration

This package provides a PostgreSQL database configuration and connection management for Go applications using `sqlx`.

### Configuration

The database can be configured either through environment variables or programmatically using the `Config` struct.

#### Environment Variables

When using `NewDefault()`, the following environment variables are supported:

| Environment Variable | Default Value       | Description                           |
| -------------------- | ------------------- | ------------------------------------- |
| `POSTGRES_HOST`      | `localhost`         | Database host address                 |
| `POSTGRES_PORT`      | `5432`              | Database port number                  |
| `POSTGRES_USER`      | `postgres`          | Database user                         |
| `POSTGRES_PASSWORD`  | `postgres`          | Database password                     |
| `POSTGRES_DB`        | `postgres`          | Database name                         |
| `MIGRATIONS_DIR`     | `../tmp/migrations` | Directory containing Goose migrations |
| `DATA_DIR`           | `../tmp`            | Directory for database backups        |
