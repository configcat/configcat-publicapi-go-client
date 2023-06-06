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

// checks if the SettingValueModelHaljsonEmbeddedConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SettingValueModelHaljsonEmbeddedConfig{}

// SettingValueModelHaljsonEmbeddedConfig struct for SettingValueModelHaljsonEmbeddedConfig
type SettingValueModelHaljsonEmbeddedConfig struct {
	Embedded *ConfigModelHaljsonEmbedded `json:"_embedded,omitempty"`
	ConfigId *string `json:"configId,omitempty"`
	Name NullableString `json:"name,omitempty"`
	Description NullableString `json:"description,omitempty"`
	Order *int32 `json:"order,omitempty"`
	MigratedConfigId NullableString `json:"migratedConfigId,omitempty"`
	EvaluationVersion *EvaluationVersion `json:"evaluationVersion,omitempty"`
	Links *ConfigModelHaljsonLinks `json:"_links,omitempty"`
}

// NewSettingValueModelHaljsonEmbeddedConfig instantiates a new SettingValueModelHaljsonEmbeddedConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSettingValueModelHaljsonEmbeddedConfig() *SettingValueModelHaljsonEmbeddedConfig {
	this := SettingValueModelHaljsonEmbeddedConfig{}
	return &this
}

// NewSettingValueModelHaljsonEmbeddedConfigWithDefaults instantiates a new SettingValueModelHaljsonEmbeddedConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSettingValueModelHaljsonEmbeddedConfigWithDefaults() *SettingValueModelHaljsonEmbeddedConfig {
	this := SettingValueModelHaljsonEmbeddedConfig{}
	return &this
}

// GetEmbedded returns the Embedded field value if set, zero value otherwise.
func (o *SettingValueModelHaljsonEmbeddedConfig) GetEmbedded() ConfigModelHaljsonEmbedded {
	if o == nil || IsNil(o.Embedded) {
		var ret ConfigModelHaljsonEmbedded
		return ret
	}
	return *o.Embedded
}

// GetEmbeddedOk returns a tuple with the Embedded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingValueModelHaljsonEmbeddedConfig) GetEmbeddedOk() (*ConfigModelHaljsonEmbedded, bool) {
	if o == nil || IsNil(o.Embedded) {
		return nil, false
	}
	return o.Embedded, true
}

// HasEmbedded returns a boolean if a field has been set.
func (o *SettingValueModelHaljsonEmbeddedConfig) HasEmbedded() bool {
	if o != nil && !IsNil(o.Embedded) {
		return true
	}

	return false
}

// SetEmbedded gets a reference to the given ConfigModelHaljsonEmbedded and assigns it to the Embedded field.
func (o *SettingValueModelHaljsonEmbeddedConfig) SetEmbedded(v ConfigModelHaljsonEmbedded) {
	o.Embedded = &v
}

// GetConfigId returns the ConfigId field value if set, zero value otherwise.
func (o *SettingValueModelHaljsonEmbeddedConfig) GetConfigId() string {
	if o == nil || IsNil(o.ConfigId) {
		var ret string
		return ret
	}
	return *o.ConfigId
}

// GetConfigIdOk returns a tuple with the ConfigId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingValueModelHaljsonEmbeddedConfig) GetConfigIdOk() (*string, bool) {
	if o == nil || IsNil(o.ConfigId) {
		return nil, false
	}
	return o.ConfigId, true
}

// HasConfigId returns a boolean if a field has been set.
func (o *SettingValueModelHaljsonEmbeddedConfig) HasConfigId() bool {
	if o != nil && !IsNil(o.ConfigId) {
		return true
	}

	return false
}

// SetConfigId gets a reference to the given string and assigns it to the ConfigId field.
func (o *SettingValueModelHaljsonEmbeddedConfig) SetConfigId(v string) {
	o.ConfigId = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SettingValueModelHaljsonEmbeddedConfig) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SettingValueModelHaljsonEmbeddedConfig) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *SettingValueModelHaljsonEmbeddedConfig) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *SettingValueModelHaljsonEmbeddedConfig) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *SettingValueModelHaljsonEmbeddedConfig) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *SettingValueModelHaljsonEmbeddedConfig) UnsetName() {
	o.Name.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SettingValueModelHaljsonEmbeddedConfig) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SettingValueModelHaljsonEmbeddedConfig) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *SettingValueModelHaljsonEmbeddedConfig) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *SettingValueModelHaljsonEmbeddedConfig) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *SettingValueModelHaljsonEmbeddedConfig) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *SettingValueModelHaljsonEmbeddedConfig) UnsetDescription() {
	o.Description.Unset()
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *SettingValueModelHaljsonEmbeddedConfig) GetOrder() int32 {
	if o == nil || IsNil(o.Order) {
		var ret int32
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingValueModelHaljsonEmbeddedConfig) GetOrderOk() (*int32, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *SettingValueModelHaljsonEmbeddedConfig) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given int32 and assigns it to the Order field.
func (o *SettingValueModelHaljsonEmbeddedConfig) SetOrder(v int32) {
	o.Order = &v
}

// GetMigratedConfigId returns the MigratedConfigId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SettingValueModelHaljsonEmbeddedConfig) GetMigratedConfigId() string {
	if o == nil || IsNil(o.MigratedConfigId.Get()) {
		var ret string
		return ret
	}
	return *o.MigratedConfigId.Get()
}

