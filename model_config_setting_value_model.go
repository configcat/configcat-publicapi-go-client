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
	"time"
)

// checks if the ConfigSettingValueModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigSettingValueModel{}

// ConfigSettingValueModel struct for ConfigSettingValueModel
type ConfigSettingValueModel struct {
	// The targeting rule collection.
	RolloutRules []RolloutRuleModel `json:"rolloutRules,omitempty"`
	// The percentage rule collection.
	RolloutPercentageItems []RolloutPercentageItemModel `json:"rolloutPercentageItems,omitempty"`
	// The value to serve. It must respect the setting type.
	Value interface{} `json:"value,omitempty"`
	Setting *SettingDataModel `json:"setting,omitempty"`
	UpdatedAt NullableTime `json:"updatedAt,omitempty"`
	LastUpdaterUserEmail NullableString `json:"lastUpdaterUserEmail,omitempty"`
	LastUpdaterUserFullName NullableString `json:"lastUpdaterUserFullName,omitempty"`
	IntegrationLinks []IntegrationLinkModel `json:"integrationLinks,omitempty"`
	SettingTags []SettingTagModel `json:"settingTags,omitempty"`
}

// NewConfigSettingValueModel instantiates a new ConfigSettingValueModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigSettingValueModel() *ConfigSettingValueModel {
	this := ConfigSettingValueModel{}
	return &this
}

// NewConfigSettingValueModelWithDefaults instantiates a new ConfigSettingValueModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigSettingValueModelWithDefaults() *ConfigSettingValueModel {
	this := ConfigSettingValueModel{}
	return &this
}

// GetRolloutRules returns the RolloutRules field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConfigSettingValueModel) GetRolloutRules() []RolloutRuleModel {
	if o == nil {
		var ret []RolloutRuleModel
		return ret
	}
	return o.RolloutRules
}

// GetRolloutRulesOk returns a tuple with the RolloutRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConfigSettingValueModel) GetRolloutRulesOk() ([]RolloutRuleModel, bool) {
	if o == nil || IsNil(o.RolloutRules) {
		return nil, false
	}
	return o.RolloutRules, true
}

// HasRolloutRules returns a boolean if a field has been set.
func (o *ConfigSettingValueModel) HasRolloutRules() bool {
	if o != nil && IsNil(o.RolloutRules) {
		return true
	}

	return false
}

// SetRolloutRules gets a reference to the given []RolloutRuleModel and assigns it to the RolloutRules field.
func (o *ConfigSettingValueModel) SetRolloutRules(v []RolloutRuleModel) {
	o.RolloutRules = v
}

// GetRolloutPercentageItems returns the RolloutPercentageItems field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConfigSettingValueModel) GetRolloutPercentageItems() []RolloutPercentageItemModel {
	if o == nil {
		var ret []RolloutPercentageItemModel
		return ret
	}
	return o.RolloutPercentageItems
}

// GetRolloutPercentageItemsOk returns a tuple with the RolloutPercentageItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConfigSettingValueModel) GetRolloutPercentageItemsOk() ([]RolloutPercentageItemModel, bool) {
	if o == nil || IsNil(o.RolloutPercentageItems) {
		return nil, false
	}
	return o.RolloutPercentageItems, true
}

// HasRolloutPercentageItems returns a boolean if a field has been set.
func (o *ConfigSettingValueModel) HasRolloutPercentageItems() bool {
	if o != nil && IsNil(o.RolloutPercentageItems) {
		return true
	}

	return false
}

