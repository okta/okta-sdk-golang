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

// OAuth2Claim struct for OAuth2Claim
type OAuth2Claim struct {
	// Specifies whether to include Claims in the token. The value is always `TRUE` for access token Claims. If the value is set to `FALSE` for an ID token claim, the Claim isn't included in the ID token when the token is requested with the access token or with the `authorization_code`. The client instead uses the access token to get Claims from the `/userinfo` endpoint.
	AlwaysIncludeInToken *bool `json:"alwaysIncludeInToken,omitempty"`
	// Specifies whether the Claim is for an access token (`RESOURCE`) or an ID token (`IDENTITY`)
	ClaimType *string `json:"claimType,omitempty"`
	Conditions *OAuth2ClaimConditions `json:"conditions,omitempty"`
	// Specifies the type of group filter if `valueType` is `GROUPS`  If `valueType` is `GROUPS`, then the groups returned are filtered according to the value of `group_filter_type`.  If you have complex filters for Groups, you can [create a Groups allowlist](https://developer.okta.com/docs/guides/customize-tokens-groups-claim/main/) to put them all in a Claim.
	GroupFilterType *string `json:"group_filter_type,omitempty"`
	// ID of the Claim
	Id *string `json:"id,omitempty"`
	// Name of the Claim
	Name *string `json:"name,omitempty"`
	Status *string `json:"status,omitempty"`
	// When `true`, indicates that Okta created the Claim
	System *bool `json:"system,omitempty"`
	// Specifies the value of the Claim. This value must be a string literal if `valueType` is `GROUPS`, and the string literal is matched with the selected `group_filter_type`. The value must be an Okta EL expression if `valueType` is `EXPRESSION`.
	Value *string `json:"value,omitempty"`
	// Specifies whether the Claim is an Okta Expression Language (EL) expression (`EXPRESSION`), a set of groups (`GROUPS`), or a system claim (`SYSTEM`)
	ValueType *string `json:"valueType,omitempty"`
	Links *LinksSelf `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OAuth2Claim OAuth2Claim

// NewOAuth2Claim instantiates a new OAuth2Claim object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuth2Claim() *OAuth2Claim {
	this := OAuth2Claim{}
	return &this
}

// NewOAuth2ClaimWithDefaults instantiates a new OAuth2Claim object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuth2ClaimWithDefaults() *OAuth2Claim {
	this := OAuth2Claim{}
	return &this
}

// GetAlwaysIncludeInToken returns the AlwaysIncludeInToken field value if set, zero value otherwise.
func (o *OAuth2Claim) GetAlwaysIncludeInToken() bool {
	if o == nil || o.AlwaysIncludeInToken == nil {
		var ret bool
		return ret
	}
	return *o.AlwaysIncludeInToken
}

// GetAlwaysIncludeInTokenOk returns a tuple with the AlwaysIncludeInToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2Claim) GetAlwaysIncludeInTokenOk() (*bool, bool) {
	if o == nil || o.AlwaysIncludeInToken == nil {
		return nil, false
	}
	return o.AlwaysIncludeInToken, true
}

// HasAlwaysIncludeInToken returns a boolean if a field has been set.
func (o *OAuth2Claim) HasAlwaysIncludeInToken() bool {
	if o != nil && o.AlwaysIncludeInToken != nil {
		return true
	}

	return false
}

// SetAlwaysIncludeInToken gets a reference to the given bool and assigns it to the AlwaysIncludeInToken field.
func (o *OAuth2Claim) SetAlwaysIncludeInToken(v bool) {
	o.AlwaysIncludeInToken = &v
}

// GetClaimType returns the ClaimType field value if set, zero value otherwise.
func (o *OAuth2Claim) GetClaimType() string {
	if o == nil || o.ClaimType == nil {
		var ret string
		return ret
	}
	return *o.ClaimType
}

// GetClaimTypeOk returns a tuple with the ClaimType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2Claim) GetClaimTypeOk() (*string, bool) {
	if o == nil || o.ClaimType == nil {
		return nil, false
	}
	return o.ClaimType, true
}

// HasClaimType returns a boolean if a field has been set.
func (o *OAuth2Claim) HasClaimType() bool {
	if o != nil && o.ClaimType != nil {
		return true
	}

	return false
}

// SetClaimType gets a reference to the given string and assigns it to the ClaimType field.
func (o *OAuth2Claim) SetClaimType(v string) {
	o.ClaimType = &v
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *OAuth2Claim) GetConditions() OAuth2ClaimConditions {
	if o == nil || o.Conditions == nil {
		var ret OAuth2ClaimConditions
		return ret
	}
	return *o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2Claim) GetConditionsOk() (*OAuth2ClaimConditions, bool) {
	if o == nil || o.Conditions == nil {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *OAuth2Claim) HasConditions() bool {
	if o != nil && o.Conditions != nil {
		return true
	}

	return false
}

// SetConditions gets a reference to the given OAuth2ClaimConditions and assigns it to the Conditions field.
func (o *OAuth2Claim) SetConditions(v OAuth2ClaimConditions) {
	o.Conditions = &v
}

// GetGroupFilterType returns the GroupFilterType field value if set, zero value otherwise.
func (o *OAuth2Claim) GetGroupFilterType() string {
	if o == nil || o.GroupFilterType == nil {
		var ret string
		return ret
	}
	return *o.GroupFilterType
}

// GetGroupFilterTypeOk returns a tuple with the GroupFilterType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2Claim) GetGroupFilterTypeOk() (*string, bool) {
	if o == nil || o.GroupFilterType == nil {
		return nil, false
	}
	return o.GroupFilterType, true
}

// HasGroupFilterType returns a boolean if a field has been set.
func (o *OAuth2Claim) HasGroupFilterType() bool {
	if o != nil && o.GroupFilterType != nil {
		return true
	}

	return false
}

// SetGroupFilterType gets a reference to the given string and assigns it to the GroupFilterType field.
func (o *OAuth2Claim) SetGroupFilterType(v string) {
	o.GroupFilterType = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OAuth2Claim) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2Claim) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OAuth2Claim) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OAuth2Claim) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OAuth2Claim) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2Claim) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OAuth2Claim) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OAuth2Claim) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OAuth2Claim) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2Claim) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OAuth2Claim) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *OAuth2Claim) SetStatus(v string) {
	o.Status = &v
}

// GetSystem returns the System field value if set, zero value otherwise.
func (o *OAuth2Claim) GetSystem() bool {
	if o == nil || o.System == nil {
		var ret bool
		return ret
	}
	return *o.System
}

// GetSystemOk returns a tuple with the System field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2Claim) GetSystemOk() (*bool, bool) {
	if o == nil || o.System == nil {
		return nil, false
	}
	return o.System, true
}

// HasSystem returns a boolean if a field has been set.
func (o *OAuth2Claim) HasSystem() bool {
	if o != nil && o.System != nil {
		return true
	}

	return false
}

// SetSystem gets a reference to the given bool and assigns it to the System field.
func (o *OAuth2Claim) SetSystem(v bool) {
	o.System = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *OAuth2Claim) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2Claim) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *OAuth2Claim) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *OAuth2Claim) SetValue(v string) {
	o.Value = &v
}

// GetValueType returns the ValueType field value if set, zero value otherwise.
func (o *OAuth2Claim) GetValueType() string {
	if o == nil || o.ValueType == nil {
		var ret string
		return ret
	}
	return *o.ValueType
}

// GetValueTypeOk returns a tuple with the ValueType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2Claim) GetValueTypeOk() (*string, bool) {
	if o == nil || o.ValueType == nil {
		return nil, false
	}
	return o.ValueType, true
}

// HasValueType returns a boolean if a field has been set.
func (o *OAuth2Claim) HasValueType() bool {
	if o != nil && o.ValueType != nil {
		return true
	}

	return false
}

// SetValueType gets a reference to the given string and assigns it to the ValueType field.
func (o *OAuth2Claim) SetValueType(v string) {
	o.ValueType = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *OAuth2Claim) GetLinks() LinksSelf {
	if o == nil || o.Links == nil {
		var ret LinksSelf
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2Claim) GetLinksOk() (*LinksSelf, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *OAuth2Claim) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinksSelf and assigns it to the Links field.
func (o *OAuth2Claim) SetLinks(v LinksSelf) {
	o.Links = &v
}

func (o OAuth2Claim) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AlwaysIncludeInToken != nil {
		toSerialize["alwaysIncludeInToken"] = o.AlwaysIncludeInToken
	}
	if o.ClaimType != nil {
		toSerialize["claimType"] = o.ClaimType
	}
	if o.Conditions != nil {
		toSerialize["conditions"] = o.Conditions
	}
	if o.GroupFilterType != nil {
		toSerialize["group_filter_type"] = o.GroupFilterType
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.System != nil {
		toSerialize["system"] = o.System
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.ValueType != nil {
		toSerialize["valueType"] = o.ValueType
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OAuth2Claim) UnmarshalJSON(bytes []byte) (err error) {
	varOAuth2Claim := _OAuth2Claim{}

	err = json.Unmarshal(bytes, &varOAuth2Claim)
	if err == nil {
		*o = OAuth2Claim(varOAuth2Claim)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "alwaysIncludeInToken")
		delete(additionalProperties, "claimType")
		delete(additionalProperties, "conditions")
		delete(additionalProperties, "group_filter_type")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "status")
		delete(additionalProperties, "system")
		delete(additionalProperties, "value")
		delete(additionalProperties, "valueType")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableOAuth2Claim struct {
	value *OAuth2Claim
	isSet bool
}

func (v NullableOAuth2Claim) Get() *OAuth2Claim {
	return v.value
}

func (v *NullableOAuth2Claim) Set(val *OAuth2Claim) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuth2Claim) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuth2Claim) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuth2Claim(val *OAuth2Claim) *NullableOAuth2Claim {
	return &NullableOAuth2Claim{value: val, isSet: true}
}

func (v NullableOAuth2Claim) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuth2Claim) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

