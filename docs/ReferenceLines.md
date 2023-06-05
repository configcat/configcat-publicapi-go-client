# ReferenceLines

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**File** | **string** | The file&#39;s name in where the code reference has been found. (Appears on the ConfigCat Dashboard) | 
**FileUrl** | Pointer to **NullableString** | The file&#39;s url. (Used to point to the file on the repository&#39;s website) | [optional] 
**PreLines** | Pointer to [**[]ReferenceLine**](ReferenceLine.md) | The lines before the actual reference line. | [optional] 
**PostLines** | Pointer to [**[]ReferenceLine**](ReferenceLine.md) | The lines after the actual reference line. | [optional] 
**ReferenceLine** | [**ReferenceLine**](ReferenceLine.md) |  | 

## Methods

### NewReferenceLines

`func NewReferenceLines(file string, referenceLine ReferenceLine, ) *ReferenceLines`

NewReferenceLines instantiates a new ReferenceLines object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReferenceLinesWithDefaults

`func NewReferenceLinesWithDefaults() *ReferenceLines`

NewReferenceLinesWithDefaults instantiates a new ReferenceLines object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFile

`func (o *ReferenceLines) GetFile() string`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *ReferenceLines) GetFileOk() (*string, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *ReferenceLines) SetFile(v string)`

SetFile sets File field to given value.


### GetFileUrl

`func (o *ReferenceLines) GetFileUrl() string`

GetFileUrl returns the FileUrl field if non-nil, zero value otherwise.

### GetFileUrlOk

`func (o *ReferenceLines) GetFileUrlOk() (*string, bool)`

GetFileUrlOk returns a tuple with the FileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUrl

`func (o *ReferenceLines) SetFileUrl(v string)`

SetFileUrl sets FileUrl field to given value.

### HasFileUrl

`func (o *ReferenceLines) HasFileUrl() bool`

HasFileUrl returns a boolean if a field has been set.

### SetFileUrlNil

`func (o *ReferenceLines) SetFileUrlNil(b bool)`

 SetFileUrlNil sets the value for FileUrl to be an explicit nil

### UnsetFileUrl
`func (o *ReferenceLines) UnsetFileUrl()`

UnsetFileUrl ensures that no value is present for FileUrl, not even an explicit nil
### GetPreLines

`func (o *ReferenceLines) GetPreLines() []ReferenceLine`

GetPreLines returns the PreLines field if non-nil, zero value otherwise.

### GetPreLinesOk

`func (o *ReferenceLines) GetPreLinesOk() (*[]ReferenceLine, bool)`

GetPreLinesOk returns a tuple with the PreLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreLines

`func (o *ReferenceLines) SetPreLines(v []ReferenceLine)`

SetPreLines sets PreLines field to given value.

### HasPreLines

`func (o *ReferenceLines) HasPreLines() bool`

HasPreLines returns a boolean if a field has been set.

### SetPreLinesNil

`func (o *ReferenceLines) SetPreLinesNil(b bool)`

 SetPreLinesNil sets the value for PreLines to be an explicit nil

### UnsetPreLines
`func (o *ReferenceLines) UnsetPreLines()`

UnsetPreLines ensures that no value is present for PreLines, not even an explicit nil
### GetPostLines

`func (o *ReferenceLines) GetPostLines() []ReferenceLine`

GetPostLines returns the PostLines field if non-nil, zero value otherwise.

### GetPostLinesOk

`func (o *ReferenceLines) GetPostLinesOk() (*[]ReferenceLine, bool)`

GetPostLinesOk returns a tuple with the PostLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostLines

`func (o *ReferenceLines) SetPostLines(v []ReferenceLine)`

SetPostLines sets PostLines field to given value.

### HasPostLines

`func (o *ReferenceLines) HasPostLines() bool`

HasPostLines returns a boolean if a field has been set.

### SetPostLinesNil

`func (o *ReferenceLines) SetPostLinesNil(b bool)`

 SetPostLinesNil sets the value for PostLines to be an explicit nil

### UnsetPostLines
`func (o *ReferenceLines) UnsetPostLines()`

UnsetPostLines ensures that no value is present for PostLines, not even an explicit nil
### GetReferenceLine

`func (o *ReferenceLines) GetReferenceLine() ReferenceLine`

GetReferenceLine returns the ReferenceLine field if non-nil, zero value otherwise.

### GetReferenceLineOk

`func (o *ReferenceLines) GetReferenceLineOk() (*ReferenceLine, bool)`

GetReferenceLineOk returns a tuple with the ReferenceLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceLine

`func (o *ReferenceLines) SetReferenceLine(v ReferenceLine)`

SetReferenceLine sets ReferenceLine field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


