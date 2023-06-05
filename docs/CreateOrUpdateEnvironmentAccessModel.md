# CreateOrUpdateEnvironmentAccessModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Color** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Order** | Pointer to **int32** |  | [optional] 
**ReasonRequired** | Pointer to **bool** |  | [optional] 
**EnvironmentAccessType** | Pointer to [**EnvironmentAccessType**](EnvironmentAccessType.md) |  | [optional] 

## Methods

### NewCreateOrUpdateEnvironmentAccessModel

`func NewCreateOrUpdateEnvironmentAccessModel() *CreateOrUpdateEnvironmentAccessModel`

NewCreateOrUpdateEnvironmentAccessModel instantiates a new CreateOrUpdateEnvironmentAccessModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrUpdateEnvironmentAccessModelWithDefaults

`func NewCreateOrUpdateEnvironmentAccessModelWithDefaults() *CreateOrUpdateEnvironmentAccessModel`

NewCreateOrUpdateEnvironmentAccessModelWithDefaults instantiates a new CreateOrUpdateEnvironmentAccessModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentId

`func (o *CreateOrUpdateEnvironmentAccessModel) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *CreateOrUpdateEnvironmentAccessModel) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *CreateOrUpdateEnvironmentAccessModel) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *CreateOrUpdateEnvironmentAccessModel) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### GetName

`func (o *CreateOrUpdateEnvironmentAccessModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateOrUpdateEnvironmentAccessModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateOrUpdateEnvironmentAccessModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateOrUpdateEnvironmentAccessModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CreateOrUpdateEnvironmentAccessModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateOrUpdateEnvironmentAccessModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetColor

`func (o *CreateOrUpdateEnvironmentAccessModel) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *CreateOrUpdateEnvironmentAccessModel) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *CreateOrUpdateEnvironmentAccessModel) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *CreateOrUpdateEnvironmentAccessModel) HasColor() bool`

HasColor returns a boolean if a field has been set.

### SetColorNil

`func (o *CreateOrUpdateEnvironmentAccessModel) SetColorNil(b bool)`

 SetColorNil sets the value for Color to be an explicit nil

### UnsetColor
`func (o *CreateOrUpdateEnvironmentAccessModel) UnsetColor()`

UnsetColor ensures that no value is present for Color, not even an explicit nil
### GetDescription

`func (o *CreateOrUpdateEnvironmentAccessModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateOrUpdateEnvironmentAccessModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateOrUpdateEnvironmentAccessModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateOrUpdateEnvironmentAccessModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateOrUpdateEnvironmentAccessModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateOrUpdateEnvironmentAccessModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetOrder

`func (o *CreateOrUpdateEnvironmentAccessModel) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *CreateOrUpdateEnvironmentAccessModel) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *CreateOrUpdateEnvironmentAccessModel) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *CreateOrUpdateEnvironmentAccessModel) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetReasonRequired

`func (o *CreateOrUpdateEnvironmentAccessModel) GetReasonRequired() bool`

GetReasonRequired returns the ReasonRequired field if non-nil, zero value otherwise.

### GetReasonRequiredOk

`func (o *CreateOrUpdateEnvironmentAccessModel) GetReasonRequiredOk() (*bool, bool)`

GetReasonRequiredOk returns a tuple with the ReasonRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonRequired

`func (o *CreateOrUpdateEnvironmentAccessModel) SetReasonRequired(v bool)`

SetReasonRequired sets ReasonRequired field to given value.

### HasReasonRequired

`func (o *CreateOrUpdateEnvironmentAccessModel) HasReasonRequired() bool`

HasReasonRequired returns a boolean if a field has been set.

### GetEnvironmentAccessType

`func (o *CreateOrUpdateEnvironmentAccessModel) GetEnvironmentAccessType() EnvironmentAccessType`

GetEnvironmentAccessType returns the EnvironmentAccessType field if non-nil, zero value otherwise.

### GetEnvironmentAccessTypeOk

`func (o *CreateOrUpdateEnvironmentAccessModel) GetEnvironmentAccessTypeOk() (*EnvironmentAccessType, bool)`

GetEnvironmentAccessTypeOk returns a tuple with the EnvironmentAccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentAccessType

`func (o *CreateOrUpdateEnvironmentAccessModel) SetEnvironmentAccessType(v EnvironmentAccessType)`

SetEnvironmentAccessType sets EnvironmentAccessType field to given value.

### HasEnvironmentAccessType

`func (o *CreateOrUpdateEnvironmentAccessModel) HasEnvironmentAccessType() bool`

HasEnvironmentAccessType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


