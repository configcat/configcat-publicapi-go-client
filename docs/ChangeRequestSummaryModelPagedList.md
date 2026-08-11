# ChangeRequestSummaryModelPagedList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Paging** | [**NullablePagingResponseInfo**](PagingResponseInfo.md) |  | 
**Data** | [**[]ChangeRequestSummaryModel**](ChangeRequestSummaryModel.md) |  | 

## Methods

### NewChangeRequestSummaryModelPagedList

`func NewChangeRequestSummaryModelPagedList(paging NullablePagingResponseInfo, data []ChangeRequestSummaryModel, ) *ChangeRequestSummaryModelPagedList`

NewChangeRequestSummaryModelPagedList instantiates a new ChangeRequestSummaryModelPagedList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangeRequestSummaryModelPagedListWithDefaults

`func NewChangeRequestSummaryModelPagedListWithDefaults() *ChangeRequestSummaryModelPagedList`

NewChangeRequestSummaryModelPagedListWithDefaults instantiates a new ChangeRequestSummaryModelPagedList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaging

`func (o *ChangeRequestSummaryModelPagedList) GetPaging() PagingResponseInfo`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *ChangeRequestSummaryModelPagedList) GetPagingOk() (*PagingResponseInfo, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *ChangeRequestSummaryModelPagedList) SetPaging(v PagingResponseInfo)`

SetPaging sets Paging field to given value.


### SetPagingNil

`func (o *ChangeRequestSummaryModelPagedList) SetPagingNil(b bool)`

 SetPagingNil sets the value for Paging to be an explicit nil

### UnsetPaging
`func (o *ChangeRequestSummaryModelPagedList) UnsetPaging()`

UnsetPaging ensures that no value is present for Paging, not even an explicit nil
### GetData

`func (o *ChangeRequestSummaryModelPagedList) GetData() []ChangeRequestSummaryModel`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ChangeRequestSummaryModelPagedList) GetDataOk() (*[]ChangeRequestSummaryModel, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ChangeRequestSummaryModelPagedList) SetData(v []ChangeRequestSummaryModel)`

SetData sets Data field to given value.


### SetDataNil

`func (o *ChangeRequestSummaryModelPagedList) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *ChangeRequestSummaryModelPagedList) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


