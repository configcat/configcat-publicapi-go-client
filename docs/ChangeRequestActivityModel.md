# ChangeRequestActivityModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChangeRequestActivityId** | **int64** | Identifier of the Change Request activity. | 
**ChangeRequestActivityType** | [**ChangeRequestActivityType**](ChangeRequestActivityType.md) |  | 
**Date** | **time.Time** | The UTC date and time when the activity occurred. | 
**UserId** | **NullableString** | Identifier of the user who triggered the activity. | 
**UserFullName** | **string** | Full name of the user who triggered the activity. | 
**UserEmail** | **NullableString** | Email of the user who triggered the activity. | 
**Details** | **string** | Detailed description of the activity. | 
**Error** | **NullableString** | Optional error message if the activity failed. | 

## Methods

### NewChangeRequestActivityModel

`func NewChangeRequestActivityModel(changeRequestActivityId int64, changeRequestActivityType ChangeRequestActivityType, date time.Time, userId NullableString, userFullName string, userEmail NullableString, details string, error_ NullableString, ) *ChangeRequestActivityModel`

NewChangeRequestActivityModel instantiates a new ChangeRequestActivityModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangeRequestActivityModelWithDefaults

`func NewChangeRequestActivityModelWithDefaults() *ChangeRequestActivityModel`

NewChangeRequestActivityModelWithDefaults instantiates a new ChangeRequestActivityModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChangeRequestActivityId

`func (o *ChangeRequestActivityModel) GetChangeRequestActivityId() int64`

GetChangeRequestActivityId returns the ChangeRequestActivityId field if non-nil, zero value otherwise.

### GetChangeRequestActivityIdOk

`func (o *ChangeRequestActivityModel) GetChangeRequestActivityIdOk() (*int64, bool)`

GetChangeRequestActivityIdOk returns a tuple with the ChangeRequestActivityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeRequestActivityId

`func (o *ChangeRequestActivityModel) SetChangeRequestActivityId(v int64)`

SetChangeRequestActivityId sets ChangeRequestActivityId field to given value.


### GetChangeRequestActivityType

`func (o *ChangeRequestActivityModel) GetChangeRequestActivityType() ChangeRequestActivityType`

GetChangeRequestActivityType returns the ChangeRequestActivityType field if non-nil, zero value otherwise.

### GetChangeRequestActivityTypeOk

`func (o *ChangeRequestActivityModel) GetChangeRequestActivityTypeOk() (*ChangeRequestActivityType, bool)`

GetChangeRequestActivityTypeOk returns a tuple with the ChangeRequestActivityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeRequestActivityType

`func (o *ChangeRequestActivityModel) SetChangeRequestActivityType(v ChangeRequestActivityType)`

SetChangeRequestActivityType sets ChangeRequestActivityType field to given value.


### GetDate

`func (o *ChangeRequestActivityModel) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *ChangeRequestActivityModel) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *ChangeRequestActivityModel) SetDate(v time.Time)`

SetDate sets Date field to given value.


### GetUserId

`func (o *ChangeRequestActivityModel) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ChangeRequestActivityModel) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ChangeRequestActivityModel) SetUserId(v string)`

SetUserId sets UserId field to given value.


### SetUserIdNil

`func (o *ChangeRequestActivityModel) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *ChangeRequestActivityModel) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetUserFullName

`func (o *ChangeRequestActivityModel) GetUserFullName() string`

GetUserFullName returns the UserFullName field if non-nil, zero value otherwise.

### GetUserFullNameOk

`func (o *ChangeRequestActivityModel) GetUserFullNameOk() (*string, bool)`

GetUserFullNameOk returns a tuple with the UserFullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserFullName

`func (o *ChangeRequestActivityModel) SetUserFullName(v string)`

SetUserFullName sets UserFullName field to given value.


### GetUserEmail

`func (o *ChangeRequestActivityModel) GetUserEmail() string`

GetUserEmail returns the UserEmail field if non-nil, zero value otherwise.

### GetUserEmailOk

`func (o *ChangeRequestActivityModel) GetUserEmailOk() (*string, bool)`

GetUserEmailOk returns a tuple with the UserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEmail

`func (o *ChangeRequestActivityModel) SetUserEmail(v string)`

SetUserEmail sets UserEmail field to given value.


### SetUserEmailNil

`func (o *ChangeRequestActivityModel) SetUserEmailNil(b bool)`

 SetUserEmailNil sets the value for UserEmail to be an explicit nil

### UnsetUserEmail
`func (o *ChangeRequestActivityModel) UnsetUserEmail()`

UnsetUserEmail ensures that no value is present for UserEmail, not even an explicit nil
### GetDetails

`func (o *ChangeRequestActivityModel) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ChangeRequestActivityModel) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ChangeRequestActivityModel) SetDetails(v string)`

SetDetails sets Details field to given value.


### GetError

`func (o *ChangeRequestActivityModel) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ChangeRequestActivityModel) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ChangeRequestActivityModel) SetError(v string)`

SetError sets Error field to given value.


### SetErrorNil

`func (o *ChangeRequestActivityModel) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *ChangeRequestActivityModel) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


