package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/marcofranssen/gothermostat/config"
	"github.com/marcofranssen/gothermostat/nest"
	"github.com/marcofranssen/gothermostat/storage"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
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
	logger, err := zap.NewDevelopment(zap.AddStacktrace(zapcore.FatalLevel))
	if err != nil {
		log.Fatalf("Can't initialize zap logger: %v", err)
	}
	defer logger.Sync()
	fmt.Printf("%v, commit %v, built at %v\n\n", version, commit, date)

	cfg := config.New(logger)

	err = cfg.Load(configFile)
	check(err)

	store = storage.NewStore("./data", cfg.Storage.MaxToKeep)

	n := nest.New(cfg)
	err = n.Authenticate()
	check(err)
	cfg.Save(configFile)

	response, err := getData(n)
	check(err)
	now := time.Now()
	logger.Info(
		"Received response",
		zap.String("userid", response.Metadata.UserID),
		zap.String("accesstoken", response.Metadata.AccessToken),
		zap.Int("clientversion", response.Metadata.ClientVersion),
	)
	printThermostatData(now, response.Devices.Thermostats, logger)
	store.SaveTemperatureResult(now, response.Devices.Thermostats)

	myContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	go schedule(myContext, n, cfg.Nest.PollInterval*time.Minute, logger)

	srv, err := NewServer(cfg.Webserver, logger)
	check(err)
	srv.Start()
}

func schedule(ctx context.Context, nest nest.Nest, refreshTime time.Duration, logger *zap.Logger) {
	ticker := time.NewTicker(refreshTime)
	go func() {
		for tick := range ticker.C {
			response, err := getData(nest)
			if err != nil {
				logger.Error("Failed to get data from nest api", zap.Error(err))
				continue
			}
			printThermostatData(tick, response.Devices.Thermostats, logger)
			store.SaveTemperatureResult(tick, response.Devices.Thermostats)
		}
	}()
	select {
	case <-ctx.Done():
		logger.Info("Stopping ticker")
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

func printThermostatData(tick time.Time, thermostats map[string]*nest.Thermostat, logger *zap.Logger) {
	for _, thermostat := range thermostats {
		logger.Info("Received thermostat data", zap.String("thermostat", thermostat.Name), zap.Float64("temperature", thermostat.AmbientTemperatureC))
	}
}
