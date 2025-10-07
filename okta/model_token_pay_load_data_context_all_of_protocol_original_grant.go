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

API version: 2025.08.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the TokenPayLoadDataContextAllOfProtocolOriginalGrant type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenPayLoadDataContextAllOfProtocolOriginalGrant{}

// TokenPayLoadDataContextAllOfProtocolOriginalGrant Information about the original token request used to get the refresh token being used, when in a refresh token request
type TokenPayLoadDataContextAllOfProtocolOriginalGrant struct {
	Request              *TokenProtocolRequest `json:"request,omitempty"`
	RefreshToken         *RefreshToken         `json:"refresh_token,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TokenPayLoadDataContextAllOfProtocolOriginalGrant TokenPayLoadDataContextAllOfProtocolOriginalGrant

// NewTokenPayLoadDataContextAllOfProtocolOriginalGrant instantiates a new TokenPayLoadDataContextAllOfProtocolOriginalGrant object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenPayLoadDataContextAllOfProtocolOriginalGrant() *TokenPayLoadDataContextAllOfProtocolOriginalGrant {
	this := TokenPayLoadDataContextAllOfProtocolOriginalGrant{}
	return &this
}

// NewTokenPayLoadDataContextAllOfProtocolOriginalGrantWithDefaults instantiates a new TokenPayLoadDataContextAllOfProtocolOriginalGrant object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenPayLoadDataContextAllOfProtocolOriginalGrantWithDefaults() *TokenPayLoadDataContextAllOfProtocolOriginalGrant {
	this := TokenPayLoadDataContextAllOfProtocolOriginalGrant{}
	return &this
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *TokenPayLoadDataContextAllOfProtocolOriginalGrant) GetRequest() TokenProtocolRequest {
	if o == nil || IsNil(o.Request) {
		var ret TokenProtocolRequest
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenPayLoadDataContextAllOfProtocolOriginalGrant) GetRequestOk() (*TokenProtocolRequest, bool) {
	if o == nil || IsNil(o.Request) {
		return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *TokenPayLoadDataContextAllOfProtocolOriginalGrant) HasRequest() bool {
	if o != nil && !IsNil(o.Request) {
		return true
	}

	return false
}

// SetRequest gets a reference to the given TokenProtocolRequest and assigns it to the Request field.
func (o *TokenPayLoadDataContextAllOfProtocolOriginalGrant) SetRequest(v TokenProtocolRequest) {
	o.Request = &v
}

// GetRefreshToken returns the RefreshToken field value if set, zero value otherwise.
func (o *TokenPayLoadDataContextAllOfProtocolOriginalGrant) GetRefreshToken() RefreshToken {
	if o == nil || IsNil(o.RefreshToken) {
		var ret RefreshToken
		return ret
	}
	return *o.RefreshToken
}

// GetRefreshTokenOk returns a tuple with the RefreshToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenPayLoadDataContextAllOfProtocolOriginalGrant) GetRefreshTokenOk() (*RefreshToken, bool) {
	if o == nil || IsNil(o.RefreshToken) {
		return nil, false
	}
	return o.RefreshToken, true
}

// HasRefreshToken returns a boolean if a field has been set.
func (o *TokenPayLoadDataContextAllOfProtocolOriginalGrant) HasRefreshToken() bool {
	if o != nil && !IsNil(o.RefreshToken) {
		return true
	}

	return false
}

// SetRefreshToken gets a reference to the given RefreshToken and assigns it to the RefreshToken field.
func (o *TokenPayLoadDataContextAllOfProtocolOriginalGrant) SetRefreshToken(v RefreshToken) {
	o.RefreshToken = &v
}

func (o TokenPayLoadDataContextAllOfProtocolOriginalGrant) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenPayLoadDataContextAllOfProtocolOriginalGrant) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Request) {
		toSerialize["request"] = o.Request
	}
	if !IsNil(o.RefreshToken) {
		toSerialize["refresh_token"] = o.RefreshToken
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TokenPayLoadDataContextAllOfProtocolOriginalGrant) UnmarshalJSON(data []byte) (err error) {
	varTokenPayLoadDataContextAllOfProtocolOriginalGrant := _TokenPayLoadDataContextAllOfProtocolOriginalGrant{}

	err = json.Unmarshal(data, &varTokenPayLoadDataContextAllOfProtocolOriginalGrant)

	if err != nil {
		return err
	}

	*o = TokenPayLoadDataContextAllOfProtocolOriginalGrant(varTokenPayLoadDataContextAllOfProtocolOriginalGrant)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "request")
		delete(additionalProperties, "refresh_token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTokenPayLoadDataContextAllOfProtocolOriginalGrant struct {
	value *TokenPayLoadDataContextAllOfProtocolOriginalGrant
	isSet bool
}

func (v NullableTokenPayLoadDataContextAllOfProtocolOriginalGrant) Get() *TokenPayLoadDataContextAllOfProtocolOriginalGrant {
	return v.value
}

func (v *NullableTokenPayLoadDataContextAllOfProtocolOriginalGrant) Set(val *TokenPayLoadDataContextAllOfProtocolOriginalGrant) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenPayLoadDataContextAllOfProtocolOriginalGrant) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenPayLoadDataContextAllOfProtocolOriginalGrant) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenPayLoadDataContextAllOfProtocolOriginalGrant(val *TokenPayLoadDataContextAllOfProtocolOriginalGrant) *NullableTokenPayLoadDataContextAllOfProtocolOriginalGrant {
	return &NullableTokenPayLoadDataContextAllOfProtocolOriginalGrant{value: val, isSet: true}
}

func (v NullableTokenPayLoadDataContextAllOfProtocolOriginalGrant) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenPayLoadDataContextAllOfProtocolOriginalGrant) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
