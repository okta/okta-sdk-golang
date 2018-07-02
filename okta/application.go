/*
* Copyright 2018 - Present Okta, Inc.
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*      http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
*/

// AUTO-GENERATED!  DO NOT EDIT FILE DIRECTLY

package okta

import (
	"bytes"
	"fmt"
	"encoding/json"
	"time"
)

type AppUser struct {
	Embedded map[string]interface{} `json:"_embedded, omitempty"`
	Links map[string]interface{} `json:"_links, omitempty"`
	Created time.Time `json:"created, omitempty"`
	Credentials AppUserCredentials `json:"credentials, omitempty"`
	ExternalId string `json:"externalId, omitempty"`
	Id string `json:"id, omitempty"`
	LastSync time.Time `json:"lastSync, omitempty"`
	LastUpdated time.Time `json:"lastUpdated, omitempty"`
	PasswordChanged time.Time `json:"passwordChanged, omitempty"`
	Profile map[string]interface{} `json:"profile, omitempty"`
	Scope string `json:"scope, omitempty"`
	Status string `json:"status, omitempty"`
	StatusChanged time.Time `json:"statusChanged, omitempty"`
	SyncState string `json:"syncState, omitempty"`
}

func (m *AppUser) WithCredentials(v AppUserCredentials) *AppUser {
	m.Credentials = v
	return m
}
func (m *AppUser) WithId(v string) *AppUser {
	m.Id = v
	return m
}
func (m *AppUser) WithProfile(v map[string]interface{}) *AppUser {
	m.Profile = v
	return m
}
func (m *AppUser) WithScope(v string) *AppUser {
	m.Scope = v
	return m
}


type AppUserCredentials struct {
	Password AppUserPasswordCredential `json:"password, omitempty"`
	UserName string `json:"userName, omitempty"`
}

func (m *AppUserCredentials) WithPassword(v AppUserPasswordCredential) *AppUserCredentials {
	m.Password = v
	return m
}
func (m *AppUserCredentials) WithUserName(v string) *AppUserCredentials {
	m.UserName = v
	return m
}


type AppUserPasswordCredential struct {
	Value string `json:"value, omitempty"`
}

func (m *AppUserPasswordCredential) WithValue(v string) *AppUserPasswordCredential {
	m.Value = v
	return m
}

type ApplicationResource resource

type Application struct {
	Embedded map[string]interface{} `json:"_embedded, omitempty"`
	Links map[string]interface{} `json:"_links, omitempty"`
	Accessibility ApplicationAccessibility `json:"accessibility, omitempty"`
	Created time.Time `json:"created, omitempty"`
	Credentials ApplicationCredentials `json:"credentials, omitempty"`
	Features map[string]interface{} `json:"features, omitempty"`
	Id string `json:"id, omitempty"`
	Label string `json:"label, omitempty"`
	LastUpdated time.Time `json:"lastUpdated, omitempty"`
	Licensing ApplicationLicensing `json:"licensing, omitempty"`
	Name string `json:"name, omitempty"`
	Settings ApplicationSettings `json:"settings, omitempty"`
	SignOnMode string `json:"signOnMode, omitempty"`
	Status string `json:"status, omitempty"`
	Visibility ApplicationVisibility `json:"visibility, omitempty"`
}

func (m *Application) WithAccessibility(v ApplicationAccessibility) *Application {
	m.Accessibility = v
	return m
}
func (m *Application) WithCredentials(v ApplicationCredentials) *Application {
	m.Credentials = v
	return m
}
func (m *Application) WithFeatures(v map[string]interface{}) *Application {
	m.Features = v
	return m
}
func (m *Application) WithLabel(v string) *Application {
	m.Label = v
	return m
}
func (m *Application) WithLicensing(v ApplicationLicensing) *Application {
	m.Licensing = v
	return m
}
func (m *Application) WithSettings(v ApplicationSettings) *Application {
	m.Settings = v
	return m
}
func (m *Application) WithSignOnMode(v string) *Application {
	m.SignOnMode = v
	return m
}
func (m *Application) WithVisibility(v ApplicationVisibility) *Application {
	m.Visibility = v
	return m
}

func (m *ApplicationResource) Activate(appId string)  {
	_, err := m.client.requestExecutor.Post("/api/v1/apps/"+appId+"/lifecycle/activate", nil)
		if err != nil  {
		fmt.Printf("error: %v", err)
	}
}

func (m *ApplicationResource) Deactivate(appId string)  {
	_, err := m.client.requestExecutor.Post("/api/v1/apps/"+appId+"/lifecycle/deactivate", nil)
		if err != nil  {
		fmt.Printf("error: %v", err)
	}
}

func (m *ApplicationResource) Listapplicationusers(appId string) AppUser {
	resp, err := m.client.requestExecutor.Get("/api/v1/apps/"+appId+"/users")
	if err != nil  {
		fmt.Printf("error: %v", err)
	}

	r := AppUser{}

	json.Unmarshal(resp, &r)

	return r
}

func (m *ApplicationResource) Assignusertoapplication(appId string, body AppUser) AppUser {
	iobytes, err := json.Marshal(body)
	if err != nil  {
		fmt.Printf("error: %v", err)
	}
	resp, err := m.client.requestExecutor.Post("/api/v1/apps/"+appId+"/users", bytes.NewReader(iobytes))
		if err != nil  {
		fmt.Printf("error: %v", err)
	}

	r := AppUser{}

	json.Unmarshal(resp, &r)

	return r
}

func (m *ApplicationResource) Getapplicationuser(appId string, userId string) AppUser {
	resp, err := m.client.requestExecutor.Get("/api/v1/apps/"+appId+"/users/"+userId+"")
	if err != nil  {
		fmt.Printf("error: %v", err)
	}

	r := AppUser{}

	json.Unmarshal(resp, &r)

	return r
}

