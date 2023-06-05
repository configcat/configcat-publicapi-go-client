# \OrganizationsApi

All URIs are relative to *https://api.configcat.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrganizations**](OrganizationsApi.md#GetOrganizations) | **Get** /v1/organizations | List Organizations



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
    openapiclient "github.com/configcat/configcat-publicapi-go-client"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsApi.GetOrganizations(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetOrganizations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizations`: []OrganizationModel
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetOrganizations`: %v\n", resp)
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
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

