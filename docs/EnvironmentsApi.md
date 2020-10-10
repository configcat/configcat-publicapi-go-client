# {{classname}}

All URIs are relative to *https://api.configcat.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEnvironment**](EnvironmentsApi.md#CreateEnvironment) | **Post** /v1/products/{productId}/environments | Create Environment
[**GetEnvironments**](EnvironmentsApi.md#GetEnvironments) | **Get** /v1/products/{productId}/environments | List Environments
[**UpdateEnvironment**](EnvironmentsApi.md#UpdateEnvironment) | **Put** /v1/environments/{environmentId} | Update Environment

# **CreateEnvironment**
> EnvironmentModel CreateEnvironment(ctx, body, productId)
Create Environment

This endpoint creates a new Environment in a specified Product  identified by the `productId` parameter, which can be obtained from the [List Products](#operation/get-products) endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateEnvironmentModel**](CreateEnvironmentModel.md)|  | 
  **productId** | [**string**](.md)| The identifier of the Product. | 

### Return type

[**EnvironmentModel**](EnvironmentModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/_*+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEnvironments**
> []EnvironmentModel GetEnvironments(ctx, productId)
List Environments

This endpoint returns the list of the Environments that belongs to the given Product identified by the `productId` parameter, which can be obtained from the [List Products](#operation/get-products) endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productId** | [**string**](.md)| The identifier of the Product. | 

### Return type

[**[]EnvironmentModel**](EnvironmentModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateEnvironment**
> EnvironmentModel UpdateEnvironment(ctx, body, environmentId)
Update Environment

This endpoint updates an Environment identified by the `environmentId` parameter.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateEnvironmentModel**](UpdateEnvironmentModel.md)|  | 
  **environmentId** | [**string**](.md)| The identifier of the Product. | 

### Return type

[**EnvironmentModel**](EnvironmentModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/_*+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

