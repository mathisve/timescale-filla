package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
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

	var inserts int
	var insertLock sync.RWMutex

	go func() {
		for {
			time.Sleep(time.Second * 1)
			insertLock.Lock()
			fmt.Printf("Inserts per second: %d\n", inserts)
			inserts = 0
			insertLock.Unlock()
		}
	}()

	for i := 0; i < 100; i++ {
		go func() {
			for {
				ctx := context.Background()
				var amount = 1000
				data := pkg.GenerateData(amount)

				err = pkg.BatchInsert(ctx, pool, data)
				if err != nil {
					log.Println(err)
					return
				}

				go func() {
					insertLock.Lock()
					inserts += amount
					insertLock.Unlock()
				}()
			}
		}()
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()

}
