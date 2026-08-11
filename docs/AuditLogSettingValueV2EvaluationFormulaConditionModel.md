# AuditLogSettingValueV2EvaluationFormulaConditionModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConditionType** | [**ConditionType**](ConditionType.md) |  | 
**ComparisonAttribute** | **NullableString** | The name of the user attribute to compare. (For User Conditions.) | 
**UserComparator** | [**NullableUserComparator**](UserComparator.md) |  | 
**ComparisonValue** | [**NullableComparisonValueModel**](ComparisonValueModel.md) |  | 
**SegmentComparator** | [**NullableSegmentComparator**](SegmentComparator.md) |  | 
**SegmentId** | **NullableString** | The ID of the segment referenced in this condition. (For Segment Conditions.) | 
**SegmentName** | **NullableString** | The name of the segment referenced in this condition. (For Segment Conditions.) | 
**PrerequisiteComparator** | [**NullablePrerequisiteComparator**](PrerequisiteComparator.md) |  | 
**PrerequisiteSettingId** | **NullableInt32** | The ID of the prerequisite setting referenced. (For Flag Conditions.) | 
**PrerequisiteSettingKey** | **NullableString** | The key identifier of the prerequisite setting. (For Flag Conditions.) | 
**PrerequisiteComparisonValue** | [**NullableValueModel**](ValueModel.md) |  | 
**PrerequisiteComparisonValuePredefinedVariationName** | **NullableString** | Optional name of the predefined variation for the prerequisite comparison value. (For Flag Conditions.) | 

## Methods

### NewAuditLogSettingValueV2EvaluationFormulaConditionModel

`func NewAuditLogSettingValueV2EvaluationFormulaConditionModel(conditionType ConditionType, comparisonAttribute NullableString, userComparator NullableUserComparator, comparisonValue NullableComparisonValueModel, segmentComparator NullableSegmentComparator, segmentId NullableString, segmentName NullableString, prerequisiteComparator NullablePrerequisiteComparator, prerequisiteSettingId NullableInt32, prerequisiteSettingKey NullableString, prerequisiteComparisonValue NullableValueModel, prerequisiteComparisonValuePredefinedVariationName NullableString, ) *AuditLogSettingValueV2EvaluationFormulaConditionModel`

NewAuditLogSettingValueV2EvaluationFormulaConditionModel instantiates a new AuditLogSettingValueV2EvaluationFormulaConditionModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditLogSettingValueV2EvaluationFormulaConditionModelWithDefaults

`func NewAuditLogSettingValueV2EvaluationFormulaConditionModelWithDefaults() *AuditLogSettingValueV2EvaluationFormulaConditionModel`

NewAuditLogSettingValueV2EvaluationFormulaConditionModelWithDefaults instantiates a new AuditLogSettingValueV2EvaluationFormulaConditionModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditionType

`func (o *AuditLogSettingValueV2EvaluationFormulaConditionModel) GetConditionType() ConditionType`

GetConditionType returns the ConditionType field if non-nil, zero value otherwise.

### GetConditionTypeOk

`func (o *AuditLogSettingValueV2EvaluationFormulaConditionModel) GetConditionTypeOk() (*ConditionType, bool)`

GetConditionTypeOk returns a tuple with the ConditionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionType

`func (o *AuditLogSettingValueV2EvaluationFormulaConditionModel) SetConditionType(v ConditionType)`

SetConditionType sets ConditionType field to given value.


### GetComparisonAttribute

`func (o *AuditLogSettingValueV2EvaluationFormulaConditionModel) GetComparisonAttribute() string`

GetComparisonAttribute returns the ComparisonAttribute field if non-nil, zero value otherwise.

### GetComparisonAttributeOk

`func (o *AuditLogSettingValueV2EvaluationFormulaConditionModel) GetComparisonAttributeOk() (*string, bool)`

GetComparisonAttributeOk returns a tuple with the ComparisonAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparisonAttribute

`func (o *AuditLogSettingValueV2EvaluationFormulaConditionModel) SetComparisonAttribute(v string)`

SetComparisonAttribute sets ComparisonAttribute field to given value.


### SetComparisonAttributeNil

`func (o *AuditLogSettingValueV2EvaluationFormulaConditionModel) SetComparisonAttributeNil(b bool)`

 SetComparisonAttributeNil sets the value for ComparisonAttribute to be an explicit nil

