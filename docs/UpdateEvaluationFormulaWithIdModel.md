# UpdateEvaluationFormulaWithIdModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultValue** | [**ValueModel**](ValueModel.md) |  | 
**TargetingRules** | Pointer to [**[]TargetingRuleModel**](TargetingRuleModel.md) | The targeting rules of the Feature Flag or Setting. | [optional] 
**PercentageEvaluationAttribute** | Pointer to **NullableString** | The user attribute used for percentage evaluation. If not set, it defaults to the &#x60;Identifier&#x60; user object attribute. | [optional] 
**SettingId** | Pointer to **int32** | The identifier of the feature flag or setting. | [optional] 

## Methods

### NewUpdateEvaluationFormulaWithIdModel

`func NewUpdateEvaluationFormulaWithIdModel(defaultValue ValueModel, ) *UpdateEvaluationFormulaWithIdModel`

NewUpdateEvaluationFormulaWithIdModel instantiates a new UpdateEvaluationFormulaWithIdModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateEvaluationFormulaWithIdModelWithDefaults

`func NewUpdateEvaluationFormulaWithIdModelWithDefaults() *UpdateEvaluationFormulaWithIdModel`

NewUpdateEvaluationFormulaWithIdModelWithDefaults instantiates a new UpdateEvaluationFormulaWithIdModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultValue

`func (o *UpdateEvaluationFormulaWithIdModel) GetDefaultValue() ValueModel`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *UpdateEvaluationFormulaWithIdModel) GetDefaultValueOk() (*ValueModel, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *UpdateEvaluationFormulaWithIdModel) SetDefaultValue(v ValueModel)`

SetDefaultValue sets DefaultValue field to given value.


### GetTargetingRules

`func (o *UpdateEvaluationFormulaWithIdModel) GetTargetingRules() []TargetingRuleModel`

GetTargetingRules returns the TargetingRules field if non-nil, zero value otherwise.

### GetTargetingRulesOk

`func (o *UpdateEvaluationFormulaWithIdModel) GetTargetingRulesOk() (*[]TargetingRuleModel, bool)`

GetTargetingRulesOk returns a tuple with the TargetingRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetingRules

`func (o *UpdateEvaluationFormulaWithIdModel) SetTargetingRules(v []TargetingRuleModel)`

SetTargetingRules sets TargetingRules field to given value.

### HasTargetingRules

`func (o *UpdateEvaluationFormulaWithIdModel) HasTargetingRules() bool`

HasTargetingRules returns a boolean if a field has been set.

### SetTargetingRulesNil

`func (o *UpdateEvaluationFormulaWithIdModel) SetTargetingRulesNil(b bool)`

 SetTargetingRulesNil sets the value for TargetingRules to be an explicit nil

### UnsetTargetingRules
`func (o *UpdateEvaluationFormulaWithIdModel) UnsetTargetingRules()`

UnsetTargetingRules ensures that no value is present for TargetingRules, not even an explicit nil
### GetPercentageEvaluationAttribute

`func (o *UpdateEvaluationFormulaWithIdModel) GetPercentageEvaluationAttribute() string`

GetPercentageEvaluationAttribute returns the PercentageEvaluationAttribute field if non-nil, zero value otherwise.

### GetPercentageEvaluationAttributeOk

`func (o *UpdateEvaluationFormulaWithIdModel) GetPercentageEvaluationAttributeOk() (*string, bool)`

GetPercentageEvaluationAttributeOk returns a tuple with the PercentageEvaluationAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageEvaluationAttribute

`func (o *UpdateEvaluationFormulaWithIdModel) SetPercentageEvaluationAttribute(v string)`

SetPercentageEvaluationAttribute sets PercentageEvaluationAttribute field to given value.

### HasPercentageEvaluationAttribute

`func (o *UpdateEvaluationFormulaWithIdModel) HasPercentageEvaluationAttribute() bool`

HasPercentageEvaluationAttribute returns a boolean if a field has been set.

### SetPercentageEvaluationAttributeNil

`func (o *UpdateEvaluationFormulaWithIdModel) SetPercentageEvaluationAttributeNil(b bool)`

 SetPercentageEvaluationAttributeNil sets the value for PercentageEvaluationAttribute to be an explicit nil

### UnsetPercentageEvaluationAttribute
`func (o *UpdateEvaluationFormulaWithIdModel) UnsetPercentageEvaluationAttribute()`

UnsetPercentageEvaluationAttribute ensures that no value is present for PercentageEvaluationAttribute, not even an explicit nil
### GetSettingId

`func (o *UpdateEvaluationFormulaWithIdModel) GetSettingId() int32`

GetSettingId returns the SettingId field if non-nil, zero value otherwise.

### GetSettingIdOk

`func (o *UpdateEvaluationFormulaWithIdModel) GetSettingIdOk() (*int32, bool)`

GetSettingIdOk returns a tuple with the SettingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingId

`func (o *UpdateEvaluationFormulaWithIdModel) SetSettingId(v int32)`

SetSettingId sets SettingId field to given value.

### HasSettingId

`func (o *UpdateEvaluationFormulaWithIdModel) HasSettingId() bool`

HasSettingId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


