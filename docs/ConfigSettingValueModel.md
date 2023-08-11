# ConfigSettingValueModel

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

## Methods

### NewConfigSettingValueModel

`func NewConfigSettingValueModel() *ConfigSettingValueModel`

NewConfigSettingValueModel instantiates a new ConfigSettingValueModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigSettingValueModelWithDefaults

`func NewConfigSettingValueModelWithDefaults() *ConfigSettingValueModel`

NewConfigSettingValueModelWithDefaults instantiates a new ConfigSettingValueModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRolloutRules

`func (o *ConfigSettingValueModel) GetRolloutRules() []RolloutRuleModel`

GetRolloutRules returns the RolloutRules field if non-nil, zero value otherwise.

### GetRolloutRulesOk

`func (o *ConfigSettingValueModel) GetRolloutRulesOk() (*[]RolloutRuleModel, bool)`

GetRolloutRulesOk returns a tuple with the RolloutRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolloutRules

`func (o *ConfigSettingValueModel) SetRolloutRules(v []RolloutRuleModel)`

SetRolloutRules sets RolloutRules field to given value.

### HasRolloutRules

`func (o *ConfigSettingValueModel) HasRolloutRules() bool`

HasRolloutRules returns a boolean if a field has been set.

### SetRolloutRulesNil

`func (o *ConfigSettingValueModel) SetRolloutRulesNil(b bool)`

 SetRolloutRulesNil sets the value for RolloutRules to be an explicit nil

### UnsetRolloutRules
`func (o *ConfigSettingValueModel) UnsetRolloutRules()`

UnsetRolloutRules ensures that no value is present for RolloutRules, not even an explicit nil
### GetRolloutPercentageItems

`func (o *ConfigSettingValueModel) GetRolloutPercentageItems() []RolloutPercentageItemModel`

GetRolloutPercentageItems returns the RolloutPercentageItems field if non-nil, zero value otherwise.

### GetRolloutPercentageItemsOk

`func (o *ConfigSettingValueModel) GetRolloutPercentageItemsOk() (*[]RolloutPercentageItemModel, bool)`

GetRolloutPercentageItemsOk returns a tuple with the RolloutPercentageItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolloutPercentageItems

`func (o *ConfigSettingValueModel) SetRolloutPercentageItems(v []RolloutPercentageItemModel)`

SetRolloutPercentageItems sets RolloutPercentageItems field to given value.

### HasRolloutPercentageItems

`func (o *ConfigSettingValueModel) HasRolloutPercentageItems() bool`

HasRolloutPercentageItems returns a boolean if a field has been set.

### SetRolloutPercentageItemsNil

`func (o *ConfigSettingValueModel) SetRolloutPercentageItemsNil(b bool)`

 SetRolloutPercentageItemsNil sets the value for RolloutPercentageItems to be an explicit nil

### UnsetRolloutPercentageItems
`func (o *ConfigSettingValueModel) UnsetRolloutPercentageItems()`

UnsetRolloutPercentageItems ensures that no value is present for RolloutPercentageItems, not even an explicit nil
### GetValue

