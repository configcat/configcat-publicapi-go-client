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

// checks if the IntegrationLinkDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IntegrationLinkDetail{}

// IntegrationLinkDetail struct for IntegrationLinkDetail
type IntegrationLinkDetail struct {
	Product *ProductModel `json:"product,omitempty"`
	Config *ConfigModel `json:"config,omitempty"`
	Environment *EnvironmentModel `json:"environment,omitempty"`
	Setting *SettingDataModel `json:"setting,omitempty"`
	ReadOnly *bool `json:"readOnly,omitempty"`
	Status NullableString `json:"status,omitempty"`
}

// NewIntegrationLinkDetail instantiates a new IntegrationLinkDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntegrationLinkDetail() *IntegrationLinkDetail {
	this := IntegrationLinkDetail{}
	return &this
}

// NewIntegrationLinkDetailWithDefaults instantiates a new IntegrationLinkDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntegrationLinkDetailWithDefaults() *IntegrationLinkDetail {
	this := IntegrationLinkDetail{}
	return &this
}

// GetProduct returns the Product field value if set, zero value otherwise.
func (o *IntegrationLinkDetail) GetProduct() ProductModel {
	if o == nil || IsNil(o.Product) {
		var ret ProductModel
		return ret
	}
	return *o.Product
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationLinkDetail) GetProductOk() (*ProductModel, bool) {
	if o == nil || IsNil(o.Product) {
		return nil, false
	}
	return o.Product, true
}

// HasProduct returns a boolean if a field has been set.
func (o *IntegrationLinkDetail) HasProduct() bool {
	if o != nil && !IsNil(o.Product) {
		return true
	}

	return false
}

// SetProduct gets a reference to the given ProductModel and assigns it to the Product field.
func (o *IntegrationLinkDetail) SetProduct(v ProductModel) {
	o.Product = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *IntegrationLinkDetail) GetConfig() ConfigModel {
	if o == nil || IsNil(o.Config) {
		var ret ConfigModel
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationLinkDetail) GetConfigOk() (*ConfigModel, bool) {
	if o == nil || IsNil(o.Config) {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *IntegrationLinkDetail) HasConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given ConfigModel and assigns it to the Config field.
func (o *IntegrationLinkDetail) SetConfig(v ConfigModel) {
	o.Config = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *IntegrationLinkDetail) GetEnvironment() EnvironmentModel {
	if o == nil || IsNil(o.Environment) {
		var ret EnvironmentModel
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationLinkDetail) GetEnvironmentOk() (*EnvironmentModel, bool) {
	if o == nil || IsNil(o.Environment) {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *IntegrationLinkDetail) HasEnvironment() bool {
	if o != nil && !IsNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given EnvironmentModel and assigns it to the Environment field.
func (o *IntegrationLinkDetail) SetEnvironment(v EnvironmentModel) {
	o.Environment = &v
}

// GetSetting returns the Setting field value if set, zero value otherwise.
func (o *IntegrationLinkDetail) GetSetting() SettingDataModel {
	if o == nil || IsNil(o.Setting) {
		var ret SettingDataModel
		return ret
	}
	return *o.Setting
}

// GetSettingOk returns a tuple with the Setting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationLinkDetail) GetSettingOk() (*SettingDataModel, bool) {
	if o == nil || IsNil(o.Setting) {
		return nil, false
	}
	return o.Setting, true
}

// HasSetting returns a boolean if a field has been set.
func (o *IntegrationLinkDetail) HasSetting() bool {
	if o != nil && !IsNil(o.Setting) {
		return true
	}

	return false
}

// SetSetting gets a reference to the given SettingDataModel and assigns it to the Setting field.
func (o *IntegrationLinkDetail) SetSetting(v SettingDataModel) {
	o.Setting = &v
}

// GetReadOnly returns the ReadOnly field value if set, zero value otherwise.
func (o *IntegrationLinkDetail) GetReadOnly() bool {
	if o == nil || IsNil(o.ReadOnly) {
		var ret bool
		return ret
	}
	return *o.ReadOnly
}

// GetReadOnlyOk returns a tuple with the ReadOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationLinkDetail) GetReadOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.ReadOnly) {
		return nil, false
	}
	return o.ReadOnly, true
}

// HasReadOnly returns a boolean if a field has been set.
func (o *IntegrationLinkDetail) HasReadOnly() bool {
	if o != nil && !IsNil(o.ReadOnly) {
		return true
	}

	return false
}

// SetReadOnly gets a reference to the given bool and assigns it to the ReadOnly field.
func (o *IntegrationLinkDetail) SetReadOnly(v bool) {
	o.ReadOnly = &v
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IntegrationLinkDetail) GetStatus() string {
	if o == nil || IsNil(o.Status.Get()) {
		var ret string
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IntegrationLinkDetail) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *IntegrationLinkDetail) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableString and assigns it to the Status field.
func (o *IntegrationLinkDetail) SetStatus(v string) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *IntegrationLinkDetail) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *IntegrationLinkDetail) UnsetStatus() {
	o.Status.Unset()
}

func (o IntegrationLinkDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IntegrationLinkDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Product) {
		toSerialize["product"] = o.Product
	}
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	if !IsNil(o.Environment) {
		toSerialize["environment"] = o.Environment
	}
	if !IsNil(o.Setting) {
		toSerialize["setting"] = o.Setting
	}
	if !IsNil(o.ReadOnly) {
		toSerialize["readOnly"] = o.ReadOnly
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	return toSerialize, nil
}

type NullableIntegrationLinkDetail struct {
	value *IntegrationLinkDetail
	isSet bool
}

func (v NullableIntegrationLinkDetail) Get() *IntegrationLinkDetail {
	return v.value
}

func (v *NullableIntegrationLinkDetail) Set(val *IntegrationLinkDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegrationLinkDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegrationLinkDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegrationLinkDetail(val *IntegrationLinkDetail) *NullableIntegrationLinkDetail {
	return &NullableIntegrationLinkDetail{value: val, isSet: true}
}

func (v NullableIntegrationLinkDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegrationLinkDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


