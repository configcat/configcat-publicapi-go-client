# RolloutPercentageItemModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Percentage** | **int64** | The percentage value for the rule. | 
**Value** | Pointer to **interface{}** | The value to serve when the user falls in the percentage rule. It must respect the setting type. | [optional] 

## Methods

### NewRolloutPercentageItemModel

`func NewRolloutPercentageItemModel(percentage int64, ) *RolloutPercentageItemModel`

NewRolloutPercentageItemModel instantiates a new RolloutPercentageItemModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRolloutPercentageItemModelWithDefaults

`func NewRolloutPercentageItemModelWithDefaults() *RolloutPercentageItemModel`

NewRolloutPercentageItemModelWithDefaults instantiates a new RolloutPercentageItemModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPercentage

`func (o *RolloutPercentageItemModel) GetPercentage() int64`

GetPercentage returns the Percentage field if non-nil, zero value otherwise.

### GetPercentageOk

`func (o *RolloutPercentageItemModel) GetPercentageOk() (*int64, bool)`

GetPercentageOk returns a tuple with the Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentage

`func (o *RolloutPercentageItemModel) SetPercentage(v int64)`

SetPercentage sets Percentage field to given value.


### GetValue

`func (o *RolloutPercentageItemModel) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RolloutPercentageItemModel) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *RolloutPercentageItemModel) SetValue(v interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *RolloutPercentageItemModel) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *RolloutPercentageItemModel) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *RolloutPercentageItemModel) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


