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

// checks if the PreferencesModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PreferencesModel{}

// PreferencesModel struct for PreferencesModel
type PreferencesModel struct {
	// Indicates that a mandatory note required for saving and publishing.
	ReasonRequired *bool `json:"reasonRequired,omitempty"`
	KeyGenerationMode *KeyGenerationMode `json:"keyGenerationMode,omitempty"`
	// Indicates whether a variation ID's must be shown on the ConfigCat Dashboard.
	ShowVariationId *bool `json:"showVariationId,omitempty"`
	// List of Environments where mandatory note must be set before saving and publishing.
	ReasonRequiredEnvironments []ReasonRequiredEnvironmentModel `json:"reasonRequiredEnvironments,omitempty"`
	// Indicates whether Feature flags and Settings must have a hint.
	MandatorySettingHint *bool `json:"mandatorySettingHint,omitempty"`
}

// NewPreferencesModel instantiates a new PreferencesModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPreferencesModel() *PreferencesModel {
	this := PreferencesModel{}
	return &this
}

// NewPreferencesModelWithDefaults instantiates a new PreferencesModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPreferencesModelWithDefaults() *PreferencesModel {
	this := PreferencesModel{}
	return &this
}

// GetReasonRequired returns the ReasonRequired field value if set, zero value otherwise.
func (o *PreferencesModel) GetReasonRequired() bool {
	if o == nil || IsNil(o.ReasonRequired) {
		var ret bool
		return ret
	}
	return *o.ReasonRequired
}

// GetReasonRequiredOk returns a tuple with the ReasonRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreferencesModel) GetReasonRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.ReasonRequired) {
		return nil, false
	}
	return o.ReasonRequired, true
}

// HasReasonRequired returns a boolean if a field has been set.
func (o *PreferencesModel) HasReasonRequired() bool {
	if o != nil && !IsNil(o.ReasonRequired) {
		return true
	}

	return false
}

// SetReasonRequired gets a reference to the given bool and assigns it to the ReasonRequired field.
func (o *PreferencesModel) SetReasonRequired(v bool) {
	o.ReasonRequired = &v
}

// GetKeyGenerationMode returns the KeyGenerationMode field value if set, zero value otherwise.
func (o *PreferencesModel) GetKeyGenerationMode() KeyGenerationMode {
	if o == nil || IsNil(o.KeyGenerationMode) {
		var ret KeyGenerationMode
		return ret
	}
	return *o.KeyGenerationMode
}

// GetKeyGenerationModeOk returns a tuple with the KeyGenerationMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreferencesModel) GetKeyGenerationModeOk() (*KeyGenerationMode, bool) {
	if o == nil || IsNil(o.KeyGenerationMode) {
		return nil, false
	}
	return o.KeyGenerationMode, true
}

// HasKeyGenerationMode returns a boolean if a field has been set.
func (o *PreferencesModel) HasKeyGenerationMode() bool {
	if o != nil && !IsNil(o.KeyGenerationMode) {
		return true
	}

	return false
}

// SetKeyGenerationMode gets a reference to the given KeyGenerationMode and assigns it to the KeyGenerationMode field.
func (o *PreferencesModel) SetKeyGenerationMode(v KeyGenerationMode) {
	o.KeyGenerationMode = &v
}

// GetShowVariationId returns the ShowVariationId field value if set, zero value otherwise.
func (o *PreferencesModel) GetShowVariationId() bool {
	if o == nil || IsNil(o.ShowVariationId) {
		var ret bool
		return ret
	}
	return *o.ShowVariationId
}

// GetShowVariationIdOk returns a tuple with the ShowVariationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreferencesModel) GetShowVariationIdOk() (*bool, bool) {
	if o == nil || IsNil(o.ShowVariationId) {
		return nil, false
	}
	return o.ShowVariationId, true
}

// HasShowVariationId returns a boolean if a field has been set.
func (o *PreferencesModel) HasShowVariationId() bool {
	if o != nil && !IsNil(o.ShowVariationId) {
		return true
	}

	return false
}

