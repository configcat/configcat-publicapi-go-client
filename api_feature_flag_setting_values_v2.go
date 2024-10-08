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


// FeatureFlagSettingValuesV2APIService FeatureFlagSettingValuesV2API service
type FeatureFlagSettingValuesV2APIService service

type FeatureFlagSettingValuesV2APIGetSettingValueV2Request struct {
	ctx context.Context
	ApiService *FeatureFlagSettingValuesV2APIService
	environmentId string
	settingId int32
}

func (r FeatureFlagSettingValuesV2APIGetSettingValueV2Request) Execute() (*SettingFormulaModel, *http.Response, error) {
	return r.ApiService.GetSettingValueV2Execute(r)
}

/*
GetSettingValueV2 Get value

This endpoint returns the value of a Feature Flag or Setting
in a specified Environment identified by the `environmentId` parameter.

The most important fields in the response are the `defaultValue`, `targetingRules`, and `percentageEvaluationAttribute`.
The `defaultValue` represents what the clients will get when the evaluation requests of our SDKs
are not matching to any of the defined Targeting Rules, or when there are no additional rules to evaluate.

The `targetingRules` represents the current
Targeting Rule configuration of the actual Feature Flag or Setting
in an **ordered** collection, which means the order of the returned rules is matching to the
evaluation order. You can read more about these rules [here](https://configcat.com/docs/targeting/targeting-overview/).

The `percentageEvaluationAttribute` represents the custom [User Object](https://configcat.com/docs/targeting/user-object/) attribute that must be used for [percentage evaluation](https://configcat.com/docs/targeting/percentage-options/) of the Feature Flag or Setting.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param environmentId The identifier of the Environment.
 @param settingId The id of the Setting.
 @return FeatureFlagSettingValuesV2APIGetSettingValueV2Request
*/
func (a *FeatureFlagSettingValuesV2APIService) GetSettingValueV2(ctx context.Context, environmentId string, settingId int32) FeatureFlagSettingValuesV2APIGetSettingValueV2Request {
	return FeatureFlagSettingValuesV2APIGetSettingValueV2Request{
		ApiService: a,
		ctx: ctx,
		environmentId: environmentId,
		settingId: settingId,
	}
}

