package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type SSLMode string

const SSLRequired SSLMode = "required"
const SSLDisable SSLMode = "disable"

type Config struct {
	Host     string
	Port     int
	DBName   string
	User     string
	Password string
	SSLMode  SSLMode
}

var pool *pgxpool.Pool

func Init(config Config) error {
	dsn := fmt.Sprintf("postgresql://%v:%v@%v:%v/%v?sslmode=%v",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.DBName,
		config.SSLMode,
	)
	var err error
	pool, err = pgxpool.New(context.Background(), dsn)
	return err
}

func GetConnection(ctx context.Context) (*pgxpool.Conn, error) {
	return pool.Acquire(ctx)
}
