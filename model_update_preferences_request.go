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

// checks if the UpdatePreferencesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdatePreferencesRequest{}

// UpdatePreferencesRequest struct for UpdatePreferencesRequest
type UpdatePreferencesRequest struct {
	// Indicates that a mandatory note is required for saving and publishing.
	ReasonRequired NullableBool `json:"reasonRequired,omitempty"`
	// Determines the Feature Flag key generation mode.
	KeyGenerationMode NullableString `json:"keyGenerationMode,omitempty"`
	// Indicates whether a variation ID's must be shown on the ConfigCat Dashboard.
	ShowVariationId NullableBool `json:"showVariationId,omitempty"`
	// Indicates whether Feature flags and Settings must have a hint.
	MandatorySettingHint NullableBool `json:"mandatorySettingHint,omitempty"`
	// List of Environments where mandatory note must be set before saving and publishing.
	ReasonRequiredEnvironments []UpdateReasonRequiredEnvironmentModel `json:"reasonRequiredEnvironments,omitempty"`
}

// NewUpdatePreferencesRequest instantiates a new UpdatePreferencesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdatePreferencesRequest() *UpdatePreferencesRequest {
	this := UpdatePreferencesRequest{}
	return &this
}

// NewUpdatePreferencesRequestWithDefaults instantiates a new UpdatePreferencesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdatePreferencesRequestWithDefaults() *UpdatePreferencesRequest {
	this := UpdatePreferencesRequest{}
	return &this
}

// GetReasonRequired returns the ReasonRequired field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdatePreferencesRequest) GetReasonRequired() bool {
	if o == nil || IsNil(o.ReasonRequired.Get()) {
		var ret bool
		return ret
	}
	return *o.ReasonRequired.Get()
}

// GetReasonRequiredOk returns a tuple with the ReasonRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdatePreferencesRequest) GetReasonRequiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReasonRequired.Get(), o.ReasonRequired.IsSet()
}

// HasReasonRequired returns a boolean if a field has been set.
func (o *UpdatePreferencesRequest) HasReasonRequired() bool {
	if o != nil && o.ReasonRequired.IsSet() {
		return true
	}

	return false
}

// SetReasonRequired gets a reference to the given NullableBool and assigns it to the ReasonRequired field.
func (o *UpdatePreferencesRequest) SetReasonRequired(v bool) {
	o.ReasonRequired.Set(&v)
}
// SetReasonRequiredNil sets the value for ReasonRequired to be an explicit nil
func (o *UpdatePreferencesRequest) SetReasonRequiredNil() {
	o.ReasonRequired.Set(nil)
}

// UnsetReasonRequired ensures that no value is present for ReasonRequired, not even an explicit nil
func (o *UpdatePreferencesRequest) UnsetReasonRequired() {
	o.ReasonRequired.Unset()
}

// GetKeyGenerationMode returns the KeyGenerationMode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdatePreferencesRequest) GetKeyGenerationMode() string {
	if o == nil || IsNil(o.KeyGenerationMode.Get()) {
		var ret string
		return ret
	}
	return *o.KeyGenerationMode.Get()
}

// GetKeyGenerationModeOk returns a tuple with the KeyGenerationMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdatePreferencesRequest) GetKeyGenerationModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.KeyGenerationMode.Get(), o.KeyGenerationMode.IsSet()
}

// HasKeyGenerationMode returns a boolean if a field has been set.
func (o *UpdatePreferencesRequest) HasKeyGenerationMode() bool {
	if o != nil && o.KeyGenerationMode.IsSet() {
		return true
	}

	return false
}

// SetKeyGenerationMode gets a reference to the given NullableString and assigns it to the KeyGenerationMode field.
func (o *UpdatePreferencesRequest) SetKeyGenerationMode(v string) {
	o.KeyGenerationMode.Set(&v)
}
// SetKeyGenerationModeNil sets the value for KeyGenerationMode to be an explicit nil
func (o *UpdatePreferencesRequest) SetKeyGenerationModeNil() {
	o.KeyGenerationMode.Set(nil)
}

// UnsetKeyGenerationMode ensures that no value is present for KeyGenerationMode, not even an explicit nil
func (o *UpdatePreferencesRequest) UnsetKeyGenerationMode() {
	o.KeyGenerationMode.Unset()
}

// GetShowVariationId returns the ShowVariationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdatePreferencesRequest) GetShowVariationId() bool {
	if o == nil || IsNil(o.ShowVariationId.Get()) {
		var ret bool
		return ret
	}
	return *o.ShowVariationId.Get()
}

// GetShowVariationIdOk returns a tuple with the ShowVariationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdatePreferencesRequest) GetShowVariationIdOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.ShowVariationId.Get(), o.ShowVariationId.IsSet()
}

// HasShowVariationId returns a boolean if a field has been set.
func (o *UpdatePreferencesRequest) HasShowVariationId() bool {
	if o != nil && o.ShowVariationId.IsSet() {
		return true
	}

	return false
}

