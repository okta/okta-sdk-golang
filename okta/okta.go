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

const Version = "0.0.0-Develop"

type Client struct {
	config *Config

	requestExecutor *RequestExecutor

	resource resource

	Application *ApplicationResource

}

type resource struct {
	client *Client
}

func NewClient(config *Config) *Client {
	if config == nil {
		config = NewConfig()
	}

	c := &Client{}
	c.config = config
	c.requestExecutor = NewRequestExecutor(nil, config)

	c.resource.client = c

	c.Application = (*ApplicationResource)(&c.resource)
	return c
}

func (c *Client) GetConfig() *Config {
	return c.config
}

func (c *Client) GetRequestExecutor() *RequestExecutor {
	return c.requestExecutor
}
