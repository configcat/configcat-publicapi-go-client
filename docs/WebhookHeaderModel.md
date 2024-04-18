# WebhookHeaderModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | The HTTP header key. | 
**Value** | **string** | The HTTP header value. | 
**IsSecure** | Pointer to **bool** | Indicates whether the header value is sensitive. | [optional] 

## Methods

### NewWebhookHeaderModel

`func NewWebhookHeaderModel(key string, value string, ) *WebhookHeaderModel`

NewWebhookHeaderModel instantiates a new WebhookHeaderModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookHeaderModelWithDefaults

`func NewWebhookHeaderModelWithDefaults() *WebhookHeaderModel`

NewWebhookHeaderModelWithDefaults instantiates a new WebhookHeaderModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *WebhookHeaderModel) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *WebhookHeaderModel) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *WebhookHeaderModel) SetKey(v string)`

SetKey sets Key field to given value.


### GetValue

`func (o *WebhookHeaderModel) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *WebhookHeaderModel) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *WebhookHeaderModel) SetValue(v string)`

SetValue sets Value field to given value.


### GetIsSecure

`func (o *WebhookHeaderModel) GetIsSecure() bool`

GetIsSecure returns the IsSecure field if non-nil, zero value otherwise.

### GetIsSecureOk

`func (o *WebhookHeaderModel) GetIsSecureOk() (*bool, bool)`

GetIsSecureOk returns a tuple with the IsSecure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSecure

`func (o *WebhookHeaderModel) SetIsSecure(v bool)`

SetIsSecure sets IsSecure field to given value.

### HasIsSecure

`func (o *WebhookHeaderModel) HasIsSecure() bool`

HasIsSecure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


