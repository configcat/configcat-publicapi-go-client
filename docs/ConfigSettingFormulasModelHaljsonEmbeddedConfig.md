# ConfigSettingFormulasModelHaljsonEmbeddedConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Embedded** | Pointer to [**ConfigModelHaljsonEmbedded**](ConfigModelHaljsonEmbedded.md) |  | [optional] 
**ConfigId** | Pointer to **string** | Identifier of the Config. | [optional] 
**Name** | Pointer to **NullableString** | Name of the Config. | [optional] 
**Description** | Pointer to **NullableString** | Description of the Config. | [optional] 
**Order** | Pointer to **int32** | The order of the Config represented on the ConfigCat Dashboard. | [optional] 
**MigratedConfigId** | Pointer to **NullableString** |  | [optional] 
**EvaluationVersion** | Pointer to [**EvaluationVersion**](EvaluationVersion.md) |  | [optional] 
**Links** | Pointer to [**ConfigModelHaljsonLinks**](ConfigModelHaljsonLinks.md) |  | [optional] 

## Methods

### NewConfigSettingFormulasModelHaljsonEmbeddedConfig

`func NewConfigSettingFormulasModelHaljsonEmbeddedConfig() *ConfigSettingFormulasModelHaljsonEmbeddedConfig`

NewConfigSettingFormulasModelHaljsonEmbeddedConfig instantiates a new ConfigSettingFormulasModelHaljsonEmbeddedConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigSettingFormulasModelHaljsonEmbeddedConfigWithDefaults

`func NewConfigSettingFormulasModelHaljsonEmbeddedConfigWithDefaults() *ConfigSettingFormulasModelHaljsonEmbeddedConfig`

NewConfigSettingFormulasModelHaljsonEmbeddedConfigWithDefaults instantiates a new ConfigSettingFormulasModelHaljsonEmbeddedConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmbedded

`func (o *ConfigSettingFormulasModelHaljsonEmbeddedConfig) GetEmbedded() ConfigModelHaljsonEmbedded`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *ConfigSettingFormulasModelHaljsonEmbeddedConfig) GetEmbeddedOk() (*ConfigModelHaljsonEmbedded, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *ConfigSettingFormulasModelHaljsonEmbeddedConfig) SetEmbedded(v ConfigModelHaljsonEmbedded)`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *ConfigSettingFormulasModelHaljsonEmbeddedConfig) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetConfigId

`func (o *ConfigSettingFormulasModelHaljsonEmbeddedConfig) GetConfigId() string`

GetConfigId returns the ConfigId field if non-nil, zero value otherwise.

### GetConfigIdOk

`func (o *ConfigSettingFormulasModelHaljsonEmbeddedConfig) GetConfigIdOk() (*string, bool)`

GetConfigIdOk returns a tuple with the ConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigId

`func (o *ConfigSettingFormulasModelHaljsonEmbeddedConfig) SetConfigId(v string)`

SetConfigId sets ConfigId field to given value.

### HasConfigId

`func (o *ConfigSettingFormulasModelHaljsonEmbeddedConfig) HasConfigId() bool`

HasConfigId returns a boolean if a field has been set.

### GetName

`func (o *ConfigSettingFormulasModelHaljsonEmbeddedConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConfigSettingFormulasModelHaljsonEmbeddedConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConfigSettingFormulasModelHaljsonEmbeddedConfig) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConfigSettingFormulasModelHaljsonEmbeddedConfig) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ConfigSettingFormulasModelHaljsonEmbeddedConfig) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ConfigSettingFormulasModelHaljsonEmbeddedConfig) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *ConfigSettingFormulasModelHaljsonEmbeddedConfig) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConfigSettingFormulasModelHaljsonEmbeddedConfig) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConfigSettingFormulasModelHaljsonEmbeddedConfig) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConfigSettingFormulasModelHaljsonEmbeddedConfig) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ConfigSettingFormulasModelHaljsonEmbeddedConfig) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ConfigSettingFormulasModelHaljsonEmbeddedConfig) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetOrder

`func (o *ConfigSettingFormulasModelHaljsonEmbeddedConfig) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ConfigSettingFormulasModelHaljsonEmbeddedConfig) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ConfigSettingFormulasModelHaljsonEmbeddedConfig) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *ConfigSettingFormulasModelHaljsonEmbeddedConfig) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetMigratedConfigId

`func (o *ConfigSettingFormulasModelHaljsonEmbeddedConfig) GetMigratedConfigId() string`

GetMigratedConfigId returns the MigratedConfigId field if non-nil, zero value otherwise.

### GetMigratedConfigIdOk

`func (o *ConfigSettingFormulasModelHaljsonEmbeddedConfig) GetMigratedConfigIdOk() (*string, bool)`

GetMigratedConfigIdOk returns a tuple with the MigratedConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigratedConfigId

`func (o *ConfigSettingFormulasModelHaljsonEmbeddedConfig) SetMigratedConfigId(v string)`

SetMigratedConfigId sets MigratedConfigId field to given value.

### HasMigratedConfigId

`func (o *ConfigSettingFormulasModelHaljsonEmbeddedConfig) HasMigratedConfigId() bool`

HasMigratedConfigId returns a boolean if a field has been set.

### SetMigratedConfigIdNil

`func (o *ConfigSettingFormulasModelHaljsonEmbeddedConfig) SetMigratedConfigIdNil(b bool)`

 SetMigratedConfigIdNil sets the value for MigratedConfigId to be an explicit nil

### UnsetMigratedConfigId
`func (o *ConfigSettingFormulasModelHaljsonEmbeddedConfig) UnsetMigratedConfigId()`

UnsetMigratedConfigId ensures that no value is present for MigratedConfigId, not even an explicit nil
### GetEvaluationVersion

`func (o *ConfigSettingFormulasModelHaljsonEmbeddedConfig) GetEvaluationVersion() EvaluationVersion`

GetEvaluationVersion returns the EvaluationVersion field if non-nil, zero value otherwise.

### GetEvaluationVersionOk

`func (o *ConfigSettingFormulasModelHaljsonEmbeddedConfig) GetEvaluationVersionOk() (*EvaluationVersion, bool)`

GetEvaluationVersionOk returns a tuple with the EvaluationVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationVersion

`func (o *ConfigSettingFormulasModelHaljsonEmbeddedConfig) SetEvaluationVersion(v EvaluationVersion)`

SetEvaluationVersion sets EvaluationVersion field to given value.

### HasEvaluationVersion

`func (o *ConfigSettingFormulasModelHaljsonEmbeddedConfig) HasEvaluationVersion() bool`

HasEvaluationVersion returns a boolean if a field has been set.

### GetLinks

`func (o *ConfigSettingFormulasModelHaljsonEmbeddedConfig) GetLinks() ConfigModelHaljsonLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ConfigSettingFormulasModelHaljsonEmbeddedConfig) GetLinksOk() (*ConfigModelHaljsonLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ConfigSettingFormulasModelHaljsonEmbeddedConfig) SetLinks(v ConfigModelHaljsonLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ConfigSettingFormulasModelHaljsonEmbeddedConfig) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


