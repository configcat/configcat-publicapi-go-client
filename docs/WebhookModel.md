# WebhookModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebhookId** | Pointer to **int32** | The identifier of the Webhook. | [optional] 
**Url** | Pointer to **NullableString** | The URL of the Webhook. | [optional] 
**HttpMethod** | Pointer to **string** | The HTTP method. | [optional] 
**Content** | Pointer to **NullableString** | The HTTP body content. | [optional] 
**WebHookHeaders** | Pointer to [**[]WebhookHeaderModel**](WebhookHeaderModel.md) | List of HTTP headers that the Webhook must send. | [optional] 
**Config** | Pointer to [**WebhookConfig**](WebhookConfig.md) |  | [optional] 
**Environment** | Pointer to [**WebhookEnvironment**](WebhookEnvironment.md) |  | [optional] 

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

`func (o *WebhookModel) GetHttpMethod() string`

GetHttpMethod returns the HttpMethod field if non-nil, zero value otherwise.

### GetHttpMethodOk

`func (o *WebhookModel) GetHttpMethodOk() (*string, bool)`

GetHttpMethodOk returns a tuple with the HttpMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpMethod

`func (o *WebhookModel) SetHttpMethod(v string)`

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
### GetConfig

`func (o *WebhookModel) GetConfig() WebhookConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *WebhookModel) GetConfigOk() (*WebhookConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *WebhookModel) SetConfig(v WebhookConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *WebhookModel) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetEnvironment

`func (o *WebhookModel) GetEnvironment() WebhookEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *WebhookModel) GetEnvironmentOk() (*WebhookEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *WebhookModel) SetEnvironment(v WebhookEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *WebhookModel) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


