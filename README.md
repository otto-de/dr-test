# Dr. Test

## Installation
Put avro schema in root directory named "schema.avsc"

Run ``docker build -t drtest .``

## Usage
Start container. Webserver starts on container port 8080 by default.

Run ``docker run -p 8080 drtest``.

Given your record's name is "Foo", your test samples can be fetched via ``/foo``.

## Without Docker
You can generate a test sample generator from multiple schemas as once.  
All you have to do is run ``go run gen/gen.go gen/avro.go --target-dir=./generated $PATH_TO_SCHEMA`` to generate the necessary methods for the webserver.
Starting the webserver ``go run webserver/cmd/app.go`` will expose one endpoint per record name.

## ToDos
### Randomizer Config
- [ ] Implemented

Make the randomizer configurable, such that
- Ranges
- Valid values
- Defaults
- etc
can be configured to be used for each field name present in the schema.

### Multiple schemas in docker container
- [ ] Implemented

Either in ``docker build`` or ``docker run`` it should be possible to load multiple avro schemas at once.

**Note**: see ["Without Docker"](Without Docker)) on how to use multiple schemas during generation, alleviating this problem. 

### Deterministic samples
- [ ] Implemented

Between two consecutive test sample fetches, data should be consistent 
given the schema, nor the requested quantity did not change.

### Generation of edge-cases
- [ ] Implemented

Edge cases for the defined entity (nullable fields, empty lists and maps, etc.) 
should be obtainable by request.

### Additional APIs
- [ ] Implemented

Test samples can also be consumed via kafka, sqs, amq, etc.

### Multiple schema technologies
- [ ] Implemented

Add protobuf, json-schema, etc. to the list of supported schema technologies.
