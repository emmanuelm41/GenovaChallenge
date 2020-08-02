package routes

import (
	"GenovaChallenge/models"
	"GenovaChallenge/satelites"
	"GenovaChallenge/workers"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
)

var w1 workers.Worker

var kn KenobiTopSecMsg
var st SatoTopSecMSg
var sw SkywalkerTopSecMsg

// TopSecretSplitRoute asdas
func TopSecretSplitRoute(res http.ResponseWriter, req *http.Request) {

	log.Printf("New request rcv in /topsecret_split. HTTP Method: %v\n", req.Method)

	if req.Method == "POST" {

		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			http.Error(res, "Body rcv is not valid. Please check it first!", http.StatusBadRequest)
			return
		}

		var myMap map[string]map[string]json.RawMessage
		json.Unmarshal(body, &myMap)

		if getType(myMap, "kenobi") {
			err = json.Unmarshal(body, &kn)
		} else if getType(myMap, "sato") {
			err = json.Unmarshal(body, &st)
		} else if getType(myMap, "skywalker") {
			err = json.Unmarshal(body, &sw)
		} else {
			http.Error(res, "Body rcv is not valid. Please check it first!", http.StatusBadRequest)
			return
		}

		if err != nil {
			http.Error(res, "Body rcv is not valid. Please check it first!", http.StatusBadRequest)
			return
		}

		log.Printf("The actual data for Kenobi is %+v\n", kn)
		log.Printf("The actual data for Sato is %+v\n", st)
		log.Printf("The actual data for Skywalker is %+v\n", sw)

		header := res.Header()
		header.Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusAccepted)
		fmt.Fprintf(res, `{"message": "Data received"}`)

	} else if req.Method == "GET" {

		if reflect.DeepEqual(kn, KenobiTopSecMsg{}) || reflect.DeepEqual(st, SatoTopSecMSg{}) || reflect.DeepEqual(sw, SkywalkerTopSecMsg{}) {
			header := res.Header()
			header.Set("Content-Type", "application/json")
			res.WriteHeader(http.StatusAccepted)
			fmt.Fprintf(res, `{"message": Not enough data"`)
			return
		}

		log.Printf("Initiating position and message process with data: %+v %+v %+v\n", kn, st, sw)

		if w1 == (workers.Worker{}) {
			sato := models.Sato{X: satelites.SatoPosX, Y: satelites.SatoPosY, Z: satelites.SatoPosZ}
			kenobi := models.Kenobi{X: satelites.KenobiPosX, Y: satelites.KenobiPosY, Z: satelites.KenobiPosZ}
			skywalker := models.Skywalker{X: satelites.SkywalkerPosX, Y: satelites.SkywalkerPosY, Z: satelites.SkywalkerPosZ}

			w1 = workers.Worker{Kenobi: kenobi, Sato: sato, Skywalker: skywalker}
		}

		x, y, z, err1 := w1.GetLocation(kn.Distance.Kenobi, st.Distance.Sato, sw.Distance.Skywalker)
		message, err2 := w1.GetMessage(kn.Message.Kenobi, st.Message.Sato, sw.Message.Skywalker)

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
