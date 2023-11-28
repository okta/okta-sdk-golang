# AuthorizationServerPolicyRuleConditions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**App** | Pointer to [**AppAndInstancePolicyRuleCondition**](AppAndInstancePolicyRuleCondition.md) |  | [optional] 
**Apps** | Pointer to [**AppInstancePolicyRuleCondition**](AppInstancePolicyRuleCondition.md) |  | [optional] 
**AuthContext** | Pointer to [**PolicyRuleAuthContextCondition**](PolicyRuleAuthContextCondition.md) |  | [optional] 
**AuthProvider** | Pointer to [**PasswordPolicyAuthenticationProviderCondition**](PasswordPolicyAuthenticationProviderCondition.md) |  | [optional] 
**BeforeScheduledAction** | Pointer to [**BeforeScheduledActionPolicyRuleCondition**](BeforeScheduledActionPolicyRuleCondition.md) |  | [optional] 
**Clients** | Pointer to [**ClientPolicyCondition**](ClientPolicyCondition.md) |  | [optional] 
**Context** | Pointer to [**ContextPolicyRuleCondition**](ContextPolicyRuleCondition.md) |  | [optional] 
**Device** | Pointer to [**DevicePolicyRuleCondition**](DevicePolicyRuleCondition.md) |  | [optional] 
**GrantTypes** | Pointer to [**GrantTypePolicyRuleCondition**](GrantTypePolicyRuleCondition.md) |  | [optional] 
**Groups** | Pointer to [**GroupPolicyRuleCondition**](GroupPolicyRuleCondition.md) |  | [optional] 
**IdentityProvider** | Pointer to [**IdentityProviderPolicyRuleCondition**](IdentityProviderPolicyRuleCondition.md) |  | [optional] 
**MdmEnrollment** | Pointer to [**MDMEnrollmentPolicyRuleCondition**](MDMEnrollmentPolicyRuleCondition.md) |  | [optional] 
**Network** | Pointer to [**PolicyNetworkCondition**](PolicyNetworkCondition.md) |  | [optional] 
**People** | Pointer to [**PolicyPeopleCondition**](PolicyPeopleCondition.md) |  | [optional] 
**Platform** | Pointer to [**PlatformPolicyRuleCondition**](PlatformPolicyRuleCondition.md) |  | [optional] 
**Risk** | Pointer to [**RiskPolicyRuleCondition**](RiskPolicyRuleCondition.md) |  | [optional] 
**RiskScore** | Pointer to [**RiskScorePolicyRuleCondition**](RiskScorePolicyRuleCondition.md) |  | [optional] 
**Scopes** | Pointer to [**OAuth2ScopesMediationPolicyRuleCondition**](OAuth2ScopesMediationPolicyRuleCondition.md) |  | [optional] 
**UserIdentifier** | Pointer to [**UserIdentifierPolicyRuleCondition**](UserIdentifierPolicyRuleCondition.md) |  | [optional] 
**Users** | Pointer to [**UserPolicyRuleCondition**](UserPolicyRuleCondition.md) |  | [optional] 
**UserStatus** | Pointer to [**UserStatusPolicyRuleCondition**](UserStatusPolicyRuleCondition.md) |  | [optional] 

## Methods

### NewAuthorizationServerPolicyRuleConditions

`func NewAuthorizationServerPolicyRuleConditions() *AuthorizationServerPolicyRuleConditions`

NewAuthorizationServerPolicyRuleConditions instantiates a new AuthorizationServerPolicyRuleConditions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizationServerPolicyRuleConditionsWithDefaults

`func NewAuthorizationServerPolicyRuleConditionsWithDefaults() *AuthorizationServerPolicyRuleConditions`

NewAuthorizationServerPolicyRuleConditionsWithDefaults instantiates a new AuthorizationServerPolicyRuleConditions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApp

`func (o *AuthorizationServerPolicyRuleConditions) GetApp() AppAndInstancePolicyRuleCondition`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *AuthorizationServerPolicyRuleConditions) GetAppOk() (*AppAndInstancePolicyRuleCondition, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *AuthorizationServerPolicyRuleConditions) SetApp(v AppAndInstancePolicyRuleCondition)`

SetApp sets App field to given value.

### HasApp

`func (o *AuthorizationServerPolicyRuleConditions) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetApps

