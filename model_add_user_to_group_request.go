/*
ConfigCat Public Management API

**Base API URL**: https://api.configcat.com  If you prefer the swagger documentation, you can find it here: [Swagger UI](https://api.configcat.com/swagger).  The purpose of this API is to access the ConfigCat platform programmatically.  You can **Create**, **Read**, **Update** and **Delete** any entities like **Feature Flags, Configs, Environments** or **Products** within ConfigCat.   The API is based on HTTP REST, uses resource-oriented URLs, status codes and supports JSON  and JSON+HAL format. Do not use this API for accessing and evaluating feature flag values. Use the [SDKs instead](https://configcat.com/docs/sdk-reference/overview).   # OpenAPI Specification  The complete specification is publicly available here: [swagger.json](v1/swagger.json).  You can use it to generate client libraries in various languages with [OpenAPI Generator](https://github.com/OpenAPITools/openapi-generator) or [Swagger Codegen](https://swagger.io/tools/swagger-codegen/) to interact with this API.  # Authentication This API uses the [Basic HTTP Authentication Scheme](https://en.wikipedia.org/wiki/Basic_access_authentication).   <!-- ReDoc-Inject: <security-definitions> -->  # Throttling and rate limits All the rate limited API calls are returning information about the current rate limit period in the following HTTP headers:  | Header | Description | | :- | :- | | X-Rate-Limit-Remaining | The maximum number of requests remaining in the current rate limit period. | | X-Rate-Limit-Reset     | The time when the current rate limit period resets.        |  When the rate limit is exceeded by a request, the API returns with a `HTTP 429 - Too many requests` status along with a `Retry-After` HTTP header. 

API version: v1
Contact: support@configcat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configcatpublicapi

import (
	"encoding/json"
)

// checks if the AddUserToGroupRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddUserToGroupRequest{}

// AddUserToGroupRequest struct for AddUserToGroupRequest
type AddUserToGroupRequest struct {
	PermissionGroupIds []int64 `json:"permissionGroupIds"`
}

// NewAddUserToGroupRequest instantiates a new AddUserToGroupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddUserToGroupRequest(permissionGroupIds []int64) *AddUserToGroupRequest {
	this := AddUserToGroupRequest{}
	this.PermissionGroupIds = permissionGroupIds
	return &this
}

// NewAddUserToGroupRequestWithDefaults instantiates a new AddUserToGroupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddUserToGroupRequestWithDefaults() *AddUserToGroupRequest {
	this := AddUserToGroupRequest{}
	return &this
}

// GetPermissionGroupIds returns the PermissionGroupIds field value
func (o *AddUserToGroupRequest) GetPermissionGroupIds() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}

	return o.PermissionGroupIds
}

// GetPermissionGroupIdsOk returns a tuple with the PermissionGroupIds field value
// and a boolean to check if the value has been set.
func (o *AddUserToGroupRequest) GetPermissionGroupIdsOk() ([]int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.PermissionGroupIds, true
}

// SetPermissionGroupIds sets field value
func (o *AddUserToGroupRequest) SetPermissionGroupIds(v []int64) {
	o.PermissionGroupIds = v
}

func (o AddUserToGroupRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddUserToGroupRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["permissionGroupIds"] = o.PermissionGroupIds
	return toSerialize, nil
}

type NullableAddUserToGroupRequest struct {
	value *AddUserToGroupRequest
	isSet bool
}

func (v NullableAddUserToGroupRequest) Get() *AddUserToGroupRequest {
	return v.value
}

func (v *NullableAddUserToGroupRequest) Set(val *AddUserToGroupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddUserToGroupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddUserToGroupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddUserToGroupRequest(val *AddUserToGroupRequest) *NullableAddUserToGroupRequest {
	return &NullableAddUserToGroupRequest{value: val, isSet: true}
}

func (v NullableAddUserToGroupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddUserToGroupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


