# SegmentConditionModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SegmentId** | **string** | The segment&#39;s identifier. | 
**Comparator** | **string** | The segment comparison operator used during the evaluation process. | 

## Methods

### NewSegmentConditionModel

`func NewSegmentConditionModel(segmentId string, comparator string, ) *SegmentConditionModel`

NewSegmentConditionModel instantiates a new SegmentConditionModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSegmentConditionModelWithDefaults

`func NewSegmentConditionModelWithDefaults() *SegmentConditionModel`

NewSegmentConditionModelWithDefaults instantiates a new SegmentConditionModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSegmentId

`func (o *SegmentConditionModel) GetSegmentId() string`

GetSegmentId returns the SegmentId field if non-nil, zero value otherwise.

### GetSegmentIdOk

`func (o *SegmentConditionModel) GetSegmentIdOk() (*string, bool)`

GetSegmentIdOk returns a tuple with the SegmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentId

`func (o *SegmentConditionModel) SetSegmentId(v string)`

SetSegmentId sets SegmentId field to given value.


### GetComparator

`func (o *SegmentConditionModel) GetComparator() string`

GetComparator returns the Comparator field if non-nil, zero value otherwise.

### GetComparatorOk

`func (o *SegmentConditionModel) GetComparatorOk() (*string, bool)`

GetComparatorOk returns a tuple with the Comparator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparator

`func (o *SegmentConditionModel) SetComparator(v string)`

SetComparator sets Comparator field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


