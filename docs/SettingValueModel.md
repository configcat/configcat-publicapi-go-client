# SettingValueModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RolloutRules** | Pointer to [**[]RolloutRuleModel**](RolloutRuleModel.md) | The targeting rule collection. | [optional] 
**RolloutPercentageItems** | Pointer to [**[]RolloutPercentageItemModel**](RolloutPercentageItemModel.md) | The percentage rule collection. | [optional] 
**Value** | Pointer to **interface{}** | The value to serve. It must respect the setting type. | [optional] 
**Setting** | Pointer to [**SettingDataModel**](SettingDataModel.md) |  | [optional] 
**UpdatedAt** | Pointer to **NullableTime** | The last updated date and time when the Feature Flag or Setting. | [optional] 
**LastUpdaterUserEmail** | Pointer to **NullableString** | The email of the user who last updated the Feature Flag or Setting. | [optional] 
**LastUpdaterUserFullName** | Pointer to **NullableString** | The name of the user who last updated the Feature Flag or Setting. | [optional] 
**IntegrationLinks** | Pointer to [**[]IntegrationLinkModel**](IntegrationLinkModel.md) | The integration links attached to the Feature Flag or Setting. | [optional] 
**SettingTags** | Pointer to [**[]SettingTagModel**](SettingTagModel.md) | The tags attached to the Feature Flag or Setting. | [optional] 
**Config** | Pointer to [**ConfigModel**](ConfigModel.md) |  | [optional] 
**Environment** | Pointer to [**EnvironmentModel**](EnvironmentModel.md) |  | [optional] 
**FeatureFlagLimitations** | Pointer to [**FeatureFlagLimitations**](FeatureFlagLimitations.md) |  | [optional] 
**ReadOnly** | Pointer to **bool** |  | [optional] 
**Links** | Pointer to [**ConfigSettingFormulasModelLinks**](ConfigSettingFormulasModelLinks.md) |  | [optional] 

## Methods

### NewSettingValueModel

`func NewSettingValueModel() *SettingValueModel`

NewSettingValueModel instantiates a new SettingValueModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettingValueModelWithDefaults

`func NewSettingValueModelWithDefaults() *SettingValueModel`

NewSettingValueModelWithDefaults instantiates a new SettingValueModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRolloutRules

`func (o *SettingValueModel) GetRolloutRules() []RolloutRuleModel`

GetRolloutRules returns the RolloutRules field if non-nil, zero value otherwise.

### GetRolloutRulesOk

`func (o *SettingValueModel) GetRolloutRulesOk() (*[]RolloutRuleModel, bool)`

GetRolloutRulesOk returns a tuple with the RolloutRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolloutRules

`func (o *SettingValueModel) SetRolloutRules(v []RolloutRuleModel)`

SetRolloutRules sets RolloutRules field to given value.

### HasRolloutRules

`func (o *SettingValueModel) HasRolloutRules() bool`

HasRolloutRules returns a boolean if a field has been set.

### SetRolloutRulesNil

`func (o *SettingValueModel) SetRolloutRulesNil(b bool)`

 SetRolloutRulesNil sets the value for RolloutRules to be an explicit nil

### UnsetRolloutRules
`func (o *SettingValueModel) UnsetRolloutRules()`

UnsetRolloutRules ensures that no value is present for RolloutRules, not even an explicit nil
### GetRolloutPercentageItems

`func (o *SettingValueModel) GetRolloutPercentageItems() []RolloutPercentageItemModel`

GetRolloutPercentageItems returns the RolloutPercentageItems field if non-nil, zero value otherwise.

### GetRolloutPercentageItemsOk

`func (o *SettingValueModel) GetRolloutPercentageItemsOk() (*[]RolloutPercentageItemModel, bool)`

GetRolloutPercentageItemsOk returns a tuple with the RolloutPercentageItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolloutPercentageItems

`func (o *SettingValueModel) SetRolloutPercentageItems(v []RolloutPercentageItemModel)`

SetRolloutPercentageItems sets RolloutPercentageItems field to given value.

### HasRolloutPercentageItems

