# WebhookModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebhookId** | Pointer to **int32** | The identifier of the Webhook. | [optional] 
**Url** | Pointer to **NullableString** | The URL of the Webhook. | [optional] 
**HttpMethod** | Pointer to [**WebHookHttpMethod**](WebHookHttpMethod.md) |  | [optional] 
**Content** | Pointer to **NullableString** | The HTTP body content. | [optional] 
**WebHookHeaders** | Pointer to [**[]WebhookHeaderModel**](WebhookHeaderModel.md) | List of HTTP headers that the Webhook must send. | [optional] 
**ConfigName** | Pointer to **NullableString** | The Config&#39;s name where the applied changes will invoke the Webhook. | [optional] 
**EnvironmentName** | Pointer to **NullableString** | The Environment&#39;s name where the applied changes will invoke the Webhook. | [optional] 

## Methods

### NewWebhookModel

`func NewWebhookModel() *WebhookModel`

NewWebhookModel instantiates a new WebhookModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookModelWithDefaults

`func NewWebhookModelWithDefaults() *WebhookModel`

NewWebhookModelWithDefaults instantiates a new WebhookModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebhookId

`func (o *WebhookModel) GetWebhookId() int32`

GetWebhookId returns the WebhookId field if non-nil, zero value otherwise.

### GetWebhookIdOk

`func (o *WebhookModel) GetWebhookIdOk() (*int32, bool)`

GetWebhookIdOk returns a tuple with the WebhookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookId

`func (o *WebhookModel) SetWebhookId(v int32)`

SetWebhookId sets WebhookId field to given value.

### HasWebhookId

`func (o *WebhookModel) HasWebhookId() bool`

HasWebhookId returns a boolean if a field has been set.

### GetUrl

`func (o *WebhookModel) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebhookModel) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WebhookModel) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *WebhookModel) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *WebhookModel) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *WebhookModel) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetHttpMethod

`func (o *WebhookModel) GetHttpMethod() WebHookHttpMethod`

GetHttpMethod returns the HttpMethod field if non-nil, zero value otherwise.

### GetHttpMethodOk

`func (o *WebhookModel) GetHttpMethodOk() (*WebHookHttpMethod, bool)`

GetHttpMethodOk returns a tuple with the HttpMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpMethod

`func (o *WebhookModel) SetHttpMethod(v WebHookHttpMethod)`

SetHttpMethod sets HttpMethod field to given value.

### HasHttpMethod

`func (o *WebhookModel) HasHttpMethod() bool`

HasHttpMethod returns a boolean if a field has been set.

### GetContent

`func (o *WebhookModel) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *WebhookModel) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *WebhookModel) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *WebhookModel) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *WebhookModel) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *WebhookModel) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetWebHookHeaders

`func (o *WebhookModel) GetWebHookHeaders() []WebhookHeaderModel`

GetWebHookHeaders returns the WebHookHeaders field if non-nil, zero value otherwise.

### GetWebHookHeadersOk

`func (o *WebhookModel) GetWebHookHeadersOk() (*[]WebhookHeaderModel, bool)`

GetWebHookHeadersOk returns a tuple with the WebHookHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebHookHeaders

`func (o *WebhookModel) SetWebHookHeaders(v []WebhookHeaderModel)`

SetWebHookHeaders sets WebHookHeaders field to given value.

### HasWebHookHeaders

`func (o *WebhookModel) HasWebHookHeaders() bool`

HasWebHookHeaders returns a boolean if a field has been set.

### SetWebHookHeadersNil

`func (o *WebhookModel) SetWebHookHeadersNil(b bool)`

 SetWebHookHeadersNil sets the value for WebHookHeaders to be an explicit nil

### UnsetWebHookHeaders
`func (o *WebhookModel) UnsetWebHookHeaders()`

UnsetWebHookHeaders ensures that no value is present for WebHookHeaders, not even an explicit nil
### GetConfigName

`func (o *WebhookModel) GetConfigName() string`

GetConfigName returns the ConfigName field if non-nil, zero value otherwise.

### GetConfigNameOk

`func (o *WebhookModel) GetConfigNameOk() (*string, bool)`

GetConfigNameOk returns a tuple with the ConfigName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigName

`func (o *WebhookModel) SetConfigName(v string)`

SetConfigName sets ConfigName field to given value.

### HasConfigName

`func (o *WebhookModel) HasConfigName() bool`

HasConfigName returns a boolean if a field has been set.

### SetConfigNameNil

`func (o *WebhookModel) SetConfigNameNil(b bool)`

 SetConfigNameNil sets the value for ConfigName to be an explicit nil

### UnsetConfigName
`func (o *WebhookModel) UnsetConfigName()`

UnsetConfigName ensures that no value is present for ConfigName, not even an explicit nil
### GetEnvironmentName

`func (o *WebhookModel) GetEnvironmentName() string`

GetEnvironmentName returns the EnvironmentName field if non-nil, zero value otherwise.

### GetEnvironmentNameOk

`func (o *WebhookModel) GetEnvironmentNameOk() (*string, bool)`

GetEnvironmentNameOk returns a tuple with the EnvironmentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentName

`func (o *WebhookModel) SetEnvironmentName(v string)`

SetEnvironmentName sets EnvironmentName field to given value.

### HasEnvironmentName

`func (o *WebhookModel) HasEnvironmentName() bool`

HasEnvironmentName returns a boolean if a field has been set.

### SetEnvironmentNameNil

`func (o *WebhookModel) SetEnvironmentNameNil(b bool)`

 SetEnvironmentNameNil sets the value for EnvironmentName to be an explicit nil

### UnsetEnvironmentName
`func (o *WebhookModel) UnsetEnvironmentName()`

UnsetEnvironmentName ensures that no value is present for EnvironmentName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


