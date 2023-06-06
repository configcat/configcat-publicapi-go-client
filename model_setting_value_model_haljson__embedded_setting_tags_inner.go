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

// checks if the SettingValueModelHaljsonEmbeddedSettingTagsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SettingValueModelHaljsonEmbeddedSettingTagsInner{}

// SettingValueModelHaljsonEmbeddedSettingTagsInner struct for SettingValueModelHaljsonEmbeddedSettingTagsInner
type SettingValueModelHaljsonEmbeddedSettingTagsInner struct {
	SettingTagId *int64 `json:"settingTagId,omitempty"`
	TagId *int64 `json:"tagId,omitempty"`
	Name NullableString `json:"name,omitempty"`
	Color NullableString `json:"color,omitempty"`
}

// NewSettingValueModelHaljsonEmbeddedSettingTagsInner instantiates a new SettingValueModelHaljsonEmbeddedSettingTagsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSettingValueModelHaljsonEmbeddedSettingTagsInner() *SettingValueModelHaljsonEmbeddedSettingTagsInner {
	this := SettingValueModelHaljsonEmbeddedSettingTagsInner{}
	return &this
}

// NewSettingValueModelHaljsonEmbeddedSettingTagsInnerWithDefaults instantiates a new SettingValueModelHaljsonEmbeddedSettingTagsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSettingValueModelHaljsonEmbeddedSettingTagsInnerWithDefaults() *SettingValueModelHaljsonEmbeddedSettingTagsInner {
	this := SettingValueModelHaljsonEmbeddedSettingTagsInner{}
	return &this
}

// GetSettingTagId returns the SettingTagId field value if set, zero value otherwise.
func (o *SettingValueModelHaljsonEmbeddedSettingTagsInner) GetSettingTagId() int64 {
	if o == nil || IsNil(o.SettingTagId) {
		var ret int64
		return ret
	}
	return *o.SettingTagId
}

// GetSettingTagIdOk returns a tuple with the SettingTagId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingValueModelHaljsonEmbeddedSettingTagsInner) GetSettingTagIdOk() (*int64, bool) {
	if o == nil || IsNil(o.SettingTagId) {
		return nil, false
	}
	return o.SettingTagId, true
}

// HasSettingTagId returns a boolean if a field has been set.
func (o *SettingValueModelHaljsonEmbeddedSettingTagsInner) HasSettingTagId() bool {
	if o != nil && !IsNil(o.SettingTagId) {
		return true
	}

	return false
}

// SetSettingTagId gets a reference to the given int64 and assigns it to the SettingTagId field.
func (o *SettingValueModelHaljsonEmbeddedSettingTagsInner) SetSettingTagId(v int64) {
	o.SettingTagId = &v
}

// GetTagId returns the TagId field value if set, zero value otherwise.
func (o *SettingValueModelHaljsonEmbeddedSettingTagsInner) GetTagId() int64 {
	if o == nil || IsNil(o.TagId) {
		var ret int64
		return ret
	}
	return *o.TagId
}

// GetTagIdOk returns a tuple with the TagId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingValueModelHaljsonEmbeddedSettingTagsInner) GetTagIdOk() (*int64, bool) {
	if o == nil || IsNil(o.TagId) {
		return nil, false
	}
	return o.TagId, true
}

// HasTagId returns a boolean if a field has been set.
func (o *SettingValueModelHaljsonEmbeddedSettingTagsInner) HasTagId() bool {
	if o != nil && !IsNil(o.TagId) {
		return true
	}

	return false
}

// SetTagId gets a reference to the given int64 and assigns it to the TagId field.
func (o *SettingValueModelHaljsonEmbeddedSettingTagsInner) SetTagId(v int64) {
	o.TagId = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SettingValueModelHaljsonEmbeddedSettingTagsInner) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SettingValueModelHaljsonEmbeddedSettingTagsInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *SettingValueModelHaljsonEmbeddedSettingTagsInner) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *SettingValueModelHaljsonEmbeddedSettingTagsInner) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *SettingValueModelHaljsonEmbeddedSettingTagsInner) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *SettingValueModelHaljsonEmbeddedSettingTagsInner) UnsetName() {
	o.Name.Unset()
}

// GetColor returns the Color field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SettingValueModelHaljsonEmbeddedSettingTagsInner) GetColor() string {
	if o == nil || IsNil(o.Color.Get()) {
		var ret string
		return ret
	}
	return *o.Color.Get()
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SettingValueModelHaljsonEmbeddedSettingTagsInner) GetColorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Color.Get(), o.Color.IsSet()
}

// HasColor returns a boolean if a field has been set.
func (o *SettingValueModelHaljsonEmbeddedSettingTagsInner) HasColor() bool {
	if o != nil && o.Color.IsSet() {
		return true
	}

	return false
}

// SetColor gets a reference to the given NullableString and assigns it to the Color field.
func (o *SettingValueModelHaljsonEmbeddedSettingTagsInner) SetColor(v string) {
	o.Color.Set(&v)
}
// SetColorNil sets the value for Color to be an explicit nil
func (o *SettingValueModelHaljsonEmbeddedSettingTagsInner) SetColorNil() {
	o.Color.Set(nil)
}

// UnsetColor ensures that no value is present for Color, not even an explicit nil
func (o *SettingValueModelHaljsonEmbeddedSettingTagsInner) UnsetColor() {
	o.Color.Unset()
}

func (o SettingValueModelHaljsonEmbeddedSettingTagsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SettingValueModelHaljsonEmbeddedSettingTagsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SettingTagId) {
		toSerialize["settingTagId"] = o.SettingTagId
	}
	if !IsNil(o.TagId) {
		toSerialize["tagId"] = o.TagId
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Color.IsSet() {
		toSerialize["color"] = o.Color.Get()
	}
	return toSerialize, nil
}

type NullableSettingValueModelHaljsonEmbeddedSettingTagsInner struct {
	value *SettingValueModelHaljsonEmbeddedSettingTagsInner
	isSet bool
}

func (v NullableSettingValueModelHaljsonEmbeddedSettingTagsInner) Get() *SettingValueModelHaljsonEmbeddedSettingTagsInner {
	return v.value
}

func (v *NullableSettingValueModelHaljsonEmbeddedSettingTagsInner) Set(val *SettingValueModelHaljsonEmbeddedSettingTagsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSettingValueModelHaljsonEmbeddedSettingTagsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSettingValueModelHaljsonEmbeddedSettingTagsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSettingValueModelHaljsonEmbeddedSettingTagsInner(val *SettingValueModelHaljsonEmbeddedSettingTagsInner) *NullableSettingValueModelHaljsonEmbeddedSettingTagsInner {
	return &NullableSettingValueModelHaljsonEmbeddedSettingTagsInner{value: val, isSet: true}
}

func (v NullableSettingValueModelHaljsonEmbeddedSettingTagsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSettingValueModelHaljsonEmbeddedSettingTagsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


