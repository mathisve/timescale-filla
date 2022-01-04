package pkg

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"time"
)

var InsertStatement = `
   INSERT INTO sensor_data (time, sensor_id, temperature, cpu, randomString) VALUES ($1, $2, $3, $4, $5);
   `

type InsertStruct struct {
	Time         time.Time
	SensorId     int
	Temperature  float64
	Cpu          float64
	RandomString string
}

// BatchInsert data in pool using batch insert
func BatchInsert(ctx context.Context, pool *pgxpool.Pool, data []InsertStruct) (err error) {
	batch := &pgx.Batch{}

	for i := range data {
		batch.Queue(InsertStatement, data[i].Time, data[i].SensorId, data[i].Temperature, data[i].Cpu, data[i].RandomString)
	}

	br := pool.SendBatch(ctx, batch)
	defer br.Close()

	for range data {
		_, err = br.Exec()
		if err != nil {
			return err
		}
	}

	return nil
}

// Insert data in pool
func Insert(ctx context.Context, pool *pgxpool.Pool, data []InsertStruct) (err error) {
	for i := range data {
		_, err = pool.Exec(ctx, InsertStatement, data[i].Time, data[i].SensorId, data[i].Temperature, data[i].Cpu, data[i].RandomString)
		if err != nil {
			return err
		}
	}
	return nil
}
