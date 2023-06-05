# CodeReferenceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigId** | **string** | The Config&#39;s identifier the scanning was performed against. | 
**Repository** | **string** | The source control repository that contains the scanned code. (Source of the repository selector on the ConfigCat Dashboard) | 
**Branch** | **string** | The source control branch on where the scan was performed. (Source of the branch selector on the ConfigCat Dashboard) | 
**CommitUrl** | Pointer to **NullableString** | The related commit&#39;s URL. (Appears on the ConfigCat Dashboard) | [optional] 
**CommitHash** | Pointer to **NullableString** | The related commit&#39;s hash. (Appears on the ConfigCat Dashboard) | [optional] 
**Uploader** | Pointer to **NullableString** | The scanning tool&#39;s name. (Appears on the ConfigCat Dashboard) | [optional] 
**ActiveBranches** | Pointer to **[]string** | The currently active branches of the repository. Each previously uploaded report that belongs to a non-reported active branch is being deleted. | [optional] 
**FlagReferences** | Pointer to [**[]FlagReference**](FlagReference.md) | The actual code reference collection. | [optional] 

## Methods

### NewCodeReferenceRequest

`func NewCodeReferenceRequest(configId string, repository string, branch string, ) *CodeReferenceRequest`

NewCodeReferenceRequest instantiates a new CodeReferenceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCodeReferenceRequestWithDefaults

`func NewCodeReferenceRequestWithDefaults() *CodeReferenceRequest`

NewCodeReferenceRequestWithDefaults instantiates a new CodeReferenceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigId

`func (o *CodeReferenceRequest) GetConfigId() string`

GetConfigId returns the ConfigId field if non-nil, zero value otherwise.

### GetConfigIdOk

`func (o *CodeReferenceRequest) GetConfigIdOk() (*string, bool)`

GetConfigIdOk returns a tuple with the ConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigId

`func (o *CodeReferenceRequest) SetConfigId(v string)`

SetConfigId sets ConfigId field to given value.


### GetRepository

`func (o *CodeReferenceRequest) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *CodeReferenceRequest) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *CodeReferenceRequest) SetRepository(v string)`

SetRepository sets Repository field to given value.


### GetBranch

`func (o *CodeReferenceRequest) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *CodeReferenceRequest) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *CodeReferenceRequest) SetBranch(v string)`

SetBranch sets Branch field to given value.


### GetCommitUrl

`func (o *CodeReferenceRequest) GetCommitUrl() string`

GetCommitUrl returns the CommitUrl field if non-nil, zero value otherwise.

### GetCommitUrlOk

`func (o *CodeReferenceRequest) GetCommitUrlOk() (*string, bool)`

GetCommitUrlOk returns a tuple with the CommitUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitUrl

`func (o *CodeReferenceRequest) SetCommitUrl(v string)`

SetCommitUrl sets CommitUrl field to given value.

### HasCommitUrl

`func (o *CodeReferenceRequest) HasCommitUrl() bool`

HasCommitUrl returns a boolean if a field has been set.

### SetCommitUrlNil

`func (o *CodeReferenceRequest) SetCommitUrlNil(b bool)`

 SetCommitUrlNil sets the value for CommitUrl to be an explicit nil

### UnsetCommitUrl
`func (o *CodeReferenceRequest) UnsetCommitUrl()`

UnsetCommitUrl ensures that no value is present for CommitUrl, not even an explicit nil
### GetCommitHash

`func (o *CodeReferenceRequest) GetCommitHash() string`

GetCommitHash returns the CommitHash field if non-nil, zero value otherwise.

### GetCommitHashOk

`func (o *CodeReferenceRequest) GetCommitHashOk() (*string, bool)`

GetCommitHashOk returns a tuple with the CommitHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitHash

`func (o *CodeReferenceRequest) SetCommitHash(v string)`

SetCommitHash sets CommitHash field to given value.

### HasCommitHash

`func (o *CodeReferenceRequest) HasCommitHash() bool`

HasCommitHash returns a boolean if a field has been set.

### SetCommitHashNil

`func (o *CodeReferenceRequest) SetCommitHashNil(b bool)`

 SetCommitHashNil sets the value for CommitHash to be an explicit nil

### UnsetCommitHash
`func (o *CodeReferenceRequest) UnsetCommitHash()`

UnsetCommitHash ensures that no value is present for CommitHash, not even an explicit nil
### GetUploader

`func (o *CodeReferenceRequest) GetUploader() string`

GetUploader returns the Uploader field if non-nil, zero value otherwise.

### GetUploaderOk

`func (o *CodeReferenceRequest) GetUploaderOk() (*string, bool)`

GetUploaderOk returns a tuple with the Uploader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploader

`func (o *CodeReferenceRequest) SetUploader(v string)`

SetUploader sets Uploader field to given value.

### HasUploader

`func (o *CodeReferenceRequest) HasUploader() bool`

HasUploader returns a boolean if a field has been set.

### SetUploaderNil

`func (o *CodeReferenceRequest) SetUploaderNil(b bool)`

 SetUploaderNil sets the value for Uploader to be an explicit nil

### UnsetUploader
`func (o *CodeReferenceRequest) UnsetUploader()`

UnsetUploader ensures that no value is present for Uploader, not even an explicit nil
### GetActiveBranches

`func (o *CodeReferenceRequest) GetActiveBranches() []string`

GetActiveBranches returns the ActiveBranches field if non-nil, zero value otherwise.

### GetActiveBranchesOk

`func (o *CodeReferenceRequest) GetActiveBranchesOk() (*[]string, bool)`

GetActiveBranchesOk returns a tuple with the ActiveBranches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveBranches

`func (o *CodeReferenceRequest) SetActiveBranches(v []string)`

SetActiveBranches sets ActiveBranches field to given value.

### HasActiveBranches

`func (o *CodeReferenceRequest) HasActiveBranches() bool`

HasActiveBranches returns a boolean if a field has been set.

### SetActiveBranchesNil

`func (o *CodeReferenceRequest) SetActiveBranchesNil(b bool)`

 SetActiveBranchesNil sets the value for ActiveBranches to be an explicit nil

### UnsetActiveBranches
`func (o *CodeReferenceRequest) UnsetActiveBranches()`

UnsetActiveBranches ensures that no value is present for ActiveBranches, not even an explicit nil
### GetFlagReferences

`func (o *CodeReferenceRequest) GetFlagReferences() []FlagReference`

GetFlagReferences returns the FlagReferences field if non-nil, zero value otherwise.

### GetFlagReferencesOk

`func (o *CodeReferenceRequest) GetFlagReferencesOk() (*[]FlagReference, bool)`

GetFlagReferencesOk returns a tuple with the FlagReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagReferences

`func (o *CodeReferenceRequest) SetFlagReferences(v []FlagReference)`

SetFlagReferences sets FlagReferences field to given value.

### HasFlagReferences

`func (o *CodeReferenceRequest) HasFlagReferences() bool`

HasFlagReferences returns a boolean if a field has been set.

### SetFlagReferencesNil

`func (o *CodeReferenceRequest) SetFlagReferencesNil(b bool)`

 SetFlagReferencesNil sets the value for FlagReferences to be an explicit nil

### UnsetFlagReferences
`func (o *CodeReferenceRequest) UnsetFlagReferences()`

UnsetFlagReferences ensures that no value is present for FlagReferences, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


