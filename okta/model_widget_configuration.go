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

// checks if the WidgetConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WidgetConfiguration{}

// WidgetConfiguration <x-lifecycle-container><x-lifecycle class=\"oie\"></x-lifecycle></x-lifecycle-container>The Sign-In Widget configuration object containing branding, features, theme tokens, and other customization settings.
type WidgetConfiguration struct {
	BrandColors *WidgetConfigurationBrandColors `json:"brandColors,omitempty"`
	// Brand name displayed on the widget
	BrandName *string                       `json:"brandName,omitempty"`
	Features  *WidgetConfigurationFeatures  `json:"features,omitempty"`
	HelpLinks *WidgetConfigurationHelpLinks `json:"helpLinks,omitempty"`
	// Support phone number displayed on the widget
	HelpSupportNumber *string `json:"helpSupportNumber,omitempty"`
	// Internationalization overrides. Keys are locale codes (e.g. \"en\", \"ja\", \"fr\"), values are objects mapping i18n keys to string overrides. Override values can't contain `<` or `>` characters.
	I18n *map[string]map[string]string `json:"i18n,omitempty"`
	// Language code for the widget
	Language *string `json:"language,omitempty"`
	// URL for the logo displayed on the Sign-In Widget
	Logo *string `json:"logo,omitempty" validate:"regexp=^https?:\\/\\/"`
	// Alt text for the logo
	LogoText *string `json:"logoText,omitempty"`
	// Custom sign-out URL
	SignOutLink *string                   `json:"signOutLink,omitempty" validate:"regexp=^https?:\\/\\/"`
	Theme       *WidgetConfigurationTheme `json:"theme,omitempty"`
}

// NewWidgetConfiguration instantiates a new WidgetConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWidgetConfiguration() *WidgetConfiguration {
	this := WidgetConfiguration{}
	return &this
}

// NewWidgetConfigurationWithDefaults instantiates a new WidgetConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWidgetConfigurationWithDefaults() *WidgetConfiguration {
	this := WidgetConfiguration{}
	return &this
}

// GetBrandColors returns the BrandColors field value if set, zero value otherwise.
func (o *WidgetConfiguration) GetBrandColors() WidgetConfigurationBrandColors {
	if o == nil || IsNil(o.BrandColors) {
		var ret WidgetConfigurationBrandColors
		return ret
	}
	return *o.BrandColors
}

// GetBrandColorsOk returns a tuple with the BrandColors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfiguration) GetBrandColorsOk() (*WidgetConfigurationBrandColors, bool) {
	if o == nil || IsNil(o.BrandColors) {
		return nil, false
	}
	return o.BrandColors, true
}

// HasBrandColors returns a boolean if a field has been set.
func (o *WidgetConfiguration) HasBrandColors() bool {
	if o != nil && !IsNil(o.BrandColors) {
		return true
	}

	return false
}

// SetBrandColors gets a reference to the given WidgetConfigurationBrandColors and assigns it to the BrandColors field.
func (o *WidgetConfiguration) SetBrandColors(v WidgetConfigurationBrandColors) {
	o.BrandColors = &v
}

// GetBrandName returns the BrandName field value if set, zero value otherwise.
func (o *WidgetConfiguration) GetBrandName() string {
	if o == nil || IsNil(o.BrandName) {
		var ret string
		return ret
	}
	return *o.BrandName
}

// GetBrandNameOk returns a tuple with the BrandName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfiguration) GetBrandNameOk() (*string, bool) {
	if o == nil || IsNil(o.BrandName) {
		return nil, false
	}
	return o.BrandName, true
}

// HasBrandName returns a boolean if a field has been set.
func (o *WidgetConfiguration) HasBrandName() bool {
	if o != nil && !IsNil(o.BrandName) {
		return true
	}

	return false
}

