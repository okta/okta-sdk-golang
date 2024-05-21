# PasswordCredentialHash

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Algorithm** | Pointer to **string** | The algorithm used to generate the hash using the password (and salt, when applicable). | [optional] 
**DigestAlgorithm** | Pointer to **string** | Algorithm used to generate the key. Only required for the PBKDF2 algorithm. | [optional] 
**IterationCount** | Pointer to **int32** | The number of iterations used when hashing passwords using PBKDF2. Must be &gt;&#x3D; 4096. Only required for PBKDF2 algorithm. | [optional] 
**KeySize** | Pointer to **int32** | Size of the derived key in bytes. Only required for PBKDF2 algorithm. | [optional] 
**Salt** | Pointer to **string** | Only required for salted hashes. For BCRYPT, this specifies Radix-64 as the encoded salt used to generate the hash, which must be 22 characters long. For other salted hashes, this specifies the Base64-encoded salt used to generate the hash. | [optional] 
**SaltOrder** | Pointer to **string** | Specifies whether salt was pre- or postfixed to the password before hashing. Only required for salted algorithms. | [optional] 
**Value** | Pointer to **string** | For SHA-512, SHA-256, SHA-1, MD5, and PBKDF2, this is the actual base64-encoded hash of the password (and salt, if used). This is the Base64-encoded &#x60;value&#x60; of the SHA-512/SHA-256/SHA-1/MD5/PBKDF2 digest that was computed by either pre-fixing or post-fixing the &#x60;salt&#x60; to the &#x60;password&#x60;, depending on the &#x60;saltOrder&#x60;. If a &#x60;salt&#x60; was not used in the &#x60;source&#x60; system, then this should just be the Base64-encoded &#x60;value&#x60; of the password&#39;s SHA-512/SHA-256/SHA-1/MD5/PBKDF2 digest. For BCRYPT, this is the actual Radix-64 encoded hashed password. | [optional] 
**WorkFactor** | Pointer to **int32** | Governs the strength of the hash and the time required to compute it. Only required for BCRYPT algorithm. | [optional] 

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

`func (o *PasswordCredentialHash) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *PasswordCredentialHash) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *PasswordCredentialHash) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.

### HasAlgorithm

`func (o *PasswordCredentialHash) HasAlgorithm() bool`

HasAlgorithm returns a boolean if a field has been set.

### GetDigestAlgorithm

`func (o *PasswordCredentialHash) GetDigestAlgorithm() string`

GetDigestAlgorithm returns the DigestAlgorithm field if non-nil, zero value otherwise.

### GetDigestAlgorithmOk

`func (o *PasswordCredentialHash) GetDigestAlgorithmOk() (*string, bool)`

GetDigestAlgorithmOk returns a tuple with the DigestAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigestAlgorithm

`func (o *PasswordCredentialHash) SetDigestAlgorithm(v string)`

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


