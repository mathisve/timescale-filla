package pkg

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
)

type ConnStruct struct {
	Username string
	Password string
	Host     string
	Port     int
	Database string
}

func (c ConnStruct) CreateConnectionString() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s", c.Username, c.Password, c.Host, c.Port, c.Database)
}

func CreateConnPool(ctx context.Context, connStr string) (*pgxpool.Pool, error) {
	return pgxpool.Connect(ctx, connStr)
}
