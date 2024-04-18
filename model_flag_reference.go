/*
ConfigCat Public Management API

The purpose of this API is to access the ConfigCat platform programmatically. You can **Create**, **Read**, **Update** and **Delete** any entities like **Feature Flags, Configs, Environments** or **Products** within ConfigCat.  **Base API URL**: https://test-api.configcat.com  If you prefer the swagger documentation, you can find it here: [Swagger UI](https://test-api.configcat.com/swagger).  The API is based on HTTP REST, uses resource-oriented URLs, status codes and supports JSON  format. Do not use this API for accessing and evaluating feature flag values. Use the [SDKs instead](https://configcat.com/docs/sdk-reference/overview).   # OpenAPI Specification  The complete specification is publicly available in the following formats:  - [OpenAPI v3](https://test-api.configcat.com/docs/v1/swagger.json) - [Swagger v2](https://test-api.configcat.com/docs/v1/swagger.v2.json)  You can use it to generate client libraries in various languages with [OpenAPI Generator](https://github.com/OpenAPITools/openapi-generator) or [Swagger Codegen](https://swagger.io/tools/swagger-codegen/) to interact with this API.  # Authentication This API uses the [Basic HTTP Authentication Scheme](https://en.wikipedia.org/wiki/Basic_access_authentication).   <!-- ReDoc-Inject: <security-definitions> -->  # Throttling and rate limits All the rate limited API calls are returning information about the current rate limit period in the following HTTP headers:  | Header | Description | | :- | :- | | X-Rate-Limit-Remaining | The maximum number of requests remaining in the current rate limit period. | | X-Rate-Limit-Reset     | The time when the current rate limit period resets.        |  When the rate limit is exceeded by a request, the API returns with a `HTTP 429 - Too many requests` status along with a `Retry-After` HTTP header. 

API version: v1
Contact: support@configcat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configcatpublicapi

import (
	"encoding/json"
)

// checks if the FlagReference type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FlagReference{}

// FlagReference struct for FlagReference
type FlagReference struct {
	// The identifier of the Feature Flag or Setting the code reference belongs to.
	SettingId int32 `json:"settingId"`
	// The actual references to the given Feature Flag or Setting.
	References []ReferenceLines `json:"references"`
}

// NewFlagReference instantiates a new FlagReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlagReference(settingId int32, references []ReferenceLines) *FlagReference {
	this := FlagReference{}
	this.SettingId = settingId
	this.References = references
	return &this
}

// NewFlagReferenceWithDefaults instantiates a new FlagReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlagReferenceWithDefaults() *FlagReference {
	this := FlagReference{}
	return &this
}

// GetSettingId returns the SettingId field value
func (o *FlagReference) GetSettingId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SettingId
}

// GetSettingIdOk returns a tuple with the SettingId field value
// and a boolean to check if the value has been set.
func (o *FlagReference) GetSettingIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SettingId, true
}

// SetSettingId sets field value
func (o *FlagReference) SetSettingId(v int32) {
	o.SettingId = v
}

// GetReferences returns the References field value
func (o *FlagReference) GetReferences() []ReferenceLines {
	if o == nil {
		var ret []ReferenceLines
		return ret
	}

	return o.References
}

// GetReferencesOk returns a tuple with the References field value
// and a boolean to check if the value has been set.
func (o *FlagReference) GetReferencesOk() ([]ReferenceLines, bool) {
	if o == nil {
		return nil, false
	}
	return o.References, true
}

// SetReferences sets field value
func (o *FlagReference) SetReferences(v []ReferenceLines) {
	o.References = v
}

func (o FlagReference) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FlagReference) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["settingId"] = o.SettingId
	toSerialize["references"] = o.References
	return toSerialize, nil
}

type NullableFlagReference struct {
	value *FlagReference
	isSet bool
}

func (v NullableFlagReference) Get() *FlagReference {
	return v.value
}

func (v *NullableFlagReference) Set(val *FlagReference) {
	v.value = val
	v.isSet = true
}

func (v NullableFlagReference) IsSet() bool {
	return v.isSet
}

func (v *NullableFlagReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlagReference(val *FlagReference) *NullableFlagReference {
	return &NullableFlagReference{value: val, isSet: true}
}

func (v NullableFlagReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlagReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


