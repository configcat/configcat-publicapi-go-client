# ProxyProfileSelectionRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | [**SelectionRuleKind**](SelectionRuleKind.md) |  | 
**ProductIdFilter** | **NullableString** | Defines the filter for matching Products by their unique identifier. | 
**ConfigIdFilter** | **NullableString** | Defines the filter for matching Configs by their unique identifier. | 
**EnvironmentIdFilter** | **NullableString** | Defines the filter for matching Environments by their unique identifier. | 
**ProductNameMatchFilter** | **NullableString** | Specifies a filter to match Product names in the proxy profile selection rule. It accepts wildcards (*). | 
**ConfigNameMatchFilter** | **NullableString** | Specifies a filter to match Config names in the proxy profile selection rule. It accepts wildcards (*). | 
**EnvironmentNameMatchFilter** | **NullableString** | Specifies a filter to match Environment names in the proxy profile selection rule. It accepts wildcards (*). | 

## Methods

### NewProxyProfileSelectionRule

`func NewProxyProfileSelectionRule(kind SelectionRuleKind, productIdFilter NullableString, configIdFilter NullableString, environmentIdFilter NullableString, productNameMatchFilter NullableString, configNameMatchFilter NullableString, environmentNameMatchFilter NullableString, ) *ProxyProfileSelectionRule`

NewProxyProfileSelectionRule instantiates a new ProxyProfileSelectionRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProxyProfileSelectionRuleWithDefaults

`func NewProxyProfileSelectionRuleWithDefaults() *ProxyProfileSelectionRule`

