# UpdateChangeRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** | The updated title of the Change Request. | 
**Reason** | Pointer to **NullableString** | The updated optional notes describing the purpose of the Change Request. | [optional] 
**ApplyAt** | Pointer to **NullableTime** | The updated optional UTC date and time when the Change Request should be applied automatically. | [optional] 
**BypassApproval** | Pointer to **NullableBool** | The updated bypass-approval flag for scheduled changes. | [optional] 

## Methods

### NewUpdateChangeRequestModel

`func NewUpdateChangeRequestModel(title string, ) *UpdateChangeRequestModel`

NewUpdateChangeRequestModel instantiates a new UpdateChangeRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateChangeRequestModelWithDefaults

`func NewUpdateChangeRequestModelWithDefaults() *UpdateChangeRequestModel`

NewUpdateChangeRequestModelWithDefaults instantiates a new UpdateChangeRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *UpdateChangeRequestModel) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UpdateChangeRequestModel) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UpdateChangeRequestModel) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetReason

`func (o *UpdateChangeRequestModel) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *UpdateChangeRequestModel) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *UpdateChangeRequestModel) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *UpdateChangeRequestModel) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *UpdateChangeRequestModel) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *UpdateChangeRequestModel) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetApplyAt

`func (o *UpdateChangeRequestModel) GetApplyAt() time.Time`

GetApplyAt returns the ApplyAt field if non-nil, zero value otherwise.

### GetApplyAtOk

`func (o *UpdateChangeRequestModel) GetApplyAtOk() (*time.Time, bool)`

GetApplyAtOk returns a tuple with the ApplyAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyAt

`func (o *UpdateChangeRequestModel) SetApplyAt(v time.Time)`

SetApplyAt sets ApplyAt field to given value.

### HasApplyAt

`func (o *UpdateChangeRequestModel) HasApplyAt() bool`

HasApplyAt returns a boolean if a field has been set.

### SetApplyAtNil

`func (o *UpdateChangeRequestModel) SetApplyAtNil(b bool)`

 SetApplyAtNil sets the value for ApplyAt to be an explicit nil

### UnsetApplyAt
`func (o *UpdateChangeRequestModel) UnsetApplyAt()`

UnsetApplyAt ensures that no value is present for ApplyAt, not even an explicit nil
### GetBypassApproval

`func (o *UpdateChangeRequestModel) GetBypassApproval() bool`

GetBypassApproval returns the BypassApproval field if non-nil, zero value otherwise.

### GetBypassApprovalOk

`func (o *UpdateChangeRequestModel) GetBypassApprovalOk() (*bool, bool)`

GetBypassApprovalOk returns a tuple with the BypassApproval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBypassApproval

`func (o *UpdateChangeRequestModel) SetBypassApproval(v bool)`

SetBypassApproval sets BypassApproval field to given value.

### HasBypassApproval

`func (o *UpdateChangeRequestModel) HasBypassApproval() bool`

HasBypassApproval returns a boolean if a field has been set.

### SetBypassApprovalNil

`func (o *UpdateChangeRequestModel) SetBypassApprovalNil(b bool)`

 SetBypassApprovalNil sets the value for BypassApproval to be an explicit nil

### UnsetBypassApproval
`func (o *UpdateChangeRequestModel) UnsetBypassApproval()`

UnsetBypassApproval ensures that no value is present for BypassApproval, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


