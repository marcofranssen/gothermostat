FROM alpine
ENTRYPOINT ["/usr/local/bin/gotherm"]
VOLUME /etc/marcofranssen/gothermostat
WORKDIR /etc/marcofranssen/gothermostat
EXPOSE 8888
COPY .gotherm.toml.dist /etc/marcofranssen/gothermostat/.gotherm.toml
COPY gotherm /usr/local/bin
# required for binaries build on windows
RUN chmod +x /usr/local/bin/gotherm
