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

// checks if the UserConditionModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserConditionModel{}

// UserConditionModel Describes a condition that is based on user attributes.
type UserConditionModel struct {
	// The user's attribute the evaluation process must take into account.
	ComparisonAttribute string `json:"comparisonAttribute"`
	Comparator UserComparator `json:"comparator"`
	ComparisonValue ComparisonValueModel `json:"comparisonValue"`
}

// NewUserConditionModel instantiates a new UserConditionModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserConditionModel(comparisonAttribute string, comparator UserComparator, comparisonValue ComparisonValueModel) *UserConditionModel {
	this := UserConditionModel{}
	this.ComparisonAttribute = comparisonAttribute
	this.Comparator = comparator
	this.ComparisonValue = comparisonValue
	return &this
}

// NewUserConditionModelWithDefaults instantiates a new UserConditionModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserConditionModelWithDefaults() *UserConditionModel {
	this := UserConditionModel{}
	return &this
}

// GetComparisonAttribute returns the ComparisonAttribute field value
func (o *UserConditionModel) GetComparisonAttribute() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ComparisonAttribute
}

// GetComparisonAttributeOk returns a tuple with the ComparisonAttribute field value
// and a boolean to check if the value has been set.
func (o *UserConditionModel) GetComparisonAttributeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ComparisonAttribute, true
}

// SetComparisonAttribute sets field value
func (o *UserConditionModel) SetComparisonAttribute(v string) {
	o.ComparisonAttribute = v
}

// GetComparator returns the Comparator field value
func (o *UserConditionModel) GetComparator() UserComparator {
	if o == nil {
		var ret UserComparator
		return ret
	}

	return o.Comparator
}

// GetComparatorOk returns a tuple with the Comparator field value
// and a boolean to check if the value has been set.
func (o *UserConditionModel) GetComparatorOk() (*UserComparator, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Comparator, true
}

// SetComparator sets field value
func (o *UserConditionModel) SetComparator(v UserComparator) {
	o.Comparator = v
}

// GetComparisonValue returns the ComparisonValue field value
func (o *UserConditionModel) GetComparisonValue() ComparisonValueModel {
	if o == nil {
		var ret ComparisonValueModel
		return ret
	}

	return o.ComparisonValue
}

// GetComparisonValueOk returns a tuple with the ComparisonValue field value
// and a boolean to check if the value has been set.
func (o *UserConditionModel) GetComparisonValueOk() (*ComparisonValueModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ComparisonValue, true
}

// SetComparisonValue sets field value
func (o *UserConditionModel) SetComparisonValue(v ComparisonValueModel) {
	o.ComparisonValue = v
}

func (o UserConditionModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserConditionModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["comparisonAttribute"] = o.ComparisonAttribute
	toSerialize["comparator"] = o.Comparator
	toSerialize["comparisonValue"] = o.ComparisonValue
	return toSerialize, nil
}

type NullableUserConditionModel struct {
	value *UserConditionModel
	isSet bool
}

func (v NullableUserConditionModel) Get() *UserConditionModel {
	return v.value
}

func (v *NullableUserConditionModel) Set(val *UserConditionModel) {
	v.value = val
	v.isSet = true
}

func (v NullableUserConditionModel) IsSet() bool {
	return v.isSet
}

func (v *NullableUserConditionModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserConditionModel(val *UserConditionModel) *NullableUserConditionModel {
	return &NullableUserConditionModel{value: val, isSet: true}
}

func (v NullableUserConditionModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserConditionModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

