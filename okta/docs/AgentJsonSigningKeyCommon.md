# AgentJsonSigningKeyCommon

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alg** | Pointer to **string** | Algorithm that&#39;s used in the JSON Web Key | [optional] 
**Use** | Pointer to **string** | Acceptable use of the JSON Web Key  You can only use signing keys for AI agents, so the value of &#x60;use&#x60; is always &#x60;sig&#x60;. | [optional] 

## Methods

### NewAgentJsonSigningKeyCommon

`func NewAgentJsonSigningKeyCommon() *AgentJsonSigningKeyCommon`

NewAgentJsonSigningKeyCommon instantiates a new AgentJsonSigningKeyCommon object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentJsonSigningKeyCommonWithDefaults

`func NewAgentJsonSigningKeyCommonWithDefaults() *AgentJsonSigningKeyCommon`

NewAgentJsonSigningKeyCommonWithDefaults instantiates a new AgentJsonSigningKeyCommon object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlg

`func (o *AgentJsonSigningKeyCommon) GetAlg() string`

GetAlg returns the Alg field if non-nil, zero value otherwise.

### GetAlgOk

`func (o *AgentJsonSigningKeyCommon) GetAlgOk() (*string, bool)`

GetAlgOk returns a tuple with the Alg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlg

`func (o *AgentJsonSigningKeyCommon) SetAlg(v string)`

SetAlg sets Alg field to given value.

### HasAlg

`func (o *AgentJsonSigningKeyCommon) HasAlg() bool`

HasAlg returns a boolean if a field has been set.

### GetUse

`func (o *AgentJsonSigningKeyCommon) GetUse() string`

GetUse returns the Use field if non-nil, zero value otherwise.

### GetUseOk

`func (o *AgentJsonSigningKeyCommon) GetUseOk() (*string, bool)`

GetUseOk returns a tuple with the Use field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUse

`func (o *AgentJsonSigningKeyCommon) SetUse(v string)`

SetUse sets Use field to given value.

### HasUse

`func (o *AgentJsonSigningKeyCommon) HasUse() bool`

HasUse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


