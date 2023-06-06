# InitialValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentId** | Pointer to **string** | The ID of the Environment where the initial value should be set. | [optional] 
**Value** | Pointer to **interface{}** | The initial value in the given Environment. It must respect the setting type. | [optional] 

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

`func (o *InitialValue) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *InitialValue) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *InitialValue) SetValue(v interface{})`

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

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


