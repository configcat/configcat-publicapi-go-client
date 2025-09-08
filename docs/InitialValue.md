# InitialValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentId** | Pointer to **string** | The ID of the Environment where the initial value must be set. | [optional] 
**Value** | Pointer to [**NullableSettingValueType**](SettingValueType.md) | The initial value in the given Environment. It must respect the setting type. In some generated clients for strictly typed languages, you may use double/float properties to handle integer values. Only one of the Value or the PredefinedVariationIndex properties can be set. | [optional] 
**PredefinedVariationIndex** | Pointer to **NullableInt32** | The initial Predefined Variation in the given Environment. The zero-based index of the Variation in the Variations list. Only one of the Value or the PredefinedVariationIndex properties can be set. | [optional] 

## Methods

### NewInitialValue

`func NewInitialValue() *InitialValue`

NewInitialValue instantiates a new InitialValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInitialValueWithDefaults

`func NewInitialValueWithDefaults() *InitialValue`

NewInitialValueWithDefaults instantiates a new InitialValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentId

`func (o *InitialValue) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *InitialValue) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *InitialValue) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *InitialValue) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### GetValue

`func (o *InitialValue) GetValue() SettingValueType`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *InitialValue) GetValueOk() (*SettingValueType, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *InitialValue) SetValue(v SettingValueType)`

SetValue sets Value field to given value.

### HasValue

`func (o *InitialValue) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *InitialValue) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *InitialValue) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetPredefinedVariationIndex

`func (o *InitialValue) GetPredefinedVariationIndex() int32`

GetPredefinedVariationIndex returns the PredefinedVariationIndex field if non-nil, zero value otherwise.

### GetPredefinedVariationIndexOk

`func (o *InitialValue) GetPredefinedVariationIndexOk() (*int32, bool)`

GetPredefinedVariationIndexOk returns a tuple with the PredefinedVariationIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredefinedVariationIndex

`func (o *InitialValue) SetPredefinedVariationIndex(v int32)`

SetPredefinedVariationIndex sets PredefinedVariationIndex field to given value.

### HasPredefinedVariationIndex

`func (o *InitialValue) HasPredefinedVariationIndex() bool`

HasPredefinedVariationIndex returns a boolean if a field has been set.

### SetPredefinedVariationIndexNil

`func (o *InitialValue) SetPredefinedVariationIndexNil(b bool)`

 SetPredefinedVariationIndexNil sets the value for PredefinedVariationIndex to be an explicit nil

### UnsetPredefinedVariationIndex
`func (o *InitialValue) UnsetPredefinedVariationIndex()`

UnsetPredefinedVariationIndex ensures that no value is present for PredefinedVariationIndex, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


