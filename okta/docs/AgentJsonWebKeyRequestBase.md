# AgentJsonWebKeyRequestBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kid** | Pointer to **string** | Unique identifier of the JSON Web Key in the AI agent&#39;s JSON Web Key Set (JWKS) | [optional] 
**Status** | Pointer to **string** | Status of the AI agent JSON Web Key | [optional] [default to "ACTIVE"]

## Methods

### NewAgentJsonWebKeyRequestBase

`func NewAgentJsonWebKeyRequestBase() *AgentJsonWebKeyRequestBase`

NewAgentJsonWebKeyRequestBase instantiates a new AgentJsonWebKeyRequestBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentJsonWebKeyRequestBaseWithDefaults

`func NewAgentJsonWebKeyRequestBaseWithDefaults() *AgentJsonWebKeyRequestBase`

NewAgentJsonWebKeyRequestBaseWithDefaults instantiates a new AgentJsonWebKeyRequestBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKid

`func (o *AgentJsonWebKeyRequestBase) GetKid() string`

GetKid returns the Kid field if non-nil, zero value otherwise.

### GetKidOk

`func (o *AgentJsonWebKeyRequestBase) GetKidOk() (*string, bool)`

GetKidOk returns a tuple with the Kid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKid

`func (o *AgentJsonWebKeyRequestBase) SetKid(v string)`

SetKid sets Kid field to given value.

### HasKid

`func (o *AgentJsonWebKeyRequestBase) HasKid() bool`

HasKid returns a boolean if a field has been set.

### GetStatus

`func (o *AgentJsonWebKeyRequestBase) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AgentJsonWebKeyRequestBase) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AgentJsonWebKeyRequestBase) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AgentJsonWebKeyRequestBase) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


