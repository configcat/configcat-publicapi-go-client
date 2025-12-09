# \OrganizationsAPI

All URIs are relative to *https://api.configcat.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrganizationLimitations**](OrganizationsAPI.md#GetOrganizationLimitations) | **Get** /v1/organizations/{organizationId}/organization-limitations | Get Organization limitations
[**GetOrganizations**](OrganizationsAPI.md#GetOrganizations) | **Get** /v1/organizations | List Organizations



## GetOrganizationLimitations

> OrganizationLimitations GetOrganizationLimitations(ctx, organizationId).Execute()

Get Organization limitations



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.GetOrganizationLimitations(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.GetOrganizationLimitations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationLimitations`: OrganizationLimitations
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.GetOrganizationLimitations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | The identifier of the Organization. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationLimitationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OrganizationLimitations**](OrganizationLimitations.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizations

> []OrganizationModel GetOrganizations(ctx).Execute()

List Organizations



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.GetOrganizations(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.GetOrganizations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizations`: []OrganizationModel
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.GetOrganizations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationsRequest struct via the builder pattern


### Return type

[**[]OrganizationModel**](OrganizationModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

