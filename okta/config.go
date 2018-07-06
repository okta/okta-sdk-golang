/*
 * Copyright 2018 - Present Okta, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package okta

import (
	"fmt"
	"github.com/go-yaml/yaml"
	"github.com/kelseyhightower/envconfig"
	"io/ioutil"
	"log"
	"os/user"
)

type Config struct {
	Okta struct {
		Client struct {
			Cache struct {
				Enabled    bool  `yaml:"enabled" envconfig:"OKTA_CLIENT_CACHE_ENABLED"`
				DefaultTtl int32 `yaml:"defaultTtl" envconfig:"OKTA_CLIENT_CACHE_DEFAULT_TTL"`
				DefaultTti int32 `yaml:"defaultTti" envconfig:"OKTA_CLIENT_CACHE_DEFAULT_TTI"`
			} `yaml:"cache"`
			Proxy struct {
				Port     int32  `yaml:"port" envconfig:"OKTA_CLIENT_PROXY_PORT"`
				Host     string `yaml:"host" envconfig:"OKTA_CLIENT_PROXY_HOST"`
				Username string `yaml:"username" envconfig:"OKTA_CLIENT_PROXY_USERNAME"`
				Password string `yaml:"password" envconfig:"OKTA_CLIENT_PROXY_PASSWORD"`
			} `yaml:"proxy"`
			ConnectionTimeout int32  `yaml:"connectionTimeout" envconfig:"OKTA_CLIENT_CONNECTION_TIMEOUT"`
			OrgUrl            string `yaml:"orgUrl" envconfig:"OKTA_CLIENT_ORGURL"`
			Token             string `yaml:"token" envconfig:"OKTA_CLIENT_TOKEN"`
		} `yaml:"client"`
	} `yaml:"okta"`
	UserAgentExtra string
}

func NewConfig() *Config {
	c := Config{}

	c.WithCache(true).
		WithCacheTtl(300).
		WithCacheTti(300).
		WithConnectionTimeout(30).
		WithUserAgentExtra("")

	c = readConfigFromSystem(c)
	c = readConfigFromApplication(c)
	c = readConfigFromEnvironment(c)

	return &c
}

func (c *Config) ReadConfigFromFile(location string) *Config {
	yamlConfig, err := ioutil.ReadFile(location)

	if err != nil {
		log.Printf("Reading configuration from file failed. %s", err.Error())
		return c
	}

	conf := Config{}
	err = yaml.Unmarshal(yamlConfig, &conf)
	if err != nil {
		return c
	}

	return &conf
}

func (c *Config) WithCache(cache bool) *Config {
	c.Okta.Client.Cache.Enabled = cache
	return c
}

func (c *Config) WithCacheTtl(i int32) *Config {
	c.Okta.Client.Cache.DefaultTtl = i
	return c
}

func (c *Config) WithCacheTti(i int32) *Config {
	c.Okta.Client.Cache.DefaultTti = i
	return c
}

func (c *Config) WithConnectionTimeout(i int32) *Config {
	c.Okta.Client.ConnectionTimeout = i
	return c
}

func (c *Config) WithProxyPort(i int32) *Config {
	c.Okta.Client.Proxy.Port = i
	return c
}

func (c *Config) WithProxyHost(host string) *Config {
	c.Okta.Client.Proxy.Host = host
	return c
}

func (c *Config) WithProxyUsername(username string) *Config {
	c.Okta.Client.Proxy.Username = username
	return c
}

func (c *Config) WithProxyPassword(pass string) *Config {
	c.Okta.Client.Proxy.Password = pass
	return c
}

func (c *Config) WithOrgUrl(url string) *Config {
	c.Okta.Client.OrgUrl = url
	return c
}

func (c *Config) WithToken(token string) *Config {
	c.Okta.Client.Token = token
	return c
}

func (c *Config) WithUserAgentExtra(s string) *Config {
	c.UserAgentExtra = s
	return c
}

func readConfigFromSystem(c Config) Config {
	currUser, err := user.Current()
	if err != nil {
		return c
	}
	if currUser.HomeDir == "" {
		return c
	}

	return *c.ReadConfigFromFile(currUser.HomeDir + "/.okta/okta.yaml")
}

func readConfigFromApplication(c Config) Config {
	return *c.ReadConfigFromFile(".okta.yaml")
}

func readConfigFromEnvironment(c Config) Config {
	err := envconfig.Process("okta", &c)
	if err != nil {
		fmt.Println("error parsing")
		return c
	}
	return c
}
