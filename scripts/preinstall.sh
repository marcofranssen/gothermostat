#!/usr/bin/env bash

# Create application user
echo -n "- Creating application user "
if [ $(id -u "gothermostat" 2>/dev/null || echo -1) -ge 0 ]; then
    echo "SKIP: User already exists"
else
    sudo useradd -r -s /bin/false gothermostat
fi

# make backup of config if exists
cp /etc/marcofranssen/gothermostat/.gotherm.toml /etc/marcofranssen/gothermostat/.gotherm.toml.bak || true 2>/dev/null
