package server

import (
	"GenovaChallenge/server/routes/topsecret"
	"GenovaChallenge/server/routes/topsecretsplit"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Start asdas
func Start(port string) {

	r := mux.NewRouter()
	r.HandleFunc("/topsecret", topsecret.RouteHandler)
	r.HandleFunc("/topsecret_split", topsecretsplit.RouteHandler)
	http.Handle("/", r)

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
