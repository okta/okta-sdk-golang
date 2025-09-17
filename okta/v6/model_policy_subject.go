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

// checks if the PolicySubject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PolicySubject{}

// PolicySubject Specifies the behavior for establishing, validating, and matching a username for an IdP user
type PolicySubject struct {
	// Optional [regular expression pattern](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Regular_expressions) used to filter untrusted IdP usernames. * As a best security practice, you should define a regular expression pattern to filter untrusted IdP usernames. This is especially important if multiple IdPs are connected to your org. The filter prevents an IdP from issuing an assertion for any user, including partners or directory users in your Okta org. * For example, the filter pattern `(\\S+@example\\.com)` allows only Users that have an `@example.com` username suffix. It rejects assertions that have any other suffix such as `@corp.example.com` or `@partner.com`. * Only `SAML2` and `OIDC` IdP providers support the `filter` property.
	Filter *string `json:"filter,omitempty"`
	// Okta user profile attribute for matching a transformed IdP username. Only for matchType `CUSTOM_ATTRIBUTE`. The `matchAttribute` must be a valid Okta user profile attribute of one of the following types: * String (with no format or 'email' format only) * Integer * Number
	MatchAttribute *string `json:"matchAttribute,omitempty"`
	// Determines the Okta user profile attribute match conditions for account linking and authentication of the transformed IdP username
	MatchType            *string                 `json:"matchType,omitempty"`
	UserNameTemplate     *PolicyUserNameTemplate `json:"userNameTemplate,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PolicySubject PolicySubject

// NewPolicySubject instantiates a new PolicySubject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicySubject() *PolicySubject {
	this := PolicySubject{}
	return &this
}

// NewPolicySubjectWithDefaults instantiates a new PolicySubject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicySubjectWithDefaults() *PolicySubject {
	this := PolicySubject{}
	return &this
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *PolicySubject) GetFilter() string {
	if o == nil || IsNil(o.Filter) {
		var ret string
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicySubject) GetFilterOk() (*string, bool) {
	if o == nil || IsNil(o.Filter) {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *PolicySubject) HasFilter() bool {
	if o != nil && !IsNil(o.Filter) {
		return true
	}

	return false
}

// SetFilter gets a reference to the given string and assigns it to the Filter field.
func (o *PolicySubject) SetFilter(v string) {
	o.Filter = &v
}

// GetMatchAttribute returns the MatchAttribute field value if set, zero value otherwise.
func (o *PolicySubject) GetMatchAttribute() string {
	if o == nil || IsNil(o.MatchAttribute) {
		var ret string
		return ret
	}
	return *o.MatchAttribute
}

// GetMatchAttributeOk returns a tuple with the MatchAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicySubject) GetMatchAttributeOk() (*string, bool) {
	if o == nil || IsNil(o.MatchAttribute) {
		return nil, false
	}
	return o.MatchAttribute, true
}

// HasMatchAttribute returns a boolean if a field has been set.
func (o *PolicySubject) HasMatchAttribute() bool {
	if o != nil && !IsNil(o.MatchAttribute) {
		return true
	}

	return false
}

// SetMatchAttribute gets a reference to the given string and assigns it to the MatchAttribute field.
func (o *PolicySubject) SetMatchAttribute(v string) {
	o.MatchAttribute = &v
}

// GetMatchType returns the MatchType field value if set, zero value otherwise.
func (o *PolicySubject) GetMatchType() string {
	if o == nil || IsNil(o.MatchType) {
		var ret string
		return ret
	}
	return *o.MatchType
}

// GetMatchTypeOk returns a tuple with the MatchType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicySubject) GetMatchTypeOk() (*string, bool) {
	if o == nil || IsNil(o.MatchType) {
		return nil, false
	}
	return o.MatchType, true
}

// HasMatchType returns a boolean if a field has been set.
func (o *PolicySubject) HasMatchType() bool {
	if o != nil && !IsNil(o.MatchType) {
		return true
	}

	return false
}

// SetMatchType gets a reference to the given string and assigns it to the MatchType field.
func (o *PolicySubject) SetMatchType(v string) {
	o.MatchType = &v
}

// GetUserNameTemplate returns the UserNameTemplate field value if set, zero value otherwise.
func (o *PolicySubject) GetUserNameTemplate() PolicyUserNameTemplate {
	if o == nil || IsNil(o.UserNameTemplate) {
		var ret PolicyUserNameTemplate
		return ret
	}
	return *o.UserNameTemplate
}

// GetUserNameTemplateOk returns a tuple with the UserNameTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicySubject) GetUserNameTemplateOk() (*PolicyUserNameTemplate, bool) {
	if o == nil || IsNil(o.UserNameTemplate) {
		return nil, false
	}
	return o.UserNameTemplate, true
}

// HasUserNameTemplate returns a boolean if a field has been set.
func (o *PolicySubject) HasUserNameTemplate() bool {
	if o != nil && !IsNil(o.UserNameTemplate) {
		return true
	}

	return false
}

// SetUserNameTemplate gets a reference to the given PolicyUserNameTemplate and assigns it to the UserNameTemplate field.
func (o *PolicySubject) SetUserNameTemplate(v PolicyUserNameTemplate) {
	o.UserNameTemplate = &v
}

func (o PolicySubject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PolicySubject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Filter) {
		toSerialize["filter"] = o.Filter
	}
	if !IsNil(o.MatchAttribute) {
		toSerialize["matchAttribute"] = o.MatchAttribute
	}
	if !IsNil(o.MatchType) {
		toSerialize["matchType"] = o.MatchType
	}
	if !IsNil(o.UserNameTemplate) {
		toSerialize["userNameTemplate"] = o.UserNameTemplate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PolicySubject) UnmarshalJSON(data []byte) (err error) {
	varPolicySubject := _PolicySubject{}

	err = json.Unmarshal(data, &varPolicySubject)

	if err != nil {
		return err
	}

	*o = PolicySubject(varPolicySubject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "filter")
		delete(additionalProperties, "matchAttribute")
		delete(additionalProperties, "matchType")
		delete(additionalProperties, "userNameTemplate")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePolicySubject struct {
	value *PolicySubject
	isSet bool
}

func (v NullablePolicySubject) Get() *PolicySubject {
	return v.value
}

func (v *NullablePolicySubject) Set(val *PolicySubject) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicySubject) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicySubject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicySubject(val *PolicySubject) *NullablePolicySubject {
	return &NullablePolicySubject{value: val, isSet: true}
}

func (v NullablePolicySubject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicySubject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
