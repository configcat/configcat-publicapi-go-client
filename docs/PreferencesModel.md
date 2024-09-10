# PreferencesModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReasonRequired** | Pointer to **bool** | Indicates that a mandatory note required for saving and publishing. | [optional] 
**KeyGenerationMode** | Pointer to **string** | Determines the Feature Flag key generation mode. | [optional] 
**ShowVariationId** | Pointer to **bool** | Indicates whether a variation ID&#39;s must be shown on the ConfigCat Dashboard. | [optional] 
**ReasonRequiredEnvironments** | Pointer to [**[]ReasonRequiredEnvironmentModel**](ReasonRequiredEnvironmentModel.md) | List of Environments where mandatory note must be set before saving and publishing. | [optional] 
**MandatorySettingHint** | Pointer to **bool** | Indicates whether Feature flags and Settings must have a hint. | [optional] 

## Methods

### NewPreferencesModel

`func NewPreferencesModel() *PreferencesModel`

NewPreferencesModel instantiates a new PreferencesModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPreferencesModelWithDefaults

`func NewPreferencesModelWithDefaults() *PreferencesModel`

NewPreferencesModelWithDefaults instantiates a new PreferencesModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReasonRequired

`func (o *PreferencesModel) GetReasonRequired() bool`

GetReasonRequired returns the ReasonRequired field if non-nil, zero value otherwise.

### GetReasonRequiredOk

`func (o *PreferencesModel) GetReasonRequiredOk() (*bool, bool)`

GetReasonRequiredOk returns a tuple with the ReasonRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonRequired

`func (o *PreferencesModel) SetReasonRequired(v bool)`

SetReasonRequired sets ReasonRequired field to given value.

### HasReasonRequired

`func (o *PreferencesModel) HasReasonRequired() bool`

HasReasonRequired returns a boolean if a field has been set.

### GetKeyGenerationMode

`func (o *PreferencesModel) GetKeyGenerationMode() string`

GetKeyGenerationMode returns the KeyGenerationMode field if non-nil, zero value otherwise.

### GetKeyGenerationModeOk

`func (o *PreferencesModel) GetKeyGenerationModeOk() (*string, bool)`

GetKeyGenerationModeOk returns a tuple with the KeyGenerationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyGenerationMode

`func (o *PreferencesModel) SetKeyGenerationMode(v string)`

SetKeyGenerationMode sets KeyGenerationMode field to given value.

### HasKeyGenerationMode

`func (o *PreferencesModel) HasKeyGenerationMode() bool`

HasKeyGenerationMode returns a boolean if a field has been set.

### GetShowVariationId

`func (o *PreferencesModel) GetShowVariationId() bool`

GetShowVariationId returns the ShowVariationId field if non-nil, zero value otherwise.

### GetShowVariationIdOk

`func (o *PreferencesModel) GetShowVariationIdOk() (*bool, bool)`

GetShowVariationIdOk returns a tuple with the ShowVariationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowVariationId

`func (o *PreferencesModel) SetShowVariationId(v bool)`

SetShowVariationId sets ShowVariationId field to given value.

### HasShowVariationId

`func (o *PreferencesModel) HasShowVariationId() bool`

HasShowVariationId returns a boolean if a field has been set.

### GetReasonRequiredEnvironments

`func (o *PreferencesModel) GetReasonRequiredEnvironments() []ReasonRequiredEnvironmentModel`

GetReasonRequiredEnvironments returns the ReasonRequiredEnvironments field if non-nil, zero value otherwise.

### GetReasonRequiredEnvironmentsOk

`func (o *PreferencesModel) GetReasonRequiredEnvironmentsOk() (*[]ReasonRequiredEnvironmentModel, bool)`

GetReasonRequiredEnvironmentsOk returns a tuple with the ReasonRequiredEnvironments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonRequiredEnvironments

`func (o *PreferencesModel) SetReasonRequiredEnvironments(v []ReasonRequiredEnvironmentModel)`

SetReasonRequiredEnvironments sets ReasonRequiredEnvironments field to given value.

### HasReasonRequiredEnvironments

`func (o *PreferencesModel) HasReasonRequiredEnvironments() bool`

HasReasonRequiredEnvironments returns a boolean if a field has been set.

### SetReasonRequiredEnvironmentsNil

`func (o *PreferencesModel) SetReasonRequiredEnvironmentsNil(b bool)`

 SetReasonRequiredEnvironmentsNil sets the value for ReasonRequiredEnvironments to be an explicit nil

### UnsetReasonRequiredEnvironments
`func (o *PreferencesModel) UnsetReasonRequiredEnvironments()`

UnsetReasonRequiredEnvironments ensures that no value is present for ReasonRequiredEnvironments, not even an explicit nil
### GetMandatorySettingHint

`func (o *PreferencesModel) GetMandatorySettingHint() bool`

GetMandatorySettingHint returns the MandatorySettingHint field if non-nil, zero value otherwise.

### GetMandatorySettingHintOk

`func (o *PreferencesModel) GetMandatorySettingHintOk() (*bool, bool)`

GetMandatorySettingHintOk returns a tuple with the MandatorySettingHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandatorySettingHint

`func (o *PreferencesModel) SetMandatorySettingHint(v bool)`

SetMandatorySettingHint sets MandatorySettingHint field to given value.

### HasMandatorySettingHint

`func (o *PreferencesModel) HasMandatorySettingHint() bool`

HasMandatorySettingHint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


