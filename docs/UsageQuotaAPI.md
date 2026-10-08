# \UsageQuotaAPI

All URIs are relative to *https://api.configcat.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrganizationUsageAndQuota**](UsageQuotaAPI.md#GetOrganizationUsageAndQuota) | **Get** /v1/organizations/{organizationId}/usage-and-quota | Get usage and quota



## GetOrganizationUsageAndQuota

> StatisticsV2Model GetOrganizationUsageAndQuota(ctx, organizationId).ProductId(productId).Execute()

Get usage and quota



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
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The identifier of the Organization.
	productId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The identifier of the Product to filter statistics for. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsageQuotaAPI.GetOrganizationUsageAndQuota(context.Background(), organizationId).ProductId(productId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageQuotaAPI.GetOrganizationUsageAndQuota``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationUsageAndQuota`: StatisticsV2Model
	fmt.Fprintf(os.Stdout, "Response from `UsageQuotaAPI.GetOrganizationUsageAndQuota`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | The identifier of the Organization. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationUsageAndQuotaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **productId** | **string** | The identifier of the Product to filter statistics for. | 

### Return type

[**StatisticsV2Model**](StatisticsV2Model.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

