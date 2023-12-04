# ConfigSettingFormulaModel

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

## Methods

### NewConfigSettingFormulaModel

`func NewConfigSettingFormulaModel() *ConfigSettingFormulaModel`

NewConfigSettingFormulaModel instantiates a new ConfigSettingFormulaModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigSettingFormulaModelWithDefaults

`func NewConfigSettingFormulaModelWithDefaults() *ConfigSettingFormulaModel`

NewConfigSettingFormulaModelWithDefaults instantiates a new ConfigSettingFormulaModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastVersionId

`func (o *ConfigSettingFormulaModel) GetLastVersionId() string`

GetLastVersionId returns the LastVersionId field if non-nil, zero value otherwise.

### GetLastVersionIdOk

`func (o *ConfigSettingFormulaModel) GetLastVersionIdOk() (*string, bool)`

GetLastVersionIdOk returns a tuple with the LastVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastVersionId

`func (o *ConfigSettingFormulaModel) SetLastVersionId(v string)`

SetLastVersionId sets LastVersionId field to given value.

### HasLastVersionId

`func (o *ConfigSettingFormulaModel) HasLastVersionId() bool`

HasLastVersionId returns a boolean if a field has been set.

### GetDefaultValue

`func (o *ConfigSettingFormulaModel) GetDefaultValue() ValueModel`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *ConfigSettingFormulaModel) GetDefaultValueOk() (*ValueModel, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *ConfigSettingFormulaModel) SetDefaultValue(v ValueModel)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *ConfigSettingFormulaModel) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetTargetingRules

`func (o *ConfigSettingFormulaModel) GetTargetingRules() []TargetingRuleModel`

GetTargetingRules returns the TargetingRules field if non-nil, zero value otherwise.

### GetTargetingRulesOk

`func (o *ConfigSettingFormulaModel) GetTargetingRulesOk() (*[]TargetingRuleModel, bool)`

GetTargetingRulesOk returns a tuple with the TargetingRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetingRules

`func (o *ConfigSettingFormulaModel) SetTargetingRules(v []TargetingRuleModel)`

SetTargetingRules sets TargetingRules field to given value.

### HasTargetingRules

`func (o *ConfigSettingFormulaModel) HasTargetingRules() bool`

HasTargetingRules returns a boolean if a field has been set.

### SetTargetingRulesNil

`func (o *ConfigSettingFormulaModel) SetTargetingRulesNil(b bool)`

 SetTargetingRulesNil sets the value for TargetingRules to be an explicit nil

### UnsetTargetingRules
`func (o *ConfigSettingFormulaModel) UnsetTargetingRules()`

UnsetTargetingRules ensures that no value is present for TargetingRules, not even an explicit nil
### GetSetting

`func (o *ConfigSettingFormulaModel) GetSetting() SettingDataModel`

GetSetting returns the Setting field if non-nil, zero value otherwise.

### GetSettingOk

`func (o *ConfigSettingFormulaModel) GetSettingOk() (*SettingDataModel, bool)`

GetSettingOk returns a tuple with the Setting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetting

`func (o *ConfigSettingFormulaModel) SetSetting(v SettingDataModel)`

SetSetting sets Setting field to given value.

### HasSetting

`func (o *ConfigSettingFormulaModel) HasSetting() bool`

HasSetting returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ConfigSettingFormulaModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ConfigSettingFormulaModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ConfigSettingFormulaModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ConfigSettingFormulaModel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *ConfigSettingFormulaModel) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *ConfigSettingFormulaModel) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetPercentageEvaluationAttribute

`func (o *ConfigSettingFormulaModel) GetPercentageEvaluationAttribute() string`

GetPercentageEvaluationAttribute returns the PercentageEvaluationAttribute field if non-nil, zero value otherwise.

### GetPercentageEvaluationAttributeOk

`func (o *ConfigSettingFormulaModel) GetPercentageEvaluationAttributeOk() (*string, bool)`

