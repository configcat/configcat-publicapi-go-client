# {{classname}}

All URIs are relative to *https://test-api.configcat.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAuditlogs**](AuditLogsApi.md#GetAuditlogs) | **Get** /v1/products/{productId}/auditlogs | List Audit logs

# **GetAuditlogs**
> []AuditLogItemModel GetAuditlogs(ctx, productId, optional)
List Audit logs

This endpoint returns the list of Audit logs for a given Product  and the result can be optionally filtered by Config and/or Environment.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productId** | [**string**](.md)| The identifier of the Product. | 
 **optional** | ***AuditLogsApiGetAuditlogsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuditLogsApiGetAuditlogsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **configId** | [**optional.Interface of string**](.md)| The identifier of the Config. | 
 **environmentId** | [**optional.Interface of string**](.md)| The identifier of the Environment. | 
 **auditLogType** | [**optional.Interface of AuditLogType**](.md)| Filter Audit logs by Audit log type. | 
 **fromUtcDateTime** | **optional.Time**| Filter Audit logs by starting UTC date. | 
 **toUtcDateTime** | **optional.Time**| Filter Audit logs by ending UTC date. | 

### Return type

[**[]AuditLogItemModel**](AuditLogItemModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

