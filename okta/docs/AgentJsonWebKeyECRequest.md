# AgentJsonWebKeyECRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Crv** | Pointer to **string** | The cryptographic curve that&#39;s used for the key pair | [optional] 
**Kty** | **string** | Cryptographic algorithm family for the certificate&#39;s key pair | 
**X** | Pointer to **string** | The public x coordinate for the elliptic curve point | [optional] 
**Y** | Pointer to **string** | The public y coordinate for the elliptic curve point | [optional] 
**Kid** | Pointer to **string** | Unique identifier of the JSON Web Key in the AI agent&#39;s JSON Web Key Set (JWKS) | [optional] 
**Status** | Pointer to **string** | Status of the AI agent JSON Web Key | [optional] [default to "ACTIVE"]
**Alg** | Pointer to **string** | Algorithm that&#39;s used in the JSON Web Key | [optional] 
**Use** | Pointer to **string** | Acceptable use of the JSON Web Key  You can only use signing keys for AI agents, so the value of &#x60;use&#x60; is always &#x60;sig&#x60;. | [optional] 

## Methods

### NewAgentJsonWebKeyECRequest

`func NewAgentJsonWebKeyECRequest(kty string, ) *AgentJsonWebKeyECRequest`

NewAgentJsonWebKeyECRequest instantiates a new AgentJsonWebKeyECRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentJsonWebKeyECRequestWithDefaults

`func NewAgentJsonWebKeyECRequestWithDefaults() *AgentJsonWebKeyECRequest`

NewAgentJsonWebKeyECRequestWithDefaults instantiates a new AgentJsonWebKeyECRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCrv

`func (o *AgentJsonWebKeyECRequest) GetCrv() string`

GetCrv returns the Crv field if non-nil, zero value otherwise.

### GetCrvOk

`func (o *AgentJsonWebKeyECRequest) GetCrvOk() (*string, bool)`

GetCrvOk returns a tuple with the Crv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrv

`func (o *AgentJsonWebKeyECRequest) SetCrv(v string)`

SetCrv sets Crv field to given value.

### HasCrv

`func (o *AgentJsonWebKeyECRequest) HasCrv() bool`

HasCrv returns a boolean if a field has been set.

### GetKty

`func (o *AgentJsonWebKeyECRequest) GetKty() string`

GetKty returns the Kty field if non-nil, zero value otherwise.

### GetKtyOk

`func (o *AgentJsonWebKeyECRequest) GetKtyOk() (*string, bool)`

GetKtyOk returns a tuple with the Kty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKty

`func (o *AgentJsonWebKeyECRequest) SetKty(v string)`

SetKty sets Kty field to given value.


### GetX

`func (o *AgentJsonWebKeyECRequest) GetX() string`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *AgentJsonWebKeyECRequest) GetXOk() (*string, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *AgentJsonWebKeyECRequest) SetX(v string)`

SetX sets X field to given value.

### HasX

`func (o *AgentJsonWebKeyECRequest) HasX() bool`

HasX returns a boolean if a field has been set.

### GetY

`func (o *AgentJsonWebKeyECRequest) GetY() string`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *AgentJsonWebKeyECRequest) GetYOk() (*string, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *AgentJsonWebKeyECRequest) SetY(v string)`

SetY sets Y field to given value.

### HasY

`func (o *AgentJsonWebKeyECRequest) HasY() bool`

HasY returns a boolean if a field has been set.

### GetKid

`func (o *AgentJsonWebKeyECRequest) GetKid() string`

GetKid returns the Kid field if non-nil, zero value otherwise.

### GetKidOk

`func (o *AgentJsonWebKeyECRequest) GetKidOk() (*string, bool)`

GetKidOk returns a tuple with the Kid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKid

`func (o *AgentJsonWebKeyECRequest) SetKid(v string)`

SetKid sets Kid field to given value.

### HasKid

`func (o *AgentJsonWebKeyECRequest) HasKid() bool`

HasKid returns a boolean if a field has been set.

### GetStatus

`func (o *AgentJsonWebKeyECRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AgentJsonWebKeyECRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AgentJsonWebKeyECRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AgentJsonWebKeyECRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAlg

`func (o *AgentJsonWebKeyECRequest) GetAlg() string`

GetAlg returns the Alg field if non-nil, zero value otherwise.

### GetAlgOk

`func (o *AgentJsonWebKeyECRequest) GetAlgOk() (*string, bool)`

GetAlgOk returns a tuple with the Alg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlg

`func (o *AgentJsonWebKeyECRequest) SetAlg(v string)`

SetAlg sets Alg field to given value.

### HasAlg

`func (o *AgentJsonWebKeyECRequest) HasAlg() bool`

HasAlg returns a boolean if a field has been set.

### GetUse

`func (o *AgentJsonWebKeyECRequest) GetUse() string`

GetUse returns the Use field if non-nil, zero value otherwise.

### GetUseOk

`func (o *AgentJsonWebKeyECRequest) GetUseOk() (*string, bool)`

GetUseOk returns a tuple with the Use field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUse

`func (o *AgentJsonWebKeyECRequest) SetUse(v string)`

SetUse sets Use field to given value.

### HasUse

`func (o *AgentJsonWebKeyECRequest) HasUse() bool`

HasUse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


