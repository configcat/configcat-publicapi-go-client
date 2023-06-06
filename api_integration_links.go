/*
ConfigCat Public Management API

**Base API URL**: https://test-api.configcat.com  If you prefer the swagger documentation, you can find it here: [Swagger UI](https://test-api.configcat.com/swagger).  The purpose of this API is to access the ConfigCat platform programmatically.  You can **Create**, **Read**, **Update** and **Delete** any entities like **Feature Flags, Configs, Environments** or **Products** within ConfigCat.   The API is based on HTTP REST, uses resource-oriented URLs, status codes and supports JSON  and JSON+HAL format. Do not use this API for accessing and evaluating feature flag values. Use the [SDKs instead](https://configcat.com/docs/sdk-reference/overview).   # OpenAPI Specification  The complete specification is publicly available here: [swagger.json](v1/swagger.json).  You can use it to generate client libraries in various languages with [OpenAPI Generator](https://github.com/OpenAPITools/openapi-generator) or [Swagger Codegen](https://swagger.io/tools/swagger-codegen/) to interact with this API.  # Authentication This API uses the [Basic HTTP Authentication Scheme](https://en.wikipedia.org/wiki/Basic_access_authentication).   <!-- ReDoc-Inject: <security-definitions> -->  # Throttling and rate limits All the rate limited API calls are returning information about the current rate limit period in the following HTTP headers:  | Header | Description | | :- | :- | | X-Rate-Limit-Remaining | The maximum number of requests remaining in the current rate limit period. | | X-Rate-Limit-Reset     | The time when the current rate limit period resets.        |  When the rate limit is exceeded by a request, the API returns with a `HTTP 429 - Too many requests` status along with a `Retry-After` HTTP header. 

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


// IntegrationLinksApiService IntegrationLinksApi service
type IntegrationLinksApiService service

type IntegrationLinksApiAddOrUpdateIntegrationLinkRequest struct {
	ctx context.Context
	ApiService *IntegrationLinksApiService
	environmentId string
	settingId int32
	integrationLinkType IntegrationLinkType
	key string
	addOrUpdateIntegrationLinkModel *AddOrUpdateIntegrationLinkModel
}

func (r IntegrationLinksApiAddOrUpdateIntegrationLinkRequest) AddOrUpdateIntegrationLinkModel(addOrUpdateIntegrationLinkModel AddOrUpdateIntegrationLinkModel) IntegrationLinksApiAddOrUpdateIntegrationLinkRequest {
	r.addOrUpdateIntegrationLinkModel = &addOrUpdateIntegrationLinkModel
	return r
}

func (r IntegrationLinksApiAddOrUpdateIntegrationLinkRequest) Execute() (*IntegrationLinkModel, *http.Response, error) {
	return r.ApiService.AddOrUpdateIntegrationLinkExecute(r)
}

/*
AddOrUpdateIntegrationLink Add or update Integration link



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param environmentId The identifier of the Environment.
 @param settingId The id of the Setting.
 @param integrationLinkType The integration link's type.
 @param key The key of the integration link.
 @return IntegrationLinksApiAddOrUpdateIntegrationLinkRequest
*/
func (a *IntegrationLinksApiService) AddOrUpdateIntegrationLink(ctx context.Context, environmentId string, settingId int32, integrationLinkType IntegrationLinkType, key string) IntegrationLinksApiAddOrUpdateIntegrationLinkRequest {
	return IntegrationLinksApiAddOrUpdateIntegrationLinkRequest{
		ApiService: a,
		ctx: ctx,
		environmentId: environmentId,
		settingId: settingId,
		integrationLinkType: integrationLinkType,
		key: key,
	}
}

