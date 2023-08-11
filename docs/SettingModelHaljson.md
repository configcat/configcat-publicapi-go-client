# SettingModelHaljson

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SettingId** | Pointer to **int32** | Identifier of the Feature Flag or Setting. | [optional] 
**Key** | Pointer to **NullableString** | Key of the Feature Flag or Setting. | [optional] 
**Name** | Pointer to **NullableString** | Name of the Feature Flag or Setting. | [optional] 
**Hint** | Pointer to **NullableString** | Description of the Feature Flag or Setting. | [optional] 
**Order** | Pointer to **int32** | The order of the Feature Flag or Setting represented on the ConfigCat Dashboard. | [optional] 
**SettingType** | Pointer to [**SettingType**](SettingType.md) |  | [optional] 
**ConfigId** | Pointer to **string** | Identifier of the Feature Flag&#39;s Config. | [optional] 
**ConfigName** | Pointer to **NullableString** | Name of the Feature Flag&#39;s Config. | [optional] 
**Embedded** | Pointer to [**SettingModelHaljsonEmbedded**](SettingModelHaljsonEmbedded.md) |  | [optional] 
**Links** | Pointer to [**ConfigSettingFormulasModelHaljsonEmbeddedEnvironmentLinks**](ConfigSettingFormulasModelHaljsonEmbeddedEnvironmentLinks.md) |  | [optional] 

## Methods

### NewSettingModelHaljson

`func NewSettingModelHaljson() *SettingModelHaljson`

NewSettingModelHaljson instantiates a new SettingModelHaljson object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettingModelHaljsonWithDefaults

`func NewSettingModelHaljsonWithDefaults() *SettingModelHaljson`

NewSettingModelHaljsonWithDefaults instantiates a new SettingModelHaljson object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSettingId

`func (o *SettingModelHaljson) GetSettingId() int32`

GetSettingId returns the SettingId field if non-nil, zero value otherwise.

### GetSettingIdOk

`func (o *SettingModelHaljson) GetSettingIdOk() (*int32, bool)`

GetSettingIdOk returns a tuple with the SettingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingId

`func (o *SettingModelHaljson) SetSettingId(v int32)`

SetSettingId sets SettingId field to given value.

### HasSettingId

`func (o *SettingModelHaljson) HasSettingId() bool`

HasSettingId returns a boolean if a field has been set.

### GetKey

`func (o *SettingModelHaljson) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SettingModelHaljson) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SettingModelHaljson) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *SettingModelHaljson) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *SettingModelHaljson) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *SettingModelHaljson) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetName

`func (o *SettingModelHaljson) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SettingModelHaljson) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SettingModelHaljson) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SettingModelHaljson) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *SettingModelHaljson) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *SettingModelHaljson) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetHint

`func (o *SettingModelHaljson) GetHint() string`

GetHint returns the Hint field if non-nil, zero value otherwise.

### GetHintOk

`func (o *SettingModelHaljson) GetHintOk() (*string, bool)`

GetHintOk returns a tuple with the Hint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHint

`func (o *SettingModelHaljson) SetHint(v string)`

SetHint sets Hint field to given value.

### HasHint

`func (o *SettingModelHaljson) HasHint() bool`

HasHint returns a boolean if a field has been set.

### SetHintNil

`func (o *SettingModelHaljson) SetHintNil(b bool)`

 SetHintNil sets the value for Hint to be an explicit nil

### UnsetHint
`func (o *SettingModelHaljson) UnsetHint()`

UnsetHint ensures that no value is present for Hint, not even an explicit nil
### GetOrder

`func (o *SettingModelHaljson) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *SettingModelHaljson) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *SettingModelHaljson) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *SettingModelHaljson) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetSettingType

`func (o *SettingModelHaljson) GetSettingType() SettingType`

GetSettingType returns the SettingType field if non-nil, zero value otherwise.

### GetSettingTypeOk

`func (o *SettingModelHaljson) GetSettingTypeOk() (*SettingType, bool)`

GetSettingTypeOk returns a tuple with the SettingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingType

`func (o *SettingModelHaljson) SetSettingType(v SettingType)`

SetSettingType sets SettingType field to given value.

### HasSettingType

`func (o *SettingModelHaljson) HasSettingType() bool`

HasSettingType returns a boolean if a field has been set.

### GetConfigId

`func (o *SettingModelHaljson) GetConfigId() string`

GetConfigId returns the ConfigId field if non-nil, zero value otherwise.

### GetConfigIdOk

`func (o *SettingModelHaljson) GetConfigIdOk() (*string, bool)`

GetConfigIdOk returns a tuple with the ConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigId

`func (o *SettingModelHaljson) SetConfigId(v string)`

SetConfigId sets ConfigId field to given value.

### HasConfigId

`func (o *SettingModelHaljson) HasConfigId() bool`

HasConfigId returns a boolean if a field has been set.

### GetConfigName

`func (o *SettingModelHaljson) GetConfigName() string`

GetConfigName returns the ConfigName field if non-nil, zero value otherwise.

### GetConfigNameOk

`func (o *SettingModelHaljson) GetConfigNameOk() (*string, bool)`

GetConfigNameOk returns a tuple with the ConfigName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigName

`func (o *SettingModelHaljson) SetConfigName(v string)`

SetConfigName sets ConfigName field to given value.

### HasConfigName

`func (o *SettingModelHaljson) HasConfigName() bool`

HasConfigName returns a boolean if a field has been set.

### SetConfigNameNil

`func (o *SettingModelHaljson) SetConfigNameNil(b bool)`

 SetConfigNameNil sets the value for ConfigName to be an explicit nil

### UnsetConfigName
`func (o *SettingModelHaljson) UnsetConfigName()`

UnsetConfigName ensures that no value is present for ConfigName, not even an explicit nil
### GetEmbedded

`func (o *SettingModelHaljson) GetEmbedded() SettingModelHaljsonEmbedded`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *SettingModelHaljson) GetEmbeddedOk() (*SettingModelHaljsonEmbedded, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *SettingModelHaljson) SetEmbedded(v SettingModelHaljsonEmbedded)`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *SettingModelHaljson) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetLinks

`func (o *SettingModelHaljson) GetLinks() ConfigSettingFormulasModelHaljsonEmbeddedEnvironmentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *SettingModelHaljson) GetLinksOk() (*ConfigSettingFormulasModelHaljsonEmbeddedEnvironmentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *SettingModelHaljson) SetLinks(v ConfigSettingFormulasModelHaljsonEmbeddedEnvironmentLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *SettingModelHaljson) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


