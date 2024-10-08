/*
ConfigCat Public Management API

**Base API URL**: https://test-api.configcat.com  If you prefer the swagger documentation, you can find it here: [Swagger UI](https://test-api.configcat.com/swagger).  The purpose of this API is to access the ConfigCat platform programmatically.  You can **Create**, **Read**, **Update** and **Delete** any entities like **Feature Flags, Configs, Environments** or **Products** within ConfigCat.   The API is based on HTTP REST, uses resource-oriented URLs, status codes and supports JSON  and JSON+HAL format. Do not use this API for accessing and evaluating feature flag values. Use the [SDKs instead](https://configcat.com/docs/sdk-reference/overview).   # OpenAPI Specification  The complete specification is publicly available in the following formats:  - [OpenAPI v3](https://test-api.configcat.com/docs/v1/swagger.json) - [Swagger v2](https://test-api.configcat.com/docs/v1/swagger.v2.json)  You can use it to generate client libraries in various languages with [OpenAPI Generator](https://github.com/OpenAPITools/openapi-generator) or [Swagger Codegen](https://swagger.io/tools/swagger-codegen/) to interact with this API.  # Authentication This API uses the [Basic HTTP Authentication Scheme](https://en.wikipedia.org/wiki/Basic_access_authentication).   <!-- ReDoc-Inject: <security-definitions> -->  # Throttling and rate limits All the rate limited API calls are returning information about the current rate limit period in the following HTTP headers:  | Header | Description | | :- | :- | | X-Rate-Limit-Remaining | The maximum number of requests remaining in the current rate limit period. | | X-Rate-Limit-Reset     | The time when the current rate limit period resets.        |  When the rate limit is exceeded by a request, the API returns with a `HTTP 429 - Too many requests` status along with a `Retry-After` HTTP header. 

API version: v1
Contact: support@configcat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configcatpublicapi

import (
	"encoding/json"
	"time"
)

// checks if the SettingFormulaModelHaljsonEmbeddedSetting type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SettingFormulaModelHaljsonEmbeddedSetting{}

// SettingFormulaModelHaljsonEmbeddedSetting Metadata of a Feature Flag or Setting.
type SettingFormulaModelHaljsonEmbeddedSetting struct {
	// Identifier of the Feature Flag or Setting.
	SettingId *int32 `json:"settingId,omitempty"`
	// Key of the Feature Flag or Setting.
	Key NullableString `json:"key,omitempty"`
	// Name of the Feature Flag or Setting.
	Name NullableString `json:"name,omitempty"`
	// Description of the Feature Flag or Setting.
	Hint NullableString `json:"hint,omitempty"`
	SettingType *SettingType `json:"settingType,omitempty"`
	// The order of the Feature Flag or Setting represented on the ConfigCat Dashboard.
	Order *int32 `json:"order,omitempty"`
	// The creation time of the Feature Flag or Setting.
	CreatedAt NullableTime `json:"createdAt,omitempty"`
	// The user's email address who created the Feature Flag or Setting.
	CreatorEmail NullableString `json:"creatorEmail,omitempty"`
	// The user's name who created the Feature Flag or Setting.
	CreatorFullName NullableString `json:"creatorFullName,omitempty"`
	IsWatching *bool `json:"isWatching,omitempty"`
}

// NewSettingFormulaModelHaljsonEmbeddedSetting instantiates a new SettingFormulaModelHaljsonEmbeddedSetting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSettingFormulaModelHaljsonEmbeddedSetting() *SettingFormulaModelHaljsonEmbeddedSetting {
	this := SettingFormulaModelHaljsonEmbeddedSetting{}
	return &this
}

// NewSettingFormulaModelHaljsonEmbeddedSettingWithDefaults instantiates a new SettingFormulaModelHaljsonEmbeddedSetting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSettingFormulaModelHaljsonEmbeddedSettingWithDefaults() *SettingFormulaModelHaljsonEmbeddedSetting {
	this := SettingFormulaModelHaljsonEmbeddedSetting{}
	return &this
}

// GetSettingId returns the SettingId field value if set, zero value otherwise.
func (o *SettingFormulaModelHaljsonEmbeddedSetting) GetSettingId() int32 {
	if o == nil || IsNil(o.SettingId) {
		var ret int32
		return ret
	}
	return *o.SettingId
}

// GetSettingIdOk returns a tuple with the SettingId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingFormulaModelHaljsonEmbeddedSetting) GetSettingIdOk() (*int32, bool) {
	if o == nil || IsNil(o.SettingId) {
		return nil, false
	}
	return o.SettingId, true
}

// HasSettingId returns a boolean if a field has been set.
func (o *SettingFormulaModelHaljsonEmbeddedSetting) HasSettingId() bool {
	if o != nil && !IsNil(o.SettingId) {
		return true
	}

	return false
}

// SetSettingId gets a reference to the given int32 and assigns it to the SettingId field.
func (o *SettingFormulaModelHaljsonEmbeddedSetting) SetSettingId(v int32) {
	o.SettingId = &v
}

// GetKey returns the Key field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SettingFormulaModelHaljsonEmbeddedSetting) GetKey() string {
	if o == nil || IsNil(o.Key.Get()) {
		var ret string
		return ret
	}
	return *o.Key.Get()
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SettingFormulaModelHaljsonEmbeddedSetting) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Key.Get(), o.Key.IsSet()
}

// HasKey returns a boolean if a field has been set.
func (o *SettingFormulaModelHaljsonEmbeddedSetting) HasKey() bool {
	if o != nil && o.Key.IsSet() {
		return true
	}

	return false
}

// SetKey gets a reference to the given NullableString and assigns it to the Key field.
func (o *SettingFormulaModelHaljsonEmbeddedSetting) SetKey(v string) {
	o.Key.Set(&v)
}
// SetKeyNil sets the value for Key to be an explicit nil
func (o *SettingFormulaModelHaljsonEmbeddedSetting) SetKeyNil() {
	o.Key.Set(nil)
}

// UnsetKey ensures that no value is present for Key, not even an explicit nil
func (o *SettingFormulaModelHaljsonEmbeddedSetting) UnsetKey() {
	o.Key.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SettingFormulaModelHaljsonEmbeddedSetting) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SettingFormulaModelHaljsonEmbeddedSetting) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *SettingFormulaModelHaljsonEmbeddedSetting) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *SettingFormulaModelHaljsonEmbeddedSetting) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *SettingFormulaModelHaljsonEmbeddedSetting) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *SettingFormulaModelHaljsonEmbeddedSetting) UnsetName() {
	o.Name.Unset()
}

// GetHint returns the Hint field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SettingFormulaModelHaljsonEmbeddedSetting) GetHint() string {
	if o == nil || IsNil(o.Hint.Get()) {
		var ret string
		return ret
	}
	return *o.Hint.Get()
}

// GetHintOk returns a tuple with the Hint field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SettingFormulaModelHaljsonEmbeddedSetting) GetHintOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Hint.Get(), o.Hint.IsSet()
}

// HasHint returns a boolean if a field has been set.
func (o *SettingFormulaModelHaljsonEmbeddedSetting) HasHint() bool {
	if o != nil && o.Hint.IsSet() {
		return true
	}

	return false
}

// SetHint gets a reference to the given NullableString and assigns it to the Hint field.
func (o *SettingFormulaModelHaljsonEmbeddedSetting) SetHint(v string) {
	o.Hint.Set(&v)
}
// SetHintNil sets the value for Hint to be an explicit nil
func (o *SettingFormulaModelHaljsonEmbeddedSetting) SetHintNil() {
	o.Hint.Set(nil)
}

// UnsetHint ensures that no value is present for Hint, not even an explicit nil
func (o *SettingFormulaModelHaljsonEmbeddedSetting) UnsetHint() {
	o.Hint.Unset()
}

// GetSettingType returns the SettingType field value if set, zero value otherwise.
func (o *SettingFormulaModelHaljsonEmbeddedSetting) GetSettingType() SettingType {
	if o == nil || IsNil(o.SettingType) {
		var ret SettingType
		return ret
	}
	return *o.SettingType
}

// GetSettingTypeOk returns a tuple with the SettingType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingFormulaModelHaljsonEmbeddedSetting) GetSettingTypeOk() (*SettingType, bool) {
	if o == nil || IsNil(o.SettingType) {
		return nil, false
	}
	return o.SettingType, true
}

// HasSettingType returns a boolean if a field has been set.
func (o *SettingFormulaModelHaljsonEmbeddedSetting) HasSettingType() bool {
	if o != nil && !IsNil(o.SettingType) {
		return true
	}

	return false
}

// SetSettingType gets a reference to the given SettingType and assigns it to the SettingType field.
func (o *SettingFormulaModelHaljsonEmbeddedSetting) SetSettingType(v SettingType) {
	o.SettingType = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *SettingFormulaModelHaljsonEmbeddedSetting) GetOrder() int32 {
	if o == nil || IsNil(o.Order) {
		var ret int32
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingFormulaModelHaljsonEmbeddedSetting) GetOrderOk() (*int32, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *SettingFormulaModelHaljsonEmbeddedSetting) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given int32 and assigns it to the Order field.
func (o *SettingFormulaModelHaljsonEmbeddedSetting) SetOrder(v int32) {
	o.Order = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SettingFormulaModelHaljsonEmbeddedSetting) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SettingFormulaModelHaljsonEmbeddedSetting) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *SettingFormulaModelHaljsonEmbeddedSetting) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableTime and assigns it to the CreatedAt field.
func (o *SettingFormulaModelHaljsonEmbeddedSetting) SetCreatedAt(v time.Time) {
	o.CreatedAt.Set(&v)
}
// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *SettingFormulaModelHaljsonEmbeddedSetting) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *SettingFormulaModelHaljsonEmbeddedSetting) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetCreatorEmail returns the CreatorEmail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SettingFormulaModelHaljsonEmbeddedSetting) GetCreatorEmail() string {
	if o == nil || IsNil(o.CreatorEmail.Get()) {
		var ret string
		return ret
	}
	return *o.CreatorEmail.Get()
}

// GetCreatorEmailOk returns a tuple with the CreatorEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SettingFormulaModelHaljsonEmbeddedSetting) GetCreatorEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatorEmail.Get(), o.CreatorEmail.IsSet()
}

// HasCreatorEmail returns a boolean if a field has been set.
func (o *SettingFormulaModelHaljsonEmbeddedSetting) HasCreatorEmail() bool {
	if o != nil && o.CreatorEmail.IsSet() {
		return true
	}

	return false
}

// SetCreatorEmail gets a reference to the given NullableString and assigns it to the CreatorEmail field.
func (o *SettingFormulaModelHaljsonEmbeddedSetting) SetCreatorEmail(v string) {
	o.CreatorEmail.Set(&v)
}
// SetCreatorEmailNil sets the value for CreatorEmail to be an explicit nil
func (o *SettingFormulaModelHaljsonEmbeddedSetting) SetCreatorEmailNil() {
	o.CreatorEmail.Set(nil)
}

// UnsetCreatorEmail ensures that no value is present for CreatorEmail, not even an explicit nil
func (o *SettingFormulaModelHaljsonEmbeddedSetting) UnsetCreatorEmail() {
	o.CreatorEmail.Unset()
}

// GetCreatorFullName returns the CreatorFullName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SettingFormulaModelHaljsonEmbeddedSetting) GetCreatorFullName() string {
	if o == nil || IsNil(o.CreatorFullName.Get()) {
		var ret string
		return ret
	}
	return *o.CreatorFullName.Get()
}

// GetCreatorFullNameOk returns a tuple with the CreatorFullName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SettingFormulaModelHaljsonEmbeddedSetting) GetCreatorFullNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatorFullName.Get(), o.CreatorFullName.IsSet()
}

// HasCreatorFullName returns a boolean if a field has been set.
func (o *SettingFormulaModelHaljsonEmbeddedSetting) HasCreatorFullName() bool {
	if o != nil && o.CreatorFullName.IsSet() {
		return true
	}

	return false
}

// SetCreatorFullName gets a reference to the given NullableString and assigns it to the CreatorFullName field.
func (o *SettingFormulaModelHaljsonEmbeddedSetting) SetCreatorFullName(v string) {
	o.CreatorFullName.Set(&v)
}
// SetCreatorFullNameNil sets the value for CreatorFullName to be an explicit nil
func (o *SettingFormulaModelHaljsonEmbeddedSetting) SetCreatorFullNameNil() {
	o.CreatorFullName.Set(nil)
}

// UnsetCreatorFullName ensures that no value is present for CreatorFullName, not even an explicit nil
func (o *SettingFormulaModelHaljsonEmbeddedSetting) UnsetCreatorFullName() {
	o.CreatorFullName.Unset()
}

// GetIsWatching returns the IsWatching field value if set, zero value otherwise.
func (o *SettingFormulaModelHaljsonEmbeddedSetting) GetIsWatching() bool {
	if o == nil || IsNil(o.IsWatching) {
		var ret bool
		return ret
	}
	return *o.IsWatching
}

// GetIsWatchingOk returns a tuple with the IsWatching field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingFormulaModelHaljsonEmbeddedSetting) GetIsWatchingOk() (*bool, bool) {
	if o == nil || IsNil(o.IsWatching) {
		return nil, false
	}
	return o.IsWatching, true
}

// HasIsWatching returns a boolean if a field has been set.
func (o *SettingFormulaModelHaljsonEmbeddedSetting) HasIsWatching() bool {
	if o != nil && !IsNil(o.IsWatching) {
		return true
	}

	return false
}

// SetIsWatching gets a reference to the given bool and assigns it to the IsWatching field.
func (o *SettingFormulaModelHaljsonEmbeddedSetting) SetIsWatching(v bool) {
	o.IsWatching = &v
}

func (o SettingFormulaModelHaljsonEmbeddedSetting) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SettingFormulaModelHaljsonEmbeddedSetting) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.SettingType) {
		toSerialize["settingType"] = o.SettingType
	}
	if !IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	if o.CreatedAt.IsSet() {
		toSerialize["createdAt"] = o.CreatedAt.Get()
	}
	if o.CreatorEmail.IsSet() {
		toSerialize["creatorEmail"] = o.CreatorEmail.Get()
	}
	if o.CreatorFullName.IsSet() {
		toSerialize["creatorFullName"] = o.CreatorFullName.Get()
	}
	if !IsNil(o.IsWatching) {
		toSerialize["isWatching"] = o.IsWatching
	}
	return toSerialize, nil
}

type NullableSettingFormulaModelHaljsonEmbeddedSetting struct {
	value *SettingFormulaModelHaljsonEmbeddedSetting
	isSet bool
}

func (v NullableSettingFormulaModelHaljsonEmbeddedSetting) Get() *SettingFormulaModelHaljsonEmbeddedSetting {
	return v.value
}

func (v *NullableSettingFormulaModelHaljsonEmbeddedSetting) Set(val *SettingFormulaModelHaljsonEmbeddedSetting) {
	v.value = val
	v.isSet = true
}

func (v NullableSettingFormulaModelHaljsonEmbeddedSetting) IsSet() bool {
	return v.isSet
}

func (v *NullableSettingFormulaModelHaljsonEmbeddedSetting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSettingFormulaModelHaljsonEmbeddedSetting(val *SettingFormulaModelHaljsonEmbeddedSetting) *NullableSettingFormulaModelHaljsonEmbeddedSetting {
	return &NullableSettingFormulaModelHaljsonEmbeddedSetting{value: val, isSet: true}
}

func (v NullableSettingFormulaModelHaljsonEmbeddedSetting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSettingFormulaModelHaljsonEmbeddedSetting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


