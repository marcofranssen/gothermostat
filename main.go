package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"
	"time"

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
	dataFile   = "./thermodata.json"
)

func main() {
	myContext, cancel := context.WithCancel(context.Background())

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

		response := getData(n)
		now := time.Now()
		fmt.Printf("UserID: %s\nAccessToken: %s\nClientVersion: %v\n", response.Metadata.UserID, response.Metadata.AccessToken, response.Metadata.ClientVersion)
		printThermostatData(now, response.Devices.Thermostats)
		saveTemperatureResult(now, response.Devices.Thermostats)

		go webserver(cfg.Webserver)
		schedule(myContext, n, 5*time.Minute)
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
			response := getData(nest)
			printThermostatData(tick, response.Devices.Thermostats)
			saveTemperatureResult(tick, response.Devices.Thermostats)
		}
	}()
	select {
	case <-ctx.Done():
		fmt.Println("Stopping ticker")
		ticker.Stop()
	}
}

func getData(myNest nest.Nest) nest.Combined {
	var response nest.Combined
	err := myNest.All(&response)
	check(err)
	return response
}

func saveTemperatureResult(tick time.Time, thermostats map[string]*nest.Thermostat) {
	data, _ := readFile(dataFile)
	var temperatureData []thermoData
	json.Unmarshal(data, &temperatureData)
	for _, thermostat := range thermostats {
		temperatureData = append(temperatureData, thermoData{
			tick,
			thermostat.Name,
			thermostat.AmbientTemperatureC,
			thermostat.AmbientTemperatureF,
		})
		bytes, err := json.Marshal(&temperatureData)
		check(err)
		err = writeFile(dataFile, bytes)
		check(err)
	}
}

func writeFile(filePath string, data []byte) error {
	file, err := os.Create(filePath)
	defer file.Close()
	if err != nil {
		return err
	}
	_, err = file.Write(data)

	return err
}

func readFile(filePath string) ([]byte, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return data, nil
}

type thermoData struct {
	Timestamp    time.Time `json:"timestamp"`
	Thermostat   string    `json:"thermostat"`
	TemperatureC float64   `json:"temperatureC"`
	TemperatureF int       `json:"temperatureF"`
}

func printThermostatData(tick time.Time, thermostats map[string]*nest.Thermostat) {
	for _, thermostat := range thermostats {
		fmt.Printf("%s | Thermostat: '%s', temperature %v\n", tick.Format(time.RFC3339), thermostat.Name, thermostat.AmbientTemperatureC)
	}
}