`func (o *ConfigSettingValueModel) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ConfigSettingValueModel) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ConfigSettingValueModel) SetValue(v interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *ConfigSettingValueModel) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *ConfigSettingValueModel) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *ConfigSettingValueModel) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetSetting

`func (o *ConfigSettingValueModel) GetSetting() SettingDataModel`

GetSetting returns the Setting field if non-nil, zero value otherwise.

### GetSettingOk

`func (o *ConfigSettingValueModel) GetSettingOk() (*SettingDataModel, bool)`

GetSettingOk returns a tuple with the Setting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetting

`func (o *ConfigSettingValueModel) SetSetting(v SettingDataModel)`

SetSetting sets Setting field to given value.

### HasSetting

`func (o *ConfigSettingValueModel) HasSetting() bool`

HasSetting returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ConfigSettingValueModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ConfigSettingValueModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ConfigSettingValueModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ConfigSettingValueModel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *ConfigSettingValueModel) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *ConfigSettingValueModel) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetLastUpdaterUserEmail

`func (o *ConfigSettingValueModel) GetLastUpdaterUserEmail() string`

GetLastUpdaterUserEmail returns the LastUpdaterUserEmail field if non-nil, zero value otherwise.

### GetLastUpdaterUserEmailOk

`func (o *ConfigSettingValueModel) GetLastUpdaterUserEmailOk() (*string, bool)`

GetLastUpdaterUserEmailOk returns a tuple with the LastUpdaterUserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdaterUserEmail

`func (o *ConfigSettingValueModel) SetLastUpdaterUserEmail(v string)`

SetLastUpdaterUserEmail sets LastUpdaterUserEmail field to given value.

### HasLastUpdaterUserEmail

`func (o *ConfigSettingValueModel) HasLastUpdaterUserEmail() bool`

HasLastUpdaterUserEmail returns a boolean if a field has been set.

### SetLastUpdaterUserEmailNil

`func (o *ConfigSettingValueModel) SetLastUpdaterUserEmailNil(b bool)`

 SetLastUpdaterUserEmailNil sets the value for LastUpdaterUserEmail to be an explicit nil

### UnsetLastUpdaterUserEmail
`func (o *ConfigSettingValueModel) UnsetLastUpdaterUserEmail()`

UnsetLastUpdaterUserEmail ensures that no value is present for LastUpdaterUserEmail, not even an explicit nil
### GetLastUpdaterUserFullName

`func (o *ConfigSettingValueModel) GetLastUpdaterUserFullName() string`

GetLastUpdaterUserFullName returns the LastUpdaterUserFullName field if non-nil, zero value otherwise.

### GetLastUpdaterUserFullNameOk

`func (o *ConfigSettingValueModel) GetLastUpdaterUserFullNameOk() (*string, bool)`

GetLastUpdaterUserFullNameOk returns a tuple with the LastUpdaterUserFullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdaterUserFullName

`func (o *ConfigSettingValueModel) SetLastUpdaterUserFullName(v string)`

SetLastUpdaterUserFullName sets LastUpdaterUserFullName field to given value.

### HasLastUpdaterUserFullName

`func (o *ConfigSettingValueModel) HasLastUpdaterUserFullName() bool`

HasLastUpdaterUserFullName returns a boolean if a field has been set.

### SetLastUpdaterUserFullNameNil

`func (o *ConfigSettingValueModel) SetLastUpdaterUserFullNameNil(b bool)`

 SetLastUpdaterUserFullNameNil sets the value for LastUpdaterUserFullName to be an explicit nil

### UnsetLastUpdaterUserFullName
`func (o *ConfigSettingValueModel) UnsetLastUpdaterUserFullName()`

UnsetLastUpdaterUserFullName ensures that no value is present for LastUpdaterUserFullName, not even an explicit nil
### GetIntegrationLinks

`func (o *ConfigSettingValueModel) GetIntegrationLinks() []IntegrationLinkModel`

GetIntegrationLinks returns the IntegrationLinks field if non-nil, zero value otherwise.

### GetIntegrationLinksOk

`func (o *ConfigSettingValueModel) GetIntegrationLinksOk() (*[]IntegrationLinkModel, bool)`

GetIntegrationLinksOk returns a tuple with the IntegrationLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationLinks

`func (o *ConfigSettingValueModel) SetIntegrationLinks(v []IntegrationLinkModel)`

SetIntegrationLinks sets IntegrationLinks field to given value.

### HasIntegrationLinks

`func (o *ConfigSettingValueModel) HasIntegrationLinks() bool`

HasIntegrationLinks returns a boolean if a field has been set.

### SetIntegrationLinksNil

`func (o *ConfigSettingValueModel) SetIntegrationLinksNil(b bool)`

 SetIntegrationLinksNil sets the value for IntegrationLinks to be an explicit nil

### UnsetIntegrationLinks
`func (o *ConfigSettingValueModel) UnsetIntegrationLinks()`

UnsetIntegrationLinks ensures that no value is present for IntegrationLinks, not even an explicit nil
### GetSettingTags

`func (o *ConfigSettingValueModel) GetSettingTags() []SettingTagModel`

GetSettingTags returns the SettingTags field if non-nil, zero value otherwise.

### GetSettingTagsOk

`func (o *ConfigSettingValueModel) GetSettingTagsOk() (*[]SettingTagModel, bool)`

GetSettingTagsOk returns a tuple with the SettingTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingTags

`func (o *ConfigSettingValueModel) SetSettingTags(v []SettingTagModel)`

SetSettingTags sets SettingTags field to given value.

### HasSettingTags

`func (o *ConfigSettingValueModel) HasSettingTags() bool`

HasSettingTags returns a boolean if a field has been set.

### SetSettingTagsNil

`func (o *ConfigSettingValueModel) SetSettingTagsNil(b bool)`

 SetSettingTagsNil sets the value for SettingTags to be an explicit nil

### UnsetSettingTags
`func (o *ConfigSettingValueModel) UnsetSettingTags()`

UnsetSettingTags ensures that no value is present for SettingTags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


