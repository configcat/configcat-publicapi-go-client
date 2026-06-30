# \AuditLogsAPI

All URIs are relative to *https://api.configcat.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAuditlogs**](AuditLogsAPI.md#GetAuditlogs) | **Get** /v1/products/{productId}/auditlogs | List Audit log items for Product
[**GetAuditlogsV2**](AuditLogsAPI.md#GetAuditlogsV2) | **Get** /v2/products/{productId}/auditlogs | List Audit log items for Product (V2)
[**GetDeletedSettings**](AuditLogsAPI.md#GetDeletedSettings) | **Get** /v1/configs/{configId}/deleted-settings | List Deleted Settings
[**GetOrganizationAuditlogs**](AuditLogsAPI.md#GetOrganizationAuditlogs) | **Get** /v1/organizations/{organizationId}/auditlogs | List Audit log items for Organization
[**GetOrganizationAuditlogsV2**](AuditLogsAPI.md#GetOrganizationAuditlogsV2) | **Get** /v2/organizations/{organizationId}/auditlogs | List Audit log items for Organization (V2)



## GetAuditlogs

> []AuditLogItemModel GetAuditlogs(ctx, productId).ConfigId(configId).EnvironmentId(environmentId).AuditLogType(auditLogType).FromUtcDateTime(fromUtcDateTime).ToUtcDateTime(toUtcDateTime).Execute()

List Audit log items for Product



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/configcat/configcat-publicapi-go-client/v3"
)

func main() {
	productId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The identifier of the Product.
	configId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The identifier of the Config. (optional)
	environmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The identifier of the Environment. (optional)
	auditLogType := openapiclient.AuditLogType("productCreated") // AuditLogType | Filter Audit logs by Audit log type. (optional)
	fromUtcDateTime := time.Now() // time.Time | Filter Audit logs by starting UTC date. (optional)
	toUtcDateTime := time.Now() // time.Time | Filter Audit logs by ending UTC date. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuditLogsAPI.GetAuditlogs(context.Background(), productId).ConfigId(configId).EnvironmentId(environmentId).AuditLogType(auditLogType).FromUtcDateTime(fromUtcDateTime).ToUtcDateTime(toUtcDateTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditLogsAPI.GetAuditlogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAuditlogs`: []AuditLogItemModel
	fmt.Fprintf(os.Stdout, "Response from `AuditLogsAPI.GetAuditlogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** | The identifier of the Product. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuditlogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **configId** | **string** | The identifier of the Config. | 
 **environmentId** | **string** | The identifier of the Environment. | 
 **auditLogType** | [**AuditLogType**](AuditLogType.md) | Filter Audit logs by Audit log type. | 
 **fromUtcDateTime** | **time.Time** | Filter Audit logs by starting UTC date. | 
 **toUtcDateTime** | **time.Time** | Filter Audit logs by ending UTC date. | 

### Return type

[**[]AuditLogItemModel**](AuditLogItemModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuditlogsV2

> AuditLogItemModelPagedList GetAuditlogsV2(ctx, productId).ConfigId(configId).EnvironmentId(environmentId).AuditLogType(auditLogType).FromUtcDateTime(fromUtcDateTime).ToUtcDateTime(toUtcDateTime).PageNumber(pageNumber).PageSize(pageSize).Execute()

List Audit log items for Product (V2)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/configcat/configcat-publicapi-go-client/v3"
)

func main() {
	productId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The identifier of the Product.
	configId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The identifier of the Config. (optional)
	environmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The identifier of the Environment. (optional)
	auditLogType := openapiclient.AuditLogType("productCreated") // AuditLogType | Filter Audit logs by Audit log type. (optional)
	fromUtcDateTime := time.Now() // time.Time | Filter Audit logs by starting UTC date. (optional)
	toUtcDateTime := time.Now() // time.Time | Filter Audit logs by ending UTC date. (optional)
	pageNumber := int32(56) // int32 | Page number (min: 1). (optional) (default to 1)
	pageSize := int32(56) // int32 | Page size (min: 1, max: 100). (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuditLogsAPI.GetAuditlogsV2(context.Background(), productId).ConfigId(configId).EnvironmentId(environmentId).AuditLogType(auditLogType).FromUtcDateTime(fromUtcDateTime).ToUtcDateTime(toUtcDateTime).PageNumber(pageNumber).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditLogsAPI.GetAuditlogsV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAuditlogsV2`: AuditLogItemModelPagedList
	fmt.Fprintf(os.Stdout, "Response from `AuditLogsAPI.GetAuditlogsV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** | The identifier of the Product. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuditlogsV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **configId** | **string** | The identifier of the Config. | 
 **environmentId** | **string** | The identifier of the Environment. | 
 **auditLogType** | [**AuditLogType**](AuditLogType.md) | Filter Audit logs by Audit log type. | 
 **fromUtcDateTime** | **time.Time** | Filter Audit logs by starting UTC date. | 
 **toUtcDateTime** | **time.Time** | Filter Audit logs by ending UTC date. | 
 **pageNumber** | **int32** | Page number (min: 1). | [default to 1]
 **pageSize** | **int32** | Page size (min: 1, max: 100). | [default to 100]

### Return type

[**AuditLogItemModelPagedList**](AuditLogItemModelPagedList.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeletedSettings

> []DeletedSettingModel GetDeletedSettings(ctx, configId).Execute()

List Deleted Settings



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
	configId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The identifier of the Config.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuditLogsAPI.GetDeletedSettings(context.Background(), configId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditLogsAPI.GetDeletedSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeletedSettings`: []DeletedSettingModel
	fmt.Fprintf(os.Stdout, "Response from `AuditLogsAPI.GetDeletedSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configId** | **string** | The identifier of the Config. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeletedSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]DeletedSettingModel**](DeletedSettingModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationAuditlogs

> []AuditLogItemModel GetOrganizationAuditlogs(ctx, organizationId).ProductId(productId).ConfigId(configId).EnvironmentId(environmentId).AuditLogType(auditLogType).FromUtcDateTime(fromUtcDateTime).ToUtcDateTime(toUtcDateTime).Execute()

List Audit log items for Organization



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/configcat/configcat-publicapi-go-client/v3"
)

func main() {
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The identifier of the Organization.
	productId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The identifier of the Product. (optional)
	configId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The identifier of the Config. (optional)
	environmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The identifier of the Environment. (optional)
	auditLogType := openapiclient.AuditLogType("productCreated") // AuditLogType | Filter Audit logs by Audit log type. (optional)
	fromUtcDateTime := time.Now() // time.Time | Filter Audit logs by starting UTC date. (optional)
	toUtcDateTime := time.Now() // time.Time | Filter Audit logs by ending UTC date. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuditLogsAPI.GetOrganizationAuditlogs(context.Background(), organizationId).ProductId(productId).ConfigId(configId).EnvironmentId(environmentId).AuditLogType(auditLogType).FromUtcDateTime(fromUtcDateTime).ToUtcDateTime(toUtcDateTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditLogsAPI.GetOrganizationAuditlogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationAuditlogs`: []AuditLogItemModel
	fmt.Fprintf(os.Stdout, "Response from `AuditLogsAPI.GetOrganizationAuditlogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | The identifier of the Organization. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationAuditlogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **productId** | **string** | The identifier of the Product. | 
 **configId** | **string** | The identifier of the Config. | 
 **environmentId** | **string** | The identifier of the Environment. | 
 **auditLogType** | [**AuditLogType**](AuditLogType.md) | Filter Audit logs by Audit log type. | 
 **fromUtcDateTime** | **time.Time** | Filter Audit logs by starting UTC date. | 
 **toUtcDateTime** | **time.Time** | Filter Audit logs by ending UTC date. | 

### Return type

[**[]AuditLogItemModel**](AuditLogItemModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationAuditlogsV2

> AuditLogItemModelPagedList GetOrganizationAuditlogsV2(ctx, organizationId).ProductId(productId).ConfigId(configId).EnvironmentId(environmentId).AuditLogType(auditLogType).FromUtcDateTime(fromUtcDateTime).ToUtcDateTime(toUtcDateTime).PageNumber(pageNumber).PageSize(pageSize).Execute()

List Audit log items for Organization (V2)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/configcat/configcat-publicapi-go-client/v3"
)

func main() {
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The identifier of the Organization.
	productId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The identifier of the Product. (optional)
	configId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The identifier of the Config. (optional)
	environmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The identifier of the Environment. (optional)
	auditLogType := openapiclient.AuditLogType("productCreated") // AuditLogType | Filter Audit logs by Audit log type. (optional)
	fromUtcDateTime := time.Now() // time.Time | Filter Audit logs by starting UTC date. (optional)
	toUtcDateTime := time.Now() // time.Time | Filter Audit logs by ending UTC date. (optional)
	pageNumber := int32(56) // int32 | Page number (min: 1). (optional) (default to 1)
	pageSize := int32(56) // int32 | Page size (min: 1, max: 100). (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuditLogsAPI.GetOrganizationAuditlogsV2(context.Background(), organizationId).ProductId(productId).ConfigId(configId).EnvironmentId(environmentId).AuditLogType(auditLogType).FromUtcDateTime(fromUtcDateTime).ToUtcDateTime(toUtcDateTime).PageNumber(pageNumber).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditLogsAPI.GetOrganizationAuditlogsV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationAuditlogsV2`: AuditLogItemModelPagedList
	fmt.Fprintf(os.Stdout, "Response from `AuditLogsAPI.GetOrganizationAuditlogsV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | The identifier of the Organization. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationAuditlogsV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **productId** | **string** | The identifier of the Product. | 
 **configId** | **string** | The identifier of the Config. | 
 **environmentId** | **string** | The identifier of the Environment. | 
 **auditLogType** | [**AuditLogType**](AuditLogType.md) | Filter Audit logs by Audit log type. | 
 **fromUtcDateTime** | **time.Time** | Filter Audit logs by starting UTC date. | 
 **toUtcDateTime** | **time.Time** | Filter Audit logs by ending UTC date. | 
 **pageNumber** | **int32** | Page number (min: 1). | [default to 1]
 **pageSize** | **int32** | Page size (min: 1, max: 100). | [default to 100]

### Return type

[**AuditLogItemModelPagedList**](AuditLogItemModelPagedList.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

