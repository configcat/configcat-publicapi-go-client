# \MembersApi

All URIs are relative to *https://api.configcat.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddMemberToGroup**](MembersApi.md#AddMemberToGroup) | **Post** /v1/organizations/{organizationId}/members/{userId} | Update Member Permissions
[**DeleteOrganizationMember**](MembersApi.md#DeleteOrganizationMember) | **Delete** /v1/organizations/{organizationId}/members/{userId} | Delete Member from Organization
[**DeleteProductMember**](MembersApi.md#DeleteProductMember) | **Delete** /v1/products/{productId}/members/{userId} | Delete Member from Product
[**GetOrganizationMembers**](MembersApi.md#GetOrganizationMembers) | **Get** /v1/organizations/{organizationId}/members | List Organization Members
[**GetProductMembers**](MembersApi.md#GetProductMembers) | **Get** /v1/products/{productId}/members | List Product Members
[**InviteMember**](MembersApi.md#InviteMember) | **Post** /v1/products/{productId}/members/invite | Invite Member



## AddMemberToGroup

> AddMemberToGroup(ctx, organizationId, userId).AddUserToGroupRequest(addUserToGroupRequest).Execute()

Update Member Permissions



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
    organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The identifier of the Organization.
    userId := "userId_example" // string | The identifier of the Member.
    addUserToGroupRequest := *openapiclient.NewAddUserToGroupRequest([]int64{int64(123)}) // AddUserToGroupRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MembersApi.AddMemberToGroup(context.Background(), organizationId, userId).AddUserToGroupRequest(addUserToGroupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MembersApi.AddMemberToGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | The identifier of the Organization. | 
**userId** | **string** | The identifier of the Member. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddMemberToGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **addUserToGroupRequest** | [**AddUserToGroupRequest**](AddUserToGroupRequest.md) |  | 

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


## DeleteOrganizationMember

> DeleteOrganizationMember(ctx, organizationId, userId).Execute()

Delete Member from Organization



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
    organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The identifier of the Organization.
    userId := "userId_example" // string | The identifier of the Member.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MembersApi.DeleteOrganizationMember(context.Background(), organizationId, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MembersApi.DeleteOrganizationMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | The identifier of the Organization. | 
**userId** | **string** | The identifier of the Member. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationMemberRequest struct via the builder pattern


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


## DeleteProductMember

> DeleteProductMember(ctx, productId, userId).Execute()

Delete Member from Product



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
    userId := "userId_example" // string | The identifier of the Member.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MembersApi.DeleteProductMember(context.Background(), productId, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MembersApi.DeleteProductMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** | The identifier of the Product. | 
**userId** | **string** | The identifier of the Member. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProductMemberRequest struct via the builder pattern


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


## GetOrganizationMembers

> []UserModel GetOrganizationMembers(ctx, organizationId).Execute()

List Organization Members



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
    organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The identifier of the Organization.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MembersApi.GetOrganizationMembers(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MembersApi.GetOrganizationMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationMembers`: []UserModel
    fmt.Fprintf(os.Stdout, "Response from `MembersApi.GetOrganizationMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | The identifier of the Organization. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]UserModel**](UserModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProductMembers

> []MemberModel GetProductMembers(ctx, productId).Execute()

List Product Members



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
    resp, r, err := apiClient.MembersApi.GetProductMembers(context.Background(), productId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MembersApi.GetProductMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProductMembers`: []MemberModel
    fmt.Fprintf(os.Stdout, "Response from `MembersApi.GetProductMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** | The identifier of the Product. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProductMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]MemberModel**](MemberModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InviteMember

> InviteMember(ctx, productId).InviteMembersRequest(inviteMembersRequest).Execute()

Invite Member



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
    inviteMembersRequest := *openapiclient.NewInviteMembersRequest([]string{"Emails_example"}, int64(123)) // InviteMembersRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MembersApi.InviteMember(context.Background(), productId).InviteMembersRequest(inviteMembersRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MembersApi.InviteMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** | The identifier of the Product. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInviteMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inviteMembersRequest** | [**InviteMembersRequest**](InviteMembersRequest.md) |  | 

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

