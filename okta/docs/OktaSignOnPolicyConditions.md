# OktaSignOnPolicyConditions

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

### NewOktaSignOnPolicyConditions

`func NewOktaSignOnPolicyConditions() *OktaSignOnPolicyConditions`

NewOktaSignOnPolicyConditions instantiates a new OktaSignOnPolicyConditions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOktaSignOnPolicyConditionsWithDefaults

`func NewOktaSignOnPolicyConditionsWithDefaults() *OktaSignOnPolicyConditions`

NewOktaSignOnPolicyConditionsWithDefaults instantiates a new OktaSignOnPolicyConditions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApp

`func (o *OktaSignOnPolicyConditions) GetApp() AppAndInstancePolicyRuleCondition`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *OktaSignOnPolicyConditions) GetAppOk() (*AppAndInstancePolicyRuleCondition, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *OktaSignOnPolicyConditions) SetApp(v AppAndInstancePolicyRuleCondition)`

SetApp sets App field to given value.

### HasApp

`func (o *OktaSignOnPolicyConditions) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetApps

`func (o *OktaSignOnPolicyConditions) GetApps() AppInstancePolicyRuleCondition`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *OktaSignOnPolicyConditions) GetAppsOk() (*AppInstancePolicyRuleCondition, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *OktaSignOnPolicyConditions) SetApps(v AppInstancePolicyRuleCondition)`

SetApps sets Apps field to given value.

### HasApps

`func (o *OktaSignOnPolicyConditions) HasApps() bool`

HasApps returns a boolean if a field has been set.

### GetAuthContext

`func (o *OktaSignOnPolicyConditions) GetAuthContext() PolicyRuleAuthContextCondition`

GetAuthContext returns the AuthContext field if non-nil, zero value otherwise.

### GetAuthContextOk

`func (o *OktaSignOnPolicyConditions) GetAuthContextOk() (*PolicyRuleAuthContextCondition, bool)`

GetAuthContextOk returns a tuple with the AuthContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthContext

`func (o *OktaSignOnPolicyConditions) SetAuthContext(v PolicyRuleAuthContextCondition)`

SetAuthContext sets AuthContext field to given value.

### HasAuthContext

`func (o *OktaSignOnPolicyConditions) HasAuthContext() bool`

HasAuthContext returns a boolean if a field has been set.

### GetAuthProvider

`func (o *OktaSignOnPolicyConditions) GetAuthProvider() PasswordPolicyAuthenticationProviderCondition`

GetAuthProvider returns the AuthProvider field if non-nil, zero value otherwise.

### GetAuthProviderOk

`func (o *OktaSignOnPolicyConditions) GetAuthProviderOk() (*PasswordPolicyAuthenticationProviderCondition, bool)`

GetAuthProviderOk returns a tuple with the AuthProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthProvider

`func (o *OktaSignOnPolicyConditions) SetAuthProvider(v PasswordPolicyAuthenticationProviderCondition)`

SetAuthProvider sets AuthProvider field to given value.

### HasAuthProvider

`func (o *OktaSignOnPolicyConditions) HasAuthProvider() bool`

HasAuthProvider returns a boolean if a field has been set.

### GetBeforeScheduledAction

`func (o *OktaSignOnPolicyConditions) GetBeforeScheduledAction() BeforeScheduledActionPolicyRuleCondition`

GetBeforeScheduledAction returns the BeforeScheduledAction field if non-nil, zero value otherwise.

### GetBeforeScheduledActionOk

`func (o *OktaSignOnPolicyConditions) GetBeforeScheduledActionOk() (*BeforeScheduledActionPolicyRuleCondition, bool)`

GetBeforeScheduledActionOk returns a tuple with the BeforeScheduledAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeforeScheduledAction

`func (o *OktaSignOnPolicyConditions) SetBeforeScheduledAction(v BeforeScheduledActionPolicyRuleCondition)`

SetBeforeScheduledAction sets BeforeScheduledAction field to given value.

### HasBeforeScheduledAction

`func (o *OktaSignOnPolicyConditions) HasBeforeScheduledAction() bool`

HasBeforeScheduledAction returns a boolean if a field has been set.

### GetClients

`func (o *OktaSignOnPolicyConditions) GetClients() ClientPolicyCondition`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *OktaSignOnPolicyConditions) GetClientsOk() (*ClientPolicyCondition, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *OktaSignOnPolicyConditions) SetClients(v ClientPolicyCondition)`

SetClients sets Clients field to given value.

### HasClients

