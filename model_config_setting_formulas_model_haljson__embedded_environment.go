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

// checks if the ConfigSettingFormulasModelHaljsonEmbeddedEnvironment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigSettingFormulasModelHaljsonEmbeddedEnvironment{}

// ConfigSettingFormulasModelHaljsonEmbeddedEnvironment struct for ConfigSettingFormulasModelHaljsonEmbeddedEnvironment
type ConfigSettingFormulasModelHaljsonEmbeddedEnvironment struct {
	Embedded *ConfigModelHaljsonEmbedded `json:"_embedded,omitempty"`
	// Identifier of the Environment.
	EnvironmentId *string `json:"environmentId,omitempty"`
	// Name of the Environment.
	Name NullableString `json:"name,omitempty"`
	// The configured color of the Environment.
	Color NullableString `json:"color,omitempty"`
	// Description of the Environment.
	Description NullableString `json:"description,omitempty"`
	// The order of the Environment represented on the ConfigCat Dashboard.
	Order *int32 `json:"order,omitempty"`
	// Determines whether a mandatory reason must be given every time when the Feature Flags or Settings in the given Environment are saved.
	ReasonRequired *bool `json:"reasonRequired,omitempty"`
	Links *ConfigSettingFormulasModelHaljsonEmbeddedEnvironmentLinks `json:"_links,omitempty"`
}

// NewConfigSettingFormulasModelHaljsonEmbeddedEnvironment instantiates a new ConfigSettingFormulasModelHaljsonEmbeddedEnvironment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigSettingFormulasModelHaljsonEmbeddedEnvironment() *ConfigSettingFormulasModelHaljsonEmbeddedEnvironment {
	this := ConfigSettingFormulasModelHaljsonEmbeddedEnvironment{}
	return &this
}

// NewConfigSettingFormulasModelHaljsonEmbeddedEnvironmentWithDefaults instantiates a new ConfigSettingFormulasModelHaljsonEmbeddedEnvironment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigSettingFormulasModelHaljsonEmbeddedEnvironmentWithDefaults() *ConfigSettingFormulasModelHaljsonEmbeddedEnvironment {
	this := ConfigSettingFormulasModelHaljsonEmbeddedEnvironment{}
	return &this
}

// GetEmbedded returns the Embedded field value if set, zero value otherwise.
func (o *ConfigSettingFormulasModelHaljsonEmbeddedEnvironment) GetEmbedded() ConfigModelHaljsonEmbedded {
	if o == nil || IsNil(o.Embedded) {
		var ret ConfigModelHaljsonEmbedded
		return ret
	}
	return *o.Embedded
}

// GetEmbeddedOk returns a tuple with the Embedded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigSettingFormulasModelHaljsonEmbeddedEnvironment) GetEmbeddedOk() (*ConfigModelHaljsonEmbedded, bool) {
	if o == nil || IsNil(o.Embedded) {
		return nil, false
	}
	return o.Embedded, true
}

// HasEmbedded returns a boolean if a field has been set.
func (o *ConfigSettingFormulasModelHaljsonEmbeddedEnvironment) HasEmbedded() bool {
	if o != nil && !IsNil(o.Embedded) {
		return true
	}

	return false
}

// SetEmbedded gets a reference to the given ConfigModelHaljsonEmbedded and assigns it to the Embedded field.
func (o *ConfigSettingFormulasModelHaljsonEmbeddedEnvironment) SetEmbedded(v ConfigModelHaljsonEmbedded) {
	o.Embedded = &v
}

// GetEnvironmentId returns the EnvironmentId field value if set, zero value otherwise.
func (o *ConfigSettingFormulasModelHaljsonEmbeddedEnvironment) GetEnvironmentId() string {
	if o == nil || IsNil(o.EnvironmentId) {
		var ret string
		return ret
	}
	return *o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigSettingFormulasModelHaljsonEmbeddedEnvironment) GetEnvironmentIdOk() (*string, bool) {
	if o == nil || IsNil(o.EnvironmentId) {
		return nil, false
	}
	return o.EnvironmentId, true
}

// HasEnvironmentId returns a boolean if a field has been set.
func (o *ConfigSettingFormulasModelHaljsonEmbeddedEnvironment) HasEnvironmentId() bool {
	if o != nil && !IsNil(o.EnvironmentId) {
		return true
	}

	return false
}

// SetEnvironmentId gets a reference to the given string and assigns it to the EnvironmentId field.
func (o *ConfigSettingFormulasModelHaljsonEmbeddedEnvironment) SetEnvironmentId(v string) {
	o.EnvironmentId = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConfigSettingFormulasModelHaljsonEmbeddedEnvironment) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConfigSettingFormulasModelHaljsonEmbeddedEnvironment) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ConfigSettingFormulasModelHaljsonEmbeddedEnvironment) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ConfigSettingFormulasModelHaljsonEmbeddedEnvironment) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *ConfigSettingFormulasModelHaljsonEmbeddedEnvironment) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ConfigSettingFormulasModelHaljsonEmbeddedEnvironment) UnsetName() {
	o.Name.Unset()
}

// GetColor returns the Color field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConfigSettingFormulasModelHaljsonEmbeddedEnvironment) GetColor() string {
	if o == nil || IsNil(o.Color.Get()) {
		var ret string
		return ret
	}
	return *o.Color.Get()
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConfigSettingFormulasModelHaljsonEmbeddedEnvironment) GetColorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Color.Get(), o.Color.IsSet()
}

