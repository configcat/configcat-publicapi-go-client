# InitialValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentId** | Pointer to **string** | The ID of the Environment where the initial value must be set. | [optional] 
**Value** | [**SettingValueType**](SettingValueType.md) | The initial value in the given Environment. It must respect the setting type. In some generated clients for strictly typed languages you may use double/float properties to handle integer values. | 

## Methods

### NewInitialValue

`func NewInitialValue(value SettingValueType, ) *InitialValue`

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


