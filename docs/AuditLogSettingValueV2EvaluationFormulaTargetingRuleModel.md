# AuditLogSettingValueV2EvaluationFormulaTargetingRuleModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | [**NullableValueModel**](ValueModel.md) |  | 
**ValuePredefinedVariationName** | **NullableString** | Optional name of the predefined variation for the value. | 
**Conditions** | [**[]AuditLogSettingValueV2EvaluationFormulaConditionModel**](AuditLogSettingValueV2EvaluationFormulaConditionModel.md) | List of conditions that must be satisfied for this targeting rule to apply. | 
**PercentageOptions** | [**[]AuditLogSettingValueV2EvaluationFormulaPercentageOptionModel**](AuditLogSettingValueV2EvaluationFormulaPercentageOptionModel.md) | List of percentage options when percentage-based user bucketing (for A/B testing) is used in the targeting rule. | 

## Methods

### NewAuditLogSettingValueV2EvaluationFormulaTargetingRuleModel

`func NewAuditLogSettingValueV2EvaluationFormulaTargetingRuleModel(value NullableValueModel, valuePredefinedVariationName NullableString, conditions []AuditLogSettingValueV2EvaluationFormulaConditionModel, percentageOptions []AuditLogSettingValueV2EvaluationFormulaPercentageOptionModel, ) *AuditLogSettingValueV2EvaluationFormulaTargetingRuleModel`

NewAuditLogSettingValueV2EvaluationFormulaTargetingRuleModel instantiates a new AuditLogSettingValueV2EvaluationFormulaTargetingRuleModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditLogSettingValueV2EvaluationFormulaTargetingRuleModelWithDefaults

`func NewAuditLogSettingValueV2EvaluationFormulaTargetingRuleModelWithDefaults() *AuditLogSettingValueV2EvaluationFormulaTargetingRuleModel`

NewAuditLogSettingValueV2EvaluationFormulaTargetingRuleModelWithDefaults instantiates a new AuditLogSettingValueV2EvaluationFormulaTargetingRuleModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *AuditLogSettingValueV2EvaluationFormulaTargetingRuleModel) GetValue() ValueModel`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AuditLogSettingValueV2EvaluationFormulaTargetingRuleModel) GetValueOk() (*ValueModel, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AuditLogSettingValueV2EvaluationFormulaTargetingRuleModel) SetValue(v ValueModel)`

SetValue sets Value field to given value.


### SetValueNil

`func (o *AuditLogSettingValueV2EvaluationFormulaTargetingRuleModel) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *AuditLogSettingValueV2EvaluationFormulaTargetingRuleModel) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetValuePredefinedVariationName

`func (o *AuditLogSettingValueV2EvaluationFormulaTargetingRuleModel) GetValuePredefinedVariationName() string`

GetValuePredefinedVariationName returns the ValuePredefinedVariationName field if non-nil, zero value otherwise.

### GetValuePredefinedVariationNameOk

`func (o *AuditLogSettingValueV2EvaluationFormulaTargetingRuleModel) GetValuePredefinedVariationNameOk() (*string, bool)`

GetValuePredefinedVariationNameOk returns a tuple with the ValuePredefinedVariationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValuePredefinedVariationName

`func (o *AuditLogSettingValueV2EvaluationFormulaTargetingRuleModel) SetValuePredefinedVariationName(v string)`

SetValuePredefinedVariationName sets ValuePredefinedVariationName field to given value.


### SetValuePredefinedVariationNameNil

`func (o *AuditLogSettingValueV2EvaluationFormulaTargetingRuleModel) SetValuePredefinedVariationNameNil(b bool)`

 SetValuePredefinedVariationNameNil sets the value for ValuePredefinedVariationName to be an explicit nil

### UnsetValuePredefinedVariationName
`func (o *AuditLogSettingValueV2EvaluationFormulaTargetingRuleModel) UnsetValuePredefinedVariationName()`

UnsetValuePredefinedVariationName ensures that no value is present for ValuePredefinedVariationName, not even an explicit nil
### GetConditions

`func (o *AuditLogSettingValueV2EvaluationFormulaTargetingRuleModel) GetConditions() []AuditLogSettingValueV2EvaluationFormulaConditionModel`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *AuditLogSettingValueV2EvaluationFormulaTargetingRuleModel) GetConditionsOk() (*[]AuditLogSettingValueV2EvaluationFormulaConditionModel, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *AuditLogSettingValueV2EvaluationFormulaTargetingRuleModel) SetConditions(v []AuditLogSettingValueV2EvaluationFormulaConditionModel)`

SetConditions sets Conditions field to given value.


### GetPercentageOptions

`func (o *AuditLogSettingValueV2EvaluationFormulaTargetingRuleModel) GetPercentageOptions() []AuditLogSettingValueV2EvaluationFormulaPercentageOptionModel`

GetPercentageOptions returns the PercentageOptions field if non-nil, zero value otherwise.

### GetPercentageOptionsOk

`func (o *AuditLogSettingValueV2EvaluationFormulaTargetingRuleModel) GetPercentageOptionsOk() (*[]AuditLogSettingValueV2EvaluationFormulaPercentageOptionModel, bool)`

GetPercentageOptionsOk returns a tuple with the PercentageOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageOptions

`func (o *AuditLogSettingValueV2EvaluationFormulaTargetingRuleModel) SetPercentageOptions(v []AuditLogSettingValueV2EvaluationFormulaPercentageOptionModel)`

SetPercentageOptions sets PercentageOptions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


