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

// checks if the PolicyUserNameTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PolicyUserNameTemplate{}

// PolicyUserNameTemplate [Okta Expression Language (EL) expression](https://developer.okta.com/docs/reference/okta-expression-language/) to generate or transform a unique username for the IdP user. * IdP user profile attributes can be referenced with the `idpuser` prefix such as `idpuser.subjectNameId`. * You must define an IdP user profile attribute before it can be referenced in an Okta EL expression. To define an IdP user attribute policy, you may need to create a new IdP instance without a base profile property. Then edit the IdP user profile to update the IdP instance with an expression that references the IdP user profile attribute that you just created.
type PolicyUserNameTemplate struct {
	Template             *string `json:"template,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PolicyUserNameTemplate PolicyUserNameTemplate

// NewPolicyUserNameTemplate instantiates a new PolicyUserNameTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyUserNameTemplate() *PolicyUserNameTemplate {
	this := PolicyUserNameTemplate{}
	return &this
}

// NewPolicyUserNameTemplateWithDefaults instantiates a new PolicyUserNameTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyUserNameTemplateWithDefaults() *PolicyUserNameTemplate {
	this := PolicyUserNameTemplate{}
	return &this
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *PolicyUserNameTemplate) GetTemplate() string {
	if o == nil || IsNil(o.Template) {
		var ret string
		return ret
	}
	return *o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyUserNameTemplate) GetTemplateOk() (*string, bool) {
	if o == nil || IsNil(o.Template) {
		return nil, false
	}
	return o.Template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *PolicyUserNameTemplate) HasTemplate() bool {
	if o != nil && !IsNil(o.Template) {
		return true
	}

	return false
}

// SetTemplate gets a reference to the given string and assigns it to the Template field.
func (o *PolicyUserNameTemplate) SetTemplate(v string) {
	o.Template = &v
}

func (o PolicyUserNameTemplate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PolicyUserNameTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Template) {
		toSerialize["template"] = o.Template
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PolicyUserNameTemplate) UnmarshalJSON(data []byte) (err error) {
	varPolicyUserNameTemplate := _PolicyUserNameTemplate{}

	err = json.Unmarshal(data, &varPolicyUserNameTemplate)

	if err != nil {
		return err
	}

	*o = PolicyUserNameTemplate(varPolicyUserNameTemplate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "template")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePolicyUserNameTemplate struct {
	value *PolicyUserNameTemplate
	isSet bool
}

func (v NullablePolicyUserNameTemplate) Get() *PolicyUserNameTemplate {
	return v.value
}

func (v *NullablePolicyUserNameTemplate) Set(val *PolicyUserNameTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyUserNameTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyUserNameTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyUserNameTemplate(val *PolicyUserNameTemplate) *NullablePolicyUserNameTemplate {
	return &NullablePolicyUserNameTemplate{value: val, isSet: true}
}

func (v NullablePolicyUserNameTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyUserNameTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
