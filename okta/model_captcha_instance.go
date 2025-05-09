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

// CAPTCHAInstance 
type CAPTCHAInstance struct {
	// The unique key for the CAPTCHA instance
	Id *string `json:"id,omitempty"`
	// The name of the CAPTCHA instance
	Name *string `json:"name,omitempty"`
	// The secret key issued from the CAPTCHA provider to perform server-side validation for a CAPTCHA token
	SecretKey *string `json:"secretKey,omitempty"`
	// The site key issued from the CAPTCHA provider to render a CAPTCHA on a page
	SiteKey *string `json:"siteKey,omitempty"`
	// The type of CAPTCHA provider
	Type *string `json:"type,omitempty"`
	Links *LinksSelf `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CAPTCHAInstance CAPTCHAInstance

// NewCAPTCHAInstance instantiates a new CAPTCHAInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCAPTCHAInstance() *CAPTCHAInstance {
	this := CAPTCHAInstance{}
	return &this
}

// NewCAPTCHAInstanceWithDefaults instantiates a new CAPTCHAInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCAPTCHAInstanceWithDefaults() *CAPTCHAInstance {
	this := CAPTCHAInstance{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CAPTCHAInstance) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CAPTCHAInstance) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CAPTCHAInstance) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CAPTCHAInstance) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CAPTCHAInstance) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CAPTCHAInstance) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CAPTCHAInstance) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CAPTCHAInstance) SetName(v string) {
	o.Name = &v
}

// GetSecretKey returns the SecretKey field value if set, zero value otherwise.
func (o *CAPTCHAInstance) GetSecretKey() string {
	if o == nil || o.SecretKey == nil {
		var ret string
		return ret
	}
	return *o.SecretKey
}

// GetSecretKeyOk returns a tuple with the SecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CAPTCHAInstance) GetSecretKeyOk() (*string, bool) {
	if o == nil || o.SecretKey == nil {
		return nil, false
	}
	return o.SecretKey, true
}

// HasSecretKey returns a boolean if a field has been set.
func (o *CAPTCHAInstance) HasSecretKey() bool {
	if o != nil && o.SecretKey != nil {
		return true
	}

	return false
}

// SetSecretKey gets a reference to the given string and assigns it to the SecretKey field.
func (o *CAPTCHAInstance) SetSecretKey(v string) {
	o.SecretKey = &v
}

// GetSiteKey returns the SiteKey field value if set, zero value otherwise.
func (o *CAPTCHAInstance) GetSiteKey() string {
	if o == nil || o.SiteKey == nil {
		var ret string
		return ret
	}
	return *o.SiteKey
}

// GetSiteKeyOk returns a tuple with the SiteKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CAPTCHAInstance) GetSiteKeyOk() (*string, bool) {
	if o == nil || o.SiteKey == nil {
		return nil, false
	}
	return o.SiteKey, true
}

// HasSiteKey returns a boolean if a field has been set.
func (o *CAPTCHAInstance) HasSiteKey() bool {
	if o != nil && o.SiteKey != nil {
		return true
	}

	return false
}

// SetSiteKey gets a reference to the given string and assigns it to the SiteKey field.
func (o *CAPTCHAInstance) SetSiteKey(v string) {
	o.SiteKey = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CAPTCHAInstance) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CAPTCHAInstance) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CAPTCHAInstance) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CAPTCHAInstance) SetType(v string) {
	o.Type = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *CAPTCHAInstance) GetLinks() LinksSelf {
	if o == nil || o.Links == nil {
		var ret LinksSelf
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CAPTCHAInstance) GetLinksOk() (*LinksSelf, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *CAPTCHAInstance) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinksSelf and assigns it to the Links field.
func (o *CAPTCHAInstance) SetLinks(v LinksSelf) {
	o.Links = &v
}

func (o CAPTCHAInstance) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.SecretKey != nil {
		toSerialize["secretKey"] = o.SecretKey
	}
	if o.SiteKey != nil {
		toSerialize["siteKey"] = o.SiteKey
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CAPTCHAInstance) UnmarshalJSON(bytes []byte) (err error) {
	varCAPTCHAInstance := _CAPTCHAInstance{}

	err = json.Unmarshal(bytes, &varCAPTCHAInstance)
	if err == nil {
		*o = CAPTCHAInstance(varCAPTCHAInstance)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "secretKey")
		delete(additionalProperties, "siteKey")
		delete(additionalProperties, "type")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableCAPTCHAInstance struct {
	value *CAPTCHAInstance
	isSet bool
}

func (v NullableCAPTCHAInstance) Get() *CAPTCHAInstance {
	return v.value
}

func (v *NullableCAPTCHAInstance) Set(val *CAPTCHAInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableCAPTCHAInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableCAPTCHAInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCAPTCHAInstance(val *CAPTCHAInstance) *NullableCAPTCHAInstance {
	return &NullableCAPTCHAInstance{value: val, isSet: true}
}

func (v NullableCAPTCHAInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCAPTCHAInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

