# SegmentListModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Product** | Pointer to [**ProductModel**](ProductModel.md) |  | [optional] 
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

## Methods

### NewSegmentListModel

`func NewSegmentListModel() *SegmentListModel`

NewSegmentListModel instantiates a new SegmentListModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSegmentListModelWithDefaults

`func NewSegmentListModelWithDefaults() *SegmentListModel`

NewSegmentListModelWithDefaults instantiates a new SegmentListModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProduct

`func (o *SegmentListModel) GetProduct() ProductModel`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *SegmentListModel) GetProductOk() (*ProductModel, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *SegmentListModel) SetProduct(v ProductModel)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *SegmentListModel) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetSegmentId

`func (o *SegmentListModel) GetSegmentId() string`

GetSegmentId returns the SegmentId field if non-nil, zero value otherwise.

### GetSegmentIdOk

`func (o *SegmentListModel) GetSegmentIdOk() (*string, bool)`

GetSegmentIdOk returns a tuple with the SegmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentId

`func (o *SegmentListModel) SetSegmentId(v string)`

SetSegmentId sets SegmentId field to given value.

### HasSegmentId

`func (o *SegmentListModel) HasSegmentId() bool`

HasSegmentId returns a boolean if a field has been set.

### GetName

`func (o *SegmentListModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SegmentListModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SegmentListModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SegmentListModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *SegmentListModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *SegmentListModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *SegmentListModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SegmentListModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SegmentListModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SegmentListModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *SegmentListModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SegmentListModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCreatorEmail

`func (o *SegmentListModel) GetCreatorEmail() string`

GetCreatorEmail returns the CreatorEmail field if non-nil, zero value otherwise.

### GetCreatorEmailOk

`func (o *SegmentListModel) GetCreatorEmailOk() (*string, bool)`

GetCreatorEmailOk returns a tuple with the CreatorEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorEmail

`func (o *SegmentListModel) SetCreatorEmail(v string)`

SetCreatorEmail sets CreatorEmail field to given value.

### HasCreatorEmail

`func (o *SegmentListModel) HasCreatorEmail() bool`

HasCreatorEmail returns a boolean if a field has been set.

### SetCreatorEmailNil

`func (o *SegmentListModel) SetCreatorEmailNil(b bool)`

 SetCreatorEmailNil sets the value for CreatorEmail to be an explicit nil

### UnsetCreatorEmail
`func (o *SegmentListModel) UnsetCreatorEmail()`

UnsetCreatorEmail ensures that no value is present for CreatorEmail, not even an explicit nil
### GetCreatorFullName

`func (o *SegmentListModel) GetCreatorFullName() string`

GetCreatorFullName returns the CreatorFullName field if non-nil, zero value otherwise.

### GetCreatorFullNameOk

`func (o *SegmentListModel) GetCreatorFullNameOk() (*string, bool)`

GetCreatorFullNameOk returns a tuple with the CreatorFullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorFullName

`func (o *SegmentListModel) SetCreatorFullName(v string)`

SetCreatorFullName sets CreatorFullName field to given value.

### HasCreatorFullName

`func (o *SegmentListModel) HasCreatorFullName() bool`

HasCreatorFullName returns a boolean if a field has been set.

### SetCreatorFullNameNil

`func (o *SegmentListModel) SetCreatorFullNameNil(b bool)`

 SetCreatorFullNameNil sets the value for CreatorFullName to be an explicit nil

### UnsetCreatorFullName
`func (o *SegmentListModel) UnsetCreatorFullName()`

UnsetCreatorFullName ensures that no value is present for CreatorFullName, not even an explicit nil
### GetCreatedAt

`func (o *SegmentListModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SegmentListModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SegmentListModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SegmentListModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastUpdaterEmail

`func (o *SegmentListModel) GetLastUpdaterEmail() string`

GetLastUpdaterEmail returns the LastUpdaterEmail field if non-nil, zero value otherwise.

### GetLastUpdaterEmailOk

`func (o *SegmentListModel) GetLastUpdaterEmailOk() (*string, bool)`

GetLastUpdaterEmailOk returns a tuple with the LastUpdaterEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdaterEmail

`func (o *SegmentListModel) SetLastUpdaterEmail(v string)`

SetLastUpdaterEmail sets LastUpdaterEmail field to given value.

### HasLastUpdaterEmail

`func (o *SegmentListModel) HasLastUpdaterEmail() bool`

HasLastUpdaterEmail returns a boolean if a field has been set.

### SetLastUpdaterEmailNil

`func (o *SegmentListModel) SetLastUpdaterEmailNil(b bool)`

 SetLastUpdaterEmailNil sets the value for LastUpdaterEmail to be an explicit nil

### UnsetLastUpdaterEmail
`func (o *SegmentListModel) UnsetLastUpdaterEmail()`

UnsetLastUpdaterEmail ensures that no value is present for LastUpdaterEmail, not even an explicit nil
### GetLastUpdaterFullName

`func (o *SegmentListModel) GetLastUpdaterFullName() string`

GetLastUpdaterFullName returns the LastUpdaterFullName field if non-nil, zero value otherwise.

### GetLastUpdaterFullNameOk

`func (o *SegmentListModel) GetLastUpdaterFullNameOk() (*string, bool)`

GetLastUpdaterFullNameOk returns a tuple with the LastUpdaterFullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdaterFullName

`func (o *SegmentListModel) SetLastUpdaterFullName(v string)`

SetLastUpdaterFullName sets LastUpdaterFullName field to given value.

### HasLastUpdaterFullName

`func (o *SegmentListModel) HasLastUpdaterFullName() bool`

HasLastUpdaterFullName returns a boolean if a field has been set.

### SetLastUpdaterFullNameNil

`func (o *SegmentListModel) SetLastUpdaterFullNameNil(b bool)`

 SetLastUpdaterFullNameNil sets the value for LastUpdaterFullName to be an explicit nil

### UnsetLastUpdaterFullName
`func (o *SegmentListModel) UnsetLastUpdaterFullName()`

UnsetLastUpdaterFullName ensures that no value is present for LastUpdaterFullName, not even an explicit nil
### GetUpdatedAt

`func (o *SegmentListModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SegmentListModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SegmentListModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SegmentListModel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUsage

`func (o *SegmentListModel) GetUsage() int32`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *SegmentListModel) GetUsageOk() (*int32, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *SegmentListModel) SetUsage(v int32)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *SegmentListModel) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


