package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"
	"timescale-filla/pkg"

	"github.com/ghodss/yaml"
)

func main() {
	// read password from file
	content, err := os.ReadFile("creds.yaml")
	if err != nil {
		log.Println(err)
		return
	}

	var cs pkg.ConnStruct
	err = yaml.Unmarshal(content, &cs)
	if err != nil {
		log.Println(err)
		return
	}

	// connect to db
	connStr := cs.CreateConnectionString()

	ctx := context.Background()
	pool, err := pkg.CreateConnPool(ctx, connStr)
	if err != nil {
		log.Println(err)
		return
	}

	defer pool.Close()

	// create tables
	err = pkg.CreateTables(ctx, pool)
	if err != nil {
		log.Println(err)
		return
	}

	var inserts int
	var insertLock sync.RWMutex

	// async insert counter
	go func() {
		for {
			time.Sleep(time.Second * 1)
			insertLock.Lock()
			fmt.Printf("Inserts per second: %d\n", inserts)
			inserts = 0
			insertLock.Unlock()
		}
	}()

	// insert goroutine
	for i := 0; i < 1; i++ {
		go func() {
			for {
				ctx := context.Background()
				var amount = 1
				data := pkg.GenerateData(amount)

				err = pkg.BatchInsert(ctx, pool, data)
				if err != nil {
					// if it fails to insert, wait one second
					// since the data is random, there is no need to retry
					log.Println(err)
					time.Sleep(time.Second * 1)

				} else {
					insertLock.Lock()
					inserts += amount
					insertLock.Unlock()
				}
			}
		}()
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()
}
