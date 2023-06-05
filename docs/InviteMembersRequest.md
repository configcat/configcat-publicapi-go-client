# InviteMembersRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Emails** | **[]string** |  | 
**PermissionGroupId** | **int64** |  | 

## Methods

### NewInviteMembersRequest

`func NewInviteMembersRequest(emails []string, permissionGroupId int64, ) *InviteMembersRequest`

NewInviteMembersRequest instantiates a new InviteMembersRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInviteMembersRequestWithDefaults

`func NewInviteMembersRequestWithDefaults() *InviteMembersRequest`

NewInviteMembersRequestWithDefaults instantiates a new InviteMembersRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmails

`func (o *InviteMembersRequest) GetEmails() []string`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *InviteMembersRequest) GetEmailsOk() (*[]string, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *InviteMembersRequest) SetEmails(v []string)`

SetEmails sets Emails field to given value.


### GetPermissionGroupId

`func (o *InviteMembersRequest) GetPermissionGroupId() int64`

GetPermissionGroupId returns the PermissionGroupId field if non-nil, zero value otherwise.

### GetPermissionGroupIdOk

`func (o *InviteMembersRequest) GetPermissionGroupIdOk() (*int64, bool)`

GetPermissionGroupIdOk returns a tuple with the PermissionGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionGroupId

`func (o *InviteMembersRequest) SetPermissionGroupId(v int64)`

SetPermissionGroupId sets PermissionGroupId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


