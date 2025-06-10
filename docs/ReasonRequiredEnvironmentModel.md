# ReasonRequiredEnvironmentModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentId** | **string** | Identifier of the Environment. | 
**ReasonRequired** | **bool** | Indicates that a mandatory note is required in this Environment for saving and publishing. | 
**EnvironmentName** | **NullableString** | Name of the Environment. | 

## Methods

### NewReasonRequiredEnvironmentModel

`func NewReasonRequiredEnvironmentModel(environmentId string, reasonRequired bool, environmentName NullableString, ) *ReasonRequiredEnvironmentModel`

NewReasonRequiredEnvironmentModel instantiates a new ReasonRequiredEnvironmentModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReasonRequiredEnvironmentModelWithDefaults

`func NewReasonRequiredEnvironmentModelWithDefaults() *ReasonRequiredEnvironmentModel`

NewReasonRequiredEnvironmentModelWithDefaults instantiates a new ReasonRequiredEnvironmentModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentId

`func (o *ReasonRequiredEnvironmentModel) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *ReasonRequiredEnvironmentModel) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *ReasonRequiredEnvironmentModel) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetReasonRequired

`func (o *ReasonRequiredEnvironmentModel) GetReasonRequired() bool`

GetReasonRequired returns the ReasonRequired field if non-nil, zero value otherwise.

### GetReasonRequiredOk

`func (o *ReasonRequiredEnvironmentModel) GetReasonRequiredOk() (*bool, bool)`

GetReasonRequiredOk returns a tuple with the ReasonRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonRequired

`func (o *ReasonRequiredEnvironmentModel) SetReasonRequired(v bool)`

SetReasonRequired sets ReasonRequired field to given value.


### GetEnvironmentName

`func (o *ReasonRequiredEnvironmentModel) GetEnvironmentName() string`

GetEnvironmentName returns the EnvironmentName field if non-nil, zero value otherwise.

### GetEnvironmentNameOk

`func (o *ReasonRequiredEnvironmentModel) GetEnvironmentNameOk() (*string, bool)`

GetEnvironmentNameOk returns a tuple with the EnvironmentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentName

`func (o *ReasonRequiredEnvironmentModel) SetEnvironmentName(v string)`

SetEnvironmentName sets EnvironmentName field to given value.


### SetEnvironmentNameNil

`func (o *ReasonRequiredEnvironmentModel) SetEnvironmentNameNil(b bool)`

 SetEnvironmentNameNil sets the value for EnvironmentName to be an explicit nil

### UnsetEnvironmentName
`func (o *ReasonRequiredEnvironmentModel) UnsetEnvironmentName()`

UnsetEnvironmentName ensures that no value is present for EnvironmentName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


