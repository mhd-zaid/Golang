package message

type Message struct {
	ID       int    `json:"id"`
	SenderID int    `json:"sender_id"`
	RoomID   int    `json:"room_id"`
	Content  string `json:"content"`
}

type InputMessage struct {
	SenderID int    `json:"sender_id"`
	RoomID   int    `json:"room_id"`
	Content  string `json:"content"`
}
