package dining

// RoomReq describes a request from a philosopher to the room.
type RoomReq struct {
	philId uint8
	ack    chan any
	Action uint8
}

// ForkReq describes a request from a philosopher to a fork.
type ForkReq struct {
	philId uint8
	Ack    chan any
	Action uint8
}