// Execute executes the request
//  @return IntegrationLinkModel
func (a *IntegrationLinksApiService) AddOrUpdateIntegrationLinkExecute(r IntegrationLinksApiAddOrUpdateIntegrationLinkRequest) (*IntegrationLinkModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *IntegrationLinkModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IntegrationLinksApiService.AddOrUpdateIntegrationLink")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{environmentId}/settings/{settingId}/integrationLinks/{integrationLinkType}/{key}"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentId"+"}", url.PathEscape(parameterValueToString(r.environmentId, "environmentId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"settingId"+"}", url.PathEscape(parameterValueToString(r.settingId, "settingId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"integrationLinkType"+"}", url.PathEscape(parameterValueToString(r.integrationLinkType, "integrationLinkType")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"key"+"}", url.PathEscape(parameterValueToString(r.key, "key")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "text/json", "application/*+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/hal+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.addOrUpdateIntegrationLinkModel
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

type IntegrationLinksApiDeleteIntegrationLinkRequest struct {
	ctx context.Context
	ApiService *IntegrationLinksApiService
	environmentId string
	settingId int32
	integrationLinkType IntegrationLinkType
	key string
}

func (r IntegrationLinksApiDeleteIntegrationLinkRequest) Execute() (*DeleteIntegrationLinkModel, *http.Response, error) {
	return r.ApiService.DeleteIntegrationLinkExecute(r)
}

/*
DeleteIntegrationLink Delete Integration link



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param environmentId The identifier of the Environment.
 @param settingId The id of the Setting.
 @param integrationLinkType The integration's type.
 @param key The key of the integration link.
 @return IntegrationLinksApiDeleteIntegrationLinkRequest
*/
func (a *IntegrationLinksApiService) DeleteIntegrationLink(ctx context.Context, environmentId string, settingId int32, integrationLinkType IntegrationLinkType, key string) IntegrationLinksApiDeleteIntegrationLinkRequest {
	return IntegrationLinksApiDeleteIntegrationLinkRequest{
		ApiService: a,
		ctx: ctx,
		environmentId: environmentId,
		settingId: settingId,
		integrationLinkType: integrationLinkType,
		key: key,
	}
}

// Execute executes the request
//  @return DeleteIntegrationLinkModel
func (a *IntegrationLinksApiService) DeleteIntegrationLinkExecute(r IntegrationLinksApiDeleteIntegrationLinkRequest) (*DeleteIntegrationLinkModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DeleteIntegrationLinkModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IntegrationLinksApiService.DeleteIntegrationLink")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{environmentId}/settings/{settingId}/integrationLinks/{integrationLinkType}/{key}"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentId"+"}", url.PathEscape(parameterValueToString(r.environmentId, "environmentId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"settingId"+"}", url.PathEscape(parameterValueToString(r.settingId, "settingId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"integrationLinkType"+"}", url.PathEscape(parameterValueToString(r.integrationLinkType, "integrationLinkType")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"key"+"}", url.PathEscape(parameterValueToString(r.key, "key")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/json", "application/hal+json"}

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

type IntegrationLinksApiGetIntegrationLinkDetailsRequest struct {
	ctx context.Context
	ApiService *IntegrationLinksApiService
	integrationLinkType IntegrationLinkType
	key string
}

func (r IntegrationLinksApiGetIntegrationLinkDetailsRequest) Execute() (*IntegrationLinkDetailsModel, *http.Response, error) {
	return r.ApiService.GetIntegrationLinkDetailsExecute(r)
}

/*
GetIntegrationLinkDetails Get Integration link



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param integrationLinkType The integration link's type.
 @param key The key of the integration link.
 @return IntegrationLinksApiGetIntegrationLinkDetailsRequest
*/
func (a *IntegrationLinksApiService) GetIntegrationLinkDetails(ctx context.Context, integrationLinkType IntegrationLinkType, key string) IntegrationLinksApiGetIntegrationLinkDetailsRequest {
	return IntegrationLinksApiGetIntegrationLinkDetailsRequest{
		ApiService: a,
		ctx: ctx,
		integrationLinkType: integrationLinkType,
		key: key,
	}
}