`func (o *OktaSignOnPolicyConditions) HasClients() bool`

HasClients returns a boolean if a field has been set.

### GetContext

`func (o *OktaSignOnPolicyConditions) GetContext() ContextPolicyRuleCondition`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *OktaSignOnPolicyConditions) GetContextOk() (*ContextPolicyRuleCondition, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *OktaSignOnPolicyConditions) SetContext(v ContextPolicyRuleCondition)`

SetContext sets Context field to given value.

### HasContext

`func (o *OktaSignOnPolicyConditions) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetDevice

`func (o *OktaSignOnPolicyConditions) GetDevice() DevicePolicyRuleCondition`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *OktaSignOnPolicyConditions) GetDeviceOk() (*DevicePolicyRuleCondition, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *OktaSignOnPolicyConditions) SetDevice(v DevicePolicyRuleCondition)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *OktaSignOnPolicyConditions) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetGrantTypes

`func (o *OktaSignOnPolicyConditions) GetGrantTypes() GrantTypePolicyRuleCondition`

GetGrantTypes returns the GrantTypes field if non-nil, zero value otherwise.

### GetGrantTypesOk

`func (o *OktaSignOnPolicyConditions) GetGrantTypesOk() (*GrantTypePolicyRuleCondition, bool)`

GetGrantTypesOk returns a tuple with the GrantTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantTypes

`func (o *OktaSignOnPolicyConditions) SetGrantTypes(v GrantTypePolicyRuleCondition)`

SetGrantTypes sets GrantTypes field to given value.

### HasGrantTypes

`func (o *OktaSignOnPolicyConditions) HasGrantTypes() bool`

HasGrantTypes returns a boolean if a field has been set.

### GetGroups

`func (o *OktaSignOnPolicyConditions) GetGroups() GroupPolicyRuleCondition`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *OktaSignOnPolicyConditions) GetGroupsOk() (*GroupPolicyRuleCondition, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *OktaSignOnPolicyConditions) SetGroups(v GroupPolicyRuleCondition)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *OktaSignOnPolicyConditions) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetIdentityProvider

`func (o *OktaSignOnPolicyConditions) GetIdentityProvider() IdentityProviderPolicyRuleCondition`

GetIdentityProvider returns the IdentityProvider field if non-nil, zero value otherwise.

### GetIdentityProviderOk

`func (o *OktaSignOnPolicyConditions) GetIdentityProviderOk() (*IdentityProviderPolicyRuleCondition, bool)`

GetIdentityProviderOk returns a tuple with the IdentityProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProvider

`func (o *OktaSignOnPolicyConditions) SetIdentityProvider(v IdentityProviderPolicyRuleCondition)`

SetIdentityProvider sets IdentityProvider field to given value.

### HasIdentityProvider

`func (o *OktaSignOnPolicyConditions) HasIdentityProvider() bool`

HasIdentityProvider returns a boolean if a field has been set.

### GetMdmEnrollment

`func (o *OktaSignOnPolicyConditions) GetMdmEnrollment() MDMEnrollmentPolicyRuleCondition`

GetMdmEnrollment returns the MdmEnrollment field if non-nil, zero value otherwise.

### GetMdmEnrollmentOk

`func (o *OktaSignOnPolicyConditions) GetMdmEnrollmentOk() (*MDMEnrollmentPolicyRuleCondition, bool)`

GetMdmEnrollmentOk returns a tuple with the MdmEnrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdmEnrollment

`func (o *OktaSignOnPolicyConditions) SetMdmEnrollment(v MDMEnrollmentPolicyRuleCondition)`

SetMdmEnrollment sets MdmEnrollment field to given value.

### HasMdmEnrollment

`func (o *OktaSignOnPolicyConditions) HasMdmEnrollment() bool`

HasMdmEnrollment returns a boolean if a field has been set.

### GetNetwork

`func (o *OktaSignOnPolicyConditions) GetNetwork() PolicyNetworkCondition`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *OktaSignOnPolicyConditions) GetNetworkOk() (*PolicyNetworkCondition, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *OktaSignOnPolicyConditions) SetNetwork(v PolicyNetworkCondition)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *OktaSignOnPolicyConditions) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetPeople

`func (o *OktaSignOnPolicyConditions) GetPeople() PolicyPeopleCondition`

GetPeople returns the People field if non-nil, zero value otherwise.

### GetPeopleOk

`func (o *OktaSignOnPolicyConditions) GetPeopleOk() (*PolicyPeopleCondition, bool)`

