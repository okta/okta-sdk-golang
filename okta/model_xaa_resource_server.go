/*
Okta Admin Management API

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

// checks if the XAAResourceServer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &XAAResourceServer{}

// XAAResourceServer XAA resource server metadata for an app instance
type XAAResourceServer struct {
	// Indicates whether the ID-JAG audience for this resource server must carry a tenant claim. When `true`, a tenant value is collected from the customer when the app is configured and is included in the ID-JAG JWT audience. When absent, this property defaults to `false`, which is the case for single-tenant resource servers.
	AudTenantRequired *bool `json:"audTenantRequired,omitempty"`
	// Indicates whether this resource server supports Client ID Metadata Documents (CIMD)
	CimdSupported *bool `json:"cimdSupported,omitempty"`
	// The issuer URL that identifies the authorization server for the app (resource server).  The `issuer` property from the OIN catalog integration can be an expression that supports the [Okta Expression Language's app properties](https://developer.okta.com/docs/reference/okta-expression-language/#application-properties), and contains the app properties that represent the customer tenant. For example, when the `issuer` on the OIN catalog integration is `https://{app.subdomain}.example.com/`, the app instance's `issuer` property is resolved to `https://mydomain.example.com` if the `subdomain` app property variable is set to `mydomain`. The fully resolved value is returned.
	Issuer *string `json:"issuer,omitempty"`
	// The resource URLs protected by this resource server
	ProtectedResources []string `json:"protectedResources,omitempty"`
	// The OAuth 2.0 scopes that this resource server accepts
	Scopes               []string `json:"scopes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _XAAResourceServer XAAResourceServer

// NewXAAResourceServer instantiates a new XAAResourceServer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewXAAResourceServer() *XAAResourceServer {
	this := XAAResourceServer{}
	return &this
}

// NewXAAResourceServerWithDefaults instantiates a new XAAResourceServer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewXAAResourceServerWithDefaults() *XAAResourceServer {
	this := XAAResourceServer{}
	return &this
}

// GetAudTenantRequired returns the AudTenantRequired field value if set, zero value otherwise.
func (o *XAAResourceServer) GetAudTenantRequired() bool {
	if o == nil || IsNil(o.AudTenantRequired) {
		var ret bool
		return ret
	}
	return *o.AudTenantRequired
}

// GetAudTenantRequiredOk returns a tuple with the AudTenantRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XAAResourceServer) GetAudTenantRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.AudTenantRequired) {
		return nil, false
	}
	return o.AudTenantRequired, true
}

// HasAudTenantRequired returns a boolean if a field has been set.
func (o *XAAResourceServer) HasAudTenantRequired() bool {
	if o != nil && !IsNil(o.AudTenantRequired) {
		return true
	}

	return false
}

// SetAudTenantRequired gets a reference to the given bool and assigns it to the AudTenantRequired field.
func (o *XAAResourceServer) SetAudTenantRequired(v bool) {
	o.AudTenantRequired = &v
}

// GetCimdSupported returns the CimdSupported field value if set, zero value otherwise.
func (o *XAAResourceServer) GetCimdSupported() bool {
	if o == nil || IsNil(o.CimdSupported) {
		var ret bool
		return ret
	}
	return *o.CimdSupported
}

// GetCimdSupportedOk returns a tuple with the CimdSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XAAResourceServer) GetCimdSupportedOk() (*bool, bool) {
	if o == nil || IsNil(o.CimdSupported) {
		return nil, false
	}
	return o.CimdSupported, true
}

// HasCimdSupported returns a boolean if a field has been set.
func (o *XAAResourceServer) HasCimdSupported() bool {
	if o != nil && !IsNil(o.CimdSupported) {
		return true
	}

	return false
}

// SetCimdSupported gets a reference to the given bool and assigns it to the CimdSupported field.
func (o *XAAResourceServer) SetCimdSupported(v bool) {
	o.CimdSupported = &v
}

// GetIssuer returns the Issuer field value if set, zero value otherwise.
func (o *XAAResourceServer) GetIssuer() string {
	if o == nil || IsNil(o.Issuer) {
		var ret string
		return ret
	}
	return *o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XAAResourceServer) GetIssuerOk() (*string, bool) {
	if o == nil || IsNil(o.Issuer) {
		return nil, false
	}
	return o.Issuer, true
}

// HasIssuer returns a boolean if a field has been set.
func (o *XAAResourceServer) HasIssuer() bool {
	if o != nil && !IsNil(o.Issuer) {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given string and assigns it to the Issuer field.
func (o *XAAResourceServer) SetIssuer(v string) {
	o.Issuer = &v
}

// GetProtectedResources returns the ProtectedResources field value if set, zero value otherwise.
func (o *XAAResourceServer) GetProtectedResources() []string {
	if o == nil || IsNil(o.ProtectedResources) {
		var ret []string
		return ret
	}
	return o.ProtectedResources
}

// GetProtectedResourcesOk returns a tuple with the ProtectedResources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XAAResourceServer) GetProtectedResourcesOk() ([]string, bool) {
	if o == nil || IsNil(o.ProtectedResources) {
		return nil, false
	}
	return o.ProtectedResources, true
}

// HasProtectedResources returns a boolean if a field has been set.
func (o *XAAResourceServer) HasProtectedResources() bool {
	if o != nil && !IsNil(o.ProtectedResources) {
		return true
	}

	return false
}

// SetProtectedResources gets a reference to the given []string and assigns it to the ProtectedResources field.
func (o *XAAResourceServer) SetProtectedResources(v []string) {
	o.ProtectedResources = v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *XAAResourceServer) GetScopes() []string {
	if o == nil || IsNil(o.Scopes) {
		var ret []string
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XAAResourceServer) GetScopesOk() ([]string, bool) {
	if o == nil || IsNil(o.Scopes) {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *XAAResourceServer) HasScopes() bool {
	if o != nil && !IsNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *XAAResourceServer) SetScopes(v []string) {
	o.Scopes = v
}

func (o XAAResourceServer) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o XAAResourceServer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AudTenantRequired) {
		toSerialize["audTenantRequired"] = o.AudTenantRequired
	}
	if !IsNil(o.CimdSupported) {
		toSerialize["cimdSupported"] = o.CimdSupported
	}
	if !IsNil(o.Issuer) {
		toSerialize["issuer"] = o.Issuer
	}
	if !IsNil(o.ProtectedResources) {
		toSerialize["protectedResources"] = o.ProtectedResources
	}
	if !IsNil(o.Scopes) {
		toSerialize["scopes"] = o.Scopes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *XAAResourceServer) UnmarshalJSON(data []byte) (err error) {
	varXAAResourceServer := _XAAResourceServer{}

	err = json.Unmarshal(data, &varXAAResourceServer)

	if err != nil {
		return err
	}

	*o = XAAResourceServer(varXAAResourceServer)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "audTenantRequired")
		delete(additionalProperties, "cimdSupported")
		delete(additionalProperties, "issuer")
		delete(additionalProperties, "protectedResources")
		delete(additionalProperties, "scopes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableXAAResourceServer struct {
	value *XAAResourceServer
	isSet bool
}

func (v NullableXAAResourceServer) Get() *XAAResourceServer {
	return v.value
}

func (v *NullableXAAResourceServer) Set(val *XAAResourceServer) {
	v.value = val
	v.isSet = true
}

func (v NullableXAAResourceServer) IsSet() bool {
	return v.isSet
}

func (v *NullableXAAResourceServer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableXAAResourceServer(val *XAAResourceServer) *NullableXAAResourceServer {
	return &NullableXAAResourceServer{value: val, isSet: true}
}

func (v NullableXAAResourceServer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableXAAResourceServer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
