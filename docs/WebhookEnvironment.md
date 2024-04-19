# WebhookEnvironment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | The Environment&#39;s name. | [optional] 
**EnvironmentId** | Pointer to **string** | The Environment&#39;s identifier. | [optional] 

## Methods

### NewWebhookEnvironment

`func NewWebhookEnvironment() *WebhookEnvironment`

NewWebhookEnvironment instantiates a new WebhookEnvironment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookEnvironmentWithDefaults

`func NewWebhookEnvironmentWithDefaults() *WebhookEnvironment`

NewWebhookEnvironmentWithDefaults instantiates a new WebhookEnvironment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WebhookEnvironment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebhookEnvironment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebhookEnvironment) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WebhookEnvironment) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *WebhookEnvironment) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *WebhookEnvironment) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetEnvironmentId

`func (o *WebhookEnvironment) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *WebhookEnvironment) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *WebhookEnvironment) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *WebhookEnvironment) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


