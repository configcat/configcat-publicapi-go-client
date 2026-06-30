# CreateSettingInitialValues

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the Feature Flag or Setting. | 
**Hint** | Pointer to **NullableString** | A short description for the setting, shown on the Dashboard UI. | [optional] 
**Tags** | Pointer to **[]int64** | The IDs of the tags which are attached to the setting. | [optional] 
**Order** | Pointer to **NullableInt32** | The order of the Setting represented on the ConfigCat Dashboard. Determined from an ascending sequence of integers. | [optional] 
**IsJson** | Pointer to **NullableBool** | Indicates whether this setting should validate string values as JSON values. | [optional] 
**Key** | **string** | The key of the Feature Flag or Setting. | 
**SettingType** | [**SettingType**](SettingType.md) |  | 
**PredefinedVariations** | Pointer to [**[]CreatePredefinedVariationModel**](CreatePredefinedVariationModel.md) | The Feature Flag or Setting&#39;s Variations. | [optional] 
**InitialValues** | Pointer to [**[]InitialValue**](InitialValue.md) | Optional, initial value of the Feature Flag or Setting in the given Environments. Only one of the SettingIdToInitFrom or the InitialValues properties can be set. | [optional] 
**SettingIdToInitFrom** | Pointer to **NullableInt32** | Optional, the SettingId to initialize the values and tags of the Feature Flag or Setting from. Only can be set if you have at least ReadOnly access in all the Environments. Only one of the SettingIdToInitFrom or the InitialValues properties can be set. | [optional] 

## Methods

### NewCreateSettingInitialValues

`func NewCreateSettingInitialValues(name string, key string, settingType SettingType, ) *CreateSettingInitialValues`

NewCreateSettingInitialValues instantiates a new CreateSettingInitialValues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSettingInitialValuesWithDefaults

`func NewCreateSettingInitialValuesWithDefaults() *CreateSettingInitialValues`

NewCreateSettingInitialValuesWithDefaults instantiates a new CreateSettingInitialValues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateSettingInitialValues) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateSettingInitialValues) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateSettingInitialValues) SetName(v string)`

SetName sets Name field to given value.


### GetHint

`func (o *CreateSettingInitialValues) GetHint() string`

GetHint returns the Hint field if non-nil, zero value otherwise.

### GetHintOk

`func (o *CreateSettingInitialValues) GetHintOk() (*string, bool)`

GetHintOk returns a tuple with the Hint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHint

`func (o *CreateSettingInitialValues) SetHint(v string)`

SetHint sets Hint field to given value.

### HasHint

`func (o *CreateSettingInitialValues) HasHint() bool`

HasHint returns a boolean if a field has been set.

### SetHintNil

`func (o *CreateSettingInitialValues) SetHintNil(b bool)`

 SetHintNil sets the value for Hint to be an explicit nil

### UnsetHint
`func (o *CreateSettingInitialValues) UnsetHint()`

UnsetHint ensures that no value is present for Hint, not even an explicit nil
### GetTags

`func (o *CreateSettingInitialValues) GetTags() []int64`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateSettingInitialValues) GetTagsOk() (*[]int64, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateSettingInitialValues) SetTags(v []int64)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateSettingInitialValues) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *CreateSettingInitialValues) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *CreateSettingInitialValues) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetOrder

`func (o *CreateSettingInitialValues) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *CreateSettingInitialValues) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *CreateSettingInitialValues) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *CreateSettingInitialValues) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### SetOrderNil

`func (o *CreateSettingInitialValues) SetOrderNil(b bool)`

 SetOrderNil sets the value for Order to be an explicit nil

### UnsetOrder
`func (o *CreateSettingInitialValues) UnsetOrder()`

UnsetOrder ensures that no value is present for Order, not even an explicit nil
### GetIsJson

`func (o *CreateSettingInitialValues) GetIsJson() bool`

GetIsJson returns the IsJson field if non-nil, zero value otherwise.

### GetIsJsonOk

`func (o *CreateSettingInitialValues) GetIsJsonOk() (*bool, bool)`

GetIsJsonOk returns a tuple with the IsJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsJson

`func (o *CreateSettingInitialValues) SetIsJson(v bool)`

SetIsJson sets IsJson field to given value.

### HasIsJson

`func (o *CreateSettingInitialValues) HasIsJson() bool`

HasIsJson returns a boolean if a field has been set.

### SetIsJsonNil

`func (o *CreateSettingInitialValues) SetIsJsonNil(b bool)`

 SetIsJsonNil sets the value for IsJson to be an explicit nil

