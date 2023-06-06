# UpdateSettingValueWithSettingIdModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RolloutRules** | Pointer to [**[]RolloutRuleModel**](RolloutRuleModel.md) | The targeting rule collection. | [optional] 
**RolloutPercentageItems** | Pointer to [**[]RolloutPercentageItemModel**](RolloutPercentageItemModel.md) | The percentage rule collection. | [optional] 
**Value** | Pointer to **interface{}** | The value to serve. It must respect the setting type. | [optional] 
**SettingId** | Pointer to **int32** | The id of the Setting. | [optional] 

## Methods

### NewUpdateSettingValueWithSettingIdModel

`func NewUpdateSettingValueWithSettingIdModel() *UpdateSettingValueWithSettingIdModel`

NewUpdateSettingValueWithSettingIdModel instantiates a new UpdateSettingValueWithSettingIdModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSettingValueWithSettingIdModelWithDefaults

`func NewUpdateSettingValueWithSettingIdModelWithDefaults() *UpdateSettingValueWithSettingIdModel`

NewUpdateSettingValueWithSettingIdModelWithDefaults instantiates a new UpdateSettingValueWithSettingIdModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRolloutRules

`func (o *UpdateSettingValueWithSettingIdModel) GetRolloutRules() []RolloutRuleModel`

GetRolloutRules returns the RolloutRules field if non-nil, zero value otherwise.

### GetRolloutRulesOk

`func (o *UpdateSettingValueWithSettingIdModel) GetRolloutRulesOk() (*[]RolloutRuleModel, bool)`

GetRolloutRulesOk returns a tuple with the RolloutRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolloutRules

`func (o *UpdateSettingValueWithSettingIdModel) SetRolloutRules(v []RolloutRuleModel)`

SetRolloutRules sets RolloutRules field to given value.

### HasRolloutRules

`func (o *UpdateSettingValueWithSettingIdModel) HasRolloutRules() bool`

HasRolloutRules returns a boolean if a field has been set.

### SetRolloutRulesNil

`func (o *UpdateSettingValueWithSettingIdModel) SetRolloutRulesNil(b bool)`

 SetRolloutRulesNil sets the value for RolloutRules to be an explicit nil

### UnsetRolloutRules
`func (o *UpdateSettingValueWithSettingIdModel) UnsetRolloutRules()`

UnsetRolloutRules ensures that no value is present for RolloutRules, not even an explicit nil
### GetRolloutPercentageItems

`func (o *UpdateSettingValueWithSettingIdModel) GetRolloutPercentageItems() []RolloutPercentageItemModel`

GetRolloutPercentageItems returns the RolloutPercentageItems field if non-nil, zero value otherwise.

### GetRolloutPercentageItemsOk

`func (o *UpdateSettingValueWithSettingIdModel) GetRolloutPercentageItemsOk() (*[]RolloutPercentageItemModel, bool)`

GetRolloutPercentageItemsOk returns a tuple with the RolloutPercentageItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolloutPercentageItems

`func (o *UpdateSettingValueWithSettingIdModel) SetRolloutPercentageItems(v []RolloutPercentageItemModel)`

SetRolloutPercentageItems sets RolloutPercentageItems field to given value.

### HasRolloutPercentageItems

`func (o *UpdateSettingValueWithSettingIdModel) HasRolloutPercentageItems() bool`

HasRolloutPercentageItems returns a boolean if a field has been set.

### SetRolloutPercentageItemsNil

`func (o *UpdateSettingValueWithSettingIdModel) SetRolloutPercentageItemsNil(b bool)`

 SetRolloutPercentageItemsNil sets the value for RolloutPercentageItems to be an explicit nil

### UnsetRolloutPercentageItems
`func (o *UpdateSettingValueWithSettingIdModel) UnsetRolloutPercentageItems()`

UnsetRolloutPercentageItems ensures that no value is present for RolloutPercentageItems, not even an explicit nil
### GetValue

`func (o *UpdateSettingValueWithSettingIdModel) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UpdateSettingValueWithSettingIdModel) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UpdateSettingValueWithSettingIdModel) SetValue(v interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *UpdateSettingValueWithSettingIdModel) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *UpdateSettingValueWithSettingIdModel) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *UpdateSettingValueWithSettingIdModel) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetSettingId

`func (o *UpdateSettingValueWithSettingIdModel) GetSettingId() int32`

GetSettingId returns the SettingId field if non-nil, zero value otherwise.

### GetSettingIdOk

`func (o *UpdateSettingValueWithSettingIdModel) GetSettingIdOk() (*int32, bool)`

GetSettingIdOk returns a tuple with the SettingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingId

`func (o *UpdateSettingValueWithSettingIdModel) SetSettingId(v int32)`

SetSettingId sets SettingId field to given value.

### HasSettingId

`func (o *UpdateSettingValueWithSettingIdModel) HasSettingId() bool`

HasSettingId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


