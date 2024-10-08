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
)

// checks if the WebhookModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookModel{}

// WebhookModel struct for WebhookModel
type WebhookModel struct {
	// The identifier of the Webhook.
	WebhookId *int32 `json:"webhookId,omitempty"`
	// The URL of the Webhook.
	Url NullableString `json:"url,omitempty"`
	HttpMethod *WebHookHttpMethod `json:"httpMethod,omitempty"`
	// The HTTP body content.
	Content NullableString `json:"content,omitempty"`
	// List of HTTP headers that the Webhook must send.
	WebHookHeaders []WebhookHeaderModel `json:"webHookHeaders,omitempty"`
	Config *WebhookConfig `json:"config,omitempty"`
	Environment *WebhookEnvironment `json:"environment,omitempty"`
}

// NewWebhookModel instantiates a new WebhookModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookModel() *WebhookModel {
	this := WebhookModel{}
	return &this
}

// NewWebhookModelWithDefaults instantiates a new WebhookModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookModelWithDefaults() *WebhookModel {
	this := WebhookModel{}
	return &this
}

// GetWebhookId returns the WebhookId field value if set, zero value otherwise.
func (o *WebhookModel) GetWebhookId() int32 {
	if o == nil || IsNil(o.WebhookId) {
		var ret int32
		return ret
	}
	return *o.WebhookId
}

// GetWebhookIdOk returns a tuple with the WebhookId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookModel) GetWebhookIdOk() (*int32, bool) {
	if o == nil || IsNil(o.WebhookId) {
		return nil, false
	}
	return o.WebhookId, true
}

// HasWebhookId returns a boolean if a field has been set.
func (o *WebhookModel) HasWebhookId() bool {
	if o != nil && !IsNil(o.WebhookId) {
		return true
	}

	return false
}

// SetWebhookId gets a reference to the given int32 and assigns it to the WebhookId field.
func (o *WebhookModel) SetWebhookId(v int32) {
	o.WebhookId = &v
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebhookModel) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhookModel) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *WebhookModel) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *WebhookModel) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *WebhookModel) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *WebhookModel) UnsetUrl() {
	o.Url.Unset()
}

// GetHttpMethod returns the HttpMethod field value if set, zero value otherwise.
func (o *WebhookModel) GetHttpMethod() WebHookHttpMethod {
	if o == nil || IsNil(o.HttpMethod) {
		var ret WebHookHttpMethod
		return ret
	}
	return *o.HttpMethod
}

// GetHttpMethodOk returns a tuple with the HttpMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookModel) GetHttpMethodOk() (*WebHookHttpMethod, bool) {
	if o == nil || IsNil(o.HttpMethod) {
		return nil, false
	}
	return o.HttpMethod, true
}

// HasHttpMethod returns a boolean if a field has been set.
func (o *WebhookModel) HasHttpMethod() bool {
	if o != nil && !IsNil(o.HttpMethod) {
		return true
	}

	return false
}

// SetHttpMethod gets a reference to the given WebHookHttpMethod and assigns it to the HttpMethod field.
func (o *WebhookModel) SetHttpMethod(v WebHookHttpMethod) {
	o.HttpMethod = &v
}

// GetContent returns the Content field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebhookModel) GetContent() string {
	if o == nil || IsNil(o.Content.Get()) {
		var ret string
		return ret
	}
	return *o.Content.Get()
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhookModel) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Content.Get(), o.Content.IsSet()
}

// HasContent returns a boolean if a field has been set.
func (o *WebhookModel) HasContent() bool {
	if o != nil && o.Content.IsSet() {
		return true
	}

	return false
}

// SetContent gets a reference to the given NullableString and assigns it to the Content field.
func (o *WebhookModel) SetContent(v string) {
	o.Content.Set(&v)
}
// SetContentNil sets the value for Content to be an explicit nil
func (o *WebhookModel) SetContentNil() {
	o.Content.Set(nil)
}

// UnsetContent ensures that no value is present for Content, not even an explicit nil
func (o *WebhookModel) UnsetContent() {
	o.Content.Unset()
}

// GetWebHookHeaders returns the WebHookHeaders field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebhookModel) GetWebHookHeaders() []WebhookHeaderModel {
	if o == nil {
		var ret []WebhookHeaderModel
		return ret
	}
	return o.WebHookHeaders
}

// GetWebHookHeadersOk returns a tuple with the WebHookHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhookModel) GetWebHookHeadersOk() ([]WebhookHeaderModel, bool) {
	if o == nil || IsNil(o.WebHookHeaders) {
		return nil, false
	}
	return o.WebHookHeaders, true
}

// HasWebHookHeaders returns a boolean if a field has been set.
func (o *WebhookModel) HasWebHookHeaders() bool {
	if o != nil && IsNil(o.WebHookHeaders) {
		return true
	}

	return false
}

// SetWebHookHeaders gets a reference to the given []WebhookHeaderModel and assigns it to the WebHookHeaders field.
func (o *WebhookModel) SetWebHookHeaders(v []WebhookHeaderModel) {
	o.WebHookHeaders = v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *WebhookModel) GetConfig() WebhookConfig {
	if o == nil || IsNil(o.Config) {
		var ret WebhookConfig
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookModel) GetConfigOk() (*WebhookConfig, bool) {
	if o == nil || IsNil(o.Config) {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *WebhookModel) HasConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given WebhookConfig and assigns it to the Config field.
func (o *WebhookModel) SetConfig(v WebhookConfig) {
	o.Config = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *WebhookModel) GetEnvironment() WebhookEnvironment {
	if o == nil || IsNil(o.Environment) {
		var ret WebhookEnvironment
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookModel) GetEnvironmentOk() (*WebhookEnvironment, bool) {
	if o == nil || IsNil(o.Environment) {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *WebhookModel) HasEnvironment() bool {
	if o != nil && !IsNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given WebhookEnvironment and assigns it to the Environment field.
func (o *WebhookModel) SetEnvironment(v WebhookEnvironment) {
	o.Environment = &v
}

func (o WebhookModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.WebhookId) {
		toSerialize["webhookId"] = o.WebhookId
	}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	if !IsNil(o.HttpMethod) {
		toSerialize["httpMethod"] = o.HttpMethod
	}
	if o.Content.IsSet() {
		toSerialize["content"] = o.Content.Get()
	}
	if o.WebHookHeaders != nil {
		toSerialize["webHookHeaders"] = o.WebHookHeaders
	}
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	if !IsNil(o.Environment) {
		toSerialize["environment"] = o.Environment
	}
	return toSerialize, nil
}

type NullableWebhookModel struct {
	value *WebhookModel
	isSet bool
}

func (v NullableWebhookModel) Get() *WebhookModel {
	return v.value
}

func (v *NullableWebhookModel) Set(val *WebhookModel) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookModel) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookModel(val *WebhookModel) *NullableWebhookModel {
	return &NullableWebhookModel{value: val, isSet: true}
}

func (v NullableWebhookModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


