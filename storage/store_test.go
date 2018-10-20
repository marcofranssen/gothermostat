package storage

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/marcofranssen/gothermostat/nest"
	. "github.com/smartystreets/goconvey/convey"
)

var storeJSON1 = []byte(`{
    "id": "b6UaPCSjpgE56SMNlaigf4Sir6gJ8Twj",
    "name": "Living Room",
    "temperatures": [
        {
            "timestamp": "2018-09-01T08:57:25.1719729+02:00",
            "ambientTemperatureC": 21,
            "ambientTemperatureF": 70,
            "targetTemperatureC": 14,
            "targetTemperatureF": 57
        }
    ]
}`)

var storeJSON2 = []byte(`{
    "id": "b6UaPCSjpgE56SMNlaigf4Sir6gJ8Twj",
    "name": "Living Room",
    "temperatures": [
        {
            "timestamp": "2018-09-01T08:57:25.1719729+02:00",
            "ambientTemperatureC": 21,
            "ambientTemperatureF": 70,
            "targetTemperatureC": 14,
            "targetTemperatureF": 57
        }
    ]
}`)

func TestUnmarshalStoredJson(t *testing.T) {
	Convey("Given one stored temperature", t, func() {
		s := NewStore("./test-data", 20)
		storedData1, err := s.unmarshal(storeJSON1)
		So(err, ShouldBeNil)
		So(storedData1, ShouldNotBeNil)
		So(len(storedData1.Temperatures), ShouldEqual, 1)
	})
}

// func TestGetTemperatureData(t *testing.T) {
// 	Convey("Given 2 stored files", t, func() {
// 		s := NewStore("./test-data")
// 		storedData1, err := s.unmarshal(storeJSON1)
// 		So(err, ShouldBeNil)
// 		So(storedData1, ShouldNotBeNil)
// 		storedData2, err := s.unmarshal(storeJSON2)
// 		So(err, ShouldBeNil)
// 		So(storedData2, ShouldNotBeNil)

// 		Convey("Retrieving temperatureData returns both thermostats", func() {
// 			result, err := s.GetTemperatureData()
// 			So(err, ShouldBeNil)
// 			So(result, ShouldNotBeNil)
// 			So(len(result["thermostats"]), ShouldEqual, 2)
// 		})
// 	})
// }

func TestAddTemperateToThermostatData(t *testing.T) {
	Convey("Given an exisiting json", t, func() {
		thermoJSON := []byte(`{
            "b6UaPCSjpgE56SMNlaigf4Sir6gJ8Twj": {
                "id": "b6UaPCSjpgE56SMNlaigf4Sir6gJ8Twj",
                "name": "Chill place",
                "ambient_temperature_c": 20.5,
                "ambient_temperature_f": 69,
                "target_temperature_c": 14,
                "target_temperature_f": 57
            }
        }`)

		var thermoData map[string]*nest.Thermostat
		err := json.Unmarshal(thermoJSON, &thermoData)
		So(err, ShouldBeNil)

		s := NewStore("./test-data", 20)
		storedData, err := s.unmarshal(storeJSON1)
		So(err, ShouldBeNil)
		So(storedData, ShouldNotBeNil)

		newData := s.updateData(storedData, thermoData["b6UaPCSjpgE56SMNlaigf4Sir6gJ8Twj"], time.Now())
		So(newData, ShouldNotBeNil)

		Convey("temperature should be added to thermostat", func() {
			So(newData.Name, ShouldEqual, "Chill place")
			So(len(newData.Temperatures), ShouldEqual, 2)
			So(newData.Temperatures[0].AmbientTemperatureC, ShouldEqual, 21)
			So(newData.Temperatures[0].AmbientTemperatureF, ShouldEqual, 70)
			So(newData.Temperatures[0].TargetTemperatureC, ShouldEqual, 14)
			So(newData.Temperatures[0].TargetTemperatureF, ShouldEqual, 57)
			So(newData.Temperatures[1].AmbientTemperatureC, ShouldEqual, 20.5)
			So(newData.Temperatures[1].AmbientTemperatureF, ShouldEqual, 69)
			So(newData.Temperatures[1].TargetTemperatureC, ShouldEqual, 14)
			So(newData.Temperatures[1].TargetTemperatureF, ShouldEqual, 57)
		})
	})
}
