# InvitationModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvitationId** | Pointer to **string** | The identifier of the Invitation. | [optional] 
**Email** | Pointer to **NullableString** | The invited user&#39;s email address. | [optional] 
**PermissionGroupId** | Pointer to **int64** | The identifier of the Permission Group the user was invited to. | [optional] 
**CreatedAt** | Pointer to **time.Time** | Creation time of the Invitation. | [optional] 
**Expired** | Pointer to **bool** | Determines whether the Invitation is expired. | [optional] 

## Methods

### NewInvitationModel

`func NewInvitationModel() *InvitationModel`

NewInvitationModel instantiates a new InvitationModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvitationModelWithDefaults

`func NewInvitationModelWithDefaults() *InvitationModel`

NewInvitationModelWithDefaults instantiates a new InvitationModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvitationId

`func (o *InvitationModel) GetInvitationId() string`

GetInvitationId returns the InvitationId field if non-nil, zero value otherwise.

### GetInvitationIdOk

`func (o *InvitationModel) GetInvitationIdOk() (*string, bool)`

GetInvitationIdOk returns a tuple with the InvitationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitationId

`func (o *InvitationModel) SetInvitationId(v string)`

SetInvitationId sets InvitationId field to given value.

### HasInvitationId

`func (o *InvitationModel) HasInvitationId() bool`

HasInvitationId returns a boolean if a field has been set.

### GetEmail

`func (o *InvitationModel) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *InvitationModel) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *InvitationModel) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *InvitationModel) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *InvitationModel) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *InvitationModel) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetPermissionGroupId

`func (o *InvitationModel) GetPermissionGroupId() int64`

GetPermissionGroupId returns the PermissionGroupId field if non-nil, zero value otherwise.

### GetPermissionGroupIdOk

`func (o *InvitationModel) GetPermissionGroupIdOk() (*int64, bool)`

GetPermissionGroupIdOk returns a tuple with the PermissionGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionGroupId

`func (o *InvitationModel) SetPermissionGroupId(v int64)`

SetPermissionGroupId sets PermissionGroupId field to given value.

### HasPermissionGroupId

`func (o *InvitationModel) HasPermissionGroupId() bool`

HasPermissionGroupId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *InvitationModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InvitationModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InvitationModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *InvitationModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetExpired

`func (o *InvitationModel) GetExpired() bool`

GetExpired returns the Expired field if non-nil, zero value otherwise.

### GetExpiredOk

`func (o *InvitationModel) GetExpiredOk() (*bool, bool)`

GetExpiredOk returns a tuple with the Expired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpired

`func (o *InvitationModel) SetExpired(v bool)`

SetExpired sets Expired field to given value.

### HasExpired

`func (o *InvitationModel) HasExpired() bool`

HasExpired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


