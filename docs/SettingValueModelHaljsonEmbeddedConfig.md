# SettingValueModelHaljsonEmbeddedConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Embedded** | Pointer to [**ConfigModelHaljsonEmbedded**](ConfigModelHaljsonEmbedded.md) |  | [optional] 
**ConfigId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Order** | Pointer to **int32** |  | [optional] 
**MigratedConfigId** | Pointer to **NullableString** |  | [optional] 
**EvaluationVersion** | Pointer to [**EvaluationVersion**](EvaluationVersion.md) |  | [optional] 
**Links** | Pointer to [**ConfigModelHaljsonLinks**](ConfigModelHaljsonLinks.md) |  | [optional] 

## Methods

### NewSettingValueModelHaljsonEmbeddedConfig

`func NewSettingValueModelHaljsonEmbeddedConfig() *SettingValueModelHaljsonEmbeddedConfig`

NewSettingValueModelHaljsonEmbeddedConfig instantiates a new SettingValueModelHaljsonEmbeddedConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettingValueModelHaljsonEmbeddedConfigWithDefaults

`func NewSettingValueModelHaljsonEmbeddedConfigWithDefaults() *SettingValueModelHaljsonEmbeddedConfig`

NewSettingValueModelHaljsonEmbeddedConfigWithDefaults instantiates a new SettingValueModelHaljsonEmbeddedConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmbedded

`func (o *SettingValueModelHaljsonEmbeddedConfig) GetEmbedded() ConfigModelHaljsonEmbedded`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *SettingValueModelHaljsonEmbeddedConfig) GetEmbeddedOk() (*ConfigModelHaljsonEmbedded, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *SettingValueModelHaljsonEmbeddedConfig) SetEmbedded(v ConfigModelHaljsonEmbedded)`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *SettingValueModelHaljsonEmbeddedConfig) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetConfigId

`func (o *SettingValueModelHaljsonEmbeddedConfig) GetConfigId() string`

GetConfigId returns the ConfigId field if non-nil, zero value otherwise.

### GetConfigIdOk

`func (o *SettingValueModelHaljsonEmbeddedConfig) GetConfigIdOk() (*string, bool)`

GetConfigIdOk returns a tuple with the ConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigId

`func (o *SettingValueModelHaljsonEmbeddedConfig) SetConfigId(v string)`

SetConfigId sets ConfigId field to given value.

### HasConfigId

`func (o *SettingValueModelHaljsonEmbeddedConfig) HasConfigId() bool`

HasConfigId returns a boolean if a field has been set.

### GetName

`func (o *SettingValueModelHaljsonEmbeddedConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SettingValueModelHaljsonEmbeddedConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SettingValueModelHaljsonEmbeddedConfig) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SettingValueModelHaljsonEmbeddedConfig) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *SettingValueModelHaljsonEmbeddedConfig) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *SettingValueModelHaljsonEmbeddedConfig) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *SettingValueModelHaljsonEmbeddedConfig) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SettingValueModelHaljsonEmbeddedConfig) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SettingValueModelHaljsonEmbeddedConfig) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SettingValueModelHaljsonEmbeddedConfig) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *SettingValueModelHaljsonEmbeddedConfig) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SettingValueModelHaljsonEmbeddedConfig) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetOrder

`func (o *SettingValueModelHaljsonEmbeddedConfig) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *SettingValueModelHaljsonEmbeddedConfig) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *SettingValueModelHaljsonEmbeddedConfig) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *SettingValueModelHaljsonEmbeddedConfig) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetMigratedConfigId

`func (o *SettingValueModelHaljsonEmbeddedConfig) GetMigratedConfigId() string`

GetMigratedConfigId returns the MigratedConfigId field if non-nil, zero value otherwise.

### GetMigratedConfigIdOk

`func (o *SettingValueModelHaljsonEmbeddedConfig) GetMigratedConfigIdOk() (*string, bool)`

GetMigratedConfigIdOk returns a tuple with the MigratedConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigratedConfigId

`func (o *SettingValueModelHaljsonEmbeddedConfig) SetMigratedConfigId(v string)`

SetMigratedConfigId sets MigratedConfigId field to given value.

### HasMigratedConfigId

`func (o *SettingValueModelHaljsonEmbeddedConfig) HasMigratedConfigId() bool`

HasMigratedConfigId returns a boolean if a field has been set.

### SetMigratedConfigIdNil

`func (o *SettingValueModelHaljsonEmbeddedConfig) SetMigratedConfigIdNil(b bool)`

 SetMigratedConfigIdNil sets the value for MigratedConfigId to be an explicit nil

### UnsetMigratedConfigId
`func (o *SettingValueModelHaljsonEmbeddedConfig) UnsetMigratedConfigId()`

UnsetMigratedConfigId ensures that no value is present for MigratedConfigId, not even an explicit nil
### GetEvaluationVersion

`func (o *SettingValueModelHaljsonEmbeddedConfig) GetEvaluationVersion() EvaluationVersion`

GetEvaluationVersion returns the EvaluationVersion field if non-nil, zero value otherwise.

### GetEvaluationVersionOk

`func (o *SettingValueModelHaljsonEmbeddedConfig) GetEvaluationVersionOk() (*EvaluationVersion, bool)`

GetEvaluationVersionOk returns a tuple with the EvaluationVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationVersion

`func (o *SettingValueModelHaljsonEmbeddedConfig) SetEvaluationVersion(v EvaluationVersion)`

SetEvaluationVersion sets EvaluationVersion field to given value.

### HasEvaluationVersion

`func (o *SettingValueModelHaljsonEmbeddedConfig) HasEvaluationVersion() bool`

HasEvaluationVersion returns a boolean if a field has been set.

### GetLinks

`func (o *SettingValueModelHaljsonEmbeddedConfig) GetLinks() ConfigModelHaljsonLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *SettingValueModelHaljsonEmbeddedConfig) GetLinksOk() (*ConfigModelHaljsonLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *SettingValueModelHaljsonEmbeddedConfig) SetLinks(v ConfigModelHaljsonLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *SettingValueModelHaljsonEmbeddedConfig) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


