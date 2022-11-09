# AccessPolicyRuleConditions

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
**Device** | Pointer to [**DeviceAccessPolicyRuleCondition**](DeviceAccessPolicyRuleCondition.md) |  | [optional] 
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
**ElCondition** | Pointer to [**AccessPolicyRuleCustomCondition**](AccessPolicyRuleCustomCondition.md) |  | [optional] 
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

### GetApp

`func (o *AccessPolicyRuleConditions) GetApp() AppAndInstancePolicyRuleCondition`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *AccessPolicyRuleConditions) GetAppOk() (*AppAndInstancePolicyRuleCondition, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *AccessPolicyRuleConditions) SetApp(v AppAndInstancePolicyRuleCondition)`

SetApp sets App field to given value.

### HasApp

`func (o *AccessPolicyRuleConditions) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetApps

`func (o *AccessPolicyRuleConditions) GetApps() AppInstancePolicyRuleCondition`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *AccessPolicyRuleConditions) GetAppsOk() (*AppInstancePolicyRuleCondition, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *AccessPolicyRuleConditions) SetApps(v AppInstancePolicyRuleCondition)`

SetApps sets Apps field to given value.

### HasApps

`func (o *AccessPolicyRuleConditions) HasApps() bool`

HasApps returns a boolean if a field has been set.

### GetAuthContext

`func (o *AccessPolicyRuleConditions) GetAuthContext() PolicyRuleAuthContextCondition`

GetAuthContext returns the AuthContext field if non-nil, zero value otherwise.

### GetAuthContextOk

`func (o *AccessPolicyRuleConditions) GetAuthContextOk() (*PolicyRuleAuthContextCondition, bool)`

GetAuthContextOk returns a tuple with the AuthContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthContext

`func (o *AccessPolicyRuleConditions) SetAuthContext(v PolicyRuleAuthContextCondition)`

SetAuthContext sets AuthContext field to given value.

### HasAuthContext

`func (o *AccessPolicyRuleConditions) HasAuthContext() bool`

HasAuthContext returns a boolean if a field has been set.

### GetAuthProvider

`func (o *AccessPolicyRuleConditions) GetAuthProvider() PasswordPolicyAuthenticationProviderCondition`

GetAuthProvider returns the AuthProvider field if non-nil, zero value otherwise.

### GetAuthProviderOk

`func (o *AccessPolicyRuleConditions) GetAuthProviderOk() (*PasswordPolicyAuthenticationProviderCondition, bool)`

GetAuthProviderOk returns a tuple with the AuthProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthProvider

`func (o *AccessPolicyRuleConditions) SetAuthProvider(v PasswordPolicyAuthenticationProviderCondition)`

SetAuthProvider sets AuthProvider field to given value.

### HasAuthProvider

`func (o *AccessPolicyRuleConditions) HasAuthProvider() bool`

HasAuthProvider returns a boolean if a field has been set.

### GetBeforeScheduledAction

`func (o *AccessPolicyRuleConditions) GetBeforeScheduledAction() BeforeScheduledActionPolicyRuleCondition`

GetBeforeScheduledAction returns the BeforeScheduledAction field if non-nil, zero value otherwise.

### GetBeforeScheduledActionOk

`func (o *AccessPolicyRuleConditions) GetBeforeScheduledActionOk() (*BeforeScheduledActionPolicyRuleCondition, bool)`

GetBeforeScheduledActionOk returns a tuple with the BeforeScheduledAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeforeScheduledAction

`func (o *AccessPolicyRuleConditions) SetBeforeScheduledAction(v BeforeScheduledActionPolicyRuleCondition)`

SetBeforeScheduledAction sets BeforeScheduledAction field to given value.

### HasBeforeScheduledAction

`func (o *AccessPolicyRuleConditions) HasBeforeScheduledAction() bool`

HasBeforeScheduledAction returns a boolean if a field has been set.

### GetClients

`func (o *AccessPolicyRuleConditions) GetClients() ClientPolicyCondition`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *AccessPolicyRuleConditions) GetClientsOk() (*ClientPolicyCondition, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *AccessPolicyRuleConditions) SetClients(v ClientPolicyCondition)`

SetClients sets Clients field to given value.

### HasClients

`func (o *AccessPolicyRuleConditions) HasClients() bool`

HasClients returns a boolean if a field has been set.

### GetContext

