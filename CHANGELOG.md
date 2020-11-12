# Changelog
Running changelog of releases since `2.0.0-rc.4`

## v2.3.0
### New Structs
- `OpenIdConnectApplicationSettingsRefreshToken`
- `OpenIdConnectRefreshTokenRotationType`

### Bug Fixes
- `NewCustomHotpUserFactor` sets factor type to `token:htop`

### Updates
- Updates `OpenIdConnectApplicationSettingsClient` properties to include `RefreshToken *OpenIdConnectApplicationSettingsRefreshToken`

## v2.2.0
### New Structs
- `OpenIdConnectApplicationSettingsClientKeys`

### Bug Fixes
- Updates `okta/OpenIdConnectApplicationSettingsClient` properties to include `jwks []JsonWebKey` (Resolves #176)

### Updates
- Removes vendor directory

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
