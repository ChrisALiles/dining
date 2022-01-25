package dining

import "fmt"

func Fork(forkId uint8, reqs chan ForkReq) {
	state := free
	for {
		req := <-reqs
		if req.Action == Quit {
			fmt.Println("Fork", forkId, "exiting")
			return
		}
		if req.Action == putdown {
			if state == free {
				panic("Fork putdown requested when not picked up")
			}
			state = free
			fmt.Println("Fork", forkId, "put down by phil", req.philId)
			continue
		}
		if state != free {
			fmt.Println("Fork", forkId, "pick up by phil", req.philId, "failed")
			req.ack <- nok
			continue
		}
		state = inuse
		fmt.Println("Fork", forkId, "picked up by phil", req.philId)
		req.ack <- ok
	}

}
