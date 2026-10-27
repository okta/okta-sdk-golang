# Okta Golang management SDK migration guide

## Migrating from 2.x to 4.x

In releases prior to version 3 we use an Open API v2 specification, and an Okta custom client generator to partially generate our SDK. A new version of the Open API specification (V3) has been released, and new well-known generators are now available and well received by the community. Planning the future of this SDK, we consider this a good opportunity to modernize by aligning with established standards for API client generation.

We acknowledge that migrating from v2 to v4 will require considerable effort, but we expect this change to benefit our customers in the long term.

With OpenAPI v3, we saw an opportunity for improvement in several areas:

* We can provide an API specification that follows the OpenAPI v3 standard and eliminate custom attributes and vendor extensions.
* Given that our specification is now compliant with OASv3, we can use well-known generators used and maintained by the community. In this case, we chose openapi-generator.tech.
* Given that we eliminated custom attributes, we can speed up our release process and let our customers access new APIs sooner.

### OktaClient vs API clients

In releases prior to version 6, you would instantiate a global `OktaClient` and access specific API clients via its properties. Now, each API has its own client and you only instantiate those clients you are interested in:

_Before_
```go
ctx, client, err := okta.NewClient(
    context.TODO(),
    okta.WithOrgUrl("https://{yourOktaDomain}"),
    okta.WithToken("{apiToken}"),
)

if err != nil {
    fmt.Printf("Error: %v\n", err)
}

user, resp, err := client.User.GetUser(ctx, "{UserId|Username|Email}")
if err != nil {
    fmt.Printf("Error Getting User: %v\n", err)
}
fmt.Printf("User: %+v\n Response: %+v\n\n",user, resp)
```

_Now_
```go
config, err := okta.NewConfiguration(
    okta.WithOrgUrl("https://{yourOktaDomain}"),
    okta.WithToken("{apiToken}"),
)
if err != nil {
    return nil, err
}
client = okta.NewAPIClient(config)

user, resp, err := client.UserAPI.GetUser(apiClient.cfg.Context, "{UserId|Username|Email}").Execute()
if err != nil {
    fmt.Printf("Error Getting User: %v\n", err)
}
fmt.Printf("User: %+v\n Response: %+v\n\n",user, resp)
```

### Manipulate Custom Attributes

Models that support dynamic properties now expose the `AdditionalProperties` property:

_Before:_

```go
profile := okta.UserProfile{}
profile["firstName"] = "John"
profile["lastName"] = "Activate"
```

_Now:_

```go
profile := make(map[string]interface{})
profile["firstName"] = "John"
profile["lastName"] = "Activate"
user.Profile.AdditionalProperties = profile
```

### Polymorphic models

In previous versions, the sdk v2 will get provided a generic interface such as App (Authenticator,Policy, etc..) where devs would then have to cast to a specific model to use. Now, you will be provide a list of model to choose from.

_Before:_

```go
retrievedApp, _, err := client.Application.GetApplication(ctx, appId, okta.NewBasicAuthApplication(), nil)
app := retrievedApp.(*okta.BasicAuthApplication)
```

_Now:_

```go
retrievedApp, _, err := client.ApplicationAPI.GetApplication(apiClient.cfg.Context, createdApp.BasicAuthApplication.GetId()).Execute()
app := retrievedApp.BasicAuthApplication
```

### Get data from _links

In previous versions, you had to manipulate raw data via the `Resource` convenience methods to access `Links`. Now, `Links` are exposed are standard properties:

_Before:_

```go
var links = createdApp.Links.(map[string]interface{})
accessPolicyLinks, ok := links["accessPolicy"]
accessPolicyLinksHref, ok := accessPolicyLinks.(map[string]interface{})
href, ok := accessPolicyLinksHref["accessPolicy"]
```

_Now:_

```go
var accessPolicyHref = createdApp.Links.AccessPolicy.Href
```

### Features parity

