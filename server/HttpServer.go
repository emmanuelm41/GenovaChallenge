package server

import (
	"GenovaChallenge/server/routes/topsecret"
	"GenovaChallenge/server/routes/topsecretsplit"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Start Iniciamos el servidor HTTP, escuchando en un puerto en particular
func Start(port string) {

	// Registramos las rutas que necesitamos
	r := mux.NewRouter()
	r.HandleFunc("/topsecret", topsecret.RouteHandler)
	r.HandleFunc("/topsecret_split", topsecretsplit.RouteHandler)
	http.Handle("/", r)

	// Iniciamos el servidor
	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
