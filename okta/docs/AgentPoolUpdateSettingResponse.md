# AgentPoolUpdateSettingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentType** | Pointer to **string** | Agent types that are being monitored | [optional] 
**ContinueOnError** | Pointer to **bool** | Continues the update even if some agents fail to update | [optional] 
**LatestVersion** | Pointer to **string** | Latest version of the agent | [optional] 
**MinimalSupportedVersion** | Pointer to **string** | Minimal version of the agent | [optional] 
**PoolId** | Pointer to **string** | ID of the agent pool that the settings apply to | [optional] [readonly] 
**PoolName** | Pointer to **string** | Pool name | [optional] 
**ReleaseChannel** | Pointer to **string** | Release channel for auto-update | [optional] 

## Methods

### NewAgentPoolUpdateSettingResponse

`func NewAgentPoolUpdateSettingResponse() *AgentPoolUpdateSettingResponse`

NewAgentPoolUpdateSettingResponse instantiates a new AgentPoolUpdateSettingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentPoolUpdateSettingResponseWithDefaults

`func NewAgentPoolUpdateSettingResponseWithDefaults() *AgentPoolUpdateSettingResponse`

NewAgentPoolUpdateSettingResponseWithDefaults instantiates a new AgentPoolUpdateSettingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentType

`func (o *AgentPoolUpdateSettingResponse) GetAgentType() string`

GetAgentType returns the AgentType field if non-nil, zero value otherwise.

### GetAgentTypeOk

`func (o *AgentPoolUpdateSettingResponse) GetAgentTypeOk() (*string, bool)`

GetAgentTypeOk returns a tuple with the AgentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentType

`func (o *AgentPoolUpdateSettingResponse) SetAgentType(v string)`

SetAgentType sets AgentType field to given value.

### HasAgentType

`func (o *AgentPoolUpdateSettingResponse) HasAgentType() bool`

HasAgentType returns a boolean if a field has been set.

### GetContinueOnError

`func (o *AgentPoolUpdateSettingResponse) GetContinueOnError() bool`

GetContinueOnError returns the ContinueOnError field if non-nil, zero value otherwise.

### GetContinueOnErrorOk

`func (o *AgentPoolUpdateSettingResponse) GetContinueOnErrorOk() (*bool, bool)`

GetContinueOnErrorOk returns a tuple with the ContinueOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueOnError

`func (o *AgentPoolUpdateSettingResponse) SetContinueOnError(v bool)`

SetContinueOnError sets ContinueOnError field to given value.

### HasContinueOnError

`func (o *AgentPoolUpdateSettingResponse) HasContinueOnError() bool`

HasContinueOnError returns a boolean if a field has been set.

### GetLatestVersion

`func (o *AgentPoolUpdateSettingResponse) GetLatestVersion() string`

GetLatestVersion returns the LatestVersion field if non-nil, zero value otherwise.

### GetLatestVersionOk

`func (o *AgentPoolUpdateSettingResponse) GetLatestVersionOk() (*string, bool)`

GetLatestVersionOk returns a tuple with the LatestVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestVersion

`func (o *AgentPoolUpdateSettingResponse) SetLatestVersion(v string)`

SetLatestVersion sets LatestVersion field to given value.

### HasLatestVersion

`func (o *AgentPoolUpdateSettingResponse) HasLatestVersion() bool`

HasLatestVersion returns a boolean if a field has been set.

### GetMinimalSupportedVersion

`func (o *AgentPoolUpdateSettingResponse) GetMinimalSupportedVersion() string`

GetMinimalSupportedVersion returns the MinimalSupportedVersion field if non-nil, zero value otherwise.

### GetMinimalSupportedVersionOk

`func (o *AgentPoolUpdateSettingResponse) GetMinimalSupportedVersionOk() (*string, bool)`

GetMinimalSupportedVersionOk returns a tuple with the MinimalSupportedVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimalSupportedVersion

`func (o *AgentPoolUpdateSettingResponse) SetMinimalSupportedVersion(v string)`

SetMinimalSupportedVersion sets MinimalSupportedVersion field to given value.

### HasMinimalSupportedVersion

`func (o *AgentPoolUpdateSettingResponse) HasMinimalSupportedVersion() bool`

HasMinimalSupportedVersion returns a boolean if a field has been set.

### GetPoolId

`func (o *AgentPoolUpdateSettingResponse) GetPoolId() string`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *AgentPoolUpdateSettingResponse) GetPoolIdOk() (*string, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *AgentPoolUpdateSettingResponse) SetPoolId(v string)`

SetPoolId sets PoolId field to given value.

### HasPoolId

`func (o *AgentPoolUpdateSettingResponse) HasPoolId() bool`

HasPoolId returns a boolean if a field has been set.

### GetPoolName

`func (o *AgentPoolUpdateSettingResponse) GetPoolName() string`

GetPoolName returns the PoolName field if non-nil, zero value otherwise.

### GetPoolNameOk

`func (o *AgentPoolUpdateSettingResponse) GetPoolNameOk() (*string, bool)`

GetPoolNameOk returns a tuple with the PoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolName

`func (o *AgentPoolUpdateSettingResponse) SetPoolName(v string)`

SetPoolName sets PoolName field to given value.

### HasPoolName

`func (o *AgentPoolUpdateSettingResponse) HasPoolName() bool`

HasPoolName returns a boolean if a field has been set.

### GetReleaseChannel

`func (o *AgentPoolUpdateSettingResponse) GetReleaseChannel() string`

GetReleaseChannel returns the ReleaseChannel field if non-nil, zero value otherwise.

### GetReleaseChannelOk

`func (o *AgentPoolUpdateSettingResponse) GetReleaseChannelOk() (*string, bool)`

GetReleaseChannelOk returns a tuple with the ReleaseChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseChannel

`func (o *AgentPoolUpdateSettingResponse) SetReleaseChannel(v string)`

SetReleaseChannel sets ReleaseChannel field to given value.

### HasReleaseChannel

`func (o *AgentPoolUpdateSettingResponse) HasReleaseChannel() bool`

HasReleaseChannel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


