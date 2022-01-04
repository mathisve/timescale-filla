package pkg

type QueryStruct struct {
	SensorId    int
	Temperature float64
	Cpu         float64
}

//
//func Query(ctx context.Context, pool *pgxpool.Pool, data []QueryStruct) (err error) {
//	// not finished
//
//	return nil
//}
