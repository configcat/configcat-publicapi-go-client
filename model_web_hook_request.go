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

// checks if the WebHookRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebHookRequest{}

// WebHookRequest struct for WebHookRequest
type WebHookRequest struct {
	// The URL of the Webhook.
	Url string `json:"url"`
	// The HTTP body content.
	Content NullableString `json:"content,omitempty"`
	HttpMethod *WebHookHttpMethod `json:"httpMethod,omitempty"`
	// List of HTTP headers.
	WebHookHeaders []WebhookHeaderModel `json:"webHookHeaders,omitempty"`
}

type _WebHookRequest WebHookRequest

// NewWebHookRequest instantiates a new WebHookRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebHookRequest(url string) *WebHookRequest {
	this := WebHookRequest{}
	this.Url = url
	return &this
}

// NewWebHookRequestWithDefaults instantiates a new WebHookRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebHookRequestWithDefaults() *WebHookRequest {
	this := WebHookRequest{}
	return &this
}

// GetUrl returns the Url field value
func (o *WebHookRequest) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *WebHookRequest) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *WebHookRequest) SetUrl(v string) {
	o.Url = v
}

// GetContent returns the Content field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebHookRequest) GetContent() string {
	if o == nil || IsNil(o.Content.Get()) {
		var ret string
		return ret
	}
	return *o.Content.Get()
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebHookRequest) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Content.Get(), o.Content.IsSet()
}

// HasContent returns a boolean if a field has been set.
func (o *WebHookRequest) HasContent() bool {
	if o != nil && o.Content.IsSet() {
		return true
	}

	return false
}

// SetContent gets a reference to the given NullableString and assigns it to the Content field.
func (o *WebHookRequest) SetContent(v string) {
	o.Content.Set(&v)
}
// SetContentNil sets the value for Content to be an explicit nil
func (o *WebHookRequest) SetContentNil() {
	o.Content.Set(nil)
}

// UnsetContent ensures that no value is present for Content, not even an explicit nil
func (o *WebHookRequest) UnsetContent() {
	o.Content.Unset()
}

// GetHttpMethod returns the HttpMethod field value if set, zero value otherwise.
func (o *WebHookRequest) GetHttpMethod() WebHookHttpMethod {
	if o == nil || IsNil(o.HttpMethod) {
		var ret WebHookHttpMethod
		return ret
	}
	return *o.HttpMethod
}

// GetHttpMethodOk returns a tuple with the HttpMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebHookRequest) GetHttpMethodOk() (*WebHookHttpMethod, bool) {
	if o == nil || IsNil(o.HttpMethod) {
		return nil, false
	}
	return o.HttpMethod, true
}

// HasHttpMethod returns a boolean if a field has been set.
func (o *WebHookRequest) HasHttpMethod() bool {
	if o != nil && !IsNil(o.HttpMethod) {
		return true
	}

	return false
}

// SetHttpMethod gets a reference to the given WebHookHttpMethod and assigns it to the HttpMethod field.
func (o *WebHookRequest) SetHttpMethod(v WebHookHttpMethod) {
	o.HttpMethod = &v
}

// GetWebHookHeaders returns the WebHookHeaders field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebHookRequest) GetWebHookHeaders() []WebhookHeaderModel {
	if o == nil {
		var ret []WebhookHeaderModel
		return ret
	}
	return o.WebHookHeaders
}

// GetWebHookHeadersOk returns a tuple with the WebHookHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebHookRequest) GetWebHookHeadersOk() ([]WebhookHeaderModel, bool) {
	if o == nil || IsNil(o.WebHookHeaders) {
		return nil, false
	}
	return o.WebHookHeaders, true
}

// HasWebHookHeaders returns a boolean if a field has been set.
func (o *WebHookRequest) HasWebHookHeaders() bool {
	if o != nil && !IsNil(o.WebHookHeaders) {
		return true
	}

	return false
}

// SetWebHookHeaders gets a reference to the given []WebhookHeaderModel and assigns it to the WebHookHeaders field.
func (o *WebHookRequest) SetWebHookHeaders(v []WebhookHeaderModel) {
	o.WebHookHeaders = v
}

func (o WebHookRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebHookRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["url"] = o.Url
	if o.Content.IsSet() {
		toSerialize["content"] = o.Content.Get()
	}
	if !IsNil(o.HttpMethod) {
		toSerialize["httpMethod"] = o.HttpMethod
	}
	if o.WebHookHeaders != nil {
		toSerialize["webHookHeaders"] = o.WebHookHeaders
	}
	return toSerialize, nil
}

func (o *WebHookRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"url",
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

	varWebHookRequest := _WebHookRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWebHookRequest)

	if err != nil {
		return err
	}

	*o = WebHookRequest(varWebHookRequest)

	return err
}

type NullableWebHookRequest struct {
	value *WebHookRequest
	isSet bool
}

func (v NullableWebHookRequest) Get() *WebHookRequest {
	return v.value
}

func (v *NullableWebHookRequest) Set(val *WebHookRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWebHookRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWebHookRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebHookRequest(val *WebHookRequest) *NullableWebHookRequest {
	return &NullableWebHookRequest{value: val, isSet: true}
}

func (v NullableWebHookRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebHookRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