// SetBrandName gets a reference to the given string and assigns it to the BrandName field.
func (o *WidgetConfiguration) SetBrandName(v string) {
	o.BrandName = &v
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *WidgetConfiguration) GetFeatures() WidgetConfigurationFeatures {
	if o == nil || IsNil(o.Features) {
		var ret WidgetConfigurationFeatures
		return ret
	}
	return *o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfiguration) GetFeaturesOk() (*WidgetConfigurationFeatures, bool) {
	if o == nil || IsNil(o.Features) {
		return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *WidgetConfiguration) HasFeatures() bool {
	if o != nil && !IsNil(o.Features) {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given WidgetConfigurationFeatures and assigns it to the Features field.
func (o *WidgetConfiguration) SetFeatures(v WidgetConfigurationFeatures) {
	o.Features = &v
}

// GetHelpLinks returns the HelpLinks field value if set, zero value otherwise.
func (o *WidgetConfiguration) GetHelpLinks() WidgetConfigurationHelpLinks {
	if o == nil || IsNil(o.HelpLinks) {
		var ret WidgetConfigurationHelpLinks
		return ret
	}
	return *o.HelpLinks
}

// GetHelpLinksOk returns a tuple with the HelpLinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfiguration) GetHelpLinksOk() (*WidgetConfigurationHelpLinks, bool) {
	if o == nil || IsNil(o.HelpLinks) {
		return nil, false
	}
	return o.HelpLinks, true
}

// HasHelpLinks returns a boolean if a field has been set.
func (o *WidgetConfiguration) HasHelpLinks() bool {
	if o != nil && !IsNil(o.HelpLinks) {
		return true
	}

	return false
}

// SetHelpLinks gets a reference to the given WidgetConfigurationHelpLinks and assigns it to the HelpLinks field.
func (o *WidgetConfiguration) SetHelpLinks(v WidgetConfigurationHelpLinks) {
	o.HelpLinks = &v
}

// GetHelpSupportNumber returns the HelpSupportNumber field value if set, zero value otherwise.
func (o *WidgetConfiguration) GetHelpSupportNumber() string {
	if o == nil || IsNil(o.HelpSupportNumber) {
		var ret string
		return ret
	}
	return *o.HelpSupportNumber
}

// GetHelpSupportNumberOk returns a tuple with the HelpSupportNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfiguration) GetHelpSupportNumberOk() (*string, bool) {
	if o == nil || IsNil(o.HelpSupportNumber) {
		return nil, false
	}
	return o.HelpSupportNumber, true
}

// HasHelpSupportNumber returns a boolean if a field has been set.
func (o *WidgetConfiguration) HasHelpSupportNumber() bool {
	if o != nil && !IsNil(o.HelpSupportNumber) {
		return true
	}

	return false
}

// SetHelpSupportNumber gets a reference to the given string and assigns it to the HelpSupportNumber field.
func (o *WidgetConfiguration) SetHelpSupportNumber(v string) {
	o.HelpSupportNumber = &v
}

// GetI18n returns the I18n field value if set, zero value otherwise.
func (o *WidgetConfiguration) GetI18n() map[string]map[string]string {
	if o == nil || IsNil(o.I18n) {
		var ret map[string]map[string]string
		return ret
	}
	return *o.I18n
}

// GetI18nOk returns a tuple with the I18n field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfiguration) GetI18nOk() (*map[string]map[string]string, bool) {
	if o == nil || IsNil(o.I18n) {
		return nil, false
	}
	return o.I18n, true
}

// HasI18n returns a boolean if a field has been set.
func (o *WidgetConfiguration) HasI18n() bool {
	if o != nil && !IsNil(o.I18n) {
		return true
	}

	return false
}

// SetI18n gets a reference to the given map[string]map[string]string and assigns it to the I18n field.
func (o *WidgetConfiguration) SetI18n(v map[string]map[string]string) {
	o.I18n = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *WidgetConfiguration) GetLanguage() string {
	if o == nil || IsNil(o.Language) {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfiguration) GetLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *WidgetConfiguration) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *WidgetConfiguration) SetLanguage(v string) {
	o.Language = &v
}

// GetLogo returns the Logo field value if set, zero value otherwise.
func (o *WidgetConfiguration) GetLogo() string {
	if o == nil || IsNil(o.Logo) {
		var ret string
		return ret
	}
	return *o.Logo
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfiguration) GetLogoOk() (*string, bool) {
	if o == nil || IsNil(o.Logo) {
		return nil, false
	}
	return o.Logo, true
}

// HasLogo returns a boolean if a field has been set.
func (o *WidgetConfiguration) HasLogo() bool {
	if o != nil && !IsNil(o.Logo) {
		return true
	}

	return false
}

// SetLogo gets a reference to the given string and assigns it to the Logo field.
func (o *WidgetConfiguration) SetLogo(v string) {
	o.Logo = &v
}

