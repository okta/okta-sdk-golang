# LogTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Detail** | Pointer to **map[string]map[string]interface{}** |  | [optional] [readonly] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Type** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewLogTransaction

`func NewLogTransaction() *LogTransaction`

NewLogTransaction instantiates a new LogTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogTransactionWithDefaults

`func NewLogTransactionWithDefaults() *LogTransaction`

NewLogTransactionWithDefaults instantiates a new LogTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetail

`func (o *LogTransaction) GetDetail() map[string]map[string]interface{}`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *LogTransaction) GetDetailOk() (*map[string]map[string]interface{}, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *LogTransaction) SetDetail(v map[string]map[string]interface{})`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *LogTransaction) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetId

`func (o *LogTransaction) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LogTransaction) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LogTransaction) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LogTransaction) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *LogTransaction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LogTransaction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LogTransaction) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *LogTransaction) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


