# OrganizationInvitationModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvitationId** | Pointer to **string** | The identifier of the Invitation. | [optional] 
**Email** | Pointer to **NullableString** | The invited user&#39;s email address. | [optional] 
**ProductId** | Pointer to **string** | The identifier of the Product the user was invited to. | [optional] 
**ProductName** | Pointer to **NullableString** | The name of the Product the user was invited to. | [optional] 
**PermissionGroupId** | Pointer to **int64** | The identifier of the Permission Group the user was invited to. | [optional] 
**CreatedAt** | Pointer to **time.Time** | Creation time of the Invitation. | [optional] 
**Expired** | Pointer to **bool** | Determines whether the Invitation is expired. | [optional] 
**Expires** | Pointer to **time.Time** | Expiration time of the Invitation. | [optional] 

## Methods

### NewOrganizationInvitationModel

`func NewOrganizationInvitationModel() *OrganizationInvitationModel`

NewOrganizationInvitationModel instantiates a new OrganizationInvitationModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationInvitationModelWithDefaults

`func NewOrganizationInvitationModelWithDefaults() *OrganizationInvitationModel`

NewOrganizationInvitationModelWithDefaults instantiates a new OrganizationInvitationModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvitationId

`func (o *OrganizationInvitationModel) GetInvitationId() string`

GetInvitationId returns the InvitationId field if non-nil, zero value otherwise.

### GetInvitationIdOk

`func (o *OrganizationInvitationModel) GetInvitationIdOk() (*string, bool)`

GetInvitationIdOk returns a tuple with the InvitationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitationId

`func (o *OrganizationInvitationModel) SetInvitationId(v string)`

SetInvitationId sets InvitationId field to given value.

### HasInvitationId

`func (o *OrganizationInvitationModel) HasInvitationId() bool`

HasInvitationId returns a boolean if a field has been set.

### GetEmail

`func (o *OrganizationInvitationModel) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *OrganizationInvitationModel) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *OrganizationInvitationModel) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *OrganizationInvitationModel) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *OrganizationInvitationModel) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *OrganizationInvitationModel) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetProductId

`func (o *OrganizationInvitationModel) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *OrganizationInvitationModel) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *OrganizationInvitationModel) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *OrganizationInvitationModel) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetProductName

`func (o *OrganizationInvitationModel) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *OrganizationInvitationModel) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *OrganizationInvitationModel) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *OrganizationInvitationModel) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### SetProductNameNil

`func (o *OrganizationInvitationModel) SetProductNameNil(b bool)`

 SetProductNameNil sets the value for ProductName to be an explicit nil

### UnsetProductName
`func (o *OrganizationInvitationModel) UnsetProductName()`

UnsetProductName ensures that no value is present for ProductName, not even an explicit nil
### GetPermissionGroupId

`func (o *OrganizationInvitationModel) GetPermissionGroupId() int64`

GetPermissionGroupId returns the PermissionGroupId field if non-nil, zero value otherwise.

### GetPermissionGroupIdOk

`func (o *OrganizationInvitationModel) GetPermissionGroupIdOk() (*int64, bool)`

GetPermissionGroupIdOk returns a tuple with the PermissionGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionGroupId

`func (o *OrganizationInvitationModel) SetPermissionGroupId(v int64)`

SetPermissionGroupId sets PermissionGroupId field to given value.

### HasPermissionGroupId

`func (o *OrganizationInvitationModel) HasPermissionGroupId() bool`

HasPermissionGroupId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *OrganizationInvitationModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrganizationInvitationModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrganizationInvitationModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *OrganizationInvitationModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetExpired

`func (o *OrganizationInvitationModel) GetExpired() bool`

GetExpired returns the Expired field if non-nil, zero value otherwise.

### GetExpiredOk

`func (o *OrganizationInvitationModel) GetExpiredOk() (*bool, bool)`

GetExpiredOk returns a tuple with the Expired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpired

`func (o *OrganizationInvitationModel) SetExpired(v bool)`

SetExpired sets Expired field to given value.

### HasExpired

`func (o *OrganizationInvitationModel) HasExpired() bool`

HasExpired returns a boolean if a field has been set.

### GetExpires

`func (o *OrganizationInvitationModel) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *OrganizationInvitationModel) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *OrganizationInvitationModel) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *OrganizationInvitationModel) HasExpires() bool`

HasExpires returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


