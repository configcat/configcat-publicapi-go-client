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

// checks if the SettingFormulaModelHaljsonEmbeddedIntegrationLinksInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SettingFormulaModelHaljsonEmbeddedIntegrationLinksInner{}

// SettingFormulaModelHaljsonEmbeddedIntegrationLinksInner struct for SettingFormulaModelHaljsonEmbeddedIntegrationLinksInner
type SettingFormulaModelHaljsonEmbeddedIntegrationLinksInner struct {
	Key NullableString `json:"key,omitempty"`
	Description NullableString `json:"description,omitempty"`
	IntegrationLinkType *IntegrationLinkType `json:"integrationLinkType,omitempty"`
	Url NullableString `json:"url,omitempty"`
}

// NewSettingFormulaModelHaljsonEmbeddedIntegrationLinksInner instantiates a new SettingFormulaModelHaljsonEmbeddedIntegrationLinksInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSettingFormulaModelHaljsonEmbeddedIntegrationLinksInner() *SettingFormulaModelHaljsonEmbeddedIntegrationLinksInner {
	this := SettingFormulaModelHaljsonEmbeddedIntegrationLinksInner{}
	return &this
}

// NewSettingFormulaModelHaljsonEmbeddedIntegrationLinksInnerWithDefaults instantiates a new SettingFormulaModelHaljsonEmbeddedIntegrationLinksInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSettingFormulaModelHaljsonEmbeddedIntegrationLinksInnerWithDefaults() *SettingFormulaModelHaljsonEmbeddedIntegrationLinksInner {
	this := SettingFormulaModelHaljsonEmbeddedIntegrationLinksInner{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SettingFormulaModelHaljsonEmbeddedIntegrationLinksInner) GetKey() string {
	if o == nil || IsNil(o.Key.Get()) {
		var ret string
		return ret
	}
	return *o.Key.Get()
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SettingFormulaModelHaljsonEmbeddedIntegrationLinksInner) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Key.Get(), o.Key.IsSet()
}

// HasKey returns a boolean if a field has been set.
func (o *SettingFormulaModelHaljsonEmbeddedIntegrationLinksInner) HasKey() bool {
	if o != nil && o.Key.IsSet() {
		return true
	}

	return false
}

// SetKey gets a reference to the given NullableString and assigns it to the Key field.
func (o *SettingFormulaModelHaljsonEmbeddedIntegrationLinksInner) SetKey(v string) {
	o.Key.Set(&v)
}
// SetKeyNil sets the value for Key to be an explicit nil
func (o *SettingFormulaModelHaljsonEmbeddedIntegrationLinksInner) SetKeyNil() {
	o.Key.Set(nil)
}

// UnsetKey ensures that no value is present for Key, not even an explicit nil
func (o *SettingFormulaModelHaljsonEmbeddedIntegrationLinksInner) UnsetKey() {
	o.Key.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SettingFormulaModelHaljsonEmbeddedIntegrationLinksInner) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SettingFormulaModelHaljsonEmbeddedIntegrationLinksInner) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *SettingFormulaModelHaljsonEmbeddedIntegrationLinksInner) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *SettingFormulaModelHaljsonEmbeddedIntegrationLinksInner) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *SettingFormulaModelHaljsonEmbeddedIntegrationLinksInner) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *SettingFormulaModelHaljsonEmbeddedIntegrationLinksInner) UnsetDescription() {
	o.Description.Unset()
}

// GetIntegrationLinkType returns the IntegrationLinkType field value if set, zero value otherwise.
func (o *SettingFormulaModelHaljsonEmbeddedIntegrationLinksInner) GetIntegrationLinkType() IntegrationLinkType {
	if o == nil || IsNil(o.IntegrationLinkType) {
		var ret IntegrationLinkType
		return ret
	}
	return *o.IntegrationLinkType
}

// GetIntegrationLinkTypeOk returns a tuple with the IntegrationLinkType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingFormulaModelHaljsonEmbeddedIntegrationLinksInner) GetIntegrationLinkTypeOk() (*IntegrationLinkType, bool) {
	if o == nil || IsNil(o.IntegrationLinkType) {
		return nil, false
	}
	return o.IntegrationLinkType, true
}

// HasIntegrationLinkType returns a boolean if a field has been set.
func (o *SettingFormulaModelHaljsonEmbeddedIntegrationLinksInner) HasIntegrationLinkType() bool {
	if o != nil && !IsNil(o.IntegrationLinkType) {
		return true
	}

	return false
}

// SetIntegrationLinkType gets a reference to the given IntegrationLinkType and assigns it to the IntegrationLinkType field.
func (o *SettingFormulaModelHaljsonEmbeddedIntegrationLinksInner) SetIntegrationLinkType(v IntegrationLinkType) {
	o.IntegrationLinkType = &v
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SettingFormulaModelHaljsonEmbeddedIntegrationLinksInner) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SettingFormulaModelHaljsonEmbeddedIntegrationLinksInner) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *SettingFormulaModelHaljsonEmbeddedIntegrationLinksInner) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *SettingFormulaModelHaljsonEmbeddedIntegrationLinksInner) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *SettingFormulaModelHaljsonEmbeddedIntegrationLinksInner) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *SettingFormulaModelHaljsonEmbeddedIntegrationLinksInner) UnsetUrl() {
	o.Url.Unset()
}

func (o SettingFormulaModelHaljsonEmbeddedIntegrationLinksInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SettingFormulaModelHaljsonEmbeddedIntegrationLinksInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Key.IsSet() {
		toSerialize["key"] = o.Key.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if !IsNil(o.IntegrationLinkType) {
		toSerialize["integrationLinkType"] = o.IntegrationLinkType
	}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	return toSerialize, nil
}

type NullableSettingFormulaModelHaljsonEmbeddedIntegrationLinksInner struct {
	value *SettingFormulaModelHaljsonEmbeddedIntegrationLinksInner
	isSet bool
}

func (v NullableSettingFormulaModelHaljsonEmbeddedIntegrationLinksInner) Get() *SettingFormulaModelHaljsonEmbeddedIntegrationLinksInner {
	return v.value
}

func (v *NullableSettingFormulaModelHaljsonEmbeddedIntegrationLinksInner) Set(val *SettingFormulaModelHaljsonEmbeddedIntegrationLinksInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSettingFormulaModelHaljsonEmbeddedIntegrationLinksInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSettingFormulaModelHaljsonEmbeddedIntegrationLinksInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSettingFormulaModelHaljsonEmbeddedIntegrationLinksInner(val *SettingFormulaModelHaljsonEmbeddedIntegrationLinksInner) *NullableSettingFormulaModelHaljsonEmbeddedIntegrationLinksInner {
	return &NullableSettingFormulaModelHaljsonEmbeddedIntegrationLinksInner{value: val, isSet: true}
}

func (v NullableSettingFormulaModelHaljsonEmbeddedIntegrationLinksInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSettingFormulaModelHaljsonEmbeddedIntegrationLinksInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

