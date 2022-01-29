// dining implements the dining philosophers problem
// described in C A R Hoare's 1978 paper.
package dining

import (
	"fmt"
	"math/rand"
	"time"
)

// Philosopher defines the life of a philosopher, enabled by
// communication with Forks and the Room.
func Philosopher(philId uint8,
	room chan RoomReq,
	roomack chan any,
	fork1 chan ForkReq,
	fork2 chan ForkReq,
	forkack chan any,
	quitreq chan any) {

	var roomreq RoomReq
	roomreq.philId = philId
	roomreq.ack = roomack

	var forkreq ForkReq
	forkreq.philId = philId
	forkreq.Ack = forkack

	for {
		// Housekeeping - check for quit signal.
		select {
		case <-quitreq:
			Log(fmt.Sprintln("Phil", philId, "is exiting"))
			return
		default:
		}
		// The life of a philosopher.
		// First, THINK.
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)

		// Now try to get into the (dining) room.
		// If the room is "full", wait and try again.
		roomreq.Action = entry
		for {
			room <- roomreq
			if roomresp := <-roomack; roomresp == ok {
				Log(fmt.Sprintln("Phil", philId, "has entered the room"))
				break
			}
			Log(fmt.Sprintln("Phil", philId, "waiting for the room"))
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
		// Now pickup the first fork.
		forkreq.Action = pickup
		for {
			fork1 <- forkreq
			if forkresp := <-forkack; forkresp == ok {
				Log(fmt.Sprintln("Phil", philId, "has picked up fork"))
				break
			}
			Log(fmt.Sprintln("Phil", philId, "is waiting for fork"))
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
		// Then the second.
		for {
			fork2 <- forkreq
			if forkresp := <-forkack; forkresp == ok {
				Log(fmt.Sprintln("Phil", philId, "has picked up second fork"))
				break
			}
			Log(fmt.Sprintln("Phil", philId, "is waiting for second fork"))
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
		// Now, eat.
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)

		// Put down the forks, leave the room and go back and think.
		forkreq.Action = putdown
		fork2 <- forkreq
		fork1 <- forkreq
		Log(fmt.Sprintln("Phil", philId, "has put down 2 forks"))
		roomreq.Action = exit
		room <- roomreq
		Log(fmt.Sprintln("Phil", philId, "has left the room"))
	}

}
