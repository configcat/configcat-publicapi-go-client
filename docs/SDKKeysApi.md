# \SDKKeysAPI

All URIs are relative to *https://test-api.configcat.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSdkKeys**](SDKKeysAPI.md#GetSdkKeys) | **Get** /v1/configs/{configId}/environments/{environmentId} | Get SDK Key



## GetSdkKeys

> SdkKeysModel GetSdkKeys(ctx, configId, environmentId).Execute()

Get SDK Key



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
	resp, r, err := apiClient.SDKKeysAPI.GetSdkKeys(context.Background(), configId, environmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SDKKeysAPI.GetSdkKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSdkKeys`: SdkKeysModel
	fmt.Fprintf(os.Stdout, "Response from `SDKKeysAPI.GetSdkKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configId** | **string** | The identifier of the Config. | 
**environmentId** | **string** | The identifier of the Environment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSdkKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SdkKeysModel**](SdkKeysModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

