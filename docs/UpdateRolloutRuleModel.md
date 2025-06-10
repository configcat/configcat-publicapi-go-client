# UpdateRolloutRuleModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComparisonAttribute** | Pointer to **NullableString** | The user attribute to compare. | [optional] 
**Comparator** | Pointer to [**NullableRolloutRuleComparator**](RolloutRuleComparator.md) |  | [optional] 
**ComparisonValue** | Pointer to **NullableString** | The value to compare against. | [optional] 
**Value** | [**SettingValueType**](SettingValueType.md) | The value to serve when the comparison matches. It must respect the setting type. In some generated clients for strictly typed languages you may use double/float properties to handle integer values. | 
**SegmentComparator** | Pointer to [**NullableSegmentComparator**](SegmentComparator.md) |  | [optional] 
**SegmentId** | Pointer to **NullableString** | The segment to compare against. | [optional] 

## Methods

### NewUpdateRolloutRuleModel

`func NewUpdateRolloutRuleModel(value SettingValueType, ) *UpdateRolloutRuleModel`

NewUpdateRolloutRuleModel instantiates a new UpdateRolloutRuleModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRolloutRuleModelWithDefaults

`func NewUpdateRolloutRuleModelWithDefaults() *UpdateRolloutRuleModel`

NewUpdateRolloutRuleModelWithDefaults instantiates a new UpdateRolloutRuleModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComparisonAttribute

`func (o *UpdateRolloutRuleModel) GetComparisonAttribute() string`

GetComparisonAttribute returns the ComparisonAttribute field if non-nil, zero value otherwise.

### GetComparisonAttributeOk

`func (o *UpdateRolloutRuleModel) GetComparisonAttributeOk() (*string, bool)`

GetComparisonAttributeOk returns a tuple with the ComparisonAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparisonAttribute

`func (o *UpdateRolloutRuleModel) SetComparisonAttribute(v string)`

SetComparisonAttribute sets ComparisonAttribute field to given value.

### HasComparisonAttribute

`func (o *UpdateRolloutRuleModel) HasComparisonAttribute() bool`

HasComparisonAttribute returns a boolean if a field has been set.

### SetComparisonAttributeNil

`func (o *UpdateRolloutRuleModel) SetComparisonAttributeNil(b bool)`

 SetComparisonAttributeNil sets the value for ComparisonAttribute to be an explicit nil

### UnsetComparisonAttribute
`func (o *UpdateRolloutRuleModel) UnsetComparisonAttribute()`

UnsetComparisonAttribute ensures that no value is present for ComparisonAttribute, not even an explicit nil
### GetComparator

`func (o *UpdateRolloutRuleModel) GetComparator() RolloutRuleComparator`

GetComparator returns the Comparator field if non-nil, zero value otherwise.

### GetComparatorOk

`func (o *UpdateRolloutRuleModel) GetComparatorOk() (*RolloutRuleComparator, bool)`

GetComparatorOk returns a tuple with the Comparator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparator

`func (o *UpdateRolloutRuleModel) SetComparator(v RolloutRuleComparator)`

SetComparator sets Comparator field to given value.

### HasComparator

`func (o *UpdateRolloutRuleModel) HasComparator() bool`

HasComparator returns a boolean if a field has been set.

### SetComparatorNil

`func (o *UpdateRolloutRuleModel) SetComparatorNil(b bool)`

 SetComparatorNil sets the value for Comparator to be an explicit nil

### UnsetComparator
`func (o *UpdateRolloutRuleModel) UnsetComparator()`

UnsetComparator ensures that no value is present for Comparator, not even an explicit nil
### GetComparisonValue

`func (o *UpdateRolloutRuleModel) GetComparisonValue() string`

GetComparisonValue returns the ComparisonValue field if non-nil, zero value otherwise.

### GetComparisonValueOk

`func (o *UpdateRolloutRuleModel) GetComparisonValueOk() (*string, bool)`

GetComparisonValueOk returns a tuple with the ComparisonValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparisonValue

`func (o *UpdateRolloutRuleModel) SetComparisonValue(v string)`

SetComparisonValue sets ComparisonValue field to given value.

### HasComparisonValue

`func (o *UpdateRolloutRuleModel) HasComparisonValue() bool`

HasComparisonValue returns a boolean if a field has been set.

### SetComparisonValueNil

`func (o *UpdateRolloutRuleModel) SetComparisonValueNil(b bool)`

 SetComparisonValueNil sets the value for ComparisonValue to be an explicit nil

### UnsetComparisonValue
`func (o *UpdateRolloutRuleModel) UnsetComparisonValue()`

UnsetComparisonValue ensures that no value is present for ComparisonValue, not even an explicit nil
### GetValue

`func (o *UpdateRolloutRuleModel) GetValue() SettingValueType`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UpdateRolloutRuleModel) GetValueOk() (*SettingValueType, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UpdateRolloutRuleModel) SetValue(v SettingValueType)`

SetValue sets Value field to given value.


### GetSegmentComparator

`func (o *UpdateRolloutRuleModel) GetSegmentComparator() SegmentComparator`

GetSegmentComparator returns the SegmentComparator field if non-nil, zero value otherwise.

### GetSegmentComparatorOk

`func (o *UpdateRolloutRuleModel) GetSegmentComparatorOk() (*SegmentComparator, bool)`

GetSegmentComparatorOk returns a tuple with the SegmentComparator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentComparator

`func (o *UpdateRolloutRuleModel) SetSegmentComparator(v SegmentComparator)`

SetSegmentComparator sets SegmentComparator field to given value.

### HasSegmentComparator

`func (o *UpdateRolloutRuleModel) HasSegmentComparator() bool`

HasSegmentComparator returns a boolean if a field has been set.

### SetSegmentComparatorNil

`func (o *UpdateRolloutRuleModel) SetSegmentComparatorNil(b bool)`

 SetSegmentComparatorNil sets the value for SegmentComparator to be an explicit nil

### UnsetSegmentComparator
`func (o *UpdateRolloutRuleModel) UnsetSegmentComparator()`

UnsetSegmentComparator ensures that no value is present for SegmentComparator, not even an explicit nil
### GetSegmentId

`func (o *UpdateRolloutRuleModel) GetSegmentId() string`

GetSegmentId returns the SegmentId field if non-nil, zero value otherwise.

### GetSegmentIdOk

`func (o *UpdateRolloutRuleModel) GetSegmentIdOk() (*string, bool)`

GetSegmentIdOk returns a tuple with the SegmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentId

`func (o *UpdateRolloutRuleModel) SetSegmentId(v string)`

SetSegmentId sets SegmentId field to given value.

### HasSegmentId

`func (o *UpdateRolloutRuleModel) HasSegmentId() bool`

HasSegmentId returns a boolean if a field has been set.

### SetSegmentIdNil

`func (o *UpdateRolloutRuleModel) SetSegmentIdNil(b bool)`

 SetSegmentIdNil sets the value for SegmentId to be an explicit nil

### UnsetSegmentId
`func (o *UpdateRolloutRuleModel) UnsetSegmentId()`

UnsetSegmentId ensures that no value is present for SegmentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


