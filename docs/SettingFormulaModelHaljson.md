# SettingFormulaModelHaljson

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastVersionId** | Pointer to **string** |  | [optional] 
**DefaultValue** | Pointer to [**ValueModel**](ValueModel.md) |  | [optional] 
**TargetingRules** | Pointer to [**[]TargetingRuleModel**](TargetingRuleModel.md) | The targeting rules of the Feature Flag or Setting. | [optional] 
**UpdatedAt** | Pointer to **NullableTime** | The last updated date and time when the Feature Flag or Setting. | [optional] 
**PercentageEvaluationAttribute** | Pointer to **NullableString** | The user attribute used for percentage evaluation. If not set, it defaults to the &#x60;Identifier&#x60; user object attribute. | [optional] 
**LastUpdaterUserEmail** | Pointer to **NullableString** | The email of the user who last updated the Feature Flag or Setting. | [optional] 
**LastUpdaterUserFullName** | Pointer to **NullableString** | The name of the user who last updated the Feature Flag or Setting. | [optional] 
**Embedded** | Pointer to [**SettingFormulaModelHaljsonEmbedded**](SettingFormulaModelHaljsonEmbedded.md) |  | [optional] 
**ReadOnly** | Pointer to **bool** |  | [optional] 
**FeatureFlagLimitations** | Pointer to [**FeatureFlagLimitations**](FeatureFlagLimitations.md) |  | [optional] 
**Links** | Pointer to [**ConfigSettingFormulasModelHaljsonEmbeddedEnvironmentLinks**](ConfigSettingFormulasModelHaljsonEmbeddedEnvironmentLinks.md) |  | [optional] 

## Methods

### NewSettingFormulaModelHaljson

`func NewSettingFormulaModelHaljson() *SettingFormulaModelHaljson`

NewSettingFormulaModelHaljson instantiates a new SettingFormulaModelHaljson object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettingFormulaModelHaljsonWithDefaults

`func NewSettingFormulaModelHaljsonWithDefaults() *SettingFormulaModelHaljson`

NewSettingFormulaModelHaljsonWithDefaults instantiates a new SettingFormulaModelHaljson object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastVersionId

`func (o *SettingFormulaModelHaljson) GetLastVersionId() string`

GetLastVersionId returns the LastVersionId field if non-nil, zero value otherwise.

### GetLastVersionIdOk

`func (o *SettingFormulaModelHaljson) GetLastVersionIdOk() (*string, bool)`

GetLastVersionIdOk returns a tuple with the LastVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastVersionId

`func (o *SettingFormulaModelHaljson) SetLastVersionId(v string)`

SetLastVersionId sets LastVersionId field to given value.

### HasLastVersionId

`func (o *SettingFormulaModelHaljson) HasLastVersionId() bool`

HasLastVersionId returns a boolean if a field has been set.

### GetDefaultValue

`func (o *SettingFormulaModelHaljson) GetDefaultValue() ValueModel`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *SettingFormulaModelHaljson) GetDefaultValueOk() (*ValueModel, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *SettingFormulaModelHaljson) SetDefaultValue(v ValueModel)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *SettingFormulaModelHaljson) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetTargetingRules

`func (o *SettingFormulaModelHaljson) GetTargetingRules() []TargetingRuleModel`

GetTargetingRules returns the TargetingRules field if non-nil, zero value otherwise.

### GetTargetingRulesOk

`func (o *SettingFormulaModelHaljson) GetTargetingRulesOk() (*[]TargetingRuleModel, bool)`

GetTargetingRulesOk returns a tuple with the TargetingRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetingRules

`func (o *SettingFormulaModelHaljson) SetTargetingRules(v []TargetingRuleModel)`

SetTargetingRules sets TargetingRules field to given value.

### HasTargetingRules

`func (o *SettingFormulaModelHaljson) HasTargetingRules() bool`

HasTargetingRules returns a boolean if a field has been set.

### SetTargetingRulesNil

`func (o *SettingFormulaModelHaljson) SetTargetingRulesNil(b bool)`

 SetTargetingRulesNil sets the value for TargetingRules to be an explicit nil

### UnsetTargetingRules
`func (o *SettingFormulaModelHaljson) UnsetTargetingRules()`

UnsetTargetingRules ensures that no value is present for TargetingRules, not even an explicit nil
### GetUpdatedAt

