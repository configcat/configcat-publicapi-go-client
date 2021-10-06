# {{classname}}

All URIs are relative to *https://api.configcat.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSettingValueBySdkkey**](FeatureFlagSettingValuesUsingSDKKeyApi.md#GetSettingValueBySdkkey) | **Get** /v1/settings/{settingKeyOrId}/value | Get value
[**ReplaceSettingValueBySdkkey**](FeatureFlagSettingValuesUsingSDKKeyApi.md#ReplaceSettingValueBySdkkey) | **Put** /v1/settings/{settingKeyOrId}/value | Replace value
[**UpdateSettingValueBySdkkey**](FeatureFlagSettingValuesUsingSDKKeyApi.md#UpdateSettingValueBySdkkey) | **Patch** /v1/settings/{settingKeyOrId}/value | Update value

# **GetSettingValueBySdkkey**
> SettingValueModel GetSettingValueBySdkkey(ctx, settingKeyOrId, optional)
Get value

This endpoint returns the value of a Feature Flag or Setting  in a specified Environment identified by the <a target=\"_blank\" rel=\"noopener noreferrer\" href=\"https://app.configcat.com/sdkkey\">SDK key</a> passed in the `X-CONFIGCAT-SDKKEY` header.  The most important attributes in the response are the `value`, `rolloutRules` and `percentageRules`. The `value` represents what the clients will get when the evaluation requests of our SDKs  are not matching to any of the defined Targeting or Percentage Rules, or when there are no additional rules to evaluate.  The `rolloutRules` and `percentageRules` attributes are representing the current  Targeting and Percentage Rules configuration of the actual Feature Flag or Setting  in an **ordered** collection, which means the order of the returned rules is matching to the evaluation order. You can read more about these rules [here](https://configcat.com/docs/advanced/targeting/).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **settingKeyOrId** | **string**| The key or id of the Setting. | 
 **optional** | ***FeatureFlagSettingValuesUsingSDKKeyApiGetSettingValueBySdkkeyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FeatureFlagSettingValuesUsingSDKKeyApiGetSettingValueBySdkkeyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xCONFIGCATSDKKEY** | **optional.String**| The ConfigCat SDK Key. (https://app.configcat.com/sdkkey) | 

### Return type

[**SettingValueModel**](SettingValueModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceSettingValueBySdkkey**
> SettingValueModel ReplaceSettingValueBySdkkey(ctx, body, settingKeyOrId, optional)
Replace value

This endpoint replaces the value of a Feature Flag or Setting  in a specified Environment identified by the <a target=\"_blank\" rel=\"noopener noreferrer\" href=\"https://app.configcat.com/sdkkey\">SDK key</a> passed in the `X-CONFIGCAT-SDKKEY` header.  Only the `value`, `rolloutRules` and `percentageRules` attributes are modifiable by this endpoint.  **Important:** As this endpoint is doing a complete replace, it's important to set every other attribute that you don't  want to change to its original state. Not listing one means that it will reset.  For example: We have the following resource. ``` {  \"rolloutPercentageItems\": [   {    \"percentage\": 30,    \"value\": true   },   {    \"percentage\": 70,    \"value\": false   }  ],  \"rolloutRules\": [],  \"value\": false } ``` If we send a replace request body as below: ``` {  \"value\": true } ``` Then besides that the default served value is set to `true`, all the Percentage Rules are deleted.  So we get a response like this: ``` {  \"rolloutPercentageItems\": [],  \"rolloutRules\": [],  \"value\": true } ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateSettingValueModel**](UpdateSettingValueModel.md)|  | 
  **settingKeyOrId** | **string**| The key or id of the Setting. | 
 **optional** | ***FeatureFlagSettingValuesUsingSDKKeyApiReplaceSettingValueBySdkkeyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FeatureFlagSettingValuesUsingSDKKeyApiReplaceSettingValueBySdkkeyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **reason** | **optional.**| The reason note for the Audit Log if the Product&#x27;s \&quot;Config changes require a reason\&quot; preference is turned on. | 
 **xCONFIGCATSDKKEY** | **optional.**| The ConfigCat SDK Key. (https://app.configcat.com/sdkkey) | 

### Return type

[**SettingValueModel**](SettingValueModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/_*+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSettingValueBySdkkey**
> SettingValueModel UpdateSettingValueBySdkkey(ctx, body, settingKeyOrId, optional)
Update value

This endpoint updates the value of a Feature Flag or Setting  with a collection of [JSON Patch](http://jsonpatch.com) operations in a specified Environment identified by the <a target=\"_blank\" rel=\"noopener noreferrer\" href=\"https://app.configcat.com/sdkkey\">SDK key</a> passed in the `X-CONFIGCAT-SDKKEY` header.  Only the `value`, `rolloutRules` and `percentageRules` attributes are modifiable by this endpoint.  The advantage of using JSON Patch is that you can describe individual update operations on a resource without touching attributes that you don't want to change. It supports collection reordering, so it also  can be used for reordering the targeting rules of a Feature Flag or Setting.  For example: We have the following resource. ``` {  \"rolloutPercentageItems\": [   {    \"percentage\": 30,    \"value\": true   },   {    \"percentage\": 70,    \"value\": false   }  ],  \"rolloutRules\": [],  \"value\": false } ``` If we send an update request body as below: ``` [  {   \"op\": \"replace\",   \"path\": \"/value\",   \"value\": true  } ] ``` Only the default served value is going to be set to `true` and all the Percentage Rules are remaining unchanged. So we get a response like this: ``` {  \"rolloutPercentageItems\": [   {    \"percentage\": 30,    \"value\": true   },   {    \"percentage\": 70,    \"value\": false   }  ],  \"rolloutRules\": [],  \"value\": true } ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]Operation**](Operation.md)|  | 
  **settingKeyOrId** | **string**| The key or id of the Setting. | 
 **optional** | ***FeatureFlagSettingValuesUsingSDKKeyApiUpdateSettingValueBySdkkeyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FeatureFlagSettingValuesUsingSDKKeyApiUpdateSettingValueBySdkkeyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **reason** | **optional.**| The reason note for the Audit Log if the Product&#x27;s \&quot;Config changes require a reason\&quot; preference is turned on. | 
 **xCONFIGCATSDKKEY** | **optional.**| The ConfigCat SDK Key. (https://app.configcat.com/sdkkey) | 

### Return type

[**SettingValueModel**](SettingValueModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

