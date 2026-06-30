# PermissionGroupEnvironmentAccessModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentId** | **string** | Identifier of the Environment. | 
**Name** | **string** | Name of the Environment. | 
**Color** | **NullableString** | Color of the Environment. | 
**Description** | **NullableString** | Description of the Environment. | 
**Order** | **int32** | The order of the Environment represented on the ConfigCat Dashboard. | 
**ReasonRequired** | **bool** | Determines whether a mandatory reason must be given every time when the Feature Flags or Settings in the given Environment are saved. | 
**ApproveRequired** | **bool** | Determines whether changes must be approved before they are applied in the given Environment. | 
**EnvironmentAccessType** | [**EnvironmentAccessType**](EnvironmentAccessType.md) |  | 

## Methods

### NewPermissionGroupEnvironmentAccessModel

`func NewPermissionGroupEnvironmentAccessModel(environmentId string, name string, color NullableString, description NullableString, order int32, reasonRequired bool, approveRequired bool, environmentAccessType EnvironmentAccessType, ) *PermissionGroupEnvironmentAccessModel`

NewPermissionGroupEnvironmentAccessModel instantiates a new PermissionGroupEnvironmentAccessModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionGroupEnvironmentAccessModelWithDefaults

`func NewPermissionGroupEnvironmentAccessModelWithDefaults() *PermissionGroupEnvironmentAccessModel`

NewPermissionGroupEnvironmentAccessModelWithDefaults instantiates a new PermissionGroupEnvironmentAccessModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentId

`func (o *PermissionGroupEnvironmentAccessModel) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *PermissionGroupEnvironmentAccessModel) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *PermissionGroupEnvironmentAccessModel) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetName

`func (o *PermissionGroupEnvironmentAccessModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PermissionGroupEnvironmentAccessModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PermissionGroupEnvironmentAccessModel) SetName(v string)`

SetName sets Name field to given value.


### GetColor

`func (o *PermissionGroupEnvironmentAccessModel) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *PermissionGroupEnvironmentAccessModel) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *PermissionGroupEnvironmentAccessModel) SetColor(v string)`

SetColor sets Color field to given value.


### SetColorNil

`func (o *PermissionGroupEnvironmentAccessModel) SetColorNil(b bool)`

 SetColorNil sets the value for Color to be an explicit nil

### UnsetColor
`func (o *PermissionGroupEnvironmentAccessModel) UnsetColor()`

UnsetColor ensures that no value is present for Color, not even an explicit nil
### GetDescription

`func (o *PermissionGroupEnvironmentAccessModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PermissionGroupEnvironmentAccessModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PermissionGroupEnvironmentAccessModel) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *PermissionGroupEnvironmentAccessModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PermissionGroupEnvironmentAccessModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetOrder

`func (o *PermissionGroupEnvironmentAccessModel) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *PermissionGroupEnvironmentAccessModel) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *PermissionGroupEnvironmentAccessModel) SetOrder(v int32)`

SetOrder sets Order field to given value.


### GetReasonRequired

`func (o *PermissionGroupEnvironmentAccessModel) GetReasonRequired() bool`

GetReasonRequired returns the ReasonRequired field if non-nil, zero value otherwise.

### GetReasonRequiredOk

`func (o *PermissionGroupEnvironmentAccessModel) GetReasonRequiredOk() (*bool, bool)`

GetReasonRequiredOk returns a tuple with the ReasonRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonRequired

`func (o *PermissionGroupEnvironmentAccessModel) SetReasonRequired(v bool)`

SetReasonRequired sets ReasonRequired field to given value.


### GetApproveRequired

`func (o *PermissionGroupEnvironmentAccessModel) GetApproveRequired() bool`

GetApproveRequired returns the ApproveRequired field if non-nil, zero value otherwise.

### GetApproveRequiredOk

`func (o *PermissionGroupEnvironmentAccessModel) GetApproveRequiredOk() (*bool, bool)`

GetApproveRequiredOk returns a tuple with the ApproveRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproveRequired

`func (o *PermissionGroupEnvironmentAccessModel) SetApproveRequired(v bool)`

SetApproveRequired sets ApproveRequired field to given value.


### GetEnvironmentAccessType

`func (o *PermissionGroupEnvironmentAccessModel) GetEnvironmentAccessType() EnvironmentAccessType`

GetEnvironmentAccessType returns the EnvironmentAccessType field if non-nil, zero value otherwise.

### GetEnvironmentAccessTypeOk

`func (o *PermissionGroupEnvironmentAccessModel) GetEnvironmentAccessTypeOk() (*EnvironmentAccessType, bool)`

GetEnvironmentAccessTypeOk returns a tuple with the EnvironmentAccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentAccessType

`func (o *PermissionGroupEnvironmentAccessModel) SetEnvironmentAccessType(v EnvironmentAccessType)`

SetEnvironmentAccessType sets EnvironmentAccessType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


