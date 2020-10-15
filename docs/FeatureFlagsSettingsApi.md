# {{classname}}

All URIs are relative to *https://api.configcat.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSetting**](FeatureFlagsSettingsApi.md#CreateSetting) | **Post** /v1/configs/{configId}/settings | Create Flag
[**DeleteSetting**](FeatureFlagsSettingsApi.md#DeleteSetting) | **Delete** /v1/settings/{settingId} | Delete Flag
[**GetSetting**](FeatureFlagsSettingsApi.md#GetSetting) | **Get** /v1/settings/{settingId} | Get Flag
[**GetSettings**](FeatureFlagsSettingsApi.md#GetSettings) | **Get** /v1/configs/{configId}/settings | List Flags
[**UpdateSetting**](FeatureFlagsSettingsApi.md#UpdateSetting) | **Patch** /v1/settings/{settingId} | Update Flag

# **CreateSetting**
> SettingModel CreateSetting(ctx, body, configId)
Create Flag

This endpoint creates a new Feature Flag or Setting in a specified Config identified by the `configId` parameter.  **Important:** The `key` attribute must be unique within the given Config.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateSettingModel**](CreateSettingModel.md)|  | 
  **configId** | [**string**](.md)| The identifier of the Config. | 

### Return type

[**SettingModel**](SettingModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/_*+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSetting**
> DeleteSetting(ctx, settingId)
Delete Flag

This endpoint removes a Feature Flag or Setting from a specified Config,  identified by the `configId` parameter.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **settingId** | **int32**| The identifier of the Setting. | 

### Return type

 (empty response body)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSetting**
> SettingModel GetSetting(ctx, settingId)
Get Flag

This endpoint returns the metadata attributes of a Feature Flag or Setting  identified by the `settingId` parameter.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **settingId** | **int32**| The identifier of the Setting. | 

### Return type

[**SettingModel**](SettingModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSettings**
> []SettingModel GetSettings(ctx, configId)
List Flags

This endpoint returns the list of the Feature Flags and Settings defined in a  specified Config, identified by the `configId` parameter.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **configId** | [**string**](.md)| The identifier of the Config. | 

### Return type

[**[]SettingModel**](SettingModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSetting**
> SettingModel UpdateSetting(ctx, body, settingId)
Update Flag

This endpoint updates the metadata of a Feature Flag or Setting  with a collection of [JSON Patch](http://jsonpatch.com) operations in a specified Config.  Only the `name`, `hint` and `tags` attributes are modifiable by this endpoint. The `tags` attribute is a simple collection of the [tag IDs](#operation/get-tags) attached to the given setting.  The advantage of using JSON Patch is that you can describe individual update operations on a resource without touching attributes that you don't want to change.  For example: We have the following resource. ``` {  \"settingId\": 5345,  \"key\": \"myAwesomeFeature\",  \"name\": \"Tihs is a naem with soem typos.\",  \"hint\": \"This Flag controls my awesome feature.\",  \"settingType\": \"boolean\",  \"tags\": [   {    \"tagId\": 0,    \"name\": \"sample tag\",    \"color\": \"whale\"   }  ] } ``` If we send an update request body as below (it changes the name and adds the already existing tag with the id 2): ``` [  {   \"op\": \"replace\",   \"path\": \"/name\",   \"value\": \"This is the name without typos.\"  },  {   \"op\": \"add\",   \"path\": \"/tags/-\",   \"value\": 2  } ] ``` Only the `name` and `tags` are going to be updated and all the other attributes are remaining unchanged. So we get a response like this: ``` {  \"settingId\": 5345,  \"key\": \"myAwesomeFeature\",  \"name\": \"This is the name without typos.\",  \"hint\": \"This Flag controls my awesome feature.\",  \"settingType\": \"boolean\",  \"tags\": [   {    \"tagId\": 0,    \"name\": \"sample tag\",    \"color\": \"whale\"   },   {    \"tagId\": 2,    \"name\": \"another tag\",    \"color\": \"koala\"   }  ] } ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]Operation**](Operation.md)|  | 
  **settingId** | **int32**| The identifier of the Setting. | 

### Return type

[**SettingModel**](SettingModel.md)

### Authorization

[Basic](../README.md#Basic)

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
 - **Accept**: application/json, application/hal+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