The following features have been ported to 4.x:

* Inline configuration, configuration via environment variables, appsettings.json or YAML files
* Manual pagination for collections
* Default retry strategy for 429 HTTP responses and ability to provide your own strategy
* Web proxy
* OAuth for Okta

## Migrating from 4.x to 5.x
Version 5x of the SDK is essentially a name change

## Migrating from 5.x to 6.x

Version 6 of the Okta Go SDK includes several improvements, API updates, and breaking changes that require code modifications.

### Prerequisites

- Go 1.24 or later (v6 requires a newer Go version)
- Update your `go.mod` file to use the new version

### Installation

Update your `go.mod` file:

```bash
go get github.com/okta/okta-sdk-golang/v6@latest
```

Or update manually in `go.mod`:
```go
require github.com/okta/okta-sdk-golang/v6 v6.0.0
```

### Breaking Changes

#### 1. Module Path Change

The import path has changed to include the major version:

**v5:**
```go
import "github.com/okta/okta-sdk-golang/v5/okta"
```

**v6:**
```go
import "github.com/okta/okta-sdk-golang/v6/okta"
```

#### 2. Client Initialization

The client initialization has been simplified:

**v5:**
```go
ctx, client, err := okta.NewClient(
    context.TODO(),
    okta.WithOrgUrl("https://dev-123456.okta.com"),
    okta.WithToken("your-api-token"),
)
```

**v6:**
```go
config := okta.NewConfiguration(
    okta.WithOrgUrl("https://dev-123456.okta.com"),
    okta.WithToken("your-api-token"),
)
client := okta.NewAPIClient(config)
```

#### 3. Context Handling

Context is now passed per-request instead of being stored in the client:

**v5:**
```go
user, resp, err := client.User.GetUser(ctx, userId)
```

**v6:**
```go
user, resp, err := client.UserAPI.GetUser(context.Background(), userId).Execute()
```

#### 4. API Method Names

API methods now follow a more consistent naming pattern with the `Execute()` pattern:

**v5:**
```go
users, resp, err := client.User.ListUsers(ctx, &query.Params{
    Limit: 10,
})
```

**v6:**
```go
users, resp, err := client.UserAPI.ListUsers(context.Background()).
    Limit(10).
    Execute()
```

#### 5. Request Builders

v6 uses a builder pattern for request configuration:

**v5:**
```go
app, resp, err := client.Application.CreateApplication(ctx,
    okta.NewBasicAuthApplication(),
    &query.Params{Activate: &activate},
)
```

**v6:**
```go
app, resp, err := client.ApplicationAPI.CreateApplication(context.Background()).
    Application(okta.NewBasicAuthApplication()).
    Activate(true).
    Execute()
```

#### 6. Error Handling

Error responses are now more structured:

**v5:**
```go
if err != nil {
    if oktaErr, ok := err.(*okta.Error); ok {
        fmt.Printf("Error Code: %s\n", oktaErr.ErrorCode)
    }
}
```

**v6:**
```go
if err != nil {
    if apiErr, ok := err.(*okta.GenericOpenAPIError); ok {
        fmt.Printf("Error: %s\n", apiErr.Error())
        fmt.Printf("Body: %s\n", apiErr.Body())
    }
}
```

#### 7. Pagination

Pagination handling has been updated:

**v5:**
```go
users, resp, err := client.User.ListUsers(ctx, nil)
for resp.HasNextPage() {
    var nextUsers []*okta.User
    resp, err = resp.Next(ctx, &nextUsers)
    users = append(users, nextUsers...)
}
```

**v6:**
```go
request := client.UserAPI.ListUsers(context.Background())
for {
    users, resp, err := request.Execute()
    if err != nil {
        break
    }
    // Process users

    if resp.HasNextPage() {
        request = request.After(resp.NextPageToken())
    } else {
        break
    }
}
```

#### 8. Model Changes

Some model types have been renamed or restructured:

