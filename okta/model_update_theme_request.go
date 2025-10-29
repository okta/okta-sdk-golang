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
	"fmt"
)

// checks if the UpdateThemeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateThemeRequest{}

// UpdateThemeRequest struct for UpdateThemeRequest
type UpdateThemeRequest struct {
	// Variant for email templates. You can publish a theme for email templates with different combinations of assets. Variants are preset combinations of those assets.
	EmailTemplateTouchPointVariant string `json:"emailTemplateTouchPointVariant"`
	// Variant for the Okta End-User Dashboard. You can publish a theme for end-user dashboard with different combinations of assets. Variants are preset combinations of those assets.
	EndUserDashboardTouchPointVariant string `json:"endUserDashboardTouchPointVariant"`
	// Variant for the error page. You can publish a theme for error page with different combinations of assets. Variants are preset combinations of those assets.
	ErrorPageTouchPointVariant string `json:"errorPageTouchPointVariant"`
	// Variant for the Okta loading page. You can publish a theme for Okta loading page with different combinations of assets. Variants are preset combinations of those assets.
	LoadingPageTouchPointVariant *string `json:"loadingPageTouchPointVariant,omitempty"`
	// Primary color contrast hex code
	PrimaryColorContrastHex *string `json:"primaryColorContrastHex,omitempty"`
	// Primary color hex code
	PrimaryColorHex string `json:"primaryColorHex"`
	// Secondary color contrast hex code
	SecondaryColorContrastHex *string `json:"secondaryColorContrastHex,omitempty"`
	// Secondary color hex code
	SecondaryColorHex string `json:"secondaryColorHex"`
	// Variant for the Okta sign-in page. You can publish a theme for sign-in page with different combinations of assets. Variants are preset combinations of those assets. > **Note:**  For a non-`OKTA_DEFAULT` variant, `primaryColorHex` is used for button background color and `primaryColorContrastHex` is used to optimize the opacity for button text.
	SignInPageTouchPointVariant string     `json:"signInPageTouchPointVariant"`
	Links                       *LinksSelf `json:"_links,omitempty"`
	AdditionalProperties        map[string]interface{}
}

type _UpdateThemeRequest UpdateThemeRequest

// NewUpdateThemeRequest instantiates a new UpdateThemeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateThemeRequest(emailTemplateTouchPointVariant string, endUserDashboardTouchPointVariant string, errorPageTouchPointVariant string, primaryColorHex string, secondaryColorHex string, signInPageTouchPointVariant string) *UpdateThemeRequest {
	this := UpdateThemeRequest{}
	this.EmailTemplateTouchPointVariant = emailTemplateTouchPointVariant
	this.EndUserDashboardTouchPointVariant = endUserDashboardTouchPointVariant
	this.ErrorPageTouchPointVariant = errorPageTouchPointVariant
	var loadingPageTouchPointVariant string = "OKTA_DEFAULT"
	this.LoadingPageTouchPointVariant = &loadingPageTouchPointVariant
	this.PrimaryColorHex = primaryColorHex
	this.SecondaryColorHex = secondaryColorHex
	this.SignInPageTouchPointVariant = signInPageTouchPointVariant
	return &this
}

// NewUpdateThemeRequestWithDefaults instantiates a new UpdateThemeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateThemeRequestWithDefaults() *UpdateThemeRequest {
	this := UpdateThemeRequest{}
	var emailTemplateTouchPointVariant string = "OKTA_DEFAULT"
	this.EmailTemplateTouchPointVariant = emailTemplateTouchPointVariant
	var endUserDashboardTouchPointVariant string = "OKTA_DEFAULT"
	this.EndUserDashboardTouchPointVariant = endUserDashboardTouchPointVariant
	var errorPageTouchPointVariant string = "OKTA_DEFAULT"
	this.ErrorPageTouchPointVariant = errorPageTouchPointVariant
	var loadingPageTouchPointVariant string = "OKTA_DEFAULT"
	this.LoadingPageTouchPointVariant = &loadingPageTouchPointVariant
	return &this
}

// GetEmailTemplateTouchPointVariant returns the EmailTemplateTouchPointVariant field value
func (o *UpdateThemeRequest) GetEmailTemplateTouchPointVariant() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EmailTemplateTouchPointVariant
}

// GetEmailTemplateTouchPointVariantOk returns a tuple with the EmailTemplateTouchPointVariant field value
// and a boolean to check if the value has been set.
func (o *UpdateThemeRequest) GetEmailTemplateTouchPointVariantOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EmailTemplateTouchPointVariant, true
}

// SetEmailTemplateTouchPointVariant sets field value
func (o *UpdateThemeRequest) SetEmailTemplateTouchPointVariant(v string) {
	o.EmailTemplateTouchPointVariant = v
}

// GetEndUserDashboardTouchPointVariant returns the EndUserDashboardTouchPointVariant field value
func (o *UpdateThemeRequest) GetEndUserDashboardTouchPointVariant() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndUserDashboardTouchPointVariant
}

