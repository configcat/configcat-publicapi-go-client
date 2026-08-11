# ChangeRequestProposedChangeModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SettingId** | **int32** | Identifier of the Setting to change. | 
**SettingKey** | **string** | Key identifier of the Setting. | 
**SettingName** | **string** | Display name of the Setting. | 
**SettingHint** | **NullableString** | Optional hint or description for the Setting. | 
**SettingType** | [**SettingType**](SettingType.md) |  | 
**HasConflict** | **bool** | Indicates whether the proposed changes to the Setting are in conflict with concurrently published changes. | 
**OriginalEvaluationFormula** | [**AuditLogSettingValueV2EvaluationFormula**](AuditLogSettingValueV2EvaluationFormula.md) |  | 
**ProposedEvaluationFormula** | [**AuditLogSettingValueV2EvaluationFormula**](AuditLogSettingValueV2EvaluationFormula.md) |  | 

## Methods

### NewChangeRequestProposedChangeModel

`func NewChangeRequestProposedChangeModel(settingId int32, settingKey string, settingName string, settingHint NullableString, settingType SettingType, hasConflict bool, originalEvaluationFormula AuditLogSettingValueV2EvaluationFormula, proposedEvaluationFormula AuditLogSettingValueV2EvaluationFormula, ) *ChangeRequestProposedChangeModel`

NewChangeRequestProposedChangeModel instantiates a new ChangeRequestProposedChangeModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangeRequestProposedChangeModelWithDefaults

`func NewChangeRequestProposedChangeModelWithDefaults() *ChangeRequestProposedChangeModel`

NewChangeRequestProposedChangeModelWithDefaults instantiates a new ChangeRequestProposedChangeModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSettingId

`func (o *ChangeRequestProposedChangeModel) GetSettingId() int32`

GetSettingId returns the SettingId field if non-nil, zero value otherwise.

### GetSettingIdOk

`func (o *ChangeRequestProposedChangeModel) GetSettingIdOk() (*int32, bool)`

GetSettingIdOk returns a tuple with the SettingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingId

`func (o *ChangeRequestProposedChangeModel) SetSettingId(v int32)`

SetSettingId sets SettingId field to given value.


### GetSettingKey

`func (o *ChangeRequestProposedChangeModel) GetSettingKey() string`

GetSettingKey returns the SettingKey field if non-nil, zero value otherwise.

### GetSettingKeyOk

`func (o *ChangeRequestProposedChangeModel) GetSettingKeyOk() (*string, bool)`

GetSettingKeyOk returns a tuple with the SettingKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingKey

`func (o *ChangeRequestProposedChangeModel) SetSettingKey(v string)`

SetSettingKey sets SettingKey field to given value.


### GetSettingName

`func (o *ChangeRequestProposedChangeModel) GetSettingName() string`

GetSettingName returns the SettingName field if non-nil, zero value otherwise.

### GetSettingNameOk

`func (o *ChangeRequestProposedChangeModel) GetSettingNameOk() (*string, bool)`

GetSettingNameOk returns a tuple with the SettingName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingName

`func (o *ChangeRequestProposedChangeModel) SetSettingName(v string)`

SetSettingName sets SettingName field to given value.


### GetSettingHint

`func (o *ChangeRequestProposedChangeModel) GetSettingHint() string`

GetSettingHint returns the SettingHint field if non-nil, zero value otherwise.

### GetSettingHintOk

`func (o *ChangeRequestProposedChangeModel) GetSettingHintOk() (*string, bool)`

GetSettingHintOk returns a tuple with the SettingHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingHint

`func (o *ChangeRequestProposedChangeModel) SetSettingHint(v string)`

SetSettingHint sets SettingHint field to given value.


### SetSettingHintNil

`func (o *ChangeRequestProposedChangeModel) SetSettingHintNil(b bool)`

 SetSettingHintNil sets the value for SettingHint to be an explicit nil

### UnsetSettingHint
`func (o *ChangeRequestProposedChangeModel) UnsetSettingHint()`

UnsetSettingHint ensures that no value is present for SettingHint, not even an explicit nil
### GetSettingType

`func (o *ChangeRequestProposedChangeModel) GetSettingType() SettingType`

GetSettingType returns the SettingType field if non-nil, zero value otherwise.

### GetSettingTypeOk

`func (o *ChangeRequestProposedChangeModel) GetSettingTypeOk() (*SettingType, bool)`

GetSettingTypeOk returns a tuple with the SettingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingType

`func (o *ChangeRequestProposedChangeModel) SetSettingType(v SettingType)`

SetSettingType sets SettingType field to given value.


### GetHasConflict

`func (o *ChangeRequestProposedChangeModel) GetHasConflict() bool`

GetHasConflict returns the HasConflict field if non-nil, zero value otherwise.

### GetHasConflictOk

`func (o *ChangeRequestProposedChangeModel) GetHasConflictOk() (*bool, bool)`

GetHasConflictOk returns a tuple with the HasConflict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasConflict

`func (o *ChangeRequestProposedChangeModel) SetHasConflict(v bool)`

SetHasConflict sets HasConflict field to given value.


### GetOriginalEvaluationFormula

`func (o *ChangeRequestProposedChangeModel) GetOriginalEvaluationFormula() AuditLogSettingValueV2EvaluationFormula`

GetOriginalEvaluationFormula returns the OriginalEvaluationFormula field if non-nil, zero value otherwise.

### GetOriginalEvaluationFormulaOk

`func (o *ChangeRequestProposedChangeModel) GetOriginalEvaluationFormulaOk() (*AuditLogSettingValueV2EvaluationFormula, bool)`

GetOriginalEvaluationFormulaOk returns a tuple with the OriginalEvaluationFormula field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalEvaluationFormula

`func (o *ChangeRequestProposedChangeModel) SetOriginalEvaluationFormula(v AuditLogSettingValueV2EvaluationFormula)`

SetOriginalEvaluationFormula sets OriginalEvaluationFormula field to given value.


### GetProposedEvaluationFormula

`func (o *ChangeRequestProposedChangeModel) GetProposedEvaluationFormula() AuditLogSettingValueV2EvaluationFormula`

GetProposedEvaluationFormula returns the ProposedEvaluationFormula field if non-nil, zero value otherwise.

### GetProposedEvaluationFormulaOk

`func (o *ChangeRequestProposedChangeModel) GetProposedEvaluationFormulaOk() (*AuditLogSettingValueV2EvaluationFormula, bool)`

GetProposedEvaluationFormulaOk returns a tuple with the ProposedEvaluationFormula field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProposedEvaluationFormula

`func (o *ChangeRequestProposedChangeModel) SetProposedEvaluationFormula(v AuditLogSettingValueV2EvaluationFormula)`

SetProposedEvaluationFormula sets ProposedEvaluationFormula field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