// GetLogoText returns the LogoText field value if set, zero value otherwise.
func (o *WidgetConfiguration) GetLogoText() string {
	if o == nil || IsNil(o.LogoText) {
		var ret string
		return ret
	}
	return *o.LogoText
}

// GetLogoTextOk returns a tuple with the LogoText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfiguration) GetLogoTextOk() (*string, bool) {
	if o == nil || IsNil(o.LogoText) {
		return nil, false
	}
	return o.LogoText, true
}

// HasLogoText returns a boolean if a field has been set.
func (o *WidgetConfiguration) HasLogoText() bool {
	if o != nil && !IsNil(o.LogoText) {
		return true
	}

	return false
}

// SetLogoText gets a reference to the given string and assigns it to the LogoText field.
func (o *WidgetConfiguration) SetLogoText(v string) {
	o.LogoText = &v
}

// GetSignOutLink returns the SignOutLink field value if set, zero value otherwise.
func (o *WidgetConfiguration) GetSignOutLink() string {
	if o == nil || IsNil(o.SignOutLink) {
		var ret string
		return ret
	}
	return *o.SignOutLink
}

// GetSignOutLinkOk returns a tuple with the SignOutLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfiguration) GetSignOutLinkOk() (*string, bool) {
	if o == nil || IsNil(o.SignOutLink) {
		return nil, false
	}
	return o.SignOutLink, true
}

// HasSignOutLink returns a boolean if a field has been set.
func (o *WidgetConfiguration) HasSignOutLink() bool {
	if o != nil && !IsNil(o.SignOutLink) {
		return true
	}

	return false
}

// SetSignOutLink gets a reference to the given string and assigns it to the SignOutLink field.
func (o *WidgetConfiguration) SetSignOutLink(v string) {
	o.SignOutLink = &v
}

// GetTheme returns the Theme field value if set, zero value otherwise.
func (o *WidgetConfiguration) GetTheme() WidgetConfigurationTheme {
	if o == nil || IsNil(o.Theme) {
		var ret WidgetConfigurationTheme
		return ret
	}
	return *o.Theme
}

// GetThemeOk returns a tuple with the Theme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfiguration) GetThemeOk() (*WidgetConfigurationTheme, bool) {
	if o == nil || IsNil(o.Theme) {
		return nil, false
	}
	return o.Theme, true
}

// HasTheme returns a boolean if a field has been set.
func (o *WidgetConfiguration) HasTheme() bool {
	if o != nil && !IsNil(o.Theme) {
		return true
	}

	return false
}

// SetTheme gets a reference to the given WidgetConfigurationTheme and assigns it to the Theme field.
func (o *WidgetConfiguration) SetTheme(v WidgetConfigurationTheme) {
	o.Theme = &v
}

func (o WidgetConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WidgetConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BrandColors) {
		toSerialize["brandColors"] = o.BrandColors
	}
	if !IsNil(o.BrandName) {
		toSerialize["brandName"] = o.BrandName
	}
	if !IsNil(o.Features) {
		toSerialize["features"] = o.Features
	}
	if !IsNil(o.HelpLinks) {
		toSerialize["helpLinks"] = o.HelpLinks
	}
	if !IsNil(o.HelpSupportNumber) {
		toSerialize["helpSupportNumber"] = o.HelpSupportNumber
	}
	if !IsNil(o.I18n) {
		toSerialize["i18n"] = o.I18n
	}
	if !IsNil(o.Language) {
		toSerialize["language"] = o.Language
	}
	if !IsNil(o.Logo) {
		toSerialize["logo"] = o.Logo
	}
	if !IsNil(o.LogoText) {
		toSerialize["logoText"] = o.LogoText
	}
	if !IsNil(o.SignOutLink) {
		toSerialize["signOutLink"] = o.SignOutLink
	}
	if !IsNil(o.Theme) {
		toSerialize["theme"] = o.Theme
	}
	return toSerialize, nil
}

type NullableWidgetConfiguration struct {
	value *WidgetConfiguration
	isSet bool
}

func (v NullableWidgetConfiguration) Get() *WidgetConfiguration {
	return v.value
}

func (v *NullableWidgetConfiguration) Set(val *WidgetConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableWidgetConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableWidgetConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWidgetConfiguration(val *WidgetConfiguration) *NullableWidgetConfiguration {
	return &NullableWidgetConfiguration{value: val, isSet: true}
}

func (v NullableWidgetConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWidgetConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