func (m *ApplicationResource) Createapplicationgroupassignment(appId string, groupId string, body ApplicationGroupAssignment) ApplicationGroupAssignment {
	iobytes, err := json.Marshal(body)
	if err != nil  {
		fmt.Printf("error: %v", err)
	}
	resp, err := m.client.requestExecutor.Put("/api/v1/apps/"+appId+"/groups/"+groupId+"", bytes.NewReader(iobytes))
		if err != nil  {
		fmt.Printf("error: %v", err)
	}

	r := ApplicationGroupAssignment{}

	json.Unmarshal(resp, &r)

	return r
}

func (m *ApplicationResource) Getapplicationgroupassignment(appId string, groupId string) ApplicationGroupAssignment {
	resp, err := m.client.requestExecutor.Get("/api/v1/apps/"+appId+"/groups/"+groupId+"")
	if err != nil  {
		fmt.Printf("error: %v", err)
	}

	r := ApplicationGroupAssignment{}

	json.Unmarshal(resp, &r)

	return r
}

func (m *ApplicationResource) Cloneapplicationkey(appId string, keyId string) JsonWebKey {
	resp, err := m.client.requestExecutor.Post("/api/v1/apps/"+appId+"/credentials/keys/"+keyId+"/clone", nil)
		if err != nil  {
		fmt.Printf("error: %v", err)
	}

	r := JsonWebKey{}

	json.Unmarshal(resp, &r)

	return r
}

func (m *ApplicationResource) Getapplicationkey(appId string, keyId string) JsonWebKey {
	resp, err := m.client.requestExecutor.Get("/api/v1/apps/"+appId+"/credentials/keys/"+keyId+"")
	if err != nil  {
		fmt.Printf("error: %v", err)
	}

	r := JsonWebKey{}

	json.Unmarshal(resp, &r)

	return r
}

func (m *ApplicationResource) Listgroupassignments(appId string) ApplicationGroupAssignment {
	resp, err := m.client.requestExecutor.Get("/api/v1/apps/"+appId+"/groups")
	if err != nil  {
		fmt.Printf("error: %v", err)
	}

	r := ApplicationGroupAssignment{}

	json.Unmarshal(resp, &r)

	return r
}

func (m *ApplicationResource) Listkeys(appId string) JsonWebKey {
	resp, err := m.client.requestExecutor.Get("/api/v1/apps/"+appId+"/credentials/keys")
	if err != nil  {
		fmt.Printf("error: %v", err)
	}

	r := JsonWebKey{}

	json.Unmarshal(resp, &r)

	return r
}


type ApplicationAccessibility struct {
	ErrorRedirectUrl string `json:"errorRedirectUrl, omitempty"`
	LoginRedirectUrl string `json:"loginRedirectUrl, omitempty"`
	SelfService bool `json:"selfService, omitempty"`
}

func (m *ApplicationAccessibility) WithErrorRedirectUrl(v string) *ApplicationAccessibility {
	m.ErrorRedirectUrl = v
	return m
}
func (m *ApplicationAccessibility) WithLoginRedirectUrl(v string) *ApplicationAccessibility {
	m.LoginRedirectUrl = v
	return m
}
func (m *ApplicationAccessibility) WithSelfService(v bool) *ApplicationAccessibility {
	m.SelfService = v
	return m
}


type ApplicationCredentials struct {
	Signing ApplicationCredentialsSigning `json:"signing, omitempty"`
	UserNameTemplate ApplicationCredentialsUsernameTemplate `json:"userNameTemplate, omitempty"`
}

func (m *ApplicationCredentials) WithSigning(v ApplicationCredentialsSigning) *ApplicationCredentials {
	m.Signing = v
	return m
}
func (m *ApplicationCredentials) WithUserNameTemplate(v ApplicationCredentialsUsernameTemplate) *ApplicationCredentials {
	m.UserNameTemplate = v
	return m
}


type ApplicationCredentialsOAuthClient struct {
	AutoKeyRotation bool `json:"autoKeyRotation, omitempty"`
	ClientId string `json:"client_id, omitempty"`
	ClientSecret string `json:"client_secret, omitempty"`
	TokenEndpointAuthMethod string `json:"token_endpoint_auth_method, omitempty"`
}

func (m *ApplicationCredentialsOAuthClient) WithAutoKeyRotation(v bool) *ApplicationCredentialsOAuthClient {
	m.AutoKeyRotation = v
	return m
}
func (m *ApplicationCredentialsOAuthClient) WithClientId(v string) *ApplicationCredentialsOAuthClient {
	m.ClientId = v
	return m
}
func (m *ApplicationCredentialsOAuthClient) WithClientSecret(v string) *ApplicationCredentialsOAuthClient {
	m.ClientSecret = v
	return m
}
func (m *ApplicationCredentialsOAuthClient) WithTokenEndpointAuthMethod(v string) *ApplicationCredentialsOAuthClient {
	m.TokenEndpointAuthMethod = v
	return m
}


type ApplicationCredentialsScheme struct {
}



type ApplicationCredentialsSigning struct {
	Kid string `json:"kid, omitempty"`
	LastRotated time.Time `json:"lastRotated, omitempty"`
	NextRotation time.Time `json:"nextRotation, omitempty"`
	RotationMode string `json:"rotationMode, omitempty"`
}

func (m *ApplicationCredentialsSigning) WithKid(v string) *ApplicationCredentialsSigning {
	m.Kid = v
	return m
}
func (m *ApplicationCredentialsSigning) WithRotationMode(v string) *ApplicationCredentialsSigning {
	m.RotationMode = v
	return m
}


