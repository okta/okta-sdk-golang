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

// AUTO-GENERATED!  DO NOT EDIT FILE DIRECTLY

package okta

import (
	"net/http"

	"github.com/okta/okta-sdk-golang/okta/cache"
)

const Version = "0.1.0"

type Client struct {
	config *Config

	requestExecutor *RequestExecutor

	resource resource

	Application *ApplicationResource
	Group       *GroupResource
	Session     *SessionResource
	User        *UserResource
	Factor      *FactorResource
}

type resource struct {
	client *Client
}

func NewClient(config *Config, httpClient *http.Client, cacheManager cache.Cache) *Client {
	if config == nil {
		config = NewConfig()
	}

	var oktaCache cache.Cache
	if !config.Okta.Client.Cache.Enabled {
		oktaCache = cache.NewNoOpCache()
	} else {
		if cacheManager == nil {
			oktaCache = cache.NewGoCache(config.Okta.Client.Cache.DefaultTtl,
				config.Okta.Client.Cache.DefaultTti)
		} else {
			oktaCache = cacheManager
		}
	}

	c := &Client{}
	c.config = config
	c.requestExecutor = NewRequestExecutor(nil, oktaCache, config)

	c.resource.client = c

	c.Application = (*ApplicationResource)(&c.resource)
	c.Group = (*GroupResource)(&c.resource)
	c.Session = (*SessionResource)(&c.resource)
	c.User = (*UserResource)(&c.resource)
	c.Factor = (*FactorResource)(&c.resource)
	return c
}

func (c *Client) GetConfig() *Config {
	return c.config
}

func (c *Client) GetRequestExecutor() *RequestExecutor {
	return c.requestExecutor
}
