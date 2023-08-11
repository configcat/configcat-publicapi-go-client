# CreateTagModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the Tag. | 
**Color** | Pointer to **NullableString** | Color of the Tag. | [optional] 

## Methods

### NewCreateTagModel

`func NewCreateTagModel(name string, ) *CreateTagModel`

NewCreateTagModel instantiates a new CreateTagModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTagModelWithDefaults

`func NewCreateTagModelWithDefaults() *CreateTagModel`

NewCreateTagModelWithDefaults instantiates a new CreateTagModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateTagModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateTagModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateTagModel) SetName(v string)`

SetName sets Name field to given value.


### GetColor

`func (o *CreateTagModel) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *CreateTagModel) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *CreateTagModel) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *CreateTagModel) HasColor() bool`

HasColor returns a boolean if a field has been set.

### SetColorNil

`func (o *CreateTagModel) SetColorNil(b bool)`

 SetColorNil sets the value for Color to be an explicit nil

### UnsetColor
`func (o *CreateTagModel) UnsetColor()`

UnsetColor ensures that no value is present for Color, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


