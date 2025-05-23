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

API version: 2024.06.1
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// PolicySubject struct for PolicySubject
type PolicySubject struct {
	Filter *string `json:"filter,omitempty"`
	Format []string `json:"format,omitempty"`
	MatchAttribute *string `json:"matchAttribute,omitempty"`
	MatchType *string `json:"matchType,omitempty"`
	UserNameTemplate *PolicyUserNameTemplate `json:"userNameTemplate,omitempty"`
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
	if o == nil || o.Filter == nil {
		var ret string
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicySubject) GetFilterOk() (*string, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *PolicySubject) HasFilter() bool {
	if o != nil && o.Filter != nil {
		return true
	}

	return false
}

// SetFilter gets a reference to the given string and assigns it to the Filter field.
func (o *PolicySubject) SetFilter(v string) {
	o.Filter = &v
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *PolicySubject) GetFormat() []string {
	if o == nil || o.Format == nil {
		var ret []string
		return ret
	}
	return o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicySubject) GetFormatOk() ([]string, bool) {
	if o == nil || o.Format == nil {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *PolicySubject) HasFormat() bool {
	if o != nil && o.Format != nil {
		return true
	}

	return false
}

// SetFormat gets a reference to the given []string and assigns it to the Format field.
func (o *PolicySubject) SetFormat(v []string) {
	o.Format = v
}

// GetMatchAttribute returns the MatchAttribute field value if set, zero value otherwise.
func (o *PolicySubject) GetMatchAttribute() string {
	if o == nil || o.MatchAttribute == nil {
		var ret string
		return ret
	}
	return *o.MatchAttribute
}

// GetMatchAttributeOk returns a tuple with the MatchAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicySubject) GetMatchAttributeOk() (*string, bool) {
	if o == nil || o.MatchAttribute == nil {
		return nil, false
	}
	return o.MatchAttribute, true
}

// HasMatchAttribute returns a boolean if a field has been set.
func (o *PolicySubject) HasMatchAttribute() bool {
	if o != nil && o.MatchAttribute != nil {
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
	if o == nil || o.MatchType == nil {
		var ret string
		return ret
	}
	return *o.MatchType
}

// GetMatchTypeOk returns a tuple with the MatchType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicySubject) GetMatchTypeOk() (*string, bool) {
	if o == nil || o.MatchType == nil {
		return nil, false
	}
	return o.MatchType, true
}

// HasMatchType returns a boolean if a field has been set.
func (o *PolicySubject) HasMatchType() bool {
	if o != nil && o.MatchType != nil {
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
	if o == nil || o.UserNameTemplate == nil {
		var ret PolicyUserNameTemplate
		return ret
	}
	return *o.UserNameTemplate
}

// GetUserNameTemplateOk returns a tuple with the UserNameTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicySubject) GetUserNameTemplateOk() (*PolicyUserNameTemplate, bool) {
	if o == nil || o.UserNameTemplate == nil {
		return nil, false
	}
	return o.UserNameTemplate, true
}

// HasUserNameTemplate returns a boolean if a field has been set.
func (o *PolicySubject) HasUserNameTemplate() bool {
	if o != nil && o.UserNameTemplate != nil {
		return true
	}

	return false
}

// SetUserNameTemplate gets a reference to the given PolicyUserNameTemplate and assigns it to the UserNameTemplate field.
func (o *PolicySubject) SetUserNameTemplate(v PolicyUserNameTemplate) {
	o.UserNameTemplate = &v
}

func (o PolicySubject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	if o.Format != nil {
		toSerialize["format"] = o.Format
	}
	if o.MatchAttribute != nil {
		toSerialize["matchAttribute"] = o.MatchAttribute
	}
	if o.MatchType != nil {
		toSerialize["matchType"] = o.MatchType
	}
	if o.UserNameTemplate != nil {
		toSerialize["userNameTemplate"] = o.UserNameTemplate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PolicySubject) UnmarshalJSON(bytes []byte) (err error) {
	varPolicySubject := _PolicySubject{}

	err = json.Unmarshal(bytes, &varPolicySubject)
	if err == nil {
		*o = PolicySubject(varPolicySubject)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "filter")
		delete(additionalProperties, "format")
		delete(additionalProperties, "matchAttribute")
		delete(additionalProperties, "matchType")
		delete(additionalProperties, "userNameTemplate")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
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

