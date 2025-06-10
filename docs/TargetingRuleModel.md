# TargetingRuleModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conditions** | [**[]ConditionModel**](ConditionModel.md) | The list of conditions that are combined with logical AND operators. It can be one of the following: - User condition - Segment condition - Prerequisite flag condition | 
**PercentageOptions** | [**[]PercentageOptionModel**](PercentageOptionModel.md) | The percentage options from where the evaluation process will choose a value based on the flag&#39;s percentage evaluation attribute. | 
**Value** | [**NullableValueModel**](ValueModel.md) |  | 

## Methods

### NewTargetingRuleModel

`func NewTargetingRuleModel(conditions []ConditionModel, percentageOptions []PercentageOptionModel, value NullableValueModel, ) *TargetingRuleModel`

NewTargetingRuleModel instantiates a new TargetingRuleModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetingRuleModelWithDefaults

`func NewTargetingRuleModelWithDefaults() *TargetingRuleModel`

NewTargetingRuleModelWithDefaults instantiates a new TargetingRuleModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditions

`func (o *TargetingRuleModel) GetConditions() []ConditionModel`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *TargetingRuleModel) GetConditionsOk() (*[]ConditionModel, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *TargetingRuleModel) SetConditions(v []ConditionModel)`

SetConditions sets Conditions field to given value.


### GetPercentageOptions

`func (o *TargetingRuleModel) GetPercentageOptions() []PercentageOptionModel`

GetPercentageOptions returns the PercentageOptions field if non-nil, zero value otherwise.

### GetPercentageOptionsOk

`func (o *TargetingRuleModel) GetPercentageOptionsOk() (*[]PercentageOptionModel, bool)`

GetPercentageOptionsOk returns a tuple with the PercentageOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageOptions

`func (o *TargetingRuleModel) SetPercentageOptions(v []PercentageOptionModel)`

SetPercentageOptions sets PercentageOptions field to given value.


### GetValue

`func (o *TargetingRuleModel) GetValue() ValueModel`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TargetingRuleModel) GetValueOk() (*ValueModel, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TargetingRuleModel) SetValue(v ValueModel)`

SetValue sets Value field to given value.


### SetValueNil

`func (o *TargetingRuleModel) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *TargetingRuleModel) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


