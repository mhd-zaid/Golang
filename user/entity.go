package user

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	RoomID   int    `json:"room_id"`
}

type InputUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
	RoomID   int    `json:"room_id"`
}