GetPercentageEvaluationAttributeOk returns a tuple with the PercentageEvaluationAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageEvaluationAttribute

`func (o *ConfigSettingFormulaModel) SetPercentageEvaluationAttribute(v string)`

SetPercentageEvaluationAttribute sets PercentageEvaluationAttribute field to given value.

### HasPercentageEvaluationAttribute

`func (o *ConfigSettingFormulaModel) HasPercentageEvaluationAttribute() bool`

HasPercentageEvaluationAttribute returns a boolean if a field has been set.

### SetPercentageEvaluationAttributeNil

`func (o *ConfigSettingFormulaModel) SetPercentageEvaluationAttributeNil(b bool)`

 SetPercentageEvaluationAttributeNil sets the value for PercentageEvaluationAttribute to be an explicit nil

### UnsetPercentageEvaluationAttribute
`func (o *ConfigSettingFormulaModel) UnsetPercentageEvaluationAttribute()`

UnsetPercentageEvaluationAttribute ensures that no value is present for PercentageEvaluationAttribute, not even an explicit nil
### GetLastUpdaterUserEmail

`func (o *ConfigSettingFormulaModel) GetLastUpdaterUserEmail() string`

GetLastUpdaterUserEmail returns the LastUpdaterUserEmail field if non-nil, zero value otherwise.

### GetLastUpdaterUserEmailOk

`func (o *ConfigSettingFormulaModel) GetLastUpdaterUserEmailOk() (*string, bool)`

GetLastUpdaterUserEmailOk returns a tuple with the LastUpdaterUserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdaterUserEmail

`func (o *ConfigSettingFormulaModel) SetLastUpdaterUserEmail(v string)`

SetLastUpdaterUserEmail sets LastUpdaterUserEmail field to given value.

### HasLastUpdaterUserEmail

`func (o *ConfigSettingFormulaModel) HasLastUpdaterUserEmail() bool`

HasLastUpdaterUserEmail returns a boolean if a field has been set.

### SetLastUpdaterUserEmailNil

`func (o *ConfigSettingFormulaModel) SetLastUpdaterUserEmailNil(b bool)`

 SetLastUpdaterUserEmailNil sets the value for LastUpdaterUserEmail to be an explicit nil

### UnsetLastUpdaterUserEmail
`func (o *ConfigSettingFormulaModel) UnsetLastUpdaterUserEmail()`

UnsetLastUpdaterUserEmail ensures that no value is present for LastUpdaterUserEmail, not even an explicit nil
### GetLastUpdaterUserFullName

`func (o *ConfigSettingFormulaModel) GetLastUpdaterUserFullName() string`

GetLastUpdaterUserFullName returns the LastUpdaterUserFullName field if non-nil, zero value otherwise.

### GetLastUpdaterUserFullNameOk

`func (o *ConfigSettingFormulaModel) GetLastUpdaterUserFullNameOk() (*string, bool)`

GetLastUpdaterUserFullNameOk returns a tuple with the LastUpdaterUserFullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdaterUserFullName

`func (o *ConfigSettingFormulaModel) SetLastUpdaterUserFullName(v string)`

SetLastUpdaterUserFullName sets LastUpdaterUserFullName field to given value.

### HasLastUpdaterUserFullName

`func (o *ConfigSettingFormulaModel) HasLastUpdaterUserFullName() bool`

HasLastUpdaterUserFullName returns a boolean if a field has been set.

### SetLastUpdaterUserFullNameNil

`func (o *ConfigSettingFormulaModel) SetLastUpdaterUserFullNameNil(b bool)`

 SetLastUpdaterUserFullNameNil sets the value for LastUpdaterUserFullName to be an explicit nil

### UnsetLastUpdaterUserFullName
`func (o *ConfigSettingFormulaModel) UnsetLastUpdaterUserFullName()`

UnsetLastUpdaterUserFullName ensures that no value is present for LastUpdaterUserFullName, not even an explicit nil
### GetIntegrationLinks

