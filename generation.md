1. Regenerate public api
Linux:
```docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli:v7.9.0 generate -i https://test-api.configcat.com/docs/v1/swagger.json -g go -o /local --git-repo-id configcat-publicapi-go-client/v2 --git-user-id configcat --additional-properties=packageName=configcatpublicapi,enumClassPrefix=true,structPrefix=true```
Windows: 
```docker run --rm -v %CD%:/local openapitools/openapi-generator-cli:v7.9.0 generate -i https://test-api.configcat.com/docs/v1/swagger.json -g go -o /local --git-repo-id configcat-publicapi-go-client/v2 --git-user-id configcat --additional-properties=packageName=configcatpublicapi,enumClassPrefix=true,structPrefix=true```
