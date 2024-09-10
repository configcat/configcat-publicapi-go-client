# AuditLogItemModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuditLogId** | Pointer to **int64** |  | [optional] 
**AuditLogDateTime** | Pointer to **time.Time** |  | [optional] 
**AuditLogTypeEnum** | Pointer to **string** |  | [optional] 
**ChangeSetId** | Pointer to **NullableString** |  | [optional] 
**Truncated** | Pointer to **bool** |  | [optional] 
**AuditLogType** | Pointer to **NullableString** |  | [optional] 
**UserEmail** | Pointer to **NullableString** |  | [optional] 
**UserName** | Pointer to **NullableString** |  | [optional] 
**Where** | Pointer to **NullableString** |  | [optional] 
**Why** | Pointer to **NullableString** |  | [optional] 
**ActionTarget** | Pointer to **NullableString** |  | [optional] 
**Details** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAuditLogItemModel

`func NewAuditLogItemModel() *AuditLogItemModel`

NewAuditLogItemModel instantiates a new AuditLogItemModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditLogItemModelWithDefaults

`func NewAuditLogItemModelWithDefaults() *AuditLogItemModel`

NewAuditLogItemModelWithDefaults instantiates a new AuditLogItemModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuditLogId

`func (o *AuditLogItemModel) GetAuditLogId() int64`

GetAuditLogId returns the AuditLogId field if non-nil, zero value otherwise.

### GetAuditLogIdOk

`func (o *AuditLogItemModel) GetAuditLogIdOk() (*int64, bool)`

GetAuditLogIdOk returns a tuple with the AuditLogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditLogId

`func (o *AuditLogItemModel) SetAuditLogId(v int64)`

SetAuditLogId sets AuditLogId field to given value.

### HasAuditLogId

`func (o *AuditLogItemModel) HasAuditLogId() bool`

HasAuditLogId returns a boolean if a field has been set.

### GetAuditLogDateTime

`func (o *AuditLogItemModel) GetAuditLogDateTime() time.Time`

GetAuditLogDateTime returns the AuditLogDateTime field if non-nil, zero value otherwise.

### GetAuditLogDateTimeOk

`func (o *AuditLogItemModel) GetAuditLogDateTimeOk() (*time.Time, bool)`

GetAuditLogDateTimeOk returns a tuple with the AuditLogDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditLogDateTime

`func (o *AuditLogItemModel) SetAuditLogDateTime(v time.Time)`

SetAuditLogDateTime sets AuditLogDateTime field to given value.

### HasAuditLogDateTime

`func (o *AuditLogItemModel) HasAuditLogDateTime() bool`

HasAuditLogDateTime returns a boolean if a field has been set.

### GetAuditLogTypeEnum

`func (o *AuditLogItemModel) GetAuditLogTypeEnum() string`

GetAuditLogTypeEnum returns the AuditLogTypeEnum field if non-nil, zero value otherwise.

### GetAuditLogTypeEnumOk

`func (o *AuditLogItemModel) GetAuditLogTypeEnumOk() (*string, bool)`

GetAuditLogTypeEnumOk returns a tuple with the AuditLogTypeEnum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditLogTypeEnum

`func (o *AuditLogItemModel) SetAuditLogTypeEnum(v string)`

SetAuditLogTypeEnum sets AuditLogTypeEnum field to given value.

### HasAuditLogTypeEnum

`func (o *AuditLogItemModel) HasAuditLogTypeEnum() bool`

HasAuditLogTypeEnum returns a boolean if a field has been set.

### GetChangeSetId

`func (o *AuditLogItemModel) GetChangeSetId() string`

GetChangeSetId returns the ChangeSetId field if non-nil, zero value otherwise.

### GetChangeSetIdOk

`func (o *AuditLogItemModel) GetChangeSetIdOk() (*string, bool)`

GetChangeSetIdOk returns a tuple with the ChangeSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeSetId

`func (o *AuditLogItemModel) SetChangeSetId(v string)`

SetChangeSetId sets ChangeSetId field to given value.

### HasChangeSetId

`func (o *AuditLogItemModel) HasChangeSetId() bool`

HasChangeSetId returns a boolean if a field has been set.

### SetChangeSetIdNil

`func (o *AuditLogItemModel) SetChangeSetIdNil(b bool)`

 SetChangeSetIdNil sets the value for ChangeSetId to be an explicit nil

### UnsetChangeSetId
`func (o *AuditLogItemModel) UnsetChangeSetId()`

UnsetChangeSetId ensures that no value is present for ChangeSetId, not even an explicit nil
### GetTruncated

`func (o *AuditLogItemModel) GetTruncated() bool`

GetTruncated returns the Truncated field if non-nil, zero value otherwise.

### GetTruncatedOk

`func (o *AuditLogItemModel) GetTruncatedOk() (*bool, bool)`

GetTruncatedOk returns a tuple with the Truncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncated

`func (o *AuditLogItemModel) SetTruncated(v bool)`

SetTruncated sets Truncated field to given value.

### HasTruncated

`func (o *AuditLogItemModel) HasTruncated() bool`

HasTruncated returns a boolean if a field has been set.

### GetAuditLogType

`func (o *AuditLogItemModel) GetAuditLogType() string`

GetAuditLogType returns the AuditLogType field if non-nil, zero value otherwise.

### GetAuditLogTypeOk

