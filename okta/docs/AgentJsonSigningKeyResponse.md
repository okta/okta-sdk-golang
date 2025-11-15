# AgentJsonSigningKeyResponse

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
**Crv** | Pointer to **string** | The cryptographic curve that&#39;s used for the key pair | [optional] 
**X** | Pointer to **string** | The public x coordinate for the elliptic curve point | [optional] 
**Y** | Pointer to **string** | The public y coordinate for the elliptic curve point | [optional] 

## Methods

### NewAgentJsonSigningKeyResponse

`func NewAgentJsonSigningKeyResponse(kty string, ) *AgentJsonSigningKeyResponse`

NewAgentJsonSigningKeyResponse instantiates a new AgentJsonSigningKeyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentJsonSigningKeyResponseWithDefaults

`func NewAgentJsonSigningKeyResponseWithDefaults() *AgentJsonSigningKeyResponse`

NewAgentJsonSigningKeyResponseWithDefaults instantiates a new AgentJsonSigningKeyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetE

`func (o *AgentJsonSigningKeyResponse) GetE() string`

GetE returns the E field if non-nil, zero value otherwise.

### GetEOk

`func (o *AgentJsonSigningKeyResponse) GetEOk() (*string, bool)`

GetEOk returns a tuple with the E field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE

`func (o *AgentJsonSigningKeyResponse) SetE(v string)`

SetE sets E field to given value.

### HasE

`func (o *AgentJsonSigningKeyResponse) HasE() bool`

HasE returns a boolean if a field has been set.

### GetKty

`func (o *AgentJsonSigningKeyResponse) GetKty() string`

GetKty returns the Kty field if non-nil, zero value otherwise.

### GetKtyOk

`func (o *AgentJsonSigningKeyResponse) GetKtyOk() (*string, bool)`

GetKtyOk returns a tuple with the Kty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKty

`func (o *AgentJsonSigningKeyResponse) SetKty(v string)`

SetKty sets Kty field to given value.


### GetN

`func (o *AgentJsonSigningKeyResponse) GetN() string`

GetN returns the N field if non-nil, zero value otherwise.

### GetNOk

`func (o *AgentJsonSigningKeyResponse) GetNOk() (*string, bool)`

GetNOk returns a tuple with the N field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN

`func (o *AgentJsonSigningKeyResponse) SetN(v string)`

SetN sets N field to given value.

### HasN

`func (o *AgentJsonSigningKeyResponse) HasN() bool`

HasN returns a boolean if a field has been set.

### GetKid

`func (o *AgentJsonSigningKeyResponse) GetKid() string`

GetKid returns the Kid field if non-nil, zero value otherwise.

### GetKidOk

`func (o *AgentJsonSigningKeyResponse) GetKidOk() (*string, bool)`

GetKidOk returns a tuple with the Kid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKid

`func (o *AgentJsonSigningKeyResponse) SetKid(v string)`

SetKid sets Kid field to given value.

### HasKid

`func (o *AgentJsonSigningKeyResponse) HasKid() bool`

HasKid returns a boolean if a field has been set.

### GetStatus

`func (o *AgentJsonSigningKeyResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AgentJsonSigningKeyResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AgentJsonSigningKeyResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AgentJsonSigningKeyResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAlg

`func (o *AgentJsonSigningKeyResponse) GetAlg() string`

GetAlg returns the Alg field if non-nil, zero value otherwise.

### GetAlgOk

`func (o *AgentJsonSigningKeyResponse) GetAlgOk() (*string, bool)`

GetAlgOk returns a tuple with the Alg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlg

`func (o *AgentJsonSigningKeyResponse) SetAlg(v string)`

SetAlg sets Alg field to given value.

### HasAlg

`func (o *AgentJsonSigningKeyResponse) HasAlg() bool`

HasAlg returns a boolean if a field has been set.

### GetUse

`func (o *AgentJsonSigningKeyResponse) GetUse() string`

GetUse returns the Use field if non-nil, zero value otherwise.

### GetUseOk

`func (o *AgentJsonSigningKeyResponse) GetUseOk() (*string, bool)`

GetUseOk returns a tuple with the Use field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUse

`func (o *AgentJsonSigningKeyResponse) SetUse(v string)`

SetUse sets Use field to given value.

### HasUse

`func (o *AgentJsonSigningKeyResponse) HasUse() bool`

HasUse returns a boolean if a field has been set.

### GetCreated

`func (o *AgentJsonSigningKeyResponse) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AgentJsonSigningKeyResponse) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AgentJsonSigningKeyResponse) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *AgentJsonSigningKeyResponse) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetId

`func (o *AgentJsonSigningKeyResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgentJsonSigningKeyResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgentJsonSigningKeyResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AgentJsonSigningKeyResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdated

`func (o *AgentJsonSigningKeyResponse) GetLastUpdated() string`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *AgentJsonSigningKeyResponse) GetLastUpdatedOk() (*string, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *AgentJsonSigningKeyResponse) SetLastUpdated(v string)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *AgentJsonSigningKeyResponse) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetLinks

`func (o *AgentJsonSigningKeyResponse) GetLinks() AgentSecretLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AgentJsonSigningKeyResponse) GetLinksOk() (*AgentSecretLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AgentJsonSigningKeyResponse) SetLinks(v AgentSecretLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AgentJsonSigningKeyResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetCrv

`func (o *AgentJsonSigningKeyResponse) GetCrv() string`

GetCrv returns the Crv field if non-nil, zero value otherwise.

### GetCrvOk

`func (o *AgentJsonSigningKeyResponse) GetCrvOk() (*string, bool)`

GetCrvOk returns a tuple with the Crv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrv

`func (o *AgentJsonSigningKeyResponse) SetCrv(v string)`

SetCrv sets Crv field to given value.

### HasCrv

`func (o *AgentJsonSigningKeyResponse) HasCrv() bool`

HasCrv returns a boolean if a field has been set.

### GetX

`func (o *AgentJsonSigningKeyResponse) GetX() string`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *AgentJsonSigningKeyResponse) GetXOk() (*string, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *AgentJsonSigningKeyResponse) SetX(v string)`

SetX sets X field to given value.

### HasX

`func (o *AgentJsonSigningKeyResponse) HasX() bool`

HasX returns a boolean if a field has been set.

### GetY

`func (o *AgentJsonSigningKeyResponse) GetY() string`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *AgentJsonSigningKeyResponse) GetYOk() (*string, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *AgentJsonSigningKeyResponse) SetY(v string)`

SetY sets Y field to given value.

### HasY

`func (o *AgentJsonSigningKeyResponse) HasY() bool`

HasY returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


