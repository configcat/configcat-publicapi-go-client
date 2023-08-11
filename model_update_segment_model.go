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
)

// checks if the UpdateSegmentModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateSegmentModel{}

// UpdateSegmentModel struct for UpdateSegmentModel
type UpdateSegmentModel struct {
	Name NullableString `json:"name,omitempty"`
	Description NullableString `json:"description,omitempty"`
	ComparisonAttribute NullableString `json:"comparisonAttribute,omitempty"`
	Comparator *RolloutRuleComparator `json:"comparator,omitempty"`
	ComparisonValue NullableString `json:"comparisonValue,omitempty"`
}

// NewUpdateSegmentModel instantiates a new UpdateSegmentModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSegmentModel() *UpdateSegmentModel {
	this := UpdateSegmentModel{}
	return &this
}

// NewUpdateSegmentModelWithDefaults instantiates a new UpdateSegmentModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSegmentModelWithDefaults() *UpdateSegmentModel {
	this := UpdateSegmentModel{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateSegmentModel) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateSegmentModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *UpdateSegmentModel) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *UpdateSegmentModel) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *UpdateSegmentModel) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *UpdateSegmentModel) UnsetName() {
	o.Name.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateSegmentModel) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateSegmentModel) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateSegmentModel) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *UpdateSegmentModel) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *UpdateSegmentModel) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *UpdateSegmentModel) UnsetDescription() {
	o.Description.Unset()
}

// GetComparisonAttribute returns the ComparisonAttribute field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateSegmentModel) GetComparisonAttribute() string {
	if o == nil || IsNil(o.ComparisonAttribute.Get()) {
		var ret string
		return ret
	}
	return *o.ComparisonAttribute.Get()
}

// GetComparisonAttributeOk returns a tuple with the ComparisonAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateSegmentModel) GetComparisonAttributeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ComparisonAttribute.Get(), o.ComparisonAttribute.IsSet()
}

// HasComparisonAttribute returns a boolean if a field has been set.
func (o *UpdateSegmentModel) HasComparisonAttribute() bool {
	if o != nil && o.ComparisonAttribute.IsSet() {
		return true
	}

	return false
}

// SetComparisonAttribute gets a reference to the given NullableString and assigns it to the ComparisonAttribute field.
func (o *UpdateSegmentModel) SetComparisonAttribute(v string) {
	o.ComparisonAttribute.Set(&v)
}
// SetComparisonAttributeNil sets the value for ComparisonAttribute to be an explicit nil
func (o *UpdateSegmentModel) SetComparisonAttributeNil() {
	o.ComparisonAttribute.Set(nil)
}

// UnsetComparisonAttribute ensures that no value is present for ComparisonAttribute, not even an explicit nil
func (o *UpdateSegmentModel) UnsetComparisonAttribute() {
	o.ComparisonAttribute.Unset()
}

// GetComparator returns the Comparator field value if set, zero value otherwise.
func (o *UpdateSegmentModel) GetComparator() RolloutRuleComparator {
	if o == nil || IsNil(o.Comparator) {
		var ret RolloutRuleComparator
		return ret
	}
	return *o.Comparator
}

// GetComparatorOk returns a tuple with the Comparator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSegmentModel) GetComparatorOk() (*RolloutRuleComparator, bool) {
	if o == nil || IsNil(o.Comparator) {
		return nil, false
	}
	return o.Comparator, true
}

// HasComparator returns a boolean if a field has been set.
func (o *UpdateSegmentModel) HasComparator() bool {
	if o != nil && !IsNil(o.Comparator) {
		return true
	}

	return false
}

// SetComparator gets a reference to the given RolloutRuleComparator and assigns it to the Comparator field.
func (o *UpdateSegmentModel) SetComparator(v RolloutRuleComparator) {
	o.Comparator = &v
}

// GetComparisonValue returns the ComparisonValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateSegmentModel) GetComparisonValue() string {
	if o == nil || IsNil(o.ComparisonValue.Get()) {
		var ret string
		return ret
	}
	return *o.ComparisonValue.Get()
}

// GetComparisonValueOk returns a tuple with the ComparisonValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateSegmentModel) GetComparisonValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ComparisonValue.Get(), o.ComparisonValue.IsSet()
}

// HasComparisonValue returns a boolean if a field has been set.
func (o *UpdateSegmentModel) HasComparisonValue() bool {
	if o != nil && o.ComparisonValue.IsSet() {
		return true
	}

	return false
}

// SetComparisonValue gets a reference to the given NullableString and assigns it to the ComparisonValue field.
func (o *UpdateSegmentModel) SetComparisonValue(v string) {
	o.ComparisonValue.Set(&v)
}
// SetComparisonValueNil sets the value for ComparisonValue to be an explicit nil
func (o *UpdateSegmentModel) SetComparisonValueNil() {
	o.ComparisonValue.Set(nil)
}

// UnsetComparisonValue ensures that no value is present for ComparisonValue, not even an explicit nil
func (o *UpdateSegmentModel) UnsetComparisonValue() {
	o.ComparisonValue.Unset()
}

func (o UpdateSegmentModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateSegmentModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.ComparisonAttribute.IsSet() {
		toSerialize["comparisonAttribute"] = o.ComparisonAttribute.Get()
	}
	if !IsNil(o.Comparator) {
		toSerialize["comparator"] = o.Comparator
	}
	if o.ComparisonValue.IsSet() {
		toSerialize["comparisonValue"] = o.ComparisonValue.Get()
	}
	return toSerialize, nil
}

type NullableUpdateSegmentModel struct {
	value *UpdateSegmentModel
	isSet bool
}

func (v NullableUpdateSegmentModel) Get() *UpdateSegmentModel {
	return v.value
}

func (v *NullableUpdateSegmentModel) Set(val *UpdateSegmentModel) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSegmentModel) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSegmentModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSegmentModel(val *UpdateSegmentModel) *NullableUpdateSegmentModel {
	return &NullableUpdateSegmentModel{value: val, isSet: true}
}

func (v NullableUpdateSegmentModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateSegmentModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


