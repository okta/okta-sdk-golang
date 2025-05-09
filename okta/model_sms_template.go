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
	"time"
)

// SmsTemplate struct for SmsTemplate
type SmsTemplate struct {
	Created *time.Time `json:"created,omitempty"`
	Id *string `json:"id,omitempty"`
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
	// Human-readable name of the Template
	Name *string `json:"name,omitempty"`
	// Text of the Template, including any [macros](https://developer.okta.com/docs/api/openapi/okta-management/management/tag/Template/)
	Template *string `json:"template,omitempty"`
	// - Template translations are optionally provided when you want to localize the SMS messages. Translations are provided as an object that contains `key:value` pairs: the language and the translated Template text. The key portion is a two-letter country code that conforms to [ISO 639-1](https://www.loc.gov/standards/iso639-2/php/code_list.php). The value is the translated SMS Template. - Just like with regular SMS Templates, the length of the SMS message can't exceed 160 characters. 
	Translations map[string]interface{} `json:"translations,omitempty"`
	// Type of the Template
	Type *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SmsTemplate SmsTemplate

// NewSmsTemplate instantiates a new SmsTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmsTemplate() *SmsTemplate {
	this := SmsTemplate{}
	return &this
}

// NewSmsTemplateWithDefaults instantiates a new SmsTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmsTemplateWithDefaults() *SmsTemplate {
	this := SmsTemplate{}
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *SmsTemplate) GetCreated() time.Time {
	if o == nil || o.Created == nil {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsTemplate) GetCreatedOk() (*time.Time, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *SmsTemplate) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *SmsTemplate) SetCreated(v time.Time) {
	o.Created = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SmsTemplate) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsTemplate) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SmsTemplate) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SmsTemplate) SetId(v string) {
	o.Id = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *SmsTemplate) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsTemplate) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *SmsTemplate) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *SmsTemplate) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SmsTemplate) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsTemplate) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SmsTemplate) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SmsTemplate) SetName(v string) {
	o.Name = &v
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *SmsTemplate) GetTemplate() string {
	if o == nil || o.Template == nil {
		var ret string
		return ret
	}
	return *o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsTemplate) GetTemplateOk() (*string, bool) {
	if o == nil || o.Template == nil {
		return nil, false
	}
	return o.Template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *SmsTemplate) HasTemplate() bool {
	if o != nil && o.Template != nil {
		return true
	}

	return false
}

// SetTemplate gets a reference to the given string and assigns it to the Template field.
func (o *SmsTemplate) SetTemplate(v string) {
	o.Template = &v
}

// GetTranslations returns the Translations field value if set, zero value otherwise.
func (o *SmsTemplate) GetTranslations() map[string]interface{} {
	if o == nil || o.Translations == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Translations
}

// GetTranslationsOk returns a tuple with the Translations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsTemplate) GetTranslationsOk() (map[string]interface{}, bool) {
	if o == nil || o.Translations == nil {
		return nil, false
	}
	return o.Translations, true
}

// HasTranslations returns a boolean if a field has been set.
func (o *SmsTemplate) HasTranslations() bool {
	if o != nil && o.Translations != nil {
		return true
	}

	return false
}

// SetTranslations gets a reference to the given map[string]interface{} and assigns it to the Translations field.
func (o *SmsTemplate) SetTranslations(v map[string]interface{}) {
	o.Translations = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SmsTemplate) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsTemplate) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SmsTemplate) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SmsTemplate) SetType(v string) {
	o.Type = &v
}

func (o SmsTemplate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LastUpdated != nil {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Template != nil {
		toSerialize["template"] = o.Template
	}
	if o.Translations != nil {
		toSerialize["translations"] = o.Translations
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SmsTemplate) UnmarshalJSON(bytes []byte) (err error) {
	varSmsTemplate := _SmsTemplate{}

	err = json.Unmarshal(bytes, &varSmsTemplate)
	if err == nil {
		*o = SmsTemplate(varSmsTemplate)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "created")
		delete(additionalProperties, "id")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "name")
		delete(additionalProperties, "template")
		delete(additionalProperties, "translations")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableSmsTemplate struct {
	value *SmsTemplate
	isSet bool
}

func (v NullableSmsTemplate) Get() *SmsTemplate {
	return v.value
}

func (v *NullableSmsTemplate) Set(val *SmsTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableSmsTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableSmsTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmsTemplate(val *SmsTemplate) *NullableSmsTemplate {
	return &NullableSmsTemplate{value: val, isSet: true}
}

func (v NullableSmsTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmsTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

