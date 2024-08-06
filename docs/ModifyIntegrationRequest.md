# ModifyIntegrationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the Integration. | 
**Parameters** | **map[string]string** | Parameters of the Integration. | 
**EnvironmentIds** | **[]string** | List of Environment IDs that are connected with this Integration. If the list is empty, all of the Environments are connected. | 
**ConfigIds** | **[]string** | List of Config IDs that are connected with this Integration. If the list is empty, all of the Configs are connected. | 

## Methods

### NewModifyIntegrationRequest

`func NewModifyIntegrationRequest(name string, parameters map[string]string, environmentIds []string, configIds []string, ) *ModifyIntegrationRequest`

NewModifyIntegrationRequest instantiates a new ModifyIntegrationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyIntegrationRequestWithDefaults

`func NewModifyIntegrationRequestWithDefaults() *ModifyIntegrationRequest`

NewModifyIntegrationRequestWithDefaults instantiates a new ModifyIntegrationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ModifyIntegrationRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModifyIntegrationRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModifyIntegrationRequest) SetName(v string)`

SetName sets Name field to given value.


### GetParameters

`func (o *ModifyIntegrationRequest) GetParameters() map[string]string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *ModifyIntegrationRequest) GetParametersOk() (*map[string]string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *ModifyIntegrationRequest) SetParameters(v map[string]string)`

SetParameters sets Parameters field to given value.


### GetEnvironmentIds

`func (o *ModifyIntegrationRequest) GetEnvironmentIds() []string`

GetEnvironmentIds returns the EnvironmentIds field if non-nil, zero value otherwise.

### GetEnvironmentIdsOk

`func (o *ModifyIntegrationRequest) GetEnvironmentIdsOk() (*[]string, bool)`

GetEnvironmentIdsOk returns a tuple with the EnvironmentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentIds

`func (o *ModifyIntegrationRequest) SetEnvironmentIds(v []string)`

SetEnvironmentIds sets EnvironmentIds field to given value.


### GetConfigIds

`func (o *ModifyIntegrationRequest) GetConfigIds() []string`

GetConfigIds returns the ConfigIds field if non-nil, zero value otherwise.

### GetConfigIdsOk

`func (o *ModifyIntegrationRequest) GetConfigIdsOk() (*[]string, bool)`

GetConfigIdsOk returns a tuple with the ConfigIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigIds

`func (o *ModifyIntegrationRequest) SetConfigIds(v []string)`

SetConfigIds sets ConfigIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


