/*
Okta Admin Management

Allows customers to easily access the Okta Management APIs

Copyright 2018 - Present Okta, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

API version: 2024.06.1
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// AccessPolicyRuleConditions struct for AccessPolicyRuleConditions
type AccessPolicyRuleConditions struct {
	App *AppAndInstancePolicyRuleCondition `json:"app,omitempty"`
	Apps *AppInstancePolicyRuleCondition `json:"apps,omitempty"`
	AuthContext *PolicyRuleAuthContextCondition `json:"authContext,omitempty"`
	AuthProvider *PasswordPolicyAuthenticationProviderCondition `json:"authProvider,omitempty"`
	BeforeScheduledAction *BeforeScheduledActionPolicyRuleCondition `json:"beforeScheduledAction,omitempty"`
	Clients *ClientPolicyCondition `json:"clients,omitempty"`
	Context *ContextPolicyRuleCondition `json:"context,omitempty"`
	Device *DeviceAccessPolicyRuleCondition `json:"device,omitempty"`
	GrantTypes *GrantTypePolicyRuleCondition `json:"grantTypes,omitempty"`
	Groups *GroupPolicyRuleCondition `json:"groups,omitempty"`
	IdentityProvider *IdentityProviderPolicyRuleCondition `json:"identityProvider,omitempty"`
	MdmEnrollment *MDMEnrollmentPolicyRuleCondition `json:"mdmEnrollment,omitempty"`
	Network *PolicyNetworkCondition `json:"network,omitempty"`
	People *PolicyPeopleCondition `json:"people,omitempty"`
	Platform *PlatformPolicyRuleCondition `json:"platform,omitempty"`
	Risk *RiskPolicyRuleCondition `json:"risk,omitempty"`
	RiskScore *RiskScorePolicyRuleCondition `json:"riskScore,omitempty"`
	Scopes *OAuth2ScopesMediationPolicyRuleCondition `json:"scopes,omitempty"`
	UserIdentifier *UserIdentifierPolicyRuleCondition `json:"userIdentifier,omitempty"`
	Users *UserPolicyRuleCondition `json:"users,omitempty"`
	UserStatus *UserStatusPolicyRuleCondition `json:"userStatus,omitempty"`
	ElCondition *AccessPolicyRuleCustomCondition `json:"elCondition,omitempty"`
	UserType *UserTypeCondition `json:"userType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccessPolicyRuleConditions AccessPolicyRuleConditions

// NewAccessPolicyRuleConditions instantiates a new AccessPolicyRuleConditions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessPolicyRuleConditions() *AccessPolicyRuleConditions {
	this := AccessPolicyRuleConditions{}
	return &this
}

// NewAccessPolicyRuleConditionsWithDefaults instantiates a new AccessPolicyRuleConditions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessPolicyRuleConditionsWithDefaults() *AccessPolicyRuleConditions {
	this := AccessPolicyRuleConditions{}
	return &this
}

// GetApp returns the App field value if set, zero value otherwise.
func (o *AccessPolicyRuleConditions) GetApp() AppAndInstancePolicyRuleCondition {
	if o == nil || o.App == nil {
		var ret AppAndInstancePolicyRuleCondition
		return ret
	}
	return *o.App
}

// GetAppOk returns a tuple with the App field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicyRuleConditions) GetAppOk() (*AppAndInstancePolicyRuleCondition, bool) {
	if o == nil || o.App == nil {
		return nil, false
	}
	return o.App, true
}

// HasApp returns a boolean if a field has been set.
func (o *AccessPolicyRuleConditions) HasApp() bool {
	if o != nil && o.App != nil {
		return true
	}

	return false
}

// SetApp gets a reference to the given AppAndInstancePolicyRuleCondition and assigns it to the App field.
func (o *AccessPolicyRuleConditions) SetApp(v AppAndInstancePolicyRuleCondition) {
	o.App = &v
}

// GetApps returns the Apps field value if set, zero value otherwise.
func (o *AccessPolicyRuleConditions) GetApps() AppInstancePolicyRuleCondition {
	if o == nil || o.Apps == nil {
		var ret AppInstancePolicyRuleCondition
		return ret
	}
	return *o.Apps
}

// GetAppsOk returns a tuple with the Apps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicyRuleConditions) GetAppsOk() (*AppInstancePolicyRuleCondition, bool) {
	if o == nil || o.Apps == nil {
		return nil, false
	}
	return o.Apps, true
}

// HasApps returns a boolean if a field has been set.
func (o *AccessPolicyRuleConditions) HasApps() bool {
	if o != nil && o.Apps != nil {
		return true
	}

	return false
}

// SetApps gets a reference to the given AppInstancePolicyRuleCondition and assigns it to the Apps field.
func (o *AccessPolicyRuleConditions) SetApps(v AppInstancePolicyRuleCondition) {
	o.Apps = &v
}

// GetAuthContext returns the AuthContext field value if set, zero value otherwise.
func (o *AccessPolicyRuleConditions) GetAuthContext() PolicyRuleAuthContextCondition {
	if o == nil || o.AuthContext == nil {
		var ret PolicyRuleAuthContextCondition
		return ret
	}
	return *o.AuthContext
}

// GetAuthContextOk returns a tuple with the AuthContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicyRuleConditions) GetAuthContextOk() (*PolicyRuleAuthContextCondition, bool) {
	if o == nil || o.AuthContext == nil {
		return nil, false
	}
	return o.AuthContext, true
}

// HasAuthContext returns a boolean if a field has been set.
func (o *AccessPolicyRuleConditions) HasAuthContext() bool {
	if o != nil && o.AuthContext != nil {
		return true
	}

	return false
}

// SetAuthContext gets a reference to the given PolicyRuleAuthContextCondition and assigns it to the AuthContext field.
func (o *AccessPolicyRuleConditions) SetAuthContext(v PolicyRuleAuthContextCondition) {
	o.AuthContext = &v
}

// GetAuthProvider returns the AuthProvider field value if set, zero value otherwise.
func (o *AccessPolicyRuleConditions) GetAuthProvider() PasswordPolicyAuthenticationProviderCondition {
	if o == nil || o.AuthProvider == nil {
		var ret PasswordPolicyAuthenticationProviderCondition
		return ret
	}
	return *o.AuthProvider
}

// GetAuthProviderOk returns a tuple with the AuthProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicyRuleConditions) GetAuthProviderOk() (*PasswordPolicyAuthenticationProviderCondition, bool) {
	if o == nil || o.AuthProvider == nil {
		return nil, false
	}
	return o.AuthProvider, true
}

// HasAuthProvider returns a boolean if a field has been set.
func (o *AccessPolicyRuleConditions) HasAuthProvider() bool {
	if o != nil && o.AuthProvider != nil {
		return true
	}

	return false
}

// SetAuthProvider gets a reference to the given PasswordPolicyAuthenticationProviderCondition and assigns it to the AuthProvider field.
func (o *AccessPolicyRuleConditions) SetAuthProvider(v PasswordPolicyAuthenticationProviderCondition) {
	o.AuthProvider = &v
}

// GetBeforeScheduledAction returns the BeforeScheduledAction field value if set, zero value otherwise.
func (o *AccessPolicyRuleConditions) GetBeforeScheduledAction() BeforeScheduledActionPolicyRuleCondition {
	if o == nil || o.BeforeScheduledAction == nil {
		var ret BeforeScheduledActionPolicyRuleCondition
		return ret
	}
	return *o.BeforeScheduledAction
}

// GetBeforeScheduledActionOk returns a tuple with the BeforeScheduledAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicyRuleConditions) GetBeforeScheduledActionOk() (*BeforeScheduledActionPolicyRuleCondition, bool) {
	if o == nil || o.BeforeScheduledAction == nil {
		return nil, false
	}
	return o.BeforeScheduledAction, true
}

// HasBeforeScheduledAction returns a boolean if a field has been set.
func (o *AccessPolicyRuleConditions) HasBeforeScheduledAction() bool {
	if o != nil && o.BeforeScheduledAction != nil {
		return true
	}

	return false
}

// SetBeforeScheduledAction gets a reference to the given BeforeScheduledActionPolicyRuleCondition and assigns it to the BeforeScheduledAction field.
func (o *AccessPolicyRuleConditions) SetBeforeScheduledAction(v BeforeScheduledActionPolicyRuleCondition) {
	o.BeforeScheduledAction = &v
}

// GetClients returns the Clients field value if set, zero value otherwise.
func (o *AccessPolicyRuleConditions) GetClients() ClientPolicyCondition {
	if o == nil || o.Clients == nil {
		var ret ClientPolicyCondition
		return ret
	}
	return *o.Clients
}

// GetClientsOk returns a tuple with the Clients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicyRuleConditions) GetClientsOk() (*ClientPolicyCondition, bool) {
	if o == nil || o.Clients == nil {
		return nil, false
	}
	return o.Clients, true
}

// HasClients returns a boolean if a field has been set.
func (o *AccessPolicyRuleConditions) HasClients() bool {
	if o != nil && o.Clients != nil {
		return true
	}

	return false
}

// SetClients gets a reference to the given ClientPolicyCondition and assigns it to the Clients field.
func (o *AccessPolicyRuleConditions) SetClients(v ClientPolicyCondition) {
	o.Clients = &v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *AccessPolicyRuleConditions) GetContext() ContextPolicyRuleCondition {
	if o == nil || o.Context == nil {
		var ret ContextPolicyRuleCondition
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicyRuleConditions) GetContextOk() (*ContextPolicyRuleCondition, bool) {
	if o == nil || o.Context == nil {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *AccessPolicyRuleConditions) HasContext() bool {
	if o != nil && o.Context != nil {
		return true
	}

	return false
}

// SetContext gets a reference to the given ContextPolicyRuleCondition and assigns it to the Context field.
func (o *AccessPolicyRuleConditions) SetContext(v ContextPolicyRuleCondition) {
	o.Context = &v
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *AccessPolicyRuleConditions) GetDevice() DeviceAccessPolicyRuleCondition {
	if o == nil || o.Device == nil {
		var ret DeviceAccessPolicyRuleCondition
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicyRuleConditions) GetDeviceOk() (*DeviceAccessPolicyRuleCondition, bool) {
	if o == nil || o.Device == nil {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *AccessPolicyRuleConditions) HasDevice() bool {
	if o != nil && o.Device != nil {
		return true
	}

	return false
}

// SetDevice gets a reference to the given DeviceAccessPolicyRuleCondition and assigns it to the Device field.
func (o *AccessPolicyRuleConditions) SetDevice(v DeviceAccessPolicyRuleCondition) {
	o.Device = &v
}

// GetGrantTypes returns the GrantTypes field value if set, zero value otherwise.
func (o *AccessPolicyRuleConditions) GetGrantTypes() GrantTypePolicyRuleCondition {
	if o == nil || o.GrantTypes == nil {
		var ret GrantTypePolicyRuleCondition
		return ret
	}
	return *o.GrantTypes
}

// GetGrantTypesOk returns a tuple with the GrantTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicyRuleConditions) GetGrantTypesOk() (*GrantTypePolicyRuleCondition, bool) {
	if o == nil || o.GrantTypes == nil {
		return nil, false
	}
	return o.GrantTypes, true
}

// HasGrantTypes returns a boolean if a field has been set.
func (o *AccessPolicyRuleConditions) HasGrantTypes() bool {
	if o != nil && o.GrantTypes != nil {
		return true
	}

	return false
}

// SetGrantTypes gets a reference to the given GrantTypePolicyRuleCondition and assigns it to the GrantTypes field.
func (o *AccessPolicyRuleConditions) SetGrantTypes(v GrantTypePolicyRuleCondition) {
	o.GrantTypes = &v
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *AccessPolicyRuleConditions) GetGroups() GroupPolicyRuleCondition {
	if o == nil || o.Groups == nil {
		var ret GroupPolicyRuleCondition
		return ret
	}
	return *o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicyRuleConditions) GetGroupsOk() (*GroupPolicyRuleCondition, bool) {
	if o == nil || o.Groups == nil {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *AccessPolicyRuleConditions) HasGroups() bool {
	if o != nil && o.Groups != nil {
		return true
	}

	return false
}

// SetGroups gets a reference to the given GroupPolicyRuleCondition and assigns it to the Groups field.
func (o *AccessPolicyRuleConditions) SetGroups(v GroupPolicyRuleCondition) {
	o.Groups = &v
}

// GetIdentityProvider returns the IdentityProvider field value if set, zero value otherwise.
func (o *AccessPolicyRuleConditions) GetIdentityProvider() IdentityProviderPolicyRuleCondition {
	if o == nil || o.IdentityProvider == nil {
		var ret IdentityProviderPolicyRuleCondition
		return ret
	}
	return *o.IdentityProvider
}

// GetIdentityProviderOk returns a tuple with the IdentityProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicyRuleConditions) GetIdentityProviderOk() (*IdentityProviderPolicyRuleCondition, bool) {
	if o == nil || o.IdentityProvider == nil {
		return nil, false
	}
	return o.IdentityProvider, true
}

// HasIdentityProvider returns a boolean if a field has been set.
func (o *AccessPolicyRuleConditions) HasIdentityProvider() bool {
	if o != nil && o.IdentityProvider != nil {
		return true
	}

	return false
}

// SetIdentityProvider gets a reference to the given IdentityProviderPolicyRuleCondition and assigns it to the IdentityProvider field.
func (o *AccessPolicyRuleConditions) SetIdentityProvider(v IdentityProviderPolicyRuleCondition) {
	o.IdentityProvider = &v
}

// GetMdmEnrollment returns the MdmEnrollment field value if set, zero value otherwise.
func (o *AccessPolicyRuleConditions) GetMdmEnrollment() MDMEnrollmentPolicyRuleCondition {
	if o == nil || o.MdmEnrollment == nil {
		var ret MDMEnrollmentPolicyRuleCondition
		return ret
	}
	return *o.MdmEnrollment
}

// GetMdmEnrollmentOk returns a tuple with the MdmEnrollment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicyRuleConditions) GetMdmEnrollmentOk() (*MDMEnrollmentPolicyRuleCondition, bool) {
	if o == nil || o.MdmEnrollment == nil {
		return nil, false
	}
	return o.MdmEnrollment, true
}

// HasMdmEnrollment returns a boolean if a field has been set.
func (o *AccessPolicyRuleConditions) HasMdmEnrollment() bool {
	if o != nil && o.MdmEnrollment != nil {
		return true
	}

	return false
}

// SetMdmEnrollment gets a reference to the given MDMEnrollmentPolicyRuleCondition and assigns it to the MdmEnrollment field.
func (o *AccessPolicyRuleConditions) SetMdmEnrollment(v MDMEnrollmentPolicyRuleCondition) {
	o.MdmEnrollment = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *AccessPolicyRuleConditions) GetNetwork() PolicyNetworkCondition {
	if o == nil || o.Network == nil {
		var ret PolicyNetworkCondition
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicyRuleConditions) GetNetworkOk() (*PolicyNetworkCondition, bool) {
	if o == nil || o.Network == nil {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *AccessPolicyRuleConditions) HasNetwork() bool {
	if o != nil && o.Network != nil {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given PolicyNetworkCondition and assigns it to the Network field.
func (o *AccessPolicyRuleConditions) SetNetwork(v PolicyNetworkCondition) {
	o.Network = &v
}

// GetPeople returns the People field value if set, zero value otherwise.
func (o *AccessPolicyRuleConditions) GetPeople() PolicyPeopleCondition {
	if o == nil || o.People == nil {
		var ret PolicyPeopleCondition
		return ret
	}
	return *o.People
}

// GetPeopleOk returns a tuple with the People field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicyRuleConditions) GetPeopleOk() (*PolicyPeopleCondition, bool) {
	if o == nil || o.People == nil {
		return nil, false
	}
	return o.People, true
}

// HasPeople returns a boolean if a field has been set.
func (o *AccessPolicyRuleConditions) HasPeople() bool {
	if o != nil && o.People != nil {
		return true
	}

	return false
}

// SetPeople gets a reference to the given PolicyPeopleCondition and assigns it to the People field.
func (o *AccessPolicyRuleConditions) SetPeople(v PolicyPeopleCondition) {
	o.People = &v
}

// GetPlatform returns the Platform field value if set, zero value otherwise.
func (o *AccessPolicyRuleConditions) GetPlatform() PlatformPolicyRuleCondition {
	if o == nil || o.Platform == nil {
		var ret PlatformPolicyRuleCondition
		return ret
	}
	return *o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicyRuleConditions) GetPlatformOk() (*PlatformPolicyRuleCondition, bool) {
	if o == nil || o.Platform == nil {
		return nil, false
	}
	return o.Platform, true
}

// HasPlatform returns a boolean if a field has been set.
func (o *AccessPolicyRuleConditions) HasPlatform() bool {
	if o != nil && o.Platform != nil {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given PlatformPolicyRuleCondition and assigns it to the Platform field.
func (o *AccessPolicyRuleConditions) SetPlatform(v PlatformPolicyRuleCondition) {
	o.Platform = &v
}

// GetRisk returns the Risk field value if set, zero value otherwise.
func (o *AccessPolicyRuleConditions) GetRisk() RiskPolicyRuleCondition {
	if o == nil || o.Risk == nil {
		var ret RiskPolicyRuleCondition
		return ret
	}
	return *o.Risk
}

// GetRiskOk returns a tuple with the Risk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicyRuleConditions) GetRiskOk() (*RiskPolicyRuleCondition, bool) {
	if o == nil || o.Risk == nil {
		return nil, false
	}
	return o.Risk, true
}

// HasRisk returns a boolean if a field has been set.
func (o *AccessPolicyRuleConditions) HasRisk() bool {
	if o != nil && o.Risk != nil {
		return true
	}

	return false
}

// SetRisk gets a reference to the given RiskPolicyRuleCondition and assigns it to the Risk field.
func (o *AccessPolicyRuleConditions) SetRisk(v RiskPolicyRuleCondition) {
	o.Risk = &v
}

// GetRiskScore returns the RiskScore field value if set, zero value otherwise.
func (o *AccessPolicyRuleConditions) GetRiskScore() RiskScorePolicyRuleCondition {
	if o == nil || o.RiskScore == nil {
		var ret RiskScorePolicyRuleCondition
		return ret
	}
	return *o.RiskScore
}

// GetRiskScoreOk returns a tuple with the RiskScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicyRuleConditions) GetRiskScoreOk() (*RiskScorePolicyRuleCondition, bool) {
	if o == nil || o.RiskScore == nil {
		return nil, false
	}
	return o.RiskScore, true
}

// HasRiskScore returns a boolean if a field has been set.
func (o *AccessPolicyRuleConditions) HasRiskScore() bool {
	if o != nil && o.RiskScore != nil {
		return true
	}

	return false
}

// SetRiskScore gets a reference to the given RiskScorePolicyRuleCondition and assigns it to the RiskScore field.
func (o *AccessPolicyRuleConditions) SetRiskScore(v RiskScorePolicyRuleCondition) {
	o.RiskScore = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *AccessPolicyRuleConditions) GetScopes() OAuth2ScopesMediationPolicyRuleCondition {
	if o == nil || o.Scopes == nil {
		var ret OAuth2ScopesMediationPolicyRuleCondition
		return ret
	}
	return *o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicyRuleConditions) GetScopesOk() (*OAuth2ScopesMediationPolicyRuleCondition, bool) {
	if o == nil || o.Scopes == nil {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *AccessPolicyRuleConditions) HasScopes() bool {
	if o != nil && o.Scopes != nil {
		return true
	}

	return false
}

// SetScopes gets a reference to the given OAuth2ScopesMediationPolicyRuleCondition and assigns it to the Scopes field.
func (o *AccessPolicyRuleConditions) SetScopes(v OAuth2ScopesMediationPolicyRuleCondition) {
	o.Scopes = &v
}

// GetUserIdentifier returns the UserIdentifier field value if set, zero value otherwise.
func (o *AccessPolicyRuleConditions) GetUserIdentifier() UserIdentifierPolicyRuleCondition {
	if o == nil || o.UserIdentifier == nil {
		var ret UserIdentifierPolicyRuleCondition
		return ret
	}
	return *o.UserIdentifier
}

// GetUserIdentifierOk returns a tuple with the UserIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicyRuleConditions) GetUserIdentifierOk() (*UserIdentifierPolicyRuleCondition, bool) {
	if o == nil || o.UserIdentifier == nil {
		return nil, false
	}
	return o.UserIdentifier, true
}

// HasUserIdentifier returns a boolean if a field has been set.
func (o *AccessPolicyRuleConditions) HasUserIdentifier() bool {
	if o != nil && o.UserIdentifier != nil {
		return true
	}

	return false
}

// SetUserIdentifier gets a reference to the given UserIdentifierPolicyRuleCondition and assigns it to the UserIdentifier field.
func (o *AccessPolicyRuleConditions) SetUserIdentifier(v UserIdentifierPolicyRuleCondition) {
	o.UserIdentifier = &v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *AccessPolicyRuleConditions) GetUsers() UserPolicyRuleCondition {
	if o == nil || o.Users == nil {
		var ret UserPolicyRuleCondition
		return ret
	}
	return *o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicyRuleConditions) GetUsersOk() (*UserPolicyRuleCondition, bool) {
	if o == nil || o.Users == nil {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *AccessPolicyRuleConditions) HasUsers() bool {
	if o != nil && o.Users != nil {
		return true
	}

	return false
}

// SetUsers gets a reference to the given UserPolicyRuleCondition and assigns it to the Users field.
func (o *AccessPolicyRuleConditions) SetUsers(v UserPolicyRuleCondition) {
	o.Users = &v
}

// GetUserStatus returns the UserStatus field value if set, zero value otherwise.
func (o *AccessPolicyRuleConditions) GetUserStatus() UserStatusPolicyRuleCondition {
	if o == nil || o.UserStatus == nil {
		var ret UserStatusPolicyRuleCondition
		return ret
	}
	return *o.UserStatus
}

// GetUserStatusOk returns a tuple with the UserStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicyRuleConditions) GetUserStatusOk() (*UserStatusPolicyRuleCondition, bool) {
	if o == nil || o.UserStatus == nil {
		return nil, false
	}
	return o.UserStatus, true
}

// HasUserStatus returns a boolean if a field has been set.
func (o *AccessPolicyRuleConditions) HasUserStatus() bool {
	if o != nil && o.UserStatus != nil {
		return true
	}

	return false
}

// SetUserStatus gets a reference to the given UserStatusPolicyRuleCondition and assigns it to the UserStatus field.
func (o *AccessPolicyRuleConditions) SetUserStatus(v UserStatusPolicyRuleCondition) {
	o.UserStatus = &v
}

// GetElCondition returns the ElCondition field value if set, zero value otherwise.
func (o *AccessPolicyRuleConditions) GetElCondition() AccessPolicyRuleCustomCondition {
	if o == nil || o.ElCondition == nil {
		var ret AccessPolicyRuleCustomCondition
		return ret
	}
	return *o.ElCondition
}

// GetElConditionOk returns a tuple with the ElCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicyRuleConditions) GetElConditionOk() (*AccessPolicyRuleCustomCondition, bool) {
	if o == nil || o.ElCondition == nil {
		return nil, false
	}
	return o.ElCondition, true
}

// HasElCondition returns a boolean if a field has been set.
func (o *AccessPolicyRuleConditions) HasElCondition() bool {
	if o != nil && o.ElCondition != nil {
		return true
	}

	return false
}

// SetElCondition gets a reference to the given AccessPolicyRuleCustomCondition and assigns it to the ElCondition field.
func (o *AccessPolicyRuleConditions) SetElCondition(v AccessPolicyRuleCustomCondition) {
	o.ElCondition = &v
}

// GetUserType returns the UserType field value if set, zero value otherwise.
func (o *AccessPolicyRuleConditions) GetUserType() UserTypeCondition {
	if o == nil || o.UserType == nil {
		var ret UserTypeCondition
		return ret
	}
	return *o.UserType
}

// GetUserTypeOk returns a tuple with the UserType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicyRuleConditions) GetUserTypeOk() (*UserTypeCondition, bool) {
	if o == nil || o.UserType == nil {
		return nil, false
	}
	return o.UserType, true
}

// HasUserType returns a boolean if a field has been set.
func (o *AccessPolicyRuleConditions) HasUserType() bool {
	if o != nil && o.UserType != nil {
		return true
	}

	return false
}

// SetUserType gets a reference to the given UserTypeCondition and assigns it to the UserType field.
func (o *AccessPolicyRuleConditions) SetUserType(v UserTypeCondition) {
	o.UserType = &v
}

func (o AccessPolicyRuleConditions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.App != nil {
		toSerialize["app"] = o.App
	}
	if o.Apps != nil {
		toSerialize["apps"] = o.Apps
	}
	if o.AuthContext != nil {
		toSerialize["authContext"] = o.AuthContext
	}
	if o.AuthProvider != nil {
		toSerialize["authProvider"] = o.AuthProvider
	}
	if o.BeforeScheduledAction != nil {
		toSerialize["beforeScheduledAction"] = o.BeforeScheduledAction
	}
	if o.Clients != nil {
		toSerialize["clients"] = o.Clients
	}
	if o.Context != nil {
		toSerialize["context"] = o.Context
	}
	if o.Device != nil {
		toSerialize["device"] = o.Device
	}
	if o.GrantTypes != nil {
		toSerialize["grantTypes"] = o.GrantTypes
	}
	if o.Groups != nil {
		toSerialize["groups"] = o.Groups
	}
	if o.IdentityProvider != nil {
		toSerialize["identityProvider"] = o.IdentityProvider
	}
	if o.MdmEnrollment != nil {
		toSerialize["mdmEnrollment"] = o.MdmEnrollment
	}
	if o.Network != nil {
		toSerialize["network"] = o.Network
	}
	if o.People != nil {
		toSerialize["people"] = o.People
	}
	if o.Platform != nil {
		toSerialize["platform"] = o.Platform
	}
	if o.Risk != nil {
		toSerialize["risk"] = o.Risk
	}
	if o.RiskScore != nil {
		toSerialize["riskScore"] = o.RiskScore
	}
	if o.Scopes != nil {
		toSerialize["scopes"] = o.Scopes
	}
	if o.UserIdentifier != nil {
		toSerialize["userIdentifier"] = o.UserIdentifier
	}
	if o.Users != nil {
		toSerialize["users"] = o.Users
	}
	if o.UserStatus != nil {
		toSerialize["userStatus"] = o.UserStatus
	}
	if o.ElCondition != nil {
		toSerialize["elCondition"] = o.ElCondition
	}
	if o.UserType != nil {
		toSerialize["userType"] = o.UserType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AccessPolicyRuleConditions) UnmarshalJSON(bytes []byte) (err error) {
	varAccessPolicyRuleConditions := _AccessPolicyRuleConditions{}

	err = json.Unmarshal(bytes, &varAccessPolicyRuleConditions)
	if err == nil {
		*o = AccessPolicyRuleConditions(varAccessPolicyRuleConditions)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "app")
		delete(additionalProperties, "apps")
		delete(additionalProperties, "authContext")
		delete(additionalProperties, "authProvider")
		delete(additionalProperties, "beforeScheduledAction")
		delete(additionalProperties, "clients")
		delete(additionalProperties, "context")
		delete(additionalProperties, "device")
		delete(additionalProperties, "grantTypes")
		delete(additionalProperties, "groups")
		delete(additionalProperties, "identityProvider")
		delete(additionalProperties, "mdmEnrollment")
		delete(additionalProperties, "network")
		delete(additionalProperties, "people")
		delete(additionalProperties, "platform")
		delete(additionalProperties, "risk")
		delete(additionalProperties, "riskScore")
		delete(additionalProperties, "scopes")
		delete(additionalProperties, "userIdentifier")
		delete(additionalProperties, "users")
		delete(additionalProperties, "userStatus")
		delete(additionalProperties, "elCondition")
		delete(additionalProperties, "userType")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableAccessPolicyRuleConditions struct {
	value *AccessPolicyRuleConditions
	isSet bool
}

func (v NullableAccessPolicyRuleConditions) Get() *AccessPolicyRuleConditions {
	return v.value
}

func (v *NullableAccessPolicyRuleConditions) Set(val *AccessPolicyRuleConditions) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessPolicyRuleConditions) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessPolicyRuleConditions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessPolicyRuleConditions(val *AccessPolicyRuleConditions) *NullableAccessPolicyRuleConditions {
	return &NullableAccessPolicyRuleConditions{value: val, isSet: true}
}

func (v NullableAccessPolicyRuleConditions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessPolicyRuleConditions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

