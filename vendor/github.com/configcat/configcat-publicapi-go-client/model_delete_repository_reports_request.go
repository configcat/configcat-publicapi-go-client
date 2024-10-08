/*
ConfigCat Public Management API

The purpose of this API is to access the ConfigCat platform programmatically. You can **Create**, **Read**, **Update** and **Delete** any entities like **Feature Flags, Configs, Environments** or **Products** within ConfigCat.  **Base API URL**: https://test-api.configcat.com  If you prefer the swagger documentation, you can find it here: [Swagger UI](https://test-api.configcat.com/swagger).  The API is based on HTTP REST, uses resource-oriented URLs, status codes and supports JSON  format.   **Important:** Do not use this API for accessing and evaluating feature flag values. Use the [SDKs](https://configcat.com/docs/sdk-reference/overview) or the [ConfigCat Proxy](https://configcat.com/docs/advanced/proxy/proxy-overview/) instead.  # OpenAPI Specification  The complete specification is publicly available in the following formats:  - [OpenAPI v3](https://test-api.configcat.com/docs/v1/swagger.json) - [Swagger v2](https://test-api.configcat.com/docs/v1/swagger.v2.json)  You can use it to generate client libraries in various languages with [OpenAPI Generator](https://github.com/OpenAPITools/openapi-generator) or [Swagger Codegen](https://swagger.io/tools/swagger-codegen/) to interact with this API.  # Authentication This API uses the [Basic HTTP Authentication Scheme](https://en.wikipedia.org/wiki/Basic_access_authentication).   <!-- ReDoc-Inject: <security-definitions> -->  # Throttling and rate limits All the rate limited API calls are returning information about the current rate limit period in the following HTTP headers:  | Header | Description | | :- | :- | | X-Rate-Limit-Remaining | The maximum number of requests remaining in the current rate limit period. | | X-Rate-Limit-Reset     | The time when the current rate limit period resets.        |  When the rate limit is exceeded by a request, the API returns with a `HTTP 429 - Too many requests` status along with a `Retry-After` HTTP header. 

API version: v1
Contact: support@configcat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configcatpublicapi

import (
	"encoding/json"
)

// checks if the DeleteRepositoryReportsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteRepositoryReportsRequest{}

// DeleteRepositoryReportsRequest struct for DeleteRepositoryReportsRequest
type DeleteRepositoryReportsRequest struct {
	// The Config's identifier from where the reports should be deleted.
	ConfigId string `json:"configId"`
	// The source control repository which's reports should be deleted.
	Repository string `json:"repository"`
	// If it's set, only this branch's reports belonging to the given repository will be deleted.
	Branch NullableString `json:"branch,omitempty"`
	// If it's set, only this setting's reports belonging to the given repository will be deleted.
	SettingId NullableInt32 `json:"settingId,omitempty"`
}

// NewDeleteRepositoryReportsRequest instantiates a new DeleteRepositoryReportsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteRepositoryReportsRequest(configId string, repository string) *DeleteRepositoryReportsRequest {
	this := DeleteRepositoryReportsRequest{}
	this.ConfigId = configId
	this.Repository = repository
	return &this
}

// NewDeleteRepositoryReportsRequestWithDefaults instantiates a new DeleteRepositoryReportsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteRepositoryReportsRequestWithDefaults() *DeleteRepositoryReportsRequest {
	this := DeleteRepositoryReportsRequest{}
	return &this
}

// GetConfigId returns the ConfigId field value
func (o *DeleteRepositoryReportsRequest) GetConfigId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConfigId
}

// GetConfigIdOk returns a tuple with the ConfigId field value
// and a boolean to check if the value has been set.
func (o *DeleteRepositoryReportsRequest) GetConfigIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfigId, true
}

// SetConfigId sets field value
func (o *DeleteRepositoryReportsRequest) SetConfigId(v string) {
	o.ConfigId = v
}

// GetRepository returns the Repository field value
func (o *DeleteRepositoryReportsRequest) GetRepository() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value
// and a boolean to check if the value has been set.
func (o *DeleteRepositoryReportsRequest) GetRepositoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Repository, true
}

// SetRepository sets field value
func (o *DeleteRepositoryReportsRequest) SetRepository(v string) {
	o.Repository = v
}

// GetBranch returns the Branch field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeleteRepositoryReportsRequest) GetBranch() string {
	if o == nil || IsNil(o.Branch.Get()) {
		var ret string
		return ret
	}
	return *o.Branch.Get()
}

// GetBranchOk returns a tuple with the Branch field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeleteRepositoryReportsRequest) GetBranchOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Branch.Get(), o.Branch.IsSet()
}

// HasBranch returns a boolean if a field has been set.
func (o *DeleteRepositoryReportsRequest) HasBranch() bool {
	if o != nil && o.Branch.IsSet() {
		return true
	}

	return false
}

// SetBranch gets a reference to the given NullableString and assigns it to the Branch field.
func (o *DeleteRepositoryReportsRequest) SetBranch(v string) {
	o.Branch.Set(&v)
}
// SetBranchNil sets the value for Branch to be an explicit nil
func (o *DeleteRepositoryReportsRequest) SetBranchNil() {
	o.Branch.Set(nil)
}

// UnsetBranch ensures that no value is present for Branch, not even an explicit nil
func (o *DeleteRepositoryReportsRequest) UnsetBranch() {
	o.Branch.Unset()
}

// GetSettingId returns the SettingId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeleteRepositoryReportsRequest) GetSettingId() int32 {
	if o == nil || IsNil(o.SettingId.Get()) {
		var ret int32
		return ret
	}
	return *o.SettingId.Get()
}

// GetSettingIdOk returns a tuple with the SettingId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeleteRepositoryReportsRequest) GetSettingIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.SettingId.Get(), o.SettingId.IsSet()
}

// HasSettingId returns a boolean if a field has been set.
func (o *DeleteRepositoryReportsRequest) HasSettingId() bool {
	if o != nil && o.SettingId.IsSet() {
		return true
	}

	return false
}

// SetSettingId gets a reference to the given NullableInt32 and assigns it to the SettingId field.
func (o *DeleteRepositoryReportsRequest) SetSettingId(v int32) {
	o.SettingId.Set(&v)
}
// SetSettingIdNil sets the value for SettingId to be an explicit nil
func (o *DeleteRepositoryReportsRequest) SetSettingIdNil() {
	o.SettingId.Set(nil)
}

// UnsetSettingId ensures that no value is present for SettingId, not even an explicit nil
func (o *DeleteRepositoryReportsRequest) UnsetSettingId() {
	o.SettingId.Unset()
}

func (o DeleteRepositoryReportsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteRepositoryReportsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["configId"] = o.ConfigId
	toSerialize["repository"] = o.Repository
	if o.Branch.IsSet() {
		toSerialize["branch"] = o.Branch.Get()
	}
	if o.SettingId.IsSet() {
		toSerialize["settingId"] = o.SettingId.Get()
	}
	return toSerialize, nil
}

type NullableDeleteRepositoryReportsRequest struct {
	value *DeleteRepositoryReportsRequest
	isSet bool
}

func (v NullableDeleteRepositoryReportsRequest) Get() *DeleteRepositoryReportsRequest {
	return v.value
}

func (v *NullableDeleteRepositoryReportsRequest) Set(val *DeleteRepositoryReportsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteRepositoryReportsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteRepositoryReportsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteRepositoryReportsRequest(val *DeleteRepositoryReportsRequest) *NullableDeleteRepositoryReportsRequest {
	return &NullableDeleteRepositoryReportsRequest{value: val, isSet: true}
}

func (v NullableDeleteRepositoryReportsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteRepositoryReportsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

