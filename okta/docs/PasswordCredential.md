# PasswordCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hash** | Pointer to [**PasswordCredentialHash**](PasswordCredentialHash.md) |  | [optional] 
**Hook** | Pointer to [**PasswordCredentialHook**](PasswordCredentialHook.md) |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 

## Methods

### NewPasswordCredential

`func NewPasswordCredential() *PasswordCredential`

NewPasswordCredential instantiates a new PasswordCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPasswordCredentialWithDefaults

`func NewPasswordCredentialWithDefaults() *PasswordCredential`

NewPasswordCredentialWithDefaults instantiates a new PasswordCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHash

`func (o *PasswordCredential) GetHash() PasswordCredentialHash`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *PasswordCredential) GetHashOk() (*PasswordCredentialHash, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *PasswordCredential) SetHash(v PasswordCredentialHash)`

SetHash sets Hash field to given value.

### HasHash

`func (o *PasswordCredential) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetHook

`func (o *PasswordCredential) GetHook() PasswordCredentialHook`

GetHook returns the Hook field if non-nil, zero value otherwise.

### GetHookOk

`func (o *PasswordCredential) GetHookOk() (*PasswordCredentialHook, bool)`

GetHookOk returns a tuple with the Hook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHook

`func (o *PasswordCredential) SetHook(v PasswordCredentialHook)`

SetHook sets Hook field to given value.

### HasHook

`func (o *PasswordCredential) HasHook() bool`

HasHook returns a boolean if a field has been set.

### GetValue

`func (o *PasswordCredential) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PasswordCredential) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PasswordCredential) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *PasswordCredential) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


