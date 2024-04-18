# \FeatureFlagSettingValuesApi

All URIs are relative to *https://test-api.configcat.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSettingValue**](FeatureFlagSettingValuesApi.md#GetSettingValue) | **Get** /v1/environments/{environmentId}/settings/{settingId}/value | Get value
[**GetSettingValues**](FeatureFlagSettingValuesApi.md#GetSettingValues) | **Get** /v1/configs/{configId}/environments/{environmentId}/values | Get values
[**PostSettingValues**](FeatureFlagSettingValuesApi.md#PostSettingValues) | **Post** /v1/configs/{configId}/environments/{environmentId}/values | Post values
[**ReplaceSettingValue**](FeatureFlagSettingValuesApi.md#ReplaceSettingValue) | **Put** /v1/environments/{environmentId}/settings/{settingId}/value | Replace value
[**UpdateSettingValue**](FeatureFlagSettingValuesApi.md#UpdateSettingValue) | **Patch** /v1/environments/{environmentId}/settings/{settingId}/value | Update value



## GetSettingValue

> SettingValueModel GetSettingValue(ctx, environmentId, settingId).Execute()

Get value



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
    environmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The identifier of the Environment.
    settingId := int32(56) // int32 | The id of the Setting.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeatureFlagSettingValuesApi.GetSettingValue(context.Background(), environmentId, settingId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagSettingValuesApi.GetSettingValue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSettingValue`: SettingValueModel
    fmt.Fprintf(os.Stdout, "Response from `FeatureFlagSettingValuesApi.GetSettingValue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | The identifier of the Environment. | 
**settingId** | **int32** | The id of the Setting. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSettingValueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SettingValueModel**](SettingValueModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSettingValues

> ConfigSettingValuesModel GetSettingValues(ctx, configId, environmentId).Execute()

Get values



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
    configId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The identifier of the Config.
    environmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The identifier of the Environment.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeatureFlagSettingValuesApi.GetSettingValues(context.Background(), configId, environmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagSettingValuesApi.GetSettingValues``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSettingValues`: ConfigSettingValuesModel
    fmt.Fprintf(os.Stdout, "Response from `FeatureFlagSettingValuesApi.GetSettingValues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configId** | **string** | The identifier of the Config. | 
**environmentId** | **string** | The identifier of the Environment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSettingValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ConfigSettingValuesModel**](ConfigSettingValuesModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSettingValues

> ConfigSettingValuesModel PostSettingValues(ctx, configId, environmentId).UpdateSettingValuesWithIdModel(updateSettingValuesWithIdModel).Reason(reason).Execute()

Post values



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
    configId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The identifier of the Config.
    environmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The identifier of the Environment.
    updateSettingValuesWithIdModel := *openapiclient.NewUpdateSettingValuesWithIdModel() // UpdateSettingValuesWithIdModel | 
    reason := "reason_example" // string | The reason note for the Audit Log if the Product's \"Config changes require a reason\" preference is turned on. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeatureFlagSettingValuesApi.PostSettingValues(context.Background(), configId, environmentId).UpdateSettingValuesWithIdModel(updateSettingValuesWithIdModel).Reason(reason).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagSettingValuesApi.PostSettingValues``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostSettingValues`: ConfigSettingValuesModel
    fmt.Fprintf(os.Stdout, "Response from `FeatureFlagSettingValuesApi.PostSettingValues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configId** | **string** | The identifier of the Config. | 
**environmentId** | **string** | The identifier of the Environment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSettingValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateSettingValuesWithIdModel** | [**UpdateSettingValuesWithIdModel**](UpdateSettingValuesWithIdModel.md) |  | 
 **reason** | **string** | The reason note for the Audit Log if the Product&#39;s \&quot;Config changes require a reason\&quot; preference is turned on. | 

### Return type

[**ConfigSettingValuesModel**](ConfigSettingValuesModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceSettingValue

> SettingValueModel ReplaceSettingValue(ctx, environmentId, settingId).UpdateSettingValueModel(updateSettingValueModel).Reason(reason).Execute()

Replace value



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
    environmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The identifier of the Environment.
    settingId := int32(56) // int32 | The id of the Setting.
    updateSettingValueModel := *openapiclient.NewUpdateSettingValueModel() // UpdateSettingValueModel | 
    reason := "reason_example" // string | The reason note for the Audit Log if the Product's \"Config changes require a reason\" preference is turned on. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeatureFlagSettingValuesApi.ReplaceSettingValue(context.Background(), environmentId, settingId).UpdateSettingValueModel(updateSettingValueModel).Reason(reason).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagSettingValuesApi.ReplaceSettingValue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceSettingValue`: SettingValueModel
    fmt.Fprintf(os.Stdout, "Response from `FeatureFlagSettingValuesApi.ReplaceSettingValue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | The identifier of the Environment. | 
**settingId** | **int32** | The id of the Setting. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceSettingValueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateSettingValueModel** | [**UpdateSettingValueModel**](UpdateSettingValueModel.md) |  | 
 **reason** | **string** | The reason note for the Audit Log if the Product&#39;s \&quot;Config changes require a reason\&quot; preference is turned on. | 

### Return type

[**SettingValueModel**](SettingValueModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSettingValue

> SettingValueModel UpdateSettingValue(ctx, environmentId, settingId).JsonPatchOperation(jsonPatchOperation).Reason(reason).Execute()

Update value



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
    environmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The identifier of the Environment.
    settingId := int32(56) // int32 | The id of the Setting.
    jsonPatchOperation := []openapiclient.JsonPatchOperation{*openapiclient.NewJsonPatchOperation(openapiclient.OperationType("unknown"), "Path_example")} // []JsonPatchOperation | 
    reason := "reason_example" // string | The reason note for the Audit Log if the Product's \"Config changes require a reason\" preference is turned on. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeatureFlagSettingValuesApi.UpdateSettingValue(context.Background(), environmentId, settingId).JsonPatchOperation(jsonPatchOperation).Reason(reason).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagSettingValuesApi.UpdateSettingValue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSettingValue`: SettingValueModel
    fmt.Fprintf(os.Stdout, "Response from `FeatureFlagSettingValuesApi.UpdateSettingValue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | The identifier of the Environment. | 
**settingId** | **int32** | The id of the Setting. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSettingValueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **jsonPatchOperation** | [**[]JsonPatchOperation**](JsonPatchOperation.md) |  | 
 **reason** | **string** | The reason note for the Audit Log if the Product&#39;s \&quot;Config changes require a reason\&quot; preference is turned on. | 

### Return type

[**SettingValueModel**](SettingValueModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

