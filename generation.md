# configcat-publicapi-go-client

## Generation on Windows
docker run --rm -it --env GOPATH=/go -v %CD%:/go/src -w /go/src swaggerapi/swagger-codegen-cli-v3:latest generate -i https://api.configcat.com/docs/v1/swagger.json -l go  -DpackageName=configcatpublicapi

docker run --rm -it --env GOPATH=/go -v %CD%:/go/src -w /go/src quay.io/goswagger/swagger generate client -f https://api.configcat.com/docs/v1/swagger.json -A configcatpublicapiclient

## Generation on Linux
docker run --rm -v ${PWD}:/local swaggerapi/swagger-codegen-cli-v3:latest generate -i https://api.configcat.com/docs/v1/swagger.json -l go -DpackageName=configcatpublicapi