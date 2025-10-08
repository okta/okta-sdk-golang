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
