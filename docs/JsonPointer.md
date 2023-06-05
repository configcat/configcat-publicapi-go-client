# JsonPointer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | Pointer to **NullableString** |  | [optional] [readonly] 
**Segments** | Pointer to [**[]PointerSegment**](PointerSegment.md) |  | [optional] [readonly] 
**IsUriEncoded** | Pointer to **bool** |  | [optional] [readonly] 
**Kind** | Pointer to [**JsonPointerKind**](JsonPointerKind.md) |  | [optional] 

## Methods

### NewJsonPointer

`func NewJsonPointer() *JsonPointer`

NewJsonPointer instantiates a new JsonPointer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJsonPointerWithDefaults

`func NewJsonPointerWithDefaults() *JsonPointer`

NewJsonPointerWithDefaults instantiates a new JsonPointer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *JsonPointer) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *JsonPointer) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *JsonPointer) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *JsonPointer) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *JsonPointer) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *JsonPointer) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetSegments

`func (o *JsonPointer) GetSegments() []PointerSegment`

GetSegments returns the Segments field if non-nil, zero value otherwise.

### GetSegmentsOk

`func (o *JsonPointer) GetSegmentsOk() (*[]PointerSegment, bool)`

GetSegmentsOk returns a tuple with the Segments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegments

`func (o *JsonPointer) SetSegments(v []PointerSegment)`

SetSegments sets Segments field to given value.

### HasSegments

`func (o *JsonPointer) HasSegments() bool`

HasSegments returns a boolean if a field has been set.

### SetSegmentsNil

`func (o *JsonPointer) SetSegmentsNil(b bool)`

 SetSegmentsNil sets the value for Segments to be an explicit nil

### UnsetSegments
`func (o *JsonPointer) UnsetSegments()`

UnsetSegments ensures that no value is present for Segments, not even an explicit nil
### GetIsUriEncoded

`func (o *JsonPointer) GetIsUriEncoded() bool`

GetIsUriEncoded returns the IsUriEncoded field if non-nil, zero value otherwise.

### GetIsUriEncodedOk

`func (o *JsonPointer) GetIsUriEncodedOk() (*bool, bool)`

GetIsUriEncodedOk returns a tuple with the IsUriEncoded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUriEncoded

`func (o *JsonPointer) SetIsUriEncoded(v bool)`

SetIsUriEncoded sets IsUriEncoded field to given value.

### HasIsUriEncoded

`func (o *JsonPointer) HasIsUriEncoded() bool`

HasIsUriEncoded returns a boolean if a field has been set.

### GetKind

`func (o *JsonPointer) GetKind() JsonPointerKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *JsonPointer) GetKindOk() (*JsonPointerKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *JsonPointer) SetKind(v JsonPointerKind)`

SetKind sets Kind field to given value.

### HasKind

`func (o *JsonPointer) HasKind() bool`

HasKind returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


