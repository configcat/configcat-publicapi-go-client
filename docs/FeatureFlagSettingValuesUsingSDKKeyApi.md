# \FeatureFlagSettingValuesUsingSDKKeyApi

All URIs are relative to *https://api.configcat.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSettingValueBySdkkey**](FeatureFlagSettingValuesUsingSDKKeyApi.md#GetSettingValueBySdkkey) | **Get** /v1/settings/{settingKeyOrId}/value | Get value
[**ReplaceSettingValueBySdkkey**](FeatureFlagSettingValuesUsingSDKKeyApi.md#ReplaceSettingValueBySdkkey) | **Put** /v1/settings/{settingKeyOrId}/value | Replace value
[**UpdateSettingValueBySdkkey**](FeatureFlagSettingValuesUsingSDKKeyApi.md#UpdateSettingValueBySdkkey) | **Patch** /v1/settings/{settingKeyOrId}/value | Update value



## GetSettingValueBySdkkey

> SettingValueModel GetSettingValueBySdkkey(ctx, settingKeyOrId).XCONFIGCATSDKKEY(xCONFIGCATSDKKEY).Execute()

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
    settingKeyOrId := "settingKeyOrId_example" // string | The key or id of the Setting.
    xCONFIGCATSDKKEY := "xCONFIGCATSDKKEY_example" // string | The ConfigCat SDK Key. (https://app.configcat.com/sdkkey) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeatureFlagSettingValuesUsingSDKKeyApi.GetSettingValueBySdkkey(context.Background(), settingKeyOrId).XCONFIGCATSDKKEY(xCONFIGCATSDKKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagSettingValuesUsingSDKKeyApi.GetSettingValueBySdkkey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSettingValueBySdkkey`: SettingValueModel
    fmt.Fprintf(os.Stdout, "Response from `FeatureFlagSettingValuesUsingSDKKeyApi.GetSettingValueBySdkkey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**settingKeyOrId** | **string** | The key or id of the Setting. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSettingValueBySdkkeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xCONFIGCATSDKKEY** | **string** | The ConfigCat SDK Key. (https://app.configcat.com/sdkkey) | 

### Return type

[**SettingValueModel**](SettingValueModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceSettingValueBySdkkey

> SettingValueModel ReplaceSettingValueBySdkkey(ctx, settingKeyOrId).UpdateSettingValueModel(updateSettingValueModel).Reason(reason).XCONFIGCATSDKKEY(xCONFIGCATSDKKEY).Execute()

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
    settingKeyOrId := "settingKeyOrId_example" // string | The key or id of the Setting.
    updateSettingValueModel := *openapiclient.NewUpdateSettingValueModel() // UpdateSettingValueModel | 
    reason := "reason_example" // string | The reason note for the Audit Log if the Product's \"Config changes require a reason\" preference is turned on. (optional)
    xCONFIGCATSDKKEY := "xCONFIGCATSDKKEY_example" // string | The ConfigCat SDK Key. (https://app.configcat.com/sdkkey) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeatureFlagSettingValuesUsingSDKKeyApi.ReplaceSettingValueBySdkkey(context.Background(), settingKeyOrId).UpdateSettingValueModel(updateSettingValueModel).Reason(reason).XCONFIGCATSDKKEY(xCONFIGCATSDKKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagSettingValuesUsingSDKKeyApi.ReplaceSettingValueBySdkkey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceSettingValueBySdkkey`: SettingValueModel
    fmt.Fprintf(os.Stdout, "Response from `FeatureFlagSettingValuesUsingSDKKeyApi.ReplaceSettingValueBySdkkey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**settingKeyOrId** | **string** | The key or id of the Setting. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceSettingValueBySdkkeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateSettingValueModel** | [**UpdateSettingValueModel**](UpdateSettingValueModel.md) |  | 
 **reason** | **string** | The reason note for the Audit Log if the Product&#39;s \&quot;Config changes require a reason\&quot; preference is turned on. | 
 **xCONFIGCATSDKKEY** | **string** | The ConfigCat SDK Key. (https://app.configcat.com/sdkkey) | 

### Return type

[**SettingValueModel**](SettingValueModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSettingValueBySdkkey

> SettingValueModel UpdateSettingValueBySdkkey(ctx, settingKeyOrId).JsonPatch(jsonPatch).Reason(reason).XCONFIGCATSDKKEY(xCONFIGCATSDKKEY).Execute()

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
    settingKeyOrId := "settingKeyOrId_example" // string | The key or id of the Setting.
    jsonPatch := *openapiclient.NewJsonPatch() // JsonPatch | 
    reason := "reason_example" // string | The reason note for the Audit Log if the Product's \"Config changes require a reason\" preference is turned on. (optional)
    xCONFIGCATSDKKEY := "xCONFIGCATSDKKEY_example" // string | The ConfigCat SDK Key. (https://app.configcat.com/sdkkey) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeatureFlagSettingValuesUsingSDKKeyApi.UpdateSettingValueBySdkkey(context.Background(), settingKeyOrId).JsonPatch(jsonPatch).Reason(reason).XCONFIGCATSDKKEY(xCONFIGCATSDKKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagSettingValuesUsingSDKKeyApi.UpdateSettingValueBySdkkey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSettingValueBySdkkey`: SettingValueModel
    fmt.Fprintf(os.Stdout, "Response from `FeatureFlagSettingValuesUsingSDKKeyApi.UpdateSettingValueBySdkkey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**settingKeyOrId** | **string** | The key or id of the Setting. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSettingValueBySdkkeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jsonPatch** | [**JsonPatch**](JsonPatch.md) |  | 
 **reason** | **string** | The reason note for the Audit Log if the Product&#39;s \&quot;Config changes require a reason\&quot; preference is turned on. | 
 **xCONFIGCATSDKKEY** | **string** | The ConfigCat SDK Key. (https://app.configcat.com/sdkkey) | 

### Return type

[**SettingValueModel**](SettingValueModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

