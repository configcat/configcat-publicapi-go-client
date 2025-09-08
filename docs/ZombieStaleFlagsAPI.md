# \ZombieStaleFlagsAPI

All URIs are relative to *https://api.configcat.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetStaleflags**](ZombieStaleFlagsAPI.md#GetStaleflags) | **Get** /v1/products/{productId}/staleflags | List Zombie (stale) flags for Product



## GetStaleflags

> StaleFlagProductModel GetStaleflags(ctx, productId).Scope(scope).StaleFlagAgeDays(staleFlagAgeDays).StaleFlagStaleInEnvironmentsType(staleFlagStaleInEnvironmentsType).IgnoredEnvironmentIds(ignoredEnvironmentIds).IgnoredTagIds(ignoredTagIds).Execute()

List Zombie (stale) flags for Product



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
	productId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The identifier of the Product.
	scope := openapiclient.StaleFlagReminderScope("all") // StaleFlagReminderScope | The scope of the report. (optional)
	staleFlagAgeDays := int32(56) // int32 | The inactivity in days after a feature flag should be considered stale. (optional)
	staleFlagStaleInEnvironmentsType := openapiclient.StaleFlagStaleInEnvironmentsType("staleInAnyEnvironments") // StaleFlagStaleInEnvironmentsType | Consider a feature flag as stale if the feature flag is stale in all/any of the environments. (optional)
	ignoredEnvironmentIds := []string{"Inner_example"} // []string | Ignore environment identifiers from the report. (optional)
	ignoredTagIds := []int64{int64(123)} // []int64 | Ignore feature flags from the report based on their tag identifiers. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ZombieStaleFlagsAPI.GetStaleflags(context.Background(), productId).Scope(scope).StaleFlagAgeDays(staleFlagAgeDays).StaleFlagStaleInEnvironmentsType(staleFlagStaleInEnvironmentsType).IgnoredEnvironmentIds(ignoredEnvironmentIds).IgnoredTagIds(ignoredTagIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ZombieStaleFlagsAPI.GetStaleflags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStaleflags`: StaleFlagProductModel
	fmt.Fprintf(os.Stdout, "Response from `ZombieStaleFlagsAPI.GetStaleflags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** | The identifier of the Product. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStaleflagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **scope** | [**StaleFlagReminderScope**](StaleFlagReminderScope.md) | The scope of the report. | 
 **staleFlagAgeDays** | **int32** | The inactivity in days after a feature flag should be considered stale. | 
 **staleFlagStaleInEnvironmentsType** | [**StaleFlagStaleInEnvironmentsType**](StaleFlagStaleInEnvironmentsType.md) | Consider a feature flag as stale if the feature flag is stale in all/any of the environments. | 
 **ignoredEnvironmentIds** | **[]string** | Ignore environment identifiers from the report. | 
 **ignoredTagIds** | **[]int64** | Ignore feature flags from the report based on their tag identifiers. | 

### Return type

[**StaleFlagProductModel**](StaleFlagProductModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

