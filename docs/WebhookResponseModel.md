# WebhookResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebhookId** | **int32** | The identifier of the Webhook. | 
**Url** | **string** | The URL of the Webhook. | 
**HttpMethod** | [**WebHookHttpMethod**](WebHookHttpMethod.md) |  | 
**Content** | **NullableString** | The HTTP body content. | 
**WebHookHeaders** | [**[]WebhookHeaderResponseModel**](WebhookHeaderResponseModel.md) | List of HTTP headers that the Webhook must send. | 
**Config** | [**WebhookConfig**](WebhookConfig.md) |  | 
**Environment** | [**WebhookEnvironment**](WebhookEnvironment.md) |  | 

## Methods

### NewWebhookResponseModel

`func NewWebhookResponseModel(webhookId int32, url string, httpMethod WebHookHttpMethod, content NullableString, webHookHeaders []WebhookHeaderResponseModel, config WebhookConfig, environment WebhookEnvironment, ) *WebhookResponseModel`

NewWebhookResponseModel instantiates a new WebhookResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookResponseModelWithDefaults

`func NewWebhookResponseModelWithDefaults() *WebhookResponseModel`

NewWebhookResponseModelWithDefaults instantiates a new WebhookResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebhookId

`func (o *WebhookResponseModel) GetWebhookId() int32`

GetWebhookId returns the WebhookId field if non-nil, zero value otherwise.

### GetWebhookIdOk

`func (o *WebhookResponseModel) GetWebhookIdOk() (*int32, bool)`

GetWebhookIdOk returns a tuple with the WebhookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookId

`func (o *WebhookResponseModel) SetWebhookId(v int32)`

SetWebhookId sets WebhookId field to given value.


### GetUrl

`func (o *WebhookResponseModel) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebhookResponseModel) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WebhookResponseModel) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetHttpMethod

`func (o *WebhookResponseModel) GetHttpMethod() WebHookHttpMethod`

GetHttpMethod returns the HttpMethod field if non-nil, zero value otherwise.

### GetHttpMethodOk

`func (o *WebhookResponseModel) GetHttpMethodOk() (*WebHookHttpMethod, bool)`

GetHttpMethodOk returns a tuple with the HttpMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpMethod

`func (o *WebhookResponseModel) SetHttpMethod(v WebHookHttpMethod)`

SetHttpMethod sets HttpMethod field to given value.


### GetContent

`func (o *WebhookResponseModel) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *WebhookResponseModel) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *WebhookResponseModel) SetContent(v string)`

SetContent sets Content field to given value.


### SetContentNil

`func (o *WebhookResponseModel) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *WebhookResponseModel) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetWebHookHeaders

`func (o *WebhookResponseModel) GetWebHookHeaders() []WebhookHeaderResponseModel`

GetWebHookHeaders returns the WebHookHeaders field if non-nil, zero value otherwise.

### GetWebHookHeadersOk

`func (o *WebhookResponseModel) GetWebHookHeadersOk() (*[]WebhookHeaderResponseModel, bool)`

GetWebHookHeadersOk returns a tuple with the WebHookHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebHookHeaders

`func (o *WebhookResponseModel) SetWebHookHeaders(v []WebhookHeaderResponseModel)`

SetWebHookHeaders sets WebHookHeaders field to given value.


### GetConfig

`func (o *WebhookResponseModel) GetConfig() WebhookConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *WebhookResponseModel) GetConfigOk() (*WebhookConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *WebhookResponseModel) SetConfig(v WebhookConfig)`

SetConfig sets Config field to given value.


### GetEnvironment

`func (o *WebhookResponseModel) GetEnvironment() WebhookEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *WebhookResponseModel) GetEnvironmentOk() (*WebhookEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *WebhookResponseModel) SetEnvironment(v WebhookEnvironment)`

SetEnvironment sets Environment field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


