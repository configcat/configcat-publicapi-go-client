# ChangeRequestSummaryModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChangeRequestId** | **int64** | Identifier of the Change Request. | 
**ChangeRequestStatus** | [**ChangeRequestStatus**](ChangeRequestStatus.md) |  | 
**NeedsAttention** | **bool** | Indicates whether the Change Request needs attention. | 
**Title** | **string** | Title of the Change Request. | 
**Reason** | **NullableString** | Optional notes describing the purpose of the Change Request. | 
**ApplyAt** | **NullableTime** | Optional UTC date and time when the Change Request should be applied automatically. | 
**CreatedAt** | **time.Time** | The UTC date and time when the Change Request was created. | 
**CreatorUserEmail** | **string** | Email of the user who created the Change Request. | 
**CreatorUserFullName** | **string** | Full name of the user who created the Change Request. | 
**CreatorUserId** | **NullableString** | Identifier of the creator user. | 
**AffectedSettingKeys** | **[]string** | List of the keys of Settings affected by this Change Request. | 
**CommentCount** | **int32** | Number of comments on the Change Request. | 
**Approved** | **bool** | Indicates whether the Change Request has been approved. | 
**ConflictCount** | **int32** | Number of conflicting settings in the Change Request. | 
**BypassApproval** | **bool** | Indicates whether approval flow is bypassed. | 
**AppliedAt** | **NullableTime** | Optional UTC date and time when the Change Request was applied. | 
**AppliedByUserId** | **NullableString** | Identifier of the user who applied the Change Request. | 
**AppliedByUserEmail** | **NullableString** | Email of the user who applied the Change Request. | 
**AppliedByUserFullName** | **NullableString** | Full name of the user who applied the Change Request. | 
**ClosedAt** | **NullableTime** | Optional UTC date and time when the Change Request was closed. | 
**ClosedByUserId** | **NullableString** | Identifier of the user who closed the Change Request. | 
**ClosedByUserEmail** | **NullableString** | Email of the user who closed the Change Request. | 
**ClosedByUserFullName** | **NullableString** | Full name of the user who closed the Change Request. | 

## Methods

### NewChangeRequestSummaryModel

`func NewChangeRequestSummaryModel(changeRequestId int64, changeRequestStatus ChangeRequestStatus, needsAttention bool, title string, reason NullableString, applyAt NullableTime, createdAt time.Time, creatorUserEmail string, creatorUserFullName string, creatorUserId NullableString, affectedSettingKeys []string, commentCount int32, approved bool, conflictCount int32, bypassApproval bool, appliedAt NullableTime, appliedByUserId NullableString, appliedByUserEmail NullableString, appliedByUserFullName NullableString, closedAt NullableTime, closedByUserId NullableString, closedByUserEmail NullableString, closedByUserFullName NullableString, ) *ChangeRequestSummaryModel`

NewChangeRequestSummaryModel instantiates a new ChangeRequestSummaryModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangeRequestSummaryModelWithDefaults

`func NewChangeRequestSummaryModelWithDefaults() *ChangeRequestSummaryModel`

NewChangeRequestSummaryModelWithDefaults instantiates a new ChangeRequestSummaryModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChangeRequestId

`func (o *ChangeRequestSummaryModel) GetChangeRequestId() int64`

GetChangeRequestId returns the ChangeRequestId field if non-nil, zero value otherwise.

### GetChangeRequestIdOk

`func (o *ChangeRequestSummaryModel) GetChangeRequestIdOk() (*int64, bool)`

GetChangeRequestIdOk returns a tuple with the ChangeRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeRequestId

`func (o *ChangeRequestSummaryModel) SetChangeRequestId(v int64)`

SetChangeRequestId sets ChangeRequestId field to given value.


### GetChangeRequestStatus

`func (o *ChangeRequestSummaryModel) GetChangeRequestStatus() ChangeRequestStatus`

GetChangeRequestStatus returns the ChangeRequestStatus field if non-nil, zero value otherwise.

