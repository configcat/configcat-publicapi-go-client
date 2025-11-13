# ProxyProfileModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProxyProfileId** | **string** | The unique identifier of the proxy profile. | 
**Name** | **string** | The name of the proxy profile. | 
**Description** | **NullableString** | The description of the proxy profile. | 
**LastAccessedAt** | **NullableTime** | The date and time when the proxy profile was last accessed. | 
**ConnectionPreferences** | [**ConnectionPreferences**](ConnectionPreferences.md) |  | 
**SdkKeySelectionRules** | [**[]ProxyProfileSelectionRule**](ProxyProfileSelectionRule.md) |  | 

## Methods

### NewProxyProfileModel

`func NewProxyProfileModel(proxyProfileId string, name string, description NullableString, lastAccessedAt NullableTime, connectionPreferences ConnectionPreferences, sdkKeySelectionRules []ProxyProfileSelectionRule, ) *ProxyProfileModel`

NewProxyProfileModel instantiates a new ProxyProfileModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProxyProfileModelWithDefaults

`func NewProxyProfileModelWithDefaults() *ProxyProfileModel`

NewProxyProfileModelWithDefaults instantiates a new ProxyProfileModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProxyProfileId

`func (o *ProxyProfileModel) GetProxyProfileId() string`

GetProxyProfileId returns the ProxyProfileId field if non-nil, zero value otherwise.

### GetProxyProfileIdOk

`func (o *ProxyProfileModel) GetProxyProfileIdOk() (*string, bool)`

GetProxyProfileIdOk returns a tuple with the ProxyProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyProfileId

`func (o *ProxyProfileModel) SetProxyProfileId(v string)`

SetProxyProfileId sets ProxyProfileId field to given value.


### GetName

`func (o *ProxyProfileModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProxyProfileModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProxyProfileModel) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ProxyProfileModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProxyProfileModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProxyProfileModel) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *ProxyProfileModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ProxyProfileModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLastAccessedAt

`func (o *ProxyProfileModel) GetLastAccessedAt() time.Time`

GetLastAccessedAt returns the LastAccessedAt field if non-nil, zero value otherwise.

### GetLastAccessedAtOk

`func (o *ProxyProfileModel) GetLastAccessedAtOk() (*time.Time, bool)`

GetLastAccessedAtOk returns a tuple with the LastAccessedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAccessedAt

`func (o *ProxyProfileModel) SetLastAccessedAt(v time.Time)`

SetLastAccessedAt sets LastAccessedAt field to given value.


### SetLastAccessedAtNil

`func (o *ProxyProfileModel) SetLastAccessedAtNil(b bool)`

 SetLastAccessedAtNil sets the value for LastAccessedAt to be an explicit nil

### UnsetLastAccessedAt
`func (o *ProxyProfileModel) UnsetLastAccessedAt()`

UnsetLastAccessedAt ensures that no value is present for LastAccessedAt, not even an explicit nil
### GetConnectionPreferences

`func (o *ProxyProfileModel) GetConnectionPreferences() ConnectionPreferences`

GetConnectionPreferences returns the ConnectionPreferences field if non-nil, zero value otherwise.

### GetConnectionPreferencesOk

`func (o *ProxyProfileModel) GetConnectionPreferencesOk() (*ConnectionPreferences, bool)`

GetConnectionPreferencesOk returns a tuple with the ConnectionPreferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionPreferences

`func (o *ProxyProfileModel) SetConnectionPreferences(v ConnectionPreferences)`

SetConnectionPreferences sets ConnectionPreferences field to given value.


### GetSdkKeySelectionRules

`func (o *ProxyProfileModel) GetSdkKeySelectionRules() []ProxyProfileSelectionRule`

GetSdkKeySelectionRules returns the SdkKeySelectionRules field if non-nil, zero value otherwise.

### GetSdkKeySelectionRulesOk

`func (o *ProxyProfileModel) GetSdkKeySelectionRulesOk() (*[]ProxyProfileSelectionRule, bool)`

GetSdkKeySelectionRulesOk returns a tuple with the SdkKeySelectionRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdkKeySelectionRules

`func (o *ProxyProfileModel) SetSdkKeySelectionRules(v []ProxyProfileSelectionRule)`

SetSdkKeySelectionRules sets SdkKeySelectionRules field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


