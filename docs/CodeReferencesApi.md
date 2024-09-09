# \CodeReferencesApi

All URIs are relative to *https://api.configcat.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1CodeReferencesDeleteReportsPost**](CodeReferencesApi.md#V1CodeReferencesDeleteReportsPost) | **Post** /v1/code-references/delete-reports | Delete Reference reports
[**V1CodeReferencesPost**](CodeReferencesApi.md#V1CodeReferencesPost) | **Post** /v1/code-references | Upload References
[**V1SettingsSettingIdCodeReferencesGet**](CodeReferencesApi.md#V1SettingsSettingIdCodeReferencesGet) | **Get** /v1/settings/{settingId}/code-references | Get References for Feature Flag or Setting



## V1CodeReferencesDeleteReportsPost

> V1CodeReferencesDeleteReportsPost(ctx).DeleteRepositoryReportsRequest(deleteRepositoryReportsRequest).Execute()

Delete Reference reports



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

Upload References



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


## V1SettingsSettingIdCodeReferencesGet

> []CodeReferenceModel V1SettingsSettingIdCodeReferencesGet(ctx, settingId).Execute()

Get References for Feature Flag or Setting



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
    settingId := int32(56) // int32 | The identifier of the Feature Flag or Setting.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CodeReferencesApi.V1SettingsSettingIdCodeReferencesGet(context.Background(), settingId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CodeReferencesApi.V1SettingsSettingIdCodeReferencesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1SettingsSettingIdCodeReferencesGet`: []CodeReferenceModel
    fmt.Fprintf(os.Stdout, "Response from `CodeReferencesApi.V1SettingsSettingIdCodeReferencesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**settingId** | **int32** | The identifier of the Feature Flag or Setting. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1SettingsSettingIdCodeReferencesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]CodeReferenceModel**](CodeReferenceModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

