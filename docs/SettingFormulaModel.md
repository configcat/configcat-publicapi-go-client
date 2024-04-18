# SettingFormulaModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastVersionId** | Pointer to **string** |  | [optional] 
**DefaultValue** | Pointer to [**ValueModel**](ValueModel.md) |  | [optional] 
**TargetingRules** | Pointer to [**[]TargetingRuleModel**](TargetingRuleModel.md) | The targeting rules of the Feature Flag or Setting. | [optional] 
**Setting** | Pointer to [**SettingDataModel**](SettingDataModel.md) |  | [optional] 
**UpdatedAt** | Pointer to **NullableTime** | The last updated date and time when the Feature Flag or Setting. | [optional] 
**PercentageEvaluationAttribute** | Pointer to **NullableString** | The user attribute used for percentage evaluation. If not set, it defaults to the &#x60;Identifier&#x60; user object attribute. | [optional] 
**LastUpdaterUserEmail** | Pointer to **NullableString** | The email of the user who last updated the Feature Flag or Setting. | [optional] 
**LastUpdaterUserFullName** | Pointer to **NullableString** | The name of the user who last updated the Feature Flag or Setting. | [optional] 
**IntegrationLinks** | Pointer to [**[]IntegrationLinkModel**](IntegrationLinkModel.md) | The integration links attached to the Feature Flag or Setting. | [optional] 
**SettingTags** | Pointer to [**[]SettingTagModel**](SettingTagModel.md) | The tags attached to the Feature Flag or Setting. | [optional] 
**SettingIdsWherePrerequisite** | Pointer to **[]int32** | List of Feature Flag and Setting IDs where the actual Feature Flag or Setting is prerequisite. | [optional] 
**Config** | Pointer to [**ConfigModel**](ConfigModel.md) |  | [optional] 
**Environment** | Pointer to [**EnvironmentModel**](EnvironmentModel.md) |  | [optional] 
**ReadOnly** | Pointer to **bool** |  | [optional] 
**FeatureFlagLimitations** | Pointer to [**FeatureFlagLimitations**](FeatureFlagLimitations.md) |  | [optional] 

## Methods

### NewSettingFormulaModel

`func NewSettingFormulaModel() *SettingFormulaModel`

NewSettingFormulaModel instantiates a new SettingFormulaModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettingFormulaModelWithDefaults

`func NewSettingFormulaModelWithDefaults() *SettingFormulaModel`

NewSettingFormulaModelWithDefaults instantiates a new SettingFormulaModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastVersionId

`func (o *SettingFormulaModel) GetLastVersionId() string`

GetLastVersionId returns the LastVersionId field if non-nil, zero value otherwise.

### GetLastVersionIdOk

`func (o *SettingFormulaModel) GetLastVersionIdOk() (*string, bool)`

GetLastVersionIdOk returns a tuple with the LastVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastVersionId

`func (o *SettingFormulaModel) SetLastVersionId(v string)`

SetLastVersionId sets LastVersionId field to given value.

### HasLastVersionId

`func (o *SettingFormulaModel) HasLastVersionId() bool`

HasLastVersionId returns a boolean if a field has been set.

### GetDefaultValue

`func (o *SettingFormulaModel) GetDefaultValue() ValueModel`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *SettingFormulaModel) GetDefaultValueOk() (*ValueModel, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *SettingFormulaModel) SetDefaultValue(v ValueModel)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *SettingFormulaModel) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetTargetingRules

`func (o *SettingFormulaModel) GetTargetingRules() []TargetingRuleModel`

GetTargetingRules returns the TargetingRules field if non-nil, zero value otherwise.

### GetTargetingRulesOk

`func (o *SettingFormulaModel) GetTargetingRulesOk() (*[]TargetingRuleModel, bool)`

GetTargetingRulesOk returns a tuple with the TargetingRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetingRules

`func (o *SettingFormulaModel) SetTargetingRules(v []TargetingRuleModel)`

SetTargetingRules sets TargetingRules field to given value.

### HasTargetingRules

`func (o *SettingFormulaModel) HasTargetingRules() bool`

HasTargetingRules returns a boolean if a field has been set.

### SetTargetingRulesNil

`func (o *SettingFormulaModel) SetTargetingRulesNil(b bool)`

 SetTargetingRulesNil sets the value for TargetingRules to be an explicit nil

### UnsetTargetingRules
`func (o *SettingFormulaModel) UnsetTargetingRules()`

UnsetTargetingRules ensures that no value is present for TargetingRules, not even an explicit nil
### GetSetting

`func (o *SettingFormulaModel) GetSetting() SettingDataModel`

GetSetting returns the Setting field if non-nil, zero value otherwise.

### GetSettingOk

`func (o *SettingFormulaModel) GetSettingOk() (*SettingDataModel, bool)`

GetSettingOk returns a tuple with the Setting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetting

`func (o *SettingFormulaModel) SetSetting(v SettingDataModel)`

