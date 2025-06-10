# WebHookRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | The URL of the Webhook. | 
**Content** | Pointer to **NullableString** | The HTTP body content. | [optional] 
**HttpMethod** | Pointer to [**NullableWebHookHttpMethod**](WebHookHttpMethod.md) |  | [optional] 
**WebHookHeaders** | Pointer to [**[]WebhookHeaderModel**](WebhookHeaderModel.md) | List of HTTP headers. | [optional] 

## Methods

### NewWebHookRequestModel

`func NewWebHookRequestModel(url string, ) *WebHookRequestModel`

NewWebHookRequestModel instantiates a new WebHookRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebHookRequestModelWithDefaults

`func NewWebHookRequestModelWithDefaults() *WebHookRequestModel`

NewWebHookRequestModelWithDefaults instantiates a new WebHookRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *WebHookRequestModel) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebHookRequestModel) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WebHookRequestModel) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetContent

`func (o *WebHookRequestModel) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *WebHookRequestModel) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *WebHookRequestModel) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *WebHookRequestModel) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *WebHookRequestModel) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *WebHookRequestModel) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetHttpMethod

`func (o *WebHookRequestModel) GetHttpMethod() WebHookHttpMethod`

GetHttpMethod returns the HttpMethod field if non-nil, zero value otherwise.

### GetHttpMethodOk

`func (o *WebHookRequestModel) GetHttpMethodOk() (*WebHookHttpMethod, bool)`

GetHttpMethodOk returns a tuple with the HttpMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpMethod

`func (o *WebHookRequestModel) SetHttpMethod(v WebHookHttpMethod)`

SetHttpMethod sets HttpMethod field to given value.

### HasHttpMethod

`func (o *WebHookRequestModel) HasHttpMethod() bool`

HasHttpMethod returns a boolean if a field has been set.

### SetHttpMethodNil

`func (o *WebHookRequestModel) SetHttpMethodNil(b bool)`

 SetHttpMethodNil sets the value for HttpMethod to be an explicit nil

### UnsetHttpMethod
`func (o *WebHookRequestModel) UnsetHttpMethod()`

UnsetHttpMethod ensures that no value is present for HttpMethod, not even an explicit nil
### GetWebHookHeaders

`func (o *WebHookRequestModel) GetWebHookHeaders() []WebhookHeaderModel`

GetWebHookHeaders returns the WebHookHeaders field if non-nil, zero value otherwise.

### GetWebHookHeadersOk

`func (o *WebHookRequestModel) GetWebHookHeadersOk() (*[]WebhookHeaderModel, bool)`

GetWebHookHeadersOk returns a tuple with the WebHookHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebHookHeaders

`func (o *WebHookRequestModel) SetWebHookHeaders(v []WebhookHeaderModel)`

SetWebHookHeaders sets WebHookHeaders field to given value.

### HasWebHookHeaders

`func (o *WebHookRequestModel) HasWebHookHeaders() bool`

HasWebHookHeaders returns a boolean if a field has been set.

### SetWebHookHeadersNil

`func (o *WebHookRequestModel) SetWebHookHeadersNil(b bool)`

 SetWebHookHeadersNil sets the value for WebHookHeaders to be an explicit nil

### UnsetWebHookHeaders
`func (o *WebHookRequestModel) UnsetWebHookHeaders()`

UnsetWebHookHeaders ensures that no value is present for WebHookHeaders, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


