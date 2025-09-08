# SettingDataV2Model

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SettingId** | **int32** | Identifier of the Feature Flag or Setting. | 
**Key** | **string** | Key of the Feature Flag or Setting. | 
**Name** | **string** | Name of the Feature Flag or Setting. | 
**Hint** | **NullableString** | Description of the Feature Flag or Setting. | 
**SettingType** | [**SettingType**](SettingType.md) |  | 
**Order** | **int32** | The order of the Feature Flag or Setting represented on the ConfigCat Dashboard. | 
**CreatedAt** | **NullableTime** | The creation time of the Feature Flag or Setting. | 
**CreatorEmail** | **NullableString** | The user&#39;s email address who created the Feature Flag or Setting. | 
**CreatorFullName** | **NullableString** | The user&#39;s name who created the Feature Flag or Setting. | 
**PredefinedVariations** | [**[]PredefinedVariationModel**](PredefinedVariationModel.md) | A collection of Variations for a Feature Flag or Setting. | 
**IsWatching** | **bool** |  | 

## Methods

### NewSettingDataV2Model

`func NewSettingDataV2Model(settingId int32, key string, name string, hint NullableString, settingType SettingType, order int32, createdAt NullableTime, creatorEmail NullableString, creatorFullName NullableString, predefinedVariations []PredefinedVariationModel, isWatching bool, ) *SettingDataV2Model`

NewSettingDataV2Model instantiates a new SettingDataV2Model object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettingDataV2ModelWithDefaults

`func NewSettingDataV2ModelWithDefaults() *SettingDataV2Model`

NewSettingDataV2ModelWithDefaults instantiates a new SettingDataV2Model object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSettingId

`func (o *SettingDataV2Model) GetSettingId() int32`

GetSettingId returns the SettingId field if non-nil, zero value otherwise.

### GetSettingIdOk

`func (o *SettingDataV2Model) GetSettingIdOk() (*int32, bool)`

GetSettingIdOk returns a tuple with the SettingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingId

`func (o *SettingDataV2Model) SetSettingId(v int32)`

SetSettingId sets SettingId field to given value.


### GetKey

`func (o *SettingDataV2Model) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SettingDataV2Model) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SettingDataV2Model) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *SettingDataV2Model) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SettingDataV2Model) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SettingDataV2Model) SetName(v string)`

SetName sets Name field to given value.


### GetHint

`func (o *SettingDataV2Model) GetHint() string`

GetHint returns the Hint field if non-nil, zero value otherwise.

### GetHintOk

`func (o *SettingDataV2Model) GetHintOk() (*string, bool)`

GetHintOk returns a tuple with the Hint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHint

`func (o *SettingDataV2Model) SetHint(v string)`

SetHint sets Hint field to given value.


### SetHintNil

`func (o *SettingDataV2Model) SetHintNil(b bool)`

 SetHintNil sets the value for Hint to be an explicit nil

### UnsetHint
`func (o *SettingDataV2Model) UnsetHint()`

UnsetHint ensures that no value is present for Hint, not even an explicit nil
### GetSettingType

`func (o *SettingDataV2Model) GetSettingType() SettingType`

GetSettingType returns the SettingType field if non-nil, zero value otherwise.

### GetSettingTypeOk

`func (o *SettingDataV2Model) GetSettingTypeOk() (*SettingType, bool)`

GetSettingTypeOk returns a tuple with the SettingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingType

`func (o *SettingDataV2Model) SetSettingType(v SettingType)`

SetSettingType sets SettingType field to given value.


### GetOrder

`func (o *SettingDataV2Model) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *SettingDataV2Model) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *SettingDataV2Model) SetOrder(v int32)`

SetOrder sets Order field to given value.


### GetCreatedAt

`func (o *SettingDataV2Model) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SettingDataV2Model) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SettingDataV2Model) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *SettingDataV2Model) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *SettingDataV2Model) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetCreatorEmail

`func (o *SettingDataV2Model) GetCreatorEmail() string`

GetCreatorEmail returns the CreatorEmail field if non-nil, zero value otherwise.

### GetCreatorEmailOk

`func (o *SettingDataV2Model) GetCreatorEmailOk() (*string, bool)`

GetCreatorEmailOk returns a tuple with the CreatorEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorEmail

`func (o *SettingDataV2Model) SetCreatorEmail(v string)`

SetCreatorEmail sets CreatorEmail field to given value.


### SetCreatorEmailNil

`func (o *SettingDataV2Model) SetCreatorEmailNil(b bool)`

 SetCreatorEmailNil sets the value for CreatorEmail to be an explicit nil

### UnsetCreatorEmail
`func (o *SettingDataV2Model) UnsetCreatorEmail()`

UnsetCreatorEmail ensures that no value is present for CreatorEmail, not even an explicit nil
### GetCreatorFullName

`func (o *SettingDataV2Model) GetCreatorFullName() string`

GetCreatorFullName returns the CreatorFullName field if non-nil, zero value otherwise.

### GetCreatorFullNameOk

`func (o *SettingDataV2Model) GetCreatorFullNameOk() (*string, bool)`

GetCreatorFullNameOk returns a tuple with the CreatorFullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorFullName

`func (o *SettingDataV2Model) SetCreatorFullName(v string)`

SetCreatorFullName sets CreatorFullName field to given value.


### SetCreatorFullNameNil

`func (o *SettingDataV2Model) SetCreatorFullNameNil(b bool)`

 SetCreatorFullNameNil sets the value for CreatorFullName to be an explicit nil

### UnsetCreatorFullName
`func (o *SettingDataV2Model) UnsetCreatorFullName()`

UnsetCreatorFullName ensures that no value is present for CreatorFullName, not even an explicit nil
### GetPredefinedVariations

`func (o *SettingDataV2Model) GetPredefinedVariations() []PredefinedVariationModel`

GetPredefinedVariations returns the PredefinedVariations field if non-nil, zero value otherwise.

### GetPredefinedVariationsOk

`func (o *SettingDataV2Model) GetPredefinedVariationsOk() (*[]PredefinedVariationModel, bool)`

GetPredefinedVariationsOk returns a tuple with the PredefinedVariations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredefinedVariations

`func (o *SettingDataV2Model) SetPredefinedVariations(v []PredefinedVariationModel)`

SetPredefinedVariations sets PredefinedVariations field to given value.


### GetIsWatching

`func (o *SettingDataV2Model) GetIsWatching() bool`

GetIsWatching returns the IsWatching field if non-nil, zero value otherwise.

### GetIsWatchingOk

`func (o *SettingDataV2Model) GetIsWatchingOk() (*bool, bool)`

GetIsWatchingOk returns a tuple with the IsWatching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWatching

`func (o *SettingDataV2Model) SetIsWatching(v bool)`

SetIsWatching sets IsWatching field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


