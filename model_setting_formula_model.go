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
	"time"
)

// checks if the SettingFormulaModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SettingFormulaModel{}

// SettingFormulaModel struct for SettingFormulaModel
type SettingFormulaModel struct {
	LastVersionId *string `json:"lastVersionId,omitempty"`
	DefaultValue *ValueModel `json:"defaultValue,omitempty"`
	// The targeting rules of the Feature Flag or Setting.
	TargetingRules []TargetingRuleModel `json:"targetingRules,omitempty"`
	Setting *SettingDataModel `json:"setting,omitempty"`
	// The last updated date and time when the Feature Flag or Setting.
	UpdatedAt NullableTime `json:"updatedAt,omitempty"`
	// The user attribute used for percentage evaluation. If not set, it defaults to the `Identifier` user object attribute.
	PercentageEvaluationAttribute NullableString `json:"percentageEvaluationAttribute,omitempty"`
	// The email of the user who last updated the Feature Flag or Setting.
	LastUpdaterUserEmail NullableString `json:"lastUpdaterUserEmail,omitempty"`
	// The name of the user who last updated the Feature Flag or Setting.
	LastUpdaterUserFullName NullableString `json:"lastUpdaterUserFullName,omitempty"`
	// The integration links attached to the Feature Flag or Setting.
	IntegrationLinks []IntegrationLinkModel `json:"integrationLinks,omitempty"`
	// The tags attached to the Feature Flag or Setting.
	SettingTags []SettingTagModel `json:"settingTags,omitempty"`
	// List of Feature Flag and Setting IDs where the actual Feature Flag or Setting is prerequisite.
	SettingIdsWherePrerequisite []int32 `json:"settingIdsWherePrerequisite,omitempty"`
	Config *ConfigModel `json:"config,omitempty"`
	Environment *EnvironmentModel `json:"environment,omitempty"`
	ReadOnly *bool `json:"readOnly,omitempty"`
	FeatureFlagLimitations *FeatureFlagLimitations `json:"featureFlagLimitations,omitempty"`
}

// NewSettingFormulaModel instantiates a new SettingFormulaModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSettingFormulaModel() *SettingFormulaModel {
	this := SettingFormulaModel{}
	return &this
}

// NewSettingFormulaModelWithDefaults instantiates a new SettingFormulaModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSettingFormulaModelWithDefaults() *SettingFormulaModel {
	this := SettingFormulaModel{}
	return &this
}

// GetLastVersionId returns the LastVersionId field value if set, zero value otherwise.
func (o *SettingFormulaModel) GetLastVersionId() string {
	if o == nil || IsNil(o.LastVersionId) {
		var ret string
		return ret
	}
	return *o.LastVersionId
}

// GetLastVersionIdOk returns a tuple with the LastVersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingFormulaModel) GetLastVersionIdOk() (*string, bool) {
	if o == nil || IsNil(o.LastVersionId) {
		return nil, false
	}
	return o.LastVersionId, true
}

// HasLastVersionId returns a boolean if a field has been set.
func (o *SettingFormulaModel) HasLastVersionId() bool {
	if o != nil && !IsNil(o.LastVersionId) {
		return true
	}

	return false
}

// SetLastVersionId gets a reference to the given string and assigns it to the LastVersionId field.
func (o *SettingFormulaModel) SetLastVersionId(v string) {
	o.LastVersionId = &v
}

// GetDefaultValue returns the DefaultValue field value if set, zero value otherwise.
func (o *SettingFormulaModel) GetDefaultValue() ValueModel {
	if o == nil || IsNil(o.DefaultValue) {
		var ret ValueModel
		return ret
	}
	return *o.DefaultValue
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingFormulaModel) GetDefaultValueOk() (*ValueModel, bool) {
	if o == nil || IsNil(o.DefaultValue) {
		return nil, false
	}
	return o.DefaultValue, true
}

