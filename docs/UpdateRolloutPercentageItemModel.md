# UpdateRolloutPercentageItemModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Percentage** | **int64** | The percentage value for the rule. | 
**Value** | [**SettingValueType**](SettingValueType.md) | The value to serve when the user falls in the percentage rule. It must respect the setting type. In some generated clients for strictly typed languages you may use double/float properties to handle integer values. | 

## Methods

### NewUpdateRolloutPercentageItemModel

`func NewUpdateRolloutPercentageItemModel(percentage int64, value SettingValueType, ) *UpdateRolloutPercentageItemModel`

NewUpdateRolloutPercentageItemModel instantiates a new UpdateRolloutPercentageItemModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRolloutPercentageItemModelWithDefaults

`func NewUpdateRolloutPercentageItemModelWithDefaults() *UpdateRolloutPercentageItemModel`

NewUpdateRolloutPercentageItemModelWithDefaults instantiates a new UpdateRolloutPercentageItemModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPercentage

`func (o *UpdateRolloutPercentageItemModel) GetPercentage() int64`

GetPercentage returns the Percentage field if non-nil, zero value otherwise.

### GetPercentageOk

`func (o *UpdateRolloutPercentageItemModel) GetPercentageOk() (*int64, bool)`

GetPercentageOk returns a tuple with the Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentage

`func (o *UpdateRolloutPercentageItemModel) SetPercentage(v int64)`

SetPercentage sets Percentage field to given value.


### GetValue

`func (o *UpdateRolloutPercentageItemModel) GetValue() SettingValueType`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UpdateRolloutPercentageItemModel) GetValueOk() (*SettingValueType, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UpdateRolloutPercentageItemModel) SetValue(v SettingValueType)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


