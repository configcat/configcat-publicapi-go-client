/*
ConfigCat Public Management API

**Base API URL**: https://test-api.configcat.com  If you prefer the swagger documentation, you can find it here: [Swagger UI](https://test-api.configcat.com/swagger).  The purpose of this API is to access the ConfigCat platform programmatically.  You can **Create**, **Read**, **Update** and **Delete** any entities like **Feature Flags, Configs, Environments** or **Products** within ConfigCat.   The API is based on HTTP REST, uses resource-oriented URLs, status codes and supports JSON  and JSON+HAL format. Do not use this API for accessing and evaluating feature flag values. Use the [SDKs instead](https://configcat.com/docs/sdk-reference/overview).   # OpenAPI Specification  The complete specification is publicly available here: [swagger.json](v1/swagger.json).  You can use it to generate client libraries in various languages with [OpenAPI Generator](https://github.com/OpenAPITools/openapi-generator) or [Swagger Codegen](https://swagger.io/tools/swagger-codegen/) to interact with this API.  # Authentication This API uses the [Basic HTTP Authentication Scheme](https://en.wikipedia.org/wiki/Basic_access_authentication).   <!-- ReDoc-Inject: <security-definitions> -->  # Throttling and rate limits All the rate limited API calls are returning information about the current rate limit period in the following HTTP headers:  | Header | Description | | :- | :- | | X-Rate-Limit-Remaining | The maximum number of requests remaining in the current rate limit period. | | X-Rate-Limit-Reset     | The time when the current rate limit period resets.        |  When the rate limit is exceeded by a request, the API returns with a `HTTP 429 - Too many requests` status along with a `Retry-After` HTTP header. 

API version: v1
Contact: support@configcat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configcatpublicapi

import (
	"encoding/json"
)

// checks if the ConnectRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConnectRequest{}

// ConnectRequest struct for ConnectRequest
type ConnectRequest struct {
	ClientKey string `json:"clientKey"`
	JiraJwtToken string `json:"jiraJwtToken"`
}

// NewConnectRequest instantiates a new ConnectRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectRequest(clientKey string, jiraJwtToken string) *ConnectRequest {
	this := ConnectRequest{}
	this.ClientKey = clientKey
	this.JiraJwtToken = jiraJwtToken
	return &this
}

// NewConnectRequestWithDefaults instantiates a new ConnectRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectRequestWithDefaults() *ConnectRequest {
	this := ConnectRequest{}
	return &this
}

// GetClientKey returns the ClientKey field value
func (o *ConnectRequest) GetClientKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientKey
}

// GetClientKeyOk returns a tuple with the ClientKey field value
// and a boolean to check if the value has been set.
func (o *ConnectRequest) GetClientKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientKey, true
}

// SetClientKey sets field value
func (o *ConnectRequest) SetClientKey(v string) {
	o.ClientKey = v
}

// GetJiraJwtToken returns the JiraJwtToken field value
func (o *ConnectRequest) GetJiraJwtToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JiraJwtToken
}

// GetJiraJwtTokenOk returns a tuple with the JiraJwtToken field value
// and a boolean to check if the value has been set.
func (o *ConnectRequest) GetJiraJwtTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JiraJwtToken, true
}

// SetJiraJwtToken sets field value
func (o *ConnectRequest) SetJiraJwtToken(v string) {
	o.JiraJwtToken = v
}

func (o ConnectRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConnectRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["clientKey"] = o.ClientKey
	toSerialize["jiraJwtToken"] = o.JiraJwtToken
	return toSerialize, nil
}

type NullableConnectRequest struct {
	value *ConnectRequest
	isSet bool
}

func (v NullableConnectRequest) Get() *ConnectRequest {
	return v.value
}

func (v *NullableConnectRequest) Set(val *ConnectRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectRequest(val *ConnectRequest) *NullableConnectRequest {
	return &NullableConnectRequest{value: val, isSet: true}
}

func (v NullableConnectRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


