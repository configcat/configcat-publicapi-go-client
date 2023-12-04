# ConfigModelHaljsonEmbeddedProduct

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Embedded** | Pointer to [**ConfigModelHaljsonEmbeddedProductEmbedded**](ConfigModelHaljsonEmbeddedProductEmbedded.md) |  | [optional] 
**ProductId** | Pointer to **string** | Identifier of the Product. | [optional] 
**Name** | Pointer to **NullableString** | Name of the Product. | [optional] 
**Description** | Pointer to **NullableString** | Description of the Product. | [optional] 
**Order** | Pointer to **int32** | The order of the Product represented on the ConfigCat Dashboard.  Determined from an ascending sequence of integers. | [optional] 
**ReasonRequired** | Pointer to **bool** | Determines whether a mandatory reason must be given every time when the Feature Flags or Settings within a Product are saved. | [optional] 
**Links** | Pointer to [**ConfigModelHaljsonEmbeddedProductLinks**](ConfigModelHaljsonEmbeddedProductLinks.md) |  | [optional] 

## Methods

### NewConfigModelHaljsonEmbeddedProduct

`func NewConfigModelHaljsonEmbeddedProduct() *ConfigModelHaljsonEmbeddedProduct`

NewConfigModelHaljsonEmbeddedProduct instantiates a new ConfigModelHaljsonEmbeddedProduct object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigModelHaljsonEmbeddedProductWithDefaults

`func NewConfigModelHaljsonEmbeddedProductWithDefaults() *ConfigModelHaljsonEmbeddedProduct`

NewConfigModelHaljsonEmbeddedProductWithDefaults instantiates a new ConfigModelHaljsonEmbeddedProduct object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmbedded

`func (o *ConfigModelHaljsonEmbeddedProduct) GetEmbedded() ConfigModelHaljsonEmbeddedProductEmbedded`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *ConfigModelHaljsonEmbeddedProduct) GetEmbeddedOk() (*ConfigModelHaljsonEmbeddedProductEmbedded, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *ConfigModelHaljsonEmbeddedProduct) SetEmbedded(v ConfigModelHaljsonEmbeddedProductEmbedded)`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *ConfigModelHaljsonEmbeddedProduct) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetProductId

`func (o *ConfigModelHaljsonEmbeddedProduct) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *ConfigModelHaljsonEmbeddedProduct) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *ConfigModelHaljsonEmbeddedProduct) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *ConfigModelHaljsonEmbeddedProduct) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetName

`func (o *ConfigModelHaljsonEmbeddedProduct) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConfigModelHaljsonEmbeddedProduct) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConfigModelHaljsonEmbeddedProduct) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConfigModelHaljsonEmbeddedProduct) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ConfigModelHaljsonEmbeddedProduct) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ConfigModelHaljsonEmbeddedProduct) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *ConfigModelHaljsonEmbeddedProduct) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConfigModelHaljsonEmbeddedProduct) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConfigModelHaljsonEmbeddedProduct) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConfigModelHaljsonEmbeddedProduct) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ConfigModelHaljsonEmbeddedProduct) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ConfigModelHaljsonEmbeddedProduct) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetOrder

`func (o *ConfigModelHaljsonEmbeddedProduct) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ConfigModelHaljsonEmbeddedProduct) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ConfigModelHaljsonEmbeddedProduct) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *ConfigModelHaljsonEmbeddedProduct) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetReasonRequired

`func (o *ConfigModelHaljsonEmbeddedProduct) GetReasonRequired() bool`

GetReasonRequired returns the ReasonRequired field if non-nil, zero value otherwise.

### GetReasonRequiredOk

`func (o *ConfigModelHaljsonEmbeddedProduct) GetReasonRequiredOk() (*bool, bool)`

GetReasonRequiredOk returns a tuple with the ReasonRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonRequired

`func (o *ConfigModelHaljsonEmbeddedProduct) SetReasonRequired(v bool)`

SetReasonRequired sets ReasonRequired field to given value.

### HasReasonRequired

`func (o *ConfigModelHaljsonEmbeddedProduct) HasReasonRequired() bool`

HasReasonRequired returns a boolean if a field has been set.

### GetLinks

`func (o *ConfigModelHaljsonEmbeddedProduct) GetLinks() ConfigModelHaljsonEmbeddedProductLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ConfigModelHaljsonEmbeddedProduct) GetLinksOk() (*ConfigModelHaljsonEmbeddedProductLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ConfigModelHaljsonEmbeddedProduct) SetLinks(v ConfigModelHaljsonEmbeddedProductLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ConfigModelHaljsonEmbeddedProduct) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


