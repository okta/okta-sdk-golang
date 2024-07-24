# AuthorizationServerJsonWebKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alg** | Pointer to **string** | The algorithm used with the Key. Valid value: &#x60;RS256&#x60; | [optional] 
**E** | Pointer to **string** | RSA key value (public exponent) for Key binding | [optional] [readonly] 
**Kid** | Pointer to **string** | Unique identifier for the key | [optional] [readonly] 
**Kty** | Pointer to **string** | Cryptographic algorithm family for the certificate&#39;s keypair. Valid value: &#x60;RSA&#x60; | [optional] [readonly] 
**N** | Pointer to **string** | RSA modulus value that is used by both the public and private keys and provides a link between them | [optional] 
**Status** | Pointer to **string** | An &#x60;ACTIVE&#x60; Key is used to sign tokens issued by the authorization server. Supported values: &#x60;ACTIVE&#x60;, &#x60;NEXT&#x60;, or &#x60;EXPIRED&#x60;&lt;br&gt; A &#x60;NEXT&#x60; Key is the next Key that the authorization server uses to sign tokens when Keys are rotated. The &#x60;NEXT&#x60; Key might not be listed if it hasn&#39;t been generated. An &#x60;EXPIRED&#x60; Key is the previous Key that the authorization server used to sign tokens. The &#x60;EXPIRED&#x60; Key might not be listed if no Key has expired or the expired Key was deleted. | [optional] 
**Use** | Pointer to **string** | Acceptable use of the key. Valid value: &#x60;sig&#x60; | [optional] [readonly] 
**Links** | Pointer to [**LinksSelf**](LinksSelf.md) |  | [optional] 

## Methods

### NewAuthorizationServerJsonWebKey

`func NewAuthorizationServerJsonWebKey() *AuthorizationServerJsonWebKey`

NewAuthorizationServerJsonWebKey instantiates a new AuthorizationServerJsonWebKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizationServerJsonWebKeyWithDefaults

`func NewAuthorizationServerJsonWebKeyWithDefaults() *AuthorizationServerJsonWebKey`

NewAuthorizationServerJsonWebKeyWithDefaults instantiates a new AuthorizationServerJsonWebKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlg

`func (o *AuthorizationServerJsonWebKey) GetAlg() string`

GetAlg returns the Alg field if non-nil, zero value otherwise.

### GetAlgOk

`func (o *AuthorizationServerJsonWebKey) GetAlgOk() (*string, bool)`

GetAlgOk returns a tuple with the Alg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlg

`func (o *AuthorizationServerJsonWebKey) SetAlg(v string)`

SetAlg sets Alg field to given value.

### HasAlg

`func (o *AuthorizationServerJsonWebKey) HasAlg() bool`

HasAlg returns a boolean if a field has been set.

### GetE

`func (o *AuthorizationServerJsonWebKey) GetE() string`

GetE returns the E field if non-nil, zero value otherwise.

### GetEOk

`func (o *AuthorizationServerJsonWebKey) GetEOk() (*string, bool)`

GetEOk returns a tuple with the E field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE

`func (o *AuthorizationServerJsonWebKey) SetE(v string)`

SetE sets E field to given value.

### HasE

`func (o *AuthorizationServerJsonWebKey) HasE() bool`

HasE returns a boolean if a field has been set.

### GetKid

`func (o *AuthorizationServerJsonWebKey) GetKid() string`

GetKid returns the Kid field if non-nil, zero value otherwise.

### GetKidOk

`func (o *AuthorizationServerJsonWebKey) GetKidOk() (*string, bool)`

GetKidOk returns a tuple with the Kid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKid

`func (o *AuthorizationServerJsonWebKey) SetKid(v string)`

SetKid sets Kid field to given value.

### HasKid

`func (o *AuthorizationServerJsonWebKey) HasKid() bool`

HasKid returns a boolean if a field has been set.

### GetKty

`func (o *AuthorizationServerJsonWebKey) GetKty() string`

GetKty returns the Kty field if non-nil, zero value otherwise.

### GetKtyOk

`func (o *AuthorizationServerJsonWebKey) GetKtyOk() (*string, bool)`

GetKtyOk returns a tuple with the Kty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKty

`func (o *AuthorizationServerJsonWebKey) SetKty(v string)`

SetKty sets Kty field to given value.

### HasKty

`func (o *AuthorizationServerJsonWebKey) HasKty() bool`

HasKty returns a boolean if a field has been set.

### GetN

`func (o *AuthorizationServerJsonWebKey) GetN() string`

GetN returns the N field if non-nil, zero value otherwise.

### GetNOk

`func (o *AuthorizationServerJsonWebKey) GetNOk() (*string, bool)`

GetNOk returns a tuple with the N field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN

`func (o *AuthorizationServerJsonWebKey) SetN(v string)`

SetN sets N field to given value.

### HasN

`func (o *AuthorizationServerJsonWebKey) HasN() bool`

HasN returns a boolean if a field has been set.

### GetStatus

`func (o *AuthorizationServerJsonWebKey) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AuthorizationServerJsonWebKey) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AuthorizationServerJsonWebKey) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AuthorizationServerJsonWebKey) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUse

`func (o *AuthorizationServerJsonWebKey) GetUse() string`

GetUse returns the Use field if non-nil, zero value otherwise.

### GetUseOk

`func (o *AuthorizationServerJsonWebKey) GetUseOk() (*string, bool)`

GetUseOk returns a tuple with the Use field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUse

`func (o *AuthorizationServerJsonWebKey) SetUse(v string)`

SetUse sets Use field to given value.

### HasUse

`func (o *AuthorizationServerJsonWebKey) HasUse() bool`

HasUse returns a boolean if a field has been set.

### GetLinks

`func (o *AuthorizationServerJsonWebKey) GetLinks() LinksSelf`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AuthorizationServerJsonWebKey) GetLinksOk() (*LinksSelf, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AuthorizationServerJsonWebKey) SetLinks(v LinksSelf)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AuthorizationServerJsonWebKey) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


