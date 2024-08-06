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

// checks if the IntegrationModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IntegrationModel{}

// IntegrationModel Details of the Integration.
type IntegrationModel struct {
	// Identifier of the Integration.
	IntegrationId *string `json:"integrationId,omitempty"`
	// Name of the Integration.
	Name NullableString `json:"name,omitempty"`
	IntegrationType *IntegrationType `json:"integrationType,omitempty"`
	// Parameters of the Integration.
	Parameters map[string]string `json:"parameters,omitempty"`
	// List of Environment IDs that are connected with this Integration. If the list is empty, all of the Environments are connected.
	IntegrationEnvironmentIds []string `json:"integrationEnvironmentIds,omitempty"`
	// List of Config IDs that are connected with this Integration. If the list is empty, all of the Configs are connected.
	IntegrationConfigIds []string `json:"integrationConfigIds,omitempty"`
}

// NewIntegrationModel instantiates a new IntegrationModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntegrationModel() *IntegrationModel {
	this := IntegrationModel{}
	return &this
}

// NewIntegrationModelWithDefaults instantiates a new IntegrationModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntegrationModelWithDefaults() *IntegrationModel {
	this := IntegrationModel{}
	return &this
}

// GetIntegrationId returns the IntegrationId field value if set, zero value otherwise.
func (o *IntegrationModel) GetIntegrationId() string {
	if o == nil || IsNil(o.IntegrationId) {
		var ret string
		return ret
	}
	return *o.IntegrationId
}

// GetIntegrationIdOk returns a tuple with the IntegrationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationModel) GetIntegrationIdOk() (*string, bool) {
	if o == nil || IsNil(o.IntegrationId) {
		return nil, false
	}
	return o.IntegrationId, true
}

// HasIntegrationId returns a boolean if a field has been set.
func (o *IntegrationModel) HasIntegrationId() bool {
	if o != nil && !IsNil(o.IntegrationId) {
		return true
	}

	return false
}

// SetIntegrationId gets a reference to the given string and assigns it to the IntegrationId field.
func (o *IntegrationModel) SetIntegrationId(v string) {
	o.IntegrationId = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IntegrationModel) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IntegrationModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *IntegrationModel) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *IntegrationModel) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *IntegrationModel) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *IntegrationModel) UnsetName() {
	o.Name.Unset()
}

// GetIntegrationType returns the IntegrationType field value if set, zero value otherwise.
func (o *IntegrationModel) GetIntegrationType() IntegrationType {
	if o == nil || IsNil(o.IntegrationType) {
		var ret IntegrationType
		return ret
	}
	return *o.IntegrationType
}

// GetIntegrationTypeOk returns a tuple with the IntegrationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationModel) GetIntegrationTypeOk() (*IntegrationType, bool) {
	if o == nil || IsNil(o.IntegrationType) {
		return nil, false
	}
	return o.IntegrationType, true
}

// HasIntegrationType returns a boolean if a field has been set.
func (o *IntegrationModel) HasIntegrationType() bool {
	if o != nil && !IsNil(o.IntegrationType) {
		return true
	}

	return false
}

// SetIntegrationType gets a reference to the given IntegrationType and assigns it to the IntegrationType field.
func (o *IntegrationModel) SetIntegrationType(v IntegrationType) {
	o.IntegrationType = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IntegrationModel) GetParameters() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IntegrationModel) GetParametersOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Parameters) {
		return nil, false
	}
	return &o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *IntegrationModel) HasParameters() bool {
	if o != nil && IsNil(o.Parameters) {
		return true
	}

	return false
}

// SetParameters gets a reference to the given map[string]string and assigns it to the Parameters field.
func (o *IntegrationModel) SetParameters(v map[string]string) {
	o.Parameters = v
}

// GetIntegrationEnvironmentIds returns the IntegrationEnvironmentIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IntegrationModel) GetIntegrationEnvironmentIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.IntegrationEnvironmentIds
}

// GetIntegrationEnvironmentIdsOk returns a tuple with the IntegrationEnvironmentIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IntegrationModel) GetIntegrationEnvironmentIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.IntegrationEnvironmentIds) {
		return nil, false
	}
	return o.IntegrationEnvironmentIds, true
}

// HasIntegrationEnvironmentIds returns a boolean if a field has been set.
func (o *IntegrationModel) HasIntegrationEnvironmentIds() bool {
	if o != nil && IsNil(o.IntegrationEnvironmentIds) {
		return true
	}

	return false
}

// SetIntegrationEnvironmentIds gets a reference to the given []string and assigns it to the IntegrationEnvironmentIds field.
func (o *IntegrationModel) SetIntegrationEnvironmentIds(v []string) {
	o.IntegrationEnvironmentIds = v
}

// GetIntegrationConfigIds returns the IntegrationConfigIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IntegrationModel) GetIntegrationConfigIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.IntegrationConfigIds
}

// GetIntegrationConfigIdsOk returns a tuple with the IntegrationConfigIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IntegrationModel) GetIntegrationConfigIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.IntegrationConfigIds) {
		return nil, false
	}
	return o.IntegrationConfigIds, true
}

// HasIntegrationConfigIds returns a boolean if a field has been set.
func (o *IntegrationModel) HasIntegrationConfigIds() bool {
	if o != nil && IsNil(o.IntegrationConfigIds) {
		return true
	}

	return false
}

// SetIntegrationConfigIds gets a reference to the given []string and assigns it to the IntegrationConfigIds field.
func (o *IntegrationModel) SetIntegrationConfigIds(v []string) {
	o.IntegrationConfigIds = v
}

func (o IntegrationModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IntegrationModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IntegrationId) {
		toSerialize["integrationId"] = o.IntegrationId
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if !IsNil(o.IntegrationType) {
		toSerialize["integrationType"] = o.IntegrationType
	}
	if o.Parameters != nil {
		toSerialize["parameters"] = o.Parameters
	}
	if o.IntegrationEnvironmentIds != nil {
		toSerialize["integrationEnvironmentIds"] = o.IntegrationEnvironmentIds
	}
	if o.IntegrationConfigIds != nil {
		toSerialize["integrationConfigIds"] = o.IntegrationConfigIds
	}
	return toSerialize, nil
}

type NullableIntegrationModel struct {
	value *IntegrationModel
	isSet bool
}

func (v NullableIntegrationModel) Get() *IntegrationModel {
	return v.value
}

func (v *NullableIntegrationModel) Set(val *IntegrationModel) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegrationModel) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegrationModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegrationModel(val *IntegrationModel) *NullableIntegrationModel {
	return &NullableIntegrationModel{value: val, isSet: true}
}

func (v NullableIntegrationModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegrationModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


