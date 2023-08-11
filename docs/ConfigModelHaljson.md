# ConfigModelHaljson

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

### NewConfigModelHaljson

`func NewConfigModelHaljson() *ConfigModelHaljson`

NewConfigModelHaljson instantiates a new ConfigModelHaljson object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigModelHaljsonWithDefaults

`func NewConfigModelHaljsonWithDefaults() *ConfigModelHaljson`

NewConfigModelHaljsonWithDefaults instantiates a new ConfigModelHaljson object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmbedded

`func (o *ConfigModelHaljson) GetEmbedded() ConfigModelHaljsonEmbedded`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *ConfigModelHaljson) GetEmbeddedOk() (*ConfigModelHaljsonEmbedded, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *ConfigModelHaljson) SetEmbedded(v ConfigModelHaljsonEmbedded)`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *ConfigModelHaljson) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetConfigId

`func (o *ConfigModelHaljson) GetConfigId() string`

GetConfigId returns the ConfigId field if non-nil, zero value otherwise.

### GetConfigIdOk

`func (o *ConfigModelHaljson) GetConfigIdOk() (*string, bool)`

GetConfigIdOk returns a tuple with the ConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigId

`func (o *ConfigModelHaljson) SetConfigId(v string)`

SetConfigId sets ConfigId field to given value.

### HasConfigId

`func (o *ConfigModelHaljson) HasConfigId() bool`

HasConfigId returns a boolean if a field has been set.

### GetName

`func (o *ConfigModelHaljson) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConfigModelHaljson) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConfigModelHaljson) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConfigModelHaljson) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ConfigModelHaljson) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ConfigModelHaljson) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *ConfigModelHaljson) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConfigModelHaljson) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConfigModelHaljson) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConfigModelHaljson) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ConfigModelHaljson) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ConfigModelHaljson) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetOrder

`func (o *ConfigModelHaljson) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ConfigModelHaljson) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ConfigModelHaljson) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *ConfigModelHaljson) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetMigratedConfigId

`func (o *ConfigModelHaljson) GetMigratedConfigId() string`

GetMigratedConfigId returns the MigratedConfigId field if non-nil, zero value otherwise.

### GetMigratedConfigIdOk

`func (o *ConfigModelHaljson) GetMigratedConfigIdOk() (*string, bool)`

GetMigratedConfigIdOk returns a tuple with the MigratedConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigratedConfigId

`func (o *ConfigModelHaljson) SetMigratedConfigId(v string)`

SetMigratedConfigId sets MigratedConfigId field to given value.

### HasMigratedConfigId

`func (o *ConfigModelHaljson) HasMigratedConfigId() bool`

HasMigratedConfigId returns a boolean if a field has been set.

### SetMigratedConfigIdNil

`func (o *ConfigModelHaljson) SetMigratedConfigIdNil(b bool)`

 SetMigratedConfigIdNil sets the value for MigratedConfigId to be an explicit nil

### UnsetMigratedConfigId
`func (o *ConfigModelHaljson) UnsetMigratedConfigId()`

UnsetMigratedConfigId ensures that no value is present for MigratedConfigId, not even an explicit nil
### GetEvaluationVersion

`func (o *ConfigModelHaljson) GetEvaluationVersion() EvaluationVersion`

GetEvaluationVersion returns the EvaluationVersion field if non-nil, zero value otherwise.

### GetEvaluationVersionOk

`func (o *ConfigModelHaljson) GetEvaluationVersionOk() (*EvaluationVersion, bool)`

GetEvaluationVersionOk returns a tuple with the EvaluationVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationVersion

`func (o *ConfigModelHaljson) SetEvaluationVersion(v EvaluationVersion)`

SetEvaluationVersion sets EvaluationVersion field to given value.

### HasEvaluationVersion

`func (o *ConfigModelHaljson) HasEvaluationVersion() bool`

HasEvaluationVersion returns a boolean if a field has been set.

### GetLinks

`func (o *ConfigModelHaljson) GetLinks() ConfigModelHaljsonLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ConfigModelHaljson) GetLinksOk() (*ConfigModelHaljsonLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ConfigModelHaljson) SetLinks(v ConfigModelHaljsonLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ConfigModelHaljson) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


