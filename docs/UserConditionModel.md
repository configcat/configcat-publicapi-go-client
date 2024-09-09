# UserConditionModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComparisonAttribute** | **string** | The User Object attribute that the condition is based on. Can be \&quot;User ID\&quot;, \&quot;Email\&quot;, \&quot;Country\&quot; or any custom attribute. | 
**Comparator** | **string** | The comparison operator which defines the relation between the comparison attribute and the comparison value. | 
**ComparisonValue** | [**ComparisonValueModel**](ComparisonValueModel.md) |  | 

## Methods

### NewUserConditionModel

`func NewUserConditionModel(comparisonAttribute string, comparator string, comparisonValue ComparisonValueModel, ) *UserConditionModel`

NewUserConditionModel instantiates a new UserConditionModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserConditionModelWithDefaults

`func NewUserConditionModelWithDefaults() *UserConditionModel`

NewUserConditionModelWithDefaults instantiates a new UserConditionModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComparisonAttribute

`func (o *UserConditionModel) GetComparisonAttribute() string`

GetComparisonAttribute returns the ComparisonAttribute field if non-nil, zero value otherwise.

### GetComparisonAttributeOk

`func (o *UserConditionModel) GetComparisonAttributeOk() (*string, bool)`

GetComparisonAttributeOk returns a tuple with the ComparisonAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparisonAttribute

`func (o *UserConditionModel) SetComparisonAttribute(v string)`

SetComparisonAttribute sets ComparisonAttribute field to given value.


### GetComparator

`func (o *UserConditionModel) GetComparator() string`

GetComparator returns the Comparator field if non-nil, zero value otherwise.

### GetComparatorOk

`func (o *UserConditionModel) GetComparatorOk() (*string, bool)`

GetComparatorOk returns a tuple with the Comparator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparator

`func (o *UserConditionModel) SetComparator(v string)`

SetComparator sets Comparator field to given value.


### GetComparisonValue

`func (o *UserConditionModel) GetComparisonValue() ComparisonValueModel`

GetComparisonValue returns the ComparisonValue field if non-nil, zero value otherwise.

### GetComparisonValueOk

`func (o *UserConditionModel) GetComparisonValueOk() (*ComparisonValueModel, bool)`

GetComparisonValueOk returns a tuple with the ComparisonValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparisonValue

`func (o *UserConditionModel) SetComparisonValue(v ComparisonValueModel)`

SetComparisonValue sets ComparisonValue field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


