``

cd src/cellar
../../bin/goagen bootstrap -d cellar/design

docker run -ti --rm -p 8888:8080 -e SWAGGER_JSON=/swagger/swagger.json -v `pwd`/swagger:/swagger swaggerapi/swagger-ui

``
