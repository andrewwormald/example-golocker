package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/andrewwormald/golocker"
	"github.com/corverroos/goku/client/logical"
)

func main() {
	ctx := context.Background()
	dbc, err := sql.Open("mysql", "golocker:@tcp(127.0.0.1:3306)/golocker?parseTime=true")
	if err != nil {
		panic(err)
	}

	err = dbc.Ping()
	if err != nil {
		panic(err)
	}

	dbc.SetMaxOpenConns(10)
	dbc.SetMaxIdleConns(5)
	dbc.SetConnMaxLifetime(time.Second * 10)

	cl := logical.New(dbc, dbc)
	client := golocker.New(ctx, dbc, cl)
	go client.SyncForever()

	m := client.NewLocker("isLeader", time.Second * 10)

	// this can be run in a for loop to switch between instances forever
	for {
		fmt.Println("waiting")
		m.Lock()
		fmt.Println("locked")
		time.Sleep(time.Second * 5)
		m.Unlock()
		fmt.Println("unlocked")
	}
}