### UnsetComparisonAttribute
`func (o *AuditLogSettingValueV2EvaluationFormulaConditionModel) UnsetComparisonAttribute()`

UnsetComparisonAttribute ensures that no value is present for ComparisonAttribute, not even an explicit nil
### GetUserComparator

`func (o *AuditLogSettingValueV2EvaluationFormulaConditionModel) GetUserComparator() UserComparator`

GetUserComparator returns the UserComparator field if non-nil, zero value otherwise.

### GetUserComparatorOk

`func (o *AuditLogSettingValueV2EvaluationFormulaConditionModel) GetUserComparatorOk() (*UserComparator, bool)`

GetUserComparatorOk returns a tuple with the UserComparator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserComparator

`func (o *AuditLogSettingValueV2EvaluationFormulaConditionModel) SetUserComparator(v UserComparator)`

SetUserComparator sets UserComparator field to given value.


### SetUserComparatorNil

`func (o *AuditLogSettingValueV2EvaluationFormulaConditionModel) SetUserComparatorNil(b bool)`

 SetUserComparatorNil sets the value for UserComparator to be an explicit nil

### UnsetUserComparator
`func (o *AuditLogSettingValueV2EvaluationFormulaConditionModel) UnsetUserComparator()`

UnsetUserComparator ensures that no value is present for UserComparator, not even an explicit nil
### GetComparisonValue

`func (o *AuditLogSettingValueV2EvaluationFormulaConditionModel) GetComparisonValue() ComparisonValueModel`

GetComparisonValue returns the ComparisonValue field if non-nil, zero value otherwise.

### GetComparisonValueOk

`func (o *AuditLogSettingValueV2EvaluationFormulaConditionModel) GetComparisonValueOk() (*ComparisonValueModel, bool)`

GetComparisonValueOk returns a tuple with the ComparisonValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparisonValue

`func (o *AuditLogSettingValueV2EvaluationFormulaConditionModel) SetComparisonValue(v ComparisonValueModel)`

SetComparisonValue sets ComparisonValue field to given value.


### SetComparisonValueNil

`func (o *AuditLogSettingValueV2EvaluationFormulaConditionModel) SetComparisonValueNil(b bool)`

 SetComparisonValueNil sets the value for ComparisonValue to be an explicit nil

### UnsetComparisonValue
`func (o *AuditLogSettingValueV2EvaluationFormulaConditionModel) UnsetComparisonValue()`

UnsetComparisonValue ensures that no value is present for ComparisonValue, not even an explicit nil
### GetSegmentComparator

`func (o *AuditLogSettingValueV2EvaluationFormulaConditionModel) GetSegmentComparator() SegmentComparator`

GetSegmentComparator returns the SegmentComparator field if non-nil, zero value otherwise.

### GetSegmentComparatorOk

`func (o *AuditLogSettingValueV2EvaluationFormulaConditionModel) GetSegmentComparatorOk() (*SegmentComparator, bool)`

GetSegmentComparatorOk returns a tuple with the SegmentComparator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentComparator

`func (o *AuditLogSettingValueV2EvaluationFormulaConditionModel) SetSegmentComparator(v SegmentComparator)`

SetSegmentComparator sets SegmentComparator field to given value.


### SetSegmentComparatorNil

`func (o *AuditLogSettingValueV2EvaluationFormulaConditionModel) SetSegmentComparatorNil(b bool)`

 SetSegmentComparatorNil sets the value for SegmentComparator to be an explicit nil

### UnsetSegmentComparator
`func (o *AuditLogSettingValueV2EvaluationFormulaConditionModel) UnsetSegmentComparator()`

UnsetSegmentComparator ensures that no value is present for SegmentComparator, not even an explicit nil
### GetSegmentId

`func (o *AuditLogSettingValueV2EvaluationFormulaConditionModel) GetSegmentId() string`

GetSegmentId returns the SegmentId field if non-nil, zero value otherwise.

### GetSegmentIdOk

`func (o *AuditLogSettingValueV2EvaluationFormulaConditionModel) GetSegmentIdOk() (*string, bool)`

GetSegmentIdOk returns a tuple with the SegmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentId

`func (o *AuditLogSettingValueV2EvaluationFormulaConditionModel) SetSegmentId(v string)`

SetSegmentId sets SegmentId field to given value.


### SetSegmentIdNil

`func (o *AuditLogSettingValueV2EvaluationFormulaConditionModel) SetSegmentIdNil(b bool)`

 SetSegmentIdNil sets the value for SegmentId to be an explicit nil

