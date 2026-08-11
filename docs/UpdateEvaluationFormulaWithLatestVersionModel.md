# UpdateEvaluationFormulaWithLatestVersionModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultValue** | [**UpdateValueModel**](UpdateValueModel.md) |  | 
**TargetingRules** | Pointer to [**[]UpdateTargetingRuleModel**](UpdateTargetingRuleModel.md) | The targeting rules of the Feature Flag or Setting. | [optional] 
**PercentageEvaluationAttribute** | Pointer to **NullableString** | The user attribute used for percentage evaluation. If not set, it defaults to the &#x60;Identifier&#x60; user object attribute. | [optional] 
**SettingId** | Pointer to **int32** | The identifier of the feature flag or setting. | [optional] 
**LatestVersionId** | Pointer to **NullableString** | The version identifier of the last change made to the Feature Flag or Setting in the Environment. It can be used to make sure concurrent updates are not overwriting each other. If provided and the version identifier does not match the current version, the update will be rejected with a 409 Conflict response. The latest version id can be acquired from the &#x60;LastVersionId&#x60; property of the response models. | [optional] 

## Methods

### NewUpdateEvaluationFormulaWithLatestVersionModel

`func NewUpdateEvaluationFormulaWithLatestVersionModel(defaultValue UpdateValueModel, ) *UpdateEvaluationFormulaWithLatestVersionModel`

NewUpdateEvaluationFormulaWithLatestVersionModel instantiates a new UpdateEvaluationFormulaWithLatestVersionModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateEvaluationFormulaWithLatestVersionModelWithDefaults

`func NewUpdateEvaluationFormulaWithLatestVersionModelWithDefaults() *UpdateEvaluationFormulaWithLatestVersionModel`

NewUpdateEvaluationFormulaWithLatestVersionModelWithDefaults instantiates a new UpdateEvaluationFormulaWithLatestVersionModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultValue

`func (o *UpdateEvaluationFormulaWithLatestVersionModel) GetDefaultValue() UpdateValueModel`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *UpdateEvaluationFormulaWithLatestVersionModel) GetDefaultValueOk() (*UpdateValueModel, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *UpdateEvaluationFormulaWithLatestVersionModel) SetDefaultValue(v UpdateValueModel)`

SetDefaultValue sets DefaultValue field to given value.


### GetTargetingRules

`func (o *UpdateEvaluationFormulaWithLatestVersionModel) GetTargetingRules() []UpdateTargetingRuleModel`

GetTargetingRules returns the TargetingRules field if non-nil, zero value otherwise.

### GetTargetingRulesOk

`func (o *UpdateEvaluationFormulaWithLatestVersionModel) GetTargetingRulesOk() (*[]UpdateTargetingRuleModel, bool)`

GetTargetingRulesOk returns a tuple with the TargetingRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetingRules

`func (o *UpdateEvaluationFormulaWithLatestVersionModel) SetTargetingRules(v []UpdateTargetingRuleModel)`

SetTargetingRules sets TargetingRules field to given value.

### HasTargetingRules

`func (o *UpdateEvaluationFormulaWithLatestVersionModel) HasTargetingRules() bool`

HasTargetingRules returns a boolean if a field has been set.

### SetTargetingRulesNil

`func (o *UpdateEvaluationFormulaWithLatestVersionModel) SetTargetingRulesNil(b bool)`

 SetTargetingRulesNil sets the value for TargetingRules to be an explicit nil

### UnsetTargetingRules
`func (o *UpdateEvaluationFormulaWithLatestVersionModel) UnsetTargetingRules()`

UnsetTargetingRules ensures that no value is present for TargetingRules, not even an explicit nil
### GetPercentageEvaluationAttribute

`func (o *UpdateEvaluationFormulaWithLatestVersionModel) GetPercentageEvaluationAttribute() string`

GetPercentageEvaluationAttribute returns the PercentageEvaluationAttribute field if non-nil, zero value otherwise.

### GetPercentageEvaluationAttributeOk

`func (o *UpdateEvaluationFormulaWithLatestVersionModel) GetPercentageEvaluationAttributeOk() (*string, bool)`

GetPercentageEvaluationAttributeOk returns a tuple with the PercentageEvaluationAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageEvaluationAttribute

`func (o *UpdateEvaluationFormulaWithLatestVersionModel) SetPercentageEvaluationAttribute(v string)`

SetPercentageEvaluationAttribute sets PercentageEvaluationAttribute field to given value.

### HasPercentageEvaluationAttribute

`func (o *UpdateEvaluationFormulaWithLatestVersionModel) HasPercentageEvaluationAttribute() bool`

HasPercentageEvaluationAttribute returns a boolean if a field has been set.

### SetPercentageEvaluationAttributeNil

`func (o *UpdateEvaluationFormulaWithLatestVersionModel) SetPercentageEvaluationAttributeNil(b bool)`

 SetPercentageEvaluationAttributeNil sets the value for PercentageEvaluationAttribute to be an explicit nil

### UnsetPercentageEvaluationAttribute
`func (o *UpdateEvaluationFormulaWithLatestVersionModel) UnsetPercentageEvaluationAttribute()`

UnsetPercentageEvaluationAttribute ensures that no value is present for PercentageEvaluationAttribute, not even an explicit nil
### GetSettingId

`func (o *UpdateEvaluationFormulaWithLatestVersionModel) GetSettingId() int32`

GetSettingId returns the SettingId field if non-nil, zero value otherwise.

### GetSettingIdOk

`func (o *UpdateEvaluationFormulaWithLatestVersionModel) GetSettingIdOk() (*int32, bool)`

GetSettingIdOk returns a tuple with the SettingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingId

`func (o *UpdateEvaluationFormulaWithLatestVersionModel) SetSettingId(v int32)`

SetSettingId sets SettingId field to given value.

### HasSettingId

`func (o *UpdateEvaluationFormulaWithLatestVersionModel) HasSettingId() bool`

HasSettingId returns a boolean if a field has been set.

### GetLatestVersionId

`func (o *UpdateEvaluationFormulaWithLatestVersionModel) GetLatestVersionId() string`

GetLatestVersionId returns the LatestVersionId field if non-nil, zero value otherwise.

### GetLatestVersionIdOk

`func (o *UpdateEvaluationFormulaWithLatestVersionModel) GetLatestVersionIdOk() (*string, bool)`

GetLatestVersionIdOk returns a tuple with the LatestVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestVersionId

`func (o *UpdateEvaluationFormulaWithLatestVersionModel) SetLatestVersionId(v string)`

SetLatestVersionId sets LatestVersionId field to given value.

### HasLatestVersionId

`func (o *UpdateEvaluationFormulaWithLatestVersionModel) HasLatestVersionId() bool`

HasLatestVersionId returns a boolean if a field has been set.

### SetLatestVersionIdNil

`func (o *UpdateEvaluationFormulaWithLatestVersionModel) SetLatestVersionIdNil(b bool)`

 SetLatestVersionIdNil sets the value for LatestVersionId to be an explicit nil

### UnsetLatestVersionId
`func (o *UpdateEvaluationFormulaWithLatestVersionModel) UnsetLatestVersionId()`

UnsetLatestVersionId ensures that no value is present for LatestVersionId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


