# ChangeRequestIssueModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChangeRequestIssueId** | **int64** | Identifier of the Change Request issue. | 
**IssueType** | [**ChangeRequestIssueType**](ChangeRequestIssueType.md) |  | 
**IssueDetails** | **string** | Description of the issue. | 
**IssueDetectedAt** | **time.Time** | The UTC date and time when the issue was detected. | 
**IssueFixedAt** | **NullableTime** | Optional UTC date and time when the issue was fixed. | 

## Methods

### NewChangeRequestIssueModel

`func NewChangeRequestIssueModel(changeRequestIssueId int64, issueType ChangeRequestIssueType, issueDetails string, issueDetectedAt time.Time, issueFixedAt NullableTime, ) *ChangeRequestIssueModel`

NewChangeRequestIssueModel instantiates a new ChangeRequestIssueModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangeRequestIssueModelWithDefaults

`func NewChangeRequestIssueModelWithDefaults() *ChangeRequestIssueModel`

NewChangeRequestIssueModelWithDefaults instantiates a new ChangeRequestIssueModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChangeRequestIssueId

`func (o *ChangeRequestIssueModel) GetChangeRequestIssueId() int64`

GetChangeRequestIssueId returns the ChangeRequestIssueId field if non-nil, zero value otherwise.

### GetChangeRequestIssueIdOk

`func (o *ChangeRequestIssueModel) GetChangeRequestIssueIdOk() (*int64, bool)`

GetChangeRequestIssueIdOk returns a tuple with the ChangeRequestIssueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeRequestIssueId

`func (o *ChangeRequestIssueModel) SetChangeRequestIssueId(v int64)`

SetChangeRequestIssueId sets ChangeRequestIssueId field to given value.


### GetIssueType

`func (o *ChangeRequestIssueModel) GetIssueType() ChangeRequestIssueType`

GetIssueType returns the IssueType field if non-nil, zero value otherwise.

### GetIssueTypeOk

`func (o *ChangeRequestIssueModel) GetIssueTypeOk() (*ChangeRequestIssueType, bool)`

GetIssueTypeOk returns a tuple with the IssueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueType

`func (o *ChangeRequestIssueModel) SetIssueType(v ChangeRequestIssueType)`

SetIssueType sets IssueType field to given value.


### GetIssueDetails

`func (o *ChangeRequestIssueModel) GetIssueDetails() string`

GetIssueDetails returns the IssueDetails field if non-nil, zero value otherwise.

### GetIssueDetailsOk

`func (o *ChangeRequestIssueModel) GetIssueDetailsOk() (*string, bool)`

GetIssueDetailsOk returns a tuple with the IssueDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDetails

`func (o *ChangeRequestIssueModel) SetIssueDetails(v string)`

SetIssueDetails sets IssueDetails field to given value.


### GetIssueDetectedAt

`func (o *ChangeRequestIssueModel) GetIssueDetectedAt() time.Time`

GetIssueDetectedAt returns the IssueDetectedAt field if non-nil, zero value otherwise.

### GetIssueDetectedAtOk

`func (o *ChangeRequestIssueModel) GetIssueDetectedAtOk() (*time.Time, bool)`

GetIssueDetectedAtOk returns a tuple with the IssueDetectedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDetectedAt

`func (o *ChangeRequestIssueModel) SetIssueDetectedAt(v time.Time)`

SetIssueDetectedAt sets IssueDetectedAt field to given value.


### GetIssueFixedAt

`func (o *ChangeRequestIssueModel) GetIssueFixedAt() time.Time`

GetIssueFixedAt returns the IssueFixedAt field if non-nil, zero value otherwise.

### GetIssueFixedAtOk

`func (o *ChangeRequestIssueModel) GetIssueFixedAtOk() (*time.Time, bool)`

GetIssueFixedAtOk returns a tuple with the IssueFixedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueFixedAt

`func (o *ChangeRequestIssueModel) SetIssueFixedAt(v time.Time)`

SetIssueFixedAt sets IssueFixedAt field to given value.


### SetIssueFixedAtNil

`func (o *ChangeRequestIssueModel) SetIssueFixedAtNil(b bool)`

 SetIssueFixedAtNil sets the value for IssueFixedAt to be an explicit nil

### UnsetIssueFixedAt
`func (o *ChangeRequestIssueModel) UnsetIssueFixedAt()`

UnsetIssueFixedAt ensures that no value is present for IssueFixedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


