# \PermissionGroupsApi

All URIs are relative to *https://api.configcat.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePermissionGroup**](PermissionGroupsApi.md#CreatePermissionGroup) | **Post** /v1/products/{productId}/permissions | Create Permission Group
[**DeletePermissionGroup**](PermissionGroupsApi.md#DeletePermissionGroup) | **Delete** /v1/permissions/{permissionGroupId} | Delete Permission Group
[**GetPermissionGroup**](PermissionGroupsApi.md#GetPermissionGroup) | **Get** /v1/permissions/{permissionGroupId} | Get Permission Group
[**GetPermissionGroups**](PermissionGroupsApi.md#GetPermissionGroups) | **Get** /v1/products/{productId}/permissions | List Permission Groups
[**UpdatePermissionGroup**](PermissionGroupsApi.md#UpdatePermissionGroup) | **Put** /v1/permissions/{permissionGroupId} | Update Permission Group



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
    openapiclient "github.com/configcat/configcat-publicapi-go-client"
)

func main() {
    productId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The identifier of the Product.
    createPermissionGroupRequest := *openapiclient.NewCreatePermissionGroupRequest("Name_example") // CreatePermissionGroupRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PermissionGroupsApi.CreatePermissionGroup(context.Background(), productId).CreatePermissionGroupRequest(createPermissionGroupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionGroupsApi.CreatePermissionGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePermissionGroup`: PermissionGroupModel
    fmt.Fprintf(os.Stdout, "Response from `PermissionGroupsApi.CreatePermissionGroup`: %v\n", resp)
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
- **Accept**: application/json, application/hal+json

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
    openapiclient "github.com/configcat/configcat-publicapi-go-client"
)

func main() {
    permissionGroupId := int64(789) // int64 | The identifier of the Permission Group.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PermissionGroupsApi.DeletePermissionGroup(context.Background(), permissionGroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionGroupsApi.DeletePermissionGroup``: %v\n", err)
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
    openapiclient "github.com/configcat/configcat-publicapi-go-client"
)

func main() {
    permissionGroupId := int64(789) // int64 | The identifier of the Permission Group.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PermissionGroupsApi.GetPermissionGroup(context.Background(), permissionGroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionGroupsApi.GetPermissionGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPermissionGroup`: PermissionGroupModel
    fmt.Fprintf(os.Stdout, "Response from `PermissionGroupsApi.GetPermissionGroup`: %v\n", resp)
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
- **Accept**: application/json, application/hal+json

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
    openapiclient "github.com/configcat/configcat-publicapi-go-client"
)

func main() {
    productId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The identifier of the Product.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PermissionGroupsApi.GetPermissionGroups(context.Background(), productId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionGroupsApi.GetPermissionGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPermissionGroups`: []PermissionGroupModel
    fmt.Fprintf(os.Stdout, "Response from `PermissionGroupsApi.GetPermissionGroups`: %v\n", resp)
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
- **Accept**: application/json, application/hal+json

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
    openapiclient "github.com/configcat/configcat-publicapi-go-client"
)

func main() {
    permissionGroupId := int64(789) // int64 | The identifier of the Permission Group.
    updatePermissionGroupRequest := *openapiclient.NewUpdatePermissionGroupRequest() // UpdatePermissionGroupRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PermissionGroupsApi.UpdatePermissionGroup(context.Background(), permissionGroupId).UpdatePermissionGroupRequest(updatePermissionGroupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionGroupsApi.UpdatePermissionGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePermissionGroup`: PermissionGroupModel
    fmt.Fprintf(os.Stdout, "Response from `PermissionGroupsApi.UpdatePermissionGroup`: %v\n", resp)
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
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

