# RolloutRuleModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComparisonAttribute** | **NullableString** | The user attribute to compare. | 
**Comparator** | [**NullableRolloutRuleComparator**](RolloutRuleComparator.md) |  | 
**ComparisonValue** | **NullableString** | The value to compare against. | 
**Value** | [**SettingValueType**](SettingValueType.md) | The value to serve when the comparison matches. It must respect the setting type. In some generated clients for strictly typed languages you may use double/float properties to handle integer values. | 
**SegmentComparator** | [**NullableSegmentComparator**](SegmentComparator.md) |  | 
**SegmentId** | **NullableString** | The segment to compare against. | 

## Methods

### NewRolloutRuleModel

`func NewRolloutRuleModel(comparisonAttribute NullableString, comparator NullableRolloutRuleComparator, comparisonValue NullableString, value SettingValueType, segmentComparator NullableSegmentComparator, segmentId NullableString, ) *RolloutRuleModel`

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


### SetComparisonAttributeNil

`func (o *RolloutRuleModel) SetComparisonAttributeNil(b bool)`

 SetComparisonAttributeNil sets the value for ComparisonAttribute to be an explicit nil

### UnsetComparisonAttribute
`func (o *RolloutRuleModel) UnsetComparisonAttribute()`

UnsetComparisonAttribute ensures that no value is present for ComparisonAttribute, not even an explicit nil
### GetComparator

`func (o *RolloutRuleModel) GetComparator() RolloutRuleComparator`

GetComparator returns the Comparator field if non-nil, zero value otherwise.

### GetComparatorOk

`func (o *RolloutRuleModel) GetComparatorOk() (*RolloutRuleComparator, bool)`

GetComparatorOk returns a tuple with the Comparator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparator

`func (o *RolloutRuleModel) SetComparator(v RolloutRuleComparator)`

SetComparator sets Comparator field to given value.


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


### SetComparisonValueNil

`func (o *RolloutRuleModel) SetComparisonValueNil(b bool)`

 SetComparisonValueNil sets the value for ComparisonValue to be an explicit nil

### UnsetComparisonValue
`func (o *RolloutRuleModel) UnsetComparisonValue()`

UnsetComparisonValue ensures that no value is present for ComparisonValue, not even an explicit nil
### GetValue

`func (o *RolloutRuleModel) GetValue() SettingValueType`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RolloutRuleModel) GetValueOk() (*SettingValueType, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *RolloutRuleModel) SetValue(v SettingValueType)`

SetValue sets Value field to given value.


### GetSegmentComparator

`func (o *RolloutRuleModel) GetSegmentComparator() SegmentComparator`

GetSegmentComparator returns the SegmentComparator field if non-nil, zero value otherwise.

### GetSegmentComparatorOk

`func (o *RolloutRuleModel) GetSegmentComparatorOk() (*SegmentComparator, bool)`

GetSegmentComparatorOk returns a tuple with the SegmentComparator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentComparator

`func (o *RolloutRuleModel) SetSegmentComparator(v SegmentComparator)`

SetSegmentComparator sets SegmentComparator field to given value.


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


### SetSegmentIdNil

`func (o *RolloutRuleModel) SetSegmentIdNil(b bool)`

 SetSegmentIdNil sets the value for SegmentId to be an explicit nil

### UnsetSegmentId
`func (o *RolloutRuleModel) UnsetSegmentId()`

UnsetSegmentId ensures that no value is present for SegmentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


