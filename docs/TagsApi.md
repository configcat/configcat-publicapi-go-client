# {{classname}}

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTag**](TagsApi.md#CreateTag) | **Post** /v1/products/{productId}/tags | Create Tag
[**DeleteTag**](TagsApi.md#DeleteTag) | **Delete** /v1/tags/{tagId} | Delete Tag
[**GetTag**](TagsApi.md#GetTag) | **Get** /v1/tags/{tagId} | Get Tag
[**GetTags**](TagsApi.md#GetTags) | **Get** /v1/products/{productId}/tags | List Tags
[**UpdateTag**](TagsApi.md#UpdateTag) | **Put** /v1/tags/{tagId} | Update Tag

# **CreateTag**
> TagModel CreateTag(ctx, body, productId)
Create Tag

This endpoint creates a new Tag in a specified Product   identified by the `productId` parameter, which can be obtained from the [List Products](#operation/get-products) endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateTagModel**](CreateTagModel.md)|  | 
  **productId** | [**string**](.md)| The identifier of the Organization. | 

### Return type

[**TagModel**](TagModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/_*+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTag**
> DeleteTag(ctx, tagId)
Delete Tag

This endpoint removes a Tag identified by the `tagId` parameter.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tagId** | **int64**| The identifier of the Tag. | 

### Return type

 (empty response body)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTag**
> TagModel GetTag(ctx, tagId)
Get Tag

This endpoint returns the metadata of a Tag   identified by the `tagId`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tagId** | **int64**| The identifier of the Tag. | 

### Return type

[**TagModel**](TagModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTags**
> []TagModel GetTags(ctx, productId)
List Tags

This endpoint returns the list of the Tags in a   specified Product, identified by the `productId` parameter.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productId** | [**string**](.md)| The identifier of the Product. | 

### Return type

[**[]TagModel**](TagModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTag**
> TagModel UpdateTag(ctx, body, tagId)
Update Tag

This endpoint updates a Tag identified by the `tagId` parameter.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateTagModel**](UpdateTagModel.md)|  | 
  **tagId** | **int64**| The identifier of the Tag. | 

### Return type

[**TagModel**](TagModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/_*+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