type ApplicationCredentialsUsernameTemplate struct {
	Suffix string `json:"suffix, omitempty"`
	Template string `json:"template, omitempty"`
	Type string `json:"type, omitempty"`
}

func (m *ApplicationCredentialsUsernameTemplate) WithSuffix(v string) *ApplicationCredentialsUsernameTemplate {
	m.Suffix = v
	return m
}
func (m *ApplicationCredentialsUsernameTemplate) WithTemplate(v string) *ApplicationCredentialsUsernameTemplate {
	m.Template = v
	return m
}
func (m *ApplicationCredentialsUsernameTemplate) WithType(v string) *ApplicationCredentialsUsernameTemplate {
	m.Type = v
	return m
}


type ApplicationGroupAssignment struct {
	Embedded map[string]interface{} `json:"_embedded, omitempty"`
	Links map[string]interface{} `json:"_links, omitempty"`
	Id string `json:"id, omitempty"`
	LastUpdated time.Time `json:"lastUpdated, omitempty"`
	Priority int64 `json:"priority, omitempty"`
	Profile map[string]interface{} `json:"profile, omitempty"`
}

func (m *ApplicationGroupAssignment) WithPriority(v int64) *ApplicationGroupAssignment {
	m.Priority = v
	return m
}
func (m *ApplicationGroupAssignment) WithProfile(v map[string]interface{}) *ApplicationGroupAssignment {
	m.Profile = v
	return m
}


type ApplicationLicensing struct {
	SeatCount int64 `json:"seatCount, omitempty"`
}

func (m *ApplicationLicensing) WithSeatCount(v int64) *ApplicationLicensing {
	m.SeatCount = v
	return m
}


type ApplicationSettings struct {
	App ApplicationSettingsApplication `json:"app, omitempty"`
	Notifications ApplicationSettingsNotifications `json:"notifications, omitempty"`
}

func (m *ApplicationSettings) WithApp(v ApplicationSettingsApplication) *ApplicationSettings {
	m.App = v
	return m
}
func (m *ApplicationSettings) WithNotifications(v ApplicationSettingsNotifications) *ApplicationSettings {
	m.Notifications = v
	return m
}


type ApplicationSettingsApplication struct {
}



type ApplicationSettingsNotifications struct {
	Vpn ApplicationSettingsNotificationsVpn `json:"vpn, omitempty"`
}

func (m *ApplicationSettingsNotifications) WithVpn(v ApplicationSettingsNotificationsVpn) *ApplicationSettingsNotifications {
	m.Vpn = v
	return m
}


type ApplicationSettingsNotificationsVpn struct {
	HelpUrl string `json:"helpUrl, omitempty"`
	Message string `json:"message, omitempty"`
	Network ApplicationSettingsNotificationsVpnNetwork `json:"network, omitempty"`
}

func (m *ApplicationSettingsNotificationsVpn) WithHelpUrl(v string) *ApplicationSettingsNotificationsVpn {
	m.HelpUrl = v
	return m
}
func (m *ApplicationSettingsNotificationsVpn) WithMessage(v string) *ApplicationSettingsNotificationsVpn {
	m.Message = v
	return m
}
func (m *ApplicationSettingsNotificationsVpn) WithNetwork(v ApplicationSettingsNotificationsVpnNetwork) *ApplicationSettingsNotificationsVpn {
	m.Network = v
	return m
}


type ApplicationSettingsNotificationsVpnNetwork struct {
	Connection string `json:"connection, omitempty"`
	Exclude map[string]interface{} `json:"exclude, omitempty"`
	Include map[string]interface{} `json:"include, omitempty"`
}

func (m *ApplicationSettingsNotificationsVpnNetwork) WithConnection(v string) *ApplicationSettingsNotificationsVpnNetwork {
	m.Connection = v
	return m
}
func (m *ApplicationSettingsNotificationsVpnNetwork) WithExclude(v map[string]interface{}) *ApplicationSettingsNotificationsVpnNetwork {
	m.Exclude = v
	return m
}
func (m *ApplicationSettingsNotificationsVpnNetwork) WithInclude(v map[string]interface{}) *ApplicationSettingsNotificationsVpnNetwork {
	m.Include = v
	return m
}


type ApplicationSignOnMode struct {
}



type ApplicationVisibility struct {
	AppLinks map[string]interface{} `json:"appLinks, omitempty"`
	AutoSubmitToolbar bool `json:"autoSubmitToolbar, omitempty"`
	Hide ApplicationVisibilityHide `json:"hide, omitempty"`
}

func (m *ApplicationVisibility) WithAppLinks(v map[string]interface{}) *ApplicationVisibility {
	m.AppLinks = v
	return m
}
func (m *ApplicationVisibility) WithAutoSubmitToolbar(v bool) *ApplicationVisibility {
	m.AutoSubmitToolbar = v
	return m
}
func (m *ApplicationVisibility) WithHide(v ApplicationVisibilityHide) *ApplicationVisibility {
	m.Hide = v
	return m
}


type ApplicationVisibilityHide struct {
	IOS bool `json:"iOS, omitempty"`
	Web bool `json:"web, omitempty"`
}

func (m *ApplicationVisibilityHide) WithIOS(v bool) *ApplicationVisibilityHide {
	m.IOS = v
	return m
}
func (m *ApplicationVisibilityHide) WithWeb(v bool) *ApplicationVisibilityHide {
	m.Web = v
	return m
}


type AutoLoginApplication struct {
	Credentials SchemeApplicationCredentials `json:"credentials, omitempty"`
	Settings AutoLoginApplicationSettings `json:"settings, omitempty"`
}

