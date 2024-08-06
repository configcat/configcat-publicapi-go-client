# IntegrationModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IntegrationId** | Pointer to **string** | Identifier of the Integration. | [optional] 
**Name** | Pointer to **NullableString** | Name of the Integration. | [optional] 
**IntegrationType** | Pointer to [**IntegrationType**](IntegrationType.md) |  | [optional] 
**Parameters** | Pointer to **map[string]string** | Parameters of the Integration. | [optional] 
**IntegrationEnvironmentIds** | Pointer to **[]string** | List of Environment IDs that are connected with this Integration. If the list is empty, all of the Environments are connected. | [optional] 
**IntegrationConfigIds** | Pointer to **[]string** | List of Config IDs that are connected with this Integration. If the list is empty, all of the Configs are connected. | [optional] 

## Methods

### NewIntegrationModel

`func NewIntegrationModel() *IntegrationModel`

NewIntegrationModel instantiates a new IntegrationModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationModelWithDefaults

`func NewIntegrationModelWithDefaults() *IntegrationModel`

NewIntegrationModelWithDefaults instantiates a new IntegrationModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIntegrationId

`func (o *IntegrationModel) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *IntegrationModel) GetIntegrationIdOk() (*string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *IntegrationModel) SetIntegrationId(v string)`

SetIntegrationId sets IntegrationId field to given value.

### HasIntegrationId

`func (o *IntegrationModel) HasIntegrationId() bool`

HasIntegrationId returns a boolean if a field has been set.

### GetName

`func (o *IntegrationModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IntegrationModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IntegrationModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IntegrationModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *IntegrationModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *IntegrationModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetIntegrationType

`func (o *IntegrationModel) GetIntegrationType() IntegrationType`

GetIntegrationType returns the IntegrationType field if non-nil, zero value otherwise.

### GetIntegrationTypeOk

`func (o *IntegrationModel) GetIntegrationTypeOk() (*IntegrationType, bool)`

GetIntegrationTypeOk returns a tuple with the IntegrationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationType

`func (o *IntegrationModel) SetIntegrationType(v IntegrationType)`

SetIntegrationType sets IntegrationType field to given value.

### HasIntegrationType

`func (o *IntegrationModel) HasIntegrationType() bool`

HasIntegrationType returns a boolean if a field has been set.

### GetParameters

`func (o *IntegrationModel) GetParameters() map[string]string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *IntegrationModel) GetParametersOk() (*map[string]string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *IntegrationModel) SetParameters(v map[string]string)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *IntegrationModel) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### SetParametersNil

`func (o *IntegrationModel) SetParametersNil(b bool)`

 SetParametersNil sets the value for Parameters to be an explicit nil

### UnsetParameters
`func (o *IntegrationModel) UnsetParameters()`

UnsetParameters ensures that no value is present for Parameters, not even an explicit nil
### GetIntegrationEnvironmentIds

`func (o *IntegrationModel) GetIntegrationEnvironmentIds() []string`

GetIntegrationEnvironmentIds returns the IntegrationEnvironmentIds field if non-nil, zero value otherwise.

### GetIntegrationEnvironmentIdsOk

`func (o *IntegrationModel) GetIntegrationEnvironmentIdsOk() (*[]string, bool)`

GetIntegrationEnvironmentIdsOk returns a tuple with the IntegrationEnvironmentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationEnvironmentIds

`func (o *IntegrationModel) SetIntegrationEnvironmentIds(v []string)`

SetIntegrationEnvironmentIds sets IntegrationEnvironmentIds field to given value.

### HasIntegrationEnvironmentIds

`func (o *IntegrationModel) HasIntegrationEnvironmentIds() bool`

HasIntegrationEnvironmentIds returns a boolean if a field has been set.

### SetIntegrationEnvironmentIdsNil

`func (o *IntegrationModel) SetIntegrationEnvironmentIdsNil(b bool)`

 SetIntegrationEnvironmentIdsNil sets the value for IntegrationEnvironmentIds to be an explicit nil

### UnsetIntegrationEnvironmentIds
`func (o *IntegrationModel) UnsetIntegrationEnvironmentIds()`

UnsetIntegrationEnvironmentIds ensures that no value is present for IntegrationEnvironmentIds, not even an explicit nil
### GetIntegrationConfigIds

`func (o *IntegrationModel) GetIntegrationConfigIds() []string`

GetIntegrationConfigIds returns the IntegrationConfigIds field if non-nil, zero value otherwise.

### GetIntegrationConfigIdsOk

`func (o *IntegrationModel) GetIntegrationConfigIdsOk() (*[]string, bool)`

GetIntegrationConfigIdsOk returns a tuple with the IntegrationConfigIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationConfigIds

`func (o *IntegrationModel) SetIntegrationConfigIds(v []string)`

SetIntegrationConfigIds sets IntegrationConfigIds field to given value.

### HasIntegrationConfigIds

`func (o *IntegrationModel) HasIntegrationConfigIds() bool`

HasIntegrationConfigIds returns a boolean if a field has been set.

### SetIntegrationConfigIdsNil

`func (o *IntegrationModel) SetIntegrationConfigIdsNil(b bool)`

 SetIntegrationConfigIdsNil sets the value for IntegrationConfigIds to be an explicit nil

### UnsetIntegrationConfigIds
`func (o *IntegrationModel) UnsetIntegrationConfigIds()`

UnsetIntegrationConfigIds ensures that no value is present for IntegrationConfigIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


