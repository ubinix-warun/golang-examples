#!/bin/bash

rm app/ client/ tool/ swagger/ bottle.go main.go  -Rf
../../bin/goagen bootstrap -d cellar/design

echo Start Swagger UI
docker run -ti --rm -p 8888:8080 -e SWAGGER_JSON=/swagger/swagger.json -v `pwd`/swagger:/swagger swaggerapi/swagger-ui