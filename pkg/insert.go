package pkg

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

var InsertStatement = `
   INSERT INTO sensor_data (time, sensor_id, temperature, cpu) VALUES ($1, $2, $3, $4);
   `

// BatchInsert data in pool using batch insert
func BatchInsert(ctx context.Context, pool *pgxpool.Pool, data []InsertSchema) (err error) {
	batch := &pgx.Batch{}

	for _, row := range data {
		batch.Queue(InsertStatement, row.Time, row.SensorId, row.Temperature, row.Cpu)
	}

	br := pool.SendBatch(ctx, batch)
	defer br.Close()

	for i := 0; i < len(data); i++ {
		_, err = br.Exec()
		if err != nil {
			return err
		}
	}

	return nil
}

// Insert data in pool
func Insert(ctx context.Context, pool *pgxpool.Pool, data []InsertSchema) (err error) {
	for _, row := range data {
		_, err = pool.Exec(ctx, InsertStatement, row.Time, row.SensorId, row.Temperature, row.Cpu)
		if err != nil {
			return err
		}
	}
	return nil
}
