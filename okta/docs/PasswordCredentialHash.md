# PasswordCredentialHash

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Algorithm** | Pointer to [**PasswordCredentialHashAlgorithm**](PasswordCredentialHashAlgorithm.md) |  | [optional] 
**DigestAlgorithm** | Pointer to [**DigestAlgorithm**](DigestAlgorithm.md) |  | [optional] 
**IterationCount** | Pointer to **int32** |  | [optional] 
**KeySize** | Pointer to **int32** |  | [optional] 
**Salt** | Pointer to **string** |  | [optional] 
**SaltOrder** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 
**WorkFactor** | Pointer to **int32** |  | [optional] 

## Methods

### NewPasswordCredentialHash

`func NewPasswordCredentialHash() *PasswordCredentialHash`

NewPasswordCredentialHash instantiates a new PasswordCredentialHash object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPasswordCredentialHashWithDefaults

`func NewPasswordCredentialHashWithDefaults() *PasswordCredentialHash`

NewPasswordCredentialHashWithDefaults instantiates a new PasswordCredentialHash object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgorithm

`func (o *PasswordCredentialHash) GetAlgorithm() PasswordCredentialHashAlgorithm`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *PasswordCredentialHash) GetAlgorithmOk() (*PasswordCredentialHashAlgorithm, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *PasswordCredentialHash) SetAlgorithm(v PasswordCredentialHashAlgorithm)`

SetAlgorithm sets Algorithm field to given value.

### HasAlgorithm

`func (o *PasswordCredentialHash) HasAlgorithm() bool`

HasAlgorithm returns a boolean if a field has been set.

### GetDigestAlgorithm

`func (o *PasswordCredentialHash) GetDigestAlgorithm() DigestAlgorithm`

GetDigestAlgorithm returns the DigestAlgorithm field if non-nil, zero value otherwise.

### GetDigestAlgorithmOk

`func (o *PasswordCredentialHash) GetDigestAlgorithmOk() (*DigestAlgorithm, bool)`

GetDigestAlgorithmOk returns a tuple with the DigestAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigestAlgorithm

`func (o *PasswordCredentialHash) SetDigestAlgorithm(v DigestAlgorithm)`

SetDigestAlgorithm sets DigestAlgorithm field to given value.

### HasDigestAlgorithm

`func (o *PasswordCredentialHash) HasDigestAlgorithm() bool`

HasDigestAlgorithm returns a boolean if a field has been set.

### GetIterationCount

`func (o *PasswordCredentialHash) GetIterationCount() int32`

GetIterationCount returns the IterationCount field if non-nil, zero value otherwise.

### GetIterationCountOk

`func (o *PasswordCredentialHash) GetIterationCountOk() (*int32, bool)`

GetIterationCountOk returns a tuple with the IterationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIterationCount

`func (o *PasswordCredentialHash) SetIterationCount(v int32)`

SetIterationCount sets IterationCount field to given value.

### HasIterationCount

`func (o *PasswordCredentialHash) HasIterationCount() bool`

HasIterationCount returns a boolean if a field has been set.

### GetKeySize

`func (o *PasswordCredentialHash) GetKeySize() int32`

GetKeySize returns the KeySize field if non-nil, zero value otherwise.

### GetKeySizeOk

`func (o *PasswordCredentialHash) GetKeySizeOk() (*int32, bool)`

GetKeySizeOk returns a tuple with the KeySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySize

`func (o *PasswordCredentialHash) SetKeySize(v int32)`

SetKeySize sets KeySize field to given value.

### HasKeySize

`func (o *PasswordCredentialHash) HasKeySize() bool`

HasKeySize returns a boolean if a field has been set.

### GetSalt

`func (o *PasswordCredentialHash) GetSalt() string`

GetSalt returns the Salt field if non-nil, zero value otherwise.

### GetSaltOk

`func (o *PasswordCredentialHash) GetSaltOk() (*string, bool)`

GetSaltOk returns a tuple with the Salt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalt

`func (o *PasswordCredentialHash) SetSalt(v string)`

SetSalt sets Salt field to given value.

### HasSalt

`func (o *PasswordCredentialHash) HasSalt() bool`

HasSalt returns a boolean if a field has been set.

### GetSaltOrder

`func (o *PasswordCredentialHash) GetSaltOrder() string`

GetSaltOrder returns the SaltOrder field if non-nil, zero value otherwise.

### GetSaltOrderOk

`func (o *PasswordCredentialHash) GetSaltOrderOk() (*string, bool)`

GetSaltOrderOk returns a tuple with the SaltOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaltOrder

`func (o *PasswordCredentialHash) SetSaltOrder(v string)`

SetSaltOrder sets SaltOrder field to given value.

### HasSaltOrder

`func (o *PasswordCredentialHash) HasSaltOrder() bool`

HasSaltOrder returns a boolean if a field has been set.

### GetValue

`func (o *PasswordCredentialHash) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PasswordCredentialHash) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PasswordCredentialHash) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *PasswordCredentialHash) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetWorkFactor

`func (o *PasswordCredentialHash) GetWorkFactor() int32`

GetWorkFactor returns the WorkFactor field if non-nil, zero value otherwise.

### GetWorkFactorOk

`func (o *PasswordCredentialHash) GetWorkFactorOk() (*int32, bool)`

GetWorkFactorOk returns a tuple with the WorkFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkFactor

`func (o *PasswordCredentialHash) SetWorkFactor(v int32)`

SetWorkFactor sets WorkFactor field to given value.

### HasWorkFactor

`func (o *PasswordCredentialHash) HasWorkFactor() bool`

HasWorkFactor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