// SetRolloutPercentageItems gets a reference to the given []RolloutPercentageItemModel and assigns it to the RolloutPercentageItems field.
func (o *ConfigSettingValueModel) SetRolloutPercentageItems(v []RolloutPercentageItemModel) {
	o.RolloutPercentageItems = v
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConfigSettingValueModel) GetValue() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConfigSettingValueModel) GetValueOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return &o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ConfigSettingValueModel) HasValue() bool {
	if o != nil && IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given interface{} and assigns it to the Value field.
func (o *ConfigSettingValueModel) SetValue(v interface{}) {
	o.Value = v
}

// GetSetting returns the Setting field value if set, zero value otherwise.
func (o *ConfigSettingValueModel) GetSetting() SettingDataModel {
	if o == nil || IsNil(o.Setting) {
		var ret SettingDataModel
		return ret
	}
	return *o.Setting
}

// GetSettingOk returns a tuple with the Setting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigSettingValueModel) GetSettingOk() (*SettingDataModel, bool) {
	if o == nil || IsNil(o.Setting) {
		return nil, false
	}
	return o.Setting, true
}

// HasSetting returns a boolean if a field has been set.
func (o *ConfigSettingValueModel) HasSetting() bool {
	if o != nil && !IsNil(o.Setting) {
		return true
	}

	return false
}

// SetSetting gets a reference to the given SettingDataModel and assigns it to the Setting field.
func (o *ConfigSettingValueModel) SetSetting(v SettingDataModel) {
	o.Setting = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConfigSettingValueModel) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConfigSettingValueModel) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ConfigSettingValueModel) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given NullableTime and assigns it to the UpdatedAt field.
func (o *ConfigSettingValueModel) SetUpdatedAt(v time.Time) {
	o.UpdatedAt.Set(&v)
}
// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil
func (o *ConfigSettingValueModel) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
func (o *ConfigSettingValueModel) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

// GetLastUpdaterUserEmail returns the LastUpdaterUserEmail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConfigSettingValueModel) GetLastUpdaterUserEmail() string {
	if o == nil || IsNil(o.LastUpdaterUserEmail.Get()) {
		var ret string
		return ret
	}
	return *o.LastUpdaterUserEmail.Get()
}

// GetLastUpdaterUserEmailOk returns a tuple with the LastUpdaterUserEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConfigSettingValueModel) GetLastUpdaterUserEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdaterUserEmail.Get(), o.LastUpdaterUserEmail.IsSet()
}

// HasLastUpdaterUserEmail returns a boolean if a field has been set.
func (o *ConfigSettingValueModel) HasLastUpdaterUserEmail() bool {
	if o != nil && o.LastUpdaterUserEmail.IsSet() {
		return true
	}

	return false
}

// SetLastUpdaterUserEmail gets a reference to the given NullableString and assigns it to the LastUpdaterUserEmail field.
func (o *ConfigSettingValueModel) SetLastUpdaterUserEmail(v string) {
	o.LastUpdaterUserEmail.Set(&v)
}
// SetLastUpdaterUserEmailNil sets the value for LastUpdaterUserEmail to be an explicit nil
func (o *ConfigSettingValueModel) SetLastUpdaterUserEmailNil() {
	o.LastUpdaterUserEmail.Set(nil)
}

// UnsetLastUpdaterUserEmail ensures that no value is present for LastUpdaterUserEmail, not even an explicit nil
func (o *ConfigSettingValueModel) UnsetLastUpdaterUserEmail() {
	o.LastUpdaterUserEmail.Unset()
}

// GetLastUpdaterUserFullName returns the LastUpdaterUserFullName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConfigSettingValueModel) GetLastUpdaterUserFullName() string {
	if o == nil || IsNil(o.LastUpdaterUserFullName.Get()) {
		var ret string
		return ret
	}
	return *o.LastUpdaterUserFullName.Get()
}

// GetLastUpdaterUserFullNameOk returns a tuple with the LastUpdaterUserFullName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConfigSettingValueModel) GetLastUpdaterUserFullNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdaterUserFullName.Get(), o.LastUpdaterUserFullName.IsSet()
}

// HasLastUpdaterUserFullName returns a boolean if a field has been set.
func (o *ConfigSettingValueModel) HasLastUpdaterUserFullName() bool {
	if o != nil && o.LastUpdaterUserFullName.IsSet() {
		return true
	}

	return false
}

