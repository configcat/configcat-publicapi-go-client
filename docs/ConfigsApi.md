# {{classname}}

All URIs are relative to *https://test-api.configcat.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConfig**](ConfigsApi.md#CreateConfig) | **Post** /v1/products/{productId}/configs | Create Config
[**DeleteConfig**](ConfigsApi.md#DeleteConfig) | **Delete** /v1/configs/{configId} | Delete Config
[**GetConfig**](ConfigsApi.md#GetConfig) | **Get** /v1/configs/{configId} | Get Config
[**GetConfigs**](ConfigsApi.md#GetConfigs) | **Get** /v1/products/{productId}/configs | List Configs
[**UpdateConfig**](ConfigsApi.md#UpdateConfig) | **Put** /v1/configs/{configId} | Update Config

# **CreateConfig**
> ConfigModel CreateConfig(ctx, body, productId)
Create Config

This endpoint creates a new Config in a specified Product  identified by the `productId` parameter, which can be obtained from the [List Products](#operation/get-products) endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateConfigRequest**](CreateConfigRequest.md)|  | 
  **productId** | [**string**](.md)| The identifier of the Product. | 

### Return type

[**ConfigModel**](ConfigModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/_*+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteConfig**
> DeleteConfig(ctx, configId)
Delete Config

This endpoint removes a Config identified by the `configId` parameter.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **configId** | [**string**](.md)| The identifier of the Config. | 

### Return type

 (empty response body)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConfig**
> ConfigModel GetConfig(ctx, configId)
Get Config

This endpoint returns the metadata of a Config identified by the `configId`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **configId** | [**string**](.md)| The identifier of the Config. | 

### Return type

[**ConfigModel**](ConfigModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConfigs**
> []ConfigModel GetConfigs(ctx, productId)
List Configs

This endpoint returns the list of the Configs that belongs to the given Product identified by the `productId` parameter, which can be obtained from the [List Products](#operation/get-products) endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productId** | [**string**](.md)| The identifier of the Product. | 

### Return type

[**[]ConfigModel**](ConfigModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateConfig**
> ConfigModel UpdateConfig(ctx, body, configId)
Update Config

This endpoint updates a Config identified by the `configId` parameter.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateConfigRequest**](UpdateConfigRequest.md)|  | 
  **configId** | [**string**](.md)| The identifier of the Config. | 

### Return type

[**ConfigModel**](ConfigModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/_*+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

