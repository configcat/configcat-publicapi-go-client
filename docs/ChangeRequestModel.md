# ChangeRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChangeRequestId** | **int64** | Identifier of the Change Request. | 
**ConfigId** | **string** | Identifier of the Config. | 
**EnvironmentId** | **string** | Identifier of the Environment. | 
**ChangeRequestStatus** | [**ChangeRequestStatus**](ChangeRequestStatus.md) |  | 
**NeedsAttention** | **bool** | Indicates whether the Change Request needs attention. | 
**Title** | **string** | Title of the Change Request. | 
**Reason** | **NullableString** | Optional notes describing the purpose of the Change Request. | 
**ApplyAt** | **NullableTime** | Optional UTC date and time when the Change Request should be applied automatically. | 
**CreatedAt** | **time.Time** | The UTC date and time when the Change Request was created. | 
**CreatorUserEmail** | **string** | Email of the user who created the Change Request. | 
**CreatorUserFullName** | **string** | Full name of the user who created the Change Request. | 
**CreatorUserId** | **NullableString** | Identifier of the creator user. | 
**SettingValues** | [**[]ChangeRequestProposedChangeModel**](ChangeRequestProposedChangeModel.md) | List of proposed setting values in the Change Request. | 
**Comments** | [**[]ChangeRequestCommentModel**](ChangeRequestCommentModel.md) | List of comments on the Change Request. | 
**Approved** | **bool** | Indicates whether the Change Request has been approved. | 
**Approvals** | [**[]ChangeRequestApprovalModel**](ChangeRequestApprovalModel.md) | List of approvals for the Change Request. | 
**Activities** | [**[]ChangeRequestActivityModel**](ChangeRequestActivityModel.md) | List of activities (history) on the Change Request. | 
**ChangeRequestIssues** | [**[]ChangeRequestIssueModel**](ChangeRequestIssueModel.md) | List of issues encountered with the Change Request. | 
**AppliedAt** | **NullableTime** | Optional UTC date and time when the Change Request was applied. | 
**AppliedByUserId** | **NullableString** | Identifier of the user who applied the Change Request. | 
**AppliedByUserEmail** | **NullableString** | Email of the user who applied the Change Request. | 
**AppliedByUserFullName** | **NullableString** | Full name of the user who applied the Change Request. | 
**ClosedAt** | **NullableTime** | Optional UTC date and time when the Change Request was closed. | 
**ClosedByUserId** | **NullableString** | Identifier of the user who closed the Change Request. | 
**ClosedByUserEmail** | **NullableString** | Email of the user who closed the Change Request. | 
**ClosedByUserFullName** | **NullableString** | Full name of the user who closed the Change Request. | 
**BypassApproval** | **bool** | Indicates whether approval flow is bypassed. | 

## Methods

### NewChangeRequestModel

`func NewChangeRequestModel(changeRequestId int64, configId string, environmentId string, changeRequestStatus ChangeRequestStatus, needsAttention bool, title string, reason NullableString, applyAt NullableTime, createdAt time.Time, creatorUserEmail string, creatorUserFullName string, creatorUserId NullableString, settingValues []ChangeRequestProposedChangeModel, comments []ChangeRequestCommentModel, approved bool, approvals []ChangeRequestApprovalModel, activities []ChangeRequestActivityModel, changeRequestIssues []ChangeRequestIssueModel, appliedAt NullableTime, appliedByUserId NullableString, appliedByUserEmail NullableString, appliedByUserFullName NullableString, closedAt NullableTime, closedByUserId NullableString, closedByUserEmail NullableString, closedByUserFullName NullableString, bypassApproval bool, ) *ChangeRequestModel`

NewChangeRequestModel instantiates a new ChangeRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangeRequestModelWithDefaults

`func NewChangeRequestModelWithDefaults() *ChangeRequestModel`

NewChangeRequestModelWithDefaults instantiates a new ChangeRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChangeRequestId

`func (o *ChangeRequestModel) GetChangeRequestId() int64`

GetChangeRequestId returns the ChangeRequestId field if non-nil, zero value otherwise.

### GetChangeRequestIdOk

`func (o *ChangeRequestModel) GetChangeRequestIdOk() (*int64, bool)`

