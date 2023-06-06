# SettingValueModelHaljson

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RolloutRules** | Pointer to [**[]RolloutRuleModel**](RolloutRuleModel.md) | The targeting rule collection. | [optional] 
**RolloutPercentageItems** | Pointer to [**[]RolloutPercentageItemModel**](RolloutPercentageItemModel.md) | The percentage rule collection. | [optional] 
**Value** | Pointer to **interface{}** | The value to serve. It must respect the setting type. | [optional] 
**UpdatedAt** | Pointer to **NullableTime** |  | [optional] 
**LastUpdaterUserEmail** | Pointer to **NullableString** |  | [optional] 
**LastUpdaterUserFullName** | Pointer to **NullableString** |  | [optional] 
**Embedded** | Pointer to [**SettingValueModelHaljsonEmbedded**](SettingValueModelHaljsonEmbedded.md) |  | [optional] 
**ReadOnly** | Pointer to **bool** |  | [optional] 
**Links** | Pointer to [**EnvironmentModelHaljsonLinks**](EnvironmentModelHaljsonLinks.md) |  | [optional] 

## Methods

### NewSettingValueModelHaljson

`func NewSettingValueModelHaljson() *SettingValueModelHaljson`

NewSettingValueModelHaljson instantiates a new SettingValueModelHaljson object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettingValueModelHaljsonWithDefaults

`func NewSettingValueModelHaljsonWithDefaults() *SettingValueModelHaljson`

NewSettingValueModelHaljsonWithDefaults instantiates a new SettingValueModelHaljson object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRolloutRules

`func (o *SettingValueModelHaljson) GetRolloutRules() []RolloutRuleModel`

GetRolloutRules returns the RolloutRules field if non-nil, zero value otherwise.

### GetRolloutRulesOk

`func (o *SettingValueModelHaljson) GetRolloutRulesOk() (*[]RolloutRuleModel, bool)`

GetRolloutRulesOk returns a tuple with the RolloutRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolloutRules

`func (o *SettingValueModelHaljson) SetRolloutRules(v []RolloutRuleModel)`

SetRolloutRules sets RolloutRules field to given value.

### HasRolloutRules

`func (o *SettingValueModelHaljson) HasRolloutRules() bool`

HasRolloutRules returns a boolean if a field has been set.

### SetRolloutRulesNil

`func (o *SettingValueModelHaljson) SetRolloutRulesNil(b bool)`

 SetRolloutRulesNil sets the value for RolloutRules to be an explicit nil

### UnsetRolloutRules
`func (o *SettingValueModelHaljson) UnsetRolloutRules()`

UnsetRolloutRules ensures that no value is present for RolloutRules, not even an explicit nil
### GetRolloutPercentageItems

`func (o *SettingValueModelHaljson) GetRolloutPercentageItems() []RolloutPercentageItemModel`

GetRolloutPercentageItems returns the RolloutPercentageItems field if non-nil, zero value otherwise.

### GetRolloutPercentageItemsOk

`func (o *SettingValueModelHaljson) GetRolloutPercentageItemsOk() (*[]RolloutPercentageItemModel, bool)`

GetRolloutPercentageItemsOk returns a tuple with the RolloutPercentageItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolloutPercentageItems

`func (o *SettingValueModelHaljson) SetRolloutPercentageItems(v []RolloutPercentageItemModel)`

SetRolloutPercentageItems sets RolloutPercentageItems field to given value.

### HasRolloutPercentageItems

`func (o *SettingValueModelHaljson) HasRolloutPercentageItems() bool`

HasRolloutPercentageItems returns a boolean if a field has been set.

### SetRolloutPercentageItemsNil

`func (o *SettingValueModelHaljson) SetRolloutPercentageItemsNil(b bool)`

 SetRolloutPercentageItemsNil sets the value for RolloutPercentageItems to be an explicit nil

### UnsetRolloutPercentageItems
`func (o *SettingValueModelHaljson) UnsetRolloutPercentageItems()`

UnsetRolloutPercentageItems ensures that no value is present for RolloutPercentageItems, not even an explicit nil
### GetValue

