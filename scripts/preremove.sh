#!/usr/bin/env bash

# Remove from systemd
systemctl stop gothermostat --quiet --force
systemctl disable gothermostat --quiet --force

systemctl daemon-reload
systemctl reset-failed