`func (o *SettingFormulaModelHaljson) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SettingFormulaModelHaljson) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SettingFormulaModelHaljson) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SettingFormulaModelHaljson) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *SettingFormulaModelHaljson) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *SettingFormulaModelHaljson) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetPercentageEvaluationAttribute

`func (o *SettingFormulaModelHaljson) GetPercentageEvaluationAttribute() string`

GetPercentageEvaluationAttribute returns the PercentageEvaluationAttribute field if non-nil, zero value otherwise.

### GetPercentageEvaluationAttributeOk

`func (o *SettingFormulaModelHaljson) GetPercentageEvaluationAttributeOk() (*string, bool)`

GetPercentageEvaluationAttributeOk returns a tuple with the PercentageEvaluationAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageEvaluationAttribute

`func (o *SettingFormulaModelHaljson) SetPercentageEvaluationAttribute(v string)`

SetPercentageEvaluationAttribute sets PercentageEvaluationAttribute field to given value.

### HasPercentageEvaluationAttribute

`func (o *SettingFormulaModelHaljson) HasPercentageEvaluationAttribute() bool`

HasPercentageEvaluationAttribute returns a boolean if a field has been set.

### SetPercentageEvaluationAttributeNil

`func (o *SettingFormulaModelHaljson) SetPercentageEvaluationAttributeNil(b bool)`

 SetPercentageEvaluationAttributeNil sets the value for PercentageEvaluationAttribute to be an explicit nil

### UnsetPercentageEvaluationAttribute
`func (o *SettingFormulaModelHaljson) UnsetPercentageEvaluationAttribute()`

UnsetPercentageEvaluationAttribute ensures that no value is present for PercentageEvaluationAttribute, not even an explicit nil
### GetLastUpdaterUserEmail

`func (o *SettingFormulaModelHaljson) GetLastUpdaterUserEmail() string`

GetLastUpdaterUserEmail returns the LastUpdaterUserEmail field if non-nil, zero value otherwise.

### GetLastUpdaterUserEmailOk

`func (o *SettingFormulaModelHaljson) GetLastUpdaterUserEmailOk() (*string, bool)`

GetLastUpdaterUserEmailOk returns a tuple with the LastUpdaterUserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdaterUserEmail

`func (o *SettingFormulaModelHaljson) SetLastUpdaterUserEmail(v string)`

SetLastUpdaterUserEmail sets LastUpdaterUserEmail field to given value.

### HasLastUpdaterUserEmail

`func (o *SettingFormulaModelHaljson) HasLastUpdaterUserEmail() bool`

HasLastUpdaterUserEmail returns a boolean if a field has been set.

### SetLastUpdaterUserEmailNil

`func (o *SettingFormulaModelHaljson) SetLastUpdaterUserEmailNil(b bool)`

 SetLastUpdaterUserEmailNil sets the value for LastUpdaterUserEmail to be an explicit nil

### UnsetLastUpdaterUserEmail
`func (o *SettingFormulaModelHaljson) UnsetLastUpdaterUserEmail()`

UnsetLastUpdaterUserEmail ensures that no value is present for LastUpdaterUserEmail, not even an explicit nil
### GetLastUpdaterUserFullName

`func (o *SettingFormulaModelHaljson) GetLastUpdaterUserFullName() string`

GetLastUpdaterUserFullName returns the LastUpdaterUserFullName field if non-nil, zero value otherwise.

### GetLastUpdaterUserFullNameOk

`func (o *SettingFormulaModelHaljson) GetLastUpdaterUserFullNameOk() (*string, bool)`

GetLastUpdaterUserFullNameOk returns a tuple with the LastUpdaterUserFullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdaterUserFullName

`func (o *SettingFormulaModelHaljson) SetLastUpdaterUserFullName(v string)`

SetLastUpdaterUserFullName sets LastUpdaterUserFullName field to given value.

### HasLastUpdaterUserFullName

`func (o *SettingFormulaModelHaljson) HasLastUpdaterUserFullName() bool`

HasLastUpdaterUserFullName returns a boolean if a field has been set.

### SetLastUpdaterUserFullNameNil

`func (o *SettingFormulaModelHaljson) SetLastUpdaterUserFullNameNil(b bool)`

 SetLastUpdaterUserFullNameNil sets the value for LastUpdaterUserFullName to be an explicit nil

### UnsetLastUpdaterUserFullName
`func (o *SettingFormulaModelHaljson) UnsetLastUpdaterUserFullName()`

UnsetLastUpdaterUserFullName ensures that no value is present for LastUpdaterUserFullName, not even an explicit nil
### GetEmbedded

`func (o *SettingFormulaModelHaljson) GetEmbedded() SettingFormulaModelHaljsonEmbedded`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *SettingFormulaModelHaljson) GetEmbeddedOk() (*SettingFormulaModelHaljsonEmbedded, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *SettingFormulaModelHaljson) SetEmbedded(v SettingFormulaModelHaljsonEmbedded)`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *SettingFormulaModelHaljson) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetReadOnly

`func (o *SettingFormulaModelHaljson) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *SettingFormulaModelHaljson) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *SettingFormulaModelHaljson) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *SettingFormulaModelHaljson) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetFeatureFlagLimitations

`func (o *SettingFormulaModelHaljson) GetFeatureFlagLimitations() FeatureFlagLimitations`

GetFeatureFlagLimitations returns the FeatureFlagLimitations field if non-nil, zero value otherwise.

### GetFeatureFlagLimitationsOk

`func (o *SettingFormulaModelHaljson) GetFeatureFlagLimitationsOk() (*FeatureFlagLimitations, bool)`

GetFeatureFlagLimitationsOk returns a tuple with the FeatureFlagLimitations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureFlagLimitations

`func (o *SettingFormulaModelHaljson) SetFeatureFlagLimitations(v FeatureFlagLimitations)`

SetFeatureFlagLimitations sets FeatureFlagLimitations field to given value.

### HasFeatureFlagLimitations

`func (o *SettingFormulaModelHaljson) HasFeatureFlagLimitations() bool`

HasFeatureFlagLimitations returns a boolean if a field has been set.

### GetLinks

`func (o *SettingFormulaModelHaljson) GetLinks() ConfigSettingFormulasModelHaljsonEmbeddedEnvironmentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *SettingFormulaModelHaljson) GetLinksOk() (*ConfigSettingFormulasModelHaljsonEmbeddedEnvironmentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *SettingFormulaModelHaljson) SetLinks(v ConfigSettingFormulasModelHaljsonEmbeddedEnvironmentLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *SettingFormulaModelHaljson) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


