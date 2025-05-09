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

// ThemeResponse struct for ThemeResponse
type ThemeResponse struct {
	BackgroundImage *string `json:"backgroundImage,omitempty"`
	// Variant for email templates. You can publish a theme for email templates with different combinations of assets. Variants are preset combinations of those assets. 
	EmailTemplateTouchPointVariant *string `json:"emailTemplateTouchPointVariant,omitempty"`
	// Variant for the Okta End-User Dashboard. You can publish a theme for end-user dashboard with different combinations of assets. Variants are preset combinations of those assets. 
	EndUserDashboardTouchPointVariant *string `json:"endUserDashboardTouchPointVariant,omitempty"`
	// Variant for the error page. You can publish a theme for error page with different combinations of assets. Variants are preset combinations of those assets. 
	ErrorPageTouchPointVariant *string `json:"errorPageTouchPointVariant,omitempty"`
	Favicon *string `json:"favicon,omitempty"`
	Id *string `json:"id,omitempty"`
	// Variant for the Okta loading page. You can publish a theme for Okta loading page with different combinations of assets. Variants are preset combinations of those assets. 
	LoadingPageTouchPointVariant *string `json:"loadingPageTouchPointVariant,omitempty"`
	Logo *string `json:"logo,omitempty"`
	// Primary color contrast hex code
	PrimaryColorContrastHex *string `json:"primaryColorContrastHex,omitempty"`
	// Primary color hex code
	PrimaryColorHex *string `json:"primaryColorHex,omitempty"`
	// Secondary color contrast hex code
	SecondaryColorContrastHex *string `json:"secondaryColorContrastHex,omitempty"`
	// Secondary color hex code
	SecondaryColorHex *string `json:"secondaryColorHex,omitempty"`
	// Variant for the Okta sign-in page. You can publish a theme for sign-in page with different combinations of assets. Variants are preset combinations of those assets. > **Note:**  For a non-`OKTA_DEFAULT` variant, `primaryColorHex` is used for button background color and `primaryColorContrastHex` is used to optimize the opacity for button text. 
	SignInPageTouchPointVariant *string `json:"signInPageTouchPointVariant,omitempty"`
	Links *LinksSelf `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ThemeResponse ThemeResponse

// NewThemeResponse instantiates a new ThemeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThemeResponse() *ThemeResponse {
	this := ThemeResponse{}
	var emailTemplateTouchPointVariant string = "OKTA_DEFAULT"
	this.EmailTemplateTouchPointVariant = &emailTemplateTouchPointVariant
	var endUserDashboardTouchPointVariant string = "OKTA_DEFAULT"
	this.EndUserDashboardTouchPointVariant = &endUserDashboardTouchPointVariant
	var errorPageTouchPointVariant string = "OKTA_DEFAULT"
	this.ErrorPageTouchPointVariant = &errorPageTouchPointVariant
	var loadingPageTouchPointVariant string = "OKTA_DEFAULT"
	this.LoadingPageTouchPointVariant = &loadingPageTouchPointVariant
	return &this
}

// NewThemeResponseWithDefaults instantiates a new ThemeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThemeResponseWithDefaults() *ThemeResponse {
	this := ThemeResponse{}
	var emailTemplateTouchPointVariant string = "OKTA_DEFAULT"
	this.EmailTemplateTouchPointVariant = &emailTemplateTouchPointVariant
	var endUserDashboardTouchPointVariant string = "OKTA_DEFAULT"
	this.EndUserDashboardTouchPointVariant = &endUserDashboardTouchPointVariant
	var errorPageTouchPointVariant string = "OKTA_DEFAULT"
	this.ErrorPageTouchPointVariant = &errorPageTouchPointVariant
	var loadingPageTouchPointVariant string = "OKTA_DEFAULT"
	this.LoadingPageTouchPointVariant = &loadingPageTouchPointVariant
	return &this
}

// GetBackgroundImage returns the BackgroundImage field value if set, zero value otherwise.
func (o *ThemeResponse) GetBackgroundImage() string {
	if o == nil || o.BackgroundImage == nil {
		var ret string
		return ret
	}
	return *o.BackgroundImage
}

// GetBackgroundImageOk returns a tuple with the BackgroundImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThemeResponse) GetBackgroundImageOk() (*string, bool) {
	if o == nil || o.BackgroundImage == nil {
		return nil, false
	}
	return o.BackgroundImage, true
}

// HasBackgroundImage returns a boolean if a field has been set.
func (o *ThemeResponse) HasBackgroundImage() bool {
	if o != nil && o.BackgroundImage != nil {
		return true
	}

	return false
}

// SetBackgroundImage gets a reference to the given string and assigns it to the BackgroundImage field.
func (o *ThemeResponse) SetBackgroundImage(v string) {
	o.BackgroundImage = &v
}

// GetEmailTemplateTouchPointVariant returns the EmailTemplateTouchPointVariant field value if set, zero value otherwise.
func (o *ThemeResponse) GetEmailTemplateTouchPointVariant() string {
	if o == nil || o.EmailTemplateTouchPointVariant == nil {
		var ret string
		return ret
	}
	return *o.EmailTemplateTouchPointVariant
}

// GetEmailTemplateTouchPointVariantOk returns a tuple with the EmailTemplateTouchPointVariant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThemeResponse) GetEmailTemplateTouchPointVariantOk() (*string, bool) {
	if o == nil || o.EmailTemplateTouchPointVariant == nil {
		return nil, false
	}
	return o.EmailTemplateTouchPointVariant, true
}

// HasEmailTemplateTouchPointVariant returns a boolean if a field has been set.
func (o *ThemeResponse) HasEmailTemplateTouchPointVariant() bool {
	if o != nil && o.EmailTemplateTouchPointVariant != nil {
		return true
	}

	return false
}

// SetEmailTemplateTouchPointVariant gets a reference to the given string and assigns it to the EmailTemplateTouchPointVariant field.
func (o *ThemeResponse) SetEmailTemplateTouchPointVariant(v string) {
	o.EmailTemplateTouchPointVariant = &v
}

// GetEndUserDashboardTouchPointVariant returns the EndUserDashboardTouchPointVariant field value if set, zero value otherwise.
func (o *ThemeResponse) GetEndUserDashboardTouchPointVariant() string {
	if o == nil || o.EndUserDashboardTouchPointVariant == nil {
		var ret string
		return ret
	}
	return *o.EndUserDashboardTouchPointVariant
}

// GetEndUserDashboardTouchPointVariantOk returns a tuple with the EndUserDashboardTouchPointVariant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThemeResponse) GetEndUserDashboardTouchPointVariantOk() (*string, bool) {
	if o == nil || o.EndUserDashboardTouchPointVariant == nil {
		return nil, false
	}
	return o.EndUserDashboardTouchPointVariant, true
}

// HasEndUserDashboardTouchPointVariant returns a boolean if a field has been set.
func (o *ThemeResponse) HasEndUserDashboardTouchPointVariant() bool {
	if o != nil && o.EndUserDashboardTouchPointVariant != nil {
		return true
	}

	return false
}

// SetEndUserDashboardTouchPointVariant gets a reference to the given string and assigns it to the EndUserDashboardTouchPointVariant field.
func (o *ThemeResponse) SetEndUserDashboardTouchPointVariant(v string) {
	o.EndUserDashboardTouchPointVariant = &v
}

// GetErrorPageTouchPointVariant returns the ErrorPageTouchPointVariant field value if set, zero value otherwise.
func (o *ThemeResponse) GetErrorPageTouchPointVariant() string {
	if o == nil || o.ErrorPageTouchPointVariant == nil {
		var ret string
		return ret
	}
	return *o.ErrorPageTouchPointVariant
}

// GetErrorPageTouchPointVariantOk returns a tuple with the ErrorPageTouchPointVariant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThemeResponse) GetErrorPageTouchPointVariantOk() (*string, bool) {
	if o == nil || o.ErrorPageTouchPointVariant == nil {
		return nil, false
	}
	return o.ErrorPageTouchPointVariant, true
}

// HasErrorPageTouchPointVariant returns a boolean if a field has been set.
func (o *ThemeResponse) HasErrorPageTouchPointVariant() bool {
	if o != nil && o.ErrorPageTouchPointVariant != nil {
		return true
	}

	return false
}

// SetErrorPageTouchPointVariant gets a reference to the given string and assigns it to the ErrorPageTouchPointVariant field.
func (o *ThemeResponse) SetErrorPageTouchPointVariant(v string) {
	o.ErrorPageTouchPointVariant = &v
}

// GetFavicon returns the Favicon field value if set, zero value otherwise.
func (o *ThemeResponse) GetFavicon() string {
	if o == nil || o.Favicon == nil {
		var ret string
		return ret
	}
	return *o.Favicon
}

// GetFaviconOk returns a tuple with the Favicon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThemeResponse) GetFaviconOk() (*string, bool) {
	if o == nil || o.Favicon == nil {
		return nil, false
	}
	return o.Favicon, true
}

// HasFavicon returns a boolean if a field has been set.
func (o *ThemeResponse) HasFavicon() bool {
	if o != nil && o.Favicon != nil {
		return true
	}

	return false
}

// SetFavicon gets a reference to the given string and assigns it to the Favicon field.
func (o *ThemeResponse) SetFavicon(v string) {
	o.Favicon = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ThemeResponse) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThemeResponse) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ThemeResponse) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ThemeResponse) SetId(v string) {
	o.Id = &v
}

// GetLoadingPageTouchPointVariant returns the LoadingPageTouchPointVariant field value if set, zero value otherwise.
func (o *ThemeResponse) GetLoadingPageTouchPointVariant() string {
	if o == nil || o.LoadingPageTouchPointVariant == nil {
		var ret string
		return ret
	}
	return *o.LoadingPageTouchPointVariant
}

// GetLoadingPageTouchPointVariantOk returns a tuple with the LoadingPageTouchPointVariant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThemeResponse) GetLoadingPageTouchPointVariantOk() (*string, bool) {
	if o == nil || o.LoadingPageTouchPointVariant == nil {
		return nil, false
	}
	return o.LoadingPageTouchPointVariant, true
}

// HasLoadingPageTouchPointVariant returns a boolean if a field has been set.
func (o *ThemeResponse) HasLoadingPageTouchPointVariant() bool {
	if o != nil && o.LoadingPageTouchPointVariant != nil {
		return true
	}

	return false
}

// SetLoadingPageTouchPointVariant gets a reference to the given string and assigns it to the LoadingPageTouchPointVariant field.
func (o *ThemeResponse) SetLoadingPageTouchPointVariant(v string) {
	o.LoadingPageTouchPointVariant = &v
}

// GetLogo returns the Logo field value if set, zero value otherwise.
func (o *ThemeResponse) GetLogo() string {
	if o == nil || o.Logo == nil {
		var ret string
		return ret
	}
	return *o.Logo
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThemeResponse) GetLogoOk() (*string, bool) {
	if o == nil || o.Logo == nil {
		return nil, false
	}
	return o.Logo, true
}

// HasLogo returns a boolean if a field has been set.
func (o *ThemeResponse) HasLogo() bool {
	if o != nil && o.Logo != nil {
		return true
	}

	return false
}

// SetLogo gets a reference to the given string and assigns it to the Logo field.
func (o *ThemeResponse) SetLogo(v string) {
	o.Logo = &v
}

// GetPrimaryColorContrastHex returns the PrimaryColorContrastHex field value if set, zero value otherwise.
func (o *ThemeResponse) GetPrimaryColorContrastHex() string {
	if o == nil || o.PrimaryColorContrastHex == nil {
		var ret string
		return ret
	}
	return *o.PrimaryColorContrastHex
}

// GetPrimaryColorContrastHexOk returns a tuple with the PrimaryColorContrastHex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThemeResponse) GetPrimaryColorContrastHexOk() (*string, bool) {
	if o == nil || o.PrimaryColorContrastHex == nil {
		return nil, false
	}
	return o.PrimaryColorContrastHex, true
}

// HasPrimaryColorContrastHex returns a boolean if a field has been set.
func (o *ThemeResponse) HasPrimaryColorContrastHex() bool {
	if o != nil && o.PrimaryColorContrastHex != nil {
		return true
	}

	return false
}

// SetPrimaryColorContrastHex gets a reference to the given string and assigns it to the PrimaryColorContrastHex field.
func (o *ThemeResponse) SetPrimaryColorContrastHex(v string) {
	o.PrimaryColorContrastHex = &v
}

// GetPrimaryColorHex returns the PrimaryColorHex field value if set, zero value otherwise.
func (o *ThemeResponse) GetPrimaryColorHex() string {
	if o == nil || o.PrimaryColorHex == nil {
		var ret string
		return ret
	}
	return *o.PrimaryColorHex
}

// GetPrimaryColorHexOk returns a tuple with the PrimaryColorHex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThemeResponse) GetPrimaryColorHexOk() (*string, bool) {
	if o == nil || o.PrimaryColorHex == nil {
		return nil, false
	}
	return o.PrimaryColorHex, true
}

// HasPrimaryColorHex returns a boolean if a field has been set.
func (o *ThemeResponse) HasPrimaryColorHex() bool {
	if o != nil && o.PrimaryColorHex != nil {
		return true
	}

	return false
}

// SetPrimaryColorHex gets a reference to the given string and assigns it to the PrimaryColorHex field.
func (o *ThemeResponse) SetPrimaryColorHex(v string) {
	o.PrimaryColorHex = &v
}

// GetSecondaryColorContrastHex returns the SecondaryColorContrastHex field value if set, zero value otherwise.
func (o *ThemeResponse) GetSecondaryColorContrastHex() string {
	if o == nil || o.SecondaryColorContrastHex == nil {
		var ret string
		return ret
	}
	return *o.SecondaryColorContrastHex
}

// GetSecondaryColorContrastHexOk returns a tuple with the SecondaryColorContrastHex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThemeResponse) GetSecondaryColorContrastHexOk() (*string, bool) {
	if o == nil || o.SecondaryColorContrastHex == nil {
		return nil, false
	}
	return o.SecondaryColorContrastHex, true
}

// HasSecondaryColorContrastHex returns a boolean if a field has been set.
func (o *ThemeResponse) HasSecondaryColorContrastHex() bool {
	if o != nil && o.SecondaryColorContrastHex != nil {
		return true
	}

	return false
}

// SetSecondaryColorContrastHex gets a reference to the given string and assigns it to the SecondaryColorContrastHex field.
func (o *ThemeResponse) SetSecondaryColorContrastHex(v string) {
	o.SecondaryColorContrastHex = &v
}

// GetSecondaryColorHex returns the SecondaryColorHex field value if set, zero value otherwise.
func (o *ThemeResponse) GetSecondaryColorHex() string {
	if o == nil || o.SecondaryColorHex == nil {
		var ret string
		return ret
	}
	return *o.SecondaryColorHex
}

// GetSecondaryColorHexOk returns a tuple with the SecondaryColorHex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThemeResponse) GetSecondaryColorHexOk() (*string, bool) {
	if o == nil || o.SecondaryColorHex == nil {
		return nil, false
	}
	return o.SecondaryColorHex, true
}

// HasSecondaryColorHex returns a boolean if a field has been set.
func (o *ThemeResponse) HasSecondaryColorHex() bool {
	if o != nil && o.SecondaryColorHex != nil {
		return true
	}

	return false
}

// SetSecondaryColorHex gets a reference to the given string and assigns it to the SecondaryColorHex field.
func (o *ThemeResponse) SetSecondaryColorHex(v string) {
	o.SecondaryColorHex = &v
}

// GetSignInPageTouchPointVariant returns the SignInPageTouchPointVariant field value if set, zero value otherwise.
func (o *ThemeResponse) GetSignInPageTouchPointVariant() string {
	if o == nil || o.SignInPageTouchPointVariant == nil {
		var ret string
		return ret
	}
	return *o.SignInPageTouchPointVariant
}

// GetSignInPageTouchPointVariantOk returns a tuple with the SignInPageTouchPointVariant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThemeResponse) GetSignInPageTouchPointVariantOk() (*string, bool) {
	if o == nil || o.SignInPageTouchPointVariant == nil {
		return nil, false
	}
	return o.SignInPageTouchPointVariant, true
}

// HasSignInPageTouchPointVariant returns a boolean if a field has been set.
func (o *ThemeResponse) HasSignInPageTouchPointVariant() bool {
	if o != nil && o.SignInPageTouchPointVariant != nil {
		return true
	}

	return false
}

// SetSignInPageTouchPointVariant gets a reference to the given string and assigns it to the SignInPageTouchPointVariant field.
func (o *ThemeResponse) SetSignInPageTouchPointVariant(v string) {
	o.SignInPageTouchPointVariant = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ThemeResponse) GetLinks() LinksSelf {
	if o == nil || o.Links == nil {
		var ret LinksSelf
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThemeResponse) GetLinksOk() (*LinksSelf, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ThemeResponse) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinksSelf and assigns it to the Links field.
func (o *ThemeResponse) SetLinks(v LinksSelf) {
	o.Links = &v
}

func (o ThemeResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BackgroundImage != nil {
		toSerialize["backgroundImage"] = o.BackgroundImage
	}
	if o.EmailTemplateTouchPointVariant != nil {
		toSerialize["emailTemplateTouchPointVariant"] = o.EmailTemplateTouchPointVariant
	}
	if o.EndUserDashboardTouchPointVariant != nil {
		toSerialize["endUserDashboardTouchPointVariant"] = o.EndUserDashboardTouchPointVariant
	}
	if o.ErrorPageTouchPointVariant != nil {
		toSerialize["errorPageTouchPointVariant"] = o.ErrorPageTouchPointVariant
	}
	if o.Favicon != nil {
		toSerialize["favicon"] = o.Favicon
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LoadingPageTouchPointVariant != nil {
		toSerialize["loadingPageTouchPointVariant"] = o.LoadingPageTouchPointVariant
	}
	if o.Logo != nil {
		toSerialize["logo"] = o.Logo
	}
	if o.PrimaryColorContrastHex != nil {
		toSerialize["primaryColorContrastHex"] = o.PrimaryColorContrastHex
	}
	if o.PrimaryColorHex != nil {
		toSerialize["primaryColorHex"] = o.PrimaryColorHex
	}
	if o.SecondaryColorContrastHex != nil {
		toSerialize["secondaryColorContrastHex"] = o.SecondaryColorContrastHex
	}
	if o.SecondaryColorHex != nil {
		toSerialize["secondaryColorHex"] = o.SecondaryColorHex
	}
	if o.SignInPageTouchPointVariant != nil {
		toSerialize["signInPageTouchPointVariant"] = o.SignInPageTouchPointVariant
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ThemeResponse) UnmarshalJSON(bytes []byte) (err error) {
	varThemeResponse := _ThemeResponse{}

	err = json.Unmarshal(bytes, &varThemeResponse)
	if err == nil {
		*o = ThemeResponse(varThemeResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "backgroundImage")
		delete(additionalProperties, "emailTemplateTouchPointVariant")
		delete(additionalProperties, "endUserDashboardTouchPointVariant")
		delete(additionalProperties, "errorPageTouchPointVariant")
		delete(additionalProperties, "favicon")
		delete(additionalProperties, "id")
		delete(additionalProperties, "loadingPageTouchPointVariant")
		delete(additionalProperties, "logo")
		delete(additionalProperties, "primaryColorContrastHex")
		delete(additionalProperties, "primaryColorHex")
		delete(additionalProperties, "secondaryColorContrastHex")
		delete(additionalProperties, "secondaryColorHex")
		delete(additionalProperties, "signInPageTouchPointVariant")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableThemeResponse struct {
	value *ThemeResponse
	isSet bool
}

func (v NullableThemeResponse) Get() *ThemeResponse {
	return v.value
}

func (v *NullableThemeResponse) Set(val *ThemeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableThemeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableThemeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThemeResponse(val *ThemeResponse) *NullableThemeResponse {
	return &NullableThemeResponse{value: val, isSet: true}
}

func (v NullableThemeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThemeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

