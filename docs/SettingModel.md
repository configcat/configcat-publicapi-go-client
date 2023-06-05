# SettingModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SettingId** | Pointer to **int32** |  | [optional] 
**Key** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Hint** | Pointer to **NullableString** |  | [optional] 
**Order** | Pointer to **int32** |  | [optional] 
**SettingType** | Pointer to [**SettingType**](SettingType.md) |  | [optional] 
**ConfigId** | Pointer to **string** |  | [optional] 
**ConfigName** | Pointer to **NullableString** |  | [optional] 
**Tags** | Pointer to [**[]TagModel**](TagModel.md) |  | [optional] 

## Methods

### NewSettingModel

`func NewSettingModel() *SettingModel`

NewSettingModel instantiates a new SettingModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettingModelWithDefaults

`func NewSettingModelWithDefaults() *SettingModel`

NewSettingModelWithDefaults instantiates a new SettingModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSettingId

`func (o *SettingModel) GetSettingId() int32`

GetSettingId returns the SettingId field if non-nil, zero value otherwise.

### GetSettingIdOk

`func (o *SettingModel) GetSettingIdOk() (*int32, bool)`

GetSettingIdOk returns a tuple with the SettingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingId

`func (o *SettingModel) SetSettingId(v int32)`

SetSettingId sets SettingId field to given value.

### HasSettingId

`func (o *SettingModel) HasSettingId() bool`

HasSettingId returns a boolean if a field has been set.

### GetKey

`func (o *SettingModel) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SettingModel) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SettingModel) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *SettingModel) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *SettingModel) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *SettingModel) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetName

`func (o *SettingModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SettingModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SettingModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SettingModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *SettingModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *SettingModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetHint

`func (o *SettingModel) GetHint() string`

GetHint returns the Hint field if non-nil, zero value otherwise.

### GetHintOk

`func (o *SettingModel) GetHintOk() (*string, bool)`

GetHintOk returns a tuple with the Hint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHint

`func (o *SettingModel) SetHint(v string)`

SetHint sets Hint field to given value.

### HasHint

`func (o *SettingModel) HasHint() bool`

HasHint returns a boolean if a field has been set.

### SetHintNil

`func (o *SettingModel) SetHintNil(b bool)`

 SetHintNil sets the value for Hint to be an explicit nil

### UnsetHint
`func (o *SettingModel) UnsetHint()`

UnsetHint ensures that no value is present for Hint, not even an explicit nil
### GetOrder

`func (o *SettingModel) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *SettingModel) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *SettingModel) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *SettingModel) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetSettingType

`func (o *SettingModel) GetSettingType() SettingType`

GetSettingType returns the SettingType field if non-nil, zero value otherwise.

### GetSettingTypeOk

`func (o *SettingModel) GetSettingTypeOk() (*SettingType, bool)`

GetSettingTypeOk returns a tuple with the SettingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingType

`func (o *SettingModel) SetSettingType(v SettingType)`

SetSettingType sets SettingType field to given value.

### HasSettingType

`func (o *SettingModel) HasSettingType() bool`

HasSettingType returns a boolean if a field has been set.

### GetConfigId

`func (o *SettingModel) GetConfigId() string`

GetConfigId returns the ConfigId field if non-nil, zero value otherwise.

### GetConfigIdOk

`func (o *SettingModel) GetConfigIdOk() (*string, bool)`

GetConfigIdOk returns a tuple with the ConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigId

`func (o *SettingModel) SetConfigId(v string)`

SetConfigId sets ConfigId field to given value.

### HasConfigId

`func (o *SettingModel) HasConfigId() bool`

HasConfigId returns a boolean if a field has been set.

### GetConfigName

`func (o *SettingModel) GetConfigName() string`

GetConfigName returns the ConfigName field if non-nil, zero value otherwise.

### GetConfigNameOk

`func (o *SettingModel) GetConfigNameOk() (*string, bool)`

GetConfigNameOk returns a tuple with the ConfigName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigName

`func (o *SettingModel) SetConfigName(v string)`

SetConfigName sets ConfigName field to given value.

### HasConfigName

`func (o *SettingModel) HasConfigName() bool`

HasConfigName returns a boolean if a field has been set.

### SetConfigNameNil

`func (o *SettingModel) SetConfigNameNil(b bool)`

 SetConfigNameNil sets the value for ConfigName to be an explicit nil

### UnsetConfigName
`func (o *SettingModel) UnsetConfigName()`

UnsetConfigName ensures that no value is present for ConfigName, not even an explicit nil
### GetTags

`func (o *SettingModel) GetTags() []TagModel`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SettingModel) GetTagsOk() (*[]TagModel, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SettingModel) SetTags(v []TagModel)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SettingModel) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *SettingModel) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *SettingModel) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