### GetChangeRequestStatusOk

`func (o *ChangeRequestSummaryModel) GetChangeRequestStatusOk() (*ChangeRequestStatus, bool)`

GetChangeRequestStatusOk returns a tuple with the ChangeRequestStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeRequestStatus

`func (o *ChangeRequestSummaryModel) SetChangeRequestStatus(v ChangeRequestStatus)`

SetChangeRequestStatus sets ChangeRequestStatus field to given value.


### GetNeedsAttention

`func (o *ChangeRequestSummaryModel) GetNeedsAttention() bool`

GetNeedsAttention returns the NeedsAttention field if non-nil, zero value otherwise.

### GetNeedsAttentionOk

`func (o *ChangeRequestSummaryModel) GetNeedsAttentionOk() (*bool, bool)`

GetNeedsAttentionOk returns a tuple with the NeedsAttention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedsAttention

`func (o *ChangeRequestSummaryModel) SetNeedsAttention(v bool)`

SetNeedsAttention sets NeedsAttention field to given value.


### GetTitle

`func (o *ChangeRequestSummaryModel) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ChangeRequestSummaryModel) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ChangeRequestSummaryModel) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetReason

`func (o *ChangeRequestSummaryModel) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ChangeRequestSummaryModel) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ChangeRequestSummaryModel) SetReason(v string)`

SetReason sets Reason field to given value.


### SetReasonNil

`func (o *ChangeRequestSummaryModel) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *ChangeRequestSummaryModel) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetApplyAt

`func (o *ChangeRequestSummaryModel) GetApplyAt() time.Time`

GetApplyAt returns the ApplyAt field if non-nil, zero value otherwise.

### GetApplyAtOk

`func (o *ChangeRequestSummaryModel) GetApplyAtOk() (*time.Time, bool)`

GetApplyAtOk returns a tuple with the ApplyAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyAt

`func (o *ChangeRequestSummaryModel) SetApplyAt(v time.Time)`

SetApplyAt sets ApplyAt field to given value.


### SetApplyAtNil

`func (o *ChangeRequestSummaryModel) SetApplyAtNil(b bool)`

 SetApplyAtNil sets the value for ApplyAt to be an explicit nil

### UnsetApplyAt
`func (o *ChangeRequestSummaryModel) UnsetApplyAt()`

UnsetApplyAt ensures that no value is present for ApplyAt, not even an explicit nil
### GetCreatedAt

