# PagingResponseInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentPage** | **int32** |  | 
**TotalPages** | **int32** |  | 
**PageSize** | **int32** |  | 
**TotalCount** | **int32** |  | 
**HasPrevious** | **bool** |  | [readonly] 
**HasNext** | **bool** |  | [readonly] 

## Methods

### NewPagingResponseInfo

`func NewPagingResponseInfo(currentPage int32, totalPages int32, pageSize int32, totalCount int32, hasPrevious bool, hasNext bool, ) *PagingResponseInfo`

NewPagingResponseInfo instantiates a new PagingResponseInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPagingResponseInfoWithDefaults

`func NewPagingResponseInfoWithDefaults() *PagingResponseInfo`

NewPagingResponseInfoWithDefaults instantiates a new PagingResponseInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentPage

`func (o *PagingResponseInfo) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *PagingResponseInfo) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *PagingResponseInfo) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.


### GetTotalPages

`func (o *PagingResponseInfo) GetTotalPages() int32`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *PagingResponseInfo) GetTotalPagesOk() (*int32, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *PagingResponseInfo) SetTotalPages(v int32)`

SetTotalPages sets TotalPages field to given value.


### GetPageSize

`func (o *PagingResponseInfo) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *PagingResponseInfo) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *PagingResponseInfo) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.


### GetTotalCount

`func (o *PagingResponseInfo) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *PagingResponseInfo) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *PagingResponseInfo) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetHasPrevious

`func (o *PagingResponseInfo) GetHasPrevious() bool`

GetHasPrevious returns the HasPrevious field if non-nil, zero value otherwise.

### GetHasPreviousOk

`func (o *PagingResponseInfo) GetHasPreviousOk() (*bool, bool)`

GetHasPreviousOk returns a tuple with the HasPrevious field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPrevious

`func (o *PagingResponseInfo) SetHasPrevious(v bool)`

SetHasPrevious sets HasPrevious field to given value.


### GetHasNext

`func (o *PagingResponseInfo) GetHasNext() bool`

GetHasNext returns the HasNext field if non-nil, zero value otherwise.

### GetHasNextOk

`func (o *PagingResponseInfo) GetHasNextOk() (*bool, bool)`

GetHasNextOk returns a tuple with the HasNext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNext

`func (o *PagingResponseInfo) SetHasNext(v bool)`

SetHasNext sets HasNext field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