**v5:**
```go
user := &okta.User{
    Profile: &okta.UserProfile{
        FirstName: "John",
        LastName:  "Doe",
        Email:     "john.doe@example.com",
        Login:     "john.doe@example.com",
    },
}
```

**v6:**
```go
profile := okta.NewUserProfile()
profile.SetFirstName("John")
profile.SetLastName("Doe")
profile.SetEmail("john.doe@example.com")
profile.SetLogin("john.doe@example.com")

user := okta.NewCreateUserRequest()
user.SetProfile(*profile)
```

### Common Migration Patterns

#### 1. User Operations

**v5:**
```go
// Create user
user, resp, err := client.User.CreateUser(ctx, body, nil)

// Update user
user, resp, err := client.User.UpdateUser(ctx, userId, body, nil)

// Delete user
resp, err := client.User.DeactivateOrDeleteUser(ctx, userId, nil)
```

**v6:**
```go
// Create user
user, resp, err := client.UserAPI.CreateUser(context.Background()).
    Body(*body).
    Execute()

// Update user
user, resp, err := client.UserAPI.UpdateUser(context.Background(), userId).
    User(*body).
    Execute()

// Delete user
resp, err := client.UserAPI.DeactivateUser(context.Background(), userId).
    Execute()
```

#### 2. Group Operations

**v5:**
```go
// Create group
group, resp, err := client.Group.CreateGroup(ctx, body)

// Add user to group
resp, err := client.Group.AddUserToGroup(ctx, groupId, userId)
```

**v6:**
```go
// Create group
group, resp, err := client.GroupAPI.CreateGroup(context.Background()).
    Group(*body).
    Execute()

// Add user to group
resp, err := client.GroupAPI.AddUserToGroup(context.Background(), groupId, userId).
    Execute()
```

#### 3. Application Operations

**v5:**
```go
// Create application
app, resp, err := client.Application.CreateApplication(ctx, body, params)

// Assign user to application
appUser, resp, err := client.Application.AssignUserToApplication(ctx, appId, body)
```

**v6:**
```go
// Create application
app, resp, err := client.ApplicationAPI.CreateApplication(context.Background()).
    Application(body).
    Execute()

// Assign user to application
appUser, resp, err := client.ApplicationAPI.AssignUserToApplication(
    context.Background(), appId).
    AppUser(*body).
    Execute()
```

### Configuration Changes

#### Rate Limiting

**v5:**
```go
_, client, err := okta.NewClient(ctx,
    okta.WithRequestTimeout(30),
    okta.WithRateLimitMaxRetries(3),
)
```

**v6:**
```go
config := okta.NewConfiguration()
config.HTTPClient = &http.Client{
    Timeout: 30 * time.Second,
}
config.RetryConfiguration = &okta.RetryConfiguration{
    MaxRetries: 3,
}
```

### Testing Changes

If you're mocking the SDK for tests:

**v5:**
```go
mockClient := &okta.Client{
    User: mockUserService,
}
```

**v6:**
```go
// Use interface-based mocking
type UserAPIInterface interface {
    GetUser(ctx context.Context, userId string) okta.ApiGetUserRequest
}
```

### Deprecations

The following features are deprecated in v6:
- Direct field access on models (use getter/setter methods)
- Query parameter objects (use builder methods)
- Some legacy authentication methods

### Features Parity

The following features have been ported to v6:
- Configuration via environment variables, YAML files, or inline
- Manual pagination for collections
- Default retry strategy for 429 HTTP responses and ability to provide your own strategy
- Web proxy support
- OAuth for Okta authentication
- Context-based request handling
- Structured error responses

## Migrating from 7.x to 8.x

Version 8 of the Okta Go SDK is generated from Okta Management API specification 2026.42.1. Some schemas changed in the specification, which changes the generated models. This section covers the changes that require code modifications.

### Installation

Update your `go.mod` file:

