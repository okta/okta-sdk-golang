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

// SAMLPayLoadDataContext struct for SAMLPayLoadDataContext
type SAMLPayLoadDataContext struct {
	Request *InlineHookRequestObject `json:"request,omitempty"`
	Session *BaseContextSession `json:"session,omitempty"`
	User *BaseContextUser `json:"user,omitempty"`
	Protocol *SAMLPayLoadDataContextAllOfProtocol `json:"protocol,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SAMLPayLoadDataContext SAMLPayLoadDataContext

// NewSAMLPayLoadDataContext instantiates a new SAMLPayLoadDataContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSAMLPayLoadDataContext() *SAMLPayLoadDataContext {
	this := SAMLPayLoadDataContext{}
	return &this
}

// NewSAMLPayLoadDataContextWithDefaults instantiates a new SAMLPayLoadDataContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSAMLPayLoadDataContextWithDefaults() *SAMLPayLoadDataContext {
	this := SAMLPayLoadDataContext{}
	return &this
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *SAMLPayLoadDataContext) GetRequest() InlineHookRequestObject {
	if o == nil || o.Request == nil {
		var ret InlineHookRequestObject
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLPayLoadDataContext) GetRequestOk() (*InlineHookRequestObject, bool) {
	if o == nil || o.Request == nil {
		return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *SAMLPayLoadDataContext) HasRequest() bool {
	if o != nil && o.Request != nil {
		return true
	}

	return false
}

// SetRequest gets a reference to the given InlineHookRequestObject and assigns it to the Request field.
func (o *SAMLPayLoadDataContext) SetRequest(v InlineHookRequestObject) {
	o.Request = &v
}

// GetSession returns the Session field value if set, zero value otherwise.
func (o *SAMLPayLoadDataContext) GetSession() BaseContextSession {
	if o == nil || o.Session == nil {
		var ret BaseContextSession
		return ret
	}
	return *o.Session
}

// GetSessionOk returns a tuple with the Session field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLPayLoadDataContext) GetSessionOk() (*BaseContextSession, bool) {
	if o == nil || o.Session == nil {
		return nil, false
	}
	return o.Session, true
}

// HasSession returns a boolean if a field has been set.
func (o *SAMLPayLoadDataContext) HasSession() bool {
	if o != nil && o.Session != nil {
		return true
	}

	return false
}

// SetSession gets a reference to the given BaseContextSession and assigns it to the Session field.
func (o *SAMLPayLoadDataContext) SetSession(v BaseContextSession) {
	o.Session = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *SAMLPayLoadDataContext) GetUser() BaseContextUser {
	if o == nil || o.User == nil {
		var ret BaseContextUser
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLPayLoadDataContext) GetUserOk() (*BaseContextUser, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *SAMLPayLoadDataContext) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given BaseContextUser and assigns it to the User field.
func (o *SAMLPayLoadDataContext) SetUser(v BaseContextUser) {
	o.User = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *SAMLPayLoadDataContext) GetProtocol() SAMLPayLoadDataContextAllOfProtocol {
	if o == nil || o.Protocol == nil {
		var ret SAMLPayLoadDataContextAllOfProtocol
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLPayLoadDataContext) GetProtocolOk() (*SAMLPayLoadDataContextAllOfProtocol, bool) {
	if o == nil || o.Protocol == nil {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *SAMLPayLoadDataContext) HasProtocol() bool {
	if o != nil && o.Protocol != nil {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given SAMLPayLoadDataContextAllOfProtocol and assigns it to the Protocol field.
func (o *SAMLPayLoadDataContext) SetProtocol(v SAMLPayLoadDataContextAllOfProtocol) {
	o.Protocol = &v
}

func (o SAMLPayLoadDataContext) MarshalJSON() ([]byte, error) {
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SAMLPayLoadDataContext) UnmarshalJSON(bytes []byte) (err error) {
	varSAMLPayLoadDataContext := _SAMLPayLoadDataContext{}

	err = json.Unmarshal(bytes, &varSAMLPayLoadDataContext)
	if err == nil {
		*o = SAMLPayLoadDataContext(varSAMLPayLoadDataContext)
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
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableSAMLPayLoadDataContext struct {
	value *SAMLPayLoadDataContext
	isSet bool
}

func (v NullableSAMLPayLoadDataContext) Get() *SAMLPayLoadDataContext {
	return v.value
}

func (v *NullableSAMLPayLoadDataContext) Set(val *SAMLPayLoadDataContext) {
	v.value = val
	v.isSet = true
}

func (v NullableSAMLPayLoadDataContext) IsSet() bool {
	return v.isSet
}

func (v *NullableSAMLPayLoadDataContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSAMLPayLoadDataContext(val *SAMLPayLoadDataContext) *NullableSAMLPayLoadDataContext {
	return &NullableSAMLPayLoadDataContext{value: val, isSet: true}
}

func (v NullableSAMLPayLoadDataContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSAMLPayLoadDataContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

