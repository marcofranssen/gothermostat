package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/marcofranssen/gothermostat/config"
)

func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func returnServerError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
}

func jsonWrite(w http.ResponseWriter, data interface{}) {
	response, err := json.Marshal(data)
	if err != nil {
		returnServerError(w, err)
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(response)
}

func api(w http.ResponseWriter, r *http.Request) {
	data, err := store.GetTemperatureData()
	if err != nil {
		returnServerError(w, err)
	} else {
		jsonWrite(w, data)
	}
}

func webserver(config config.WebserverConfig) error {
	http.Handle("/", http.FileServer(http.Dir("./web/build")))
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/api/thermostat-data", api)

	fmt.Printf("Starting webserver at %s\n", config.Address)
	return http.ListenAndServe(config.Address, nil)
}