// GetEndUserDashboardTouchPointVariantOk returns a tuple with the EndUserDashboardTouchPointVariant field value
// and a boolean to check if the value has been set.
func (o *UpdateThemeRequest) GetEndUserDashboardTouchPointVariantOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndUserDashboardTouchPointVariant, true
}

// SetEndUserDashboardTouchPointVariant sets field value
func (o *UpdateThemeRequest) SetEndUserDashboardTouchPointVariant(v string) {
	o.EndUserDashboardTouchPointVariant = v
}

// GetErrorPageTouchPointVariant returns the ErrorPageTouchPointVariant field value
func (o *UpdateThemeRequest) GetErrorPageTouchPointVariant() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorPageTouchPointVariant
}

// GetErrorPageTouchPointVariantOk returns a tuple with the ErrorPageTouchPointVariant field value
// and a boolean to check if the value has been set.
func (o *UpdateThemeRequest) GetErrorPageTouchPointVariantOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorPageTouchPointVariant, true
}

// SetErrorPageTouchPointVariant sets field value
func (o *UpdateThemeRequest) SetErrorPageTouchPointVariant(v string) {
	o.ErrorPageTouchPointVariant = v
}

// GetLoadingPageTouchPointVariant returns the LoadingPageTouchPointVariant field value if set, zero value otherwise.
func (o *UpdateThemeRequest) GetLoadingPageTouchPointVariant() string {
	if o == nil || IsNil(o.LoadingPageTouchPointVariant) {
		var ret string
		return ret
	}
	return *o.LoadingPageTouchPointVariant
}

// GetLoadingPageTouchPointVariantOk returns a tuple with the LoadingPageTouchPointVariant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateThemeRequest) GetLoadingPageTouchPointVariantOk() (*string, bool) {
	if o == nil || IsNil(o.LoadingPageTouchPointVariant) {
		return nil, false
	}
	return o.LoadingPageTouchPointVariant, true
}

// HasLoadingPageTouchPointVariant returns a boolean if a field has been set.
func (o *UpdateThemeRequest) HasLoadingPageTouchPointVariant() bool {
	if o != nil && !IsNil(o.LoadingPageTouchPointVariant) {
		return true
	}

	return false
}

// SetLoadingPageTouchPointVariant gets a reference to the given string and assigns it to the LoadingPageTouchPointVariant field.
func (o *UpdateThemeRequest) SetLoadingPageTouchPointVariant(v string) {
	o.LoadingPageTouchPointVariant = &v
}

// GetPrimaryColorContrastHex returns the PrimaryColorContrastHex field value if set, zero value otherwise.
func (o *UpdateThemeRequest) GetPrimaryColorContrastHex() string {
	if o == nil || IsNil(o.PrimaryColorContrastHex) {
		var ret string
		return ret
	}
	return *o.PrimaryColorContrastHex
}

// GetPrimaryColorContrastHexOk returns a tuple with the PrimaryColorContrastHex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateThemeRequest) GetPrimaryColorContrastHexOk() (*string, bool) {
	if o == nil || IsNil(o.PrimaryColorContrastHex) {
		return nil, false
	}
	return o.PrimaryColorContrastHex, true
}

// HasPrimaryColorContrastHex returns a boolean if a field has been set.
func (o *UpdateThemeRequest) HasPrimaryColorContrastHex() bool {
	if o != nil && !IsNil(o.PrimaryColorContrastHex) {
		return true
	}

	return false
}

// SetPrimaryColorContrastHex gets a reference to the given string and assigns it to the PrimaryColorContrastHex field.
func (o *UpdateThemeRequest) SetPrimaryColorContrastHex(v string) {
	o.PrimaryColorContrastHex = &v
}

// GetPrimaryColorHex returns the PrimaryColorHex field value
func (o *UpdateThemeRequest) GetPrimaryColorHex() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrimaryColorHex
}

// GetPrimaryColorHexOk returns a tuple with the PrimaryColorHex field value
// and a boolean to check if the value has been set.
func (o *UpdateThemeRequest) GetPrimaryColorHexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrimaryColorHex, true
}

// SetPrimaryColorHex sets field value
func (o *UpdateThemeRequest) SetPrimaryColorHex(v string) {
	o.PrimaryColorHex = v
}

// GetSecondaryColorContrastHex returns the SecondaryColorContrastHex field value if set, zero value otherwise.
func (o *UpdateThemeRequest) GetSecondaryColorContrastHex() string {
	if o == nil || IsNil(o.SecondaryColorContrastHex) {
		var ret string
		return ret
	}
	return *o.SecondaryColorContrastHex
}

// GetSecondaryColorContrastHexOk returns a tuple with the SecondaryColorContrastHex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateThemeRequest) GetSecondaryColorContrastHexOk() (*string, bool) {
	if o == nil || IsNil(o.SecondaryColorContrastHex) {
		return nil, false
	}
	return o.SecondaryColorContrastHex, true
}

