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

// checks if the TargetingRuleModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TargetingRuleModel{}

// TargetingRuleModel struct for TargetingRuleModel
type TargetingRuleModel struct {
	// The list of conditions that are combined with logical AND operators.  It can be one of the following:  - User condition  - Segment condition  - Prerequisite flag condition
	Conditions []ConditionModel `json:"conditions,omitempty"`
	// The percentage options from where the evaluation process will choose a value based on the flag's percentage evaluation attribute.
	PercentageOptions []PercentageOptionModel `json:"percentageOptions,omitempty"`
	Value *ValueModel `json:"value,omitempty"`
}

// NewTargetingRuleModel instantiates a new TargetingRuleModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTargetingRuleModel() *TargetingRuleModel {
	this := TargetingRuleModel{}
	return &this
}

// NewTargetingRuleModelWithDefaults instantiates a new TargetingRuleModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTargetingRuleModelWithDefaults() *TargetingRuleModel {
	this := TargetingRuleModel{}
	return &this
}

// GetConditions returns the Conditions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TargetingRuleModel) GetConditions() []ConditionModel {
	if o == nil {
		var ret []ConditionModel
		return ret
	}
	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TargetingRuleModel) GetConditionsOk() ([]ConditionModel, bool) {
	if o == nil || IsNil(o.Conditions) {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *TargetingRuleModel) HasConditions() bool {
	if o != nil && IsNil(o.Conditions) {
		return true
	}

	return false
}

// SetConditions gets a reference to the given []ConditionModel and assigns it to the Conditions field.
func (o *TargetingRuleModel) SetConditions(v []ConditionModel) {
	o.Conditions = v
}

// GetPercentageOptions returns the PercentageOptions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TargetingRuleModel) GetPercentageOptions() []PercentageOptionModel {
	if o == nil {
		var ret []PercentageOptionModel
		return ret
	}
	return o.PercentageOptions
}

// GetPercentageOptionsOk returns a tuple with the PercentageOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TargetingRuleModel) GetPercentageOptionsOk() ([]PercentageOptionModel, bool) {
	if o == nil || IsNil(o.PercentageOptions) {
		return nil, false
	}
	return o.PercentageOptions, true
}

// HasPercentageOptions returns a boolean if a field has been set.
func (o *TargetingRuleModel) HasPercentageOptions() bool {
	if o != nil && IsNil(o.PercentageOptions) {
		return true
	}

	return false
}

// SetPercentageOptions gets a reference to the given []PercentageOptionModel and assigns it to the PercentageOptions field.
func (o *TargetingRuleModel) SetPercentageOptions(v []PercentageOptionModel) {
	o.PercentageOptions = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *TargetingRuleModel) GetValue() ValueModel {
	if o == nil || IsNil(o.Value) {
		var ret ValueModel
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetingRuleModel) GetValueOk() (*ValueModel, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *TargetingRuleModel) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given ValueModel and assigns it to the Value field.
func (o *TargetingRuleModel) SetValue(v ValueModel) {
	o.Value = &v
}

func (o TargetingRuleModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TargetingRuleModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Conditions != nil {
		toSerialize["conditions"] = o.Conditions
	}
	if o.PercentageOptions != nil {
		toSerialize["percentageOptions"] = o.PercentageOptions
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableTargetingRuleModel struct {
	value *TargetingRuleModel
	isSet bool
}

func (v NullableTargetingRuleModel) Get() *TargetingRuleModel {
	return v.value
}

func (v *NullableTargetingRuleModel) Set(val *TargetingRuleModel) {
	v.value = val
	v.isSet = true
}

func (v NullableTargetingRuleModel) IsSet() bool {
	return v.isSet
}

func (v *NullableTargetingRuleModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTargetingRuleModel(val *TargetingRuleModel) *NullableTargetingRuleModel {
	return &NullableTargetingRuleModel{value: val, isSet: true}
}

func (v NullableTargetingRuleModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTargetingRuleModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


