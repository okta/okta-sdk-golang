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

// TokenPayLoadDataContextAllOfProtocol Details of the authentication protocol
type TokenPayLoadDataContextAllOfProtocol struct {
	// The type of authentication protocol used
	Type *string `json:"type,omitempty"`
	Request *TokenProtocolRequest `json:"request,omitempty"`
	OriginalGrant *TokenPayLoadDataContextAllOfProtocolOriginalGrant `json:"OriginalGrant,omitempty"`
	Issuer *TokenPayLoadDataContextAllOfProtocolIssuer `json:"issuer,omitempty"`
	Client *TokenPayLoadDataContextAllOfProtocolClient `json:"client,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TokenPayLoadDataContextAllOfProtocol TokenPayLoadDataContextAllOfProtocol

// NewTokenPayLoadDataContextAllOfProtocol instantiates a new TokenPayLoadDataContextAllOfProtocol object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenPayLoadDataContextAllOfProtocol() *TokenPayLoadDataContextAllOfProtocol {
	this := TokenPayLoadDataContextAllOfProtocol{}
	return &this
}

// NewTokenPayLoadDataContextAllOfProtocolWithDefaults instantiates a new TokenPayLoadDataContextAllOfProtocol object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenPayLoadDataContextAllOfProtocolWithDefaults() *TokenPayLoadDataContextAllOfProtocol {
	this := TokenPayLoadDataContextAllOfProtocol{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TokenPayLoadDataContextAllOfProtocol) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenPayLoadDataContextAllOfProtocol) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TokenPayLoadDataContextAllOfProtocol) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *TokenPayLoadDataContextAllOfProtocol) SetType(v string) {
	o.Type = &v
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *TokenPayLoadDataContextAllOfProtocol) GetRequest() TokenProtocolRequest {
	if o == nil || o.Request == nil {
		var ret TokenProtocolRequest
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenPayLoadDataContextAllOfProtocol) GetRequestOk() (*TokenProtocolRequest, bool) {
	if o == nil || o.Request == nil {
		return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *TokenPayLoadDataContextAllOfProtocol) HasRequest() bool {
	if o != nil && o.Request != nil {
		return true
	}

	return false
}

// SetRequest gets a reference to the given TokenProtocolRequest and assigns it to the Request field.
func (o *TokenPayLoadDataContextAllOfProtocol) SetRequest(v TokenProtocolRequest) {
	o.Request = &v
}

// GetOriginalGrant returns the OriginalGrant field value if set, zero value otherwise.
func (o *TokenPayLoadDataContextAllOfProtocol) GetOriginalGrant() TokenPayLoadDataContextAllOfProtocolOriginalGrant {
	if o == nil || o.OriginalGrant == nil {
		var ret TokenPayLoadDataContextAllOfProtocolOriginalGrant
		return ret
	}
	return *o.OriginalGrant
}

// GetOriginalGrantOk returns a tuple with the OriginalGrant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenPayLoadDataContextAllOfProtocol) GetOriginalGrantOk() (*TokenPayLoadDataContextAllOfProtocolOriginalGrant, bool) {
	if o == nil || o.OriginalGrant == nil {
		return nil, false
	}
	return o.OriginalGrant, true
}

// HasOriginalGrant returns a boolean if a field has been set.
func (o *TokenPayLoadDataContextAllOfProtocol) HasOriginalGrant() bool {
	if o != nil && o.OriginalGrant != nil {
		return true
	}

	return false
}

// SetOriginalGrant gets a reference to the given TokenPayLoadDataContextAllOfProtocolOriginalGrant and assigns it to the OriginalGrant field.
func (o *TokenPayLoadDataContextAllOfProtocol) SetOriginalGrant(v TokenPayLoadDataContextAllOfProtocolOriginalGrant) {
	o.OriginalGrant = &v
}

// GetIssuer returns the Issuer field value if set, zero value otherwise.
func (o *TokenPayLoadDataContextAllOfProtocol) GetIssuer() TokenPayLoadDataContextAllOfProtocolIssuer {
	if o == nil || o.Issuer == nil {
		var ret TokenPayLoadDataContextAllOfProtocolIssuer
		return ret
	}
	return *o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenPayLoadDataContextAllOfProtocol) GetIssuerOk() (*TokenPayLoadDataContextAllOfProtocolIssuer, bool) {
	if o == nil || o.Issuer == nil {
		return nil, false
	}
	return o.Issuer, true
}

// HasIssuer returns a boolean if a field has been set.
func (o *TokenPayLoadDataContextAllOfProtocol) HasIssuer() bool {
	if o != nil && o.Issuer != nil {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given TokenPayLoadDataContextAllOfProtocolIssuer and assigns it to the Issuer field.
func (o *TokenPayLoadDataContextAllOfProtocol) SetIssuer(v TokenPayLoadDataContextAllOfProtocolIssuer) {
	o.Issuer = &v
}

// GetClient returns the Client field value if set, zero value otherwise.
func (o *TokenPayLoadDataContextAllOfProtocol) GetClient() TokenPayLoadDataContextAllOfProtocolClient {
	if o == nil || o.Client == nil {
		var ret TokenPayLoadDataContextAllOfProtocolClient
		return ret
	}
	return *o.Client
}

// GetClientOk returns a tuple with the Client field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenPayLoadDataContextAllOfProtocol) GetClientOk() (*TokenPayLoadDataContextAllOfProtocolClient, bool) {
	if o == nil || o.Client == nil {
		return nil, false
	}
	return o.Client, true
}

// HasClient returns a boolean if a field has been set.
func (o *TokenPayLoadDataContextAllOfProtocol) HasClient() bool {
	if o != nil && o.Client != nil {
		return true
	}

	return false
}

// SetClient gets a reference to the given TokenPayLoadDataContextAllOfProtocolClient and assigns it to the Client field.
func (o *TokenPayLoadDataContextAllOfProtocol) SetClient(v TokenPayLoadDataContextAllOfProtocolClient) {
	o.Client = &v
}

func (o TokenPayLoadDataContextAllOfProtocol) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Request != nil {
		toSerialize["request"] = o.Request
	}
	if o.OriginalGrant != nil {
		toSerialize["OriginalGrant"] = o.OriginalGrant
	}
	if o.Issuer != nil {
		toSerialize["issuer"] = o.Issuer
	}
	if o.Client != nil {
		toSerialize["client"] = o.Client
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TokenPayLoadDataContextAllOfProtocol) UnmarshalJSON(bytes []byte) (err error) {
	varTokenPayLoadDataContextAllOfProtocol := _TokenPayLoadDataContextAllOfProtocol{}

	err = json.Unmarshal(bytes, &varTokenPayLoadDataContextAllOfProtocol)
	if err == nil {
		*o = TokenPayLoadDataContextAllOfProtocol(varTokenPayLoadDataContextAllOfProtocol)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "request")
		delete(additionalProperties, "OriginalGrant")
		delete(additionalProperties, "issuer")
		delete(additionalProperties, "client")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableTokenPayLoadDataContextAllOfProtocol struct {
	value *TokenPayLoadDataContextAllOfProtocol
	isSet bool
}

func (v NullableTokenPayLoadDataContextAllOfProtocol) Get() *TokenPayLoadDataContextAllOfProtocol {
	return v.value
}

func (v *NullableTokenPayLoadDataContextAllOfProtocol) Set(val *TokenPayLoadDataContextAllOfProtocol) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenPayLoadDataContextAllOfProtocol) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenPayLoadDataContextAllOfProtocol) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenPayLoadDataContextAllOfProtocol(val *TokenPayLoadDataContextAllOfProtocol) *NullableTokenPayLoadDataContextAllOfProtocol {
	return &NullableTokenPayLoadDataContextAllOfProtocol{value: val, isSet: true}
}

func (v NullableTokenPayLoadDataContextAllOfProtocol) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenPayLoadDataContextAllOfProtocol) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