`func (o *ChangeRequestSummaryModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ChangeRequestSummaryModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ChangeRequestSummaryModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatorUserEmail

`func (o *ChangeRequestSummaryModel) GetCreatorUserEmail() string`

GetCreatorUserEmail returns the CreatorUserEmail field if non-nil, zero value otherwise.

### GetCreatorUserEmailOk

`func (o *ChangeRequestSummaryModel) GetCreatorUserEmailOk() (*string, bool)`

GetCreatorUserEmailOk returns a tuple with the CreatorUserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorUserEmail

`func (o *ChangeRequestSummaryModel) SetCreatorUserEmail(v string)`

SetCreatorUserEmail sets CreatorUserEmail field to given value.


### GetCreatorUserFullName

`func (o *ChangeRequestSummaryModel) GetCreatorUserFullName() string`

GetCreatorUserFullName returns the CreatorUserFullName field if non-nil, zero value otherwise.

### GetCreatorUserFullNameOk

`func (o *ChangeRequestSummaryModel) GetCreatorUserFullNameOk() (*string, bool)`

GetCreatorUserFullNameOk returns a tuple with the CreatorUserFullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorUserFullName

`func (o *ChangeRequestSummaryModel) SetCreatorUserFullName(v string)`

SetCreatorUserFullName sets CreatorUserFullName field to given value.


### GetCreatorUserId

`func (o *ChangeRequestSummaryModel) GetCreatorUserId() string`

GetCreatorUserId returns the CreatorUserId field if non-nil, zero value otherwise.

### GetCreatorUserIdOk

`func (o *ChangeRequestSummaryModel) GetCreatorUserIdOk() (*string, bool)`

GetCreatorUserIdOk returns a tuple with the CreatorUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorUserId

`func (o *ChangeRequestSummaryModel) SetCreatorUserId(v string)`

SetCreatorUserId sets CreatorUserId field to given value.


### SetCreatorUserIdNil

`func (o *ChangeRequestSummaryModel) SetCreatorUserIdNil(b bool)`

 SetCreatorUserIdNil sets the value for CreatorUserId to be an explicit nil

### UnsetCreatorUserId
`func (o *ChangeRequestSummaryModel) UnsetCreatorUserId()`

UnsetCreatorUserId ensures that no value is present for CreatorUserId, not even an explicit nil
### GetAffectedSettingKeys

`func (o *ChangeRequestSummaryModel) GetAffectedSettingKeys() []string`

GetAffectedSettingKeys returns the AffectedSettingKeys field if non-nil, zero value otherwise.

### GetAffectedSettingKeysOk

`func (o *ChangeRequestSummaryModel) GetAffectedSettingKeysOk() (*[]string, bool)`

GetAffectedSettingKeysOk returns a tuple with the AffectedSettingKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedSettingKeys

`func (o *ChangeRequestSummaryModel) SetAffectedSettingKeys(v []string)`

SetAffectedSettingKeys sets AffectedSettingKeys field to given value.


### GetCommentCount

`func (o *ChangeRequestSummaryModel) GetCommentCount() int32`

GetCommentCount returns the CommentCount field if non-nil, zero value otherwise.

### GetCommentCountOk

`func (o *ChangeRequestSummaryModel) GetCommentCountOk() (*int32, bool)`

GetCommentCountOk returns a tuple with the CommentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentCount

`func (o *ChangeRequestSummaryModel) SetCommentCount(v int32)`

SetCommentCount sets CommentCount field to given value.


### GetApproved

`func (o *ChangeRequestSummaryModel) GetApproved() bool`

GetApproved returns the Approved field if non-nil, zero value otherwise.

### GetApprovedOk

`func (o *ChangeRequestSummaryModel) GetApprovedOk() (*bool, bool)`

GetApprovedOk returns a tuple with the Approved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproved

`func (o *ChangeRequestSummaryModel) SetApproved(v bool)`

SetApproved sets Approved field to given value.


### GetConflictCount

`func (o *ChangeRequestSummaryModel) GetConflictCount() int32`

GetConflictCount returns the ConflictCount field if non-nil, zero value otherwise.

### GetConflictCountOk

`func (o *ChangeRequestSummaryModel) GetConflictCountOk() (*int32, bool)`

GetConflictCountOk returns a tuple with the ConflictCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflictCount

`func (o *ChangeRequestSummaryModel) SetConflictCount(v int32)`

SetConflictCount sets ConflictCount field to given value.


### GetBypassApproval

`func (o *ChangeRequestSummaryModel) GetBypassApproval() bool`

GetBypassApproval returns the BypassApproval field if non-nil, zero value otherwise.

### GetBypassApprovalOk

`func (o *ChangeRequestSummaryModel) GetBypassApprovalOk() (*bool, bool)`

GetBypassApprovalOk returns a tuple with the BypassApproval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBypassApproval

`func (o *ChangeRequestSummaryModel) SetBypassApproval(v bool)`

SetBypassApproval sets BypassApproval field to given value.


### GetAppliedAt

`func (o *ChangeRequestSummaryModel) GetAppliedAt() time.Time`

GetAppliedAt returns the AppliedAt field if non-nil, zero value otherwise.

### GetAppliedAtOk

`func (o *ChangeRequestSummaryModel) GetAppliedAtOk() (*time.Time, bool)`

GetAppliedAtOk returns a tuple with the AppliedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedAt

`func (o *ChangeRequestSummaryModel) SetAppliedAt(v time.Time)`

SetAppliedAt sets AppliedAt field to given value.


### SetAppliedAtNil

`func (o *ChangeRequestSummaryModel) SetAppliedAtNil(b bool)`

 SetAppliedAtNil sets the value for AppliedAt to be an explicit nil

### UnsetAppliedAt
`func (o *ChangeRequestSummaryModel) UnsetAppliedAt()`

UnsetAppliedAt ensures that no value is present for AppliedAt, not even an explicit nil
### GetAppliedByUserId

`func (o *ChangeRequestSummaryModel) GetAppliedByUserId() string`

GetAppliedByUserId returns the AppliedByUserId field if non-nil, zero value otherwise.

### GetAppliedByUserIdOk

`func (o *ChangeRequestSummaryModel) GetAppliedByUserIdOk() (*string, bool)`

GetAppliedByUserIdOk returns a tuple with the AppliedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedByUserId

`func (o *ChangeRequestSummaryModel) SetAppliedByUserId(v string)`

SetAppliedByUserId sets AppliedByUserId field to given value.


### SetAppliedByUserIdNil

`func (o *ChangeRequestSummaryModel) SetAppliedByUserIdNil(b bool)`

 SetAppliedByUserIdNil sets the value for AppliedByUserId to be an explicit nil

### UnsetAppliedByUserId
`func (o *ChangeRequestSummaryModel) UnsetAppliedByUserId()`

UnsetAppliedByUserId ensures that no value is present for AppliedByUserId, not even an explicit nil
### GetAppliedByUserEmail

`func (o *ChangeRequestSummaryModel) GetAppliedByUserEmail() string`

GetAppliedByUserEmail returns the AppliedByUserEmail field if non-nil, zero value otherwise.

### GetAppliedByUserEmailOk

`func (o *ChangeRequestSummaryModel) GetAppliedByUserEmailOk() (*string, bool)`

GetAppliedByUserEmailOk returns a tuple with the AppliedByUserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedByUserEmail

`func (o *ChangeRequestSummaryModel) SetAppliedByUserEmail(v string)`

SetAppliedByUserEmail sets AppliedByUserEmail field to given value.


### SetAppliedByUserEmailNil

`func (o *ChangeRequestSummaryModel) SetAppliedByUserEmailNil(b bool)`

 SetAppliedByUserEmailNil sets the value for AppliedByUserEmail to be an explicit nil

### UnsetAppliedByUserEmail
`func (o *ChangeRequestSummaryModel) UnsetAppliedByUserEmail()`

UnsetAppliedByUserEmail ensures that no value is present for AppliedByUserEmail, not even an explicit nil
### GetAppliedByUserFullName

`func (o *ChangeRequestSummaryModel) GetAppliedByUserFullName() string`

GetAppliedByUserFullName returns the AppliedByUserFullName field if non-nil, zero value otherwise.

### GetAppliedByUserFullNameOk

`func (o *ChangeRequestSummaryModel) GetAppliedByUserFullNameOk() (*string, bool)`

GetAppliedByUserFullNameOk returns a tuple with the AppliedByUserFullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedByUserFullName

`func (o *ChangeRequestSummaryModel) SetAppliedByUserFullName(v string)`

SetAppliedByUserFullName sets AppliedByUserFullName field to given value.


### SetAppliedByUserFullNameNil

`func (o *ChangeRequestSummaryModel) SetAppliedByUserFullNameNil(b bool)`

 SetAppliedByUserFullNameNil sets the value for AppliedByUserFullName to be an explicit nil

### UnsetAppliedByUserFullName
`func (o *ChangeRequestSummaryModel) UnsetAppliedByUserFullName()`

UnsetAppliedByUserFullName ensures that no value is present for AppliedByUserFullName, not even an explicit nil
### GetClosedAt

`func (o *ChangeRequestSummaryModel) GetClosedAt() time.Time`

GetClosedAt returns the ClosedAt field if non-nil, zero value otherwise.

### GetClosedAtOk

`func (o *ChangeRequestSummaryModel) GetClosedAtOk() (*time.Time, bool)`

GetClosedAtOk returns a tuple with the ClosedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedAt

`func (o *ChangeRequestSummaryModel) SetClosedAt(v time.Time)`

SetClosedAt sets ClosedAt field to given value.


### SetClosedAtNil

`func (o *ChangeRequestSummaryModel) SetClosedAtNil(b bool)`

 SetClosedAtNil sets the value for ClosedAt to be an explicit nil

### UnsetClosedAt
`func (o *ChangeRequestSummaryModel) UnsetClosedAt()`

UnsetClosedAt ensures that no value is present for ClosedAt, not even an explicit nil
### GetClosedByUserId

`func (o *ChangeRequestSummaryModel) GetClosedByUserId() string`

GetClosedByUserId returns the ClosedByUserId field if non-nil, zero value otherwise.

### GetClosedByUserIdOk

`func (o *ChangeRequestSummaryModel) GetClosedByUserIdOk() (*string, bool)`

GetClosedByUserIdOk returns a tuple with the ClosedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedByUserId

`func (o *ChangeRequestSummaryModel) SetClosedByUserId(v string)`

SetClosedByUserId sets ClosedByUserId field to given value.


### SetClosedByUserIdNil

`func (o *ChangeRequestSummaryModel) SetClosedByUserIdNil(b bool)`

 SetClosedByUserIdNil sets the value for ClosedByUserId to be an explicit nil

### UnsetClosedByUserId
`func (o *ChangeRequestSummaryModel) UnsetClosedByUserId()`

UnsetClosedByUserId ensures that no value is present for ClosedByUserId, not even an explicit nil
### GetClosedByUserEmail

`func (o *ChangeRequestSummaryModel) GetClosedByUserEmail() string`

GetClosedByUserEmail returns the ClosedByUserEmail field if non-nil, zero value otherwise.

### GetClosedByUserEmailOk

`func (o *ChangeRequestSummaryModel) GetClosedByUserEmailOk() (*string, bool)`

GetClosedByUserEmailOk returns a tuple with the ClosedByUserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedByUserEmail

`func (o *ChangeRequestSummaryModel) SetClosedByUserEmail(v string)`

SetClosedByUserEmail sets ClosedByUserEmail field to given value.


### SetClosedByUserEmailNil

`func (o *ChangeRequestSummaryModel) SetClosedByUserEmailNil(b bool)`

 SetClosedByUserEmailNil sets the value for ClosedByUserEmail to be an explicit nil

### UnsetClosedByUserEmail
`func (o *ChangeRequestSummaryModel) UnsetClosedByUserEmail()`

UnsetClosedByUserEmail ensures that no value is present for ClosedByUserEmail, not even an explicit nil
### GetClosedByUserFullName

`func (o *ChangeRequestSummaryModel) GetClosedByUserFullName() string`

GetClosedByUserFullName returns the ClosedByUserFullName field if non-nil, zero value otherwise.

### GetClosedByUserFullNameOk

`func (o *ChangeRequestSummaryModel) GetClosedByUserFullNameOk() (*string, bool)`

GetClosedByUserFullNameOk returns a tuple with the ClosedByUserFullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedByUserFullName

`func (o *ChangeRequestSummaryModel) SetClosedByUserFullName(v string)`

SetClosedByUserFullName sets ClosedByUserFullName field to given value.


### SetClosedByUserFullNameNil

`func (o *ChangeRequestSummaryModel) SetClosedByUserFullNameNil(b bool)`

 SetClosedByUserFullNameNil sets the value for ClosedByUserFullName to be an explicit nil

### UnsetClosedByUserFullName
`func (o *ChangeRequestSummaryModel) UnsetClosedByUserFullName()`

UnsetClosedByUserFullName ensures that no value is present for ClosedByUserFullName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


