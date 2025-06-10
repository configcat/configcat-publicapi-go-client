# CodeReferenceModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Branch** | **NullableString** | The source control branch on where the scan was performed. (Source of the branch selector on the ConfigCat Dashboard) | 
**References** | [**[]ReferenceLinesModel**](ReferenceLinesModel.md) | The actual references to the given Feature Flag or Setting. | 
**CommitUrl** | **NullableString** | The related commit&#39;s URL. | 
**CommitHash** | **NullableString** | The related commit&#39;s hash. | 
**SyncedAt** | **time.Time** | The date and time when the reference report was uploaded. | 
**Repository** | **NullableString** | The source control repository that contains the scanned code. | 
**CodeReferenceId** | **string** | The identifier of the reference report. | 
**Uploader** | **NullableString** | The code reference scanning tool&#39;s name. | 

## Methods

### NewCodeReferenceModel

`func NewCodeReferenceModel(branch NullableString, references []ReferenceLinesModel, commitUrl NullableString, commitHash NullableString, syncedAt time.Time, repository NullableString, codeReferenceId string, uploader NullableString, ) *CodeReferenceModel`

NewCodeReferenceModel instantiates a new CodeReferenceModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCodeReferenceModelWithDefaults

`func NewCodeReferenceModelWithDefaults() *CodeReferenceModel`

NewCodeReferenceModelWithDefaults instantiates a new CodeReferenceModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranch

`func (o *CodeReferenceModel) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *CodeReferenceModel) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *CodeReferenceModel) SetBranch(v string)`

SetBranch sets Branch field to given value.


### SetBranchNil

`func (o *CodeReferenceModel) SetBranchNil(b bool)`

 SetBranchNil sets the value for Branch to be an explicit nil

### UnsetBranch
`func (o *CodeReferenceModel) UnsetBranch()`

UnsetBranch ensures that no value is present for Branch, not even an explicit nil
### GetReferences

`func (o *CodeReferenceModel) GetReferences() []ReferenceLinesModel`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *CodeReferenceModel) GetReferencesOk() (*[]ReferenceLinesModel, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *CodeReferenceModel) SetReferences(v []ReferenceLinesModel)`

SetReferences sets References field to given value.


### SetReferencesNil

`func (o *CodeReferenceModel) SetReferencesNil(b bool)`

 SetReferencesNil sets the value for References to be an explicit nil

### UnsetReferences
`func (o *CodeReferenceModel) UnsetReferences()`

UnsetReferences ensures that no value is present for References, not even an explicit nil
### GetCommitUrl

`func (o *CodeReferenceModel) GetCommitUrl() string`

GetCommitUrl returns the CommitUrl field if non-nil, zero value otherwise.

### GetCommitUrlOk

`func (o *CodeReferenceModel) GetCommitUrlOk() (*string, bool)`

GetCommitUrlOk returns a tuple with the CommitUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitUrl

`func (o *CodeReferenceModel) SetCommitUrl(v string)`

SetCommitUrl sets CommitUrl field to given value.


### SetCommitUrlNil

`func (o *CodeReferenceModel) SetCommitUrlNil(b bool)`

 SetCommitUrlNil sets the value for CommitUrl to be an explicit nil

### UnsetCommitUrl
`func (o *CodeReferenceModel) UnsetCommitUrl()`

UnsetCommitUrl ensures that no value is present for CommitUrl, not even an explicit nil
### GetCommitHash

`func (o *CodeReferenceModel) GetCommitHash() string`

GetCommitHash returns the CommitHash field if non-nil, zero value otherwise.

### GetCommitHashOk

`func (o *CodeReferenceModel) GetCommitHashOk() (*string, bool)`

GetCommitHashOk returns a tuple with the CommitHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitHash

`func (o *CodeReferenceModel) SetCommitHash(v string)`

SetCommitHash sets CommitHash field to given value.


### SetCommitHashNil

`func (o *CodeReferenceModel) SetCommitHashNil(b bool)`

 SetCommitHashNil sets the value for CommitHash to be an explicit nil

### UnsetCommitHash
`func (o *CodeReferenceModel) UnsetCommitHash()`

UnsetCommitHash ensures that no value is present for CommitHash, not even an explicit nil
### GetSyncedAt

`func (o *CodeReferenceModel) GetSyncedAt() time.Time`

GetSyncedAt returns the SyncedAt field if non-nil, zero value otherwise.

### GetSyncedAtOk

`func (o *CodeReferenceModel) GetSyncedAtOk() (*time.Time, bool)`

GetSyncedAtOk returns a tuple with the SyncedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncedAt

`func (o *CodeReferenceModel) SetSyncedAt(v time.Time)`

SetSyncedAt sets SyncedAt field to given value.


### GetRepository

`func (o *CodeReferenceModel) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *CodeReferenceModel) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *CodeReferenceModel) SetRepository(v string)`

SetRepository sets Repository field to given value.


### SetRepositoryNil

`func (o *CodeReferenceModel) SetRepositoryNil(b bool)`

 SetRepositoryNil sets the value for Repository to be an explicit nil

### UnsetRepository
`func (o *CodeReferenceModel) UnsetRepository()`

UnsetRepository ensures that no value is present for Repository, not even an explicit nil
### GetCodeReferenceId

`func (o *CodeReferenceModel) GetCodeReferenceId() string`

GetCodeReferenceId returns the CodeReferenceId field if non-nil, zero value otherwise.

### GetCodeReferenceIdOk

`func (o *CodeReferenceModel) GetCodeReferenceIdOk() (*string, bool)`

GetCodeReferenceIdOk returns a tuple with the CodeReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeReferenceId

`func (o *CodeReferenceModel) SetCodeReferenceId(v string)`

SetCodeReferenceId sets CodeReferenceId field to given value.


### GetUploader

`func (o *CodeReferenceModel) GetUploader() string`

GetUploader returns the Uploader field if non-nil, zero value otherwise.

### GetUploaderOk

`func (o *CodeReferenceModel) GetUploaderOk() (*string, bool)`

GetUploaderOk returns a tuple with the Uploader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploader

`func (o *CodeReferenceModel) SetUploader(v string)`

SetUploader sets Uploader field to given value.


### SetUploaderNil

`func (o *CodeReferenceModel) SetUploaderNil(b bool)`

 SetUploaderNil sets the value for Uploader to be an explicit nil

### UnsetUploader
`func (o *CodeReferenceModel) UnsetUploader()`

UnsetUploader ensures that no value is present for Uploader, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


