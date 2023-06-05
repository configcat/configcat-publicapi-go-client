# DeleteRepositoryReportsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigId** | **string** | The Config&#39;s identifier from where the reports should be deleted. | 
**Repository** | **string** | The source control repository which&#39;s reports should be deleted. | 
**Branch** | Pointer to **NullableString** | If it&#39;s set, only this branch&#39;s reports belonging to the given repository will be deleted. | [optional] 
**SettingId** | Pointer to **NullableInt32** | If it&#39;s set, only this setting&#39;s reports belonging to the given repository will be deleted. | [optional] 

## Methods

### NewDeleteRepositoryReportsRequest

`func NewDeleteRepositoryReportsRequest(configId string, repository string, ) *DeleteRepositoryReportsRequest`

NewDeleteRepositoryReportsRequest instantiates a new DeleteRepositoryReportsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteRepositoryReportsRequestWithDefaults

`func NewDeleteRepositoryReportsRequestWithDefaults() *DeleteRepositoryReportsRequest`

NewDeleteRepositoryReportsRequestWithDefaults instantiates a new DeleteRepositoryReportsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigId

`func (o *DeleteRepositoryReportsRequest) GetConfigId() string`

GetConfigId returns the ConfigId field if non-nil, zero value otherwise.

### GetConfigIdOk

`func (o *DeleteRepositoryReportsRequest) GetConfigIdOk() (*string, bool)`

GetConfigIdOk returns a tuple with the ConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigId

`func (o *DeleteRepositoryReportsRequest) SetConfigId(v string)`

SetConfigId sets ConfigId field to given value.


### GetRepository

`func (o *DeleteRepositoryReportsRequest) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *DeleteRepositoryReportsRequest) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *DeleteRepositoryReportsRequest) SetRepository(v string)`

SetRepository sets Repository field to given value.


### GetBranch

`func (o *DeleteRepositoryReportsRequest) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *DeleteRepositoryReportsRequest) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *DeleteRepositoryReportsRequest) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *DeleteRepositoryReportsRequest) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### SetBranchNil

`func (o *DeleteRepositoryReportsRequest) SetBranchNil(b bool)`

 SetBranchNil sets the value for Branch to be an explicit nil

### UnsetBranch
`func (o *DeleteRepositoryReportsRequest) UnsetBranch()`

UnsetBranch ensures that no value is present for Branch, not even an explicit nil
### GetSettingId

`func (o *DeleteRepositoryReportsRequest) GetSettingId() int32`

GetSettingId returns the SettingId field if non-nil, zero value otherwise.

### GetSettingIdOk

`func (o *DeleteRepositoryReportsRequest) GetSettingIdOk() (*int32, bool)`

GetSettingIdOk returns a tuple with the SettingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingId

`func (o *DeleteRepositoryReportsRequest) SetSettingId(v int32)`

SetSettingId sets SettingId field to given value.

### HasSettingId

`func (o *DeleteRepositoryReportsRequest) HasSettingId() bool`

HasSettingId returns a boolean if a field has been set.

### SetSettingIdNil

`func (o *DeleteRepositoryReportsRequest) SetSettingIdNil(b bool)`

 SetSettingIdNil sets the value for SettingId to be an explicit nil

### UnsetSettingId
`func (o *DeleteRepositoryReportsRequest) UnsetSettingId()`

UnsetSettingId ensures that no value is present for SettingId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


