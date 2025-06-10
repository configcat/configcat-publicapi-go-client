# SegmentModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Product** | [**ProductModel**](ProductModel.md) |  | 
**SegmentId** | **string** | Identifier of the Segment. | 
**Name** | **string** | Name of the Segment. | 
**Description** | **NullableString** | Description of the Segment. | 
**CreatorEmail** | **NullableString** | The email of the user who created the Segment. | 
**CreatorFullName** | **NullableString** | The name of the user who created the Segment. | 
**CreatedAt** | **time.Time** | The date and time when the Segment was created. | 
**LastUpdaterEmail** | **NullableString** | The email of the user who last updated the Segment. | 
**LastUpdaterFullName** | **NullableString** | The name of the user who last updated the Segment. | 
**UpdatedAt** | **time.Time** | The date and time when the Segment was last updated. | 
**ComparisonAttribute** | **string** | The user&#39;s attribute the evaluation process must take into account. | 
**Comparator** | [**RolloutRuleComparator**](RolloutRuleComparator.md) |  | 
**ComparisonValue** | **string** | The value to compare with the given user attribute&#39;s value. | 

## Methods

### NewSegmentModel

`func NewSegmentModel(product ProductModel, segmentId string, name string, description NullableString, creatorEmail NullableString, creatorFullName NullableString, createdAt time.Time, lastUpdaterEmail NullableString, lastUpdaterFullName NullableString, updatedAt time.Time, comparisonAttribute string, comparator RolloutRuleComparator, comparisonValue string, ) *SegmentModel`

NewSegmentModel instantiates a new SegmentModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSegmentModelWithDefaults

`func NewSegmentModelWithDefaults() *SegmentModel`

NewSegmentModelWithDefaults instantiates a new SegmentModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProduct

`func (o *SegmentModel) GetProduct() ProductModel`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *SegmentModel) GetProductOk() (*ProductModel, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *SegmentModel) SetProduct(v ProductModel)`

SetProduct sets Product field to given value.


### GetSegmentId

`func (o *SegmentModel) GetSegmentId() string`

GetSegmentId returns the SegmentId field if non-nil, zero value otherwise.

### GetSegmentIdOk

`func (o *SegmentModel) GetSegmentIdOk() (*string, bool)`

GetSegmentIdOk returns a tuple with the SegmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentId

`func (o *SegmentModel) SetSegmentId(v string)`

SetSegmentId sets SegmentId field to given value.


### GetName

`func (o *SegmentModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SegmentModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SegmentModel) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *SegmentModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SegmentModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SegmentModel) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *SegmentModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SegmentModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCreatorEmail

`func (o *SegmentModel) GetCreatorEmail() string`

GetCreatorEmail returns the CreatorEmail field if non-nil, zero value otherwise.

### GetCreatorEmailOk

`func (o *SegmentModel) GetCreatorEmailOk() (*string, bool)`

GetCreatorEmailOk returns a tuple with the CreatorEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorEmail

`func (o *SegmentModel) SetCreatorEmail(v string)`

SetCreatorEmail sets CreatorEmail field to given value.


### SetCreatorEmailNil

`func (o *SegmentModel) SetCreatorEmailNil(b bool)`

 SetCreatorEmailNil sets the value for CreatorEmail to be an explicit nil

### UnsetCreatorEmail
`func (o *SegmentModel) UnsetCreatorEmail()`

UnsetCreatorEmail ensures that no value is present for CreatorEmail, not even an explicit nil
### GetCreatorFullName

`func (o *SegmentModel) GetCreatorFullName() string`

GetCreatorFullName returns the CreatorFullName field if non-nil, zero value otherwise.

### GetCreatorFullNameOk

`func (o *SegmentModel) GetCreatorFullNameOk() (*string, bool)`

GetCreatorFullNameOk returns a tuple with the CreatorFullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorFullName

`func (o *SegmentModel) SetCreatorFullName(v string)`

SetCreatorFullName sets CreatorFullName field to given value.


### SetCreatorFullNameNil

