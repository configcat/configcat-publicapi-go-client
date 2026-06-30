# PredefinedVariationWithUsagesModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | [**PredefinedVariationValueModel**](PredefinedVariationValueModel.md) |  | 
**Name** | **NullableString** | The name of the Predefined Variation, shown on the Dashboard UI. If not set, the Value will be shown. | 
**Hint** | **NullableString** | The name of the Predefined Variation, shown on the Dashboard UI. If not set, the Value will be shown. | 
**PredefinedVariationId** | **string** | The Predefined Variation&#39;s identifier. | 
**Usages** | [**[]PredefinedVariationUsageModel**](PredefinedVariationUsageModel.md) | The Feature Flag or Setting Variation&#39;s usages in the given Environments. | 
**UsagesInOtherEnvironments** | **int32** | The Feature Flag or Setting Variation&#39;s usages in the Environments you don&#39;t have access to. | 
**ChangeRequestUsages** | [**[]PredefinedVariationChangeRequestUsageModel**](PredefinedVariationChangeRequestUsageModel.md) | The Feature Flag or Setting Variation&#39;s usages in the given Change Requests. | 
**ChangeRequestUsagesInOtherEnvironments** | **int32** | The Feature Flag or Setting Variation&#39;s usages in the Change Requests you don&#39;t have access to. | 

## Methods

### NewPredefinedVariationWithUsagesModel

`func NewPredefinedVariationWithUsagesModel(value PredefinedVariationValueModel, name NullableString, hint NullableString, predefinedVariationId string, usages []PredefinedVariationUsageModel, usagesInOtherEnvironments int32, changeRequestUsages []PredefinedVariationChangeRequestUsageModel, changeRequestUsagesInOtherEnvironments int32, ) *PredefinedVariationWithUsagesModel`

NewPredefinedVariationWithUsagesModel instantiates a new PredefinedVariationWithUsagesModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPredefinedVariationWithUsagesModelWithDefaults

`func NewPredefinedVariationWithUsagesModelWithDefaults() *PredefinedVariationWithUsagesModel`

NewPredefinedVariationWithUsagesModelWithDefaults instantiates a new PredefinedVariationWithUsagesModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *PredefinedVariationWithUsagesModel) GetValue() PredefinedVariationValueModel`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PredefinedVariationWithUsagesModel) GetValueOk() (*PredefinedVariationValueModel, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PredefinedVariationWithUsagesModel) SetValue(v PredefinedVariationValueModel)`

SetValue sets Value field to given value.


### GetName

`func (o *PredefinedVariationWithUsagesModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PredefinedVariationWithUsagesModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PredefinedVariationWithUsagesModel) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *PredefinedVariationWithUsagesModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PredefinedVariationWithUsagesModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetHint

`func (o *PredefinedVariationWithUsagesModel) GetHint() string`

GetHint returns the Hint field if non-nil, zero value otherwise.

### GetHintOk

`func (o *PredefinedVariationWithUsagesModel) GetHintOk() (*string, bool)`

GetHintOk returns a tuple with the Hint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHint

`func (o *PredefinedVariationWithUsagesModel) SetHint(v string)`

SetHint sets Hint field to given value.


### SetHintNil

`func (o *PredefinedVariationWithUsagesModel) SetHintNil(b bool)`

 SetHintNil sets the value for Hint to be an explicit nil

### UnsetHint
`func (o *PredefinedVariationWithUsagesModel) UnsetHint()`

UnsetHint ensures that no value is present for Hint, not even an explicit nil
### GetPredefinedVariationId

`func (o *PredefinedVariationWithUsagesModel) GetPredefinedVariationId() string`

GetPredefinedVariationId returns the PredefinedVariationId field if non-nil, zero value otherwise.

### GetPredefinedVariationIdOk

`func (o *PredefinedVariationWithUsagesModel) GetPredefinedVariationIdOk() (*string, bool)`

GetPredefinedVariationIdOk returns a tuple with the PredefinedVariationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredefinedVariationId

`func (o *PredefinedVariationWithUsagesModel) SetPredefinedVariationId(v string)`

SetPredefinedVariationId sets PredefinedVariationId field to given value.


### GetUsages

`func (o *PredefinedVariationWithUsagesModel) GetUsages() []PredefinedVariationUsageModel`

GetUsages returns the Usages field if non-nil, zero value otherwise.

### GetUsagesOk

`func (o *PredefinedVariationWithUsagesModel) GetUsagesOk() (*[]PredefinedVariationUsageModel, bool)`

GetUsagesOk returns a tuple with the Usages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsages

`func (o *PredefinedVariationWithUsagesModel) SetUsages(v []PredefinedVariationUsageModel)`

SetUsages sets Usages field to given value.


### GetUsagesInOtherEnvironments

`func (o *PredefinedVariationWithUsagesModel) GetUsagesInOtherEnvironments() int32`

GetUsagesInOtherEnvironments returns the UsagesInOtherEnvironments field if non-nil, zero value otherwise.

### GetUsagesInOtherEnvironmentsOk

`func (o *PredefinedVariationWithUsagesModel) GetUsagesInOtherEnvironmentsOk() (*int32, bool)`

GetUsagesInOtherEnvironmentsOk returns a tuple with the UsagesInOtherEnvironments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsagesInOtherEnvironments

`func (o *PredefinedVariationWithUsagesModel) SetUsagesInOtherEnvironments(v int32)`

SetUsagesInOtherEnvironments sets UsagesInOtherEnvironments field to given value.


### GetChangeRequestUsages

`func (o *PredefinedVariationWithUsagesModel) GetChangeRequestUsages() []PredefinedVariationChangeRequestUsageModel`

GetChangeRequestUsages returns the ChangeRequestUsages field if non-nil, zero value otherwise.

### GetChangeRequestUsagesOk

`func (o *PredefinedVariationWithUsagesModel) GetChangeRequestUsagesOk() (*[]PredefinedVariationChangeRequestUsageModel, bool)`

GetChangeRequestUsagesOk returns a tuple with the ChangeRequestUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeRequestUsages

`func (o *PredefinedVariationWithUsagesModel) SetChangeRequestUsages(v []PredefinedVariationChangeRequestUsageModel)`

SetChangeRequestUsages sets ChangeRequestUsages field to given value.


### GetChangeRequestUsagesInOtherEnvironments

`func (o *PredefinedVariationWithUsagesModel) GetChangeRequestUsagesInOtherEnvironments() int32`

GetChangeRequestUsagesInOtherEnvironments returns the ChangeRequestUsagesInOtherEnvironments field if non-nil, zero value otherwise.

### GetChangeRequestUsagesInOtherEnvironmentsOk

`func (o *PredefinedVariationWithUsagesModel) GetChangeRequestUsagesInOtherEnvironmentsOk() (*int32, bool)`

GetChangeRequestUsagesInOtherEnvironmentsOk returns a tuple with the ChangeRequestUsagesInOtherEnvironments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeRequestUsagesInOtherEnvironments

`func (o *PredefinedVariationWithUsagesModel) SetChangeRequestUsagesInOtherEnvironments(v int32)`

SetChangeRequestUsagesInOtherEnvironments sets ChangeRequestUsagesInOtherEnvironments field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


