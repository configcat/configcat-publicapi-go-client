# ApproveRequiredEnvironmentModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentId** | **string** | Identifier of the Environment. | 
**ApproveRequired** | **bool** | Indicates that a mandatory approval is required in this Environment before changes are applied. | 
**EnvironmentName** | **NullableString** | Name of the Environment. | 

## Methods

### NewApproveRequiredEnvironmentModel

`func NewApproveRequiredEnvironmentModel(environmentId string, approveRequired bool, environmentName NullableString, ) *ApproveRequiredEnvironmentModel`

NewApproveRequiredEnvironmentModel instantiates a new ApproveRequiredEnvironmentModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApproveRequiredEnvironmentModelWithDefaults

`func NewApproveRequiredEnvironmentModelWithDefaults() *ApproveRequiredEnvironmentModel`

NewApproveRequiredEnvironmentModelWithDefaults instantiates a new ApproveRequiredEnvironmentModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentId

`func (o *ApproveRequiredEnvironmentModel) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *ApproveRequiredEnvironmentModel) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *ApproveRequiredEnvironmentModel) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetApproveRequired

`func (o *ApproveRequiredEnvironmentModel) GetApproveRequired() bool`

GetApproveRequired returns the ApproveRequired field if non-nil, zero value otherwise.

### GetApproveRequiredOk

`func (o *ApproveRequiredEnvironmentModel) GetApproveRequiredOk() (*bool, bool)`

GetApproveRequiredOk returns a tuple with the ApproveRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproveRequired

`func (o *ApproveRequiredEnvironmentModel) SetApproveRequired(v bool)`

SetApproveRequired sets ApproveRequired field to given value.


### GetEnvironmentName

`func (o *ApproveRequiredEnvironmentModel) GetEnvironmentName() string`

GetEnvironmentName returns the EnvironmentName field if non-nil, zero value otherwise.

### GetEnvironmentNameOk

`func (o *ApproveRequiredEnvironmentModel) GetEnvironmentNameOk() (*string, bool)`

GetEnvironmentNameOk returns a tuple with the EnvironmentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentName

`func (o *ApproveRequiredEnvironmentModel) SetEnvironmentName(v string)`

SetEnvironmentName sets EnvironmentName field to given value.


### SetEnvironmentNameNil

`func (o *ApproveRequiredEnvironmentModel) SetEnvironmentNameNil(b bool)`

 SetEnvironmentNameNil sets the value for EnvironmentName to be an explicit nil

### UnsetEnvironmentName
`func (o *ApproveRequiredEnvironmentModel) UnsetEnvironmentName()`

UnsetEnvironmentName ensures that no value is present for EnvironmentName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


