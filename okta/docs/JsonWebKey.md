# JsonWebKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alg** | Pointer to **string** | The algorithm used with the Key. Valid value: &#x60;RS256&#x60; | [optional] 
**Created** | Pointer to **time.Time** | Timestamp when the object was created | [optional] [readonly] 
**E** | Pointer to **string** | RSA key value (public exponent) for Key binding | [optional] [readonly] 
**ExpiresAt** | Pointer to **time.Time** | Timestamp when the certificate expires | [optional] [readonly] 
**KeyOps** | Pointer to **[]string** | Identifies the operation(s) for which the key is intended to be used | [optional] 
**Kid** | Pointer to **string** | Unique identifier for the certificate | [optional] [readonly] 
**Kty** | Pointer to **string** | Cryptographic algorithm family for the certificate&#39;s keypair. Valid value: &#x60;RSA&#x60; | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** | Timestamp when the object was last updated | [optional] [readonly] 
**N** | Pointer to **string** | RSA modulus value that is used by both the public and private keys and provides a link between them | [optional] 
**Status** | Pointer to **string** | An &#x60;ACTIVE&#x60; Key is used to sign tokens issued by the authorization server. Supported values: &#x60;ACTIVE&#x60;, &#x60;NEXT&#x60;, or &#x60;EXPIRED&#x60;&lt;br&gt; A &#x60;NEXT&#x60; Key is the next Key that the authorization server uses to sign tokens when Keys are rotated. The &#x60;NEXT&#x60; Key might not be listed if it hasn&#39;t been generated yet. An &#x60;EXPIRED&#x60; Key is the previous Key that the authorization server used to sign tokens. The &#x60;EXPIRED&#x60; Key might not be listed if no Key has expired or the expired Key was deleted. | [optional] 
**Use** | Pointer to **string** | Acceptable use of the certificate. Valid value: &#x60;sig&#x60; | [optional] [readonly] 
**X5c** | Pointer to **[]string** | X.509 certificate chain that contains a chain of one or more certificates | [optional] [readonly] 
**X5t** | Pointer to **string** | X.509 certificate SHA-1 thumbprint, which is the base64url-encoded SHA-1 thumbprint (digest) of the DER encoding of an X.509 certificate | [optional] [readonly] 
**X5tS256** | Pointer to **string** | X.509 certificate SHA-256 thumbprint, which is the base64url-encoded SHA-256 thumbprint (digest) of the DER encoding of an X.509 certificate | [optional] [readonly] 
**X5u** | Pointer to **string** | A URI that refers to a resource for the X.509 public key certificate or certificate chain corresponding to the key used to digitally sign the JWS (JSON Web Signature) | [optional] [readonly] 
**Links** | Pointer to [**LinksSelf**](LinksSelf.md) |  | [optional] 

## Methods

### NewJsonWebKey

`func NewJsonWebKey() *JsonWebKey`

NewJsonWebKey instantiates a new JsonWebKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJsonWebKeyWithDefaults

`func NewJsonWebKeyWithDefaults() *JsonWebKey`

NewJsonWebKeyWithDefaults instantiates a new JsonWebKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlg

`func (o *JsonWebKey) GetAlg() string`

GetAlg returns the Alg field if non-nil, zero value otherwise.

### GetAlgOk

`func (o *JsonWebKey) GetAlgOk() (*string, bool)`

GetAlgOk returns a tuple with the Alg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlg

`func (o *JsonWebKey) SetAlg(v string)`

SetAlg sets Alg field to given value.

### HasAlg

`func (o *JsonWebKey) HasAlg() bool`

HasAlg returns a boolean if a field has been set.

### GetCreated

`func (o *JsonWebKey) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *JsonWebKey) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *JsonWebKey) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *JsonWebKey) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetE

`func (o *JsonWebKey) GetE() string`

GetE returns the E field if non-nil, zero value otherwise.

### GetEOk

`func (o *JsonWebKey) GetEOk() (*string, bool)`

GetEOk returns a tuple with the E field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE

`func (o *JsonWebKey) SetE(v string)`

SetE sets E field to given value.

### HasE

`func (o *JsonWebKey) HasE() bool`

HasE returns a boolean if a field has been set.

### GetExpiresAt

`func (o *JsonWebKey) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *JsonWebKey) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *JsonWebKey) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *JsonWebKey) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetKeyOps

`func (o *JsonWebKey) GetKeyOps() []string`

GetKeyOps returns the KeyOps field if non-nil, zero value otherwise.

### GetKeyOpsOk

`func (o *JsonWebKey) GetKeyOpsOk() (*[]string, bool)`

GetKeyOpsOk returns a tuple with the KeyOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyOps

`func (o *JsonWebKey) SetKeyOps(v []string)`

SetKeyOps sets KeyOps field to given value.

