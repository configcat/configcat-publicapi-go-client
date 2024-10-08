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


// FeatureFlagsSettingsApiService FeatureFlagsSettingsApi service
type FeatureFlagsSettingsApiService service

type FeatureFlagsSettingsApiCreateSettingRequest struct {
	ctx context.Context
	ApiService *FeatureFlagsSettingsApiService
	configId string
	createSettingInitialValues *CreateSettingInitialValues
}

func (r FeatureFlagsSettingsApiCreateSettingRequest) CreateSettingInitialValues(createSettingInitialValues CreateSettingInitialValues) FeatureFlagsSettingsApiCreateSettingRequest {
	r.createSettingInitialValues = &createSettingInitialValues
	return r
}

func (r FeatureFlagsSettingsApiCreateSettingRequest) Execute() (*SettingModel, *http.Response, error) {
	return r.ApiService.CreateSettingExecute(r)
}

/*
CreateSetting Create Flag

This endpoint creates a new Feature Flag or Setting in a specified Config
identified by the `configId` parameter.

**Important:** The `key` attribute must be unique within the given Config.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param configId The identifier of the Config.
 @return FeatureFlagsSettingsApiCreateSettingRequest
*/
func (a *FeatureFlagsSettingsApiService) CreateSetting(ctx context.Context, configId string) FeatureFlagsSettingsApiCreateSettingRequest {
	return FeatureFlagsSettingsApiCreateSettingRequest{
		ApiService: a,
		ctx: ctx,
		configId: configId,
	}
}

