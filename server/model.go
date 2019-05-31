package server

// ModelPong describe the response http to a ping method
type ModelPong struct {
	Pong bool `json:"pong"`
}
