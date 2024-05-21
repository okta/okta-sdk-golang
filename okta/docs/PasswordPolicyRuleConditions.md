# PasswordPolicyRuleConditions

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

### NewPasswordPolicyRuleConditions

`func NewPasswordPolicyRuleConditions() *PasswordPolicyRuleConditions`

NewPasswordPolicyRuleConditions instantiates a new PasswordPolicyRuleConditions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPasswordPolicyRuleConditionsWithDefaults

`func NewPasswordPolicyRuleConditionsWithDefaults() *PasswordPolicyRuleConditions`

NewPasswordPolicyRuleConditionsWithDefaults instantiates a new PasswordPolicyRuleConditions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApp

`func (o *PasswordPolicyRuleConditions) GetApp() AppAndInstancePolicyRuleCondition`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *PasswordPolicyRuleConditions) GetAppOk() (*AppAndInstancePolicyRuleCondition, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *PasswordPolicyRuleConditions) SetApp(v AppAndInstancePolicyRuleCondition)`

SetApp sets App field to given value.

### HasApp

`func (o *PasswordPolicyRuleConditions) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetApps

`func (o *PasswordPolicyRuleConditions) GetApps() AppInstancePolicyRuleCondition`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *PasswordPolicyRuleConditions) GetAppsOk() (*AppInstancePolicyRuleCondition, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *PasswordPolicyRuleConditions) SetApps(v AppInstancePolicyRuleCondition)`

SetApps sets Apps field to given value.

### HasApps

`func (o *PasswordPolicyRuleConditions) HasApps() bool`

HasApps returns a boolean if a field has been set.

### GetAuthContext

`func (o *PasswordPolicyRuleConditions) GetAuthContext() PolicyRuleAuthContextCondition`

GetAuthContext returns the AuthContext field if non-nil, zero value otherwise.

### GetAuthContextOk

`func (o *PasswordPolicyRuleConditions) GetAuthContextOk() (*PolicyRuleAuthContextCondition, bool)`

GetAuthContextOk returns a tuple with the AuthContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthContext

`func (o *PasswordPolicyRuleConditions) SetAuthContext(v PolicyRuleAuthContextCondition)`

SetAuthContext sets AuthContext field to given value.

### HasAuthContext

`func (o *PasswordPolicyRuleConditions) HasAuthContext() bool`

HasAuthContext returns a boolean if a field has been set.

### GetAuthProvider

`func (o *PasswordPolicyRuleConditions) GetAuthProvider() PasswordPolicyAuthenticationProviderCondition`

GetAuthProvider returns the AuthProvider field if non-nil, zero value otherwise.

### GetAuthProviderOk

`func (o *PasswordPolicyRuleConditions) GetAuthProviderOk() (*PasswordPolicyAuthenticationProviderCondition, bool)`

GetAuthProviderOk returns a tuple with the AuthProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthProvider

`func (o *PasswordPolicyRuleConditions) SetAuthProvider(v PasswordPolicyAuthenticationProviderCondition)`

SetAuthProvider sets AuthProvider field to given value.

### HasAuthProvider

`func (o *PasswordPolicyRuleConditions) HasAuthProvider() bool`

HasAuthProvider returns a boolean if a field has been set.

### GetBeforeScheduledAction

`func (o *PasswordPolicyRuleConditions) GetBeforeScheduledAction() BeforeScheduledActionPolicyRuleCondition`

GetBeforeScheduledAction returns the BeforeScheduledAction field if non-nil, zero value otherwise.

### GetBeforeScheduledActionOk

`func (o *PasswordPolicyRuleConditions) GetBeforeScheduledActionOk() (*BeforeScheduledActionPolicyRuleCondition, bool)`

GetBeforeScheduledActionOk returns a tuple with the BeforeScheduledAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeforeScheduledAction

`func (o *PasswordPolicyRuleConditions) SetBeforeScheduledAction(v BeforeScheduledActionPolicyRuleCondition)`

SetBeforeScheduledAction sets BeforeScheduledAction field to given value.

### HasBeforeScheduledAction

`func (o *PasswordPolicyRuleConditions) HasBeforeScheduledAction() bool`

HasBeforeScheduledAction returns a boolean if a field has been set.

### GetClients

`func (o *PasswordPolicyRuleConditions) GetClients() ClientPolicyCondition`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *PasswordPolicyRuleConditions) GetClientsOk() (*ClientPolicyCondition, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *PasswordPolicyRuleConditions) SetClients(v ClientPolicyCondition)`

SetClients sets Clients field to given value.

