# AgentJsonWebKeyRsaResponse

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
**Created** | Pointer to **string** | Timestamp of when the AI agent JSON Web Key was created | [optional] [readonly] 
**Id** | Pointer to **string** | The unique ID of the AI agent JSON Web Key | [optional] [readonly] 
**LastUpdated** | Pointer to **string** | Timestamp of when the AI agent JSON Web Key was last updated | [optional] [readonly] 
**Links** | Pointer to [**AgentSecretLinks**](AgentSecretLinks.md) |  | [optional] 

## Methods

### NewAgentJsonWebKeyRsaResponse

`func NewAgentJsonWebKeyRsaResponse(kty string, ) *AgentJsonWebKeyRsaResponse`

NewAgentJsonWebKeyRsaResponse instantiates a new AgentJsonWebKeyRsaResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentJsonWebKeyRsaResponseWithDefaults

`func NewAgentJsonWebKeyRsaResponseWithDefaults() *AgentJsonWebKeyRsaResponse`

NewAgentJsonWebKeyRsaResponseWithDefaults instantiates a new AgentJsonWebKeyRsaResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetE

`func (o *AgentJsonWebKeyRsaResponse) GetE() string`

GetE returns the E field if non-nil, zero value otherwise.

### GetEOk

`func (o *AgentJsonWebKeyRsaResponse) GetEOk() (*string, bool)`

GetEOk returns a tuple with the E field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE

`func (o *AgentJsonWebKeyRsaResponse) SetE(v string)`

SetE sets E field to given value.

### HasE

`func (o *AgentJsonWebKeyRsaResponse) HasE() bool`

HasE returns a boolean if a field has been set.

### GetKty

`func (o *AgentJsonWebKeyRsaResponse) GetKty() string`

GetKty returns the Kty field if non-nil, zero value otherwise.

### GetKtyOk

`func (o *AgentJsonWebKeyRsaResponse) GetKtyOk() (*string, bool)`

GetKtyOk returns a tuple with the Kty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKty

`func (o *AgentJsonWebKeyRsaResponse) SetKty(v string)`

SetKty sets Kty field to given value.


### GetN

`func (o *AgentJsonWebKeyRsaResponse) GetN() string`

GetN returns the N field if non-nil, zero value otherwise.

### GetNOk

`func (o *AgentJsonWebKeyRsaResponse) GetNOk() (*string, bool)`

GetNOk returns a tuple with the N field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN

`func (o *AgentJsonWebKeyRsaResponse) SetN(v string)`

SetN sets N field to given value.

### HasN

`func (o *AgentJsonWebKeyRsaResponse) HasN() bool`

HasN returns a boolean if a field has been set.

### GetKid

`func (o *AgentJsonWebKeyRsaResponse) GetKid() string`

GetKid returns the Kid field if non-nil, zero value otherwise.

### GetKidOk

`func (o *AgentJsonWebKeyRsaResponse) GetKidOk() (*string, bool)`

GetKidOk returns a tuple with the Kid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKid

`func (o *AgentJsonWebKeyRsaResponse) SetKid(v string)`

SetKid sets Kid field to given value.

### HasKid

`func (o *AgentJsonWebKeyRsaResponse) HasKid() bool`

HasKid returns a boolean if a field has been set.

### GetStatus

`func (o *AgentJsonWebKeyRsaResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AgentJsonWebKeyRsaResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AgentJsonWebKeyRsaResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AgentJsonWebKeyRsaResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAlg

`func (o *AgentJsonWebKeyRsaResponse) GetAlg() string`

GetAlg returns the Alg field if non-nil, zero value otherwise.

### GetAlgOk

`func (o *AgentJsonWebKeyRsaResponse) GetAlgOk() (*string, bool)`

GetAlgOk returns a tuple with the Alg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlg

`func (o *AgentJsonWebKeyRsaResponse) SetAlg(v string)`

SetAlg sets Alg field to given value.

### HasAlg

`func (o *AgentJsonWebKeyRsaResponse) HasAlg() bool`

HasAlg returns a boolean if a field has been set.

### GetUse

`func (o *AgentJsonWebKeyRsaResponse) GetUse() string`

GetUse returns the Use field if non-nil, zero value otherwise.

### GetUseOk

`func (o *AgentJsonWebKeyRsaResponse) GetUseOk() (*string, bool)`

GetUseOk returns a tuple with the Use field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUse

`func (o *AgentJsonWebKeyRsaResponse) SetUse(v string)`

SetUse sets Use field to given value.

### HasUse

`func (o *AgentJsonWebKeyRsaResponse) HasUse() bool`

HasUse returns a boolean if a field has been set.

### GetCreated

`func (o *AgentJsonWebKeyRsaResponse) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AgentJsonWebKeyRsaResponse) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AgentJsonWebKeyRsaResponse) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *AgentJsonWebKeyRsaResponse) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetId

`func (o *AgentJsonWebKeyRsaResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgentJsonWebKeyRsaResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgentJsonWebKeyRsaResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AgentJsonWebKeyRsaResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdated

`func (o *AgentJsonWebKeyRsaResponse) GetLastUpdated() string`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *AgentJsonWebKeyRsaResponse) GetLastUpdatedOk() (*string, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *AgentJsonWebKeyRsaResponse) SetLastUpdated(v string)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *AgentJsonWebKeyRsaResponse) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetLinks

`func (o *AgentJsonWebKeyRsaResponse) GetLinks() AgentSecretLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AgentJsonWebKeyRsaResponse) GetLinksOk() (*AgentSecretLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AgentJsonWebKeyRsaResponse) SetLinks(v AgentSecretLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AgentJsonWebKeyRsaResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


