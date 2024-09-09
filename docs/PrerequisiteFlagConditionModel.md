# PrerequisiteFlagConditionModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrerequisiteSettingId** | **int32** | The prerequisite flag&#39;s identifier. | 
**Comparator** | **string** | Prerequisite flag comparison operator used during the evaluation process. | 
**PrerequisiteComparisonValue** | [**ValueModel**](ValueModel.md) |  | 

## Methods

### NewPrerequisiteFlagConditionModel

`func NewPrerequisiteFlagConditionModel(prerequisiteSettingId int32, comparator string, prerequisiteComparisonValue ValueModel, ) *PrerequisiteFlagConditionModel`

NewPrerequisiteFlagConditionModel instantiates a new PrerequisiteFlagConditionModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrerequisiteFlagConditionModelWithDefaults

`func NewPrerequisiteFlagConditionModelWithDefaults() *PrerequisiteFlagConditionModel`

NewPrerequisiteFlagConditionModelWithDefaults instantiates a new PrerequisiteFlagConditionModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrerequisiteSettingId

`func (o *PrerequisiteFlagConditionModel) GetPrerequisiteSettingId() int32`

GetPrerequisiteSettingId returns the PrerequisiteSettingId field if non-nil, zero value otherwise.

### GetPrerequisiteSettingIdOk

`func (o *PrerequisiteFlagConditionModel) GetPrerequisiteSettingIdOk() (*int32, bool)`

GetPrerequisiteSettingIdOk returns a tuple with the PrerequisiteSettingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrerequisiteSettingId

`func (o *PrerequisiteFlagConditionModel) SetPrerequisiteSettingId(v int32)`

SetPrerequisiteSettingId sets PrerequisiteSettingId field to given value.


### GetComparator

`func (o *PrerequisiteFlagConditionModel) GetComparator() string`

GetComparator returns the Comparator field if non-nil, zero value otherwise.

### GetComparatorOk

`func (o *PrerequisiteFlagConditionModel) GetComparatorOk() (*string, bool)`

GetComparatorOk returns a tuple with the Comparator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparator

`func (o *PrerequisiteFlagConditionModel) SetComparator(v string)`

SetComparator sets Comparator field to given value.


### GetPrerequisiteComparisonValue

`func (o *PrerequisiteFlagConditionModel) GetPrerequisiteComparisonValue() ValueModel`

GetPrerequisiteComparisonValue returns the PrerequisiteComparisonValue field if non-nil, zero value otherwise.

### GetPrerequisiteComparisonValueOk

`func (o *PrerequisiteFlagConditionModel) GetPrerequisiteComparisonValueOk() (*ValueModel, bool)`

GetPrerequisiteComparisonValueOk returns a tuple with the PrerequisiteComparisonValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrerequisiteComparisonValue

`func (o *PrerequisiteFlagConditionModel) SetPrerequisiteComparisonValue(v ValueModel)`

SetPrerequisiteComparisonValue sets PrerequisiteComparisonValue field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


