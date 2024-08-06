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

// checks if the IntegrationsModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IntegrationsModel{}

// IntegrationsModel struct for IntegrationsModel
type IntegrationsModel struct {
	// The Integrations of the Product.
	Integrations []IntegrationModel `json:"integrations,omitempty"`
}

// NewIntegrationsModel instantiates a new IntegrationsModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntegrationsModel() *IntegrationsModel {
	this := IntegrationsModel{}
	return &this
}

// NewIntegrationsModelWithDefaults instantiates a new IntegrationsModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntegrationsModelWithDefaults() *IntegrationsModel {
	this := IntegrationsModel{}
	return &this
}

// GetIntegrations returns the Integrations field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IntegrationsModel) GetIntegrations() []IntegrationModel {
	if o == nil {
		var ret []IntegrationModel
		return ret
	}
	return o.Integrations
}

// GetIntegrationsOk returns a tuple with the Integrations field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IntegrationsModel) GetIntegrationsOk() ([]IntegrationModel, bool) {
	if o == nil || IsNil(o.Integrations) {
		return nil, false
	}
	return o.Integrations, true
}

// HasIntegrations returns a boolean if a field has been set.
func (o *IntegrationsModel) HasIntegrations() bool {
	if o != nil && !IsNil(o.Integrations) {
		return true
	}

	return false
}

// SetIntegrations gets a reference to the given []IntegrationModel and assigns it to the Integrations field.
func (o *IntegrationsModel) SetIntegrations(v []IntegrationModel) {
	o.Integrations = v
}

func (o IntegrationsModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IntegrationsModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Integrations != nil {
		toSerialize["integrations"] = o.Integrations
	}
	return toSerialize, nil
}

type NullableIntegrationsModel struct {
	value *IntegrationsModel
	isSet bool
}

func (v NullableIntegrationsModel) Get() *IntegrationsModel {
	return v.value
}

func (v *NullableIntegrationsModel) Set(val *IntegrationsModel) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegrationsModel) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegrationsModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegrationsModel(val *IntegrationsModel) *NullableIntegrationsModel {
	return &NullableIntegrationsModel{value: val, isSet: true}
}

func (v NullableIntegrationsModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegrationsModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


