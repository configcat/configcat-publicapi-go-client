/*
ConfigCat Public Management API

The purpose of this API is to access the ConfigCat platform programmatically. You can **Create**, **Read**, **Update** and **Delete** any entities like **Feature Flags, Configs, Environments** or **Products** within ConfigCat.  **Base API URL**: https://test-api.configcat.com  If you prefer the swagger documentation, you can find it here: [Swagger UI](https://test-api.configcat.com/swagger).  The API is based on HTTP REST, uses resource-oriented URLs, status codes and supports JSON  format. Do not use this API for accessing and evaluating feature flag values. Use the [SDKs instead](https://configcat.com/docs/sdk-reference/overview).   # OpenAPI Specification  The complete specification is publicly available in the following formats:  - [OpenAPI v3](https://test-api.configcat.com/docs/v1/swagger.json) - [Swagger v2](https://test-api.configcat.com/docs/v1/swagger.v2.json)  You can use it to generate client libraries in various languages with [OpenAPI Generator](https://github.com/OpenAPITools/openapi-generator) or [Swagger Codegen](https://swagger.io/tools/swagger-codegen/) to interact with this API.  # Authentication This API uses the [Basic HTTP Authentication Scheme](https://en.wikipedia.org/wiki/Basic_access_authentication).   <!-- ReDoc-Inject: <security-definitions> -->  # Throttling and rate limits All the rate limited API calls are returning information about the current rate limit period in the following HTTP headers:  | Header | Description | | :- | :- | | X-Rate-Limit-Remaining | The maximum number of requests remaining in the current rate limit period. | | X-Rate-Limit-Reset     | The time when the current rate limit period resets.        |  When the rate limit is exceeded by a request, the API returns with a `HTTP 429 - Too many requests` status along with a `Retry-After` HTTP header. 

API version: v1
Contact: support@configcat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configcatpublicapi

import (
	"encoding/json"
)

// checks if the ConfigModelLinksSelf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigModelLinksSelf{}

// ConfigModelLinksSelf struct for ConfigModelLinksSelf
type ConfigModelLinksSelf struct {
	Href *string `json:"href,omitempty"`
}

// NewConfigModelLinksSelf instantiates a new ConfigModelLinksSelf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigModelLinksSelf() *ConfigModelLinksSelf {
	this := ConfigModelLinksSelf{}
	return &this
}

// NewConfigModelLinksSelfWithDefaults instantiates a new ConfigModelLinksSelf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigModelLinksSelfWithDefaults() *ConfigModelLinksSelf {
	this := ConfigModelLinksSelf{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *ConfigModelLinksSelf) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigModelLinksSelf) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *ConfigModelLinksSelf) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *ConfigModelLinksSelf) SetHref(v string) {
	o.Href = &v
}

func (o ConfigModelLinksSelf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigModelLinksSelf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	return toSerialize, nil
}

type NullableConfigModelLinksSelf struct {
	value *ConfigModelLinksSelf
	isSet bool
}

func (v NullableConfigModelLinksSelf) Get() *ConfigModelLinksSelf {
	return v.value
}

func (v *NullableConfigModelLinksSelf) Set(val *ConfigModelLinksSelf) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigModelLinksSelf) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigModelLinksSelf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigModelLinksSelf(val *ConfigModelLinksSelf) *NullableConfigModelLinksSelf {
	return &NullableConfigModelLinksSelf{value: val, isSet: true}
}

func (v NullableConfigModelLinksSelf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigModelLinksSelf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


