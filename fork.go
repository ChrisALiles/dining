package dining

import (
	"fmt"
)

func Fork(forkId uint8, reqs chan ForkReq) {
	state := free
	for {
		req := <-reqs
		if req.Action == Quit {
			Log(fmt.Sprintln("Fork", forkId, "exiting"))
			return
		}
		if req.Action == putdown {
			if state == free {
				panic("Fork putdown requested when not picked up")
			}
			state = free
			Log(fmt.Sprintln("Fork", forkId, "put down by phil", req.philId))
			continue
		}
		if state != free {
			Log(fmt.Sprintln("Fork", forkId, "pick up by phil", req.philId, "failed"))
			req.ack <- nok
			continue
		}
		state = inuse
		Log(fmt.Sprintln("Fork", forkId, "picked up by phil", req.philId))
		req.ack <- ok
	}

}
