# OrganizationMemberModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **NullableString** | Identifier of the Organization Admin. | [optional] 
**FullName** | Pointer to **NullableString** | Name of the Organization Admin. | [optional] 
**Email** | Pointer to **NullableString** | Email of the OrganizationAdmin. | [optional] 
**TwoFactorEnabled** | Pointer to **bool** | Determines whether 2FA is enabled for the Organization Admin. | [optional] 
**Permissions** | Pointer to [**[]OrganizationPermissionModel**](OrganizationPermissionModel.md) | The permissions of the Member. | [optional] 

## Methods

### NewOrganizationMemberModel

`func NewOrganizationMemberModel() *OrganizationMemberModel`

NewOrganizationMemberModel instantiates a new OrganizationMemberModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationMemberModelWithDefaults

`func NewOrganizationMemberModelWithDefaults() *OrganizationMemberModel`

NewOrganizationMemberModelWithDefaults instantiates a new OrganizationMemberModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *OrganizationMemberModel) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *OrganizationMemberModel) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *OrganizationMemberModel) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *OrganizationMemberModel) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *OrganizationMemberModel) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *OrganizationMemberModel) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetFullName

`func (o *OrganizationMemberModel) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *OrganizationMemberModel) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *OrganizationMemberModel) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *OrganizationMemberModel) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### SetFullNameNil

`func (o *OrganizationMemberModel) SetFullNameNil(b bool)`

 SetFullNameNil sets the value for FullName to be an explicit nil

### UnsetFullName
`func (o *OrganizationMemberModel) UnsetFullName()`

UnsetFullName ensures that no value is present for FullName, not even an explicit nil
### GetEmail

`func (o *OrganizationMemberModel) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *OrganizationMemberModel) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *OrganizationMemberModel) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *OrganizationMemberModel) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *OrganizationMemberModel) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *OrganizationMemberModel) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetTwoFactorEnabled

`func (o *OrganizationMemberModel) GetTwoFactorEnabled() bool`

GetTwoFactorEnabled returns the TwoFactorEnabled field if non-nil, zero value otherwise.

### GetTwoFactorEnabledOk

`func (o *OrganizationMemberModel) GetTwoFactorEnabledOk() (*bool, bool)`

GetTwoFactorEnabledOk returns a tuple with the TwoFactorEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoFactorEnabled

`func (o *OrganizationMemberModel) SetTwoFactorEnabled(v bool)`

SetTwoFactorEnabled sets TwoFactorEnabled field to given value.

### HasTwoFactorEnabled

`func (o *OrganizationMemberModel) HasTwoFactorEnabled() bool`

HasTwoFactorEnabled returns a boolean if a field has been set.

### GetPermissions

`func (o *OrganizationMemberModel) GetPermissions() []OrganizationPermissionModel`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *OrganizationMemberModel) GetPermissionsOk() (*[]OrganizationPermissionModel, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *OrganizationMemberModel) SetPermissions(v []OrganizationPermissionModel)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *OrganizationMemberModel) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### SetPermissionsNil

`func (o *OrganizationMemberModel) SetPermissionsNil(b bool)`

 SetPermissionsNil sets the value for Permissions to be an explicit nil

### UnsetPermissions
`func (o *OrganizationMemberModel) UnsetPermissions()`

UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


