# CreateEnvironmentModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the Environment. | 
**Color** | Pointer to **NullableString** | The color of the Environment. RGB or HTML color codes are allowed. | [optional] 
**Description** | Pointer to **NullableString** | The description of the Environment. | [optional] 
**Order** | Pointer to **NullableInt32** | The order of the Environment represented on the ConfigCat Dashboard.  Determined from an ascending sequence of integers. | [optional] 

## Methods

### NewCreateEnvironmentModel

`func NewCreateEnvironmentModel(name string, ) *CreateEnvironmentModel`

NewCreateEnvironmentModel instantiates a new CreateEnvironmentModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateEnvironmentModelWithDefaults

`func NewCreateEnvironmentModelWithDefaults() *CreateEnvironmentModel`

NewCreateEnvironmentModelWithDefaults instantiates a new CreateEnvironmentModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateEnvironmentModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateEnvironmentModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateEnvironmentModel) SetName(v string)`

SetName sets Name field to given value.


### GetColor

`func (o *CreateEnvironmentModel) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *CreateEnvironmentModel) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *CreateEnvironmentModel) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *CreateEnvironmentModel) HasColor() bool`

HasColor returns a boolean if a field has been set.

### SetColorNil

`func (o *CreateEnvironmentModel) SetColorNil(b bool)`

 SetColorNil sets the value for Color to be an explicit nil

### UnsetColor
`func (o *CreateEnvironmentModel) UnsetColor()`

UnsetColor ensures that no value is present for Color, not even an explicit nil
### GetDescription

`func (o *CreateEnvironmentModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateEnvironmentModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateEnvironmentModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateEnvironmentModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateEnvironmentModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateEnvironmentModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetOrder

`func (o *CreateEnvironmentModel) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *CreateEnvironmentModel) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *CreateEnvironmentModel) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *CreateEnvironmentModel) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### SetOrderNil

`func (o *CreateEnvironmentModel) SetOrderNil(b bool)`

 SetOrderNil sets the value for Order to be an explicit nil

### UnsetOrder
`func (o *CreateEnvironmentModel) UnsetOrder()`

UnsetOrder ensures that no value is present for Order, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


