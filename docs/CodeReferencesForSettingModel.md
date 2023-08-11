# CodeReferencesForSettingModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Setting** | Pointer to [**SettingModel**](SettingModel.md) |  | [optional] 
**CodeReferences** | Pointer to [**[]CodeReferenceModel**](CodeReferenceModel.md) | List of Code references that belongs to the Feature Flag or Setting. | [optional] 

## Methods

### NewCodeReferencesForSettingModel

`func NewCodeReferencesForSettingModel() *CodeReferencesForSettingModel`

NewCodeReferencesForSettingModel instantiates a new CodeReferencesForSettingModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCodeReferencesForSettingModelWithDefaults

`func NewCodeReferencesForSettingModelWithDefaults() *CodeReferencesForSettingModel`

NewCodeReferencesForSettingModelWithDefaults instantiates a new CodeReferencesForSettingModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSetting

`func (o *CodeReferencesForSettingModel) GetSetting() SettingModel`

GetSetting returns the Setting field if non-nil, zero value otherwise.

### GetSettingOk

`func (o *CodeReferencesForSettingModel) GetSettingOk() (*SettingModel, bool)`

GetSettingOk returns a tuple with the Setting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetting

`func (o *CodeReferencesForSettingModel) SetSetting(v SettingModel)`

SetSetting sets Setting field to given value.

### HasSetting

`func (o *CodeReferencesForSettingModel) HasSetting() bool`

HasSetting returns a boolean if a field has been set.

### GetCodeReferences

`func (o *CodeReferencesForSettingModel) GetCodeReferences() []CodeReferenceModel`

GetCodeReferences returns the CodeReferences field if non-nil, zero value otherwise.

### GetCodeReferencesOk

`func (o *CodeReferencesForSettingModel) GetCodeReferencesOk() (*[]CodeReferenceModel, bool)`

GetCodeReferencesOk returns a tuple with the CodeReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeReferences

`func (o *CodeReferencesForSettingModel) SetCodeReferences(v []CodeReferenceModel)`

SetCodeReferences sets CodeReferences field to given value.

### HasCodeReferences

`func (o *CodeReferencesForSettingModel) HasCodeReferences() bool`

HasCodeReferences returns a boolean if a field has been set.

### SetCodeReferencesNil

`func (o *CodeReferencesForSettingModel) SetCodeReferencesNil(b bool)`

 SetCodeReferencesNil sets the value for CodeReferences to be an explicit nil

### UnsetCodeReferences
`func (o *CodeReferencesForSettingModel) UnsetCodeReferences()`

UnsetCodeReferences ensures that no value is present for CodeReferences, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


