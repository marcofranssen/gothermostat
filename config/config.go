package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// WebserverConfig holds webserver settings
type WebserverConfig struct {
	Address string `json:"address"`
}

// StorageConfig holds storage settings
type StorageConfig struct {
	MaxToKeep int `json:"maxToKeep"`
}

// Config the configuration from configuration.json
type Config struct {
	Webserver    WebserverConfig `json:"webserver"`
	Storage      StorageConfig   `json:"storage"`
	AuthURL      string          `json:"authUrl"`
	TokenURL     string          `json:"tokenUrl"`
	ClientID     string          `json:"clientId"`
	ClientSecret string          `json:"clientSecret"`
	AuthCode     string          `json:"authCode"`
	AccessToken  string          `json:"accessToken,omitempty"`
}

// New Create a new instance of the configuration object
func New() *Config {
	return &Config{}
}

// Load Load the config file in the given filePath
func (c *Config) Load(filePath string) error {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, c)
	return err
}

// Save Save the config in the given filePath
func (c *Config) Save(filePath string) error {
	fmt.Println("Saving config")
	data, err := JSONMarshal(c)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filePath, data, 0644)
	return err
}

// JSONMarshal customized serializer to prevent escapeHtml in urls
func JSONMarshal(t interface{}) ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "  ")
	err := encoder.Encode(t)
	return buffer.Bytes(), err
}
