# CreateChangeRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** | The title of the Change Request. | 
**Reason** | Pointer to **NullableString** | The optional notes describing the purpose of the Change Request. This will appear in the Audit Log (in the Notes section when you expand the corresponding entry) upon applying the change request. | [optional] 
**ApplyAt** | Pointer to **NullableTime** | The optional UTC date and time when the scheduled Change Request should be applied automatically. | [optional] 
**BypassApproval** | Pointer to **bool** | When true, bypasses required approval checks for scheduled changes. | [optional] 
**ProposedChanges** | Pointer to [**[]CreateChangeRequestProposedChangeModel**](CreateChangeRequestProposedChangeModel.md) | The list of models describing the proposed changes to the Settings included in the new Change Request. | [optional] 

## Methods

### NewCreateChangeRequestModel

`func NewCreateChangeRequestModel(title string, ) *CreateChangeRequestModel`

NewCreateChangeRequestModel instantiates a new CreateChangeRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateChangeRequestModelWithDefaults

`func NewCreateChangeRequestModelWithDefaults() *CreateChangeRequestModel`

NewCreateChangeRequestModelWithDefaults instantiates a new CreateChangeRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *CreateChangeRequestModel) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateChangeRequestModel) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateChangeRequestModel) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetReason

`func (o *CreateChangeRequestModel) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CreateChangeRequestModel) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CreateChangeRequestModel) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *CreateChangeRequestModel) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *CreateChangeRequestModel) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *CreateChangeRequestModel) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetApplyAt

`func (o *CreateChangeRequestModel) GetApplyAt() time.Time`

GetApplyAt returns the ApplyAt field if non-nil, zero value otherwise.

### GetApplyAtOk

`func (o *CreateChangeRequestModel) GetApplyAtOk() (*time.Time, bool)`

GetApplyAtOk returns a tuple with the ApplyAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyAt

`func (o *CreateChangeRequestModel) SetApplyAt(v time.Time)`

SetApplyAt sets ApplyAt field to given value.

### HasApplyAt

`func (o *CreateChangeRequestModel) HasApplyAt() bool`

HasApplyAt returns a boolean if a field has been set.

### SetApplyAtNil

`func (o *CreateChangeRequestModel) SetApplyAtNil(b bool)`

 SetApplyAtNil sets the value for ApplyAt to be an explicit nil

### UnsetApplyAt
`func (o *CreateChangeRequestModel) UnsetApplyAt()`

UnsetApplyAt ensures that no value is present for ApplyAt, not even an explicit nil
### GetBypassApproval

`func (o *CreateChangeRequestModel) GetBypassApproval() bool`

GetBypassApproval returns the BypassApproval field if non-nil, zero value otherwise.

### GetBypassApprovalOk

`func (o *CreateChangeRequestModel) GetBypassApprovalOk() (*bool, bool)`

GetBypassApprovalOk returns a tuple with the BypassApproval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBypassApproval

`func (o *CreateChangeRequestModel) SetBypassApproval(v bool)`

SetBypassApproval sets BypassApproval field to given value.

### HasBypassApproval

`func (o *CreateChangeRequestModel) HasBypassApproval() bool`

HasBypassApproval returns a boolean if a field has been set.

### GetProposedChanges

`func (o *CreateChangeRequestModel) GetProposedChanges() []CreateChangeRequestProposedChangeModel`

GetProposedChanges returns the ProposedChanges field if non-nil, zero value otherwise.

### GetProposedChangesOk

`func (o *CreateChangeRequestModel) GetProposedChangesOk() (*[]CreateChangeRequestProposedChangeModel, bool)`

GetProposedChangesOk returns a tuple with the ProposedChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProposedChanges

`func (o *CreateChangeRequestModel) SetProposedChanges(v []CreateChangeRequestProposedChangeModel)`

SetProposedChanges sets ProposedChanges field to given value.

### HasProposedChanges

`func (o *CreateChangeRequestModel) HasProposedChanges() bool`

HasProposedChanges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


