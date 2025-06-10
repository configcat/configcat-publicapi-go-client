# ConfigSettingFormulasModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | [**ConfigModel**](ConfigModel.md) |  | 
**Environment** | [**EnvironmentModel**](EnvironmentModel.md) |  | 
**ReadOnly** | **bool** |  | 
**SettingFormulas** | [**[]ConfigSettingFormulaModel**](ConfigSettingFormulaModel.md) | Evaluation descriptors of each updated Feature Flag and Setting. | 
**FeatureFlagLimitations** | [**FeatureFlagLimitations**](FeatureFlagLimitations.md) |  | 

## Methods

### NewConfigSettingFormulasModel

`func NewConfigSettingFormulasModel(config ConfigModel, environment EnvironmentModel, readOnly bool, settingFormulas []ConfigSettingFormulaModel, featureFlagLimitations FeatureFlagLimitations, ) *ConfigSettingFormulasModel`

NewConfigSettingFormulasModel instantiates a new ConfigSettingFormulasModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigSettingFormulasModelWithDefaults

`func NewConfigSettingFormulasModelWithDefaults() *ConfigSettingFormulasModel`

NewConfigSettingFormulasModelWithDefaults instantiates a new ConfigSettingFormulasModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *ConfigSettingFormulasModel) GetConfig() ConfigModel`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ConfigSettingFormulasModel) GetConfigOk() (*ConfigModel, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ConfigSettingFormulasModel) SetConfig(v ConfigModel)`

SetConfig sets Config field to given value.


### GetEnvironment

`func (o *ConfigSettingFormulasModel) GetEnvironment() EnvironmentModel`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ConfigSettingFormulasModel) GetEnvironmentOk() (*EnvironmentModel, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ConfigSettingFormulasModel) SetEnvironment(v EnvironmentModel)`

SetEnvironment sets Environment field to given value.


### GetReadOnly

`func (o *ConfigSettingFormulasModel) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *ConfigSettingFormulasModel) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *ConfigSettingFormulasModel) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.


### GetSettingFormulas

`func (o *ConfigSettingFormulasModel) GetSettingFormulas() []ConfigSettingFormulaModel`

GetSettingFormulas returns the SettingFormulas field if non-nil, zero value otherwise.

### GetSettingFormulasOk

`func (o *ConfigSettingFormulasModel) GetSettingFormulasOk() (*[]ConfigSettingFormulaModel, bool)`

GetSettingFormulasOk returns a tuple with the SettingFormulas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingFormulas

`func (o *ConfigSettingFormulasModel) SetSettingFormulas(v []ConfigSettingFormulaModel)`

SetSettingFormulas sets SettingFormulas field to given value.


### GetFeatureFlagLimitations

`func (o *ConfigSettingFormulasModel) GetFeatureFlagLimitations() FeatureFlagLimitations`

GetFeatureFlagLimitations returns the FeatureFlagLimitations field if non-nil, zero value otherwise.

### GetFeatureFlagLimitationsOk

`func (o *ConfigSettingFormulasModel) GetFeatureFlagLimitationsOk() (*FeatureFlagLimitations, bool)`

GetFeatureFlagLimitationsOk returns a tuple with the FeatureFlagLimitations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureFlagLimitations

`func (o *ConfigSettingFormulasModel) SetFeatureFlagLimitations(v FeatureFlagLimitations)`

SetFeatureFlagLimitations sets FeatureFlagLimitations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


