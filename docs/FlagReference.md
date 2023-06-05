# FlagReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SettingId** | **int32** | The identifier of the Feature Flag or Setting the code reference belongs to. | 
**References** | [**[]ReferenceLines**](ReferenceLines.md) | The actual references to the given Feature Flag or Setting. | 

## Methods

### NewFlagReference

`func NewFlagReference(settingId int32, references []ReferenceLines, ) *FlagReference`

NewFlagReference instantiates a new FlagReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlagReferenceWithDefaults

`func NewFlagReferenceWithDefaults() *FlagReference`

NewFlagReferenceWithDefaults instantiates a new FlagReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSettingId

`func (o *FlagReference) GetSettingId() int32`

GetSettingId returns the SettingId field if non-nil, zero value otherwise.

### GetSettingIdOk

`func (o *FlagReference) GetSettingIdOk() (*int32, bool)`

GetSettingIdOk returns a tuple with the SettingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingId

`func (o *FlagReference) SetSettingId(v int32)`

SetSettingId sets SettingId field to given value.


### GetReferences

`func (o *FlagReference) GetReferences() []ReferenceLines`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *FlagReference) GetReferencesOk() (*[]ReferenceLines, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *FlagReference) SetReferences(v []ReferenceLines)`

SetReferences sets References field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


