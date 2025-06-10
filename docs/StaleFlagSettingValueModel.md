# StaleFlagSettingValueModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentId** | **string** | Identifier of the Environment. | 
**UpdatedAt** | **NullableTime** | The last updated date and time when the Feature Flag or Setting. | 
**IsStale** | **bool** | Is the feature flag considered stale in the environment. | 

## Methods

### NewStaleFlagSettingValueModel

`func NewStaleFlagSettingValueModel(environmentId string, updatedAt NullableTime, isStale bool, ) *StaleFlagSettingValueModel`

NewStaleFlagSettingValueModel instantiates a new StaleFlagSettingValueModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStaleFlagSettingValueModelWithDefaults

`func NewStaleFlagSettingValueModelWithDefaults() *StaleFlagSettingValueModel`

NewStaleFlagSettingValueModelWithDefaults instantiates a new StaleFlagSettingValueModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentId

`func (o *StaleFlagSettingValueModel) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *StaleFlagSettingValueModel) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *StaleFlagSettingValueModel) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetUpdatedAt

`func (o *StaleFlagSettingValueModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *StaleFlagSettingValueModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *StaleFlagSettingValueModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### SetUpdatedAtNil

`func (o *StaleFlagSettingValueModel) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *StaleFlagSettingValueModel) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetIsStale

`func (o *StaleFlagSettingValueModel) GetIsStale() bool`

GetIsStale returns the IsStale field if non-nil, zero value otherwise.

### GetIsStaleOk

`func (o *StaleFlagSettingValueModel) GetIsStaleOk() (*bool, bool)`

GetIsStaleOk returns a tuple with the IsStale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStale

`func (o *StaleFlagSettingValueModel) SetIsStale(v bool)`

SetIsStale sets IsStale field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


