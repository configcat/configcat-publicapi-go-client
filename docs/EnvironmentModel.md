# EnvironmentModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Product** | [**ProductModel**](ProductModel.md) |  | 
**EnvironmentId** | **string** | Identifier of the Environment. | 
**Name** | **string** | Name of the Environment. | 
**Color** | **NullableString** | The configured color of the Environment. | 
**Description** | **NullableString** | Description of the Environment. | 
**Order** | **int32** | The order of the Environment represented on the ConfigCat Dashboard. | 
**ReasonRequired** | **bool** | Determines whether a mandatory reason must be given every time when the Feature Flags or Settings in the given Environment are saved. | 

## Methods

### NewEnvironmentModel

`func NewEnvironmentModel(product ProductModel, environmentId string, name string, color NullableString, description NullableString, order int32, reasonRequired bool, ) *EnvironmentModel`

NewEnvironmentModel instantiates a new EnvironmentModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentModelWithDefaults

`func NewEnvironmentModelWithDefaults() *EnvironmentModel`

NewEnvironmentModelWithDefaults instantiates a new EnvironmentModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProduct

`func (o *EnvironmentModel) GetProduct() ProductModel`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *EnvironmentModel) GetProductOk() (*ProductModel, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *EnvironmentModel) SetProduct(v ProductModel)`

SetProduct sets Product field to given value.


### GetEnvironmentId

`func (o *EnvironmentModel) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *EnvironmentModel) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *EnvironmentModel) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetName

`func (o *EnvironmentModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnvironmentModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnvironmentModel) SetName(v string)`

SetName sets Name field to given value.


### GetColor

`func (o *EnvironmentModel) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *EnvironmentModel) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *EnvironmentModel) SetColor(v string)`

SetColor sets Color field to given value.


### SetColorNil

`func (o *EnvironmentModel) SetColorNil(b bool)`

 SetColorNil sets the value for Color to be an explicit nil

### UnsetColor
`func (o *EnvironmentModel) UnsetColor()`

UnsetColor ensures that no value is present for Color, not even an explicit nil
### GetDescription

`func (o *EnvironmentModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EnvironmentModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EnvironmentModel) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *EnvironmentModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *EnvironmentModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetOrder

`func (o *EnvironmentModel) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *EnvironmentModel) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *EnvironmentModel) SetOrder(v int32)`

SetOrder sets Order field to given value.


### GetReasonRequired

`func (o *EnvironmentModel) GetReasonRequired() bool`

GetReasonRequired returns the ReasonRequired field if non-nil, zero value otherwise.

### GetReasonRequiredOk

`func (o *EnvironmentModel) GetReasonRequiredOk() (*bool, bool)`

GetReasonRequiredOk returns a tuple with the ReasonRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonRequired

`func (o *EnvironmentModel) SetReasonRequired(v bool)`

SetReasonRequired sets ReasonRequired field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


