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

// checks if the UpdateSettingValueWithSettingIdModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateSettingValueWithSettingIdModel{}

// UpdateSettingValueWithSettingIdModel struct for UpdateSettingValueWithSettingIdModel
type UpdateSettingValueWithSettingIdModel struct {
	// The targeting rule collection.
	RolloutRules []RolloutRuleModel `json:"rolloutRules,omitempty"`
	// The percentage rule collection.
	RolloutPercentageItems []RolloutPercentageItemModel `json:"rolloutPercentageItems,omitempty"`
	// The value to serve. It must respect the setting type.
	Value interface{} `json:"value,omitempty"`
	// The id of the Setting.
	SettingId *int32 `json:"settingId,omitempty"`
}

// NewUpdateSettingValueWithSettingIdModel instantiates a new UpdateSettingValueWithSettingIdModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSettingValueWithSettingIdModel() *UpdateSettingValueWithSettingIdModel {
	this := UpdateSettingValueWithSettingIdModel{}
	return &this
}

// NewUpdateSettingValueWithSettingIdModelWithDefaults instantiates a new UpdateSettingValueWithSettingIdModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSettingValueWithSettingIdModelWithDefaults() *UpdateSettingValueWithSettingIdModel {
	this := UpdateSettingValueWithSettingIdModel{}
	return &this
}

// GetRolloutRules returns the RolloutRules field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateSettingValueWithSettingIdModel) GetRolloutRules() []RolloutRuleModel {
	if o == nil {
		var ret []RolloutRuleModel
		return ret
	}
	return o.RolloutRules
}

// GetRolloutRulesOk returns a tuple with the RolloutRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateSettingValueWithSettingIdModel) GetRolloutRulesOk() ([]RolloutRuleModel, bool) {
	if o == nil || IsNil(o.RolloutRules) {
		return nil, false
	}
	return o.RolloutRules, true
}

// HasRolloutRules returns a boolean if a field has been set.
func (o *UpdateSettingValueWithSettingIdModel) HasRolloutRules() bool {
	if o != nil && IsNil(o.RolloutRules) {
		return true
	}

	return false
}

// SetRolloutRules gets a reference to the given []RolloutRuleModel and assigns it to the RolloutRules field.
func (o *UpdateSettingValueWithSettingIdModel) SetRolloutRules(v []RolloutRuleModel) {
	o.RolloutRules = v
}

// GetRolloutPercentageItems returns the RolloutPercentageItems field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateSettingValueWithSettingIdModel) GetRolloutPercentageItems() []RolloutPercentageItemModel {
	if o == nil {
		var ret []RolloutPercentageItemModel
		return ret
	}
	return o.RolloutPercentageItems
}

// GetRolloutPercentageItemsOk returns a tuple with the RolloutPercentageItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateSettingValueWithSettingIdModel) GetRolloutPercentageItemsOk() ([]RolloutPercentageItemModel, bool) {
	if o == nil || IsNil(o.RolloutPercentageItems) {
		return nil, false
	}
	return o.RolloutPercentageItems, true
}

// HasRolloutPercentageItems returns a boolean if a field has been set.
func (o *UpdateSettingValueWithSettingIdModel) HasRolloutPercentageItems() bool {
	if o != nil && IsNil(o.RolloutPercentageItems) {
		return true
	}

	return false
}

// SetRolloutPercentageItems gets a reference to the given []RolloutPercentageItemModel and assigns it to the RolloutPercentageItems field.
func (o *UpdateSettingValueWithSettingIdModel) SetRolloutPercentageItems(v []RolloutPercentageItemModel) {
	o.RolloutPercentageItems = v
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateSettingValueWithSettingIdModel) GetValue() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateSettingValueWithSettingIdModel) GetValueOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return &o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *UpdateSettingValueWithSettingIdModel) HasValue() bool {
	if o != nil && IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given interface{} and assigns it to the Value field.
func (o *UpdateSettingValueWithSettingIdModel) SetValue(v interface{}) {
	o.Value = v
}

// GetSettingId returns the SettingId field value if set, zero value otherwise.
func (o *UpdateSettingValueWithSettingIdModel) GetSettingId() int32 {
	if o == nil || IsNil(o.SettingId) {
		var ret int32
		return ret
	}
	return *o.SettingId
}

// GetSettingIdOk returns a tuple with the SettingId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSettingValueWithSettingIdModel) GetSettingIdOk() (*int32, bool) {
	if o == nil || IsNil(o.SettingId) {
		return nil, false
	}
	return o.SettingId, true
}

// HasSettingId returns a boolean if a field has been set.
func (o *UpdateSettingValueWithSettingIdModel) HasSettingId() bool {
	if o != nil && !IsNil(o.SettingId) {
		return true
	}

	return false
}

// SetSettingId gets a reference to the given int32 and assigns it to the SettingId field.
func (o *UpdateSettingValueWithSettingIdModel) SetSettingId(v int32) {
	o.SettingId = &v
}

func (o UpdateSettingValueWithSettingIdModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateSettingValueWithSettingIdModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.RolloutRules != nil {
		toSerialize["rolloutRules"] = o.RolloutRules
	}
	if o.RolloutPercentageItems != nil {
		toSerialize["rolloutPercentageItems"] = o.RolloutPercentageItems
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.SettingId) {
		toSerialize["settingId"] = o.SettingId
	}
	return toSerialize, nil
}

type NullableUpdateSettingValueWithSettingIdModel struct {
	value *UpdateSettingValueWithSettingIdModel
	isSet bool
}

func (v NullableUpdateSettingValueWithSettingIdModel) Get() *UpdateSettingValueWithSettingIdModel {
	return v.value
}

func (v *NullableUpdateSettingValueWithSettingIdModel) Set(val *UpdateSettingValueWithSettingIdModel) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSettingValueWithSettingIdModel) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSettingValueWithSettingIdModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSettingValueWithSettingIdModel(val *UpdateSettingValueWithSettingIdModel) *NullableUpdateSettingValueWithSettingIdModel {
	return &NullableUpdateSettingValueWithSettingIdModel{value: val, isSet: true}
}

func (v NullableUpdateSettingValueWithSettingIdModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateSettingValueWithSettingIdModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


