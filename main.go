package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/marcofranssen/gothermostat/config"
	"github.com/marcofranssen/gothermostat/nest"
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
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGKILL)

	go func() {
		cfg := config.New()

		err := cfg.Load(configFile)
		check(err)

		n := nest.New(cfg)
		err = n.Authenticate()
		check(err)
		cfg.Save(configFile)

		var response nest.Combined
		err = n.All(&response)
		check(err)

		fmt.Printf("UserID: %s\nAccessToken: %s\nClientVersion: %v\n", response.Metadata.UserID, response.Metadata.AccessToken, response.Metadata.ClientVersion)
		printThermostatData(response.Devices.Thermostats)
	}()

	fmt.Println("Waiting for you to close")
	fmt.Println()

	sig := <-sigs
	fmt.Println()
	fmt.Println(sig)

	<-done
	fmt.Println("Exiting")
}

func printThermostatData(thermostats map[string]*nest.Thermostat) {
	for _, thermostat := range thermostats {
		fmt.Printf("Thermostat: '%s', temperature %v", thermostat.Name, thermostat.AmbientTemperatureC)
	}
}
