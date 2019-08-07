FROM alpine
ENTRYPOINT ["/usr/local/bin/gothermostat"]
VOLUME /etc/marcofranssen/gothermostat
WORKDIR /etc/marcofranssen/gothermostat
EXPOSE 8888
COPY .gotherm.toml /etc/marcofranssen/gothermostat/.gotherm.toml
COPY gothermostat /usr/local/bin
# required for binaries build on windows
RUN chmod +x /usr/local/bin/gothermostat
