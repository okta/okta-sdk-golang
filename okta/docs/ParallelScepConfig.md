# ParallelScepConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChallengeType** | **string** | The challenge type for a configuration created during a renewal. Only &#x60;DELEGATED&#x60; is supported for now — a static or dynamic clone would need a freshly generated challenge secret, which this operation can&#39;t return. Migrate those configurations with &#x60;rolloverConfigIds&#x60; instead. | 
**ConfigInfo** | Pointer to [**RegistrationAuthorityConfigInfo**](RegistrationAuthorityConfigInfo.md) |  | [optional] 
**Name** | Pointer to **string** | The name for the new configuration. Give it one that distinguishes it from the source, so the two are easy to tell apart while both are live. | [optional] 
**Protocol** | **string** | The enrollment protocol the configuration serves | 
**SourceConfigId** | **string** | The ID of the existing configuration to clone. It stays bound to the certificate it&#39;s already on. | 

## Methods

### NewParallelScepConfig

`func NewParallelScepConfig(challengeType string, protocol string, sourceConfigId string, ) *ParallelScepConfig`

NewParallelScepConfig instantiates a new ParallelScepConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParallelScepConfigWithDefaults

`func NewParallelScepConfigWithDefaults() *ParallelScepConfig`

NewParallelScepConfigWithDefaults instantiates a new ParallelScepConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChallengeType

`func (o *ParallelScepConfig) GetChallengeType() string`

GetChallengeType returns the ChallengeType field if non-nil, zero value otherwise.

### GetChallengeTypeOk

`func (o *ParallelScepConfig) GetChallengeTypeOk() (*string, bool)`

GetChallengeTypeOk returns a tuple with the ChallengeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallengeType

`func (o *ParallelScepConfig) SetChallengeType(v string)`

SetChallengeType sets ChallengeType field to given value.


### GetConfigInfo

`func (o *ParallelScepConfig) GetConfigInfo() RegistrationAuthorityConfigInfo`

GetConfigInfo returns the ConfigInfo field if non-nil, zero value otherwise.

### GetConfigInfoOk

`func (o *ParallelScepConfig) GetConfigInfoOk() (*RegistrationAuthorityConfigInfo, bool)`

GetConfigInfoOk returns a tuple with the ConfigInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigInfo

`func (o *ParallelScepConfig) SetConfigInfo(v RegistrationAuthorityConfigInfo)`

SetConfigInfo sets ConfigInfo field to given value.

### HasConfigInfo

`func (o *ParallelScepConfig) HasConfigInfo() bool`

HasConfigInfo returns a boolean if a field has been set.

### GetName

`func (o *ParallelScepConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ParallelScepConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ParallelScepConfig) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ParallelScepConfig) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProtocol

`func (o *ParallelScepConfig) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *ParallelScepConfig) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *ParallelScepConfig) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.


### GetSourceConfigId

`func (o *ParallelScepConfig) GetSourceConfigId() string`

GetSourceConfigId returns the SourceConfigId field if non-nil, zero value otherwise.

### GetSourceConfigIdOk

`func (o *ParallelScepConfig) GetSourceConfigIdOk() (*string, bool)`

GetSourceConfigIdOk returns a tuple with the SourceConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceConfigId

`func (o *ParallelScepConfig) SetSourceConfigId(v string)`

SetSourceConfigId sets SourceConfigId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


