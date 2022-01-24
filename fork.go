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
			state = free
			fmt.Println("Fork", forkId, "free")
			continue
		}
		if state != free {
			req.ack <- nok
			continue
		}
		state = inuse
		fmt.Println("Fork", forkId, "in use")
		req.ack <- ok
	}

}