`func (o *AuthorizationServerPolicyRuleConditions) GetApps() AppInstancePolicyRuleCondition`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *AuthorizationServerPolicyRuleConditions) GetAppsOk() (*AppInstancePolicyRuleCondition, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *AuthorizationServerPolicyRuleConditions) SetApps(v AppInstancePolicyRuleCondition)`

SetApps sets Apps field to given value.

### HasApps

`func (o *AuthorizationServerPolicyRuleConditions) HasApps() bool`

HasApps returns a boolean if a field has been set.

### GetAuthContext

`func (o *AuthorizationServerPolicyRuleConditions) GetAuthContext() PolicyRuleAuthContextCondition`

GetAuthContext returns the AuthContext field if non-nil, zero value otherwise.

### GetAuthContextOk

`func (o *AuthorizationServerPolicyRuleConditions) GetAuthContextOk() (*PolicyRuleAuthContextCondition, bool)`

GetAuthContextOk returns a tuple with the AuthContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthContext

`func (o *AuthorizationServerPolicyRuleConditions) SetAuthContext(v PolicyRuleAuthContextCondition)`

SetAuthContext sets AuthContext field to given value.

### HasAuthContext

`func (o *AuthorizationServerPolicyRuleConditions) HasAuthContext() bool`

HasAuthContext returns a boolean if a field has been set.

### GetAuthProvider

`func (o *AuthorizationServerPolicyRuleConditions) GetAuthProvider() PasswordPolicyAuthenticationProviderCondition`

GetAuthProvider returns the AuthProvider field if non-nil, zero value otherwise.

### GetAuthProviderOk

`func (o *AuthorizationServerPolicyRuleConditions) GetAuthProviderOk() (*PasswordPolicyAuthenticationProviderCondition, bool)`

GetAuthProviderOk returns a tuple with the AuthProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthProvider

`func (o *AuthorizationServerPolicyRuleConditions) SetAuthProvider(v PasswordPolicyAuthenticationProviderCondition)`

SetAuthProvider sets AuthProvider field to given value.

### HasAuthProvider

`func (o *AuthorizationServerPolicyRuleConditions) HasAuthProvider() bool`

HasAuthProvider returns a boolean if a field has been set.

### GetBeforeScheduledAction

`func (o *AuthorizationServerPolicyRuleConditions) GetBeforeScheduledAction() BeforeScheduledActionPolicyRuleCondition`

GetBeforeScheduledAction returns the BeforeScheduledAction field if non-nil, zero value otherwise.

### GetBeforeScheduledActionOk

`func (o *AuthorizationServerPolicyRuleConditions) GetBeforeScheduledActionOk() (*BeforeScheduledActionPolicyRuleCondition, bool)`

GetBeforeScheduledActionOk returns a tuple with the BeforeScheduledAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeforeScheduledAction

`func (o *AuthorizationServerPolicyRuleConditions) SetBeforeScheduledAction(v BeforeScheduledActionPolicyRuleCondition)`

SetBeforeScheduledAction sets BeforeScheduledAction field to given value.

### HasBeforeScheduledAction

`func (o *AuthorizationServerPolicyRuleConditions) HasBeforeScheduledAction() bool`

HasBeforeScheduledAction returns a boolean if a field has been set.

### GetClients

`func (o *AuthorizationServerPolicyRuleConditions) GetClients() ClientPolicyCondition`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *AuthorizationServerPolicyRuleConditions) GetClientsOk() (*ClientPolicyCondition, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *AuthorizationServerPolicyRuleConditions) SetClients(v ClientPolicyCondition)`

SetClients sets Clients field to given value.

### HasClients

`func (o *AuthorizationServerPolicyRuleConditions) HasClients() bool`

HasClients returns a boolean if a field has been set.

### GetContext

`func (o *AuthorizationServerPolicyRuleConditions) GetContext() ContextPolicyRuleCondition`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *AuthorizationServerPolicyRuleConditions) GetContextOk() (*ContextPolicyRuleCondition, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *AuthorizationServerPolicyRuleConditions) SetContext(v ContextPolicyRuleCondition)`

SetContext sets Context field to given value.

### HasContext

`func (o *AuthorizationServerPolicyRuleConditions) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetDevice

`func (o *AuthorizationServerPolicyRuleConditions) GetDevice() DevicePolicyRuleCondition`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *AuthorizationServerPolicyRuleConditions) GetDeviceOk() (*DevicePolicyRuleCondition, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *AuthorizationServerPolicyRuleConditions) SetDevice(v DevicePolicyRuleCondition)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *AuthorizationServerPolicyRuleConditions) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetGrantTypes

