# Go Thermostat

[![Build Status](https://travis-ci.org/marcofranssen/gothermostat.svg?branch=master)](https://travis-ci.org/marcofranssen/gothermostat)
[![Software License](https://img.shields.io/badge/License-MIT-orange.svg?style=flat-square)](https://github.com/marcofranssen/gothermostat/blob/master/LICENSE.md)
[![GoDoc](https://godoc.org/github.com/marcofranssen/gothermostat?status.svg)](https://godoc.org/github.com/marcofranssen/gothermostat)
[![Coverage Status](http://codecov.io/github/marcofranssen/gothermostat/coverage.svg?branch=master)](http://codecov.io/github/marcofranssen/gothermostat?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/marcofranssen/gothermostat)](https://goreportcard.com/report/github.com/marcofranssen/gothermostat)

This project enables to interact with your nest thermostat.

## Build

```bash
go build .
```

## Configure

```bash
cp dist.config.json config.json
```

Fill out your clientId, clientSecret, authCode.

### How to get your nest token

https://codelabs.developers.google.com/codelabs/wwn-api-quickstart/#2

## Run

```
./gothermostat
```

## Demo

![web](doc/web.jpg)
