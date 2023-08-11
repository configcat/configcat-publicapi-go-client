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

// checks if the ConfigModelHaljsonEmbeddedProductEmbeddedOrganization type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigModelHaljsonEmbeddedProductEmbeddedOrganization{}

// ConfigModelHaljsonEmbeddedProductEmbeddedOrganization struct for ConfigModelHaljsonEmbeddedProductEmbeddedOrganization
type ConfigModelHaljsonEmbeddedProductEmbeddedOrganization struct {
	// Identifier of the Organization.
	OrganizationId *string `json:"organizationId,omitempty"`
	// Name of the Organization.
	Name NullableString `json:"name,omitempty"`
	Links *ConfigModelHaljsonEmbeddedProductEmbeddedOrganizationLinks `json:"_links,omitempty"`
}

// NewConfigModelHaljsonEmbeddedProductEmbeddedOrganization instantiates a new ConfigModelHaljsonEmbeddedProductEmbeddedOrganization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigModelHaljsonEmbeddedProductEmbeddedOrganization() *ConfigModelHaljsonEmbeddedProductEmbeddedOrganization {
	this := ConfigModelHaljsonEmbeddedProductEmbeddedOrganization{}
	return &this
}

// NewConfigModelHaljsonEmbeddedProductEmbeddedOrganizationWithDefaults instantiates a new ConfigModelHaljsonEmbeddedProductEmbeddedOrganization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigModelHaljsonEmbeddedProductEmbeddedOrganizationWithDefaults() *ConfigModelHaljsonEmbeddedProductEmbeddedOrganization {
	this := ConfigModelHaljsonEmbeddedProductEmbeddedOrganization{}
	return &this
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *ConfigModelHaljsonEmbeddedProductEmbeddedOrganization) GetOrganizationId() string {
	if o == nil || IsNil(o.OrganizationId) {
		var ret string
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigModelHaljsonEmbeddedProductEmbeddedOrganization) GetOrganizationIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrganizationId) {
		return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *ConfigModelHaljsonEmbeddedProductEmbeddedOrganization) HasOrganizationId() bool {
	if o != nil && !IsNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given string and assigns it to the OrganizationId field.
func (o *ConfigModelHaljsonEmbeddedProductEmbeddedOrganization) SetOrganizationId(v string) {
	o.OrganizationId = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConfigModelHaljsonEmbeddedProductEmbeddedOrganization) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConfigModelHaljsonEmbeddedProductEmbeddedOrganization) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ConfigModelHaljsonEmbeddedProductEmbeddedOrganization) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ConfigModelHaljsonEmbeddedProductEmbeddedOrganization) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *ConfigModelHaljsonEmbeddedProductEmbeddedOrganization) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ConfigModelHaljsonEmbeddedProductEmbeddedOrganization) UnsetName() {
	o.Name.Unset()
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ConfigModelHaljsonEmbeddedProductEmbeddedOrganization) GetLinks() ConfigModelHaljsonEmbeddedProductEmbeddedOrganizationLinks {
	if o == nil || IsNil(o.Links) {
		var ret ConfigModelHaljsonEmbeddedProductEmbeddedOrganizationLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigModelHaljsonEmbeddedProductEmbeddedOrganization) GetLinksOk() (*ConfigModelHaljsonEmbeddedProductEmbeddedOrganizationLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ConfigModelHaljsonEmbeddedProductEmbeddedOrganization) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ConfigModelHaljsonEmbeddedProductEmbeddedOrganizationLinks and assigns it to the Links field.
func (o *ConfigModelHaljsonEmbeddedProductEmbeddedOrganization) SetLinks(v ConfigModelHaljsonEmbeddedProductEmbeddedOrganizationLinks) {
	o.Links = &v
}

func (o ConfigModelHaljsonEmbeddedProductEmbeddedOrganization) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigModelHaljsonEmbeddedProductEmbeddedOrganization) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OrganizationId) {
		toSerialize["organizationId"] = o.OrganizationId
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableConfigModelHaljsonEmbeddedProductEmbeddedOrganization struct {
	value *ConfigModelHaljsonEmbeddedProductEmbeddedOrganization
	isSet bool
}

func (v NullableConfigModelHaljsonEmbeddedProductEmbeddedOrganization) Get() *ConfigModelHaljsonEmbeddedProductEmbeddedOrganization {
	return v.value
}

func (v *NullableConfigModelHaljsonEmbeddedProductEmbeddedOrganization) Set(val *ConfigModelHaljsonEmbeddedProductEmbeddedOrganization) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigModelHaljsonEmbeddedProductEmbeddedOrganization) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigModelHaljsonEmbeddedProductEmbeddedOrganization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigModelHaljsonEmbeddedProductEmbeddedOrganization(val *ConfigModelHaljsonEmbeddedProductEmbeddedOrganization) *NullableConfigModelHaljsonEmbeddedProductEmbeddedOrganization {
	return &NullableConfigModelHaljsonEmbeddedProductEmbeddedOrganization{value: val, isSet: true}
}

func (v NullableConfigModelHaljsonEmbeddedProductEmbeddedOrganization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigModelHaljsonEmbeddedProductEmbeddedOrganization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