func (m *AutoLoginApplication) WithCredentials(v SchemeApplicationCredentials) *AutoLoginApplication {
	m.Credentials = v
	return m
}
func (m *AutoLoginApplication) WithSettings(v AutoLoginApplicationSettings) *AutoLoginApplication {
	m.Settings = v
	return m
}


type AutoLoginApplicationSettings struct {
	SignOn AutoLoginApplicationSettingsSignOn `json:"signOn, omitempty"`
}

func (m *AutoLoginApplicationSettings) WithSignOn(v AutoLoginApplicationSettingsSignOn) *AutoLoginApplicationSettings {
	m.SignOn = v
	return m
}


type AutoLoginApplicationSettingsSignOn struct {
	LoginUrl string `json:"loginUrl, omitempty"`
	RedirectUrl string `json:"redirectUrl, omitempty"`
}

func (m *AutoLoginApplicationSettingsSignOn) WithLoginUrl(v string) *AutoLoginApplicationSettingsSignOn {
	m.LoginUrl = v
	return m
}
func (m *AutoLoginApplicationSettingsSignOn) WithRedirectUrl(v string) *AutoLoginApplicationSettingsSignOn {
	m.RedirectUrl = v
	return m
}


type BasicApplicationSettings struct {
	App BasicApplicationSettingsApplication `json:"app, omitempty"`
}

func (m *BasicApplicationSettings) WithApp(v BasicApplicationSettingsApplication) *BasicApplicationSettings {
	m.App = v
	return m
}


type BasicApplicationSettingsApplication struct {
	AuthURL string `json:"authURL, omitempty"`
	Url string `json:"url, omitempty"`
}

func (m *BasicApplicationSettingsApplication) WithAuthURL(v string) *BasicApplicationSettingsApplication {
	m.AuthURL = v
	return m
}
func (m *BasicApplicationSettingsApplication) WithUrl(v string) *BasicApplicationSettingsApplication {
	m.Url = v
	return m
}


type BasicAuthApplication struct {
	Credentials SchemeApplicationCredentials `json:"credentials, omitempty"`
	Name string `json:"name, omitempty"`
	Settings BasicApplicationSettings `json:"settings, omitempty"`
}

func (m *BasicAuthApplication) WithCredentials(v SchemeApplicationCredentials) *BasicAuthApplication {
	m.Credentials = v
	return m
}
func (m *BasicAuthApplication) WithName(v string) *BasicAuthApplication {
	m.Name = v
	return m
}
func (m *BasicAuthApplication) WithSettings(v BasicApplicationSettings) *BasicAuthApplication {
	m.Settings = v
	return m
}


type BookmarkApplication struct {
	Name string `json:"name, omitempty"`
	Settings BookmarkApplicationSettings `json:"settings, omitempty"`
}

func (m *BookmarkApplication) WithName(v string) *BookmarkApplication {
	m.Name = v
	return m
}
func (m *BookmarkApplication) WithSettings(v BookmarkApplicationSettings) *BookmarkApplication {
	m.Settings = v
	return m
}


type BookmarkApplicationSettings struct {
	App BookmarkApplicationSettingsApplication `json:"app, omitempty"`
}

func (m *BookmarkApplicationSettings) WithApp(v BookmarkApplicationSettingsApplication) *BookmarkApplicationSettings {
	m.App = v
	return m
}


type BookmarkApplicationSettingsApplication struct {
	RequestIntegration bool `json:"requestIntegration, omitempty"`
	Url string `json:"url, omitempty"`
}

func (m *BookmarkApplicationSettingsApplication) WithRequestIntegration(v bool) *BookmarkApplicationSettingsApplication {
	m.RequestIntegration = v
	return m
}
func (m *BookmarkApplicationSettingsApplication) WithUrl(v string) *BookmarkApplicationSettingsApplication {
	m.Url = v
	return m
}


type BrowserPluginApplication struct {
	Credentials SchemeApplicationCredentials `json:"credentials, omitempty"`
}

func (m *BrowserPluginApplication) WithCredentials(v SchemeApplicationCredentials) *BrowserPluginApplication {
	m.Credentials = v
	return m
}


type JsonWebKey struct {
	Alg string `json:"alg, omitempty"`
	Created time.Time `json:"created, omitempty"`
	E string `json:"e, omitempty"`
	ExpiresAt time.Time `json:"expiresAt, omitempty"`
	KeyOps map[string]interface{} `json:"key_ops, omitempty"`
	Kid string `json:"kid, omitempty"`
	Kty string `json:"kty, omitempty"`
	LastUpdated time.Time `json:"lastUpdated, omitempty"`
	N string `json:"n, omitempty"`
	Status string `json:"status, omitempty"`
	Use string `json:"use, omitempty"`
	X5c map[string]interface{} `json:"x5c, omitempty"`
	X5t string `json:"x5t, omitempty"`
	X5tS256 string `json:"x5t#S256, omitempty"`
	X5u string `json:"x5u, omitempty"`
}



type OAuthApplicationCredentials struct {
	OauthClient ApplicationCredentialsOAuthClient `json:"oauthClient, omitempty"`
}

func (m *OAuthApplicationCredentials) WithOauthClient(v ApplicationCredentialsOAuthClient) *OAuthApplicationCredentials {
	m.OauthClient = v
	return m
}


type OAuthEndpointAuthenticationMethod struct {
}



type OAuthGrantType struct {
}



type OAuthResponseType struct {
}



type OpenIdConnectApplication struct {
	Credentials OAuthApplicationCredentials `json:"credentials, omitempty"`
	Name string `json:"name, omitempty"`
	Settings OpenIdConnectApplicationSettings `json:"settings, omitempty"`
}

