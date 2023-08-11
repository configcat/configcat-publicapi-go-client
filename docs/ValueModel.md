# ValueModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BoolValue** | Pointer to **NullableBool** | The served value in case of a boolean Feature Flag. | [optional] 
**StringValue** | Pointer to **NullableString** | The served value in case of a text Setting. | [optional] 
**IntValue** | Pointer to **NullableInt32** | The served value in case of a whole number Setting. | [optional] 
**DoubleValue** | Pointer to **NullableFloat64** | The served value in case of a decimal number Setting. | [optional] 

## Methods

### NewValueModel

`func NewValueModel() *ValueModel`

NewValueModel instantiates a new ValueModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValueModelWithDefaults

`func NewValueModelWithDefaults() *ValueModel`

NewValueModelWithDefaults instantiates a new ValueModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBoolValue

`func (o *ValueModel) GetBoolValue() bool`

GetBoolValue returns the BoolValue field if non-nil, zero value otherwise.

### GetBoolValueOk

`func (o *ValueModel) GetBoolValueOk() (*bool, bool)`

GetBoolValueOk returns a tuple with the BoolValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoolValue

`func (o *ValueModel) SetBoolValue(v bool)`

SetBoolValue sets BoolValue field to given value.

### HasBoolValue

`func (o *ValueModel) HasBoolValue() bool`

HasBoolValue returns a boolean if a field has been set.

### SetBoolValueNil

`func (o *ValueModel) SetBoolValueNil(b bool)`

 SetBoolValueNil sets the value for BoolValue to be an explicit nil

### UnsetBoolValue
`func (o *ValueModel) UnsetBoolValue()`

UnsetBoolValue ensures that no value is present for BoolValue, not even an explicit nil
### GetStringValue

`func (o *ValueModel) GetStringValue() string`

GetStringValue returns the StringValue field if non-nil, zero value otherwise.

### GetStringValueOk

`func (o *ValueModel) GetStringValueOk() (*string, bool)`

GetStringValueOk returns a tuple with the StringValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStringValue

`func (o *ValueModel) SetStringValue(v string)`

SetStringValue sets StringValue field to given value.

### HasStringValue

`func (o *ValueModel) HasStringValue() bool`

HasStringValue returns a boolean if a field has been set.

### SetStringValueNil

`func (o *ValueModel) SetStringValueNil(b bool)`

 SetStringValueNil sets the value for StringValue to be an explicit nil

### UnsetStringValue
`func (o *ValueModel) UnsetStringValue()`

UnsetStringValue ensures that no value is present for StringValue, not even an explicit nil
### GetIntValue

`func (o *ValueModel) GetIntValue() int32`

GetIntValue returns the IntValue field if non-nil, zero value otherwise.

### GetIntValueOk

`func (o *ValueModel) GetIntValueOk() (*int32, bool)`

GetIntValueOk returns a tuple with the IntValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntValue

`func (o *ValueModel) SetIntValue(v int32)`

SetIntValue sets IntValue field to given value.

### HasIntValue

`func (o *ValueModel) HasIntValue() bool`

HasIntValue returns a boolean if a field has been set.

### SetIntValueNil

`func (o *ValueModel) SetIntValueNil(b bool)`

 SetIntValueNil sets the value for IntValue to be an explicit nil

### UnsetIntValue
`func (o *ValueModel) UnsetIntValue()`

UnsetIntValue ensures that no value is present for IntValue, not even an explicit nil
### GetDoubleValue

`func (o *ValueModel) GetDoubleValue() float64`

GetDoubleValue returns the DoubleValue field if non-nil, zero value otherwise.

### GetDoubleValueOk

`func (o *ValueModel) GetDoubleValueOk() (*float64, bool)`

GetDoubleValueOk returns a tuple with the DoubleValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoubleValue

`func (o *ValueModel) SetDoubleValue(v float64)`

SetDoubleValue sets DoubleValue field to given value.

### HasDoubleValue

`func (o *ValueModel) HasDoubleValue() bool`

HasDoubleValue returns a boolean if a field has been set.

### SetDoubleValueNil

`func (o *ValueModel) SetDoubleValueNil(b bool)`

 SetDoubleValueNil sets the value for DoubleValue to be an explicit nil

### UnsetDoubleValue
`func (o *ValueModel) UnsetDoubleValue()`

UnsetDoubleValue ensures that no value is present for DoubleValue, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


