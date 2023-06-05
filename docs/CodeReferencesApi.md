# \CodeReferencesApi

All URIs are relative to *https://api.configcat.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1CodeReferencesDeleteReportsPost**](CodeReferencesApi.md#V1CodeReferencesDeleteReportsPost) | **Post** /v1/code-references/delete-reports | 
[**V1CodeReferencesPost**](CodeReferencesApi.md#V1CodeReferencesPost) | **Post** /v1/code-references | 



## V1CodeReferencesDeleteReportsPost

> V1CodeReferencesDeleteReportsPost(ctx).DeleteRepositoryReportsRequest(deleteRepositoryReportsRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/configcat/configcat-publicapi-go-client"
)

func main() {
    deleteRepositoryReportsRequest := *openapiclient.NewDeleteRepositoryReportsRequest("ConfigId_example", "Repository_example") // DeleteRepositoryReportsRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CodeReferencesApi.V1CodeReferencesDeleteReportsPost(context.Background()).DeleteRepositoryReportsRequest(deleteRepositoryReportsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CodeReferencesApi.V1CodeReferencesDeleteReportsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1CodeReferencesDeleteReportsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteRepositoryReportsRequest** | [**DeleteRepositoryReportsRequest**](DeleteRepositoryReportsRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CodeReferencesPost

> V1CodeReferencesPost(ctx).CodeReferenceRequest(codeReferenceRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/configcat/configcat-publicapi-go-client"
)

func main() {
    codeReferenceRequest := *openapiclient.NewCodeReferenceRequest("ConfigId_example", "Repository_example", "Branch_example") // CodeReferenceRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CodeReferencesApi.V1CodeReferencesPost(context.Background()).CodeReferenceRequest(codeReferenceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CodeReferencesApi.V1CodeReferencesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1CodeReferencesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **codeReferenceRequest** | [**CodeReferenceRequest**](CodeReferenceRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

