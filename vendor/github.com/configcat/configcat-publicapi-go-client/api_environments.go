/*
ConfigCat Public Management API

The purpose of this API is to access the ConfigCat platform programmatically. You can **Create**, **Read**, **Update** and **Delete** any entities like **Feature Flags, Configs, Environments** or **Products** within ConfigCat.  **Base API URL**: https://test-api.configcat.com  If you prefer the swagger documentation, you can find it here: [Swagger UI](https://test-api.configcat.com/swagger).  The API is based on HTTP REST, uses resource-oriented URLs, status codes and supports JSON  format.   **Important:** Do not use this API for accessing and evaluating feature flag values. Use the [SDKs](https://configcat.com/docs/sdk-reference/overview) or the [ConfigCat Proxy](https://configcat.com/docs/advanced/proxy/proxy-overview/) instead.  # OpenAPI Specification  The complete specification is publicly available in the following formats:  - [OpenAPI v3](https://test-api.configcat.com/docs/v1/swagger.json) - [Swagger v2](https://test-api.configcat.com/docs/v1/swagger.v2.json)  You can use it to generate client libraries in various languages with [OpenAPI Generator](https://github.com/OpenAPITools/openapi-generator) or [Swagger Codegen](https://swagger.io/tools/swagger-codegen/) to interact with this API.  # Authentication This API uses the [Basic HTTP Authentication Scheme](https://en.wikipedia.org/wiki/Basic_access_authentication).   <!-- ReDoc-Inject: <security-definitions> -->  # Throttling and rate limits All the rate limited API calls are returning information about the current rate limit period in the following HTTP headers:  | Header | Description | | :- | :- | | X-Rate-Limit-Remaining | The maximum number of requests remaining in the current rate limit period. | | X-Rate-Limit-Reset     | The time when the current rate limit period resets.        |  When the rate limit is exceeded by a request, the API returns with a `HTTP 429 - Too many requests` status along with a `Retry-After` HTTP header. 

API version: v1
Contact: support@configcat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configcatpublicapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// EnvironmentsApiService EnvironmentsApi service
type EnvironmentsApiService service

type EnvironmentsApiCreateEnvironmentRequest struct {
	ctx context.Context
	ApiService *EnvironmentsApiService
	productId string
	createEnvironmentModel *CreateEnvironmentModel
}

func (r EnvironmentsApiCreateEnvironmentRequest) CreateEnvironmentModel(createEnvironmentModel CreateEnvironmentModel) EnvironmentsApiCreateEnvironmentRequest {
	r.createEnvironmentModel = &createEnvironmentModel
	return r
}

func (r EnvironmentsApiCreateEnvironmentRequest) Execute() (*EnvironmentModel, *http.Response, error) {
	return r.ApiService.CreateEnvironmentExecute(r)
}

/*
CreateEnvironment Create Environment

This endpoint creates a new Environment in a specified Product 
identified by the `productId` parameter, which can be obtained from the [List Products](#operation/get-products) endpoint.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param productId The identifier of the Product.
 @return EnvironmentsApiCreateEnvironmentRequest
*/
func (a *EnvironmentsApiService) CreateEnvironment(ctx context.Context, productId string) EnvironmentsApiCreateEnvironmentRequest {
	return EnvironmentsApiCreateEnvironmentRequest{
		ApiService: a,
		ctx: ctx,
		productId: productId,
	}
}

