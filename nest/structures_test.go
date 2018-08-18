package nest

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCombined(t *testing.T) {
	Convey("Given combined json response", t, func() {
		combined := &Combined{}
		err := json.Unmarshal(combinedResponse(), combined)
		So(err, ShouldBeNil)

		Convey("We should have thermostats", func() {
			So(len(combined.Devices.Thermostats), ShouldEqual, 1)
			for key, value := range combined.Devices.Thermostats {
				So(key, ShouldEqual, combined.Devices.Thermostats[key].DeviceID)
				checkFields(value)
			}
		})

		Convey("We should have structures", func() {
			So(len(combined.Structures), ShouldEqual, 2)
			for key, value := range combined.Structures {
				So(key, ShouldEqual, combined.Structures[key].StructureID)
				checkFields(value)
			}
		})

		Convey("We should have metadata", func() {
			So(combined.Metadata, ShouldNotBeEmpty)
			So(combined.Metadata.AccessToken, ShouldEqual, "c.nyL7tUSHUkIksJu0fLs9RA97Zjt9W95yj10MgmlT8qDzaXhgYDh1DvxgY2uJGmBUCkLSr8rEleVla50jyAotkJstmyPvd7jhW2qZKWuIUmTWejeGUNxgRbaP570iA5cRSSWxGq5vc1dthBb1")
			So(combined.Metadata.ClientVersion, ShouldEqual, 1)
			So(combined.Metadata.UserID, ShouldEqual, "z.1.1.PzJKI35K0I5erCaRf20BNBNUzJg/DKXa1iBHtOyfTl5=")
		})
	})
}

func TestAccess(t *testing.T) {
	Convey("Given a JSON object with an access token", t, func() {
		access := &token{}
		err := json.Unmarshal(tokenResponse(), access)
		So(err, ShouldBeNil)
	})
}

func checkFields(value interface{}) {
	s := reflect.ValueOf(value).Elem()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		value := f.Interface()
		switch value.(type) {
		case string:
			So(value.(string), ShouldNotBeEmpty)
		case bool:
			So(value.(bool), ShouldBeTrue)
		case int:
			So(value.(int), ShouldNotBeNil)
		case float64:
			So(value.(float64), ShouldNotBeNil)
		case time.Time:
			So(value.(time.Time), ShouldNotBeNil)
		}
	}
}

func tokenResponse() []byte {
	return []byte(`{
		"access_token": "c.FmDPkzyzaQe...",
		"expires_in": 315360000
	}`)
}

