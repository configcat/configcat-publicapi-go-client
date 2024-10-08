# \MembersAPI

All URIs are relative to *https://test-api.configcat.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddMemberToGroup**](MembersAPI.md#AddMemberToGroup) | **Post** /v1/organizations/{organizationId}/members/{userId} | Update Member Permissions
[**DeleteInvitation**](MembersAPI.md#DeleteInvitation) | **Delete** /v1/invitations/{invitationId} | Delete Invitation
[**DeleteOrganizationMember**](MembersAPI.md#DeleteOrganizationMember) | **Delete** /v1/organizations/{organizationId}/members/{userId} | Delete Member from Organization
[**DeleteProductMember**](MembersAPI.md#DeleteProductMember) | **Delete** /v1/products/{productId}/members/{userId} | Delete Member from Product
[**GetOrganizationMembers**](MembersAPI.md#GetOrganizationMembers) | **Get** /v1/organizations/{organizationId}/members | List Organization Members
[**GetOrganizationMembersV2**](MembersAPI.md#GetOrganizationMembersV2) | **Get** /v2/organizations/{organizationId}/members | List Organization Members
[**GetPendingInvitations**](MembersAPI.md#GetPendingInvitations) | **Get** /v1/products/{productId}/invitations | List Pending Invitations in Product
[**GetPendingInvitationsOrg**](MembersAPI.md#GetPendingInvitationsOrg) | **Get** /v1/organizations/{organizationId}/invitations | List Pending Invitations in Organization
[**GetProductMembers**](MembersAPI.md#GetProductMembers) | **Get** /v1/products/{productId}/members | List Product Members
[**InviteMember**](MembersAPI.md#InviteMember) | **Post** /v1/products/{productId}/members/invite | Invite Member



## AddMemberToGroup

> AddMemberToGroup(ctx, organizationId, userId).UpdateMemberPermissionsRequest(updateMemberPermissionsRequest).Execute()

Update Member Permissions



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
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The identifier of the Organization.
	userId := "userId_example" // string | The identifier of the Member.
	updateMemberPermissionsRequest := *openapiclient.NewUpdateMemberPermissionsRequest() // UpdateMemberPermissionsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MembersAPI.AddMemberToGroup(context.Background(), organizationId, userId).UpdateMemberPermissionsRequest(updateMemberPermissionsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MembersAPI.AddMemberToGroup``: %v\n", err)
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


 **updateMemberPermissionsRequest** | [**UpdateMemberPermissionsRequest**](UpdateMemberPermissionsRequest.md) |  | 

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


## DeleteInvitation

> DeleteInvitation(ctx, invitationId).Execute()

Delete Invitation



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
	invitationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The identifier of the Invitation.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MembersAPI.DeleteInvitation(context.Background(), invitationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MembersAPI.DeleteInvitation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invitationId** | **string** | The identifier of the Invitation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInvitationRequest struct via the builder pattern


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
	openapiclient "github.com/configcat/configcat-publicapi-go-client/v2"
)

func main() {
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The identifier of the Organization.
	userId := "userId_example" // string | The identifier of the Member.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MembersAPI.DeleteOrganizationMember(context.Background(), organizationId, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MembersAPI.DeleteOrganizationMember``: %v\n", err)
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
	openapiclient "github.com/configcat/configcat-publicapi-go-client/v2"
)

func main() {
	productId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The identifier of the Product.
	userId := "userId_example" // string | The identifier of the Member.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MembersAPI.DeleteProductMember(context.Background(), productId, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MembersAPI.DeleteProductMember``: %v\n", err)
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
	openapiclient "github.com/configcat/configcat-publicapi-go-client/v2"
)

func main() {
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The identifier of the Organization.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MembersAPI.GetOrganizationMembers(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MembersAPI.GetOrganizationMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationMembers`: []UserModel
	fmt.Fprintf(os.Stdout, "Response from `MembersAPI.GetOrganizationMembers`: %v\n", resp)
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
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationMembersV2

> OrganizationMembersModel GetOrganizationMembersV2(ctx, organizationId).Execute()

List Organization Members



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
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The identifier of the Organization.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MembersAPI.GetOrganizationMembersV2(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MembersAPI.GetOrganizationMembersV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationMembersV2`: OrganizationMembersModel
	fmt.Fprintf(os.Stdout, "Response from `MembersAPI.GetOrganizationMembersV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | The identifier of the Organization. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationMembersV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OrganizationMembersModel**](OrganizationMembersModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPendingInvitations

> []InvitationModel GetPendingInvitations(ctx, productId).Execute()

List Pending Invitations in Product



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
	resp, r, err := apiClient.MembersAPI.GetPendingInvitations(context.Background(), productId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MembersAPI.GetPendingInvitations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPendingInvitations`: []InvitationModel
	fmt.Fprintf(os.Stdout, "Response from `MembersAPI.GetPendingInvitations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** | The identifier of the Product. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPendingInvitationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]InvitationModel**](InvitationModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPendingInvitationsOrg

> []OrganizationInvitationModel GetPendingInvitationsOrg(ctx, organizationId).Execute()

List Pending Invitations in Organization



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
	organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The identifier of the Organization.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MembersAPI.GetPendingInvitationsOrg(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MembersAPI.GetPendingInvitationsOrg``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPendingInvitationsOrg`: []OrganizationInvitationModel
	fmt.Fprintf(os.Stdout, "Response from `MembersAPI.GetPendingInvitationsOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | The identifier of the Organization. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPendingInvitationsOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]OrganizationInvitationModel**](OrganizationInvitationModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

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
	openapiclient "github.com/configcat/configcat-publicapi-go-client/v2"
)

func main() {
	productId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The identifier of the Product.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MembersAPI.GetProductMembers(context.Background(), productId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MembersAPI.GetProductMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProductMembers`: []MemberModel
	fmt.Fprintf(os.Stdout, "Response from `MembersAPI.GetProductMembers`: %v\n", resp)
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
- **Accept**: application/json

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
	openapiclient "github.com/configcat/configcat-publicapi-go-client/v2"
)

func main() {
	productId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The identifier of the Product.
	inviteMembersRequest := *openapiclient.NewInviteMembersRequest([]string{"Emails_example"}, int64(123)) // InviteMembersRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MembersAPI.InviteMember(context.Background(), productId).InviteMembersRequest(inviteMembersRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MembersAPI.InviteMember``: %v\n", err)
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

