# OrganizationMonthlyStatisticV2Model

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | **time.Time** | The date for which the aggregate statistics are reported. | 
**MillionRequestCount** | **float64** | The total request volume in millions for the period. | 
**ResponseMegaBytes** | **float64** | The total network traffic in megabytes for the period. | 
**OverLimit** | **bool** | Indicates whether the request quota was exceeded. | 
**OverNetworkTrafficLimit** | **bool** | Indicates whether the network traffic quota was exceeded. | 
**MillionRequestLimitPerMonth** | **int32** | The monthly request quota limit in millions. | 
**NetworkTrafficGigaByteLimitPerMonth** | **int32** | The monthly network traffic quota limit in gigabytes. | 
**PublicApiCallCount** | **int64** | The number of Public API calls recorded for the period. | 

## Methods

### NewOrganizationMonthlyStatisticV2Model

`func NewOrganizationMonthlyStatisticV2Model(date time.Time, millionRequestCount float64, responseMegaBytes float64, overLimit bool, overNetworkTrafficLimit bool, millionRequestLimitPerMonth int32, networkTrafficGigaByteLimitPerMonth int32, publicApiCallCount int64, ) *OrganizationMonthlyStatisticV2Model`

NewOrganizationMonthlyStatisticV2Model instantiates a new OrganizationMonthlyStatisticV2Model object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationMonthlyStatisticV2ModelWithDefaults

`func NewOrganizationMonthlyStatisticV2ModelWithDefaults() *OrganizationMonthlyStatisticV2Model`

NewOrganizationMonthlyStatisticV2ModelWithDefaults instantiates a new OrganizationMonthlyStatisticV2Model object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *OrganizationMonthlyStatisticV2Model) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *OrganizationMonthlyStatisticV2Model) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *OrganizationMonthlyStatisticV2Model) SetDate(v time.Time)`

SetDate sets Date field to given value.


### GetMillionRequestCount

`func (o *OrganizationMonthlyStatisticV2Model) GetMillionRequestCount() float64`

GetMillionRequestCount returns the MillionRequestCount field if non-nil, zero value otherwise.

### GetMillionRequestCountOk

`func (o *OrganizationMonthlyStatisticV2Model) GetMillionRequestCountOk() (*float64, bool)`

GetMillionRequestCountOk returns a tuple with the MillionRequestCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMillionRequestCount

`func (o *OrganizationMonthlyStatisticV2Model) SetMillionRequestCount(v float64)`

SetMillionRequestCount sets MillionRequestCount field to given value.


### GetResponseMegaBytes

`func (o *OrganizationMonthlyStatisticV2Model) GetResponseMegaBytes() float64`

GetResponseMegaBytes returns the ResponseMegaBytes field if non-nil, zero value otherwise.

### GetResponseMegaBytesOk

`func (o *OrganizationMonthlyStatisticV2Model) GetResponseMegaBytesOk() (*float64, bool)`

GetResponseMegaBytesOk returns a tuple with the ResponseMegaBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseMegaBytes

`func (o *OrganizationMonthlyStatisticV2Model) SetResponseMegaBytes(v float64)`

SetResponseMegaBytes sets ResponseMegaBytes field to given value.


### GetOverLimit

`func (o *OrganizationMonthlyStatisticV2Model) GetOverLimit() bool`

GetOverLimit returns the OverLimit field if non-nil, zero value otherwise.

### GetOverLimitOk

`func (o *OrganizationMonthlyStatisticV2Model) GetOverLimitOk() (*bool, bool)`

GetOverLimitOk returns a tuple with the OverLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverLimit

`func (o *OrganizationMonthlyStatisticV2Model) SetOverLimit(v bool)`

SetOverLimit sets OverLimit field to given value.


### GetOverNetworkTrafficLimit

`func (o *OrganizationMonthlyStatisticV2Model) GetOverNetworkTrafficLimit() bool`

GetOverNetworkTrafficLimit returns the OverNetworkTrafficLimit field if non-nil, zero value otherwise.

### GetOverNetworkTrafficLimitOk

`func (o *OrganizationMonthlyStatisticV2Model) GetOverNetworkTrafficLimitOk() (*bool, bool)`

GetOverNetworkTrafficLimitOk returns a tuple with the OverNetworkTrafficLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverNetworkTrafficLimit

`func (o *OrganizationMonthlyStatisticV2Model) SetOverNetworkTrafficLimit(v bool)`

SetOverNetworkTrafficLimit sets OverNetworkTrafficLimit field to given value.


### GetMillionRequestLimitPerMonth

`func (o *OrganizationMonthlyStatisticV2Model) GetMillionRequestLimitPerMonth() int32`

GetMillionRequestLimitPerMonth returns the MillionRequestLimitPerMonth field if non-nil, zero value otherwise.

### GetMillionRequestLimitPerMonthOk

`func (o *OrganizationMonthlyStatisticV2Model) GetMillionRequestLimitPerMonthOk() (*int32, bool)`

GetMillionRequestLimitPerMonthOk returns a tuple with the MillionRequestLimitPerMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMillionRequestLimitPerMonth

`func (o *OrganizationMonthlyStatisticV2Model) SetMillionRequestLimitPerMonth(v int32)`

SetMillionRequestLimitPerMonth sets MillionRequestLimitPerMonth field to given value.


### GetNetworkTrafficGigaByteLimitPerMonth

`func (o *OrganizationMonthlyStatisticV2Model) GetNetworkTrafficGigaByteLimitPerMonth() int32`

GetNetworkTrafficGigaByteLimitPerMonth returns the NetworkTrafficGigaByteLimitPerMonth field if non-nil, zero value otherwise.

### GetNetworkTrafficGigaByteLimitPerMonthOk

`func (o *OrganizationMonthlyStatisticV2Model) GetNetworkTrafficGigaByteLimitPerMonthOk() (*int32, bool)`

GetNetworkTrafficGigaByteLimitPerMonthOk returns a tuple with the NetworkTrafficGigaByteLimitPerMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkTrafficGigaByteLimitPerMonth

`func (o *OrganizationMonthlyStatisticV2Model) SetNetworkTrafficGigaByteLimitPerMonth(v int32)`

SetNetworkTrafficGigaByteLimitPerMonth sets NetworkTrafficGigaByteLimitPerMonth field to given value.


### GetPublicApiCallCount

`func (o *OrganizationMonthlyStatisticV2Model) GetPublicApiCallCount() int64`

GetPublicApiCallCount returns the PublicApiCallCount field if non-nil, zero value otherwise.

### GetPublicApiCallCountOk

`func (o *OrganizationMonthlyStatisticV2Model) GetPublicApiCallCountOk() (*int64, bool)`

GetPublicApiCallCountOk returns a tuple with the PublicApiCallCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicApiCallCount

`func (o *OrganizationMonthlyStatisticV2Model) SetPublicApiCallCount(v int64)`

SetPublicApiCallCount sets PublicApiCallCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


