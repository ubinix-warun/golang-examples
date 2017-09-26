```
docker run -d --name sonarqube -p 9000:9000 -p 9092:9092 sonarqube


docker cp ~/Downloads/sonar-golang-plugin-1.2.10.jar sonarqube:/extensions/plugins

docker restart sonarqube 

sonar-scanner

--------------------------

go get -u gopkg.in/alecthomas/gometalinter.v1
./bin/gometalinter.v1 --install
./bin/gometalinter.v1 --checkstyle > report.xml

#go get github.com/axw/gocov/...
#go get github.com/AlekSi/gocov-xml

# on src/echo
#    go test -coverprofile=cover.out
#    gocov convert cover.out | gocov-xml > coverage.xml

#   go test -covermode=count -coverprofile=cover.out

go get -u github.com/jstemmer/go-junit-report
go get github.com/stretchr/testify/assert

go test -v ./... | go-junit-report > test.xml


```