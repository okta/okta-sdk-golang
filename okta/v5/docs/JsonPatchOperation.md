# JsonPatchOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Op** | Pointer to **string** | The operation (PATCH action) | [optional] 
**Path** | Pointer to **string** | The resource path of the attribute to update | [optional] 
**Value** | Pointer to **map[string]interface{}** | The update operation value | [optional] 

## Methods

### NewJsonPatchOperation

`func NewJsonPatchOperation() *JsonPatchOperation`

NewJsonPatchOperation instantiates a new JsonPatchOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJsonPatchOperationWithDefaults

`func NewJsonPatchOperationWithDefaults() *JsonPatchOperation`

NewJsonPatchOperationWithDefaults instantiates a new JsonPatchOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOp

`func (o *JsonPatchOperation) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *JsonPatchOperation) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *JsonPatchOperation) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *JsonPatchOperation) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetPath

`func (o *JsonPatchOperation) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *JsonPatchOperation) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *JsonPatchOperation) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *JsonPatchOperation) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetValue

`func (o *JsonPatchOperation) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *JsonPatchOperation) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *JsonPatchOperation) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *JsonPatchOperation) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