### HasClients

`func (o *PasswordPolicyRuleConditions) HasClients() bool`

HasClients returns a boolean if a field has been set.

### GetContext

`func (o *PasswordPolicyRuleConditions) GetContext() ContextPolicyRuleCondition`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *PasswordPolicyRuleConditions) GetContextOk() (*ContextPolicyRuleCondition, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *PasswordPolicyRuleConditions) SetContext(v ContextPolicyRuleCondition)`

SetContext sets Context field to given value.

### HasContext

`func (o *PasswordPolicyRuleConditions) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetDevice

`func (o *PasswordPolicyRuleConditions) GetDevice() DevicePolicyRuleCondition`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *PasswordPolicyRuleConditions) GetDeviceOk() (*DevicePolicyRuleCondition, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *PasswordPolicyRuleConditions) SetDevice(v DevicePolicyRuleCondition)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *PasswordPolicyRuleConditions) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetGrantTypes

`func (o *PasswordPolicyRuleConditions) GetGrantTypes() GrantTypePolicyRuleCondition`

GetGrantTypes returns the GrantTypes field if non-nil, zero value otherwise.

### GetGrantTypesOk

`func (o *PasswordPolicyRuleConditions) GetGrantTypesOk() (*GrantTypePolicyRuleCondition, bool)`

GetGrantTypesOk returns a tuple with the GrantTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantTypes

`func (o *PasswordPolicyRuleConditions) SetGrantTypes(v GrantTypePolicyRuleCondition)`

SetGrantTypes sets GrantTypes field to given value.

### HasGrantTypes

`func (o *PasswordPolicyRuleConditions) HasGrantTypes() bool`

HasGrantTypes returns a boolean if a field has been set.

### GetGroups

`func (o *PasswordPolicyRuleConditions) GetGroups() GroupPolicyRuleCondition`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *PasswordPolicyRuleConditions) GetGroupsOk() (*GroupPolicyRuleCondition, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *PasswordPolicyRuleConditions) SetGroups(v GroupPolicyRuleCondition)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *PasswordPolicyRuleConditions) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetIdentityProvider

`func (o *PasswordPolicyRuleConditions) GetIdentityProvider() IdentityProviderPolicyRuleCondition`

GetIdentityProvider returns the IdentityProvider field if non-nil, zero value otherwise.

### GetIdentityProviderOk

`func (o *PasswordPolicyRuleConditions) GetIdentityProviderOk() (*IdentityProviderPolicyRuleCondition, bool)`

GetIdentityProviderOk returns a tuple with the IdentityProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProvider

`func (o *PasswordPolicyRuleConditions) SetIdentityProvider(v IdentityProviderPolicyRuleCondition)`

SetIdentityProvider sets IdentityProvider field to given value.

### HasIdentityProvider

`func (o *PasswordPolicyRuleConditions) HasIdentityProvider() bool`

HasIdentityProvider returns a boolean if a field has been set.

### GetMdmEnrollment

`func (o *PasswordPolicyRuleConditions) GetMdmEnrollment() MDMEnrollmentPolicyRuleCondition`

GetMdmEnrollment returns the MdmEnrollment field if non-nil, zero value otherwise.

### GetMdmEnrollmentOk

`func (o *PasswordPolicyRuleConditions) GetMdmEnrollmentOk() (*MDMEnrollmentPolicyRuleCondition, bool)`

GetMdmEnrollmentOk returns a tuple with the MdmEnrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdmEnrollment

`func (o *PasswordPolicyRuleConditions) SetMdmEnrollment(v MDMEnrollmentPolicyRuleCondition)`

SetMdmEnrollment sets MdmEnrollment field to given value.

### HasMdmEnrollment

`func (o *PasswordPolicyRuleConditions) HasMdmEnrollment() bool`

HasMdmEnrollment returns a boolean if a field has been set.

### GetNetwork

`func (o *PasswordPolicyRuleConditions) GetNetwork() PolicyNetworkCondition`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *PasswordPolicyRuleConditions) GetNetworkOk() (*PolicyNetworkCondition, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *PasswordPolicyRuleConditions) SetNetwork(v PolicyNetworkCondition)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *PasswordPolicyRuleConditions) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetPeople

`func (o *PasswordPolicyRuleConditions) GetPeople() PolicyPeopleCondition`

GetPeople returns the People field if non-nil, zero value otherwise.

### GetPeopleOk

`func (o *PasswordPolicyRuleConditions) GetPeopleOk() (*PolicyPeopleCondition, bool)`

