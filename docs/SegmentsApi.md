# {{classname}}

All URIs are relative to *https://api.configcat.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSegment**](SegmentsApi.md#CreateSegment) | **Post** /v1/products/{productId}/segments | Create Segment
[**DeleteSegment**](SegmentsApi.md#DeleteSegment) | **Delete** /v1/segments/{segmentId} | Delete Segment
[**GetSegment**](SegmentsApi.md#GetSegment) | **Get** /v1/segments/{segmentId} | Get Segment
[**GetSegments**](SegmentsApi.md#GetSegments) | **Get** /v1/products/{productId}/segments | List Segments
[**UpdateSegment**](SegmentsApi.md#UpdateSegment) | **Put** /v1/segments/{segmentId} | Update Segment

# **CreateSegment**
> SegmentModel CreateSegment(ctx, body, productId)
Create Segment

This endpoint creates a new Segment in a specified Product  identified by the `productId` parameter, which can be obtained from the [List Products](#operation/get-products) endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateSegmentModel**](CreateSegmentModel.md)|  | 
  **productId** | [**string**](.md)| The identifier of the Product. | 

### Return type

[**SegmentModel**](SegmentModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/_*+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSegment**
> DeleteSegment(ctx, segmentId)
Delete Segment

This endpoint removes a Segment identified by the `segmentId` parameter.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **segmentId** | [**string**](.md)| The identifier of the Segment. | 

### Return type

 (empty response body)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSegment**
> SegmentModel GetSegment(ctx, segmentId)
Get Segment

This endpoint returns the metadata of a Segment identified by the `segmentId`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **segmentId** | [**string**](.md)| The identifier of the Segment. | 

### Return type

[**SegmentModel**](SegmentModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSegments**
> []SegmentListModel GetSegments(ctx, productId)
List Segments

This endpoint returns the list of the Segments that belongs to the given Product identified by the `productId` parameter, which can be obtained from the [List Products](#operation/get-products) endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productId** | [**string**](.md)| The identifier of the Product. | 

### Return type

[**[]SegmentListModel**](SegmentListModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSegment**
> SegmentModel UpdateSegment(ctx, body, segmentId)
Update Segment

This endpoint updates a Segment identified by the `segmentId` parameter.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateSegmentModel**](UpdateSegmentModel.md)|  | 
  **segmentId** | [**string**](.md)| The identifier of the Segment. | 

### Return type

[**SegmentModel**](SegmentModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/_*+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

