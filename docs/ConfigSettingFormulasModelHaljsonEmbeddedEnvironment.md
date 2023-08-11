# ConfigSettingFormulasModelHaljsonEmbeddedEnvironment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Embedded** | Pointer to [**ConfigModelHaljsonEmbedded**](ConfigModelHaljsonEmbedded.md) |  | [optional] 
**EnvironmentId** | Pointer to **string** | Identifier of the Environment. | [optional] 
**Name** | Pointer to **NullableString** | Name of the Environment. | [optional] 
**Color** | Pointer to **NullableString** | The configured color of the Environment. | [optional] 
**Description** | Pointer to **NullableString** | Description of the Environment. | [optional] 
**Order** | Pointer to **int32** | The order of the Environment represented on the ConfigCat Dashboard. | [optional] 
**ReasonRequired** | Pointer to **bool** | Determines whether a mandatory reason must be given every time when the Feature Flags or Settings in the given Environment are saved. | [optional] 
**Links** | Pointer to [**ConfigSettingFormulasModelHaljsonEmbeddedEnvironmentLinks**](ConfigSettingFormulasModelHaljsonEmbeddedEnvironmentLinks.md) |  | [optional] 

## Methods

### NewConfigSettingFormulasModelHaljsonEmbeddedEnvironment

`func NewConfigSettingFormulasModelHaljsonEmbeddedEnvironment() *ConfigSettingFormulasModelHaljsonEmbeddedEnvironment`

NewConfigSettingFormulasModelHaljsonEmbeddedEnvironment instantiates a new ConfigSettingFormulasModelHaljsonEmbeddedEnvironment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigSettingFormulasModelHaljsonEmbeddedEnvironmentWithDefaults

`func NewConfigSettingFormulasModelHaljsonEmbeddedEnvironmentWithDefaults() *ConfigSettingFormulasModelHaljsonEmbeddedEnvironment`

NewConfigSettingFormulasModelHaljsonEmbeddedEnvironmentWithDefaults instantiates a new ConfigSettingFormulasModelHaljsonEmbeddedEnvironment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmbedded

`func (o *ConfigSettingFormulasModelHaljsonEmbeddedEnvironment) GetEmbedded() ConfigModelHaljsonEmbedded`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *ConfigSettingFormulasModelHaljsonEmbeddedEnvironment) GetEmbeddedOk() (*ConfigModelHaljsonEmbedded, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *ConfigSettingFormulasModelHaljsonEmbeddedEnvironment) SetEmbedded(v ConfigModelHaljsonEmbedded)`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *ConfigSettingFormulasModelHaljsonEmbeddedEnvironment) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *ConfigSettingFormulasModelHaljsonEmbeddedEnvironment) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *ConfigSettingFormulasModelHaljsonEmbeddedEnvironment) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *ConfigSettingFormulasModelHaljsonEmbeddedEnvironment) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *ConfigSettingFormulasModelHaljsonEmbeddedEnvironment) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### GetName

`func (o *ConfigSettingFormulasModelHaljsonEmbeddedEnvironment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConfigSettingFormulasModelHaljsonEmbeddedEnvironment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConfigSettingFormulasModelHaljsonEmbeddedEnvironment) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConfigSettingFormulasModelHaljsonEmbeddedEnvironment) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ConfigSettingFormulasModelHaljsonEmbeddedEnvironment) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ConfigSettingFormulasModelHaljsonEmbeddedEnvironment) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetColor

`func (o *ConfigSettingFormulasModelHaljsonEmbeddedEnvironment) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *ConfigSettingFormulasModelHaljsonEmbeddedEnvironment) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *ConfigSettingFormulasModelHaljsonEmbeddedEnvironment) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *ConfigSettingFormulasModelHaljsonEmbeddedEnvironment) HasColor() bool`

HasColor returns a boolean if a field has been set.

### SetColorNil

`func (o *ConfigSettingFormulasModelHaljsonEmbeddedEnvironment) SetColorNil(b bool)`

 SetColorNil sets the value for Color to be an explicit nil

### UnsetColor
`func (o *ConfigSettingFormulasModelHaljsonEmbeddedEnvironment) UnsetColor()`

UnsetColor ensures that no value is present for Color, not even an explicit nil
### GetDescription

`func (o *ConfigSettingFormulasModelHaljsonEmbeddedEnvironment) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConfigSettingFormulasModelHaljsonEmbeddedEnvironment) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConfigSettingFormulasModelHaljsonEmbeddedEnvironment) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConfigSettingFormulasModelHaljsonEmbeddedEnvironment) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ConfigSettingFormulasModelHaljsonEmbeddedEnvironment) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ConfigSettingFormulasModelHaljsonEmbeddedEnvironment) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetOrder

`func (o *ConfigSettingFormulasModelHaljsonEmbeddedEnvironment) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ConfigSettingFormulasModelHaljsonEmbeddedEnvironment) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ConfigSettingFormulasModelHaljsonEmbeddedEnvironment) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *ConfigSettingFormulasModelHaljsonEmbeddedEnvironment) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetReasonRequired

`func (o *ConfigSettingFormulasModelHaljsonEmbeddedEnvironment) GetReasonRequired() bool`

GetReasonRequired returns the ReasonRequired field if non-nil, zero value otherwise.

### GetReasonRequiredOk

`func (o *ConfigSettingFormulasModelHaljsonEmbeddedEnvironment) GetReasonRequiredOk() (*bool, bool)`

GetReasonRequiredOk returns a tuple with the ReasonRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonRequired

`func (o *ConfigSettingFormulasModelHaljsonEmbeddedEnvironment) SetReasonRequired(v bool)`

SetReasonRequired sets ReasonRequired field to given value.

### HasReasonRequired

`func (o *ConfigSettingFormulasModelHaljsonEmbeddedEnvironment) HasReasonRequired() bool`

HasReasonRequired returns a boolean if a field has been set.

### GetLinks

`func (o *ConfigSettingFormulasModelHaljsonEmbeddedEnvironment) GetLinks() ConfigSettingFormulasModelHaljsonEmbeddedEnvironmentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ConfigSettingFormulasModelHaljsonEmbeddedEnvironment) GetLinksOk() (*ConfigSettingFormulasModelHaljsonEmbeddedEnvironmentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ConfigSettingFormulasModelHaljsonEmbeddedEnvironment) SetLinks(v ConfigSettingFormulasModelHaljsonEmbeddedEnvironmentLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ConfigSettingFormulasModelHaljsonEmbeddedEnvironment) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


