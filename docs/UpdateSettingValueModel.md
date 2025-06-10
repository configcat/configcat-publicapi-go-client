# UpdateSettingValueModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RolloutRules** | Pointer to [**[]UpdateRolloutRuleModel**](UpdateRolloutRuleModel.md) | The targeting rule collection. | [optional] 
**RolloutPercentageItems** | Pointer to [**[]UpdateRolloutPercentageItemModel**](UpdateRolloutPercentageItemModel.md) | The percentage rule collection. | [optional] 
**Value** | [**SettingValueType**](SettingValueType.md) | The value to serve. It must respect the setting type. In some generated clients for strictly typed languages you may use double/float properties to handle integer values. | 

## Methods

### NewUpdateSettingValueModel

`func NewUpdateSettingValueModel(value SettingValueType, ) *UpdateSettingValueModel`

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

`func (o *UpdateSettingValueModel) GetRolloutRules() []UpdateRolloutRuleModel`

GetRolloutRules returns the RolloutRules field if non-nil, zero value otherwise.

### GetRolloutRulesOk

`func (o *UpdateSettingValueModel) GetRolloutRulesOk() (*[]UpdateRolloutRuleModel, bool)`

GetRolloutRulesOk returns a tuple with the RolloutRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolloutRules

`func (o *UpdateSettingValueModel) SetRolloutRules(v []UpdateRolloutRuleModel)`

SetRolloutRules sets RolloutRules field to given value.

### HasRolloutRules

`func (o *UpdateSettingValueModel) HasRolloutRules() bool`

HasRolloutRules returns a boolean if a field has been set.

### GetRolloutPercentageItems

`func (o *UpdateSettingValueModel) GetRolloutPercentageItems() []UpdateRolloutPercentageItemModel`

GetRolloutPercentageItems returns the RolloutPercentageItems field if non-nil, zero value otherwise.

### GetRolloutPercentageItemsOk

`func (o *UpdateSettingValueModel) GetRolloutPercentageItemsOk() (*[]UpdateRolloutPercentageItemModel, bool)`

GetRolloutPercentageItemsOk returns a tuple with the RolloutPercentageItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolloutPercentageItems

`func (o *UpdateSettingValueModel) SetRolloutPercentageItems(v []UpdateRolloutPercentageItemModel)`

SetRolloutPercentageItems sets RolloutPercentageItems field to given value.

### HasRolloutPercentageItems

`func (o *UpdateSettingValueModel) HasRolloutPercentageItems() bool`

HasRolloutPercentageItems returns a boolean if a field has been set.

### GetValue

`func (o *UpdateSettingValueModel) GetValue() SettingValueType`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UpdateSettingValueModel) GetValueOk() (*SettingValueType, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UpdateSettingValueModel) SetValue(v SettingValueType)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