`func (o *AuthorizationServerPolicyRuleConditions) GetGrantTypes() GrantTypePolicyRuleCondition`

GetGrantTypes returns the GrantTypes field if non-nil, zero value otherwise.

### GetGrantTypesOk

`func (o *AuthorizationServerPolicyRuleConditions) GetGrantTypesOk() (*GrantTypePolicyRuleCondition, bool)`

GetGrantTypesOk returns a tuple with the GrantTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantTypes

`func (o *AuthorizationServerPolicyRuleConditions) SetGrantTypes(v GrantTypePolicyRuleCondition)`

SetGrantTypes sets GrantTypes field to given value.

### HasGrantTypes

`func (o *AuthorizationServerPolicyRuleConditions) HasGrantTypes() bool`

HasGrantTypes returns a boolean if a field has been set.

### GetGroups

`func (o *AuthorizationServerPolicyRuleConditions) GetGroups() GroupPolicyRuleCondition`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *AuthorizationServerPolicyRuleConditions) GetGroupsOk() (*GroupPolicyRuleCondition, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *AuthorizationServerPolicyRuleConditions) SetGroups(v GroupPolicyRuleCondition)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *AuthorizationServerPolicyRuleConditions) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetIdentityProvider

`func (o *AuthorizationServerPolicyRuleConditions) GetIdentityProvider() IdentityProviderPolicyRuleCondition`

GetIdentityProvider returns the IdentityProvider field if non-nil, zero value otherwise.

### GetIdentityProviderOk

`func (o *AuthorizationServerPolicyRuleConditions) GetIdentityProviderOk() (*IdentityProviderPolicyRuleCondition, bool)`

GetIdentityProviderOk returns a tuple with the IdentityProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProvider

`func (o *AuthorizationServerPolicyRuleConditions) SetIdentityProvider(v IdentityProviderPolicyRuleCondition)`

SetIdentityProvider sets IdentityProvider field to given value.

### HasIdentityProvider

`func (o *AuthorizationServerPolicyRuleConditions) HasIdentityProvider() bool`

HasIdentityProvider returns a boolean if a field has been set.

### GetMdmEnrollment

`func (o *AuthorizationServerPolicyRuleConditions) GetMdmEnrollment() MDMEnrollmentPolicyRuleCondition`

GetMdmEnrollment returns the MdmEnrollment field if non-nil, zero value otherwise.

### GetMdmEnrollmentOk

`func (o *AuthorizationServerPolicyRuleConditions) GetMdmEnrollmentOk() (*MDMEnrollmentPolicyRuleCondition, bool)`

GetMdmEnrollmentOk returns a tuple with the MdmEnrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdmEnrollment

`func (o *AuthorizationServerPolicyRuleConditions) SetMdmEnrollment(v MDMEnrollmentPolicyRuleCondition)`

SetMdmEnrollment sets MdmEnrollment field to given value.

### HasMdmEnrollment

`func (o *AuthorizationServerPolicyRuleConditions) HasMdmEnrollment() bool`

HasMdmEnrollment returns a boolean if a field has been set.

### GetNetwork

`func (o *AuthorizationServerPolicyRuleConditions) GetNetwork() PolicyNetworkCondition`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *AuthorizationServerPolicyRuleConditions) GetNetworkOk() (*PolicyNetworkCondition, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *AuthorizationServerPolicyRuleConditions) SetNetwork(v PolicyNetworkCondition)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *AuthorizationServerPolicyRuleConditions) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetPeople

`func (o *AuthorizationServerPolicyRuleConditions) GetPeople() PolicyPeopleCondition`

GetPeople returns the People field if non-nil, zero value otherwise.

### GetPeopleOk

`func (o *AuthorizationServerPolicyRuleConditions) GetPeopleOk() (*PolicyPeopleCondition, bool)`

GetPeopleOk returns a tuple with the People field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeople

`func (o *AuthorizationServerPolicyRuleConditions) SetPeople(v PolicyPeopleCondition)`

SetPeople sets People field to given value.

### HasPeople

`func (o *AuthorizationServerPolicyRuleConditions) HasPeople() bool`

HasPeople returns a boolean if a field has been set.

### GetPlatform

