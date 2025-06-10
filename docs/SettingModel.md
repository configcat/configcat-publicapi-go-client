# SettingModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SettingId** | **int32** | Identifier of the Feature Flag or Setting. | 
**Key** | **string** | Key of the Feature Flag or Setting. | 
**Name** | **string** | Name of the Feature Flag or Setting. | 
**Hint** | **NullableString** | Description of the Feature Flag or Setting. | 
**Order** | **int32** | The order of the Feature Flag or Setting represented on the ConfigCat Dashboard. | 
**SettingType** | [**SettingType**](SettingType.md) |  | 
**ConfigId** | **string** | Identifier of the Feature Flag&#39;s Config. | 
**ConfigName** | **string** | Name of the Feature Flag&#39;s Config. | 
**CreatedAt** | **NullableTime** | The creation time of the Feature Flag or Setting. | 
**Tags** | [**[]TagModel**](TagModel.md) | The tags attached to the Feature Flag or Setting. | 

## Methods

### NewSettingModel

`func NewSettingModel(settingId int32, key string, name string, hint NullableString, order int32, settingType SettingType, configId string, configName string, createdAt NullableTime, tags []TagModel, ) *SettingModel`

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


### GetCreatedAt

`func (o *SettingModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SettingModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SettingModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *SettingModel) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *SettingModel) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


