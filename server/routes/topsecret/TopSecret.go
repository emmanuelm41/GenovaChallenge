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

// RouteHandler Implementa la logica de negocio para la ruta /topsecret
func RouteHandler(res http.ResponseWriter, req *http.Request) {
	var topSecretMsg Msg

	header := res.Header()
	header.Set("Content-Type", "application/json")

	log.Printf("New request rcv in /topsecret. HTTP Method: %v\n", req.Method)

	if req.Method != "POST" {
		log.Printf("HTTP Method not supported")

		res.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(res, `{"message": "Method is not supported."}`)
		return
	}

	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&topSecretMsg)
	if err != nil {
		log.Printf("Error found trying to read the body")

		res.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(res, `{"message": "Msg rcv is not valid"}`)
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
		log.Printf("Calculation process failed")

		res.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(res, `{"message":"Coud not apply process to rcv data"`)
	} else {
		log.Printf("Sending response back to client")

		res.WriteHeader(http.StatusAccepted)
		fmt.Fprintf(res, `{"message": "%s", "position": { "x": %f, "y": %f, "z": %f }}`, message, x, y, z)
	}

}
