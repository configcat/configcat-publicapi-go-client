# UpdatePreferencesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReasonRequired** | Pointer to **NullableBool** | Indicates that a mandatory note is required for saving and publishing. | [optional] 
**KeyGenerationMode** | Pointer to [**KeyGenerationMode**](KeyGenerationMode.md) |  | [optional] 
**ShowVariationId** | Pointer to **NullableBool** | Indicates whether a variation ID&#39;s must be shown on the ConfigCat Dashboard. | [optional] 
**MandatorySettingHint** | Pointer to **NullableBool** | Indicates whether Feature flags and Settings must have a hint. | [optional] 
**ReasonRequiredEnvironments** | Pointer to [**[]UpdateReasonRequiredEnvironmentModel**](UpdateReasonRequiredEnvironmentModel.md) | List of Environments where mandatory note must be set before saving and publishing. | [optional] 

## Methods

### NewUpdatePreferencesRequest

`func NewUpdatePreferencesRequest() *UpdatePreferencesRequest`

NewUpdatePreferencesRequest instantiates a new UpdatePreferencesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePreferencesRequestWithDefaults

`func NewUpdatePreferencesRequestWithDefaults() *UpdatePreferencesRequest`

NewUpdatePreferencesRequestWithDefaults instantiates a new UpdatePreferencesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReasonRequired

`func (o *UpdatePreferencesRequest) GetReasonRequired() bool`

GetReasonRequired returns the ReasonRequired field if non-nil, zero value otherwise.

### GetReasonRequiredOk

`func (o *UpdatePreferencesRequest) GetReasonRequiredOk() (*bool, bool)`

GetReasonRequiredOk returns a tuple with the ReasonRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonRequired

`func (o *UpdatePreferencesRequest) SetReasonRequired(v bool)`

SetReasonRequired sets ReasonRequired field to given value.

### HasReasonRequired

`func (o *UpdatePreferencesRequest) HasReasonRequired() bool`

HasReasonRequired returns a boolean if a field has been set.

### SetReasonRequiredNil

`func (o *UpdatePreferencesRequest) SetReasonRequiredNil(b bool)`

 SetReasonRequiredNil sets the value for ReasonRequired to be an explicit nil

### UnsetReasonRequired
`func (o *UpdatePreferencesRequest) UnsetReasonRequired()`

UnsetReasonRequired ensures that no value is present for ReasonRequired, not even an explicit nil
### GetKeyGenerationMode

`func (o *UpdatePreferencesRequest) GetKeyGenerationMode() KeyGenerationMode`

GetKeyGenerationMode returns the KeyGenerationMode field if non-nil, zero value otherwise.

### GetKeyGenerationModeOk

`func (o *UpdatePreferencesRequest) GetKeyGenerationModeOk() (*KeyGenerationMode, bool)`

GetKeyGenerationModeOk returns a tuple with the KeyGenerationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyGenerationMode

`func (o *UpdatePreferencesRequest) SetKeyGenerationMode(v KeyGenerationMode)`

SetKeyGenerationMode sets KeyGenerationMode field to given value.

### HasKeyGenerationMode

`func (o *UpdatePreferencesRequest) HasKeyGenerationMode() bool`

HasKeyGenerationMode returns a boolean if a field has been set.

### GetShowVariationId

`func (o *UpdatePreferencesRequest) GetShowVariationId() bool`

GetShowVariationId returns the ShowVariationId field if non-nil, zero value otherwise.

### GetShowVariationIdOk

`func (o *UpdatePreferencesRequest) GetShowVariationIdOk() (*bool, bool)`

GetShowVariationIdOk returns a tuple with the ShowVariationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowVariationId

`func (o *UpdatePreferencesRequest) SetShowVariationId(v bool)`

SetShowVariationId sets ShowVariationId field to given value.

### HasShowVariationId

`func (o *UpdatePreferencesRequest) HasShowVariationId() bool`

HasShowVariationId returns a boolean if a field has been set.

### SetShowVariationIdNil

`func (o *UpdatePreferencesRequest) SetShowVariationIdNil(b bool)`

 SetShowVariationIdNil sets the value for ShowVariationId to be an explicit nil

### UnsetShowVariationId
`func (o *UpdatePreferencesRequest) UnsetShowVariationId()`

UnsetShowVariationId ensures that no value is present for ShowVariationId, not even an explicit nil
### GetMandatorySettingHint

`func (o *UpdatePreferencesRequest) GetMandatorySettingHint() bool`

GetMandatorySettingHint returns the MandatorySettingHint field if non-nil, zero value otherwise.

### GetMandatorySettingHintOk

`func (o *UpdatePreferencesRequest) GetMandatorySettingHintOk() (*bool, bool)`

GetMandatorySettingHintOk returns a tuple with the MandatorySettingHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandatorySettingHint

`func (o *UpdatePreferencesRequest) SetMandatorySettingHint(v bool)`

SetMandatorySettingHint sets MandatorySettingHint field to given value.

### HasMandatorySettingHint

`func (o *UpdatePreferencesRequest) HasMandatorySettingHint() bool`

HasMandatorySettingHint returns a boolean if a field has been set.

### SetMandatorySettingHintNil

`func (o *UpdatePreferencesRequest) SetMandatorySettingHintNil(b bool)`

 SetMandatorySettingHintNil sets the value for MandatorySettingHint to be an explicit nil

### UnsetMandatorySettingHint
`func (o *UpdatePreferencesRequest) UnsetMandatorySettingHint()`

UnsetMandatorySettingHint ensures that no value is present for MandatorySettingHint, not even an explicit nil
### GetReasonRequiredEnvironments

`func (o *UpdatePreferencesRequest) GetReasonRequiredEnvironments() []UpdateReasonRequiredEnvironmentModel`

GetReasonRequiredEnvironments returns the ReasonRequiredEnvironments field if non-nil, zero value otherwise.

### GetReasonRequiredEnvironmentsOk

`func (o *UpdatePreferencesRequest) GetReasonRequiredEnvironmentsOk() (*[]UpdateReasonRequiredEnvironmentModel, bool)`

GetReasonRequiredEnvironmentsOk returns a tuple with the ReasonRequiredEnvironments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonRequiredEnvironments

`func (o *UpdatePreferencesRequest) SetReasonRequiredEnvironments(v []UpdateReasonRequiredEnvironmentModel)`

SetReasonRequiredEnvironments sets ReasonRequiredEnvironments field to given value.

### HasReasonRequiredEnvironments

`func (o *UpdatePreferencesRequest) HasReasonRequiredEnvironments() bool`

HasReasonRequiredEnvironments returns a boolean if a field has been set.

### SetReasonRequiredEnvironmentsNil

`func (o *UpdatePreferencesRequest) SetReasonRequiredEnvironmentsNil(b bool)`

 SetReasonRequiredEnvironmentsNil sets the value for ReasonRequiredEnvironments to be an explicit nil

### UnsetReasonRequiredEnvironments
`func (o *UpdatePreferencesRequest) UnsetReasonRequiredEnvironments()`

UnsetReasonRequiredEnvironments ensures that no value is present for ReasonRequiredEnvironments, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


