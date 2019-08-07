package storage

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/marcofranssen/gothermostat/nest"
)

// Store storage structure
type Store struct {
	thermoDataStorage string
	maxToKeep         int
}

// Thermostat contains all stored temperatures for a thermostat
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
func NewStore(storageLocation string, maxToKeep int) *Store {
	return &Store{
		storageLocation,
		maxToKeep,
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
		if s.checkTempChanged(thermoData, v) {
			newData := s.updateData(thermoData, v, tick)
			newData = s.removeFromStart(&newData, s.maxToKeep)
			buffer := &bytes.Buffer{}
			encoder := json.NewEncoder(buffer)
			encoder.SetEscapeHTML(false)
			encoder.SetIndent("", "  ")
			err := encoder.Encode(newData)

			if err != nil {
				return err
			}
			err = s.writeFile(fileName, buffer.Bytes())
			return err
		}
	}
	return nil
}

func (s *Store) checkTempChanged(storedData *Thermostat, thermostat *nest.Thermostat) bool {
	lastMeasure := storedData.Temperatures[len(storedData.Temperatures)-1]
	return lastMeasure.AmbientTemperatureC != thermostat.AmbientTemperatureC ||
		lastMeasure.TargetTemperatureC != thermostat.TargetTemperatureC ||
		lastMeasure.AmbientTemperatureF != thermostat.AmbientTemperatureF ||
		lastMeasure.TargetTemperatureF != thermostat.TargetTemperatureF
}

func (s *Store) updateData(storedData *Thermostat, thermostat *nest.Thermostat, tick time.Time) Thermostat {
	storedData.Name = thermostat.Name
	storedData.Temperatures = append(storedData.Temperatures, s.temp(thermostat, tick))

	return *storedData
}

func (s *Store) removeFromStart(storedData *Thermostat, maxToKeep int) Thermostat {
	totalTemps := len(storedData.Temperatures)
	toRemove := totalTemps - maxToKeep

	if toRemove > 0 {
		storedData.Temperatures = storedData.Temperatures[toRemove:]
	}

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