// HasSecondaryColorContrastHex returns a boolean if a field has been set.
func (o *UpdateThemeRequest) HasSecondaryColorContrastHex() bool {
	if o != nil && !IsNil(o.SecondaryColorContrastHex) {
		return true
	}

	return false
}

// SetSecondaryColorContrastHex gets a reference to the given string and assigns it to the SecondaryColorContrastHex field.
func (o *UpdateThemeRequest) SetSecondaryColorContrastHex(v string) {
	o.SecondaryColorContrastHex = &v
}

// GetSecondaryColorHex returns the SecondaryColorHex field value
func (o *UpdateThemeRequest) GetSecondaryColorHex() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SecondaryColorHex
}

// GetSecondaryColorHexOk returns a tuple with the SecondaryColorHex field value
// and a boolean to check if the value has been set.
func (o *UpdateThemeRequest) GetSecondaryColorHexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecondaryColorHex, true
}

// SetSecondaryColorHex sets field value
func (o *UpdateThemeRequest) SetSecondaryColorHex(v string) {
	o.SecondaryColorHex = v
}

// GetSignInPageTouchPointVariant returns the SignInPageTouchPointVariant field value
func (o *UpdateThemeRequest) GetSignInPageTouchPointVariant() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SignInPageTouchPointVariant
}

// GetSignInPageTouchPointVariantOk returns a tuple with the SignInPageTouchPointVariant field value
// and a boolean to check if the value has been set.
func (o *UpdateThemeRequest) GetSignInPageTouchPointVariantOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SignInPageTouchPointVariant, true
}

// SetSignInPageTouchPointVariant sets field value
func (o *UpdateThemeRequest) SetSignInPageTouchPointVariant(v string) {
	o.SignInPageTouchPointVariant = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *UpdateThemeRequest) GetLinks() LinksSelf {
	if o == nil || IsNil(o.Links) {
		var ret LinksSelf
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateThemeRequest) GetLinksOk() (*LinksSelf, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *UpdateThemeRequest) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinksSelf and assigns it to the Links field.
func (o *UpdateThemeRequest) SetLinks(v LinksSelf) {
	o.Links = &v
}

func (o UpdateThemeRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateThemeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["emailTemplateTouchPointVariant"] = o.EmailTemplateTouchPointVariant
	toSerialize["endUserDashboardTouchPointVariant"] = o.EndUserDashboardTouchPointVariant
	toSerialize["errorPageTouchPointVariant"] = o.ErrorPageTouchPointVariant
	if !IsNil(o.LoadingPageTouchPointVariant) {
		toSerialize["loadingPageTouchPointVariant"] = o.LoadingPageTouchPointVariant
	}
	if !IsNil(o.PrimaryColorContrastHex) {
		toSerialize["primaryColorContrastHex"] = o.PrimaryColorContrastHex
	}
	toSerialize["primaryColorHex"] = o.PrimaryColorHex
	if !IsNil(o.SecondaryColorContrastHex) {
		toSerialize["secondaryColorContrastHex"] = o.SecondaryColorContrastHex
	}
	toSerialize["secondaryColorHex"] = o.SecondaryColorHex
	toSerialize["signInPageTouchPointVariant"] = o.SignInPageTouchPointVariant
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateThemeRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"emailTemplateTouchPointVariant",
		"endUserDashboardTouchPointVariant",
		"errorPageTouchPointVariant",
		"primaryColorHex",
		"secondaryColorHex",
		"signInPageTouchPointVariant",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varUpdateThemeRequest := _UpdateThemeRequest{}

	err = json.Unmarshal(data, &varUpdateThemeRequest)

	if err != nil {
		return err
	}

	*o = UpdateThemeRequest(varUpdateThemeRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "emailTemplateTouchPointVariant")
		delete(additionalProperties, "endUserDashboardTouchPointVariant")
		delete(additionalProperties, "errorPageTouchPointVariant")
		delete(additionalProperties, "loadingPageTouchPointVariant")
		delete(additionalProperties, "primaryColorContrastHex")
		delete(additionalProperties, "primaryColorHex")
		delete(additionalProperties, "secondaryColorContrastHex")
		delete(additionalProperties, "secondaryColorHex")
		delete(additionalProperties, "signInPageTouchPointVariant")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateThemeRequest struct {
	value *UpdateThemeRequest
	isSet bool
}

func (v NullableUpdateThemeRequest) Get() *UpdateThemeRequest {
	return v.value
}

func (v *NullableUpdateThemeRequest) Set(val *UpdateThemeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateThemeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateThemeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateThemeRequest(val *UpdateThemeRequest) *NullableUpdateThemeRequest {
	return &NullableUpdateThemeRequest{value: val, isSet: true}
}

func (v NullableUpdateThemeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateThemeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
