# DetailedStatisticV2Model

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | **time.Time** | The date for which the usage was recorded. | 
**ProductId** | **string** | The identifier of the Product associated with the usage. | 
**ProductName** | **string** | The name of the Product associated with the usage. | 
**ConfigId** | **string** | The identifier of the Config associated with the usage. | 
**ConfigName** | **string** | The name of the Config associated with the usage. | 
**EnvironmentId** | **NullableString** | The identifier of the Environment associated with the usage, if available. | 
**EnvironmentName** | **string** | The name of the Environment associated with the usage. | 
**Sdk** | **string** | The SDK type that generated the usage. | 
**SdkKey** | **string** | The SDK key used for the request. | 
**RequestCount** | **int64** | The number of requests recorded for the entry. | 
**ResponseKiloBytes** | **float64** | The total response payload size in kilobytes for the entry. | 

## Methods

### NewDetailedStatisticV2Model

`func NewDetailedStatisticV2Model(date time.Time, productId string, productName string, configId string, configName string, environmentId NullableString, environmentName string, sdk string, sdkKey string, requestCount int64, responseKiloBytes float64, ) *DetailedStatisticV2Model`

NewDetailedStatisticV2Model instantiates a new DetailedStatisticV2Model object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDetailedStatisticV2ModelWithDefaults

`func NewDetailedStatisticV2ModelWithDefaults() *DetailedStatisticV2Model`

NewDetailedStatisticV2ModelWithDefaults instantiates a new DetailedStatisticV2Model object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *DetailedStatisticV2Model) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *DetailedStatisticV2Model) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *DetailedStatisticV2Model) SetDate(v time.Time)`

SetDate sets Date field to given value.


### GetProductId

`func (o *DetailedStatisticV2Model) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *DetailedStatisticV2Model) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *DetailedStatisticV2Model) SetProductId(v string)`

SetProductId sets ProductId field to given value.


### GetProductName

`func (o *DetailedStatisticV2Model) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *DetailedStatisticV2Model) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *DetailedStatisticV2Model) SetProductName(v string)`

SetProductName sets ProductName field to given value.


### GetConfigId

`func (o *DetailedStatisticV2Model) GetConfigId() string`

GetConfigId returns the ConfigId field if non-nil, zero value otherwise.

### GetConfigIdOk

`func (o *DetailedStatisticV2Model) GetConfigIdOk() (*string, bool)`

GetConfigIdOk returns a tuple with the ConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigId

`func (o *DetailedStatisticV2Model) SetConfigId(v string)`

SetConfigId sets ConfigId field to given value.


### GetConfigName

`func (o *DetailedStatisticV2Model) GetConfigName() string`

GetConfigName returns the ConfigName field if non-nil, zero value otherwise.

### GetConfigNameOk

`func (o *DetailedStatisticV2Model) GetConfigNameOk() (*string, bool)`

GetConfigNameOk returns a tuple with the ConfigName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigName

`func (o *DetailedStatisticV2Model) SetConfigName(v string)`

SetConfigName sets ConfigName field to given value.


### GetEnvironmentId

`func (o *DetailedStatisticV2Model) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *DetailedStatisticV2Model) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *DetailedStatisticV2Model) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### SetEnvironmentIdNil

`func (o *DetailedStatisticV2Model) SetEnvironmentIdNil(b bool)`

 SetEnvironmentIdNil sets the value for EnvironmentId to be an explicit nil

### UnsetEnvironmentId
`func (o *DetailedStatisticV2Model) UnsetEnvironmentId()`

UnsetEnvironmentId ensures that no value is present for EnvironmentId, not even an explicit nil
### GetEnvironmentName

`func (o *DetailedStatisticV2Model) GetEnvironmentName() string`

GetEnvironmentName returns the EnvironmentName field if non-nil, zero value otherwise.

### GetEnvironmentNameOk

`func (o *DetailedStatisticV2Model) GetEnvironmentNameOk() (*string, bool)`

GetEnvironmentNameOk returns a tuple with the EnvironmentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentName

`func (o *DetailedStatisticV2Model) SetEnvironmentName(v string)`

SetEnvironmentName sets EnvironmentName field to given value.


### GetSdk

`func (o *DetailedStatisticV2Model) GetSdk() string`

GetSdk returns the Sdk field if non-nil, zero value otherwise.

### GetSdkOk

`func (o *DetailedStatisticV2Model) GetSdkOk() (*string, bool)`

GetSdkOk returns a tuple with the Sdk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdk

`func (o *DetailedStatisticV2Model) SetSdk(v string)`

SetSdk sets Sdk field to given value.


### GetSdkKey

`func (o *DetailedStatisticV2Model) GetSdkKey() string`

GetSdkKey returns the SdkKey field if non-nil, zero value otherwise.

### GetSdkKeyOk

`func (o *DetailedStatisticV2Model) GetSdkKeyOk() (*string, bool)`

GetSdkKeyOk returns a tuple with the SdkKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdkKey

`func (o *DetailedStatisticV2Model) SetSdkKey(v string)`

SetSdkKey sets SdkKey field to given value.


### GetRequestCount

`func (o *DetailedStatisticV2Model) GetRequestCount() int64`

GetRequestCount returns the RequestCount field if non-nil, zero value otherwise.

### GetRequestCountOk

`func (o *DetailedStatisticV2Model) GetRequestCountOk() (*int64, bool)`

GetRequestCountOk returns a tuple with the RequestCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCount

`func (o *DetailedStatisticV2Model) SetRequestCount(v int64)`

SetRequestCount sets RequestCount field to given value.


### GetResponseKiloBytes

`func (o *DetailedStatisticV2Model) GetResponseKiloBytes() float64`

GetResponseKiloBytes returns the ResponseKiloBytes field if non-nil, zero value otherwise.

### GetResponseKiloBytesOk

`func (o *DetailedStatisticV2Model) GetResponseKiloBytesOk() (*float64, bool)`

GetResponseKiloBytesOk returns a tuple with the ResponseKiloBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseKiloBytes

`func (o *DetailedStatisticV2Model) SetResponseKiloBytes(v float64)`

SetResponseKiloBytes sets ResponseKiloBytes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


