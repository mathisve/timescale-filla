package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"timescale-filla/pkg"
)

func main() {
	connStr := pkg.ConnStruct{
		Username: "postgres",
		Password: "password",
		Host:     "localhost",
		Port:     5432,
		Database: "postgres",
	}.CreateConnectionString()

	ctx := context.Background()
	pool, err := pkg.CreateConnPool(ctx, connStr)
	if err != nil {
		log.Println(err)
		return
	}

	defer pool.Close()

	err = pkg.CreateTables(ctx, pool)
	if err != nil {
		log.Println(err)
		return
	}

	for i := 0; i <= 100; i++ {
		go func() {
			for {
				ctx := context.Background()
				var amount = 100
				data := pkg.GenerateData(amount)

				err = pkg.Insert(ctx, pool, data)
				if err != nil {
					log.Println(err)
					return
				}

				fmt.Printf("Inserted %d rows of data!\n", amount)
			}
		}()
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()

}
