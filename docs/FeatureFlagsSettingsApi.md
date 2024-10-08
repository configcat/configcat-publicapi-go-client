# \FeatureFlagsSettingsAPI

All URIs are relative to *https://test-api.configcat.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSetting**](FeatureFlagsSettingsAPI.md#CreateSetting) | **Post** /v1/configs/{configId}/settings | Create Flag
[**DeleteSetting**](FeatureFlagsSettingsAPI.md#DeleteSetting) | **Delete** /v1/settings/{settingId} | Delete Flag
[**GetSetting**](FeatureFlagsSettingsAPI.md#GetSetting) | **Get** /v1/settings/{settingId} | Get Flag
[**GetSettings**](FeatureFlagsSettingsAPI.md#GetSettings) | **Get** /v1/configs/{configId}/settings | List Flags
[**ReplaceSetting**](FeatureFlagsSettingsAPI.md#ReplaceSetting) | **Put** /v1/settings/{settingId} | Replace Flag
[**UpdateSetting**](FeatureFlagsSettingsAPI.md#UpdateSetting) | **Patch** /v1/settings/{settingId} | Update Flag



## CreateSetting

> SettingModel CreateSetting(ctx, configId).CreateSettingInitialValues(createSettingInitialValues).Execute()

Create Flag



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
	createSettingInitialValues := *openapiclient.NewCreateSettingInitialValues("Key_example", "Name_example", openapiclient.SettingType("boolean")) // CreateSettingInitialValues | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FeatureFlagsSettingsAPI.CreateSetting(context.Background(), configId).CreateSettingInitialValues(createSettingInitialValues).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagsSettingsAPI.CreateSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSetting`: SettingModel
	fmt.Fprintf(os.Stdout, "Response from `FeatureFlagsSettingsAPI.CreateSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configId** | **string** | The identifier of the Config. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createSettingInitialValues** | [**CreateSettingInitialValues**](CreateSettingInitialValues.md) |  | 

### Return type

[**SettingModel**](SettingModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSetting

> DeleteSetting(ctx, settingId).Execute()

Delete Flag



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
	settingId := int32(56) // int32 | The identifier of the Setting.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FeatureFlagsSettingsAPI.DeleteSetting(context.Background(), settingId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagsSettingsAPI.DeleteSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**settingId** | **int32** | The identifier of the Setting. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSetting

> SettingModel GetSetting(ctx, settingId).Execute()

Get Flag



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
	settingId := int32(56) // int32 | The identifier of the Setting.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FeatureFlagsSettingsAPI.GetSetting(context.Background(), settingId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagsSettingsAPI.GetSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSetting`: SettingModel
	fmt.Fprintf(os.Stdout, "Response from `FeatureFlagsSettingsAPI.GetSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**settingId** | **int32** | The identifier of the Setting. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SettingModel**](SettingModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSettings

> []SettingModel GetSettings(ctx, configId).Execute()

List Flags



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FeatureFlagsSettingsAPI.GetSettings(context.Background(), configId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagsSettingsAPI.GetSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSettings`: []SettingModel
	fmt.Fprintf(os.Stdout, "Response from `FeatureFlagsSettingsAPI.GetSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configId** | **string** | The identifier of the Config. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]SettingModel**](SettingModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceSetting

> SettingModel ReplaceSetting(ctx, settingId).ReplaceSettingModel(replaceSettingModel).Execute()

Replace Flag



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
	settingId := int32(56) // int32 | The identifier of the Setting.
	replaceSettingModel := *openapiclient.NewReplaceSettingModel("Name_example") // ReplaceSettingModel | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FeatureFlagsSettingsAPI.ReplaceSetting(context.Background(), settingId).ReplaceSettingModel(replaceSettingModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagsSettingsAPI.ReplaceSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceSetting`: SettingModel
	fmt.Fprintf(os.Stdout, "Response from `FeatureFlagsSettingsAPI.ReplaceSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**settingId** | **int32** | The identifier of the Setting. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **replaceSettingModel** | [**ReplaceSettingModel**](ReplaceSettingModel.md) |  | 

### Return type

[**SettingModel**](SettingModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSetting

> SettingModel UpdateSetting(ctx, settingId).JsonPatchOperation(jsonPatchOperation).Execute()

Update Flag



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
	settingId := int32(56) // int32 | The identifier of the Setting.
	jsonPatchOperation := []openapiclient.JsonPatchOperation{*openapiclient.NewJsonPatchOperation(openapiclient.OperationType("unknown"), "Path_example")} // []JsonPatchOperation | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FeatureFlagsSettingsAPI.UpdateSetting(context.Background(), settingId).JsonPatchOperation(jsonPatchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagsSettingsAPI.UpdateSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSetting`: SettingModel
	fmt.Fprintf(os.Stdout, "Response from `FeatureFlagsSettingsAPI.UpdateSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**settingId** | **int32** | The identifier of the Setting. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jsonPatchOperation** | [**[]JsonPatchOperation**](JsonPatchOperation.md) |  | 

### Return type

[**SettingModel**](SettingModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

