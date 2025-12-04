# EnvironmentAccessModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentId** | **string** | Identifier of the Environment. | 
**Name** | **string** | Name of the Environment. | 
**Color** | **NullableString** | Color of the Environment. | 
**Description** | **NullableString** | Description of the Environment. | 
**Order** | **int32** | The order of the Environment represented on the ConfigCat Dashboard. | 
**ReasonRequired** | **bool** | Determines whether a mandatory reason must be given every time when the Feature Flags or Settings in the given Environment are saved. | 
**EnvironmentAccessType** | [**EnvironmentAccessType**](EnvironmentAccessType.md) |  | 

## Methods

### NewEnvironmentAccessModel

`func NewEnvironmentAccessModel(environmentId string, name string, color NullableString, description NullableString, order int32, reasonRequired bool, environmentAccessType EnvironmentAccessType, ) *EnvironmentAccessModel`

NewEnvironmentAccessModel instantiates a new EnvironmentAccessModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentAccessModelWithDefaults

`func NewEnvironmentAccessModelWithDefaults() *EnvironmentAccessModel`

NewEnvironmentAccessModelWithDefaults instantiates a new EnvironmentAccessModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentId

`func (o *EnvironmentAccessModel) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *EnvironmentAccessModel) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *EnvironmentAccessModel) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetName

`func (o *EnvironmentAccessModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnvironmentAccessModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnvironmentAccessModel) SetName(v string)`

SetName sets Name field to given value.


### GetColor

`func (o *EnvironmentAccessModel) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *EnvironmentAccessModel) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *EnvironmentAccessModel) SetColor(v string)`

SetColor sets Color field to given value.


### SetColorNil

`func (o *EnvironmentAccessModel) SetColorNil(b bool)`

 SetColorNil sets the value for Color to be an explicit nil

### UnsetColor
`func (o *EnvironmentAccessModel) UnsetColor()`

UnsetColor ensures that no value is present for Color, not even an explicit nil
### GetDescription

`func (o *EnvironmentAccessModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EnvironmentAccessModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EnvironmentAccessModel) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *EnvironmentAccessModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *EnvironmentAccessModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetOrder

`func (o *EnvironmentAccessModel) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *EnvironmentAccessModel) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *EnvironmentAccessModel) SetOrder(v int32)`

SetOrder sets Order field to given value.


### GetReasonRequired

`func (o *EnvironmentAccessModel) GetReasonRequired() bool`

GetReasonRequired returns the ReasonRequired field if non-nil, zero value otherwise.

### GetReasonRequiredOk

`func (o *EnvironmentAccessModel) GetReasonRequiredOk() (*bool, bool)`

GetReasonRequiredOk returns a tuple with the ReasonRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonRequired

`func (o *EnvironmentAccessModel) SetReasonRequired(v bool)`

SetReasonRequired sets ReasonRequired field to given value.


### GetEnvironmentAccessType

`func (o *EnvironmentAccessModel) GetEnvironmentAccessType() EnvironmentAccessType`

GetEnvironmentAccessType returns the EnvironmentAccessType field if non-nil, zero value otherwise.

### GetEnvironmentAccessTypeOk

`func (o *EnvironmentAccessModel) GetEnvironmentAccessTypeOk() (*EnvironmentAccessType, bool)`

GetEnvironmentAccessTypeOk returns a tuple with the EnvironmentAccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentAccessType

`func (o *EnvironmentAccessModel) SetEnvironmentAccessType(v EnvironmentAccessType)`

SetEnvironmentAccessType sets EnvironmentAccessType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


