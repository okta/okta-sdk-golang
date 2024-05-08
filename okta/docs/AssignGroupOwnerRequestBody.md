# AssignGroupOwnerRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The &#x60;id&#x60; of the group owner | [optional] 
**Type** | Pointer to **string** | The entity type of the owner | [optional] 

## Methods

### NewAssignGroupOwnerRequestBody

`func NewAssignGroupOwnerRequestBody() *AssignGroupOwnerRequestBody`

NewAssignGroupOwnerRequestBody instantiates a new AssignGroupOwnerRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssignGroupOwnerRequestBodyWithDefaults

`func NewAssignGroupOwnerRequestBodyWithDefaults() *AssignGroupOwnerRequestBody`

NewAssignGroupOwnerRequestBodyWithDefaults instantiates a new AssignGroupOwnerRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AssignGroupOwnerRequestBody) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AssignGroupOwnerRequestBody) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AssignGroupOwnerRequestBody) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AssignGroupOwnerRequestBody) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *AssignGroupOwnerRequestBody) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AssignGroupOwnerRequestBody) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AssignGroupOwnerRequestBody) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AssignGroupOwnerRequestBody) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


