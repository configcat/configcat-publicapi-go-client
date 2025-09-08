# CreatePredefinedVariationModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | [**CreatePredefinedVariationValueModel**](CreatePredefinedVariationValueModel.md) |  | 
**Name** | Pointer to **NullableString** | The name of the Feature Flag or Predefined Variation, shown on the Dashboard UI. If not set, the Value will be shown. | [optional] 
**Hint** | Pointer to **NullableString** | The name of the Feature Flag or Predefined Variation, shown on the Dashboard UI. If not set, the Value will be shown. | [optional] 

## Methods

### NewCreatePredefinedVariationModel

`func NewCreatePredefinedVariationModel(value CreatePredefinedVariationValueModel, ) *CreatePredefinedVariationModel`

NewCreatePredefinedVariationModel instantiates a new CreatePredefinedVariationModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePredefinedVariationModelWithDefaults

`func NewCreatePredefinedVariationModelWithDefaults() *CreatePredefinedVariationModel`

NewCreatePredefinedVariationModelWithDefaults instantiates a new CreatePredefinedVariationModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *CreatePredefinedVariationModel) GetValue() CreatePredefinedVariationValueModel`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CreatePredefinedVariationModel) GetValueOk() (*CreatePredefinedVariationValueModel, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CreatePredefinedVariationModel) SetValue(v CreatePredefinedVariationValueModel)`

SetValue sets Value field to given value.


### GetName

`func (o *CreatePredefinedVariationModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreatePredefinedVariationModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreatePredefinedVariationModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreatePredefinedVariationModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CreatePredefinedVariationModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreatePredefinedVariationModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetHint

`func (o *CreatePredefinedVariationModel) GetHint() string`

GetHint returns the Hint field if non-nil, zero value otherwise.

### GetHintOk

`func (o *CreatePredefinedVariationModel) GetHintOk() (*string, bool)`

GetHintOk returns a tuple with the Hint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHint

`func (o *CreatePredefinedVariationModel) SetHint(v string)`

SetHint sets Hint field to given value.

### HasHint

`func (o *CreatePredefinedVariationModel) HasHint() bool`

HasHint returns a boolean if a field has been set.

### SetHintNil

`func (o *CreatePredefinedVariationModel) SetHintNil(b bool)`

 SetHintNil sets the value for Hint to be an explicit nil

### UnsetHint
`func (o *CreatePredefinedVariationModel) UnsetHint()`

UnsetHint ensures that no value is present for Hint, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


