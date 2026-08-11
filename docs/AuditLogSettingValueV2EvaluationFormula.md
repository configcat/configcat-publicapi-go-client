# AuditLogSettingValueV2EvaluationFormula

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultValue** | [**ValueModel**](ValueModel.md) |  | 
**DefaultValuePredefinedVariationName** | **NullableString** | Optional name of the predefined variation for the default value. | 
**TargetingRules** | [**[]AuditLogSettingValueV2EvaluationFormulaTargetingRuleModel**](AuditLogSettingValueV2EvaluationFormulaTargetingRuleModel.md) | List of targeting rules that determine when different values should be returned. | 
**PercentageEvaluationAttribute** | **NullableString** | Optional name of the user attribute for percentage-based user bucketing (for A/B testing). | 

## Methods

### NewAuditLogSettingValueV2EvaluationFormula

`func NewAuditLogSettingValueV2EvaluationFormula(defaultValue ValueModel, defaultValuePredefinedVariationName NullableString, targetingRules []AuditLogSettingValueV2EvaluationFormulaTargetingRuleModel, percentageEvaluationAttribute NullableString, ) *AuditLogSettingValueV2EvaluationFormula`

NewAuditLogSettingValueV2EvaluationFormula instantiates a new AuditLogSettingValueV2EvaluationFormula object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditLogSettingValueV2EvaluationFormulaWithDefaults

`func NewAuditLogSettingValueV2EvaluationFormulaWithDefaults() *AuditLogSettingValueV2EvaluationFormula`

NewAuditLogSettingValueV2EvaluationFormulaWithDefaults instantiates a new AuditLogSettingValueV2EvaluationFormula object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultValue

`func (o *AuditLogSettingValueV2EvaluationFormula) GetDefaultValue() ValueModel`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *AuditLogSettingValueV2EvaluationFormula) GetDefaultValueOk() (*ValueModel, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *AuditLogSettingValueV2EvaluationFormula) SetDefaultValue(v ValueModel)`

SetDefaultValue sets DefaultValue field to given value.


### GetDefaultValuePredefinedVariationName

`func (o *AuditLogSettingValueV2EvaluationFormula) GetDefaultValuePredefinedVariationName() string`

GetDefaultValuePredefinedVariationName returns the DefaultValuePredefinedVariationName field if non-nil, zero value otherwise.

### GetDefaultValuePredefinedVariationNameOk

`func (o *AuditLogSettingValueV2EvaluationFormula) GetDefaultValuePredefinedVariationNameOk() (*string, bool)`

GetDefaultValuePredefinedVariationNameOk returns a tuple with the DefaultValuePredefinedVariationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValuePredefinedVariationName

`func (o *AuditLogSettingValueV2EvaluationFormula) SetDefaultValuePredefinedVariationName(v string)`

SetDefaultValuePredefinedVariationName sets DefaultValuePredefinedVariationName field to given value.


### SetDefaultValuePredefinedVariationNameNil

`func (o *AuditLogSettingValueV2EvaluationFormula) SetDefaultValuePredefinedVariationNameNil(b bool)`

 SetDefaultValuePredefinedVariationNameNil sets the value for DefaultValuePredefinedVariationName to be an explicit nil

### UnsetDefaultValuePredefinedVariationName
`func (o *AuditLogSettingValueV2EvaluationFormula) UnsetDefaultValuePredefinedVariationName()`

UnsetDefaultValuePredefinedVariationName ensures that no value is present for DefaultValuePredefinedVariationName, not even an explicit nil
### GetTargetingRules

`func (o *AuditLogSettingValueV2EvaluationFormula) GetTargetingRules() []AuditLogSettingValueV2EvaluationFormulaTargetingRuleModel`

GetTargetingRules returns the TargetingRules field if non-nil, zero value otherwise.

### GetTargetingRulesOk

`func (o *AuditLogSettingValueV2EvaluationFormula) GetTargetingRulesOk() (*[]AuditLogSettingValueV2EvaluationFormulaTargetingRuleModel, bool)`

GetTargetingRulesOk returns a tuple with the TargetingRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetingRules

`func (o *AuditLogSettingValueV2EvaluationFormula) SetTargetingRules(v []AuditLogSettingValueV2EvaluationFormulaTargetingRuleModel)`

SetTargetingRules sets TargetingRules field to given value.


### GetPercentageEvaluationAttribute

`func (o *AuditLogSettingValueV2EvaluationFormula) GetPercentageEvaluationAttribute() string`

GetPercentageEvaluationAttribute returns the PercentageEvaluationAttribute field if non-nil, zero value otherwise.

### GetPercentageEvaluationAttributeOk

`func (o *AuditLogSettingValueV2EvaluationFormula) GetPercentageEvaluationAttributeOk() (*string, bool)`

GetPercentageEvaluationAttributeOk returns a tuple with the PercentageEvaluationAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageEvaluationAttribute

`func (o *AuditLogSettingValueV2EvaluationFormula) SetPercentageEvaluationAttribute(v string)`

SetPercentageEvaluationAttribute sets PercentageEvaluationAttribute field to given value.


### SetPercentageEvaluationAttributeNil

`func (o *AuditLogSettingValueV2EvaluationFormula) SetPercentageEvaluationAttributeNil(b bool)`

 SetPercentageEvaluationAttributeNil sets the value for PercentageEvaluationAttribute to be an explicit nil

### UnsetPercentageEvaluationAttribute
`func (o *AuditLogSettingValueV2EvaluationFormula) UnsetPercentageEvaluationAttribute()`

UnsetPercentageEvaluationAttribute ensures that no value is present for PercentageEvaluationAttribute, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


