# {{classname}}

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProduct**](ProductsApi.md#CreateProduct) | **Post** /v1/organizations/{organizationId}/products | Create Product
[**DeleteProduct**](ProductsApi.md#DeleteProduct) | **Delete** /v1/products/{productId} | Delete Product
[**GetProduct**](ProductsApi.md#GetProduct) | **Get** /v1/products/{productId} | Get Product
[**GetProducts**](ProductsApi.md#GetProducts) | **Get** /v1/products | List Products
[**UpdateProduct**](ProductsApi.md#UpdateProduct) | **Put** /v1/products/{productId} | Update Product

# **CreateProduct**
> ProductModel CreateProduct(ctx, body, organizationId)
Create Product

This endpoint creates a new Product in a specified Organization   identified by the `organizationId` parameter, which can be obtained from the [List Organizations](#operation/get-organizations) endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateProductRequest**](CreateProductRequest.md)|  | 
  **organizationId** | [**string**](.md)| The identifier of the Organization. | 

### Return type

[**ProductModel**](ProductModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/_*+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteProduct**
> DeleteProduct(ctx, productId)
Delete Product

This endpoint removes a Product identified by the `productId` parameter.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productId** | [**string**](.md)| The identifier of the Product. | 

### Return type

 (empty response body)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProduct**
> ProductModel GetProduct(ctx, productId)
Get Product

This endpoint returns the metadata of a Product   identified by the `productId`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productId** | [**string**](.md)| The identifier of the Product. | 

### Return type

[**ProductModel**](ProductModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProducts**
> []ProductModel GetProducts(ctx, )
List Products

This endpoint returns the list of the Products that belongs to the user.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]ProductModel**](ProductModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateProduct**
> ProductModel UpdateProduct(ctx, body, productId)
Update Product

This endpoint updates a Product identified by the `productId` parameter.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateProductRequest**](UpdateProductRequest.md)|  | 
  **productId** | [**string**](.md)| The identifier of the Product. | 

### Return type

[**ProductModel**](ProductModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/_*+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

