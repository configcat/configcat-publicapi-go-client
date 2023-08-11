# ComparisonValueListModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | **string** | The actual comparison value. | 
**Hint** | Pointer to **NullableString** | An optional hint for the comparison value. | [optional] 

## Methods

### NewComparisonValueListModel

`func NewComparisonValueListModel(value string, ) *ComparisonValueListModel`

NewComparisonValueListModel instantiates a new ComparisonValueListModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComparisonValueListModelWithDefaults

`func NewComparisonValueListModelWithDefaults() *ComparisonValueListModel`

NewComparisonValueListModelWithDefaults instantiates a new ComparisonValueListModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *ComparisonValueListModel) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ComparisonValueListModel) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ComparisonValueListModel) SetValue(v string)`

SetValue sets Value field to given value.


### GetHint

`func (o *ComparisonValueListModel) GetHint() string`

GetHint returns the Hint field if non-nil, zero value otherwise.

### GetHintOk

`func (o *ComparisonValueListModel) GetHintOk() (*string, bool)`

GetHintOk returns a tuple with the Hint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHint

`func (o *ComparisonValueListModel) SetHint(v string)`

SetHint sets Hint field to given value.

### HasHint

`func (o *ComparisonValueListModel) HasHint() bool`

HasHint returns a boolean if a field has been set.

### SetHintNil

`func (o *ComparisonValueListModel) SetHintNil(b bool)`

 SetHintNil sets the value for Hint to be an explicit nil

### UnsetHint
`func (o *ComparisonValueListModel) UnsetHint()`

UnsetHint ensures that no value is present for Hint, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


