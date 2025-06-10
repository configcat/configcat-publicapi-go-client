# ConfigSettingValueModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Setting** | [**SettingDataModel**](SettingDataModel.md) |  | 
**UpdatedAt** | **NullableTime** | The last updated date and time when the Feature Flag or Setting. | 
**LastUpdaterUserEmail** | **NullableString** | The email of the user who last updated the Feature Flag or Setting. | 
**LastUpdaterUserFullName** | **NullableString** | The name of the user who last updated the Feature Flag or Setting. | 
**IntegrationLinks** | [**[]IntegrationLinkModel**](IntegrationLinkModel.md) | The integration links attached to the Feature Flag or Setting. | 
**SettingTags** | [**[]SettingTagModel**](SettingTagModel.md) | The tags attached to the Feature Flag or Setting. | 
**RolloutRules** | [**[]RolloutRuleModel**](RolloutRuleModel.md) | The targeting rule collection. | 
**RolloutPercentageItems** | [**[]RolloutPercentageItemModel**](RolloutPercentageItemModel.md) | The percentage rule collection. | 
**Value** | [**SettingValueType**](SettingValueType.md) | The value to serve. It must respect the setting type. In some generated clients for strictly typed languages you may use double/float properties to handle integer values. | 

## Methods

### NewConfigSettingValueModel

`func NewConfigSettingValueModel(setting SettingDataModel, updatedAt NullableTime, lastUpdaterUserEmail NullableString, lastUpdaterUserFullName NullableString, integrationLinks []IntegrationLinkModel, settingTags []SettingTagModel, rolloutRules []RolloutRuleModel, rolloutPercentageItems []RolloutPercentageItemModel, value SettingValueType, ) *ConfigSettingValueModel`

NewConfigSettingValueModel instantiates a new ConfigSettingValueModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigSettingValueModelWithDefaults

`func NewConfigSettingValueModelWithDefaults() *ConfigSettingValueModel`

NewConfigSettingValueModelWithDefaults instantiates a new ConfigSettingValueModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


### GetValue

`func (o *ConfigSettingValueModel) GetValue() SettingValueType`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ConfigSettingValueModel) GetValueOk() (*SettingValueType, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ConfigSettingValueModel) SetValue(v SettingValueType)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


