# UpdateSettingValueModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RolloutRules** | Pointer to [**[]RolloutRuleModel**](RolloutRuleModel.md) | The targeting rule collection. | [optional] 
**RolloutPercentageItems** | Pointer to [**[]RolloutPercentageItemModel**](RolloutPercentageItemModel.md) | The percentage rule collection. | [optional] 
**Value** | Pointer to **interface{}** | The value to serve. It must respect the setting type. | [optional] 

## Methods

### NewUpdateSettingValueModel

`func NewUpdateSettingValueModel() *UpdateSettingValueModel`

NewUpdateSettingValueModel instantiates a new UpdateSettingValueModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSettingValueModelWithDefaults

`func NewUpdateSettingValueModelWithDefaults() *UpdateSettingValueModel`

NewUpdateSettingValueModelWithDefaults instantiates a new UpdateSettingValueModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRolloutRules

`func (o *UpdateSettingValueModel) GetRolloutRules() []RolloutRuleModel`

GetRolloutRules returns the RolloutRules field if non-nil, zero value otherwise.

### GetRolloutRulesOk

`func (o *UpdateSettingValueModel) GetRolloutRulesOk() (*[]RolloutRuleModel, bool)`

GetRolloutRulesOk returns a tuple with the RolloutRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolloutRules

`func (o *UpdateSettingValueModel) SetRolloutRules(v []RolloutRuleModel)`

SetRolloutRules sets RolloutRules field to given value.

### HasRolloutRules

`func (o *UpdateSettingValueModel) HasRolloutRules() bool`

HasRolloutRules returns a boolean if a field has been set.

### SetRolloutRulesNil

`func (o *UpdateSettingValueModel) SetRolloutRulesNil(b bool)`

 SetRolloutRulesNil sets the value for RolloutRules to be an explicit nil

### UnsetRolloutRules
`func (o *UpdateSettingValueModel) UnsetRolloutRules()`

UnsetRolloutRules ensures that no value is present for RolloutRules, not even an explicit nil
### GetRolloutPercentageItems

`func (o *UpdateSettingValueModel) GetRolloutPercentageItems() []RolloutPercentageItemModel`

GetRolloutPercentageItems returns the RolloutPercentageItems field if non-nil, zero value otherwise.

### GetRolloutPercentageItemsOk

`func (o *UpdateSettingValueModel) GetRolloutPercentageItemsOk() (*[]RolloutPercentageItemModel, bool)`

GetRolloutPercentageItemsOk returns a tuple with the RolloutPercentageItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolloutPercentageItems

`func (o *UpdateSettingValueModel) SetRolloutPercentageItems(v []RolloutPercentageItemModel)`

SetRolloutPercentageItems sets RolloutPercentageItems field to given value.

### HasRolloutPercentageItems

`func (o *UpdateSettingValueModel) HasRolloutPercentageItems() bool`

HasRolloutPercentageItems returns a boolean if a field has been set.

### SetRolloutPercentageItemsNil

`func (o *UpdateSettingValueModel) SetRolloutPercentageItemsNil(b bool)`

 SetRolloutPercentageItemsNil sets the value for RolloutPercentageItems to be an explicit nil

### UnsetRolloutPercentageItems
`func (o *UpdateSettingValueModel) UnsetRolloutPercentageItems()`

UnsetRolloutPercentageItems ensures that no value is present for RolloutPercentageItems, not even an explicit nil
### GetValue

`func (o *UpdateSettingValueModel) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UpdateSettingValueModel) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UpdateSettingValueModel) SetValue(v interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *UpdateSettingValueModel) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *UpdateSettingValueModel) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *UpdateSettingValueModel) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


