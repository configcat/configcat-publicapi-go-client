# OrganizationMembersModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Admins** | Pointer to [**[]OrganizationAdminModel**](OrganizationAdminModel.md) | List of Organization Admins. | [optional] 
**Members** | Pointer to [**[]OrganizationMemberModel**](OrganizationMemberModel.md) | List of Organization Members. | [optional] 

## Methods

### NewOrganizationMembersModel

`func NewOrganizationMembersModel() *OrganizationMembersModel`

NewOrganizationMembersModel instantiates a new OrganizationMembersModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationMembersModelWithDefaults

`func NewOrganizationMembersModelWithDefaults() *OrganizationMembersModel`

NewOrganizationMembersModelWithDefaults instantiates a new OrganizationMembersModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdmins

`func (o *OrganizationMembersModel) GetAdmins() []OrganizationAdminModel`

GetAdmins returns the Admins field if non-nil, zero value otherwise.

### GetAdminsOk

`func (o *OrganizationMembersModel) GetAdminsOk() (*[]OrganizationAdminModel, bool)`

GetAdminsOk returns a tuple with the Admins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmins

`func (o *OrganizationMembersModel) SetAdmins(v []OrganizationAdminModel)`

SetAdmins sets Admins field to given value.

### HasAdmins

`func (o *OrganizationMembersModel) HasAdmins() bool`

HasAdmins returns a boolean if a field has been set.

### SetAdminsNil

`func (o *OrganizationMembersModel) SetAdminsNil(b bool)`

 SetAdminsNil sets the value for Admins to be an explicit nil

### UnsetAdmins
`func (o *OrganizationMembersModel) UnsetAdmins()`

UnsetAdmins ensures that no value is present for Admins, not even an explicit nil
### GetMembers

`func (o *OrganizationMembersModel) GetMembers() []OrganizationMemberModel`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *OrganizationMembersModel) GetMembersOk() (*[]OrganizationMemberModel, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *OrganizationMembersModel) SetMembers(v []OrganizationMemberModel)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *OrganizationMembersModel) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### SetMembersNil

`func (o *OrganizationMembersModel) SetMembersNil(b bool)`

 SetMembersNil sets the value for Members to be an explicit nil

### UnsetMembers
`func (o *OrganizationMembersModel) UnsetMembers()`

UnsetMembers ensures that no value is present for Members, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


