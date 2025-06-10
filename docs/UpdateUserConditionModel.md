# UpdateUserConditionModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComparisonAttribute** | **string** | The User Object attribute that the condition is based on. Can be \&quot;User ID\&quot;, \&quot;Email\&quot;, \&quot;Country\&quot; or any custom attribute. | 
**Comparator** | [**UserComparator**](UserComparator.md) |  | 
**ComparisonValue** | [**UpdateComparisonValueModel**](UpdateComparisonValueModel.md) |  | 

## Methods

### NewUpdateUserConditionModel

`func NewUpdateUserConditionModel(comparisonAttribute string, comparator UserComparator, comparisonValue UpdateComparisonValueModel, ) *UpdateUserConditionModel`

NewUpdateUserConditionModel instantiates a new UpdateUserConditionModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateUserConditionModelWithDefaults

`func NewUpdateUserConditionModelWithDefaults() *UpdateUserConditionModel`

NewUpdateUserConditionModelWithDefaults instantiates a new UpdateUserConditionModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComparisonAttribute

`func (o *UpdateUserConditionModel) GetComparisonAttribute() string`

GetComparisonAttribute returns the ComparisonAttribute field if non-nil, zero value otherwise.

### GetComparisonAttributeOk

`func (o *UpdateUserConditionModel) GetComparisonAttributeOk() (*string, bool)`

GetComparisonAttributeOk returns a tuple with the ComparisonAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparisonAttribute

`func (o *UpdateUserConditionModel) SetComparisonAttribute(v string)`

SetComparisonAttribute sets ComparisonAttribute field to given value.


### GetComparator

`func (o *UpdateUserConditionModel) GetComparator() UserComparator`

GetComparator returns the Comparator field if non-nil, zero value otherwise.

### GetComparatorOk

`func (o *UpdateUserConditionModel) GetComparatorOk() (*UserComparator, bool)`

GetComparatorOk returns a tuple with the Comparator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparator

`func (o *UpdateUserConditionModel) SetComparator(v UserComparator)`

SetComparator sets Comparator field to given value.


### GetComparisonValue

`func (o *UpdateUserConditionModel) GetComparisonValue() UpdateComparisonValueModel`

GetComparisonValue returns the ComparisonValue field if non-nil, zero value otherwise.

### GetComparisonValueOk

`func (o *UpdateUserConditionModel) GetComparisonValueOk() (*UpdateComparisonValueModel, bool)`

GetComparisonValueOk returns a tuple with the ComparisonValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparisonValue

`func (o *UpdateUserConditionModel) SetComparisonValue(v UpdateComparisonValueModel)`

SetComparisonValue sets ComparisonValue field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


