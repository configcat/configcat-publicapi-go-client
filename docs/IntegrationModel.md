# IntegrationModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Product** | Pointer to [**ProductModel**](ProductModel.md) |  | [optional] 
**IntegrationId** | Pointer to **string** | Identifier of the Integration. | [optional] 
**Name** | Pointer to **NullableString** | Name of the Integration. | [optional] 
**IntegrationType** | Pointer to [**IntegrationType**](IntegrationType.md) |  | [optional] 
**Parameters** | Pointer to **map[string]string** | Parameters of the Integration. | [optional] 
**EnvironmentIds** | Pointer to **[]string** | List of Environment IDs that are connected with this Integration. If the list is empty, all of the Environments are connected. | [optional] 
**ConfigIds** | Pointer to **[]string** | List of Config IDs that are connected with this Integration. If the list is empty, all of the Configs are connected. | [optional] 

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

### GetProduct

`func (o *IntegrationModel) GetProduct() ProductModel`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *IntegrationModel) GetProductOk() (*ProductModel, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *IntegrationModel) SetProduct(v ProductModel)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *IntegrationModel) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

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
### GetEnvironmentIds

`func (o *IntegrationModel) GetEnvironmentIds() []string`

GetEnvironmentIds returns the EnvironmentIds field if non-nil, zero value otherwise.

### GetEnvironmentIdsOk

`func (o *IntegrationModel) GetEnvironmentIdsOk() (*[]string, bool)`

GetEnvironmentIdsOk returns a tuple with the EnvironmentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentIds

`func (o *IntegrationModel) SetEnvironmentIds(v []string)`

SetEnvironmentIds sets EnvironmentIds field to given value.

### HasEnvironmentIds

`func (o *IntegrationModel) HasEnvironmentIds() bool`

HasEnvironmentIds returns a boolean if a field has been set.

### SetEnvironmentIdsNil

`func (o *IntegrationModel) SetEnvironmentIdsNil(b bool)`

 SetEnvironmentIdsNil sets the value for EnvironmentIds to be an explicit nil

### UnsetEnvironmentIds
`func (o *IntegrationModel) UnsetEnvironmentIds()`

UnsetEnvironmentIds ensures that no value is present for EnvironmentIds, not even an explicit nil
### GetConfigIds

`func (o *IntegrationModel) GetConfigIds() []string`

GetConfigIds returns the ConfigIds field if non-nil, zero value otherwise.

### GetConfigIdsOk

`func (o *IntegrationModel) GetConfigIdsOk() (*[]string, bool)`

GetConfigIdsOk returns a tuple with the ConfigIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigIds

`func (o *IntegrationModel) SetConfigIds(v []string)`

SetConfigIds sets ConfigIds field to given value.

### HasConfigIds

`func (o *IntegrationModel) HasConfigIds() bool`

HasConfigIds returns a boolean if a field has been set.

### SetConfigIdsNil

`func (o *IntegrationModel) SetConfigIdsNil(b bool)`

 SetConfigIdsNil sets the value for ConfigIds to be an explicit nil

### UnsetConfigIds
`func (o *IntegrationModel) UnsetConfigIds()`

UnsetConfigIds ensures that no value is present for ConfigIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


