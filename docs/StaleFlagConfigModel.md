# StaleFlagConfigModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigId** | **string** | Identifier of the Config. | 
**Name** | **string** | Name of the Config. | 
**EvaluationVersion** | [**EvaluationVersion**](EvaluationVersion.md) |  | 
**HasCodeReferences** | **bool** | Config has code references uploaded. | 
**Settings** | [**[]StaleFlagSettingModel**](StaleFlagSettingModel.md) | Stale feature flags. | 

## Methods

### NewStaleFlagConfigModel

`func NewStaleFlagConfigModel(configId string, name string, evaluationVersion EvaluationVersion, hasCodeReferences bool, settings []StaleFlagSettingModel, ) *StaleFlagConfigModel`

NewStaleFlagConfigModel instantiates a new StaleFlagConfigModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStaleFlagConfigModelWithDefaults

`func NewStaleFlagConfigModelWithDefaults() *StaleFlagConfigModel`

NewStaleFlagConfigModelWithDefaults instantiates a new StaleFlagConfigModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigId

`func (o *StaleFlagConfigModel) GetConfigId() string`

GetConfigId returns the ConfigId field if non-nil, zero value otherwise.

### GetConfigIdOk

`func (o *StaleFlagConfigModel) GetConfigIdOk() (*string, bool)`

GetConfigIdOk returns a tuple with the ConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigId

`func (o *StaleFlagConfigModel) SetConfigId(v string)`

SetConfigId sets ConfigId field to given value.


### GetName

`func (o *StaleFlagConfigModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StaleFlagConfigModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StaleFlagConfigModel) SetName(v string)`

SetName sets Name field to given value.


### GetEvaluationVersion

`func (o *StaleFlagConfigModel) GetEvaluationVersion() EvaluationVersion`

GetEvaluationVersion returns the EvaluationVersion field if non-nil, zero value otherwise.

### GetEvaluationVersionOk

`func (o *StaleFlagConfigModel) GetEvaluationVersionOk() (*EvaluationVersion, bool)`

GetEvaluationVersionOk returns a tuple with the EvaluationVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationVersion

`func (o *StaleFlagConfigModel) SetEvaluationVersion(v EvaluationVersion)`

SetEvaluationVersion sets EvaluationVersion field to given value.


### GetHasCodeReferences

`func (o *StaleFlagConfigModel) GetHasCodeReferences() bool`

GetHasCodeReferences returns the HasCodeReferences field if non-nil, zero value otherwise.

### GetHasCodeReferencesOk

`func (o *StaleFlagConfigModel) GetHasCodeReferencesOk() (*bool, bool)`

GetHasCodeReferencesOk returns a tuple with the HasCodeReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasCodeReferences

`func (o *StaleFlagConfigModel) SetHasCodeReferences(v bool)`

SetHasCodeReferences sets HasCodeReferences field to given value.


### GetSettings

`func (o *StaleFlagConfigModel) GetSettings() []StaleFlagSettingModel`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *StaleFlagConfigModel) GetSettingsOk() (*[]StaleFlagSettingModel, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *StaleFlagConfigModel) SetSettings(v []StaleFlagSettingModel)`

SetSettings sets Settings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


