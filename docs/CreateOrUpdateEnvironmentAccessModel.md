# CreateOrUpdateEnvironmentAccessModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentId** | Pointer to **string** | Identifier of the Environment. | [optional] 
**EnvironmentAccessType** | Pointer to **string** | Represent the environment specific Feature Management permission. | [optional] 

## Methods

### NewCreateOrUpdateEnvironmentAccessModel

`func NewCreateOrUpdateEnvironmentAccessModel() *CreateOrUpdateEnvironmentAccessModel`

NewCreateOrUpdateEnvironmentAccessModel instantiates a new CreateOrUpdateEnvironmentAccessModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrUpdateEnvironmentAccessModelWithDefaults

`func NewCreateOrUpdateEnvironmentAccessModelWithDefaults() *CreateOrUpdateEnvironmentAccessModel`

NewCreateOrUpdateEnvironmentAccessModelWithDefaults instantiates a new CreateOrUpdateEnvironmentAccessModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentId

`func (o *CreateOrUpdateEnvironmentAccessModel) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *CreateOrUpdateEnvironmentAccessModel) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *CreateOrUpdateEnvironmentAccessModel) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *CreateOrUpdateEnvironmentAccessModel) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### GetEnvironmentAccessType

`func (o *CreateOrUpdateEnvironmentAccessModel) GetEnvironmentAccessType() string`

GetEnvironmentAccessType returns the EnvironmentAccessType field if non-nil, zero value otherwise.

### GetEnvironmentAccessTypeOk

`func (o *CreateOrUpdateEnvironmentAccessModel) GetEnvironmentAccessTypeOk() (*string, bool)`

GetEnvironmentAccessTypeOk returns a tuple with the EnvironmentAccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentAccessType

`func (o *CreateOrUpdateEnvironmentAccessModel) SetEnvironmentAccessType(v string)`

SetEnvironmentAccessType sets EnvironmentAccessType field to given value.

### HasEnvironmentAccessType

`func (o *CreateOrUpdateEnvironmentAccessModel) HasEnvironmentAccessType() bool`

HasEnvironmentAccessType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


