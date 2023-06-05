# SettingValueModelHaljsonEmbeddedSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SettingId** | Pointer to **int32** |  | [optional] 
**Key** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Hint** | Pointer to **NullableString** |  | [optional] 
**SettingType** | Pointer to [**SettingType**](SettingType.md) |  | [optional] 
**Order** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **NullableTime** |  | [optional] 
**CreatorEmail** | Pointer to **NullableString** |  | [optional] 
**CreatorFullName** | Pointer to **NullableString** |  | [optional] 
**IsWatching** | Pointer to **bool** |  | [optional] 

## Methods

### NewSettingValueModelHaljsonEmbeddedSetting

`func NewSettingValueModelHaljsonEmbeddedSetting() *SettingValueModelHaljsonEmbeddedSetting`

NewSettingValueModelHaljsonEmbeddedSetting instantiates a new SettingValueModelHaljsonEmbeddedSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettingValueModelHaljsonEmbeddedSettingWithDefaults

`func NewSettingValueModelHaljsonEmbeddedSettingWithDefaults() *SettingValueModelHaljsonEmbeddedSetting`

NewSettingValueModelHaljsonEmbeddedSettingWithDefaults instantiates a new SettingValueModelHaljsonEmbeddedSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSettingId

`func (o *SettingValueModelHaljsonEmbeddedSetting) GetSettingId() int32`

GetSettingId returns the SettingId field if non-nil, zero value otherwise.

### GetSettingIdOk

`func (o *SettingValueModelHaljsonEmbeddedSetting) GetSettingIdOk() (*int32, bool)`

GetSettingIdOk returns a tuple with the SettingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingId

`func (o *SettingValueModelHaljsonEmbeddedSetting) SetSettingId(v int32)`

SetSettingId sets SettingId field to given value.

### HasSettingId

`func (o *SettingValueModelHaljsonEmbeddedSetting) HasSettingId() bool`

HasSettingId returns a boolean if a field has been set.

### GetKey

`func (o *SettingValueModelHaljsonEmbeddedSetting) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SettingValueModelHaljsonEmbeddedSetting) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SettingValueModelHaljsonEmbeddedSetting) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *SettingValueModelHaljsonEmbeddedSetting) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *SettingValueModelHaljsonEmbeddedSetting) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *SettingValueModelHaljsonEmbeddedSetting) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetName

`func (o *SettingValueModelHaljsonEmbeddedSetting) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SettingValueModelHaljsonEmbeddedSetting) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SettingValueModelHaljsonEmbeddedSetting) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SettingValueModelHaljsonEmbeddedSetting) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *SettingValueModelHaljsonEmbeddedSetting) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *SettingValueModelHaljsonEmbeddedSetting) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetHint

`func (o *SettingValueModelHaljsonEmbeddedSetting) GetHint() string`

GetHint returns the Hint field if non-nil, zero value otherwise.

### GetHintOk

`func (o *SettingValueModelHaljsonEmbeddedSetting) GetHintOk() (*string, bool)`

GetHintOk returns a tuple with the Hint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHint

`func (o *SettingValueModelHaljsonEmbeddedSetting) SetHint(v string)`

SetHint sets Hint field to given value.

### HasHint

`func (o *SettingValueModelHaljsonEmbeddedSetting) HasHint() bool`

HasHint returns a boolean if a field has been set.

### SetHintNil

`func (o *SettingValueModelHaljsonEmbeddedSetting) SetHintNil(b bool)`

 SetHintNil sets the value for Hint to be an explicit nil

### UnsetHint
`func (o *SettingValueModelHaljsonEmbeddedSetting) UnsetHint()`

UnsetHint ensures that no value is present for Hint, not even an explicit nil
### GetSettingType

`func (o *SettingValueModelHaljsonEmbeddedSetting) GetSettingType() SettingType`

GetSettingType returns the SettingType field if non-nil, zero value otherwise.

### GetSettingTypeOk

`func (o *SettingValueModelHaljsonEmbeddedSetting) GetSettingTypeOk() (*SettingType, bool)`

GetSettingTypeOk returns a tuple with the SettingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingType

