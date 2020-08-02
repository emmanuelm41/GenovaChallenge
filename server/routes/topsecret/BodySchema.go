package topsecret

// Distances asdasd
type Distances struct {
	Kenobi    float64 `json:"kenobi"`
	Sato      float64 `json:"sato"`
	Skywalker float64 `json:"skywalker"`
}

// Messages asdasdasd
type Messages struct {
	Kenobi    []string `json:"kenobi"`
	Sato      []string `json:"sato"`
	Skywalker []string `json:"skywalker"`
}

// Msg asdasd
type Msg struct {
	Distance Distances `json:"distance"`
	Message  Messages  `json:"message"`
}
