# ProxyProfileSdkKeyItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrimarySdkKey** | **string** | The primary SDK Key of the Config / Environment pair selected for the Proxy Profile. | 
**SecondarySdkKey** | **NullableString** | The secondary SDK Key of the Config / Environment pair selected for the Proxy Profile. | 
**SdkId** | **string** | The SDK ID identifying the Config / Environment pair selected for the Proxy Profile. | 
**ConfigId** | **string** | The identifier of the Config associated with the SDK key. | 
**EnvironmentId** | **string** | The identifier of the Environment associated with the SDK key. | 

## Methods

### NewProxyProfileSdkKeyItem

`func NewProxyProfileSdkKeyItem(primarySdkKey string, secondarySdkKey NullableString, sdkId string, configId string, environmentId string, ) *ProxyProfileSdkKeyItem`

NewProxyProfileSdkKeyItem instantiates a new ProxyProfileSdkKeyItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProxyProfileSdkKeyItemWithDefaults

`func NewProxyProfileSdkKeyItemWithDefaults() *ProxyProfileSdkKeyItem`

NewProxyProfileSdkKeyItemWithDefaults instantiates a new ProxyProfileSdkKeyItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrimarySdkKey

`func (o *ProxyProfileSdkKeyItem) GetPrimarySdkKey() string`

GetPrimarySdkKey returns the PrimarySdkKey field if non-nil, zero value otherwise.

### GetPrimarySdkKeyOk

`func (o *ProxyProfileSdkKeyItem) GetPrimarySdkKeyOk() (*string, bool)`

GetPrimarySdkKeyOk returns a tuple with the PrimarySdkKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimarySdkKey

`func (o *ProxyProfileSdkKeyItem) SetPrimarySdkKey(v string)`

SetPrimarySdkKey sets PrimarySdkKey field to given value.


### GetSecondarySdkKey

`func (o *ProxyProfileSdkKeyItem) GetSecondarySdkKey() string`

GetSecondarySdkKey returns the SecondarySdkKey field if non-nil, zero value otherwise.

### GetSecondarySdkKeyOk

`func (o *ProxyProfileSdkKeyItem) GetSecondarySdkKeyOk() (*string, bool)`

GetSecondarySdkKeyOk returns a tuple with the SecondarySdkKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondarySdkKey

`func (o *ProxyProfileSdkKeyItem) SetSecondarySdkKey(v string)`

SetSecondarySdkKey sets SecondarySdkKey field to given value.


### SetSecondarySdkKeyNil

`func (o *ProxyProfileSdkKeyItem) SetSecondarySdkKeyNil(b bool)`

 SetSecondarySdkKeyNil sets the value for SecondarySdkKey to be an explicit nil

### UnsetSecondarySdkKey
`func (o *ProxyProfileSdkKeyItem) UnsetSecondarySdkKey()`

UnsetSecondarySdkKey ensures that no value is present for SecondarySdkKey, not even an explicit nil
### GetSdkId

`func (o *ProxyProfileSdkKeyItem) GetSdkId() string`

GetSdkId returns the SdkId field if non-nil, zero value otherwise.

### GetSdkIdOk

`func (o *ProxyProfileSdkKeyItem) GetSdkIdOk() (*string, bool)`

GetSdkIdOk returns a tuple with the SdkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdkId

`func (o *ProxyProfileSdkKeyItem) SetSdkId(v string)`

SetSdkId sets SdkId field to given value.


### GetConfigId

`func (o *ProxyProfileSdkKeyItem) GetConfigId() string`

GetConfigId returns the ConfigId field if non-nil, zero value otherwise.

### GetConfigIdOk

`func (o *ProxyProfileSdkKeyItem) GetConfigIdOk() (*string, bool)`

GetConfigIdOk returns a tuple with the ConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigId

`func (o *ProxyProfileSdkKeyItem) SetConfigId(v string)`

SetConfigId sets ConfigId field to given value.


### GetEnvironmentId

`func (o *ProxyProfileSdkKeyItem) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *ProxyProfileSdkKeyItem) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *ProxyProfileSdkKeyItem) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


