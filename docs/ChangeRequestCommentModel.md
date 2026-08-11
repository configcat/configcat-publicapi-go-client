# ChangeRequestCommentModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChangeRequestCommentId** | **int64** | Identifier of the Change Request comment. | 
**Body** | **string** | Body content of the comment. | 
**CreatedAt** | **time.Time** | The UTC date and time when the comment was created. | 
**UserEmail** | **string** | Email of the user who created the comment. | 
**UserFullName** | **string** | Full name of the user who created the comment. | 
**EditedAt** | **NullableTime** | Optional UTC date and time when the comment was last edited. | 

## Methods

### NewChangeRequestCommentModel

`func NewChangeRequestCommentModel(changeRequestCommentId int64, body string, createdAt time.Time, userEmail string, userFullName string, editedAt NullableTime, ) *ChangeRequestCommentModel`

NewChangeRequestCommentModel instantiates a new ChangeRequestCommentModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangeRequestCommentModelWithDefaults

`func NewChangeRequestCommentModelWithDefaults() *ChangeRequestCommentModel`

NewChangeRequestCommentModelWithDefaults instantiates a new ChangeRequestCommentModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChangeRequestCommentId

`func (o *ChangeRequestCommentModel) GetChangeRequestCommentId() int64`

GetChangeRequestCommentId returns the ChangeRequestCommentId field if non-nil, zero value otherwise.

### GetChangeRequestCommentIdOk

`func (o *ChangeRequestCommentModel) GetChangeRequestCommentIdOk() (*int64, bool)`

GetChangeRequestCommentIdOk returns a tuple with the ChangeRequestCommentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeRequestCommentId

`func (o *ChangeRequestCommentModel) SetChangeRequestCommentId(v int64)`

SetChangeRequestCommentId sets ChangeRequestCommentId field to given value.


### GetBody

`func (o *ChangeRequestCommentModel) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *ChangeRequestCommentModel) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *ChangeRequestCommentModel) SetBody(v string)`

SetBody sets Body field to given value.


### GetCreatedAt

`func (o *ChangeRequestCommentModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ChangeRequestCommentModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ChangeRequestCommentModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUserEmail

`func (o *ChangeRequestCommentModel) GetUserEmail() string`

GetUserEmail returns the UserEmail field if non-nil, zero value otherwise.

### GetUserEmailOk

`func (o *ChangeRequestCommentModel) GetUserEmailOk() (*string, bool)`

GetUserEmailOk returns a tuple with the UserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEmail

`func (o *ChangeRequestCommentModel) SetUserEmail(v string)`

SetUserEmail sets UserEmail field to given value.


### GetUserFullName

`func (o *ChangeRequestCommentModel) GetUserFullName() string`

GetUserFullName returns the UserFullName field if non-nil, zero value otherwise.

### GetUserFullNameOk

`func (o *ChangeRequestCommentModel) GetUserFullNameOk() (*string, bool)`

GetUserFullNameOk returns a tuple with the UserFullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserFullName

`func (o *ChangeRequestCommentModel) SetUserFullName(v string)`

SetUserFullName sets UserFullName field to given value.


### GetEditedAt

`func (o *ChangeRequestCommentModel) GetEditedAt() time.Time`

GetEditedAt returns the EditedAt field if non-nil, zero value otherwise.

### GetEditedAtOk

`func (o *ChangeRequestCommentModel) GetEditedAtOk() (*time.Time, bool)`

GetEditedAtOk returns a tuple with the EditedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditedAt

`func (o *ChangeRequestCommentModel) SetEditedAt(v time.Time)`

SetEditedAt sets EditedAt field to given value.


### SetEditedAtNil

`func (o *ChangeRequestCommentModel) SetEditedAtNil(b bool)`

 SetEditedAtNil sets the value for EditedAt to be an explicit nil

### UnsetEditedAt
`func (o *ChangeRequestCommentModel) UnsetEditedAt()`

UnsetEditedAt ensures that no value is present for EditedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


