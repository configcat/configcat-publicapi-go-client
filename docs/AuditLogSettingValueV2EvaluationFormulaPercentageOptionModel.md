# AuditLogSettingValueV2EvaluationFormulaPercentageOptionModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Percentage** | **int32** | The percentage of users (0-100) who should receive the value. | 
**Value** | [**ValueModel**](ValueModel.md) |  | 
**ValuePredefinedVariationName** | **NullableString** | Optional name of the predefined variation for this percentage option. | 

## Methods

### NewAuditLogSettingValueV2EvaluationFormulaPercentageOptionModel

`func NewAuditLogSettingValueV2EvaluationFormulaPercentageOptionModel(percentage int32, value ValueModel, valuePredefinedVariationName NullableString, ) *AuditLogSettingValueV2EvaluationFormulaPercentageOptionModel`

NewAuditLogSettingValueV2EvaluationFormulaPercentageOptionModel instantiates a new AuditLogSettingValueV2EvaluationFormulaPercentageOptionModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditLogSettingValueV2EvaluationFormulaPercentageOptionModelWithDefaults

`func NewAuditLogSettingValueV2EvaluationFormulaPercentageOptionModelWithDefaults() *AuditLogSettingValueV2EvaluationFormulaPercentageOptionModel`

NewAuditLogSettingValueV2EvaluationFormulaPercentageOptionModelWithDefaults instantiates a new AuditLogSettingValueV2EvaluationFormulaPercentageOptionModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPercentage

`func (o *AuditLogSettingValueV2EvaluationFormulaPercentageOptionModel) GetPercentage() int32`

GetPercentage returns the Percentage field if non-nil, zero value otherwise.

### GetPercentageOk

`func (o *AuditLogSettingValueV2EvaluationFormulaPercentageOptionModel) GetPercentageOk() (*int32, bool)`

GetPercentageOk returns a tuple with the Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentage

`func (o *AuditLogSettingValueV2EvaluationFormulaPercentageOptionModel) SetPercentage(v int32)`

SetPercentage sets Percentage field to given value.


### GetValue

`func (o *AuditLogSettingValueV2EvaluationFormulaPercentageOptionModel) GetValue() ValueModel`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AuditLogSettingValueV2EvaluationFormulaPercentageOptionModel) GetValueOk() (*ValueModel, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AuditLogSettingValueV2EvaluationFormulaPercentageOptionModel) SetValue(v ValueModel)`

SetValue sets Value field to given value.


### GetValuePredefinedVariationName

`func (o *AuditLogSettingValueV2EvaluationFormulaPercentageOptionModel) GetValuePredefinedVariationName() string`

GetValuePredefinedVariationName returns the ValuePredefinedVariationName field if non-nil, zero value otherwise.

### GetValuePredefinedVariationNameOk

`func (o *AuditLogSettingValueV2EvaluationFormulaPercentageOptionModel) GetValuePredefinedVariationNameOk() (*string, bool)`

GetValuePredefinedVariationNameOk returns a tuple with the ValuePredefinedVariationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValuePredefinedVariationName

`func (o *AuditLogSettingValueV2EvaluationFormulaPercentageOptionModel) SetValuePredefinedVariationName(v string)`

SetValuePredefinedVariationName sets ValuePredefinedVariationName field to given value.


### SetValuePredefinedVariationNameNil

`func (o *AuditLogSettingValueV2EvaluationFormulaPercentageOptionModel) SetValuePredefinedVariationNameNil(b bool)`

 SetValuePredefinedVariationNameNil sets the value for ValuePredefinedVariationName to be an explicit nil

### UnsetValuePredefinedVariationName
`func (o *AuditLogSettingValueV2EvaluationFormulaPercentageOptionModel) UnsetValuePredefinedVariationName()`

UnsetValuePredefinedVariationName ensures that no value is present for ValuePredefinedVariationName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


