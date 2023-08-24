# ConfigSettingValuesModelHaljson

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Embedded** | Pointer to [**ConfigSettingFormulasModelHaljsonEmbedded**](ConfigSettingFormulasModelHaljsonEmbedded.md) |  | [optional] 
**ReadOnly** | Pointer to **bool** |  | [optional] 
**SettingValues** | Pointer to [**[]ConfigSettingValueModel**](ConfigSettingValueModel.md) |  | [optional] 
**FeatureFlagLimitations** | Pointer to [**FeatureFlagLimitations**](FeatureFlagLimitations.md) |  | [optional] 
**Links** | Pointer to [**ConfigSettingFormulasModelHaljsonEmbeddedEnvironmentLinks**](ConfigSettingFormulasModelHaljsonEmbeddedEnvironmentLinks.md) |  | [optional] 

## Methods

### NewConfigSettingValuesModelHaljson

`func NewConfigSettingValuesModelHaljson() *ConfigSettingValuesModelHaljson`

NewConfigSettingValuesModelHaljson instantiates a new ConfigSettingValuesModelHaljson object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigSettingValuesModelHaljsonWithDefaults

`func NewConfigSettingValuesModelHaljsonWithDefaults() *ConfigSettingValuesModelHaljson`

NewConfigSettingValuesModelHaljsonWithDefaults instantiates a new ConfigSettingValuesModelHaljson object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmbedded

`func (o *ConfigSettingValuesModelHaljson) GetEmbedded() ConfigSettingFormulasModelHaljsonEmbedded`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *ConfigSettingValuesModelHaljson) GetEmbeddedOk() (*ConfigSettingFormulasModelHaljsonEmbedded, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *ConfigSettingValuesModelHaljson) SetEmbedded(v ConfigSettingFormulasModelHaljsonEmbedded)`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *ConfigSettingValuesModelHaljson) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetReadOnly

`func (o *ConfigSettingValuesModelHaljson) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *ConfigSettingValuesModelHaljson) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *ConfigSettingValuesModelHaljson) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *ConfigSettingValuesModelHaljson) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetSettingValues

`func (o *ConfigSettingValuesModelHaljson) GetSettingValues() []ConfigSettingValueModel`

GetSettingValues returns the SettingValues field if non-nil, zero value otherwise.

### GetSettingValuesOk

`func (o *ConfigSettingValuesModelHaljson) GetSettingValuesOk() (*[]ConfigSettingValueModel, bool)`

GetSettingValuesOk returns a tuple with the SettingValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingValues

`func (o *ConfigSettingValuesModelHaljson) SetSettingValues(v []ConfigSettingValueModel)`

SetSettingValues sets SettingValues field to given value.

### HasSettingValues

`func (o *ConfigSettingValuesModelHaljson) HasSettingValues() bool`

HasSettingValues returns a boolean if a field has been set.

### SetSettingValuesNil

`func (o *ConfigSettingValuesModelHaljson) SetSettingValuesNil(b bool)`

 SetSettingValuesNil sets the value for SettingValues to be an explicit nil

### UnsetSettingValues
`func (o *ConfigSettingValuesModelHaljson) UnsetSettingValues()`

UnsetSettingValues ensures that no value is present for SettingValues, not even an explicit nil
### GetFeatureFlagLimitations

`func (o *ConfigSettingValuesModelHaljson) GetFeatureFlagLimitations() FeatureFlagLimitations`

GetFeatureFlagLimitations returns the FeatureFlagLimitations field if non-nil, zero value otherwise.

### GetFeatureFlagLimitationsOk

`func (o *ConfigSettingValuesModelHaljson) GetFeatureFlagLimitationsOk() (*FeatureFlagLimitations, bool)`

GetFeatureFlagLimitationsOk returns a tuple with the FeatureFlagLimitations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureFlagLimitations

`func (o *ConfigSettingValuesModelHaljson) SetFeatureFlagLimitations(v FeatureFlagLimitations)`

SetFeatureFlagLimitations sets FeatureFlagLimitations field to given value.

### HasFeatureFlagLimitations

`func (o *ConfigSettingValuesModelHaljson) HasFeatureFlagLimitations() bool`

HasFeatureFlagLimitations returns a boolean if a field has been set.

### GetLinks

`func (o *ConfigSettingValuesModelHaljson) GetLinks() ConfigSettingFormulasModelHaljsonEmbeddedEnvironmentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ConfigSettingValuesModelHaljson) GetLinksOk() (*ConfigSettingFormulasModelHaljsonEmbeddedEnvironmentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ConfigSettingValuesModelHaljson) SetLinks(v ConfigSettingFormulasModelHaljsonEmbeddedEnvironmentLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ConfigSettingValuesModelHaljson) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


