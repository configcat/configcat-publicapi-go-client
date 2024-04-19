# ReplaceSettingModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hint** | Pointer to **NullableString** | A short description for the setting, shown on the Dashboard UI. | [optional] 
**Tags** | Pointer to **[]int64** | The IDs of the tags which are attached to the setting. | [optional] 
**Order** | Pointer to **NullableInt32** | The order of the Setting represented on the ConfigCat Dashboard.  Determined from an ascending sequence of integers. | [optional] 
**Name** | **string** | The name of the Feature Flag or Setting. | 

## Methods

### NewReplaceSettingModel

`func NewReplaceSettingModel(name string, ) *ReplaceSettingModel`

NewReplaceSettingModel instantiates a new ReplaceSettingModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplaceSettingModelWithDefaults

`func NewReplaceSettingModelWithDefaults() *ReplaceSettingModel`

NewReplaceSettingModelWithDefaults instantiates a new ReplaceSettingModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHint

`func (o *ReplaceSettingModel) GetHint() string`

GetHint returns the Hint field if non-nil, zero value otherwise.

### GetHintOk

`func (o *ReplaceSettingModel) GetHintOk() (*string, bool)`

GetHintOk returns a tuple with the Hint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHint

`func (o *ReplaceSettingModel) SetHint(v string)`

SetHint sets Hint field to given value.

### HasHint

`func (o *ReplaceSettingModel) HasHint() bool`

HasHint returns a boolean if a field has been set.

### SetHintNil

`func (o *ReplaceSettingModel) SetHintNil(b bool)`

 SetHintNil sets the value for Hint to be an explicit nil

### UnsetHint
`func (o *ReplaceSettingModel) UnsetHint()`

UnsetHint ensures that no value is present for Hint, not even an explicit nil
### GetTags

`func (o *ReplaceSettingModel) GetTags() []int64`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ReplaceSettingModel) GetTagsOk() (*[]int64, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ReplaceSettingModel) SetTags(v []int64)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ReplaceSettingModel) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *ReplaceSettingModel) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *ReplaceSettingModel) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetOrder

`func (o *ReplaceSettingModel) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ReplaceSettingModel) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ReplaceSettingModel) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *ReplaceSettingModel) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### SetOrderNil

`func (o *ReplaceSettingModel) SetOrderNil(b bool)`

 SetOrderNil sets the value for Order to be an explicit nil

### UnsetOrder
`func (o *ReplaceSettingModel) UnsetOrder()`

UnsetOrder ensures that no value is present for Order, not even an explicit nil
### GetName

`func (o *ReplaceSettingModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReplaceSettingModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReplaceSettingModel) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


