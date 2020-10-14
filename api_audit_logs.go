
/*
 * ConfigCat Public Management API
 *
 * **Base API URL**: https://test-api.configcat.com  If you prefer the swagger documentation, you can find it here: [Swagger UI](https://api.configcat.com/swagger).  The purpose of this API is to access the ConfigCat platform programmatically.  You can **Create**, **Read**, **Update** and **Delete** any entities like **Feature Flags, Configs, Environments** or **Products** within ConfigCat.   The API is based on HTTP REST, uses resource-oriented URLs, status codes and supports JSON  and JSON+HAL format. Do not use this API for accessing and evaluating feature flag values. Use the [SDKs instead](https://configcat.com/docs/sdk-reference/overview).   # OpenAPI Specification  The complete specification is publicly available here: [swagger.json](v1/swagger.json).  You can use it to generate client libraries in various languages with [OpenAPI Generator](https://github.com/OpenAPITools/openapi-generator) or [Swagger Codegen](https://swagger.io/tools/swagger-codegen/) to interact with this API.  # Authentication This API uses the [Basic HTTP Authentication Scheme](https://en.wikipedia.org/wiki/Basic_access_authentication).   <!-- ReDoc-Inject: <security-definitions> -->  # Throttling and rate limits All the rate limited API calls are returning information about the current rate limit period in the following HTTP headers:  | Header | Description | | :- | :- | | X-Rate-Limit-Remaining | The maximum number of requests remaining in the current rate limit period. | | X-Rate-Limit-Reset     | The time when the current rate limit period resets.        |  When the rate limit is exceeded by a request, the API returns with a `HTTP 429 - Too many requests` status along with a `Retry-After` HTTP header. 
 *
 * API version: v1
 * Contact: support@configcat.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package configcatpublicapi

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"fmt"
	"github.com/antihax/optional"
)

// Linger please
var (
	_ context.Context
)

type AuditLogsApiService service
/*
AuditLogsApiService List Audit logs
This endpoint returns the list of Audit logs for a given Product  and the result can be optionally filtered by Config and/or Environment.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param productId The identifier of the Product.
 * @param optional nil or *AuditLogsApiGetAuditlogsOpts - Optional Parameters:
     * @param "ConfigId" (optional.Interface of string) -  The identifier of the Config.
     * @param "EnvironmentId" (optional.Interface of string) -  The identifier of the Environment.
     * @param "AuditLogType" (optional.Interface of AuditLogType) -  Filter Audit logs by Audit log type.
     * @param "FromUtcDateTime" (optional.Time) -  Filter Audit logs by starting UTC date.
     * @param "ToUtcDateTime" (optional.Time) -  Filter Audit logs by ending UTC date.
@return []AuditLogItemModel
*/

type AuditLogsApiGetAuditlogsOpts struct {
    ConfigId optional.Interface
    EnvironmentId optional.Interface
    AuditLogType optional.Interface
    FromUtcDateTime optional.Time
    ToUtcDateTime optional.Time
}

func (a *AuditLogsApiService) GetAuditlogs(ctx context.Context, productId string, localVarOptionals *AuditLogsApiGetAuditlogsOpts) ([]AuditLogItemModel, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue []AuditLogItemModel
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v1/products/{productId}/auditlogs"
	localVarPath = strings.Replace(localVarPath, "{"+"productId"+"}", fmt.Sprintf("%v", productId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.ConfigId.IsSet() {
		localVarQueryParams.Add("configId", parameterToString(localVarOptionals.ConfigId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.EnvironmentId.IsSet() {
		localVarQueryParams.Add("environmentId", parameterToString(localVarOptionals.EnvironmentId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.AuditLogType.IsSet() {
		localVarQueryParams.Add("auditLogType", parameterToString(localVarOptionals.AuditLogType.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FromUtcDateTime.IsSet() {
		localVarQueryParams.Add("fromUtcDateTime", parameterToString(localVarOptionals.FromUtcDateTime.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ToUtcDateTime.IsSet() {
		localVarQueryParams.Add("toUtcDateTime", parameterToString(localVarOptionals.ToUtcDateTime.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json", "application/hal+json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v []AuditLogItemModel
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
