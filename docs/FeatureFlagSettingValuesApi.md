# {{classname}}

All URIs are relative to *https://api.configcat.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSettingValue**](FeatureFlagSettingValuesApi.md#GetSettingValue) | **Get** /v1/environments/{environmentId}/settings/{settingId}/value | Get value
[**GetSettingValues**](FeatureFlagSettingValuesApi.md#GetSettingValues) | **Get** /v1/configs/{configId}/environments/{environmentId}/values | Get values
[**PostSettingValues**](FeatureFlagSettingValuesApi.md#PostSettingValues) | **Post** /v1/configs/{configId}/environments/{environmentId}/values | Post values
[**ReplaceSettingValue**](FeatureFlagSettingValuesApi.md#ReplaceSettingValue) | **Put** /v1/environments/{environmentId}/settings/{settingId}/value | Replace value
[**UpdateSettingValue**](FeatureFlagSettingValuesApi.md#UpdateSettingValue) | **Patch** /v1/environments/{environmentId}/settings/{settingId}/value | Update value

# **GetSettingValue**
> SettingValueModel GetSettingValue(ctx, environmentId, settingId)
Get value

This endpoint returns the value of a Feature Flag or Setting  in a specified Environment identified by the `environmentId` parameter.  The most important attributes in the response are the `value`, `rolloutRules` and `percentageRules`. The `value` represents what the clients will get when the evaluation requests of our SDKs  are not matching to any of the defined Targeting or Percentage Rules, or when there are no additional rules to evaluate.  The `rolloutRules` and `percentageRules` attributes are representing the current  Targeting and Percentage Rules configuration of the actual Feature Flag or Setting  in an **ordered** collection, which means the order of the returned rules is matching to the evaluation order. You can read more about these rules [here](https://configcat.com/docs/advanced/targeting/).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **environmentId** | [**string**](.md)| The identifier of the Environment. | 
  **settingId** | **int32**| The id of the Setting. | 

### Return type

[**SettingValueModel**](SettingValueModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSettingValues**
> ConfigSettingValuesModel GetSettingValues(ctx, configId, environmentId)
Get values

This endpoint returns the value of a specified Config's Feature Flags or Settings identified by the `configId` parameter in a specified Environment identified by the `environmentId` parameter.  The most important attributes in the response are the `value`, `rolloutRules` and `percentageRules`. The `value` represents what the clients will get when the evaluation requests of our SDKs  are not matching to any of the defined Targeting or Percentage Rules, or when there are no additional rules to evaluate.  The `rolloutRules` and `percentageRules` attributes are representing the current  Targeting and Percentage Rules configuration of the actual Feature Flag or Setting  in an **ordered** collection, which means the order of the returned rules is matching to the evaluation order. You can read more about these rules [here](https://configcat.com/docs/advanced/targeting/).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **configId** | [**string**](.md)| The identifier of the Config. | 
  **environmentId** | [**string**](.md)| The identifier of the Environment. | 

### Return type

[**ConfigSettingValuesModel**](ConfigSettingValuesModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostSettingValues**
> ConfigSettingValuesModel PostSettingValues(ctx, body, configId, environmentId, optional)
Post values

This endpoint replaces the values of a specified Config's Feature Flags or Settings identified by the `configId` parameter in a specified Environment identified by the `environmentId` parameter.  Only the `value`, `rolloutRules` and `percentageRules` attributes are modifiable by this endpoint.  **Important:** As this endpoint is doing a complete replace, it's important to set every other attribute that you don't  want to change in its original state. Not listing one means that it will reset.  For example: We have the following resource. ``` {     \"settingValues\": [   {    \"rolloutPercentageItems\": [     {      \"percentage\": 30,      \"value\": true     },     {      \"percentage\": 70,      \"value\": false     }    ],    \"rolloutRules\": [],    \"value\": false,    \"settingId\": 1   }  ] } ``` If we send a replace request body as below: ``` {   \"settingValues\": [   {    \"value\": true,    \"settingId\": 1   }  ] } ``` Then besides that the default value is set to `true`, all the Percentage Rules are deleted.  So we get a response like this: ``` {  \"settingValues\": [   {    \"rolloutPercentageItems\": [],    \"rolloutRules\": [],    \"value\": true,    \"setting\":     {     \"settingId\": 1    }   }  ] } ```  The `rolloutRules` property describes two types of rules:  - **Targeting rules**: When you want to add or update a targenting rule, the `comparator`, `comparisonAttribute`, and `comparisonValue` members are required. - **Segment rules**: When you want to add add or update a segment rule, the `segmentId` which identifies the desired segment and the `segmentComparator` members are required.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateSettingValuesWithIdModel**](UpdateSettingValuesWithIdModel.md)|  | 
  **configId** | [**string**](.md)| The identifier of the Config. | 
  **environmentId** | [**string**](.md)| The identifier of the Environment. | 
 **optional** | ***FeatureFlagSettingValuesApiPostSettingValuesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FeatureFlagSettingValuesApiPostSettingValuesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **reason** | **optional.**| The reason note for the Audit Log if the Product&#x27;s \&quot;Config changes require a reason\&quot; preference is turned on. | 

### Return type

[**ConfigSettingValuesModel**](ConfigSettingValuesModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/_*+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceSettingValue**
> SettingValueModel ReplaceSettingValue(ctx, body, environmentId, settingId, optional)
Replace value

This endpoint replaces the whole value of a Feature Flag or Setting in a specified Environment.  Only the `value`, `rolloutRules` and `percentageRules` attributes are modifiable by this endpoint.  **Important:** As this endpoint is doing a complete replace, it's important to set every other attribute that you don't  want to change in its original state. Not listing one means that it will reset.  For example: We have the following resource. ``` {  \"rolloutPercentageItems\": [   {    \"percentage\": 30,    \"value\": true   },   {    \"percentage\": 70,    \"value\": false   }  ],  \"rolloutRules\": [],  \"value\": false } ``` If we send a replace request body as below: ``` {  \"value\": true } ``` Then besides that the default value is set to `true`, all the Percentage Rules are deleted.  So we get a response like this: ``` {  \"rolloutPercentageItems\": [],  \"rolloutRules\": [],  \"value\": true } ```  The `rolloutRules` property describes two types of rules:  - **Targeting rules**: When you want to add or update a targenting rule, the `comparator`, `comparisonAttribute`, and `comparisonValue` members are required. - **Segment rules**: When you want to add add or update a segment rule, the `segmentId` which identifies the desired segment and the `segmentComparator` members are required.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateSettingValueModel**](UpdateSettingValueModel.md)|  | 
  **environmentId** | [**string**](.md)| The identifier of the Environment. | 
  **settingId** | **int32**| The id of the Setting. | 
 **optional** | ***FeatureFlagSettingValuesApiReplaceSettingValueOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FeatureFlagSettingValuesApiReplaceSettingValueOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **reason** | **optional.**| The reason note for the Audit Log if the Product&#x27;s \&quot;Config changes require a reason\&quot; preference is turned on. | 

### Return type

[**SettingValueModel**](SettingValueModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/_*+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSettingValue**
> SettingValueModel UpdateSettingValue(ctx, body, environmentId, settingId, optional)
Update value

This endpoint updates the value of a Feature Flag or Setting  with a collection of [JSON Patch](http://jsonpatch.com) operations in a specified Environment.  Only the `value`, `rolloutRules` and `percentageRules` attributes are modifiable by this endpoint.  The advantage of using JSON Patch is that you can describe individual update operations on a resource without touching attributes that you don't want to change. It supports collection reordering, so it also  can be used for reordering the targeting rules of a Feature Flag or Setting.  For example: We have the following resource. ``` {  \"rolloutPercentageItems\": [   {    \"percentage\": 30,    \"value\": true   },   {    \"percentage\": 70,    \"value\": false   }  ],  \"rolloutRules\": [],  \"value\": false } ``` If we send an update request body as below: ``` [  {   \"op\": \"replace\",   \"path\": \"/value\",   \"value\": true  } ] ``` Only the default value is going to be set to `true` and all the Percentage Rules are remaining unchanged. So we get a response like this: ``` {  \"rolloutPercentageItems\": [   {    \"percentage\": 30,    \"value\": true   },   {    \"percentage\": 70,    \"value\": false   }  ],  \"rolloutRules\": [],  \"value\": true } ```  The `rolloutRules` property describes two types of rules:  - **Targeting rules**: When you want to add or update a targenting rule, the `comparator`, `comparisonAttribute`, and `comparisonValue` members are required. - **Segment rules**: When you want to add add or update a segment rule, the `segmentId` which identifies the desired segment and the `segmentComparator` members are required.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**JsonPatch**](JsonPatch.md)|  | 
  **environmentId** | [**string**](.md)| The identifier of the Environment. | 
  **settingId** | **int32**| The id of the Setting. | 
 **optional** | ***FeatureFlagSettingValuesApiUpdateSettingValueOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FeatureFlagSettingValuesApiUpdateSettingValueOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **reason** | **optional.**| The reason note for the Audit Log if the Product&#x27;s \&quot;Config changes require a reason\&quot; preference is turned on. | 

### Return type

[**SettingValueModel**](SettingValueModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/_*+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