// SetShowVariationId gets a reference to the given bool and assigns it to the ShowVariationId field.
func (o *PreferencesModel) SetShowVariationId(v bool) {
	o.ShowVariationId = &v
}

// GetReasonRequiredEnvironments returns the ReasonRequiredEnvironments field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PreferencesModel) GetReasonRequiredEnvironments() []ReasonRequiredEnvironmentModel {
	if o == nil {
		var ret []ReasonRequiredEnvironmentModel
		return ret
	}
	return o.ReasonRequiredEnvironments
}

// GetReasonRequiredEnvironmentsOk returns a tuple with the ReasonRequiredEnvironments field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PreferencesModel) GetReasonRequiredEnvironmentsOk() ([]ReasonRequiredEnvironmentModel, bool) {
	if o == nil || IsNil(o.ReasonRequiredEnvironments) {
		return nil, false
	}
	return o.ReasonRequiredEnvironments, true
}

// HasReasonRequiredEnvironments returns a boolean if a field has been set.
func (o *PreferencesModel) HasReasonRequiredEnvironments() bool {
	if o != nil && !IsNil(o.ReasonRequiredEnvironments) {
		return true
	}

	return false
}

// SetReasonRequiredEnvironments gets a reference to the given []ReasonRequiredEnvironmentModel and assigns it to the ReasonRequiredEnvironments field.
func (o *PreferencesModel) SetReasonRequiredEnvironments(v []ReasonRequiredEnvironmentModel) {
	o.ReasonRequiredEnvironments = v
}

// GetMandatorySettingHint returns the MandatorySettingHint field value if set, zero value otherwise.
func (o *PreferencesModel) GetMandatorySettingHint() bool {
	if o == nil || IsNil(o.MandatorySettingHint) {
		var ret bool
		return ret
	}
	return *o.MandatorySettingHint
}

// GetMandatorySettingHintOk returns a tuple with the MandatorySettingHint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreferencesModel) GetMandatorySettingHintOk() (*bool, bool) {
	if o == nil || IsNil(o.MandatorySettingHint) {
		return nil, false
	}
	return o.MandatorySettingHint, true
}

// HasMandatorySettingHint returns a boolean if a field has been set.
func (o *PreferencesModel) HasMandatorySettingHint() bool {
	if o != nil && !IsNil(o.MandatorySettingHint) {
		return true
	}

	return false
}

// SetMandatorySettingHint gets a reference to the given bool and assigns it to the MandatorySettingHint field.
func (o *PreferencesModel) SetMandatorySettingHint(v bool) {
	o.MandatorySettingHint = &v
}

func (o PreferencesModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PreferencesModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ReasonRequired) {
		toSerialize["reasonRequired"] = o.ReasonRequired
	}
	if !IsNil(o.KeyGenerationMode) {
		toSerialize["keyGenerationMode"] = o.KeyGenerationMode
	}
	if !IsNil(o.ShowVariationId) {
		toSerialize["showVariationId"] = o.ShowVariationId
	}
	if o.ReasonRequiredEnvironments != nil {
		toSerialize["reasonRequiredEnvironments"] = o.ReasonRequiredEnvironments
	}
	if !IsNil(o.MandatorySettingHint) {
		toSerialize["mandatorySettingHint"] = o.MandatorySettingHint
	}
	return toSerialize, nil
}

type NullablePreferencesModel struct {
	value *PreferencesModel
	isSet bool
}

func (v NullablePreferencesModel) Get() *PreferencesModel {
	return v.value
}

func (v *NullablePreferencesModel) Set(val *PreferencesModel) {
	v.value = val
	v.isSet = true
}

func (v NullablePreferencesModel) IsSet() bool {
	return v.isSet
}

func (v *NullablePreferencesModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePreferencesModel(val *PreferencesModel) *NullablePreferencesModel {
	return &NullablePreferencesModel{value: val, isSet: true}
}

func (v NullablePreferencesModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePreferencesModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


