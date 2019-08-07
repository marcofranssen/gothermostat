#!/usr/bin/env bash

echo "Application permissions"
chown -R gothermostat:gothermostat {/etc,/var/log}/marcofranssen/gothermostat
chmod +s {/etc,/var/log}/marcofranssen/gothermostat

# restore config backup if exists
cp /etc/marcofranssen/gothermostat/.gotherm.toml.bak /etc/marcofranssen/gothermostat/.gotherm.toml 2> /dev/null