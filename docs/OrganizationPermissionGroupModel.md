# OrganizationPermissionGroupModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PermissionGroupId** | Pointer to **int64** | Identifier of the Member&#39;s Permission Group. | [optional] 
**Name** | Pointer to **NullableString** | Name of the Member&#39;s Permission Group. | [optional] 

## Methods

### NewOrganizationPermissionGroupModel

`func NewOrganizationPermissionGroupModel() *OrganizationPermissionGroupModel`

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

### HasPermissionGroupId

`func (o *OrganizationPermissionGroupModel) HasPermissionGroupId() bool`

HasPermissionGroupId returns a boolean if a field has been set.

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

### HasName

`func (o *OrganizationPermissionGroupModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *OrganizationPermissionGroupModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *OrganizationPermissionGroupModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