`func (o *AuthorizationServerPolicyRuleConditions) GetPlatform() PlatformPolicyRuleCondition`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *AuthorizationServerPolicyRuleConditions) GetPlatformOk() (*PlatformPolicyRuleCondition, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *AuthorizationServerPolicyRuleConditions) SetPlatform(v PlatformPolicyRuleCondition)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *AuthorizationServerPolicyRuleConditions) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetRisk

`func (o *AuthorizationServerPolicyRuleConditions) GetRisk() RiskPolicyRuleCondition`

GetRisk returns the Risk field if non-nil, zero value otherwise.

### GetRiskOk

`func (o *AuthorizationServerPolicyRuleConditions) GetRiskOk() (*RiskPolicyRuleCondition, bool)`

GetRiskOk returns a tuple with the Risk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRisk

`func (o *AuthorizationServerPolicyRuleConditions) SetRisk(v RiskPolicyRuleCondition)`

SetRisk sets Risk field to given value.

### HasRisk

`func (o *AuthorizationServerPolicyRuleConditions) HasRisk() bool`

HasRisk returns a boolean if a field has been set.

### GetRiskScore

`func (o *AuthorizationServerPolicyRuleConditions) GetRiskScore() RiskScorePolicyRuleCondition`

GetRiskScore returns the RiskScore field if non-nil, zero value otherwise.

### GetRiskScoreOk

`func (o *AuthorizationServerPolicyRuleConditions) GetRiskScoreOk() (*RiskScorePolicyRuleCondition, bool)`

GetRiskScoreOk returns a tuple with the RiskScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskScore

`func (o *AuthorizationServerPolicyRuleConditions) SetRiskScore(v RiskScorePolicyRuleCondition)`

SetRiskScore sets RiskScore field to given value.

### HasRiskScore

`func (o *AuthorizationServerPolicyRuleConditions) HasRiskScore() bool`

HasRiskScore returns a boolean if a field has been set.

### GetScopes

`func (o *AuthorizationServerPolicyRuleConditions) GetScopes() OAuth2ScopesMediationPolicyRuleCondition`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *AuthorizationServerPolicyRuleConditions) GetScopesOk() (*OAuth2ScopesMediationPolicyRuleCondition, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *AuthorizationServerPolicyRuleConditions) SetScopes(v OAuth2ScopesMediationPolicyRuleCondition)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *AuthorizationServerPolicyRuleConditions) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetUserIdentifier

`func (o *AuthorizationServerPolicyRuleConditions) GetUserIdentifier() UserIdentifierPolicyRuleCondition`

GetUserIdentifier returns the UserIdentifier field if non-nil, zero value otherwise.

### GetUserIdentifierOk

`func (o *AuthorizationServerPolicyRuleConditions) GetUserIdentifierOk() (*UserIdentifierPolicyRuleCondition, bool)`

GetUserIdentifierOk returns a tuple with the UserIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdentifier

`func (o *AuthorizationServerPolicyRuleConditions) SetUserIdentifier(v UserIdentifierPolicyRuleCondition)`

SetUserIdentifier sets UserIdentifier field to given value.

### HasUserIdentifier

`func (o *AuthorizationServerPolicyRuleConditions) HasUserIdentifier() bool`

HasUserIdentifier returns a boolean if a field has been set.

### GetUsers

`func (o *AuthorizationServerPolicyRuleConditions) GetUsers() UserPolicyRuleCondition`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *AuthorizationServerPolicyRuleConditions) GetUsersOk() (*UserPolicyRuleCondition, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *AuthorizationServerPolicyRuleConditions) SetUsers(v UserPolicyRuleCondition)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *AuthorizationServerPolicyRuleConditions) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetUserStatus

`func (o *AuthorizationServerPolicyRuleConditions) GetUserStatus() UserStatusPolicyRuleCondition`

GetUserStatus returns the UserStatus field if non-nil, zero value otherwise.

### GetUserStatusOk

`func (o *AuthorizationServerPolicyRuleConditions) GetUserStatusOk() (*UserStatusPolicyRuleCondition, bool)`

GetUserStatusOk returns a tuple with the UserStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserStatus

`func (o *AuthorizationServerPolicyRuleConditions) SetUserStatus(v UserStatusPolicyRuleCondition)`

SetUserStatus sets UserStatus field to given value.

### HasUserStatus

`func (o *AuthorizationServerPolicyRuleConditions) HasUserStatus() bool`

HasUserStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


