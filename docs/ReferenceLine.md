# ReferenceLine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LineText** | Pointer to **NullableString** | The content of the reference line. | [optional] 
**LineNumber** | **int32** | The line number of the reference line. | 

## Methods

### NewReferenceLine

`func NewReferenceLine(lineNumber int32, ) *ReferenceLine`

NewReferenceLine instantiates a new ReferenceLine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReferenceLineWithDefaults

`func NewReferenceLineWithDefaults() *ReferenceLine`

NewReferenceLineWithDefaults instantiates a new ReferenceLine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLineText

`func (o *ReferenceLine) GetLineText() string`

GetLineText returns the LineText field if non-nil, zero value otherwise.

### GetLineTextOk

`func (o *ReferenceLine) GetLineTextOk() (*string, bool)`

GetLineTextOk returns a tuple with the LineText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineText

`func (o *ReferenceLine) SetLineText(v string)`

SetLineText sets LineText field to given value.

### HasLineText

`func (o *ReferenceLine) HasLineText() bool`

HasLineText returns a boolean if a field has been set.

### SetLineTextNil

`func (o *ReferenceLine) SetLineTextNil(b bool)`

 SetLineTextNil sets the value for LineText to be an explicit nil

### UnsetLineText
`func (o *ReferenceLine) UnsetLineText()`

UnsetLineText ensures that no value is present for LineText, not even an explicit nil
### GetLineNumber

`func (o *ReferenceLine) GetLineNumber() int32`

GetLineNumber returns the LineNumber field if non-nil, zero value otherwise.

### GetLineNumberOk

`func (o *ReferenceLine) GetLineNumberOk() (*int32, bool)`

GetLineNumberOk returns a tuple with the LineNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineNumber

`func (o *ReferenceLine) SetLineNumber(v int32)`

SetLineNumber sets LineNumber field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


