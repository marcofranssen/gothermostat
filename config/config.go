package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Config the configuration from configuration.json
type Config struct {
	AuthURL      string `json:"authUrl"`
	TokenURL     string `json:"tokenUrl"`
	ClientID     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
	AuthCode     string `json:"authCode"`
	AccessToken  string `json:"accessToken,omitempty"`
}

// New Create a new instance of the configuration object
func New() *Config {
	return &Config{}
}

func (c *Config) Load(file string) error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, c)
	return err
}

func (c *Config) Save(file string) error {
	fmt.Println("Saving config")
	data, err := jsonMarshal(c)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(file, data, 0644)
	return err
}

// jsonMarshal customized serializer to prevent escapeHtml in urls
func jsonMarshal(t interface{}) ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "  ")
	err := encoder.Encode(t)
	return buffer.Bytes(), err
}
