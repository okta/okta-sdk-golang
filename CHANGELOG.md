# Changelog
Running changelog of releases since `2.0.0-rc.4`

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
