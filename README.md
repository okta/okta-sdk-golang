[<img src="https://aws1.discourse-cdn.com/standard14/uploads/oktadev/original/1X/0c6402653dfb70edc661d4976a43a46f33e5e919.png" align="right" width="256px"/>](https://devforum.okta.com/)
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

The SDK is compatible with Go version 1.12.x and up. For SDK v2 and above, you
must use [Go Modules](https://blog.golang.org/using-go-modules) to install the
SDK.

### Install current release

To install the Okta Golang SDK in your project:
  - Create a module file by running `go mod init`
    - You can skip this step if you already use `go mod`
  - Run `go get github.com/okta/okta-sdk-golang/v2@latest`. This will add
    the SDK to your `go.mod` file.
  - Import the package in your project with `import
   "github.com/okta/okta-sdk-golang/v2/okta"`

### Installing legacy version

Although we do not suggest using the 1.x version of the SDK, you can still use
it. Version 1.x is *retiring* and will not be supported past March 4, 2021. It
will likely remain working after that date, but you should make a plan to
migrate to the new 2.x version.

You can install v1 of the SDK by running `go get
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
	"github.com/okta/okta-sdk-golang/v2/okta"
)

func main() {
  ctx, client, err := okta.NewClient(
    context.TODO(),
    okta.WithOrgUrl("https://{yourOktaDomain}"),
    okta.WithToken("{apiToken}"),
  )

  if err != nil {
    fmt.Errorf("Error: %v", err)
  }

  fmt.Printf("Context: %+v\n Client: %+v\n",ctx, client)
)
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


## Connection Retry / Rate Limiting

By default this SDK retries requests that are returned with a 429 exception. To
disable this functionality set `OKTA_CLIENT_REQUESTTIMEOUT` and
`OKTA_CLIENT_RATELIMIT_MAXRETRIES` to `0`.

Setting only one of the values to zero disables that check. Meaning, by
default, four retry attempts will be made. If you set
`OKTA_CLIENT_REQUESTTIMEOUT` to 45 seconds and
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
	"github.com/okta/okta-sdk-golang/v2/okta"
)

func main() {
  ctx, client, err := okta.NewClient(
    context,
    okta.WithOrgUrl("https://{yourOktaDomain}"),
    okta.WithToken("{apiToken}"),
    okta.WithRequestTimeout(45),
    okta.WithRateLimitMaxRetries(3),
  )

  if err != nil {
    fmt.Errorf("Error: %v", err)
  }

  fmt.Printf("Context: %+v\n Client: %+v\n",ctx, client)
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
	"github.com/okta/okta-sdk-golang/v2/okta"
)

func main() {
  ctx, client, err := okta.NewClient(
    context.TODO(),
    okta.WithOrgUrl("https://{yourOktaDomain}"),
    okta.WithToken("{apiToken}"),
  )

  if err != nil {
    fmt.Errorf("Error: %v", err)
  }

  fmt.Printf("Context: %+v\n Client: %+v\n",ctx, client)

  user, resp, err := client.User.GetUser(ctx, "{UserId|Username|Email}")
  if err != nil {
    fmt.Errorf("Error Getting User: %v", err)
  }
  fmt.Printf("User: %+v\n Response: %+v\n\n",user, resp)
}
```

### List all Users

```go
import (
	"fmt"
	"context"
	"github.com/okta/okta-sdk-golang/v2/okta"
)

func main() {
  ctx, client, err := okta.NewClient(
    context.TODO(),
    okta.WithOrgUrl("https://{yourOktaDomain}"),
    okta.WithToken("{apiToken}"),
  )

  if err != nil {
    fmt.Errorf("Error: %v", err)
  }

  fmt.Printf("Context: %+v\n Client: %+v\n",ctx, client)

  users, resp, err := client.User.ListUsers(ctx, nil)
  if err != nil {
    fmt.Errorf("Error Getting Users: %v", err)
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
	"github.com/okta/okta-sdk-golang/v2/okta"
	"github.com/okta/okta-sdk-golang/v2/okta/query"
)

func main() {
  ctx, client, err := okta.NewClient(
    context.TODO(),
    okta.WithOrgUrl("https://{yourOktaDomain}"),
    okta.WithToken("{apiToken}"),
  )

  if err != nil {
    fmt.Errorf("Error: %v", err)
  }

  fmt.Printf("Context: %+v\n Client: %+v\n",ctx, client)

  filter := query.NewQueryParams(query.WithFilter("status eq \"ACTIVE\""))

	filteredUsers, resp, err := client.User.ListUsers(ctx, filter)
	if err != nil {
		fmt.Errorf("Error Getting Users: %v", err)
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
	"github.com/okta/okta-sdk-golang/v2/okta"
	"github.com/okta/okta-sdk-golang/v2/okta/query"
)

func main() {
  ctx, client, err := okta.NewClient(
    context.TODO(),
    okta.WithOrgUrl("https://{yourOktaDomain}"),
    okta.WithToken("{apiToken}"),
  )

  if err != nil {
    fmt.Errorf("Error: %v", err)
  }

  fmt.Printf("Context: %+v\n Client: %+v\n",ctx, client)

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


	// activate user on create
	queryParam := query.NewQueryParams(query.WithActivate(true))


	user, resp, err := client.User.CreateUser(ctx, createUserRequest, queryParam)
	if err != nil {
		fmt.Errorf("Error Creating User: %v", err)
	}

	fmt.Printf("User: %+v\n Response: %+v\n\n", user, resp)
}
```

### Update a User

```go
import (
	"fmt"
	"context"
	"github.com/okta/okta-sdk-golang/v2/okta"
	"github.com/okta/okta-sdk-golang/v2/okta/query"
)

func main() {
  ctx, client, err := okta.NewClient(
    context.TODO(),
    okta.WithOrgUrl("https://{yourOktaDomain}"),
    okta.WithToken("{apiToken}"),
  )

  if err != nil {
    fmt.Errorf("Error: %v", err)
  }

  fmt.Printf("Context: %+v\n Client: %+v\n",ctx, client)

  	userToUpdate, resp, err := client.User.GetUser(ctx, "{userId}")
	if err != nil {
		fmt.Errorf("Error Getting User to update: %v", err)
	}

	fmt.Printf("User to update: %+v\n Response: %+v\n\n", userToUpdate, resp)

	newProfile := *userToUpdate.Profile
	newProfile["nickName"] = "Kylo Ren"
	updateUser := &okta.User{
		Profile: &newProfile,
	}

	updatedUser, resp, err := client.User.UpdateUser(ctx, userToUpdate.Id, *updateUser, nil)
	if err != nil {
		fmt.Errorf("Error updating user: %v", err)
	}

	fmt.Printf("Updated User: %+v\n Response: %+v\n\n", updatedUser, resp)
}
```

### Get and set custom attributes

Custom attributes must first be defined in the Okta profile editor. Then, you
can work with custom attributes on a user the same as any other profile attribute

### Remove a User

You must first deactivate the user, and then you can delete the user.

```go
import (
	"fmt"
	"context"
	"github.com/okta/okta-sdk-golang/v2/okta"
)

func main() {
  ctx, client, err := okta.NewClient(
    context.TODO(),
    okta.WithOrgUrl("https://{yourOktaDomain}"),
    okta.WithToken("{apiToken}"),
  )

  if err != nil {
    fmt.Errorf("Error: %v", err)
  }

  fmt.Printf("Context: %+v\n Client: %+v\n",ctx, client)

  resp, err := client.User.DeactivateUser(ctx, "00u14ffhw5szVqide0h8", nil)
	if err != nil {
		fmt.Errorf("Error deactivating user: %v", err)
	}

	fmt.Printf("Response: %+v\n\n", resp)

	resp, err = client.User.DeactivateOrDeleteUser(ctx, "00u14ffhw5szVqide0h8", nil)
	if err != nil {
		fmt.Errorf("Error deleting user: %v", err)
	}

	fmt.Printf("Response: %+v\n\n", resp)
}
```

### List a User's Groups

```go
import (
	"fmt"
	"context"
	"github.com/okta/okta-sdk-golang/v2/okta"
)

func main() {
  ctx, client, err := okta.NewClient(
    context.TODO(),
    okta.WithOrgUrl("https://{yourOktaDomain}"),
    okta.WithToken("{apiToken}"),
  )

  if err != nil {
    fmt.Errorf("Error: %v", err)
  }

  fmt.Printf("Context: %+v\n Client: %+v\n",ctx, client)

  groups, resp, err := client.User.ListUserGroups(ctx, "00u14fg9ff4MExj5f0h8")
	if err != nil {
		fmt.Errorf("Error getting group list for user: %v", err)
	}

	fmt.Printf("Groups: %+v\n Response: %+v\n\n", groups, resp)

	for index, group := range groups {
		fmt.Printf("Group %d: %v\n", index, group)
	}
}
```

### Create a Group

```go
import (
	"fmt"
	"context"
	"github.com/okta/okta-sdk-golang/v2/okta"
)

func main() {
  ctx, client, err := okta.NewClient(
    context.TODO(),
    okta.WithOrgUrl("https://{yourOktaDomain}"),
    okta.WithToken("{apiToken}"),
  )

  if err != nil {
    fmt.Errorf("Error: %v", err)
  }

  fmt.Printf("Context: %+v\n Client: %+v\n",ctx, client)

  groupProfile := &okta.GroupProfile{
		Name: "Sith",
	}
	groupToCreate := &okta.Group{
		Profile: groupProfile,
	}
	group, resp, err := client.Group.CreateGroup(ctx, *groupToCreate)
	if err != nil {
		fmt.Errorf("Error creating group: %v", err)
	}

	fmt.Printf("Created Group: %+v\n Response: %+v\n\n", group, resp)
}
```

### Add a User to a Group

```go
import (
	"fmt"
	"context"
	"github.com/okta/okta-sdk-golang/v2/okta"
)

func main() {
  ctx, client, err := okta.NewClient(
    context.TODO(),
    okta.WithOrgUrl("https://{yourOktaDomain}"),
    okta.WithToken("{apiToken}"),
  )

  if err != nil {
    fmt.Errorf("Error: %v", err)
  }

  fmt.Printf("Context: %+v\n Client: %+v\n",ctx, client)

  resp, err := client.Group.AddUserToGroup(ctx, "{groupId}", "{userId}")
	if err != nil {
		fmt.Errorf("Error adding user to group: %v", err)
	}

	fmt.Printf("Response: %+v\n\n", resp)
}
```

### List all Applications

```go
import (
	"fmt"
	"context"
	"github.com/okta/okta-sdk-golang/v2/okta"
)

func main() {
  ctx, client, err := okta.NewClient(
    context.TODO(),
    okta.WithOrgUrl("https://{yourOktaDomain}"),
    okta.WithToken("{apiToken}"),
  )

  if err != nil {
    fmt.Errorf("Error: %v", err)
  }

  fmt.Printf("Context: %+v\n Client: %+v\n",ctx, client)

  applicationList, resp, err := client.Application.ListApplications(ctx, nil)
	if err != nil {
		fmt.Errorf("Error listing applications: %v", err)
	}

	fmt.Printf("ApplicationList: %+v\n Response: %+v\n\n", applicationList, resp)

	// Listing applications is mapped and returned as a interface. Once you
	// get the list of applications, find the one you want to work on, and
	// make a `GET` request with that ID and the concrete application type
	for _, app := range applicationList {
		if app.(*okta.Application).Name == "oidc_client" {
			application, resp, err := client.Application.GetApplication(ctx, app.(*okta.Application).Id, okta.NewOpenIdConnectApplication(), nil)
			if err != nil {
				fmt.Errorf("Error getting application: %v", err)
			}

			fmt.Printf("Concrete Application: %+v\n Response: %+v\n\n", application, resp)
		}
	}
}
```

### Get an Application

```go
import (
	"fmt"
	"context"
	"github.com/okta/okta-sdk-golang/v2/okta"
)

func main() {
  ctx, client, err := okta.NewClient(
    context.TODO(),
    okta.WithOrgUrl("https://{yourOktaDomain}"),
    okta.WithToken("{apiToken}"),
  )

  if err != nil {
    fmt.Errorf("Error: %v", err)
  }

  fmt.Printf("Context: %+v\n Client: %+v\n",ctx, client)

  application, resp, err := client.Application.GetApplication(ctx, "0oaswjmkbtlpBDWpu0h7", okta.NewOpenIdConnectApplication(), nil)
	if err != nil {
		fmt.Errorf("Error getting application: %v", err)
	}

	fmt.Printf("Application: %+v\n Response: %+v\n\n", application, resp)

}
```

### Create an Open ID Connect Application

```go
import (
	"fmt"
	"context"
	"github.com/okta/okta-sdk-golang/v2/okta"
)

func main() {
  ctx, client, err := okta.NewClient(
    context.TODO(),
    okta.WithOrgUrl("https://{yourOktaDomain}"),
    okta.WithToken("{apiToken}"),
  )

  if err != nil {
    fmt.Errorf("Error: %v", err)
  }

  fmt.Printf("Context: %+v\n Client: %+v\n",ctx, client)

  swaApplicationSettingsApplication := okta.SwaApplicationSettingsApplication{
		ButtonField: "btn-login",
		PasswordField: "txtbox-password",
		UsernameField: "txtbox-username",
		Url: "http://example.com/login.html",
	}
	swaApplicationSettings := okta.SwaApplicationSettings{
		App: &swaApplicationSettingsApplication,
	}

	swaApp := okta.NewSwaApplication()
	swaApp.Label = "Empire Internals"
	swaApp.Settings = &swaApplicationSettings

	application, resp, err := client.Application.CreateApplication(ctx, swaApp, nil)
	if err != nil {
		fmt.Errorf("Error creating application: %v", err)
	}

	fmt.Printf("Application: %+v\n Response: %+v\n\n", application, resp)
}
```

### Call other API endpoints

Not every API endpoint is represented by a method in this library. You can call
any Okta management API endpoint using this generic syntax:

```go
import (
	"fmt"
	"context"
  "time"
	"github.com/okta/okta-sdk-golang/v2/okta"
)

func main() {
  ctx, client, err := okta.NewClient(
    context.TODO(),
    okta.WithOrgUrl("https://{yourOktaDomain}"),
    okta.WithToken("{apiToken}"),
  )

  if err != nil {
    fmt.Errorf("Error: %v", err)
  }

  fmt.Printf("Context: %+v\n Client: %+v\n",ctx, client)

  // The URL you want to call (org url will be automatically prefixed)
  url := "/api/v1/authorizationServers"

  // Set up the body of the request with structs and data
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

  // create a new request using the cloned request executor
	req, err := client.CloneRequestExecutor().NewRequest("POST", url, as)
	if err != nil {
		fmt.Printf("Error creating new request: %v", err)
	}

  // Make the request
	var authServer *AuthorizationServer
	resp, err := client.CloneRequestExecutor().Do(ctx, req, &authServer)
	if err != nil {
		fmt.Printf("Error executing request: %v", err)
	}

	fmt.Printf("Authorization Server: %v\n Response: %v\n\n", authServer, resp)
}
```

### Access Request Executor

If you need to gain access to the request executor, we have provided a method
off the `Client` to do so.

```go
re := client.CloneRequestExecutor()
```

Doing this will provide you with the ability to create your own requests for
the Okta API and call the `Do` method that handles all of the headers for you
based on the configuration.

## Configuration reference

This library looks for configuration in the following sources:

0. An `okta.yaml` file in a `.okta` folder in the current user's home directory
   (`~/.okta/okta.yaml` or `%userprofile\.okta\okta.yaml`)
0. A `.okta.yaml` file in the application or project's root directory
0. Environment variables
0. Configuration explicitly passed to the constructor (see the example in
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

## Upgrading to 2.0.x

The main purpose of this version is to include all documented,
`application/json` content-type endpoints to the SDK. During this update we
have made many changes to method names, as well as method signatures.

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

### OAuth 2.0

Okta allows you to interact with Okta APIs using scoped OAuth 2.0 access
tokens. Each access token enables the bearer to perform specific actions on
specific Okta endpoints, with that ability controlled by which scopes the
access token contains.

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
ctx := context.TODO()
ctx, client, err := okta.NewClient(ctx,
  okta.WithOrgUrl("https://{yourOktaDomain}"),
  okta.WithAuthorizationMode("PrivateKey"),
  okta.WithClientId("{{clientId}}"),
  okta.WithScopes(([]string{"okta.users.manage"})),
  okta.WithPrivateKey({{PEM PRIVATE KEY BLOCK}}) //when pasting blocks, use backticks and remove all space at beginning of each line.
)
if err != nil {
  fmt.Errorf("Error: %v", err)
}

fmt.Printf("Context: %+v\n Client: %+v\n\n",ctx, client)
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
ctx := context.TODO()
ctx, client, err := okta.NewClient(ctx,
  okta.WithOrgUrl("https://{yourOktaDomain}"),
  okta.WithAuthorizationMode("PrivateKey"),
  okta.WithClientId("{client_id}"),
  okta.WithScopes([]string{"{scopes}"}),
  okta.WithPrivateKey("{private_key}"),
  )

```

### Extending the Client

When calling `okta.NewClient()` we allow for you to pass custom instances of
`http.Client` and `cache.Cache`.

```go
myClient := &http.Client{}

myCache := NewCustomCacheDriver()

ctx, client, err := okta.NewClient(
  context.TODO(),
  okta.WithOrgUrl("https://{yourOktaDomain}"),
  okta.WithToken("{apiToken}"),
  okta.WithHttpClient(myClient),
  okta.WithCacheManager(myCache)
)
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

### Refreshing Cache for Specific Call

If you have an issue where you do a `GET`, then a `DELETE`, and then re-issue a
`GET` to the original endpoint, you may have an issue with the cache returning
with the deleted resource. An example of this is listing application users,
delete an application user, and then listing them again.

You can solve this by running `client.CloneRequestExecutor().RefreshNext()`
before your second `ListApplicationUsers` call, which will tell the call to
delete the cache for this endpoint and make a new call.

```go
appUserList, resp, err := client.Application.ListApplicationsUsers(ctx, nil)

client.Application.DeleteApplicationUser(context.TODO(), appId, appUser.Id, nil)

client.CloneRequestExecutor().RefreshNext()
appUserList, resp, err := client.Application.ListApplicationsUsers(ctx, nil)

```

### Pagination

If your request comes back with more than the default or set limit, you can
request the next page.

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

In most cases, you won't need to build the SDK from source. If you want to
build it yourself, you'll need these prerequisites:

- Clone the repo
- Run `make build` from the root of the project

## Contributing

We're happy to accept contributions and PRs! Please see the [contribution
guide](/okta/okta-sdk-golang/blob/master/CONTRIBUTING.md) to understand how to
structure a contribution.

[devforum]: https://devforum.okta.com/
[sdkapiref]: https://godoc.org/github.com/okta/okta-sdk-golang/okta
[lang-landing]: https://developer.okta.com/code/go/
[github-issues]: /okta/okta-sdk-golang/issues
[github-releases]: /okta/okta-sdk-golang/releases
