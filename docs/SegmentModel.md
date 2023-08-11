# SegmentModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Product** | Pointer to [**ProductModel**](ProductModel.md) |  | [optional] 
**SegmentId** | Pointer to **string** | Identifier of the Segment. | [optional] 
**Name** | Pointer to **NullableString** | Name of the Segment. | [optional] 
**Description** | Pointer to **NullableString** | Description of the Segment. | [optional] 
**CreatorEmail** | Pointer to **NullableString** | The email of the user who created the Segment. | [optional] 
**CreatorFullName** | Pointer to **NullableString** | The name of the user who created the Segment. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The date and time when the Segment was created. | [optional] 
**LastUpdaterEmail** | Pointer to **NullableString** | The email of the user who last updated the Segment. | [optional] 
**LastUpdaterFullName** | Pointer to **NullableString** | The name of the user who last updated the Segment. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The date and time when the Segment was last updated. | [optional] 
**ComparisonAttribute** | Pointer to **NullableString** | The user&#39;s attribute the evaluation process must take into account. | [optional] 
**Comparator** | Pointer to [**RolloutRuleComparator**](RolloutRuleComparator.md) |  | [optional] 
**ComparisonValue** | Pointer to **NullableString** | The value to compare with the given user attribute&#39;s value. | [optional] 

## Methods

### NewSegmentModel

`func NewSegmentModel() *SegmentModel`

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

### HasProduct

`func (o *SegmentModel) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

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

### HasSegmentId

`func (o *SegmentModel) HasSegmentId() bool`

HasSegmentId returns a boolean if a field has been set.

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

### HasName

`func (o *SegmentModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *SegmentModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *SegmentModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
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

### HasDescription

`func (o *SegmentModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

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

### HasCreatorEmail

`func (o *SegmentModel) HasCreatorEmail() bool`

HasCreatorEmail returns a boolean if a field has been set.

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

### HasCreatorFullName

`func (o *SegmentModel) HasCreatorFullName() bool`

HasCreatorFullName returns a boolean if a field has been set.

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

### HasCreatedAt

`func (o *SegmentModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

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

### HasLastUpdaterEmail

`func (o *SegmentModel) HasLastUpdaterEmail() bool`

HasLastUpdaterEmail returns a boolean if a field has been set.

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

### HasLastUpdaterFullName

`func (o *SegmentModel) HasLastUpdaterFullName() bool`

HasLastUpdaterFullName returns a boolean if a field has been set.

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

### HasUpdatedAt

`func (o *SegmentModel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

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

### HasComparisonAttribute

`func (o *SegmentModel) HasComparisonAttribute() bool`

HasComparisonAttribute returns a boolean if a field has been set.

### SetComparisonAttributeNil

`func (o *SegmentModel) SetComparisonAttributeNil(b bool)`

 SetComparisonAttributeNil sets the value for ComparisonAttribute to be an explicit nil

### UnsetComparisonAttribute
`func (o *SegmentModel) UnsetComparisonAttribute()`

UnsetComparisonAttribute ensures that no value is present for ComparisonAttribute, not even an explicit nil
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

### HasComparator

`func (o *SegmentModel) HasComparator() bool`

HasComparator returns a boolean if a field has been set.

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

### HasComparisonValue

`func (o *SegmentModel) HasComparisonValue() bool`

HasComparisonValue returns a boolean if a field has been set.

### SetComparisonValueNil

`func (o *SegmentModel) SetComparisonValueNil(b bool)`

 SetComparisonValueNil sets the value for ComparisonValue to be an explicit nil

### UnsetComparisonValue
`func (o *SegmentModel) UnsetComparisonValue()`

UnsetComparisonValue ensures that no value is present for ComparisonValue, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


