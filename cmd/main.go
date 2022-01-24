package main

import (
	"fmt"
	"time"

	"github.com/ChrisALiles/dining"
)

func main() {

	var roomreq dining.RoomReq
	var forkreq dining.ForkReq

	// Channels from room to each philosopher.
	roomPhil0 := make(chan any)
	roomPhil1 := make(chan any)
	roomPhil2 := make(chan any)
	roomPhil3 := make(chan any)
	roomPhil4 := make(chan any)

	// Channels from (any) fork to each philosopher.
	forkPhil0 := make(chan any)
	forkPhil1 := make(chan any)
	forkPhil2 := make(chan any)
	forkPhil3 := make(chan any)
	forkPhil4 := make(chan any)

	// Channels from a philosopher to each fork.
	philFork0 := make(chan dining.ForkReq)
	philFork1 := make(chan dining.ForkReq)
	philFork2 := make(chan dining.ForkReq)
	philFork3 := make(chan dining.ForkReq)
	philFork4 := make(chan dining.ForkReq)

	philRoom := make(chan dining.RoomReq)

	quitreq := make(chan any)

	timeout := time.After(1 * time.Second)

	// Activate the room.
	go dining.Room(philRoom)

	// Activate the philosophers.
	go dining.Philosopher(0,
		philRoom,
		roomPhil0,
		philFork0,
		philFork1,
		forkPhil0,
		quitreq)
	go dining.Philosopher(1,
		philRoom,
		roomPhil1,
		philFork1,
		philFork2,
		forkPhil1,
		quitreq)
	go dining.Philosopher(2,
		philRoom,
		roomPhil2,
		philFork2,
		philFork3,
		forkPhil2,
		quitreq)
	go dining.Philosopher(3,
		philRoom,
		roomPhil3,
		philFork3,
		philFork4,
		forkPhil3,
		quitreq)
	go dining.Philosopher(4,
		philRoom,
		roomPhil4,
		philFork4,
		philFork0,
		forkPhil4,
		quitreq)

	// Activate the forks.
	go dining.Fork(0, philFork0)
	go dining.Fork(1, philFork1)
	go dining.Fork(2, philFork2)
	go dining.Fork(3, philFork3)
	go dining.Fork(4, philFork4)

	// Wait to be timed out.
	<-timeout
	fmt.Println("TIMED OUT")
	// Brig everything to a halt.
	// Ask the philosophers to quit first.
	for i := 1; i < 6; i++ {
		quitreq <- true
	}
	// Now the forks and the room.
	forkreq.Action = dining.Quit
	roomreq.Action = dining.Quit
	philFork0 <- forkreq
	philFork1 <- forkreq
	philFork2 <- forkreq
	philFork3 <- forkreq
	philFork4 <- forkreq

	philRoom <- roomreq
}
