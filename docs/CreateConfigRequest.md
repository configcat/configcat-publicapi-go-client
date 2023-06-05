# CreateConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**EvaluationVersion** | Pointer to [**EvaluationVersion**](EvaluationVersion.md) |  | [optional] 

## Methods

### NewCreateConfigRequest

`func NewCreateConfigRequest(name string, ) *CreateConfigRequest`

NewCreateConfigRequest instantiates a new CreateConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateConfigRequestWithDefaults

`func NewCreateConfigRequestWithDefaults() *CreateConfigRequest`

NewCreateConfigRequestWithDefaults instantiates a new CreateConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateConfigRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateConfigRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateConfigRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateConfigRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateConfigRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateConfigRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateConfigRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateConfigRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateConfigRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEvaluationVersion

`func (o *CreateConfigRequest) GetEvaluationVersion() EvaluationVersion`

GetEvaluationVersion returns the EvaluationVersion field if non-nil, zero value otherwise.

### GetEvaluationVersionOk

`func (o *CreateConfigRequest) GetEvaluationVersionOk() (*EvaluationVersion, bool)`

GetEvaluationVersionOk returns a tuple with the EvaluationVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationVersion

`func (o *CreateConfigRequest) SetEvaluationVersion(v EvaluationVersion)`

SetEvaluationVersion sets EvaluationVersion field to given value.

### HasEvaluationVersion

`func (o *CreateConfigRequest) HasEvaluationVersion() bool`

HasEvaluationVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


