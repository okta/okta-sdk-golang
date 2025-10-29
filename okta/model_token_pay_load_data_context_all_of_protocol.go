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

// checks if the TokenPayLoadDataContextAllOfProtocol type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenPayLoadDataContextAllOfProtocol{}

// TokenPayLoadDataContextAllOfProtocol Details of the authentication protocol
type TokenPayLoadDataContextAllOfProtocol struct {
	// The type of authentication protocol used
	Type                 *string                                            `json:"type,omitempty"`
	Request              *TokenProtocolRequest                              `json:"request,omitempty"`
	OriginalGrant        *TokenPayLoadDataContextAllOfProtocolOriginalGrant `json:"OriginalGrant,omitempty"`
	Issuer               *TokenPayLoadDataContextAllOfProtocolIssuer        `json:"issuer,omitempty"`
	Client               *TokenPayLoadDataContextAllOfProtocolClient        `json:"client,omitempty"`
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
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenPayLoadDataContextAllOfProtocol) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TokenPayLoadDataContextAllOfProtocol) HasType() bool {
	if o != nil && !IsNil(o.Type) {
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
	if o == nil || IsNil(o.Request) {
		var ret TokenProtocolRequest
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenPayLoadDataContextAllOfProtocol) GetRequestOk() (*TokenProtocolRequest, bool) {
	if o == nil || IsNil(o.Request) {
		return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *TokenPayLoadDataContextAllOfProtocol) HasRequest() bool {
	if o != nil && !IsNil(o.Request) {
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
	if o == nil || IsNil(o.OriginalGrant) {
		var ret TokenPayLoadDataContextAllOfProtocolOriginalGrant
		return ret
	}
	return *o.OriginalGrant
}

// GetOriginalGrantOk returns a tuple with the OriginalGrant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenPayLoadDataContextAllOfProtocol) GetOriginalGrantOk() (*TokenPayLoadDataContextAllOfProtocolOriginalGrant, bool) {
	if o == nil || IsNil(o.OriginalGrant) {
		return nil, false
	}
	return o.OriginalGrant, true
}

// HasOriginalGrant returns a boolean if a field has been set.
func (o *TokenPayLoadDataContextAllOfProtocol) HasOriginalGrant() bool {
	if o != nil && !IsNil(o.OriginalGrant) {
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
	if o == nil || IsNil(o.Issuer) {
		var ret TokenPayLoadDataContextAllOfProtocolIssuer
		return ret
	}
	return *o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenPayLoadDataContextAllOfProtocol) GetIssuerOk() (*TokenPayLoadDataContextAllOfProtocolIssuer, bool) {
	if o == nil || IsNil(o.Issuer) {
		return nil, false
	}
	return o.Issuer, true
}

// HasIssuer returns a boolean if a field has been set.
func (o *TokenPayLoadDataContextAllOfProtocol) HasIssuer() bool {
	if o != nil && !IsNil(o.Issuer) {
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
	if o == nil || IsNil(o.Client) {
		var ret TokenPayLoadDataContextAllOfProtocolClient
		return ret
	}
	return *o.Client
}

// GetClientOk returns a tuple with the Client field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenPayLoadDataContextAllOfProtocol) GetClientOk() (*TokenPayLoadDataContextAllOfProtocolClient, bool) {
	if o == nil || IsNil(o.Client) {
		return nil, false
	}
	return o.Client, true
}

// HasClient returns a boolean if a field has been set.
func (o *TokenPayLoadDataContextAllOfProtocol) HasClient() bool {
	if o != nil && !IsNil(o.Client) {
		return true
	}

	return false
}

// SetClient gets a reference to the given TokenPayLoadDataContextAllOfProtocolClient and assigns it to the Client field.
func (o *TokenPayLoadDataContextAllOfProtocol) SetClient(v TokenPayLoadDataContextAllOfProtocolClient) {
	o.Client = &v
}

func (o TokenPayLoadDataContextAllOfProtocol) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenPayLoadDataContextAllOfProtocol) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Request) {
		toSerialize["request"] = o.Request
	}
	if !IsNil(o.OriginalGrant) {
		toSerialize["OriginalGrant"] = o.OriginalGrant
	}
	if !IsNil(o.Issuer) {
		toSerialize["issuer"] = o.Issuer
	}
	if !IsNil(o.Client) {
		toSerialize["client"] = o.Client
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TokenPayLoadDataContextAllOfProtocol) UnmarshalJSON(data []byte) (err error) {
	varTokenPayLoadDataContextAllOfProtocol := _TokenPayLoadDataContextAllOfProtocol{}

	err = json.Unmarshal(data, &varTokenPayLoadDataContextAllOfProtocol)

	if err != nil {
		return err
	}

	*o = TokenPayLoadDataContextAllOfProtocol(varTokenPayLoadDataContextAllOfProtocol)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "request")
		delete(additionalProperties, "OriginalGrant")
		delete(additionalProperties, "issuer")
		delete(additionalProperties, "client")
		o.AdditionalProperties = additionalProperties
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
