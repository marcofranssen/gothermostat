package main

import (
	"fmt"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	cfg := &config{}

	err := cfg.load(configFile)
	check(err)

	nest := &nest{config: cfg}
	err = nest.authenticate()
	check(err)
	cfg.save(configFile)

	var response Combined
	err = nest.All(&response)
	check(err)

	fmt.Printf("UserID: %s\nAccessToken: %s\nClientVersion: %v\n", response.Metadata.UserID, response.Metadata.AccessToken, response.Metadata.ClientVersion)
	for _, thermostat := range response.Devices.Thermostats {
		fmt.Printf("Thermostat: '%s', temperature %v", thermostat.Name, thermostat.AmbientTemperatureC)
	}

}
