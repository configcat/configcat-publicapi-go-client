# ConfigSettingFormulasModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to [**ConfigModel**](ConfigModel.md) |  | [optional] 
**Environment** | Pointer to [**EnvironmentModel**](EnvironmentModel.md) |  | [optional] 
**ReadOnly** | Pointer to **bool** |  | [optional] 
**SettingFormulas** | Pointer to [**[]ConfigSettingFormulaModel**](ConfigSettingFormulaModel.md) | Evaluation descriptors of each updated Feature Flag and Setting. | [optional] 
**FeatureFlagLimitations** | Pointer to [**FeatureFlagLimitations**](FeatureFlagLimitations.md) |  | [optional] 
**Links** | Pointer to [**ConfigSettingFormulasModelLinks**](ConfigSettingFormulasModelLinks.md) |  | [optional] 

## Methods

### NewConfigSettingFormulasModel

`func NewConfigSettingFormulasModel() *ConfigSettingFormulasModel`

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

### HasConfig

`func (o *ConfigSettingFormulasModel) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

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

### HasEnvironment

`func (o *ConfigSettingFormulasModel) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

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

### HasReadOnly

`func (o *ConfigSettingFormulasModel) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

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

### HasSettingFormulas

`func (o *ConfigSettingFormulasModel) HasSettingFormulas() bool`

HasSettingFormulas returns a boolean if a field has been set.

### SetSettingFormulasNil

`func (o *ConfigSettingFormulasModel) SetSettingFormulasNil(b bool)`

 SetSettingFormulasNil sets the value for SettingFormulas to be an explicit nil

### UnsetSettingFormulas
`func (o *ConfigSettingFormulasModel) UnsetSettingFormulas()`

UnsetSettingFormulas ensures that no value is present for SettingFormulas, not even an explicit nil
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

### HasFeatureFlagLimitations

`func (o *ConfigSettingFormulasModel) HasFeatureFlagLimitations() bool`

HasFeatureFlagLimitations returns a boolean if a field has been set.

### GetLinks

`func (o *ConfigSettingFormulasModel) GetLinks() ConfigSettingFormulasModelLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ConfigSettingFormulasModel) GetLinksOk() (*ConfigSettingFormulasModelLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ConfigSettingFormulasModel) SetLinks(v ConfigSettingFormulasModelLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ConfigSettingFormulasModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