func (m *OpenIdConnectApplication) WithCredentials(v OAuthApplicationCredentials) *OpenIdConnectApplication {
	m.Credentials = v
	return m
}
func (m *OpenIdConnectApplication) WithName(v string) *OpenIdConnectApplication {
	m.Name = v
	return m
}
func (m *OpenIdConnectApplication) WithSettings(v OpenIdConnectApplicationSettings) *OpenIdConnectApplication {
	m.Settings = v
	return m
}


type OpenIdConnectApplicationConsentMethod struct {
}



type OpenIdConnectApplicationSettings struct {
	OauthClient OpenIdConnectApplicationSettingsClient `json:"oauthClient, omitempty"`
}

func (m *OpenIdConnectApplicationSettings) WithOauthClient(v OpenIdConnectApplicationSettingsClient) *OpenIdConnectApplicationSettings {
	m.OauthClient = v
	return m
}


type OpenIdConnectApplicationSettingsClient struct {
	ApplicationType string `json:"application_type, omitempty"`
	ClientUri string `json:"client_uri, omitempty"`
	ConsentMethod string `json:"consent_method, omitempty"`
	GrantTypes map[string]interface{} `json:"grant_types, omitempty"`
	LogoUri string `json:"logo_uri, omitempty"`
	PolicyUri string `json:"policy_uri, omitempty"`
	RedirectUris map[string]interface{} `json:"redirect_uris, omitempty"`
	ResponseTypes map[string]interface{} `json:"response_types, omitempty"`
	TosUri string `json:"tos_uri, omitempty"`
}

func (m *OpenIdConnectApplicationSettingsClient) WithApplicationType(v string) *OpenIdConnectApplicationSettingsClient {
	m.ApplicationType = v
	return m
}
func (m *OpenIdConnectApplicationSettingsClient) WithClientUri(v string) *OpenIdConnectApplicationSettingsClient {
	m.ClientUri = v
	return m
}
func (m *OpenIdConnectApplicationSettingsClient) WithConsentMethod(v string) *OpenIdConnectApplicationSettingsClient {
	m.ConsentMethod = v
	return m
}
func (m *OpenIdConnectApplicationSettingsClient) WithGrantTypes(v map[string]interface{}) *OpenIdConnectApplicationSettingsClient {
	m.GrantTypes = v
	return m
}
func (m *OpenIdConnectApplicationSettingsClient) WithLogoUri(v string) *OpenIdConnectApplicationSettingsClient {
	m.LogoUri = v
	return m
}
func (m *OpenIdConnectApplicationSettingsClient) WithPolicyUri(v string) *OpenIdConnectApplicationSettingsClient {
	m.PolicyUri = v
	return m
}
func (m *OpenIdConnectApplicationSettingsClient) WithRedirectUris(v map[string]interface{}) *OpenIdConnectApplicationSettingsClient {
	m.RedirectUris = v
	return m
}
func (m *OpenIdConnectApplicationSettingsClient) WithResponseTypes(v map[string]interface{}) *OpenIdConnectApplicationSettingsClient {
	m.ResponseTypes = v
	return m
}
func (m *OpenIdConnectApplicationSettingsClient) WithTosUri(v string) *OpenIdConnectApplicationSettingsClient {
	m.TosUri = v
	return m
}


type OpenIdConnectApplicationType struct {
}



type SamlApplication struct {
	Settings SamlApplicationSettings `json:"settings, omitempty"`
}

func (m *SamlApplication) WithSettings(v SamlApplicationSettings) *SamlApplication {
	m.Settings = v
	return m
}


type SamlApplicationSettings struct {
	SignOn SamlApplicationSettingsSignOn `json:"signOn, omitempty"`
}

func (m *SamlApplicationSettings) WithSignOn(v SamlApplicationSettingsSignOn) *SamlApplicationSettings {
	m.SignOn = v
	return m
}


type SamlApplicationSettingsSignOn struct {
	AssertionSigned bool `json:"assertionSigned, omitempty"`
	AttributeStatements map[string]interface{} `json:"attributeStatements, omitempty"`
	Audience string `json:"audience, omitempty"`
	AudienceOverride string `json:"audienceOverride, omitempty"`
	AuthnContextClassRef string `json:"authnContextClassRef, omitempty"`
	DefaultRelayState string `json:"defaultRelayState, omitempty"`
	Destination string `json:"destination, omitempty"`
	DestinationOverride string `json:"destinationOverride, omitempty"`
	DigestAlgorithm string `json:"digestAlgorithm, omitempty"`
	HonorForceAuthn bool `json:"honorForceAuthn, omitempty"`
	IdpIssuer string `json:"idpIssuer, omitempty"`
	Recipient string `json:"recipient, omitempty"`
	RecipientOverride string `json:"recipientOverride, omitempty"`
	RequestCompressed bool `json:"requestCompressed, omitempty"`
	ResponseSigned bool `json:"responseSigned, omitempty"`
	SignatureAlgorithm string `json:"signatureAlgorithm, omitempty"`
	SpIssuer string `json:"spIssuer, omitempty"`
	SsoAcsUrl string `json:"ssoAcsUrl, omitempty"`
	SsoAcsUrlOverride string `json:"ssoAcsUrlOverride, omitempty"`
	SubjectNameIdFormat string `json:"subjectNameIdFormat, omitempty"`
	SubjectNameIdTemplate string `json:"subjectNameIdTemplate, omitempty"`
}

