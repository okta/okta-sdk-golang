/*
Okta Admin Management

Allows customers to easily access the Okta Management APIs

Copyright 2025 - Present Okta, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

API version: 2025.10.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the TokenProtocolRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenProtocolRequest{}

// TokenProtocolRequest Details of the token request
type TokenProtocolRequest struct {
	// The ID of the client associated with the token
	ClientId *string `json:"client_id,omitempty"`
	// Determines the mechanism Okta uses to authorize the creation of the tokens.
	GrantType *string `json:"grant_type,omitempty"`
	// Specifies the callback location where the authorization was sent
	RedirectUri *string `json:"redirect_uri,omitempty"`
	// The authorization response mode
	ResponseMode *string `json:"response_mode,omitempty"`
	// The authorization response type
	ResponseType *string `json:"response_type,omitempty"`
	// The scopes requested
	Scope                *string `json:"scope,omitempty"`
	State                *string `json:"state,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TokenProtocolRequest TokenProtocolRequest

// NewTokenProtocolRequest instantiates a new TokenProtocolRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenProtocolRequest() *TokenProtocolRequest {
	this := TokenProtocolRequest{}
	return &this
}

// NewTokenProtocolRequestWithDefaults instantiates a new TokenProtocolRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenProtocolRequestWithDefaults() *TokenProtocolRequest {
	this := TokenProtocolRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *TokenProtocolRequest) GetClientId() string {
	if o == nil || IsNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenProtocolRequest) GetClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *TokenProtocolRequest) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *TokenProtocolRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetGrantType returns the GrantType field value if set, zero value otherwise.
func (o *TokenProtocolRequest) GetGrantType() string {
	if o == nil || IsNil(o.GrantType) {
		var ret string
		return ret
	}
	return *o.GrantType
}

// GetGrantTypeOk returns a tuple with the GrantType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenProtocolRequest) GetGrantTypeOk() (*string, bool) {
	if o == nil || IsNil(o.GrantType) {
		return nil, false
	}
	return o.GrantType, true
}

// HasGrantType returns a boolean if a field has been set.
func (o *TokenProtocolRequest) HasGrantType() bool {
	if o != nil && !IsNil(o.GrantType) {
		return true
	}

	return false
}

// SetGrantType gets a reference to the given string and assigns it to the GrantType field.
func (o *TokenProtocolRequest) SetGrantType(v string) {
	o.GrantType = &v
}

// GetRedirectUri returns the RedirectUri field value if set, zero value otherwise.
func (o *TokenProtocolRequest) GetRedirectUri() string {
	if o == nil || IsNil(o.RedirectUri) {
		var ret string
		return ret
	}
	return *o.RedirectUri
}

// GetRedirectUriOk returns a tuple with the RedirectUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenProtocolRequest) GetRedirectUriOk() (*string, bool) {
	if o == nil || IsNil(o.RedirectUri) {
		return nil, false
	}
	return o.RedirectUri, true
}

// HasRedirectUri returns a boolean if a field has been set.
func (o *TokenProtocolRequest) HasRedirectUri() bool {
	if o != nil && !IsNil(o.RedirectUri) {
		return true
	}

	return false
}

// SetRedirectUri gets a reference to the given string and assigns it to the RedirectUri field.
func (o *TokenProtocolRequest) SetRedirectUri(v string) {
	o.RedirectUri = &v
}

// GetResponseMode returns the ResponseMode field value if set, zero value otherwise.
func (o *TokenProtocolRequest) GetResponseMode() string {
	if o == nil || IsNil(o.ResponseMode) {
		var ret string
		return ret
	}
	return *o.ResponseMode
}

// GetResponseModeOk returns a tuple with the ResponseMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenProtocolRequest) GetResponseModeOk() (*string, bool) {
	if o == nil || IsNil(o.ResponseMode) {
		return nil, false
	}
	return o.ResponseMode, true
}

// HasResponseMode returns a boolean if a field has been set.
func (o *TokenProtocolRequest) HasResponseMode() bool {
	if o != nil && !IsNil(o.ResponseMode) {
		return true
	}

	return false
}

// SetResponseMode gets a reference to the given string and assigns it to the ResponseMode field.
func (o *TokenProtocolRequest) SetResponseMode(v string) {
	o.ResponseMode = &v
}

// GetResponseType returns the ResponseType field value if set, zero value otherwise.
func (o *TokenProtocolRequest) GetResponseType() string {
	if o == nil || IsNil(o.ResponseType) {
		var ret string
		return ret
	}
	return *o.ResponseType
}

// GetResponseTypeOk returns a tuple with the ResponseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenProtocolRequest) GetResponseTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ResponseType) {
		return nil, false
	}
	return o.ResponseType, true
}

// HasResponseType returns a boolean if a field has been set.
func (o *TokenProtocolRequest) HasResponseType() bool {
	if o != nil && !IsNil(o.ResponseType) {
		return true
	}

	return false
}

// SetResponseType gets a reference to the given string and assigns it to the ResponseType field.
func (o *TokenProtocolRequest) SetResponseType(v string) {
	o.ResponseType = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *TokenProtocolRequest) GetScope() string {
	if o == nil || IsNil(o.Scope) {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenProtocolRequest) GetScopeOk() (*string, bool) {
	if o == nil || IsNil(o.Scope) {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *TokenProtocolRequest) HasScope() bool {
	if o != nil && !IsNil(o.Scope) {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *TokenProtocolRequest) SetScope(v string) {
	o.Scope = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *TokenProtocolRequest) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenProtocolRequest) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *TokenProtocolRequest) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *TokenProtocolRequest) SetState(v string) {
	o.State = &v
}

func (o TokenProtocolRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenProtocolRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClientId) {
		toSerialize["client_id"] = o.ClientId
	}
	if !IsNil(o.GrantType) {
		toSerialize["grant_type"] = o.GrantType
	}
	if !IsNil(o.RedirectUri) {
		toSerialize["redirect_uri"] = o.RedirectUri
	}
	if !IsNil(o.ResponseMode) {
		toSerialize["response_mode"] = o.ResponseMode
	}
	if !IsNil(o.ResponseType) {
		toSerialize["response_type"] = o.ResponseType
	}
	if !IsNil(o.Scope) {
		toSerialize["scope"] = o.Scope
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TokenProtocolRequest) UnmarshalJSON(data []byte) (err error) {
	varTokenProtocolRequest := _TokenProtocolRequest{}

	err = json.Unmarshal(data, &varTokenProtocolRequest)

	if err != nil {
		return err
	}

	*o = TokenProtocolRequest(varTokenProtocolRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "client_id")
		delete(additionalProperties, "grant_type")
		delete(additionalProperties, "redirect_uri")
		delete(additionalProperties, "response_mode")
		delete(additionalProperties, "response_type")
		delete(additionalProperties, "scope")
		delete(additionalProperties, "state")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTokenProtocolRequest struct {
	value *TokenProtocolRequest
	isSet bool
}

func (v NullableTokenProtocolRequest) Get() *TokenProtocolRequest {
	return v.value
}

func (v *NullableTokenProtocolRequest) Set(val *TokenProtocolRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenProtocolRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenProtocolRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenProtocolRequest(val *TokenProtocolRequest) *NullableTokenProtocolRequest {
	return &NullableTokenProtocolRequest{value: val, isSet: true}
}

func (v NullableTokenProtocolRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenProtocolRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
