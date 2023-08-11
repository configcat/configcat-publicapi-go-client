# MemberModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **NullableString** | Identifier of the Member. | [optional] 
**ProductId** | Pointer to **string** | Identifier of the Product where the Member has access. | [optional] 
**PermissionGroupId** | Pointer to **int64** | Identifier of the Member&#39;s Permission Group. | [optional] 
**FullName** | Pointer to **NullableString** | Name of the Member. | [optional] 
**Email** | Pointer to **NullableString** | Email of the Member. | [optional] 

## Methods

### NewMemberModel

`func NewMemberModel() *MemberModel`

NewMemberModel instantiates a new MemberModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberModelWithDefaults

`func NewMemberModelWithDefaults() *MemberModel`

NewMemberModelWithDefaults instantiates a new MemberModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *MemberModel) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *MemberModel) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *MemberModel) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *MemberModel) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *MemberModel) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *MemberModel) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetProductId

`func (o *MemberModel) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *MemberModel) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *MemberModel) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *MemberModel) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetPermissionGroupId

`func (o *MemberModel) GetPermissionGroupId() int64`

GetPermissionGroupId returns the PermissionGroupId field if non-nil, zero value otherwise.

### GetPermissionGroupIdOk

`func (o *MemberModel) GetPermissionGroupIdOk() (*int64, bool)`

GetPermissionGroupIdOk returns a tuple with the PermissionGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionGroupId

`func (o *MemberModel) SetPermissionGroupId(v int64)`

SetPermissionGroupId sets PermissionGroupId field to given value.

### HasPermissionGroupId

`func (o *MemberModel) HasPermissionGroupId() bool`

HasPermissionGroupId returns a boolean if a field has been set.

### GetFullName

`func (o *MemberModel) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *MemberModel) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *MemberModel) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *MemberModel) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### SetFullNameNil

`func (o *MemberModel) SetFullNameNil(b bool)`

 SetFullNameNil sets the value for FullName to be an explicit nil

### UnsetFullName
`func (o *MemberModel) UnsetFullName()`

UnsetFullName ensures that no value is present for FullName, not even an explicit nil
### GetEmail

`func (o *MemberModel) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *MemberModel) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *MemberModel) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *MemberModel) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *MemberModel) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *MemberModel) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


