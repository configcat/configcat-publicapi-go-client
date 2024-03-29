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

// checks if the ConfigModelHaljsonLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigModelHaljsonLinks{}

// ConfigModelHaljsonLinks struct for ConfigModelHaljsonLinks
type ConfigModelHaljsonLinks struct {
	Self *string `json:"self,omitempty"`
	Settings *string `json:"settings,omitempty"`
}

// NewConfigModelHaljsonLinks instantiates a new ConfigModelHaljsonLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigModelHaljsonLinks() *ConfigModelHaljsonLinks {
	this := ConfigModelHaljsonLinks{}
	return &this
}

// NewConfigModelHaljsonLinksWithDefaults instantiates a new ConfigModelHaljsonLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigModelHaljsonLinksWithDefaults() *ConfigModelHaljsonLinks {
	this := ConfigModelHaljsonLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *ConfigModelHaljsonLinks) GetSelf() string {
	if o == nil || IsNil(o.Self) {
		var ret string
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigModelHaljsonLinks) GetSelfOk() (*string, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *ConfigModelHaljsonLinks) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given string and assigns it to the Self field.
func (o *ConfigModelHaljsonLinks) SetSelf(v string) {
	o.Self = &v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *ConfigModelHaljsonLinks) GetSettings() string {
	if o == nil || IsNil(o.Settings) {
		var ret string
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigModelHaljsonLinks) GetSettingsOk() (*string, bool) {
	if o == nil || IsNil(o.Settings) {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *ConfigModelHaljsonLinks) HasSettings() bool {
	if o != nil && !IsNil(o.Settings) {
		return true
	}

	return false
}

// SetSettings gets a reference to the given string and assigns it to the Settings field.
func (o *ConfigModelHaljsonLinks) SetSettings(v string) {
	o.Settings = &v
}

func (o ConfigModelHaljsonLinks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigModelHaljsonLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}
	if !IsNil(o.Settings) {
		toSerialize["settings"] = o.Settings
	}
	return toSerialize, nil
}

type NullableConfigModelHaljsonLinks struct {
	value *ConfigModelHaljsonLinks
	isSet bool
}

func (v NullableConfigModelHaljsonLinks) Get() *ConfigModelHaljsonLinks {
	return v.value
}

func (v *NullableConfigModelHaljsonLinks) Set(val *ConfigModelHaljsonLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigModelHaljsonLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigModelHaljsonLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigModelHaljsonLinks(val *ConfigModelHaljsonLinks) *NullableConfigModelHaljsonLinks {
	return &NullableConfigModelHaljsonLinks{value: val, isSet: true}
}

func (v NullableConfigModelHaljsonLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigModelHaljsonLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


