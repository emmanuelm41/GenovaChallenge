package topsecret

import (
	"GenovaChallenge/models"
	"GenovaChallenge/satelites"
	"GenovaChallenge/workers"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var w2 workers.Worker

// RouteHandler asdas
func RouteHandler(res http.ResponseWriter, req *http.Request) {
	var topSecretMsg Msg

	log.Printf("New request rcv in /topsecret. HTTP Method: %v\n", req.Method)

	if req.Method != "POST" {
		http.Error(res, "Method is not supported.", http.StatusNotFound)
		return
	}

	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&topSecretMsg)
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
