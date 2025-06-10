# IntegrationLinkDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Product** | [**ProductModel**](ProductModel.md) |  | 
**Config** | [**ConfigModel**](ConfigModel.md) |  | 
**Environment** | [**EnvironmentModel**](EnvironmentModel.md) |  | 
**Setting** | [**SettingDataModel**](SettingDataModel.md) |  | 
**ReadOnly** | **bool** |  | 
**Status** | **string** |  | 

## Methods

### NewIntegrationLinkDetail

`func NewIntegrationLinkDetail(product ProductModel, config ConfigModel, environment EnvironmentModel, setting SettingDataModel, readOnly bool, status string, ) *IntegrationLinkDetail`

NewIntegrationLinkDetail instantiates a new IntegrationLinkDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationLinkDetailWithDefaults

`func NewIntegrationLinkDetailWithDefaults() *IntegrationLinkDetail`

NewIntegrationLinkDetailWithDefaults instantiates a new IntegrationLinkDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProduct

`func (o *IntegrationLinkDetail) GetProduct() ProductModel`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *IntegrationLinkDetail) GetProductOk() (*ProductModel, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *IntegrationLinkDetail) SetProduct(v ProductModel)`

SetProduct sets Product field to given value.


### GetConfig

`func (o *IntegrationLinkDetail) GetConfig() ConfigModel`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *IntegrationLinkDetail) GetConfigOk() (*ConfigModel, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *IntegrationLinkDetail) SetConfig(v ConfigModel)`

SetConfig sets Config field to given value.


### GetEnvironment

`func (o *IntegrationLinkDetail) GetEnvironment() EnvironmentModel`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *IntegrationLinkDetail) GetEnvironmentOk() (*EnvironmentModel, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *IntegrationLinkDetail) SetEnvironment(v EnvironmentModel)`

SetEnvironment sets Environment field to given value.


### GetSetting

`func (o *IntegrationLinkDetail) GetSetting() SettingDataModel`

GetSetting returns the Setting field if non-nil, zero value otherwise.

### GetSettingOk

`func (o *IntegrationLinkDetail) GetSettingOk() (*SettingDataModel, bool)`

GetSettingOk returns a tuple with the Setting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetting

`func (o *IntegrationLinkDetail) SetSetting(v SettingDataModel)`

SetSetting sets Setting field to given value.


### GetReadOnly

`func (o *IntegrationLinkDetail) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *IntegrationLinkDetail) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *IntegrationLinkDetail) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.


### GetStatus

`func (o *IntegrationLinkDetail) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IntegrationLinkDetail) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IntegrationLinkDetail) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


