package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/luismiguel010/match/middlew"
	"github.com/luismiguel010/match/routers"
	"github.com/rs/cors"
)

func Handler() {
	router := mux.NewRouter()

	router.HandleFunc("/match", middlew.CheckFiles(routers.MakeMatch)).Methods(http.MethodPost)
	router.HandleFunc("/generate", routers.GeneratorFiles).Methods(http.MethodPost)

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
