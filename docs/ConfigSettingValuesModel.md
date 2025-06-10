# ConfigSettingValuesModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | [**ConfigModel**](ConfigModel.md) |  | 
**Environment** | [**EnvironmentModel**](EnvironmentModel.md) |  | 
**ReadOnly** | **bool** |  | 
**SettingValues** | [**[]ConfigSettingValueModel**](ConfigSettingValueModel.md) |  | 
**FeatureFlagLimitations** | [**FeatureFlagLimitations**](FeatureFlagLimitations.md) |  | 

## Methods

### NewConfigSettingValuesModel

`func NewConfigSettingValuesModel(config ConfigModel, environment EnvironmentModel, readOnly bool, settingValues []ConfigSettingValueModel, featureFlagLimitations FeatureFlagLimitations, ) *ConfigSettingValuesModel`

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


