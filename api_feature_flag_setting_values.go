/*
ConfigCat Public Management API

**Base API URL**: https://test-api.configcat.com  If you prefer the swagger documentation, you can find it here: [Swagger UI](https://test-api.configcat.com/swagger).  The purpose of this API is to access the ConfigCat platform programmatically.  You can **Create**, **Read**, **Update** and **Delete** any entities like **Feature Flags, Configs, Environments** or **Products** within ConfigCat.   The API is based on HTTP REST, uses resource-oriented URLs, status codes and supports JSON  and JSON+HAL format. Do not use this API for accessing and evaluating feature flag values. Use the [SDKs instead](https://configcat.com/docs/sdk-reference/overview).   # OpenAPI Specification  The complete specification is publicly available in the following formats:  - [OpenAPI v3](https://test-api.configcat.com/docs/v1/swagger.json) - [Swagger v2](https://test-api.configcat.com/docs/v1/swagger.v2.json)  You can use it to generate client libraries in various languages with [OpenAPI Generator](https://github.com/OpenAPITools/openapi-generator) or [Swagger Codegen](https://swagger.io/tools/swagger-codegen/) to interact with this API.  # Authentication This API uses the [Basic HTTP Authentication Scheme](https://en.wikipedia.org/wiki/Basic_access_authentication).   <!-- ReDoc-Inject: <security-definitions> -->  # Throttling and rate limits All the rate limited API calls are returning information about the current rate limit period in the following HTTP headers:  | Header | Description | | :- | :- | | X-Rate-Limit-Remaining | The maximum number of requests remaining in the current rate limit period. | | X-Rate-Limit-Reset     | The time when the current rate limit period resets.        |  When the rate limit is exceeded by a request, the API returns with a `HTTP 429 - Too many requests` status along with a `Retry-After` HTTP header. 

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


// FeatureFlagSettingValuesApiService FeatureFlagSettingValuesApi service
type FeatureFlagSettingValuesApiService service

type FeatureFlagSettingValuesApiGetSettingValueRequest struct {
	ctx context.Context
	ApiService *FeatureFlagSettingValuesApiService
	environmentId string
	settingId int32
}

func (r FeatureFlagSettingValuesApiGetSettingValueRequest) Execute() (*SettingValueModel, *http.Response, error) {
	return r.ApiService.GetSettingValueExecute(r)
}

/*
GetSettingValue Get value

This endpoint returns the value of a Feature Flag or Setting 
in a specified Environment identified by the `environmentId` parameter.

The most important attributes in the response are the `value`, `rolloutRules` and `percentageRules`.
The `value` represents what the clients will get when the evaluation requests of our SDKs 
are not matching to any of the defined Targeting or Percentage Rules, or when there are no additional rules to evaluate.

The `rolloutRules` and `percentageRules` attributes are representing the current 
Targeting and Percentage Rules configuration of the actual Feature Flag or Setting 
in an **ordered** collection, which means the order of the returned rules is matching to the
evaluation order. You can read more about these rules [here](https://configcat.com/docs/advanced/targeting/).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param environmentId The identifier of the Environment.
 @param settingId The id of the Setting.
 @return FeatureFlagSettingValuesApiGetSettingValueRequest
*/
func (a *FeatureFlagSettingValuesApiService) GetSettingValue(ctx context.Context, environmentId string, settingId int32) FeatureFlagSettingValuesApiGetSettingValueRequest {
	return FeatureFlagSettingValuesApiGetSettingValueRequest{
		ApiService: a,
		ctx: ctx,
		environmentId: environmentId,
		settingId: settingId,
	}
}

// Execute executes the request
//  @return SettingValueModel
func (a *FeatureFlagSettingValuesApiService) GetSettingValueExecute(r FeatureFlagSettingValuesApiGetSettingValueRequest) (*SettingValueModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SettingValueModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FeatureFlagSettingValuesApiService.GetSettingValue")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{environmentId}/settings/{settingId}/value"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentId"+"}", url.PathEscape(parameterValueToString(r.environmentId, "environmentId")), -1)
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

type FeatureFlagSettingValuesApiGetSettingValuesRequest struct {
	ctx context.Context
	ApiService *FeatureFlagSettingValuesApiService
	configId string
	environmentId string
}

func (r FeatureFlagSettingValuesApiGetSettingValuesRequest) Execute() (*ConfigSettingValuesModel, *http.Response, error) {
	return r.ApiService.GetSettingValuesExecute(r)
}

