# {{classname}}

All URIs are relative to *https://api.configcat.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetConfigs**](ConfigsApi.md#GetConfigs) | **Get** /v1/products/{productId}/configs | List Configs

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

