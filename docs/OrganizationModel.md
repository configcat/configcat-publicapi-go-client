# OrganizationModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrganizationId** | Pointer to **string** | Identifier of the Organization. | [optional] 
**Name** | Pointer to **NullableString** | Name of the Organization. | [optional] 
**Links** | Pointer to [**OrganizationModelLinks**](OrganizationModelLinks.md) |  | [optional] 

## Methods

### NewOrganizationModel

`func NewOrganizationModel() *OrganizationModel`

NewOrganizationModel instantiates a new OrganizationModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationModelWithDefaults

`func NewOrganizationModelWithDefaults() *OrganizationModel`

NewOrganizationModelWithDefaults instantiates a new OrganizationModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganizationId

`func (o *OrganizationModel) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *OrganizationModel) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *OrganizationModel) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *OrganizationModel) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetName

`func (o *OrganizationModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrganizationModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *OrganizationModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *OrganizationModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetLinks

`func (o *OrganizationModel) GetLinks() OrganizationModelLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *OrganizationModel) GetLinksOk() (*OrganizationModelLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *OrganizationModel) SetLinks(v OrganizationModelLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *OrganizationModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


