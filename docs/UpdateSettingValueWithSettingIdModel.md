# UpdateSettingValueWithSettingIdModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RolloutRules** | Pointer to [**[]UpdateRolloutRuleModel**](UpdateRolloutRuleModel.md) | The targeting rule collection. | [optional] 
**RolloutPercentageItems** | Pointer to [**[]UpdateRolloutPercentageItemModel**](UpdateRolloutPercentageItemModel.md) | The percentage rule collection. | [optional] 
**Value** | [**SettingValueType**](SettingValueType.md) | The value to serve. It must respect the setting type. In some generated clients for strictly typed languages you may use double/float properties to handle integer values. | 
**SettingId** | Pointer to **int32** | The id of the Setting. | [optional] 

## Methods

### NewUpdateSettingValueWithSettingIdModel

`func NewUpdateSettingValueWithSettingIdModel(value SettingValueType, ) *UpdateSettingValueWithSettingIdModel`

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

`func (o *UpdateSettingValueWithSettingIdModel) GetRolloutRules() []UpdateRolloutRuleModel`

GetRolloutRules returns the RolloutRules field if non-nil, zero value otherwise.

### GetRolloutRulesOk

`func (o *UpdateSettingValueWithSettingIdModel) GetRolloutRulesOk() (*[]UpdateRolloutRuleModel, bool)`

GetRolloutRulesOk returns a tuple with the RolloutRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolloutRules

`func (o *UpdateSettingValueWithSettingIdModel) SetRolloutRules(v []UpdateRolloutRuleModel)`

SetRolloutRules sets RolloutRules field to given value.

### HasRolloutRules

`func (o *UpdateSettingValueWithSettingIdModel) HasRolloutRules() bool`

HasRolloutRules returns a boolean if a field has been set.

### GetRolloutPercentageItems

`func (o *UpdateSettingValueWithSettingIdModel) GetRolloutPercentageItems() []UpdateRolloutPercentageItemModel`

GetRolloutPercentageItems returns the RolloutPercentageItems field if non-nil, zero value otherwise.

### GetRolloutPercentageItemsOk

`func (o *UpdateSettingValueWithSettingIdModel) GetRolloutPercentageItemsOk() (*[]UpdateRolloutPercentageItemModel, bool)`

GetRolloutPercentageItemsOk returns a tuple with the RolloutPercentageItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolloutPercentageItems

`func (o *UpdateSettingValueWithSettingIdModel) SetRolloutPercentageItems(v []UpdateRolloutPercentageItemModel)`

SetRolloutPercentageItems sets RolloutPercentageItems field to given value.

### HasRolloutPercentageItems

`func (o *UpdateSettingValueWithSettingIdModel) HasRolloutPercentageItems() bool`

HasRolloutPercentageItems returns a boolean if a field has been set.

### GetValue

`func (o *UpdateSettingValueWithSettingIdModel) GetValue() SettingValueType`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UpdateSettingValueWithSettingIdModel) GetValueOk() (*SettingValueType, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UpdateSettingValueWithSettingIdModel) SetValue(v SettingValueType)`

SetValue sets Value field to given value.


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


