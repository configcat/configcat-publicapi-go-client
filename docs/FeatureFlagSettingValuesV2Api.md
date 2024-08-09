# \FeatureFlagSettingValuesV2API

All URIs are relative to *https://test-api.configcat.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSettingValueV2**](FeatureFlagSettingValuesV2API.md#GetSettingValueV2) | **Get** /v2/environments/{environmentId}/settings/{settingId}/value | Get value
[**GetSettingValuesV2**](FeatureFlagSettingValuesV2API.md#GetSettingValuesV2) | **Get** /v2/configs/{configId}/environments/{environmentId}/values | Get values
[**PostSettingValuesV2**](FeatureFlagSettingValuesV2API.md#PostSettingValuesV2) | **Post** /v2/configs/{configId}/environments/{environmentId}/values | Post values
[**ReplaceSettingValueV2**](FeatureFlagSettingValuesV2API.md#ReplaceSettingValueV2) | **Put** /v2/environments/{environmentId}/settings/{settingId}/value | Replace value
[**UpdateSettingValueV2**](FeatureFlagSettingValuesV2API.md#UpdateSettingValueV2) | **Patch** /v2/environments/{environmentId}/settings/{settingId}/value | Update value



## GetSettingValueV2

> SettingFormulaModel GetSettingValueV2(ctx, environmentId, settingId).Execute()

Get value



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/configcat/configcat-publicapi-go-client/v2"
)

func main() {
	environmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The identifier of the Environment.
	settingId := int32(56) // int32 | The id of the Setting.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FeatureFlagSettingValuesV2API.GetSettingValueV2(context.Background(), environmentId, settingId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagSettingValuesV2API.GetSettingValueV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSettingValueV2`: SettingFormulaModel
	fmt.Fprintf(os.Stdout, "Response from `FeatureFlagSettingValuesV2API.GetSettingValueV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | The identifier of the Environment. | 
**settingId** | **int32** | The id of the Setting. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSettingValueV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## GetSettingValuesV2

> ConfigSettingFormulasModel GetSettingValuesV2(ctx, configId, environmentId).Execute()

Get values



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/configcat/configcat-publicapi-go-client/v2"
)

func main() {
	configId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The identifier of the Config.
	environmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The identifier of the Environment.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FeatureFlagSettingValuesV2API.GetSettingValuesV2(context.Background(), configId, environmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagSettingValuesV2API.GetSettingValuesV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSettingValuesV2`: ConfigSettingFormulasModel
	fmt.Fprintf(os.Stdout, "Response from `FeatureFlagSettingValuesV2API.GetSettingValuesV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configId** | **string** | The identifier of the Config. | 
**environmentId** | **string** | The identifier of the Environment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSettingValuesV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ConfigSettingFormulasModel**](ConfigSettingFormulasModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSettingValuesV2

> ConfigSettingFormulasModel PostSettingValuesV2(ctx, configId, environmentId).UpdateEvaluationFormulasModel(updateEvaluationFormulasModel).Reason(reason).Execute()

Post values



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/configcat/configcat-publicapi-go-client/v2"
)

func main() {
	configId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The identifier of the Config.
	environmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The identifier of the Environment.
	updateEvaluationFormulasModel := *openapiclient.NewUpdateEvaluationFormulasModel() // UpdateEvaluationFormulasModel | 
	reason := "reason_example" // string | The reason note for the Audit Log if the Product's \"Config changes require a reason\" preference is turned on. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FeatureFlagSettingValuesV2API.PostSettingValuesV2(context.Background(), configId, environmentId).UpdateEvaluationFormulasModel(updateEvaluationFormulasModel).Reason(reason).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagSettingValuesV2API.PostSettingValuesV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSettingValuesV2`: ConfigSettingFormulasModel
	fmt.Fprintf(os.Stdout, "Response from `FeatureFlagSettingValuesV2API.PostSettingValuesV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configId** | **string** | The identifier of the Config. | 
**environmentId** | **string** | The identifier of the Environment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSettingValuesV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateEvaluationFormulasModel** | [**UpdateEvaluationFormulasModel**](UpdateEvaluationFormulasModel.md) |  | 
 **reason** | **string** | The reason note for the Audit Log if the Product&#39;s \&quot;Config changes require a reason\&quot; preference is turned on. | 

### Return type

[**ConfigSettingFormulasModel**](ConfigSettingFormulasModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceSettingValueV2

> SettingFormulaModel ReplaceSettingValueV2(ctx, environmentId, settingId).UpdateEvaluationFormulaModel(updateEvaluationFormulaModel).Reason(reason).Execute()

Replace value



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/configcat/configcat-publicapi-go-client/v2"
)

func main() {
	environmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The identifier of the Environment.
	settingId := int32(56) // int32 | The id of the Setting.
	updateEvaluationFormulaModel := *openapiclient.NewUpdateEvaluationFormulaModel(*openapiclient.NewValueModel()) // UpdateEvaluationFormulaModel | 
	reason := "reason_example" // string | The reason note for the Audit Log if the Product's \"Config changes require a reason\" preference is turned on. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FeatureFlagSettingValuesV2API.ReplaceSettingValueV2(context.Background(), environmentId, settingId).UpdateEvaluationFormulaModel(updateEvaluationFormulaModel).Reason(reason).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagSettingValuesV2API.ReplaceSettingValueV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceSettingValueV2`: SettingFormulaModel
	fmt.Fprintf(os.Stdout, "Response from `FeatureFlagSettingValuesV2API.ReplaceSettingValueV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | The identifier of the Environment. | 
**settingId** | **int32** | The id of the Setting. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceSettingValueV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateEvaluationFormulaModel** | [**UpdateEvaluationFormulaModel**](UpdateEvaluationFormulaModel.md) |  | 
 **reason** | **string** | The reason note for the Audit Log if the Product&#39;s \&quot;Config changes require a reason\&quot; preference is turned on. | 

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


## UpdateSettingValueV2

> SettingFormulaModel UpdateSettingValueV2(ctx, environmentId, settingId).JsonPatchOperation(jsonPatchOperation).Reason(reason).Execute()

Update value



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/configcat/configcat-publicapi-go-client/v2"
)

func main() {
	environmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The identifier of the Environment.
	settingId := int32(56) // int32 | The id of the Setting.
	jsonPatchOperation := []openapiclient.JsonPatchOperation{*openapiclient.NewJsonPatchOperation(openapiclient.OperationType("unknown"), "Path_example")} // []JsonPatchOperation | 
	reason := "reason_example" // string | The reason note for the Audit Log if the Product's \"Config changes require a reason\" preference is turned on. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FeatureFlagSettingValuesV2API.UpdateSettingValueV2(context.Background(), environmentId, settingId).JsonPatchOperation(jsonPatchOperation).Reason(reason).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagSettingValuesV2API.UpdateSettingValueV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSettingValueV2`: SettingFormulaModel
	fmt.Fprintf(os.Stdout, "Response from `FeatureFlagSettingValuesV2API.UpdateSettingValueV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | The identifier of the Environment. | 
**settingId** | **int32** | The id of the Setting. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSettingValueV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **jsonPatchOperation** | [**[]JsonPatchOperation**](JsonPatchOperation.md) |  | 
 **reason** | **string** | The reason note for the Audit Log if the Product&#39;s \&quot;Config changes require a reason\&quot; preference is turned on. | 

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

