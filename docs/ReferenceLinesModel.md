# ReferenceLinesModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**File** | **string** | The file&#39;s name in where the code reference has been found. (Appears on the ConfigCat Dashboard) | 
**FileUrl** | **NullableString** | The file&#39;s url. (Used to point to the file on the repository&#39;s website) | 
**PreLines** | [**[]ReferenceLineModel**](ReferenceLineModel.md) | The lines before the actual reference line. | 
**PostLines** | [**[]ReferenceLineModel**](ReferenceLineModel.md) | The lines after the actual reference line. | 
**ReferenceLine** | [**ReferenceLineModel**](ReferenceLineModel.md) |  | 

## Methods

### NewReferenceLinesModel

`func NewReferenceLinesModel(file string, fileUrl NullableString, preLines []ReferenceLineModel, postLines []ReferenceLineModel, referenceLine ReferenceLineModel, ) *ReferenceLinesModel`

NewReferenceLinesModel instantiates a new ReferenceLinesModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReferenceLinesModelWithDefaults

`func NewReferenceLinesModelWithDefaults() *ReferenceLinesModel`

NewReferenceLinesModelWithDefaults instantiates a new ReferenceLinesModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFile

`func (o *ReferenceLinesModel) GetFile() string`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *ReferenceLinesModel) GetFileOk() (*string, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *ReferenceLinesModel) SetFile(v string)`

SetFile sets File field to given value.


### GetFileUrl

`func (o *ReferenceLinesModel) GetFileUrl() string`

GetFileUrl returns the FileUrl field if non-nil, zero value otherwise.

### GetFileUrlOk

`func (o *ReferenceLinesModel) GetFileUrlOk() (*string, bool)`

GetFileUrlOk returns a tuple with the FileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUrl

`func (o *ReferenceLinesModel) SetFileUrl(v string)`

SetFileUrl sets FileUrl field to given value.


### SetFileUrlNil

`func (o *ReferenceLinesModel) SetFileUrlNil(b bool)`

 SetFileUrlNil sets the value for FileUrl to be an explicit nil

### UnsetFileUrl
`func (o *ReferenceLinesModel) UnsetFileUrl()`

UnsetFileUrl ensures that no value is present for FileUrl, not even an explicit nil
### GetPreLines

`func (o *ReferenceLinesModel) GetPreLines() []ReferenceLineModel`

GetPreLines returns the PreLines field if non-nil, zero value otherwise.

### GetPreLinesOk

`func (o *ReferenceLinesModel) GetPreLinesOk() (*[]ReferenceLineModel, bool)`

GetPreLinesOk returns a tuple with the PreLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreLines

`func (o *ReferenceLinesModel) SetPreLines(v []ReferenceLineModel)`

SetPreLines sets PreLines field to given value.


### GetPostLines

`func (o *ReferenceLinesModel) GetPostLines() []ReferenceLineModel`

GetPostLines returns the PostLines field if non-nil, zero value otherwise.

### GetPostLinesOk

`func (o *ReferenceLinesModel) GetPostLinesOk() (*[]ReferenceLineModel, bool)`

GetPostLinesOk returns a tuple with the PostLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostLines

`func (o *ReferenceLinesModel) SetPostLines(v []ReferenceLineModel)`

SetPostLines sets PostLines field to given value.


### GetReferenceLine

`func (o *ReferenceLinesModel) GetReferenceLine() ReferenceLineModel`

GetReferenceLine returns the ReferenceLine field if non-nil, zero value otherwise.

### GetReferenceLineOk

`func (o *ReferenceLinesModel) GetReferenceLineOk() (*ReferenceLineModel, bool)`

GetReferenceLineOk returns a tuple with the ReferenceLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceLine

`func (o *ReferenceLinesModel) SetReferenceLine(v ReferenceLineModel)`

SetReferenceLine sets ReferenceLine field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