SetSetting sets Setting field to given value.

### HasSetting

`func (o *SettingFormulaModel) HasSetting() bool`

HasSetting returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *SettingFormulaModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SettingFormulaModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SettingFormulaModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SettingFormulaModel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *SettingFormulaModel) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *SettingFormulaModel) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetPercentageEvaluationAttribute

`func (o *SettingFormulaModel) GetPercentageEvaluationAttribute() string`

GetPercentageEvaluationAttribute returns the PercentageEvaluationAttribute field if non-nil, zero value otherwise.

### GetPercentageEvaluationAttributeOk

`func (o *SettingFormulaModel) GetPercentageEvaluationAttributeOk() (*string, bool)`

GetPercentageEvaluationAttributeOk returns a tuple with the PercentageEvaluationAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageEvaluationAttribute

`func (o *SettingFormulaModel) SetPercentageEvaluationAttribute(v string)`

SetPercentageEvaluationAttribute sets PercentageEvaluationAttribute field to given value.

### HasPercentageEvaluationAttribute

`func (o *SettingFormulaModel) HasPercentageEvaluationAttribute() bool`

HasPercentageEvaluationAttribute returns a boolean if a field has been set.

### SetPercentageEvaluationAttributeNil

`func (o *SettingFormulaModel) SetPercentageEvaluationAttributeNil(b bool)`

 SetPercentageEvaluationAttributeNil sets the value for PercentageEvaluationAttribute to be an explicit nil

### UnsetPercentageEvaluationAttribute
`func (o *SettingFormulaModel) UnsetPercentageEvaluationAttribute()`

UnsetPercentageEvaluationAttribute ensures that no value is present for PercentageEvaluationAttribute, not even an explicit nil
### GetLastUpdaterUserEmail

`func (o *SettingFormulaModel) GetLastUpdaterUserEmail() string`

GetLastUpdaterUserEmail returns the LastUpdaterUserEmail field if non-nil, zero value otherwise.

### GetLastUpdaterUserEmailOk

`func (o *SettingFormulaModel) GetLastUpdaterUserEmailOk() (*string, bool)`

GetLastUpdaterUserEmailOk returns a tuple with the LastUpdaterUserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdaterUserEmail

`func (o *SettingFormulaModel) SetLastUpdaterUserEmail(v string)`

SetLastUpdaterUserEmail sets LastUpdaterUserEmail field to given value.

### HasLastUpdaterUserEmail

`func (o *SettingFormulaModel) HasLastUpdaterUserEmail() bool`

HasLastUpdaterUserEmail returns a boolean if a field has been set.

### SetLastUpdaterUserEmailNil

`func (o *SettingFormulaModel) SetLastUpdaterUserEmailNil(b bool)`

 SetLastUpdaterUserEmailNil sets the value for LastUpdaterUserEmail to be an explicit nil

### UnsetLastUpdaterUserEmail
`func (o *SettingFormulaModel) UnsetLastUpdaterUserEmail()`

UnsetLastUpdaterUserEmail ensures that no value is present for LastUpdaterUserEmail, not even an explicit nil
### GetLastUpdaterUserFullName

`func (o *SettingFormulaModel) GetLastUpdaterUserFullName() string`

GetLastUpdaterUserFullName returns the LastUpdaterUserFullName field if non-nil, zero value otherwise.

### GetLastUpdaterUserFullNameOk

`func (o *SettingFormulaModel) GetLastUpdaterUserFullNameOk() (*string, bool)`

GetLastUpdaterUserFullNameOk returns a tuple with the LastUpdaterUserFullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdaterUserFullName

`func (o *SettingFormulaModel) SetLastUpdaterUserFullName(v string)`

SetLastUpdaterUserFullName sets LastUpdaterUserFullName field to given value.

### HasLastUpdaterUserFullName

`func (o *SettingFormulaModel) HasLastUpdaterUserFullName() bool`

HasLastUpdaterUserFullName returns a boolean if a field has been set.

### SetLastUpdaterUserFullNameNil

`func (o *SettingFormulaModel) SetLastUpdaterUserFullNameNil(b bool)`

 SetLastUpdaterUserFullNameNil sets the value for LastUpdaterUserFullName to be an explicit nil

### UnsetLastUpdaterUserFullName
`func (o *SettingFormulaModel) UnsetLastUpdaterUserFullName()`

UnsetLastUpdaterUserFullName ensures that no value is present for LastUpdaterUserFullName, not even an explicit nil
### GetIntegrationLinks

`func (o *SettingFormulaModel) GetIntegrationLinks() []IntegrationLinkModel`

GetIntegrationLinks returns the IntegrationLinks field if non-nil, zero value otherwise.

### GetIntegrationLinksOk

`func (o *SettingFormulaModel) GetIntegrationLinksOk() (*[]IntegrationLinkModel, bool)`

GetIntegrationLinksOk returns a tuple with the IntegrationLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationLinks

