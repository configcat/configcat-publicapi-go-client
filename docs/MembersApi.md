# {{classname}}

All URIs are relative to *https://api.configcat.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteOrganizationMember**](MembersApi.md#DeleteOrganizationMember) | **Delete** /v1/organizations/{organizationId}/members/{userId} | Delete Member from Organization
[**DeleteProductMember**](MembersApi.md#DeleteProductMember) | **Delete** /v1/products/{productId}/members/{userId} | Delete Member from Product
[**GetOrganizationMembers**](MembersApi.md#GetOrganizationMembers) | **Get** /v1/organizations/{organizationId}/members | List Organization Members
[**GetProductMembers**](MembersApi.md#GetProductMembers) | **Get** /v1/products/{productId}/members | List Product Members
[**InviteMember**](MembersApi.md#InviteMember) | **Post** /v1/products/{productId}/members/invite | Invite Member

# **DeleteOrganizationMember**
> DeleteOrganizationMember(ctx, organizationId, userId)
Delete Member from Organization

This endpoint removes a Member identified by the `userId` from the  given Organization identified by the `organizationId` parameter.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **organizationId** | [**string**](.md)| The identifier of the Organization. | 
  **userId** | **string**| The identifier of the Member. | 

### Return type

 (empty response body)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteProductMember**
> DeleteProductMember(ctx, productId, userId)
Delete Member from Product

This endpoint removes a Member identified by the `userId` from the  given Product identified by the `productId` parameter.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productId** | [**string**](.md)| The identifier of the Product. | 
  **userId** | **string**| The identifier of the Member. | 

### Return type

 (empty response body)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrganizationMembers**
> []UserModel GetOrganizationMembers(ctx, organizationId)
List Organization Members

This endpoint returns the list of Members that belongs  to the given Organization, identified by the `organizationId` parameter.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **organizationId** | [**string**](.md)| The identifier of the Organization. | 

### Return type

[**[]UserModel**](UserModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProductMembers**
> []MemberModel GetProductMembers(ctx, productId)
List Product Members

This endpoint returns the list of Members that belongs  to the given Product, identified by the `productId` parameter.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productId** | [**string**](.md)| The identifier of the Product. | 

### Return type

[**[]MemberModel**](MemberModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InviteMember**
> InviteMember(ctx, body, productId)
Invite Member

This endpoint invites a Member into the given Product identified by the `productId` parameter.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**InviteMembersRequest**](InviteMembersRequest.md)|  | 
  **productId** | [**string**](.md)| The identifier of the Product. | 

### Return type

 (empty response body)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/_*+json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