// SetShowVariationId gets a reference to the given NullableBool and assigns it to the ShowVariationId field.
func (o *UpdatePreferencesRequest) SetShowVariationId(v bool) {
	o.ShowVariationId.Set(&v)
}
// SetShowVariationIdNil sets the value for ShowVariationId to be an explicit nil
func (o *UpdatePreferencesRequest) SetShowVariationIdNil() {
	o.ShowVariationId.Set(nil)
}

// UnsetShowVariationId ensures that no value is present for ShowVariationId, not even an explicit nil
func (o *UpdatePreferencesRequest) UnsetShowVariationId() {
	o.ShowVariationId.Unset()
}

// GetMandatorySettingHint returns the MandatorySettingHint field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdatePreferencesRequest) GetMandatorySettingHint() bool {
	if o == nil || IsNil(o.MandatorySettingHint.Get()) {
		var ret bool
		return ret
	}
	return *o.MandatorySettingHint.Get()
}

// GetMandatorySettingHintOk returns a tuple with the MandatorySettingHint field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdatePreferencesRequest) GetMandatorySettingHintOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.MandatorySettingHint.Get(), o.MandatorySettingHint.IsSet()
}

// HasMandatorySettingHint returns a boolean if a field has been set.
func (o *UpdatePreferencesRequest) HasMandatorySettingHint() bool {
	if o != nil && o.MandatorySettingHint.IsSet() {
		return true
	}

	return false
}

// SetMandatorySettingHint gets a reference to the given NullableBool and assigns it to the MandatorySettingHint field.
func (o *UpdatePreferencesRequest) SetMandatorySettingHint(v bool) {
	o.MandatorySettingHint.Set(&v)
}
// SetMandatorySettingHintNil sets the value for MandatorySettingHint to be an explicit nil
func (o *UpdatePreferencesRequest) SetMandatorySettingHintNil() {
	o.MandatorySettingHint.Set(nil)
}

// UnsetMandatorySettingHint ensures that no value is present for MandatorySettingHint, not even an explicit nil
func (o *UpdatePreferencesRequest) UnsetMandatorySettingHint() {
	o.MandatorySettingHint.Unset()
}

// GetReasonRequiredEnvironments returns the ReasonRequiredEnvironments field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdatePreferencesRequest) GetReasonRequiredEnvironments() []UpdateReasonRequiredEnvironmentModel {
	if o == nil {
		var ret []UpdateReasonRequiredEnvironmentModel
		return ret
	}
	return o.ReasonRequiredEnvironments
}

// GetReasonRequiredEnvironmentsOk returns a tuple with the ReasonRequiredEnvironments field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdatePreferencesRequest) GetReasonRequiredEnvironmentsOk() ([]UpdateReasonRequiredEnvironmentModel, bool) {
	if o == nil || IsNil(o.ReasonRequiredEnvironments) {
		return nil, false
	}
	return o.ReasonRequiredEnvironments, true
}

// HasReasonRequiredEnvironments returns a boolean if a field has been set.
func (o *UpdatePreferencesRequest) HasReasonRequiredEnvironments() bool {
	if o != nil && IsNil(o.ReasonRequiredEnvironments) {
		return true
	}

	return false
}

// SetReasonRequiredEnvironments gets a reference to the given []UpdateReasonRequiredEnvironmentModel and assigns it to the ReasonRequiredEnvironments field.
func (o *UpdatePreferencesRequest) SetReasonRequiredEnvironments(v []UpdateReasonRequiredEnvironmentModel) {
	o.ReasonRequiredEnvironments = v
}

func (o UpdatePreferencesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdatePreferencesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ReasonRequired.IsSet() {
		toSerialize["reasonRequired"] = o.ReasonRequired.Get()
	}
	if o.KeyGenerationMode.IsSet() {
		toSerialize["keyGenerationMode"] = o.KeyGenerationMode.Get()
	}
	if o.ShowVariationId.IsSet() {
		toSerialize["showVariationId"] = o.ShowVariationId.Get()
	}
	if o.MandatorySettingHint.IsSet() {
		toSerialize["mandatorySettingHint"] = o.MandatorySettingHint.Get()
	}
	if o.ReasonRequiredEnvironments != nil {
		toSerialize["reasonRequiredEnvironments"] = o.ReasonRequiredEnvironments
	}
	return toSerialize, nil
}

type NullableUpdatePreferencesRequest struct {
	value *UpdatePreferencesRequest
	isSet bool
}

func (v NullableUpdatePreferencesRequest) Get() *UpdatePreferencesRequest {
	return v.value
}

func (v *NullableUpdatePreferencesRequest) Set(val *UpdatePreferencesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdatePreferencesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdatePreferencesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdatePreferencesRequest(val *UpdatePreferencesRequest) *NullableUpdatePreferencesRequest {
	return &NullableUpdatePreferencesRequest{value: val, isSet: true}
}

func (v NullableUpdatePreferencesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdatePreferencesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


