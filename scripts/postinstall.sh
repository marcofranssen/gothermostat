#!/usr/bin/env bash

echo "- Set application permissions"
chown -R root:adm /var/log/marcofranssen
chown -R gothermostat:gothermostat {/etc,/var/log}/marcofranssen/gothermostat
chmod +x /usr/local/bin/gotherm

# restore config backup if exists
echo "- Restoring config"
cp /etc/marcofranssen/gothermostat/.gotherm.toml.bak /etc/marcofranssen/gothermostat/.gotherm.toml || true 2>/dev/null