// SetLastUpdaterUserFullName gets a reference to the given NullableString and assigns it to the LastUpdaterUserFullName field.
func (o *ConfigSettingValueModel) SetLastUpdaterUserFullName(v string) {
	o.LastUpdaterUserFullName.Set(&v)
}
// SetLastUpdaterUserFullNameNil sets the value for LastUpdaterUserFullName to be an explicit nil
func (o *ConfigSettingValueModel) SetLastUpdaterUserFullNameNil() {
	o.LastUpdaterUserFullName.Set(nil)
}

// UnsetLastUpdaterUserFullName ensures that no value is present for LastUpdaterUserFullName, not even an explicit nil
func (o *ConfigSettingValueModel) UnsetLastUpdaterUserFullName() {
	o.LastUpdaterUserFullName.Unset()
}

// GetIntegrationLinks returns the IntegrationLinks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConfigSettingValueModel) GetIntegrationLinks() []IntegrationLinkModel {
	if o == nil {
		var ret []IntegrationLinkModel
		return ret
	}
	return o.IntegrationLinks
}

// GetIntegrationLinksOk returns a tuple with the IntegrationLinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConfigSettingValueModel) GetIntegrationLinksOk() ([]IntegrationLinkModel, bool) {
	if o == nil || IsNil(o.IntegrationLinks) {
		return nil, false
	}
	return o.IntegrationLinks, true
}

// HasIntegrationLinks returns a boolean if a field has been set.
func (o *ConfigSettingValueModel) HasIntegrationLinks() bool {
	if o != nil && IsNil(o.IntegrationLinks) {
		return true
	}

	return false
}

// SetIntegrationLinks gets a reference to the given []IntegrationLinkModel and assigns it to the IntegrationLinks field.
func (o *ConfigSettingValueModel) SetIntegrationLinks(v []IntegrationLinkModel) {
	o.IntegrationLinks = v
}

// GetSettingTags returns the SettingTags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConfigSettingValueModel) GetSettingTags() []SettingTagModel {
	if o == nil {
		var ret []SettingTagModel
		return ret
	}
	return o.SettingTags
}

// GetSettingTagsOk returns a tuple with the SettingTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConfigSettingValueModel) GetSettingTagsOk() ([]SettingTagModel, bool) {
	if o == nil || IsNil(o.SettingTags) {
		return nil, false
	}
	return o.SettingTags, true
}

// HasSettingTags returns a boolean if a field has been set.
func (o *ConfigSettingValueModel) HasSettingTags() bool {
	if o != nil && IsNil(o.SettingTags) {
		return true
	}

	return false
}

// SetSettingTags gets a reference to the given []SettingTagModel and assigns it to the SettingTags field.
func (o *ConfigSettingValueModel) SetSettingTags(v []SettingTagModel) {
	o.SettingTags = v
}

func (o ConfigSettingValueModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigSettingValueModel) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Setting) {
		toSerialize["setting"] = o.Setting
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
	if o.IntegrationLinks != nil {
		toSerialize["integrationLinks"] = o.IntegrationLinks
	}
	if o.SettingTags != nil {
		toSerialize["settingTags"] = o.SettingTags
	}
	return toSerialize, nil
}

type NullableConfigSettingValueModel struct {
	value *ConfigSettingValueModel
	isSet bool
}

func (v NullableConfigSettingValueModel) Get() *ConfigSettingValueModel {
	return v.value
}

func (v *NullableConfigSettingValueModel) Set(val *ConfigSettingValueModel) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigSettingValueModel) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigSettingValueModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigSettingValueModel(val *ConfigSettingValueModel) *NullableConfigSettingValueModel {
	return &NullableConfigSettingValueModel{value: val, isSet: true}
}

func (v NullableConfigSettingValueModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigSettingValueModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


