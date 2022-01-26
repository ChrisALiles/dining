package dining

type RoomReq struct {
	philId uint8
	ack    chan any
	Action uint8
}
type ForkReq struct {
	philId uint8
	ack    chan any
	Action uint8
}

const (
	entry = iota
	exit
	pickup
	putdown
	Quit
	ok
	nok
	free
	inuse

	logfile = "./diningout"
)
