# SegmentModelHaljson

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Embedded** | Pointer to [**ConfigModelHaljsonEmbedded**](ConfigModelHaljsonEmbedded.md) |  | [optional] 
**SegmentId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**CreatorEmail** | Pointer to **NullableString** |  | [optional] 
**CreatorFullName** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**LastUpdaterEmail** | Pointer to **NullableString** |  | [optional] 
**LastUpdaterFullName** | Pointer to **NullableString** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**ComparisonAttribute** | Pointer to **NullableString** |  | [optional] 
**Comparator** | Pointer to [**RolloutRuleComparator**](RolloutRuleComparator.md) |  | [optional] 
**ComparisonValue** | Pointer to **NullableString** |  | [optional] 
**Links** | Pointer to [**EnvironmentModelHaljsonLinks**](EnvironmentModelHaljsonLinks.md) |  | [optional] 

## Methods

### NewSegmentModelHaljson

`func NewSegmentModelHaljson() *SegmentModelHaljson`

NewSegmentModelHaljson instantiates a new SegmentModelHaljson object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSegmentModelHaljsonWithDefaults

`func NewSegmentModelHaljsonWithDefaults() *SegmentModelHaljson`

NewSegmentModelHaljsonWithDefaults instantiates a new SegmentModelHaljson object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmbedded

`func (o *SegmentModelHaljson) GetEmbedded() ConfigModelHaljsonEmbedded`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *SegmentModelHaljson) GetEmbeddedOk() (*ConfigModelHaljsonEmbedded, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *SegmentModelHaljson) SetEmbedded(v ConfigModelHaljsonEmbedded)`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *SegmentModelHaljson) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetSegmentId

`func (o *SegmentModelHaljson) GetSegmentId() string`

GetSegmentId returns the SegmentId field if non-nil, zero value otherwise.

### GetSegmentIdOk

`func (o *SegmentModelHaljson) GetSegmentIdOk() (*string, bool)`

GetSegmentIdOk returns a tuple with the SegmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentId

`func (o *SegmentModelHaljson) SetSegmentId(v string)`

SetSegmentId sets SegmentId field to given value.

### HasSegmentId

`func (o *SegmentModelHaljson) HasSegmentId() bool`

HasSegmentId returns a boolean if a field has been set.

### GetName

`func (o *SegmentModelHaljson) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SegmentModelHaljson) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SegmentModelHaljson) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SegmentModelHaljson) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *SegmentModelHaljson) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *SegmentModelHaljson) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *SegmentModelHaljson) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SegmentModelHaljson) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SegmentModelHaljson) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SegmentModelHaljson) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *SegmentModelHaljson) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SegmentModelHaljson) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCreatorEmail

`func (o *SegmentModelHaljson) GetCreatorEmail() string`

GetCreatorEmail returns the CreatorEmail field if non-nil, zero value otherwise.

### GetCreatorEmailOk

`func (o *SegmentModelHaljson) GetCreatorEmailOk() (*string, bool)`

GetCreatorEmailOk returns a tuple with the CreatorEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorEmail

`func (o *SegmentModelHaljson) SetCreatorEmail(v string)`

SetCreatorEmail sets CreatorEmail field to given value.

### HasCreatorEmail

`func (o *SegmentModelHaljson) HasCreatorEmail() bool`

HasCreatorEmail returns a boolean if a field has been set.

### SetCreatorEmailNil

`func (o *SegmentModelHaljson) SetCreatorEmailNil(b bool)`

 SetCreatorEmailNil sets the value for CreatorEmail to be an explicit nil

### UnsetCreatorEmail
`func (o *SegmentModelHaljson) UnsetCreatorEmail()`

UnsetCreatorEmail ensures that no value is present for CreatorEmail, not even an explicit nil
### GetCreatorFullName

`func (o *SegmentModelHaljson) GetCreatorFullName() string`

GetCreatorFullName returns the CreatorFullName field if non-nil, zero value otherwise.

### GetCreatorFullNameOk

`func (o *SegmentModelHaljson) GetCreatorFullNameOk() (*string, bool)`

GetCreatorFullNameOk returns a tuple with the CreatorFullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorFullName

`func (o *SegmentModelHaljson) SetCreatorFullName(v string)`

SetCreatorFullName sets CreatorFullName field to given value.

### HasCreatorFullName

`func (o *SegmentModelHaljson) HasCreatorFullName() bool`

HasCreatorFullName returns a boolean if a field has been set.

### SetCreatorFullNameNil

`func (o *SegmentModelHaljson) SetCreatorFullNameNil(b bool)`

 SetCreatorFullNameNil sets the value for CreatorFullName to be an explicit nil

### UnsetCreatorFullName
`func (o *SegmentModelHaljson) UnsetCreatorFullName()`

UnsetCreatorFullName ensures that no value is present for CreatorFullName, not even an explicit nil
### GetCreatedAt

