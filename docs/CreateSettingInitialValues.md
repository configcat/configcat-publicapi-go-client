# CreateSettingInitialValues

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hint** | Pointer to **NullableString** | A short description for the setting, shown on the Dashboard UI. | [optional] 
**Tags** | Pointer to **[]int64** | The IDs of the tags which are attached to the setting. | [optional] 
**Order** | Pointer to **NullableInt32** | The order of the Setting represented on the ConfigCat Dashboard.  Determined from an ascending sequence of integers. | [optional] 
**Key** | **string** | The key of the Feature Flag or Setting. | 
**Name** | **string** | The name of the Feature Flag or Setting. | 
**SettingType** | [**SettingType**](SettingType.md) |  | 
**InitialValues** | Pointer to [**[]InitialValue**](InitialValue.md) | Optional, initial value of the Feature Flag or Setting in the given Environments. | [optional] 

## Methods

### NewCreateSettingInitialValues

`func NewCreateSettingInitialValues(key string, name string, settingType SettingType, ) *CreateSettingInitialValues`

NewCreateSettingInitialValues instantiates a new CreateSettingInitialValues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSettingInitialValuesWithDefaults

`func NewCreateSettingInitialValuesWithDefaults() *CreateSettingInitialValues`

NewCreateSettingInitialValuesWithDefaults instantiates a new CreateSettingInitialValues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHint

`func (o *CreateSettingInitialValues) GetHint() string`

GetHint returns the Hint field if non-nil, zero value otherwise.

### GetHintOk

`func (o *CreateSettingInitialValues) GetHintOk() (*string, bool)`

GetHintOk returns a tuple with the Hint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHint

`func (o *CreateSettingInitialValues) SetHint(v string)`

SetHint sets Hint field to given value.

### HasHint

`func (o *CreateSettingInitialValues) HasHint() bool`

HasHint returns a boolean if a field has been set.

### SetHintNil

`func (o *CreateSettingInitialValues) SetHintNil(b bool)`

 SetHintNil sets the value for Hint to be an explicit nil

### UnsetHint
`func (o *CreateSettingInitialValues) UnsetHint()`

UnsetHint ensures that no value is present for Hint, not even an explicit nil
### GetTags

`func (o *CreateSettingInitialValues) GetTags() []int64`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateSettingInitialValues) GetTagsOk() (*[]int64, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateSettingInitialValues) SetTags(v []int64)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateSettingInitialValues) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *CreateSettingInitialValues) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *CreateSettingInitialValues) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetOrder

`func (o *CreateSettingInitialValues) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *CreateSettingInitialValues) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *CreateSettingInitialValues) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *CreateSettingInitialValues) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### SetOrderNil

`func (o *CreateSettingInitialValues) SetOrderNil(b bool)`

 SetOrderNil sets the value for Order to be an explicit nil

### UnsetOrder
`func (o *CreateSettingInitialValues) UnsetOrder()`

UnsetOrder ensures that no value is present for Order, not even an explicit nil
### GetKey

`func (o *CreateSettingInitialValues) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CreateSettingInitialValues) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CreateSettingInitialValues) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *CreateSettingInitialValues) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateSettingInitialValues) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateSettingInitialValues) SetName(v string)`

SetName sets Name field to given value.


### GetSettingType

`func (o *CreateSettingInitialValues) GetSettingType() SettingType`

GetSettingType returns the SettingType field if non-nil, zero value otherwise.

### GetSettingTypeOk

`func (o *CreateSettingInitialValues) GetSettingTypeOk() (*SettingType, bool)`

GetSettingTypeOk returns a tuple with the SettingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingType

`func (o *CreateSettingInitialValues) SetSettingType(v SettingType)`

SetSettingType sets SettingType field to given value.


### GetInitialValues

`func (o *CreateSettingInitialValues) GetInitialValues() []InitialValue`

GetInitialValues returns the InitialValues field if non-nil, zero value otherwise.

### GetInitialValuesOk

`func (o *CreateSettingInitialValues) GetInitialValuesOk() (*[]InitialValue, bool)`

GetInitialValuesOk returns a tuple with the InitialValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialValues

`func (o *CreateSettingInitialValues) SetInitialValues(v []InitialValue)`

SetInitialValues sets InitialValues field to given value.

### HasInitialValues

`func (o *CreateSettingInitialValues) HasInitialValues() bool`

HasInitialValues returns a boolean if a field has been set.

### SetInitialValuesNil

`func (o *CreateSettingInitialValues) SetInitialValuesNil(b bool)`

 SetInitialValuesNil sets the value for InitialValues to be an explicit nil

### UnsetInitialValues
`func (o *CreateSettingInitialValues) UnsetInitialValues()`

UnsetInitialValues ensures that no value is present for InitialValues, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