`func (o *SettingValueModel) HasRolloutPercentageItems() bool`

HasRolloutPercentageItems returns a boolean if a field has been set.

### SetRolloutPercentageItemsNil

`func (o *SettingValueModel) SetRolloutPercentageItemsNil(b bool)`

 SetRolloutPercentageItemsNil sets the value for RolloutPercentageItems to be an explicit nil

### UnsetRolloutPercentageItems
`func (o *SettingValueModel) UnsetRolloutPercentageItems()`

UnsetRolloutPercentageItems ensures that no value is present for RolloutPercentageItems, not even an explicit nil
### GetValue

`func (o *SettingValueModel) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SettingValueModel) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SettingValueModel) SetValue(v interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *SettingValueModel) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *SettingValueModel) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *SettingValueModel) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetSetting

`func (o *SettingValueModel) GetSetting() SettingDataModel`

GetSetting returns the Setting field if non-nil, zero value otherwise.

### GetSettingOk

`func (o *SettingValueModel) GetSettingOk() (*SettingDataModel, bool)`

GetSettingOk returns a tuple with the Setting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetting

`func (o *SettingValueModel) SetSetting(v SettingDataModel)`

SetSetting sets Setting field to given value.

### HasSetting

`func (o *SettingValueModel) HasSetting() bool`

HasSetting returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *SettingValueModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SettingValueModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SettingValueModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SettingValueModel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *SettingValueModel) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *SettingValueModel) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetLastUpdaterUserEmail

`func (o *SettingValueModel) GetLastUpdaterUserEmail() string`

GetLastUpdaterUserEmail returns the LastUpdaterUserEmail field if non-nil, zero value otherwise.

### GetLastUpdaterUserEmailOk

`func (o *SettingValueModel) GetLastUpdaterUserEmailOk() (*string, bool)`

GetLastUpdaterUserEmailOk returns a tuple with the LastUpdaterUserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdaterUserEmail

`func (o *SettingValueModel) SetLastUpdaterUserEmail(v string)`

SetLastUpdaterUserEmail sets LastUpdaterUserEmail field to given value.

### HasLastUpdaterUserEmail

`func (o *SettingValueModel) HasLastUpdaterUserEmail() bool`

HasLastUpdaterUserEmail returns a boolean if a field has been set.

### SetLastUpdaterUserEmailNil

`func (o *SettingValueModel) SetLastUpdaterUserEmailNil(b bool)`

 SetLastUpdaterUserEmailNil sets the value for LastUpdaterUserEmail to be an explicit nil

### UnsetLastUpdaterUserEmail
`func (o *SettingValueModel) UnsetLastUpdaterUserEmail()`

UnsetLastUpdaterUserEmail ensures that no value is present for LastUpdaterUserEmail, not even an explicit nil
### GetLastUpdaterUserFullName

`func (o *SettingValueModel) GetLastUpdaterUserFullName() string`

GetLastUpdaterUserFullName returns the LastUpdaterUserFullName field if non-nil, zero value otherwise.

### GetLastUpdaterUserFullNameOk

`func (o *SettingValueModel) GetLastUpdaterUserFullNameOk() (*string, bool)`

GetLastUpdaterUserFullNameOk returns a tuple with the LastUpdaterUserFullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdaterUserFullName

`func (o *SettingValueModel) SetLastUpdaterUserFullName(v string)`

SetLastUpdaterUserFullName sets LastUpdaterUserFullName field to given value.

### HasLastUpdaterUserFullName

`func (o *SettingValueModel) HasLastUpdaterUserFullName() bool`

HasLastUpdaterUserFullName returns a boolean if a field has been set.

### SetLastUpdaterUserFullNameNil

`func (o *SettingValueModel) SetLastUpdaterUserFullNameNil(b bool)`

 SetLastUpdaterUserFullNameNil sets the value for LastUpdaterUserFullName to be an explicit nil

### UnsetLastUpdaterUserFullName
`func (o *SettingValueModel) UnsetLastUpdaterUserFullName()`

UnsetLastUpdaterUserFullName ensures that no value is present for LastUpdaterUserFullName, not even an explicit nil
### GetIntegrationLinks

`func (o *SettingValueModel) GetIntegrationLinks() []IntegrationLinkModel`

GetIntegrationLinks returns the IntegrationLinks field if non-nil, zero value otherwise.

### GetIntegrationLinksOk

`func (o *SettingValueModel) GetIntegrationLinksOk() (*[]IntegrationLinkModel, bool)`

GetIntegrationLinksOk returns a tuple with the IntegrationLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationLinks

`func (o *SettingValueModel) SetIntegrationLinks(v []IntegrationLinkModel)`

SetIntegrationLinks sets IntegrationLinks field to given value.

### HasIntegrationLinks

`func (o *SettingValueModel) HasIntegrationLinks() bool`

HasIntegrationLinks returns a boolean if a field has been set.

### SetIntegrationLinksNil

`func (o *SettingValueModel) SetIntegrationLinksNil(b bool)`

 SetIntegrationLinksNil sets the value for IntegrationLinks to be an explicit nil

### UnsetIntegrationLinks
`func (o *SettingValueModel) UnsetIntegrationLinks()`

UnsetIntegrationLinks ensures that no value is present for IntegrationLinks, not even an explicit nil
### GetSettingTags

`func (o *SettingValueModel) GetSettingTags() []SettingTagModel`

GetSettingTags returns the SettingTags field if non-nil, zero value otherwise.

### GetSettingTagsOk

`func (o *SettingValueModel) GetSettingTagsOk() (*[]SettingTagModel, bool)`

GetSettingTagsOk returns a tuple with the SettingTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingTags

`func (o *SettingValueModel) SetSettingTags(v []SettingTagModel)`

SetSettingTags sets SettingTags field to given value.

### HasSettingTags

`func (o *SettingValueModel) HasSettingTags() bool`

HasSettingTags returns a boolean if a field has been set.

### SetSettingTagsNil

`func (o *SettingValueModel) SetSettingTagsNil(b bool)`

 SetSettingTagsNil sets the value for SettingTags to be an explicit nil

### UnsetSettingTags
`func (o *SettingValueModel) UnsetSettingTags()`

UnsetSettingTags ensures that no value is present for SettingTags, not even an explicit nil
### GetConfig

`func (o *SettingValueModel) GetConfig() ConfigModel`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *SettingValueModel) GetConfigOk() (*ConfigModel, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *SettingValueModel) SetConfig(v ConfigModel)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *SettingValueModel) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetEnvironment

`func (o *SettingValueModel) GetEnvironment() EnvironmentModel`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *SettingValueModel) GetEnvironmentOk() (*EnvironmentModel, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *SettingValueModel) SetEnvironment(v EnvironmentModel)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *SettingValueModel) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetFeatureFlagLimitations

`func (o *SettingValueModel) GetFeatureFlagLimitations() FeatureFlagLimitations`

GetFeatureFlagLimitations returns the FeatureFlagLimitations field if non-nil, zero value otherwise.

### GetFeatureFlagLimitationsOk

`func (o *SettingValueModel) GetFeatureFlagLimitationsOk() (*FeatureFlagLimitations, bool)`

GetFeatureFlagLimitationsOk returns a tuple with the FeatureFlagLimitations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureFlagLimitations

`func (o *SettingValueModel) SetFeatureFlagLimitations(v FeatureFlagLimitations)`

SetFeatureFlagLimitations sets FeatureFlagLimitations field to given value.

### HasFeatureFlagLimitations

`func (o *SettingValueModel) HasFeatureFlagLimitations() bool`

HasFeatureFlagLimitations returns a boolean if a field has been set.

### GetReadOnly

`func (o *SettingValueModel) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *SettingValueModel) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *SettingValueModel) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *SettingValueModel) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetLinks

`func (o *SettingValueModel) GetLinks() ConfigSettingFormulasModelLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *SettingValueModel) GetLinksOk() (*ConfigSettingFormulasModelLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *SettingValueModel) SetLinks(v ConfigSettingFormulasModelLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *SettingValueModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


