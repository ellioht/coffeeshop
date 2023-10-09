package coffee

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	pool *pgxpool.Pool
}

func NewCoffeeDB(pool *pgxpool.Pool) *DB {
	return &DB{
		pool: pool,
	}
}