func (m *SamlApplicationSettingsSignOn) WithAssertionSigned(v bool) *SamlApplicationSettingsSignOn {
	m.AssertionSigned = v
	return m
}
func (m *SamlApplicationSettingsSignOn) WithAttributeStatements(v map[string]interface{}) *SamlApplicationSettingsSignOn {
	m.AttributeStatements = v
	return m
}
func (m *SamlApplicationSettingsSignOn) WithAudience(v string) *SamlApplicationSettingsSignOn {
	m.Audience = v
	return m
}
func (m *SamlApplicationSettingsSignOn) WithAudienceOverride(v string) *SamlApplicationSettingsSignOn {
	m.AudienceOverride = v
	return m
}
func (m *SamlApplicationSettingsSignOn) WithAuthnContextClassRef(v string) *SamlApplicationSettingsSignOn {
	m.AuthnContextClassRef = v
	return m
}
func (m *SamlApplicationSettingsSignOn) WithDefaultRelayState(v string) *SamlApplicationSettingsSignOn {
	m.DefaultRelayState = v
	return m
}
func (m *SamlApplicationSettingsSignOn) WithDestination(v string) *SamlApplicationSettingsSignOn {
	m.Destination = v
	return m
}
func (m *SamlApplicationSettingsSignOn) WithDestinationOverride(v string) *SamlApplicationSettingsSignOn {
	m.DestinationOverride = v
	return m
}
func (m *SamlApplicationSettingsSignOn) WithDigestAlgorithm(v string) *SamlApplicationSettingsSignOn {
	m.DigestAlgorithm = v
	return m
}
func (m *SamlApplicationSettingsSignOn) WithHonorForceAuthn(v bool) *SamlApplicationSettingsSignOn {
	m.HonorForceAuthn = v
	return m
}
func (m *SamlApplicationSettingsSignOn) WithIdpIssuer(v string) *SamlApplicationSettingsSignOn {
	m.IdpIssuer = v
	return m
}
func (m *SamlApplicationSettingsSignOn) WithRecipient(v string) *SamlApplicationSettingsSignOn {
	m.Recipient = v
	return m
}
func (m *SamlApplicationSettingsSignOn) WithRecipientOverride(v string) *SamlApplicationSettingsSignOn {
	m.RecipientOverride = v
	return m
}
func (m *SamlApplicationSettingsSignOn) WithRequestCompressed(v bool) *SamlApplicationSettingsSignOn {
	m.RequestCompressed = v
	return m
}
func (m *SamlApplicationSettingsSignOn) WithResponseSigned(v bool) *SamlApplicationSettingsSignOn {
	m.ResponseSigned = v
	return m
}
func (m *SamlApplicationSettingsSignOn) WithSignatureAlgorithm(v string) *SamlApplicationSettingsSignOn {
	m.SignatureAlgorithm = v
	return m
}
func (m *SamlApplicationSettingsSignOn) WithSpIssuer(v string) *SamlApplicationSettingsSignOn {
	m.SpIssuer = v
	return m
}
func (m *SamlApplicationSettingsSignOn) WithSsoAcsUrl(v string) *SamlApplicationSettingsSignOn {
	m.SsoAcsUrl = v
	return m
}
func (m *SamlApplicationSettingsSignOn) WithSsoAcsUrlOverride(v string) *SamlApplicationSettingsSignOn {
	m.SsoAcsUrlOverride = v
	return m
}
func (m *SamlApplicationSettingsSignOn) WithSubjectNameIdFormat(v string) *SamlApplicationSettingsSignOn {
	m.SubjectNameIdFormat = v
	return m
}
func (m *SamlApplicationSettingsSignOn) WithSubjectNameIdTemplate(v string) *SamlApplicationSettingsSignOn {
	m.SubjectNameIdTemplate = v
	return m
}


type SamlAttributeStatement struct {
	Name string `json:"name, omitempty"`
	Namespace string `json:"namespace, omitempty"`
	Type string `json:"type, omitempty"`
	Values map[string]interface{} `json:"values, omitempty"`
}

func (m *SamlAttributeStatement) WithName(v string) *SamlAttributeStatement {
	m.Name = v
	return m
}
func (m *SamlAttributeStatement) WithNamespace(v string) *SamlAttributeStatement {
	m.Namespace = v
	return m
}
func (m *SamlAttributeStatement) WithType(v string) *SamlAttributeStatement {
	m.Type = v
	return m
}
func (m *SamlAttributeStatement) WithValues(v map[string]interface{}) *SamlAttributeStatement {
	m.Values = v
	return m
}


type SchemeApplicationCredentials struct {
	Password PasswordCredential `json:"password, omitempty"`
	RevealPassword bool `json:"revealPassword, omitempty"`
	Scheme string `json:"scheme, omitempty"`
	Signing ApplicationCredentialsSigning `json:"signing, omitempty"`
	UserName string `json:"userName, omitempty"`
}

func (m *SchemeApplicationCredentials) WithPassword(v PasswordCredential) *SchemeApplicationCredentials {
	m.Password = v
	return m
}
func (m *SchemeApplicationCredentials) WithRevealPassword(v bool) *SchemeApplicationCredentials {
	m.RevealPassword = v
	return m
}
func (m *SchemeApplicationCredentials) WithScheme(v string) *SchemeApplicationCredentials {
	m.Scheme = v
	return m
}
func (m *SchemeApplicationCredentials) WithSigning(v ApplicationCredentialsSigning) *SchemeApplicationCredentials {
	m.Signing = v
	return m
}
func (m *SchemeApplicationCredentials) WithUserName(v string) *SchemeApplicationCredentials {
	m.UserName = v
	return m
}


type SecurePasswordStoreApplication struct {
	Credentials SchemeApplicationCredentials `json:"credentials, omitempty"`
	Name string `json:"name, omitempty"`
	Settings SecurePasswordStoreApplicationSettings `json:"settings, omitempty"`
}

