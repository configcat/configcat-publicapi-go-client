# \FeatureFlagSettingValuesUsingSDKKeyV2API

All URIs are relative to *https://api.configcat.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSettingValueBySdkkeyV2**](FeatureFlagSettingValuesUsingSDKKeyV2API.md#GetSettingValueBySdkkeyV2) | **Get** /v2/settings/{settingKeyOrId}/value | Get value
[**ReplaceSettingValueBySdkkeyV2**](FeatureFlagSettingValuesUsingSDKKeyV2API.md#ReplaceSettingValueBySdkkeyV2) | **Put** /v2/settings/{settingKeyOrId}/value | Replace value
[**UpdateSettingValueBySdkkeyV2**](FeatureFlagSettingValuesUsingSDKKeyV2API.md#UpdateSettingValueBySdkkeyV2) | **Patch** /v2/settings/{settingKeyOrId}/value | Update value



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
	openapiclient "github.com/configcat/configcat-publicapi-go-client/v3"
)

func main() {
	settingKeyOrId := "settingKeyOrId_example" // string | The key or id of the Setting.
	xCONFIGCATSDKKEY := "xCONFIGCATSDKKEY_example" // string | The ConfigCat SDK Key. (https://app.configcat.com/sdkkey) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FeatureFlagSettingValuesUsingSDKKeyV2API.GetSettingValueBySdkkeyV2(context.Background(), settingKeyOrId).XCONFIGCATSDKKEY(xCONFIGCATSDKKEY).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagSettingValuesUsingSDKKeyV2API.GetSettingValueBySdkkeyV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSettingValueBySdkkeyV2`: SettingFormulaModel
	fmt.Fprintf(os.Stdout, "Response from `FeatureFlagSettingValuesUsingSDKKeyV2API.GetSettingValueBySdkkeyV2`: %v\n", resp)
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

> SettingFormulaModel ReplaceSettingValueBySdkkeyV2(ctx, settingKeyOrId).UpdateEvaluationFormulaModel(updateEvaluationFormulaModel).Reason(reason).BypassApproval(bypassApproval).LatestVersionId(latestVersionId).XCONFIGCATSDKKEY(xCONFIGCATSDKKEY).Execute()

Replace value



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/configcat/configcat-publicapi-go-client/v3"
)

func main() {
	settingKeyOrId := "settingKeyOrId_example" // string | The key or id of the Setting.
	updateEvaluationFormulaModel := *openapiclient.NewUpdateEvaluationFormulaModel(*openapiclient.NewUpdateValueModel()) // UpdateEvaluationFormulaModel | 
	reason := "reason_example" // string | The reason note for the Audit Log if the Product's \"Config changes require a reason\" preference is turned on. (optional)
	bypassApproval := true // bool | Whether to bypass the approval process and directly apply the change. This is only applicable for users with bypass approval permission. (optional)
	latestVersionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional. The version identifier of the last change made to the Feature Flag or Setting in the Environment. It can be used to make sure concurrent updates are not overwriting each other. If provided and the version identifier does not match the current version, the update will be rejected with a 409 Conflict response. The latest version id can be acquired from the `LastVersionId` property of the response models. (optional)
	xCONFIGCATSDKKEY := "xCONFIGCATSDKKEY_example" // string | The ConfigCat SDK Key. (https://app.configcat.com/sdkkey) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FeatureFlagSettingValuesUsingSDKKeyV2API.ReplaceSettingValueBySdkkeyV2(context.Background(), settingKeyOrId).UpdateEvaluationFormulaModel(updateEvaluationFormulaModel).Reason(reason).BypassApproval(bypassApproval).LatestVersionId(latestVersionId).XCONFIGCATSDKKEY(xCONFIGCATSDKKEY).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagSettingValuesUsingSDKKeyV2API.ReplaceSettingValueBySdkkeyV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceSettingValueBySdkkeyV2`: SettingFormulaModel
	fmt.Fprintf(os.Stdout, "Response from `FeatureFlagSettingValuesUsingSDKKeyV2API.ReplaceSettingValueBySdkkeyV2`: %v\n", resp)
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
 **bypassApproval** | **bool** | Whether to bypass the approval process and directly apply the change. This is only applicable for users with bypass approval permission. | 
 **latestVersionId** | **string** | Optional. The version identifier of the last change made to the Feature Flag or Setting in the Environment. It can be used to make sure concurrent updates are not overwriting each other. If provided and the version identifier does not match the current version, the update will be rejected with a 409 Conflict response. The latest version id can be acquired from the &#x60;LastVersionId&#x60; property of the response models. | 
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

> SettingFormulaModel UpdateSettingValueBySdkkeyV2(ctx, settingKeyOrId).JsonPatchOperation(jsonPatchOperation).Reason(reason).BypassApproval(bypassApproval).LatestVersionId(latestVersionId).XCONFIGCATSDKKEY(xCONFIGCATSDKKEY).Execute()

Update value



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/configcat/configcat-publicapi-go-client/v3"
)

func main() {
	settingKeyOrId := "settingKeyOrId_example" // string | The key or id of the Setting.
	jsonPatchOperation := []openapiclient.JsonPatchOperation{*openapiclient.NewJsonPatchOperation(openapiclient.OperationType("unknown"), "Path_example")} // []JsonPatchOperation | 
	reason := "reason_example" // string | The reason note for the Audit Log if the Product's \"Config changes require a reason\" preference is turned on. (optional)
	bypassApproval := true // bool | Whether to bypass the approval process and directly apply the change. This is only applicable for users with bypass approval permission. (optional)
	latestVersionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional. The version identifier of the last change made to the Feature Flag or Setting in the Environment. It can be used to make sure concurrent updates are not overwriting each other. If provided and the version identifier does not match the current version, the update will be rejected with a 409 Conflict response. The latest version id can be acquired from the `LastVersionId` property of the response models. (optional)
	xCONFIGCATSDKKEY := "xCONFIGCATSDKKEY_example" // string | The ConfigCat SDK Key. (https://app.configcat.com/sdkkey) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FeatureFlagSettingValuesUsingSDKKeyV2API.UpdateSettingValueBySdkkeyV2(context.Background(), settingKeyOrId).JsonPatchOperation(jsonPatchOperation).Reason(reason).BypassApproval(bypassApproval).LatestVersionId(latestVersionId).XCONFIGCATSDKKEY(xCONFIGCATSDKKEY).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagSettingValuesUsingSDKKeyV2API.UpdateSettingValueBySdkkeyV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSettingValueBySdkkeyV2`: SettingFormulaModel
	fmt.Fprintf(os.Stdout, "Response from `FeatureFlagSettingValuesUsingSDKKeyV2API.UpdateSettingValueBySdkkeyV2`: %v\n", resp)
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
 **bypassApproval** | **bool** | Whether to bypass the approval process and directly apply the change. This is only applicable for users with bypass approval permission. | 
 **latestVersionId** | **string** | Optional. The version identifier of the last change made to the Feature Flag or Setting in the Environment. It can be used to make sure concurrent updates are not overwriting each other. If provided and the version identifier does not match the current version, the update will be rejected with a 409 Conflict response. The latest version id can be acquired from the &#x60;LastVersionId&#x60; property of the response models. | 
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

