[<img src="https://devforum.okta.com/uploads/oktadev/original/1X/bf54a16b5fda189e4ad2706fb57cbb7a1e5b8deb.png" align="right" width="256px"/>](https://devforum.okta.com/)
[![Build Status](https://img.shields.io/travis/okta/okta-sdk-golang.svg?logo=travis)](https://travis-ci.org/okta/okta-sdk-golang)
[![GitHub](https://img.shields.io/github/license/okta/okta-sdk-github.svg)](https://github.com/okta/okta-sdk-golang)
![Beta Release](https://img.shields.io/badge/Beta-Unstable-yellow.svg)
[![License](https://img.shields.io/github/license/okta/okta-sdk-golang.svg)](https://opensource.org/licenses/Apache-2.0)
[![Support](https://img.shields.io/badge/support-Developer%20Forum-blue.svg)][devforum]
[![API Reference](https://img.shields.io/badge/docs-reference-lightgrey.svg)][sdkapiref]

# Okta Golang management SDK

> :warning: Beta alert! This library is in beta. See [release status](#release-status) for more information.

* [Release status](#release-status)
* [Need help?](#need-help)
* [Getting started](#getting-started)
* [Usage guide](#usage-guide)
* [Configuration reference](#configuration-reference)
* [Building the SDK](#building-the-sdk)
* [Contributing](#contributing)

This repository contains the Okta management SDK for Golang. This SDK can be used in your server-side code to interact with the Okta management API and

* Create and update users with the [Users API](https://developer.okta.com/docs/api/resources/users)
* Add security factors to users with the [Factors API](https://developer.okta.com/docs/api/resources/factors)
* Manage groups with the [Groups API](https://developer.okta.com/docs/api/resources/groups)
* Manage applications with the [Apps API](https://developer.okta.com/docs/api/resources/apps)
* Much more!

We also publish these libraries for Golang:

* [JWT Verifier](https://github.com/okta/okta-jwt-verifier-golang)

You can learn more on the [Okta + Golang](lang-landing) page in our documentation.

## Release status

This library uses semantic versioning and follows Okta's [library version policy](https://developer.okta.com/code/library-versions/).

:warning: Beta alert! This library is in beta. Breaking changes may be introduced at minor versions in the 0.x range. Please lock your dependency to a specific version until this library reaches 1.x.

| Version | Status                    |
| ------- | ------------------------- |
| 0.x  | :heavy_check_mark: Beta Release               |

The latest release can always be found on the [releases page][github-releases].

## Need help?

If you run into problems using the SDK, you can

* Ask questions on the [Okta Developer Forums][devforum]
* Post [issues][github-issues] here on GitHub (for code errors)

## Getting started

Follow these steps to install the Okta Golang SDK in your project:
- Run `go get github.com/okta/okta-sdk-golang`

You'll also need

* An Okta account, called an _organization_ (sign up for a free [developer organization](https://developer.okta.com/signup) if you need one)
* An [API token](https://developer.okta.com/docs/api/getting_started/getting_a_token)

Construct a client instance by passing it your Okta domain name and API token:

```
config := okta.NewConfig().WithOrgUrl("{yourOktaDomain}").WithToken("{apiToken}")
client := okta.NewClient(config, nil, nil)
```

Hard-coding the Okta domain and API token works for quick tests, but for real projects you should use a more secure way of storing these values (such as environment variables). This library supports a few different configuration sources, covered in the [configuration reference](#configuration-reference) section.

### Extending the Client
When calling `okta.NewClient()` we allow for you to pass custom instances of `http.Client` and `cache.Cache`.

```
myClient := &http.Client{}

myCache := NewCustomCacheDriver()

config := okta.NewConfig().WithOrgUrl("{yourOktaDomain}").WithToken("{apiToken}")
client := okta.NewClient(config, myClient, myCache)
```


### Extending or Creating New Cache Manager
You can create a custom cache driver by implementing `cache.Cache`

```
type CustomCacheDriver struct {
}

func NewCustomCacheDriver() Cache {
	return CustomCacheDriver{}
}

func (c CustomCacheDriver) Get(key string) *http.Response {
	return nil
}

func (c CustomCacheDriver) Set(key string, value *http.Response) {}

func (c CustomCacheDriver) Delete(key string) {}

func (c CustomCacheDriver) Clear() {}

func (c CustomCacheDriver) Has(key string) bool {
	return false
}
```

## Usage guide

These examples will help you understand how to use this library. You can also browse the full [API reference documentation][sdkapiref].

Once you initialize a `client`, you can call methods to make requests to the Okta API. Most methods are grouped by the API endpoint they belong to. For example, methods that call the [Users API](https://developer.okta.com/docs/api/resources/users) are organized under `client.User`.

### Authenticate a User
This library should be used with the Okta management API. To call the [Authentication API](https://developer.okta.com/docs/api/resources/authn), you should construct your own HTTP requests.

### Get a User
```
config := okta.NewConfig()
client := okta.NewClient(config, nil, nil)
user, resp, err := client.User.GetUser(user.Id, nil)
```

### List all Users
```
config := okta.NewConfig()
client := okta.NewClient(config, nil, nil)
users, resp, err := client.User.ListUsers()
```

### Filter or search for Users
```
config := okta.NewConfig()
client := okta.NewClient(config, nil, nil)

filter := query.NewQueryParams(query.WithFilter("status eq \"ACTIVE\""))

users, resp, err := client.User.ListUsers(filter)
```

### Create a User
```
config := okta.NewConfig()
client := okta.NewClient(config, nil, nil)

p := &okta.PasswordCredential{
		Value: "Abcd1234",
}
uc := &okta.UserCredentials{
		Password: p,
}
profile := okta.UserProfile{}
profile["firstName"] = "John"
profile["lastName"] = "Activate"
profile["email"] = "john-activate@example.com"
profile["login"] = "john-activate@example.com"
u := &okta.User{
		Credentials: uc,
		Profile:     &profile,
}

user, resp, err := client.User.CreateUser(*u, nil)
```

### Update a User
```
config := okta.NewConfig()
client := okta.NewClient(config, nil, nil)

newProfile := *user.Profile
newProfile["nickName"] = "Batman"
updatedUser := &okta.User{
  Profile: &newProfile,
}
user, resp, err := client.User.UpdateUser(user.Id, *updatedUser, nil)
```

### Get and set custom attributes
Custom attributes must first be defined in the Okta profile editor. Then, you can work with custom attributes on a user:
```
config := okta.NewConfig()
client := okta.NewClient(config, nil, nil)
user, resp, err := client.User.GetUser(user.Id, nil)

nickName = user.Profile["nickName"]
```

### Remove a User
You must first deactivate the user, and then you can delete the user.
```
config := okta.NewConfig()
client := okta.NewClient(config, nil, nil)
resp, err := client.User.DeactivateUser(user.Id, nil)

resp, err := client.User.DeactivateOrDeleteUser(user.Id, nil)
```

### List a User's Groups
```
config := okta.NewConfig()
client := okta.NewClient(config, nil, nil)

groups, resp, err := client.User.ListUserGroups(user.Id, nil)
```

### Create a Group
```
config := okta.NewConfig()
client := okta.NewClient(config, nil, nil)

gp := &okta.GroupProfile{
  Name: "Get Test Group",
}
g := &okta.Group{
  Profile: gp,
}
group, resp, err := client.Group.CreateGroup(*g, nil)
```

### Add a User to a Group
```
config := okta.NewConfig()
client := okta.NewClient(config, nil, nil)

resp, err := client.Group.AddUserToGroup(group.Id, user.Id, nil)
```

> Factors and Applications are currently not fully functioning.  The SDK will be able to read factors and applications from the API, but they are not set to the correct types yet.  This is coming soon.
### List a User's enrolled Factors
```
config := okta.NewConfig()
client := okta.NewClient(config, nil, nil)

factors, resp, err := client.Factor.ListFactors(group.Id, user.Id, nil)
```

### Enroll a User in a new Factor
```
config := okta.NewConfig()
client := okta.NewClient(config, nil, nil)

factorProfile := okta.SmsFactorProfile{
  PhoneNumber: "5551234567",
}
factor := okta.SmsFactor{
  Profile: &factorProfile,
}
factor, resp, err := client.Factor.AddFactor(user.Id, factor, nil)
```

### Activate a Factor
```
config := okta.NewConfig()
client := okta.NewClient(config, nil, nil)

factor, resp, err := client.Factor.ActivateFactor(user.Id, factor.Id, nil)
```

### Verify a Factor
```
config := okta.NewConfig()
client := okta.NewClient(config, nil, nil)

verifyFactorRequest := okta.VerifyFactorRequest{
  PassCode: "123456"
}
verifyFactorResp, resp, err := client.Factor.VerifyFactor(user.Id, factor.Id, verifyFactorRequest, nil)
```

### List all Applications
```
config := okta.NewConfig()
client := okta.NewClient(config, nil, nil)

applications, resp, err := client.Application.ListApplications(nil)
```

### Get an Application
```
config := okta.NewConfig()
client := okta.NewClient(config, nil, nil)

application, resp, err := client.Applicaiton.GetApplication(applicationId, nil)
```

### Create a SWA Application
```
config := okta.NewConfig()
client := okta.NewClient(config, nil, nil)

swaAppSettingsApp := okta.SwaApplicationSettingsApplication{
    ButtonField:   "btn-login",
    PasswordField: "txtbox-password",
    UsernameField: "txtbox-username",
    Url:           "https://example.com/login.html",
    LoginUrlRegex: "REGEX_EXPRESSION",
}
swaAppSettings := okta.SwaApplicationSettings{
    App: &swaAppSettingsApp,
}
swaApp := okta.SwaApplication{
    Name:     "Test App",
    Settings: &swaAppSettings,
}
application, resp, err := client.Application.CreateApplication(swaApp, nil)
```

### Call other API endpoints
Not every API endpoint is represented by a method in this library. You can call any Okta management API endpoint using this generic syntax:
```
config := okta.NewConfig()
client := okta.NewClient(config, nil, nil)

url := "https://golang.oktapreview.com/api/v1/authorizationServers

type Signing struct {
	RotationMode string     `json:"rotationMode,omitempty"`
	LastRotated  *time.Time `json:"lastRotated,omitempty"`
	NextRotation *time.Time `json:"nextRotation,omitempty"`
	Kid          string     `json:"kid,omitempty"`
}

type Credentials struct {
	Signing *Signing `json:"signing,omitempty"`
}

type AuthorizationServer struct {
	Id          string       `json:"id,omitempty"`
	Name        string       `json:"name,omitempty"`
	Description string       `json:"description,omitempty"`
	Audiences   []string     `json:"audiences,omitempty"`
	Issuer      string       `json:"issuer,omitempty"`
	IssuerMode  string       `json:"issuerMode,omitempty"`
	Status      string       `json:"status,omitempty"`
	Created     *time.Time   `json:"created,omitempty"`
	LastUpdated *time.Time   `json:"lastUpdated,omitempty"`
	Credentials *Credentials `json:"credentials,omitempty"`
	Embedded    interface{}  `json:"_embedded,omitempty"`
	Links       interface{}  `json:"_links,omitempty"`
}

as := AuthorizationServer{
  Name:        "Sample Authorization Server",
  Description: "Sample Authorization Server description",
  Audiences:   []string{"api://default"},
}

req, err := m.client.requestExecutor.NewRequest("POST", url, as)
if err != nil {
  return nil, nil, err
}

var authServer *AuthorizationServer
resp, err := m.client.requestExecutor.Do(req, &authServer)
if err != nil {
  return nil, resp, err
}
return authServer, resp, nil
```

## Configuration reference

This library looks for configuration in the following sources:

0. An `okta.yaml` file in a `.okta` folder in the current user's home directory (`~/.okta/okta.yaml` or `%userprofile\.okta\okta.yaml`)
0. An `okta.yaml` file in a `.okta` folder in the application or project's root directory
0. Environment variables
0. Configuration explicitly passed to the constructor (see the example in [Getting started](#getting-started))

Higher numbers win. In other words, configuration passed via the constructor will override configuration found in environment variables, which will override configuration in `okta.yaml` (if any), and so on.

### YAML configuration

The full YAML configuration looks like:

```yaml
okta:
  client:
    connectionTimeout: 30 # seconds
    orgUrl: "https://{yourOktaDomain}"
    proxy:
      port: null
      host: null
      username: null
      password: null
    token: {apiToken}
```

### Environment variables

Each one of the configuration values above can be turned into an environment variable name with the `_` (underscore) character:

* `OKTA_CLIENT_CONNECTIONTIMEOUT`
* `OKTA_CLIENT_TOKEN`
* and so on

## Building the SDK

In most cases, you won't need to build the SDK from source. If you want to build it yourself, you'll need these prerequisites:

- Clone the repo
- Run `make build` from the root of the project

## Contributing

We're happy to accept contributions and PRs! Please see the [contribution guide](contributing.md) to understand how to structure a contribution.


[devforum]: https://devforum.okta.com/
[sdkapiref]: https://godoc.org/github.com/okta/okta-sdk-golang/okta
[lang-landing]: https://developer.okta.com/code/golang/
[github-issues]: /issues
[github-releases]: /releases
