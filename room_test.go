package dining

import "testing"

func TestRoom(t *testing.T) {
	log := make(chan string)
	go Logger(log, true)
	var req RoomReq
	reqs := make(chan RoomReq)
	ack := make(chan any)
	quitack := make(chan any)
	go Room(reqs, quitack)
	t.Run("One entry", func(t *testing.T) {
		req.Action = entry
		req.ack = ack
		reqs <- req
		if r := <-ack; r != ok {
			t.Errorf("ack is %v", r)
		}
	})
	t.Run("5 entries", func(t *testing.T) {
		for i := 1; i < 4; i++ {
			reqs <- req
			<-ack
		}
		reqs <- req
		if r := <-ack; r != nok {
			t.Errorf("ack is %v", r)
		}
	})
	t.Run("Exit then entry", func(t *testing.T) {
		req.Action = exit
		reqs <- req
		req.Action = entry
		reqs <- req
		if r := <-ack; r != ok {
			t.Errorf("ack is %v", r)
		}
	})
	t.Run("Quit", func(t *testing.T) {
		req.Action = Quit
		reqs <- req
		close(reqs)
		<-quitack // Is this really a test?
	})
}
