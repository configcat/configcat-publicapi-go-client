# UpdatePrerequisiteFlagConditionModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrerequisiteSettingId** | **int32** | The prerequisite flag&#39;s identifier. | 
**Comparator** | [**PrerequisiteComparator**](PrerequisiteComparator.md) |  | 
**PrerequisiteComparisonValue** | [**UpdateValueModel**](UpdateValueModel.md) |  | 

## Methods

### NewUpdatePrerequisiteFlagConditionModel

`func NewUpdatePrerequisiteFlagConditionModel(prerequisiteSettingId int32, comparator PrerequisiteComparator, prerequisiteComparisonValue UpdateValueModel, ) *UpdatePrerequisiteFlagConditionModel`

NewUpdatePrerequisiteFlagConditionModel instantiates a new UpdatePrerequisiteFlagConditionModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePrerequisiteFlagConditionModelWithDefaults

`func NewUpdatePrerequisiteFlagConditionModelWithDefaults() *UpdatePrerequisiteFlagConditionModel`

NewUpdatePrerequisiteFlagConditionModelWithDefaults instantiates a new UpdatePrerequisiteFlagConditionModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrerequisiteSettingId

`func (o *UpdatePrerequisiteFlagConditionModel) GetPrerequisiteSettingId() int32`

GetPrerequisiteSettingId returns the PrerequisiteSettingId field if non-nil, zero value otherwise.

### GetPrerequisiteSettingIdOk

`func (o *UpdatePrerequisiteFlagConditionModel) GetPrerequisiteSettingIdOk() (*int32, bool)`

GetPrerequisiteSettingIdOk returns a tuple with the PrerequisiteSettingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrerequisiteSettingId

`func (o *UpdatePrerequisiteFlagConditionModel) SetPrerequisiteSettingId(v int32)`

SetPrerequisiteSettingId sets PrerequisiteSettingId field to given value.


### GetComparator

`func (o *UpdatePrerequisiteFlagConditionModel) GetComparator() PrerequisiteComparator`

GetComparator returns the Comparator field if non-nil, zero value otherwise.

### GetComparatorOk

`func (o *UpdatePrerequisiteFlagConditionModel) GetComparatorOk() (*PrerequisiteComparator, bool)`

GetComparatorOk returns a tuple with the Comparator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparator

`func (o *UpdatePrerequisiteFlagConditionModel) SetComparator(v PrerequisiteComparator)`

SetComparator sets Comparator field to given value.


### GetPrerequisiteComparisonValue

`func (o *UpdatePrerequisiteFlagConditionModel) GetPrerequisiteComparisonValue() UpdateValueModel`

GetPrerequisiteComparisonValue returns the PrerequisiteComparisonValue field if non-nil, zero value otherwise.

### GetPrerequisiteComparisonValueOk

`func (o *UpdatePrerequisiteFlagConditionModel) GetPrerequisiteComparisonValueOk() (*UpdateValueModel, bool)`

GetPrerequisiteComparisonValueOk returns a tuple with the PrerequisiteComparisonValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrerequisiteComparisonValue

`func (o *UpdatePrerequisiteFlagConditionModel) SetPrerequisiteComparisonValue(v UpdateValueModel)`

SetPrerequisiteComparisonValue sets PrerequisiteComparisonValue field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