`func (o *SegmentModel) SetCreatorFullNameNil(b bool)`

 SetCreatorFullNameNil sets the value for CreatorFullName to be an explicit nil

### UnsetCreatorFullName
`func (o *SegmentModel) UnsetCreatorFullName()`

UnsetCreatorFullName ensures that no value is present for CreatorFullName, not even an explicit nil
### GetCreatedAt

`func (o *SegmentModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SegmentModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SegmentModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetLastUpdaterEmail

`func (o *SegmentModel) GetLastUpdaterEmail() string`

GetLastUpdaterEmail returns the LastUpdaterEmail field if non-nil, zero value otherwise.

### GetLastUpdaterEmailOk

`func (o *SegmentModel) GetLastUpdaterEmailOk() (*string, bool)`

GetLastUpdaterEmailOk returns a tuple with the LastUpdaterEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdaterEmail

`func (o *SegmentModel) SetLastUpdaterEmail(v string)`

SetLastUpdaterEmail sets LastUpdaterEmail field to given value.


### SetLastUpdaterEmailNil

`func (o *SegmentModel) SetLastUpdaterEmailNil(b bool)`

 SetLastUpdaterEmailNil sets the value for LastUpdaterEmail to be an explicit nil

### UnsetLastUpdaterEmail
`func (o *SegmentModel) UnsetLastUpdaterEmail()`

UnsetLastUpdaterEmail ensures that no value is present for LastUpdaterEmail, not even an explicit nil
### GetLastUpdaterFullName

`func (o *SegmentModel) GetLastUpdaterFullName() string`

GetLastUpdaterFullName returns the LastUpdaterFullName field if non-nil, zero value otherwise.

### GetLastUpdaterFullNameOk

`func (o *SegmentModel) GetLastUpdaterFullNameOk() (*string, bool)`

GetLastUpdaterFullNameOk returns a tuple with the LastUpdaterFullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdaterFullName

`func (o *SegmentModel) SetLastUpdaterFullName(v string)`

SetLastUpdaterFullName sets LastUpdaterFullName field to given value.


### SetLastUpdaterFullNameNil

`func (o *SegmentModel) SetLastUpdaterFullNameNil(b bool)`

 SetLastUpdaterFullNameNil sets the value for LastUpdaterFullName to be an explicit nil

### UnsetLastUpdaterFullName
`func (o *SegmentModel) UnsetLastUpdaterFullName()`

UnsetLastUpdaterFullName ensures that no value is present for LastUpdaterFullName, not even an explicit nil
### GetUpdatedAt

`func (o *SegmentModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SegmentModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SegmentModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetComparisonAttribute

`func (o *SegmentModel) GetComparisonAttribute() string`

GetComparisonAttribute returns the ComparisonAttribute field if non-nil, zero value otherwise.

### GetComparisonAttributeOk

`func (o *SegmentModel) GetComparisonAttributeOk() (*string, bool)`

GetComparisonAttributeOk returns a tuple with the ComparisonAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparisonAttribute

`func (o *SegmentModel) SetComparisonAttribute(v string)`

SetComparisonAttribute sets ComparisonAttribute field to given value.


### GetComparator

`func (o *SegmentModel) GetComparator() RolloutRuleComparator`

GetComparator returns the Comparator field if non-nil, zero value otherwise.

### GetComparatorOk

`func (o *SegmentModel) GetComparatorOk() (*RolloutRuleComparator, bool)`

GetComparatorOk returns a tuple with the Comparator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparator

`func (o *SegmentModel) SetComparator(v RolloutRuleComparator)`

SetComparator sets Comparator field to given value.


### GetComparisonValue

`func (o *SegmentModel) GetComparisonValue() string`

GetComparisonValue returns the ComparisonValue field if non-nil, zero value otherwise.

### GetComparisonValueOk

`func (o *SegmentModel) GetComparisonValueOk() (*string, bool)`

GetComparisonValueOk returns a tuple with the ComparisonValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparisonValue

`func (o *SegmentModel) SetComparisonValue(v string)`

SetComparisonValue sets ComparisonValue field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