// HasColor returns a boolean if a field has been set.
func (o *ConfigSettingFormulasModelHaljsonEmbeddedEnvironment) HasColor() bool {
	if o != nil && o.Color.IsSet() {
		return true
	}

	return false
}

// SetColor gets a reference to the given NullableString and assigns it to the Color field.
func (o *ConfigSettingFormulasModelHaljsonEmbeddedEnvironment) SetColor(v string) {
	o.Color.Set(&v)
}
// SetColorNil sets the value for Color to be an explicit nil
func (o *ConfigSettingFormulasModelHaljsonEmbeddedEnvironment) SetColorNil() {
	o.Color.Set(nil)
}

// UnsetColor ensures that no value is present for Color, not even an explicit nil
func (o *ConfigSettingFormulasModelHaljsonEmbeddedEnvironment) UnsetColor() {
	o.Color.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConfigSettingFormulasModelHaljsonEmbeddedEnvironment) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConfigSettingFormulasModelHaljsonEmbeddedEnvironment) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ConfigSettingFormulasModelHaljsonEmbeddedEnvironment) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *ConfigSettingFormulasModelHaljsonEmbeddedEnvironment) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *ConfigSettingFormulasModelHaljsonEmbeddedEnvironment) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *ConfigSettingFormulasModelHaljsonEmbeddedEnvironment) UnsetDescription() {
	o.Description.Unset()
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *ConfigSettingFormulasModelHaljsonEmbeddedEnvironment) GetOrder() int32 {
	if o == nil || IsNil(o.Order) {
		var ret int32
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigSettingFormulasModelHaljsonEmbeddedEnvironment) GetOrderOk() (*int32, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *ConfigSettingFormulasModelHaljsonEmbeddedEnvironment) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given int32 and assigns it to the Order field.
func (o *ConfigSettingFormulasModelHaljsonEmbeddedEnvironment) SetOrder(v int32) {
	o.Order = &v
}

// GetReasonRequired returns the ReasonRequired field value if set, zero value otherwise.
func (o *ConfigSettingFormulasModelHaljsonEmbeddedEnvironment) GetReasonRequired() bool {
	if o == nil || IsNil(o.ReasonRequired) {
		var ret bool
		return ret
	}
	return *o.ReasonRequired
}

// GetReasonRequiredOk returns a tuple with the ReasonRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigSettingFormulasModelHaljsonEmbeddedEnvironment) GetReasonRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.ReasonRequired) {
		return nil, false
	}
	return o.ReasonRequired, true
}

// HasReasonRequired returns a boolean if a field has been set.
func (o *ConfigSettingFormulasModelHaljsonEmbeddedEnvironment) HasReasonRequired() bool {
	if o != nil && !IsNil(o.ReasonRequired) {
		return true
	}

	return false
}

// SetReasonRequired gets a reference to the given bool and assigns it to the ReasonRequired field.
func (o *ConfigSettingFormulasModelHaljsonEmbeddedEnvironment) SetReasonRequired(v bool) {
	o.ReasonRequired = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ConfigSettingFormulasModelHaljsonEmbeddedEnvironment) GetLinks() ConfigSettingFormulasModelHaljsonEmbeddedEnvironmentLinks {
	if o == nil || IsNil(o.Links) {
		var ret ConfigSettingFormulasModelHaljsonEmbeddedEnvironmentLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigSettingFormulasModelHaljsonEmbeddedEnvironment) GetLinksOk() (*ConfigSettingFormulasModelHaljsonEmbeddedEnvironmentLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ConfigSettingFormulasModelHaljsonEmbeddedEnvironment) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ConfigSettingFormulasModelHaljsonEmbeddedEnvironmentLinks and assigns it to the Links field.
func (o *ConfigSettingFormulasModelHaljsonEmbeddedEnvironment) SetLinks(v ConfigSettingFormulasModelHaljsonEmbeddedEnvironmentLinks) {
	o.Links = &v
}

func (o ConfigSettingFormulasModelHaljsonEmbeddedEnvironment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigSettingFormulasModelHaljsonEmbeddedEnvironment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Embedded) {
		toSerialize["_embedded"] = o.Embedded
	}
	if !IsNil(o.EnvironmentId) {
		toSerialize["environmentId"] = o.EnvironmentId
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Color.IsSet() {
		toSerialize["color"] = o.Color.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if !IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	if !IsNil(o.ReasonRequired) {
		toSerialize["reasonRequired"] = o.ReasonRequired
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableConfigSettingFormulasModelHaljsonEmbeddedEnvironment struct {
	value *ConfigSettingFormulasModelHaljsonEmbeddedEnvironment
	isSet bool
}

func (v NullableConfigSettingFormulasModelHaljsonEmbeddedEnvironment) Get() *ConfigSettingFormulasModelHaljsonEmbeddedEnvironment {
	return v.value
}

func (v *NullableConfigSettingFormulasModelHaljsonEmbeddedEnvironment) Set(val *ConfigSettingFormulasModelHaljsonEmbeddedEnvironment) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigSettingFormulasModelHaljsonEmbeddedEnvironment) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigSettingFormulasModelHaljsonEmbeddedEnvironment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigSettingFormulasModelHaljsonEmbeddedEnvironment(val *ConfigSettingFormulasModelHaljsonEmbeddedEnvironment) *NullableConfigSettingFormulasModelHaljsonEmbeddedEnvironment {
	return &NullableConfigSettingFormulasModelHaljsonEmbeddedEnvironment{value: val, isSet: true}
}

func (v NullableConfigSettingFormulasModelHaljsonEmbeddedEnvironment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigSettingFormulasModelHaljsonEmbeddedEnvironment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


