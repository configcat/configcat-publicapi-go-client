# UpdateSegmentConditionModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SegmentId** | **string** | The segment&#39;s identifier. | 
**Comparator** | [**SegmentComparator**](SegmentComparator.md) |  | 

## Methods

### NewUpdateSegmentConditionModel

`func NewUpdateSegmentConditionModel(segmentId string, comparator SegmentComparator, ) *UpdateSegmentConditionModel`

NewUpdateSegmentConditionModel instantiates a new UpdateSegmentConditionModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSegmentConditionModelWithDefaults

`func NewUpdateSegmentConditionModelWithDefaults() *UpdateSegmentConditionModel`

NewUpdateSegmentConditionModelWithDefaults instantiates a new UpdateSegmentConditionModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSegmentId

`func (o *UpdateSegmentConditionModel) GetSegmentId() string`

GetSegmentId returns the SegmentId field if non-nil, zero value otherwise.

### GetSegmentIdOk

`func (o *UpdateSegmentConditionModel) GetSegmentIdOk() (*string, bool)`

GetSegmentIdOk returns a tuple with the SegmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentId

`func (o *UpdateSegmentConditionModel) SetSegmentId(v string)`

SetSegmentId sets SegmentId field to given value.


### GetComparator

`func (o *UpdateSegmentConditionModel) GetComparator() SegmentComparator`

GetComparator returns the Comparator field if non-nil, zero value otherwise.

### GetComparatorOk

`func (o *UpdateSegmentConditionModel) GetComparatorOk() (*SegmentComparator, bool)`

GetComparatorOk returns a tuple with the Comparator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparator

`func (o *UpdateSegmentConditionModel) SetComparator(v SegmentComparator)`

SetComparator sets Comparator field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


