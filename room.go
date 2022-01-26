package dining

import (
	"fmt"
)

func Room(reqs chan RoomReq, ack chan any) {
	var occupancy uint8

	for {
		req := <-reqs
		if req.Action == Quit {
			// Drain the request channel.
			for range reqs {
			}
			Log(fmt.Sprintln("Room exiting"))
			ack <- true
			return
		}
		if req.Action == exit {
			if occupancy == 0 {
				panic("Room about to have negative occupancy")
			}
			occupancy--
			Log(fmt.Sprintln("Room occupancy decremented to", occupancy))
			continue
		}
		if occupancy == 4 {
			req.ack <- nok
			continue
		}
		occupancy++
		Log(fmt.Sprintln("Room occupancy incremented to", occupancy))
		req.ack <- ok
	}
}