`func (o *SettingValueModelHaljsonEmbeddedSetting) SetSettingType(v SettingType)`

SetSettingType sets SettingType field to given value.

### HasSettingType

`func (o *SettingValueModelHaljsonEmbeddedSetting) HasSettingType() bool`

HasSettingType returns a boolean if a field has been set.

### GetOrder

`func (o *SettingValueModelHaljsonEmbeddedSetting) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *SettingValueModelHaljsonEmbeddedSetting) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *SettingValueModelHaljsonEmbeddedSetting) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *SettingValueModelHaljsonEmbeddedSetting) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SettingValueModelHaljsonEmbeddedSetting) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SettingValueModelHaljsonEmbeddedSetting) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SettingValueModelHaljsonEmbeddedSetting) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SettingValueModelHaljsonEmbeddedSetting) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *SettingValueModelHaljsonEmbeddedSetting) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *SettingValueModelHaljsonEmbeddedSetting) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetCreatorEmail

`func (o *SettingValueModelHaljsonEmbeddedSetting) GetCreatorEmail() string`

GetCreatorEmail returns the CreatorEmail field if non-nil, zero value otherwise.

### GetCreatorEmailOk

`func (o *SettingValueModelHaljsonEmbeddedSetting) GetCreatorEmailOk() (*string, bool)`

GetCreatorEmailOk returns a tuple with the CreatorEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorEmail

`func (o *SettingValueModelHaljsonEmbeddedSetting) SetCreatorEmail(v string)`

SetCreatorEmail sets CreatorEmail field to given value.

### HasCreatorEmail

`func (o *SettingValueModelHaljsonEmbeddedSetting) HasCreatorEmail() bool`

HasCreatorEmail returns a boolean if a field has been set.

### SetCreatorEmailNil

`func (o *SettingValueModelHaljsonEmbeddedSetting) SetCreatorEmailNil(b bool)`

 SetCreatorEmailNil sets the value for CreatorEmail to be an explicit nil

### UnsetCreatorEmail
`func (o *SettingValueModelHaljsonEmbeddedSetting) UnsetCreatorEmail()`

UnsetCreatorEmail ensures that no value is present for CreatorEmail, not even an explicit nil
### GetCreatorFullName

`func (o *SettingValueModelHaljsonEmbeddedSetting) GetCreatorFullName() string`

GetCreatorFullName returns the CreatorFullName field if non-nil, zero value otherwise.

### GetCreatorFullNameOk

`func (o *SettingValueModelHaljsonEmbeddedSetting) GetCreatorFullNameOk() (*string, bool)`

GetCreatorFullNameOk returns a tuple with the CreatorFullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorFullName

`func (o *SettingValueModelHaljsonEmbeddedSetting) SetCreatorFullName(v string)`

SetCreatorFullName sets CreatorFullName field to given value.

### HasCreatorFullName

`func (o *SettingValueModelHaljsonEmbeddedSetting) HasCreatorFullName() bool`

HasCreatorFullName returns a boolean if a field has been set.

### SetCreatorFullNameNil

`func (o *SettingValueModelHaljsonEmbeddedSetting) SetCreatorFullNameNil(b bool)`

 SetCreatorFullNameNil sets the value for CreatorFullName to be an explicit nil

### UnsetCreatorFullName
`func (o *SettingValueModelHaljsonEmbeddedSetting) UnsetCreatorFullName()`

UnsetCreatorFullName ensures that no value is present for CreatorFullName, not even an explicit nil
### GetIsWatching

`func (o *SettingValueModelHaljsonEmbeddedSetting) GetIsWatching() bool`

GetIsWatching returns the IsWatching field if non-nil, zero value otherwise.

### GetIsWatchingOk

`func (o *SettingValueModelHaljsonEmbeddedSetting) GetIsWatchingOk() (*bool, bool)`

GetIsWatchingOk returns a tuple with the IsWatching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWatching

`func (o *SettingValueModelHaljsonEmbeddedSetting) SetIsWatching(v bool)`

SetIsWatching sets IsWatching field to given value.

### HasIsWatching

`func (o *SettingValueModelHaljsonEmbeddedSetting) HasIsWatching() bool`

HasIsWatching returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


