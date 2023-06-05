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

// checks if the AddOrUpdateJiraIntegrationLinkModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddOrUpdateJiraIntegrationLinkModel{}

// AddOrUpdateJiraIntegrationLinkModel struct for AddOrUpdateJiraIntegrationLinkModel
type AddOrUpdateJiraIntegrationLinkModel struct {
	JiraJwtToken string `json:"jiraJwtToken"`
	ClientKey string `json:"clientKey"`
	Description NullableString `json:"description,omitempty"`
	Url NullableString `json:"url,omitempty"`
}

// NewAddOrUpdateJiraIntegrationLinkModel instantiates a new AddOrUpdateJiraIntegrationLinkModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddOrUpdateJiraIntegrationLinkModel(jiraJwtToken string, clientKey string) *AddOrUpdateJiraIntegrationLinkModel {
	this := AddOrUpdateJiraIntegrationLinkModel{}
	this.JiraJwtToken = jiraJwtToken
	this.ClientKey = clientKey
	return &this
}

// NewAddOrUpdateJiraIntegrationLinkModelWithDefaults instantiates a new AddOrUpdateJiraIntegrationLinkModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddOrUpdateJiraIntegrationLinkModelWithDefaults() *AddOrUpdateJiraIntegrationLinkModel {
	this := AddOrUpdateJiraIntegrationLinkModel{}
	return &this
}

// GetJiraJwtToken returns the JiraJwtToken field value
func (o *AddOrUpdateJiraIntegrationLinkModel) GetJiraJwtToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JiraJwtToken
}

// GetJiraJwtTokenOk returns a tuple with the JiraJwtToken field value
// and a boolean to check if the value has been set.
func (o *AddOrUpdateJiraIntegrationLinkModel) GetJiraJwtTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JiraJwtToken, true
}

// SetJiraJwtToken sets field value
func (o *AddOrUpdateJiraIntegrationLinkModel) SetJiraJwtToken(v string) {
	o.JiraJwtToken = v
}

// GetClientKey returns the ClientKey field value
func (o *AddOrUpdateJiraIntegrationLinkModel) GetClientKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientKey
}

// GetClientKeyOk returns a tuple with the ClientKey field value
// and a boolean to check if the value has been set.
func (o *AddOrUpdateJiraIntegrationLinkModel) GetClientKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientKey, true
}

// SetClientKey sets field value
func (o *AddOrUpdateJiraIntegrationLinkModel) SetClientKey(v string) {
	o.ClientKey = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AddOrUpdateJiraIntegrationLinkModel) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AddOrUpdateJiraIntegrationLinkModel) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *AddOrUpdateJiraIntegrationLinkModel) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *AddOrUpdateJiraIntegrationLinkModel) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *AddOrUpdateJiraIntegrationLinkModel) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *AddOrUpdateJiraIntegrationLinkModel) UnsetDescription() {
	o.Description.Unset()
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AddOrUpdateJiraIntegrationLinkModel) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AddOrUpdateJiraIntegrationLinkModel) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *AddOrUpdateJiraIntegrationLinkModel) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *AddOrUpdateJiraIntegrationLinkModel) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *AddOrUpdateJiraIntegrationLinkModel) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *AddOrUpdateJiraIntegrationLinkModel) UnsetUrl() {
	o.Url.Unset()
}

func (o AddOrUpdateJiraIntegrationLinkModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddOrUpdateJiraIntegrationLinkModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["jiraJwtToken"] = o.JiraJwtToken
	toSerialize["clientKey"] = o.ClientKey
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	return toSerialize, nil
}

type NullableAddOrUpdateJiraIntegrationLinkModel struct {
	value *AddOrUpdateJiraIntegrationLinkModel
	isSet bool
}

func (v NullableAddOrUpdateJiraIntegrationLinkModel) Get() *AddOrUpdateJiraIntegrationLinkModel {
	return v.value
}

func (v *NullableAddOrUpdateJiraIntegrationLinkModel) Set(val *AddOrUpdateJiraIntegrationLinkModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAddOrUpdateJiraIntegrationLinkModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAddOrUpdateJiraIntegrationLinkModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddOrUpdateJiraIntegrationLinkModel(val *AddOrUpdateJiraIntegrationLinkModel) *NullableAddOrUpdateJiraIntegrationLinkModel {
	return &NullableAddOrUpdateJiraIntegrationLinkModel{value: val, isSet: true}
}

func (v NullableAddOrUpdateJiraIntegrationLinkModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddOrUpdateJiraIntegrationLinkModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


