# UpdateTargetingRuleModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conditions** | Pointer to [**[]UpdateConditionModel**](UpdateConditionModel.md) | The list of conditions that are combined with logical AND operators. It can be one of the following: - User condition - Segment condition - Prerequisite flag condition | [optional] 
**PercentageOptions** | Pointer to [**[]UpdatePercentageOptionModel**](UpdatePercentageOptionModel.md) | The percentage options from where the evaluation process will choose a value based on the flag&#39;s percentage evaluation attribute. | [optional] 
**Value** | Pointer to [**NullableUpdateValueModel**](UpdateValueModel.md) |  | [optional] 

## Methods

### NewUpdateTargetingRuleModel

`func NewUpdateTargetingRuleModel() *UpdateTargetingRuleModel`

NewUpdateTargetingRuleModel instantiates a new UpdateTargetingRuleModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTargetingRuleModelWithDefaults

`func NewUpdateTargetingRuleModelWithDefaults() *UpdateTargetingRuleModel`

NewUpdateTargetingRuleModelWithDefaults instantiates a new UpdateTargetingRuleModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditions

`func (o *UpdateTargetingRuleModel) GetConditions() []UpdateConditionModel`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *UpdateTargetingRuleModel) GetConditionsOk() (*[]UpdateConditionModel, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *UpdateTargetingRuleModel) SetConditions(v []UpdateConditionModel)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *UpdateTargetingRuleModel) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### SetConditionsNil

`func (o *UpdateTargetingRuleModel) SetConditionsNil(b bool)`

 SetConditionsNil sets the value for Conditions to be an explicit nil

### UnsetConditions
`func (o *UpdateTargetingRuleModel) UnsetConditions()`

UnsetConditions ensures that no value is present for Conditions, not even an explicit nil
### GetPercentageOptions

`func (o *UpdateTargetingRuleModel) GetPercentageOptions() []UpdatePercentageOptionModel`

GetPercentageOptions returns the PercentageOptions field if non-nil, zero value otherwise.

### GetPercentageOptionsOk

`func (o *UpdateTargetingRuleModel) GetPercentageOptionsOk() (*[]UpdatePercentageOptionModel, bool)`

GetPercentageOptionsOk returns a tuple with the PercentageOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageOptions

`func (o *UpdateTargetingRuleModel) SetPercentageOptions(v []UpdatePercentageOptionModel)`

SetPercentageOptions sets PercentageOptions field to given value.

### HasPercentageOptions

`func (o *UpdateTargetingRuleModel) HasPercentageOptions() bool`

HasPercentageOptions returns a boolean if a field has been set.

### SetPercentageOptionsNil

`func (o *UpdateTargetingRuleModel) SetPercentageOptionsNil(b bool)`

 SetPercentageOptionsNil sets the value for PercentageOptions to be an explicit nil

### UnsetPercentageOptions
`func (o *UpdateTargetingRuleModel) UnsetPercentageOptions()`

UnsetPercentageOptions ensures that no value is present for PercentageOptions, not even an explicit nil
### GetValue

`func (o *UpdateTargetingRuleModel) GetValue() UpdateValueModel`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UpdateTargetingRuleModel) GetValueOk() (*UpdateValueModel, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UpdateTargetingRuleModel) SetValue(v UpdateValueModel)`

SetValue sets Value field to given value.

### HasValue

`func (o *UpdateTargetingRuleModel) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *UpdateTargetingRuleModel) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *UpdateTargetingRuleModel) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


