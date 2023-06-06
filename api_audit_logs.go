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
	"time"
)


// AuditLogsApiService AuditLogsApi service
type AuditLogsApiService service

type AuditLogsApiGetAuditlogsRequest struct {
	ctx context.Context
	ApiService *AuditLogsApiService
	productId string
	configId *string
	environmentId *string
	auditLogType *AuditLogType
	fromUtcDateTime *time.Time
	toUtcDateTime *time.Time
}

// The identifier of the Config.
func (r AuditLogsApiGetAuditlogsRequest) ConfigId(configId string) AuditLogsApiGetAuditlogsRequest {
	r.configId = &configId
	return r
}

// The identifier of the Environment.
func (r AuditLogsApiGetAuditlogsRequest) EnvironmentId(environmentId string) AuditLogsApiGetAuditlogsRequest {
	r.environmentId = &environmentId
	return r
}

// Filter Audit logs by Audit log type.
func (r AuditLogsApiGetAuditlogsRequest) AuditLogType(auditLogType AuditLogType) AuditLogsApiGetAuditlogsRequest {
	r.auditLogType = &auditLogType
	return r
}

// Filter Audit logs by starting UTC date.
func (r AuditLogsApiGetAuditlogsRequest) FromUtcDateTime(fromUtcDateTime time.Time) AuditLogsApiGetAuditlogsRequest {
	r.fromUtcDateTime = &fromUtcDateTime
	return r
}

// Filter Audit logs by ending UTC date.
func (r AuditLogsApiGetAuditlogsRequest) ToUtcDateTime(toUtcDateTime time.Time) AuditLogsApiGetAuditlogsRequest {
	r.toUtcDateTime = &toUtcDateTime
	return r
}

func (r AuditLogsApiGetAuditlogsRequest) Execute() ([]AuditLogItemModel, *http.Response, error) {
	return r.ApiService.GetAuditlogsExecute(r)
}

/*
GetAuditlogs List Audit log items for Product

This endpoint returns the list of Audit log items for a given Product 
and the result can be optionally filtered by Config and/or Environment.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param productId The identifier of the Product.
 @return AuditLogsApiGetAuditlogsRequest
*/
func (a *AuditLogsApiService) GetAuditlogs(ctx context.Context, productId string) AuditLogsApiGetAuditlogsRequest {
	return AuditLogsApiGetAuditlogsRequest{
		ApiService: a,
		ctx: ctx,
		productId: productId,
	}
}

// Execute executes the request
//  @return []AuditLogItemModel
func (a *AuditLogsApiService) GetAuditlogsExecute(r AuditLogsApiGetAuditlogsRequest) ([]AuditLogItemModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []AuditLogItemModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuditLogsApiService.GetAuditlogs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/products/{productId}/auditlogs"
	localVarPath = strings.Replace(localVarPath, "{"+"productId"+"}", url.PathEscape(parameterValueToString(r.productId, "productId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.configId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "configId", r.configId, "")
	}
	if r.environmentId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "environmentId", r.environmentId, "")
	}
	if r.auditLogType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "auditLogType", r.auditLogType, "")
	}
	if r.fromUtcDateTime != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fromUtcDateTime", r.fromUtcDateTime, "")
	}
	if r.toUtcDateTime != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "toUtcDateTime", r.toUtcDateTime, "")
	}
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

type AuditLogsApiGetDeletedSettingsRequest struct {
	ctx context.Context
	ApiService *AuditLogsApiService
	configId string
}

func (r AuditLogsApiGetDeletedSettingsRequest) Execute() ([]SettingModel, *http.Response, error) {
	return r.ApiService.GetDeletedSettingsExecute(r)
}

/*
GetDeletedSettings List Deleted Settings

This endpoint returns the list of Feature Flags and Settings that were deleted from the given Config.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param configId The identifier of the Config.
 @return AuditLogsApiGetDeletedSettingsRequest
*/
func (a *AuditLogsApiService) GetDeletedSettings(ctx context.Context, configId string) AuditLogsApiGetDeletedSettingsRequest {
	return AuditLogsApiGetDeletedSettingsRequest{
		ApiService: a,
		ctx: ctx,
		configId: configId,
	}
}

