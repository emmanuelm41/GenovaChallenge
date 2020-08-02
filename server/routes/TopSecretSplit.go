package routes

import (
	"GenovaChallenge/models"
	"GenovaChallenge/satelites"
	"GenovaChallenge/workers"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/imdario/mergo"
)

var topSecretMsg TopSecretMsg
var w1 workers.Worker

// TopSecretSplitRoute asdas
func TopSecretSplitRoute(res http.ResponseWriter, req *http.Request) {

	log.Printf("New request rcv in /topsecret. HTTP Method: %v\n", req.Method)

	if req.Method == "POST" {

		var pTopSecretMsg TopSecretMsg
		err := json.NewDecoder(req.Body).Decode(&pTopSecretMsg)
		if err != nil {
			http.Error(res, "Body rcv is not valid. Please check it first!", http.StatusBadRequest)
			return
		}

		log.Printf("The rcv body is %+v\n", pTopSecretMsg)

		mergo.Merge(&topSecretMsg, pTopSecretMsg)

		log.Printf("The actual data is %+v\n", topSecretMsg)

		header := res.Header()
		header.Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusAccepted)
		fmt.Fprintf(res, `{"message": "Data received"}`)

	} else if req.Method == "GET" {

		log.Printf("Initiating position and message process with data: %+v\n", topSecretMsg)

		if w1 == (workers.Worker{}) {
			sato := models.Sato{X: satelites.SatoPosX, Y: satelites.SatoPosY, Z: satelites.SatoPosZ}
			kenobi := models.Kenobi{X: satelites.KenobiPosX, Y: satelites.KenobiPosY, Z: satelites.KenobiPosZ}
			skywalker := models.Skywalker{X: satelites.SkywalkerPosX, Y: satelites.SkywalkerPosY, Z: satelites.SkywalkerPosZ}

			w1 = workers.Worker{Kenobi: kenobi, Sato: sato, Skywalker: skywalker}
		}

		x, y, z, err1 := w1.GetLocation(topSecretMsg.Distance.Kenobi, topSecretMsg.Distance.Sato, topSecretMsg.Distance.Skywalker)
		message, err2 := w1.GetMessage(topSecretMsg.Message.Kenobi, topSecretMsg.Message.Sato, topSecretMsg.Message.Skywalker)

		if err1 != nil || err2 != nil {
			http.Error(res, "", http.StatusNotFound)
		} else {
			header := res.Header()
			header.Set("Content-Type", "application/json")
			res.WriteHeader(http.StatusAccepted)
			fmt.Fprintf(res, `{"message": "%s", "position": { "x": %f, "y": %f, "z": %f }}`, message, x, y, z)
		}
	} else {
		http.Error(res, "Method is not supported.", http.StatusNotFound)
		return
	}

}
