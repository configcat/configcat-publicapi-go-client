# SettingFormulaModel

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
**Config** | [**ConfigModel**](ConfigModel.md) |  | 
**Environment** | [**EnvironmentModel**](EnvironmentModel.md) |  | 
**ReadOnly** | **bool** |  | 
**FeatureFlagLimitations** | [**FeatureFlagLimitations**](FeatureFlagLimitations.md) |  | 

## Methods

### NewSettingFormulaModel

`func NewSettingFormulaModel(lastVersionId string, defaultValue ValueModel, targetingRules []TargetingRuleModel, setting SettingDataV2Model, updatedAt NullableTime, percentageEvaluationAttribute NullableString, lastUpdaterUserEmail NullableString, lastUpdaterUserFullName NullableString, integrationLinks []IntegrationLinkModel, settingTags []SettingTagModel, settingIdsWherePrerequisite []int32, config ConfigModel, environment EnvironmentModel, readOnly bool, featureFlagLimitations FeatureFlagLimitations, ) *SettingFormulaModel`

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


### GetSetting

`func (o *SettingFormulaModel) GetSetting() SettingDataV2Model`

GetSetting returns the Setting field if non-nil, zero value otherwise.

### GetSettingOk

`func (o *SettingFormulaModel) GetSettingOk() (*SettingDataV2Model, bool)`

GetSettingOk returns a tuple with the Setting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetting

`func (o *SettingFormulaModel) SetSetting(v SettingDataV2Model)`

SetSetting sets Setting field to given value.


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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


