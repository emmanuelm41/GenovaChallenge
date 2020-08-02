package server

import (
	"GenovaChallenge/server/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Start asdas
func Start(port string) {

	r := mux.NewRouter()
	r.HandleFunc("/topsecret", routes.TopSecretRoute)
	r.HandleFunc("/topsecret_split", routes.TopSecretSplitRoute)
	http.Handle("/", r)

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
