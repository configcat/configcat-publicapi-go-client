/*
ConfigCat Public Management API

**Base API URL**: https://test-api.configcat.com  If you prefer the swagger documentation, you can find it here: [Swagger UI](https://test-api.configcat.com/swagger).  The purpose of this API is to access the ConfigCat platform programmatically.  You can **Create**, **Read**, **Update** and **Delete** any entities like **Feature Flags, Configs, Environments** or **Products** within ConfigCat.   The API is based on HTTP REST, uses resource-oriented URLs, status codes and supports JSON  and JSON+HAL format. Do not use this API for accessing and evaluating feature flag values. Use the [SDKs instead](https://configcat.com/docs/sdk-reference/overview).   # OpenAPI Specification  The complete specification is publicly available in the following formats:  - [OpenAPI v3](https://test-api.configcat.com/docs/v1/swagger.json) - [Swagger v2](https://test-api.configcat.com/docs/v1/swagger.v2.json)  You can use it to generate client libraries in various languages with [OpenAPI Generator](https://github.com/OpenAPITools/openapi-generator) or [Swagger Codegen](https://swagger.io/tools/swagger-codegen/) to interact with this API.  # Authentication This API uses the [Basic HTTP Authentication Scheme](https://en.wikipedia.org/wiki/Basic_access_authentication).   <!-- ReDoc-Inject: <security-definitions> -->  # Throttling and rate limits All the rate limited API calls are returning information about the current rate limit period in the following HTTP headers:  | Header | Description | | :- | :- | | X-Rate-Limit-Remaining | The maximum number of requests remaining in the current rate limit period. | | X-Rate-Limit-Reset     | The time when the current rate limit period resets.        |  When the rate limit is exceeded by a request, the API returns with a `HTTP 429 - Too many requests` status along with a `Retry-After` HTTP header. 

API version: v1
Contact: support@configcat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configcatpublicapi

import (
	"encoding/json"
	"time"
)

// checks if the SettingValueModelHaljson type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SettingValueModelHaljson{}

// SettingValueModelHaljson struct for SettingValueModelHaljson
type SettingValueModelHaljson struct {
	// The targeting rule collection.
	RolloutRules []RolloutRuleModel `json:"rolloutRules,omitempty"`
	// The percentage rule collection.
	RolloutPercentageItems []RolloutPercentageItemModel `json:"rolloutPercentageItems,omitempty"`
	// The value to serve. It must respect the setting type.
	Value interface{} `json:"value,omitempty"`
	// The last updated date and time when the Feature Flag or Setting.
	UpdatedAt NullableTime `json:"updatedAt,omitempty"`
	// The email of the user who last updated the Feature Flag or Setting.
	LastUpdaterUserEmail NullableString `json:"lastUpdaterUserEmail,omitempty"`
	// The name of the user who last updated the Feature Flag or Setting.
	LastUpdaterUserFullName NullableString `json:"lastUpdaterUserFullName,omitempty"`
	Embedded *SettingFormulaModelHaljsonEmbedded `json:"_embedded,omitempty"`
	ReadOnly *bool `json:"readOnly,omitempty"`
	Links *ConfigSettingFormulasModelHaljsonEmbeddedEnvironmentLinks `json:"_links,omitempty"`
}

// NewSettingValueModelHaljson instantiates a new SettingValueModelHaljson object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSettingValueModelHaljson() *SettingValueModelHaljson {
	this := SettingValueModelHaljson{}
	return &this
}

// NewSettingValueModelHaljsonWithDefaults instantiates a new SettingValueModelHaljson object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSettingValueModelHaljsonWithDefaults() *SettingValueModelHaljson {
	this := SettingValueModelHaljson{}
	return &this
}

// GetRolloutRules returns the RolloutRules field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SettingValueModelHaljson) GetRolloutRules() []RolloutRuleModel {
	if o == nil {
		var ret []RolloutRuleModel
		return ret
	}
	return o.RolloutRules
}

// GetRolloutRulesOk returns a tuple with the RolloutRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SettingValueModelHaljson) GetRolloutRulesOk() ([]RolloutRuleModel, bool) {
	if o == nil || IsNil(o.RolloutRules) {
		return nil, false
	}
	return o.RolloutRules, true
}

// HasRolloutRules returns a boolean if a field has been set.
func (o *SettingValueModelHaljson) HasRolloutRules() bool {
	if o != nil && IsNil(o.RolloutRules) {
		return true
	}

	return false
}

// SetRolloutRules gets a reference to the given []RolloutRuleModel and assigns it to the RolloutRules field.
func (o *SettingValueModelHaljson) SetRolloutRules(v []RolloutRuleModel) {
	o.RolloutRules = v
}

// GetRolloutPercentageItems returns the RolloutPercentageItems field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SettingValueModelHaljson) GetRolloutPercentageItems() []RolloutPercentageItemModel {
	if o == nil {
		var ret []RolloutPercentageItemModel
		return ret
	}
	return o.RolloutPercentageItems
}

// GetRolloutPercentageItemsOk returns a tuple with the RolloutPercentageItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SettingValueModelHaljson) GetRolloutPercentageItemsOk() ([]RolloutPercentageItemModel, bool) {
	if o == nil || IsNil(o.RolloutPercentageItems) {
		return nil, false
	}
	return o.RolloutPercentageItems, true
}

// HasRolloutPercentageItems returns a boolean if a field has been set.
func (o *SettingValueModelHaljson) HasRolloutPercentageItems() bool {
	if o != nil && IsNil(o.RolloutPercentageItems) {
		return true
	}

	return false
}

// SetRolloutPercentageItems gets a reference to the given []RolloutPercentageItemModel and assigns it to the RolloutPercentageItems field.
func (o *SettingValueModelHaljson) SetRolloutPercentageItems(v []RolloutPercentageItemModel) {
	o.RolloutPercentageItems = v
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SettingValueModelHaljson) GetValue() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SettingValueModelHaljson) GetValueOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return &o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *SettingValueModelHaljson) HasValue() bool {
	if o != nil && IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given interface{} and assigns it to the Value field.
func (o *SettingValueModelHaljson) SetValue(v interface{}) {
	o.Value = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SettingValueModelHaljson) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SettingValueModelHaljson) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *SettingValueModelHaljson) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given NullableTime and assigns it to the UpdatedAt field.
func (o *SettingValueModelHaljson) SetUpdatedAt(v time.Time) {
	o.UpdatedAt.Set(&v)
}
// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil
func (o *SettingValueModelHaljson) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
func (o *SettingValueModelHaljson) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

// GetLastUpdaterUserEmail returns the LastUpdaterUserEmail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SettingValueModelHaljson) GetLastUpdaterUserEmail() string {
	if o == nil || IsNil(o.LastUpdaterUserEmail.Get()) {
		var ret string
		return ret
	}
	return *o.LastUpdaterUserEmail.Get()
}

// GetLastUpdaterUserEmailOk returns a tuple with the LastUpdaterUserEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SettingValueModelHaljson) GetLastUpdaterUserEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdaterUserEmail.Get(), o.LastUpdaterUserEmail.IsSet()
}

// HasLastUpdaterUserEmail returns a boolean if a field has been set.
func (o *SettingValueModelHaljson) HasLastUpdaterUserEmail() bool {
	if o != nil && o.LastUpdaterUserEmail.IsSet() {
		return true
	}

	return false
}

// SetLastUpdaterUserEmail gets a reference to the given NullableString and assigns it to the LastUpdaterUserEmail field.
func (o *SettingValueModelHaljson) SetLastUpdaterUserEmail(v string) {
	o.LastUpdaterUserEmail.Set(&v)
}
// SetLastUpdaterUserEmailNil sets the value for LastUpdaterUserEmail to be an explicit nil
func (o *SettingValueModelHaljson) SetLastUpdaterUserEmailNil() {
	o.LastUpdaterUserEmail.Set(nil)
}

// UnsetLastUpdaterUserEmail ensures that no value is present for LastUpdaterUserEmail, not even an explicit nil
func (o *SettingValueModelHaljson) UnsetLastUpdaterUserEmail() {
	o.LastUpdaterUserEmail.Unset()
}

// GetLastUpdaterUserFullName returns the LastUpdaterUserFullName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SettingValueModelHaljson) GetLastUpdaterUserFullName() string {
	if o == nil || IsNil(o.LastUpdaterUserFullName.Get()) {
		var ret string
		return ret
	}
	return *o.LastUpdaterUserFullName.Get()
}

// GetLastUpdaterUserFullNameOk returns a tuple with the LastUpdaterUserFullName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SettingValueModelHaljson) GetLastUpdaterUserFullNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdaterUserFullName.Get(), o.LastUpdaterUserFullName.IsSet()
}

// HasLastUpdaterUserFullName returns a boolean if a field has been set.
func (o *SettingValueModelHaljson) HasLastUpdaterUserFullName() bool {
	if o != nil && o.LastUpdaterUserFullName.IsSet() {
		return true
	}

	return false
}

// SetLastUpdaterUserFullName gets a reference to the given NullableString and assigns it to the LastUpdaterUserFullName field.
func (o *SettingValueModelHaljson) SetLastUpdaterUserFullName(v string) {
	o.LastUpdaterUserFullName.Set(&v)
}
// SetLastUpdaterUserFullNameNil sets the value for LastUpdaterUserFullName to be an explicit nil
func (o *SettingValueModelHaljson) SetLastUpdaterUserFullNameNil() {
	o.LastUpdaterUserFullName.Set(nil)
}

// UnsetLastUpdaterUserFullName ensures that no value is present for LastUpdaterUserFullName, not even an explicit nil
func (o *SettingValueModelHaljson) UnsetLastUpdaterUserFullName() {
	o.LastUpdaterUserFullName.Unset()
}

// GetEmbedded returns the Embedded field value if set, zero value otherwise.
func (o *SettingValueModelHaljson) GetEmbedded() SettingFormulaModelHaljsonEmbedded {
	if o == nil || IsNil(o.Embedded) {
		var ret SettingFormulaModelHaljsonEmbedded
		return ret
	}
	return *o.Embedded
}

// GetEmbeddedOk returns a tuple with the Embedded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingValueModelHaljson) GetEmbeddedOk() (*SettingFormulaModelHaljsonEmbedded, bool) {
	if o == nil || IsNil(o.Embedded) {
		return nil, false
	}
	return o.Embedded, true
}

// HasEmbedded returns a boolean if a field has been set.
func (o *SettingValueModelHaljson) HasEmbedded() bool {
	if o != nil && !IsNil(o.Embedded) {
		return true
	}

	return false
}

// SetEmbedded gets a reference to the given SettingFormulaModelHaljsonEmbedded and assigns it to the Embedded field.
func (o *SettingValueModelHaljson) SetEmbedded(v SettingFormulaModelHaljsonEmbedded) {
	o.Embedded = &v
}

// GetReadOnly returns the ReadOnly field value if set, zero value otherwise.
func (o *SettingValueModelHaljson) GetReadOnly() bool {
	if o == nil || IsNil(o.ReadOnly) {
		var ret bool
		return ret
	}
	return *o.ReadOnly
}

// GetReadOnlyOk returns a tuple with the ReadOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingValueModelHaljson) GetReadOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.ReadOnly) {
		return nil, false
	}
	return o.ReadOnly, true
}

// HasReadOnly returns a boolean if a field has been set.
func (o *SettingValueModelHaljson) HasReadOnly() bool {
	if o != nil && !IsNil(o.ReadOnly) {
		return true
	}

	return false
}

// SetReadOnly gets a reference to the given bool and assigns it to the ReadOnly field.
func (o *SettingValueModelHaljson) SetReadOnly(v bool) {
	o.ReadOnly = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *SettingValueModelHaljson) GetLinks() ConfigSettingFormulasModelHaljsonEmbeddedEnvironmentLinks {
	if o == nil || IsNil(o.Links) {
		var ret ConfigSettingFormulasModelHaljsonEmbeddedEnvironmentLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingValueModelHaljson) GetLinksOk() (*ConfigSettingFormulasModelHaljsonEmbeddedEnvironmentLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *SettingValueModelHaljson) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ConfigSettingFormulasModelHaljsonEmbeddedEnvironmentLinks and assigns it to the Links field.
func (o *SettingValueModelHaljson) SetLinks(v ConfigSettingFormulasModelHaljsonEmbeddedEnvironmentLinks) {
	o.Links = &v
}

func (o SettingValueModelHaljson) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SettingValueModelHaljson) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.RolloutRules != nil {
		toSerialize["rolloutRules"] = o.RolloutRules
	}
	if o.RolloutPercentageItems != nil {
		toSerialize["rolloutPercentageItems"] = o.RolloutPercentageItems
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.UpdatedAt.IsSet() {
		toSerialize["updatedAt"] = o.UpdatedAt.Get()
	}
	if o.LastUpdaterUserEmail.IsSet() {
		toSerialize["lastUpdaterUserEmail"] = o.LastUpdaterUserEmail.Get()
	}
	if o.LastUpdaterUserFullName.IsSet() {
		toSerialize["lastUpdaterUserFullName"] = o.LastUpdaterUserFullName.Get()
	}
	if !IsNil(o.Embedded) {
		toSerialize["_embedded"] = o.Embedded
	}
	if !IsNil(o.ReadOnly) {
		toSerialize["readOnly"] = o.ReadOnly
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableSettingValueModelHaljson struct {
	value *SettingValueModelHaljson
	isSet bool
}

func (v NullableSettingValueModelHaljson) Get() *SettingValueModelHaljson {
	return v.value
}

func (v *NullableSettingValueModelHaljson) Set(val *SettingValueModelHaljson) {
	v.value = val
	v.isSet = true
}

func (v NullableSettingValueModelHaljson) IsSet() bool {
	return v.isSet
}

func (v *NullableSettingValueModelHaljson) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSettingValueModelHaljson(val *SettingValueModelHaljson) *NullableSettingValueModelHaljson {
	return &NullableSettingValueModelHaljson{value: val, isSet: true}
}

func (v NullableSettingValueModelHaljson) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSettingValueModelHaljson) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


