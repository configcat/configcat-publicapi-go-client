# WebhookNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebhookProxyUrl** | **string** | The webhook proxy URL for receiving config JSON change notifications. | 
**SigningKey1** | **string** | The primary signing key used for verifying the authenticity of webhook requests. | 
**SigningKey2** | **NullableString** | The secondary signing key used for verifying the authenticity of webhook requests. | 

## Methods

### NewWebhookNotification

`func NewWebhookNotification(webhookProxyUrl string, signingKey1 string, signingKey2 NullableString, ) *WebhookNotification`

NewWebhookNotification instantiates a new WebhookNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookNotificationWithDefaults

`func NewWebhookNotificationWithDefaults() *WebhookNotification`

NewWebhookNotificationWithDefaults instantiates a new WebhookNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebhookProxyUrl

`func (o *WebhookNotification) GetWebhookProxyUrl() string`

GetWebhookProxyUrl returns the WebhookProxyUrl field if non-nil, zero value otherwise.

### GetWebhookProxyUrlOk

`func (o *WebhookNotification) GetWebhookProxyUrlOk() (*string, bool)`

GetWebhookProxyUrlOk returns a tuple with the WebhookProxyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookProxyUrl

`func (o *WebhookNotification) SetWebhookProxyUrl(v string)`

SetWebhookProxyUrl sets WebhookProxyUrl field to given value.


### GetSigningKey1

`func (o *WebhookNotification) GetSigningKey1() string`

GetSigningKey1 returns the SigningKey1 field if non-nil, zero value otherwise.

### GetSigningKey1Ok

`func (o *WebhookNotification) GetSigningKey1Ok() (*string, bool)`

GetSigningKey1Ok returns a tuple with the SigningKey1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningKey1

`func (o *WebhookNotification) SetSigningKey1(v string)`

SetSigningKey1 sets SigningKey1 field to given value.


### GetSigningKey2

`func (o *WebhookNotification) GetSigningKey2() string`

GetSigningKey2 returns the SigningKey2 field if non-nil, zero value otherwise.

### GetSigningKey2Ok

`func (o *WebhookNotification) GetSigningKey2Ok() (*string, bool)`

GetSigningKey2Ok returns a tuple with the SigningKey2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningKey2

`func (o *WebhookNotification) SetSigningKey2(v string)`

SetSigningKey2 sets SigningKey2 field to given value.


### SetSigningKey2Nil

`func (o *WebhookNotification) SetSigningKey2Nil(b bool)`

 SetSigningKey2Nil sets the value for SigningKey2 to be an explicit nil

### UnsetSigningKey2
`func (o *WebhookNotification) UnsetSigningKey2()`

UnsetSigningKey2 ensures that no value is present for SigningKey2, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