### UnsetSegmentId
`func (o *AuditLogSettingValueV2EvaluationFormulaConditionModel) UnsetSegmentId()`

UnsetSegmentId ensures that no value is present for SegmentId, not even an explicit nil
### GetSegmentName

`func (o *AuditLogSettingValueV2EvaluationFormulaConditionModel) GetSegmentName() string`

GetSegmentName returns the SegmentName field if non-nil, zero value otherwise.

### GetSegmentNameOk

`func (o *AuditLogSettingValueV2EvaluationFormulaConditionModel) GetSegmentNameOk() (*string, bool)`

GetSegmentNameOk returns a tuple with the SegmentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentName

`func (o *AuditLogSettingValueV2EvaluationFormulaConditionModel) SetSegmentName(v string)`

SetSegmentName sets SegmentName field to given value.


### SetSegmentNameNil

`func (o *AuditLogSettingValueV2EvaluationFormulaConditionModel) SetSegmentNameNil(b bool)`

 SetSegmentNameNil sets the value for SegmentName to be an explicit nil

### UnsetSegmentName
`func (o *AuditLogSettingValueV2EvaluationFormulaConditionModel) UnsetSegmentName()`

UnsetSegmentName ensures that no value is present for SegmentName, not even an explicit nil
### GetPrerequisiteComparator

`func (o *AuditLogSettingValueV2EvaluationFormulaConditionModel) GetPrerequisiteComparator() PrerequisiteComparator`

GetPrerequisiteComparator returns the PrerequisiteComparator field if non-nil, zero value otherwise.

### GetPrerequisiteComparatorOk

`func (o *AuditLogSettingValueV2EvaluationFormulaConditionModel) GetPrerequisiteComparatorOk() (*PrerequisiteComparator, bool)`

GetPrerequisiteComparatorOk returns a tuple with the PrerequisiteComparator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrerequisiteComparator

`func (o *AuditLogSettingValueV2EvaluationFormulaConditionModel) SetPrerequisiteComparator(v PrerequisiteComparator)`

SetPrerequisiteComparator sets PrerequisiteComparator field to given value.


### SetPrerequisiteComparatorNil

`func (o *AuditLogSettingValueV2EvaluationFormulaConditionModel) SetPrerequisiteComparatorNil(b bool)`

 SetPrerequisiteComparatorNil sets the value for PrerequisiteComparator to be an explicit nil

### UnsetPrerequisiteComparator
`func (o *AuditLogSettingValueV2EvaluationFormulaConditionModel) UnsetPrerequisiteComparator()`

UnsetPrerequisiteComparator ensures that no value is present for PrerequisiteComparator, not even an explicit nil
### GetPrerequisiteSettingId

`func (o *AuditLogSettingValueV2EvaluationFormulaConditionModel) GetPrerequisiteSettingId() int32`

GetPrerequisiteSettingId returns the PrerequisiteSettingId field if non-nil, zero value otherwise.

### GetPrerequisiteSettingIdOk

`func (o *AuditLogSettingValueV2EvaluationFormulaConditionModel) GetPrerequisiteSettingIdOk() (*int32, bool)`

GetPrerequisiteSettingIdOk returns a tuple with the PrerequisiteSettingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrerequisiteSettingId

`func (o *AuditLogSettingValueV2EvaluationFormulaConditionModel) SetPrerequisiteSettingId(v int32)`

SetPrerequisiteSettingId sets PrerequisiteSettingId field to given value.


### SetPrerequisiteSettingIdNil

`func (o *AuditLogSettingValueV2EvaluationFormulaConditionModel) SetPrerequisiteSettingIdNil(b bool)`

 SetPrerequisiteSettingIdNil sets the value for PrerequisiteSettingId to be an explicit nil

### UnsetPrerequisiteSettingId
`func (o *AuditLogSettingValueV2EvaluationFormulaConditionModel) UnsetPrerequisiteSettingId()`

UnsetPrerequisiteSettingId ensures that no value is present for PrerequisiteSettingId, not even an explicit nil
### GetPrerequisiteSettingKey

`func (o *AuditLogSettingValueV2EvaluationFormulaConditionModel) GetPrerequisiteSettingKey() string`

GetPrerequisiteSettingKey returns the PrerequisiteSettingKey field if non-nil, zero value otherwise.

### GetPrerequisiteSettingKeyOk

