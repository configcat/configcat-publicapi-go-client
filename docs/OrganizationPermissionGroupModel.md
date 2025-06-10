# OrganizationPermissionGroupModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PermissionGroupId** | **int64** | Identifier of the Member&#39;s Permission Group. | 
**Name** | **string** | Name of the Member&#39;s Permission Group. | 

## Methods

### NewOrganizationPermissionGroupModel

`func NewOrganizationPermissionGroupModel(permissionGroupId int64, name string, ) *OrganizationPermissionGroupModel`

NewOrganizationPermissionGroupModel instantiates a new OrganizationPermissionGroupModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationPermissionGroupModelWithDefaults

`func NewOrganizationPermissionGroupModelWithDefaults() *OrganizationPermissionGroupModel`

NewOrganizationPermissionGroupModelWithDefaults instantiates a new OrganizationPermissionGroupModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermissionGroupId

`func (o *OrganizationPermissionGroupModel) GetPermissionGroupId() int64`

GetPermissionGroupId returns the PermissionGroupId field if non-nil, zero value otherwise.

### GetPermissionGroupIdOk

`func (o *OrganizationPermissionGroupModel) GetPermissionGroupIdOk() (*int64, bool)`

GetPermissionGroupIdOk returns a tuple with the PermissionGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionGroupId

`func (o *OrganizationPermissionGroupModel) SetPermissionGroupId(v int64)`

SetPermissionGroupId sets PermissionGroupId field to given value.


### GetName

`func (o *OrganizationPermissionGroupModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationPermissionGroupModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationPermissionGroupModel) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


