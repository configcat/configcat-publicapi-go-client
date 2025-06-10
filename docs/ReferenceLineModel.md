# ReferenceLineModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LineText** | **NullableString** | The content of the reference line. | 
**LineNumber** | **int32** | The line number. | 

## Methods

### NewReferenceLineModel

`func NewReferenceLineModel(lineText NullableString, lineNumber int32, ) *ReferenceLineModel`

NewReferenceLineModel instantiates a new ReferenceLineModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReferenceLineModelWithDefaults

`func NewReferenceLineModelWithDefaults() *ReferenceLineModel`

NewReferenceLineModelWithDefaults instantiates a new ReferenceLineModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLineText

`func (o *ReferenceLineModel) GetLineText() string`

GetLineText returns the LineText field if non-nil, zero value otherwise.

### GetLineTextOk

`func (o *ReferenceLineModel) GetLineTextOk() (*string, bool)`

GetLineTextOk returns a tuple with the LineText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineText

`func (o *ReferenceLineModel) SetLineText(v string)`

SetLineText sets LineText field to given value.


### SetLineTextNil

`func (o *ReferenceLineModel) SetLineTextNil(b bool)`

 SetLineTextNil sets the value for LineText to be an explicit nil

### UnsetLineText
`func (o *ReferenceLineModel) UnsetLineText()`

UnsetLineText ensures that no value is present for LineText, not even an explicit nil
### GetLineNumber

`func (o *ReferenceLineModel) GetLineNumber() int32`

GetLineNumber returns the LineNumber field if non-nil, zero value otherwise.

### GetLineNumberOk

`func (o *ReferenceLineModel) GetLineNumberOk() (*int32, bool)`

GetLineNumberOk returns a tuple with the LineNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineNumber

`func (o *ReferenceLineModel) SetLineNumber(v int32)`

SetLineNumber sets LineNumber field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