`func (o *AccessPolicyRuleConditions) GetContext() ContextPolicyRuleCondition`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *AccessPolicyRuleConditions) GetContextOk() (*ContextPolicyRuleCondition, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *AccessPolicyRuleConditions) SetContext(v ContextPolicyRuleCondition)`

SetContext sets Context field to given value.

### HasContext

`func (o *AccessPolicyRuleConditions) HasContext() bool`

HasContext returns a boolean if a field has been set.

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

### GetGrantTypes

`func (o *AccessPolicyRuleConditions) GetGrantTypes() GrantTypePolicyRuleCondition`

GetGrantTypes returns the GrantTypes field if non-nil, zero value otherwise.

### GetGrantTypesOk

`func (o *AccessPolicyRuleConditions) GetGrantTypesOk() (*GrantTypePolicyRuleCondition, bool)`

GetGrantTypesOk returns a tuple with the GrantTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantTypes

`func (o *AccessPolicyRuleConditions) SetGrantTypes(v GrantTypePolicyRuleCondition)`

SetGrantTypes sets GrantTypes field to given value.

### HasGrantTypes

`func (o *AccessPolicyRuleConditions) HasGrantTypes() bool`

HasGrantTypes returns a boolean if a field has been set.

### GetGroups

`func (o *AccessPolicyRuleConditions) GetGroups() GroupPolicyRuleCondition`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *AccessPolicyRuleConditions) GetGroupsOk() (*GroupPolicyRuleCondition, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *AccessPolicyRuleConditions) SetGroups(v GroupPolicyRuleCondition)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *AccessPolicyRuleConditions) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetIdentityProvider

`func (o *AccessPolicyRuleConditions) GetIdentityProvider() IdentityProviderPolicyRuleCondition`

GetIdentityProvider returns the IdentityProvider field if non-nil, zero value otherwise.

### GetIdentityProviderOk

`func (o *AccessPolicyRuleConditions) GetIdentityProviderOk() (*IdentityProviderPolicyRuleCondition, bool)`

GetIdentityProviderOk returns a tuple with the IdentityProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProvider

`func (o *AccessPolicyRuleConditions) SetIdentityProvider(v IdentityProviderPolicyRuleCondition)`

SetIdentityProvider sets IdentityProvider field to given value.

### HasIdentityProvider

`func (o *AccessPolicyRuleConditions) HasIdentityProvider() bool`

HasIdentityProvider returns a boolean if a field has been set.

### GetMdmEnrollment

`func (o *AccessPolicyRuleConditions) GetMdmEnrollment() MDMEnrollmentPolicyRuleCondition`

GetMdmEnrollment returns the MdmEnrollment field if non-nil, zero value otherwise.

### GetMdmEnrollmentOk

`func (o *AccessPolicyRuleConditions) GetMdmEnrollmentOk() (*MDMEnrollmentPolicyRuleCondition, bool)`

GetMdmEnrollmentOk returns a tuple with the MdmEnrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdmEnrollment

`func (o *AccessPolicyRuleConditions) SetMdmEnrollment(v MDMEnrollmentPolicyRuleCondition)`

SetMdmEnrollment sets MdmEnrollment field to given value.

### HasMdmEnrollment

`func (o *AccessPolicyRuleConditions) HasMdmEnrollment() bool`

HasMdmEnrollment returns a boolean if a field has been set.

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

### GetRisk

`func (o *AccessPolicyRuleConditions) GetRisk() RiskPolicyRuleCondition`

GetRisk returns the Risk field if non-nil, zero value otherwise.

### GetRiskOk

`func (o *AccessPolicyRuleConditions) GetRiskOk() (*RiskPolicyRuleCondition, bool)`

GetRiskOk returns a tuple with the Risk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRisk

`func (o *AccessPolicyRuleConditions) SetRisk(v RiskPolicyRuleCondition)`

SetRisk sets Risk field to given value.

### HasRisk

`func (o *AccessPolicyRuleConditions) HasRisk() bool`

HasRisk returns a boolean if a field has been set.

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

### GetScopes

`func (o *AccessPolicyRuleConditions) GetScopes() OAuth2ScopesMediationPolicyRuleCondition`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *AccessPolicyRuleConditions) GetScopesOk() (*OAuth2ScopesMediationPolicyRuleCondition, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *AccessPolicyRuleConditions) SetScopes(v OAuth2ScopesMediationPolicyRuleCondition)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *AccessPolicyRuleConditions) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetUserIdentifier

`func (o *AccessPolicyRuleConditions) GetUserIdentifier() UserIdentifierPolicyRuleCondition`

GetUserIdentifier returns the UserIdentifier field if non-nil, zero value otherwise.

### GetUserIdentifierOk

`func (o *AccessPolicyRuleConditions) GetUserIdentifierOk() (*UserIdentifierPolicyRuleCondition, bool)`

GetUserIdentifierOk returns a tuple with the UserIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdentifier

`func (o *AccessPolicyRuleConditions) SetUserIdentifier(v UserIdentifierPolicyRuleCondition)`

SetUserIdentifier sets UserIdentifier field to given value.

### HasUserIdentifier

`func (o *AccessPolicyRuleConditions) HasUserIdentifier() bool`

HasUserIdentifier returns a boolean if a field has been set.

### GetUsers

`func (o *AccessPolicyRuleConditions) GetUsers() UserPolicyRuleCondition`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *AccessPolicyRuleConditions) GetUsersOk() (*UserPolicyRuleCondition, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *AccessPolicyRuleConditions) SetUsers(v UserPolicyRuleCondition)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *AccessPolicyRuleConditions) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetUserStatus

`func (o *AccessPolicyRuleConditions) GetUserStatus() UserStatusPolicyRuleCondition`

GetUserStatus returns the UserStatus field if non-nil, zero value otherwise.

### GetUserStatusOk

`func (o *AccessPolicyRuleConditions) GetUserStatusOk() (*UserStatusPolicyRuleCondition, bool)`

GetUserStatusOk returns a tuple with the UserStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserStatus

`func (o *AccessPolicyRuleConditions) SetUserStatus(v UserStatusPolicyRuleCondition)`

SetUserStatus sets UserStatus field to given value.

### HasUserStatus

`func (o *AccessPolicyRuleConditions) HasUserStatus() bool`

HasUserStatus returns a boolean if a field has been set.

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