`func (o *SegmentModelHaljson) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SegmentModelHaljson) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SegmentModelHaljson) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SegmentModelHaljson) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastUpdaterEmail

`func (o *SegmentModelHaljson) GetLastUpdaterEmail() string`

GetLastUpdaterEmail returns the LastUpdaterEmail field if non-nil, zero value otherwise.

### GetLastUpdaterEmailOk

`func (o *SegmentModelHaljson) GetLastUpdaterEmailOk() (*string, bool)`

GetLastUpdaterEmailOk returns a tuple with the LastUpdaterEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdaterEmail

`func (o *SegmentModelHaljson) SetLastUpdaterEmail(v string)`

SetLastUpdaterEmail sets LastUpdaterEmail field to given value.

### HasLastUpdaterEmail

`func (o *SegmentModelHaljson) HasLastUpdaterEmail() bool`

HasLastUpdaterEmail returns a boolean if a field has been set.

### SetLastUpdaterEmailNil

`func (o *SegmentModelHaljson) SetLastUpdaterEmailNil(b bool)`

 SetLastUpdaterEmailNil sets the value for LastUpdaterEmail to be an explicit nil

### UnsetLastUpdaterEmail
`func (o *SegmentModelHaljson) UnsetLastUpdaterEmail()`

UnsetLastUpdaterEmail ensures that no value is present for LastUpdaterEmail, not even an explicit nil
### GetLastUpdaterFullName

`func (o *SegmentModelHaljson) GetLastUpdaterFullName() string`

GetLastUpdaterFullName returns the LastUpdaterFullName field if non-nil, zero value otherwise.

### GetLastUpdaterFullNameOk

`func (o *SegmentModelHaljson) GetLastUpdaterFullNameOk() (*string, bool)`

GetLastUpdaterFullNameOk returns a tuple with the LastUpdaterFullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdaterFullName

`func (o *SegmentModelHaljson) SetLastUpdaterFullName(v string)`

SetLastUpdaterFullName sets LastUpdaterFullName field to given value.

### HasLastUpdaterFullName

`func (o *SegmentModelHaljson) HasLastUpdaterFullName() bool`

HasLastUpdaterFullName returns a boolean if a field has been set.

### SetLastUpdaterFullNameNil

`func (o *SegmentModelHaljson) SetLastUpdaterFullNameNil(b bool)`

 SetLastUpdaterFullNameNil sets the value for LastUpdaterFullName to be an explicit nil

### UnsetLastUpdaterFullName
`func (o *SegmentModelHaljson) UnsetLastUpdaterFullName()`

UnsetLastUpdaterFullName ensures that no value is present for LastUpdaterFullName, not even an explicit nil
### GetUpdatedAt

`func (o *SegmentModelHaljson) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SegmentModelHaljson) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SegmentModelHaljson) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SegmentModelHaljson) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetComparisonAttribute

`func (o *SegmentModelHaljson) GetComparisonAttribute() string`

GetComparisonAttribute returns the ComparisonAttribute field if non-nil, zero value otherwise.

### GetComparisonAttributeOk

`func (o *SegmentModelHaljson) GetComparisonAttributeOk() (*string, bool)`

GetComparisonAttributeOk returns a tuple with the ComparisonAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparisonAttribute

`func (o *SegmentModelHaljson) SetComparisonAttribute(v string)`

SetComparisonAttribute sets ComparisonAttribute field to given value.

### HasComparisonAttribute

`func (o *SegmentModelHaljson) HasComparisonAttribute() bool`

HasComparisonAttribute returns a boolean if a field has been set.

### SetComparisonAttributeNil

`func (o *SegmentModelHaljson) SetComparisonAttributeNil(b bool)`

 SetComparisonAttributeNil sets the value for ComparisonAttribute to be an explicit nil

### UnsetComparisonAttribute
`func (o *SegmentModelHaljson) UnsetComparisonAttribute()`

UnsetComparisonAttribute ensures that no value is present for ComparisonAttribute, not even an explicit nil
### GetComparator

`func (o *SegmentModelHaljson) GetComparator() RolloutRuleComparator`

GetComparator returns the Comparator field if non-nil, zero value otherwise.

### GetComparatorOk

`func (o *SegmentModelHaljson) GetComparatorOk() (*RolloutRuleComparator, bool)`

GetComparatorOk returns a tuple with the Comparator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparator

`func (o *SegmentModelHaljson) SetComparator(v RolloutRuleComparator)`

SetComparator sets Comparator field to given value.

### HasComparator

`func (o *SegmentModelHaljson) HasComparator() bool`

HasComparator returns a boolean if a field has been set.

### GetComparisonValue

`func (o *SegmentModelHaljson) GetComparisonValue() string`

GetComparisonValue returns the ComparisonValue field if non-nil, zero value otherwise.

### GetComparisonValueOk

`func (o *SegmentModelHaljson) GetComparisonValueOk() (*string, bool)`

GetComparisonValueOk returns a tuple with the ComparisonValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparisonValue

`func (o *SegmentModelHaljson) SetComparisonValue(v string)`

SetComparisonValue sets ComparisonValue field to given value.

### HasComparisonValue

`func (o *SegmentModelHaljson) HasComparisonValue() bool`

HasComparisonValue returns a boolean if a field has been set.

### SetComparisonValueNil

`func (o *SegmentModelHaljson) SetComparisonValueNil(b bool)`

 SetComparisonValueNil sets the value for ComparisonValue to be an explicit nil

### UnsetComparisonValue
`func (o *SegmentModelHaljson) UnsetComparisonValue()`

UnsetComparisonValue ensures that no value is present for ComparisonValue, not even an explicit nil
### GetLinks

`func (o *SegmentModelHaljson) GetLinks() EnvironmentModelHaljsonLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *SegmentModelHaljson) GetLinksOk() (*EnvironmentModelHaljsonLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *SegmentModelHaljson) SetLinks(v EnvironmentModelHaljsonLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *SegmentModelHaljson) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


