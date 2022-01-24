package dining

import "fmt"

func Room(reqs chan RoomReq) {
	var occupancy uint8

	for {
		req := <-reqs
		if req.Action == Quit {
			fmt.Println("Room exiting")
			return
		}
		if req.Action == exit {
			occupancy--
			fmt.Println("Room occupancy decremented", occupancy)
			continue
		}
		if occupancy == 4 {
			req.ack <- nok
		} else {
			occupancy++
			fmt.Println("Room occupancy incremented", occupancy)
			req.ack <- ok
		}
	}
}
