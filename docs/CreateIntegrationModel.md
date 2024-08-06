# CreateIntegrationModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IntegrationType** | [**IntegrationType**](IntegrationType.md) |  | 
**Name** | **string** | Name of the Integration. | 
**Parameters** | **map[string]string** | Parameters of the Integration. | 
**EnvironmentIds** | **[]string** | List of Environment IDs that are connected with this Integration. If the list is empty, all of the Environments are connected. | 
**ConfigIds** | **[]string** | List of Config IDs that are connected with this Integration. If the list is empty, all of the Configs are connected. | 

## Methods

### NewCreateIntegrationModel

`func NewCreateIntegrationModel(integrationType IntegrationType, name string, parameters map[string]string, environmentIds []string, configIds []string, ) *CreateIntegrationModel`

NewCreateIntegrationModel instantiates a new CreateIntegrationModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateIntegrationModelWithDefaults

`func NewCreateIntegrationModelWithDefaults() *CreateIntegrationModel`

NewCreateIntegrationModelWithDefaults instantiates a new CreateIntegrationModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIntegrationType

`func (o *CreateIntegrationModel) GetIntegrationType() IntegrationType`

GetIntegrationType returns the IntegrationType field if non-nil, zero value otherwise.

### GetIntegrationTypeOk

`func (o *CreateIntegrationModel) GetIntegrationTypeOk() (*IntegrationType, bool)`

GetIntegrationTypeOk returns a tuple with the IntegrationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationType

`func (o *CreateIntegrationModel) SetIntegrationType(v IntegrationType)`

SetIntegrationType sets IntegrationType field to given value.


### GetName

`func (o *CreateIntegrationModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateIntegrationModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateIntegrationModel) SetName(v string)`

SetName sets Name field to given value.


### GetParameters

`func (o *CreateIntegrationModel) GetParameters() map[string]string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *CreateIntegrationModel) GetParametersOk() (*map[string]string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *CreateIntegrationModel) SetParameters(v map[string]string)`

SetParameters sets Parameters field to given value.


### GetEnvironmentIds

`func (o *CreateIntegrationModel) GetEnvironmentIds() []string`

GetEnvironmentIds returns the EnvironmentIds field if non-nil, zero value otherwise.

### GetEnvironmentIdsOk

`func (o *CreateIntegrationModel) GetEnvironmentIdsOk() (*[]string, bool)`

GetEnvironmentIdsOk returns a tuple with the EnvironmentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentIds

`func (o *CreateIntegrationModel) SetEnvironmentIds(v []string)`

SetEnvironmentIds sets EnvironmentIds field to given value.


### GetConfigIds

`func (o *CreateIntegrationModel) GetConfigIds() []string`

GetConfigIds returns the ConfigIds field if non-nil, zero value otherwise.

### GetConfigIdsOk

`func (o *CreateIntegrationModel) GetConfigIdsOk() (*[]string, bool)`

GetConfigIdsOk returns a tuple with the ConfigIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigIds

`func (o *CreateIntegrationModel) SetConfigIds(v []string)`

SetConfigIds sets ConfigIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


