# DeletedSettingModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | Key of the Feature Flag or Setting. | [readonly] 
**Name** | **string** | Name of the Feature Flag or Setting. | [readonly] 
**Hint** | **NullableString** | Description of the Feature Flag or Setting. | [readonly] 
**SettingType** | [**SettingType**](SettingType.md) |  | 

## Methods

### NewDeletedSettingModel

`func NewDeletedSettingModel(key string, name string, hint NullableString, settingType SettingType, ) *DeletedSettingModel`

NewDeletedSettingModel instantiates a new DeletedSettingModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeletedSettingModelWithDefaults

`func NewDeletedSettingModelWithDefaults() *DeletedSettingModel`

NewDeletedSettingModelWithDefaults instantiates a new DeletedSettingModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *DeletedSettingModel) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *DeletedSettingModel) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *DeletedSettingModel) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *DeletedSettingModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeletedSettingModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeletedSettingModel) SetName(v string)`

SetName sets Name field to given value.


### GetHint

`func (o *DeletedSettingModel) GetHint() string`

GetHint returns the Hint field if non-nil, zero value otherwise.

### GetHintOk

`func (o *DeletedSettingModel) GetHintOk() (*string, bool)`

GetHintOk returns a tuple with the Hint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHint

`func (o *DeletedSettingModel) SetHint(v string)`

SetHint sets Hint field to given value.


### SetHintNil

`func (o *DeletedSettingModel) SetHintNil(b bool)`

 SetHintNil sets the value for Hint to be an explicit nil

### UnsetHint
`func (o *DeletedSettingModel) UnsetHint()`

UnsetHint ensures that no value is present for Hint, not even an explicit nil
### GetSettingType

`func (o *DeletedSettingModel) GetSettingType() SettingType`

GetSettingType returns the SettingType field if non-nil, zero value otherwise.

### GetSettingTypeOk

`func (o *DeletedSettingModel) GetSettingTypeOk() (*SettingType, bool)`

GetSettingTypeOk returns a tuple with the SettingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingType

`func (o *DeletedSettingModel) SetSettingType(v SettingType)`

SetSettingType sets SettingType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


