# AccessPolicyRuleConditions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Device** | Pointer to [**DeviceAccessPolicyRuleCondition**](DeviceAccessPolicyRuleCondition.md) |  | [optional] 
**ElCondition** | Pointer to [**AccessPolicyRuleCustomCondition**](AccessPolicyRuleCustomCondition.md) |  | [optional] 
**Network** | Pointer to [**PolicyNetworkCondition**](PolicyNetworkCondition.md) |  | [optional] 
**People** | Pointer to [**PolicyPeopleCondition**](PolicyPeopleCondition.md) |  | [optional] 
**Platform** | Pointer to [**PlatformPolicyRuleCondition**](PlatformPolicyRuleCondition.md) |  | [optional] 
**RiskScore** | Pointer to [**RiskScorePolicyRuleCondition**](RiskScorePolicyRuleCondition.md) |  | [optional] 
**UserType** | Pointer to [**UserTypeCondition**](UserTypeCondition.md) |  | [optional] 

## Methods

### NewAccessPolicyRuleConditions

`func NewAccessPolicyRuleConditions() *AccessPolicyRuleConditions`

NewAccessPolicyRuleConditions instantiates a new AccessPolicyRuleConditions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessPolicyRuleConditionsWithDefaults

`func NewAccessPolicyRuleConditionsWithDefaults() *AccessPolicyRuleConditions`

NewAccessPolicyRuleConditionsWithDefaults instantiates a new AccessPolicyRuleConditions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevice

`func (o *AccessPolicyRuleConditions) GetDevice() DeviceAccessPolicyRuleCondition`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *AccessPolicyRuleConditions) GetDeviceOk() (*DeviceAccessPolicyRuleCondition, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *AccessPolicyRuleConditions) SetDevice(v DeviceAccessPolicyRuleCondition)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *AccessPolicyRuleConditions) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetElCondition

`func (o *AccessPolicyRuleConditions) GetElCondition() AccessPolicyRuleCustomCondition`

GetElCondition returns the ElCondition field if non-nil, zero value otherwise.

### GetElConditionOk

`func (o *AccessPolicyRuleConditions) GetElConditionOk() (*AccessPolicyRuleCustomCondition, bool)`

GetElConditionOk returns a tuple with the ElCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElCondition

`func (o *AccessPolicyRuleConditions) SetElCondition(v AccessPolicyRuleCustomCondition)`

SetElCondition sets ElCondition field to given value.

### HasElCondition

`func (o *AccessPolicyRuleConditions) HasElCondition() bool`

HasElCondition returns a boolean if a field has been set.

### GetNetwork

`func (o *AccessPolicyRuleConditions) GetNetwork() PolicyNetworkCondition`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *AccessPolicyRuleConditions) GetNetworkOk() (*PolicyNetworkCondition, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *AccessPolicyRuleConditions) SetNetwork(v PolicyNetworkCondition)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *AccessPolicyRuleConditions) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetPeople

`func (o *AccessPolicyRuleConditions) GetPeople() PolicyPeopleCondition`

GetPeople returns the People field if non-nil, zero value otherwise.

### GetPeopleOk

`func (o *AccessPolicyRuleConditions) GetPeopleOk() (*PolicyPeopleCondition, bool)`

GetPeopleOk returns a tuple with the People field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeople

`func (o *AccessPolicyRuleConditions) SetPeople(v PolicyPeopleCondition)`

SetPeople sets People field to given value.

### HasPeople

`func (o *AccessPolicyRuleConditions) HasPeople() bool`

HasPeople returns a boolean if a field has been set.

### GetPlatform

`func (o *AccessPolicyRuleConditions) GetPlatform() PlatformPolicyRuleCondition`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *AccessPolicyRuleConditions) GetPlatformOk() (*PlatformPolicyRuleCondition, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *AccessPolicyRuleConditions) SetPlatform(v PlatformPolicyRuleCondition)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *AccessPolicyRuleConditions) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetRiskScore

`func (o *AccessPolicyRuleConditions) GetRiskScore() RiskScorePolicyRuleCondition`

GetRiskScore returns the RiskScore field if non-nil, zero value otherwise.

### GetRiskScoreOk

`func (o *AccessPolicyRuleConditions) GetRiskScoreOk() (*RiskScorePolicyRuleCondition, bool)`

GetRiskScoreOk returns a tuple with the RiskScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskScore

`func (o *AccessPolicyRuleConditions) SetRiskScore(v RiskScorePolicyRuleCondition)`

SetRiskScore sets RiskScore field to given value.

### HasRiskScore

`func (o *AccessPolicyRuleConditions) HasRiskScore() bool`

HasRiskScore returns a boolean if a field has been set.

### GetUserType

`func (o *AccessPolicyRuleConditions) GetUserType() UserTypeCondition`

GetUserType returns the UserType field if non-nil, zero value otherwise.

### GetUserTypeOk

`func (o *AccessPolicyRuleConditions) GetUserTypeOk() (*UserTypeCondition, bool)`

GetUserTypeOk returns a tuple with the UserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserType

`func (o *AccessPolicyRuleConditions) SetUserType(v UserTypeCondition)`

SetUserType sets UserType field to given value.

### HasUserType

`func (o *AccessPolicyRuleConditions) HasUserType() bool`

HasUserType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


