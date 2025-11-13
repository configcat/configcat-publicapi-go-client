# ProxyProfileSdkKeysRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]ProxyProfileSdkKeyRequestItem**](ProxyProfileSdkKeyRequestItem.md) | The list of Config / Environment pairs. By only setting a &#x60;configId&#x60; means all Environments of that Config will be included. Similarly, by only setting an &#x60;environmentId&#x60; means all Configs with that Environment will be included. | 

## Methods

### NewProxyProfileSdkKeysRequest

`func NewProxyProfileSdkKeysRequest(items []ProxyProfileSdkKeyRequestItem, ) *ProxyProfileSdkKeysRequest`

NewProxyProfileSdkKeysRequest instantiates a new ProxyProfileSdkKeysRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProxyProfileSdkKeysRequestWithDefaults

`func NewProxyProfileSdkKeysRequestWithDefaults() *ProxyProfileSdkKeysRequest`

NewProxyProfileSdkKeysRequestWithDefaults instantiates a new ProxyProfileSdkKeysRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ProxyProfileSdkKeysRequest) GetItems() []ProxyProfileSdkKeyRequestItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ProxyProfileSdkKeysRequest) GetItemsOk() (*[]ProxyProfileSdkKeyRequestItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ProxyProfileSdkKeysRequest) SetItems(v []ProxyProfileSdkKeyRequestItem)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


