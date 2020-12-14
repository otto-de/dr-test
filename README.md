# Dr. Test

## Installation
Create the docker image using the following command:

```shell
docker build -t drtest .
```

## Usage
### Using run.sh
`run.sh` is meant to simplify the usage of `dr-test`. After building the docker image 
using the step described above, call the script with the schema files you want to serve 
data for.

```shell
./run.sh drtest:latest my-schema1.avsc my-schema2.avsc
```

The script will then call the docker image with the schema files mounted into the
expected directory.

### Using docker run
If you would like to use `docker run` directly, feel free to do so.

You can either call `docker run` and pass your schema files as mounted files:

```shell
docker run -v /path/to/my-schema1.avsc:/opt/schemafiles/my-schema1.avsc drtest:latest
```

Or to pass multiple schema files:

```shell
docker run -v /path/to/my-schema1.avsc:/opt/schemafiles/my-schema1.avsc -v /other/path/to/my-schema2.avsc:/opt/schemafiles/my-schema2.avsc drtest:latest
```

You can also mount an entire local directory. `dr-test` will use all files with an `.avsc` extension. 

The webserver starts on container port 8080 by default. You can overwrite it using docker's `-p` flag.

Given your record's name is `Foo`, your test samples can be fetched via ``/foo``.

```shell
curl localhost:8080/foo
```

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
- [x] Implemented

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
