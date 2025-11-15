# AgentJsonSigningKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**E** | Pointer to **string** | The public exponent of the RSA key, represented as a Base64URL-encoded string.  This value is used in combination with the modulus (&#x60;n&#x60;) to verify signatures and encrypt data. | [optional] 
**Kty** | **string** | Cryptographic algorithm family for the certificate&#39;s key pair | 
**N** | Pointer to **string** | The modulus of the RSA public key, represented as a Base64URL-encoded string.  This is the primary component of the RSA key and, with the exponent (&#x60;e&#x60;), is used for cryptographic signature verification and encryption. | [optional] 
**Kid** | Pointer to **string** | Unique identifier of the JSON Web Key in the AI agent&#39;s JSON Web Key Set (JWKS) | [optional] 
**Status** | Pointer to **string** | Status of the AI agent JSON Web Key | [optional] [default to "ACTIVE"]
**Alg** | Pointer to **string** | Algorithm that&#39;s used in the JSON Web Key | [optional] 
**Use** | Pointer to **string** | Acceptable use of the JSON Web Key  You can only use signing keys for AI agents, so the value of &#x60;use&#x60; is always &#x60;sig&#x60;. | [optional] 
**Crv** | Pointer to **string** | The cryptographic curve that&#39;s used for the key pair | [optional] 
**X** | Pointer to **string** | The public x coordinate for the elliptic curve point | [optional] 
**Y** | Pointer to **string** | The public y coordinate for the elliptic curve point | [optional] 

## Methods

### NewAgentJsonSigningKeyRequest

`func NewAgentJsonSigningKeyRequest(kty string, ) *AgentJsonSigningKeyRequest`

NewAgentJsonSigningKeyRequest instantiates a new AgentJsonSigningKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentJsonSigningKeyRequestWithDefaults

`func NewAgentJsonSigningKeyRequestWithDefaults() *AgentJsonSigningKeyRequest`

NewAgentJsonSigningKeyRequestWithDefaults instantiates a new AgentJsonSigningKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetE

`func (o *AgentJsonSigningKeyRequest) GetE() string`

GetE returns the E field if non-nil, zero value otherwise.

### GetEOk

`func (o *AgentJsonSigningKeyRequest) GetEOk() (*string, bool)`

GetEOk returns a tuple with the E field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE

`func (o *AgentJsonSigningKeyRequest) SetE(v string)`

SetE sets E field to given value.

### HasE

`func (o *AgentJsonSigningKeyRequest) HasE() bool`

HasE returns a boolean if a field has been set.

### GetKty

`func (o *AgentJsonSigningKeyRequest) GetKty() string`

GetKty returns the Kty field if non-nil, zero value otherwise.

### GetKtyOk

`func (o *AgentJsonSigningKeyRequest) GetKtyOk() (*string, bool)`

GetKtyOk returns a tuple with the Kty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKty

`func (o *AgentJsonSigningKeyRequest) SetKty(v string)`

SetKty sets Kty field to given value.


### GetN

`func (o *AgentJsonSigningKeyRequest) GetN() string`

GetN returns the N field if non-nil, zero value otherwise.

### GetNOk

`func (o *AgentJsonSigningKeyRequest) GetNOk() (*string, bool)`

GetNOk returns a tuple with the N field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN

`func (o *AgentJsonSigningKeyRequest) SetN(v string)`

SetN sets N field to given value.

### HasN

`func (o *AgentJsonSigningKeyRequest) HasN() bool`

HasN returns a boolean if a field has been set.

### GetKid

`func (o *AgentJsonSigningKeyRequest) GetKid() string`

GetKid returns the Kid field if non-nil, zero value otherwise.

### GetKidOk

`func (o *AgentJsonSigningKeyRequest) GetKidOk() (*string, bool)`

GetKidOk returns a tuple with the Kid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKid

`func (o *AgentJsonSigningKeyRequest) SetKid(v string)`

SetKid sets Kid field to given value.

### HasKid

`func (o *AgentJsonSigningKeyRequest) HasKid() bool`

HasKid returns a boolean if a field has been set.

### GetStatus

`func (o *AgentJsonSigningKeyRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AgentJsonSigningKeyRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AgentJsonSigningKeyRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AgentJsonSigningKeyRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAlg

`func (o *AgentJsonSigningKeyRequest) GetAlg() string`

GetAlg returns the Alg field if non-nil, zero value otherwise.

### GetAlgOk

`func (o *AgentJsonSigningKeyRequest) GetAlgOk() (*string, bool)`

GetAlgOk returns a tuple with the Alg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlg

`func (o *AgentJsonSigningKeyRequest) SetAlg(v string)`

SetAlg sets Alg field to given value.

### HasAlg

`func (o *AgentJsonSigningKeyRequest) HasAlg() bool`

HasAlg returns a boolean if a field has been set.

### GetUse

`func (o *AgentJsonSigningKeyRequest) GetUse() string`

GetUse returns the Use field if non-nil, zero value otherwise.

### GetUseOk

`func (o *AgentJsonSigningKeyRequest) GetUseOk() (*string, bool)`

GetUseOk returns a tuple with the Use field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUse

`func (o *AgentJsonSigningKeyRequest) SetUse(v string)`

SetUse sets Use field to given value.

### HasUse

`func (o *AgentJsonSigningKeyRequest) HasUse() bool`

HasUse returns a boolean if a field has been set.

### GetCrv

`func (o *AgentJsonSigningKeyRequest) GetCrv() string`

GetCrv returns the Crv field if non-nil, zero value otherwise.

### GetCrvOk

`func (o *AgentJsonSigningKeyRequest) GetCrvOk() (*string, bool)`

GetCrvOk returns a tuple with the Crv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrv

`func (o *AgentJsonSigningKeyRequest) SetCrv(v string)`

SetCrv sets Crv field to given value.

### HasCrv

`func (o *AgentJsonSigningKeyRequest) HasCrv() bool`

HasCrv returns a boolean if a field has been set.

### GetX

`func (o *AgentJsonSigningKeyRequest) GetX() string`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *AgentJsonSigningKeyRequest) GetXOk() (*string, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *AgentJsonSigningKeyRequest) SetX(v string)`

SetX sets X field to given value.

### HasX

`func (o *AgentJsonSigningKeyRequest) HasX() bool`

HasX returns a boolean if a field has been set.

### GetY

`func (o *AgentJsonSigningKeyRequest) GetY() string`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *AgentJsonSigningKeyRequest) GetYOk() (*string, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *AgentJsonSigningKeyRequest) SetY(v string)`

SetY sets Y field to given value.

### HasY

`func (o *AgentJsonSigningKeyRequest) HasY() bool`

HasY returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


