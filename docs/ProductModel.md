# ProductModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Organization** | [**OrganizationModel**](OrganizationModel.md) |  | 
**ProductId** | **string** | Identifier of the Product. | 
**Name** | **string** | Name of the Product. | 
**Description** | **NullableString** | Description of the Product. | 
**Order** | **int32** | The order of the Product represented on the ConfigCat Dashboard. Determined from an ascending sequence of integers. | 
**ReasonRequired** | **bool** | Determines whether a mandatory reason must be given every time when the Feature Flags or Settings within a Product are saved. | 

## Methods

### NewProductModel

`func NewProductModel(organization OrganizationModel, productId string, name string, description NullableString, order int32, reasonRequired bool, ) *ProductModel`

NewProductModel instantiates a new ProductModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductModelWithDefaults

`func NewProductModelWithDefaults() *ProductModel`

NewProductModelWithDefaults instantiates a new ProductModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganization

`func (o *ProductModel) GetOrganization() OrganizationModel`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *ProductModel) GetOrganizationOk() (*OrganizationModel, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *ProductModel) SetOrganization(v OrganizationModel)`

SetOrganization sets Organization field to given value.


### GetProductId

`func (o *ProductModel) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *ProductModel) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *ProductModel) SetProductId(v string)`

SetProductId sets ProductId field to given value.


### GetName

`func (o *ProductModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProductModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProductModel) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ProductModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProductModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProductModel) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *ProductModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ProductModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetOrder

`func (o *ProductModel) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ProductModel) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ProductModel) SetOrder(v int32)`

SetOrder sets Order field to given value.


### GetReasonRequired

`func (o *ProductModel) GetReasonRequired() bool`

GetReasonRequired returns the ReasonRequired field if non-nil, zero value otherwise.

### GetReasonRequiredOk

`func (o *ProductModel) GetReasonRequiredOk() (*bool, bool)`

GetReasonRequiredOk returns a tuple with the ReasonRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonRequired

`func (o *ProductModel) SetReasonRequired(v bool)`

SetReasonRequired sets ReasonRequired field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