// HasDefaultValue returns a boolean if a field has been set.
func (o *SettingFormulaModel) HasDefaultValue() bool {
	if o != nil && !IsNil(o.DefaultValue) {
		return true
	}

	return false
}

// SetDefaultValue gets a reference to the given ValueModel and assigns it to the DefaultValue field.
func (o *SettingFormulaModel) SetDefaultValue(v ValueModel) {
	o.DefaultValue = &v
}

// GetTargetingRules returns the TargetingRules field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SettingFormulaModel) GetTargetingRules() []TargetingRuleModel {
	if o == nil {
		var ret []TargetingRuleModel
		return ret
	}
	return o.TargetingRules
}

// GetTargetingRulesOk returns a tuple with the TargetingRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SettingFormulaModel) GetTargetingRulesOk() ([]TargetingRuleModel, bool) {
	if o == nil || IsNil(o.TargetingRules) {
		return nil, false
	}
	return o.TargetingRules, true
}

// HasTargetingRules returns a boolean if a field has been set.
func (o *SettingFormulaModel) HasTargetingRules() bool {
	if o != nil && !IsNil(o.TargetingRules) {
		return true
	}

	return false
}

// SetTargetingRules gets a reference to the given []TargetingRuleModel and assigns it to the TargetingRules field.
func (o *SettingFormulaModel) SetTargetingRules(v []TargetingRuleModel) {
	o.TargetingRules = v
}

// GetSetting returns the Setting field value if set, zero value otherwise.
func (o *SettingFormulaModel) GetSetting() SettingDataModel {
	if o == nil || IsNil(o.Setting) {
		var ret SettingDataModel
		return ret
	}
	return *o.Setting
}

// GetSettingOk returns a tuple with the Setting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingFormulaModel) GetSettingOk() (*SettingDataModel, bool) {
	if o == nil || IsNil(o.Setting) {
		return nil, false
	}
	return o.Setting, true
}

// HasSetting returns a boolean if a field has been set.
func (o *SettingFormulaModel) HasSetting() bool {
	if o != nil && !IsNil(o.Setting) {
		return true
	}

	return false
}

// SetSetting gets a reference to the given SettingDataModel and assigns it to the Setting field.
func (o *SettingFormulaModel) SetSetting(v SettingDataModel) {
	o.Setting = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SettingFormulaModel) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SettingFormulaModel) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *SettingFormulaModel) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given NullableTime and assigns it to the UpdatedAt field.
func (o *SettingFormulaModel) SetUpdatedAt(v time.Time) {
	o.UpdatedAt.Set(&v)
}
// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil
func (o *SettingFormulaModel) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
func (o *SettingFormulaModel) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

// GetPercentageEvaluationAttribute returns the PercentageEvaluationAttribute field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SettingFormulaModel) GetPercentageEvaluationAttribute() string {
	if o == nil || IsNil(o.PercentageEvaluationAttribute.Get()) {
		var ret string
		return ret
	}
	return *o.PercentageEvaluationAttribute.Get()
}

// GetPercentageEvaluationAttributeOk returns a tuple with the PercentageEvaluationAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SettingFormulaModel) GetPercentageEvaluationAttributeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PercentageEvaluationAttribute.Get(), o.PercentageEvaluationAttribute.IsSet()
}

// HasPercentageEvaluationAttribute returns a boolean if a field has been set.
func (o *SettingFormulaModel) HasPercentageEvaluationAttribute() bool {
	if o != nil && o.PercentageEvaluationAttribute.IsSet() {
		return true
	}

	return false
}

// SetPercentageEvaluationAttribute gets a reference to the given NullableString and assigns it to the PercentageEvaluationAttribute field.
func (o *SettingFormulaModel) SetPercentageEvaluationAttribute(v string) {
	o.PercentageEvaluationAttribute.Set(&v)
}
// SetPercentageEvaluationAttributeNil sets the value for PercentageEvaluationAttribute to be an explicit nil
func (o *SettingFormulaModel) SetPercentageEvaluationAttributeNil() {
	o.PercentageEvaluationAttribute.Set(nil)
}

