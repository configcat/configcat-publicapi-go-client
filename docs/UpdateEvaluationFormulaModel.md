# UpdateEvaluationFormulaModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultValue** | [**ValueModel**](ValueModel.md) |  | 
**TargetingRules** | Pointer to [**[]TargetingRuleModel**](TargetingRuleModel.md) | The targeting rules of the Feature Flag or Setting. | [optional] 
**PercentageEvaluationAttribute** | Pointer to **NullableString** | The user attribute used for percentage evaluation. If not set, it defaults to the &#x60;Identifier&#x60; user object attribute. | [optional] 

## Methods

### NewUpdateEvaluationFormulaModel

`func NewUpdateEvaluationFormulaModel(defaultValue ValueModel, ) *UpdateEvaluationFormulaModel`

NewUpdateEvaluationFormulaModel instantiates a new UpdateEvaluationFormulaModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateEvaluationFormulaModelWithDefaults

`func NewUpdateEvaluationFormulaModelWithDefaults() *UpdateEvaluationFormulaModel`

NewUpdateEvaluationFormulaModelWithDefaults instantiates a new UpdateEvaluationFormulaModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultValue

`func (o *UpdateEvaluationFormulaModel) GetDefaultValue() ValueModel`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *UpdateEvaluationFormulaModel) GetDefaultValueOk() (*ValueModel, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *UpdateEvaluationFormulaModel) SetDefaultValue(v ValueModel)`

SetDefaultValue sets DefaultValue field to given value.


### GetTargetingRules

`func (o *UpdateEvaluationFormulaModel) GetTargetingRules() []TargetingRuleModel`

GetTargetingRules returns the TargetingRules field if non-nil, zero value otherwise.

### GetTargetingRulesOk

`func (o *UpdateEvaluationFormulaModel) GetTargetingRulesOk() (*[]TargetingRuleModel, bool)`

GetTargetingRulesOk returns a tuple with the TargetingRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetingRules

`func (o *UpdateEvaluationFormulaModel) SetTargetingRules(v []TargetingRuleModel)`

SetTargetingRules sets TargetingRules field to given value.

### HasTargetingRules

`func (o *UpdateEvaluationFormulaModel) HasTargetingRules() bool`

HasTargetingRules returns a boolean if a field has been set.

### SetTargetingRulesNil

`func (o *UpdateEvaluationFormulaModel) SetTargetingRulesNil(b bool)`

 SetTargetingRulesNil sets the value for TargetingRules to be an explicit nil

### UnsetTargetingRules
`func (o *UpdateEvaluationFormulaModel) UnsetTargetingRules()`

UnsetTargetingRules ensures that no value is present for TargetingRules, not even an explicit nil
### GetPercentageEvaluationAttribute

`func (o *UpdateEvaluationFormulaModel) GetPercentageEvaluationAttribute() string`

GetPercentageEvaluationAttribute returns the PercentageEvaluationAttribute field if non-nil, zero value otherwise.

### GetPercentageEvaluationAttributeOk

`func (o *UpdateEvaluationFormulaModel) GetPercentageEvaluationAttributeOk() (*string, bool)`

GetPercentageEvaluationAttributeOk returns a tuple with the PercentageEvaluationAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageEvaluationAttribute

`func (o *UpdateEvaluationFormulaModel) SetPercentageEvaluationAttribute(v string)`

SetPercentageEvaluationAttribute sets PercentageEvaluationAttribute field to given value.

### HasPercentageEvaluationAttribute

`func (o *UpdateEvaluationFormulaModel) HasPercentageEvaluationAttribute() bool`

HasPercentageEvaluationAttribute returns a boolean if a field has been set.

### SetPercentageEvaluationAttributeNil

`func (o *UpdateEvaluationFormulaModel) SetPercentageEvaluationAttributeNil(b bool)`

 SetPercentageEvaluationAttributeNil sets the value for PercentageEvaluationAttribute to be an explicit nil

### UnsetPercentageEvaluationAttribute
`func (o *UpdateEvaluationFormulaModel) UnsetPercentageEvaluationAttribute()`

UnsetPercentageEvaluationAttribute ensures that no value is present for PercentageEvaluationAttribute, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


