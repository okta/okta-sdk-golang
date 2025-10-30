[<img src="https://cdn.brandfolder.io/2VK5Y09C/at/bb3mqsj5ssrgxtc5fbvtx/Logo-H_Developer-blue.svg" align="right" width="384px"/>](https://developer.okta.com/)
[![License: Apache 2.0](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
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

This repository contains the Okta management SDK for Golang. This SDK can be
used in your server-side code to interact with the Okta management API and

* Create and update users with the [Users API](https://developer.okta.com/docs/api/resources/users)
* Add security factors to users with the [Factors API](https://developer.okta.com/docs/api/resources/factors)
* Manage groups with the [Groups API](https://developer.okta.com/docs/api/resources/groups)
* Manage applications with the [Apps API](https://developer.okta.com/docs/api/resources/apps)
* Much more!

We also publish these libraries for Golang:

* [JWT Verifier](https://github.com/okta/okta-jwt-verifier-golang)

You can learn more on the [Okta + Golang][lang-landing] page in our documentation.

## Release status

This library uses semantic versioning and follows Okta's [library version
policy](https://developer.okta.com/code/library-versions/).

| Version | Status                                                       |
| ------- |--------------------------------------------------------------|
| 0.x     | :warning: Beta Release (Retired)                             |
| 1.x     | :warning: Retired
| 2.x     | :warning: Retired                                            |
| 3.x     | :warning: Retired                                            |
| 4.x     | :warning: Retiring                                           |
| 5.x     | :heavy_check_mark: Release                                   |
| 6.x     | :heavy_check_mark: Release ([migration guide](MIGRATING.md)) |

The latest release can always be found on the [releases page][github-releases].

## Need help?

If you run into problems using the SDK, you can

* Ask questions on the [Okta Developer Forums][devforum]
* Post [issues][github-issues] here on GitHub (for code errors)

## Getting started

The SDK is compatible with Go version 1.24.x and up. For SDK v2 and above, you
must use [Go Modules](https://blog.golang.org/using-go-modules) to install the
SDK.

### Install current release

To install the Okta Golang SDK in your project:
  - Create a module file by running `go mod init`
    - You can skip this step if you already use `go mod`
  - Run `go get github.com/okta/okta-sdk-golang/v6@latest`. This will add
    the SDK to your `go.mod` file.
  - Import the package in your project with `import
   "github.com/okta/okta-sdk-golang/v6/okta"`

### Installing legacy version

Although we do not suggest using the 2.x, 3.x, 4.x, and 5.x version of the SDK, you can still use
it. The earlier versions are either retired or in an unsupported state. It
will likely remain working currently, but you should make a plan to
migrate to the new 6.x version.

You can install the latest version of the SDK by running `go get
github.com/okta/okta-sdk-golang@latest` and import the package in your project
with `import "github.com/okta/okta-sdk-golang"`

### You'll also need

* An Okta account, called an _organization_ (sign up for a free [developer
  organization](https://developer.okta.com/signup) if you need one)
* An [API
  token](https://developer.okta.com/docs/api/getting_started/getting_a_token)

### Initialize a client

Construct a client instance by passing it your Okta domain name and API token:

```go
import (
	"fmt"
	"context"
	"github.com/okta/okta-sdk-golang/v6/okta"
)

func main() {
  config, err := okta.NewConfiguration(
    okta.WithOrgUrl("https://{yourOktaDomain}"),
    okta.WithToken("{apiToken}"),
  )
  if err != nil {
    fmt.Printf("Error: %v\n", err)
  }
  client := okta.NewAPIClient(config)
}
```

Hard-coding the Okta domain and API token works for quick tests, but for real
projects you should use a more secure way of storing these values (such as
environment variables). This library supports a few different configuration
sources, covered in the [configuration reference](#configuration-reference)
  section.

## Usage guide

These examples will help you understand how to use this library. You can also
browse the full [API reference documentation][sdkapiref].

Once you initialize a `client`, you can call methods to make requests to the
Okta API. Most methods are grouped by the API endpoint they belong to. For
example, methods that call the [Users
API](https://developer.okta.com/docs/api/resources/users) are organized under
`client.User`.

## Caching

In the default configuration the client utilizes a memory cache that has a time
to live on its cached values. See [Configuration Setter
Object](#configuration-setter-object)  `WithCache(cache bool)`,
`WithCacheTtl(i int32)`, and `WithCacheTti(i int32)`.  This helps to
keep HTTP requests to the Okta API at a minimum. In the case where the client
needs to be certain it is accessing recent data; for instance, list items,
delete an item, then list items again; be sure to make use of the refresh next
facility to clear the request cache. See [Refreshing Cache for Specific
Call](#refreshing-cache-for-specific-call). To completely disable the request
memory cache configure the client with `WithCache(false)`.

NOTE: Regardless of cache manager, Access Tokens from OAuth requests are always
cached.

## Connection Retry / Rate Limiting

By default this SDK retries requests that are returned with a 429 exception. To
disable this functionality set `OKTA_CLIENT_REQUEST_TIMEOUT` and
`OKTA_CLIENT_RATELIMIT_MAXRETRIES` to `0`.

Setting only one of the values to zero disables that check. Meaning, by
default, four retry attempts will be made. If you set
`OKTA_CLIENT_REQUEST_TIMEOUT` to 45 seconds and
`OKTA_CLIENT_RATELIMIT_MAXRETRIES` to `0`. This SDK will continue to retry
indefinitely for 45 seconds. If both values are non zero, this SDK attempts to
retry until either of the conditions are met (not both).

We use the Date header from the server to calculate the delta, as it's more
reliable than system time. But always add 1 second to account for some clock
skew in our service:

```go
backoff_seconds = header['X-Rate-Limit-Reset'] - header['Date'] + 1s
```

If the `backoff_seconds` calculation exceeds the request timeout, the initial
429 response will be allowed through without additional attempts.

When creating your client, you can pass in these settings like you would with
any other configuration.

```go
import (
	"fmt"
	"context"
	"github.com/okta/okta-sdk-golang/v6/okta"
)

func main() {
  config, err := okta.NewConfiguration(
    okta.WithOrgUrl("https://{yourOktaDomain}"),
    okta.WithToken("{apiToken}"),
    okta.WithRequestTimeout(45),
    okta.WithRateLimitMaxRetries(3),
  )
  if err != nil {
    fmt.Printf("Error: %v\n", err)
  }
  client := okta.NewAPIClient(config)
}
```

### Authenticate a User

This library should only be used with the Okta management API. To call the
[Authentication API](https://developer.okta.com/docs/api/resources/authn), you
should construct your own HTTP requests.

### Get a User

```go
import (
	"fmt"
	"context"
	"github.com/okta/okta-sdk-golang/v6/okta"
)

func main() {
  config, err := okta.NewConfiguration(
    okta.WithOrgUrl("https://{yourOktaDomain}"),
    okta.WithToken("{apiToken}"),
  )
  if err != nil {
    fmt.Printf("Error: %v\n", err)
  }
  client := okta.NewAPIClient(config)

  user, resp, err := client.UserAPI.GetUser(client.GetConfig().Context, "{UserId|Username|Email}").Execute()
  if err != nil {
      fmt.Printf("Error Getting User: %v\n", err)
  }
  fmt.Printf("User: %+v\n Response: %+v\n\n",user, resp)
}
```

### List all Users

```go
import (
	"fmt"
	"context"
	"github.com/okta/okta-sdk-golang/v6/okta"
)

func main() {
  config, err := okta.NewConfiguration(
    okta.WithOrgUrl("https://{yourOktaDomain}"),
    okta.WithToken("{apiToken}"),
  )
  if err != nil {
    fmt.Printf("Error: %v\n", err)
  }
  client := okta.NewAPIClient(config)

  users, resp, err := client.UserAPI.ListUsers(client.GetConfig().Context).Execute()
  if err != nil {
    fmt.Printf("Error Getting Users: %v\n", err)
  }
  fmt.Printf("Users: %+v\n Response: %+v\n\n",users, resp)
  for index, user := range users {
    fmt.Printf("User %d: %+v\n", index, user)
  }
}
```

### Filter or search for Users

```go
import (
	"fmt"
	"context"
	"github.com/okta/okta-sdk-golang/v6/okta"
)

func main() {
  config, err := okta.NewConfiguration(
    okta.WithOrgUrl("https://{yourOktaDomain}"),
    okta.WithToken("{apiToken}"),
  )
  if err != nil {
    fmt.Printf("Error: %v\n", err)
  }
  client := okta.NewAPIClient(config)

  users, resp, err := client.UserAPI.ListUsers(client.GetConfig().Context).Filter("status eq \"ACTIVE\"").Execute()
  if err != nil {
    fmt.Printf("Error Getting Users: %v\n", err)
  }

	fmt.Printf("Filtered Users: %+v\n Response: %+v\n\n", filteredUsers, resp)

	for index, user := range filteredUsers {
		fmt.Printf("User %d: %+v\n", index, user)
	}
}
```

### Create a User

```go
import (
	"fmt"
	"context"
	"github.com/okta/okta-sdk-golang/v6/okta"
)

func main() {
  config, err := okta.NewConfiguration(
    okta.WithOrgUrl("https://{yourOktaDomain}"),
    okta.WithToken("{apiToken}"),
  )
  if err != nil {
    fmt.Printf("Error: %v\n", err)
  }
  client := okta.NewAPIClient(config)

  password := &okta.PasswordCredential{
		Value: "Abcd1234!",
	}

	userCredentials := &okta.UserCredentials{
		Password: password,
	}

	profile := okta.UserProfile{}
	profile["firstName"] = "Ben"
	profile["lastName"] = "Solo"
	profile["email"] = "ben-solo@example.com"
	profile["login"] = "ben-solo@example.com"

	createUserRequest := okta.CreateUserRequest{
		Credentials: userCredentials,
		Profile: &profile,
	}

  users, resp, err := client.UserAPI.CreateUser(client.GetConfig().Context).Body(createUserRequest).Activate(true).Execute()
  if err != nil {
    fmt.Printf("Error Creating Users: %v\n", err)
  }

	fmt.Printf("User: %+v\n Response: %+v\n\n", user, resp)
}
```

### Update a User

```go
import (
	"fmt"
	"context"
	"github.com/okta/okta-sdk-golang/v6/okta"
)

func main() {
  config, err := okta.NewConfiguration(
    okta.WithOrgUrl("https://{yourOktaDomain}"),
    okta.WithToken("{apiToken}"),
  )
  if err != nil {
    fmt.Printf("Error: %v\n", err)
  }
  client := okta.NewAPIClient(config)

  userToUpdate, resp, err := client.UserAPI.GetUser(client.GetConfig().Context, "{userId}")
	if err != nil {
		fmt.Printf("Error Getting User to update: %v\n", err)
	}

	fmt.Printf("User to update: %+v\n Response: %+v\n\n", userToUpdate, resp)

	newProfile := *userToUpdate.Profile
	newProfile["nickName"] = "Kylo Ren"
	updateUser := &okta.UpdateUserRequest{
		Profile: &newProfile,
	}

	updatedUser, resp, err := client.UserAPI.UpdateUser(client.GetConfig().Context, userToUpdate.Id).Body(updateUser).Execute()
	if err != nil {
		fmt.Printf("Error updating user: %v\n", err)
	}

	fmt.Printf("Updated User: %+v\n Response: %+v\n\n", updatedUser, resp)
}
```

### Get and set custom attributes

Custom attributes must first be defined in the Okta profile editor. Then, you
can work with custom attributes on a user the same as any other profile attribute

### Delete a User

To delete a user permanently. Only users that have a `DEPROVISIONED` status can be deleted.`

```go
import (
	"fmt"
	"context"
	"github.com/okta/okta-sdk-golang/v6/okta"
)

func main() {
  config, err := okta.NewConfiguration(
    okta.WithOrgUrl("https://{yourOktaDomain}"),
    okta.WithToken("{apiToken}"),
  )
  if err != nil {
    fmt.Printf("Error: %v\n", err)
  }
  client := okta.NewAPIClient(config)

	resp, err = client.UserAPI.DeleteUser(client.GetConfig().Context, "00u14ffhw5szVqide0h8").Execute()
	if err != nil {
		fmt.Printf("Error deleting user: %v\n", err)
	}

	fmt.Printf("Response: %+v\n\n", resp)
}
```

### Add a Group

```go
import (
	"fmt"
	"context"
	"github.com/okta/okta-sdk-golang/v6/okta"
)

func main() {
  config, err := okta.NewConfiguration(
    okta.WithOrgUrl("https://{yourOktaDomain}"),
    okta.WithToken("{apiToken}"),
  )
  if err != nil {
    fmt.Printf("Error: %v\n", err)
  }

  client := okta.NewAPIClient(config)

  grpProfile := okta.NewOktaUserGroupProfile()
	grpProfile.SetDescription("This is a group of developers")
	grpProfile.SetName("Developers")
	addGrpRequest := okta.AddGroupRequest{
		Profile: grpProfile,
	}
	// Add the group using the GroupAPI
	group, resp, err := client.GroupAPI.AddGroup(client.GetConfig().Context).Group(addGrpRequest).Execute()
	if err != nil {
		fmt.Printf("Error creating group: %v\n", err)
	}

	fmt.Printf("Created Group: %+v\n Response: %+v\n\n", group, resp)
}
```

### Add a User to a Group

```go
import (
	"fmt"
	"context"
	"github.com/okta/okta-sdk-golang/v6/okta"
)

func main() {
  config, err := okta.NewConfiguration(
    okta.WithOrgUrl("https://{yourOktaDomain}"),
    okta.WithToken("{apiToken}"),
  )
  if err != nil {
    fmt.Printf("Error: %v\n", err)
  }
  client := okta.NewAPIClient(config)

  resp, err := client.GroupAPI.AssignUserToGroup(client.GetConfig().Context, "{groupId}", "{userId}").Execute()
	if err != nil {
		fmt.Printf("Error adding user to group: %v\n", err)
	}

	fmt.Printf("Response: %+v\n\n", resp)
}
```

### List all Applications

```go
import (
	"fmt"
	"context"
	"github.com/okta/okta-sdk-golang/v6/okta"
)

func main() {
  config, err := okta.NewConfiguration(
    okta.WithOrgUrl("https://{yourOktaDomain}"),
    okta.WithToken("{apiToken}"),
  )
  if err != nil {
    fmt.Printf("Error: %v\n", err)
  }
  client := okta.NewAPIClient(config)

  applicationList, resp, err := client.ApplicationAPI.ListApplications(client.GetConfig().Context).Execute()
	if err != nil {
		fmt.Printf("Error listing applications: %v\n", err)
	}

	fmt.Printf("ApplicationList: %+v\n Response: %+v\n\n", applicationList, resp)
}
```

### Get an Application

```go
import (
	"fmt"
	"context"
	"github.com/okta/okta-sdk-golang/v6/okta"
)

func main() {
  config, err := okta.NewConfiguration(
    okta.WithOrgUrl("https://{yourOktaDomain}"),
    okta.WithToken("{apiToken}"),
  )
  if err != nil {
    fmt.Printf("Error: %v\n", err)
  }
  client := okta.NewAPIClient(config)

  application, resp, err := client.ApplicationAPI.GetApplication(client.GetConfig().Context, "0oaswjmkbtlpBDWpu0h7").Execute()
	if err != nil {
		fmt.Printf("Error getting application: %v\n", err)
	}

	fmt.Printf("Application: %+v\n Response: %+v\n\n", application, resp)

}
```

### Create an OpenID Connect(OIDC) Application

```go
import (
	"fmt"
	"context"
	"github.com/okta/okta-sdk-golang/v6/okta"
)

func main() {
  config, err := okta.NewConfiguration(
    okta.WithOrgUrl("https://{yourOktaDomain}"),
    okta.WithToken("{apiToken}"),
  )
  if err != nil {
    fmt.Printf("Error: %v\n", err)
  }
  client := okta.NewAPIClient(config)

  settingClient := okta.NewOpenIdConnectApplicationSettingsClient([]string{"grantTypes"})
  settingClient.SetClientUri("https://example.com/client")
  settingClient.SetLogoUri("https://example.com/assets/images/logo-new.png")
  settingClient.SetResponseTypes([]string{"token", "id_token", "code"})
  settingClient.SetRedirectUris([]string{"https://example.com/oauth2/callback", "myapp://callback"})
  settingClient.SetPostLogoutRedirectUris([]string{"https://example.com/postlogout", "myapp://postlogoutcallback"})
  settingClient.SetGrantTypes([]string{"implicit", "authorization_code"})
  settingClient.SetApplicationType("native")
  settingClient.SetTosUri("https://example.com/client/tos")
  settingClient.SetPolicyUri("https://example.com/client/policy")
  setting := okta.NewOpenIdConnectApplicationSettings()
  setting.SetOauthClient(*settingClient)
  credClient := okta.NewApplicationCredentialsOAuthClient()
  credClient.SetTokenEndpointAuthMethod("client_secret_post")
  credClient.SetClientId("id")
  credClient.SetAutoKeyRotation(true)
  credentials := okta.NewOAuthApplicationCredentials()
  credentials.SetOauthClient(*credClient)
  res := okta.OpenIdConnectApplication{}
  res.SetSettings(*setting)
  res.SetCredentials(*credentials)
  res.SetName("oidc_client")
  res.SetSignOnMode("OPENID_CONNECT")
  res.SetLabel("label")

  application, resp, err := client.ApplicationAPI.CreateApplication(client.GetConfig().Context).Application(okta.ListApplications200ResponseInner{OpenIdConnectApplication: &res}).Execute()
  if err != nil {
      fmt.Printf("Error creating application: %v\n", err)
  }

  fmt.Printf("Application: %+v\n Response: %+v\n\n", application, resp)
}
```

## Configuration reference

This library looks for configuration in the following sources:

1. An `okta.yaml` file in a `.okta` folder in the current user's home directory
   (`~/.okta/okta.yaml` or `%userprofile\.okta\okta.yaml`)
2. A `.okta.yaml` file in the application or project's root directory
3. Environment variables
4. Configuration explicitly passed to the constructor (see the example in
   [Getting started](#getting-started))

Higher numbers win. In other words, configuration passed via the constructor
will override configuration found in environment variables, which will override
configuration in `okta.yaml` (if any), and so on.

### YAML configuration

When you use an API Token instead of OAuth 2.0 the full YAML configuration
looks like:

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
    privateKeyId: "{JWK key id (kid}" # needed if Okta service application has more then a single JWK registered
    requestTimeout: 0 # seconds
    rateLimit:
      maxRetries: 4
```

### Environment variables

Each one of the configuration values above can be turned into an environment
variable name with the `_` (underscore) character:

* `OKTA_CLIENT_CONNECTIONTIMEOUT`
* `OKTA_CLIENT_TOKEN`
* and so on

### Configuration Setter Object

The client is configured with a configuration setter object passed to the `NewClient` function.

| function | description |
|----------|-------------|
| WithCache(cache bool) | Use request memory cache |
| WithCacheManager(cacheManager cache.Cache) | Use custom cache object that implements the `cache.Cache` interface |
| WithCacheTtl(i int32) | Cache time to live in seconds |
| WithCacheTti(i int32) | Cache clean up interval in seconds |
| WithConnectionTimeout(i int64) | HTTP connection timeout in seconds |
| WithProxyPort(i int32) | HTTP proxy port |
| WithProxyHost(host string) | HTTP proxy host |
| WithProxyUsername(username string) | HTTP proxy username |
| WithProxyPassword(pass string) | HTTP proxy password |
| WithOrgUrl(url string) | Okta organization URL |
| WithToken(token string) | Okta API token |
| WithUserAgentExtra(userAgent string) | Append additional information to the HTTP User-Agent |
| WithHttpClientPtr(httpClient *http.Client) | pointer to custom net/http client |
| WithTestingDisableHttpsCheck(httpsCheck bool) | Disable net/http SSL checks |
| WithRequestTimeout(requestTimeout int64) | HTTP request time out in seconds |
| WithRateLimitMaxRetries(maxRetries int32) | Number of request retries when http request times out |
| WithRateLimitMaxBackOff(maxBackoff int64) | Max amount of time to wait on request back off |
| WithAuthorizationMode(authzMode string) | Okta API auth mode, `SSWS` (Okta based), `PrivateKey` (OAuth app based) or `JWT` (OAuth app based) |
| WithScopes(scopes []string) | Okta API app scopes |
| WithPrivateKey(privateKey string) | Private key value |
| WithPrivateKeyId(privateKeyId string) | Private key id (kid) value |
| WithPrivateKeySigner(signer jose.Signer) | Custom private key signer implementing the `jose.Signer` interface |

### Okta Client Base Configuration

The Okta Client's base configuration starts at

| config setting |
|----------------|
| WithConnectionTimeout(60) |
| WithCache(true) |
| WithCacheTtl(300) |
| WithCacheTti(300),
| WithUserAgentExtra("") |
| WithTestingDisableHttpsCheck(false) |
| WithRequestTimeout(0) |
| WithRateLimitMaxBackOff(30) |
| WithRateLimitMaxRetries(2) |
| WithAuthorizationMode("SSWS") |

### Context

Every method that calls the API now has the ability to pass `context.Context`
to it as the first parameter. If you do not have a context or do not know which
context to use, you can pass `context.TODO()` to the methods.

### Method changes

We have spent time during this update making sure we become a little more
uniform with naming of methods. This will require you to update some of your
calls to the SDK with the new names.

All methods now specify the `Accept` and `Content-Type` headers when creating a
new request. This allows for future use of the SDK to handle multiple `Accept`
types.

### OAuth 2.0 With Private Key

Okta allows you to interact with Okta APIs using scoped OAuth 2.0 access
tokens. Each access token enables the bearer to perform specific actions on
specific Okta endpoints, with that ability controlled by which scopes the
access token contains.

Access Tokens are always cached and respect the `expires_in` value of an access
token response.

This SDK supports this feature only for service-to-service applications. Check
out [our
guides](https://developer.okta.com/docs/guides/implement-oauth-for-okta/overview/)
to learn more about how to register a new service application using a private
and public key pair.

Key pairs can only be used to manage certain Okta endpoints, using the scopes listed
[here](https://developer.okta.com/docs/guides/implement-oauth-for-okta/main/#scopes-and-supported-endpoints).
To manage Okta endpoints that are not listed
[here](https://developer.okta.com/docs/guides/implement-oauth-for-okta/main/#scopes-and-supported-endpoints),
an API Token belonging to a user with appropriate permissions must be used instead.

When using a keypair, you won't need an API Token because the SDK will
request an access token for you. In order to use OAuth 2.0, construct a client
instance by passing the following parameters:

OAuth 2.0 flow now supports [DPoP](https://developer.okta.com/docs/guides/dpop/main/)

```go
config, err := okta.NewConfiguration(
  okta.WithOrgUrl("https://{yourOktaDomain}"),
  okta.WithAuthorizationMode("PrivateKey"),
  okta.WithClientId("{{clientId}}"),
  okta.WithScopes(([]string{"okta.users.manage"})),
  okta.WithPrivateKey({{PEM PRIVATE KEY BLOCK}}), //when pasting blocks, use backticks and remove all space at beginning of each line.
  okta.WithPrivateKeyId("{{private key id}}"),    //needed if Okta service application has more then a single JWK registered
)
if err != nil {
  fmt.Printf("Error: %v\n", err)
}
client := okta.NewAPIClient(config)
```

#### Example

Let's say the need is to authenticate a script that will run in a pipeline (or
any other automated way), and instead of using an API Token (that is bound
to a user) the goal is to use a service app.

*A public/private key pair is required to do so.*

These are the requirements:

  - a public/private key par in JWT format
    ([reference](https://developer.okta.com/docs/guides/implement-oauth-for-okta-serviceapp/create-publicprivate-keypair/))
  - a service app that uses the created key
    ([reference](https://developer.okta.com/docs/guides/implement-oauth-for-okta-serviceapp/create-serviceapp-grantscopes/))
  - store the private key in a PEM format
    ([reference](https://www.npmjs.com/package/pem-jwk))

To store the PEM formatted key with new lines in a JSON file, the multiple
lines need to be one-line formatted by joining them with the "\n" character.
The `awk` command makes this formatting quick and precise:

```bash
awk 'NF {sub(/\r/, ""); printf "%s\\n",$0;}' private.pem
```

##### Steps

The [mkjwk](https://mkjwk.org/) can be used to create your keys. If generating
production keys, only use `mkjwk` running locally after you have audited their
code.

Save the three files, for example named as *public-key*,
*public-private-keypair* and *public-private-keypair-set*.

Now create the PEM formatted private key.
`[pem-jwk](https://www.npmjs.com/package/pem-jwk)` can be utilized to do the
PEM formatting. Be sure to audit the `pem-jwk` code before trusting it with
production values.

```bash
pem-jwk public-private-keypair > private.pem
```

Create the service app following Okta's guide ["Create a service app and grant
scopes > Create a service
app"](https://developer.okta.com/docs/guides/implement-oauth-for-okta-serviceapp/create-serviceapp-grantscopes/#create-a-service-app),
using *public-private-keypair-set*.

Use the Okta web console to grant the scopes as usual.

To complete our example the PEM formatted private key will be stored in a JSON
file so your app can read it at run time.  If this is the PEM formatted key:

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

... then the value that must be added to the JSON file is:

```bash
----BEGIN RSA PRIVATE KEY-----\nMIIEowIBAAKCAQEAmyX8wdrHK1ycOMeXNg3NOMQvebnfQp+3L5OaaiX16/+tLbwb\nJTZDYh0EXLySMVsduRxC/1PQdPuI6x50TdkoB3C4JMuU968uJqkFp7fXXy5SMAej\nHAyF67cY51dx15ztvakRNJPhhI5WaC20RfR/eow0IH5lGI3czcvTCChGau5qLue3\nHqNDYFY+U3xhOlavSDdtmuxpIFsDycn/OjYjsV4lzyRrOArqtVV/kXHKx04T6A1x\nSc99999999999999999999999999999999999999999999999999EGekHlUAIUpw\nTqzdddddddddddddddddddddddddddddddddddddddddddddddddd5rNLLe5C7p/\nLMta1rzm5TPYwazIbiMkFLAW02ToNAs9LGgDP+VRCZskl6+LuaA5XGabpi09ka7x\nvJ88888888888888888888888888888888888888888888888888JDmIfMSU1tEw\nHmthd7qcwrx29ectcGHyQaX6iyYlQiBKCto+VwPrUq/qDCPYMIyqCTxAGTPLWQAz\nTqzdddddddddddddddddddddddddddddddddddddddddddddddddd5rNLLe5C7p/\nHpDowa9Q+CBO5CEcPW4w9gsCgYA68a+82YtsiyYjdwSzscKIrw4ht3qAZQhGO4Id\nH70kN1CkhHUNFf5UuRHJJ+s3BKlawBUwZaKD5KQ+lpnmrwqfArWu+3HNpX3LIPs5\nvJ88888888888888888888888888888888888888888888888888JDmIfMSU1tEw\ntATmAQKBgCKBkyHmEaS8tEAXcRI26oHOwTZAj6tJp2ODrEcrWtT0bR8wjaEdASdM\nTqzdddddddddddddddddddddddddddddddddddddddddddddddddd5rNLLe5C7p/\n4b+xIYHz8dxmWPsZ7C4WbW9pw3Wn1Du/uvImwt0f4Jp6IPZM9vltXz0Dh7Yv5/SE\nvJ88888888888888888888888888888888888888888888888888JDmIfMSU1tEw\ntATmAQKBgCKBkyHmEaS8tEAXcRI26oHOwTZAj6tJp2ODrEcrWtT0bR8wjaEdASdM\nTqzdddddddddddddddddddddddddddddddddddddddddddddddddd5rNLLe5C7p/\nvJ88888888888888888888888888888888888888888888888888JDmIfMSU1tEw\ntATmAQKBgCKBkyHmEaS8tEAXcRI26oHOwTZAj6tJp2ODrEcrWtT0bR8wjaEdASdM\nDKctGohIQ/ujUD9wzSvlaSZjBcKWw27yN0HiEBn+whKmO76PT7NFAQv/TG8ou3NE\nftlYhgBkwRwRfk7lEvmaTvJugd5g1E/9DAXTajlRYdohGubVz+2G\n-----END RSA PRIVATE KEY-----
```

This way, if a JSON file like this one is created:

```yaml
{
    "Oktadomain" : "yourorg.okta.com",
    "clientId" : "yourclientid",
    "privateKey" : "----BEGIN RSA PRIVATE KEY-----\nMIIEowIBAAKCAQEAmyX8wdrHK1ycOMeXNg3NOMQvebnfQp+3L5OaaiX16/+tLbwb\nJTZDYh0EXLySMVsduRxC/1PQdPuI6x50TdkoB3C4JMuU968uJqkFp7fXXy5SMAej\nHAyF67cY51dx15ztvakRNJPhhI5WaC20RfR/eow0IH5lGI3czcvTCChGau5qLue3\nHqNDYFY+U3xhOlavSDdtmuxpIFsDycn/OjYjsV4lzyRrOArqtVV/kXHKx04T6A1x\nSc99999999999999999999999999999999999999999999999999EGekHlUAIUpw\nTqzdddddddddddddddddddddddddddddddddddddddddddddddddd5rNLLe5C7p/\nLMta1rzm5TPYwazIbiMkFLAW02ToNAs9LGgDP+VRCZskl6+LuaA5XGabpi09ka7x\nvJ88888888888888888888888888888888888888888888888888JDmIfMSU1tEw\nHmthd7qcwrx29ectcGHyQaX6iyYlQiBKCto+VwPrUq/qDCPYMIyqCTxAGTPLWQAz\nTqzdddddddddddddddddddddddddddddddddddddddddddddddddd5rNLLe5C7p/\nHpDowa9Q+CBO5CEcPW4w9gsCgYA68a+82YtsiyYjdwSzscKIrw4ht3qAZQhGO4Id\nH70kN1CkhHUNFf5UuRHJJ+s3BKlawBUwZaKD5KQ+lpnmrwqfArWu+3HNpX3LIPs5\nvJ88888888888888888888888888888888888888888888888888JDmIfMSU1tEw\ntATmAQKBgCKBkyHmEaS8tEAXcRI26oHOwTZAj6tJp2ODrEcrWtT0bR8wjaEdASdM\nTqzdddddddddddddddddddddddddddddddddddddddddddddddddd5rNLLe5C7p/\n4b+xIYHz8dxmWPsZ7C4WbW9pw3Wn1Du/uvImwt0f4Jp6IPZM9vltXz0Dh7Yv5/SE\nvJ88888888888888888888888888888888888888888888888888JDmIfMSU1tEw\ntATmAQKBgCKBkyHmEaS8tEAXcRI26oHOwTZAj6tJp2ODrEcrWtT0bR8wjaEdASdM\nTqzdddddddddddddddddddddddddddddddddddddddddddddddddd5rNLLe5C7p/\nvJ88888888888888888888888888888888888888888888888888JDmIfMSU1tEw\ntATmAQKBgCKBkyHmEaS8tEAXcRI26oHOwTZAj6tJp2ODrEcrWtT0bR8wjaEdASdM\nDKctGohIQ/ujUD9wzSvlaSZjBcKWw27yN0HiEBn+whKmO76PT7NFAQv/TG8ou3NE\nftlYhgBkwRwRfk7lEvmaTvJugd5g1E/9DAXTajlRYdohGubVz+2G\n-----END RSA PRIVATE KEY-----"
}
```

The file can be read from Go and used directly in the client creation:

```go
config, err := okta.NewConfiguration(
  okta.WithOrgUrl("https://{yourOktaDomain}"),
  okta.WithAuthorizationMode("PrivateKey"),
  okta.WithClientId("{client_id}"),
  okta.WithScopes([]string{"{scopes}"}),
  okta.WithPrivateKey("{private_key}"),
  okta.WithPrivateKeyId("{private key id}"), //needed if Okta service application has more then a single JWK registered
)
if err != nil {
  fmt.Printf("Error: %v\n", err)
}
client := okta.NewAPIClient(config)
```

### OAuth 2.0 With JWT Key
Okta allows you to interact with Okta APIs using scoped OAuth 2.0 access
tokens. Each access token enables the bearer to perform specific actions on
specific Okta endpoints, with that ability controlled by which scopes the
access token contains.

Access Tokens are always cached and respect the `expires_in` value of an access
token response.

This SDK supports this feature only for service-to-service applications. Check
out [our
guides](https://developer.okta.com/docs/guides/implement-oauth-for-okta/overview/)
to learn more about how to register a new service application using a private
and public key pair. Otherwise, follow the example steps at the end of this
topic.

When using this approach you won't need an API Token because the SDK will
request an access token for you. In order to use OAuth 2.0, construct a client
instance by passing the following parameters:

```go
config, err := okta.NewConfiguration(
  okta.WithOrgUrl("https://{yourOktaDomain}"),
  okta.WithAuthorizationMode("JWT"),
  okta.WithClientAssertion("{{clientAssertion}}"),
  okta.WithScopes(([]string{"okta.users.manage"})),
)
if err != nil {
  fmt.Printf("Error: %v\n", err)
}
client := okta.NewAPIClient(config)
```

This is very similar to PrivateKey Authorization Mode with a caveat, instead of providing public/privatekey pair, you can use a pre-signed JWT instead

### OAuth 2.0 With Bearer Token

Okta SDK supports authorization using a `Bearer` token. A bearer token is an
ephemeral token issued by Okta via supported Okta apps. Exchanging provisioned
credentials for a bearer token may involve external dependencies that are out
of scope for the SDK to support natively.

#### Bearer Token Scope

Bearer tokens are scoped to an application and not to the greater organization
as is enabled when the SDK is initialized with an Okta SSWS token or an OAuth
private key. Therefore a bearer token will not have the permissions to perform
all of the API calls for organization management through the SDK. The scope of
the bearer token is determined by the scope configured on the Okta app __and__
the scope requested during the authorization. Care should be taken by the
implementers to configure the Okta app with the necessary scopes for the
integration.

#### Implementation Steps

1. Create an Okta app based on your requirements. The following table lists the apps types that issue `Bearer` tokens.

    | Sign-in method        | Application Type        | Details                                                                                                                                                                                                                                         |
    |-----------------------|-------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
    | OIDC - OpenID Connect | Web Application         | [Authorization code flow with client secret](https://developer.okta.com/docs/guides/implement-grant-type/authcode/main/)                                                                                                                        |
    | OIDC - OpenID Connect | Single-Page Application | [Authorization code flow with PKCE](https://developer.okta.com/docs/guides/implement-grant-type/authcodepkce/main/)                                                                                                                             |
    | OIDC - OpenID Connect | Native Application      | [Authorization code flow with client secret](https://developer.okta.com/docs/guides/implement-grant-type/authcode/main/) or [Authorization code flow with PKCE](https://developer.okta.com/docs/guides/implement-grant-type/authcodepkce/main/) |
    | API Services          | N/A                     | Natively supported by SDK using `PrivateKey` Authorization mode                                                                                                                                                                                 |

2. Make a call to the [Org Authorization Server](https://developer.okta.com/docs/concepts/auth-servers/#org-authorization-server) endpoint to get the authorization code.
3. Exchange authorization code for a `Bearer` token.
4. Instantiate and use Okta client with `Bearer` token.
    ```go
    config, err := okta.NewConfiguration(
        okta.WithOrgUrl("https://{yourOktaDomain}"),
        okta.WithAuthorizationMode("Bearer"),
        okta.WithClientId("{client_id}"),
        okta.WithToken("{token}"),
    )
    if err != nil {
      fmt.Printf("Error: %v\n", err)
    }
    client := okta.NewAPIClient(config)
    ```
    A bearer token is scoped implicitly, so there is no need to provide
    `okta.WithScopes()` config setter method when initializing the Okta client.

### Extending the Client

When calling `okta.NewConfiguration()` we allow for you to pass custom instances of
`http.Client` and `cache.Cache`.

```go
myClient := &http.Client{}

myCache := NewCustomCacheDriver()

config, err := okta.NewConfiguration(
  okta.WithOrgUrl("https://{yourOktaDomain}"),
  okta.WithToken("{apiToken}"),
  okta.WithHttpClient(myClient),
  okta.WithCacheManager(myCache)
)
if err != nil {
  fmt.Printf("Error: %v\n", err)
}
client := okta.NewAPIClient(config)
```


### Extending or Creating New Cache Manager

You can create a custom cache driver by implementing `cache.Cache`

```go
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

### Pagination

If your request comes back with more than the default or set limit, you can
request the next page.

Example of listing users 1 at a time:

```go
users, resp, err := client.UserAPI.ListUsers(ctx, query)
// Do something with your users until you run out of users to iterate.
if resp.HasNextPage() {
  var nextUserSet []okta.User
  resp, err = resp.Next(&user2)
}
```

## Contributing

We're happy to accept contributions and PRs! Please see the [contribution
guide](CONTRIBUTING.md) to understand how to structure a contribution.

[devforum]: https://devforum.okta.com/
[sdkapiref]: https://godoc.org/github.com/okta/okta-sdk-golang/v6/okta
[lang-landing]: https://developer.okta.com/code/go/
[github-issues]: /okta/okta-sdk-golang/issues
[github-releases]: /okta/okta-sdk-golang/releases
