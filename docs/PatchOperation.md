# PatchOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Op** | Pointer to [**OperationType**](OperationType.md) |  | [optional] 
**From** | Pointer to [**JsonPointer**](JsonPointer.md) |  | [optional] 
**Path** | Pointer to [**JsonPointer**](JsonPointer.md) |  | [optional] 
**Value** | Pointer to [**JsonNode**](JsonNode.md) |  | [optional] 

## Methods

### NewPatchOperation

`func NewPatchOperation() *PatchOperation`

NewPatchOperation instantiates a new PatchOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchOperationWithDefaults

`func NewPatchOperationWithDefaults() *PatchOperation`

NewPatchOperationWithDefaults instantiates a new PatchOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOp

`func (o *PatchOperation) GetOp() OperationType`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *PatchOperation) GetOpOk() (*OperationType, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *PatchOperation) SetOp(v OperationType)`

SetOp sets Op field to given value.

### HasOp

`func (o *PatchOperation) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetFrom

`func (o *PatchOperation) GetFrom() JsonPointer`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *PatchOperation) GetFromOk() (*JsonPointer, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *PatchOperation) SetFrom(v JsonPointer)`

SetFrom sets From field to given value.

### HasFrom

`func (o *PatchOperation) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetPath

`func (o *PatchOperation) GetPath() JsonPointer`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *PatchOperation) GetPathOk() (*JsonPointer, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *PatchOperation) SetPath(v JsonPointer)`

SetPath sets Path field to given value.

### HasPath

`func (o *PatchOperation) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetValue

`func (o *PatchOperation) GetValue() JsonNode`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PatchOperation) GetValueOk() (*JsonNode, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PatchOperation) SetValue(v JsonNode)`

SetValue sets Value field to given value.

### HasValue

`func (o *PatchOperation) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