func (m *SecurePasswordStoreApplication) WithCredentials(v SchemeApplicationCredentials) *SecurePasswordStoreApplication {
	m.Credentials = v
	return m
}
func (m *SecurePasswordStoreApplication) WithName(v string) *SecurePasswordStoreApplication {
	m.Name = v
	return m
}
func (m *SecurePasswordStoreApplication) WithSettings(v SecurePasswordStoreApplicationSettings) *SecurePasswordStoreApplication {
	m.Settings = v
	return m
}


type SecurePasswordStoreApplicationSettings struct {
	App SecurePasswordStoreApplicationSettingsApplication `json:"app, omitempty"`
}

func (m *SecurePasswordStoreApplicationSettings) WithApp(v SecurePasswordStoreApplicationSettingsApplication) *SecurePasswordStoreApplicationSettings {
	m.App = v
	return m
}


type SecurePasswordStoreApplicationSettingsApplication struct {
	OptionalField1 string `json:"optionalField1, omitempty"`
	OptionalField1Value string `json:"optionalField1Value, omitempty"`
	OptionalField2 string `json:"optionalField2, omitempty"`
	OptionalField2Value string `json:"optionalField2Value, omitempty"`
	OptionalField3 string `json:"optionalField3, omitempty"`
	OptionalField3Value string `json:"optionalField3Value, omitempty"`
	PasswordField string `json:"passwordField, omitempty"`
	Url string `json:"url, omitempty"`
	UsernameField string `json:"usernameField, omitempty"`
}

func (m *SecurePasswordStoreApplicationSettingsApplication) WithOptionalField1(v string) *SecurePasswordStoreApplicationSettingsApplication {
	m.OptionalField1 = v
	return m
}
func (m *SecurePasswordStoreApplicationSettingsApplication) WithOptionalField1Value(v string) *SecurePasswordStoreApplicationSettingsApplication {
	m.OptionalField1Value = v
	return m
}
func (m *SecurePasswordStoreApplicationSettingsApplication) WithOptionalField2(v string) *SecurePasswordStoreApplicationSettingsApplication {
	m.OptionalField2 = v
	return m
}
func (m *SecurePasswordStoreApplicationSettingsApplication) WithOptionalField2Value(v string) *SecurePasswordStoreApplicationSettingsApplication {
	m.OptionalField2Value = v
	return m
}
func (m *SecurePasswordStoreApplicationSettingsApplication) WithOptionalField3(v string) *SecurePasswordStoreApplicationSettingsApplication {
	m.OptionalField3 = v
	return m
}
func (m *SecurePasswordStoreApplicationSettingsApplication) WithOptionalField3Value(v string) *SecurePasswordStoreApplicationSettingsApplication {
	m.OptionalField3Value = v
	return m
}
func (m *SecurePasswordStoreApplicationSettingsApplication) WithPasswordField(v string) *SecurePasswordStoreApplicationSettingsApplication {
	m.PasswordField = v
	return m
}
func (m *SecurePasswordStoreApplicationSettingsApplication) WithUrl(v string) *SecurePasswordStoreApplicationSettingsApplication {
	m.Url = v
	return m
}
func (m *SecurePasswordStoreApplicationSettingsApplication) WithUsernameField(v string) *SecurePasswordStoreApplicationSettingsApplication {
	m.UsernameField = v
	return m
}


type SwaApplication struct {
	Name string `json:"name, omitempty"`
	Settings SwaApplicationSettings `json:"settings, omitempty"`
}

func (m *SwaApplication) WithName(v string) *SwaApplication {
	m.Name = v
	return m
}
func (m *SwaApplication) WithSettings(v SwaApplicationSettings) *SwaApplication {
	m.Settings = v
	return m
}


type SwaApplicationSettings struct {
	App SwaApplicationSettingsApplication `json:"app, omitempty"`
}

func (m *SwaApplicationSettings) WithApp(v SwaApplicationSettingsApplication) *SwaApplicationSettings {
	m.App = v
	return m
}


type SwaApplicationSettingsApplication struct {
	ButtonField string `json:"buttonField, omitempty"`
	LoginUrlRegex string `json:"loginUrlRegex, omitempty"`
	PasswordField string `json:"passwordField, omitempty"`
	Url string `json:"url, omitempty"`
	UsernameField string `json:"usernameField, omitempty"`
}

func (m *SwaApplicationSettingsApplication) WithButtonField(v string) *SwaApplicationSettingsApplication {
	m.ButtonField = v
	return m
}
func (m *SwaApplicationSettingsApplication) WithLoginUrlRegex(v string) *SwaApplicationSettingsApplication {
	m.LoginUrlRegex = v
	return m
}
func (m *SwaApplicationSettingsApplication) WithPasswordField(v string) *SwaApplicationSettingsApplication {
	m.PasswordField = v
	return m
}
func (m *SwaApplicationSettingsApplication) WithUrl(v string) *SwaApplicationSettingsApplication {
	m.Url = v
	return m
}
func (m *SwaApplicationSettingsApplication) WithUsernameField(v string) *SwaApplicationSettingsApplication {
	m.UsernameField = v
	return m
}


type SwaThreeFieldApplication struct {
	Name string `json:"name, omitempty"`
	Settings SwaThreeFieldApplicationSettings `json:"settings, omitempty"`
}

func (m *SwaThreeFieldApplication) WithName(v string) *SwaThreeFieldApplication {
	m.Name = v
	return m
}
func (m *SwaThreeFieldApplication) WithSettings(v SwaThreeFieldApplicationSettings) *SwaThreeFieldApplication {
	m.Settings = v
	return m
}


type SwaThreeFieldApplicationSettings struct {
	App SwaThreeFieldApplicationSettingsApplication `json:"app, omitempty"`
}

func (m *SwaThreeFieldApplicationSettings) WithApp(v SwaThreeFieldApplicationSettingsApplication) *SwaThreeFieldApplicationSettings {
	m.App = v
	return m
}


