# Changelog
Running changelog of releases since `2.0.0-rc.4`

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


## v2.0.0-rc.4
- Updated Resource Executor to handle nil pointer dereference (#125, fix suggested by @johanbrandhorst)
- Updates JWT Signer Package issues (#121, fix suggested by @Jmainguy)
