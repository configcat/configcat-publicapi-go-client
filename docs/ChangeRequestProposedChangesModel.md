# ChangeRequestProposedChangesModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeatureFlagLimitations** | [**FeatureFlagLimitations**](FeatureFlagLimitations.md) |  | 
**ProposedChanges** | [**[]ConfigSettingFormulaModel**](ConfigSettingFormulaModel.md) | Collection of evaluation formulas for each Setting included in the Change Request. | 

## Methods

### NewChangeRequestProposedChangesModel

`func NewChangeRequestProposedChangesModel(featureFlagLimitations FeatureFlagLimitations, proposedChanges []ConfigSettingFormulaModel, ) *ChangeRequestProposedChangesModel`

NewChangeRequestProposedChangesModel instantiates a new ChangeRequestProposedChangesModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangeRequestProposedChangesModelWithDefaults

`func NewChangeRequestProposedChangesModelWithDefaults() *ChangeRequestProposedChangesModel`

NewChangeRequestProposedChangesModelWithDefaults instantiates a new ChangeRequestProposedChangesModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeatureFlagLimitations

`func (o *ChangeRequestProposedChangesModel) GetFeatureFlagLimitations() FeatureFlagLimitations`

GetFeatureFlagLimitations returns the FeatureFlagLimitations field if non-nil, zero value otherwise.

### GetFeatureFlagLimitationsOk

`func (o *ChangeRequestProposedChangesModel) GetFeatureFlagLimitationsOk() (*FeatureFlagLimitations, bool)`

GetFeatureFlagLimitationsOk returns a tuple with the FeatureFlagLimitations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureFlagLimitations

`func (o *ChangeRequestProposedChangesModel) SetFeatureFlagLimitations(v FeatureFlagLimitations)`

SetFeatureFlagLimitations sets FeatureFlagLimitations field to given value.


### GetProposedChanges

`func (o *ChangeRequestProposedChangesModel) GetProposedChanges() []ConfigSettingFormulaModel`

GetProposedChanges returns the ProposedChanges field if non-nil, zero value otherwise.

### GetProposedChangesOk

`func (o *ChangeRequestProposedChangesModel) GetProposedChangesOk() (*[]ConfigSettingFormulaModel, bool)`

GetProposedChangesOk returns a tuple with the ProposedChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProposedChanges

`func (o *ChangeRequestProposedChangesModel) SetProposedChanges(v []ConfigSettingFormulaModel)`

SetProposedChanges sets ProposedChanges field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


