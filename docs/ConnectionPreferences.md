# ConnectionPreferences

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SdkPollInterval** | **int32** | The SDK poll interval in seconds that determines how often SDKs should fetch config JSON updates. | 
**WebhookNotification** | [**NullableWebhookNotification**](WebhookNotification.md) |  | 

## Methods

### NewConnectionPreferences

`func NewConnectionPreferences(sdkPollInterval int32, webhookNotification NullableWebhookNotification, ) *ConnectionPreferences`

NewConnectionPreferences instantiates a new ConnectionPreferences object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionPreferencesWithDefaults

`func NewConnectionPreferencesWithDefaults() *ConnectionPreferences`

NewConnectionPreferencesWithDefaults instantiates a new ConnectionPreferences object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSdkPollInterval

`func (o *ConnectionPreferences) GetSdkPollInterval() int32`

GetSdkPollInterval returns the SdkPollInterval field if non-nil, zero value otherwise.

### GetSdkPollIntervalOk

`func (o *ConnectionPreferences) GetSdkPollIntervalOk() (*int32, bool)`

GetSdkPollIntervalOk returns a tuple with the SdkPollInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdkPollInterval

`func (o *ConnectionPreferences) SetSdkPollInterval(v int32)`

SetSdkPollInterval sets SdkPollInterval field to given value.


### GetWebhookNotification

`func (o *ConnectionPreferences) GetWebhookNotification() WebhookNotification`

GetWebhookNotification returns the WebhookNotification field if non-nil, zero value otherwise.

### GetWebhookNotificationOk

`func (o *ConnectionPreferences) GetWebhookNotificationOk() (*WebhookNotification, bool)`

GetWebhookNotificationOk returns a tuple with the WebhookNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookNotification

`func (o *ConnectionPreferences) SetWebhookNotification(v WebhookNotification)`

SetWebhookNotification sets WebhookNotification field to given value.


### SetWebhookNotificationNil

`func (o *ConnectionPreferences) SetWebhookNotificationNil(b bool)`

 SetWebhookNotificationNil sets the value for WebhookNotification to be an explicit nil

### UnsetWebhookNotification
`func (o *ConnectionPreferences) UnsetWebhookNotification()`

UnsetWebhookNotification ensures that no value is present for WebhookNotification, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


