# AgentJsonWebKeyECResponse

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
**Created** | Pointer to **string** | Timestamp of when the AI agent JSON Web Key was created | [optional] [readonly] 
**Id** | Pointer to **string** | The unique ID of the AI agent JSON Web Key | [optional] [readonly] 
**LastUpdated** | Pointer to **string** | Timestamp of when the AI agent JSON Web Key was last updated | [optional] [readonly] 
**Links** | Pointer to [**AgentSecretLinks**](AgentSecretLinks.md) |  | [optional] 

## Methods

### NewAgentJsonWebKeyECResponse

`func NewAgentJsonWebKeyECResponse(kty string, ) *AgentJsonWebKeyECResponse`

NewAgentJsonWebKeyECResponse instantiates a new AgentJsonWebKeyECResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentJsonWebKeyECResponseWithDefaults

`func NewAgentJsonWebKeyECResponseWithDefaults() *AgentJsonWebKeyECResponse`

NewAgentJsonWebKeyECResponseWithDefaults instantiates a new AgentJsonWebKeyECResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCrv

`func (o *AgentJsonWebKeyECResponse) GetCrv() string`

GetCrv returns the Crv field if non-nil, zero value otherwise.

### GetCrvOk

`func (o *AgentJsonWebKeyECResponse) GetCrvOk() (*string, bool)`

GetCrvOk returns a tuple with the Crv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrv

`func (o *AgentJsonWebKeyECResponse) SetCrv(v string)`

SetCrv sets Crv field to given value.

### HasCrv

`func (o *AgentJsonWebKeyECResponse) HasCrv() bool`

HasCrv returns a boolean if a field has been set.

### GetKty

`func (o *AgentJsonWebKeyECResponse) GetKty() string`

GetKty returns the Kty field if non-nil, zero value otherwise.

### GetKtyOk

`func (o *AgentJsonWebKeyECResponse) GetKtyOk() (*string, bool)`

GetKtyOk returns a tuple with the Kty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKty

`func (o *AgentJsonWebKeyECResponse) SetKty(v string)`

SetKty sets Kty field to given value.


### GetX

`func (o *AgentJsonWebKeyECResponse) GetX() string`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *AgentJsonWebKeyECResponse) GetXOk() (*string, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *AgentJsonWebKeyECResponse) SetX(v string)`

SetX sets X field to given value.

### HasX

`func (o *AgentJsonWebKeyECResponse) HasX() bool`

HasX returns a boolean if a field has been set.

### GetY

`func (o *AgentJsonWebKeyECResponse) GetY() string`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *AgentJsonWebKeyECResponse) GetYOk() (*string, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *AgentJsonWebKeyECResponse) SetY(v string)`

SetY sets Y field to given value.

### HasY

`func (o *AgentJsonWebKeyECResponse) HasY() bool`

HasY returns a boolean if a field has been set.

### GetKid

`func (o *AgentJsonWebKeyECResponse) GetKid() string`

GetKid returns the Kid field if non-nil, zero value otherwise.

### GetKidOk

`func (o *AgentJsonWebKeyECResponse) GetKidOk() (*string, bool)`

GetKidOk returns a tuple with the Kid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKid

`func (o *AgentJsonWebKeyECResponse) SetKid(v string)`

SetKid sets Kid field to given value.

### HasKid

`func (o *AgentJsonWebKeyECResponse) HasKid() bool`

HasKid returns a boolean if a field has been set.

### GetStatus

`func (o *AgentJsonWebKeyECResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AgentJsonWebKeyECResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AgentJsonWebKeyECResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AgentJsonWebKeyECResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAlg

`func (o *AgentJsonWebKeyECResponse) GetAlg() string`

GetAlg returns the Alg field if non-nil, zero value otherwise.

### GetAlgOk

`func (o *AgentJsonWebKeyECResponse) GetAlgOk() (*string, bool)`

GetAlgOk returns a tuple with the Alg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlg

`func (o *AgentJsonWebKeyECResponse) SetAlg(v string)`

SetAlg sets Alg field to given value.

### HasAlg

`func (o *AgentJsonWebKeyECResponse) HasAlg() bool`

HasAlg returns a boolean if a field has been set.

### GetUse

`func (o *AgentJsonWebKeyECResponse) GetUse() string`

GetUse returns the Use field if non-nil, zero value otherwise.

### GetUseOk

`func (o *AgentJsonWebKeyECResponse) GetUseOk() (*string, bool)`

GetUseOk returns a tuple with the Use field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUse

`func (o *AgentJsonWebKeyECResponse) SetUse(v string)`

SetUse sets Use field to given value.

### HasUse

`func (o *AgentJsonWebKeyECResponse) HasUse() bool`

HasUse returns a boolean if a field has been set.

### GetCreated

`func (o *AgentJsonWebKeyECResponse) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AgentJsonWebKeyECResponse) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AgentJsonWebKeyECResponse) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *AgentJsonWebKeyECResponse) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetId

`func (o *AgentJsonWebKeyECResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgentJsonWebKeyECResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgentJsonWebKeyECResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AgentJsonWebKeyECResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdated

`func (o *AgentJsonWebKeyECResponse) GetLastUpdated() string`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *AgentJsonWebKeyECResponse) GetLastUpdatedOk() (*string, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *AgentJsonWebKeyECResponse) SetLastUpdated(v string)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *AgentJsonWebKeyECResponse) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetLinks

`func (o *AgentJsonWebKeyECResponse) GetLinks() AgentSecretLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AgentJsonWebKeyECResponse) GetLinksOk() (*AgentSecretLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AgentJsonWebKeyECResponse) SetLinks(v AgentSecretLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AgentJsonWebKeyECResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


