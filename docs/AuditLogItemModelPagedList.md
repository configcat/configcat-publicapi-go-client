# AuditLogItemModelPagedList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Paging** | [**NullablePagingResponseInfo**](PagingResponseInfo.md) |  | 
**Data** | [**[]AuditLogItemModel**](AuditLogItemModel.md) |  | 

## Methods

### NewAuditLogItemModelPagedList

`func NewAuditLogItemModelPagedList(paging NullablePagingResponseInfo, data []AuditLogItemModel, ) *AuditLogItemModelPagedList`

NewAuditLogItemModelPagedList instantiates a new AuditLogItemModelPagedList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditLogItemModelPagedListWithDefaults

`func NewAuditLogItemModelPagedListWithDefaults() *AuditLogItemModelPagedList`

NewAuditLogItemModelPagedListWithDefaults instantiates a new AuditLogItemModelPagedList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaging

`func (o *AuditLogItemModelPagedList) GetPaging() PagingResponseInfo`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *AuditLogItemModelPagedList) GetPagingOk() (*PagingResponseInfo, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *AuditLogItemModelPagedList) SetPaging(v PagingResponseInfo)`

SetPaging sets Paging field to given value.


### SetPagingNil

`func (o *AuditLogItemModelPagedList) SetPagingNil(b bool)`

 SetPagingNil sets the value for Paging to be an explicit nil

### UnsetPaging
`func (o *AuditLogItemModelPagedList) UnsetPaging()`

UnsetPaging ensures that no value is present for Paging, not even an explicit nil
### GetData

`func (o *AuditLogItemModelPagedList) GetData() []AuditLogItemModel`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AuditLogItemModelPagedList) GetDataOk() (*[]AuditLogItemModel, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AuditLogItemModelPagedList) SetData(v []AuditLogItemModel)`

SetData sets Data field to given value.


### SetDataNil

`func (o *AuditLogItemModelPagedList) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *AuditLogItemModelPagedList) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


