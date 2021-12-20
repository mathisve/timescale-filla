package pkg

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

var InsertStatement = `
   INSERT INTO sensor_data (time, sensor_id, temperature, cpu) VALUES ($1, $2, $3, $4);
   `

func Insert(ctx context.Context, pool *pgxpool.Pool, data []InsertSchema) (err error) {

	for _, row := range data {
		_, err = pool.Exec(ctx, InsertStatement, row.Time, row.SensorId, row.Temperature, row.Cpu)
		if err != nil {
			return err
		}
	}
	return nil
}
