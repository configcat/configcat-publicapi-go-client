/*
ConfigCat Public Management API

The purpose of this API is to access the ConfigCat platform programmatically. You can **Create**, **Read**, **Update** and **Delete** any entities like **Feature Flags, Configs, Environments** or **Products** within ConfigCat.  **Base API URL**: https://api.configcat.com  If you prefer the swagger documentation, you can find it here: [Swagger UI](https://api.configcat.com/swagger).  The API is based on HTTP REST, uses resource-oriented URLs, status codes and supports JSON  format.   **Important:** Do not use this API for accessing and evaluating feature flag values. Use the [SDKs](https://configcat.com/docs/sdk-reference/overview) or the [ConfigCat Proxy](https://configcat.com/docs/advanced/proxy/proxy-overview/) instead.  # OpenAPI Specification  The complete specification is publicly available in the following formats:  - [OpenAPI v3](https://api.configcat.com/docs/v1/swagger.json) - [Swagger v2](https://api.configcat.com/docs/v1/swagger.v2.json)  You can use it to generate client libraries in various languages with [OpenAPI Generator](https://github.com/OpenAPITools/openapi-generator) or [Swagger Codegen](https://swagger.io/tools/swagger-codegen/) to interact with this API.  # Authentication This API uses the [Basic HTTP Authentication Scheme](https://en.wikipedia.org/wiki/Basic_access_authentication).   <!-- ReDoc-Inject: <security-definitions> -->  # Throttling and rate limits All the rate limited API calls are returning information about the current rate limit period in the following HTTP headers:  | Header | Description | | :- | :- | | X-Rate-Limit-Remaining | The maximum number of requests remaining in the current rate limit period. | | X-Rate-Limit-Reset     | The time when the current rate limit period resets.        |  When the rate limit is exceeded by a request, the API returns with a `HTTP 429 - Too many requests` status along with a `Retry-After` HTTP header. 

API version: v1
Contact: support@configcat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configcatpublicapi

import (
	"encoding/json"
)

// checks if the WebhookSigningKeysModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookSigningKeysModel{}

// WebhookSigningKeysModel struct for WebhookSigningKeysModel
type WebhookSigningKeysModel struct {
	// The first signing key.
	Key1 NullableString `json:"key1,omitempty"`
	// The second signing key.
	Key2 NullableString `json:"key2,omitempty"`
}

// NewWebhookSigningKeysModel instantiates a new WebhookSigningKeysModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookSigningKeysModel() *WebhookSigningKeysModel {
	this := WebhookSigningKeysModel{}
	return &this
}

// NewWebhookSigningKeysModelWithDefaults instantiates a new WebhookSigningKeysModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookSigningKeysModelWithDefaults() *WebhookSigningKeysModel {
	this := WebhookSigningKeysModel{}
	return &this
}

// GetKey1 returns the Key1 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebhookSigningKeysModel) GetKey1() string {
	if o == nil || IsNil(o.Key1.Get()) {
		var ret string
		return ret
	}
	return *o.Key1.Get()
}

// GetKey1Ok returns a tuple with the Key1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhookSigningKeysModel) GetKey1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Key1.Get(), o.Key1.IsSet()
}

// HasKey1 returns a boolean if a field has been set.
func (o *WebhookSigningKeysModel) HasKey1() bool {
	if o != nil && o.Key1.IsSet() {
		return true
	}

	return false
}

// SetKey1 gets a reference to the given NullableString and assigns it to the Key1 field.
func (o *WebhookSigningKeysModel) SetKey1(v string) {
	o.Key1.Set(&v)
}
// SetKey1Nil sets the value for Key1 to be an explicit nil
func (o *WebhookSigningKeysModel) SetKey1Nil() {
	o.Key1.Set(nil)
}

// UnsetKey1 ensures that no value is present for Key1, not even an explicit nil
func (o *WebhookSigningKeysModel) UnsetKey1() {
	o.Key1.Unset()
}

// GetKey2 returns the Key2 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebhookSigningKeysModel) GetKey2() string {
	if o == nil || IsNil(o.Key2.Get()) {
		var ret string
		return ret
	}
	return *o.Key2.Get()
}

// GetKey2Ok returns a tuple with the Key2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhookSigningKeysModel) GetKey2Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Key2.Get(), o.Key2.IsSet()
}

// HasKey2 returns a boolean if a field has been set.
func (o *WebhookSigningKeysModel) HasKey2() bool {
	if o != nil && o.Key2.IsSet() {
		return true
	}

	return false
}

// SetKey2 gets a reference to the given NullableString and assigns it to the Key2 field.
func (o *WebhookSigningKeysModel) SetKey2(v string) {
	o.Key2.Set(&v)
}
// SetKey2Nil sets the value for Key2 to be an explicit nil
func (o *WebhookSigningKeysModel) SetKey2Nil() {
	o.Key2.Set(nil)
}

// UnsetKey2 ensures that no value is present for Key2, not even an explicit nil
func (o *WebhookSigningKeysModel) UnsetKey2() {
	o.Key2.Unset()
}

func (o WebhookSigningKeysModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookSigningKeysModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Key1.IsSet() {
		toSerialize["key1"] = o.Key1.Get()
	}
	if o.Key2.IsSet() {
		toSerialize["key2"] = o.Key2.Get()
	}
	return toSerialize, nil
}

type NullableWebhookSigningKeysModel struct {
	value *WebhookSigningKeysModel
	isSet bool
}

func (v NullableWebhookSigningKeysModel) Get() *WebhookSigningKeysModel {
	return v.value
}

func (v *NullableWebhookSigningKeysModel) Set(val *WebhookSigningKeysModel) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookSigningKeysModel) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookSigningKeysModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookSigningKeysModel(val *WebhookSigningKeysModel) *NullableWebhookSigningKeysModel {
	return &NullableWebhookSigningKeysModel{value: val, isSet: true}
}

func (v NullableWebhookSigningKeysModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookSigningKeysModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