// Execute executes the request
//  @return SettingFormulaModel
func (a *FeatureFlagSettingValuesV2APIService) GetSettingValueV2Execute(r FeatureFlagSettingValuesV2APIGetSettingValueV2Request) (*SettingFormulaModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SettingFormulaModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FeatureFlagSettingValuesV2APIService.GetSettingValueV2")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/environments/{environmentId}/settings/{settingId}/value"
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

type FeatureFlagSettingValuesV2APIGetSettingValuesV2Request struct {
	ctx context.Context
	ApiService *FeatureFlagSettingValuesV2APIService
	configId string
	environmentId string
}

func (r FeatureFlagSettingValuesV2APIGetSettingValuesV2Request) Execute() (*ConfigSettingFormulasModel, *http.Response, error) {
	return r.ApiService.GetSettingValuesV2Execute(r)
}

/*
GetSettingValuesV2 Get values

This endpoint returns all Feature Flag and Setting values of a Config identified by the `configId` parameter
in a specified Environment identified by the `environmentId` parameter.

The most important fields in the response are the `defaultValue`, `targetingRules`.
The `defaultValue` represents what the clients will get when the evaluation requests of our SDKs
are not matching to any of the defined Targeting Rules, or when there are no additional rules to evaluate.

The `targetingRules` represents the current
Targeting Rule configuration of the actual Feature Flag or Setting
in an **ordered** collection, which means the order of the returned rules is matching to the
evaluation order. You can read more about these rules [here](https://configcat.com/docs/targeting/targeting-overview/).

The `percentageEvaluationAttribute` represents the custom [User Object](https://configcat.com/docs/targeting/user-object/) attribute that must be used for [percentage evaluation](https://configcat.com/docs/targeting/percentage-options/) of the Feature Flag or Setting.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param configId The identifier of the Config.
 @param environmentId The identifier of the Environment.
 @return FeatureFlagSettingValuesV2APIGetSettingValuesV2Request
*/
func (a *FeatureFlagSettingValuesV2APIService) GetSettingValuesV2(ctx context.Context, configId string, environmentId string) FeatureFlagSettingValuesV2APIGetSettingValuesV2Request {
	return FeatureFlagSettingValuesV2APIGetSettingValuesV2Request{
		ApiService: a,
		ctx: ctx,
		configId: configId,
		environmentId: environmentId,
	}
}

// Execute executes the request
//  @return ConfigSettingFormulasModel
func (a *FeatureFlagSettingValuesV2APIService) GetSettingValuesV2Execute(r FeatureFlagSettingValuesV2APIGetSettingValuesV2Request) (*ConfigSettingFormulasModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ConfigSettingFormulasModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FeatureFlagSettingValuesV2APIService.GetSettingValuesV2")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/configs/{configId}/environments/{environmentId}/values"
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

type FeatureFlagSettingValuesV2APIPostSettingValuesV2Request struct {
	ctx context.Context
	ApiService *FeatureFlagSettingValuesV2APIService
	configId string
	environmentId string
	updateEvaluationFormulasModel *UpdateEvaluationFormulasModel
	reason *string
}

func (r FeatureFlagSettingValuesV2APIPostSettingValuesV2Request) UpdateEvaluationFormulasModel(updateEvaluationFormulasModel UpdateEvaluationFormulasModel) FeatureFlagSettingValuesV2APIPostSettingValuesV2Request {
	r.updateEvaluationFormulasModel = &updateEvaluationFormulasModel
	return r
}

// The reason note for the Audit Log if the Product&#39;s \&quot;Config changes require a reason\&quot; preference is turned on.
func (r FeatureFlagSettingValuesV2APIPostSettingValuesV2Request) Reason(reason string) FeatureFlagSettingValuesV2APIPostSettingValuesV2Request {
	r.reason = &reason
	return r
}

func (r FeatureFlagSettingValuesV2APIPostSettingValuesV2Request) Execute() (*ConfigSettingFormulasModel, *http.Response, error) {
	return r.ApiService.PostSettingValuesV2Execute(r)
}

/*
PostSettingValuesV2 Post values

This endpoint batch updates the Feature Flags and Settings of a Config identified by the `configId` parameter
in a specified Environment identified by the `environmentId` parameter.

Only those Feature Flags and Settings are updated which are part of the request, all the others are left untouched.

**Important:** As this endpoint is doing a complete replace on those Feature Flags and Settings, which are set in the request. 
It's important to set every other field that you don't want to change in its original state. Not listing a field means that it will reset.

For example: We have the following resource of a Feature Flag.
```json
{
  "settingFormulas": [
    {
      "defaultValue": {
        "boolValue": false
      },
      "targetingRules": [
        {
          "conditions": [
            {
              "userCondition": {
                "comparisonAttribute": "Email",
                "comparator": "sensitiveTextEquals",
                "comparisonValue": {
                  "stringValue": "test@example.com"
                }
              }
            }
          ],
          "percentageOptions": [],
          "value": {
            "boolValue": true
          }
        }
      ],
      "settingId": 1
    }
  ]
}
```
If we send a batch replace request body as below:
```json
{ 
  "updateFormulas": [
    {
      "defaultValue": {
        "boolValue": false
      },
      "settingId": 1
    }
  ]
}
```
Then besides that the default value is set to `true`, all Targeting Rules of the related Feature Flag are deleted.
So we get a response like this:
```json
{
  "settingFormulas": [
    {
      "defaultValue": {
        "boolValue": false
      },
      "targetingRules": [],
      "setting": 
      {
        "settingId": 1
      }
    }
  ]
}
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param configId The identifier of the Config.
 @param environmentId The identifier of the Environment.
 @return FeatureFlagSettingValuesV2APIPostSettingValuesV2Request
*/
func (a *FeatureFlagSettingValuesV2APIService) PostSettingValuesV2(ctx context.Context, configId string, environmentId string) FeatureFlagSettingValuesV2APIPostSettingValuesV2Request {
	return FeatureFlagSettingValuesV2APIPostSettingValuesV2Request{
		ApiService: a,
		ctx: ctx,
		configId: configId,
		environmentId: environmentId,
	}
}

// Execute executes the request
//  @return ConfigSettingFormulasModel
func (a *FeatureFlagSettingValuesV2APIService) PostSettingValuesV2Execute(r FeatureFlagSettingValuesV2APIPostSettingValuesV2Request) (*ConfigSettingFormulasModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ConfigSettingFormulasModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FeatureFlagSettingValuesV2APIService.PostSettingValuesV2")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/configs/{configId}/environments/{environmentId}/values"
	localVarPath = strings.Replace(localVarPath, "{"+"configId"+"}", url.PathEscape(parameterValueToString(r.configId, "configId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"environmentId"+"}", url.PathEscape(parameterValueToString(r.environmentId, "environmentId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.updateEvaluationFormulasModel == nil {
		return localVarReturnValue, nil, reportError("updateEvaluationFormulasModel is required and must be specified")
	}

	if r.reason != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "reason", r.reason, "form", "")
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
	localVarPostBody = r.updateEvaluationFormulasModel
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

type FeatureFlagSettingValuesV2APIReplaceSettingValueV2Request struct {
	ctx context.Context
	ApiService *FeatureFlagSettingValuesV2APIService
	environmentId string
	settingId int32
	updateEvaluationFormulaModel *UpdateEvaluationFormulaModel
	reason *string
}

func (r FeatureFlagSettingValuesV2APIReplaceSettingValueV2Request) UpdateEvaluationFormulaModel(updateEvaluationFormulaModel UpdateEvaluationFormulaModel) FeatureFlagSettingValuesV2APIReplaceSettingValueV2Request {
	r.updateEvaluationFormulaModel = &updateEvaluationFormulaModel
	return r
}

// The reason note for the Audit Log if the Product&#39;s \&quot;Config changes require a reason\&quot; preference is turned on.
func (r FeatureFlagSettingValuesV2APIReplaceSettingValueV2Request) Reason(reason string) FeatureFlagSettingValuesV2APIReplaceSettingValueV2Request {
	r.reason = &reason
	return r
}

func (r FeatureFlagSettingValuesV2APIReplaceSettingValueV2Request) Execute() (*SettingFormulaModel, *http.Response, error) {
	return r.ApiService.ReplaceSettingValueV2Execute(r)
}

/*
ReplaceSettingValueV2 Replace value

This endpoint replaces the value and the Targeting Rules of a Feature Flag or Setting
in a specified Environment identified by the <a target="_blank" rel="noopener noreferrer" href="https://app.configcat.com/sdkkey">SDK key</a> passed in the `X-CONFIGCAT-SDKKEY` header.

Only the `defaultValue`, `targetingRules`, and `percentageEvaluationAttribute` fields are modifiable by this endpoint.

**Important:** As this endpoint is doing a complete replace, it's important to set every other field that you don't
want to change to its original state. Not listing one means it will reset.

For example: We have the following resource of a Feature Flag.
```json
{
  "defaultValue": {
    "boolValue": false
  },
  "targetingRules": [
    {
      "conditions": [
        {
          "userCondition": {
            "comparisonAttribute": "Email",
            "comparator": "sensitiveTextEquals",
            "comparisonValue": {
              "stringValue": "test@example.com"
            }
          }
        }
      ],
      "percentageOptions": [],
      "value": {
        "boolValue": true
      }
    }
  ]
}
```
If we send a replace request body as below:
```json
{
  "defaultValue": {
    "boolValue": true
  }
}
```
Then besides that the default served value is set to `true`, all the Targeting Rules are deleted.
So we get a response like this:
```json
{
  "defaultValue": {
    "boolValue": true
  },
  "targetingRules": []
}
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param environmentId The identifier of the Environment.
 @param settingId The id of the Setting.
 @return FeatureFlagSettingValuesV2APIReplaceSettingValueV2Request
*/
func (a *FeatureFlagSettingValuesV2APIService) ReplaceSettingValueV2(ctx context.Context, environmentId string, settingId int32) FeatureFlagSettingValuesV2APIReplaceSettingValueV2Request {
	return FeatureFlagSettingValuesV2APIReplaceSettingValueV2Request{
		ApiService: a,
		ctx: ctx,
		environmentId: environmentId,
		settingId: settingId,
	}
}

// Execute executes the request
//  @return SettingFormulaModel
func (a *FeatureFlagSettingValuesV2APIService) ReplaceSettingValueV2Execute(r FeatureFlagSettingValuesV2APIReplaceSettingValueV2Request) (*SettingFormulaModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SettingFormulaModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FeatureFlagSettingValuesV2APIService.ReplaceSettingValueV2")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/environments/{environmentId}/settings/{settingId}/value"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentId"+"}", url.PathEscape(parameterValueToString(r.environmentId, "environmentId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"settingId"+"}", url.PathEscape(parameterValueToString(r.settingId, "settingId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.updateEvaluationFormulaModel == nil {
		return localVarReturnValue, nil, reportError("updateEvaluationFormulaModel is required and must be specified")
	}

	if r.reason != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "reason", r.reason, "form", "")
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
	localVarPostBody = r.updateEvaluationFormulaModel
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

type FeatureFlagSettingValuesV2APIUpdateSettingValueV2Request struct {
	ctx context.Context
	ApiService *FeatureFlagSettingValuesV2APIService
	environmentId string
	settingId int32
	jsonPatchOperation *[]JsonPatchOperation
	reason *string
}

func (r FeatureFlagSettingValuesV2APIUpdateSettingValueV2Request) JsonPatchOperation(jsonPatchOperation []JsonPatchOperation) FeatureFlagSettingValuesV2APIUpdateSettingValueV2Request {
	r.jsonPatchOperation = &jsonPatchOperation
	return r
}

// The reason note for the Audit Log if the Product&#39;s \&quot;Config changes require a reason\&quot; preference is turned on.
func (r FeatureFlagSettingValuesV2APIUpdateSettingValueV2Request) Reason(reason string) FeatureFlagSettingValuesV2APIUpdateSettingValueV2Request {
	r.reason = &reason
	return r
}

func (r FeatureFlagSettingValuesV2APIUpdateSettingValueV2Request) Execute() (*SettingFormulaModel, *http.Response, error) {
	return r.ApiService.UpdateSettingValueV2Execute(r)
}

/*
UpdateSettingValueV2 Update value

This endpoint updates the value of a Feature Flag or Setting
with a collection of [JSON Patch](https://jsonpatch.com) operations in a specified Environment.

Only the `defaultValue`, `targetingRules`, and `percentageEvaluationAttribute` fields are modifiable by this endpoint.

The advantage of using JSON Patch is that you can describe individual update operations on a resource
without touching attributes that you don't want to change. It supports collection reordering, so it also
can be used for reordering the targeting rules of a Feature Flag or Setting.

For example: We have the following resource of a Feature Flag.
```json
{
  "defaultValue": {
    "boolValue": false
  },
  "targetingRules": [
    {
      "conditions": [
        {
          "userCondition": {
            "comparisonAttribute": "Email",
            "comparator": "sensitiveTextEquals",
            "comparisonValue": {
              "stringValue": "test@example.com"
            }
          }
        }
      ],
      "percentageOptions": [],
      "value": {
        "boolValue": true
      }
    }
  ]
}
```
If we send an update request body as below:
```json
[
  {
    "op": "replace",
    "path": "/targetingRules/0/value/boolValue",
    "value": true
  }
]
```
Only the first Targeting Rule's `value` is going to be set to `false` and all the other fields are remaining unchanged.

So we get a response like this:
```json
{
  "defaultValue": {
    "boolValue": false
  },
  "targetingRules": [
    {
      "conditions": [
        {
          "userCondition": {
            "comparisonAttribute": "Email",
            "comparator": "sensitiveTextEquals",
            "comparisonValue": {
              "stringValue": "test@example.com"
            }
          }
        }
      ],
      "percentageOptions": [],
      "value": {
        "boolValue": false
      }
    }
  ]
}
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param environmentId The identifier of the Environment.
 @param settingId The id of the Setting.
 @return FeatureFlagSettingValuesV2APIUpdateSettingValueV2Request
*/
func (a *FeatureFlagSettingValuesV2APIService) UpdateSettingValueV2(ctx context.Context, environmentId string, settingId int32) FeatureFlagSettingValuesV2APIUpdateSettingValueV2Request {
	return FeatureFlagSettingValuesV2APIUpdateSettingValueV2Request{
		ApiService: a,
		ctx: ctx,
		environmentId: environmentId,
		settingId: settingId,
	}
}

// Execute executes the request
//  @return SettingFormulaModel
func (a *FeatureFlagSettingValuesV2APIService) UpdateSettingValueV2Execute(r FeatureFlagSettingValuesV2APIUpdateSettingValueV2Request) (*SettingFormulaModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SettingFormulaModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FeatureFlagSettingValuesV2APIService.UpdateSettingValueV2")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/environments/{environmentId}/settings/{settingId}/value"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentId"+"}", url.PathEscape(parameterValueToString(r.environmentId, "environmentId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"settingId"+"}", url.PathEscape(parameterValueToString(r.settingId, "settingId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.jsonPatchOperation == nil {
		return localVarReturnValue, nil, reportError("jsonPatchOperation is required and must be specified")
	}

	if r.reason != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "reason", r.reason, "form", "")
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
