# HookKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **NullableTime** | Timestamp when the key was created | [optional] [readonly] 
**Id** | Pointer to **string** | The unique identifier for the key | [optional] [readonly] 
**IsUsed** | Pointer to **bool** | Whether this key is currently in use by other applications | [optional] [readonly] 
**KeyId** | Pointer to **string** | The alias of the public key | [optional] [readonly] 
**LastUpdated** | Pointer to **NullableTime** | Timestamp when the key was updated | [optional] [readonly] 
**Name** | Pointer to **string** | Display name of the key | [optional] 

## Methods

### NewHookKey

`func NewHookKey() *HookKey`

NewHookKey instantiates a new HookKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHookKeyWithDefaults

`func NewHookKeyWithDefaults() *HookKey`

NewHookKeyWithDefaults instantiates a new HookKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *HookKey) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *HookKey) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *HookKey) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *HookKey) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreatedNil

`func (o *HookKey) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *HookKey) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetId

`func (o *HookKey) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HookKey) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HookKey) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HookKey) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsUsed

`func (o *HookKey) GetIsUsed() bool`

GetIsUsed returns the IsUsed field if non-nil, zero value otherwise.

### GetIsUsedOk

`func (o *HookKey) GetIsUsedOk() (*bool, bool)`

GetIsUsedOk returns a tuple with the IsUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUsed

`func (o *HookKey) SetIsUsed(v bool)`

SetIsUsed sets IsUsed field to given value.

### HasIsUsed

`func (o *HookKey) HasIsUsed() bool`

HasIsUsed returns a boolean if a field has been set.

### GetKeyId

`func (o *HookKey) GetKeyId() string`

GetKeyId returns the KeyId field if non-nil, zero value otherwise.

### GetKeyIdOk

`func (o *HookKey) GetKeyIdOk() (*string, bool)`

GetKeyIdOk returns a tuple with the KeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyId

`func (o *HookKey) SetKeyId(v string)`

SetKeyId sets KeyId field to given value.

### HasKeyId

`func (o *HookKey) HasKeyId() bool`

HasKeyId returns a boolean if a field has been set.

### GetLastUpdated

`func (o *HookKey) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *HookKey) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *HookKey) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *HookKey) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### SetLastUpdatedNil

`func (o *HookKey) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *HookKey) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetName

`func (o *HookKey) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HookKey) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HookKey) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HookKey) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


