/*
ConfigCat Public Management API

The purpose of this API is to access the ConfigCat platform programmatically. You can **Create**, **Read**, **Update** and **Delete** any entities like **Feature Flags, Configs, Environments** or **Products** within ConfigCat.  **Base API URL**: https://api.configcat.com  If you prefer the swagger documentation, you can find it here: [Swagger UI](https://api.configcat.com/swagger).  The API is based on HTTP REST, uses resource-oriented URLs, status codes and supports JSON  format.   **Important:** Do not use this API for accessing and evaluating feature flag values. Use the [SDKs](https://configcat.com/docs/sdk-reference/overview) or the [ConfigCat Proxy](https://configcat.com/docs/advanced/proxy/proxy-overview/) instead.  # OpenAPI Specification  The complete specification is publicly available in the following formats:  - [OpenAPI v3](https://api.configcat.com/docs/v1/swagger.json) - [Swagger v2](https://api.configcat.com/docs/v1/swagger.v2.json)  You can use it to generate client libraries in various languages with [OpenAPI Generator](https://github.com/OpenAPITools/openapi-generator) or [Swagger Codegen](https://swagger.io/tools/swagger-codegen/) to interact with this API.  # Authentication This API uses the [Basic HTTP Authentication Scheme](https://en.wikipedia.org/wiki/Basic_access_authentication).   <!-- ReDoc-Inject: <security-definitions> -->  # Throttling and rate limits All the rate limited API calls are returning information about the current rate limit period in the following HTTP headers:  | Header | Description | | :- | :- | | X-Rate-Limit-Remaining | The maximum number of requests remaining in the current rate limit period. | | X-Rate-Limit-Reset     | The time when the current rate limit period resets.        |  When the rate limit is exceeded by a request, the API returns with a `HTTP 429 - Too many requests` status along with a `Retry-After` HTTP header. 

API version: v1
Contact: support@configcat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configcatpublicapi

import (
	"encoding/json"
)

// checks if the ModifyIntegrationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModifyIntegrationRequest{}

// ModifyIntegrationRequest struct for ModifyIntegrationRequest
type ModifyIntegrationRequest struct {
	// Name of the Integration.
	Name string `json:"name"`
	// Parameters of the Integration.
	Parameters map[string]string `json:"parameters"`
	// List of Environment IDs that are connected with this Integration. If the list is empty, all of the Environments are connected.
	EnvironmentIds []string `json:"environmentIds"`
	// List of Config IDs that are connected with this Integration. If the list is empty, all of the Configs are connected.
	ConfigIds []string `json:"configIds"`
}

// NewModifyIntegrationRequest instantiates a new ModifyIntegrationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModifyIntegrationRequest(name string, parameters map[string]string, environmentIds []string, configIds []string) *ModifyIntegrationRequest {
	this := ModifyIntegrationRequest{}
	this.Name = name
	this.Parameters = parameters
	this.EnvironmentIds = environmentIds
	this.ConfigIds = configIds
	return &this
}

// NewModifyIntegrationRequestWithDefaults instantiates a new ModifyIntegrationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModifyIntegrationRequestWithDefaults() *ModifyIntegrationRequest {
	this := ModifyIntegrationRequest{}
	return &this
}

// GetName returns the Name field value
func (o *ModifyIntegrationRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ModifyIntegrationRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ModifyIntegrationRequest) SetName(v string) {
	o.Name = v
}

// GetParameters returns the Parameters field value
func (o *ModifyIntegrationRequest) GetParameters() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value
// and a boolean to check if the value has been set.
func (o *ModifyIntegrationRequest) GetParametersOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Parameters, true
}

// SetParameters sets field value
func (o *ModifyIntegrationRequest) SetParameters(v map[string]string) {
	o.Parameters = v
}

// GetEnvironmentIds returns the EnvironmentIds field value
func (o *ModifyIntegrationRequest) GetEnvironmentIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.EnvironmentIds
}

// GetEnvironmentIdsOk returns a tuple with the EnvironmentIds field value
// and a boolean to check if the value has been set.
func (o *ModifyIntegrationRequest) GetEnvironmentIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnvironmentIds, true
}

// SetEnvironmentIds sets field value
func (o *ModifyIntegrationRequest) SetEnvironmentIds(v []string) {
	o.EnvironmentIds = v
}

// GetConfigIds returns the ConfigIds field value
func (o *ModifyIntegrationRequest) GetConfigIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ConfigIds
}

// GetConfigIdsOk returns a tuple with the ConfigIds field value
// and a boolean to check if the value has been set.
func (o *ModifyIntegrationRequest) GetConfigIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigIds, true
}

// SetConfigIds sets field value
func (o *ModifyIntegrationRequest) SetConfigIds(v []string) {
	o.ConfigIds = v
}

func (o ModifyIntegrationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModifyIntegrationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["parameters"] = o.Parameters
	toSerialize["environmentIds"] = o.EnvironmentIds
	toSerialize["configIds"] = o.ConfigIds
	return toSerialize, nil
}

type NullableModifyIntegrationRequest struct {
	value *ModifyIntegrationRequest
	isSet bool
}

func (v NullableModifyIntegrationRequest) Get() *ModifyIntegrationRequest {
	return v.value
}

func (v *NullableModifyIntegrationRequest) Set(val *ModifyIntegrationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModifyIntegrationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModifyIntegrationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModifyIntegrationRequest(val *ModifyIntegrationRequest) *NullableModifyIntegrationRequest {
	return &NullableModifyIntegrationRequest{value: val, isSet: true}
}

func (v NullableModifyIntegrationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModifyIntegrationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