// UnsetPercentageEvaluationAttribute ensures that no value is present for PercentageEvaluationAttribute, not even an explicit nil
func (o *SettingFormulaModel) UnsetPercentageEvaluationAttribute() {
	o.PercentageEvaluationAttribute.Unset()
}

// GetLastUpdaterUserEmail returns the LastUpdaterUserEmail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SettingFormulaModel) GetLastUpdaterUserEmail() string {
	if o == nil || IsNil(o.LastUpdaterUserEmail.Get()) {
		var ret string
		return ret
	}
	return *o.LastUpdaterUserEmail.Get()
}

// GetLastUpdaterUserEmailOk returns a tuple with the LastUpdaterUserEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SettingFormulaModel) GetLastUpdaterUserEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdaterUserEmail.Get(), o.LastUpdaterUserEmail.IsSet()
}

// HasLastUpdaterUserEmail returns a boolean if a field has been set.
func (o *SettingFormulaModel) HasLastUpdaterUserEmail() bool {
	if o != nil && o.LastUpdaterUserEmail.IsSet() {
		return true
	}

	return false
}

// SetLastUpdaterUserEmail gets a reference to the given NullableString and assigns it to the LastUpdaterUserEmail field.
func (o *SettingFormulaModel) SetLastUpdaterUserEmail(v string) {
	o.LastUpdaterUserEmail.Set(&v)
}
// SetLastUpdaterUserEmailNil sets the value for LastUpdaterUserEmail to be an explicit nil
func (o *SettingFormulaModel) SetLastUpdaterUserEmailNil() {
	o.LastUpdaterUserEmail.Set(nil)
}

// UnsetLastUpdaterUserEmail ensures that no value is present for LastUpdaterUserEmail, not even an explicit nil
func (o *SettingFormulaModel) UnsetLastUpdaterUserEmail() {
	o.LastUpdaterUserEmail.Unset()
}

// GetLastUpdaterUserFullName returns the LastUpdaterUserFullName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SettingFormulaModel) GetLastUpdaterUserFullName() string {
	if o == nil || IsNil(o.LastUpdaterUserFullName.Get()) {
		var ret string
		return ret
	}
	return *o.LastUpdaterUserFullName.Get()
}

// GetLastUpdaterUserFullNameOk returns a tuple with the LastUpdaterUserFullName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SettingFormulaModel) GetLastUpdaterUserFullNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdaterUserFullName.Get(), o.LastUpdaterUserFullName.IsSet()
}

// HasLastUpdaterUserFullName returns a boolean if a field has been set.
func (o *SettingFormulaModel) HasLastUpdaterUserFullName() bool {
	if o != nil && o.LastUpdaterUserFullName.IsSet() {
		return true
	}

	return false
}

// SetLastUpdaterUserFullName gets a reference to the given NullableString and assigns it to the LastUpdaterUserFullName field.
func (o *SettingFormulaModel) SetLastUpdaterUserFullName(v string) {
	o.LastUpdaterUserFullName.Set(&v)
}
// SetLastUpdaterUserFullNameNil sets the value for LastUpdaterUserFullName to be an explicit nil
func (o *SettingFormulaModel) SetLastUpdaterUserFullNameNil() {
	o.LastUpdaterUserFullName.Set(nil)
}

// UnsetLastUpdaterUserFullName ensures that no value is present for LastUpdaterUserFullName, not even an explicit nil
func (o *SettingFormulaModel) UnsetLastUpdaterUserFullName() {
	o.LastUpdaterUserFullName.Unset()
}

// GetIntegrationLinks returns the IntegrationLinks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SettingFormulaModel) GetIntegrationLinks() []IntegrationLinkModel {
	if o == nil {
		var ret []IntegrationLinkModel
		return ret
	}
	return o.IntegrationLinks
}

