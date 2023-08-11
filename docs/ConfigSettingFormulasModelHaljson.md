# ConfigSettingFormulasModelHaljson

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Embedded** | Pointer to [**ConfigSettingFormulasModelHaljsonEmbedded**](ConfigSettingFormulasModelHaljsonEmbedded.md) |  | [optional] 
**ReadOnly** | Pointer to **bool** |  | [optional] 
**SettingFormulas** | Pointer to [**[]ConfigSettingFormulaModel**](ConfigSettingFormulaModel.md) | Evaluation descriptors of each updated Feature Flag and Setting. | [optional] 
**Links** | Pointer to [**ConfigSettingFormulasModelHaljsonEmbeddedEnvironmentLinks**](ConfigSettingFormulasModelHaljsonEmbeddedEnvironmentLinks.md) |  | [optional] 

## Methods

### NewConfigSettingFormulasModelHaljson

`func NewConfigSettingFormulasModelHaljson() *ConfigSettingFormulasModelHaljson`

NewConfigSettingFormulasModelHaljson instantiates a new ConfigSettingFormulasModelHaljson object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigSettingFormulasModelHaljsonWithDefaults

`func NewConfigSettingFormulasModelHaljsonWithDefaults() *ConfigSettingFormulasModelHaljson`

NewConfigSettingFormulasModelHaljsonWithDefaults instantiates a new ConfigSettingFormulasModelHaljson object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmbedded

`func (o *ConfigSettingFormulasModelHaljson) GetEmbedded() ConfigSettingFormulasModelHaljsonEmbedded`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *ConfigSettingFormulasModelHaljson) GetEmbeddedOk() (*ConfigSettingFormulasModelHaljsonEmbedded, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *ConfigSettingFormulasModelHaljson) SetEmbedded(v ConfigSettingFormulasModelHaljsonEmbedded)`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *ConfigSettingFormulasModelHaljson) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetReadOnly

`func (o *ConfigSettingFormulasModelHaljson) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *ConfigSettingFormulasModelHaljson) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *ConfigSettingFormulasModelHaljson) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *ConfigSettingFormulasModelHaljson) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetSettingFormulas

`func (o *ConfigSettingFormulasModelHaljson) GetSettingFormulas() []ConfigSettingFormulaModel`

GetSettingFormulas returns the SettingFormulas field if non-nil, zero value otherwise.

### GetSettingFormulasOk

`func (o *ConfigSettingFormulasModelHaljson) GetSettingFormulasOk() (*[]ConfigSettingFormulaModel, bool)`

GetSettingFormulasOk returns a tuple with the SettingFormulas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingFormulas

`func (o *ConfigSettingFormulasModelHaljson) SetSettingFormulas(v []ConfigSettingFormulaModel)`

SetSettingFormulas sets SettingFormulas field to given value.

### HasSettingFormulas

`func (o *ConfigSettingFormulasModelHaljson) HasSettingFormulas() bool`

HasSettingFormulas returns a boolean if a field has been set.

### SetSettingFormulasNil

`func (o *ConfigSettingFormulasModelHaljson) SetSettingFormulasNil(b bool)`

 SetSettingFormulasNil sets the value for SettingFormulas to be an explicit nil

### UnsetSettingFormulas
`func (o *ConfigSettingFormulasModelHaljson) UnsetSettingFormulas()`

UnsetSettingFormulas ensures that no value is present for SettingFormulas, not even an explicit nil
### GetLinks

`func (o *ConfigSettingFormulasModelHaljson) GetLinks() ConfigSettingFormulasModelHaljsonEmbeddedEnvironmentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ConfigSettingFormulasModelHaljson) GetLinksOk() (*ConfigSettingFormulasModelHaljsonEmbeddedEnvironmentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ConfigSettingFormulasModelHaljson) SetLinks(v ConfigSettingFormulasModelHaljsonEmbeddedEnvironmentLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ConfigSettingFormulasModelHaljson) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