GetPeopleOk returns a tuple with the People field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeople

`func (o *OktaSignOnPolicyConditions) SetPeople(v PolicyPeopleCondition)`

SetPeople sets People field to given value.

### HasPeople

`func (o *OktaSignOnPolicyConditions) HasPeople() bool`

HasPeople returns a boolean if a field has been set.

### GetPlatform

`func (o *OktaSignOnPolicyConditions) GetPlatform() PlatformPolicyRuleCondition`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *OktaSignOnPolicyConditions) GetPlatformOk() (*PlatformPolicyRuleCondition, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *OktaSignOnPolicyConditions) SetPlatform(v PlatformPolicyRuleCondition)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *OktaSignOnPolicyConditions) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetRisk

`func (o *OktaSignOnPolicyConditions) GetRisk() RiskPolicyRuleCondition`

GetRisk returns the Risk field if non-nil, zero value otherwise.

### GetRiskOk

`func (o *OktaSignOnPolicyConditions) GetRiskOk() (*RiskPolicyRuleCondition, bool)`

GetRiskOk returns a tuple with the Risk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRisk

`func (o *OktaSignOnPolicyConditions) SetRisk(v RiskPolicyRuleCondition)`

SetRisk sets Risk field to given value.

### HasRisk

`func (o *OktaSignOnPolicyConditions) HasRisk() bool`

HasRisk returns a boolean if a field has been set.

### GetRiskScore

`func (o *OktaSignOnPolicyConditions) GetRiskScore() RiskScorePolicyRuleCondition`

GetRiskScore returns the RiskScore field if non-nil, zero value otherwise.

### GetRiskScoreOk

`func (o *OktaSignOnPolicyConditions) GetRiskScoreOk() (*RiskScorePolicyRuleCondition, bool)`

GetRiskScoreOk returns a tuple with the RiskScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskScore

`func (o *OktaSignOnPolicyConditions) SetRiskScore(v RiskScorePolicyRuleCondition)`

SetRiskScore sets RiskScore field to given value.

### HasRiskScore

`func (o *OktaSignOnPolicyConditions) HasRiskScore() bool`

HasRiskScore returns a boolean if a field has been set.

### GetScopes

`func (o *OktaSignOnPolicyConditions) GetScopes() OAuth2ScopesMediationPolicyRuleCondition`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *OktaSignOnPolicyConditions) GetScopesOk() (*OAuth2ScopesMediationPolicyRuleCondition, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *OktaSignOnPolicyConditions) SetScopes(v OAuth2ScopesMediationPolicyRuleCondition)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *OktaSignOnPolicyConditions) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetUserIdentifier

`func (o *OktaSignOnPolicyConditions) GetUserIdentifier() UserIdentifierPolicyRuleCondition`

GetUserIdentifier returns the UserIdentifier field if non-nil, zero value otherwise.

### GetUserIdentifierOk

`func (o *OktaSignOnPolicyConditions) GetUserIdentifierOk() (*UserIdentifierPolicyRuleCondition, bool)`

GetUserIdentifierOk returns a tuple with the UserIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdentifier

`func (o *OktaSignOnPolicyConditions) SetUserIdentifier(v UserIdentifierPolicyRuleCondition)`

SetUserIdentifier sets UserIdentifier field to given value.

### HasUserIdentifier

`func (o *OktaSignOnPolicyConditions) HasUserIdentifier() bool`

HasUserIdentifier returns a boolean if a field has been set.

### GetUsers

`func (o *OktaSignOnPolicyConditions) GetUsers() UserPolicyRuleCondition`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *OktaSignOnPolicyConditions) GetUsersOk() (*UserPolicyRuleCondition, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *OktaSignOnPolicyConditions) SetUsers(v UserPolicyRuleCondition)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *OktaSignOnPolicyConditions) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetUserStatus

`func (o *OktaSignOnPolicyConditions) GetUserStatus() UserStatusPolicyRuleCondition`

GetUserStatus returns the UserStatus field if non-nil, zero value otherwise.

### GetUserStatusOk

`func (o *OktaSignOnPolicyConditions) GetUserStatusOk() (*UserStatusPolicyRuleCondition, bool)`

GetUserStatusOk returns a tuple with the UserStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserStatus

`func (o *OktaSignOnPolicyConditions) SetUserStatus(v UserStatusPolicyRuleCondition)`

SetUserStatus sets UserStatus field to given value.

### HasUserStatus

`func (o *OktaSignOnPolicyConditions) HasUserStatus() bool`

HasUserStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


