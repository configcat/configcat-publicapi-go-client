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

// checks if the UpdateEvaluationFormulaModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateEvaluationFormulaModel{}

// UpdateEvaluationFormulaModel struct for UpdateEvaluationFormulaModel
type UpdateEvaluationFormulaModel struct {
	DefaultValue ValueModel `json:"defaultValue"`
	// The targeting rules of the Feature Flag or Setting.
	TargetingRules []TargetingRuleModel `json:"targetingRules,omitempty"`
	// The user attribute used for percentage evaluation. If not set, it defaults to the `Identifier` user object attribute.
	PercentageEvaluationAttribute NullableString `json:"percentageEvaluationAttribute,omitempty"`
}

// NewUpdateEvaluationFormulaModel instantiates a new UpdateEvaluationFormulaModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateEvaluationFormulaModel(defaultValue ValueModel) *UpdateEvaluationFormulaModel {
	this := UpdateEvaluationFormulaModel{}
	this.DefaultValue = defaultValue
	return &this
}

// NewUpdateEvaluationFormulaModelWithDefaults instantiates a new UpdateEvaluationFormulaModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateEvaluationFormulaModelWithDefaults() *UpdateEvaluationFormulaModel {
	this := UpdateEvaluationFormulaModel{}
	return &this
}

// GetDefaultValue returns the DefaultValue field value
func (o *UpdateEvaluationFormulaModel) GetDefaultValue() ValueModel {
	if o == nil {
		var ret ValueModel
		return ret
	}

	return o.DefaultValue
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value
// and a boolean to check if the value has been set.
func (o *UpdateEvaluationFormulaModel) GetDefaultValueOk() (*ValueModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultValue, true
}

// SetDefaultValue sets field value
func (o *UpdateEvaluationFormulaModel) SetDefaultValue(v ValueModel) {
	o.DefaultValue = v
}

// GetTargetingRules returns the TargetingRules field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateEvaluationFormulaModel) GetTargetingRules() []TargetingRuleModel {
	if o == nil {
		var ret []TargetingRuleModel
		return ret
	}
	return o.TargetingRules
}

// GetTargetingRulesOk returns a tuple with the TargetingRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateEvaluationFormulaModel) GetTargetingRulesOk() ([]TargetingRuleModel, bool) {
	if o == nil || IsNil(o.TargetingRules) {
		return nil, false
	}
	return o.TargetingRules, true
}

// HasTargetingRules returns a boolean if a field has been set.
func (o *UpdateEvaluationFormulaModel) HasTargetingRules() bool {
	if o != nil && IsNil(o.TargetingRules) {
		return true
	}

	return false
}

// SetTargetingRules gets a reference to the given []TargetingRuleModel and assigns it to the TargetingRules field.
func (o *UpdateEvaluationFormulaModel) SetTargetingRules(v []TargetingRuleModel) {
	o.TargetingRules = v
}

// GetPercentageEvaluationAttribute returns the PercentageEvaluationAttribute field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateEvaluationFormulaModel) GetPercentageEvaluationAttribute() string {
	if o == nil || IsNil(o.PercentageEvaluationAttribute.Get()) {
		var ret string
		return ret
	}
	return *o.PercentageEvaluationAttribute.Get()
}

// GetPercentageEvaluationAttributeOk returns a tuple with the PercentageEvaluationAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateEvaluationFormulaModel) GetPercentageEvaluationAttributeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PercentageEvaluationAttribute.Get(), o.PercentageEvaluationAttribute.IsSet()
}

// HasPercentageEvaluationAttribute returns a boolean if a field has been set.
func (o *UpdateEvaluationFormulaModel) HasPercentageEvaluationAttribute() bool {
	if o != nil && o.PercentageEvaluationAttribute.IsSet() {
		return true
	}

	return false
}

// SetPercentageEvaluationAttribute gets a reference to the given NullableString and assigns it to the PercentageEvaluationAttribute field.
func (o *UpdateEvaluationFormulaModel) SetPercentageEvaluationAttribute(v string) {
	o.PercentageEvaluationAttribute.Set(&v)
}
// SetPercentageEvaluationAttributeNil sets the value for PercentageEvaluationAttribute to be an explicit nil
func (o *UpdateEvaluationFormulaModel) SetPercentageEvaluationAttributeNil() {
	o.PercentageEvaluationAttribute.Set(nil)
}

// UnsetPercentageEvaluationAttribute ensures that no value is present for PercentageEvaluationAttribute, not even an explicit nil
func (o *UpdateEvaluationFormulaModel) UnsetPercentageEvaluationAttribute() {
	o.PercentageEvaluationAttribute.Unset()
}

func (o UpdateEvaluationFormulaModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateEvaluationFormulaModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["defaultValue"] = o.DefaultValue
	if o.TargetingRules != nil {
		toSerialize["targetingRules"] = o.TargetingRules
	}
	if o.PercentageEvaluationAttribute.IsSet() {
		toSerialize["percentageEvaluationAttribute"] = o.PercentageEvaluationAttribute.Get()
	}
	return toSerialize, nil
}

type NullableUpdateEvaluationFormulaModel struct {
	value *UpdateEvaluationFormulaModel
	isSet bool
}

func (v NullableUpdateEvaluationFormulaModel) Get() *UpdateEvaluationFormulaModel {
	return v.value
}

func (v *NullableUpdateEvaluationFormulaModel) Set(val *UpdateEvaluationFormulaModel) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateEvaluationFormulaModel) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateEvaluationFormulaModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateEvaluationFormulaModel(val *UpdateEvaluationFormulaModel) *NullableUpdateEvaluationFormulaModel {
	return &NullableUpdateEvaluationFormulaModel{value: val, isSet: true}
}

func (v NullableUpdateEvaluationFormulaModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateEvaluationFormulaModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


