# UserSchemaAttributeMaster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Priority** | Pointer to [**[]UserSchemaAttributeMasterPriority**](UserSchemaAttributeMasterPriority.md) |  | [optional] 
**Type** | Pointer to [**UserSchemaAttributeMasterType**](UserSchemaAttributeMasterType.md) |  | [optional] 

## Methods

### NewUserSchemaAttributeMaster

`func NewUserSchemaAttributeMaster() *UserSchemaAttributeMaster`

NewUserSchemaAttributeMaster instantiates a new UserSchemaAttributeMaster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSchemaAttributeMasterWithDefaults

`func NewUserSchemaAttributeMasterWithDefaults() *UserSchemaAttributeMaster`

NewUserSchemaAttributeMasterWithDefaults instantiates a new UserSchemaAttributeMaster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPriority

`func (o *UserSchemaAttributeMaster) GetPriority() []UserSchemaAttributeMasterPriority`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *UserSchemaAttributeMaster) GetPriorityOk() (*[]UserSchemaAttributeMasterPriority, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *UserSchemaAttributeMaster) SetPriority(v []UserSchemaAttributeMasterPriority)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *UserSchemaAttributeMaster) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetType

`func (o *UserSchemaAttributeMaster) GetType() UserSchemaAttributeMasterType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserSchemaAttributeMaster) GetTypeOk() (*UserSchemaAttributeMasterType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserSchemaAttributeMaster) SetType(v UserSchemaAttributeMasterType)`

SetType sets Type field to given value.

### HasType

`func (o *UserSchemaAttributeMaster) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


