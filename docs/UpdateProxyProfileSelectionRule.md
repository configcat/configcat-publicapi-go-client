# UpdateProxyProfileSelectionRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to [**SelectionRuleKind**](SelectionRuleKind.md) |  | [optional] 
**ProductIdFilter** | Pointer to **NullableString** | Defines the filter for matching Products by their unique identifier. | [optional] 
**ConfigIdFilter** | Pointer to **NullableString** | Defines the filter for matching Configs by their unique identifier. | [optional] 
**EnvironmentIdFilter** | Pointer to **NullableString** | Defines the filter for matching Environments by their unique identifier. | [optional] 
**ProductNameMatchFilter** | Pointer to **NullableString** | Specifies a filter to match Product names in the proxy profile selection rule. It accepts wildcards (*). | [optional] 
**ConfigNameMatchFilter** | Pointer to **NullableString** | Specifies a filter to match Config names in the proxy profile selection rule. It accepts wildcards (*). | [optional] 
**EnvironmentNameMatchFilter** | Pointer to **NullableString** | Specifies a filter to match Environment names in the proxy profile selection rule. It accepts wildcards (*). | [optional] 

## Methods

### NewUpdateProxyProfileSelectionRule

`func NewUpdateProxyProfileSelectionRule() *UpdateProxyProfileSelectionRule`

NewUpdateProxyProfileSelectionRule instantiates a new UpdateProxyProfileSelectionRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateProxyProfileSelectionRuleWithDefaults

`func NewUpdateProxyProfileSelectionRuleWithDefaults() *UpdateProxyProfileSelectionRule`

NewUpdateProxyProfileSelectionRuleWithDefaults instantiates a new UpdateProxyProfileSelectionRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *UpdateProxyProfileSelectionRule) GetKind() SelectionRuleKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *UpdateProxyProfileSelectionRule) GetKindOk() (*SelectionRuleKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *UpdateProxyProfileSelectionRule) SetKind(v SelectionRuleKind)`

SetKind sets Kind field to given value.

### HasKind

`func (o *UpdateProxyProfileSelectionRule) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetProductIdFilter

`func (o *UpdateProxyProfileSelectionRule) GetProductIdFilter() string`

GetProductIdFilter returns the ProductIdFilter field if non-nil, zero value otherwise.

### GetProductIdFilterOk

`func (o *UpdateProxyProfileSelectionRule) GetProductIdFilterOk() (*string, bool)`

GetProductIdFilterOk returns a tuple with the ProductIdFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductIdFilter

`func (o *UpdateProxyProfileSelectionRule) SetProductIdFilter(v string)`

SetProductIdFilter sets ProductIdFilter field to given value.

### HasProductIdFilter

`func (o *UpdateProxyProfileSelectionRule) HasProductIdFilter() bool`

HasProductIdFilter returns a boolean if a field has been set.

### SetProductIdFilterNil

`func (o *UpdateProxyProfileSelectionRule) SetProductIdFilterNil(b bool)`

 SetProductIdFilterNil sets the value for ProductIdFilter to be an explicit nil

### UnsetProductIdFilter
`func (o *UpdateProxyProfileSelectionRule) UnsetProductIdFilter()`

UnsetProductIdFilter ensures that no value is present for ProductIdFilter, not even an explicit nil
### GetConfigIdFilter

`func (o *UpdateProxyProfileSelectionRule) GetConfigIdFilter() string`

GetConfigIdFilter returns the ConfigIdFilter field if non-nil, zero value otherwise.

### GetConfigIdFilterOk

`func (o *UpdateProxyProfileSelectionRule) GetConfigIdFilterOk() (*string, bool)`

GetConfigIdFilterOk returns a tuple with the ConfigIdFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigIdFilter

`func (o *UpdateProxyProfileSelectionRule) SetConfigIdFilter(v string)`

SetConfigIdFilter sets ConfigIdFilter field to given value.

### HasConfigIdFilter

`func (o *UpdateProxyProfileSelectionRule) HasConfigIdFilter() bool`

HasConfigIdFilter returns a boolean if a field has been set.

### SetConfigIdFilterNil

`func (o *UpdateProxyProfileSelectionRule) SetConfigIdFilterNil(b bool)`

 SetConfigIdFilterNil sets the value for ConfigIdFilter to be an explicit nil

### UnsetConfigIdFilter
`func (o *UpdateProxyProfileSelectionRule) UnsetConfigIdFilter()`

UnsetConfigIdFilter ensures that no value is present for ConfigIdFilter, not even an explicit nil
### GetEnvironmentIdFilter

`func (o *UpdateProxyProfileSelectionRule) GetEnvironmentIdFilter() string`

GetEnvironmentIdFilter returns the EnvironmentIdFilter field if non-nil, zero value otherwise.

### GetEnvironmentIdFilterOk

`func (o *UpdateProxyProfileSelectionRule) GetEnvironmentIdFilterOk() (*string, bool)`

GetEnvironmentIdFilterOk returns a tuple with the EnvironmentIdFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentIdFilter

`func (o *UpdateProxyProfileSelectionRule) SetEnvironmentIdFilter(v string)`

