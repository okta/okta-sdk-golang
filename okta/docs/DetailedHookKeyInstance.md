# DetailedHookKeyInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **NullableTime** | Timestamp when the key was created | [optional] [readonly] 
**Id** | Pointer to **string** | The unique Okta ID of this key record | [optional] [readonly] 
**IsUsed** | Pointer to **bool** | Whether this key is currently in use by other applications | [optional] [readonly] 
**KeyId** | Pointer to **string** | The alias of the public key | [optional] [readonly] 
**LastUpdated** | Pointer to **NullableTime** | Timestamp when the key was updated | [optional] [readonly] 
**Name** | Pointer to **string** | Display name of the key | [optional] 
**Embedded** | Pointer to [**Embedded**](Embedded.md) |  | [optional] 

## Methods

### NewDetailedHookKeyInstance

`func NewDetailedHookKeyInstance() *DetailedHookKeyInstance`

NewDetailedHookKeyInstance instantiates a new DetailedHookKeyInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDetailedHookKeyInstanceWithDefaults

`func NewDetailedHookKeyInstanceWithDefaults() *DetailedHookKeyInstance`

NewDetailedHookKeyInstanceWithDefaults instantiates a new DetailedHookKeyInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *DetailedHookKeyInstance) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *DetailedHookKeyInstance) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *DetailedHookKeyInstance) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *DetailedHookKeyInstance) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreatedNil

`func (o *DetailedHookKeyInstance) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *DetailedHookKeyInstance) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetId

`func (o *DetailedHookKeyInstance) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DetailedHookKeyInstance) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DetailedHookKeyInstance) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DetailedHookKeyInstance) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsUsed

`func (o *DetailedHookKeyInstance) GetIsUsed() bool`

GetIsUsed returns the IsUsed field if non-nil, zero value otherwise.

### GetIsUsedOk

`func (o *DetailedHookKeyInstance) GetIsUsedOk() (*bool, bool)`

GetIsUsedOk returns a tuple with the IsUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUsed

`func (o *DetailedHookKeyInstance) SetIsUsed(v bool)`

SetIsUsed sets IsUsed field to given value.

### HasIsUsed

`func (o *DetailedHookKeyInstance) HasIsUsed() bool`

HasIsUsed returns a boolean if a field has been set.

### GetKeyId

`func (o *DetailedHookKeyInstance) GetKeyId() string`

GetKeyId returns the KeyId field if non-nil, zero value otherwise.

### GetKeyIdOk

`func (o *DetailedHookKeyInstance) GetKeyIdOk() (*string, bool)`

GetKeyIdOk returns a tuple with the KeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyId

`func (o *DetailedHookKeyInstance) SetKeyId(v string)`

SetKeyId sets KeyId field to given value.

### HasKeyId

`func (o *DetailedHookKeyInstance) HasKeyId() bool`

HasKeyId returns a boolean if a field has been set.

### GetLastUpdated

`func (o *DetailedHookKeyInstance) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *DetailedHookKeyInstance) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *DetailedHookKeyInstance) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *DetailedHookKeyInstance) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### SetLastUpdatedNil

`func (o *DetailedHookKeyInstance) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *DetailedHookKeyInstance) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetName

`func (o *DetailedHookKeyInstance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DetailedHookKeyInstance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DetailedHookKeyInstance) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DetailedHookKeyInstance) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEmbedded

`func (o *DetailedHookKeyInstance) GetEmbedded() Embedded`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *DetailedHookKeyInstance) GetEmbeddedOk() (*Embedded, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *DetailedHookKeyInstance) SetEmbedded(v Embedded)`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *DetailedHookKeyInstance) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


