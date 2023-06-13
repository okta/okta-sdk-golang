# AgentPoolUpdateSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentType** | Pointer to [**AgentType**](AgentType.md) |  | [optional] 
**ContinueOnError** | Pointer to **bool** |  | [optional] 
**LatestVersion** | Pointer to **string** |  | [optional] 
**MinimalSupportedVersion** | Pointer to **string** |  | [optional] 
**PoolId** | Pointer to **string** |  | [optional] [readonly] 
**PoolName** | Pointer to **string** |  | [optional] 
**ReleaseChannel** | Pointer to [**ReleaseChannel**](ReleaseChannel.md) |  | [optional] 

## Methods

### NewAgentPoolUpdateSetting

`func NewAgentPoolUpdateSetting() *AgentPoolUpdateSetting`

NewAgentPoolUpdateSetting instantiates a new AgentPoolUpdateSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentPoolUpdateSettingWithDefaults

`func NewAgentPoolUpdateSettingWithDefaults() *AgentPoolUpdateSetting`

NewAgentPoolUpdateSettingWithDefaults instantiates a new AgentPoolUpdateSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentType

`func (o *AgentPoolUpdateSetting) GetAgentType() AgentType`

GetAgentType returns the AgentType field if non-nil, zero value otherwise.

### GetAgentTypeOk

`func (o *AgentPoolUpdateSetting) GetAgentTypeOk() (*AgentType, bool)`

GetAgentTypeOk returns a tuple with the AgentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentType

`func (o *AgentPoolUpdateSetting) SetAgentType(v AgentType)`

SetAgentType sets AgentType field to given value.

### HasAgentType

`func (o *AgentPoolUpdateSetting) HasAgentType() bool`

HasAgentType returns a boolean if a field has been set.

### GetContinueOnError

`func (o *AgentPoolUpdateSetting) GetContinueOnError() bool`

GetContinueOnError returns the ContinueOnError field if non-nil, zero value otherwise.

### GetContinueOnErrorOk

`func (o *AgentPoolUpdateSetting) GetContinueOnErrorOk() (*bool, bool)`

GetContinueOnErrorOk returns a tuple with the ContinueOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueOnError

`func (o *AgentPoolUpdateSetting) SetContinueOnError(v bool)`

SetContinueOnError sets ContinueOnError field to given value.

### HasContinueOnError

`func (o *AgentPoolUpdateSetting) HasContinueOnError() bool`

HasContinueOnError returns a boolean if a field has been set.

### GetLatestVersion

`func (o *AgentPoolUpdateSetting) GetLatestVersion() string`

GetLatestVersion returns the LatestVersion field if non-nil, zero value otherwise.

### GetLatestVersionOk

`func (o *AgentPoolUpdateSetting) GetLatestVersionOk() (*string, bool)`

GetLatestVersionOk returns a tuple with the LatestVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestVersion

`func (o *AgentPoolUpdateSetting) SetLatestVersion(v string)`

SetLatestVersion sets LatestVersion field to given value.

### HasLatestVersion

`func (o *AgentPoolUpdateSetting) HasLatestVersion() bool`

HasLatestVersion returns a boolean if a field has been set.

### GetMinimalSupportedVersion

`func (o *AgentPoolUpdateSetting) GetMinimalSupportedVersion() string`

GetMinimalSupportedVersion returns the MinimalSupportedVersion field if non-nil, zero value otherwise.

### GetMinimalSupportedVersionOk

`func (o *AgentPoolUpdateSetting) GetMinimalSupportedVersionOk() (*string, bool)`

GetMinimalSupportedVersionOk returns a tuple with the MinimalSupportedVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimalSupportedVersion

`func (o *AgentPoolUpdateSetting) SetMinimalSupportedVersion(v string)`

SetMinimalSupportedVersion sets MinimalSupportedVersion field to given value.

### HasMinimalSupportedVersion

`func (o *AgentPoolUpdateSetting) HasMinimalSupportedVersion() bool`

HasMinimalSupportedVersion returns a boolean if a field has been set.

### GetPoolId

`func (o *AgentPoolUpdateSetting) GetPoolId() string`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *AgentPoolUpdateSetting) GetPoolIdOk() (*string, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *AgentPoolUpdateSetting) SetPoolId(v string)`

SetPoolId sets PoolId field to given value.

### HasPoolId

`func (o *AgentPoolUpdateSetting) HasPoolId() bool`

HasPoolId returns a boolean if a field has been set.

### GetPoolName

`func (o *AgentPoolUpdateSetting) GetPoolName() string`

GetPoolName returns the PoolName field if non-nil, zero value otherwise.

### GetPoolNameOk

`func (o *AgentPoolUpdateSetting) GetPoolNameOk() (*string, bool)`

GetPoolNameOk returns a tuple with the PoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolName

`func (o *AgentPoolUpdateSetting) SetPoolName(v string)`

SetPoolName sets PoolName field to given value.

### HasPoolName

`func (o *AgentPoolUpdateSetting) HasPoolName() bool`

HasPoolName returns a boolean if a field has been set.

### GetReleaseChannel

`func (o *AgentPoolUpdateSetting) GetReleaseChannel() ReleaseChannel`

GetReleaseChannel returns the ReleaseChannel field if non-nil, zero value otherwise.

### GetReleaseChannelOk

`func (o *AgentPoolUpdateSetting) GetReleaseChannelOk() (*ReleaseChannel, bool)`

GetReleaseChannelOk returns a tuple with the ReleaseChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseChannel

`func (o *AgentPoolUpdateSetting) SetReleaseChannel(v ReleaseChannel)`

SetReleaseChannel sets ReleaseChannel field to given value.

### HasReleaseChannel

`func (o *AgentPoolUpdateSetting) HasReleaseChannel() bool`

HasReleaseChannel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


