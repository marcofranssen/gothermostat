package main

import (
	"fmt"

	"github.com/marcofranssen/gothermostat/config"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

const (
	configFile = "./config.json"
)

func main() {
	cfg := config.New()

	err := cfg.Load(configFile)
	check(err)

	nest := &nest{config: cfg}
	err = nest.authenticate()
	check(err)
	cfg.Save(configFile)

	var response Combined
	err = nest.All(&response)
	check(err)

	fmt.Printf("UserID: %s\nAccessToken: %s\nClientVersion: %v\n", response.Metadata.UserID, response.Metadata.AccessToken, response.Metadata.ClientVersion)
	for _, thermostat := range response.Devices.Thermostats {
		fmt.Printf("Thermostat: '%s', temperature %v", thermostat.Name, thermostat.AmbientTemperatureC)
	}

}