GetChangeRequestIdOk returns a tuple with the ChangeRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeRequestId

`func (o *ChangeRequestModel) SetChangeRequestId(v int64)`

SetChangeRequestId sets ChangeRequestId field to given value.


### GetConfigId

`func (o *ChangeRequestModel) GetConfigId() string`

GetConfigId returns the ConfigId field if non-nil, zero value otherwise.

### GetConfigIdOk

`func (o *ChangeRequestModel) GetConfigIdOk() (*string, bool)`

GetConfigIdOk returns a tuple with the ConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigId

`func (o *ChangeRequestModel) SetConfigId(v string)`

SetConfigId sets ConfigId field to given value.


### GetEnvironmentId

`func (o *ChangeRequestModel) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *ChangeRequestModel) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *ChangeRequestModel) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetChangeRequestStatus

`func (o *ChangeRequestModel) GetChangeRequestStatus() ChangeRequestStatus`

GetChangeRequestStatus returns the ChangeRequestStatus field if non-nil, zero value otherwise.

### GetChangeRequestStatusOk

`func (o *ChangeRequestModel) GetChangeRequestStatusOk() (*ChangeRequestStatus, bool)`

GetChangeRequestStatusOk returns a tuple with the ChangeRequestStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeRequestStatus

`func (o *ChangeRequestModel) SetChangeRequestStatus(v ChangeRequestStatus)`

SetChangeRequestStatus sets ChangeRequestStatus field to given value.


### GetNeedsAttention

`func (o *ChangeRequestModel) GetNeedsAttention() bool`

GetNeedsAttention returns the NeedsAttention field if non-nil, zero value otherwise.

### GetNeedsAttentionOk

`func (o *ChangeRequestModel) GetNeedsAttentionOk() (*bool, bool)`

GetNeedsAttentionOk returns a tuple with the NeedsAttention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedsAttention

`func (o *ChangeRequestModel) SetNeedsAttention(v bool)`

SetNeedsAttention sets NeedsAttention field to given value.


### GetTitle

`func (o *ChangeRequestModel) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ChangeRequestModel) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ChangeRequestModel) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetReason

`func (o *ChangeRequestModel) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ChangeRequestModel) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ChangeRequestModel) SetReason(v string)`

SetReason sets Reason field to given value.


### SetReasonNil

