package main

import (
	"context"
	"database/sql"
	"time"

	"github.com/andrewwormald/golocker"
	"github.com/corverroos/goku/client/logical"
	"github.com/luno/jettison/log"
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
	instanceLocker := golocker.New(ctx, "leader_election", dbc, cl)
	go instanceLocker.SyncForever()

	m := instanceLocker.NewLocker("isLeader", time.Second * 10)

	// this can be run in a for loop to switch between instances forever
	log.Info(ctx, "waiting")
	m.Lock()
	log.Info(ctx, "I am now the leader")
	time.Sleep(time.Second * 5)
	m.Unlock()
	log.Info(ctx, "okay bye now")
}