`func (o *ConfigSettingFormulaModel) GetIntegrationLinks() []IntegrationLinkModel`

GetIntegrationLinks returns the IntegrationLinks field if non-nil, zero value otherwise.

### GetIntegrationLinksOk

`func (o *ConfigSettingFormulaModel) GetIntegrationLinksOk() (*[]IntegrationLinkModel, bool)`

GetIntegrationLinksOk returns a tuple with the IntegrationLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationLinks

`func (o *ConfigSettingFormulaModel) SetIntegrationLinks(v []IntegrationLinkModel)`

SetIntegrationLinks sets IntegrationLinks field to given value.

### HasIntegrationLinks

`func (o *ConfigSettingFormulaModel) HasIntegrationLinks() bool`

HasIntegrationLinks returns a boolean if a field has been set.

### SetIntegrationLinksNil

`func (o *ConfigSettingFormulaModel) SetIntegrationLinksNil(b bool)`

 SetIntegrationLinksNil sets the value for IntegrationLinks to be an explicit nil

### UnsetIntegrationLinks
`func (o *ConfigSettingFormulaModel) UnsetIntegrationLinks()`

UnsetIntegrationLinks ensures that no value is present for IntegrationLinks, not even an explicit nil
### GetSettingTags

`func (o *ConfigSettingFormulaModel) GetSettingTags() []SettingTagModel`

GetSettingTags returns the SettingTags field if non-nil, zero value otherwise.

### GetSettingTagsOk

`func (o *ConfigSettingFormulaModel) GetSettingTagsOk() (*[]SettingTagModel, bool)`

GetSettingTagsOk returns a tuple with the SettingTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingTags

`func (o *ConfigSettingFormulaModel) SetSettingTags(v []SettingTagModel)`

SetSettingTags sets SettingTags field to given value.

### HasSettingTags

`func (o *ConfigSettingFormulaModel) HasSettingTags() bool`

HasSettingTags returns a boolean if a field has been set.

### SetSettingTagsNil

`func (o *ConfigSettingFormulaModel) SetSettingTagsNil(b bool)`

 SetSettingTagsNil sets the value for SettingTags to be an explicit nil

### UnsetSettingTags
`func (o *ConfigSettingFormulaModel) UnsetSettingTags()`

UnsetSettingTags ensures that no value is present for SettingTags, not even an explicit nil
### GetSettingIdsWherePrerequisite

`func (o *ConfigSettingFormulaModel) GetSettingIdsWherePrerequisite() []int32`

GetSettingIdsWherePrerequisite returns the SettingIdsWherePrerequisite field if non-nil, zero value otherwise.

### GetSettingIdsWherePrerequisiteOk

`func (o *ConfigSettingFormulaModel) GetSettingIdsWherePrerequisiteOk() (*[]int32, bool)`

GetSettingIdsWherePrerequisiteOk returns a tuple with the SettingIdsWherePrerequisite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingIdsWherePrerequisite

`func (o *ConfigSettingFormulaModel) SetSettingIdsWherePrerequisite(v []int32)`

SetSettingIdsWherePrerequisite sets SettingIdsWherePrerequisite field to given value.

### HasSettingIdsWherePrerequisite

`func (o *ConfigSettingFormulaModel) HasSettingIdsWherePrerequisite() bool`

HasSettingIdsWherePrerequisite returns a boolean if a field has been set.

### SetSettingIdsWherePrerequisiteNil

`func (o *ConfigSettingFormulaModel) SetSettingIdsWherePrerequisiteNil(b bool)`

 SetSettingIdsWherePrerequisiteNil sets the value for SettingIdsWherePrerequisite to be an explicit nil

### UnsetSettingIdsWherePrerequisite
`func (o *ConfigSettingFormulaModel) UnsetSettingIdsWherePrerequisite()`

UnsetSettingIdsWherePrerequisite ensures that no value is present for SettingIdsWherePrerequisite, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


