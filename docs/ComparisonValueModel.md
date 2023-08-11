# ComparisonValueModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StringValue** | Pointer to **NullableString** | The string representation of the comparison value. | [optional] 
**DoubleValue** | Pointer to **NullableFloat64** | The number representation of the comparison value. | [optional] 
**ListValue** | Pointer to [**[]ComparisonValueListModel**](ComparisonValueListModel.md) | The list representation of the comparison value. | [optional] 

## Methods

### NewComparisonValueModel

`func NewComparisonValueModel() *ComparisonValueModel`

NewComparisonValueModel instantiates a new ComparisonValueModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComparisonValueModelWithDefaults

`func NewComparisonValueModelWithDefaults() *ComparisonValueModel`

NewComparisonValueModelWithDefaults instantiates a new ComparisonValueModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStringValue

`func (o *ComparisonValueModel) GetStringValue() string`

GetStringValue returns the StringValue field if non-nil, zero value otherwise.

### GetStringValueOk

`func (o *ComparisonValueModel) GetStringValueOk() (*string, bool)`

GetStringValueOk returns a tuple with the StringValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStringValue

`func (o *ComparisonValueModel) SetStringValue(v string)`

SetStringValue sets StringValue field to given value.

### HasStringValue

`func (o *ComparisonValueModel) HasStringValue() bool`

HasStringValue returns a boolean if a field has been set.

### SetStringValueNil

`func (o *ComparisonValueModel) SetStringValueNil(b bool)`

 SetStringValueNil sets the value for StringValue to be an explicit nil

### UnsetStringValue
`func (o *ComparisonValueModel) UnsetStringValue()`

UnsetStringValue ensures that no value is present for StringValue, not even an explicit nil
### GetDoubleValue

`func (o *ComparisonValueModel) GetDoubleValue() float64`

GetDoubleValue returns the DoubleValue field if non-nil, zero value otherwise.

### GetDoubleValueOk

`func (o *ComparisonValueModel) GetDoubleValueOk() (*float64, bool)`

GetDoubleValueOk returns a tuple with the DoubleValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoubleValue

`func (o *ComparisonValueModel) SetDoubleValue(v float64)`

SetDoubleValue sets DoubleValue field to given value.

### HasDoubleValue

`func (o *ComparisonValueModel) HasDoubleValue() bool`

HasDoubleValue returns a boolean if a field has been set.

### SetDoubleValueNil

`func (o *ComparisonValueModel) SetDoubleValueNil(b bool)`

 SetDoubleValueNil sets the value for DoubleValue to be an explicit nil

### UnsetDoubleValue
`func (o *ComparisonValueModel) UnsetDoubleValue()`

UnsetDoubleValue ensures that no value is present for DoubleValue, not even an explicit nil
### GetListValue

`func (o *ComparisonValueModel) GetListValue() []ComparisonValueListModel`

GetListValue returns the ListValue field if non-nil, zero value otherwise.

### GetListValueOk

`func (o *ComparisonValueModel) GetListValueOk() (*[]ComparisonValueListModel, bool)`

GetListValueOk returns a tuple with the ListValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListValue

`func (o *ComparisonValueModel) SetListValue(v []ComparisonValueListModel)`

SetListValue sets ListValue field to given value.

### HasListValue

`func (o *ComparisonValueModel) HasListValue() bool`

HasListValue returns a boolean if a field has been set.

### SetListValueNil

`func (o *ComparisonValueModel) SetListValueNil(b bool)`

 SetListValueNil sets the value for ListValue to be an explicit nil

### UnsetListValue
`func (o *ComparisonValueModel) UnsetListValue()`

UnsetListValue ensures that no value is present for ListValue, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


