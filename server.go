package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/marcofranssen/gothermostat/config"
)

func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func api(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile(dataFile)
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		w.Write(data)
	}
}

func webserver(config config.WebserverConfig) error {
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/api/thermostat-data", api)

	fmt.Printf("Starting webserver at %s\n", config.Address)
	return http.ListenAndServe(config.Address, nil)
}