// Execute executes the request
//  @return SettingModel
func (a *FeatureFlagsSettingsApiService) CreateSettingExecute(r FeatureFlagsSettingsApiCreateSettingRequest) (*SettingModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SettingModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FeatureFlagsSettingsApiService.CreateSetting")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/configs/{configId}/settings"
	localVarPath = strings.Replace(localVarPath, "{"+"configId"+"}", url.PathEscape(parameterValueToString(r.configId, "configId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createSettingInitialValues == nil {
		return localVarReturnValue, nil, reportError("createSettingInitialValues is required and must be specified")
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
	localVarPostBody = r.createSettingInitialValues
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

type FeatureFlagsSettingsApiDeleteSettingRequest struct {
	ctx context.Context
	ApiService *FeatureFlagsSettingsApiService
	settingId int32
}

func (r FeatureFlagsSettingsApiDeleteSettingRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteSettingExecute(r)
}

/*
DeleteSetting Delete Flag

This endpoint removes a Feature Flag or Setting from a specified Config, 
identified by the `configId` parameter.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param settingId The identifier of the Setting.
 @return FeatureFlagsSettingsApiDeleteSettingRequest
*/
func (a *FeatureFlagsSettingsApiService) DeleteSetting(ctx context.Context, settingId int32) FeatureFlagsSettingsApiDeleteSettingRequest {
	return FeatureFlagsSettingsApiDeleteSettingRequest{
		ApiService: a,
		ctx: ctx,
		settingId: settingId,
	}
}

// Execute executes the request
func (a *FeatureFlagsSettingsApiService) DeleteSettingExecute(r FeatureFlagsSettingsApiDeleteSettingRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FeatureFlagsSettingsApiService.DeleteSetting")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/settings/{settingId}"
	localVarPath = strings.Replace(localVarPath, "{"+"settingId"+"}", url.PathEscape(parameterValueToString(r.settingId, "settingId")), -1)

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

type FeatureFlagsSettingsApiGetSettingRequest struct {
	ctx context.Context
	ApiService *FeatureFlagsSettingsApiService
	settingId int32
}

func (r FeatureFlagsSettingsApiGetSettingRequest) Execute() (*SettingModel, *http.Response, error) {
	return r.ApiService.GetSettingExecute(r)
}

/*
GetSetting Get Flag

This endpoint returns the metadata attributes of a Feature Flag or Setting 
identified by the `settingId` parameter.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param settingId The identifier of the Setting.
 @return FeatureFlagsSettingsApiGetSettingRequest
*/
func (a *FeatureFlagsSettingsApiService) GetSetting(ctx context.Context, settingId int32) FeatureFlagsSettingsApiGetSettingRequest {
	return FeatureFlagsSettingsApiGetSettingRequest{
		ApiService: a,
		ctx: ctx,
		settingId: settingId,
	}
}

// Execute executes the request
//  @return SettingModel
func (a *FeatureFlagsSettingsApiService) GetSettingExecute(r FeatureFlagsSettingsApiGetSettingRequest) (*SettingModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SettingModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FeatureFlagsSettingsApiService.GetSetting")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/settings/{settingId}"
	localVarPath = strings.Replace(localVarPath, "{"+"settingId"+"}", url.PathEscape(parameterValueToString(r.settingId, "settingId")), -1)

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

type FeatureFlagsSettingsApiGetSettingsRequest struct {
	ctx context.Context
	ApiService *FeatureFlagsSettingsApiService
	configId string
}

func (r FeatureFlagsSettingsApiGetSettingsRequest) Execute() ([]SettingModel, *http.Response, error) {
	return r.ApiService.GetSettingsExecute(r)
}

/*
GetSettings List Flags

This endpoint returns the list of the Feature Flags and Settings defined in a 
specified Config, identified by the `configId` parameter.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param configId The identifier of the Config.
 @return FeatureFlagsSettingsApiGetSettingsRequest
*/
func (a *FeatureFlagsSettingsApiService) GetSettings(ctx context.Context, configId string) FeatureFlagsSettingsApiGetSettingsRequest {
	return FeatureFlagsSettingsApiGetSettingsRequest{
		ApiService: a,
		ctx: ctx,
		configId: configId,
	}
}

// Execute executes the request
//  @return []SettingModel
func (a *FeatureFlagsSettingsApiService) GetSettingsExecute(r FeatureFlagsSettingsApiGetSettingsRequest) ([]SettingModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []SettingModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FeatureFlagsSettingsApiService.GetSettings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/configs/{configId}/settings"
	localVarPath = strings.Replace(localVarPath, "{"+"configId"+"}", url.PathEscape(parameterValueToString(r.configId, "configId")), -1)

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

type FeatureFlagsSettingsApiReplaceSettingRequest struct {
	ctx context.Context
	ApiService *FeatureFlagsSettingsApiService
	settingId int32
	replaceSettingModel *ReplaceSettingModel
}

func (r FeatureFlagsSettingsApiReplaceSettingRequest) ReplaceSettingModel(replaceSettingModel ReplaceSettingModel) FeatureFlagsSettingsApiReplaceSettingRequest {
	r.replaceSettingModel = &replaceSettingModel
	return r
}

func (r FeatureFlagsSettingsApiReplaceSettingRequest) Execute() (*SettingModel, *http.Response, error) {
	return r.ApiService.ReplaceSettingExecute(r)
}

/*
ReplaceSetting Replace Flag

This endpoint replaces the whole value of a Feature Flag or Setting
identified by the `settingId` parameter.

**Important:** As this endpoint is doing a complete replace, it's important to set every other attribute that you don't 
want to change in its original state. Not listing one means it will reset.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param settingId The identifier of the Setting.
 @return FeatureFlagsSettingsApiReplaceSettingRequest
*/
func (a *FeatureFlagsSettingsApiService) ReplaceSetting(ctx context.Context, settingId int32) FeatureFlagsSettingsApiReplaceSettingRequest {
	return FeatureFlagsSettingsApiReplaceSettingRequest{
		ApiService: a,
		ctx: ctx,
		settingId: settingId,
	}
}

// Execute executes the request
//  @return SettingModel
func (a *FeatureFlagsSettingsApiService) ReplaceSettingExecute(r FeatureFlagsSettingsApiReplaceSettingRequest) (*SettingModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SettingModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FeatureFlagsSettingsApiService.ReplaceSetting")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/settings/{settingId}"
	localVarPath = strings.Replace(localVarPath, "{"+"settingId"+"}", url.PathEscape(parameterValueToString(r.settingId, "settingId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.replaceSettingModel == nil {
		return localVarReturnValue, nil, reportError("replaceSettingModel is required and must be specified")
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
	localVarPostBody = r.replaceSettingModel
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

type FeatureFlagsSettingsApiUpdateSettingRequest struct {
	ctx context.Context
	ApiService *FeatureFlagsSettingsApiService
	settingId int32
	jsonPatchOperation *[]JsonPatchOperation
}

func (r FeatureFlagsSettingsApiUpdateSettingRequest) JsonPatchOperation(jsonPatchOperation []JsonPatchOperation) FeatureFlagsSettingsApiUpdateSettingRequest {
	r.jsonPatchOperation = &jsonPatchOperation
	return r
}

func (r FeatureFlagsSettingsApiUpdateSettingRequest) Execute() (*SettingModel, *http.Response, error) {
	return r.ApiService.UpdateSettingExecute(r)
}

/*
UpdateSetting Update Flag

This endpoint updates the metadata of a Feature Flag or Setting 
with a collection of [JSON Patch](https://jsonpatch.com) operations in a specified Config.

Only the `name`, `hint` and `tags` attributes are modifiable by this endpoint.
The `tags` attribute is a simple collection of the [tag IDs](#operation/get-tags) attached to the given setting.

The advantage of using JSON Patch is that you can describe individual update operations on a resource
without touching attributes that you don't want to change.

For example: We have the following resource.
```json
{
  "settingId": 5345,
  "key": "myGrandFeature",
  "name": "Tihs is a naem with soem typos.",
  "hint": "This flag controls my grandioso feature.",
  "settingType": "boolean",
  "tags": [
    {
      "tagId": 0, 
      "name": "sample tag", 
      "color": "whale"
    }
  ]
}
```
If we send an update request body as below (it changes the `name` and adds the already existing tag with the id `2`):
```json
[
  {
    "op": "replace", 
    "path": "/name", 
    "value": "This is the name without typos."
  }, 
  {
    "op": "add", 
    "path": "/tags/-", 
    "value": 2
  }
]
```
Only the `name` and `tags` are updated and all the other attributes remain unchanged.
So we get a response like this:
```json
{
  "settingId": 5345, 
  "key": "myGrandFeature", 
  "name": "This is the name without typos.", 
  "hint": "This flag controls my grandioso feature.", 
  "settingType": "boolean", 
  "tags": [
    {
      "tagId": 0, 
      "name": "sample tag", 
      "color": "whale"
    }, 
    {
      "tagId": 2, 
      "name": "another tag", 
      "color": "koala"
    }
  ]
}
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param settingId The identifier of the Setting.
 @return FeatureFlagsSettingsApiUpdateSettingRequest
*/
func (a *FeatureFlagsSettingsApiService) UpdateSetting(ctx context.Context, settingId int32) FeatureFlagsSettingsApiUpdateSettingRequest {
	return FeatureFlagsSettingsApiUpdateSettingRequest{
		ApiService: a,
		ctx: ctx,
		settingId: settingId,
	}
}

// Execute executes the request
//  @return SettingModel
func (a *FeatureFlagsSettingsApiService) UpdateSettingExecute(r FeatureFlagsSettingsApiUpdateSettingRequest) (*SettingModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SettingModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FeatureFlagsSettingsApiService.UpdateSetting")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/settings/{settingId}"
	localVarPath = strings.Replace(localVarPath, "{"+"settingId"+"}", url.PathEscape(parameterValueToString(r.settingId, "settingId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.jsonPatchOperation == nil {
		return localVarReturnValue, nil, reportError("jsonPatchOperation is required and must be specified")
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
	localVarPostBody = r.jsonPatchOperation
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