# UpdatePredefinedVariationModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | [**UpdatePredefinedVariationValueModel**](UpdatePredefinedVariationValueModel.md) |  | 
**Name** | Pointer to **NullableString** | The name of the Predefined Variation, shown on the Dashboard UI. If not set, the Value will be shown. | [optional] 
**Hint** | Pointer to **NullableString** | The name of the Predefined Variation, shown on the Dashboard UI. If not set, the Value will be shown. | [optional] 
**PredefinedVariationId** | Pointer to **NullableString** | The Predefined Variation&#39;s identifier to update. Omit the value if you want to add a new predefined variation. | [optional] 

## Methods

### NewUpdatePredefinedVariationModel

`func NewUpdatePredefinedVariationModel(value UpdatePredefinedVariationValueModel, ) *UpdatePredefinedVariationModel`

NewUpdatePredefinedVariationModel instantiates a new UpdatePredefinedVariationModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePredefinedVariationModelWithDefaults

`func NewUpdatePredefinedVariationModelWithDefaults() *UpdatePredefinedVariationModel`

NewUpdatePredefinedVariationModelWithDefaults instantiates a new UpdatePredefinedVariationModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *UpdatePredefinedVariationModel) GetValue() UpdatePredefinedVariationValueModel`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UpdatePredefinedVariationModel) GetValueOk() (*UpdatePredefinedVariationValueModel, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UpdatePredefinedVariationModel) SetValue(v UpdatePredefinedVariationValueModel)`

SetValue sets Value field to given value.


### GetName

`func (o *UpdatePredefinedVariationModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdatePredefinedVariationModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdatePredefinedVariationModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdatePredefinedVariationModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UpdatePredefinedVariationModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdatePredefinedVariationModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetHint

`func (o *UpdatePredefinedVariationModel) GetHint() string`

GetHint returns the Hint field if non-nil, zero value otherwise.

### GetHintOk

`func (o *UpdatePredefinedVariationModel) GetHintOk() (*string, bool)`

GetHintOk returns a tuple with the Hint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHint

`func (o *UpdatePredefinedVariationModel) SetHint(v string)`

SetHint sets Hint field to given value.

### HasHint

`func (o *UpdatePredefinedVariationModel) HasHint() bool`

HasHint returns a boolean if a field has been set.

### SetHintNil

`func (o *UpdatePredefinedVariationModel) SetHintNil(b bool)`

 SetHintNil sets the value for Hint to be an explicit nil

### UnsetHint
`func (o *UpdatePredefinedVariationModel) UnsetHint()`

UnsetHint ensures that no value is present for Hint, not even an explicit nil
### GetPredefinedVariationId

`func (o *UpdatePredefinedVariationModel) GetPredefinedVariationId() string`

GetPredefinedVariationId returns the PredefinedVariationId field if non-nil, zero value otherwise.

### GetPredefinedVariationIdOk

`func (o *UpdatePredefinedVariationModel) GetPredefinedVariationIdOk() (*string, bool)`

GetPredefinedVariationIdOk returns a tuple with the PredefinedVariationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredefinedVariationId

`func (o *UpdatePredefinedVariationModel) SetPredefinedVariationId(v string)`

SetPredefinedVariationId sets PredefinedVariationId field to given value.

### HasPredefinedVariationId

`func (o *UpdatePredefinedVariationModel) HasPredefinedVariationId() bool`

HasPredefinedVariationId returns a boolean if a field has been set.

### SetPredefinedVariationIdNil

`func (o *UpdatePredefinedVariationModel) SetPredefinedVariationIdNil(b bool)`

 SetPredefinedVariationIdNil sets the value for PredefinedVariationId to be an explicit nil

### UnsetPredefinedVariationId
`func (o *UpdatePredefinedVariationModel) UnsetPredefinedVariationId()`

UnsetPredefinedVariationId ensures that no value is present for PredefinedVariationId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


