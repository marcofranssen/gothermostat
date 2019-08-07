package main

import (
	"github.com/marcofranssen/gothermostat/cmd"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	v := cmd.VersionInfo{
		Version: version,
		Commit:  commit,
		Date:    cmd.ParseDate(date),
	}
	cmd.Execute(v)
}
