# SettingValueModelHaljson

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReadOnly** | **bool** |  | [optional] [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**LastUpdaterUserEmail** | **string** |  | [optional] [default to null]
**LastUpdaterUserFullName** | **string** |  | [optional] [default to null]
**Embedded** | [***SettingValueModelhaljsonEmbedded**](SettingValueModelhaljson__embedded.md) |  | [optional] [default to null]
**RolloutRules** | [**[]RolloutRuleModel**](RolloutRuleModel.md) | The targeting rule collection. | [optional] [default to null]
**RolloutPercentageItems** | [**[]RolloutPercentageItemModel**](RolloutPercentageItemModel.md) | The percentage rule collection. | [optional] [default to null]
**Value** | [***interface{}**](interface{}.md) | The value to serve. It must respect the setting type. | [optional] [default to null]
**Links** | [***EnvironmentModelhaljsonLinks**](EnvironmentModelhaljson__links.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