GetPeopleOk returns a tuple with the People field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeople

`func (o *PasswordPolicyRuleConditions) SetPeople(v PolicyPeopleCondition)`

SetPeople sets People field to given value.

### HasPeople

`func (o *PasswordPolicyRuleConditions) HasPeople() bool`

HasPeople returns a boolean if a field has been set.

### GetPlatform

`func (o *PasswordPolicyRuleConditions) GetPlatform() PlatformPolicyRuleCondition`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *PasswordPolicyRuleConditions) GetPlatformOk() (*PlatformPolicyRuleCondition, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *PasswordPolicyRuleConditions) SetPlatform(v PlatformPolicyRuleCondition)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *PasswordPolicyRuleConditions) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetRisk

`func (o *PasswordPolicyRuleConditions) GetRisk() RiskPolicyRuleCondition`

GetRisk returns the Risk field if non-nil, zero value otherwise.

### GetRiskOk

`func (o *PasswordPolicyRuleConditions) GetRiskOk() (*RiskPolicyRuleCondition, bool)`

GetRiskOk returns a tuple with the Risk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRisk

`func (o *PasswordPolicyRuleConditions) SetRisk(v RiskPolicyRuleCondition)`

SetRisk sets Risk field to given value.

### HasRisk

`func (o *PasswordPolicyRuleConditions) HasRisk() bool`

HasRisk returns a boolean if a field has been set.

### GetRiskScore

`func (o *PasswordPolicyRuleConditions) GetRiskScore() RiskScorePolicyRuleCondition`

GetRiskScore returns the RiskScore field if non-nil, zero value otherwise.

### GetRiskScoreOk

`func (o *PasswordPolicyRuleConditions) GetRiskScoreOk() (*RiskScorePolicyRuleCondition, bool)`

GetRiskScoreOk returns a tuple with the RiskScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskScore

`func (o *PasswordPolicyRuleConditions) SetRiskScore(v RiskScorePolicyRuleCondition)`

SetRiskScore sets RiskScore field to given value.

### HasRiskScore

`func (o *PasswordPolicyRuleConditions) HasRiskScore() bool`

HasRiskScore returns a boolean if a field has been set.

### GetScopes

`func (o *PasswordPolicyRuleConditions) GetScopes() OAuth2ScopesMediationPolicyRuleCondition`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *PasswordPolicyRuleConditions) GetScopesOk() (*OAuth2ScopesMediationPolicyRuleCondition, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *PasswordPolicyRuleConditions) SetScopes(v OAuth2ScopesMediationPolicyRuleCondition)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *PasswordPolicyRuleConditions) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetUserIdentifier

`func (o *PasswordPolicyRuleConditions) GetUserIdentifier() UserIdentifierPolicyRuleCondition`

GetUserIdentifier returns the UserIdentifier field if non-nil, zero value otherwise.

### GetUserIdentifierOk

`func (o *PasswordPolicyRuleConditions) GetUserIdentifierOk() (*UserIdentifierPolicyRuleCondition, bool)`

GetUserIdentifierOk returns a tuple with the UserIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdentifier

`func (o *PasswordPolicyRuleConditions) SetUserIdentifier(v UserIdentifierPolicyRuleCondition)`

SetUserIdentifier sets UserIdentifier field to given value.

### HasUserIdentifier

`func (o *PasswordPolicyRuleConditions) HasUserIdentifier() bool`

HasUserIdentifier returns a boolean if a field has been set.

### GetUsers

`func (o *PasswordPolicyRuleConditions) GetUsers() UserPolicyRuleCondition`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *PasswordPolicyRuleConditions) GetUsersOk() (*UserPolicyRuleCondition, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *PasswordPolicyRuleConditions) SetUsers(v UserPolicyRuleCondition)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *PasswordPolicyRuleConditions) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetUserStatus

`func (o *PasswordPolicyRuleConditions) GetUserStatus() UserStatusPolicyRuleCondition`

GetUserStatus returns the UserStatus field if non-nil, zero value otherwise.

### GetUserStatusOk

`func (o *PasswordPolicyRuleConditions) GetUserStatusOk() (*UserStatusPolicyRuleCondition, bool)`

GetUserStatusOk returns a tuple with the UserStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserStatus

`func (o *PasswordPolicyRuleConditions) SetUserStatus(v UserStatusPolicyRuleCondition)`

SetUserStatus sets UserStatus field to given value.

### HasUserStatus

`func (o *PasswordPolicyRuleConditions) HasUserStatus() bool`

HasUserStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


