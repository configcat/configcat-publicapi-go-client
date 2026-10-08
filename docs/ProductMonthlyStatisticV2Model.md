# ProductMonthlyStatisticV2Model

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductId** | **string** | The identifier of the Product associated with the statistics. | 
**Date** | **time.Time** | The date for which the statistics are reported. | 
**MillionRequestCount** | **float64** | The total request volume in millions for the Product. | 
**ResponseMegaBytes** | **float64** | The total network traffic in megabytes for the Product. | 

## Methods

### NewProductMonthlyStatisticV2Model

`func NewProductMonthlyStatisticV2Model(productId string, date time.Time, millionRequestCount float64, responseMegaBytes float64, ) *ProductMonthlyStatisticV2Model`

NewProductMonthlyStatisticV2Model instantiates a new ProductMonthlyStatisticV2Model object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductMonthlyStatisticV2ModelWithDefaults

`func NewProductMonthlyStatisticV2ModelWithDefaults() *ProductMonthlyStatisticV2Model`

NewProductMonthlyStatisticV2ModelWithDefaults instantiates a new ProductMonthlyStatisticV2Model object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductId

`func (o *ProductMonthlyStatisticV2Model) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *ProductMonthlyStatisticV2Model) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *ProductMonthlyStatisticV2Model) SetProductId(v string)`

SetProductId sets ProductId field to given value.


### GetDate

`func (o *ProductMonthlyStatisticV2Model) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *ProductMonthlyStatisticV2Model) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *ProductMonthlyStatisticV2Model) SetDate(v time.Time)`

SetDate sets Date field to given value.


### GetMillionRequestCount

`func (o *ProductMonthlyStatisticV2Model) GetMillionRequestCount() float64`

GetMillionRequestCount returns the MillionRequestCount field if non-nil, zero value otherwise.

### GetMillionRequestCountOk

`func (o *ProductMonthlyStatisticV2Model) GetMillionRequestCountOk() (*float64, bool)`

GetMillionRequestCountOk returns a tuple with the MillionRequestCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMillionRequestCount

`func (o *ProductMonthlyStatisticV2Model) SetMillionRequestCount(v float64)`

SetMillionRequestCount sets MillionRequestCount field to given value.


### GetResponseMegaBytes

`func (o *ProductMonthlyStatisticV2Model) GetResponseMegaBytes() float64`

GetResponseMegaBytes returns the ResponseMegaBytes field if non-nil, zero value otherwise.

### GetResponseMegaBytesOk

`func (o *ProductMonthlyStatisticV2Model) GetResponseMegaBytesOk() (*float64, bool)`

GetResponseMegaBytesOk returns a tuple with the ResponseMegaBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseMegaBytes

`func (o *ProductMonthlyStatisticV2Model) SetResponseMegaBytes(v float64)`

SetResponseMegaBytes sets ResponseMegaBytes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


