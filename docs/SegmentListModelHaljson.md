# SegmentListModelHaljson

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
**Usage** | Pointer to **int32** |  | [optional] 
**Links** | Pointer to [**EnvironmentModelHaljsonLinks**](EnvironmentModelHaljsonLinks.md) |  | [optional] 

## Methods

### NewSegmentListModelHaljson

`func NewSegmentListModelHaljson() *SegmentListModelHaljson`

NewSegmentListModelHaljson instantiates a new SegmentListModelHaljson object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSegmentListModelHaljsonWithDefaults

`func NewSegmentListModelHaljsonWithDefaults() *SegmentListModelHaljson`

NewSegmentListModelHaljsonWithDefaults instantiates a new SegmentListModelHaljson object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmbedded

`func (o *SegmentListModelHaljson) GetEmbedded() ConfigModelHaljsonEmbedded`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *SegmentListModelHaljson) GetEmbeddedOk() (*ConfigModelHaljsonEmbedded, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *SegmentListModelHaljson) SetEmbedded(v ConfigModelHaljsonEmbedded)`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *SegmentListModelHaljson) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetSegmentId

`func (o *SegmentListModelHaljson) GetSegmentId() string`

GetSegmentId returns the SegmentId field if non-nil, zero value otherwise.

### GetSegmentIdOk

`func (o *SegmentListModelHaljson) GetSegmentIdOk() (*string, bool)`

GetSegmentIdOk returns a tuple with the SegmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentId

`func (o *SegmentListModelHaljson) SetSegmentId(v string)`

SetSegmentId sets SegmentId field to given value.

### HasSegmentId

`func (o *SegmentListModelHaljson) HasSegmentId() bool`

HasSegmentId returns a boolean if a field has been set.

### GetName

`func (o *SegmentListModelHaljson) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SegmentListModelHaljson) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SegmentListModelHaljson) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SegmentListModelHaljson) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *SegmentListModelHaljson) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *SegmentListModelHaljson) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *SegmentListModelHaljson) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SegmentListModelHaljson) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SegmentListModelHaljson) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SegmentListModelHaljson) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *SegmentListModelHaljson) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SegmentListModelHaljson) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCreatorEmail

`func (o *SegmentListModelHaljson) GetCreatorEmail() string`

GetCreatorEmail returns the CreatorEmail field if non-nil, zero value otherwise.

### GetCreatorEmailOk

`func (o *SegmentListModelHaljson) GetCreatorEmailOk() (*string, bool)`

GetCreatorEmailOk returns a tuple with the CreatorEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorEmail

`func (o *SegmentListModelHaljson) SetCreatorEmail(v string)`

SetCreatorEmail sets CreatorEmail field to given value.

### HasCreatorEmail

`func (o *SegmentListModelHaljson) HasCreatorEmail() bool`

HasCreatorEmail returns a boolean if a field has been set.

### SetCreatorEmailNil

`func (o *SegmentListModelHaljson) SetCreatorEmailNil(b bool)`

 SetCreatorEmailNil sets the value for CreatorEmail to be an explicit nil

### UnsetCreatorEmail
`func (o *SegmentListModelHaljson) UnsetCreatorEmail()`

UnsetCreatorEmail ensures that no value is present for CreatorEmail, not even an explicit nil
### GetCreatorFullName

`func (o *SegmentListModelHaljson) GetCreatorFullName() string`

GetCreatorFullName returns the CreatorFullName field if non-nil, zero value otherwise.

### GetCreatorFullNameOk

`func (o *SegmentListModelHaljson) GetCreatorFullNameOk() (*string, bool)`

GetCreatorFullNameOk returns a tuple with the CreatorFullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorFullName

`func (o *SegmentListModelHaljson) SetCreatorFullName(v string)`

SetCreatorFullName sets CreatorFullName field to given value.

### HasCreatorFullName

`func (o *SegmentListModelHaljson) HasCreatorFullName() bool`

