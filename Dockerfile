FROM golang:alpine

WORKDIR /opt/drtest

ADD . .

# ADD schema

RUN cd /opt/drtest/gen && go run main.go avro.go

CMD ["go", "run", "/opt/drtest/webserver/cmd/app.go"]