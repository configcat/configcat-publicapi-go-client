# EnvironmentAccessModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentId** | Pointer to **string** | Identifier of the Environment. | [optional] 
**Name** | Pointer to **NullableString** | Name of the Environment. | [optional] 
**Color** | Pointer to **NullableString** | Color of the Environment. | [optional] 
**Description** | Pointer to **NullableString** | Description of the Environment. | [optional] 
**Order** | Pointer to **int32** | The order of the Environment represented on the ConfigCat Dashboard. | [optional] 
**ReasonRequired** | Pointer to **bool** | Determines whether a mandatory reason must be given every time when the Feature Flags or Settings in the given Environment are saved. | [optional] 
**EnvironmentAccessType** | Pointer to [**EnvironmentAccessType**](EnvironmentAccessType.md) |  | [optional] 

## Methods

### NewEnvironmentAccessModel

`func NewEnvironmentAccessModel() *EnvironmentAccessModel`

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

### HasEnvironmentId

`func (o *EnvironmentAccessModel) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

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

### HasName

`func (o *EnvironmentAccessModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *EnvironmentAccessModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *EnvironmentAccessModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
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

### HasColor

`func (o *EnvironmentAccessModel) HasColor() bool`

HasColor returns a boolean if a field has been set.

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

### HasDescription

`func (o *EnvironmentAccessModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

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

### HasOrder

`func (o *EnvironmentAccessModel) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

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

### HasReasonRequired

`func (o *EnvironmentAccessModel) HasReasonRequired() bool`

HasReasonRequired returns a boolean if a field has been set.

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

### HasEnvironmentAccessType

`func (o *EnvironmentAccessModel) HasEnvironmentAccessType() bool`

HasEnvironmentAccessType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


