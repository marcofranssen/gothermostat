package main

import (
	"fmt"
	"net/http"

	"github.com/marcofranssen/gothermostat/config"
)

func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func api(w http.ResponseWriter, r *http.Request) {
	data, err := store.GetTemperatureData()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		w.Write(data)
	}
}

func webserver(config config.WebserverConfig) error {
	http.Handle("/", http.FileServer(http.Dir("./web/build")))
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/api/thermostat-data", api)

	fmt.Printf("Starting webserver at %s\n", config.Address)
	return http.ListenAndServe(config.Address, nil)
}
