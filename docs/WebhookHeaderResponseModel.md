# WebhookHeaderResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | The HTTP header key. | 
**Value** | **string** | The HTTP header value. | 
**IsSecure** | **bool** | Indicates whether the header value is sensitive. | 

## Methods

### NewWebhookHeaderResponseModel

`func NewWebhookHeaderResponseModel(key string, value string, isSecure bool, ) *WebhookHeaderResponseModel`

NewWebhookHeaderResponseModel instantiates a new WebhookHeaderResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookHeaderResponseModelWithDefaults

`func NewWebhookHeaderResponseModelWithDefaults() *WebhookHeaderResponseModel`

NewWebhookHeaderResponseModelWithDefaults instantiates a new WebhookHeaderResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *WebhookHeaderResponseModel) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *WebhookHeaderResponseModel) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *WebhookHeaderResponseModel) SetKey(v string)`

SetKey sets Key field to given value.


### GetValue

`func (o *WebhookHeaderResponseModel) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *WebhookHeaderResponseModel) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *WebhookHeaderResponseModel) SetValue(v string)`

SetValue sets Value field to given value.


### GetIsSecure

`func (o *WebhookHeaderResponseModel) GetIsSecure() bool`

GetIsSecure returns the IsSecure field if non-nil, zero value otherwise.

### GetIsSecureOk

`func (o *WebhookHeaderResponseModel) GetIsSecureOk() (*bool, bool)`

GetIsSecureOk returns a tuple with the IsSecure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSecure

`func (o *WebhookHeaderResponseModel) SetIsSecure(v bool)`

SetIsSecure sets IsSecure field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


