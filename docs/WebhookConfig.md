# WebhookConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | The Config&#39;s name. | [optional] 
**ConfigId** | Pointer to **string** | The Config&#39;s identifier. | [optional] 

## Methods

### NewWebhookConfig

`func NewWebhookConfig() *WebhookConfig`

NewWebhookConfig instantiates a new WebhookConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookConfigWithDefaults

`func NewWebhookConfigWithDefaults() *WebhookConfig`

NewWebhookConfigWithDefaults instantiates a new WebhookConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WebhookConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebhookConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebhookConfig) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WebhookConfig) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *WebhookConfig) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *WebhookConfig) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetConfigId

`func (o *WebhookConfig) GetConfigId() string`

GetConfigId returns the ConfigId field if non-nil, zero value otherwise.

### GetConfigIdOk

`func (o *WebhookConfig) GetConfigIdOk() (*string, bool)`

GetConfigIdOk returns a tuple with the ConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigId

`func (o *WebhookConfig) SetConfigId(v string)`

SetConfigId sets ConfigId field to given value.

### HasConfigId

`func (o *WebhookConfig) HasConfigId() bool`

HasConfigId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