### HasKeyOps

`func (o *JsonWebKey) HasKeyOps() bool`

HasKeyOps returns a boolean if a field has been set.

### GetKid

`func (o *JsonWebKey) GetKid() string`

GetKid returns the Kid field if non-nil, zero value otherwise.

### GetKidOk

`func (o *JsonWebKey) GetKidOk() (*string, bool)`

GetKidOk returns a tuple with the Kid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKid

`func (o *JsonWebKey) SetKid(v string)`

SetKid sets Kid field to given value.

### HasKid

`func (o *JsonWebKey) HasKid() bool`

HasKid returns a boolean if a field has been set.

### GetKty

`func (o *JsonWebKey) GetKty() string`

GetKty returns the Kty field if non-nil, zero value otherwise.

### GetKtyOk

`func (o *JsonWebKey) GetKtyOk() (*string, bool)`

GetKtyOk returns a tuple with the Kty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKty

`func (o *JsonWebKey) SetKty(v string)`

SetKty sets Kty field to given value.

### HasKty

`func (o *JsonWebKey) HasKty() bool`

HasKty returns a boolean if a field has been set.

### GetLastUpdated

`func (o *JsonWebKey) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *JsonWebKey) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *JsonWebKey) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *JsonWebKey) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetN

`func (o *JsonWebKey) GetN() string`

GetN returns the N field if non-nil, zero value otherwise.

### GetNOk

`func (o *JsonWebKey) GetNOk() (*string, bool)`

GetNOk returns a tuple with the N field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN

`func (o *JsonWebKey) SetN(v string)`

SetN sets N field to given value.

### HasN

`func (o *JsonWebKey) HasN() bool`

HasN returns a boolean if a field has been set.

### GetStatus

`func (o *JsonWebKey) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *JsonWebKey) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *JsonWebKey) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *JsonWebKey) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUse

`func (o *JsonWebKey) GetUse() string`

GetUse returns the Use field if non-nil, zero value otherwise.

### GetUseOk

`func (o *JsonWebKey) GetUseOk() (*string, bool)`

GetUseOk returns a tuple with the Use field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUse

`func (o *JsonWebKey) SetUse(v string)`

SetUse sets Use field to given value.

### HasUse

`func (o *JsonWebKey) HasUse() bool`

HasUse returns a boolean if a field has been set.

### GetX5c

`func (o *JsonWebKey) GetX5c() []string`

GetX5c returns the X5c field if non-nil, zero value otherwise.

### GetX5cOk

`func (o *JsonWebKey) GetX5cOk() (*[]string, bool)`

GetX5cOk returns a tuple with the X5c field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX5c

`func (o *JsonWebKey) SetX5c(v []string)`

SetX5c sets X5c field to given value.

### HasX5c

`func (o *JsonWebKey) HasX5c() bool`

HasX5c returns a boolean if a field has been set.

### GetX5t

`func (o *JsonWebKey) GetX5t() string`

GetX5t returns the X5t field if non-nil, zero value otherwise.

### GetX5tOk

`func (o *JsonWebKey) GetX5tOk() (*string, bool)`

GetX5tOk returns a tuple with the X5t field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX5t

`func (o *JsonWebKey) SetX5t(v string)`

SetX5t sets X5t field to given value.

### HasX5t

`func (o *JsonWebKey) HasX5t() bool`

HasX5t returns a boolean if a field has been set.

### GetX5tS256

`func (o *JsonWebKey) GetX5tS256() string`

GetX5tS256 returns the X5tS256 field if non-nil, zero value otherwise.

### GetX5tS256Ok

`func (o *JsonWebKey) GetX5tS256Ok() (*string, bool)`

GetX5tS256Ok returns a tuple with the X5tS256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX5tS256

`func (o *JsonWebKey) SetX5tS256(v string)`

SetX5tS256 sets X5tS256 field to given value.

### HasX5tS256

`func (o *JsonWebKey) HasX5tS256() bool`

HasX5tS256 returns a boolean if a field has been set.

### GetX5u

`func (o *JsonWebKey) GetX5u() string`

GetX5u returns the X5u field if non-nil, zero value otherwise.

### GetX5uOk

`func (o *JsonWebKey) GetX5uOk() (*string, bool)`

GetX5uOk returns a tuple with the X5u field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX5u

`func (o *JsonWebKey) SetX5u(v string)`

SetX5u sets X5u field to given value.

### HasX5u

`func (o *JsonWebKey) HasX5u() bool`

HasX5u returns a boolean if a field has been set.

### GetLinks

`func (o *JsonWebKey) GetLinks() LinksSelf`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *JsonWebKey) GetLinksOk() (*LinksSelf, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *JsonWebKey) SetLinks(v LinksSelf)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *JsonWebKey) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


