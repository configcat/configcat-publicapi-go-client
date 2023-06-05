# JsonNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Options** | Pointer to [**JsonNodeOptions**](JsonNodeOptions.md) |  | [optional] 
**Parent** | Pointer to [**JsonNode**](JsonNode.md) |  | [optional] 
**Root** | Pointer to [**JsonNode**](JsonNode.md) |  | [optional] 

## Methods

### NewJsonNode

`func NewJsonNode() *JsonNode`

NewJsonNode instantiates a new JsonNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJsonNodeWithDefaults

`func NewJsonNodeWithDefaults() *JsonNode`

NewJsonNodeWithDefaults instantiates a new JsonNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOptions

`func (o *JsonNode) GetOptions() JsonNodeOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *JsonNode) GetOptionsOk() (*JsonNodeOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *JsonNode) SetOptions(v JsonNodeOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *JsonNode) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetParent

`func (o *JsonNode) GetParent() JsonNode`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *JsonNode) GetParentOk() (*JsonNode, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *JsonNode) SetParent(v JsonNode)`

SetParent sets Parent field to given value.

### HasParent

`func (o *JsonNode) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetRoot

`func (o *JsonNode) GetRoot() JsonNode`

GetRoot returns the Root field if non-nil, zero value otherwise.

### GetRootOk

`func (o *JsonNode) GetRootOk() (*JsonNode, bool)`

GetRootOk returns a tuple with the Root field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoot

`func (o *JsonNode) SetRoot(v JsonNode)`

SetRoot sets Root field to given value.

### HasRoot

`func (o *JsonNode) HasRoot() bool`

HasRoot returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


