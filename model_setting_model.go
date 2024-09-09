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

// checks if the SettingModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SettingModel{}

// SettingModel Metadata of a Feature Flag or Setting.
type SettingModel struct {
	// Identifier of the Feature Flag or Setting.
	SettingId *int32 `json:"settingId,omitempty"`
	// Key of the Feature Flag or Setting.
	Key NullableString `json:"key,omitempty"`
	// Name of the Feature Flag or Setting.
	Name NullableString `json:"name,omitempty"`
	// Description of the Feature Flag or Setting.
	Hint NullableString `json:"hint,omitempty"`
	// The order of the Feature Flag or Setting represented on the ConfigCat Dashboard.
	Order *int32 `json:"order,omitempty"`
	SettingType *SettingType `json:"settingType,omitempty"`
	// Identifier of the Feature Flag's Config.
	ConfigId *string `json:"configId,omitempty"`
	// Name of the Feature Flag's Config.
	ConfigName NullableString `json:"configName,omitempty"`
	// The tags attached to the Feature Flag or Setting.
	Tags []TagModel `json:"tags,omitempty"`
}

// NewSettingModel instantiates a new SettingModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSettingModel() *SettingModel {
	this := SettingModel{}
	return &this
}

// NewSettingModelWithDefaults instantiates a new SettingModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSettingModelWithDefaults() *SettingModel {
	this := SettingModel{}
	return &this
}

// GetSettingId returns the SettingId field value if set, zero value otherwise.
func (o *SettingModel) GetSettingId() int32 {
	if o == nil || IsNil(o.SettingId) {
		var ret int32
		return ret
	}
	return *o.SettingId
}

// GetSettingIdOk returns a tuple with the SettingId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingModel) GetSettingIdOk() (*int32, bool) {
	if o == nil || IsNil(o.SettingId) {
		return nil, false
	}
	return o.SettingId, true
}

// HasSettingId returns a boolean if a field has been set.
func (o *SettingModel) HasSettingId() bool {
	if o != nil && !IsNil(o.SettingId) {
		return true
	}

	return false
}

// SetSettingId gets a reference to the given int32 and assigns it to the SettingId field.
func (o *SettingModel) SetSettingId(v int32) {
	o.SettingId = &v
}

// GetKey returns the Key field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SettingModel) GetKey() string {
	if o == nil || IsNil(o.Key.Get()) {
		var ret string
		return ret
	}
	return *o.Key.Get()
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SettingModel) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Key.Get(), o.Key.IsSet()
}

// HasKey returns a boolean if a field has been set.
func (o *SettingModel) HasKey() bool {
	if o != nil && o.Key.IsSet() {
		return true
	}

	return false
}

// SetKey gets a reference to the given NullableString and assigns it to the Key field.
func (o *SettingModel) SetKey(v string) {
	o.Key.Set(&v)
}
// SetKeyNil sets the value for Key to be an explicit nil
func (o *SettingModel) SetKeyNil() {
	o.Key.Set(nil)
}

// UnsetKey ensures that no value is present for Key, not even an explicit nil
func (o *SettingModel) UnsetKey() {
	o.Key.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SettingModel) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SettingModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *SettingModel) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *SettingModel) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *SettingModel) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *SettingModel) UnsetName() {
	o.Name.Unset()
}

// GetHint returns the Hint field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SettingModel) GetHint() string {
	if o == nil || IsNil(o.Hint.Get()) {
		var ret string
		return ret
	}
	return *o.Hint.Get()
}

// GetHintOk returns a tuple with the Hint field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SettingModel) GetHintOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Hint.Get(), o.Hint.IsSet()
}

// HasHint returns a boolean if a field has been set.
func (o *SettingModel) HasHint() bool {
	if o != nil && o.Hint.IsSet() {
		return true
	}

	return false
}

// SetHint gets a reference to the given NullableString and assigns it to the Hint field.
func (o *SettingModel) SetHint(v string) {
	o.Hint.Set(&v)
}
// SetHintNil sets the value for Hint to be an explicit nil
func (o *SettingModel) SetHintNil() {
	o.Hint.Set(nil)
}

