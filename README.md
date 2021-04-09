[<img src="https://devforum.okta.com/uploads/oktadev/original/1X/bf54a16b5fda189e4ad2706fb57cbb7a1e5b8deb.png" align="right" width="256px"/>](https://devforum.okta.com/)
[![Build Status](https://img.shields.io/travis/okta/okta-sdk-golang.svg?logo=travis)](https://travis-ci.org/okta/okta-sdk-golang)
[![License](https://img.shields.io/github/license/okta/okta-sdk-golang.svg)](https://opensource.org/licenses/Apache-2.0)
[![Support](https://img.shields.io/badge/support-Developer%20Forum-blue.svg)][devforum]
[![API Reference](https://img.shields.io/badge/docs-reference-lightgrey.svg)][sdkapiref]

# Okta Golang management SDK
* [Release status](#release-status)
* [Need help?](#need-help)
* [Getting started](#getting-started)
* [Usage guide](#usage-guide)
* [Configuration reference](#configuration-reference)
* [Upgrading Guide](#upgrading-to-20x)
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

You can learn more on the [Okta + Golang][lang-landing] page in our documentation.

## Release status

This library uses semantic versioning and follows Okta's [library version policy](https://developer.okta.com/code/library-versions/).

| Version | Status                             |
| ------- | ---------------------------------- |
| 0.x     |  :warning: Beta Release (Retired)  |
| 1.x     |  :warning: Retiring on 2021-03-04  |
| 2.x     |  :heavy_check_mark: Release        |

The latest release can always be found on the [releases page][github-releases].

## Need help?

If you run into problems using the SDK, you can

* Ask questions on the [Okta Developer Forums][devforum]
* Post [issues][github-issues] here on GitHub (for code errors)

## Getting started

The SDK is compatible with Go version 1.12.x and up. For SDK v2 and above, you must use [Go Modules](https://blog.golang.org/using-go-modules) to install the SDK.

### Install current release
To install the Okta Golang SDK in your project:
  - Create a module file by running `go mod init`
    - You can skip this step if you already use `go mod`
  - Run `go get github.com/okta/okta-sdk-golang/v2`. This will add the SDK to your `go.mod` file.
  - Import the package in your project with `import "github.com/okta/okta-sdk-golang/v2/okta"`

### Installing legacy version
Although we do not suggest using the 1.x version of the SDK, you can still use it. Version 1.x is *retiring* and will not be supported past March 4, 2021. it will likely remain working after that date, but you should make a plan to migrate to the new 2.x version.

You can install v1 of the sdk by running `go get github.com/okta/okta-sdk-golang` and import the package in your project with `import "github.com/okta/okta-sdk-golang"`

### You'll also need

* An Okta account, called an _organization_ (sign up for a free [developer organization](https://developer.okta.com/signup) if you need one)
* An [API token](https://developer.okta.com/docs/api/getting_started/getting_a_token)

### Initialize a client

Construct a client instance by passing it your Okta domain name and API token:

```
ctx, client, err := okta.NewClient(context, okta.WithOrgUrl("https://{yourOktaDomain}"), okta.WithToken("{apiToken}"))
```

Hard-coding the Okta domain and API token works for quick tests, but for real projects you should use a more secure way of storing these values (such as environment variables). This library supports a few different configuration sources, covered in the [configuration reference](#configuration-reference) section.

## Usage guide

These examples will help you understand how to use this library. You can also browse the full [API reference documentation][sdkapiref].

Once you initialize a `client`, you can call methods to make requests to the Okta API. Most methods are grouped by the API endpoint they belong to. For example, methods that call the [Users API](https://developer.okta.com/docs/api/resources/users) are organized under `client.User`.


## Connection Retry / Rate Limiting
By default this SDK retries requests that are returned with a 429 exception. To disable this functionality set `OKTA_CLIENT_REQUESTTIMEOUT` and `OKTA_CLIENT_RATELIMIT_MAXRETRIES` to 0.

Setting only one of the values to zero disables that check. Meaning, by default, four retry attempts will be made. If you set `OKTA_CLIENT_REQUESTTIMEOUT` to 45 seconds and `OKTA_CLIENT_RATELIMIT_MAXRETRIES` to 0. This SDK will continue to retry indefinitely for 45 seconds. If both values are non zero, this SDK attempts to retry until either of the conditions are met (not both).

We use the Date header from the server to calculate the delta, as it's more reliable than system time. But always add 1 second to account for some clock skew in our service:

```
backoff_seconds = header['X-Rate-Limit-Reset'] - header['Date'] + 1s
```

If the `backoff_seconds` calculation exceeds the request timeout, the initial 429 response will be allowed through without additional attempts.

When creating your client, you can pass in these settings like you would with any other configuration.

```go
ctx, client, err := okta.NewClient(context, okta.WithRequestTimeout(45), okta.WithRateLimitMaxRetries(3))
```

### Authenticate a User
This library should only be used with the Okta management API. To call the [Authentication API](https://developer.okta.com/docs/api/resources/authn), you should construct your own HTTP requests.

### Get a User
```
ctx, client, err := okta.NewClient(context, okta.WithOrgUrl("https://{yourOktaDomain}"), okta.WithToken("{apiToken}"))
user, resp, err := client.User.GetUser(user.Id, nil)
```

### List all Users
```
ctx, client, err := okta.NewClient(context, okta.WithOrgUrl("https://{yourOktaDomain}"), okta.WithToken("{apiToken}"))
users, resp, err := client.User.ListUsers()
```

### Filter or search for Users
```
ctx, client, err := okta.NewClient(context, okta.WithOrgUrl("https://{yourOktaDomain}"), okta.WithToken("{apiToken}"))

filter := query.NewQueryParams(query.WithFilter("status eq \"ACTIVE\""))

users, resp, err := client.User.ListUsers(filter)
```

### Create a User
```
ctx, client, err := okta.NewClient(context, okta.WithOrgUrl("https://{yourOktaDomain}"), okta.WithToken("{apiToken}"))

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
ctx, client, err := okta.NewClient(context, okta.WithOrgUrl("https://{yourOktaDomain}"), okta.WithToken("{apiToken}"))

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
ctx, client, err := okta.NewClient(context, okta.WithOrgUrl("https://{yourOktaDomain}"), okta.WithToken("{apiToken}"))
user, resp, err := client.User.GetUser(user.Id, nil)

nickName = user.Profile["nickName"]
```

### Remove a User
You must first deactivate the user, and then you can delete the user.
```
ctx, client, err := okta.NewClient(context, okta.WithOrgUrl("https://{yourOktaDomain}"), okta.WithToken("{apiToken}"))
resp, err := client.User.DeactivateUser(user.Id, nil)

resp, err := client.User.DeactivateOrDeleteUser(user.Id, nil)
```

### List a User's Groups
```
ctx, client, err := okta.NewClient(context, okta.WithOrgUrl("https://{yourOktaDomain}"), okta.WithToken("{apiToken}"))

groups, resp, err := client.User.ListUserGroups(user.Id, nil)
```

### Create a Group
```
ctx, client, err := okta.NewClient(context, okta.WithOrgUrl("https://{yourOktaDomain}"), okta.WithToken("{apiToken}"))

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
ctx, client, err := okta.NewClient(context, okta.WithOrgUrl("https://{yourOktaDomain}"), okta.WithToken("{apiToken}"))

resp, err := client.Group.AddUserToGroup(group.Id, user.Id, nil)
```

### List a User's enrolled Factors
```
ctx, client, err := okta.NewClient(context, okta.WithOrgUrl("https://{yourOktaDomain}"), okta.WithToken("{apiToken}"))

allowedFactors, resp, err := client.Factor.ListSupportedFactors(user.Id)
```

### Enroll a User in a new Factor
```
ctx, client, err := okta.NewClient(context, okta.WithOrgUrl("https://{yourOktaDomain}"), okta.WithToken("{apiToken}"))

factorProfile := okta.NewSmsFactorProfile()
factorProfile.PhoneNumber = "5551234567"

factor := okta.NewSmsFactor()
factor.Profile = factorProfile

addedFactor, resp, err := client.Factor.AddFactor(user.Id, factor, nil)
```

### Activate a Factor
```
ctx, client, err := okta.NewClient(context, okta.WithOrgUrl("https://{yourOktaDomain}"), okta.WithToken("{apiToken}"))

factor, resp, err := client.Factor.ActivateFactor(user.Id, factor.Id, nil)
```

### Verify a Factor
```
ctx, client, err := okta.NewClient(context, okta.WithOrgUrl("https://{yourOktaDomain}"), okta.WithToken("{apiToken}"))

verifyFactorRequest := okta.VerifyFactorRequest{
  PassCode: "123456"
}
verifyFactorResp, resp, err := client.Factor.VerifyFactor(user.Id, factor.Id, verifyFactorRequest, nil)
```

### List all Applications
```
ctx, client, err := okta.NewClient(context, okta.WithOrgUrl("https://{yourOktaDomain}"), okta.WithToken("{apiToken}"))

applications, resp, err := client.Application.ListApplications(nil)

//applications will need to be cast from the interface into its concrete form before you can use it.
for _, a := range applications {
		if a.(*okta.Application).Name == "bookmark" {
			if a.(*okta.Application).Id == app2.(okta.BookmarkApplication).Id {
				application :=  *a.(*okta.BookmarkApplication) //This will cast it to a Bookmark Application
			}
		}
		// continue for each type you want to work with.
	}
```

### Get an Application
```
ctx, client, err := okta.NewClient(context, okta.WithOrgUrl("https://{yourOktaDomain}"), okta.WithToken("{apiToken}"))

//Getting a Basic Auth Application
application, resp, err = client.Application.GetApplication(appId, okta.NewBasicAuthApplication(), nil)

//To use the application, you must cast it to the type.
app := application.(*okta.BasicAuthApplication)
```

### Create a SWA Application
```
ctx, client, err := okta.NewClient(context, okta.WithOrgUrl("https://{yourOktaDomain}"), okta.WithToken("{apiToken}"))

swaAppSettingsApp := newSwaApplicationSettingsApplication()
swaAppSettingsApp.ButtonField = "btn-login"
swaAppSettingsApp.PasswordField = "txtbox-password"
swaAppSettingsApp.UsernameField = "txtbox-username"
swaAppSettingsApp.Url = "https://example.com/login.html"
swaAppSettingsApp.LoginUrlRegex = "REGEX_EXPRESSION"

swaAppSettings := newSwaApplicationSettings()
swaAppSettings.App = &swaAppSettingsApp

swaApp := newSwaApplication()
swaApp.Label = "Test App"
swaApp.Settings = &swaAppSettings

application, resp, err := client.Application.CreateApplication(swaApp, nil)
```

### Call other API endpoints
Not every API endpoint is represented by a method in this library. You can call any Okta management API endpoint using this generic syntax:
```
ctx, client, err := okta.NewClient(context, okta.WithOrgUrl("https://{yourOktaDomain}"), okta.WithToken("{apiToken}"))

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

### Access Request Executor
If you need to gain access to the request executor, we have provided a method off the `Client` to do so.

```go
re := client.GetRequestExecutor()
```

Doing this will provide you with the ability to create your own requests for the Okta API and call the `Do` method that handles all of the headers for you based on the configuration.

## Configuration reference

This library looks for configuration in the following sources:

0. An `okta.yaml` file in a `.okta` folder in the current user's home directory (`~/.okta/okta.yaml` or `%userprofile\.okta\okta.yaml`)
0. A `.okta.yaml` file in the application or project's root directory
0. Environment variables
0. Configuration explicitly passed to the constructor (see the example in [Getting started](#getting-started))

Higher numbers win. In other words, configuration passed via the constructor will override configuration found in environment variables, which will override configuration in `okta.yaml` (if any), and so on.

### YAML configuration

When you use an API Token instead of OAuth 2.0 the full YAML configuration looks like:

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

When you use OAuth 2.0 the full YAML configuration looks like:

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
    authorizationMode: "PrivateKey"
    clientId: "{yourClientId}"
    scopes:
      - scope.1
      - scope.2
    privateKey: |
        -----BEGIN RSA PRIVATE KEY-----
        MIIEogIBAAKCAQEAl4F5CrP6Wu2kKwH1Z+CNBdo0iteHhVRIXeHdeoqIB1iXvuv4
        THQdM5PIlot6XmeV1KUKuzw2ewDeb5zcasA4QHPcSVh2+KzbttPQ+RUXCUAr5t+r
        0r6gBc5Dy1IPjCFsqsPJXFwqe3RzUb...
        -----END RSA PRIVATE KEY-----
    requestTimeout: 0 # seconds
    rateLimit:
      maxRetries: 4
```

### Environment variables

Each one of the configuration values above can be turned into an environment variable name with the `_` (underscore) character:

* `OKTA_CLIENT_CONNECTIONTIMEOUT`
* `OKTA_CLIENT_TOKEN`
* and so on

## Upgrading to 2.0.x
The main purpose of this version is to include all documented, application/json endpoints
to the SDK. During this update we have made many changes to method names, as well as method signatures.

### Context
Every method that calls the API now has the ability to pass `context.Context` to it as the first parameter. If you do not have a context or do not know which context to use, you can pass `context.TODO()` to the methods.

### Method changes
We have spent time during this update making sure we become a little more uniform with naming of methods. This will require you to update some of your calls to the SDK with the new names.

All methods now specify the `Accept` and `Content-Type` headers when creating a new request. This allows for future use of the SDK to handle multiple `Accept` types.

### OAuth 2.0

Okta allows you to interact with Okta APIs using scoped OAuth 2.0 access tokens. Each access token enables the bearer to perform specific actions on specific Okta endpoints, with that ability controlled by which scopes the access token contains.

This SDK supports this feature only for service-to-service applications. Check out [our guides](https://developer.okta.com/docs/guides/implement-oauth-for-okta/overview/) to learn more about how to register a new service application using a private and public key pair, or follow the example steps at the end of this topic.

When using this approach you won't need an API Token because the SDK will request an access token for you. In order to use OAuth 2.0, construct a client instance by passing the following parameters:

```
ctx, client, err := okta.NewClient(context,
  okta.WithAuthorizationMode("PrivateKey"),
  okta.WithClientId("{{clientId}}"),
  okta.WithScopes(([]string{"okta.users.manage"})),
  okta.WithPrivateKey({{PEM PRIVATE KEY BLOCK}}) //when pasting blocks, use backticks and remove all space at beginning of each line.
)
```

#### Example

Let's say the need is to authenticate a script that will run in a pipeline (or any other automated way), and instead of using an API Token (that is bound to a user) the goal is to use a service app.

*To do this a public/private key pair will be needed.*

These are the requirements:

  - a public/private key par in JWT format [more here](https://developer.okta.com/docs/guides/implement-oauth-for-okta-serviceapp/create-publicprivate-keypair/)
  - a service app that uses the created key [more here](https://developer.okta.com/docs/guides/implement-oauth-for-okta-serviceapp/create-serviceapp-grantscopes/)
  - store the private key in a PEM format [more here](https://www.npmjs.com/package/pem-jwk)

To store the PEM formatted key in a json file the multiple lines need to be one-line formatted joined with "\n". It can be accomplished with the following command:

```bash
awk 'NF {sub(/\r/, ""); printf "%s\\n",$0;}' private.pem
```

##### Steps

Go to [this tool](https://mkjwk.org/) and create your keys. (It is recommend to user your own instance when creating keys for prod)

Save the three files in your PC, i.e. in *public-key*, *public-private-keypair* and *public-private-keypair-set* files.

Now create the PEM formatted private key. It can be done using [this tool](https://www.npmjs.com/package/pem-jwk) as follows:

```bash
pem-jwk public-private-keypair > private.pem
```

Create the service app as stated [here](https://developer.okta.com/docs/guides/implement-oauth-for-okta-serviceapp/create-serviceapp-grantscopes/#create-a-service-app), using *public-private-keypair-set*.

To grant the scopes, you can go to web console and do this in the usual way.

Now, let's say the PEM formatted private key will be stored in a json file, so your app can read it during the run time. It needs to be stored as a one-line.

If this is the PEM formatted key:

```bash
----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAmyX8wdrHK1ycOMeXNg3NOMQvebnfQp+3L5OaaiX16/+tLbwb
JTZDYh0EXLySMVsduRxC/1PQdPuI6x50TdkoB3C4JMuU968uJqkFp7fXXy5SMAej
HAyF67cY51dx15ztvakRNJPhhI5WaC20RfR/eow0IH5lGI3czcvTCChGau5qLue3
HqNDYFY+U3xhOlavSDdtmuxpIFsDycn/OjYjsV4lzyRrOArqtVV/kXHKx04T6A1x
Sc99999999999999999999999999999999999999999999999999EGekHlUAIUpw
Tqzdddddddddddddddddddddddddddddddddddddddddddddddddd5rNLLe5C7p/
LMta1rzm5TPYwazIbiMkFLAW02ToNAs9LGgDP+VRCZskl6+LuaA5XGabpi09ka7x
vJ88888888888888888888888888888888888888888888888888JDmIfMSU1tEw
Hmthd7qcwrx29ectcGHyQaX6iyYlQiBKCto+VwPrUq/qDCPYMIyqCTxAGTPLWQAz
Tqzdddddddddddddddddddddddddddddddddddddddddddddddddd5rNLLe5C7p/
HpDowa9Q+CBO5CEcPW4w9gsCgYA68a+82YtsiyYjdwSzscKIrw4ht3qAZQhGO4Id
H70kN1CkhHUNFf5UuRHJJ+s3BKlawBUwZaKD5KQ+lpnmrwqfArWu+3HNpX3LIPs5
vJ88888888888888888888888888888888888888888888888888JDmIfMSU1tEw
tATmAQKBgCKBkyHmEaS8tEAXcRI26oHOwTZAj6tJp2ODrEcrWtT0bR8wjaEdASdM
Tqzdddddddddddddddddddddddddddddddddddddddddddddddddd5rNLLe5C7p/
4b+xIYHz8dxmWPsZ7C4WbW9pw3Wn1Du/uvImwt0f4Jp6IPZM9vltXz0Dh7Yv5/SE
vJ88888888888888888888888888888888888888888888888888JDmIfMSU1tEw
tATmAQKBgCKBkyHmEaS8tEAXcRI26oHOwTZAj6tJp2ODrEcrWtT0bR8wjaEdASdM
Tqzdddddddddddddddddddddddddddddddddddddddddddddddddd5rNLLe5C7p/
vJ88888888888888888888888888888888888888888888888888JDmIfMSU1tEw
tATmAQKBgCKBkyHmEaS8tEAXcRI26oHOwTZAj6tJp2ODrEcrWtT0bR8wjaEdASdM
DKctGohIQ/ujUD9wzSvlaSZjBcKWw27yN0HiEBn+whKmO76PT7NFAQv/TG8ou3NE
ftlYhgBkwRwRfk7lEvmaTvJugd5g1E/9DAXTajlRYdohGubVz+2G
-----END RSA PRIVATE KEY-----
```

...then the value that must be added to the json file is:

```bash
----BEGIN RSA PRIVATE KEY-----\nMIIEowIBAAKCAQEAmyX8wdrHK1ycOMeXNg3NOMQvebnfQp+3L5OaaiX16/+tLbwb\nJTZDYh0EXLySMVsduRxC/1PQdPuI6x50TdkoB3C4JMuU968uJqkFp7fXXy5SMAej\nHAyF67cY51dx15ztvakRNJPhhI5WaC20RfR/eow0IH5lGI3czcvTCChGau5qLue3\nHqNDYFY+U3xhOlavSDdtmuxpIFsDycn/OjYjsV4lzyRrOArqtVV/kXHKx04T6A1x\nSc99999999999999999999999999999999999999999999999999EGekHlUAIUpw\nTqzdddddddddddddddddddddddddddddddddddddddddddddddddd5rNLLe5C7p/\nLMta1rzm5TPYwazIbiMkFLAW02ToNAs9LGgDP+VRCZskl6+LuaA5XGabpi09ka7x\nvJ88888888888888888888888888888888888888888888888888JDmIfMSU1tEw\nHmthd7qcwrx29ectcGHyQaX6iyYlQiBKCto+VwPrUq/qDCPYMIyqCTxAGTPLWQAz\nTqzdddddddddddddddddddddddddddddddddddddddddddddddddd5rNLLe5C7p/\nHpDowa9Q+CBO5CEcPW4w9gsCgYA68a+82YtsiyYjdwSzscKIrw4ht3qAZQhGO4Id\nH70kN1CkhHUNFf5UuRHJJ+s3BKlawBUwZaKD5KQ+lpnmrwqfArWu+3HNpX3LIPs5\nvJ88888888888888888888888888888888888888888888888888JDmIfMSU1tEw\ntATmAQKBgCKBkyHmEaS8tEAXcRI26oHOwTZAj6tJp2ODrEcrWtT0bR8wjaEdASdM\nTqzdddddddddddddddddddddddddddddddddddddddddddddddddd5rNLLe5C7p/\n4b+xIYHz8dxmWPsZ7C4WbW9pw3Wn1Du/uvImwt0f4Jp6IPZM9vltXz0Dh7Yv5/SE\nvJ88888888888888888888888888888888888888888888888888JDmIfMSU1tEw\ntATmAQKBgCKBkyHmEaS8tEAXcRI26oHOwTZAj6tJp2ODrEcrWtT0bR8wjaEdASdM\nTqzdddddddddddddddddddddddddddddddddddddddddddddddddd5rNLLe5C7p/\nvJ88888888888888888888888888888888888888888888888888JDmIfMSU1tEw\ntATmAQKBgCKBkyHmEaS8tEAXcRI26oHOwTZAj6tJp2ODrEcrWtT0bR8wjaEdASdM\nDKctGohIQ/ujUD9wzSvlaSZjBcKWw27yN0HiEBn+whKmO76PT7NFAQv/TG8ou3NE\nftlYhgBkwRwRfk7lEvmaTvJugd5g1E/9DAXTajlRYdohGubVz+2G\n-----END RSA PRIVATE KEY-----
```

This way, if a json file like this one is created:

```yaml
{
    "Oktadomain" : "yourorg.okta.com",
    "clientId" : "yourclientid",
    "privateKey" : "----BEGIN RSA PRIVATE KEY-----\nMIIEowIBAAKCAQEAmyX8wdrHK1ycOMeXNg3NOMQvebnfQp+3L5OaaiX16/+tLbwb\nJTZDYh0EXLySMVsduRxC/1PQdPuI6x50TdkoB3C4JMuU968uJqkFp7fXXy5SMAej\nHAyF67cY51dx15ztvakRNJPhhI5WaC20RfR/eow0IH5lGI3czcvTCChGau5qLue3\nHqNDYFY+U3xhOlavSDdtmuxpIFsDycn/OjYjsV4lzyRrOArqtVV/kXHKx04T6A1x\nSc99999999999999999999999999999999999999999999999999EGekHlUAIUpw\nTqzdddddddddddddddddddddddddddddddddddddddddddddddddd5rNLLe5C7p/\nLMta1rzm5TPYwazIbiMkFLAW02ToNAs9LGgDP+VRCZskl6+LuaA5XGabpi09ka7x\nvJ88888888888888888888888888888888888888888888888888JDmIfMSU1tEw\nHmthd7qcwrx29ectcGHyQaX6iyYlQiBKCto+VwPrUq/qDCPYMIyqCTxAGTPLWQAz\nTqzdddddddddddddddddddddddddddddddddddddddddddddddddd5rNLLe5C7p/\nHpDowa9Q+CBO5CEcPW4w9gsCgYA68a+82YtsiyYjdwSzscKIrw4ht3qAZQhGO4Id\nH70kN1CkhHUNFf5UuRHJJ+s3BKlawBUwZaKD5KQ+lpnmrwqfArWu+3HNpX3LIPs5\nvJ88888888888888888888888888888888888888888888888888JDmIfMSU1tEw\ntATmAQKBgCKBkyHmEaS8tEAXcRI26oHOwTZAj6tJp2ODrEcrWtT0bR8wjaEdASdM\nTqzdddddddddddddddddddddddddddddddddddddddddddddddddd5rNLLe5C7p/\n4b+xIYHz8dxmWPsZ7C4WbW9pw3Wn1Du/uvImwt0f4Jp6IPZM9vltXz0Dh7Yv5/SE\nvJ88888888888888888888888888888888888888888888888888JDmIfMSU1tEw\ntATmAQKBgCKBkyHmEaS8tEAXcRI26oHOwTZAj6tJp2ODrEcrWtT0bR8wjaEdASdM\nTqzdddddddddddddddddddddddddddddddddddddddddddddddddd5rNLLe5C7p/\nvJ88888888888888888888888888888888888888888888888888JDmIfMSU1tEw\ntATmAQKBgCKBkyHmEaS8tEAXcRI26oHOwTZAj6tJp2ODrEcrWtT0bR8wjaEdASdM\nDKctGohIQ/ujUD9wzSvlaSZjBcKWw27yN0HiEBn+whKmO76PT7NFAQv/TG8ou3NE\nftlYhgBkwRwRfk7lEvmaTvJugd5g1E/9DAXTajlRYdohGubVz+2G\n-----END RSA PRIVATE KEY-----"
}
```

This can be read from Go and used directly in the client creation:

```go
  ctx, client, err := okta.NewClient(context,
    okta.WithOrgUrl("https://"+Oktadomain),
    okta.WithAuthorizationMode("PrivateKey"),
    okta.WithClientId(clientId),
    okta.WithScopes(scopes),
    okta.WithPrivateKey(privateKey),
    )
```
### Extending the Client
When calling `okta.NewClient()` we allow for you to pass custom instances of `http.Client` and `cache.Cache`.

```
myClient := &http.Client{}

myCache := NewCustomCacheDriver()

ctx, client, err := okta.NewClient(context, okta.WithOrgUrl("https://{yourOktaDomain}"), okta.WithToken("{apiToken}"), okta.WithHttpClient(myClient), okta.WithCacheManager(myCache))
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

### Refreshing Cache for Specific Call
If you have an issue where you do a `GET`, then a `DELETE`, and then re-issue a `GET` to the original endpoint, you may have an issue with the cache returning with the deleted resource. An example of this is listing application users, delete an application user, and then listing them again.

You can solve this by running `client.GetRequestExecutor().RefreshNext()` before your second `ListApplicationUsers` call, which will tell the call to delete the cache for this endpoint and make a new call.

```go
appUserList, _, _ = client.Application.ListApplicationUsers(context.TODO(), appId, nil)

client.Application.DeleteApplicationUser(context.TODO(), appId, appUser.Id, nil)

client.GetRequestExecutor().RefreshNext()
appUserList, _, _ = client.Application.ListApplicationUsers(context.TODO(), appId, nil)
```

### Pagination
If your request comes back with more than the default or set limit, you can request the next page.

Example of listing users 1 at a time:
```go
query := query.NewQueryParams(query.WithLimit(1))
users, resp, err := client.User.ListUsers(ctx, query)
// Do something with your users until you run out of users to iterate.
if resp.HasNextPage() {
  var nextUserSet []*okta.User
  resp, err = resp.Next(ctx, &nextUserSet)
}
```

## Building the SDK

In most cases, you won't need to build the SDK from source. If you want to build it yourself, you'll need these prerequisites:

- Clone the repo
- Run `make build` from the root of the project

## Contributing

We're happy to accept contributions and PRs! Please see the [contribution guide](/okta/okta-sdk-golang/blob/master/CONTRIBUTING.md) to understand how to structure a contribution.


[devforum]: https://devforum.okta.com/
[sdkapiref]: https://godoc.org/github.com/okta/okta-sdk-golang/okta
[lang-landing]: https://developer.okta.com/code/go/
[github-issues]: /okta/okta-sdk-golang/issues
[github-releases]: /okta/okta-sdk-golang/releases
