# TagModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Product** | Pointer to [**ProductModel**](ProductModel.md) |  | [optional] 
**TagId** | Pointer to **int64** | Identifier of the Tag. | [optional] 
**Name** | Pointer to **NullableString** | Name of the Tag. | [optional] 
**Color** | Pointer to **NullableString** | The configured color of the Tag. | [optional] 

## Methods

### NewTagModel

`func NewTagModel() *TagModel`

NewTagModel instantiates a new TagModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagModelWithDefaults

`func NewTagModelWithDefaults() *TagModel`

NewTagModelWithDefaults instantiates a new TagModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProduct

`func (o *TagModel) GetProduct() ProductModel`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *TagModel) GetProductOk() (*ProductModel, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *TagModel) SetProduct(v ProductModel)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *TagModel) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetTagId

`func (o *TagModel) GetTagId() int64`

GetTagId returns the TagId field if non-nil, zero value otherwise.

### GetTagIdOk

`func (o *TagModel) GetTagIdOk() (*int64, bool)`

GetTagIdOk returns a tuple with the TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagId

`func (o *TagModel) SetTagId(v int64)`

SetTagId sets TagId field to given value.

### HasTagId

`func (o *TagModel) HasTagId() bool`

HasTagId returns a boolean if a field has been set.

### GetName

`func (o *TagModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TagModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TagModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TagModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TagModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TagModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetColor

`func (o *TagModel) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *TagModel) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *TagModel) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *TagModel) HasColor() bool`

HasColor returns a boolean if a field has been set.

### SetColorNil

`func (o *TagModel) SetColorNil(b bool)`

 SetColorNil sets the value for Color to be an explicit nil

### UnsetColor
`func (o *TagModel) UnsetColor()`

UnsetColor ensures that no value is present for Color, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


