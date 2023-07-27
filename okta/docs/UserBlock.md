# UserBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppliesTo** | Pointer to **string** |  | [optional] [readonly] 
**Type** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewUserBlock

`func NewUserBlock() *UserBlock`

NewUserBlock instantiates a new UserBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserBlockWithDefaults

`func NewUserBlockWithDefaults() *UserBlock`

NewUserBlockWithDefaults instantiates a new UserBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppliesTo

`func (o *UserBlock) GetAppliesTo() string`

GetAppliesTo returns the AppliesTo field if non-nil, zero value otherwise.

### GetAppliesToOk

`func (o *UserBlock) GetAppliesToOk() (*string, bool)`

GetAppliesToOk returns a tuple with the AppliesTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliesTo

`func (o *UserBlock) SetAppliesTo(v string)`

SetAppliesTo sets AppliesTo field to given value.

### HasAppliesTo

`func (o *UserBlock) HasAppliesTo() bool`

HasAppliesTo returns a boolean if a field has been set.

### GetType

`func (o *UserBlock) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserBlock) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserBlock) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UserBlock) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


