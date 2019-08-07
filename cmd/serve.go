package cmd

import (
	"context"
	"log"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/marcofranssen/gothermostat/nest"
	"github.com/marcofranssen/gothermostat/server"
	"github.com/marcofranssen/gothermostat/storage"
)

const (
	listenAddr = "listenAddr"
)

var store *storage.Store

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Runs the webserver",
	Run: func(cmd *cobra.Command, args []string) {
		logger, err := zap.NewDevelopment(zap.AddStacktrace(zapcore.FatalLevel))
		if err != nil {
			log.Fatalf("Can't initialize zap logger: %v", err)
		}
		defer logger.Sync()

		cmd.Println()
		cmd.Print(version)

		cfgStorage := viper.Sub("storage")
		store = storage.NewStore("./data", cfgStorage.GetInt("maxToKeep"))

		n := nest.New(viper.Sub("nest.api"))
		err = n.Authenticate()
		check(err)
		err = viper.WriteConfig()
		check(err)
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

		nestCfg := viper.Sub("nest")
		pollInterval := nestCfg.GetDuration("pollInterval")
		go schedule(myContext, n, pollInterval*time.Minute, logger)

		srv, err := server.NewServer(viper.Sub("server"), store, logger)
		check(err)
		srv.Start()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.Flags().String(listenAddr, ":8888", "server listen address")
	viper.BindPFlag("server.listenaddr", serveCmd.Flags().Lookup(listenAddr))
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

func check(err error) {
	if err != nil {
		panic(err)
	}
}
