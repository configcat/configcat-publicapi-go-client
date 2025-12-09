# OrganizationLimitations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxPercentageOptionCount** | **int32** | Maximum number of percentage options a Feature Flag or Setting can have within a targeting rule. | 
**MaxTargetingRuleCount** | **int32** | Maximum number of targeting rules a Feature Flag or Setting can have. | 
**MaxComparisonValueLength** | **int32** | Maximum length of a text comparison value. | 
**MaxComparisonValueListLength** | **int32** | Maximum item count of a list comparison value. | 
**MaxComparisonValueListItemLength** | **int32** | Maximum length of a list comparison value&#39;s item. | 
**MaxStringFlagValueLength** | **int32** | Maximum length of a text Setting&#39;s value. | 
**MaxConditionPerTargetingRuleCount** | **int32** | Maximum number of &#x60;AND&#x60; conditions a Feature Flag or Setting can have within a targeting rule. | 
**MaxPredefinedVariations** | **int32** | The maximum number of predefined variations allowed for a Feature Flag or Setting. | 

## Methods

### NewOrganizationLimitations

`func NewOrganizationLimitations(maxPercentageOptionCount int32, maxTargetingRuleCount int32, maxComparisonValueLength int32, maxComparisonValueListLength int32, maxComparisonValueListItemLength int32, maxStringFlagValueLength int32, maxConditionPerTargetingRuleCount int32, maxPredefinedVariations int32, ) *OrganizationLimitations`

NewOrganizationLimitations instantiates a new OrganizationLimitations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationLimitationsWithDefaults

`func NewOrganizationLimitationsWithDefaults() *OrganizationLimitations`

NewOrganizationLimitationsWithDefaults instantiates a new OrganizationLimitations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxPercentageOptionCount

`func (o *OrganizationLimitations) GetMaxPercentageOptionCount() int32`

GetMaxPercentageOptionCount returns the MaxPercentageOptionCount field if non-nil, zero value otherwise.

### GetMaxPercentageOptionCountOk

`func (o *OrganizationLimitations) GetMaxPercentageOptionCountOk() (*int32, bool)`

GetMaxPercentageOptionCountOk returns a tuple with the MaxPercentageOptionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPercentageOptionCount

`func (o *OrganizationLimitations) SetMaxPercentageOptionCount(v int32)`

SetMaxPercentageOptionCount sets MaxPercentageOptionCount field to given value.


### GetMaxTargetingRuleCount

`func (o *OrganizationLimitations) GetMaxTargetingRuleCount() int32`

GetMaxTargetingRuleCount returns the MaxTargetingRuleCount field if non-nil, zero value otherwise.

### GetMaxTargetingRuleCountOk

`func (o *OrganizationLimitations) GetMaxTargetingRuleCountOk() (*int32, bool)`

GetMaxTargetingRuleCountOk returns a tuple with the MaxTargetingRuleCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTargetingRuleCount

`func (o *OrganizationLimitations) SetMaxTargetingRuleCount(v int32)`

SetMaxTargetingRuleCount sets MaxTargetingRuleCount field to given value.


### GetMaxComparisonValueLength

`func (o *OrganizationLimitations) GetMaxComparisonValueLength() int32`

GetMaxComparisonValueLength returns the MaxComparisonValueLength field if non-nil, zero value otherwise.

### GetMaxComparisonValueLengthOk

`func (o *OrganizationLimitations) GetMaxComparisonValueLengthOk() (*int32, bool)`

GetMaxComparisonValueLengthOk returns a tuple with the MaxComparisonValueLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxComparisonValueLength

`func (o *OrganizationLimitations) SetMaxComparisonValueLength(v int32)`

SetMaxComparisonValueLength sets MaxComparisonValueLength field to given value.


### GetMaxComparisonValueListLength

`func (o *OrganizationLimitations) GetMaxComparisonValueListLength() int32`

GetMaxComparisonValueListLength returns the MaxComparisonValueListLength field if non-nil, zero value otherwise.

### GetMaxComparisonValueListLengthOk

`func (o *OrganizationLimitations) GetMaxComparisonValueListLengthOk() (*int32, bool)`

GetMaxComparisonValueListLengthOk returns a tuple with the MaxComparisonValueListLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxComparisonValueListLength

`func (o *OrganizationLimitations) SetMaxComparisonValueListLength(v int32)`

SetMaxComparisonValueListLength sets MaxComparisonValueListLength field to given value.


### GetMaxComparisonValueListItemLength

`func (o *OrganizationLimitations) GetMaxComparisonValueListItemLength() int32`

GetMaxComparisonValueListItemLength returns the MaxComparisonValueListItemLength field if non-nil, zero value otherwise.

### GetMaxComparisonValueListItemLengthOk

`func (o *OrganizationLimitations) GetMaxComparisonValueListItemLengthOk() (*int32, bool)`

GetMaxComparisonValueListItemLengthOk returns a tuple with the MaxComparisonValueListItemLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxComparisonValueListItemLength

`func (o *OrganizationLimitations) SetMaxComparisonValueListItemLength(v int32)`

SetMaxComparisonValueListItemLength sets MaxComparisonValueListItemLength field to given value.


### GetMaxStringFlagValueLength

`func (o *OrganizationLimitations) GetMaxStringFlagValueLength() int32`

GetMaxStringFlagValueLength returns the MaxStringFlagValueLength field if non-nil, zero value otherwise.

### GetMaxStringFlagValueLengthOk

`func (o *OrganizationLimitations) GetMaxStringFlagValueLengthOk() (*int32, bool)`

GetMaxStringFlagValueLengthOk returns a tuple with the MaxStringFlagValueLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStringFlagValueLength

`func (o *OrganizationLimitations) SetMaxStringFlagValueLength(v int32)`

SetMaxStringFlagValueLength sets MaxStringFlagValueLength field to given value.


### GetMaxConditionPerTargetingRuleCount

`func (o *OrganizationLimitations) GetMaxConditionPerTargetingRuleCount() int32`

GetMaxConditionPerTargetingRuleCount returns the MaxConditionPerTargetingRuleCount field if non-nil, zero value otherwise.

### GetMaxConditionPerTargetingRuleCountOk

`func (o *OrganizationLimitations) GetMaxConditionPerTargetingRuleCountOk() (*int32, bool)`

GetMaxConditionPerTargetingRuleCountOk returns a tuple with the MaxConditionPerTargetingRuleCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConditionPerTargetingRuleCount

`func (o *OrganizationLimitations) SetMaxConditionPerTargetingRuleCount(v int32)`

SetMaxConditionPerTargetingRuleCount sets MaxConditionPerTargetingRuleCount field to given value.


### GetMaxPredefinedVariations

`func (o *OrganizationLimitations) GetMaxPredefinedVariations() int32`

GetMaxPredefinedVariations returns the MaxPredefinedVariations field if non-nil, zero value otherwise.

### GetMaxPredefinedVariationsOk

`func (o *OrganizationLimitations) GetMaxPredefinedVariationsOk() (*int32, bool)`

GetMaxPredefinedVariationsOk returns a tuple with the MaxPredefinedVariations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPredefinedVariations

`func (o *OrganizationLimitations) SetMaxPredefinedVariations(v int32)`

SetMaxPredefinedVariations sets MaxPredefinedVariations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


