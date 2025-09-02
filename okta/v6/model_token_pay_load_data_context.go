/*
Okta Admin Management

Allows customers to easily access the Okta Management APIs

Copyright 2018 - Present Okta, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// TokenPayLoadDataContext struct for TokenPayLoadDataContext
type TokenPayLoadDataContext struct {
	Request *InlineHookRequestObject `json:"request,omitempty"`
	Session *BaseContextSession `json:"session,omitempty"`
	User *BaseContextUser `json:"user,omitempty"`
	Protocol *TokenPayLoadDataContextAllOfProtocol `json:"protocol,omitempty"`
	Policy *TokenPayLoadDataContextAllOfPolicy `json:"policy,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TokenPayLoadDataContext TokenPayLoadDataContext

// NewTokenPayLoadDataContext instantiates a new TokenPayLoadDataContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenPayLoadDataContext() *TokenPayLoadDataContext {
	this := TokenPayLoadDataContext{}
	return &this
}

// NewTokenPayLoadDataContextWithDefaults instantiates a new TokenPayLoadDataContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenPayLoadDataContextWithDefaults() *TokenPayLoadDataContext {
	this := TokenPayLoadDataContext{}
	return &this
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *TokenPayLoadDataContext) GetRequest() InlineHookRequestObject {
	if o == nil || o.Request == nil {
		var ret InlineHookRequestObject
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenPayLoadDataContext) GetRequestOk() (*InlineHookRequestObject, bool) {
	if o == nil || o.Request == nil {
		return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *TokenPayLoadDataContext) HasRequest() bool {
	if o != nil && o.Request != nil {
		return true
	}

	return false
}

// SetRequest gets a reference to the given InlineHookRequestObject and assigns it to the Request field.
func (o *TokenPayLoadDataContext) SetRequest(v InlineHookRequestObject) {
	o.Request = &v
}

// GetSession returns the Session field value if set, zero value otherwise.
func (o *TokenPayLoadDataContext) GetSession() BaseContextSession {
	if o == nil || o.Session == nil {
		var ret BaseContextSession
		return ret
	}
	return *o.Session
}

// GetSessionOk returns a tuple with the Session field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenPayLoadDataContext) GetSessionOk() (*BaseContextSession, bool) {
	if o == nil || o.Session == nil {
		return nil, false
	}
	return o.Session, true
}

// HasSession returns a boolean if a field has been set.
func (o *TokenPayLoadDataContext) HasSession() bool {
	if o != nil && o.Session != nil {
		return true
	}

	return false
}

// SetSession gets a reference to the given BaseContextSession and assigns it to the Session field.
func (o *TokenPayLoadDataContext) SetSession(v BaseContextSession) {
	o.Session = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *TokenPayLoadDataContext) GetUser() BaseContextUser {
	if o == nil || o.User == nil {
		var ret BaseContextUser
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenPayLoadDataContext) GetUserOk() (*BaseContextUser, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *TokenPayLoadDataContext) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given BaseContextUser and assigns it to the User field.
func (o *TokenPayLoadDataContext) SetUser(v BaseContextUser) {
	o.User = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *TokenPayLoadDataContext) GetProtocol() TokenPayLoadDataContextAllOfProtocol {
	if o == nil || o.Protocol == nil {
		var ret TokenPayLoadDataContextAllOfProtocol
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenPayLoadDataContext) GetProtocolOk() (*TokenPayLoadDataContextAllOfProtocol, bool) {
	if o == nil || o.Protocol == nil {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *TokenPayLoadDataContext) HasProtocol() bool {
	if o != nil && o.Protocol != nil {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given TokenPayLoadDataContextAllOfProtocol and assigns it to the Protocol field.
func (o *TokenPayLoadDataContext) SetProtocol(v TokenPayLoadDataContextAllOfProtocol) {
	o.Protocol = &v
}

// GetPolicy returns the Policy field value if set, zero value otherwise.
func (o *TokenPayLoadDataContext) GetPolicy() TokenPayLoadDataContextAllOfPolicy {
	if o == nil || o.Policy == nil {
		var ret TokenPayLoadDataContextAllOfPolicy
		return ret
	}
	return *o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenPayLoadDataContext) GetPolicyOk() (*TokenPayLoadDataContextAllOfPolicy, bool) {
	if o == nil || o.Policy == nil {
		return nil, false
	}
	return o.Policy, true
}

// HasPolicy returns a boolean if a field has been set.
func (o *TokenPayLoadDataContext) HasPolicy() bool {
	if o != nil && o.Policy != nil {
		return true
	}

	return false
}

// SetPolicy gets a reference to the given TokenPayLoadDataContextAllOfPolicy and assigns it to the Policy field.
func (o *TokenPayLoadDataContext) SetPolicy(v TokenPayLoadDataContextAllOfPolicy) {
	o.Policy = &v
}

func (o TokenPayLoadDataContext) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Request != nil {
		toSerialize["request"] = o.Request
	}
	if o.Session != nil {
		toSerialize["session"] = o.Session
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	if o.Protocol != nil {
		toSerialize["protocol"] = o.Protocol
	}
	if o.Policy != nil {
		toSerialize["policy"] = o.Policy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TokenPayLoadDataContext) UnmarshalJSON(bytes []byte) (err error) {
	varTokenPayLoadDataContext := _TokenPayLoadDataContext{}

	err = json.Unmarshal(bytes, &varTokenPayLoadDataContext)
	if err == nil {
		*o = TokenPayLoadDataContext(varTokenPayLoadDataContext)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "request")
		delete(additionalProperties, "session")
		delete(additionalProperties, "user")
		delete(additionalProperties, "protocol")
		delete(additionalProperties, "policy")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableTokenPayLoadDataContext struct {
	value *TokenPayLoadDataContext
	isSet bool
}

func (v NullableTokenPayLoadDataContext) Get() *TokenPayLoadDataContext {
	return v.value
}

func (v *NullableTokenPayLoadDataContext) Set(val *TokenPayLoadDataContext) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenPayLoadDataContext) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenPayLoadDataContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenPayLoadDataContext(val *TokenPayLoadDataContext) *NullableTokenPayLoadDataContext {
	return &NullableTokenPayLoadDataContext{value: val, isSet: true}
}

func (v NullableTokenPayLoadDataContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenPayLoadDataContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