// Execute executes the request
//  @return EnvironmentModel
func (a *EnvironmentsApiService) CreateEnvironmentExecute(r EnvironmentsApiCreateEnvironmentRequest) (*EnvironmentModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EnvironmentModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EnvironmentsApiService.CreateEnvironment")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/products/{productId}/environments"
	localVarPath = strings.Replace(localVarPath, "{"+"productId"+"}", url.PathEscape(parameterValueToString(r.productId, "productId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createEnvironmentModel == nil {
		return localVarReturnValue, nil, reportError("createEnvironmentModel is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "text/json", "application/*+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.createEnvironmentModel
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type EnvironmentsApiDeleteEnvironmentRequest struct {
	ctx context.Context
	ApiService *EnvironmentsApiService
	environmentId string
	cleanupAuditLogs *bool
}

// An optional flag which indicates whether the audit log records related to the environment should be deleted or not.
func (r EnvironmentsApiDeleteEnvironmentRequest) CleanupAuditLogs(cleanupAuditLogs bool) EnvironmentsApiDeleteEnvironmentRequest {
	r.cleanupAuditLogs = &cleanupAuditLogs
	return r
}

func (r EnvironmentsApiDeleteEnvironmentRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteEnvironmentExecute(r)
}

/*
DeleteEnvironment Delete Environment

This endpoint removes an Environment identified by the `environmentId` parameter.
If the `cleanupAuditLogs` flag is set to true, it also deletes the audit log records related to the environment
(except for the `Created a new environment` and `Deleted an environment` records).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param environmentId The identifier of the Environment.
 @return EnvironmentsApiDeleteEnvironmentRequest
*/
func (a *EnvironmentsApiService) DeleteEnvironment(ctx context.Context, environmentId string) EnvironmentsApiDeleteEnvironmentRequest {
	return EnvironmentsApiDeleteEnvironmentRequest{
		ApiService: a,
		ctx: ctx,
		environmentId: environmentId,
	}
}

// Execute executes the request
func (a *EnvironmentsApiService) DeleteEnvironmentExecute(r EnvironmentsApiDeleteEnvironmentRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EnvironmentsApiService.DeleteEnvironment")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{environmentId}"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentId"+"}", url.PathEscape(parameterValueToString(r.environmentId, "environmentId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.cleanupAuditLogs != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cleanupAuditLogs", r.cleanupAuditLogs, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type EnvironmentsApiGetEnvironmentRequest struct {
	ctx context.Context
	ApiService *EnvironmentsApiService
	environmentId string
}

func (r EnvironmentsApiGetEnvironmentRequest) Execute() (*EnvironmentModel, *http.Response, error) {
	return r.ApiService.GetEnvironmentExecute(r)
}

/*
GetEnvironment Get Environment

This endpoint returns the metadata of an Environment 
identified by the `environmentId`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param environmentId The identifier of the Environment.
 @return EnvironmentsApiGetEnvironmentRequest
*/
func (a *EnvironmentsApiService) GetEnvironment(ctx context.Context, environmentId string) EnvironmentsApiGetEnvironmentRequest {
	return EnvironmentsApiGetEnvironmentRequest{
		ApiService: a,
		ctx: ctx,
		environmentId: environmentId,
	}
}

// Execute executes the request
//  @return EnvironmentModel
func (a *EnvironmentsApiService) GetEnvironmentExecute(r EnvironmentsApiGetEnvironmentRequest) (*EnvironmentModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EnvironmentModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EnvironmentsApiService.GetEnvironment")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{environmentId}"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentId"+"}", url.PathEscape(parameterValueToString(r.environmentId, "environmentId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type EnvironmentsApiGetEnvironmentsRequest struct {
	ctx context.Context
	ApiService *EnvironmentsApiService
	productId string
}

func (r EnvironmentsApiGetEnvironmentsRequest) Execute() ([]EnvironmentModel, *http.Response, error) {
	return r.ApiService.GetEnvironmentsExecute(r)
}

/*
GetEnvironments List Environments

This endpoint returns the list of the Environments that belongs to the given Product identified by the
`productId` parameter, which can be obtained from the [List Products](#operation/get-products) endpoint.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param productId The identifier of the Product.
 @return EnvironmentsApiGetEnvironmentsRequest
*/
func (a *EnvironmentsApiService) GetEnvironments(ctx context.Context, productId string) EnvironmentsApiGetEnvironmentsRequest {
	return EnvironmentsApiGetEnvironmentsRequest{
		ApiService: a,
		ctx: ctx,
		productId: productId,
	}
}

// Execute executes the request
//  @return []EnvironmentModel
func (a *EnvironmentsApiService) GetEnvironmentsExecute(r EnvironmentsApiGetEnvironmentsRequest) ([]EnvironmentModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []EnvironmentModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EnvironmentsApiService.GetEnvironments")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/products/{productId}/environments"
	localVarPath = strings.Replace(localVarPath, "{"+"productId"+"}", url.PathEscape(parameterValueToString(r.productId, "productId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type EnvironmentsApiUpdateEnvironmentRequest struct {
	ctx context.Context
	ApiService *EnvironmentsApiService
	environmentId string
	updateEnvironmentModel *UpdateEnvironmentModel
}

func (r EnvironmentsApiUpdateEnvironmentRequest) UpdateEnvironmentModel(updateEnvironmentModel UpdateEnvironmentModel) EnvironmentsApiUpdateEnvironmentRequest {
	r.updateEnvironmentModel = &updateEnvironmentModel
	return r
}

func (r EnvironmentsApiUpdateEnvironmentRequest) Execute() (*EnvironmentModel, *http.Response, error) {
	return r.ApiService.UpdateEnvironmentExecute(r)
}

/*
UpdateEnvironment Update Environment

This endpoint updates an Environment identified by the `environmentId` parameter.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param environmentId The identifier of the Environment.
 @return EnvironmentsApiUpdateEnvironmentRequest
*/
func (a *EnvironmentsApiService) UpdateEnvironment(ctx context.Context, environmentId string) EnvironmentsApiUpdateEnvironmentRequest {
	return EnvironmentsApiUpdateEnvironmentRequest{
		ApiService: a,
		ctx: ctx,
		environmentId: environmentId,
	}
}

// Execute executes the request
//  @return EnvironmentModel
func (a *EnvironmentsApiService) UpdateEnvironmentExecute(r EnvironmentsApiUpdateEnvironmentRequest) (*EnvironmentModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EnvironmentModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EnvironmentsApiService.UpdateEnvironment")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{environmentId}"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentId"+"}", url.PathEscape(parameterValueToString(r.environmentId, "environmentId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.updateEnvironmentModel == nil {
		return localVarReturnValue, nil, reportError("updateEnvironmentModel is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "text/json", "application/*+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.updateEnvironmentModel
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
