# StaleFlagSettingModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SettingId** | **int32** | Identifier of the Feature Flag or Setting. | 
**Name** | **string** | Name of the Feature Flag or Setting. | 
**Key** | **string** | Key of the Feature Flag or Setting. | 
**Hint** | **NullableString** | Description of the Feature Flag or Setting. | 
**HasCodeReferences** | **bool** | Feature Flag or Setting has code references uploaded. | 
**Tags** | [**[]StaleFlagSettingTagModel**](StaleFlagSettingTagModel.md) | The tags&#39; identifiers attached to the Feature Flag or Setting. | 
**SettingValues** | [**[]StaleFlagSettingValueModel**](StaleFlagSettingValueModel.md) | Environment level feature flag stale data. | 

## Methods

### NewStaleFlagSettingModel

`func NewStaleFlagSettingModel(settingId int32, name string, key string, hint NullableString, hasCodeReferences bool, tags []StaleFlagSettingTagModel, settingValues []StaleFlagSettingValueModel, ) *StaleFlagSettingModel`

NewStaleFlagSettingModel instantiates a new StaleFlagSettingModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStaleFlagSettingModelWithDefaults

`func NewStaleFlagSettingModelWithDefaults() *StaleFlagSettingModel`

NewStaleFlagSettingModelWithDefaults instantiates a new StaleFlagSettingModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSettingId

`func (o *StaleFlagSettingModel) GetSettingId() int32`

GetSettingId returns the SettingId field if non-nil, zero value otherwise.

### GetSettingIdOk

`func (o *StaleFlagSettingModel) GetSettingIdOk() (*int32, bool)`

GetSettingIdOk returns a tuple with the SettingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingId

`func (o *StaleFlagSettingModel) SetSettingId(v int32)`

SetSettingId sets SettingId field to given value.


### GetName

`func (o *StaleFlagSettingModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StaleFlagSettingModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StaleFlagSettingModel) SetName(v string)`

SetName sets Name field to given value.


### GetKey

`func (o *StaleFlagSettingModel) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *StaleFlagSettingModel) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *StaleFlagSettingModel) SetKey(v string)`

SetKey sets Key field to given value.


### GetHint

`func (o *StaleFlagSettingModel) GetHint() string`

GetHint returns the Hint field if non-nil, zero value otherwise.

### GetHintOk

`func (o *StaleFlagSettingModel) GetHintOk() (*string, bool)`

GetHintOk returns a tuple with the Hint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHint

`func (o *StaleFlagSettingModel) SetHint(v string)`

SetHint sets Hint field to given value.


### SetHintNil

`func (o *StaleFlagSettingModel) SetHintNil(b bool)`

 SetHintNil sets the value for Hint to be an explicit nil

### UnsetHint
`func (o *StaleFlagSettingModel) UnsetHint()`

UnsetHint ensures that no value is present for Hint, not even an explicit nil
### GetHasCodeReferences

`func (o *StaleFlagSettingModel) GetHasCodeReferences() bool`

GetHasCodeReferences returns the HasCodeReferences field if non-nil, zero value otherwise.

### GetHasCodeReferencesOk

`func (o *StaleFlagSettingModel) GetHasCodeReferencesOk() (*bool, bool)`

GetHasCodeReferencesOk returns a tuple with the HasCodeReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasCodeReferences

`func (o *StaleFlagSettingModel) SetHasCodeReferences(v bool)`

SetHasCodeReferences sets HasCodeReferences field to given value.


### GetTags

`func (o *StaleFlagSettingModel) GetTags() []StaleFlagSettingTagModel`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *StaleFlagSettingModel) GetTagsOk() (*[]StaleFlagSettingTagModel, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *StaleFlagSettingModel) SetTags(v []StaleFlagSettingTagModel)`

SetTags sets Tags field to given value.


### GetSettingValues

`func (o *StaleFlagSettingModel) GetSettingValues() []StaleFlagSettingValueModel`

GetSettingValues returns the SettingValues field if non-nil, zero value otherwise.

### GetSettingValuesOk

`func (o *StaleFlagSettingModel) GetSettingValuesOk() (*[]StaleFlagSettingValueModel, bool)`

GetSettingValuesOk returns a tuple with the SettingValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingValues

`func (o *StaleFlagSettingModel) SetSettingValues(v []StaleFlagSettingValueModel)`

SetSettingValues sets SettingValues field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


