# MeModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **NullableString** |  | [optional] [readonly] 
**FullName** | Pointer to **NullableString** |  | [optional] [readonly] 

## Methods

### NewMeModel

`func NewMeModel() *MeModel`

NewMeModel instantiates a new MeModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeModelWithDefaults

`func NewMeModelWithDefaults() *MeModel`

NewMeModelWithDefaults instantiates a new MeModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *MeModel) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *MeModel) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *MeModel) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *MeModel) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *MeModel) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *MeModel) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetFullName

`func (o *MeModel) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *MeModel) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *MeModel) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *MeModel) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### SetFullNameNil

`func (o *MeModel) SetFullNameNil(b bool)`

 SetFullNameNil sets the value for FullName to be an explicit nil

### UnsetFullName
`func (o *MeModel) UnsetFullName()`

UnsetFullName ensures that no value is present for FullName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