// GetIntegrationLinksOk returns a tuple with the IntegrationLinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SettingFormulaModel) GetIntegrationLinksOk() ([]IntegrationLinkModel, bool) {
	if o == nil || IsNil(o.IntegrationLinks) {
		return nil, false
	}
	return o.IntegrationLinks, true
}

// HasIntegrationLinks returns a boolean if a field has been set.
func (o *SettingFormulaModel) HasIntegrationLinks() bool {
	if o != nil && !IsNil(o.IntegrationLinks) {
		return true
	}

	return false
}

// SetIntegrationLinks gets a reference to the given []IntegrationLinkModel and assigns it to the IntegrationLinks field.
func (o *SettingFormulaModel) SetIntegrationLinks(v []IntegrationLinkModel) {
	o.IntegrationLinks = v
}

// GetSettingTags returns the SettingTags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SettingFormulaModel) GetSettingTags() []SettingTagModel {
	if o == nil {
		var ret []SettingTagModel
		return ret
	}
	return o.SettingTags
}

// GetSettingTagsOk returns a tuple with the SettingTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SettingFormulaModel) GetSettingTagsOk() ([]SettingTagModel, bool) {
	if o == nil || IsNil(o.SettingTags) {
		return nil, false
	}
	return o.SettingTags, true
}

// HasSettingTags returns a boolean if a field has been set.
func (o *SettingFormulaModel) HasSettingTags() bool {
	if o != nil && !IsNil(o.SettingTags) {
		return true
	}

	return false
}

// SetSettingTags gets a reference to the given []SettingTagModel and assigns it to the SettingTags field.
func (o *SettingFormulaModel) SetSettingTags(v []SettingTagModel) {
	o.SettingTags = v
}

// GetSettingIdsWherePrerequisite returns the SettingIdsWherePrerequisite field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SettingFormulaModel) GetSettingIdsWherePrerequisite() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}
	return o.SettingIdsWherePrerequisite
}

// GetSettingIdsWherePrerequisiteOk returns a tuple with the SettingIdsWherePrerequisite field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SettingFormulaModel) GetSettingIdsWherePrerequisiteOk() ([]int32, bool) {
	if o == nil || IsNil(o.SettingIdsWherePrerequisite) {
		return nil, false
	}
	return o.SettingIdsWherePrerequisite, true
}

// HasSettingIdsWherePrerequisite returns a boolean if a field has been set.
func (o *SettingFormulaModel) HasSettingIdsWherePrerequisite() bool {
	if o != nil && !IsNil(o.SettingIdsWherePrerequisite) {
		return true
	}

	return false
}

