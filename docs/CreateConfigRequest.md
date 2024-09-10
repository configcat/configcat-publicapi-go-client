# CreateConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the Config. | 
**Description** | Pointer to **NullableString** | The description of the Config. | [optional] 
**Order** | Pointer to **NullableInt32** | The order of the Config represented on the ConfigCat Dashboard.  Determined from an ascending sequence of integers. | [optional] 
**EvaluationVersion** | Pointer to **string** | Determines the evaluation version of a Config.  Using &#x60;v2&#x60; enables the new features of Config V2 (https://configcat.com/docs/advanced/config-v2). | [optional] 

## Methods

### NewCreateConfigRequest

`func NewCreateConfigRequest(name string, ) *CreateConfigRequest`

NewCreateConfigRequest instantiates a new CreateConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateConfigRequestWithDefaults

`func NewCreateConfigRequestWithDefaults() *CreateConfigRequest`

NewCreateConfigRequestWithDefaults instantiates a new CreateConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateConfigRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateConfigRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateConfigRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateConfigRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateConfigRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateConfigRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateConfigRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateConfigRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateConfigRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetOrder

`func (o *CreateConfigRequest) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *CreateConfigRequest) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *CreateConfigRequest) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *CreateConfigRequest) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### SetOrderNil

`func (o *CreateConfigRequest) SetOrderNil(b bool)`

 SetOrderNil sets the value for Order to be an explicit nil

### UnsetOrder
`func (o *CreateConfigRequest) UnsetOrder()`

UnsetOrder ensures that no value is present for Order, not even an explicit nil
### GetEvaluationVersion

`func (o *CreateConfigRequest) GetEvaluationVersion() string`

GetEvaluationVersion returns the EvaluationVersion field if non-nil, zero value otherwise.

### GetEvaluationVersionOk

`func (o *CreateConfigRequest) GetEvaluationVersionOk() (*string, bool)`

GetEvaluationVersionOk returns a tuple with the EvaluationVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationVersion

`func (o *CreateConfigRequest) SetEvaluationVersion(v string)`

SetEvaluationVersion sets EvaluationVersion field to given value.

### HasEvaluationVersion

`func (o *CreateConfigRequest) HasEvaluationVersion() bool`

HasEvaluationVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


