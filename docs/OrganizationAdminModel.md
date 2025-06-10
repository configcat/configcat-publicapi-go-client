# OrganizationAdminModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **string** | Identifier of the Organization Admin. | 
**FullName** | **string** | Name of the Organization Admin. | 
**Email** | **string** | Email of the OrganizationAdmin. | 
**TwoFactorEnabled** | **bool** | Determines whether 2FA is enabled for the Organization Admin. | 

## Methods

### NewOrganizationAdminModel

`func NewOrganizationAdminModel(userId string, fullName string, email string, twoFactorEnabled bool, ) *OrganizationAdminModel`

NewOrganizationAdminModel instantiates a new OrganizationAdminModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationAdminModelWithDefaults

`func NewOrganizationAdminModelWithDefaults() *OrganizationAdminModel`

NewOrganizationAdminModelWithDefaults instantiates a new OrganizationAdminModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *OrganizationAdminModel) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *OrganizationAdminModel) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *OrganizationAdminModel) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetFullName

`func (o *OrganizationAdminModel) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *OrganizationAdminModel) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *OrganizationAdminModel) SetFullName(v string)`

SetFullName sets FullName field to given value.


### GetEmail

`func (o *OrganizationAdminModel) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *OrganizationAdminModel) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *OrganizationAdminModel) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetTwoFactorEnabled

`func (o *OrganizationAdminModel) GetTwoFactorEnabled() bool`

GetTwoFactorEnabled returns the TwoFactorEnabled field if non-nil, zero value otherwise.

### GetTwoFactorEnabledOk

`func (o *OrganizationAdminModel) GetTwoFactorEnabledOk() (*bool, bool)`

GetTwoFactorEnabledOk returns a tuple with the TwoFactorEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoFactorEnabled

`func (o *OrganizationAdminModel) SetTwoFactorEnabled(v bool)`

SetTwoFactorEnabled sets TwoFactorEnabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


