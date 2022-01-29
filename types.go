package dining

// RoomReq is a request from a philosopher to the room.
type RoomReq struct {
	philId uint8
	ack    chan any
	Action uint8
}

// ForkReq is a request from a philosopher to a fork.
type ForkReq struct {
	philId uint8
	Ack    chan any
	Action uint8
}
