# IntegrationModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Product** | [**ProductModel**](ProductModel.md) |  | 
**IntegrationId** | **string** | Identifier of the Integration. | 
**Name** | **string** | Name of the Integration. | 
**IntegrationType** | [**IntegrationType**](IntegrationType.md) |  | 
**Parameters** | **map[string]string** | Parameters of the Integration. | 
**EnvironmentIds** | **[]string** | List of Environment IDs that are connected with this Integration. If the list is empty, all of the Environments are connected. | 
**ConfigIds** | **[]string** | List of Config IDs that are connected with this Integration. If the list is empty, all of the Configs are connected. | 

## Methods

### NewIntegrationModel

`func NewIntegrationModel(product ProductModel, integrationId string, name string, integrationType IntegrationType, parameters map[string]string, environmentIds []string, configIds []string, ) *IntegrationModel`

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


