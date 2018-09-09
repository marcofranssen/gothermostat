package storage

import (
	"encoding/json"
	"fmt"
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

type Thermostat struct {
	ID           string        `json:"id"`
	Name         string        `json:"name"`
	Temperatures []temperature `json:"temperatures"`
}

type temperature struct {
	Timestamp           time.Time `json:"timestamp"`
	AmbientTemperatureC float64   `json:"ambientTemperatureC"`
	AmbientTemperatureF int       `json:"ambientTemperatureF"`
	TargetTemperatureC  float64   `json:"targetTemperatureC"`
	TargetTemperatureF  int       `json:"targetTemperatureF"`
}

// NewStore create a new store for persisting data
func NewStore(storageLocation string) *Store {
	return &Store{
		storageLocation,
	}
}

// GetTemperatureData reads the temperatureData from storage
func (s *Store) GetTemperatureData() (map[string][]Thermostat, error) {
	fileInfo, err := ioutil.ReadDir(s.thermoDataStorage)
	if err != nil {
		return nil, err
	}

	result := make(map[string][]Thermostat)
	result["thermostats"] = make([]Thermostat, 0, len(fileInfo))

	for _, file := range fileInfo {
		jsonBytes, err := s.readFile(s.thermoDataStorage + "/" + file.Name())
		if err != nil {
			return result, nil
		}
		data, err := s.unmarshal(jsonBytes)
		result["thermostats"] = append(result["thermostats"], *data)
	}

	return result, nil
}

// SaveTemperatureResult persists temperature data
func (s *Store) SaveTemperatureResult(tick time.Time, thermostats map[string]*nest.Thermostat) error {
	for k, v := range thermostats {
		fileName := fmt.Sprintf("%s/%s.json", s.thermoDataStorage, k)
		data, err := s.readFile(fileName)
		if err != nil {
			return err
		}
		thermoData, err := s.unmarshal(data)
		if err != nil {
			return err
		}
		newData := s.updateData(thermoData, v, tick)
		bytes, err := config.JsonMarshal(&newData)

		if err != nil {
			return err
		}
		err = s.writeFile(fileName, bytes)
		return err
	}
	return nil
}

func (s *Store) updateData(storedData *Thermostat, thermostat *nest.Thermostat, tick time.Time) Thermostat {
	storedData.Name = thermostat.Name
	storedData.Temperatures = append(storedData.Temperatures, s.temp(thermostat, tick))

	return *storedData
}

func (s *Store) temp(t *nest.Thermostat, tick time.Time) temperature {
	return temperature{
		tick,
		t.AmbientTemperatureC,
		t.AmbientTemperatureF,
		t.TargetTemperatureC,
		t.TargetTemperatureF,
	}
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

func (s *Store) unmarshal(data []byte) (*Thermostat, error) {
	var thermoData Thermostat
	err := json.Unmarshal(data, &thermoData)
	if err != nil {
		return nil, err
	}
	return &thermoData, nil
}
