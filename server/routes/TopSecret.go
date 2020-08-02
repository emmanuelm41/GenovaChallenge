package routes

import (
	"GenovaChallenge/models"
	"GenovaChallenge/satelites"
	"GenovaChallenge/workers"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

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

// TopSecretMsg asdasd
type TopSecretMsg struct {
	Distance Distances `json:"distance"`
	Message  Messages  `json:"message"`
}

var w2 workers.Worker

// TopSecretRoute asdas
func TopSecretRoute(res http.ResponseWriter, req *http.Request) {
	var topSecretMsg TopSecretMsg

	log.Printf("New request rcv in /topsecret. HTTP Method: %v\n", req.Method)

	if req.Method != "POST" {
		http.Error(res, "Method is not supported.", http.StatusNotFound)
		return
	}

	err := json.NewDecoder(req.Body).Decode(&topSecretMsg)
	if err != nil {
		http.Error(res, "Body rcv is not valid. Please check it first!", http.StatusBadRequest)
		return
	}

	log.Printf("The rcv body is %+v\n", topSecretMsg)

	if w2 == (workers.Worker{}) {
		sato := models.Sato{X: satelites.SatoPosX, Y: satelites.SatoPosY, Z: satelites.SatoPosZ}
		kenobi := models.Kenobi{X: satelites.KenobiPosX, Y: satelites.KenobiPosY, Z: satelites.KenobiPosZ}
		skywalker := models.Skywalker{X: satelites.SkywalkerPosX, Y: satelites.SkywalkerPosY, Z: satelites.SkywalkerPosZ}

		w2 = workers.Worker{Kenobi: kenobi, Sato: sato, Skywalker: skywalker}
	}

	x, y, z, err1 := w2.GetLocation(topSecretMsg.Distance.Kenobi, topSecretMsg.Distance.Sato, topSecretMsg.Distance.Skywalker)
	message, err2 := w2.GetMessage(topSecretMsg.Message.Kenobi, topSecretMsg.Message.Sato, topSecretMsg.Message.Skywalker)

	if err1 != nil || err2 != nil {
		http.Error(res, "", http.StatusNotFound)
	} else {
		header := res.Header()
		header.Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusAccepted)
		fmt.Fprintf(res, `{"message": "%s", "position": { "x": %f, "y": %f, "z": %f }}`, message, x, y, z)
	}

}
