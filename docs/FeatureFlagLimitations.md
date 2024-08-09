# FeatureFlagLimitations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxPercentageOptionCount** | Pointer to **int32** | Maximum number of percentage options a Feature Flag or Setting can have within a targeting rule. | [optional] 
**MaxTargetingRuleCount** | Pointer to **int32** | Maximum number of targeting rules a Feature Flag or Setting can have. | [optional] 
**MaxComparisonValueLength** | Pointer to **int32** | Maximum length of a text comparison value. | [optional] 
**MaxComparisonValueListLength** | Pointer to **int32** | Maximum item count of a list comparison value. | [optional] 
**MaxComparisonValueListItemLength** | Pointer to **int32** | Maximum length of a list comparison value&#39;s item. | [optional] 
**MaxStringFlagValueLength** | Pointer to **int32** | Maximum length of a text Setting&#39;s value. | [optional] 
**MaxConditionPerTargetingRuleCount** | Pointer to **int32** | Maximum number of &#x60;AND&#x60; conditions a Feature Flag or Setting can have within a targeting rule. | [optional] 

## Methods

### NewFeatureFlagLimitations

`func NewFeatureFlagLimitations() *FeatureFlagLimitations`

NewFeatureFlagLimitations instantiates a new FeatureFlagLimitations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatureFlagLimitationsWithDefaults

`func NewFeatureFlagLimitationsWithDefaults() *FeatureFlagLimitations`

NewFeatureFlagLimitationsWithDefaults instantiates a new FeatureFlagLimitations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxPercentageOptionCount

`func (o *FeatureFlagLimitations) GetMaxPercentageOptionCount() int32`

GetMaxPercentageOptionCount returns the MaxPercentageOptionCount field if non-nil, zero value otherwise.

### GetMaxPercentageOptionCountOk

`func (o *FeatureFlagLimitations) GetMaxPercentageOptionCountOk() (*int32, bool)`

GetMaxPercentageOptionCountOk returns a tuple with the MaxPercentageOptionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPercentageOptionCount

`func (o *FeatureFlagLimitations) SetMaxPercentageOptionCount(v int32)`

SetMaxPercentageOptionCount sets MaxPercentageOptionCount field to given value.

### HasMaxPercentageOptionCount

`func (o *FeatureFlagLimitations) HasMaxPercentageOptionCount() bool`

HasMaxPercentageOptionCount returns a boolean if a field has been set.

### GetMaxTargetingRuleCount

`func (o *FeatureFlagLimitations) GetMaxTargetingRuleCount() int32`

GetMaxTargetingRuleCount returns the MaxTargetingRuleCount field if non-nil, zero value otherwise.

### GetMaxTargetingRuleCountOk

`func (o *FeatureFlagLimitations) GetMaxTargetingRuleCountOk() (*int32, bool)`

GetMaxTargetingRuleCountOk returns a tuple with the MaxTargetingRuleCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTargetingRuleCount

`func (o *FeatureFlagLimitations) SetMaxTargetingRuleCount(v int32)`

SetMaxTargetingRuleCount sets MaxTargetingRuleCount field to given value.

### HasMaxTargetingRuleCount

`func (o *FeatureFlagLimitations) HasMaxTargetingRuleCount() bool`

HasMaxTargetingRuleCount returns a boolean if a field has been set.

### GetMaxComparisonValueLength

`func (o *FeatureFlagLimitations) GetMaxComparisonValueLength() int32`

GetMaxComparisonValueLength returns the MaxComparisonValueLength field if non-nil, zero value otherwise.

### GetMaxComparisonValueLengthOk

`func (o *FeatureFlagLimitations) GetMaxComparisonValueLengthOk() (*int32, bool)`

GetMaxComparisonValueLengthOk returns a tuple with the MaxComparisonValueLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxComparisonValueLength

`func (o *FeatureFlagLimitations) SetMaxComparisonValueLength(v int32)`

SetMaxComparisonValueLength sets MaxComparisonValueLength field to given value.

### HasMaxComparisonValueLength