SetEnvironmentIdFilter sets EnvironmentIdFilter field to given value.

### HasEnvironmentIdFilter

`func (o *UpdateProxyProfileSelectionRule) HasEnvironmentIdFilter() bool`

HasEnvironmentIdFilter returns a boolean if a field has been set.

### SetEnvironmentIdFilterNil

`func (o *UpdateProxyProfileSelectionRule) SetEnvironmentIdFilterNil(b bool)`

 SetEnvironmentIdFilterNil sets the value for EnvironmentIdFilter to be an explicit nil

### UnsetEnvironmentIdFilter
`func (o *UpdateProxyProfileSelectionRule) UnsetEnvironmentIdFilter()`

UnsetEnvironmentIdFilter ensures that no value is present for EnvironmentIdFilter, not even an explicit nil
### GetProductNameMatchFilter

`func (o *UpdateProxyProfileSelectionRule) GetProductNameMatchFilter() string`

GetProductNameMatchFilter returns the ProductNameMatchFilter field if non-nil, zero value otherwise.

### GetProductNameMatchFilterOk

`func (o *UpdateProxyProfileSelectionRule) GetProductNameMatchFilterOk() (*string, bool)`

GetProductNameMatchFilterOk returns a tuple with the ProductNameMatchFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductNameMatchFilter

`func (o *UpdateProxyProfileSelectionRule) SetProductNameMatchFilter(v string)`

SetProductNameMatchFilter sets ProductNameMatchFilter field to given value.

### HasProductNameMatchFilter

`func (o *UpdateProxyProfileSelectionRule) HasProductNameMatchFilter() bool`

HasProductNameMatchFilter returns a boolean if a field has been set.

### SetProductNameMatchFilterNil

`func (o *UpdateProxyProfileSelectionRule) SetProductNameMatchFilterNil(b bool)`

 SetProductNameMatchFilterNil sets the value for ProductNameMatchFilter to be an explicit nil

### UnsetProductNameMatchFilter
`func (o *UpdateProxyProfileSelectionRule) UnsetProductNameMatchFilter()`

UnsetProductNameMatchFilter ensures that no value is present for ProductNameMatchFilter, not even an explicit nil
### GetConfigNameMatchFilter

`func (o *UpdateProxyProfileSelectionRule) GetConfigNameMatchFilter() string`

GetConfigNameMatchFilter returns the ConfigNameMatchFilter field if non-nil, zero value otherwise.

### GetConfigNameMatchFilterOk

`func (o *UpdateProxyProfileSelectionRule) GetConfigNameMatchFilterOk() (*string, bool)`

GetConfigNameMatchFilterOk returns a tuple with the ConfigNameMatchFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigNameMatchFilter

`func (o *UpdateProxyProfileSelectionRule) SetConfigNameMatchFilter(v string)`

SetConfigNameMatchFilter sets ConfigNameMatchFilter field to given value.

### HasConfigNameMatchFilter

`func (o *UpdateProxyProfileSelectionRule) HasConfigNameMatchFilter() bool`

HasConfigNameMatchFilter returns a boolean if a field has been set.

### SetConfigNameMatchFilterNil

`func (o *UpdateProxyProfileSelectionRule) SetConfigNameMatchFilterNil(b bool)`

 SetConfigNameMatchFilterNil sets the value for ConfigNameMatchFilter to be an explicit nil

### UnsetConfigNameMatchFilter
`func (o *UpdateProxyProfileSelectionRule) UnsetConfigNameMatchFilter()`

UnsetConfigNameMatchFilter ensures that no value is present for ConfigNameMatchFilter, not even an explicit nil
### GetEnvironmentNameMatchFilter

`func (o *UpdateProxyProfileSelectionRule) GetEnvironmentNameMatchFilter() string`

GetEnvironmentNameMatchFilter returns the EnvironmentNameMatchFilter field if non-nil, zero value otherwise.

### GetEnvironmentNameMatchFilterOk

`func (o *UpdateProxyProfileSelectionRule) GetEnvironmentNameMatchFilterOk() (*string, bool)`

GetEnvironmentNameMatchFilterOk returns a tuple with the EnvironmentNameMatchFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentNameMatchFilter

`func (o *UpdateProxyProfileSelectionRule) SetEnvironmentNameMatchFilter(v string)`

SetEnvironmentNameMatchFilter sets EnvironmentNameMatchFilter field to given value.

### HasEnvironmentNameMatchFilter

`func (o *UpdateProxyProfileSelectionRule) HasEnvironmentNameMatchFilter() bool`

HasEnvironmentNameMatchFilter returns a boolean if a field has been set.

### SetEnvironmentNameMatchFilterNil

`func (o *UpdateProxyProfileSelectionRule) SetEnvironmentNameMatchFilterNil(b bool)`

 SetEnvironmentNameMatchFilterNil sets the value for EnvironmentNameMatchFilter to be an explicit nil

### UnsetEnvironmentNameMatchFilter
`func (o *UpdateProxyProfileSelectionRule) UnsetEnvironmentNameMatchFilter()`

UnsetEnvironmentNameMatchFilter ensures that no value is present for EnvironmentNameMatchFilter, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


