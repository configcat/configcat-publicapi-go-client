# \FeatureFlagSettingValuesUsingSDKKeyV2Api

All URIs are relative to *https://api.configcat.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSettingValueBySdkkeyV2**](FeatureFlagSettingValuesUsingSDKKeyV2Api.md#GetSettingValueBySdkkeyV2) | **Get** /v2/settings/{settingKeyOrId}/value | Get value
[**ReplaceSettingValueBySdkkeyV2**](FeatureFlagSettingValuesUsingSDKKeyV2Api.md#ReplaceSettingValueBySdkkeyV2) | **Put** /v2/settings/{settingKeyOrId}/value | Replace value
[**UpdateSettingValueBySdkkeyV2**](FeatureFlagSettingValuesUsingSDKKeyV2Api.md#UpdateSettingValueBySdkkeyV2) | **Patch** /v2/settings/{settingKeyOrId}/value | Update value



## GetSettingValueBySdkkeyV2

> SettingFormulaModel GetSettingValueBySdkkeyV2(ctx, settingKeyOrId).XCONFIGCATSDKKEY(xCONFIGCATSDKKEY).Execute()

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
    resp, r, err := apiClient.FeatureFlagSettingValuesUsingSDKKeyV2Api.GetSettingValueBySdkkeyV2(context.Background(), settingKeyOrId).XCONFIGCATSDKKEY(xCONFIGCATSDKKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagSettingValuesUsingSDKKeyV2Api.GetSettingValueBySdkkeyV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSettingValueBySdkkeyV2`: SettingFormulaModel
    fmt.Fprintf(os.Stdout, "Response from `FeatureFlagSettingValuesUsingSDKKeyV2Api.GetSettingValueBySdkkeyV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**settingKeyOrId** | **string** | The key or id of the Setting. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSettingValueBySdkkeyV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xCONFIGCATSDKKEY** | **string** | The ConfigCat SDK Key. (https://app.configcat.com/sdkkey) | 

### Return type

[**SettingFormulaModel**](SettingFormulaModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceSettingValueBySdkkeyV2

> SettingFormulaModel ReplaceSettingValueBySdkkeyV2(ctx, settingKeyOrId).UpdateEvaluationFormulaModel(updateEvaluationFormulaModel).Reason(reason).XCONFIGCATSDKKEY(xCONFIGCATSDKKEY).Execute()

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
    updateEvaluationFormulaModel := *openapiclient.NewUpdateEvaluationFormulaModel(*openapiclient.NewValueModel()) // UpdateEvaluationFormulaModel | 
    reason := "reason_example" // string | The reason note for the Audit Log if the Product's \"Config changes require a reason\" preference is turned on. (optional)
    xCONFIGCATSDKKEY := "xCONFIGCATSDKKEY_example" // string | The ConfigCat SDK Key. (https://app.configcat.com/sdkkey) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeatureFlagSettingValuesUsingSDKKeyV2Api.ReplaceSettingValueBySdkkeyV2(context.Background(), settingKeyOrId).UpdateEvaluationFormulaModel(updateEvaluationFormulaModel).Reason(reason).XCONFIGCATSDKKEY(xCONFIGCATSDKKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagSettingValuesUsingSDKKeyV2Api.ReplaceSettingValueBySdkkeyV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceSettingValueBySdkkeyV2`: SettingFormulaModel
    fmt.Fprintf(os.Stdout, "Response from `FeatureFlagSettingValuesUsingSDKKeyV2Api.ReplaceSettingValueBySdkkeyV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**settingKeyOrId** | **string** | The key or id of the Setting. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceSettingValueBySdkkeyV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateEvaluationFormulaModel** | [**UpdateEvaluationFormulaModel**](UpdateEvaluationFormulaModel.md) |  | 
 **reason** | **string** | The reason note for the Audit Log if the Product&#39;s \&quot;Config changes require a reason\&quot; preference is turned on. | 
 **xCONFIGCATSDKKEY** | **string** | The ConfigCat SDK Key. (https://app.configcat.com/sdkkey) | 

### Return type

[**SettingFormulaModel**](SettingFormulaModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSettingValueBySdkkeyV2

> SettingFormulaModel UpdateSettingValueBySdkkeyV2(ctx, settingKeyOrId).JsonPatchOperation(jsonPatchOperation).Reason(reason).XCONFIGCATSDKKEY(xCONFIGCATSDKKEY).Execute()

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
    jsonPatchOperation := []openapiclient.JsonPatchOperation{*openapiclient.NewJsonPatchOperation("Op_example", "Path_example")} // []JsonPatchOperation | 
    reason := "reason_example" // string | The reason note for the Audit Log if the Product's \"Config changes require a reason\" preference is turned on. (optional)
    xCONFIGCATSDKKEY := "xCONFIGCATSDKKEY_example" // string | The ConfigCat SDK Key. (https://app.configcat.com/sdkkey) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeatureFlagSettingValuesUsingSDKKeyV2Api.UpdateSettingValueBySdkkeyV2(context.Background(), settingKeyOrId).JsonPatchOperation(jsonPatchOperation).Reason(reason).XCONFIGCATSDKKEY(xCONFIGCATSDKKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagSettingValuesUsingSDKKeyV2Api.UpdateSettingValueBySdkkeyV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSettingValueBySdkkeyV2`: SettingFormulaModel
    fmt.Fprintf(os.Stdout, "Response from `FeatureFlagSettingValuesUsingSDKKeyV2Api.UpdateSettingValueBySdkkeyV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**settingKeyOrId** | **string** | The key or id of the Setting. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSettingValueBySdkkeyV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jsonPatchOperation** | [**[]JsonPatchOperation**](JsonPatchOperation.md) |  | 
 **reason** | **string** | The reason note for the Audit Log if the Product&#39;s \&quot;Config changes require a reason\&quot; preference is turned on. | 
 **xCONFIGCATSDKKEY** | **string** | The ConfigCat SDK Key. (https://app.configcat.com/sdkkey) | 

### Return type

[**SettingFormulaModel**](SettingFormulaModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

