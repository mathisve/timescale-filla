package pkg

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type ConnStruct struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Database string `json:"database"`
}

func (c ConnStruct) CreateConnectionString() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s", c.Username, c.Password, c.Host, c.Port, c.Database)
}

func CreateConnPool(ctx context.Context, connStr string) (*pgxpool.Pool, error) {
	return pgxpool.Connect(ctx, connStr)
}