```bash
go get github.com/okta/okta-sdk-golang/v8@latest
```

Update your import paths:

```go
import "github.com/okta/okta-sdk-golang/v8/okta"
```

### Breaking Changes

#### 1. `SamlApplicationSettingsSignOn` boolean fields are now optional

The `allowMultipleAcsEndpoints`, `assertionSigned`, `honorForceAuthn`, `requestCompressed`, and `responseSigned` properties are no longer required in the specification. API responses can omit them, so the SDK no longer fails to unmarshal SAML apps that leave them out ([#594](https://github.com/okta/okta-sdk-golang/issues/594)).

As a result:
- `NewSamlApplicationSettingsSignOn` no longer takes any arguments.
- The fields changed from `bool` to `*bool` and are omitted from the JSON request when not set.

**v7:**
```go
signOn := okta.NewSamlApplicationSettingsSignOn(false, true, false, false, true)

if signOn.AssertionSigned {
    // ...
}
```

**v8:**
```go
signOn := okta.NewSamlApplicationSettingsSignOn()
signOn.SetAllowMultipleAcsEndpoints(false)
signOn.SetAssertionSigned(true)
signOn.SetHonorForceAuthn(false)
signOn.SetRequestCompressed(false)
signOn.SetResponseSigned(true)

// GetAssertionSigned returns false when the field isn't set.
// Use HasAssertionSigned to check whether the field was present in the response.
if signOn.GetAssertionSigned() {
    // ...
}
```

#### 2. `Scopes` moved from `ProfileEnrollmentPolicyRuleAction` to `PreRegistrationInlineHook`

The `Scopes` field and its `GetScopes`/`SetScopes` methods were removed from `ProfileEnrollmentPolicyRuleAction`. Scopes are now set on each `PreRegistrationInlineHook` and specify which registration flows trigger that hook (`SELF_SERVICE_REGISTRATION`, `PROGRESSIVE_PROFILING`, or both). `scopes` is required when `preRegistrationInlineHooks` is set.

**v7:**
```go
hook := okta.NewPreRegistrationInlineHook()
hook.SetInlineHookId("{inlineHookId}")

action := okta.NewProfileEnrollmentPolicyRuleAction()
action.SetPreRegistrationInlineHooks([]okta.PreRegistrationInlineHook{*hook})
action.SetScopes([]string{"SELF_SERVICE_REGISTRATION"})
```

**v8:**
```go
hook := okta.NewPreRegistrationInlineHook()
hook.SetInlineHookId("{inlineHookId}")
hook.SetScopes([]string{"SELF_SERVICE_REGISTRATION", "PROGRESSIVE_PROFILING"})

action := okta.NewProfileEnrollmentPolicyRuleAction()
action.SetPreRegistrationInlineHooks([]okta.PreRegistrationInlineHook{*hook})
```

#### 3. `AppInstanceSearchType` has a new required `Count` field

`AppInstanceSearchType` now includes `Count`, the number of app instances of that type in the org. For admins with scoped app access, the count only includes apps they can view. Because `count` is required, `NewAppInstanceSearchType` takes it as the first argument.

**v7:**
```go
searchType := okta.NewAppInstanceSearchType("{type}")
```

**v8:**
```go
searchType := okta.NewAppInstanceSearchType(0, "{type}")
count := searchType.GetCount()
```

#### 4. `ShowSignInWithAPasskeyButton` no longer defaults to `false`

`NewAuthenticatorMethodWebAuthnAllOfSettings` and `NewAuthenticatorMethodWebAuthnAllOfSettingsWithDefaults` no longer set `ShowSignInWithAPasskeyButton` to `false`. The field stays `nil` and is omitted from the request unless you set it. If your code relied on the SDK sending `false`, set it explicitly:

```go
settings := okta.NewAuthenticatorMethodWebAuthnAllOfSettings()
settings.SetShowSignInWithAPasskeyButton(false)
```
