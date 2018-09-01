package storage

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"time"

	"github.com/marcofranssen/gothermostat/config"
	"github.com/marcofranssen/gothermostat/nest"
)

// Store storage structure
type Store struct {
	thermoDataStorage string
}

type thermoData struct {
	Timestamp           time.Time `json:"timestamp"`
	Thermostat          string    `json:"thermostat"`
	AmbientTemperatureC float64   `json:"ambientTemperatureC"`
	AmbientTemperatureF int       `json:"ambientTemperatureF"`
	TargetTemperatureC  float64   `json:"targetTemperatureC"`
	TargetTemperatureF  int       `json:"targetTemperatureF"`
}

// NewStore create a new store for persisting data
func NewStore(storageLocation string) *Store {
	return &Store{
		storageLocation + "/thermoData.json",
	}
}

// SaveTemperatureResult persists temperature data
func (s *Store) SaveTemperatureResult(tick time.Time, thermostats map[string]*nest.Thermostat) error {
	data, err := s.readFile(s.thermoDataStorage)
	if err != nil {
		return err
	}

	var temperatureData []thermoData
	json.Unmarshal(data, &temperatureData)
	for _, thermostat := range thermostats {
		temperatureData = append(temperatureData, thermoData{
			tick,
			thermostat.Name,
			thermostat.AmbientTemperatureC,
			thermostat.AmbientTemperatureF,
			thermostat.TargetTemperatureC,
			thermostat.TargetTemperatureF,
		})
		bytes, err := config.JsonMarshal(&temperatureData)

		if err != nil {
			return err
		}
		err = s.writeFile(s.thermoDataStorage, bytes)
		if err != nil {
			return err
		}
	}
	return nil
}

// GetTemperatureData reads the temperatureData from storage
func (s *Store) GetTemperatureData() ([]byte, error) {
	return s.readFile(s.thermoDataStorage)
}

func (s *Store) writeFile(filePath string, data []byte) error {
	file, err := os.Create(filePath)
	defer file.Close()
	if err != nil {
		return err
	}
	_, err = file.Write(data)

	return err
}

func (s *Store) readFile(filePath string) ([]byte, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return data, nil
}
