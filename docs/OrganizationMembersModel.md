# OrganizationMembersModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Admins** | [**[]OrganizationAdminModel**](OrganizationAdminModel.md) | List of Organization Admins. | 
**BillingManagers** | [**[]OrganizationAdminModel**](OrganizationAdminModel.md) | List of Billing Managers. | 
**Members** | [**[]OrganizationMemberModel**](OrganizationMemberModel.md) | List of Organization Members. | 

## Methods

### NewOrganizationMembersModel

`func NewOrganizationMembersModel(admins []OrganizationAdminModel, billingManagers []OrganizationAdminModel, members []OrganizationMemberModel, ) *OrganizationMembersModel`

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


### GetBillingManagers

`func (o *OrganizationMembersModel) GetBillingManagers() []OrganizationAdminModel`

GetBillingManagers returns the BillingManagers field if non-nil, zero value otherwise.

### GetBillingManagersOk

`func (o *OrganizationMembersModel) GetBillingManagersOk() (*[]OrganizationAdminModel, bool)`

GetBillingManagersOk returns a tuple with the BillingManagers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingManagers

`func (o *OrganizationMembersModel) SetBillingManagers(v []OrganizationAdminModel)`

SetBillingManagers sets BillingManagers field to given value.


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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