// Execute executes the request
//  @return []SettingModel
func (a *AuditLogsApiService) GetDeletedSettingsExecute(r AuditLogsApiGetDeletedSettingsRequest) ([]SettingModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []SettingModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuditLogsApiService.GetDeletedSettings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/configs/{configId}/deleted-settings"
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

type AuditLogsApiGetOrganizationAuditlogsRequest struct {
	ctx context.Context
	ApiService *AuditLogsApiService
	organizationId string
	productId *string
	configId *string
	environmentId *string
	auditLogType *AuditLogType
	fromUtcDateTime *time.Time
	toUtcDateTime *time.Time
}

// The identifier of the Product.
func (r AuditLogsApiGetOrganizationAuditlogsRequest) ProductId(productId string) AuditLogsApiGetOrganizationAuditlogsRequest {
	r.productId = &productId
	return r
}

// The identifier of the Config.
func (r AuditLogsApiGetOrganizationAuditlogsRequest) ConfigId(configId string) AuditLogsApiGetOrganizationAuditlogsRequest {
	r.configId = &configId
	return r
}

// The identifier of the Environment.
func (r AuditLogsApiGetOrganizationAuditlogsRequest) EnvironmentId(environmentId string) AuditLogsApiGetOrganizationAuditlogsRequest {
	r.environmentId = &environmentId
	return r
}

// Filter Audit logs by Audit log type.
func (r AuditLogsApiGetOrganizationAuditlogsRequest) AuditLogType(auditLogType AuditLogType) AuditLogsApiGetOrganizationAuditlogsRequest {
	r.auditLogType = &auditLogType
	return r
}

// Filter Audit logs by starting UTC date.
func (r AuditLogsApiGetOrganizationAuditlogsRequest) FromUtcDateTime(fromUtcDateTime time.Time) AuditLogsApiGetOrganizationAuditlogsRequest {
	r.fromUtcDateTime = &fromUtcDateTime
	return r
}

// Filter Audit logs by ending UTC date.
func (r AuditLogsApiGetOrganizationAuditlogsRequest) ToUtcDateTime(toUtcDateTime time.Time) AuditLogsApiGetOrganizationAuditlogsRequest {
	r.toUtcDateTime = &toUtcDateTime
	return r
}

func (r AuditLogsApiGetOrganizationAuditlogsRequest) Execute() ([]AuditLogItemModel, *http.Response, error) {
	return r.ApiService.GetOrganizationAuditlogsExecute(r)
}

/*
GetOrganizationAuditlogs List Audit log items for Organization

This endpoint returns the list of Audit log items for a given Organization 
and the result can be optionally filtered by Product and/or Config and/or Environment.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId The identifier of the Organization.
 @return AuditLogsApiGetOrganizationAuditlogsRequest
*/
func (a *AuditLogsApiService) GetOrganizationAuditlogs(ctx context.Context, organizationId string) AuditLogsApiGetOrganizationAuditlogsRequest {
	return AuditLogsApiGetOrganizationAuditlogsRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []AuditLogItemModel
func (a *AuditLogsApiService) GetOrganizationAuditlogsExecute(r AuditLogsApiGetOrganizationAuditlogsRequest) ([]AuditLogItemModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []AuditLogItemModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuditLogsApiService.GetOrganizationAuditlogs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/organizations/{organizationId}/auditlogs"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterValueToString(r.organizationId, "organizationId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.productId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "productId", r.productId, "")
	}
	if r.configId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "configId", r.configId, "")
	}
	if r.environmentId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "environmentId", r.environmentId, "")
	}
	if r.auditLogType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "auditLogType", r.auditLogType, "")
	}
	if r.fromUtcDateTime != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fromUtcDateTime", r.fromUtcDateTime, "")
	}
	if r.toUtcDateTime != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "toUtcDateTime", r.toUtcDateTime, "")
	}
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
