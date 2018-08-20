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
	"fmt"

	"github.com/okta/okta-sdk-golang/okta/cache"
	"github.com/okta/okta-sdk-golang/okta/query"
)

const Version = "0.0.0-Develop"

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

func NewClient(config *Config) *Client {
	if config == nil {
		config = NewConfig()
	}

	var oktaCache cache.Cache

	if !config.Okta.Client.Cache.Enabled {
		oktaCache = cache.NewNoOpCache()
	} else {
		oktaCache = cache.NewMapCache(config.Okta.Client.Cache.DefaultTtl, config.Okta.Client.Cache.DefaultTti)
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

func (c *Client) ListApplications(qp *query.Params) ([]*Application, *Response, error) {
	url := fmt.Sprintf("/api/v1/apps")
	if qp != nil {
		url = url + qp.String()
	}
	req, err := c.requestExecutor.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var application []*Application
	resp, err := c.requestExecutor.Do(req, &application)
	if err != nil {
		return nil, resp, err
	}
	return application, resp, nil
}
func (c *Client) CreateApplication(body Application, qp *query.Params) (*Application, *Response, error) {
	url := fmt.Sprintf("/api/v1/apps")
	if qp != nil {
		url = url + qp.String()
	}
	req, err := c.requestExecutor.NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}

	var application *Application
	resp, err := c.requestExecutor.Do(req, &application)
	if err != nil {
		return nil, resp, err
	}
	return application, resp, nil
}
func (c *Client) DeleteApplication(appId string, qp *query.Params) (*Response, error) {
	url := fmt.Sprintf("/api/v1/apps/%v", appId)
	if qp != nil {
		url = url + qp.String()
	}
	req, err := c.requestExecutor.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.requestExecutor.Do(req, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
func (c *Client) GetApplication(appId string, qp *query.Params) (*Application, *Response, error) {
	url := fmt.Sprintf("/api/v1/apps/%v", appId)
	if qp != nil {
		url = url + qp.String()
	}
	req, err := c.requestExecutor.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var application *Application
	resp, err := c.requestExecutor.Do(req, &application)
	if err != nil {
		return nil, resp, err
	}
	return application, resp, nil
}
func (c *Client) UpdateApplication(appId string, body Application, qp *query.Params) (*Application, *Response, error) {
	url := fmt.Sprintf("/api/v1/apps/%v", appId)
	if qp != nil {
		url = url + qp.String()
	}
	req, err := c.requestExecutor.NewRequest("PUT", url, body)
	if err != nil {
		return nil, nil, err
	}

	var application *Application
	resp, err := c.requestExecutor.Do(req, &application)
	if err != nil {
		return nil, resp, err
	}
	return application, resp, nil
}
func (c *Client) ListApplicationKeys(appId string, qp *query.Params) ([]*JsonWebKey, *Response, error) {
	url := fmt.Sprintf("/api/v1/apps/%v/credentials/keys", appId)
	if qp != nil {
		url = url + qp.String()
	}
	req, err := c.requestExecutor.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var jsonWebKey []*JsonWebKey
	resp, err := c.requestExecutor.Do(req, &jsonWebKey)
	if err != nil {
		return nil, resp, err
	}
	return jsonWebKey, resp, nil
}
func (c *Client) GetApplicationKey(appId string, keyId string, qp *query.Params) (*JsonWebKey, *Response, error) {
	url := fmt.Sprintf("/api/v1/apps/%v/credentials/keys/%v", appId, keyId)
	if qp != nil {
		url = url + qp.String()
	}
	req, err := c.requestExecutor.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var jsonWebKey *JsonWebKey
	resp, err := c.requestExecutor.Do(req, &jsonWebKey)
	if err != nil {
		return nil, resp, err
	}
	return jsonWebKey, resp, nil
}
func (c *Client) CloneApplicationKey(appId string, keyId string, qp *query.Params) (*JsonWebKey, *Response, error) {
	url := fmt.Sprintf("/api/v1/apps/%v/credentials/keys/%v/clone", appId, keyId)
	if qp != nil {
		url = url + qp.String()
	}
	req, err := c.requestExecutor.NewRequest("POST", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var jsonWebKey *JsonWebKey
	resp, err := c.requestExecutor.Do(req, &jsonWebKey)
	if err != nil {
		return nil, resp, err
	}
	return jsonWebKey, resp, nil
}
func (c *Client) ListApplicationGroupAssignments(appId string, qp *query.Params) ([]*ApplicationGroupAssignment, *Response, error) {
	url := fmt.Sprintf("/api/v1/apps/%v/groups", appId)
	if qp != nil {
		url = url + qp.String()
	}
	req, err := c.requestExecutor.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var applicationGroupAssignment []*ApplicationGroupAssignment
	resp, err := c.requestExecutor.Do(req, &applicationGroupAssignment)
	if err != nil {
		return nil, resp, err
	}
	return applicationGroupAssignment, resp, nil
}
func (c *Client) DeleteApplicationGroupAssignment(appId string, groupId string, qp *query.Params) (*Response, error) {
	url := fmt.Sprintf("/api/v1/apps/%v/groups/%v", appId, groupId)
	if qp != nil {
		url = url + qp.String()
	}
	req, err := c.requestExecutor.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.requestExecutor.Do(req, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
func (c *Client) GetApplicationGroupAssignment(appId string, groupId string, qp *query.Params) (*ApplicationGroupAssignment, *Response, error) {
	url := fmt.Sprintf("/api/v1/apps/%v/groups/%v", appId, groupId)
	if qp != nil {
		url = url + qp.String()
	}
	req, err := c.requestExecutor.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var applicationGroupAssignment *ApplicationGroupAssignment
	resp, err := c.requestExecutor.Do(req, &applicationGroupAssignment)
	if err != nil {
		return nil, resp, err
	}
	return applicationGroupAssignment, resp, nil
}
func (c *Client) CreateApplicationGroupAssignment(appId string, groupId string, body ApplicationGroupAssignment, qp *query.Params) (*ApplicationGroupAssignment, *Response, error) {
	url := fmt.Sprintf("/api/v1/apps/%v/groups/%v", appId, groupId)
	if qp != nil {
		url = url + qp.String()
	}
	req, err := c.requestExecutor.NewRequest("PUT", url, body)
	if err != nil {
		return nil, nil, err
	}

	var applicationGroupAssignment *ApplicationGroupAssignment
	resp, err := c.requestExecutor.Do(req, &applicationGroupAssignment)
	if err != nil {
		return nil, resp, err
	}
	return applicationGroupAssignment, resp, nil
}
func (c *Client) ActivateApplication(appId string, qp *query.Params) (*Response, error) {
	url := fmt.Sprintf("/api/v1/apps/%v/lifecycle/activate", appId)
	if qp != nil {
		url = url + qp.String()
	}
	req, err := c.requestExecutor.NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.requestExecutor.Do(req, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
func (c *Client) DeactivateApplication(appId string, qp *query.Params) (*Response, error) {
	url := fmt.Sprintf("/api/v1/apps/%v/lifecycle/deactivate", appId)
	if qp != nil {
		url = url + qp.String()
	}
	req, err := c.requestExecutor.NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.requestExecutor.Do(req, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
func (c *Client) ListApplicationUsers(appId string, qp *query.Params) ([]*AppUser, *Response, error) {
	url := fmt.Sprintf("/api/v1/apps/%v/users", appId)
	if qp != nil {
		url = url + qp.String()
	}
	req, err := c.requestExecutor.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var appUser []*AppUser
	resp, err := c.requestExecutor.Do(req, &appUser)
	if err != nil {
		return nil, resp, err
	}
	return appUser, resp, nil
}
func (c *Client) AssignUserToApplication(appId string, body AppUser, qp *query.Params) (*AppUser, *Response, error) {
	url := fmt.Sprintf("/api/v1/apps/%v/users", appId)
	if qp != nil {
		url = url + qp.String()
	}
	req, err := c.requestExecutor.NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}

	var appUser *AppUser
	resp, err := c.requestExecutor.Do(req, &appUser)
	if err != nil {
		return nil, resp, err
	}
	return appUser, resp, nil
}
func (c *Client) DeleteApplicationUser(appId string, userId string, qp *query.Params) (*Response, error) {
	url := fmt.Sprintf("/api/v1/apps/%v/users/%v", appId, userId)
	if qp != nil {
		url = url + qp.String()
	}
	req, err := c.requestExecutor.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.requestExecutor.Do(req, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
func (c *Client) GetApplicationUser(appId string, userId string, qp *query.Params) (*AppUser, *Response, error) {
	url := fmt.Sprintf("/api/v1/apps/%v/users/%v", appId, userId)
	if qp != nil {
		url = url + qp.String()
	}
	req, err := c.requestExecutor.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var appUser *AppUser
	resp, err := c.requestExecutor.Do(req, &appUser)
	if err != nil {
		return nil, resp, err
	}
	return appUser, resp, nil
}
func (c *Client) UpdateApplicationUser(appId string, userId string, body AppUser, qp *query.Params) (*AppUser, *Response, error) {
	url := fmt.Sprintf("/api/v1/apps/%v/users/%v", appId, userId)
	if qp != nil {
		url = url + qp.String()
	}
	req, err := c.requestExecutor.NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}

	var appUser *AppUser
	resp, err := c.requestExecutor.Do(req, &appUser)
	if err != nil {
		return nil, resp, err
	}
	return appUser, resp, nil
}
func (c *Client) ListGroups(qp *query.Params) ([]*Group, *Response, error) {
	url := fmt.Sprintf("/api/v1/groups")
	if qp != nil {
		url = url + qp.String()
	}
	req, err := c.requestExecutor.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var group []*Group
	resp, err := c.requestExecutor.Do(req, &group)
	if err != nil {
		return nil, resp, err
	}
	return group, resp, nil
}
func (c *Client) CreateGroup(body Group, qp *query.Params) (*Group, *Response, error) {
	url := fmt.Sprintf("/api/v1/groups")
	if qp != nil {
		url = url + qp.String()
	}
	req, err := c.requestExecutor.NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}

	var group *Group
	resp, err := c.requestExecutor.Do(req, &group)
	if err != nil {
		return nil, resp, err
	}
	return group, resp, nil
}
func (c *Client) ListRules(qp *query.Params) ([]*GroupRule, *Response, error) {
	url := fmt.Sprintf("/api/v1/groups/rules")
	if qp != nil {
		url = url + qp.String()
	}
	req, err := c.requestExecutor.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var groupRule []*GroupRule
	resp, err := c.requestExecutor.Do(req, &groupRule)
	if err != nil {
		return nil, resp, err
	}
	return groupRule, resp, nil
}
func (c *Client) CreateRule(body GroupRule, qp *query.Params) (*GroupRule, *Response, error) {
	url := fmt.Sprintf("/api/v1/groups/rules")
	if qp != nil {
		url = url + qp.String()
	}
	req, err := c.requestExecutor.NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}

	var groupRule *GroupRule
	resp, err := c.requestExecutor.Do(req, &groupRule)
	if err != nil {
		return nil, resp, err
	}
	return groupRule, resp, nil
}
func (c *Client) DeleteRule(ruleId string, qp *query.Params) (*Response, error) {
	url := fmt.Sprintf("/api/v1/groups/rules/%v", ruleId)
	if qp != nil {
		url = url + qp.String()
	}
	req, err := c.requestExecutor.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.requestExecutor.Do(req, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
func (c *Client) GetRule(ruleId string, qp *query.Params) (*GroupRule, *Response, error) {
	url := fmt.Sprintf("/api/v1/groups/rules/%v", ruleId)
	if qp != nil {
		url = url + qp.String()
	}
	req, err := c.requestExecutor.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var groupRule *GroupRule
	resp, err := c.requestExecutor.Do(req, &groupRule)
	if err != nil {
		return nil, resp, err
	}
	return groupRule, resp, nil
}
func (c *Client) UpdateRule(ruleId string, body GroupRule, qp *query.Params) (*GroupRule, *Response, error) {
	url := fmt.Sprintf("/api/v1/groups/rules/%v", ruleId)
	if qp != nil {
		url = url + qp.String()
	}
	req, err := c.requestExecutor.NewRequest("PUT", url, body)
	if err != nil {
		return nil, nil, err
	}

	var groupRule *GroupRule
	resp, err := c.requestExecutor.Do(req, &groupRule)
	if err != nil {
		return nil, resp, err
	}
	return groupRule, resp, nil
}
func (c *Client) ActivateRule(ruleId string, qp *query.Params) (*Response, error) {
	url := fmt.Sprintf("/api/v1/groups/rules/%v/lifecycle/activate", ruleId)
	if qp != nil {
		url = url + qp.String()
	}
	req, err := c.requestExecutor.NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.requestExecutor.Do(req, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
func (c *Client) DeactivateRule(ruleId string, qp *query.Params) (*Response, error) {
	url := fmt.Sprintf("/api/v1/groups/rules/%v/lifecycle/deactivate", ruleId)
	if qp != nil {
		url = url + qp.String()
	}
	req, err := c.requestExecutor.NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.requestExecutor.Do(req, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
func (c *Client) DeleteGroup(groupId string, qp *query.Params) (*Response, error) {
	url := fmt.Sprintf("/api/v1/groups/%v", groupId)
	if qp != nil {
		url = url + qp.String()
	}
	req, err := c.requestExecutor.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.requestExecutor.Do(req, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
func (c *Client) GetGroup(groupId string, qp *query.Params) (*Group, *Response, error) {
	url := fmt.Sprintf("/api/v1/groups/%v", groupId)
	if qp != nil {
		url = url + qp.String()
	}
	req, err := c.requestExecutor.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var group *Group
	resp, err := c.requestExecutor.Do(req, &group)
	if err != nil {
		return nil, resp, err
	}
	return group, resp, nil
}
func (c *Client) UpdateGroup(groupId string, body Group, qp *query.Params) (*Group, *Response, error) {
	url := fmt.Sprintf("/api/v1/groups/%v", groupId)
	if qp != nil {
		url = url + qp.String()
	}
	req, err := c.requestExecutor.NewRequest("PUT", url, body)
	if err != nil {
		return nil, nil, err
	}

	var group *Group
	resp, err := c.requestExecutor.Do(req, &group)
	if err != nil {
		return nil, resp, err
	}
	return group, resp, nil
}
func (c *Client) ListGroupUsers(groupId string, qp *query.Params) ([]*User, *Response, error) {
	url := fmt.Sprintf("/api/v1/groups/%v/users", groupId)
	if qp != nil {
		url = url + qp.String()
	}
	req, err := c.requestExecutor.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var user []*User
	resp, err := c.requestExecutor.Do(req, &user)
	if err != nil {
		return nil, resp, err
	}
	return user, resp, nil
}
func (c *Client) RemoveGroupUser(groupId string, userId string, qp *query.Params) (*Response, error) {
	url := fmt.Sprintf("/api/v1/groups/%v/users/%v", groupId, userId)
	if qp != nil {
		url = url + qp.String()
	}
	req, err := c.requestExecutor.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.requestExecutor.Do(req, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
func (c *Client) AddUserToGroup(groupId string, userId string, qp *query.Params) (*Response, error) {
	url := fmt.Sprintf("/api/v1/groups/%v/users/%v", groupId, userId)
	if qp != nil {
		url = url + qp.String()
	}
	req, err := c.requestExecutor.NewRequest("PUT", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.requestExecutor.Do(req, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
func (c *Client) GetLogs(qp *query.Params) ([]*LogEvent, *Response, error) {
	url := fmt.Sprintf("/api/v1/logs")
	if qp != nil {
		url = url + qp.String()
	}
	req, err := c.requestExecutor.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var logEvent []*LogEvent
	resp, err := c.requestExecutor.Do(req, &logEvent)
	if err != nil {
		return nil, resp, err
	}
	return logEvent, resp, nil
}
func (c *Client) CreateSession(body CreateSessionRequest, qp *query.Params) (*Session, *Response, error) {
	url := fmt.Sprintf("/api/v1/sessions")
	if qp != nil {
		url = url + qp.String()
	}
	req, err := c.requestExecutor.NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}

	var session *Session
	resp, err := c.requestExecutor.Do(req, &session)
	if err != nil {
		return nil, resp, err
	}
	return session, resp, nil
}
func (c *Client) EndSession(sessionId string, qp *query.Params) (*Response, error) {
	url := fmt.Sprintf("/api/v1/sessions/%v", sessionId)
	if qp != nil {
		url = url + qp.String()
	}
	req, err := c.requestExecutor.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.requestExecutor.Do(req, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
func (c *Client) GetSession(sessionId string, qp *query.Params) (*Session, *Response, error) {
	url := fmt.Sprintf("/api/v1/sessions/%v", sessionId)
	if qp != nil {
		url = url + qp.String()
	}
	req, err := c.requestExecutor.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var session *Session
	resp, err := c.requestExecutor.Do(req, &session)
	if err != nil {
		return nil, resp, err
	}
	return session, resp, nil
}
func (c *Client) RefreshSession(sessionId string, qp *query.Params) (*Session, *Response, error) {
	url := fmt.Sprintf("/api/v1/sessions/%v/lifecycle/refresh", sessionId)
	if qp != nil {
		url = url + qp.String()
	}
	req, err := c.requestExecutor.NewRequest("POST", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var session *Session
	resp, err := c.requestExecutor.Do(req, &session)
	if err != nil {
		return nil, resp, err
	}
	return session, resp, nil
}
func (c *Client) ListUsers(qp *query.Params) ([]*User, *Response, error) {
	url := fmt.Sprintf("/api/v1/users")
	if qp != nil {
		url = url + qp.String()
	}
	req, err := c.requestExecutor.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var user []*User
	resp, err := c.requestExecutor.Do(req, &user)
	if err != nil {
		return nil, resp, err
	}
	return user, resp, nil
}
func (c *Client) CreateUser(body User, qp *query.Params) (*User, *Response, error) {
	url := fmt.Sprintf("/api/v1/users")
	if qp != nil {
		url = url + qp.String()
	}
	req, err := c.requestExecutor.NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}

	var user *User
	resp, err := c.requestExecutor.Do(req, &user)
	if err != nil {
		return nil, resp, err
	}
	return user, resp, nil
}
func (c *Client) DeactivateOrDeleteUser(userId string, qp *query.Params) (*Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v", userId)
	if qp != nil {
		url = url + qp.String()
	}
	req, err := c.requestExecutor.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.requestExecutor.Do(req, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
func (c *Client) GetUser(userId string, qp *query.Params) (*User, *Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v", userId)
	if qp != nil {
		url = url + qp.String()
	}
	req, err := c.requestExecutor.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var user *User
	resp, err := c.requestExecutor.Do(req, &user)
	if err != nil {
		return nil, resp, err
	}
	return user, resp, nil
}
func (c *Client) UpdateUser(userId string, body User, qp *query.Params) (*User, *Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v", userId)
	if qp != nil {
		url = url + qp.String()
	}
	req, err := c.requestExecutor.NewRequest("PUT", url, body)
	if err != nil {
		return nil, nil, err
	}

	var user *User
	resp, err := c.requestExecutor.Do(req, &user)
	if err != nil {
		return nil, resp, err
	}
	return user, resp, nil
}
func (c *Client) ListAppLinks(userId string, qp *query.Params) ([]*AppLink, *Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/appLinks", userId)
	if qp != nil {
		url = url + qp.String()
	}
	req, err := c.requestExecutor.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var appLink []*AppLink
	resp, err := c.requestExecutor.Do(req, &appLink)
	if err != nil {
		return nil, resp, err
	}
	return appLink, resp, nil
}
func (c *Client) ChangePassword(userId string, body ChangePasswordRequest, qp *query.Params) (*UserCredentials, *Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/credentials/change_password", userId)
	if qp != nil {
		url = url + qp.String()
	}
	req, err := c.requestExecutor.NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}

	var userCredentials *UserCredentials
	resp, err := c.requestExecutor.Do(req, &userCredentials)
	if err != nil {
		return nil, resp, err
	}
	return userCredentials, resp, nil
}
func (c *Client) ChangeRecoveryQuestion(userId string, body UserCredentials, qp *query.Params) (*UserCredentials, *Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/credentials/change_recovery_question", userId)
	if qp != nil {
		url = url + qp.String()
	}
	req, err := c.requestExecutor.NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}

	var userCredentials *UserCredentials
	resp, err := c.requestExecutor.Do(req, &userCredentials)
	if err != nil {
		return nil, resp, err
	}
	return userCredentials, resp, nil
}
func (c *Client) ForgotPassword(userId string, body UserCredentials, qp *query.Params) (*ForgotPasswordResponse, *Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/credentials/forgot_password", userId)
	if qp != nil {
		url = url + qp.String()
	}
	req, err := c.requestExecutor.NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}

	var forgotPasswordResponse *ForgotPasswordResponse
	resp, err := c.requestExecutor.Do(req, &forgotPasswordResponse)
	if err != nil {
		return nil, resp, err
	}
	return forgotPasswordResponse, resp, nil
}
func (c *Client) ListFactors(userId string, qp *query.Params) ([]*Factor, *Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/factors", userId)
	if qp != nil {
		url = url + qp.String()
	}
	req, err := c.requestExecutor.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var factor []*Factor
	resp, err := c.requestExecutor.Do(req, &factor)
	if err != nil {
		return nil, resp, err
	}
	return factor, resp, nil
}
func (c *Client) AddFactor(userId string, body Factor, qp *query.Params) (*Factor, *Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/factors", userId)
	if qp != nil {
		url = url + qp.String()
	}
	req, err := c.requestExecutor.NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}

	var factor *Factor
	resp, err := c.requestExecutor.Do(req, &factor)
	if err != nil {
		return nil, resp, err
	}
	return factor, resp, nil
}
func (c *Client) ListSupportedFactors(userId string, qp *query.Params) ([]*Factor, *Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/factors/catalog", userId)
	if qp != nil {
		url = url + qp.String()
	}
	req, err := c.requestExecutor.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var factor []*Factor
	resp, err := c.requestExecutor.Do(req, &factor)
	if err != nil {
		return nil, resp, err
	}
	return factor, resp, nil
}
func (c *Client) ListSupportedSecurityQuestions(userId string, qp *query.Params) ([]*SecurityQuestion, *Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/factors/questions", userId)
	if qp != nil {
		url = url + qp.String()
	}
	req, err := c.requestExecutor.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var securityQuestion []*SecurityQuestion
	resp, err := c.requestExecutor.Do(req, &securityQuestion)
	if err != nil {
		return nil, resp, err
	}
	return securityQuestion, resp, nil
}
func (c *Client) DeleteFactor(userId string, factorId string, qp *query.Params) (*Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/factors/%v", userId, factorId)
	if qp != nil {
		url = url + qp.String()
	}
	req, err := c.requestExecutor.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.requestExecutor.Do(req, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
func (c *Client) GetFactor(userId string, factorId string, qp *query.Params) (*Factor, *Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/factors/%v", userId, factorId)
	if qp != nil {
		url = url + qp.String()
	}
	req, err := c.requestExecutor.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var factor *Factor
	resp, err := c.requestExecutor.Do(req, &factor)
	if err != nil {
		return nil, resp, err
	}
	return factor, resp, nil
}
func (c *Client) ActivateFactor(userId string, factorId string, body VerifyFactorRequest, qp *query.Params) (*Factor, *Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/factors/%v/lifecycle/activate", userId, factorId)
	if qp != nil {
		url = url + qp.String()
	}
	req, err := c.requestExecutor.NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}

	var factor *Factor
	resp, err := c.requestExecutor.Do(req, &factor)
	if err != nil {
		return nil, resp, err
	}
	return factor, resp, nil
}
func (c *Client) VerifyFactor(userId string, factorId string, body VerifyFactorRequest, qp *query.Params) (*VerifyFactorResponse, *Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/factors/%v/verify", userId, factorId)
	if qp != nil {
		url = url + qp.String()
	}
	req, err := c.requestExecutor.NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}

	var verifyFactorResponse *VerifyFactorResponse
	resp, err := c.requestExecutor.Do(req, &verifyFactorResponse)
	if err != nil {
		return nil, resp, err
	}
	return verifyFactorResponse, resp, nil
}
func (c *Client) ListUserGroups(userId string, qp *query.Params) ([]*Group, *Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/groups", userId)
	if qp != nil {
		url = url + qp.String()
	}
	req, err := c.requestExecutor.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var group []*Group
	resp, err := c.requestExecutor.Do(req, &group)
	if err != nil {
		return nil, resp, err
	}
	return group, resp, nil
}
func (c *Client) ActivateUser(userId string, qp *query.Params) (*UserActivationToken, *Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/lifecycle/activate", userId)
	if qp != nil {
		url = url + qp.String()
	}
	req, err := c.requestExecutor.NewRequest("POST", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var userActivationToken *UserActivationToken
	resp, err := c.requestExecutor.Do(req, &userActivationToken)
	if err != nil {
		return nil, resp, err
	}
	return userActivationToken, resp, nil
}
func (c *Client) DeactivateUser(userId string, qp *query.Params) (*Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/lifecycle/deactivate", userId)
	if qp != nil {
		url = url + qp.String()
	}
	req, err := c.requestExecutor.NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.requestExecutor.Do(req, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
func (c *Client) ExpirePassword(userId string, qp *query.Params) (*TempPassword, *Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/lifecycle/expire_password", userId)
	if qp != nil {
		url = url + qp.String()
	}
	req, err := c.requestExecutor.NewRequest("POST", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var tempPassword *TempPassword
	resp, err := c.requestExecutor.Do(req, &tempPassword)
	if err != nil {
		return nil, resp, err
	}
	return tempPassword, resp, nil
}
func (c *Client) ResetAllFactors(userId string, qp *query.Params) (*Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/lifecycle/reset_factors", userId)
	if qp != nil {
		url = url + qp.String()
	}
	req, err := c.requestExecutor.NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.requestExecutor.Do(req, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
func (c *Client) ResetPassword(userId string, qp *query.Params) (*ResetPasswordToken, *Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/lifecycle/reset_password", userId)
	if qp != nil {
		url = url + qp.String()
	}
	req, err := c.requestExecutor.NewRequest("POST", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var resetPasswordToken *ResetPasswordToken
	resp, err := c.requestExecutor.Do(req, &resetPasswordToken)
	if err != nil {
		return nil, resp, err
	}
	return resetPasswordToken, resp, nil
}
func (c *Client) SuspendUser(userId string, qp *query.Params) (*Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/lifecycle/suspend", userId)
	if qp != nil {
		url = url + qp.String()
	}
	req, err := c.requestExecutor.NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.requestExecutor.Do(req, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
func (c *Client) UnlockUser(userId string, qp *query.Params) (*Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/lifecycle/unlock", userId)
	if qp != nil {
		url = url + qp.String()
	}
	req, err := c.requestExecutor.NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.requestExecutor.Do(req, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
func (c *Client) UnsuspendUser(userId string, qp *query.Params) (*Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/lifecycle/unsuspend", userId)
	if qp != nil {
		url = url + qp.String()
	}
	req, err := c.requestExecutor.NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.requestExecutor.Do(req, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
func (c *Client) ListAssignedRoles(userId string, qp *query.Params) ([]*Role, *Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/roles", userId)
	if qp != nil {
		url = url + qp.String()
	}
	req, err := c.requestExecutor.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var role []*Role
	resp, err := c.requestExecutor.Do(req, &role)
	if err != nil {
		return nil, resp, err
	}
	return role, resp, nil
}
func (c *Client) AddRoleToUser(userId string, body Role, qp *query.Params) (*Role, *Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/roles", userId)
	if qp != nil {
		url = url + qp.String()
	}
	req, err := c.requestExecutor.NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}

	var role *Role
	resp, err := c.requestExecutor.Do(req, &role)
	if err != nil {
		return nil, resp, err
	}
	return role, resp, nil
}
func (c *Client) RemoveRoleFromUser(userId string, roleId string, qp *query.Params) (*Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/roles/%v", userId, roleId)
	if qp != nil {
		url = url + qp.String()
	}
	req, err := c.requestExecutor.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.requestExecutor.Do(req, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
func (c *Client) ListGroupTargetsForRole(userId string, roleId string, qp *query.Params) ([]*Group, *Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/roles/%v/targets/groups", userId, roleId)
	if qp != nil {
		url = url + qp.String()
	}
	req, err := c.requestExecutor.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var group []*Group
	resp, err := c.requestExecutor.Do(req, &group)
	if err != nil {
		return nil, resp, err
	}
	return group, resp, nil
}
func (c *Client) RemoveGroupTargetFromRole(userId string, roleId string, groupId string, qp *query.Params) (*Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/roles/%v/targets/groups/%v", userId, roleId, groupId)
	if qp != nil {
		url = url + qp.String()
	}
	req, err := c.requestExecutor.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.requestExecutor.Do(req, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
func (c *Client) AddGroupTargetToRole(userId string, roleId string, groupId string, qp *query.Params) (*Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/roles/%v/targets/groups/%v", userId, roleId, groupId)
	if qp != nil {
		url = url + qp.String()
	}
	req, err := c.requestExecutor.NewRequest("PUT", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.requestExecutor.Do(req, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
func (c *Client) EndAllUserSessions(userId string, qp *query.Params) (*Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/sessions", userId)
	if qp != nil {
		url = url + qp.String()
	}
	req, err := c.requestExecutor.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.requestExecutor.Do(req, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