// GetMigratedConfigIdOk returns a tuple with the MigratedConfigId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SettingValueModelHaljsonEmbeddedConfig) GetMigratedConfigIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MigratedConfigId.Get(), o.MigratedConfigId.IsSet()
}

// HasMigratedConfigId returns a boolean if a field has been set.
func (o *SettingValueModelHaljsonEmbeddedConfig) HasMigratedConfigId() bool {
	if o != nil && o.MigratedConfigId.IsSet() {
		return true
	}

	return false
}

// SetMigratedConfigId gets a reference to the given NullableString and assigns it to the MigratedConfigId field.
func (o *SettingValueModelHaljsonEmbeddedConfig) SetMigratedConfigId(v string) {
	o.MigratedConfigId.Set(&v)
}
// SetMigratedConfigIdNil sets the value for MigratedConfigId to be an explicit nil
func (o *SettingValueModelHaljsonEmbeddedConfig) SetMigratedConfigIdNil() {
	o.MigratedConfigId.Set(nil)
}

// UnsetMigratedConfigId ensures that no value is present for MigratedConfigId, not even an explicit nil
func (o *SettingValueModelHaljsonEmbeddedConfig) UnsetMigratedConfigId() {
	o.MigratedConfigId.Unset()
}

// GetEvaluationVersion returns the EvaluationVersion field value if set, zero value otherwise.
func (o *SettingValueModelHaljsonEmbeddedConfig) GetEvaluationVersion() EvaluationVersion {
	if o == nil || IsNil(o.EvaluationVersion) {
		var ret EvaluationVersion
		return ret
	}
	return *o.EvaluationVersion
}

// GetEvaluationVersionOk returns a tuple with the EvaluationVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingValueModelHaljsonEmbeddedConfig) GetEvaluationVersionOk() (*EvaluationVersion, bool) {
	if o == nil || IsNil(o.EvaluationVersion) {
		return nil, false
	}
	return o.EvaluationVersion, true
}

// HasEvaluationVersion returns a boolean if a field has been set.
func (o *SettingValueModelHaljsonEmbeddedConfig) HasEvaluationVersion() bool {
	if o != nil && !IsNil(o.EvaluationVersion) {
		return true
	}

	return false
}

// SetEvaluationVersion gets a reference to the given EvaluationVersion and assigns it to the EvaluationVersion field.
func (o *SettingValueModelHaljsonEmbeddedConfig) SetEvaluationVersion(v EvaluationVersion) {
	o.EvaluationVersion = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *SettingValueModelHaljsonEmbeddedConfig) GetLinks() ConfigModelHaljsonLinks {
	if o == nil || IsNil(o.Links) {
		var ret ConfigModelHaljsonLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingValueModelHaljsonEmbeddedConfig) GetLinksOk() (*ConfigModelHaljsonLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *SettingValueModelHaljsonEmbeddedConfig) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ConfigModelHaljsonLinks and assigns it to the Links field.
func (o *SettingValueModelHaljsonEmbeddedConfig) SetLinks(v ConfigModelHaljsonLinks) {
	o.Links = &v
}

func (o SettingValueModelHaljsonEmbeddedConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SettingValueModelHaljsonEmbeddedConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Embedded) {
		toSerialize["_embedded"] = o.Embedded
	}
	if !IsNil(o.ConfigId) {
		toSerialize["configId"] = o.ConfigId
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if !IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	if o.MigratedConfigId.IsSet() {
		toSerialize["migratedConfigId"] = o.MigratedConfigId.Get()
	}
	if !IsNil(o.EvaluationVersion) {
		toSerialize["evaluationVersion"] = o.EvaluationVersion
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableSettingValueModelHaljsonEmbeddedConfig struct {
	value *SettingValueModelHaljsonEmbeddedConfig
	isSet bool
}

func (v NullableSettingValueModelHaljsonEmbeddedConfig) Get() *SettingValueModelHaljsonEmbeddedConfig {
	return v.value
}

func (v *NullableSettingValueModelHaljsonEmbeddedConfig) Set(val *SettingValueModelHaljsonEmbeddedConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableSettingValueModelHaljsonEmbeddedConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableSettingValueModelHaljsonEmbeddedConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSettingValueModelHaljsonEmbeddedConfig(val *SettingValueModelHaljsonEmbeddedConfig) *NullableSettingValueModelHaljsonEmbeddedConfig {
	return &NullableSettingValueModelHaljsonEmbeddedConfig{value: val, isSet: true}
}

func (v NullableSettingValueModelHaljsonEmbeddedConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSettingValueModelHaljsonEmbeddedConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


