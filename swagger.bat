@echo off
echo.
docker run --rm -it --env GOPATH=/go -v %CD%:/go/src -w /go/src swaggerapi/swagger-codegen-cli-v3:latest %*