`func (o *SettingValueModelHaljson) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SettingValueModelHaljson) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SettingValueModelHaljson) SetValue(v interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *SettingValueModelHaljson) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *SettingValueModelHaljson) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *SettingValueModelHaljson) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetUpdatedAt

`func (o *SettingValueModelHaljson) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SettingValueModelHaljson) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SettingValueModelHaljson) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SettingValueModelHaljson) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *SettingValueModelHaljson) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *SettingValueModelHaljson) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetLastUpdaterUserEmail

`func (o *SettingValueModelHaljson) GetLastUpdaterUserEmail() string`

GetLastUpdaterUserEmail returns the LastUpdaterUserEmail field if non-nil, zero value otherwise.

### GetLastUpdaterUserEmailOk

`func (o *SettingValueModelHaljson) GetLastUpdaterUserEmailOk() (*string, bool)`

GetLastUpdaterUserEmailOk returns a tuple with the LastUpdaterUserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdaterUserEmail

`func (o *SettingValueModelHaljson) SetLastUpdaterUserEmail(v string)`

SetLastUpdaterUserEmail sets LastUpdaterUserEmail field to given value.

### HasLastUpdaterUserEmail

`func (o *SettingValueModelHaljson) HasLastUpdaterUserEmail() bool`

HasLastUpdaterUserEmail returns a boolean if a field has been set.

### SetLastUpdaterUserEmailNil

`func (o *SettingValueModelHaljson) SetLastUpdaterUserEmailNil(b bool)`

 SetLastUpdaterUserEmailNil sets the value for LastUpdaterUserEmail to be an explicit nil

### UnsetLastUpdaterUserEmail
`func (o *SettingValueModelHaljson) UnsetLastUpdaterUserEmail()`

UnsetLastUpdaterUserEmail ensures that no value is present for LastUpdaterUserEmail, not even an explicit nil
### GetLastUpdaterUserFullName

`func (o *SettingValueModelHaljson) GetLastUpdaterUserFullName() string`

GetLastUpdaterUserFullName returns the LastUpdaterUserFullName field if non-nil, zero value otherwise.

### GetLastUpdaterUserFullNameOk

`func (o *SettingValueModelHaljson) GetLastUpdaterUserFullNameOk() (*string, bool)`

GetLastUpdaterUserFullNameOk returns a tuple with the LastUpdaterUserFullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdaterUserFullName

`func (o *SettingValueModelHaljson) SetLastUpdaterUserFullName(v string)`

SetLastUpdaterUserFullName sets LastUpdaterUserFullName field to given value.

### HasLastUpdaterUserFullName

`func (o *SettingValueModelHaljson) HasLastUpdaterUserFullName() bool`

HasLastUpdaterUserFullName returns a boolean if a field has been set.

### SetLastUpdaterUserFullNameNil

`func (o *SettingValueModelHaljson) SetLastUpdaterUserFullNameNil(b bool)`

 SetLastUpdaterUserFullNameNil sets the value for LastUpdaterUserFullName to be an explicit nil

### UnsetLastUpdaterUserFullName
`func (o *SettingValueModelHaljson) UnsetLastUpdaterUserFullName()`

UnsetLastUpdaterUserFullName ensures that no value is present for LastUpdaterUserFullName, not even an explicit nil
### GetEmbedded

`func (o *SettingValueModelHaljson) GetEmbedded() SettingValueModelHaljsonEmbedded`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *SettingValueModelHaljson) GetEmbeddedOk() (*SettingValueModelHaljsonEmbedded, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *SettingValueModelHaljson) SetEmbedded(v SettingValueModelHaljsonEmbedded)`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *SettingValueModelHaljson) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetReadOnly

`func (o *SettingValueModelHaljson) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *SettingValueModelHaljson) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *SettingValueModelHaljson) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *SettingValueModelHaljson) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetLinks

`func (o *SettingValueModelHaljson) GetLinks() EnvironmentModelHaljsonLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *SettingValueModelHaljson) GetLinksOk() (*EnvironmentModelHaljsonLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *SettingValueModelHaljson) SetLinks(v EnvironmentModelHaljsonLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *SettingValueModelHaljson) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


