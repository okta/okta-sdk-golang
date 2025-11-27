# AgentJsonWebKeyRsaRequest

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

## Methods

### NewAgentJsonWebKeyRsaRequest

`func NewAgentJsonWebKeyRsaRequest(kty string, ) *AgentJsonWebKeyRsaRequest`

NewAgentJsonWebKeyRsaRequest instantiates a new AgentJsonWebKeyRsaRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentJsonWebKeyRsaRequestWithDefaults

`func NewAgentJsonWebKeyRsaRequestWithDefaults() *AgentJsonWebKeyRsaRequest`

NewAgentJsonWebKeyRsaRequestWithDefaults instantiates a new AgentJsonWebKeyRsaRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetE

`func (o *AgentJsonWebKeyRsaRequest) GetE() string`

GetE returns the E field if non-nil, zero value otherwise.

### GetEOk

`func (o *AgentJsonWebKeyRsaRequest) GetEOk() (*string, bool)`

GetEOk returns a tuple with the E field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE

`func (o *AgentJsonWebKeyRsaRequest) SetE(v string)`

SetE sets E field to given value.

### HasE

`func (o *AgentJsonWebKeyRsaRequest) HasE() bool`

HasE returns a boolean if a field has been set.

### GetKty

`func (o *AgentJsonWebKeyRsaRequest) GetKty() string`

GetKty returns the Kty field if non-nil, zero value otherwise.

### GetKtyOk

`func (o *AgentJsonWebKeyRsaRequest) GetKtyOk() (*string, bool)`

GetKtyOk returns a tuple with the Kty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKty

`func (o *AgentJsonWebKeyRsaRequest) SetKty(v string)`

SetKty sets Kty field to given value.


### GetN

`func (o *AgentJsonWebKeyRsaRequest) GetN() string`

GetN returns the N field if non-nil, zero value otherwise.

### GetNOk

`func (o *AgentJsonWebKeyRsaRequest) GetNOk() (*string, bool)`

GetNOk returns a tuple with the N field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN

`func (o *AgentJsonWebKeyRsaRequest) SetN(v string)`

SetN sets N field to given value.

### HasN

`func (o *AgentJsonWebKeyRsaRequest) HasN() bool`

HasN returns a boolean if a field has been set.

### GetKid

`func (o *AgentJsonWebKeyRsaRequest) GetKid() string`

GetKid returns the Kid field if non-nil, zero value otherwise.

### GetKidOk

`func (o *AgentJsonWebKeyRsaRequest) GetKidOk() (*string, bool)`

GetKidOk returns a tuple with the Kid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKid

`func (o *AgentJsonWebKeyRsaRequest) SetKid(v string)`

SetKid sets Kid field to given value.

### HasKid

`func (o *AgentJsonWebKeyRsaRequest) HasKid() bool`

HasKid returns a boolean if a field has been set.

### GetStatus

`func (o *AgentJsonWebKeyRsaRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AgentJsonWebKeyRsaRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AgentJsonWebKeyRsaRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AgentJsonWebKeyRsaRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAlg

`func (o *AgentJsonWebKeyRsaRequest) GetAlg() string`

GetAlg returns the Alg field if non-nil, zero value otherwise.

### GetAlgOk

`func (o *AgentJsonWebKeyRsaRequest) GetAlgOk() (*string, bool)`

GetAlgOk returns a tuple with the Alg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlg

`func (o *AgentJsonWebKeyRsaRequest) SetAlg(v string)`

SetAlg sets Alg field to given value.

### HasAlg

`func (o *AgentJsonWebKeyRsaRequest) HasAlg() bool`

HasAlg returns a boolean if a field has been set.

### GetUse

`func (o *AgentJsonWebKeyRsaRequest) GetUse() string`

GetUse returns the Use field if non-nil, zero value otherwise.

### GetUseOk

`func (o *AgentJsonWebKeyRsaRequest) GetUseOk() (*string, bool)`

GetUseOk returns a tuple with the Use field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUse

`func (o *AgentJsonWebKeyRsaRequest) SetUse(v string)`

SetUse sets Use field to given value.

### HasUse

`func (o *AgentJsonWebKeyRsaRequest) HasUse() bool`

HasUse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