NewProxyProfileSelectionRuleWithDefaults instantiates a new ProxyProfileSelectionRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *ProxyProfileSelectionRule) GetKind() SelectionRuleKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ProxyProfileSelectionRule) GetKindOk() (*SelectionRuleKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ProxyProfileSelectionRule) SetKind(v SelectionRuleKind)`

SetKind sets Kind field to given value.


### GetProductIdFilter

`func (o *ProxyProfileSelectionRule) GetProductIdFilter() string`

GetProductIdFilter returns the ProductIdFilter field if non-nil, zero value otherwise.

### GetProductIdFilterOk

`func (o *ProxyProfileSelectionRule) GetProductIdFilterOk() (*string, bool)`

GetProductIdFilterOk returns a tuple with the ProductIdFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductIdFilter

`func (o *ProxyProfileSelectionRule) SetProductIdFilter(v string)`

SetProductIdFilter sets ProductIdFilter field to given value.


### SetProductIdFilterNil

`func (o *ProxyProfileSelectionRule) SetProductIdFilterNil(b bool)`

 SetProductIdFilterNil sets the value for ProductIdFilter to be an explicit nil

### UnsetProductIdFilter
`func (o *ProxyProfileSelectionRule) UnsetProductIdFilter()`

UnsetProductIdFilter ensures that no value is present for ProductIdFilter, not even an explicit nil
### GetConfigIdFilter

`func (o *ProxyProfileSelectionRule) GetConfigIdFilter() string`

GetConfigIdFilter returns the ConfigIdFilter field if non-nil, zero value otherwise.

### GetConfigIdFilterOk

`func (o *ProxyProfileSelectionRule) GetConfigIdFilterOk() (*string, bool)`

GetConfigIdFilterOk returns a tuple with the ConfigIdFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigIdFilter

`func (o *ProxyProfileSelectionRule) SetConfigIdFilter(v string)`

SetConfigIdFilter sets ConfigIdFilter field to given value.


### SetConfigIdFilterNil

`func (o *ProxyProfileSelectionRule) SetConfigIdFilterNil(b bool)`

 SetConfigIdFilterNil sets the value for ConfigIdFilter to be an explicit nil

### UnsetConfigIdFilter
`func (o *ProxyProfileSelectionRule) UnsetConfigIdFilter()`

UnsetConfigIdFilter ensures that no value is present for ConfigIdFilter, not even an explicit nil
### GetEnvironmentIdFilter

`func (o *ProxyProfileSelectionRule) GetEnvironmentIdFilter() string`

GetEnvironmentIdFilter returns the EnvironmentIdFilter field if non-nil, zero value otherwise.

### GetEnvironmentIdFilterOk

`func (o *ProxyProfileSelectionRule) GetEnvironmentIdFilterOk() (*string, bool)`

GetEnvironmentIdFilterOk returns a tuple with the EnvironmentIdFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentIdFilter

`func (o *ProxyProfileSelectionRule) SetEnvironmentIdFilter(v string)`

SetEnvironmentIdFilter sets EnvironmentIdFilter field to given value.


### SetEnvironmentIdFilterNil

`func (o *ProxyProfileSelectionRule) SetEnvironmentIdFilterNil(b bool)`

 SetEnvironmentIdFilterNil sets the value for EnvironmentIdFilter to be an explicit nil

### UnsetEnvironmentIdFilter
`func (o *ProxyProfileSelectionRule) UnsetEnvironmentIdFilter()`

UnsetEnvironmentIdFilter ensures that no value is present for EnvironmentIdFilter, not even an explicit nil
### GetProductNameMatchFilter

`func (o *ProxyProfileSelectionRule) GetProductNameMatchFilter() string`

GetProductNameMatchFilter returns the ProductNameMatchFilter field if non-nil, zero value otherwise.

### GetProductNameMatchFilterOk

`func (o *ProxyProfileSelectionRule) GetProductNameMatchFilterOk() (*string, bool)`

GetProductNameMatchFilterOk returns a tuple with the ProductNameMatchFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductNameMatchFilter

`func (o *ProxyProfileSelectionRule) SetProductNameMatchFilter(v string)`

SetProductNameMatchFilter sets ProductNameMatchFilter field to given value.


### SetProductNameMatchFilterNil

`func (o *ProxyProfileSelectionRule) SetProductNameMatchFilterNil(b bool)`

 SetProductNameMatchFilterNil sets the value for ProductNameMatchFilter to be an explicit nil

### UnsetProductNameMatchFilter
`func (o *ProxyProfileSelectionRule) UnsetProductNameMatchFilter()`

UnsetProductNameMatchFilter ensures that no value is present for ProductNameMatchFilter, not even an explicit nil
### GetConfigNameMatchFilter

`func (o *ProxyProfileSelectionRule) GetConfigNameMatchFilter() string`

GetConfigNameMatchFilter returns the ConfigNameMatchFilter field if non-nil, zero value otherwise.

### GetConfigNameMatchFilterOk

`func (o *ProxyProfileSelectionRule) GetConfigNameMatchFilterOk() (*string, bool)`

GetConfigNameMatchFilterOk returns a tuple with the ConfigNameMatchFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigNameMatchFilter

`func (o *ProxyProfileSelectionRule) SetConfigNameMatchFilter(v string)`

SetConfigNameMatchFilter sets ConfigNameMatchFilter field to given value.


### SetConfigNameMatchFilterNil

`func (o *ProxyProfileSelectionRule) SetConfigNameMatchFilterNil(b bool)`

 SetConfigNameMatchFilterNil sets the value for ConfigNameMatchFilter to be an explicit nil

### UnsetConfigNameMatchFilter
`func (o *ProxyProfileSelectionRule) UnsetConfigNameMatchFilter()`

UnsetConfigNameMatchFilter ensures that no value is present for ConfigNameMatchFilter, not even an explicit nil
### GetEnvironmentNameMatchFilter

`func (o *ProxyProfileSelectionRule) GetEnvironmentNameMatchFilter() string`

GetEnvironmentNameMatchFilter returns the EnvironmentNameMatchFilter field if non-nil, zero value otherwise.

### GetEnvironmentNameMatchFilterOk

`func (o *ProxyProfileSelectionRule) GetEnvironmentNameMatchFilterOk() (*string, bool)`

GetEnvironmentNameMatchFilterOk returns a tuple with the EnvironmentNameMatchFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentNameMatchFilter

`func (o *ProxyProfileSelectionRule) SetEnvironmentNameMatchFilter(v string)`

SetEnvironmentNameMatchFilter sets EnvironmentNameMatchFilter field to given value.


### SetEnvironmentNameMatchFilterNil

`func (o *ProxyProfileSelectionRule) SetEnvironmentNameMatchFilterNil(b bool)`

 SetEnvironmentNameMatchFilterNil sets the value for EnvironmentNameMatchFilter to be an explicit nil

### UnsetEnvironmentNameMatchFilter
`func (o *ProxyProfileSelectionRule) UnsetEnvironmentNameMatchFilter()`

UnsetEnvironmentNameMatchFilter ensures that no value is present for EnvironmentNameMatchFilter, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


