package config

import (
	"io/ioutil"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

const configFilePath = "../dist.config.json"

func TestLoadConfig(t *testing.T) {
	Convey("Given loading a config file", t, func() {
		config := New()
		So(config, ShouldNotBeNil)
		config.Load(configFilePath)

		Convey("Config values should be there", func() {
			So(config.Webserver, ShouldNotBeEmpty)
			So(config.AuthURL, ShouldNotBeEmpty)
			So(config.TokenURL, ShouldNotBeEmpty)
			So(config.ClientID, ShouldBeEmpty)
			So(config.ClientSecret, ShouldBeEmpty)
			So(config.AuthCode, ShouldBeEmpty)
			So(config.AccessToken, ShouldBeEmpty)
		})
	})
}
func TestSaveConfig(t *testing.T) {
	Convey("Given saving a config file", t, func() {
		config := New()
		config.Load(configFilePath)
		So(config, ShouldNotBeNil)

		Convey("Given the config is saved", func() {
			loadedConfigJSON, _ := JSONMarshal(config)
			config.Save(configFilePath)
			savedConfigJSON, _ := ioutil.ReadFile(configFilePath)
			So(string(loadedConfigJSON), ShouldEqual, string(savedConfigJSON))
		})
	})
}
