# UpdateMemberPermissionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PermissionGroupIds** | Pointer to **[]int64** | List of Permission Group identifiers to where the Member should be added. | [optional] 
**IsAdmin** | Pointer to **NullableBool** | Indicates that the member must be Organization Admin. | [optional] 
**IsBillingManager** | Pointer to **NullableBool** | Indicates that the member must be Billing Manager. | [optional] 
**RemoveFromPermissionGroupsWhereIdNotSet** | Pointer to **bool** | When &#x60;true&#x60;, the member will be removed from those Permission Groups that are not listed in the &#x60;permissionGroupIds&#x60; field. | [optional] 

## Methods

### NewUpdateMemberPermissionsRequest

`func NewUpdateMemberPermissionsRequest() *UpdateMemberPermissionsRequest`

NewUpdateMemberPermissionsRequest instantiates a new UpdateMemberPermissionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMemberPermissionsRequestWithDefaults

`func NewUpdateMemberPermissionsRequestWithDefaults() *UpdateMemberPermissionsRequest`

NewUpdateMemberPermissionsRequestWithDefaults instantiates a new UpdateMemberPermissionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermissionGroupIds

`func (o *UpdateMemberPermissionsRequest) GetPermissionGroupIds() []int64`

GetPermissionGroupIds returns the PermissionGroupIds field if non-nil, zero value otherwise.

### GetPermissionGroupIdsOk

`func (o *UpdateMemberPermissionsRequest) GetPermissionGroupIdsOk() (*[]int64, bool)`

GetPermissionGroupIdsOk returns a tuple with the PermissionGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionGroupIds

`func (o *UpdateMemberPermissionsRequest) SetPermissionGroupIds(v []int64)`

SetPermissionGroupIds sets PermissionGroupIds field to given value.

### HasPermissionGroupIds

`func (o *UpdateMemberPermissionsRequest) HasPermissionGroupIds() bool`

HasPermissionGroupIds returns a boolean if a field has been set.

### SetPermissionGroupIdsNil

`func (o *UpdateMemberPermissionsRequest) SetPermissionGroupIdsNil(b bool)`

 SetPermissionGroupIdsNil sets the value for PermissionGroupIds to be an explicit nil

### UnsetPermissionGroupIds
`func (o *UpdateMemberPermissionsRequest) UnsetPermissionGroupIds()`

UnsetPermissionGroupIds ensures that no value is present for PermissionGroupIds, not even an explicit nil
### GetIsAdmin

`func (o *UpdateMemberPermissionsRequest) GetIsAdmin() bool`

GetIsAdmin returns the IsAdmin field if non-nil, zero value otherwise.

### GetIsAdminOk

`func (o *UpdateMemberPermissionsRequest) GetIsAdminOk() (*bool, bool)`

GetIsAdminOk returns a tuple with the IsAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdmin

`func (o *UpdateMemberPermissionsRequest) SetIsAdmin(v bool)`

SetIsAdmin sets IsAdmin field to given value.

### HasIsAdmin

`func (o *UpdateMemberPermissionsRequest) HasIsAdmin() bool`

HasIsAdmin returns a boolean if a field has been set.

### SetIsAdminNil

`func (o *UpdateMemberPermissionsRequest) SetIsAdminNil(b bool)`

 SetIsAdminNil sets the value for IsAdmin to be an explicit nil

### UnsetIsAdmin
`func (o *UpdateMemberPermissionsRequest) UnsetIsAdmin()`

UnsetIsAdmin ensures that no value is present for IsAdmin, not even an explicit nil
### GetIsBillingManager

`func (o *UpdateMemberPermissionsRequest) GetIsBillingManager() bool`

GetIsBillingManager returns the IsBillingManager field if non-nil, zero value otherwise.

### GetIsBillingManagerOk

`func (o *UpdateMemberPermissionsRequest) GetIsBillingManagerOk() (*bool, bool)`

GetIsBillingManagerOk returns a tuple with the IsBillingManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBillingManager

`func (o *UpdateMemberPermissionsRequest) SetIsBillingManager(v bool)`

SetIsBillingManager sets IsBillingManager field to given value.

### HasIsBillingManager

`func (o *UpdateMemberPermissionsRequest) HasIsBillingManager() bool`

HasIsBillingManager returns a boolean if a field has been set.

### SetIsBillingManagerNil

`func (o *UpdateMemberPermissionsRequest) SetIsBillingManagerNil(b bool)`

 SetIsBillingManagerNil sets the value for IsBillingManager to be an explicit nil

### UnsetIsBillingManager
`func (o *UpdateMemberPermissionsRequest) UnsetIsBillingManager()`

UnsetIsBillingManager ensures that no value is present for IsBillingManager, not even an explicit nil
### GetRemoveFromPermissionGroupsWhereIdNotSet

`func (o *UpdateMemberPermissionsRequest) GetRemoveFromPermissionGroupsWhereIdNotSet() bool`

GetRemoveFromPermissionGroupsWhereIdNotSet returns the RemoveFromPermissionGroupsWhereIdNotSet field if non-nil, zero value otherwise.

### GetRemoveFromPermissionGroupsWhereIdNotSetOk

`func (o *UpdateMemberPermissionsRequest) GetRemoveFromPermissionGroupsWhereIdNotSetOk() (*bool, bool)`

GetRemoveFromPermissionGroupsWhereIdNotSetOk returns a tuple with the RemoveFromPermissionGroupsWhereIdNotSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveFromPermissionGroupsWhereIdNotSet

`func (o *UpdateMemberPermissionsRequest) SetRemoveFromPermissionGroupsWhereIdNotSet(v bool)`

SetRemoveFromPermissionGroupsWhereIdNotSet sets RemoveFromPermissionGroupsWhereIdNotSet field to given value.

### HasRemoveFromPermissionGroupsWhereIdNotSet

`func (o *UpdateMemberPermissionsRequest) HasRemoveFromPermissionGroupsWhereIdNotSet() bool`

HasRemoveFromPermissionGroupsWhereIdNotSet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