// UnsetHint ensures that no value is present for Hint, not even an explicit nil
func (o *SettingModel) UnsetHint() {
	o.Hint.Unset()
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *SettingModel) GetOrder() int32 {
	if o == nil || IsNil(o.Order) {
		var ret int32
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingModel) GetOrderOk() (*int32, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *SettingModel) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given int32 and assigns it to the Order field.
func (o *SettingModel) SetOrder(v int32) {
	o.Order = &v
}

// GetSettingType returns the SettingType field value if set, zero value otherwise.
func (o *SettingModel) GetSettingType() SettingType {
	if o == nil || IsNil(o.SettingType) {
		var ret SettingType
		return ret
	}
	return *o.SettingType
}

// GetSettingTypeOk returns a tuple with the SettingType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingModel) GetSettingTypeOk() (*SettingType, bool) {
	if o == nil || IsNil(o.SettingType) {
		return nil, false
	}
	return o.SettingType, true
}

// HasSettingType returns a boolean if a field has been set.
func (o *SettingModel) HasSettingType() bool {
	if o != nil && !IsNil(o.SettingType) {
		return true
	}

	return false
}

// SetSettingType gets a reference to the given SettingType and assigns it to the SettingType field.
func (o *SettingModel) SetSettingType(v SettingType) {
	o.SettingType = &v
}

// GetConfigId returns the ConfigId field value if set, zero value otherwise.
func (o *SettingModel) GetConfigId() string {
	if o == nil || IsNil(o.ConfigId) {
		var ret string
		return ret
	}
	return *o.ConfigId
}

// GetConfigIdOk returns a tuple with the ConfigId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingModel) GetConfigIdOk() (*string, bool) {
	if o == nil || IsNil(o.ConfigId) {
		return nil, false
	}
	return o.ConfigId, true
}

// HasConfigId returns a boolean if a field has been set.
func (o *SettingModel) HasConfigId() bool {
	if o != nil && !IsNil(o.ConfigId) {
		return true
	}

	return false
}

// SetConfigId gets a reference to the given string and assigns it to the ConfigId field.
func (o *SettingModel) SetConfigId(v string) {
	o.ConfigId = &v
}

// GetConfigName returns the ConfigName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SettingModel) GetConfigName() string {
	if o == nil || IsNil(o.ConfigName.Get()) {
		var ret string
		return ret
	}
	return *o.ConfigName.Get()
}

// GetConfigNameOk returns a tuple with the ConfigName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SettingModel) GetConfigNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigName.Get(), o.ConfigName.IsSet()
}

// HasConfigName returns a boolean if a field has been set.
func (o *SettingModel) HasConfigName() bool {
	if o != nil && o.ConfigName.IsSet() {
		return true
	}

	return false
}

// SetConfigName gets a reference to the given NullableString and assigns it to the ConfigName field.
func (o *SettingModel) SetConfigName(v string) {
	o.ConfigName.Set(&v)
}
// SetConfigNameNil sets the value for ConfigName to be an explicit nil
func (o *SettingModel) SetConfigNameNil() {
	o.ConfigName.Set(nil)
}

// UnsetConfigName ensures that no value is present for ConfigName, not even an explicit nil
func (o *SettingModel) UnsetConfigName() {
	o.ConfigName.Unset()
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SettingModel) GetTags() []TagModel {
	if o == nil {
		var ret []TagModel
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SettingModel) GetTagsOk() ([]TagModel, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *SettingModel) HasTags() bool {
	if o != nil && IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []TagModel and assigns it to the Tags field.
func (o *SettingModel) SetTags(v []TagModel) {
	o.Tags = v
}

func (o SettingModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SettingModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SettingId) {
		toSerialize["settingId"] = o.SettingId
	}
	if o.Key.IsSet() {
		toSerialize["key"] = o.Key.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Hint.IsSet() {
		toSerialize["hint"] = o.Hint.Get()
	}
	if !IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	if !IsNil(o.SettingType) {
		toSerialize["settingType"] = o.SettingType
	}
	if !IsNil(o.ConfigId) {
		toSerialize["configId"] = o.ConfigId
	}
	if o.ConfigName.IsSet() {
		toSerialize["configName"] = o.ConfigName.Get()
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	return toSerialize, nil
}

type NullableSettingModel struct {
	value *SettingModel
	isSet bool
}

func (v NullableSettingModel) Get() *SettingModel {
	return v.value
}

func (v *NullableSettingModel) Set(val *SettingModel) {
	v.value = val
	v.isSet = true
}

func (v NullableSettingModel) IsSet() bool {
	return v.isSet
}

func (v *NullableSettingModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSettingModel(val *SettingModel) *NullableSettingModel {
	return &NullableSettingModel{value: val, isSet: true}
}

func (v NullableSettingModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSettingModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


