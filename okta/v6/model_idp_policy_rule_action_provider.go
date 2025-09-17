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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the IdpPolicyRuleActionProvider type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdpPolicyRuleActionProvider{}

// IdpPolicyRuleActionProvider struct for IdpPolicyRuleActionProvider
type IdpPolicyRuleActionProvider struct {
	// IdP types of `OKTA`, `AgentlessDSSO`, and `IWA` don't require an ID.
	Id *string `json:"id,omitempty"`
	// Provider `name` in Okta. Optional. Supported in `IDENTITY ENGINE`.
	Name *string `json:"name,omitempty"`
	// The IdP object's `type` property identifies the social or enterprise IdP used for authentication. Each IdP uses a specific protocol, therefore the `protocol` object must correspond with the IdP `type`. If the protocol is OAuth 2.0-based, the `protocol` object's `scopes` property must also correspond with the scopes supported by the IdP `type`. For policy actions supported by each IdP type, see [IdP type policy actions](https://developer.okta.com/docs/api/openapi/okta-management/management/tag/IdentityProvider/#tag/IdentityProvider/operation/createIdentityProvider!path=policy&t=request).  | Type               | Description                                                                                                                                           | Corresponding protocol | Corresponding protocol scopes                                         | | ------------------ | ----------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------- | --------------------------------------------------------------------  | | `AMAZON`           | [Amazon](https://developer.amazon.com/settings/console/registration?return_to=/)&nbsp;as the IdP                                        | OpenID Connect         | `profile`, `profile:user_id`                                          | | `APPLE`            | [Apple](https://developer.apple.com/sign-in-with-apple/)&nbsp;as the IdP                                                                | OpenID Connect         | `names`, `email`, `openid`                                            | | `DISCORD`          | [Discord](https://discord.com/login)&nbsp;as the IdP                                                                                    | OAuth 2.0              | `identify`, `email`                                                   | | `FACEBOOK`         | [Facebook](https://developers.facebook.com)&nbsp;as the IdP                                                                             | OAuth 2.0              | `public_profile`, `email`                                             | | `GITHUB`           | [GitHub](https://github.com/join)&nbsp;as the IdP                                                                                       | OAuth 2.0              | `user`                                                                | | `GITLAB`           | [GitLab](https://gitlab.com/users/sign_in)&nbsp;as the IdP                                                                              | OpenID Connect         | `openid`, `read_user`, `profile`, `email`                             | | `GOOGLE`           | [Google](https://accounts.google.com/signup)&nbsp;as the IdP                                                                            | OpenID Connect         | `openid`, `email`, `profile`                                          | | `IDV_PERSONA`      | [Persona](https://app.withpersona.com/dashboard/login)&nbsp;as the IDV IdP                                                              | ID verification        |                                                                       | | `IDV_CLEAR`        | [CLEAR Verified](https://www.clearme.com/)&nbsp;as the IDV IdP                                                                          | ID verification        | `openid`, `profile`, `identity_assurance`                             | | `IDV_INCODE`       | [Incode](https://incode.com/)&nbsp;as the IDV IdP                                                                                       | ID verification        | `openid`, `profile`, `identity_assurance`                             | | `LINKEDIN`         | [LinkedIn](https://developer.linkedin.com/)&nbsp;as the IdP                                                                             | OAuth 2.0              | `r_emailaddress`, `r_liteprofile`                                     | | `LOGINGOV`         | [Login.gov](https://developers.login.gov/)&nbsp;as the IdP                                                                              | OpenID Connect         | `email`, `profile`, `profile:name`                                    | | `LOGINGOV_SANDBOX` | [Login.gov's identity sandbox](https://developers.login.gov/testing/)&nbsp;as the IdP                                                   | OpenID Connect         | `email`, `profile`, `profile:name`                                    | | `MICROSOFT`        | [Microsoft Enterprise SSO](https://azure.microsoft.com/)&nbsp;as the IdP                                                                | OpenID Connect         | `openid`, `email`, `profile`, `https://graph.microsoft.com/User.Read` | | `OIDC`             | IdP that supports [OpenID Connect](https://openid.net/specs/openid-connect-core-1_0.html)                                               | OpenID Connect         | `openid`, `email`, `profile`                                          | | `PAYPAL`           | [Paypal](https://www.paypal.com/signin)&nbsp;as the IdP                                                                                 | OpenID Connect         | `openid`, `email`, `profile`                                          | | `PAYPAL_SANDBOX`   | [Paypal Sandbox](https://developer.paypal.com/tools/sandbox/)&nbsp;as the IdP                                                           | OpenID Connect         | `openid`, `email`, `profile`                                          | | `SALESFORCE`       | [SalesForce](https://login.salesforce.com/)&nbsp;as the IdP                                                                             | OAuth 2.0              | `id`, `email`, `profile`                                              | | `SAML2`            | Enterprise IdP that supports the [SAML 2.0 Web Browser SSO Profile](https://docs.oasis-open.org/security/saml/v2.0/saml-profiles-2.0-os.pdf)| SAML 2.0  |                                                                                | | `SPOTIFY`          | [Spotify](https://developer.spotify.com/)&nbsp;as the IdP                                                                               | OpenID Connect         | `user-read-email`, `user-read-private`                                | | `X509`             | [Smart Card IdP](https://tools.ietf.org/html/rfc5280)                                                                                   | Mutual TLS             |                                                                       | | `XERO`             | [Xero](https://www.xero.com/us/signup/api/)&nbsp;as the IdP                                                                             | OpenID Connect         | `openid`, `profile`, `email`                                          | | `YAHOO`            | [Yahoo](https://login.yahoo.com/)&nbsp;as the IdP                                                                                       | OpenID Connect         | `openid`, `profile`, `email`                                          | | `YAHOOJP`          | [Yahoo Japan](https://login.yahoo.co.jp/config/login)&nbsp;as the IdP                                                                   | OpenID Connect         | `openid`, `profile`, `email`                                          | | `OKTA_INTEGRATION`             | IdP that supports the [OpenID Connect](https://openid.net/specs/openid-connect-core-1_0.html) Org2Org IdP                                               | OpenID Connect         | `openid`, `email`, `profile`                                          |
	Type                 *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IdpPolicyRuleActionProvider IdpPolicyRuleActionProvider

// NewIdpPolicyRuleActionProvider instantiates a new IdpPolicyRuleActionProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdpPolicyRuleActionProvider() *IdpPolicyRuleActionProvider {
	this := IdpPolicyRuleActionProvider{}
	return &this
}

