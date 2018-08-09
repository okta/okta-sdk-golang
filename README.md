# Okta Management SDK for Golang
[![Build Status](https://img.shields.io/travis/okta/okta-sdk-golang/master.svg)](https://travis-ci.org/okta/okta-sdk-golang)
[![License](https://img.shields.io/github/license/okta/okta-sdk-golang.svg)](https://opensource.org/licenses/Apache-2.0)
[![Support](https://img.shields.io/badge/support-Developer%20Forum-blue.svg)](https://devforum.okta.com/)

## :warning: :construction: Alpha Preview :construction: :warning:
This library is under development and is currently in 0.0.x version series.  Breaking changes may be introduced at minor versions in the 0.x range.  Please lock your dependency to a specific version until this library reaches 1.x.

Need help? [Okta Developer Forum].

## Configuration
Configuration details are pulled from 5 different sources. The following is the order the SDK looks for configuration items and builds on top of all the defaults.

1. Default configuration.
2. okta.yaml file from ~/.okta directory.
3. .okta.yaml file from application directory.
4. Environment variables.
5. Configuration provided with the Config struct

### Creating a New Configuration
The sdk will provide a default configuration when calling for a new config struct. This will return all okta defaults which you can then build on top of as long as your dont have any other configuration files or options as listed above.

```
oktaConfig := okta.NewConfig()
```
The above will result in the following being set for the `oktaConfig`

```
{
  Okta: {
    Client: {
      Cache: {
        Enabled:true
        DefaultTtl:300
        DefaultTti:300
      }
      Proxy: {
        Port:0
        Host:
        Username:
        Password:
      }
      ConnectionTimeout:30
      OrgUrl:
      Token:
    }
  }
}
```

### Modifying the Configuration
There are a few ways to modify the default configuration at this point.

#### Creating a Global okta.yaml file
Within the users home directory, you need to create a `.okta/okta.yaml` file. This file should include the below, using any of the values you want to override, the default config with. This file can include the defaults as well

```
okta:
  client:
    cache:
      enabled: true
      defaultTtl: 300 # seconds
      defaultTti: 300
      caches: { }
    connectionTimeout: 30 # seconds
    orgUrl: "https://dev-123.oktapreview.com/"
    proxy:
      port: null
      host: null
      username: null
      password: null
    token: xxxxx
```

#### Creating a Local .okta.yaml file
The same rules apply as above, but this file should be called `.okta.yaml` and live inside fo your application root directory

#### Using Environment Variables
Environment variables are a great secure way of storing these credentials. The environment variable should be all upper case and `_` seperated.  As an example, if you wanted to create a Default TTL for the cache, your environment variable would be `OKTA_CLIENT_CACHE_DEFAULT_TTL`

#### Overriding Config with Methods
As the last step in providing overrides for your config values, the following would overwrite any previous set configuration options.

```
oktaConfig := okta.NewConfig().
                  WithToken("{YourAPIToken}").
                  WithOrgUrl("https://{YourOrgUrl}.com").
                  WithCache(false). // this would disable your cache
                  WithConnectionTimeout(15) //this sets your connection timeout to 15 seconds

```

[Okta Developer Forum]: https://devforum.okta.com/