// Execute executes the request
//  @return IntegrationLinkDetailsModel
func (a *IntegrationLinksApiService) GetIntegrationLinkDetailsExecute(r IntegrationLinksApiGetIntegrationLinkDetailsRequest) (*IntegrationLinkDetailsModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *IntegrationLinkDetailsModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IntegrationLinksApiService.GetIntegrationLinkDetails")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/integrationLink/{integrationLinkType}/{key}/details"
	localVarPath = strings.Replace(localVarPath, "{"+"integrationLinkType"+"}", url.PathEscape(parameterValueToString(r.integrationLinkType, "integrationLinkType")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"key"+"}", url.PathEscape(parameterValueToString(r.key, "key")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/json", "application/hal+json"}

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

type IntegrationLinksApiJiraAddOrUpdateIntegrationLinkRequest struct {
	ctx context.Context
	ApiService *IntegrationLinksApiService
	environmentId string
	settingId int32
	key string
	addOrUpdateJiraIntegrationLinkModel *AddOrUpdateJiraIntegrationLinkModel
}

func (r IntegrationLinksApiJiraAddOrUpdateIntegrationLinkRequest) AddOrUpdateJiraIntegrationLinkModel(addOrUpdateJiraIntegrationLinkModel AddOrUpdateJiraIntegrationLinkModel) IntegrationLinksApiJiraAddOrUpdateIntegrationLinkRequest {
	r.addOrUpdateJiraIntegrationLinkModel = &addOrUpdateJiraIntegrationLinkModel
	return r
}

func (r IntegrationLinksApiJiraAddOrUpdateIntegrationLinkRequest) Execute() (*IntegrationLinkModel, *http.Response, error) {
	return r.ApiService.JiraAddOrUpdateIntegrationLinkExecute(r)
}

/*
JiraAddOrUpdateIntegrationLink Method for JiraAddOrUpdateIntegrationLink

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param environmentId The identifier of the Environment.
 @param settingId The id of the Setting.
 @param key The key of the integration link.
 @return IntegrationLinksApiJiraAddOrUpdateIntegrationLinkRequest
*/
func (a *IntegrationLinksApiService) JiraAddOrUpdateIntegrationLink(ctx context.Context, environmentId string, settingId int32, key string) IntegrationLinksApiJiraAddOrUpdateIntegrationLinkRequest {
	return IntegrationLinksApiJiraAddOrUpdateIntegrationLinkRequest{
		ApiService: a,
		ctx: ctx,
		environmentId: environmentId,
		settingId: settingId,
		key: key,
	}
}

// Execute executes the request
//  @return IntegrationLinkModel
func (a *IntegrationLinksApiService) JiraAddOrUpdateIntegrationLinkExecute(r IntegrationLinksApiJiraAddOrUpdateIntegrationLinkRequest) (*IntegrationLinkModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *IntegrationLinkModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IntegrationLinksApiService.JiraAddOrUpdateIntegrationLink")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/jira/environments/{environmentId}/settings/{settingId}/integrationLinks/{key}"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentId"+"}", url.PathEscape(parameterValueToString(r.environmentId, "environmentId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"settingId"+"}", url.PathEscape(parameterValueToString(r.settingId, "settingId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"key"+"}", url.PathEscape(parameterValueToString(r.key, "key")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "text/json", "application/*+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/hal+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.addOrUpdateJiraIntegrationLinkModel
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

type IntegrationLinksApiV1JiraConnectPostRequest struct {
	ctx context.Context
	ApiService *IntegrationLinksApiService
	connectRequest *ConnectRequest
}

func (r IntegrationLinksApiV1JiraConnectPostRequest) ConnectRequest(connectRequest ConnectRequest) IntegrationLinksApiV1JiraConnectPostRequest {
	r.connectRequest = &connectRequest
	return r
}

func (r IntegrationLinksApiV1JiraConnectPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.V1JiraConnectPostExecute(r)
}

/*
V1JiraConnectPost Method for V1JiraConnectPost

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return IntegrationLinksApiV1JiraConnectPostRequest
*/
func (a *IntegrationLinksApiService) V1JiraConnectPost(ctx context.Context) IntegrationLinksApiV1JiraConnectPostRequest {
	return IntegrationLinksApiV1JiraConnectPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *IntegrationLinksApiService) V1JiraConnectPostExecute(r IntegrationLinksApiV1JiraConnectPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IntegrationLinksApiService.V1JiraConnectPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/jira/Connect"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "text/json", "application/*+json"}

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
	// body params
	localVarPostBody = r.connectRequest
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
