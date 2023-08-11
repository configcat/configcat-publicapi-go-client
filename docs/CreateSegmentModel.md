# CreateSegmentModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the Segment. | 
**Description** | Pointer to **NullableString** | Description of the Segment. | [optional] 
**ComparisonAttribute** | **string** | The user&#39;s attribute the evaluation process must take into account. | 
**Comparator** | [**RolloutRuleComparator**](RolloutRuleComparator.md) |  | 
**ComparisonValue** | **string** | The value to compare with the given user attribute&#39;s value. | 

## Methods

### NewCreateSegmentModel

`func NewCreateSegmentModel(name string, comparisonAttribute string, comparator RolloutRuleComparator, comparisonValue string, ) *CreateSegmentModel`

NewCreateSegmentModel instantiates a new CreateSegmentModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSegmentModelWithDefaults

`func NewCreateSegmentModelWithDefaults() *CreateSegmentModel`

NewCreateSegmentModelWithDefaults instantiates a new CreateSegmentModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateSegmentModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateSegmentModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateSegmentModel) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateSegmentModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateSegmentModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateSegmentModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateSegmentModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateSegmentModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateSegmentModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetComparisonAttribute

`func (o *CreateSegmentModel) GetComparisonAttribute() string`

GetComparisonAttribute returns the ComparisonAttribute field if non-nil, zero value otherwise.

### GetComparisonAttributeOk

`func (o *CreateSegmentModel) GetComparisonAttributeOk() (*string, bool)`

GetComparisonAttributeOk returns a tuple with the ComparisonAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparisonAttribute

`func (o *CreateSegmentModel) SetComparisonAttribute(v string)`

SetComparisonAttribute sets ComparisonAttribute field to given value.


### GetComparator

`func (o *CreateSegmentModel) GetComparator() RolloutRuleComparator`

GetComparator returns the Comparator field if non-nil, zero value otherwise.

### GetComparatorOk

`func (o *CreateSegmentModel) GetComparatorOk() (*RolloutRuleComparator, bool)`

GetComparatorOk returns a tuple with the Comparator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparator

`func (o *CreateSegmentModel) SetComparator(v RolloutRuleComparator)`

SetComparator sets Comparator field to given value.


### GetComparisonValue

`func (o *CreateSegmentModel) GetComparisonValue() string`

GetComparisonValue returns the ComparisonValue field if non-nil, zero value otherwise.

### GetComparisonValueOk

`func (o *CreateSegmentModel) GetComparisonValueOk() (*string, bool)`

GetComparisonValueOk returns a tuple with the ComparisonValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparisonValue

`func (o *CreateSegmentModel) SetComparisonValue(v string)`

SetComparisonValue sets ComparisonValue field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


