package internal

import (
	"TSIS1/TSIS1/pkg"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	log.Println("entering health check end point")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "API is up and running")
}

func Players(w http.ResponseWriter, r *http.Request) {
	log.Println("entering persons end point")
	var response pkg.Response
	players := pkg.PrepareResponse()

	response.Players = players

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return
	}

	w.Write(jsonResponse)
}

func StartServer() {
	//create a new router
	router := mux.NewRouter()
	log.Println("creating routes")
	//specify endpoints
	router.HandleFunc("/health-check", HealthCheck).Methods("GET")
	router.HandleFunc("/players", Players).Methods("GET")
	router.HandleFunc("/players/{number}", PlayerInfo).Methods("GET")
	http.Handle("/", router)

	//start and listen to requests
	http.ListenAndServe(":8080", router)

}

func PlayerInfo(w http.ResponseWriter, r *http.Request) {
	log.Println("entering player info end point")
	vars := mux.Vars(r)
	playerNumber, err := strconv.Atoi(vars["number"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid player number")
		return
	}

	pResponse := pkg.PrepareResponse()

	var foundPlayer *pkg.Football
	for _, p := range pResponse {
		if p.Number == playerNumber {
			foundPlayer = &p
			break
		}
	}

	if foundPlayer == nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Player with number %d not found", playerNumber)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonResponse, err := json.Marshal(foundPlayer)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error marshalling player info")
		return
	}

	w.Write(jsonResponse)
}
