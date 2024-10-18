# Changelog
Running changelog of releases since `2.0.0-rc.4`

## v5.0.2
 - Fix unmarshalling error for Group owner object lastUpdated (#486) Thanks [@duytiennguyen-okta]

## v5.0.1
 - Fix unmarshalling error for Log object that does not have additional properties (#484) Thanks [@duytiennguyen-okta]

## v5.0.0
 - Update newest spec (#463) Thanks [@duytiennguyen-okta]

## v4.1.2
 - Fix panic issue when using access token is empty (#466) Thanks [@duytiennguyen-okta]
 - Fix object that does not have additional properties (#466) Thanks [@duytiennguyen-okta]

## v4.1.1
 - Fix panic issue when using bearer token (#463) Thanks [@duytiennguyen-okta]
 - Fix object that does not have additional properties (#463) Thanks [@duytiennguyen-okta]

## v4.1.0
 - Add support for dpop (#454) Thanks [@duytiennguyen-okta]
 - Fix object that does not have additional properties (#456) Thanks [@duytiennguyen-okta]
 - Fix header accept issue that only accept application/json (#457) Thanks [@duytiennguyen-okta]

## v4.0.0
 - Release v4 version of the sdk base on openapi spec v3 (#4) Thanks [@duytiennguyen-okta]

## v2.19.0
 - Upgrade dependency (#378) Thanks [@duytiennguyen-okta]

## v2.18.0
 - Added `WithHttpInterceptorAndHttpClientPtr` (#371) Thanks @dnelson27
 - Handle network errors and don't cancel context in request before reading response (#369) Thanks @ttimonen

## v2.17.0

 - Revert integer values representation as integer pointers
   - Original change was too agreesive for a minor release
   - All integer fields have a parallel pointer field that can be used as an alternative to deal with 0 values during JSON marshaling
     - Example `PasswordPolicy` model has fields int64 `Priority` and *int64 `PriorityPtr`
     - This change will not break code from release v2.15.0 and earlier
 - New models
   - `MultifactorEnrollmentPolicy`
   - `MultifactorEnrollmentPolicyAuthenticatorSettings`
   - `MultifactorEnrollmentPolicyAuthenticatorStatus`
   - `MultifactorEnrollmentPolicyAuthenticatorType`
   - `MultifactorEnrollmentPolicySettings`
   - `MultifactorEnrollmentPolicySettingsType`
 - New methods / endpoints
   - `Application` has `PreviewSAMLAppMetadata` method for `GET /api/v1/apps/{applicationId}/sso/saml/metadata` endpoint.
 - Added parameters
   - `Default` bool on `AuthorizationServer` model
 
## v2.16.0

 - All API integers, which previously had been typed as `int64`, are now typed
   `*int64` so that serialization of actual `0` values are not ignored when
   omitempty tags are present for serialization.

## v2.15.0

### Additions

 - Added `POST /api/v1/authenticators` for `CreateAuthenticator` method to `AuthenticatorResource` model by @monde in  https://github.com/okta/okta-management-openapi-spec/pull/154
 - Added `/api/v1/apps/${applicationId}/credentials/secrets` endpoints and `ClientSecret` model by @monde in https://github.com/okta/okta-management-openapi-spec/pull/146
 - Added property `pkce_required` to the `ApplicationCredentialsOAuthClient` model by @monde in https://github.com/okta/okta-management-openapi-spec/pull/145
 - Added `uiSchemaId` property to the `ProfileEnrollmentPolicyRuleAction` model by @emanor-okta in https://github.com/okta/okta-management-openapi-spec/pull/144

### Bug fixes

 - Corrected parameters order
   - `/api/v1/authorizationServers/{authServerId}/policies/{policyId}/rules`
   - `/api/v1/authorizationServers/{authServerId}/policies/{policyId}/rules/{ruleId}`
   - By @monde https://github.com/okta/okta-management-openapi-spec/pull/131

## v2.14.0

### Updates

 - Added property `inactivityPeriod` to `VerificationMethod` model [openapi-2.14.0](https://github.com/okta/okta-management-openapi-spec/releases/tag/openapi-2.14.0)
 - OAuth tokens are cached regardless of what request cache manager is utilized on the client.

## v2.13.0

Okta SDK Golang release in conjunction with [v2.13.0 of the Okta OpenAPI Specification (Management APIs)](https://github.com/okta/okta-management-openapi-spec/releases/tag/openapi-2.13.0)

### Updates

 - Added properties `host`, `integrationKey`, and `secretKey` to `AuthenticatorProviderConfiguration` model to support Duo Authenticator [openapi-2.13.0](https://github.com/okta/okta-management-openapi-spec/releases/tag/openapi-2.13.0)
 - Renamed query parameter `search` to `filter` on `GET /api/v1/groups` endpoint [openapi-2.13.0](https://github.com/okta/okta-management-openapi-spec/releases/tag/openapi-2.13.0)
 - New feature, SDK client can authenticate with Bearer token orienting the SDK client to the security scope of a specific application [#310](https://github.com/okta/okta-sdk-golang/pull/310)
 - Added ability to specify JWT kid header  [#308](https://github.com/okta/okta-sdk-golang/pull/308)

### Bug fixes

 - Schema API calls with one of / enum values will have correct type when calling API [#312](https://github.com/okta/okta-sdk-golang/pull/312)


## v2.12.1
  - Corrected out of order method and API parameters
    - `AuthorizationServerResource`
      - `ListAuthorizationServerPolicyRules(ctx context.Context, authServerId string, policyId string)`
      - `CreateAuthorizationServerPolicyRule(ctx context.Context, authServerId string, policyId string, body AuthorizationServerPolicyRule)`
  - Corrected `ListPolicies` generation for new `Policies` interface.

## v2.12.0

### Updates
  - Update code generation to v2.12.0 of Okta's openapi spec
  - Update to use trusted origins IFRAME_EMBED - IframeEmbedScopeAllowedApps in Scope
  - Completed missing code generation of Policy related models
    - Policy related models implement a `Policies` interface
    - Correctly handle pass by reference in `CreatePolicy`, `UpdatePolicy`

### Bug fixes
  - Corrected env var typo on `OKTA_CLIENT_REQUEST_TIMEOUT` in README
  - Corrected code generation pass by referrence of a Factor instance in `EnrollFactor`, `ActivateFactor`, and `VerifyFactor`
  - Provider query parameter can be boolean or string
  - Corrected out of order method and API parameters
    - `AuthorizationServerResource`
      - `DeleteAuthorizationServerPolicyRule(ctx context.Context, authServerId string, policyId string, ruleId string)`
      - `GetAuthorizationServerPolicyRule(ctx context.Context, authServerId string, policyId string, ruleId string)`
      - `UpdateAuthorizationServerPolicyRule(ctx context.Context, authServerId string, policyId string, ruleId string, body AuthorizationServerPolicyRule)`
    - `AuthorizationServerPolicyRuleResource`
      - `UpdateAuthorizationServerPolicyRule(ctx context.Context, authServerId string, policyId string, ruleId string, body AuthorizationServerPolicyRule)`
      - `DeleteAuthorizationServerPolicyRule(ctx context.Context, authServerId string, policyId string, ruleId string)`

## v2.11.2

### Updates
  - Added missing `_links` property to `ResponseLinks` model

## v2.11.1

### Updates
  - Changes `sendTestEmail` response to a `204 no content` instead of `200 success`
  - Adds `privateKey` and `teamId` properties to `IdentityProviderCredentialsSigning` to support Apple devices
  - Adds to `APPLE` to the enums in `FactorProvider' and `LogCredentialProvider

## v2.11.0

### Updates

- New models
  - `EmailTemplate`
  - `EmailTemplateContent`
  - `EmailTemplateCustomization`
  - `EmailTemplateCustomizationRequest`
  - `EmailTemplateResource`
  - `EmailTemplateTestRequest`
  - `EmailTemplateTouchPointVariant`
  - `IdpPolicyRuleAction`
  - `IdpPolicyRuleActionProvider`

## v2.10.1

### Updates

- Policy is first class parent with child policy types.

### Bug fixes

- Special case unmarshal for SocialAuthToken nullable expiresAt being blank from Okta API

## v2.10.0

Release of openapispec 2.10.0

## Supported Endpoints
* /api/v1/apps/{appId}/connections/default
* /api/v1/apps/{appId}/connections/default/lifecycle/activate
* /api/v1/apps/{appId}/connections/default/lifecycle/deactivate
* /api/v1/apps/{appId}/features
* /api/v1/apps/{appId}/features/{name}
* /api/v1/apps/{appId}/logo
* /api/v1/policies/{policyId}/rules/{ruleId}/lifecycle/activate
* /api/v1/policies/{policyId}/rules/{ruleId}/lifecycle/deactivate
* /api/v1/roles/{roleTypeOrRoleId}/subscriptions
* /api/v1/roles/{roleTypeOrRoleId}/subscriptions/{notificationType}
* /api/v1/roles/{roleTypeOrRoleId}/subscriptions/{notificationType}/subscribe
* /api/v1/roles/{roleTypeOrRoleId}/subscriptions/{notificationType}/unsubscribe
* /api/v1/users/{userId}/subscriptions
* /api/v1/users/{userId}/subscriptions/{notificationType}
* /api/v1/users/{userId}/subscriptions/{notificationType}/subscribe
* /api/v1/users/{userId}/subscriptions/{notificationType}/unsubscribe

## New Methods

## ApplicationResource
* ActivateDefaultProvisioningConnectionForApplication
* DeactivateDefaultProvisioningConnectionForApplication
* GetDefaultProvisioningConnectionForApplication
* GetFeatureForApplication
* ListFeaturesForApplication
* SetDefaultProvisioningConnectionForApplication
* UpdateFeatureForApplication
* UploadApplicationLogo

### BrandResource
* UploadBrandThemeBackgroundImage
* UploadBrandThemeFavicon
* UploadBrandThemeLogo

### OrgSettingResource
* UpdateOrgLogo

### SubscriptionResource
* GetRoleSubscriptionByNotificationType(ctx context.Context, roleTypeOrRoleId string, notificationType string) (*Subscription, *Response, error) {
* ListRoleSubscriptions
* SubscribeRoleSubscriptionByNotificationType
* SubscribeUserSubscriptionByNotificationType
* UnsubscribeRoleSubscriptionByNotificationType
* UnsubscribeUserSubscriptionByNotificationType

### UserResource
* GetUserSubscriptionByNotificationType
* ListUserSubscriptions

## New Constructors:
* NewApplicationFeature
* NewCapabilitiesCreateObject
* NewCapabilitiesObject
* NewCapabilitiesUpdateObject
* NewLifecycleCreateSettingObject
* NewLifecycleDeactivateSettingObject
* NewOrg2OrgApplication
* NewOrg2OrgApplicationSettings
* NewOrg2OrgApplicationSettingsApp
* NewPasswordSettingObject
* NewProfileSettingObject
* NewProvisioningConnection
* NewProvisioningConnectionProfile
* NewProvisioningConnectionRequest

## New Models and/or Updated models
* ApplicationFeature
* CapabilitiesCreateObject
* CapabilitiesObject
* CapabilitiesUpdateObject
* ChangeEnum
* LifecycleCreateSettingObject
* LifecycleDeactivateSettingObject
* NotificationType
* Org2OrgApplication
* Org2OrgApplicationSettings
* Org2OrgApplicationSettingsApp
* PasswordSettingObject
* ProfileSettingObject
* ProvisioningConnection
* ProvisioningConnectionAuthScheme
* ProvisioningConnectionProfile
* ProvisioningConnectionRequest
* ProvisioningConnectionStatus
* SeedEnum
* Subscription
* SubscriptionStatus


## v2.9.2

### Updates
- Adjustments were made the attributes on the `DeviceAccessPolicyRuleCondition` type.
- Adding `GroupProfileMap` attribute to `GroupProfile` struct to support additional properties. See #268


## v2.9.1

### New Types / Models
- `AccessPolicy`
- `AccessPolicyConstraint`
- `AccessPolicyConstraints`
- `AccessPolicyRule`
- `AccessPolicyRuleActions`
- `AccessPolicyRuleApplicationSignOn`
- `AccessPolicyRuleConditions`
- `AccessPolicyRuleCustomCondition`
- `AuthenticatorProvider`
- `AuthenticatorProviderConfiguration`
- `AuthenticatorProviderConfigurationUserNamePlate`
- `BrandResource`
- `Brand`
- `ChannelBinding`
- `Compliance`
- `DeviceAccessPolicyRuleCondition`
- `EmailTemplateTouchPointVariant`
- `EndUserDashboardTouchPointVariant`
- `ErrorPageTouchPointVariant`
- `FipsEnum`
- `ImageUploadResponse`
- `KnowledgeConstraint`
- `PossessionConstraint`
- `PreRegistrationInlineHook`
- `ProfileEnrollmentPolicy`
- `ProfileEnrollmentPolicyRule`
- `ProfileEnrollmentPolicyRuleAction`
- `ProfileEnrollmentPolicyRuleActions`
- `ProfileEnrollmentPolicyRuleActivationRequirement`
- `ProfileEnrollmentPolicyRuleProfileAttribute`
- `RequiredEnum`
- `SignInPageTouchPointVariant`
- `ThemeResource`
- `Theme`
- `ThemeResponse`
- `UserTypeCondition`
- `UserVerificationEnum`
- `VerificationMethod`

### New Methods
- `AuthenticatorResource`
  - `GetAuthenticator`
  - `UpdateAuthenticator`
  - `ListAuthenticators`
- `BrandResource`
  - `GetBrand`
  - `UpdateBrand`
  - `ListBrands`
  - `ListBrandThemes`
  - `GetBrandTheme`
  - `UpdateBrandTheme`
  - `DeleteBrandThemeBackgroundImage`
  - `UploadBrandThemeBackgroundImage`
  - `DeleteBrandThemeFavicon`
  - `UploadBrandThemeFavicon`
  - `DeleteBrandThemeLogo`
  - `UploadBrandThemeLogo`
- `OrgSettingResource`
  - `UpdateOrgLogo`
  - `GetOrgPreferences`
  - `HideOktaUIFooter`
  - `ShowOktaUIFooter`
  - `OptInUsersToOktaCommunicationEmails`
  - `OptOutUsersFromOktaCommunicationEmails`
  - `GetOrgOktaSupportSettings`
  - `ExtendOktaSupport`
  - `GrantOktaSupport`
  - `RevokeOktaSupport`
- `ThemeResource`
  - `GetBrandTheme`
  - `UpdateBrandTheme`

### Updates
- Updated Authenticators integration tests
- Integration tests on `IdentityProviderResource`

## v2.8.0

### Updates
- methods on client's Authenticator resource for `/api/v1/authenticators` endpoints and their integration tests
- methods on client's OrgSetting resource for `/api/v1/org` endpoints and their integration tests
- edge case type fixes in golang code generator
- `make pull-spec` pulls from https://github.com/okta/okta-management-openapi-spec master or branch set by `OPENAPI_SPEC_BRANCH` env variable
- golang generator sets api version from spec.json or by env variable override `OKTA_SDK_GOLANG_VERISON`
- integration test improvements

## v2.6.2

### Updates
- Okta API Group Schema support #256 thanks @bogdanprodan-okta

### Bug fixes
- Fix race condition #257 thanks @bogdanprodan-okta, @maxkondr

## v2.6.1

### Bug fixes
- Pass previous query parameters to the NextPage URL #252 @bogdanprodan-okta

## v2.6.0

### Updates
- Corrected UserSchemaBase properties field as UserSchemaAttribute #248 @bogdanprodan-okta
- Added ApplicationSettingsNotes to all apps

## v2.5.2

### Updates
- Added more fields to the error struct for different error response patterns from API #246 @bogdanprodan-okta

### Bug fixes
- Fixed error handling for 400/401/403 errors

## v2.5.1

### Updates
- Hardcode to version 2.5.1 off of main branch of Okta's openapi spec https://github.com/okta/okta-management-openapi-spec

### Bug fixes
- Fixed code generation inconsistencies from Okta's open API spec #243 @bogdanprodan-okta
- Fixed/improved app instances interface implementations, removed 'omitemty' tags #241 @bogdanprodan-okta
- Fixed failed access token requests are not retried. #194 @adw1n
- Fixed `client.User.CreateUser` example in README. #231 @Sanil2108

## v2.5.0

### Updates
- Update generated code to version 2.5.0 of Okta's openapi spec https://github.com/okta/okta-management-openapi-spec
- Refine Makefile's flexibility

## v2.3.1

### Updates
- Support opaque private keys (#225) @aborrello
- Improve OAuth 2.0 servapp doc (#221) @juanmatias
- Removed omitempty from ApplicationGroupAssignment priority field (#218) @bogdanprodan-okta
- Added `WithHttpClientPtr`, `WithHttpClientPtr` (#214) @bogdanprodan-okta
- Remove omitempty of 'default_scope' field (#212) @bogdanprodan-okta
- Add exported OKTA_ACCESS_TOKEN constant (#213) @bretterer
- Adds Stalebot workflow (#210) @bretterer

### Bug fixes
- Fixed typo in error message of "too many requests" #230 @halmai

## v2.3.0

### New Structs
- AuthorizationServerPolicy
- AuthorizationServerPolicyRule

### Bug fixes
- Fixed response body parsing
- Fix error handling for 401 error
- Retry on EOF errors from HTTP client

### Updates
- Update `SamlApplicationSettingsSignOn` properties to include `SingleLogout` and `SpCertificate`
- Update `InlineHookChannelConfig` properties to include `Method`
- Update `OpenIdConnectApplicationSettingsClient` properties to include `IdpInitiatedLogin`

## v2.2.1
### Bug Fixes
- Update default connection Timeout to `60` seconds (#185)
- Http Client now uses `Timeout` setting for connection timeout (#185)
- SDK no longer cancels context (#185)

## v2.2.0
### New Structs
- `OpenIdConnectApplicationSettingsClientKeys`
- `OpenIdConnectApplicationSettingsRefreshToken`
- `OpenIdConnectRefreshTokenRotationType`

### Bug Fixes
- Updates `okta/OpenIdConnectApplicationSettingsClient` properties to include `jwks []JsonWebKey` (Resolves #176)
- `NewCustomHotpUserFactor` sets factor type to `token:htop`

### Updates
- Removes vendor directory
- Updates `OpenIdConnectApplicationSettingsClient` properties to include `RefreshToken *OpenIdConnectApplicationSettingsRefreshToken`

## v2.1.0
### New Structs
- `AcsEndpoint`
- `CatalogApplication`
- `CatalogApplicationStatus`
- `CustomHotpUserFactor`
- `CustomHotpUserFactorProfile`
- `InlineHookStatus`
- `InlineHookType`
- `WebAuthnUserFactor`
- `WebAuthnUserFactorProfile`

### Bug Fixes
- Updates `okta/config.go` to use `github.com/okta/okta-sdk-golang/v2/okta/cache` (#167, Resolves #149)
- Updates `okta/config.go` to use `int64` for `ConnectionTimeout` and `RequestTimeout` (#166)
- `ListApplicationTargetsForApplicationAdministratorRoleForGroup` now returns correct `[]CatalogApplication`
- `ListApplicationTargetsForApplicationAdministratorRoleForUser` now returns correct `[]CatalogApplication`
- `JsonWebKey.ExpiresAt` changed to correctly use `*time.Time`
- `verifyUserFactorResponse.ExpiresAt` changed to correctly use `*time.Time`
- Correctly reads `.okta.yaml` from project root directory
- Updated typo in `PasswordCredentialHash` struct and changed `WorkerFactor` to the correct `WorkFactor`

### Updates
- Updates to @okta/openapi 2.1.1 (And associated Generator Updates)
- Added `UserId` to `OAuth2ScopeConsentGrant` struct
- **Now Returns error instead of panic from config in `okta.go`** (Resolves #122)
- `RequestExecutor` returns error if `privPem` is null when using PrivateKey authentication
- Added `AcsEndpoints` and `AllowMultipleAcsEndpoints` properties to `SamlApplicationSettingsSignOn`
- Added `FilterType` and `FilterValue` to `SamlAttributeStatement` (Resolves #76)
- Generated file header in line with Go standard (Resolves #118)

## v2.0.0
- Official Release of v2 Okta Golang SDK (June 4, 2020)
- `CSR` Model has been renamed to `Csr`
- `CSRMetadata` Model has been renamed to `CsrMetadata`
- `CSRMetadataSubject` Model has been renamed to `CsrMetadataSubject`
- `CSRMetadataSubjectAltNames` Model has been renamed to `CsrMetadataSubjectAltNames`
- `GenerateCsrForIdentityProvider` now sets accept header correctly to `application/json`

## v2.0.0-rc.5
- Remove `PreviewSamlMetadataForApplication` method from `ApplicationResource`
- `Kid` Query Param removed
- `KeepCurrent` Query Param removed
- `ActivateFactorRequest` struct added
- New Method `GenerateApplicationKey` added to `ApplicationResource`
- `FactorResultType` struct added
- `OpenIdConnectApplicationIssuerMode` struct added
- `InitiateLoginUri` property (type string) added to `OpenIdConnectApplicationSettingsClient` struct
- `IssuerMode` property (type string) added to `OpenIdConnectApplicationSettingsClient` struct
- `Format` property (type []string) added to `PolicySubject` struct
- `ExpiresAt` property (type *time.Time) added to `PushUserFactor` struct
- `FactorResult` property (type string) added to `PushUserFactor` struct
- Remove `ClearCurrentUserSession` from `UserResource`
- `UserFactorResource` method `ActivateFactor` now takes `ActivateFactorRequest` as the `body` param instead of `VerifyFactorRequest`
- `NextPassCode` property (type string) added to `VerifyFactorRequest` struct
- Update Readme to show how to access `RequestExecutor` (Fixes #83)
- Fixes Cache EOF (Fixes #103 and Fixes #36)
- Update all enum objects to `type string` (Fixes #95)
- Implement Pagination

## v2.0.0-rc.4
- Updated Resource Executor to handle nil pointer dereference (#125, fix suggested by @johanbrandhorst)
- Updates JWT Signer Package issues (#121, fix suggested by @Jmainguy)
