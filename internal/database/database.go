package database

import (
	"context"
	"fmt"
	"time"

	"github.com/d-darac/lagra/internal/database/sqlc"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Database wraps a pgx connection pool and generated sqlc queries.
type Database struct {
	pool    *pgxpool.Pool
	Queries *sqlc.Queries
}

// Config holds settings used to connect to the PostgreSQL instance.
type Config struct {
	Host, User, Password, DBName, SSLMode string
	Port                                  int
}

// New creates a new Database instance by parsing the given configuration,
// establishing a pgx connection pool, and verifying connectivity with a ping.
// It returns an initialized Database containing the pool and sqlc queries.
func New(ctx context.Context, cfg Config) (Database, error) {
	connStr := fmt.Sprintf(
		"postgresql://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.SSLMode,
	)

	pgxcfg, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return Database{}, fmt.Errorf("Couldn't parse config: %v", err)
	}
	pgxcfg.MaxConnLifetime = 5 * time.Minute

	pool, err := pgxpool.NewWithConfig(ctx, pgxcfg)
	if err != nil {
		return Database{}, fmt.Errorf("Couldn't create a connection pool: %v", err)
	}

	if err := pool.Ping(ctx); err != nil {
		return Database{}, fmt.Errorf("Couldn't ping database: %v", err)
	}

	return Database{pool: pool, Queries: sqlc.New(pool)}, nil
}

// Close terminates the underlying connection pool.
func (db *Database) Close() {
	db.pool.Close()
}

// BeginTx starts a new transaction on the connection pool.
func (db *Database) BeginTx(ctx context.Context) (pgx.Tx, error) {
	return db.pool.Begin(ctx)
}

func (db *Database) Pool() *pgxpool.Pool {
	return db.pool
}
