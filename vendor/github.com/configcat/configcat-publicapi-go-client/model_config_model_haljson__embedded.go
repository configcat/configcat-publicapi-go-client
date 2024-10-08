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

// checks if the ConfigModelHaljsonEmbedded type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigModelHaljsonEmbedded{}

// ConfigModelHaljsonEmbedded struct for ConfigModelHaljsonEmbedded
type ConfigModelHaljsonEmbedded struct {
	Product *ConfigModelHaljsonEmbeddedProduct `json:"product,omitempty"`
}

// NewConfigModelHaljsonEmbedded instantiates a new ConfigModelHaljsonEmbedded object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigModelHaljsonEmbedded() *ConfigModelHaljsonEmbedded {
	this := ConfigModelHaljsonEmbedded{}
	return &this
}

// NewConfigModelHaljsonEmbeddedWithDefaults instantiates a new ConfigModelHaljsonEmbedded object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigModelHaljsonEmbeddedWithDefaults() *ConfigModelHaljsonEmbedded {
	this := ConfigModelHaljsonEmbedded{}
	return &this
}

// GetProduct returns the Product field value if set, zero value otherwise.
func (o *ConfigModelHaljsonEmbedded) GetProduct() ConfigModelHaljsonEmbeddedProduct {
	if o == nil || IsNil(o.Product) {
		var ret ConfigModelHaljsonEmbeddedProduct
		return ret
	}
	return *o.Product
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigModelHaljsonEmbedded) GetProductOk() (*ConfigModelHaljsonEmbeddedProduct, bool) {
	if o == nil || IsNil(o.Product) {
		return nil, false
	}
	return o.Product, true
}

// HasProduct returns a boolean if a field has been set.
func (o *ConfigModelHaljsonEmbedded) HasProduct() bool {
	if o != nil && !IsNil(o.Product) {
		return true
	}

	return false
}

// SetProduct gets a reference to the given ConfigModelHaljsonEmbeddedProduct and assigns it to the Product field.
func (o *ConfigModelHaljsonEmbedded) SetProduct(v ConfigModelHaljsonEmbeddedProduct) {
	o.Product = &v
}

func (o ConfigModelHaljsonEmbedded) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigModelHaljsonEmbedded) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Product) {
		toSerialize["product"] = o.Product
	}
	return toSerialize, nil
}

type NullableConfigModelHaljsonEmbedded struct {
	value *ConfigModelHaljsonEmbedded
	isSet bool
}

func (v NullableConfigModelHaljsonEmbedded) Get() *ConfigModelHaljsonEmbedded {
	return v.value
}

func (v *NullableConfigModelHaljsonEmbedded) Set(val *ConfigModelHaljsonEmbedded) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigModelHaljsonEmbedded) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigModelHaljsonEmbedded) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigModelHaljsonEmbedded(val *ConfigModelHaljsonEmbedded) *NullableConfigModelHaljsonEmbedded {
	return &NullableConfigModelHaljsonEmbedded{value: val, isSet: true}
}

func (v NullableConfigModelHaljsonEmbedded) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigModelHaljsonEmbedded) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


