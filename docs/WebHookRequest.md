# WebHookRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | The URL of the Webhook. | 
**Content** | Pointer to **NullableString** | The HTTP body content. | [optional] 
**HttpMethod** | Pointer to [**WebHookHttpMethod**](WebHookHttpMethod.md) |  | [optional] 
**WebHookHeaders** | Pointer to [**[]WebhookHeaderModel**](WebhookHeaderModel.md) | List of HTTP headers. | [optional] 

## Methods

### NewWebHookRequest

`func NewWebHookRequest(url string, ) *WebHookRequest`

NewWebHookRequest instantiates a new WebHookRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebHookRequestWithDefaults

`func NewWebHookRequestWithDefaults() *WebHookRequest`

NewWebHookRequestWithDefaults instantiates a new WebHookRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *WebHookRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebHookRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WebHookRequest) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetContent

`func (o *WebHookRequest) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *WebHookRequest) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *WebHookRequest) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *WebHookRequest) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *WebHookRequest) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *WebHookRequest) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetHttpMethod

`func (o *WebHookRequest) GetHttpMethod() WebHookHttpMethod`

GetHttpMethod returns the HttpMethod field if non-nil, zero value otherwise.

### GetHttpMethodOk

`func (o *WebHookRequest) GetHttpMethodOk() (*WebHookHttpMethod, bool)`

GetHttpMethodOk returns a tuple with the HttpMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpMethod

`func (o *WebHookRequest) SetHttpMethod(v WebHookHttpMethod)`

SetHttpMethod sets HttpMethod field to given value.

### HasHttpMethod

`func (o *WebHookRequest) HasHttpMethod() bool`

HasHttpMethod returns a boolean if a field has been set.

### GetWebHookHeaders

`func (o *WebHookRequest) GetWebHookHeaders() []WebhookHeaderModel`

GetWebHookHeaders returns the WebHookHeaders field if non-nil, zero value otherwise.

### GetWebHookHeadersOk

`func (o *WebHookRequest) GetWebHookHeadersOk() (*[]WebhookHeaderModel, bool)`

GetWebHookHeadersOk returns a tuple with the WebHookHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebHookHeaders

`func (o *WebHookRequest) SetWebHookHeaders(v []WebhookHeaderModel)`

SetWebHookHeaders sets WebHookHeaders field to given value.

### HasWebHookHeaders

`func (o *WebHookRequest) HasWebHookHeaders() bool`

HasWebHookHeaders returns a boolean if a field has been set.

### SetWebHookHeadersNil

`func (o *WebHookRequest) SetWebHookHeadersNil(b bool)`

 SetWebHookHeadersNil sets the value for WebHookHeaders to be an explicit nil

### UnsetWebHookHeaders
`func (o *WebHookRequest) UnsetWebHookHeaders()`

UnsetWebHookHeaders ensures that no value is present for WebHookHeaders, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


