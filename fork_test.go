package dining

import "testing"

func TestFork(t *testing.T) {
	log := make(chan string)
	go Logger(log)
	var req ForkReq
	reqs := make(chan ForkReq)
	ack := make(chan any)
	go Fork(4, reqs)
	t.Run("Pickup", func(t *testing.T) {
		req.Action = pickup
		req.ack = ack
		reqs <- req
		if r := <-ack; r != ok {
			t.Errorf("ack is %v", r)
		}
	})
	t.Run("Pickup again", func(t *testing.T) {
		reqs <- req
		if r := <-ack; r != nok {
			t.Errorf("ack is %v", r)
		}
	})
	t.Run("Putdown then pickup", func(t *testing.T) {
		req.Action = putdown
		reqs <- req
		req.Action = pickup
		reqs <- req
		if r := <-ack; r != ok {
			t.Errorf("ack is %v", r)
		}
	})
	t.Run("Quit", func(t *testing.T) {
		req.Action = Quit
		reqs <- req
		select {
		case reqs <- req:
			t.Errorf("Fork did not quit")
		default:
		}

	})
}
