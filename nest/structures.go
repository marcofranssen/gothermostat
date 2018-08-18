package nest

import "time"

type apiError struct {
	Error       string `json:"error,omitempty"`
	Description string `json:"error_description,omitempty"`
	InstanceID  string `json:"instance_id,omitempty"`
}

type token struct {
	AccessToken string `json:"access_token,omitempty"`
	ExpiresIn   int    `json:"expires_in,omitempty"`
}

// Devices Nest device collection holding Thermostats
type Devices struct {
	Thermostats map[string]*Thermostat `json:"thermostats,omitempty"`
}

// Thermostat Nest thermostat holding thermostat data and settings
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

// Where Structure location
type Where struct {
	WhereID string `json:"where_id"`
	Name    string `json:"name"`
}

// Structure Nest structure
type Structure struct {
	Name        string            `json:"name"`
	CountryCode string            `json:"country_code"`
	TimeZone    string            `json:"time_zone"`
	Away        string            `json:"away"`
	Thermostats []string          `json:"thermostats"`
	StructureID string            `json:"structure_id"`
	Wheres      map[string]*Where `json:"wheres"`
}

// Metadata Nest API metadata holding client data
type Metadata struct {
	AccessToken   string `json:"access_token"`
	ClientVersion int    `json:"client_version"`
	UserID        string `json:"user_id"`
}

// Combined Combines all Devices, Structures and Metadata coming from Nest API
type Combined struct {
	Devices    Devices               `json:"devices"`
	Structures map[string]*Structure `json:"structures"`
	Metadata   Metadata              `json:"metadata"`
}