`func (o *ChangeRequestModel) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *ChangeRequestModel) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetApplyAt

`func (o *ChangeRequestModel) GetApplyAt() time.Time`

GetApplyAt returns the ApplyAt field if non-nil, zero value otherwise.

### GetApplyAtOk

`func (o *ChangeRequestModel) GetApplyAtOk() (*time.Time, bool)`

GetApplyAtOk returns a tuple with the ApplyAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyAt

`func (o *ChangeRequestModel) SetApplyAt(v time.Time)`

SetApplyAt sets ApplyAt field to given value.


### SetApplyAtNil

`func (o *ChangeRequestModel) SetApplyAtNil(b bool)`

 SetApplyAtNil sets the value for ApplyAt to be an explicit nil

### UnsetApplyAt
`func (o *ChangeRequestModel) UnsetApplyAt()`

UnsetApplyAt ensures that no value is present for ApplyAt, not even an explicit nil
### GetCreatedAt

`func (o *ChangeRequestModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ChangeRequestModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ChangeRequestModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatorUserEmail

`func (o *ChangeRequestModel) GetCreatorUserEmail() string`

GetCreatorUserEmail returns the CreatorUserEmail field if non-nil, zero value otherwise.

### GetCreatorUserEmailOk

`func (o *ChangeRequestModel) GetCreatorUserEmailOk() (*string, bool)`

GetCreatorUserEmailOk returns a tuple with the CreatorUserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorUserEmail

`func (o *ChangeRequestModel) SetCreatorUserEmail(v string)`

SetCreatorUserEmail sets CreatorUserEmail field to given value.


### GetCreatorUserFullName

`func (o *ChangeRequestModel) GetCreatorUserFullName() string`

GetCreatorUserFullName returns the CreatorUserFullName field if non-nil, zero value otherwise.

### GetCreatorUserFullNameOk

`func (o *ChangeRequestModel) GetCreatorUserFullNameOk() (*string, bool)`

GetCreatorUserFullNameOk returns a tuple with the CreatorUserFullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorUserFullName

`func (o *ChangeRequestModel) SetCreatorUserFullName(v string)`

SetCreatorUserFullName sets CreatorUserFullName field to given value.


### GetCreatorUserId

`func (o *ChangeRequestModel) GetCreatorUserId() string`

GetCreatorUserId returns the CreatorUserId field if non-nil, zero value otherwise.

### GetCreatorUserIdOk

`func (o *ChangeRequestModel) GetCreatorUserIdOk() (*string, bool)`

GetCreatorUserIdOk returns a tuple with the CreatorUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorUserId

`func (o *ChangeRequestModel) SetCreatorUserId(v string)`

SetCreatorUserId sets CreatorUserId field to given value.


### SetCreatorUserIdNil

`func (o *ChangeRequestModel) SetCreatorUserIdNil(b bool)`

 SetCreatorUserIdNil sets the value for CreatorUserId to be an explicit nil

### UnsetCreatorUserId
`func (o *ChangeRequestModel) UnsetCreatorUserId()`

UnsetCreatorUserId ensures that no value is present for CreatorUserId, not even an explicit nil
### GetSettingValues

`func (o *ChangeRequestModel) GetSettingValues() []ChangeRequestProposedChangeModel`

GetSettingValues returns the SettingValues field if non-nil, zero value otherwise.

### GetSettingValuesOk

`func (o *ChangeRequestModel) GetSettingValuesOk() (*[]ChangeRequestProposedChangeModel, bool)`

GetSettingValuesOk returns a tuple with the SettingValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingValues

`func (o *ChangeRequestModel) SetSettingValues(v []ChangeRequestProposedChangeModel)`

SetSettingValues sets SettingValues field to given value.


### GetComments

`func (o *ChangeRequestModel) GetComments() []ChangeRequestCommentModel`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *ChangeRequestModel) GetCommentsOk() (*[]ChangeRequestCommentModel, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *ChangeRequestModel) SetComments(v []ChangeRequestCommentModel)`

SetComments sets Comments field to given value.


### GetApproved

`func (o *ChangeRequestModel) GetApproved() bool`

GetApproved returns the Approved field if non-nil, zero value otherwise.

### GetApprovedOk

`func (o *ChangeRequestModel) GetApprovedOk() (*bool, bool)`

GetApprovedOk returns a tuple with the Approved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproved

`func (o *ChangeRequestModel) SetApproved(v bool)`

SetApproved sets Approved field to given value.


### GetApprovals

`func (o *ChangeRequestModel) GetApprovals() []ChangeRequestApprovalModel`

GetApprovals returns the Approvals field if non-nil, zero value otherwise.

### GetApprovalsOk

`func (o *ChangeRequestModel) GetApprovalsOk() (*[]ChangeRequestApprovalModel, bool)`

GetApprovalsOk returns a tuple with the Approvals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovals

`func (o *ChangeRequestModel) SetApprovals(v []ChangeRequestApprovalModel)`

SetApprovals sets Approvals field to given value.


### GetActivities

`func (o *ChangeRequestModel) GetActivities() []ChangeRequestActivityModel`

GetActivities returns the Activities field if non-nil, zero value otherwise.

### GetActivitiesOk

`func (o *ChangeRequestModel) GetActivitiesOk() (*[]ChangeRequestActivityModel, bool)`

GetActivitiesOk returns a tuple with the Activities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivities

`func (o *ChangeRequestModel) SetActivities(v []ChangeRequestActivityModel)`

SetActivities sets Activities field to given value.


### GetChangeRequestIssues

`func (o *ChangeRequestModel) GetChangeRequestIssues() []ChangeRequestIssueModel`

GetChangeRequestIssues returns the ChangeRequestIssues field if non-nil, zero value otherwise.

### GetChangeRequestIssuesOk

`func (o *ChangeRequestModel) GetChangeRequestIssuesOk() (*[]ChangeRequestIssueModel, bool)`

GetChangeRequestIssuesOk returns a tuple with the ChangeRequestIssues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeRequestIssues

`func (o *ChangeRequestModel) SetChangeRequestIssues(v []ChangeRequestIssueModel)`

SetChangeRequestIssues sets ChangeRequestIssues field to given value.


### GetAppliedAt

`func (o *ChangeRequestModel) GetAppliedAt() time.Time`

GetAppliedAt returns the AppliedAt field if non-nil, zero value otherwise.

### GetAppliedAtOk

`func (o *ChangeRequestModel) GetAppliedAtOk() (*time.Time, bool)`

GetAppliedAtOk returns a tuple with the AppliedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedAt

`func (o *ChangeRequestModel) SetAppliedAt(v time.Time)`

SetAppliedAt sets AppliedAt field to given value.


### SetAppliedAtNil

`func (o *ChangeRequestModel) SetAppliedAtNil(b bool)`

 SetAppliedAtNil sets the value for AppliedAt to be an explicit nil

### UnsetAppliedAt
`func (o *ChangeRequestModel) UnsetAppliedAt()`

UnsetAppliedAt ensures that no value is present for AppliedAt, not even an explicit nil
### GetAppliedByUserId

`func (o *ChangeRequestModel) GetAppliedByUserId() string`

GetAppliedByUserId returns the AppliedByUserId field if non-nil, zero value otherwise.

### GetAppliedByUserIdOk

`func (o *ChangeRequestModel) GetAppliedByUserIdOk() (*string, bool)`

GetAppliedByUserIdOk returns a tuple with the AppliedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedByUserId

`func (o *ChangeRequestModel) SetAppliedByUserId(v string)`

SetAppliedByUserId sets AppliedByUserId field to given value.


### SetAppliedByUserIdNil

`func (o *ChangeRequestModel) SetAppliedByUserIdNil(b bool)`

 SetAppliedByUserIdNil sets the value for AppliedByUserId to be an explicit nil

### UnsetAppliedByUserId
`func (o *ChangeRequestModel) UnsetAppliedByUserId()`

UnsetAppliedByUserId ensures that no value is present for AppliedByUserId, not even an explicit nil
### GetAppliedByUserEmail

`func (o *ChangeRequestModel) GetAppliedByUserEmail() string`

GetAppliedByUserEmail returns the AppliedByUserEmail field if non-nil, zero value otherwise.

### GetAppliedByUserEmailOk

`func (o *ChangeRequestModel) GetAppliedByUserEmailOk() (*string, bool)`

GetAppliedByUserEmailOk returns a tuple with the AppliedByUserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedByUserEmail

`func (o *ChangeRequestModel) SetAppliedByUserEmail(v string)`

SetAppliedByUserEmail sets AppliedByUserEmail field to given value.


### SetAppliedByUserEmailNil

`func (o *ChangeRequestModel) SetAppliedByUserEmailNil(b bool)`

 SetAppliedByUserEmailNil sets the value for AppliedByUserEmail to be an explicit nil

### UnsetAppliedByUserEmail
`func (o *ChangeRequestModel) UnsetAppliedByUserEmail()`

UnsetAppliedByUserEmail ensures that no value is present for AppliedByUserEmail, not even an explicit nil
### GetAppliedByUserFullName

`func (o *ChangeRequestModel) GetAppliedByUserFullName() string`

GetAppliedByUserFullName returns the AppliedByUserFullName field if non-nil, zero value otherwise.

### GetAppliedByUserFullNameOk

`func (o *ChangeRequestModel) GetAppliedByUserFullNameOk() (*string, bool)`

GetAppliedByUserFullNameOk returns a tuple with the AppliedByUserFullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedByUserFullName

`func (o *ChangeRequestModel) SetAppliedByUserFullName(v string)`

SetAppliedByUserFullName sets AppliedByUserFullName field to given value.


### SetAppliedByUserFullNameNil

`func (o *ChangeRequestModel) SetAppliedByUserFullNameNil(b bool)`

 SetAppliedByUserFullNameNil sets the value for AppliedByUserFullName to be an explicit nil

### UnsetAppliedByUserFullName
`func (o *ChangeRequestModel) UnsetAppliedByUserFullName()`

UnsetAppliedByUserFullName ensures that no value is present for AppliedByUserFullName, not even an explicit nil
### GetClosedAt

`func (o *ChangeRequestModel) GetClosedAt() time.Time`

GetClosedAt returns the ClosedAt field if non-nil, zero value otherwise.

### GetClosedAtOk

`func (o *ChangeRequestModel) GetClosedAtOk() (*time.Time, bool)`

GetClosedAtOk returns a tuple with the ClosedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedAt

`func (o *ChangeRequestModel) SetClosedAt(v time.Time)`

SetClosedAt sets ClosedAt field to given value.


### SetClosedAtNil

`func (o *ChangeRequestModel) SetClosedAtNil(b bool)`

 SetClosedAtNil sets the value for ClosedAt to be an explicit nil

### UnsetClosedAt
`func (o *ChangeRequestModel) UnsetClosedAt()`

UnsetClosedAt ensures that no value is present for ClosedAt, not even an explicit nil
### GetClosedByUserId

`func (o *ChangeRequestModel) GetClosedByUserId() string`

GetClosedByUserId returns the ClosedByUserId field if non-nil, zero value otherwise.

### GetClosedByUserIdOk

`func (o *ChangeRequestModel) GetClosedByUserIdOk() (*string, bool)`

GetClosedByUserIdOk returns a tuple with the ClosedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedByUserId

`func (o *ChangeRequestModel) SetClosedByUserId(v string)`

SetClosedByUserId sets ClosedByUserId field to given value.


### SetClosedByUserIdNil

`func (o *ChangeRequestModel) SetClosedByUserIdNil(b bool)`

 SetClosedByUserIdNil sets the value for ClosedByUserId to be an explicit nil

### UnsetClosedByUserId
`func (o *ChangeRequestModel) UnsetClosedByUserId()`

UnsetClosedByUserId ensures that no value is present for ClosedByUserId, not even an explicit nil
### GetClosedByUserEmail

`func (o *ChangeRequestModel) GetClosedByUserEmail() string`

GetClosedByUserEmail returns the ClosedByUserEmail field if non-nil, zero value otherwise.

### GetClosedByUserEmailOk

`func (o *ChangeRequestModel) GetClosedByUserEmailOk() (*string, bool)`

GetClosedByUserEmailOk returns a tuple with the ClosedByUserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedByUserEmail

`func (o *ChangeRequestModel) SetClosedByUserEmail(v string)`

SetClosedByUserEmail sets ClosedByUserEmail field to given value.


### SetClosedByUserEmailNil

`func (o *ChangeRequestModel) SetClosedByUserEmailNil(b bool)`

 SetClosedByUserEmailNil sets the value for ClosedByUserEmail to be an explicit nil

### UnsetClosedByUserEmail
`func (o *ChangeRequestModel) UnsetClosedByUserEmail()`

UnsetClosedByUserEmail ensures that no value is present for ClosedByUserEmail, not even an explicit nil
### GetClosedByUserFullName

`func (o *ChangeRequestModel) GetClosedByUserFullName() string`

GetClosedByUserFullName returns the ClosedByUserFullName field if non-nil, zero value otherwise.

### GetClosedByUserFullNameOk

`func (o *ChangeRequestModel) GetClosedByUserFullNameOk() (*string, bool)`

GetClosedByUserFullNameOk returns a tuple with the ClosedByUserFullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedByUserFullName

`func (o *ChangeRequestModel) SetClosedByUserFullName(v string)`

SetClosedByUserFullName sets ClosedByUserFullName field to given value.


### SetClosedByUserFullNameNil

`func (o *ChangeRequestModel) SetClosedByUserFullNameNil(b bool)`

 SetClosedByUserFullNameNil sets the value for ClosedByUserFullName to be an explicit nil

### UnsetClosedByUserFullName
`func (o *ChangeRequestModel) UnsetClosedByUserFullName()`

UnsetClosedByUserFullName ensures that no value is present for ClosedByUserFullName, not even an explicit nil
### GetBypassApproval

`func (o *ChangeRequestModel) GetBypassApproval() bool`

GetBypassApproval returns the BypassApproval field if non-nil, zero value otherwise.

### GetBypassApprovalOk

`func (o *ChangeRequestModel) GetBypassApprovalOk() (*bool, bool)`

GetBypassApprovalOk returns a tuple with the BypassApproval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBypassApproval

`func (o *ChangeRequestModel) SetBypassApproval(v bool)`

SetBypassApproval sets BypassApproval field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


