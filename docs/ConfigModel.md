# ConfigModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Product** | Pointer to [**ProductModel**](ProductModel.md) |  | [optional] 
**ConfigId** | Pointer to **string** | Identifier of the Config. | [optional] 
**Name** | Pointer to **NullableString** | Name of the Config. | [optional] 
**Description** | Pointer to **NullableString** | Description of the Config. | [optional] 
**Order** | Pointer to **int32** | The order of the Config represented on the ConfigCat Dashboard. | [optional] 
**MigratedConfigId** | Pointer to **NullableString** |  | [optional] 
**EvaluationVersion** | Pointer to **string** | Determines the evaluation version of a Config.  Using &#x60;v2&#x60; enables the new features of Config V2 (https://configcat.com/docs/advanced/config-v2). | [optional] 

## Methods

### NewConfigModel

`func NewConfigModel() *ConfigModel`

NewConfigModel instantiates a new ConfigModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigModelWithDefaults

`func NewConfigModelWithDefaults() *ConfigModel`

NewConfigModelWithDefaults instantiates a new ConfigModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProduct

`func (o *ConfigModel) GetProduct() ProductModel`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *ConfigModel) GetProductOk() (*ProductModel, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *ConfigModel) SetProduct(v ProductModel)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *ConfigModel) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetConfigId

`func (o *ConfigModel) GetConfigId() string`

GetConfigId returns the ConfigId field if non-nil, zero value otherwise.

### GetConfigIdOk

`func (o *ConfigModel) GetConfigIdOk() (*string, bool)`

GetConfigIdOk returns a tuple with the ConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigId

`func (o *ConfigModel) SetConfigId(v string)`

SetConfigId sets ConfigId field to given value.

### HasConfigId

`func (o *ConfigModel) HasConfigId() bool`

HasConfigId returns a boolean if a field has been set.

### GetName

`func (o *ConfigModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConfigModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConfigModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConfigModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ConfigModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ConfigModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *ConfigModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConfigModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConfigModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConfigModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ConfigModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ConfigModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetOrder

`func (o *ConfigModel) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ConfigModel) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ConfigModel) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *ConfigModel) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetMigratedConfigId

`func (o *ConfigModel) GetMigratedConfigId() string`

GetMigratedConfigId returns the MigratedConfigId field if non-nil, zero value otherwise.

### GetMigratedConfigIdOk

`func (o *ConfigModel) GetMigratedConfigIdOk() (*string, bool)`

GetMigratedConfigIdOk returns a tuple with the MigratedConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigratedConfigId

`func (o *ConfigModel) SetMigratedConfigId(v string)`

SetMigratedConfigId sets MigratedConfigId field to given value.

### HasMigratedConfigId

`func (o *ConfigModel) HasMigratedConfigId() bool`

HasMigratedConfigId returns a boolean if a field has been set.

### SetMigratedConfigIdNil

`func (o *ConfigModel) SetMigratedConfigIdNil(b bool)`

 SetMigratedConfigIdNil sets the value for MigratedConfigId to be an explicit nil

### UnsetMigratedConfigId
`func (o *ConfigModel) UnsetMigratedConfigId()`

UnsetMigratedConfigId ensures that no value is present for MigratedConfigId, not even an explicit nil
### GetEvaluationVersion

`func (o *ConfigModel) GetEvaluationVersion() string`

GetEvaluationVersion returns the EvaluationVersion field if non-nil, zero value otherwise.

### GetEvaluationVersionOk

`func (o *ConfigModel) GetEvaluationVersionOk() (*string, bool)`

GetEvaluationVersionOk returns a tuple with the EvaluationVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationVersion

`func (o *ConfigModel) SetEvaluationVersion(v string)`

SetEvaluationVersion sets EvaluationVersion field to given value.

### HasEvaluationVersion

`func (o *ConfigModel) HasEvaluationVersion() bool`

HasEvaluationVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


