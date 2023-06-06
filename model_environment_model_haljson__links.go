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

// checks if the EnvironmentModelHaljsonLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnvironmentModelHaljsonLinks{}

// EnvironmentModelHaljsonLinks struct for EnvironmentModelHaljsonLinks
type EnvironmentModelHaljsonLinks struct {
	Self *string `json:"self,omitempty"`
}

// NewEnvironmentModelHaljsonLinks instantiates a new EnvironmentModelHaljsonLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentModelHaljsonLinks() *EnvironmentModelHaljsonLinks {
	this := EnvironmentModelHaljsonLinks{}
	return &this
}

// NewEnvironmentModelHaljsonLinksWithDefaults instantiates a new EnvironmentModelHaljsonLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentModelHaljsonLinksWithDefaults() *EnvironmentModelHaljsonLinks {
	this := EnvironmentModelHaljsonLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *EnvironmentModelHaljsonLinks) GetSelf() string {
	if o == nil || IsNil(o.Self) {
		var ret string
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentModelHaljsonLinks) GetSelfOk() (*string, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *EnvironmentModelHaljsonLinks) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given string and assigns it to the Self field.
func (o *EnvironmentModelHaljsonLinks) SetSelf(v string) {
	o.Self = &v
}

func (o EnvironmentModelHaljsonLinks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnvironmentModelHaljsonLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}
	return toSerialize, nil
}

type NullableEnvironmentModelHaljsonLinks struct {
	value *EnvironmentModelHaljsonLinks
	isSet bool
}

func (v NullableEnvironmentModelHaljsonLinks) Get() *EnvironmentModelHaljsonLinks {
	return v.value
}

func (v *NullableEnvironmentModelHaljsonLinks) Set(val *EnvironmentModelHaljsonLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentModelHaljsonLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentModelHaljsonLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentModelHaljsonLinks(val *EnvironmentModelHaljsonLinks) *NullableEnvironmentModelHaljsonLinks {
	return &NullableEnvironmentModelHaljsonLinks{value: val, isSet: true}
}

func (v NullableEnvironmentModelHaljsonLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentModelHaljsonLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