// NewIdpPolicyRuleActionProviderWithDefaults instantiates a new IdpPolicyRuleActionProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdpPolicyRuleActionProviderWithDefaults() *IdpPolicyRuleActionProvider {
	this := IdpPolicyRuleActionProvider{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IdpPolicyRuleActionProvider) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpPolicyRuleActionProvider) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IdpPolicyRuleActionProvider) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IdpPolicyRuleActionProvider) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IdpPolicyRuleActionProvider) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpPolicyRuleActionProvider) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IdpPolicyRuleActionProvider) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IdpPolicyRuleActionProvider) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *IdpPolicyRuleActionProvider) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpPolicyRuleActionProvider) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *IdpPolicyRuleActionProvider) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *IdpPolicyRuleActionProvider) SetType(v string) {
	o.Type = &v
}

func (o IdpPolicyRuleActionProvider) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdpPolicyRuleActionProvider) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IdpPolicyRuleActionProvider) UnmarshalJSON(data []byte) (err error) {
	varIdpPolicyRuleActionProvider := _IdpPolicyRuleActionProvider{}

	err = json.Unmarshal(data, &varIdpPolicyRuleActionProvider)

	if err != nil {
		return err
	}

	*o = IdpPolicyRuleActionProvider(varIdpPolicyRuleActionProvider)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIdpPolicyRuleActionProvider struct {
	value *IdpPolicyRuleActionProvider
	isSet bool
}

func (v NullableIdpPolicyRuleActionProvider) Get() *IdpPolicyRuleActionProvider {
	return v.value
}

func (v *NullableIdpPolicyRuleActionProvider) Set(val *IdpPolicyRuleActionProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableIdpPolicyRuleActionProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableIdpPolicyRuleActionProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdpPolicyRuleActionProvider(val *IdpPolicyRuleActionProvider) *NullableIdpPolicyRuleActionProvider {
	return &NullableIdpPolicyRuleActionProvider{value: val, isSet: true}
}

func (v NullableIdpPolicyRuleActionProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdpPolicyRuleActionProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
