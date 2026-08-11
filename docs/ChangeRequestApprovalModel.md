# ChangeRequestApprovalModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChangeRequestApprovalId** | **int64** | Identifier of the Change Request approval. | 
**UserEmail** | **string** | Email of the user who approved the Change Request. | 
**UserFullName** | **string** | Full name of the user who approved the Change Request. | 
**ApprovedAt** | **time.Time** | The UTC date and time when the Change Request was approved. | 
**DismissedAt** | **NullableTime** | Optional UTC date and time when the approval was dismissed. | 

## Methods

### NewChangeRequestApprovalModel

`func NewChangeRequestApprovalModel(changeRequestApprovalId int64, userEmail string, userFullName string, approvedAt time.Time, dismissedAt NullableTime, ) *ChangeRequestApprovalModel`

NewChangeRequestApprovalModel instantiates a new ChangeRequestApprovalModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangeRequestApprovalModelWithDefaults

`func NewChangeRequestApprovalModelWithDefaults() *ChangeRequestApprovalModel`

NewChangeRequestApprovalModelWithDefaults instantiates a new ChangeRequestApprovalModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChangeRequestApprovalId

`func (o *ChangeRequestApprovalModel) GetChangeRequestApprovalId() int64`

GetChangeRequestApprovalId returns the ChangeRequestApprovalId field if non-nil, zero value otherwise.

### GetChangeRequestApprovalIdOk

`func (o *ChangeRequestApprovalModel) GetChangeRequestApprovalIdOk() (*int64, bool)`

GetChangeRequestApprovalIdOk returns a tuple with the ChangeRequestApprovalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeRequestApprovalId

`func (o *ChangeRequestApprovalModel) SetChangeRequestApprovalId(v int64)`

SetChangeRequestApprovalId sets ChangeRequestApprovalId field to given value.


### GetUserEmail

`func (o *ChangeRequestApprovalModel) GetUserEmail() string`

GetUserEmail returns the UserEmail field if non-nil, zero value otherwise.

### GetUserEmailOk

`func (o *ChangeRequestApprovalModel) GetUserEmailOk() (*string, bool)`

GetUserEmailOk returns a tuple with the UserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEmail

`func (o *ChangeRequestApprovalModel) SetUserEmail(v string)`

SetUserEmail sets UserEmail field to given value.


### GetUserFullName

`func (o *ChangeRequestApprovalModel) GetUserFullName() string`

GetUserFullName returns the UserFullName field if non-nil, zero value otherwise.

### GetUserFullNameOk

`func (o *ChangeRequestApprovalModel) GetUserFullNameOk() (*string, bool)`

GetUserFullNameOk returns a tuple with the UserFullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserFullName

`func (o *ChangeRequestApprovalModel) SetUserFullName(v string)`

SetUserFullName sets UserFullName field to given value.


### GetApprovedAt

`func (o *ChangeRequestApprovalModel) GetApprovedAt() time.Time`

GetApprovedAt returns the ApprovedAt field if non-nil, zero value otherwise.

### GetApprovedAtOk

`func (o *ChangeRequestApprovalModel) GetApprovedAtOk() (*time.Time, bool)`

GetApprovedAtOk returns a tuple with the ApprovedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedAt

`func (o *ChangeRequestApprovalModel) SetApprovedAt(v time.Time)`

SetApprovedAt sets ApprovedAt field to given value.


### GetDismissedAt

`func (o *ChangeRequestApprovalModel) GetDismissedAt() time.Time`

GetDismissedAt returns the DismissedAt field if non-nil, zero value otherwise.

### GetDismissedAtOk

`func (o *ChangeRequestApprovalModel) GetDismissedAtOk() (*time.Time, bool)`

GetDismissedAtOk returns a tuple with the DismissedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDismissedAt

`func (o *ChangeRequestApprovalModel) SetDismissedAt(v time.Time)`

SetDismissedAt sets DismissedAt field to given value.


### SetDismissedAtNil

`func (o *ChangeRequestApprovalModel) SetDismissedAtNil(b bool)`

 SetDismissedAtNil sets the value for DismissedAt to be an explicit nil

### UnsetDismissedAt
`func (o *ChangeRequestApprovalModel) UnsetDismissedAt()`

UnsetDismissedAt ensures that no value is present for DismissedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


