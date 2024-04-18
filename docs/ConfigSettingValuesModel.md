# ConfigSettingValuesModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to [**ConfigModel**](ConfigModel.md) |  | [optional] 
**Environment** | Pointer to [**EnvironmentModel**](EnvironmentModel.md) |  | [optional] 
**ReadOnly** | Pointer to **bool** |  | [optional] 
**SettingValues** | Pointer to [**[]ConfigSettingValueModel**](ConfigSettingValueModel.md) |  | [optional] 
**FeatureFlagLimitations** | Pointer to [**FeatureFlagLimitations**](FeatureFlagLimitations.md) |  | [optional] 
**Links** | Pointer to [**ConfigSettingFormulasModelLinks**](ConfigSettingFormulasModelLinks.md) |  | [optional] 

## Methods

### NewConfigSettingValuesModel

`func NewConfigSettingValuesModel() *ConfigSettingValuesModel`

NewConfigSettingValuesModel instantiates a new ConfigSettingValuesModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigSettingValuesModelWithDefaults

`func NewConfigSettingValuesModelWithDefaults() *ConfigSettingValuesModel`

NewConfigSettingValuesModelWithDefaults instantiates a new ConfigSettingValuesModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *ConfigSettingValuesModel) GetConfig() ConfigModel`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ConfigSettingValuesModel) GetConfigOk() (*ConfigModel, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ConfigSettingValuesModel) SetConfig(v ConfigModel)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ConfigSettingValuesModel) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetEnvironment

`func (o *ConfigSettingValuesModel) GetEnvironment() EnvironmentModel`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ConfigSettingValuesModel) GetEnvironmentOk() (*EnvironmentModel, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ConfigSettingValuesModel) SetEnvironment(v EnvironmentModel)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *ConfigSettingValuesModel) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetReadOnly

`func (o *ConfigSettingValuesModel) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *ConfigSettingValuesModel) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *ConfigSettingValuesModel) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *ConfigSettingValuesModel) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetSettingValues

`func (o *ConfigSettingValuesModel) GetSettingValues() []ConfigSettingValueModel`

GetSettingValues returns the SettingValues field if non-nil, zero value otherwise.

### GetSettingValuesOk

`func (o *ConfigSettingValuesModel) GetSettingValuesOk() (*[]ConfigSettingValueModel, bool)`

GetSettingValuesOk returns a tuple with the SettingValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingValues

`func (o *ConfigSettingValuesModel) SetSettingValues(v []ConfigSettingValueModel)`

SetSettingValues sets SettingValues field to given value.

### HasSettingValues

`func (o *ConfigSettingValuesModel) HasSettingValues() bool`

HasSettingValues returns a boolean if a field has been set.

### SetSettingValuesNil

`func (o *ConfigSettingValuesModel) SetSettingValuesNil(b bool)`

 SetSettingValuesNil sets the value for SettingValues to be an explicit nil

### UnsetSettingValues
`func (o *ConfigSettingValuesModel) UnsetSettingValues()`

UnsetSettingValues ensures that no value is present for SettingValues, not even an explicit nil
### GetFeatureFlagLimitations

`func (o *ConfigSettingValuesModel) GetFeatureFlagLimitations() FeatureFlagLimitations`

GetFeatureFlagLimitations returns the FeatureFlagLimitations field if non-nil, zero value otherwise.

### GetFeatureFlagLimitationsOk

`func (o *ConfigSettingValuesModel) GetFeatureFlagLimitationsOk() (*FeatureFlagLimitations, bool)`

GetFeatureFlagLimitationsOk returns a tuple with the FeatureFlagLimitations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureFlagLimitations

`func (o *ConfigSettingValuesModel) SetFeatureFlagLimitations(v FeatureFlagLimitations)`

SetFeatureFlagLimitations sets FeatureFlagLimitations field to given value.

### HasFeatureFlagLimitations

`func (o *ConfigSettingValuesModel) HasFeatureFlagLimitations() bool`

HasFeatureFlagLimitations returns a boolean if a field has been set.

### GetLinks

`func (o *ConfigSettingValuesModel) GetLinks() ConfigSettingFormulasModelLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ConfigSettingValuesModel) GetLinksOk() (*ConfigSettingFormulasModelLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ConfigSettingValuesModel) SetLinks(v ConfigSettingFormulasModelLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ConfigSettingValuesModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


