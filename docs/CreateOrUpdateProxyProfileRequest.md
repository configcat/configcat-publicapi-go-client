# CreateOrUpdateProxyProfileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the proxy profile. | 
**Description** | Pointer to **NullableString** | The description of the proxy profile. | [optional] 
**ConnectionPreferences** | Pointer to [**NullableCreateOrUpdateConnectionPreferences**](CreateOrUpdateConnectionPreferences.md) |  | [optional] 
**SdkKeySelectionRules** | Pointer to [**[]UpdateProxyProfileSelectionRule**](UpdateProxyProfileSelectionRule.md) | A collection of selection rules that determine the SDK keys applicable for a proxy profile. | [optional] 

## Methods

### NewCreateOrUpdateProxyProfileRequest

`func NewCreateOrUpdateProxyProfileRequest(name string, ) *CreateOrUpdateProxyProfileRequest`

NewCreateOrUpdateProxyProfileRequest instantiates a new CreateOrUpdateProxyProfileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrUpdateProxyProfileRequestWithDefaults

`func NewCreateOrUpdateProxyProfileRequestWithDefaults() *CreateOrUpdateProxyProfileRequest`

NewCreateOrUpdateProxyProfileRequestWithDefaults instantiates a new CreateOrUpdateProxyProfileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateOrUpdateProxyProfileRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateOrUpdateProxyProfileRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateOrUpdateProxyProfileRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateOrUpdateProxyProfileRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateOrUpdateProxyProfileRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateOrUpdateProxyProfileRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateOrUpdateProxyProfileRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateOrUpdateProxyProfileRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateOrUpdateProxyProfileRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetConnectionPreferences

`func (o *CreateOrUpdateProxyProfileRequest) GetConnectionPreferences() CreateOrUpdateConnectionPreferences`

GetConnectionPreferences returns the ConnectionPreferences field if non-nil, zero value otherwise.

### GetConnectionPreferencesOk

`func (o *CreateOrUpdateProxyProfileRequest) GetConnectionPreferencesOk() (*CreateOrUpdateConnectionPreferences, bool)`

GetConnectionPreferencesOk returns a tuple with the ConnectionPreferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionPreferences

`func (o *CreateOrUpdateProxyProfileRequest) SetConnectionPreferences(v CreateOrUpdateConnectionPreferences)`

SetConnectionPreferences sets ConnectionPreferences field to given value.

### HasConnectionPreferences

`func (o *CreateOrUpdateProxyProfileRequest) HasConnectionPreferences() bool`

HasConnectionPreferences returns a boolean if a field has been set.

### SetConnectionPreferencesNil

`func (o *CreateOrUpdateProxyProfileRequest) SetConnectionPreferencesNil(b bool)`

 SetConnectionPreferencesNil sets the value for ConnectionPreferences to be an explicit nil

### UnsetConnectionPreferences
`func (o *CreateOrUpdateProxyProfileRequest) UnsetConnectionPreferences()`

UnsetConnectionPreferences ensures that no value is present for ConnectionPreferences, not even an explicit nil
### GetSdkKeySelectionRules

`func (o *CreateOrUpdateProxyProfileRequest) GetSdkKeySelectionRules() []UpdateProxyProfileSelectionRule`

GetSdkKeySelectionRules returns the SdkKeySelectionRules field if non-nil, zero value otherwise.

### GetSdkKeySelectionRulesOk

`func (o *CreateOrUpdateProxyProfileRequest) GetSdkKeySelectionRulesOk() (*[]UpdateProxyProfileSelectionRule, bool)`

GetSdkKeySelectionRulesOk returns a tuple with the SdkKeySelectionRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdkKeySelectionRules

`func (o *CreateOrUpdateProxyProfileRequest) SetSdkKeySelectionRules(v []UpdateProxyProfileSelectionRule)`

SetSdkKeySelectionRules sets SdkKeySelectionRules field to given value.

### HasSdkKeySelectionRules

`func (o *CreateOrUpdateProxyProfileRequest) HasSdkKeySelectionRules() bool`

HasSdkKeySelectionRules returns a boolean if a field has been set.

### SetSdkKeySelectionRulesNil

`func (o *CreateOrUpdateProxyProfileRequest) SetSdkKeySelectionRulesNil(b bool)`

 SetSdkKeySelectionRulesNil sets the value for SdkKeySelectionRules to be an explicit nil

### UnsetSdkKeySelectionRules
`func (o *CreateOrUpdateProxyProfileRequest) UnsetSdkKeySelectionRules()`

UnsetSdkKeySelectionRules ensures that no value is present for SdkKeySelectionRules, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


