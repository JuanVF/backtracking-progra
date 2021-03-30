package sockets

// Message will be the structure used as a
// message container to be send to the user
type Message struct {
	ID     int    `json:"ID"`
	Number int    `json:"number"`
	Json   string `json:"json"`
}