`func (o *AuditLogItemModel) GetAuditLogTypeOk() (*string, bool)`

GetAuditLogTypeOk returns a tuple with the AuditLogType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditLogType

`func (o *AuditLogItemModel) SetAuditLogType(v string)`

SetAuditLogType sets AuditLogType field to given value.

### HasAuditLogType

`func (o *AuditLogItemModel) HasAuditLogType() bool`

HasAuditLogType returns a boolean if a field has been set.

### SetAuditLogTypeNil

`func (o *AuditLogItemModel) SetAuditLogTypeNil(b bool)`

 SetAuditLogTypeNil sets the value for AuditLogType to be an explicit nil

### UnsetAuditLogType
`func (o *AuditLogItemModel) UnsetAuditLogType()`

UnsetAuditLogType ensures that no value is present for AuditLogType, not even an explicit nil
### GetUserEmail

`func (o *AuditLogItemModel) GetUserEmail() string`

GetUserEmail returns the UserEmail field if non-nil, zero value otherwise.

### GetUserEmailOk

`func (o *AuditLogItemModel) GetUserEmailOk() (*string, bool)`

GetUserEmailOk returns a tuple with the UserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEmail

`func (o *AuditLogItemModel) SetUserEmail(v string)`

SetUserEmail sets UserEmail field to given value.

### HasUserEmail

`func (o *AuditLogItemModel) HasUserEmail() bool`

HasUserEmail returns a boolean if a field has been set.

### SetUserEmailNil

`func (o *AuditLogItemModel) SetUserEmailNil(b bool)`

 SetUserEmailNil sets the value for UserEmail to be an explicit nil

### UnsetUserEmail
`func (o *AuditLogItemModel) UnsetUserEmail()`

UnsetUserEmail ensures that no value is present for UserEmail, not even an explicit nil
### GetUserName

`func (o *AuditLogItemModel) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *AuditLogItemModel) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *AuditLogItemModel) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *AuditLogItemModel) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### SetUserNameNil

`func (o *AuditLogItemModel) SetUserNameNil(b bool)`

 SetUserNameNil sets the value for UserName to be an explicit nil

### UnsetUserName
`func (o *AuditLogItemModel) UnsetUserName()`

UnsetUserName ensures that no value is present for UserName, not even an explicit nil
### GetWhere

`func (o *AuditLogItemModel) GetWhere() string`

GetWhere returns the Where field if non-nil, zero value otherwise.

### GetWhereOk

`func (o *AuditLogItemModel) GetWhereOk() (*string, bool)`

GetWhereOk returns a tuple with the Where field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhere

`func (o *AuditLogItemModel) SetWhere(v string)`

SetWhere sets Where field to given value.

### HasWhere

`func (o *AuditLogItemModel) HasWhere() bool`

HasWhere returns a boolean if a field has been set.

### SetWhereNil

`func (o *AuditLogItemModel) SetWhereNil(b bool)`

 SetWhereNil sets the value for Where to be an explicit nil

### UnsetWhere
`func (o *AuditLogItemModel) UnsetWhere()`

UnsetWhere ensures that no value is present for Where, not even an explicit nil
### GetWhy

`func (o *AuditLogItemModel) GetWhy() string`

GetWhy returns the Why field if non-nil, zero value otherwise.

### GetWhyOk

`func (o *AuditLogItemModel) GetWhyOk() (*string, bool)`

GetWhyOk returns a tuple with the Why field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhy

`func (o *AuditLogItemModel) SetWhy(v string)`

SetWhy sets Why field to given value.

### HasWhy

`func (o *AuditLogItemModel) HasWhy() bool`

HasWhy returns a boolean if a field has been set.

### SetWhyNil

`func (o *AuditLogItemModel) SetWhyNil(b bool)`

 SetWhyNil sets the value for Why to be an explicit nil

### UnsetWhy
`func (o *AuditLogItemModel) UnsetWhy()`

UnsetWhy ensures that no value is present for Why, not even an explicit nil
### GetActionTarget

`func (o *AuditLogItemModel) GetActionTarget() string`

GetActionTarget returns the ActionTarget field if non-nil, zero value otherwise.

### GetActionTargetOk

`func (o *AuditLogItemModel) GetActionTargetOk() (*string, bool)`

GetActionTargetOk returns a tuple with the ActionTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionTarget

`func (o *AuditLogItemModel) SetActionTarget(v string)`

SetActionTarget sets ActionTarget field to given value.

### HasActionTarget

`func (o *AuditLogItemModel) HasActionTarget() bool`

HasActionTarget returns a boolean if a field has been set.

### SetActionTargetNil

`func (o *AuditLogItemModel) SetActionTargetNil(b bool)`

 SetActionTargetNil sets the value for ActionTarget to be an explicit nil

### UnsetActionTarget
`func (o *AuditLogItemModel) UnsetActionTarget()`

UnsetActionTarget ensures that no value is present for ActionTarget, not even an explicit nil
### GetDetails

`func (o *AuditLogItemModel) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *AuditLogItemModel) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *AuditLogItemModel) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *AuditLogItemModel) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### SetDetailsNil

`func (o *AuditLogItemModel) SetDetailsNil(b bool)`

 SetDetailsNil sets the value for Details to be an explicit nil

### UnsetDetails
`func (o *AuditLogItemModel) UnsetDetails()`

UnsetDetails ensures that no value is present for Details, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


