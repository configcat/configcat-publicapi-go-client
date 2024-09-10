# RolloutRuleModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComparisonAttribute** | Pointer to **NullableString** | The user attribute to compare. | [optional] 
**Comparator** | Pointer to **NullableString** |  | [optional] 
**ComparisonValue** | Pointer to **NullableString** | The value to compare against. | [optional] 
**Value** | Pointer to **interface{}** | The value to serve when the comparison matches. It must respect the setting type. | [optional] 
**SegmentComparator** | Pointer to **NullableString** | The segment comparison operator. | [optional] 
**SegmentId** | Pointer to **NullableString** | The segment to compare against. | [optional] 

## Methods

### NewRolloutRuleModel

`func NewRolloutRuleModel() *RolloutRuleModel`

NewRolloutRuleModel instantiates a new RolloutRuleModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRolloutRuleModelWithDefaults

`func NewRolloutRuleModelWithDefaults() *RolloutRuleModel`

NewRolloutRuleModelWithDefaults instantiates a new RolloutRuleModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComparisonAttribute

`func (o *RolloutRuleModel) GetComparisonAttribute() string`

GetComparisonAttribute returns the ComparisonAttribute field if non-nil, zero value otherwise.

### GetComparisonAttributeOk

`func (o *RolloutRuleModel) GetComparisonAttributeOk() (*string, bool)`

GetComparisonAttributeOk returns a tuple with the ComparisonAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparisonAttribute

`func (o *RolloutRuleModel) SetComparisonAttribute(v string)`

SetComparisonAttribute sets ComparisonAttribute field to given value.

### HasComparisonAttribute

`func (o *RolloutRuleModel) HasComparisonAttribute() bool`

HasComparisonAttribute returns a boolean if a field has been set.

### SetComparisonAttributeNil

`func (o *RolloutRuleModel) SetComparisonAttributeNil(b bool)`

 SetComparisonAttributeNil sets the value for ComparisonAttribute to be an explicit nil

### UnsetComparisonAttribute
`func (o *RolloutRuleModel) UnsetComparisonAttribute()`

UnsetComparisonAttribute ensures that no value is present for ComparisonAttribute, not even an explicit nil
### GetComparator

`func (o *RolloutRuleModel) GetComparator() string`

GetComparator returns the Comparator field if non-nil, zero value otherwise.

### GetComparatorOk

`func (o *RolloutRuleModel) GetComparatorOk() (*string, bool)`

GetComparatorOk returns a tuple with the Comparator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparator

`func (o *RolloutRuleModel) SetComparator(v string)`

SetComparator sets Comparator field to given value.

### HasComparator

`func (o *RolloutRuleModel) HasComparator() bool`

HasComparator returns a boolean if a field has been set.

### SetComparatorNil

`func (o *RolloutRuleModel) SetComparatorNil(b bool)`

 SetComparatorNil sets the value for Comparator to be an explicit nil

### UnsetComparator
`func (o *RolloutRuleModel) UnsetComparator()`

UnsetComparator ensures that no value is present for Comparator, not even an explicit nil
### GetComparisonValue

`func (o *RolloutRuleModel) GetComparisonValue() string`

GetComparisonValue returns the ComparisonValue field if non-nil, zero value otherwise.

### GetComparisonValueOk

`func (o *RolloutRuleModel) GetComparisonValueOk() (*string, bool)`

GetComparisonValueOk returns a tuple with the ComparisonValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparisonValue

`func (o *RolloutRuleModel) SetComparisonValue(v string)`

SetComparisonValue sets ComparisonValue field to given value.

### HasComparisonValue

`func (o *RolloutRuleModel) HasComparisonValue() bool`

HasComparisonValue returns a boolean if a field has been set.

### SetComparisonValueNil

`func (o *RolloutRuleModel) SetComparisonValueNil(b bool)`

 SetComparisonValueNil sets the value for ComparisonValue to be an explicit nil

### UnsetComparisonValue
`func (o *RolloutRuleModel) UnsetComparisonValue()`

UnsetComparisonValue ensures that no value is present for ComparisonValue, not even an explicit nil
### GetValue

`func (o *RolloutRuleModel) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RolloutRuleModel) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *RolloutRuleModel) SetValue(v interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *RolloutRuleModel) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *RolloutRuleModel) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *RolloutRuleModel) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetSegmentComparator

`func (o *RolloutRuleModel) GetSegmentComparator() string`

GetSegmentComparator returns the SegmentComparator field if non-nil, zero value otherwise.

### GetSegmentComparatorOk

`func (o *RolloutRuleModel) GetSegmentComparatorOk() (*string, bool)`

GetSegmentComparatorOk returns a tuple with the SegmentComparator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentComparator

`func (o *RolloutRuleModel) SetSegmentComparator(v string)`

SetSegmentComparator sets SegmentComparator field to given value.

### HasSegmentComparator

`func (o *RolloutRuleModel) HasSegmentComparator() bool`

HasSegmentComparator returns a boolean if a field has been set.

### SetSegmentComparatorNil

`func (o *RolloutRuleModel) SetSegmentComparatorNil(b bool)`

 SetSegmentComparatorNil sets the value for SegmentComparator to be an explicit nil

### UnsetSegmentComparator
`func (o *RolloutRuleModel) UnsetSegmentComparator()`

UnsetSegmentComparator ensures that no value is present for SegmentComparator, not even an explicit nil
### GetSegmentId

`func (o *RolloutRuleModel) GetSegmentId() string`

GetSegmentId returns the SegmentId field if non-nil, zero value otherwise.

### GetSegmentIdOk

`func (o *RolloutRuleModel) GetSegmentIdOk() (*string, bool)`

GetSegmentIdOk returns a tuple with the SegmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentId

`func (o *RolloutRuleModel) SetSegmentId(v string)`

SetSegmentId sets SegmentId field to given value.

### HasSegmentId

`func (o *RolloutRuleModel) HasSegmentId() bool`

HasSegmentId returns a boolean if a field has been set.

### SetSegmentIdNil

`func (o *RolloutRuleModel) SetSegmentIdNil(b bool)`

 SetSegmentIdNil sets the value for SegmentId to be an explicit nil

### UnsetSegmentId
`func (o *RolloutRuleModel) UnsetSegmentId()`

UnsetSegmentId ensures that no value is present for SegmentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


