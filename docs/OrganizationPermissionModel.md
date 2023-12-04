# OrganizationPermissionModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Product** | Pointer to [**OrganizationProductModel**](OrganizationProductModel.md) |  | [optional] 
**PermissionGroup** | Pointer to [**OrganizationPermissionGroupModel**](OrganizationPermissionGroupModel.md) |  | [optional] 

## Methods

### NewOrganizationPermissionModel

`func NewOrganizationPermissionModel() *OrganizationPermissionModel`

NewOrganizationPermissionModel instantiates a new OrganizationPermissionModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationPermissionModelWithDefaults

`func NewOrganizationPermissionModelWithDefaults() *OrganizationPermissionModel`

NewOrganizationPermissionModelWithDefaults instantiates a new OrganizationPermissionModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProduct

`func (o *OrganizationPermissionModel) GetProduct() OrganizationProductModel`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *OrganizationPermissionModel) GetProductOk() (*OrganizationProductModel, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *OrganizationPermissionModel) SetProduct(v OrganizationProductModel)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *OrganizationPermissionModel) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetPermissionGroup

`func (o *OrganizationPermissionModel) GetPermissionGroup() OrganizationPermissionGroupModel`

GetPermissionGroup returns the PermissionGroup field if non-nil, zero value otherwise.

### GetPermissionGroupOk

`func (o *OrganizationPermissionModel) GetPermissionGroupOk() (*OrganizationPermissionGroupModel, bool)`

GetPermissionGroupOk returns a tuple with the PermissionGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionGroup

`func (o *OrganizationPermissionModel) SetPermissionGroup(v OrganizationPermissionGroupModel)`

SetPermissionGroup sets PermissionGroup field to given value.

### HasPermissionGroup

`func (o *OrganizationPermissionModel) HasPermissionGroup() bool`

HasPermissionGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


