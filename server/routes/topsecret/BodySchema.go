package topsecret

// Distances Distancia del emisor a cada satelite
type Distances struct {
	Kenobi    float64 `json:"kenobi"`
	Sato      float64 `json:"sato"`
	Skywalker float64 `json:"skywalker"`
}

// Messages Mensajes recibidos del emisor en cada satelite
type Messages struct {
	Kenobi    []string `json:"kenobi"`
	Sato      []string `json:"sato"`
	Skywalker []string `json:"skywalker"`
}

// Msg Distancia del emisor y mensajes recibidos a cada satelite
type Msg struct {
	Distance Distances `json:"distance"`
	Message  Messages  `json:"message"`
}
