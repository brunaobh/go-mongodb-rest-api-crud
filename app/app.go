package main

import (
	"os"
	"encoding/json"
	"log"
	"net/http"

	"gopkg.in/mgo.v2/bson"
	"github.com/gorilla/mux"

	. "github.com/user/app/config"
	. "github.com/user/app/dao"
	. "github.com/user/app/models"
)

var config = Config{}
var dao = FlightsDAO{}

// GET list of flights
func AllFlights(w http.ResponseWriter, r *http.Request) {
	flights, err := dao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	retResponse(w, http.StatusOK, flights)
}

// GET a flight by its ID
func FindFlightEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	flight, err := dao.FindById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Flight ID")
		return
	}
	retResponse(w, http.StatusOK, flight)
}

// POST a new flight
func CreateFlight(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var flight Flight
	if err := json.NewDecoder(r.Body).Decode(&flight); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	flight.ID = bson.NewObjectId()
	if err := dao.Insert(flight); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	retResponse(w, http.StatusCreated, flight)
}

// PUT update an existing flight
func UpdateFlight(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var flight Flight
	if err := json.NewDecoder(r.Body).Decode(&flight); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Update(flight); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	retResponse(w, http.StatusOK, map[string]string{"result": "success"})
}

// DELETE an existing flight
func DeleteFlight(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var flight Flight
	if err := json.NewDecoder(r.Body).Decode(&flight); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Delete(flight); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	retResponse(w, http.StatusOK, map[string]string{"result": "success"})
}

func Health(w http.ResponseWriter, r *http.Request) {
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	retResponse(w, http.StatusOK, map[string]string{"server":name,"result": "success"})
} 

func respondWithError(w http.ResponseWriter, code int, msg string) {
	retResponse(w, code, map[string]string{"error": msg})
}

func retResponse(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

// Define HTTP request routes
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/health", Health)
	r.HandleFunc("/flights", AllFlights).Methods("GET")
	r.HandleFunc("/flights", CreateFlight).Methods("POST")
	r.HandleFunc("/flights", UpdateFlight).Methods("PUT")
	r.HandleFunc("/flights", DeleteFlight).Methods("DELETE")
	r.HandleFunc("/flights/{id}", FindFlightEndpoint).Methods("GET")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
