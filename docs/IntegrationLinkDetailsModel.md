# IntegrationLinkDetailsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Details** | Pointer to [**[]IntegrationLinkDetail**](IntegrationLinkDetail.md) |  | [optional] [readonly] 
**AllIntegrationLinkCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewIntegrationLinkDetailsModel

`func NewIntegrationLinkDetailsModel() *IntegrationLinkDetailsModel`

NewIntegrationLinkDetailsModel instantiates a new IntegrationLinkDetailsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationLinkDetailsModelWithDefaults

`func NewIntegrationLinkDetailsModelWithDefaults() *IntegrationLinkDetailsModel`

NewIntegrationLinkDetailsModelWithDefaults instantiates a new IntegrationLinkDetailsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetails

`func (o *IntegrationLinkDetailsModel) GetDetails() []IntegrationLinkDetail`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *IntegrationLinkDetailsModel) GetDetailsOk() (*[]IntegrationLinkDetail, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *IntegrationLinkDetailsModel) SetDetails(v []IntegrationLinkDetail)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *IntegrationLinkDetailsModel) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### SetDetailsNil

`func (o *IntegrationLinkDetailsModel) SetDetailsNil(b bool)`

 SetDetailsNil sets the value for Details to be an explicit nil

### UnsetDetails
`func (o *IntegrationLinkDetailsModel) UnsetDetails()`

UnsetDetails ensures that no value is present for Details, not even an explicit nil
### GetAllIntegrationLinkCount

`func (o *IntegrationLinkDetailsModel) GetAllIntegrationLinkCount() int32`

GetAllIntegrationLinkCount returns the AllIntegrationLinkCount field if non-nil, zero value otherwise.

### GetAllIntegrationLinkCountOk

`func (o *IntegrationLinkDetailsModel) GetAllIntegrationLinkCountOk() (*int32, bool)`

GetAllIntegrationLinkCountOk returns a tuple with the AllIntegrationLinkCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllIntegrationLinkCount

`func (o *IntegrationLinkDetailsModel) SetAllIntegrationLinkCount(v int32)`

SetAllIntegrationLinkCount sets AllIntegrationLinkCount field to given value.

### HasAllIntegrationLinkCount

`func (o *IntegrationLinkDetailsModel) HasAllIntegrationLinkCount() bool`

HasAllIntegrationLinkCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


