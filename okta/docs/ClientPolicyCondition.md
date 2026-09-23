# ClientPolicyCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cimd** | Pointer to [**[]CimdClientCondition**](CimdClientCondition.md) | &lt;x-lifecycle-container&gt;&lt;x-lifecycle class&#x3D;\&quot;oie\&quot;&gt;&lt;/x-lifecycle&gt;&lt;/x-lifecycle-container&gt;The CIMD-registered apps or AI agents to include in the policy, identified by type and ID | [optional] 
**Include** | Pointer to **[]string** | Which clients are included in the policy | [optional] 

## Methods

### NewClientPolicyCondition

`func NewClientPolicyCondition() *ClientPolicyCondition`

NewClientPolicyCondition instantiates a new ClientPolicyCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientPolicyConditionWithDefaults

`func NewClientPolicyConditionWithDefaults() *ClientPolicyCondition`

NewClientPolicyConditionWithDefaults instantiates a new ClientPolicyCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCimd

`func (o *ClientPolicyCondition) GetCimd() []CimdClientCondition`

GetCimd returns the Cimd field if non-nil, zero value otherwise.

### GetCimdOk

`func (o *ClientPolicyCondition) GetCimdOk() (*[]CimdClientCondition, bool)`

GetCimdOk returns a tuple with the Cimd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCimd

`func (o *ClientPolicyCondition) SetCimd(v []CimdClientCondition)`

SetCimd sets Cimd field to given value.

### HasCimd

`func (o *ClientPolicyCondition) HasCimd() bool`

HasCimd returns a boolean if a field has been set.

### GetInclude

`func (o *ClientPolicyCondition) GetInclude() []string`

GetInclude returns the Include field if non-nil, zero value otherwise.

### GetIncludeOk

`func (o *ClientPolicyCondition) GetIncludeOk() (*[]string, bool)`

GetIncludeOk returns a tuple with the Include field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInclude

`func (o *ClientPolicyCondition) SetInclude(v []string)`

SetInclude sets Include field to given value.

### HasInclude

`func (o *ClientPolicyCondition) HasInclude() bool`

HasInclude returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


