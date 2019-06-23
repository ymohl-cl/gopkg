package model

// Pong model to describe the response http to a ping method
type Pong struct {
	Pong bool `json:"pong"`
}
