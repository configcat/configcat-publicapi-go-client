# ConfigSettingFormulaModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastVersionId** | **string** |  | 
**DefaultValue** | [**ValueModel**](ValueModel.md) |  | 
**TargetingRules** | [**[]TargetingRuleModel**](TargetingRuleModel.md) | The targeting rules of the Feature Flag or Setting. | 
**Setting** | [**SettingDataV2Model**](SettingDataV2Model.md) |  | 
**UpdatedAt** | **NullableTime** | The last updated date and time when the Feature Flag or Setting. | 
**PercentageEvaluationAttribute** | **NullableString** | The user attribute used for percentage evaluation. If not set, it defaults to the &#x60;Identifier&#x60; user object attribute. | 
**LastUpdaterUserEmail** | **NullableString** | The email of the user who last updated the Feature Flag or Setting. | 
**LastUpdaterUserFullName** | **NullableString** | The name of the user who last updated the Feature Flag or Setting. | 
**IntegrationLinks** | [**[]IntegrationLinkModel**](IntegrationLinkModel.md) | The integration links attached to the Feature Flag or Setting. | 
**SettingTags** | [**[]SettingTagModel**](SettingTagModel.md) | The tags attached to the Feature Flag or Setting. | 
**SettingIdsWherePrerequisite** | **[]int32** | List of Feature Flag and Setting IDs where the actual Feature Flag or Setting is prerequisite. | 

## Methods

### NewConfigSettingFormulaModel

`func NewConfigSettingFormulaModel(lastVersionId string, defaultValue ValueModel, targetingRules []TargetingRuleModel, setting SettingDataV2Model, updatedAt NullableTime, percentageEvaluationAttribute NullableString, lastUpdaterUserEmail NullableString, lastUpdaterUserFullName NullableString, integrationLinks []IntegrationLinkModel, settingTags []SettingTagModel, settingIdsWherePrerequisite []int32, ) *ConfigSettingFormulaModel`

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


### GetSetting

`func (o *ConfigSettingFormulaModel) GetSetting() SettingDataV2Model`

GetSetting returns the Setting field if non-nil, zero value otherwise.

### GetSettingOk

`func (o *ConfigSettingFormulaModel) GetSettingOk() (*SettingDataV2Model, bool)`

GetSettingOk returns a tuple with the Setting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetting

`func (o *ConfigSettingFormulaModel) SetSetting(v SettingDataV2Model)`

SetSetting sets Setting field to given value.


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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