`func (o *FeatureFlagLimitations) HasMaxComparisonValueLength() bool`

HasMaxComparisonValueLength returns a boolean if a field has been set.

### GetMaxComparisonValueListLength

`func (o *FeatureFlagLimitations) GetMaxComparisonValueListLength() int32`

GetMaxComparisonValueListLength returns the MaxComparisonValueListLength field if non-nil, zero value otherwise.

### GetMaxComparisonValueListLengthOk

`func (o *FeatureFlagLimitations) GetMaxComparisonValueListLengthOk() (*int32, bool)`

GetMaxComparisonValueListLengthOk returns a tuple with the MaxComparisonValueListLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxComparisonValueListLength

`func (o *FeatureFlagLimitations) SetMaxComparisonValueListLength(v int32)`

SetMaxComparisonValueListLength sets MaxComparisonValueListLength field to given value.

### HasMaxComparisonValueListLength

`func (o *FeatureFlagLimitations) HasMaxComparisonValueListLength() bool`

HasMaxComparisonValueListLength returns a boolean if a field has been set.

### GetMaxComparisonValueListItemLength

`func (o *FeatureFlagLimitations) GetMaxComparisonValueListItemLength() int32`

GetMaxComparisonValueListItemLength returns the MaxComparisonValueListItemLength field if non-nil, zero value otherwise.

### GetMaxComparisonValueListItemLengthOk

`func (o *FeatureFlagLimitations) GetMaxComparisonValueListItemLengthOk() (*int32, bool)`

GetMaxComparisonValueListItemLengthOk returns a tuple with the MaxComparisonValueListItemLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxComparisonValueListItemLength

`func (o *FeatureFlagLimitations) SetMaxComparisonValueListItemLength(v int32)`

SetMaxComparisonValueListItemLength sets MaxComparisonValueListItemLength field to given value.

### HasMaxComparisonValueListItemLength

`func (o *FeatureFlagLimitations) HasMaxComparisonValueListItemLength() bool`

HasMaxComparisonValueListItemLength returns a boolean if a field has been set.

### GetMaxStringFlagValueLength

`func (o *FeatureFlagLimitations) GetMaxStringFlagValueLength() int32`

GetMaxStringFlagValueLength returns the MaxStringFlagValueLength field if non-nil, zero value otherwise.

### GetMaxStringFlagValueLengthOk

`func (o *FeatureFlagLimitations) GetMaxStringFlagValueLengthOk() (*int32, bool)`

GetMaxStringFlagValueLengthOk returns a tuple with the MaxStringFlagValueLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStringFlagValueLength

`func (o *FeatureFlagLimitations) SetMaxStringFlagValueLength(v int32)`

SetMaxStringFlagValueLength sets MaxStringFlagValueLength field to given value.

### HasMaxStringFlagValueLength

`func (o *FeatureFlagLimitations) HasMaxStringFlagValueLength() bool`

HasMaxStringFlagValueLength returns a boolean if a field has been set.

### GetMaxConditionPerTargetingRuleCount

`func (o *FeatureFlagLimitations) GetMaxConditionPerTargetingRuleCount() int32`

GetMaxConditionPerTargetingRuleCount returns the MaxConditionPerTargetingRuleCount field if non-nil, zero value otherwise.

### GetMaxConditionPerTargetingRuleCountOk

`func (o *FeatureFlagLimitations) GetMaxConditionPerTargetingRuleCountOk() (*int32, bool)`

GetMaxConditionPerTargetingRuleCountOk returns a tuple with the MaxConditionPerTargetingRuleCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConditionPerTargetingRuleCount

`func (o *FeatureFlagLimitations) SetMaxConditionPerTargetingRuleCount(v int32)`

SetMaxConditionPerTargetingRuleCount sets MaxConditionPerTargetingRuleCount field to given value.

### HasMaxConditionPerTargetingRuleCount

`func (o *FeatureFlagLimitations) HasMaxConditionPerTargetingRuleCount() bool`

HasMaxConditionPerTargetingRuleCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


