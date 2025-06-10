# ReferenceLinesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**File** | **string** | The file&#39;s name in where the code reference has been found. (Appears on the ConfigCat Dashboard) | 
**FileUrl** | Pointer to **NullableString** | The file&#39;s url. (Used to point to the file on the repository&#39;s website) | [optional] 
**PreLines** | Pointer to [**[]ReferenceLineRequest**](ReferenceLineRequest.md) | The lines before the actual reference line. | [optional] 
**PostLines** | Pointer to [**[]ReferenceLineRequest**](ReferenceLineRequest.md) | The lines after the actual reference line. | [optional] 
**ReferenceLine** | [**ReferenceLineRequest**](ReferenceLineRequest.md) |  | 

## Methods

### NewReferenceLinesRequest

`func NewReferenceLinesRequest(file string, referenceLine ReferenceLineRequest, ) *ReferenceLinesRequest`

NewReferenceLinesRequest instantiates a new ReferenceLinesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReferenceLinesRequestWithDefaults

`func NewReferenceLinesRequestWithDefaults() *ReferenceLinesRequest`

NewReferenceLinesRequestWithDefaults instantiates a new ReferenceLinesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFile

`func (o *ReferenceLinesRequest) GetFile() string`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *ReferenceLinesRequest) GetFileOk() (*string, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *ReferenceLinesRequest) SetFile(v string)`

SetFile sets File field to given value.


### GetFileUrl

`func (o *ReferenceLinesRequest) GetFileUrl() string`

GetFileUrl returns the FileUrl field if non-nil, zero value otherwise.

### GetFileUrlOk

`func (o *ReferenceLinesRequest) GetFileUrlOk() (*string, bool)`

GetFileUrlOk returns a tuple with the FileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUrl

`func (o *ReferenceLinesRequest) SetFileUrl(v string)`

SetFileUrl sets FileUrl field to given value.

### HasFileUrl

`func (o *ReferenceLinesRequest) HasFileUrl() bool`

HasFileUrl returns a boolean if a field has been set.

### SetFileUrlNil

`func (o *ReferenceLinesRequest) SetFileUrlNil(b bool)`

 SetFileUrlNil sets the value for FileUrl to be an explicit nil

### UnsetFileUrl
`func (o *ReferenceLinesRequest) UnsetFileUrl()`

UnsetFileUrl ensures that no value is present for FileUrl, not even an explicit nil
### GetPreLines

`func (o *ReferenceLinesRequest) GetPreLines() []ReferenceLineRequest`

GetPreLines returns the PreLines field if non-nil, zero value otherwise.

### GetPreLinesOk

`func (o *ReferenceLinesRequest) GetPreLinesOk() (*[]ReferenceLineRequest, bool)`

GetPreLinesOk returns a tuple with the PreLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreLines

`func (o *ReferenceLinesRequest) SetPreLines(v []ReferenceLineRequest)`

SetPreLines sets PreLines field to given value.

### HasPreLines

`func (o *ReferenceLinesRequest) HasPreLines() bool`

HasPreLines returns a boolean if a field has been set.

### GetPostLines

`func (o *ReferenceLinesRequest) GetPostLines() []ReferenceLineRequest`

GetPostLines returns the PostLines field if non-nil, zero value otherwise.

### GetPostLinesOk

`func (o *ReferenceLinesRequest) GetPostLinesOk() (*[]ReferenceLineRequest, bool)`

GetPostLinesOk returns a tuple with the PostLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostLines

`func (o *ReferenceLinesRequest) SetPostLines(v []ReferenceLineRequest)`

SetPostLines sets PostLines field to given value.

### HasPostLines

`func (o *ReferenceLinesRequest) HasPostLines() bool`

HasPostLines returns a boolean if a field has been set.

### GetReferenceLine

`func (o *ReferenceLinesRequest) GetReferenceLine() ReferenceLineRequest`

GetReferenceLine returns the ReferenceLine field if non-nil, zero value otherwise.

### GetReferenceLineOk

`func (o *ReferenceLinesRequest) GetReferenceLineOk() (*ReferenceLineRequest, bool)`

GetReferenceLineOk returns a tuple with the ReferenceLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceLine

`func (o *ReferenceLinesRequest) SetReferenceLine(v ReferenceLineRequest)`

SetReferenceLine sets ReferenceLine field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