/*
GetSettingValues Get values

This endpoint returns the value of a specified Config's Feature Flags or Settings identified by the `configId` parameter
in a specified Environment identified by the `environmentId` parameter.

The most important attributes in the response are the `value`, `rolloutRules` and `percentageRules`.
The `value` represents what the clients will get when the evaluation requests of our SDKs 
are not matching to any of the defined Targeting or Percentage Rules, or when there are no additional rules to evaluate.

The `rolloutRules` and `percentageRules` attributes are representing the current 
Targeting and Percentage Rules configuration of the actual Feature Flag or Setting 
in an **ordered** collection, which means the order of the returned rules is matching to the
evaluation order. You can read more about these rules [here](https://configcat.com/docs/advanced/targeting/).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param configId The identifier of the Config.
 @param environmentId The identifier of the Environment.
 @return FeatureFlagSettingValuesApiGetSettingValuesRequest
*/
func (a *FeatureFlagSettingValuesApiService) GetSettingValues(ctx context.Context, configId string, environmentId string) FeatureFlagSettingValuesApiGetSettingValuesRequest {
	return FeatureFlagSettingValuesApiGetSettingValuesRequest{
		ApiService: a,
		ctx: ctx,
		configId: configId,
		environmentId: environmentId,
	}
}

// Execute executes the request
//  @return ConfigSettingValuesModel
func (a *FeatureFlagSettingValuesApiService) GetSettingValuesExecute(r FeatureFlagSettingValuesApiGetSettingValuesRequest) (*ConfigSettingValuesModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ConfigSettingValuesModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FeatureFlagSettingValuesApiService.GetSettingValues")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/configs/{configId}/environments/{environmentId}/values"
	localVarPath = strings.Replace(localVarPath, "{"+"configId"+"}", url.PathEscape(parameterValueToString(r.configId, "configId")), -1)
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

type FeatureFlagSettingValuesApiPostSettingValuesRequest struct {
	ctx context.Context
	ApiService *FeatureFlagSettingValuesApiService
	configId string
	environmentId string
	updateSettingValuesWithIdModel *UpdateSettingValuesWithIdModel
	reason *string
}

func (r FeatureFlagSettingValuesApiPostSettingValuesRequest) UpdateSettingValuesWithIdModel(updateSettingValuesWithIdModel UpdateSettingValuesWithIdModel) FeatureFlagSettingValuesApiPostSettingValuesRequest {
	r.updateSettingValuesWithIdModel = &updateSettingValuesWithIdModel
	return r
}

// The reason note for the Audit Log if the Product&#39;s \&quot;Config changes require a reason\&quot; preference is turned on.
func (r FeatureFlagSettingValuesApiPostSettingValuesRequest) Reason(reason string) FeatureFlagSettingValuesApiPostSettingValuesRequest {
	r.reason = &reason
	return r
}

func (r FeatureFlagSettingValuesApiPostSettingValuesRequest) Execute() (*ConfigSettingValuesModel, *http.Response, error) {
	return r.ApiService.PostSettingValuesExecute(r)
}

/*
PostSettingValues Post values

This endpoint replaces the values of a specified Config's Feature Flags or Settings identified by the `configId` parameter
in a specified Environment identified by the `environmentId` parameter.

Only the `value`, `rolloutRules` and `percentageRules` attributes are modifiable by this endpoint.

**Important:** As this endpoint is doing a complete replace, it's important to set every other attribute that you don't 
want to change in its original state. Not listing one means that it will reset.

For example: We have the following resource.
```
{
    "settingValues": [
		{
			"rolloutPercentageItems": [
				{
					"percentage": 30,
					"value": true
				},
				{
					"percentage": 70,
					"value": false
				}
			],
			"rolloutRules": [],
			"value": false,
			"settingId": 1
		}
	]
}
```
If we send a replace request body as below:
```
{ 
	"settingValues": [
		{
			"value": true,
			"settingId": 1
		}
	]
}
```
Then besides that the default value is set to `true`, all the Percentage Rules are deleted. 
So we get a response like this:
```
{
	"settingValues": [
		{
			"rolloutPercentageItems": [],
			"rolloutRules": [],
			"value": true,
			"setting": 
			{
				"settingId": 1
			}
		}
	]
}
```

The `rolloutRules` property describes two types of rules:

- **Targeting rules**: When you want to add or update a targeting rule, the `comparator`, `comparisonAttribute`, and `comparisonValue` members are required.
- **Segment rules**: When you want to add add or update a segment rule, the `segmentId` which identifies the desired segment and the `segmentComparator` members are required.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param configId The identifier of the Config.
 @param environmentId The identifier of the Environment.
 @return FeatureFlagSettingValuesApiPostSettingValuesRequest
*/
func (a *FeatureFlagSettingValuesApiService) PostSettingValues(ctx context.Context, configId string, environmentId string) FeatureFlagSettingValuesApiPostSettingValuesRequest {
	return FeatureFlagSettingValuesApiPostSettingValuesRequest{
		ApiService: a,
		ctx: ctx,
		configId: configId,
		environmentId: environmentId,
	}
}

// Execute executes the request
//  @return ConfigSettingValuesModel
func (a *FeatureFlagSettingValuesApiService) PostSettingValuesExecute(r FeatureFlagSettingValuesApiPostSettingValuesRequest) (*ConfigSettingValuesModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ConfigSettingValuesModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FeatureFlagSettingValuesApiService.PostSettingValues")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/configs/{configId}/environments/{environmentId}/values"
	localVarPath = strings.Replace(localVarPath, "{"+"configId"+"}", url.PathEscape(parameterValueToString(r.configId, "configId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"environmentId"+"}", url.PathEscape(parameterValueToString(r.environmentId, "environmentId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.updateSettingValuesWithIdModel == nil {
		return localVarReturnValue, nil, reportError("updateSettingValuesWithIdModel is required and must be specified")
	}

	if r.reason != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "reason", r.reason, "")
	}
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
	localVarPostBody = r.updateSettingValuesWithIdModel
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

type FeatureFlagSettingValuesApiReplaceSettingValueRequest struct {
	ctx context.Context
	ApiService *FeatureFlagSettingValuesApiService
	environmentId string
	settingId int32
	updateSettingValueModel *UpdateSettingValueModel
	reason *string
}

func (r FeatureFlagSettingValuesApiReplaceSettingValueRequest) UpdateSettingValueModel(updateSettingValueModel UpdateSettingValueModel) FeatureFlagSettingValuesApiReplaceSettingValueRequest {
	r.updateSettingValueModel = &updateSettingValueModel
	return r
}

// The reason note for the Audit Log if the Product&#39;s \&quot;Config changes require a reason\&quot; preference is turned on.
func (r FeatureFlagSettingValuesApiReplaceSettingValueRequest) Reason(reason string) FeatureFlagSettingValuesApiReplaceSettingValueRequest {
	r.reason = &reason
	return r
}

func (r FeatureFlagSettingValuesApiReplaceSettingValueRequest) Execute() (*SettingValueModel, *http.Response, error) {
	return r.ApiService.ReplaceSettingValueExecute(r)
}

/*
ReplaceSettingValue Replace value

This endpoint replaces the whole value of a Feature Flag or Setting in a specified Environment.

Only the `value`, `rolloutRules` and `percentageRules` attributes are modifiable by this endpoint.

**Important:** As this endpoint is doing a complete replace, it's important to set every other attribute that you don't 
want to change in its original state. Not listing one means that it will reset.

For example: We have the following resource.
```
{
	"rolloutPercentageItems": [
		{
			"percentage": 30,
			"value": true
		},
		{
			"percentage": 70,
			"value": false
		}
	],
	"rolloutRules": [],
	"value": false
}
```
If we send a replace request body as below:
```
{
	"value": true
}
```
Then besides that the default value is set to `true`, all the Percentage Rules are deleted. 
So we get a response like this:
```
{
	"rolloutPercentageItems": [],
	"rolloutRules": [],
	"value": true
}
```

The `rolloutRules` property describes two types of rules:

- **Targeting rules**: When you want to add or update a targeting rule, the `comparator`, `comparisonAttribute`, and `comparisonValue` members are required.
- **Segment rules**: When you want to add add or update a segment rule, the `segmentId` which identifies the desired segment and the `segmentComparator` members are required.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param environmentId The identifier of the Environment.
 @param settingId The id of the Setting.
 @return FeatureFlagSettingValuesApiReplaceSettingValueRequest
*/
func (a *FeatureFlagSettingValuesApiService) ReplaceSettingValue(ctx context.Context, environmentId string, settingId int32) FeatureFlagSettingValuesApiReplaceSettingValueRequest {
	return FeatureFlagSettingValuesApiReplaceSettingValueRequest{
		ApiService: a,
		ctx: ctx,
		environmentId: environmentId,
		settingId: settingId,
	}
}

// Execute executes the request
//  @return SettingValueModel
func (a *FeatureFlagSettingValuesApiService) ReplaceSettingValueExecute(r FeatureFlagSettingValuesApiReplaceSettingValueRequest) (*SettingValueModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SettingValueModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FeatureFlagSettingValuesApiService.ReplaceSettingValue")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{environmentId}/settings/{settingId}/value"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentId"+"}", url.PathEscape(parameterValueToString(r.environmentId, "environmentId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"settingId"+"}", url.PathEscape(parameterValueToString(r.settingId, "settingId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.updateSettingValueModel == nil {
		return localVarReturnValue, nil, reportError("updateSettingValueModel is required and must be specified")
	}

	if r.reason != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "reason", r.reason, "")
	}
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
	localVarPostBody = r.updateSettingValueModel
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

type FeatureFlagSettingValuesApiUpdateSettingValueRequest struct {
	ctx context.Context
	ApiService *FeatureFlagSettingValuesApiService
	environmentId string
	settingId int32
	jsonPatchOperation *[]JsonPatchOperation
	reason *string
}

func (r FeatureFlagSettingValuesApiUpdateSettingValueRequest) JsonPatchOperation(jsonPatchOperation []JsonPatchOperation) FeatureFlagSettingValuesApiUpdateSettingValueRequest {
	r.jsonPatchOperation = &jsonPatchOperation
	return r
}

// The reason note for the Audit Log if the Product&#39;s \&quot;Config changes require a reason\&quot; preference is turned on.
func (r FeatureFlagSettingValuesApiUpdateSettingValueRequest) Reason(reason string) FeatureFlagSettingValuesApiUpdateSettingValueRequest {
	r.reason = &reason
	return r
}

func (r FeatureFlagSettingValuesApiUpdateSettingValueRequest) Execute() (*SettingValueModel, *http.Response, error) {
	return r.ApiService.UpdateSettingValueExecute(r)
}

/*
UpdateSettingValue Update value

This endpoint updates the value of a Feature Flag or Setting 
with a collection of [JSON Patch](http://jsonpatch.com) operations in a specified Environment.

Only the `value`, `rolloutRules` and `percentageRules` attributes are modifiable by this endpoint.

The advantage of using JSON Patch is that you can describe individual update operations on a resource
without touching attributes that you don't want to change. It supports collection reordering, so it also 
can be used for reordering the targeting rules of a Feature Flag or Setting.

For example: We have the following resource.
```
{
	"rolloutPercentageItems": [
		{
			"percentage": 30,
			"value": true
		},
		{
			"percentage": 70,
			"value": false
		}
	],
	"rolloutRules": [],
	"value": false
}
```
If we send an update request body as below:
```
[
	{
		"op": "replace",
		"path": "/value",
		"value": true
	}
]
```
Only the default value is going to be set to `true` and all the Percentage Rules are remaining unchanged.
So we get a response like this:
```
{
	"rolloutPercentageItems": [
		{
			"percentage": 30,
			"value": true
		},
		{
			"percentage": 70,
			"value": false
		}
	],
	"rolloutRules": [],
	"value": true
}
```

The `rolloutRules` property describes two types of rules:

- **Targeting rules**: When you want to add or update a targeting rule, the `comparator`, `comparisonAttribute`, and `comparisonValue` members are required.
- **Segment rules**: When you want to add add or update a segment rule, the `segmentId` which identifies the desired segment and the `segmentComparator` members are required.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param environmentId The identifier of the Environment.
 @param settingId The id of the Setting.
 @return FeatureFlagSettingValuesApiUpdateSettingValueRequest
*/
func (a *FeatureFlagSettingValuesApiService) UpdateSettingValue(ctx context.Context, environmentId string, settingId int32) FeatureFlagSettingValuesApiUpdateSettingValueRequest {
	return FeatureFlagSettingValuesApiUpdateSettingValueRequest{
		ApiService: a,
		ctx: ctx,
		environmentId: environmentId,
		settingId: settingId,
	}
}

// Execute executes the request
//  @return SettingValueModel
func (a *FeatureFlagSettingValuesApiService) UpdateSettingValueExecute(r FeatureFlagSettingValuesApiUpdateSettingValueRequest) (*SettingValueModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SettingValueModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FeatureFlagSettingValuesApiService.UpdateSettingValue")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{environmentId}/settings/{settingId}/value"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentId"+"}", url.PathEscape(parameterValueToString(r.environmentId, "environmentId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"settingId"+"}", url.PathEscape(parameterValueToString(r.settingId, "settingId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.jsonPatchOperation == nil {
		return localVarReturnValue, nil, reportError("jsonPatchOperation is required and must be specified")
	}

	if r.reason != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "reason", r.reason, "")
	}
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
