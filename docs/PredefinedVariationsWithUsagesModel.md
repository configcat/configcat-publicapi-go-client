# PredefinedVariationsWithUsagesModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SettingKey** | **string** | Key of the Feature Flag or Setting. | 
**SettingType** | [**SettingType**](SettingType.md) |  | 
**PredefinedVariations** | [**[]PredefinedVariationWithUsagesModel**](PredefinedVariationWithUsagesModel.md) | The Feature Flag or Setting&#39;s Variations. | 
**Environments** | [**[]PredefinedVariationEnvironmentModel**](PredefinedVariationEnvironmentModel.md) | The Environment descriptors for the Variations&#39; usages. | 
**MaxPredefinedVariations** | **int32** | The maximum number of predefined variations allowed for the Feature Flag or Setting. | 

## Methods

### NewPredefinedVariationsWithUsagesModel

`func NewPredefinedVariationsWithUsagesModel(settingKey string, settingType SettingType, predefinedVariations []PredefinedVariationWithUsagesModel, environments []PredefinedVariationEnvironmentModel, maxPredefinedVariations int32, ) *PredefinedVariationsWithUsagesModel`

NewPredefinedVariationsWithUsagesModel instantiates a new PredefinedVariationsWithUsagesModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPredefinedVariationsWithUsagesModelWithDefaults

`func NewPredefinedVariationsWithUsagesModelWithDefaults() *PredefinedVariationsWithUsagesModel`

NewPredefinedVariationsWithUsagesModelWithDefaults instantiates a new PredefinedVariationsWithUsagesModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSettingKey

`func (o *PredefinedVariationsWithUsagesModel) GetSettingKey() string`

GetSettingKey returns the SettingKey field if non-nil, zero value otherwise.

### GetSettingKeyOk

`func (o *PredefinedVariationsWithUsagesModel) GetSettingKeyOk() (*string, bool)`

GetSettingKeyOk returns a tuple with the SettingKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingKey

`func (o *PredefinedVariationsWithUsagesModel) SetSettingKey(v string)`

SetSettingKey sets SettingKey field to given value.


### GetSettingType

`func (o *PredefinedVariationsWithUsagesModel) GetSettingType() SettingType`

GetSettingType returns the SettingType field if non-nil, zero value otherwise.

### GetSettingTypeOk

`func (o *PredefinedVariationsWithUsagesModel) GetSettingTypeOk() (*SettingType, bool)`

GetSettingTypeOk returns a tuple with the SettingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingType

`func (o *PredefinedVariationsWithUsagesModel) SetSettingType(v SettingType)`

SetSettingType sets SettingType field to given value.


### GetPredefinedVariations

`func (o *PredefinedVariationsWithUsagesModel) GetPredefinedVariations() []PredefinedVariationWithUsagesModel`

GetPredefinedVariations returns the PredefinedVariations field if non-nil, zero value otherwise.

### GetPredefinedVariationsOk

`func (o *PredefinedVariationsWithUsagesModel) GetPredefinedVariationsOk() (*[]PredefinedVariationWithUsagesModel, bool)`

GetPredefinedVariationsOk returns a tuple with the PredefinedVariations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredefinedVariations

`func (o *PredefinedVariationsWithUsagesModel) SetPredefinedVariations(v []PredefinedVariationWithUsagesModel)`

SetPredefinedVariations sets PredefinedVariations field to given value.


### GetEnvironments

`func (o *PredefinedVariationsWithUsagesModel) GetEnvironments() []PredefinedVariationEnvironmentModel`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *PredefinedVariationsWithUsagesModel) GetEnvironmentsOk() (*[]PredefinedVariationEnvironmentModel, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *PredefinedVariationsWithUsagesModel) SetEnvironments(v []PredefinedVariationEnvironmentModel)`

SetEnvironments sets Environments field to given value.


### GetMaxPredefinedVariations

`func (o *PredefinedVariationsWithUsagesModel) GetMaxPredefinedVariations() int32`

GetMaxPredefinedVariations returns the MaxPredefinedVariations field if non-nil, zero value otherwise.

### GetMaxPredefinedVariationsOk

`func (o *PredefinedVariationsWithUsagesModel) GetMaxPredefinedVariationsOk() (*int32, bool)`

GetMaxPredefinedVariationsOk returns a tuple with the MaxPredefinedVariations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPredefinedVariations

`func (o *PredefinedVariationsWithUsagesModel) SetMaxPredefinedVariations(v int32)`

SetMaxPredefinedVariations sets MaxPredefinedVariations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