type SwaThreeFieldApplicationSettingsApplication struct {
	ButtonSelector string `json:"buttonSelector, omitempty"`
	ExtraFieldSelector string `json:"extraFieldSelector, omitempty"`
	ExtraFieldValue string `json:"extraFieldValue, omitempty"`
	LoginUrlRegex string `json:"loginUrlRegex, omitempty"`
	PasswordSelector string `json:"passwordSelector, omitempty"`
	TargetUrl string `json:"targetUrl, omitempty"`
	UserNameSelector string `json:"userNameSelector, omitempty"`
}

func (m *SwaThreeFieldApplicationSettingsApplication) WithButtonSelector(v string) *SwaThreeFieldApplicationSettingsApplication {
	m.ButtonSelector = v
	return m
}
func (m *SwaThreeFieldApplicationSettingsApplication) WithExtraFieldSelector(v string) *SwaThreeFieldApplicationSettingsApplication {
	m.ExtraFieldSelector = v
	return m
}
func (m *SwaThreeFieldApplicationSettingsApplication) WithExtraFieldValue(v string) *SwaThreeFieldApplicationSettingsApplication {
	m.ExtraFieldValue = v
	return m
}
func (m *SwaThreeFieldApplicationSettingsApplication) WithLoginUrlRegex(v string) *SwaThreeFieldApplicationSettingsApplication {
	m.LoginUrlRegex = v
	return m
}
func (m *SwaThreeFieldApplicationSettingsApplication) WithPasswordSelector(v string) *SwaThreeFieldApplicationSettingsApplication {
	m.PasswordSelector = v
	return m
}
func (m *SwaThreeFieldApplicationSettingsApplication) WithTargetUrl(v string) *SwaThreeFieldApplicationSettingsApplication {
	m.TargetUrl = v
	return m
}
func (m *SwaThreeFieldApplicationSettingsApplication) WithUserNameSelector(v string) *SwaThreeFieldApplicationSettingsApplication {
	m.UserNameSelector = v
	return m
}


type WsFederationApplication struct {
	Name string `json:"name, omitempty"`
	Settings WsFederationApplicationSettings `json:"settings, omitempty"`
}

func (m *WsFederationApplication) WithName(v string) *WsFederationApplication {
	m.Name = v
	return m
}
func (m *WsFederationApplication) WithSettings(v WsFederationApplicationSettings) *WsFederationApplication {
	m.Settings = v
	return m
}


type WsFederationApplicationSettings struct {
	App WsFederationApplicationSettingsApplication `json:"app, omitempty"`
}

func (m *WsFederationApplicationSettings) WithApp(v WsFederationApplicationSettingsApplication) *WsFederationApplicationSettings {
	m.App = v
	return m
}


type WsFederationApplicationSettingsApplication struct {
	AttributeStatements string `json:"attributeStatements, omitempty"`
	AudienceRestriction string `json:"audienceRestriction, omitempty"`
	AuthnContextClassRef string `json:"authnContextClassRef, omitempty"`
	GroupFilter string `json:"groupFilter, omitempty"`
	GroupName string `json:"groupName, omitempty"`
	GroupValueFormat string `json:"groupValueFormat, omitempty"`
	NameIDFormat string `json:"nameIDFormat, omitempty"`
	Realm string `json:"realm, omitempty"`
	SiteURL string `json:"siteURL, omitempty"`
	UsernameAttribute string `json:"usernameAttribute, omitempty"`
	WReplyOverride bool `json:"wReplyOverride, omitempty"`
	WReplyURL string `json:"wReplyURL, omitempty"`
}

func (m *WsFederationApplicationSettingsApplication) WithAttributeStatements(v string) *WsFederationApplicationSettingsApplication {
	m.AttributeStatements = v
	return m
}
func (m *WsFederationApplicationSettingsApplication) WithAudienceRestriction(v string) *WsFederationApplicationSettingsApplication {
	m.AudienceRestriction = v
	return m
}
func (m *WsFederationApplicationSettingsApplication) WithAuthnContextClassRef(v string) *WsFederationApplicationSettingsApplication {
	m.AuthnContextClassRef = v
	return m
}
func (m *WsFederationApplicationSettingsApplication) WithGroupFilter(v string) *WsFederationApplicationSettingsApplication {
	m.GroupFilter = v
	return m
}
func (m *WsFederationApplicationSettingsApplication) WithGroupName(v string) *WsFederationApplicationSettingsApplication {
	m.GroupName = v
	return m
}
func (m *WsFederationApplicationSettingsApplication) WithGroupValueFormat(v string) *WsFederationApplicationSettingsApplication {
	m.GroupValueFormat = v
	return m
}
func (m *WsFederationApplicationSettingsApplication) WithNameIDFormat(v string) *WsFederationApplicationSettingsApplication {
	m.NameIDFormat = v
	return m
}
func (m *WsFederationApplicationSettingsApplication) WithRealm(v string) *WsFederationApplicationSettingsApplication {
	m.Realm = v
	return m
}
func (m *WsFederationApplicationSettingsApplication) WithSiteURL(v string) *WsFederationApplicationSettingsApplication {
	m.SiteURL = v
	return m
}
func (m *WsFederationApplicationSettingsApplication) WithUsernameAttribute(v string) *WsFederationApplicationSettingsApplication {
	m.UsernameAttribute = v
	return m
}
func (m *WsFederationApplicationSettingsApplication) WithWReplyOverride(v bool) *WsFederationApplicationSettingsApplication {
	m.WReplyOverride = v
	return m
}
func (m *WsFederationApplicationSettingsApplication) WithWReplyURL(v string) *WsFederationApplicationSettingsApplication {
	m.WReplyURL = v
	return m
}




