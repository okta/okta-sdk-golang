# IdpDiscoveryPolicyRuleCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**App** | Pointer to [**AppAndInstancePolicyRuleCondition**](AppAndInstancePolicyRuleCondition.md) |  | [optional] 
**Network** | Pointer to [**PolicyNetworkCondition**](PolicyNetworkCondition.md) |  | [optional] 
**UserIdentifier** | Pointer to [**UserIdentifierPolicyRuleCondition**](UserIdentifierPolicyRuleCondition.md) |  | [optional] 
**Platform** | Pointer to [**PlatformPolicyRuleCondition**](PlatformPolicyRuleCondition.md) |  | [optional] 

## Methods

### NewIdpDiscoveryPolicyRuleCondition

`func NewIdpDiscoveryPolicyRuleCondition() *IdpDiscoveryPolicyRuleCondition`

NewIdpDiscoveryPolicyRuleCondition instantiates a new IdpDiscoveryPolicyRuleCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdpDiscoveryPolicyRuleConditionWithDefaults

`func NewIdpDiscoveryPolicyRuleConditionWithDefaults() *IdpDiscoveryPolicyRuleCondition`

NewIdpDiscoveryPolicyRuleConditionWithDefaults instantiates a new IdpDiscoveryPolicyRuleCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApp

`func (o *IdpDiscoveryPolicyRuleCondition) GetApp() AppAndInstancePolicyRuleCondition`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *IdpDiscoveryPolicyRuleCondition) GetAppOk() (*AppAndInstancePolicyRuleCondition, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *IdpDiscoveryPolicyRuleCondition) SetApp(v AppAndInstancePolicyRuleCondition)`

SetApp sets App field to given value.

### HasApp

`func (o *IdpDiscoveryPolicyRuleCondition) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetNetwork

`func (o *IdpDiscoveryPolicyRuleCondition) GetNetwork() PolicyNetworkCondition`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *IdpDiscoveryPolicyRuleCondition) GetNetworkOk() (*PolicyNetworkCondition, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *IdpDiscoveryPolicyRuleCondition) SetNetwork(v PolicyNetworkCondition)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *IdpDiscoveryPolicyRuleCondition) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetUserIdentifier

`func (o *IdpDiscoveryPolicyRuleCondition) GetUserIdentifier() UserIdentifierPolicyRuleCondition`

GetUserIdentifier returns the UserIdentifier field if non-nil, zero value otherwise.

### GetUserIdentifierOk

`func (o *IdpDiscoveryPolicyRuleCondition) GetUserIdentifierOk() (*UserIdentifierPolicyRuleCondition, bool)`

GetUserIdentifierOk returns a tuple with the UserIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdentifier

`func (o *IdpDiscoveryPolicyRuleCondition) SetUserIdentifier(v UserIdentifierPolicyRuleCondition)`

SetUserIdentifier sets UserIdentifier field to given value.

### HasUserIdentifier

`func (o *IdpDiscoveryPolicyRuleCondition) HasUserIdentifier() bool`

HasUserIdentifier returns a boolean if a field has been set.

### GetPlatform

`func (o *IdpDiscoveryPolicyRuleCondition) GetPlatform() PlatformPolicyRuleCondition`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *IdpDiscoveryPolicyRuleCondition) GetPlatformOk() (*PlatformPolicyRuleCondition, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *IdpDiscoveryPolicyRuleCondition) SetPlatform(v PlatformPolicyRuleCondition)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *IdpDiscoveryPolicyRuleCondition) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


