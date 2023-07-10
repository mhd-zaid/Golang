package room

type Room struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type InputRoom struct {
	Name string `json:"name"`
}
