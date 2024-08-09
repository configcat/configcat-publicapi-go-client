# \PermissionGroupsAPI

All URIs are relative to *https://test-api.configcat.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePermissionGroup**](PermissionGroupsAPI.md#CreatePermissionGroup) | **Post** /v1/products/{productId}/permissions | Create Permission Group
[**DeletePermissionGroup**](PermissionGroupsAPI.md#DeletePermissionGroup) | **Delete** /v1/permissions/{permissionGroupId} | Delete Permission Group
[**GetPermissionGroup**](PermissionGroupsAPI.md#GetPermissionGroup) | **Get** /v1/permissions/{permissionGroupId} | Get Permission Group
[**GetPermissionGroups**](PermissionGroupsAPI.md#GetPermissionGroups) | **Get** /v1/products/{productId}/permissions | List Permission Groups
[**UpdatePermissionGroup**](PermissionGroupsAPI.md#UpdatePermissionGroup) | **Put** /v1/permissions/{permissionGroupId} | Update Permission Group



## CreatePermissionGroup

> PermissionGroupModel CreatePermissionGroup(ctx, productId).CreatePermissionGroupRequest(createPermissionGroupRequest).Execute()

Create Permission Group



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
	productId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The identifier of the Product.
	createPermissionGroupRequest := *openapiclient.NewCreatePermissionGroupRequest("Name_example") // CreatePermissionGroupRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PermissionGroupsAPI.CreatePermissionGroup(context.Background(), productId).CreatePermissionGroupRequest(createPermissionGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionGroupsAPI.CreatePermissionGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePermissionGroup`: PermissionGroupModel
	fmt.Fprintf(os.Stdout, "Response from `PermissionGroupsAPI.CreatePermissionGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** | The identifier of the Product. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePermissionGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createPermissionGroupRequest** | [**CreatePermissionGroupRequest**](CreatePermissionGroupRequest.md) |  | 

### Return type

[**PermissionGroupModel**](PermissionGroupModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePermissionGroup

> DeletePermissionGroup(ctx, permissionGroupId).Execute()

Delete Permission Group



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
	permissionGroupId := int64(789) // int64 | The identifier of the Permission Group.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PermissionGroupsAPI.DeletePermissionGroup(context.Background(), permissionGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionGroupsAPI.DeletePermissionGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**permissionGroupId** | **int64** | The identifier of the Permission Group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePermissionGroupRequest struct via the builder pattern


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


## GetPermissionGroup

> PermissionGroupModel GetPermissionGroup(ctx, permissionGroupId).Execute()

Get Permission Group



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
	permissionGroupId := int64(789) // int64 | The identifier of the Permission Group.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PermissionGroupsAPI.GetPermissionGroup(context.Background(), permissionGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionGroupsAPI.GetPermissionGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPermissionGroup`: PermissionGroupModel
	fmt.Fprintf(os.Stdout, "Response from `PermissionGroupsAPI.GetPermissionGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**permissionGroupId** | **int64** | The identifier of the Permission Group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPermissionGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PermissionGroupModel**](PermissionGroupModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPermissionGroups

> []PermissionGroupModel GetPermissionGroups(ctx, productId).Execute()

List Permission Groups



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
	productId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The identifier of the Product.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PermissionGroupsAPI.GetPermissionGroups(context.Background(), productId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionGroupsAPI.GetPermissionGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPermissionGroups`: []PermissionGroupModel
	fmt.Fprintf(os.Stdout, "Response from `PermissionGroupsAPI.GetPermissionGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** | The identifier of the Product. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPermissionGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]PermissionGroupModel**](PermissionGroupModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePermissionGroup

> PermissionGroupModel UpdatePermissionGroup(ctx, permissionGroupId).UpdatePermissionGroupRequest(updatePermissionGroupRequest).Execute()

Update Permission Group



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
	permissionGroupId := int64(789) // int64 | The identifier of the Permission Group.
	updatePermissionGroupRequest := *openapiclient.NewUpdatePermissionGroupRequest() // UpdatePermissionGroupRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PermissionGroupsAPI.UpdatePermissionGroup(context.Background(), permissionGroupId).UpdatePermissionGroupRequest(updatePermissionGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionGroupsAPI.UpdatePermissionGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePermissionGroup`: PermissionGroupModel
	fmt.Fprintf(os.Stdout, "Response from `PermissionGroupsAPI.UpdatePermissionGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**permissionGroupId** | **int64** | The identifier of the Permission Group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePermissionGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatePermissionGroupRequest** | [**UpdatePermissionGroupRequest**](UpdatePermissionGroupRequest.md) |  | 

### Return type

[**PermissionGroupModel**](PermissionGroupModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