`func (o *SettingFormulaModel) SetIntegrationLinks(v []IntegrationLinkModel)`

SetIntegrationLinks sets IntegrationLinks field to given value.

### HasIntegrationLinks

`func (o *SettingFormulaModel) HasIntegrationLinks() bool`

HasIntegrationLinks returns a boolean if a field has been set.

### SetIntegrationLinksNil

`func (o *SettingFormulaModel) SetIntegrationLinksNil(b bool)`

 SetIntegrationLinksNil sets the value for IntegrationLinks to be an explicit nil

### UnsetIntegrationLinks
`func (o *SettingFormulaModel) UnsetIntegrationLinks()`

UnsetIntegrationLinks ensures that no value is present for IntegrationLinks, not even an explicit nil
### GetSettingTags

`func (o *SettingFormulaModel) GetSettingTags() []SettingTagModel`

GetSettingTags returns the SettingTags field if non-nil, zero value otherwise.

### GetSettingTagsOk

`func (o *SettingFormulaModel) GetSettingTagsOk() (*[]SettingTagModel, bool)`

GetSettingTagsOk returns a tuple with the SettingTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingTags

`func (o *SettingFormulaModel) SetSettingTags(v []SettingTagModel)`

SetSettingTags sets SettingTags field to given value.

### HasSettingTags

`func (o *SettingFormulaModel) HasSettingTags() bool`

HasSettingTags returns a boolean if a field has been set.

### SetSettingTagsNil

`func (o *SettingFormulaModel) SetSettingTagsNil(b bool)`

 SetSettingTagsNil sets the value for SettingTags to be an explicit nil

### UnsetSettingTags
`func (o *SettingFormulaModel) UnsetSettingTags()`

UnsetSettingTags ensures that no value is present for SettingTags, not even an explicit nil
### GetSettingIdsWherePrerequisite

`func (o *SettingFormulaModel) GetSettingIdsWherePrerequisite() []int32`

GetSettingIdsWherePrerequisite returns the SettingIdsWherePrerequisite field if non-nil, zero value otherwise.

### GetSettingIdsWherePrerequisiteOk

`func (o *SettingFormulaModel) GetSettingIdsWherePrerequisiteOk() (*[]int32, bool)`

GetSettingIdsWherePrerequisiteOk returns a tuple with the SettingIdsWherePrerequisite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingIdsWherePrerequisite

`func (o *SettingFormulaModel) SetSettingIdsWherePrerequisite(v []int32)`

SetSettingIdsWherePrerequisite sets SettingIdsWherePrerequisite field to given value.

### HasSettingIdsWherePrerequisite

`func (o *SettingFormulaModel) HasSettingIdsWherePrerequisite() bool`

HasSettingIdsWherePrerequisite returns a boolean if a field has been set.

### SetSettingIdsWherePrerequisiteNil

`func (o *SettingFormulaModel) SetSettingIdsWherePrerequisiteNil(b bool)`

 SetSettingIdsWherePrerequisiteNil sets the value for SettingIdsWherePrerequisite to be an explicit nil

### UnsetSettingIdsWherePrerequisite
`func (o *SettingFormulaModel) UnsetSettingIdsWherePrerequisite()`

UnsetSettingIdsWherePrerequisite ensures that no value is present for SettingIdsWherePrerequisite, not even an explicit nil
### GetConfig

`func (o *SettingFormulaModel) GetConfig() ConfigModel`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *SettingFormulaModel) GetConfigOk() (*ConfigModel, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *SettingFormulaModel) SetConfig(v ConfigModel)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *SettingFormulaModel) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetEnvironment

`func (o *SettingFormulaModel) GetEnvironment() EnvironmentModel`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *SettingFormulaModel) GetEnvironmentOk() (*EnvironmentModel, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *SettingFormulaModel) SetEnvironment(v EnvironmentModel)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *SettingFormulaModel) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetReadOnly

`func (o *SettingFormulaModel) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *SettingFormulaModel) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *SettingFormulaModel) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *SettingFormulaModel) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetFeatureFlagLimitations

`func (o *SettingFormulaModel) GetFeatureFlagLimitations() FeatureFlagLimitations`

GetFeatureFlagLimitations returns the FeatureFlagLimitations field if non-nil, zero value otherwise.

### GetFeatureFlagLimitationsOk

`func (o *SettingFormulaModel) GetFeatureFlagLimitationsOk() (*FeatureFlagLimitations, bool)`

GetFeatureFlagLimitationsOk returns a tuple with the FeatureFlagLimitations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureFlagLimitations

`func (o *SettingFormulaModel) SetFeatureFlagLimitations(v FeatureFlagLimitations)`

SetFeatureFlagLimitations sets FeatureFlagLimitations field to given value.

### HasFeatureFlagLimitations

`func (o *SettingFormulaModel) HasFeatureFlagLimitations() bool`

HasFeatureFlagLimitations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


