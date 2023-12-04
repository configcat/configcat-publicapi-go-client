# OrganizationProductModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductId** | Pointer to **string** | Identifier of the Member&#39;s Product. | [optional] 
**Name** | Pointer to **NullableString** | Name of the Member&#39;s Product. | [optional] 

## Methods

### NewOrganizationProductModel

`func NewOrganizationProductModel() *OrganizationProductModel`

NewOrganizationProductModel instantiates a new OrganizationProductModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationProductModelWithDefaults

`func NewOrganizationProductModelWithDefaults() *OrganizationProductModel`

NewOrganizationProductModelWithDefaults instantiates a new OrganizationProductModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductId

`func (o *OrganizationProductModel) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *OrganizationProductModel) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *OrganizationProductModel) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *OrganizationProductModel) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetName

`func (o *OrganizationProductModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationProductModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationProductModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrganizationProductModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *OrganizationProductModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *OrganizationProductModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


