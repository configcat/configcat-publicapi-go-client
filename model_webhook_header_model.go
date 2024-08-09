/*
ConfigCat Public Management API

The purpose of this API is to access the ConfigCat platform programmatically. You can **Create**, **Read**, **Update** and **Delete** any entities like **Feature Flags, Configs, Environments** or **Products** within ConfigCat.  **Base API URL**: https://test-api.configcat.com  If you prefer the swagger documentation, you can find it here: [Swagger UI](https://test-api.configcat.com/swagger).  The API is based on HTTP REST, uses resource-oriented URLs, status codes and supports JSON  format.   **Important:** Do not use this API for accessing and evaluating feature flag values. Use the [SDKs](https://configcat.com/docs/sdk-reference/overview) or the [ConfigCat Proxy](https://configcat.com/docs/advanced/proxy/proxy-overview/) instead.  # OpenAPI Specification  The complete specification is publicly available in the following formats:  - [OpenAPI v3](https://test-api.configcat.com/docs/v1/swagger.json) - [Swagger v2](https://test-api.configcat.com/docs/v1/swagger.v2.json)  You can use it to generate client libraries in various languages with [OpenAPI Generator](https://github.com/OpenAPITools/openapi-generator) or [Swagger Codegen](https://swagger.io/tools/swagger-codegen/) to interact with this API.  # Authentication This API uses the [Basic HTTP Authentication Scheme](https://en.wikipedia.org/wiki/Basic_access_authentication).   <!-- ReDoc-Inject: <security-definitions> -->  # Throttling and rate limits All the rate limited API calls are returning information about the current rate limit period in the following HTTP headers:  | Header | Description | | :- | :- | | X-Rate-Limit-Remaining | The maximum number of requests remaining in the current rate limit period. | | X-Rate-Limit-Reset     | The time when the current rate limit period resets.        |  When the rate limit is exceeded by a request, the API returns with a `HTTP 429 - Too many requests` status along with a `Retry-After` HTTP header. 

API version: v1
Contact: support@configcat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configcatpublicapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the WebhookHeaderModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookHeaderModel{}

// WebhookHeaderModel struct for WebhookHeaderModel
type WebhookHeaderModel struct {
	// The HTTP header key.
	Key string `json:"key"`
	// The HTTP header value.
	Value string `json:"value"`
	// Indicates whether the header value is sensitive.
	IsSecure *bool `json:"isSecure,omitempty"`
}

type _WebhookHeaderModel WebhookHeaderModel

// NewWebhookHeaderModel instantiates a new WebhookHeaderModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookHeaderModel(key string, value string) *WebhookHeaderModel {
	this := WebhookHeaderModel{}
	this.Key = key
	this.Value = value
	return &this
}

// NewWebhookHeaderModelWithDefaults instantiates a new WebhookHeaderModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookHeaderModelWithDefaults() *WebhookHeaderModel {
	this := WebhookHeaderModel{}
	return &this
}

// GetKey returns the Key field value
func (o *WebhookHeaderModel) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *WebhookHeaderModel) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *WebhookHeaderModel) SetKey(v string) {
	o.Key = v
}

// GetValue returns the Value field value
func (o *WebhookHeaderModel) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *WebhookHeaderModel) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *WebhookHeaderModel) SetValue(v string) {
	o.Value = v
}

// GetIsSecure returns the IsSecure field value if set, zero value otherwise.
func (o *WebhookHeaderModel) GetIsSecure() bool {
	if o == nil || IsNil(o.IsSecure) {
		var ret bool
		return ret
	}
	return *o.IsSecure
}

// GetIsSecureOk returns a tuple with the IsSecure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookHeaderModel) GetIsSecureOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSecure) {
		return nil, false
	}
	return o.IsSecure, true
}

// HasIsSecure returns a boolean if a field has been set.
func (o *WebhookHeaderModel) HasIsSecure() bool {
	if o != nil && !IsNil(o.IsSecure) {
		return true
	}

	return false
}

// SetIsSecure gets a reference to the given bool and assigns it to the IsSecure field.
func (o *WebhookHeaderModel) SetIsSecure(v bool) {
	o.IsSecure = &v
}

func (o WebhookHeaderModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookHeaderModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	toSerialize["value"] = o.Value
	if !IsNil(o.IsSecure) {
		toSerialize["isSecure"] = o.IsSecure
	}
	return toSerialize, nil
}

func (o *WebhookHeaderModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"key",
		"value",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varWebhookHeaderModel := _WebhookHeaderModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWebhookHeaderModel)

	if err != nil {
		return err
	}

	*o = WebhookHeaderModel(varWebhookHeaderModel)

	return err
}

type NullableWebhookHeaderModel struct {
	value *WebhookHeaderModel
	isSet bool
}

func (v NullableWebhookHeaderModel) Get() *WebhookHeaderModel {
	return v.value
}

func (v *NullableWebhookHeaderModel) Set(val *WebhookHeaderModel) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookHeaderModel) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookHeaderModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookHeaderModel(val *WebhookHeaderModel) *NullableWebhookHeaderModel {
	return &NullableWebhookHeaderModel{value: val, isSet: true}
}

func (v NullableWebhookHeaderModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookHeaderModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