HasCreatorFullName returns a boolean if a field has been set.

### SetCreatorFullNameNil

`func (o *SegmentListModelHaljson) SetCreatorFullNameNil(b bool)`

 SetCreatorFullNameNil sets the value for CreatorFullName to be an explicit nil

### UnsetCreatorFullName
`func (o *SegmentListModelHaljson) UnsetCreatorFullName()`

UnsetCreatorFullName ensures that no value is present for CreatorFullName, not even an explicit nil
### GetCreatedAt

`func (o *SegmentListModelHaljson) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SegmentListModelHaljson) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SegmentListModelHaljson) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SegmentListModelHaljson) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastUpdaterEmail

`func (o *SegmentListModelHaljson) GetLastUpdaterEmail() string`

GetLastUpdaterEmail returns the LastUpdaterEmail field if non-nil, zero value otherwise.

### GetLastUpdaterEmailOk

`func (o *SegmentListModelHaljson) GetLastUpdaterEmailOk() (*string, bool)`

GetLastUpdaterEmailOk returns a tuple with the LastUpdaterEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdaterEmail

`func (o *SegmentListModelHaljson) SetLastUpdaterEmail(v string)`

SetLastUpdaterEmail sets LastUpdaterEmail field to given value.

### HasLastUpdaterEmail

`func (o *SegmentListModelHaljson) HasLastUpdaterEmail() bool`

HasLastUpdaterEmail returns a boolean if a field has been set.

### SetLastUpdaterEmailNil

`func (o *SegmentListModelHaljson) SetLastUpdaterEmailNil(b bool)`

 SetLastUpdaterEmailNil sets the value for LastUpdaterEmail to be an explicit nil

### UnsetLastUpdaterEmail
`func (o *SegmentListModelHaljson) UnsetLastUpdaterEmail()`

UnsetLastUpdaterEmail ensures that no value is present for LastUpdaterEmail, not even an explicit nil
### GetLastUpdaterFullName

`func (o *SegmentListModelHaljson) GetLastUpdaterFullName() string`

GetLastUpdaterFullName returns the LastUpdaterFullName field if non-nil, zero value otherwise.

### GetLastUpdaterFullNameOk

`func (o *SegmentListModelHaljson) GetLastUpdaterFullNameOk() (*string, bool)`

GetLastUpdaterFullNameOk returns a tuple with the LastUpdaterFullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdaterFullName

`func (o *SegmentListModelHaljson) SetLastUpdaterFullName(v string)`

SetLastUpdaterFullName sets LastUpdaterFullName field to given value.

### HasLastUpdaterFullName

`func (o *SegmentListModelHaljson) HasLastUpdaterFullName() bool`

HasLastUpdaterFullName returns a boolean if a field has been set.

### SetLastUpdaterFullNameNil

`func (o *SegmentListModelHaljson) SetLastUpdaterFullNameNil(b bool)`

 SetLastUpdaterFullNameNil sets the value for LastUpdaterFullName to be an explicit nil

### UnsetLastUpdaterFullName
`func (o *SegmentListModelHaljson) UnsetLastUpdaterFullName()`

UnsetLastUpdaterFullName ensures that no value is present for LastUpdaterFullName, not even an explicit nil
### GetUpdatedAt

`func (o *SegmentListModelHaljson) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SegmentListModelHaljson) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SegmentListModelHaljson) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SegmentListModelHaljson) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUsage

`func (o *SegmentListModelHaljson) GetUsage() int32`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *SegmentListModelHaljson) GetUsageOk() (*int32, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *SegmentListModelHaljson) SetUsage(v int32)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *SegmentListModelHaljson) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetLinks

`func (o *SegmentListModelHaljson) GetLinks() EnvironmentModelHaljsonLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *SegmentListModelHaljson) GetLinksOk() (*EnvironmentModelHaljsonLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *SegmentListModelHaljson) SetLinks(v EnvironmentModelHaljsonLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *SegmentListModelHaljson) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