func combinedResponse() []byte {
	return []byte(`{
		"devices": {
		  "thermostats": {
			"b6UaPCSjpgE56SMNlaigf4Sir6gJ8Tej": {
			  "humidity": 45,
			  "locale": "nl-NL",
			  "temperature_scale": "C",
			  "is_using_emergency_heat": true,
			  "has_fan": true,
			  "software_version": "5.8.2-1",
			  "has_leaf": true,
			  "where_id": "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQij75PUxJPcbe",
			  "device_id": "b6UaPCSjpgE56SMNlaigf4Sir6gJ8Tej",
			  "name": "Living Room",
			  "can_heat": true,
			  "can_cool": true,
			  "target_temperature_c": 14,
			  "target_temperature_f": 58,
			  "target_temperature_high_c": 24,
			  "target_temperature_high_f": 75,
			  "target_temperature_low_c": 20,
			  "target_temperature_low_f": 68,
			  "ambient_temperature_c": 24,
			  "ambient_temperature_f": 76,
			  "away_temperature_high_c": 24,
			  "away_temperature_high_f": 76,
			  "away_temperature_low_c": 10,
			  "away_temperature_low_f": 50,
			  "eco_temperature_high_c": 24,
			  "eco_temperature_high_f": 76,
			  "eco_temperature_low_c": 10,
			  "eco_temperature_low_f": 50,
			  "is_locked": true,
			  "locked_temp_min_c": 20,
			  "locked_temp_min_f": 68,
			  "locked_temp_max_c": 22,
			  "locked_temp_max_f": 72,
			  "sunlight_correction_active": true,
			  "sunlight_correction_enabled": true,
			  "structure_id": "SjDpoU_d5Mo2fwDY8BMGpeDKosmdMfHh3kLlqAh9E8yNjv2wT-oQOS",
			  "fan_timer_active": true,
			  "fan_timer_timeout": "1970-01-01T00:00:00Z",
			  "fan_timer_duration": 15,
			  "previous_hvac_mode": "heat",
			  "hvac_mode": "off",
			  "time_to_target": "~0",
			  "time_to_target_training": "ready",
			  "where_name": "Living Room",
			  "label": "LR",
			  "name_long": "Living Room Thermostat",
			  "is_online": true,
			  "last_connection": "2018-08-18T07:50:00.027Z",
			  "hvac_state": "off"
			}
		  }
		},
		"structures": {
		  "NrE8wCFsUDb_p22EUUXBDk4qUtPkRNp8Vyjb_SHh8Y63xdkW4CvmQg": {
			"name": "Structure 1",
			"country_code": "US",
			"time_zone": "America/Los_Angeles",
			"away": "home",
			"thermostats": null,
			"structure_id": "NrE8wCFsUDb_p22EUUXBDk4qUtPkRNp8Vyjb_SHh8Y63xdkW4CvmQg",
			"wheres": {
			  "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQg9OBBMX7NqtQ": {
				"where_id": "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQg9OBBMX7NqtQ",
				"name": "Master Bedroom"
			  },
			  "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQgHr_XOOg8ZQg": {
				"where_id": "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQgHr_XOOg8ZQg",
				"name": "Office"
			  },
			  "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQgSP4DIDcRIPg": {
				"where_id": "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQgSP4DIDcRIPg",
				"name": "Bathroom"
			  },
			  "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQgqiM7kRPw_3A": {
				"where_id": "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQgqiM7kRPw_3A",
				"name": "Den"
			  },
			  "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQhCjWot27XtxQ": {
				"where_id": "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQhCjWot27XtxQ",
				"name": "Guest House"
			  },
			  "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQh_N94ixM1XdA": {
				"where_id": "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQh_N94ixM1XdA",
				"name": "Guest Room"
			  },
			  "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQhbcWh2ukrk5A": {
				"where_id": "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQhbcWh2ukrk5A",
				"name": "Kids Room"
			  },
			  "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQhdHBj4PpxtEA": {
				"where_id": "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQhdHBj4PpxtEA",
				"name": "Outside"
			  },
			  "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQhgzm27SZXAkA": {
				"where_id": "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQhgzm27SZXAkA",
				"name": "Downstairs"
			  },
			  "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQhqH490onwrhQ": {
				"where_id": "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQhqH490onwrhQ",
				"name": "Back Door"
			  },
			  "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQhqU8S7STHTAQ": {
				"where_id": "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQhqU8S7STHTAQ",
				"name": "Kitchen"
			  },
			  "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQhvrmxEcxAf9A": {
				"where_id": "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQhvrmxEcxAf9A",
				"name": "Front Yard"
			  },
			  "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQi3VmCvBH-USQ": {
				"where_id": "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQi3VmCvBH-USQ",
				"name": "Family Room"
			  },
			  "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQi8S6Rj4fuJKg": {
				"where_id": "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQi8S6Rj4fuJKg",
				"name": "Upstairs"
			  },
			  "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQiNhgsgbVHMrg": {
				"where_id": "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQiNhgsgbVHMrg",
				"name": "Dining Room"
			  },
			  "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQiU63lO67Bk2Q": {
				"where_id": "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQiU63lO67Bk2Q",
				"name": "Shed"
			  },
			  "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQiYq4LdKX1UKA": {
				"where_id": "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQiYq4LdKX1UKA",
				"name": "Bedroom"
			  },
			  "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQibdAM63xMoDA": {
				"where_id": "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQibdAM63xMoDA",
				"name": "Side Door"
			  },
			  "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQie7-CdzH181Q": {
				"where_id": "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQie7-CdzH181Q",
				"name": "Basement"
			  },
			  "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQihTmB0lIGb_A": {
				"where_id": "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQihTmB0lIGb_A",
				"name": "Entryway"
			  },
			  "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQij75PUxJPxvw": {
				"where_id": "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQij75PUxJPxvw",
				"name": "Living Room"
			  },
			  "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQj-wj-PPrbHEw": {
				"where_id": "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQj-wj-PPrbHEw",
				"name": "Patio"
			  },
			  "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQj3BdanUmyLog": {
				"where_id": "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQj3BdanUmyLog",
				"name": "Driveway"
			  },
			  "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQj4BCcZfmhP5w": {
				"where_id": "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQj4BCcZfmhP5w",
				"name": "Hallway"
			  },
			  "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQjHWx7lWmV_hQ": {
				"where_id": "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQjHWx7lWmV_hQ",
				"name": "Deck"
			  },
			  "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQjK2SeBSkpl7A": {
				"where_id": "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQjK2SeBSkpl7A",
				"name": "Backyard"
			  },
			  "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQjRlx7Ne9uNkQ": {
				"where_id": "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQjRlx7Ne9uNkQ",
				"name": "Garage"
			  },
			  "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQjWLlSZqRVq5A": {
				"where_id": "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQjWLlSZqRVq5A",
				"name": "Attic"
			  },
			  "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQjoLYdCprvRiA": {
				"where_id": "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQjoLYdCprvRiA",
				"name": "Front Door"
			  }
			}
		  },
		  "SjDpoU_d5Mo2fwDY8BMGpeDKosmdMfHh3kLlqAh9E8yNjv2wT-oQOS": {
			"name": "Thuis",
			"country_code": "NL",
			"time_zone": "Europe/Amsterdam",
			"away": "away",
			"thermostats": [
			  "b6UaPCSjpgE56SMNlaigf4Sir6gJ8Twh"
			],
			"structure_id": "SjDpoU_d5Mo2fwDY8BMGpeDKosmdMfHh3kLlqAh9E8yNjv2wT-oQOS",
			"wheres": {
			  "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQg9OBBMX7NqtQ": {
				"where_id": "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQg9OBBMX7NqtQ",
				"name": "Master Bedroom"
			  },
			  "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQgHr_XOOg8ZQg": {
				"where_id": "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQgHr_XOOg8ZQg",
				"name": "Office"
			  },
			  "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQgSP4DIDcRIPg": {
				"where_id": "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQgSP4DIDcRIPg",
				"name": "Bathroom"
			  },
			  "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQgqiM7kRPw_3A": {
				"where_id": "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQgqiM7kRPw_3A",
				"name": "Den"
			  },
			  "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQhCjWot27XtxQ": {
				"where_id": "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQhCjWot27XtxQ",
				"name": "Guest House"
			  },
			  "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQh_N94ixM1XdA": {
				"where_id": "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQh_N94ixM1XdA",
				"name": "Guest Room"
			  },
			  "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQhbcWh2ukrk5A": {
				"where_id": "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQhbcWh2ukrk5A",
				"name": "Kids Room"
			  },
			  "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQhdHBj4PpxtEA": {
				"where_id": "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQhdHBj4PpxtEA",
				"name": "Outside"
			  },
			  "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQhgzm27SZXAkA": {
				"where_id": "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQhgzm27SZXAkA",
				"name": "Downstairs"
			  },
			  "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQhqH490onwrhQ": {
				"where_id": "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQhqH490onwrhQ",
				"name": "Back Door"
			  },
			  "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQhqU8S7STHTAQ": {
				"where_id": "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQhqU8S7STHTAQ",
				"name": "Kitchen"
			  },
			  "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQhvrmxEcxAf9A": {
				"where_id": "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQhvrmxEcxAf9A",
				"name": "Front Yard"
			  },
			  "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQi3VmCvBH-USQ": {
				"where_id": "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQi3VmCvBH-USQ",
				"name": "Family Room"
			  },
			  "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQi8S6Rj4fuJKg": {
				"where_id": "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQi8S6Rj4fuJKg",
				"name": "Upstairs"
			  },
			  "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQiNhgsgbVHMrg": {
				"where_id": "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQiNhgsgbVHMrg",
				"name": "Dining Room"
			  },
			  "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQiU63lO67Bk2Q": {
				"where_id": "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQiU63lO67Bk2Q",
				"name": "Shed"
			  },
			  "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQiYq4LdKX1UKA": {
				"where_id": "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQiYq4LdKX1UKA",
				"name": "Bedroom"
			  },
			  "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQibdAM63xMoDA": {
				"where_id": "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQibdAM63xMoDA",
				"name": "Side Door"
			  },
			  "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQie7-CdzH181Q": {
				"where_id": "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQie7-CdzH181Q",
				"name": "Basement"
			  },
			  "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQihTmB0lIGb_A": {
				"where_id": "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQihTmB0lIGb_A",
				"name": "Entryway"
			  },
			  "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQij75PUxJPxvw": {
				"where_id": "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQij75PUxJPxvw",
				"name": "Living Room"
			  },
			  "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQj-wj-PPrbHEw": {
				"where_id": "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQj-wj-PPrbHEw",
				"name": "Patio"
			  },
			  "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQj3BdanUmyLog": {
				"where_id": "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQj3BdanUmyLog",
				"name": "Driveway"
			  },
			  "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQj4BCcZfmhP5w": {
				"where_id": "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQj4BCcZfmhP5w",
				"name": "Hallway"
			  },
			  "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQjHWx7lWmV_hQ": {
				"where_id": "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQjHWx7lWmV_hQ",
				"name": "Deck"
			  },
			  "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQjK2SeBSkpl7A": {
				"where_id": "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQjK2SeBSkpl7A",
				"name": "Backyard"
			  },
			  "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQjRlx7Ne9uNkQ": {
				"where_id": "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQjRlx7Ne9uNkQ",
				"name": "Garage"
			  },
			  "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQjWLlSZqRVq5A": {
				"where_id": "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQjWLlSZqRVq5A",
				"name": "Attic"
			  },
			  "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQjoLYdCprvRiA": {
				"where_id": "WOcupKgmH5c-rA8W9IGGAWBtG8amcjJ89gBj5agMHQjoLYdCprvRiA",
				"name": "Front Door"
			  }
			}
		  }
		},
		"metadata": {
		  "access_token": "c.nyL7tUSHUkIksJu0fLs9RA97Zjt9W95yj10MgmlT8qDzaXhgYDh1DvxgY2uJGmBUCkLSr8rEleVla50jyAotkJstmyPvd7jhW2qZKWuIUmTWejeGUNxgRbaP570iA5cRSSWxGq5vc1dthBb1",
		  "client_version": 1,
		  "user_id": "z.1.1.PzJKI35K0I5erCaRf20BNBNUzJg/DKXa1iBHtOyfTl5="
		}
	  }`)
}
