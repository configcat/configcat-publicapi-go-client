# StaleFlagProductModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductId** | **string** | Identifier of the Product. | 
**Name** | **string** | Name of the Product. | 
**Configs** | [**[]StaleFlagConfigModel**](StaleFlagConfigModel.md) | Configs that contain stale feature flags. | 
**Environments** | [**[]StaleFlagEnvironmentModel**](StaleFlagEnvironmentModel.md) | Environment list. | 

## Methods

### NewStaleFlagProductModel

`func NewStaleFlagProductModel(productId string, name string, configs []StaleFlagConfigModel, environments []StaleFlagEnvironmentModel, ) *StaleFlagProductModel`

NewStaleFlagProductModel instantiates a new StaleFlagProductModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStaleFlagProductModelWithDefaults

`func NewStaleFlagProductModelWithDefaults() *StaleFlagProductModel`

NewStaleFlagProductModelWithDefaults instantiates a new StaleFlagProductModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductId

`func (o *StaleFlagProductModel) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *StaleFlagProductModel) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *StaleFlagProductModel) SetProductId(v string)`

SetProductId sets ProductId field to given value.


### GetName

`func (o *StaleFlagProductModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StaleFlagProductModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StaleFlagProductModel) SetName(v string)`

SetName sets Name field to given value.


### GetConfigs

`func (o *StaleFlagProductModel) GetConfigs() []StaleFlagConfigModel`

GetConfigs returns the Configs field if non-nil, zero value otherwise.

### GetConfigsOk

`func (o *StaleFlagProductModel) GetConfigsOk() (*[]StaleFlagConfigModel, bool)`

GetConfigsOk returns a tuple with the Configs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigs

`func (o *StaleFlagProductModel) SetConfigs(v []StaleFlagConfigModel)`

SetConfigs sets Configs field to given value.


### GetEnvironments

`func (o *StaleFlagProductModel) GetEnvironments() []StaleFlagEnvironmentModel`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *StaleFlagProductModel) GetEnvironmentsOk() (*[]StaleFlagEnvironmentModel, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *StaleFlagProductModel) SetEnvironments(v []StaleFlagEnvironmentModel)`

SetEnvironments sets Environments field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


