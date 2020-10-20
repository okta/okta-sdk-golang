# Changelog
Running changelog of releases since `2.0.0-rc.4`

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
