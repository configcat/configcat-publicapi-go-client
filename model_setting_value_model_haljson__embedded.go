/*
ConfigCat Public Management API

**Base API URL**: https://api.configcat.com  If you prefer the swagger documentation, you can find it here: [Swagger UI](https://api.configcat.com/swagger).  The purpose of this API is to access the ConfigCat platform programmatically.  You can **Create**, **Read**, **Update** and **Delete** any entities like **Feature Flags, Configs, Environments** or **Products** within ConfigCat.   The API is based on HTTP REST, uses resource-oriented URLs, status codes and supports JSON  and JSON+HAL format. Do not use this API for accessing and evaluating feature flag values. Use the [SDKs instead](https://configcat.com/docs/sdk-reference/overview).   # OpenAPI Specification  The complete specification is publicly available here: [swagger.json](v1/swagger.json).  You can use it to generate client libraries in various languages with [OpenAPI Generator](https://github.com/OpenAPITools/openapi-generator) or [Swagger Codegen](https://swagger.io/tools/swagger-codegen/) to interact with this API.  # Authentication This API uses the [Basic HTTP Authentication Scheme](https://en.wikipedia.org/wiki/Basic_access_authentication).   <!-- ReDoc-Inject: <security-definitions> -->  # Throttling and rate limits All the rate limited API calls are returning information about the current rate limit period in the following HTTP headers:  | Header | Description | | :- | :- | | X-Rate-Limit-Remaining | The maximum number of requests remaining in the current rate limit period. | | X-Rate-Limit-Reset     | The time when the current rate limit period resets.        |  When the rate limit is exceeded by a request, the API returns with a `HTTP 429 - Too many requests` status along with a `Retry-After` HTTP header. 

API version: v1
Contact: support@configcat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configcatpublicapi

import (
	"encoding/json"
)

// checks if the SettingValueModelHaljsonEmbedded type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SettingValueModelHaljsonEmbedded{}

// SettingValueModelHaljsonEmbedded struct for SettingValueModelHaljsonEmbedded
type SettingValueModelHaljsonEmbedded struct {
	Setting *SettingValueModelHaljsonEmbeddedSetting `json:"setting,omitempty"`
	Config *SettingValueModelHaljsonEmbeddedConfig `json:"config,omitempty"`
	Environment *SettingValueModelHaljsonEmbeddedEnvironment `json:"environment,omitempty"`
	IntegrationLinks []SettingValueModelHaljsonEmbeddedIntegrationLinksInner `json:"integrationLinks,omitempty"`
	SettingTags []SettingValueModelHaljsonEmbeddedSettingTagsInner `json:"settingTags,omitempty"`
}

// NewSettingValueModelHaljsonEmbedded instantiates a new SettingValueModelHaljsonEmbedded object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSettingValueModelHaljsonEmbedded() *SettingValueModelHaljsonEmbedded {
	this := SettingValueModelHaljsonEmbedded{}
	return &this
}

// NewSettingValueModelHaljsonEmbeddedWithDefaults instantiates a new SettingValueModelHaljsonEmbedded object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSettingValueModelHaljsonEmbeddedWithDefaults() *SettingValueModelHaljsonEmbedded {
	this := SettingValueModelHaljsonEmbedded{}
	return &this
}

// GetSetting returns the Setting field value if set, zero value otherwise.
func (o *SettingValueModelHaljsonEmbedded) GetSetting() SettingValueModelHaljsonEmbeddedSetting {
	if o == nil || IsNil(o.Setting) {
		var ret SettingValueModelHaljsonEmbeddedSetting
		return ret
	}
	return *o.Setting
}

// GetSettingOk returns a tuple with the Setting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingValueModelHaljsonEmbedded) GetSettingOk() (*SettingValueModelHaljsonEmbeddedSetting, bool) {
	if o == nil || IsNil(o.Setting) {
		return nil, false
	}
	return o.Setting, true
}

// HasSetting returns a boolean if a field has been set.
func (o *SettingValueModelHaljsonEmbedded) HasSetting() bool {
	if o != nil && !IsNil(o.Setting) {
		return true
	}

	return false
}

// SetSetting gets a reference to the given SettingValueModelHaljsonEmbeddedSetting and assigns it to the Setting field.
func (o *SettingValueModelHaljsonEmbedded) SetSetting(v SettingValueModelHaljsonEmbeddedSetting) {
	o.Setting = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *SettingValueModelHaljsonEmbedded) GetConfig() SettingValueModelHaljsonEmbeddedConfig {
	if o == nil || IsNil(o.Config) {
		var ret SettingValueModelHaljsonEmbeddedConfig
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingValueModelHaljsonEmbedded) GetConfigOk() (*SettingValueModelHaljsonEmbeddedConfig, bool) {
	if o == nil || IsNil(o.Config) {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *SettingValueModelHaljsonEmbedded) HasConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given SettingValueModelHaljsonEmbeddedConfig and assigns it to the Config field.
func (o *SettingValueModelHaljsonEmbedded) SetConfig(v SettingValueModelHaljsonEmbeddedConfig) {
	o.Config = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *SettingValueModelHaljsonEmbedded) GetEnvironment() SettingValueModelHaljsonEmbeddedEnvironment {
	if o == nil || IsNil(o.Environment) {
		var ret SettingValueModelHaljsonEmbeddedEnvironment
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingValueModelHaljsonEmbedded) GetEnvironmentOk() (*SettingValueModelHaljsonEmbeddedEnvironment, bool) {
	if o == nil || IsNil(o.Environment) {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *SettingValueModelHaljsonEmbedded) HasEnvironment() bool {
	if o != nil && !IsNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given SettingValueModelHaljsonEmbeddedEnvironment and assigns it to the Environment field.
func (o *SettingValueModelHaljsonEmbedded) SetEnvironment(v SettingValueModelHaljsonEmbeddedEnvironment) {
	o.Environment = &v
}

// GetIntegrationLinks returns the IntegrationLinks field value if set, zero value otherwise.
func (o *SettingValueModelHaljsonEmbedded) GetIntegrationLinks() []SettingValueModelHaljsonEmbeddedIntegrationLinksInner {
	if o == nil || IsNil(o.IntegrationLinks) {
		var ret []SettingValueModelHaljsonEmbeddedIntegrationLinksInner
		return ret
	}
	return o.IntegrationLinks
}

// GetIntegrationLinksOk returns a tuple with the IntegrationLinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingValueModelHaljsonEmbedded) GetIntegrationLinksOk() ([]SettingValueModelHaljsonEmbeddedIntegrationLinksInner, bool) {
	if o == nil || IsNil(o.IntegrationLinks) {
		return nil, false
	}
	return o.IntegrationLinks, true
}

// HasIntegrationLinks returns a boolean if a field has been set.
func (o *SettingValueModelHaljsonEmbedded) HasIntegrationLinks() bool {
	if o != nil && !IsNil(o.IntegrationLinks) {
		return true
	}

	return false
}

// SetIntegrationLinks gets a reference to the given []SettingValueModelHaljsonEmbeddedIntegrationLinksInner and assigns it to the IntegrationLinks field.
func (o *SettingValueModelHaljsonEmbedded) SetIntegrationLinks(v []SettingValueModelHaljsonEmbeddedIntegrationLinksInner) {
	o.IntegrationLinks = v
}

// GetSettingTags returns the SettingTags field value if set, zero value otherwise.
func (o *SettingValueModelHaljsonEmbedded) GetSettingTags() []SettingValueModelHaljsonEmbeddedSettingTagsInner {
	if o == nil || IsNil(o.SettingTags) {
		var ret []SettingValueModelHaljsonEmbeddedSettingTagsInner
		return ret
	}
	return o.SettingTags
}

// GetSettingTagsOk returns a tuple with the SettingTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingValueModelHaljsonEmbedded) GetSettingTagsOk() ([]SettingValueModelHaljsonEmbeddedSettingTagsInner, bool) {
	if o == nil || IsNil(o.SettingTags) {
		return nil, false
	}
	return o.SettingTags, true
}

// HasSettingTags returns a boolean if a field has been set.
func (o *SettingValueModelHaljsonEmbedded) HasSettingTags() bool {
	if o != nil && !IsNil(o.SettingTags) {
		return true
	}

	return false
}

// SetSettingTags gets a reference to the given []SettingValueModelHaljsonEmbeddedSettingTagsInner and assigns it to the SettingTags field.
func (o *SettingValueModelHaljsonEmbedded) SetSettingTags(v []SettingValueModelHaljsonEmbeddedSettingTagsInner) {
	o.SettingTags = v
}

func (o SettingValueModelHaljsonEmbedded) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SettingValueModelHaljsonEmbedded) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Setting) {
		toSerialize["setting"] = o.Setting
	}
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	if !IsNil(o.Environment) {
		toSerialize["environment"] = o.Environment
	}
	if !IsNil(o.IntegrationLinks) {
		toSerialize["integrationLinks"] = o.IntegrationLinks
	}
	if !IsNil(o.SettingTags) {
		toSerialize["settingTags"] = o.SettingTags
	}
	return toSerialize, nil
}

type NullableSettingValueModelHaljsonEmbedded struct {
	value *SettingValueModelHaljsonEmbedded
	isSet bool
}

func (v NullableSettingValueModelHaljsonEmbedded) Get() *SettingValueModelHaljsonEmbedded {
	return v.value
}

func (v *NullableSettingValueModelHaljsonEmbedded) Set(val *SettingValueModelHaljsonEmbedded) {
	v.value = val
	v.isSet = true
}

func (v NullableSettingValueModelHaljsonEmbedded) IsSet() bool {
	return v.isSet
}

func (v *NullableSettingValueModelHaljsonEmbedded) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSettingValueModelHaljsonEmbedded(val *SettingValueModelHaljsonEmbedded) *NullableSettingValueModelHaljsonEmbedded {
	return &NullableSettingValueModelHaljsonEmbedded{value: val, isSet: true}
}

func (v NullableSettingValueModelHaljsonEmbedded) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSettingValueModelHaljsonEmbedded) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


