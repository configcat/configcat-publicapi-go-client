# \IntegrationLinksAPI

All URIs are relative to *https://test-api.configcat.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddOrUpdateIntegrationLink**](IntegrationLinksAPI.md#AddOrUpdateIntegrationLink) | **Post** /v1/environments/{environmentId}/settings/{settingId}/integrationLinks/{integrationLinkType}/{key} | Add or update Integration link
[**DeleteIntegrationLink**](IntegrationLinksAPI.md#DeleteIntegrationLink) | **Delete** /v1/environments/{environmentId}/settings/{settingId}/integrationLinks/{integrationLinkType}/{key} | Delete Integration link
[**GetIntegrationLinkDetails**](IntegrationLinksAPI.md#GetIntegrationLinkDetails) | **Get** /v1/integrationLink/{integrationLinkType}/{key}/details | Get Integration link
[**JiraAddOrUpdateIntegrationLink**](IntegrationLinksAPI.md#JiraAddOrUpdateIntegrationLink) | **Post** /v1/jira/environments/{environmentId}/settings/{settingId}/integrationLinks/{key} | 
[**JiraConnect**](IntegrationLinksAPI.md#JiraConnect) | **Post** /v1/jira/connect | 



## AddOrUpdateIntegrationLink

> IntegrationLinkModel AddOrUpdateIntegrationLink(ctx, environmentId, settingId, integrationLinkType, key).AddOrUpdateIntegrationLinkModel(addOrUpdateIntegrationLinkModel).Execute()

Add or update Integration link



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
	integrationLinkType := openapiclient.IntegrationLinkType("trello") // IntegrationLinkType | The integration link's type.
	key := "key_example" // string | The key of the integration link.
	addOrUpdateIntegrationLinkModel := *openapiclient.NewAddOrUpdateIntegrationLinkModel() // AddOrUpdateIntegrationLinkModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationLinksAPI.AddOrUpdateIntegrationLink(context.Background(), environmentId, settingId, integrationLinkType, key).AddOrUpdateIntegrationLinkModel(addOrUpdateIntegrationLinkModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationLinksAPI.AddOrUpdateIntegrationLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddOrUpdateIntegrationLink`: IntegrationLinkModel
	fmt.Fprintf(os.Stdout, "Response from `IntegrationLinksAPI.AddOrUpdateIntegrationLink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | The identifier of the Environment. | 
**settingId** | **int32** | The id of the Setting. | 
**integrationLinkType** | [**IntegrationLinkType**](.md) | The integration link&#39;s type. | 
**key** | **string** | The key of the integration link. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddOrUpdateIntegrationLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **addOrUpdateIntegrationLinkModel** | [**AddOrUpdateIntegrationLinkModel**](AddOrUpdateIntegrationLinkModel.md) |  | 

### Return type

[**IntegrationLinkModel**](IntegrationLinkModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIntegrationLink

> DeleteIntegrationLinkModel DeleteIntegrationLink(ctx, environmentId, settingId, integrationLinkType, key).Execute()

Delete Integration link



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
	integrationLinkType := openapiclient.IntegrationLinkType("trello") // IntegrationLinkType | The integration's type.
	key := "key_example" // string | The key of the integration link.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationLinksAPI.DeleteIntegrationLink(context.Background(), environmentId, settingId, integrationLinkType, key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationLinksAPI.DeleteIntegrationLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteIntegrationLink`: DeleteIntegrationLinkModel
	fmt.Fprintf(os.Stdout, "Response from `IntegrationLinksAPI.DeleteIntegrationLink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | The identifier of the Environment. | 
**settingId** | **int32** | The id of the Setting. | 
**integrationLinkType** | [**IntegrationLinkType**](.md) | The integration&#39;s type. | 
**key** | **string** | The key of the integration link. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIntegrationLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**DeleteIntegrationLinkModel**](DeleteIntegrationLinkModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIntegrationLinkDetails

> IntegrationLinkDetailsModel GetIntegrationLinkDetails(ctx, integrationLinkType, key).Execute()

Get Integration link



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
	integrationLinkType := openapiclient.IntegrationLinkType("trello") // IntegrationLinkType | The integration link's type.
	key := "key_example" // string | The key of the integration link.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationLinksAPI.GetIntegrationLinkDetails(context.Background(), integrationLinkType, key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationLinksAPI.GetIntegrationLinkDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIntegrationLinkDetails`: IntegrationLinkDetailsModel
	fmt.Fprintf(os.Stdout, "Response from `IntegrationLinksAPI.GetIntegrationLinkDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationLinkType** | [**IntegrationLinkType**](.md) | The integration link&#39;s type. | 
**key** | **string** | The key of the integration link. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationLinkDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**IntegrationLinkDetailsModel**](IntegrationLinkDetailsModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JiraAddOrUpdateIntegrationLink

> IntegrationLinkModel JiraAddOrUpdateIntegrationLink(ctx, environmentId, settingId, key).AddOrUpdateJiraIntegrationLinkModel(addOrUpdateJiraIntegrationLinkModel).Execute()



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
	key := "key_example" // string | The key of the integration link.
	addOrUpdateJiraIntegrationLinkModel := *openapiclient.NewAddOrUpdateJiraIntegrationLinkModel("JiraJwtToken_example", "ClientKey_example") // AddOrUpdateJiraIntegrationLinkModel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationLinksAPI.JiraAddOrUpdateIntegrationLink(context.Background(), environmentId, settingId, key).AddOrUpdateJiraIntegrationLinkModel(addOrUpdateJiraIntegrationLinkModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationLinksAPI.JiraAddOrUpdateIntegrationLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `JiraAddOrUpdateIntegrationLink`: IntegrationLinkModel
	fmt.Fprintf(os.Stdout, "Response from `IntegrationLinksAPI.JiraAddOrUpdateIntegrationLink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | The identifier of the Environment. | 
**settingId** | **int32** | The id of the Setting. | 
**key** | **string** | The key of the integration link. | 

### Other Parameters

Other parameters are passed through a pointer to a apiJiraAddOrUpdateIntegrationLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **addOrUpdateJiraIntegrationLinkModel** | [**AddOrUpdateJiraIntegrationLinkModel**](AddOrUpdateJiraIntegrationLinkModel.md) |  | 

### Return type

[**IntegrationLinkModel**](IntegrationLinkModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JiraConnect

> JiraConnect(ctx).ConnectRequest(connectRequest).Execute()



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
	connectRequest := *openapiclient.NewConnectRequest("ClientKey_example", "JiraJwtToken_example") // ConnectRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IntegrationLinksAPI.JiraConnect(context.Background()).ConnectRequest(connectRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationLinksAPI.JiraConnect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiJiraConnectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **connectRequest** | [**ConnectRequest**](ConnectRequest.md) |  | 

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

