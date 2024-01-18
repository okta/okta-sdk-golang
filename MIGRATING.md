# Okta Golang management SDK migration guide

## Migrating from 2.x to 3.x

In releases prior to version 3 we use an Open API v2 specification, and an Okta custom client generator to partially generate our SDK. A new version of the Open API specification (V3) has been released, and new well-known generators are now available and well received by the community. Planning the future of this SDK, we consider this a good opportunity to modernize by aligning with established standards for API client generation. 

We acknowledge that migrating from v2 to v3 will require considerable effort, but we expect this change to benefit our customers in the long term.

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

The following features have been ported to 6.x:

* Inline configuration, configuration via environment variables, appsettings.json or YAML files
* Manual pagination for collections
* Default retry strategy for 429 HTTP responses and ability to provide your own strategy
* Web proxy 
* OAuth for Okta
