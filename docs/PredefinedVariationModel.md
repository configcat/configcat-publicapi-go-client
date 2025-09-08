# PredefinedVariationModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | [**PredefinedVariationValueModel**](PredefinedVariationValueModel.md) |  | 
**Name** | **NullableString** | The name of the Feature Flag or Predefined Variation, shown on the Dashboard UI. If not set, the Value will be shown. | 
**Hint** | **NullableString** | The name of the Feature Flag or Predefined Variation, shown on the Dashboard UI. If not set, the Value will be shown. | 
**PredefinedVariationId** | **string** | The Feature Flag or Predefined Variation&#39;s identifier. | 

## Methods

### NewPredefinedVariationModel

`func NewPredefinedVariationModel(value PredefinedVariationValueModel, name NullableString, hint NullableString, predefinedVariationId string, ) *PredefinedVariationModel`

NewPredefinedVariationModel instantiates a new PredefinedVariationModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPredefinedVariationModelWithDefaults

`func NewPredefinedVariationModelWithDefaults() *PredefinedVariationModel`

NewPredefinedVariationModelWithDefaults instantiates a new PredefinedVariationModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *PredefinedVariationModel) GetValue() PredefinedVariationValueModel`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PredefinedVariationModel) GetValueOk() (*PredefinedVariationValueModel, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PredefinedVariationModel) SetValue(v PredefinedVariationValueModel)`

SetValue sets Value field to given value.


### GetName

`func (o *PredefinedVariationModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PredefinedVariationModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PredefinedVariationModel) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *PredefinedVariationModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PredefinedVariationModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetHint

`func (o *PredefinedVariationModel) GetHint() string`

GetHint returns the Hint field if non-nil, zero value otherwise.

### GetHintOk

`func (o *PredefinedVariationModel) GetHintOk() (*string, bool)`

GetHintOk returns a tuple with the Hint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHint

`func (o *PredefinedVariationModel) SetHint(v string)`

SetHint sets Hint field to given value.


### SetHintNil

`func (o *PredefinedVariationModel) SetHintNil(b bool)`

 SetHintNil sets the value for Hint to be an explicit nil

### UnsetHint
`func (o *PredefinedVariationModel) UnsetHint()`

UnsetHint ensures that no value is present for Hint, not even an explicit nil
### GetPredefinedVariationId

`func (o *PredefinedVariationModel) GetPredefinedVariationId() string`

GetPredefinedVariationId returns the PredefinedVariationId field if non-nil, zero value otherwise.

### GetPredefinedVariationIdOk

`func (o *PredefinedVariationModel) GetPredefinedVariationIdOk() (*string, bool)`

GetPredefinedVariationIdOk returns a tuple with the PredefinedVariationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredefinedVariationId

`func (o *PredefinedVariationModel) SetPredefinedVariationId(v string)`

SetPredefinedVariationId sets PredefinedVariationId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


