# {{classname}}

All URIs are relative to *https://api.configcat.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddOrUpdateIntegrationLink**](IntegrationLinksApi.md#AddOrUpdateIntegrationLink) | **Post** /v1/environments/{environmentId}/settings/{settingId}/integrationLinks/{integrationLinkType}/{key} | Add or update Integration link
[**DeleteIntegrationLink**](IntegrationLinksApi.md#DeleteIntegrationLink) | **Delete** /v1/environments/{environmentId}/settings/{settingId}/integrationLinks/{integrationLinkType}/{key} | Delete Integration link
[**GetIntegrationLinkDetails**](IntegrationLinksApi.md#GetIntegrationLinkDetails) | **Get** /v1/integrationLink/{integrationLinkType}/{key}/details | Get Integration link

# **AddOrUpdateIntegrationLink**
> IntegrationLinkModel AddOrUpdateIntegrationLink(ctx, environmentId, settingId, integrationLinkType, key, optional)
Add or update Integration link

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **environmentId** | [**string**](.md)| The identifier of the Environment. | 
  **settingId** | **int32**| The id of the Setting. | 
  **integrationLinkType** | [**IntegrationLinkType**](.md)| The integration link&#x27;s type. | 
  **key** | **string**| The key of the integration link. | 
 **optional** | ***IntegrationLinksApiAddOrUpdateIntegrationLinkOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IntegrationLinksApiAddOrUpdateIntegrationLinkOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **body** | [**optional.Interface of AddOrUpdateIntegrationLinkModel**](AddOrUpdateIntegrationLinkModel.md)|  | 

### Return type

[**IntegrationLinkModel**](IntegrationLinkModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/_*+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteIntegrationLink**
> DeleteIntegrationLinkModel DeleteIntegrationLink(ctx, environmentId, settingId, integrationLinkType, key)
Delete Integration link

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **environmentId** | [**string**](.md)| The identifier of the Environment. | 
  **settingId** | **int32**| The id of the Setting. | 
  **integrationLinkType** | [**IntegrationLinkType**](.md)| The integration&#x27;s type. | 
  **key** | **string**| The key of the integration link. | 

### Return type

[**DeleteIntegrationLinkModel**](DeleteIntegrationLinkModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIntegrationLinkDetails**
> IntegrationLinkDetailsModel GetIntegrationLinkDetails(ctx, integrationLinkType, key)
Get Integration link

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **integrationLinkType** | [**IntegrationLinkType**](.md)| The integration link&#x27;s type. | 
  **key** | **string**| The key of the integration link. | 

### Return type

[**IntegrationLinkDetailsModel**](IntegrationLinkDetailsModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

