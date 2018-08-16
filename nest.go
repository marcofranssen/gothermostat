package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/marcofranssen/gothermostat/config"
)

type apiError struct {
	Error       string `json:"error,omitempty"`
	Description string `json:"error_description,omitempty"`
	InstanceID  string `json:"instance_id,omitempty"`
}

type tokenResponse struct {
	AccessToken string `json:"access_token,omitempty"`
	ExpiresIn   int    `json:"expires_in,omitempty"`
}

type Devices struct {
	Thermostats map[string]*Thermostat `json:"thermostats,omitempty"`
}

type Thermostat struct {
	Humidity                  int       `json:"humidity"`
	Locale                    string    `json:"locale"`
	TemperatureScale          string    `json:"temperature_scale"`
	IsUsingEmergencyHeat      bool      `json:"is_using_emergency_heat"`
	HasFan                    bool      `json:"has_fan"`
	SoftwareVersion           string    `json:"software_version"`
	HasLeaf                   bool      `json:"has_leaf"`
	WhereID                   string    `json:"where_id"`
	DeviceID                  string    `json:"device_id"`
	Name                      string    `json:"name"`
	CanHeat                   bool      `json:"can_heat"`
	CanCool                   bool      `json:"can_cool"`
	TargetTemperatureC        float64   `json:"target_temperature_c"`
	TargetTemperatureF        int       `json:"target_temperature_f"`
	TargetTemperatureHighC    float64   `json:"target_temperature_high_c"`
	TargetTemperatureHighF    int       `json:"target_temperature_high_f"`
	TargetTemperatureLowC     float64   `json:"target_temperature_low_c"`
	TargetTemperatureLowF     int       `json:"target_temperature_low_f"`
	AmbientTemperatureC       float64   `json:"ambient_temperature_c"`
	AmbientTemperatureF       int       `json:"ambient_temperature_f"`
	AwayTemperatureHighC      float64   `json:"away_temperature_high_c"`
	AwayTemperatureHighF      int       `json:"away_temperature_high_f"`
	AwayTemperatureLowC       float64   `json:"away_temperature_low_c"`
	AwayTemperatureLowF       int       `json:"away_temperature_low_f"`
	EcoTemperatureHighC       float64   `json:"eco_temperature_high_c"`
	EcoTemperatureHighF       int       `json:"eco_temperature_high_f"`
	EcoTemperatureLowC        float64   `json:"eco_temperature_low_c"`
	EcoTemperatureLowF        int       `json:"eco_temperature_low_f"`
	IsLocked                  bool      `json:"is_locked"`
	LockedTempMinC            float64   `json:"locked_temp_min_c"`
	LockedTempMinF            int       `json:"locked_temp_min_f"`
	LockedTempMaxC            float64   `json:"locked_temp_max_c"`
	LockedTempMaxF            int       `json:"locked_temp_max_f"`
	SunlightCorrectionActive  bool      `json:"sunlight_correction_active"`
	SunlightCorrectionEnabled bool      `json:"sunlight_correction_enabled"`
	StructureID               string    `json:"structure_id"`
	FanTimerActive            bool      `json:"fan_timer_active"`
	FanTimerTimeout           time.Time `json:"fan_timer_timeout"`
	FanTimerDuration          int       `json:"fan_timer_duration"`
	PreviousHvacMode          string    `json:"previous_hvac_mode"`
	HvacMode                  string    `json:"hvac_mode"`
	TimeToTarget              string    `json:"time_to_target"`
	TimeToTargetTraining      string    `json:"time_to_target_training"`
	WhereName                 string    `json:"where_name"`
	Label                     string    `json:"label"`
	NameLong                  string    `json:"name_long"`
	IsOnline                  bool      `json:"is_online"`
	LastConnection            time.Time `json:"last_connection"`
	HvacState                 string    `json:"hvac_state"`
}

type Where struct {
	WhereID string `json:"where_id"`
	Name    string `json:"name"`
}
type Structure struct {
	Name        string            `json:"name"`
	CountryCode string            `json:"country_code"`
	TimeZone    string            `json:"time_zone"`
	Away        string            `json:"away"`
	Thermostats []string          `json:"thermostats"`
	StructureID string            `json:"structure_id"`
	Wheres      map[string]*Where `json:"wheres"`
}

type Metadata struct {
	AccessToken   string `json:"access_token"`
	ClientVersion int    `json:"client_version"`
	UserID        string `json:"user_id"`
}

type Combined struct {
	Devices    Devices               `json:"devices"`
	Structures map[string]*Structure `json:"structures"`
	Metadata   Metadata              `json:"metadata"`
}

type nest struct {
	config *config.Config
}

func (n *nest) authenticate() error {
	if len(n.config.AuthCode) <= 0 {
		fmt.Printf("Go to %s and get a authCode and put it in your config file.\n", n.config.AuthURL)
	}

	if len(n.config.AccessToken) <= 0 {
		tokenResp, err := getAccessToken(n.config)
		if err != nil {
			return err
		}

		n.config.AccessToken = tokenResp.AccessToken

		fmt.Println(tokenResp)
	}
	return nil
}

func (n *nest) Devices(devices *Devices) error {
	return n.get("/devices", devices)
}

func (n *nest) All(devices *Combined) error {
	return n.get("/", devices)
}

func (n *nest) get(path string, response interface{}) error {
	client := http.Client{
		CheckRedirect: checkRedirect,
	}

	url := fmt.Sprintf("https://developer-api.nest.com%s", path)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("User-Agent", "")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", n.config.AccessToken))

	var resp *http.Response
	resp, err = client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("Api request failed: %s\n%s", resp.Status, body)
	}

	err = json.Unmarshal(body, response)
	return err
}

func checkRedirect(redirRequest *http.Request, via []*http.Request) error {
	// Go's http.DefaultClient does not forward headers when a redirect 3xx
	// response is received. Thus, the header (which in this case contains the
	// Authorization token) needs to be passed forward to the redirect
	// destinations.
	redirRequest.Header = via[0].Header

	// Go's http.DefaultClient allows 10 redirects before returning an
	// an error. We have mimicked this default behavior.s
	if len(via) >= 10 {
		return errors.New("stopped after 10 redirects")
	}
	return nil
}

func getAccessToken(cfg *config.Config) (tokenResponse, error) {
	var tokenResp tokenResponse
	authURL := fmt.Sprintf(cfg.TokenURL+"?client_id=%s&client_secret=%s&code=%s&grant_type=authorization_code", cfg.ClientID, cfg.ClientSecret, cfg.AuthCode)
	resp, err := http.Post(authURL, "x-www-form-urlencoded", nil)
	if err != nil {
		return tokenResp, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return tokenResp, err
	}

	if resp.StatusCode != 200 {
		return tokenResp, fmt.Errorf("accesstoken failed: %s\n%s", resp.Status, body)
	}

	err = json.Unmarshal(body, &tokenResp)
	return tokenResp, err
}