### UnsetIsJson
`func (o *CreateSettingInitialValues) UnsetIsJson()`

UnsetIsJson ensures that no value is present for IsJson, not even an explicit nil
### GetKey

`func (o *CreateSettingInitialValues) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CreateSettingInitialValues) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CreateSettingInitialValues) SetKey(v string)`

SetKey sets Key field to given value.


### GetSettingType

`func (o *CreateSettingInitialValues) GetSettingType() SettingType`

GetSettingType returns the SettingType field if non-nil, zero value otherwise.

### GetSettingTypeOk

`func (o *CreateSettingInitialValues) GetSettingTypeOk() (*SettingType, bool)`

GetSettingTypeOk returns a tuple with the SettingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingType

`func (o *CreateSettingInitialValues) SetSettingType(v SettingType)`

SetSettingType sets SettingType field to given value.


### GetPredefinedVariations

`func (o *CreateSettingInitialValues) GetPredefinedVariations() []CreatePredefinedVariationModel`

GetPredefinedVariations returns the PredefinedVariations field if non-nil, zero value otherwise.

### GetPredefinedVariationsOk

`func (o *CreateSettingInitialValues) GetPredefinedVariationsOk() (*[]CreatePredefinedVariationModel, bool)`

GetPredefinedVariationsOk returns a tuple with the PredefinedVariations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredefinedVariations

`func (o *CreateSettingInitialValues) SetPredefinedVariations(v []CreatePredefinedVariationModel)`

SetPredefinedVariations sets PredefinedVariations field to given value.

### HasPredefinedVariations

`func (o *CreateSettingInitialValues) HasPredefinedVariations() bool`

HasPredefinedVariations returns a boolean if a field has been set.

### SetPredefinedVariationsNil

`func (o *CreateSettingInitialValues) SetPredefinedVariationsNil(b bool)`

 SetPredefinedVariationsNil sets the value for PredefinedVariations to be an explicit nil

### UnsetPredefinedVariations
`func (o *CreateSettingInitialValues) UnsetPredefinedVariations()`

UnsetPredefinedVariations ensures that no value is present for PredefinedVariations, not even an explicit nil
### GetInitialValues

`func (o *CreateSettingInitialValues) GetInitialValues() []InitialValue`

GetInitialValues returns the InitialValues field if non-nil, zero value otherwise.

### GetInitialValuesOk

`func (o *CreateSettingInitialValues) GetInitialValuesOk() (*[]InitialValue, bool)`

GetInitialValuesOk returns a tuple with the InitialValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialValues

`func (o *CreateSettingInitialValues) SetInitialValues(v []InitialValue)`

SetInitialValues sets InitialValues field to given value.

### HasInitialValues

`func (o *CreateSettingInitialValues) HasInitialValues() bool`

HasInitialValues returns a boolean if a field has been set.

### SetInitialValuesNil

`func (o *CreateSettingInitialValues) SetInitialValuesNil(b bool)`

 SetInitialValuesNil sets the value for InitialValues to be an explicit nil

### UnsetInitialValues
`func (o *CreateSettingInitialValues) UnsetInitialValues()`

UnsetInitialValues ensures that no value is present for InitialValues, not even an explicit nil
### GetSettingIdToInitFrom

`func (o *CreateSettingInitialValues) GetSettingIdToInitFrom() int32`

GetSettingIdToInitFrom returns the SettingIdToInitFrom field if non-nil, zero value otherwise.

### GetSettingIdToInitFromOk

`func (o *CreateSettingInitialValues) GetSettingIdToInitFromOk() (*int32, bool)`

GetSettingIdToInitFromOk returns a tuple with the SettingIdToInitFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingIdToInitFrom

`func (o *CreateSettingInitialValues) SetSettingIdToInitFrom(v int32)`

SetSettingIdToInitFrom sets SettingIdToInitFrom field to given value.

### HasSettingIdToInitFrom

`func (o *CreateSettingInitialValues) HasSettingIdToInitFrom() bool`

HasSettingIdToInitFrom returns a boolean if a field has been set.

### SetSettingIdToInitFromNil

`func (o *CreateSettingInitialValues) SetSettingIdToInitFromNil(b bool)`

 SetSettingIdToInitFromNil sets the value for SettingIdToInitFrom to be an explicit nil

### UnsetSettingIdToInitFrom
`func (o *CreateSettingInitialValues) UnsetSettingIdToInitFrom()`

UnsetSettingIdToInitFrom ensures that no value is present for SettingIdToInitFrom, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


