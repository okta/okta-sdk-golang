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

// checks if the PostAuthKeepMeSignedInPrompt type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostAuthKeepMeSignedInPrompt{}

// PostAuthKeepMeSignedInPrompt struct for PostAuthKeepMeSignedInPrompt
type PostAuthKeepMeSignedInPrompt struct {
	// The label on the accept button when prompting for Stay signed in
	AcceptButtonText *string `json:"acceptButtonText,omitempty"`
	// The label on the reject button when prompting for Stay signed in
	RejectButtonText *string `json:"rejectButtonText,omitempty"`
	// The subtitle on the Sign-In Widget when prompting for Stay signed in
	Subtitle *string `json:"subtitle,omitempty"`
	// The title on the Sign-In Widget when prompting for Stay signed in
	Title                *string `json:"title,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PostAuthKeepMeSignedInPrompt PostAuthKeepMeSignedInPrompt

// NewPostAuthKeepMeSignedInPrompt instantiates a new PostAuthKeepMeSignedInPrompt object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostAuthKeepMeSignedInPrompt() *PostAuthKeepMeSignedInPrompt {
	this := PostAuthKeepMeSignedInPrompt{}
	return &this
}

// NewPostAuthKeepMeSignedInPromptWithDefaults instantiates a new PostAuthKeepMeSignedInPrompt object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostAuthKeepMeSignedInPromptWithDefaults() *PostAuthKeepMeSignedInPrompt {
	this := PostAuthKeepMeSignedInPrompt{}
	return &this
}

// GetAcceptButtonText returns the AcceptButtonText field value if set, zero value otherwise.
func (o *PostAuthKeepMeSignedInPrompt) GetAcceptButtonText() string {
	if o == nil || IsNil(o.AcceptButtonText) {
		var ret string
		return ret
	}
	return *o.AcceptButtonText
}

// GetAcceptButtonTextOk returns a tuple with the AcceptButtonText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostAuthKeepMeSignedInPrompt) GetAcceptButtonTextOk() (*string, bool) {
	if o == nil || IsNil(o.AcceptButtonText) {
		return nil, false
	}
	return o.AcceptButtonText, true
}

// HasAcceptButtonText returns a boolean if a field has been set.
func (o *PostAuthKeepMeSignedInPrompt) HasAcceptButtonText() bool {
	if o != nil && !IsNil(o.AcceptButtonText) {
		return true
	}

	return false
}

// SetAcceptButtonText gets a reference to the given string and assigns it to the AcceptButtonText field.
func (o *PostAuthKeepMeSignedInPrompt) SetAcceptButtonText(v string) {
	o.AcceptButtonText = &v
}

// GetRejectButtonText returns the RejectButtonText field value if set, zero value otherwise.
func (o *PostAuthKeepMeSignedInPrompt) GetRejectButtonText() string {
	if o == nil || IsNil(o.RejectButtonText) {
		var ret string
		return ret
	}
	return *o.RejectButtonText
}

// GetRejectButtonTextOk returns a tuple with the RejectButtonText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostAuthKeepMeSignedInPrompt) GetRejectButtonTextOk() (*string, bool) {
	if o == nil || IsNil(o.RejectButtonText) {
		return nil, false
	}
	return o.RejectButtonText, true
}

// HasRejectButtonText returns a boolean if a field has been set.
func (o *PostAuthKeepMeSignedInPrompt) HasRejectButtonText() bool {
	if o != nil && !IsNil(o.RejectButtonText) {
		return true
	}

	return false
}

// SetRejectButtonText gets a reference to the given string and assigns it to the RejectButtonText field.
func (o *PostAuthKeepMeSignedInPrompt) SetRejectButtonText(v string) {
	o.RejectButtonText = &v
}

// GetSubtitle returns the Subtitle field value if set, zero value otherwise.
func (o *PostAuthKeepMeSignedInPrompt) GetSubtitle() string {
	if o == nil || IsNil(o.Subtitle) {
		var ret string
		return ret
	}
	return *o.Subtitle
}

// GetSubtitleOk returns a tuple with the Subtitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostAuthKeepMeSignedInPrompt) GetSubtitleOk() (*string, bool) {
	if o == nil || IsNil(o.Subtitle) {
		return nil, false
	}
	return o.Subtitle, true
}

// HasSubtitle returns a boolean if a field has been set.
func (o *PostAuthKeepMeSignedInPrompt) HasSubtitle() bool {
	if o != nil && !IsNil(o.Subtitle) {
		return true
	}

	return false
}

// SetSubtitle gets a reference to the given string and assigns it to the Subtitle field.
func (o *PostAuthKeepMeSignedInPrompt) SetSubtitle(v string) {
	o.Subtitle = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *PostAuthKeepMeSignedInPrompt) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostAuthKeepMeSignedInPrompt) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *PostAuthKeepMeSignedInPrompt) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *PostAuthKeepMeSignedInPrompt) SetTitle(v string) {
	o.Title = &v
}

func (o PostAuthKeepMeSignedInPrompt) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostAuthKeepMeSignedInPrompt) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AcceptButtonText) {
		toSerialize["acceptButtonText"] = o.AcceptButtonText
	}
	if !IsNil(o.RejectButtonText) {
		toSerialize["rejectButtonText"] = o.RejectButtonText
	}
	if !IsNil(o.Subtitle) {
		toSerialize["subtitle"] = o.Subtitle
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PostAuthKeepMeSignedInPrompt) UnmarshalJSON(data []byte) (err error) {
	varPostAuthKeepMeSignedInPrompt := _PostAuthKeepMeSignedInPrompt{}

	err = json.Unmarshal(data, &varPostAuthKeepMeSignedInPrompt)

	if err != nil {
		return err
	}

	*o = PostAuthKeepMeSignedInPrompt(varPostAuthKeepMeSignedInPrompt)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "acceptButtonText")
		delete(additionalProperties, "rejectButtonText")
		delete(additionalProperties, "subtitle")
		delete(additionalProperties, "title")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePostAuthKeepMeSignedInPrompt struct {
	value *PostAuthKeepMeSignedInPrompt
	isSet bool
}

func (v NullablePostAuthKeepMeSignedInPrompt) Get() *PostAuthKeepMeSignedInPrompt {
	return v.value
}

func (v *NullablePostAuthKeepMeSignedInPrompt) Set(val *PostAuthKeepMeSignedInPrompt) {
	v.value = val
	v.isSet = true
}

func (v NullablePostAuthKeepMeSignedInPrompt) IsSet() bool {
	return v.isSet
}

func (v *NullablePostAuthKeepMeSignedInPrompt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostAuthKeepMeSignedInPrompt(val *PostAuthKeepMeSignedInPrompt) *NullablePostAuthKeepMeSignedInPrompt {
	return &NullablePostAuthKeepMeSignedInPrompt{value: val, isSet: true}
}

func (v NullablePostAuthKeepMeSignedInPrompt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostAuthKeepMeSignedInPrompt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