// SetSettingIdsWherePrerequisite gets a reference to the given []int32 and assigns it to the SettingIdsWherePrerequisite field.
func (o *SettingFormulaModel) SetSettingIdsWherePrerequisite(v []int32) {
	o.SettingIdsWherePrerequisite = v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *SettingFormulaModel) GetConfig() ConfigModel {
	if o == nil || IsNil(o.Config) {
		var ret ConfigModel
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingFormulaModel) GetConfigOk() (*ConfigModel, bool) {
	if o == nil || IsNil(o.Config) {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *SettingFormulaModel) HasConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given ConfigModel and assigns it to the Config field.
func (o *SettingFormulaModel) SetConfig(v ConfigModel) {
	o.Config = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *SettingFormulaModel) GetEnvironment() EnvironmentModel {
	if o == nil || IsNil(o.Environment) {
		var ret EnvironmentModel
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingFormulaModel) GetEnvironmentOk() (*EnvironmentModel, bool) {
	if o == nil || IsNil(o.Environment) {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *SettingFormulaModel) HasEnvironment() bool {
	if o != nil && !IsNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given EnvironmentModel and assigns it to the Environment field.
func (o *SettingFormulaModel) SetEnvironment(v EnvironmentModel) {
	o.Environment = &v
}

// GetReadOnly returns the ReadOnly field value if set, zero value otherwise.
func (o *SettingFormulaModel) GetReadOnly() bool {
	if o == nil || IsNil(o.ReadOnly) {
		var ret bool
		return ret
	}
	return *o.ReadOnly
}

// GetReadOnlyOk returns a tuple with the ReadOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingFormulaModel) GetReadOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.ReadOnly) {
		return nil, false
	}
	return o.ReadOnly, true
}

// HasReadOnly returns a boolean if a field has been set.
func (o *SettingFormulaModel) HasReadOnly() bool {
	if o != nil && !IsNil(o.ReadOnly) {
		return true
	}

	return false
}

// SetReadOnly gets a reference to the given bool and assigns it to the ReadOnly field.
func (o *SettingFormulaModel) SetReadOnly(v bool) {
	o.ReadOnly = &v
}

// GetFeatureFlagLimitations returns the FeatureFlagLimitations field value if set, zero value otherwise.
func (o *SettingFormulaModel) GetFeatureFlagLimitations() FeatureFlagLimitations {
	if o == nil || IsNil(o.FeatureFlagLimitations) {
		var ret FeatureFlagLimitations
		return ret
	}
	return *o.FeatureFlagLimitations
}

// GetFeatureFlagLimitationsOk returns a tuple with the FeatureFlagLimitations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingFormulaModel) GetFeatureFlagLimitationsOk() (*FeatureFlagLimitations, bool) {
	if o == nil || IsNil(o.FeatureFlagLimitations) {
		return nil, false
	}
	return o.FeatureFlagLimitations, true
}

// HasFeatureFlagLimitations returns a boolean if a field has been set.
func (o *SettingFormulaModel) HasFeatureFlagLimitations() bool {
	if o != nil && !IsNil(o.FeatureFlagLimitations) {
		return true
	}

	return false
}

// SetFeatureFlagLimitations gets a reference to the given FeatureFlagLimitations and assigns it to the FeatureFlagLimitations field.
func (o *SettingFormulaModel) SetFeatureFlagLimitations(v FeatureFlagLimitations) {
	o.FeatureFlagLimitations = &v
}

func (o SettingFormulaModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SettingFormulaModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LastVersionId) {
		toSerialize["lastVersionId"] = o.LastVersionId
	}
	if !IsNil(o.DefaultValue) {
		toSerialize["defaultValue"] = o.DefaultValue
	}
	if o.TargetingRules != nil {
		toSerialize["targetingRules"] = o.TargetingRules
	}
	if !IsNil(o.Setting) {
		toSerialize["setting"] = o.Setting
	}
	if o.UpdatedAt.IsSet() {
		toSerialize["updatedAt"] = o.UpdatedAt.Get()
	}
	if o.PercentageEvaluationAttribute.IsSet() {
		toSerialize["percentageEvaluationAttribute"] = o.PercentageEvaluationAttribute.Get()
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
	if o.SettingIdsWherePrerequisite != nil {
		toSerialize["settingIdsWherePrerequisite"] = o.SettingIdsWherePrerequisite
	}
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	if !IsNil(o.Environment) {
		toSerialize["environment"] = o.Environment
	}
	if !IsNil(o.ReadOnly) {
		toSerialize["readOnly"] = o.ReadOnly
	}
	if !IsNil(o.FeatureFlagLimitations) {
		toSerialize["featureFlagLimitations"] = o.FeatureFlagLimitations
	}
	return toSerialize, nil
}

type NullableSettingFormulaModel struct {
	value *SettingFormulaModel
	isSet bool
}

func (v NullableSettingFormulaModel) Get() *SettingFormulaModel {
	return v.value
}

func (v *NullableSettingFormulaModel) Set(val *SettingFormulaModel) {
	v.value = val
	v.isSet = true
}

func (v NullableSettingFormulaModel) IsSet() bool {
	return v.isSet
}

func (v *NullableSettingFormulaModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSettingFormulaModel(val *SettingFormulaModel) *NullableSettingFormulaModel {
	return &NullableSettingFormulaModel{value: val, isSet: true}
}

func (v NullableSettingFormulaModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSettingFormulaModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


