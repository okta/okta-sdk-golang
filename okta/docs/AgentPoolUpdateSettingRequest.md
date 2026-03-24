# AgentPoolUpdateSettingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentType** | **string** | Agent types that are being monitored | 
**ContinueOnError** | Pointer to **bool** | Continues the update even if some agents fail to update | [optional] 
**LatestVersion** | Pointer to **string** | Latest version of the agent | [optional] 
**MinimalSupportedVersion** | Pointer to **string** | Minimal version of the agent | [optional] 
**PoolId** | Pointer to **string** | ID of the agent pool that the settings apply to | [optional] [readonly] 
**PoolName** | Pointer to **string** | Pool name | [optional] 
**ReleaseChannel** | Pointer to **string** | Release channel for auto-update | [optional] 

## Methods

### NewAgentPoolUpdateSettingRequest

`func NewAgentPoolUpdateSettingRequest(agentType string, ) *AgentPoolUpdateSettingRequest`

NewAgentPoolUpdateSettingRequest instantiates a new AgentPoolUpdateSettingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentPoolUpdateSettingRequestWithDefaults

`func NewAgentPoolUpdateSettingRequestWithDefaults() *AgentPoolUpdateSettingRequest`

NewAgentPoolUpdateSettingRequestWithDefaults instantiates a new AgentPoolUpdateSettingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentType

`func (o *AgentPoolUpdateSettingRequest) GetAgentType() string`

GetAgentType returns the AgentType field if non-nil, zero value otherwise.

### GetAgentTypeOk

`func (o *AgentPoolUpdateSettingRequest) GetAgentTypeOk() (*string, bool)`

GetAgentTypeOk returns a tuple with the AgentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentType

`func (o *AgentPoolUpdateSettingRequest) SetAgentType(v string)`

SetAgentType sets AgentType field to given value.


### GetContinueOnError

`func (o *AgentPoolUpdateSettingRequest) GetContinueOnError() bool`

GetContinueOnError returns the ContinueOnError field if non-nil, zero value otherwise.

### GetContinueOnErrorOk

`func (o *AgentPoolUpdateSettingRequest) GetContinueOnErrorOk() (*bool, bool)`

GetContinueOnErrorOk returns a tuple with the ContinueOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueOnError

`func (o *AgentPoolUpdateSettingRequest) SetContinueOnError(v bool)`

SetContinueOnError sets ContinueOnError field to given value.

### HasContinueOnError

`func (o *AgentPoolUpdateSettingRequest) HasContinueOnError() bool`

HasContinueOnError returns a boolean if a field has been set.

### GetLatestVersion

`func (o *AgentPoolUpdateSettingRequest) GetLatestVersion() string`

GetLatestVersion returns the LatestVersion field if non-nil, zero value otherwise.

### GetLatestVersionOk

`func (o *AgentPoolUpdateSettingRequest) GetLatestVersionOk() (*string, bool)`

GetLatestVersionOk returns a tuple with the LatestVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestVersion

`func (o *AgentPoolUpdateSettingRequest) SetLatestVersion(v string)`

SetLatestVersion sets LatestVersion field to given value.

### HasLatestVersion

`func (o *AgentPoolUpdateSettingRequest) HasLatestVersion() bool`

HasLatestVersion returns a boolean if a field has been set.

### GetMinimalSupportedVersion

`func (o *AgentPoolUpdateSettingRequest) GetMinimalSupportedVersion() string`

GetMinimalSupportedVersion returns the MinimalSupportedVersion field if non-nil, zero value otherwise.

### GetMinimalSupportedVersionOk

`func (o *AgentPoolUpdateSettingRequest) GetMinimalSupportedVersionOk() (*string, bool)`

GetMinimalSupportedVersionOk returns a tuple with the MinimalSupportedVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimalSupportedVersion

`func (o *AgentPoolUpdateSettingRequest) SetMinimalSupportedVersion(v string)`

SetMinimalSupportedVersion sets MinimalSupportedVersion field to given value.

### HasMinimalSupportedVersion

`func (o *AgentPoolUpdateSettingRequest) HasMinimalSupportedVersion() bool`

HasMinimalSupportedVersion returns a boolean if a field has been set.

### GetPoolId

`func (o *AgentPoolUpdateSettingRequest) GetPoolId() string`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *AgentPoolUpdateSettingRequest) GetPoolIdOk() (*string, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *AgentPoolUpdateSettingRequest) SetPoolId(v string)`

SetPoolId sets PoolId field to given value.

### HasPoolId

`func (o *AgentPoolUpdateSettingRequest) HasPoolId() bool`

HasPoolId returns a boolean if a field has been set.

### GetPoolName

`func (o *AgentPoolUpdateSettingRequest) GetPoolName() string`

GetPoolName returns the PoolName field if non-nil, zero value otherwise.

### GetPoolNameOk

`func (o *AgentPoolUpdateSettingRequest) GetPoolNameOk() (*string, bool)`

GetPoolNameOk returns a tuple with the PoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolName

`func (o *AgentPoolUpdateSettingRequest) SetPoolName(v string)`

SetPoolName sets PoolName field to given value.

### HasPoolName

`func (o *AgentPoolUpdateSettingRequest) HasPoolName() bool`

HasPoolName returns a boolean if a field has been set.

### GetReleaseChannel

`func (o *AgentPoolUpdateSettingRequest) GetReleaseChannel() string`

GetReleaseChannel returns the ReleaseChannel field if non-nil, zero value otherwise.

### GetReleaseChannelOk

`func (o *AgentPoolUpdateSettingRequest) GetReleaseChannelOk() (*string, bool)`

GetReleaseChannelOk returns a tuple with the ReleaseChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseChannel

`func (o *AgentPoolUpdateSettingRequest) SetReleaseChannel(v string)`

SetReleaseChannel sets ReleaseChannel field to given value.

### HasReleaseChannel

`func (o *AgentPoolUpdateSettingRequest) HasReleaseChannel() bool`

HasReleaseChannel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


