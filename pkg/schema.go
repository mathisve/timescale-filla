package pkg

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

var TABLE_SQL = `CREATE TABLE IF NOT EXISTS sensors (
		id SERIAL PRIMARY KEY,
		type VARCHAR(50), 
		location VARCHAR(50)
	);
	CREATE TABLE IF NOT EXISTS sensor_data (
		time TIMESTAMPTZ NOT NULL,
		sensor_id INTEGER,
		temperature DOUBLE PRECISION,
		cpu DOUBLE PRECISION,
		randomString TEXT,
		
		FOREIGN KEY (sensor_id) REFERENCES sensors (id)
	);
	SELECT create_hypertable('sensor_data', 'time', if_not_exists => TRUE);
   `

func CreateTables(ctx context.Context, pool *pgxpool.Pool) (err error) {

	_, err = pool.Exec(ctx, TABLE_SQL)
	if err != nil {
		return err
	}

	return InsertSensorTypes(ctx, pool)
}

func InsertSensorTypes(ctx context.Context, pool *pgxpool.Pool) (err error) {
	for i := range SensorTypes {
		insertQuery := `INSERT INTO sensors (type, location) VALUES ($1, $2);`

		_, err = pool.Exec(ctx, insertQuery, SensorTypes[i], SensorLocations[i])
		if err != nil {
			return err
		}
	}

	return nil
}
