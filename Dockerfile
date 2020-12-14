FROM golang:alpine

WORKDIR /opt/drtest

COPY . .

RUN chmod +x entrypoint.sh

RUN go mod download

VOLUME /opt/schemafiles

RUN cd gen && go build -o drtest-gen && mv drtest-gen /opt/drtest

EXPOSE 8080

ENTRYPOINT ["/opt/drtest/entrypoint.sh"]
