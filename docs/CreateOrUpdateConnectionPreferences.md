# CreateOrUpdateConnectionPreferences

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SdkPollInterval** | Pointer to **NullableInt32** | The SDK poll interval in seconds. If not specified, a default value (60) will be used. | [optional] 
**WebhookNotification** | Pointer to [**NullableCreateOrUpdateWebhookNotification**](CreateOrUpdateWebhookNotification.md) |  | [optional] 

## Methods

### NewCreateOrUpdateConnectionPreferences

`func NewCreateOrUpdateConnectionPreferences() *CreateOrUpdateConnectionPreferences`

NewCreateOrUpdateConnectionPreferences instantiates a new CreateOrUpdateConnectionPreferences object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrUpdateConnectionPreferencesWithDefaults

`func NewCreateOrUpdateConnectionPreferencesWithDefaults() *CreateOrUpdateConnectionPreferences`

NewCreateOrUpdateConnectionPreferencesWithDefaults instantiates a new CreateOrUpdateConnectionPreferences object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSdkPollInterval

`func (o *CreateOrUpdateConnectionPreferences) GetSdkPollInterval() int32`

GetSdkPollInterval returns the SdkPollInterval field if non-nil, zero value otherwise.

### GetSdkPollIntervalOk

`func (o *CreateOrUpdateConnectionPreferences) GetSdkPollIntervalOk() (*int32, bool)`

GetSdkPollIntervalOk returns a tuple with the SdkPollInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdkPollInterval

`func (o *CreateOrUpdateConnectionPreferences) SetSdkPollInterval(v int32)`

SetSdkPollInterval sets SdkPollInterval field to given value.

### HasSdkPollInterval

`func (o *CreateOrUpdateConnectionPreferences) HasSdkPollInterval() bool`

HasSdkPollInterval returns a boolean if a field has been set.

### SetSdkPollIntervalNil

`func (o *CreateOrUpdateConnectionPreferences) SetSdkPollIntervalNil(b bool)`

 SetSdkPollIntervalNil sets the value for SdkPollInterval to be an explicit nil

### UnsetSdkPollInterval
`func (o *CreateOrUpdateConnectionPreferences) UnsetSdkPollInterval()`

UnsetSdkPollInterval ensures that no value is present for SdkPollInterval, not even an explicit nil
### GetWebhookNotification

`func (o *CreateOrUpdateConnectionPreferences) GetWebhookNotification() CreateOrUpdateWebhookNotification`

GetWebhookNotification returns the WebhookNotification field if non-nil, zero value otherwise.

### GetWebhookNotificationOk

`func (o *CreateOrUpdateConnectionPreferences) GetWebhookNotificationOk() (*CreateOrUpdateWebhookNotification, bool)`

GetWebhookNotificationOk returns a tuple with the WebhookNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookNotification

`func (o *CreateOrUpdateConnectionPreferences) SetWebhookNotification(v CreateOrUpdateWebhookNotification)`

SetWebhookNotification sets WebhookNotification field to given value.

### HasWebhookNotification

`func (o *CreateOrUpdateConnectionPreferences) HasWebhookNotification() bool`

HasWebhookNotification returns a boolean if a field has been set.

### SetWebhookNotificationNil

`func (o *CreateOrUpdateConnectionPreferences) SetWebhookNotificationNil(b bool)`

 SetWebhookNotificationNil sets the value for WebhookNotification to be an explicit nil

### UnsetWebhookNotification
`func (o *CreateOrUpdateConnectionPreferences) UnsetWebhookNotification()`

UnsetWebhookNotification ensures that no value is present for WebhookNotification, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


