# {{classname}}

All URIs are relative to *https://api.configcat.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSdkKeys**](SDKKeysApi.md#GetSdkKeys) | **Get** /v1/configs/{configId}/environments/{environmentId} | Get SDK Key

# **GetSdkKeys**
> SdkKeysModel GetSdkKeys(ctx, configId, environmentId)
Get SDK Key

This endpoint returns the SDK Key for your Config in a specified Environment.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **configId** | [**string**](.md)| The identifier of the Config. | 
  **environmentId** | [**string**](.md)| The identifier of the Environment. | 

### Return type

[**SdkKeysModel**](SdkKeysModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

