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

// checks if the CreateSegmentModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSegmentModel{}

// CreateSegmentModel struct for CreateSegmentModel
type CreateSegmentModel struct {
	Name string `json:"name"`
	Description NullableString `json:"description,omitempty"`
	ComparisonAttribute string `json:"comparisonAttribute"`
	Comparator RolloutRuleComparator `json:"comparator"`
	ComparisonValue string `json:"comparisonValue"`
}

// NewCreateSegmentModel instantiates a new CreateSegmentModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSegmentModel(name string, comparisonAttribute string, comparator RolloutRuleComparator, comparisonValue string) *CreateSegmentModel {
	this := CreateSegmentModel{}
	this.Name = name
	this.ComparisonAttribute = comparisonAttribute
	this.Comparator = comparator
	this.ComparisonValue = comparisonValue
	return &this
}

// NewCreateSegmentModelWithDefaults instantiates a new CreateSegmentModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSegmentModelWithDefaults() *CreateSegmentModel {
	this := CreateSegmentModel{}
	return &this
}

// GetName returns the Name field value
func (o *CreateSegmentModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateSegmentModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateSegmentModel) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateSegmentModel) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateSegmentModel) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateSegmentModel) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *CreateSegmentModel) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *CreateSegmentModel) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *CreateSegmentModel) UnsetDescription() {
	o.Description.Unset()
}

// GetComparisonAttribute returns the ComparisonAttribute field value
func (o *CreateSegmentModel) GetComparisonAttribute() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ComparisonAttribute
}

// GetComparisonAttributeOk returns a tuple with the ComparisonAttribute field value
// and a boolean to check if the value has been set.
func (o *CreateSegmentModel) GetComparisonAttributeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ComparisonAttribute, true
}

// SetComparisonAttribute sets field value
func (o *CreateSegmentModel) SetComparisonAttribute(v string) {
	o.ComparisonAttribute = v
}

// GetComparator returns the Comparator field value
func (o *CreateSegmentModel) GetComparator() RolloutRuleComparator {
	if o == nil {
		var ret RolloutRuleComparator
		return ret
	}

	return o.Comparator
}

// GetComparatorOk returns a tuple with the Comparator field value
// and a boolean to check if the value has been set.
func (o *CreateSegmentModel) GetComparatorOk() (*RolloutRuleComparator, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Comparator, true
}

// SetComparator sets field value
func (o *CreateSegmentModel) SetComparator(v RolloutRuleComparator) {
	o.Comparator = v
}

// GetComparisonValue returns the ComparisonValue field value
func (o *CreateSegmentModel) GetComparisonValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ComparisonValue
}

// GetComparisonValueOk returns a tuple with the ComparisonValue field value
// and a boolean to check if the value has been set.
func (o *CreateSegmentModel) GetComparisonValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ComparisonValue, true
}

// SetComparisonValue sets field value
func (o *CreateSegmentModel) SetComparisonValue(v string) {
	o.ComparisonValue = v
}

func (o CreateSegmentModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateSegmentModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	toSerialize["comparisonAttribute"] = o.ComparisonAttribute
	toSerialize["comparator"] = o.Comparator
	toSerialize["comparisonValue"] = o.ComparisonValue
	return toSerialize, nil
}

type NullableCreateSegmentModel struct {
	value *CreateSegmentModel
	isSet bool
}

func (v NullableCreateSegmentModel) Get() *CreateSegmentModel {
	return v.value
}

func (v *NullableCreateSegmentModel) Set(val *CreateSegmentModel) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSegmentModel) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSegmentModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSegmentModel(val *CreateSegmentModel) *NullableCreateSegmentModel {
	return &NullableCreateSegmentModel{value: val, isSet: true}
}

func (v NullableCreateSegmentModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSegmentModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