`func (o *AuditLogSettingValueV2EvaluationFormulaConditionModel) GetPrerequisiteSettingKeyOk() (*string, bool)`

GetPrerequisiteSettingKeyOk returns a tuple with the PrerequisiteSettingKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrerequisiteSettingKey

`func (o *AuditLogSettingValueV2EvaluationFormulaConditionModel) SetPrerequisiteSettingKey(v string)`

SetPrerequisiteSettingKey sets PrerequisiteSettingKey field to given value.


### SetPrerequisiteSettingKeyNil

`func (o *AuditLogSettingValueV2EvaluationFormulaConditionModel) SetPrerequisiteSettingKeyNil(b bool)`

 SetPrerequisiteSettingKeyNil sets the value for PrerequisiteSettingKey to be an explicit nil

### UnsetPrerequisiteSettingKey
`func (o *AuditLogSettingValueV2EvaluationFormulaConditionModel) UnsetPrerequisiteSettingKey()`

UnsetPrerequisiteSettingKey ensures that no value is present for PrerequisiteSettingKey, not even an explicit nil
### GetPrerequisiteComparisonValue

`func (o *AuditLogSettingValueV2EvaluationFormulaConditionModel) GetPrerequisiteComparisonValue() ValueModel`

GetPrerequisiteComparisonValue returns the PrerequisiteComparisonValue field if non-nil, zero value otherwise.

### GetPrerequisiteComparisonValueOk

`func (o *AuditLogSettingValueV2EvaluationFormulaConditionModel) GetPrerequisiteComparisonValueOk() (*ValueModel, bool)`

GetPrerequisiteComparisonValueOk returns a tuple with the PrerequisiteComparisonValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrerequisiteComparisonValue

`func (o *AuditLogSettingValueV2EvaluationFormulaConditionModel) SetPrerequisiteComparisonValue(v ValueModel)`

SetPrerequisiteComparisonValue sets PrerequisiteComparisonValue field to given value.


### SetPrerequisiteComparisonValueNil

`func (o *AuditLogSettingValueV2EvaluationFormulaConditionModel) SetPrerequisiteComparisonValueNil(b bool)`

 SetPrerequisiteComparisonValueNil sets the value for PrerequisiteComparisonValue to be an explicit nil

### UnsetPrerequisiteComparisonValue
`func (o *AuditLogSettingValueV2EvaluationFormulaConditionModel) UnsetPrerequisiteComparisonValue()`

UnsetPrerequisiteComparisonValue ensures that no value is present for PrerequisiteComparisonValue, not even an explicit nil
### GetPrerequisiteComparisonValuePredefinedVariationName

`func (o *AuditLogSettingValueV2EvaluationFormulaConditionModel) GetPrerequisiteComparisonValuePredefinedVariationName() string`

GetPrerequisiteComparisonValuePredefinedVariationName returns the PrerequisiteComparisonValuePredefinedVariationName field if non-nil, zero value otherwise.

### GetPrerequisiteComparisonValuePredefinedVariationNameOk

`func (o *AuditLogSettingValueV2EvaluationFormulaConditionModel) GetPrerequisiteComparisonValuePredefinedVariationNameOk() (*string, bool)`

GetPrerequisiteComparisonValuePredefinedVariationNameOk returns a tuple with the PrerequisiteComparisonValuePredefinedVariationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrerequisiteComparisonValuePredefinedVariationName

`func (o *AuditLogSettingValueV2EvaluationFormulaConditionModel) SetPrerequisiteComparisonValuePredefinedVariationName(v string)`

SetPrerequisiteComparisonValuePredefinedVariationName sets PrerequisiteComparisonValuePredefinedVariationName field to given value.


### SetPrerequisiteComparisonValuePredefinedVariationNameNil

`func (o *AuditLogSettingValueV2EvaluationFormulaConditionModel) SetPrerequisiteComparisonValuePredefinedVariationNameNil(b bool)`

 SetPrerequisiteComparisonValuePredefinedVariationNameNil sets the value for PrerequisiteComparisonValuePredefinedVariationName to be an explicit nil

### UnsetPrerequisiteComparisonValuePredefinedVariationName
`func (o *AuditLogSettingValueV2EvaluationFormulaConditionModel) UnsetPrerequisiteComparisonValuePredefinedVariationName()`

UnsetPrerequisiteComparisonValuePredefinedVariationName ensures that no value is present for PrerequisiteComparisonValuePredefinedVariationName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


