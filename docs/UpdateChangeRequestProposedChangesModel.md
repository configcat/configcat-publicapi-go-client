# UpdateChangeRequestProposedChangesModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProposedChanges** | Pointer to [**[]UpdateEvaluationFormulaWithLatestVersionModel**](UpdateEvaluationFormulaWithLatestVersionModel.md) | The setting values to update on the Change Request. | [optional] 
**Forced** | Pointer to **bool** | When true, skips conflict (LatestVersionId) checking. | [optional] 

## Methods

### NewUpdateChangeRequestProposedChangesModel

`func NewUpdateChangeRequestProposedChangesModel() *UpdateChangeRequestProposedChangesModel`

NewUpdateChangeRequestProposedChangesModel instantiates a new UpdateChangeRequestProposedChangesModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateChangeRequestProposedChangesModelWithDefaults

`func NewUpdateChangeRequestProposedChangesModelWithDefaults() *UpdateChangeRequestProposedChangesModel`

NewUpdateChangeRequestProposedChangesModelWithDefaults instantiates a new UpdateChangeRequestProposedChangesModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProposedChanges

`func (o *UpdateChangeRequestProposedChangesModel) GetProposedChanges() []UpdateEvaluationFormulaWithLatestVersionModel`

GetProposedChanges returns the ProposedChanges field if non-nil, zero value otherwise.

### GetProposedChangesOk

`func (o *UpdateChangeRequestProposedChangesModel) GetProposedChangesOk() (*[]UpdateEvaluationFormulaWithLatestVersionModel, bool)`

GetProposedChangesOk returns a tuple with the ProposedChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProposedChanges

`func (o *UpdateChangeRequestProposedChangesModel) SetProposedChanges(v []UpdateEvaluationFormulaWithLatestVersionModel)`

SetProposedChanges sets ProposedChanges field to given value.

### HasProposedChanges

`func (o *UpdateChangeRequestProposedChangesModel) HasProposedChanges() bool`

HasProposedChanges returns a boolean if a field has been set.

### GetForced

`func (o *UpdateChangeRequestProposedChangesModel) GetForced() bool`

GetForced returns the Forced field if non-nil, zero value otherwise.

### GetForcedOk

`func (o *UpdateChangeRequestProposedChangesModel) GetForcedOk() (*bool, bool)`

GetForcedOk returns a tuple with the Forced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForced

`func (o *UpdateChangeRequestProposedChangesModel) SetForced(v bool)`

SetForced sets Forced field to given value.

### HasForced

`func (o *UpdateChangeRequestProposedChangesModel) HasForced() bool`

HasForced returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


