package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/marcofranssen/gothermostat/config"
	"github.com/marcofranssen/gothermostat/nest"
	"github.com/marcofranssen/gothermostat/storage"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

const (
	configFile = "./config.json"
)

var (
	store   *storage.Store
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	fmt.Printf("%v, commit %v, built at %v\n\n", version, commit, date)

	myContext, cancel := context.WithCancel(context.Background())

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGKILL)

	go func() {
		cfg := config.New()

		err := cfg.Load(configFile)
		check(err)

		store = storage.NewStore("./data", cfg.Storage.MaxToKeep)

		n := nest.New(cfg)
		err = n.Authenticate()
		check(err)
		cfg.Save(configFile)

		response, err := getData(n)
		check(err)
		now := time.Now()
		fmt.Printf("UserID: %s\nAccessToken: %s\nClientVersion: %v\n", response.Metadata.UserID, response.Metadata.AccessToken, response.Metadata.ClientVersion)
		printThermostatData(now, response.Devices.Thermostats)
		store.SaveTemperatureResult(now, response.Devices.Thermostats)

		go webserver(cfg.Webserver)
		schedule(myContext, n, cfg.Nest.PollInterval*time.Minute)
	}()

	fmt.Println("Waiting for you to close")
	fmt.Println()

	sig := <-sigs
	fmt.Println()
	fmt.Println(sig)
	cancel()

	<-done
	fmt.Println("Exiting")
}

func schedule(ctx context.Context, nest nest.Nest, refreshTime time.Duration) {
	ticker := time.NewTicker(refreshTime)
	go func() {
		for tick := range ticker.C {
			response, err := getData(nest)
			if err != nil {
				return
			}
			printThermostatData(tick, response.Devices.Thermostats)
			store.SaveTemperatureResult(tick, response.Devices.Thermostats)
		}
	}()
	select {
	case <-ctx.Done():
		fmt.Println("Stopping ticker")
		ticker.Stop()
	}
}

func getData(myNest nest.Nest) (*nest.Combined, error) {
	var response nest.Combined
	err := myNest.All(&response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func printThermostatData(tick time.Time, thermostats map[string]*nest.Thermostat) {
	for _, thermostat := range thermostats {
		fmt.Printf("%s | Thermostat: '%s', temperature %v\n", tick.Format(time.RFC3339), thermostat.Name, thermostat.AmbientTemperatureC)
	}
}
