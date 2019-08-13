#!/usr/bin/env bash

echo "- Application permissions"
chown -R root:adm /var/log/marcofranssen
chown -R gothermostat:gothermostat {/etc,/var/log}/marcofranssen/gothermostat
# chmod +s {/etc,/var/log}/marcofranssen/gothermostat
chmod +x /usr/local/bin/gotherm

# restore config backup if exists
cp /etc/marcofranssen/gothermostat/.gotherm.toml.bak /etc/marcofranssen/gothermostat/.gotherm.toml || true 2>/dev/null
