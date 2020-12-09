FROM golang:alpine

WORKDIR /opt/drtest

ADD . .
RUN go mod download

# TODO: Dynamic loading of multiple schema datas
RUN go run gen/gen.go gen/avro.go --target-dir=./generated /opt/drtest/schema.avsc

CMD ["go", "run", "/opt/drtest/webserver/cmd/app.go"]