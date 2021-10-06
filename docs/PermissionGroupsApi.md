# {{classname}}

All URIs are relative to *https://api.configcat.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePermissionGroup**](PermissionGroupsApi.md#CreatePermissionGroup) | **Post** /v1/products/{productId}/permissions | Create Permission Group
[**DeletePermissionGroup**](PermissionGroupsApi.md#DeletePermissionGroup) | **Delete** /v1/permissions/{permissionGroupId} | Delete Permission Group
[**GetPermissionGroup**](PermissionGroupsApi.md#GetPermissionGroup) | **Get** /v1/permissions/{permissionGroupId} | Get Permission Group
[**GetPermissionGroups**](PermissionGroupsApi.md#GetPermissionGroups) | **Get** /v1/products/{productId}/permissions | List Permission Groups
[**UpdatePermissionGroup**](PermissionGroupsApi.md#UpdatePermissionGroup) | **Put** /v1/permissions/{permissionGroupId} | Update Permission Group

# **CreatePermissionGroup**
> PermissionGroupModel CreatePermissionGroup(ctx, body, productId)
Create Permission Group

This endpoint creates a new Permission Group in a specified Product  identified by the `productId` parameter, which can be obtained from the [List Products](#operation/get-products) endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreatePermissionGroupRequest**](CreatePermissionGroupRequest.md)|  | 
  **productId** | [**string**](.md)| The identifier of the Product. | 

### Return type

[**PermissionGroupModel**](PermissionGroupModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/_*+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePermissionGroup**
> DeletePermissionGroup(ctx, permissionGroupId)
Delete Permission Group

This endpoint removes a Permission Group identified by the `permissionGroupId` parameter.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **permissionGroupId** | **int64**| The identifier of the Permission Group. | 

### Return type

 (empty response body)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPermissionGroup**
> PermissionGroupModel GetPermissionGroup(ctx, permissionGroupId)
Get Permission Group

This endpoint returns the metadata of a Permission Group  identified by the `permissionGroupId`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **permissionGroupId** | **int64**| The identifier of the Permission Group. | 

### Return type

[**PermissionGroupModel**](PermissionGroupModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPermissionGroups**
> []PermissionGroupModel GetPermissionGroups(ctx, productId)
List Permission Groups

This endpoint returns the list of the Permission Groups that belongs to the given Product identified by the `productId` parameter, which can be obtained from the [List Products](#operation/get-products) endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productId** | [**string**](.md)| The identifier of the Product. | 

### Return type

[**[]PermissionGroupModel**](PermissionGroupModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePermissionGroup**
> PermissionGroupModel UpdatePermissionGroup(ctx, body, permissionGroupId)
Update Permission Group

This endpoint updates a Permission Group identified by the `permissionGroupId` parameter.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdatePermissionGroupRequest**](UpdatePermissionGroupRequest.md)|  | 
  **permissionGroupId** | **int64**| The identifier of the Permission Group. | 

### Return type

[**PermissionGroupModel**](PermissionGroupModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/_*+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

