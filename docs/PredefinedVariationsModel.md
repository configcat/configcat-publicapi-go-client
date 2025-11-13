# PredefinedVariationsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PredefinedVariations** | [**[]PredefinedVariationModel**](PredefinedVariationModel.md) | The Feature Flag or Setting&#39;s Variations. | 
**MaxPredefinedVariations** | **int32** | The maximum number of predefined variations allowed for the Feature Flag or Setting. | 

## Methods

### NewPredefinedVariationsModel

`func NewPredefinedVariationsModel(predefinedVariations []PredefinedVariationModel, maxPredefinedVariations int32, ) *PredefinedVariationsModel`

NewPredefinedVariationsModel instantiates a new PredefinedVariationsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPredefinedVariationsModelWithDefaults

`func NewPredefinedVariationsModelWithDefaults() *PredefinedVariationsModel`

NewPredefinedVariationsModelWithDefaults instantiates a new PredefinedVariationsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPredefinedVariations

`func (o *PredefinedVariationsModel) GetPredefinedVariations() []PredefinedVariationModel`

GetPredefinedVariations returns the PredefinedVariations field if non-nil, zero value otherwise.

### GetPredefinedVariationsOk

`func (o *PredefinedVariationsModel) GetPredefinedVariationsOk() (*[]PredefinedVariationModel, bool)`

GetPredefinedVariationsOk returns a tuple with the PredefinedVariations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredefinedVariations

`func (o *PredefinedVariationsModel) SetPredefinedVariations(v []PredefinedVariationModel)`

SetPredefinedVariations sets PredefinedVariations field to given value.


### GetMaxPredefinedVariations

`func (o *PredefinedVariationsModel) GetMaxPredefinedVariations() int32`

GetMaxPredefinedVariations returns the MaxPredefinedVariations field if non-nil, zero value otherwise.

### GetMaxPredefinedVariationsOk

`func (o *PredefinedVariationsModel) GetMaxPredefinedVariationsOk() (*int32, bool)`

GetMaxPredefinedVariationsOk returns a tuple with the MaxPredefinedVariations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPredefinedVariations

`func (o *PredefinedVariationsModel) SetMaxPredefinedVariations(v int32)`

SetMaxPredefinedVariations sets MaxPredefinedVariations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


