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

API version: 2025.08.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
	"fmt"
	"time"
)

// checks if the Application type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Application{}

// Application struct for Application
type Application struct {
	Accessibility *ApplicationAccessibility `json:"accessibility,omitempty"`
	// Timestamp when the application object was created
	Created *time.Time `json:"created,omitempty"`
	// Enabled app features > **Note:** See [Application Features](/openapi/okta-management/management/tag/ApplicationFeatures/) for app provisioning features.
	Features []string `json:"features,omitempty"`
	// Unique ID for the app instance
	Id *string `json:"id,omitempty"`
	// User-defined display name for app
	Label string `json:"label"`
	// Timestamp when the application object was last updated
	LastUpdated *time.Time            `json:"lastUpdated,omitempty"`
	Licensing   *ApplicationLicensing `json:"licensing,omitempty"`
	// The Okta resource name (ORN) for the current app instance
	Orn *string `json:"orn,omitempty"`
	// Contains any valid JSON schema for specifying properties that can be referenced from a request (only available to OAuth 2.0 client apps). For example, add an app manager contact email address or define an allowlist of groups that you can then reference using the Okta Expression Language `getFilteredGroups` function.  > **Notes:** > * `profile` isn't encrypted, so don't store sensitive data in it. > * `profile` doesn't limit the level of nesting in the JSON schema you created, but there is a practical size limit. Okta recommends a JSON schema size of 1 MB or less for best performance.
	Profile map[string]interface{} `json:"profile,omitempty"`
	// Authentication mode for the app  | signOnMode | Description | | ---------- | ----------- | | AUTO_LOGIN | Secure Web Authentication (SWA) | | BASIC_AUTH | HTTP Basic Authentication with Okta Browser Plugin | | BOOKMARK | Just a bookmark (no-authentication) | | BROWSER_PLUGIN | Secure Web Authentication (SWA) with Okta Browser Plugin | | OPENID_CONNECT | Federated Authentication with OpenID Connect (OIDC) | | SAML_1_1 | Federated Authentication with SAML 1.1 WebSSO (not supported for custom apps) | | SAML_2_0 | Federated Authentication with SAML 2.0 WebSSO | | SECURE_PASSWORD_STORE | Secure Web Authentication (SWA) with POST (plugin not required) | | WS_FEDERATION | Federated Authentication with WS-Federation Passive Requestor Profile |  Select the `signOnMode` for your custom app:
	SignOnMode string `json:"signOnMode"`
	// App instance status
	Status               *string                     `json:"status,omitempty"`
	UniversalLogout      *ApplicationUniversalLogout `json:"universalLogout,omitempty"`
	Visibility           *ApplicationVisibility      `json:"visibility,omitempty"`
	Embedded             *ApplicationEmbedded        `json:"_embedded,omitempty"`
	Links                *ApplicationLinks           `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Application Application

// NewApplication instantiates a new Application object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplication(label string, signOnMode string) *Application {
	this := Application{}
	this.Label = label
	this.SignOnMode = signOnMode
	return &this
}

// NewApplicationWithDefaults instantiates a new Application object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationWithDefaults() *Application {
	this := Application{}
	return &this
}

// GetAccessibility returns the Accessibility field value if set, zero value otherwise.
func (o *Application) GetAccessibility() ApplicationAccessibility {
	if o == nil || IsNil(o.Accessibility) {
		var ret ApplicationAccessibility
		return ret
	}
	return *o.Accessibility
}

// GetAccessibilityOk returns a tuple with the Accessibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetAccessibilityOk() (*ApplicationAccessibility, bool) {
	if o == nil || IsNil(o.Accessibility) {
		return nil, false
	}
	return o.Accessibility, true
}

// HasAccessibility returns a boolean if a field has been set.
func (o *Application) HasAccessibility() bool {
	if o != nil && !IsNil(o.Accessibility) {
		return true
	}

	return false
}

// SetAccessibility gets a reference to the given ApplicationAccessibility and assigns it to the Accessibility field.
func (o *Application) SetAccessibility(v ApplicationAccessibility) {
	o.Accessibility = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *Application) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *Application) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *Application) SetCreated(v time.Time) {
	o.Created = &v
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *Application) GetFeatures() []string {
	if o == nil || IsNil(o.Features) {
		var ret []string
		return ret
	}
	return o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetFeaturesOk() ([]string, bool) {
	if o == nil || IsNil(o.Features) {
		return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *Application) HasFeatures() bool {
	if o != nil && !IsNil(o.Features) {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given []string and assigns it to the Features field.
func (o *Application) SetFeatures(v []string) {
	o.Features = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Application) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Application) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Application) SetId(v string) {
	o.Id = &v
}

// GetLabel returns the Label field value
func (o *Application) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *Application) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *Application) SetLabel(v string) {
	o.Label = v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *Application) GetLastUpdated() time.Time {
	if o == nil || IsNil(o.LastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *Application) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *Application) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetLicensing returns the Licensing field value if set, zero value otherwise.
func (o *Application) GetLicensing() ApplicationLicensing {
	if o == nil || IsNil(o.Licensing) {
		var ret ApplicationLicensing
		return ret
	}
	return *o.Licensing
}

// GetLicensingOk returns a tuple with the Licensing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetLicensingOk() (*ApplicationLicensing, bool) {
	if o == nil || IsNil(o.Licensing) {
		return nil, false
	}
	return o.Licensing, true
}

// HasLicensing returns a boolean if a field has been set.
func (o *Application) HasLicensing() bool {
	if o != nil && !IsNil(o.Licensing) {
		return true
	}

	return false
}

// SetLicensing gets a reference to the given ApplicationLicensing and assigns it to the Licensing field.
func (o *Application) SetLicensing(v ApplicationLicensing) {
	o.Licensing = &v
}

// GetOrn returns the Orn field value if set, zero value otherwise.
func (o *Application) GetOrn() string {
	if o == nil || IsNil(o.Orn) {
		var ret string
		return ret
	}
	return *o.Orn
}

// GetOrnOk returns a tuple with the Orn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetOrnOk() (*string, bool) {
	if o == nil || IsNil(o.Orn) {
		return nil, false
	}
	return o.Orn, true
}

// HasOrn returns a boolean if a field has been set.
func (o *Application) HasOrn() bool {
	if o != nil && !IsNil(o.Orn) {
		return true
	}

	return false
}

// SetOrn gets a reference to the given string and assigns it to the Orn field.
func (o *Application) SetOrn(v string) {
	o.Orn = &v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *Application) GetProfile() map[string]interface{} {
	if o == nil || IsNil(o.Profile) {
		var ret map[string]interface{}
		return ret
	}
	return o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetProfileOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Profile) {
		return map[string]interface{}{}, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *Application) HasProfile() bool {
	if o != nil && !IsNil(o.Profile) {
		return true
	}

	return false
}

// SetProfile gets a reference to the given map[string]interface{} and assigns it to the Profile field.
func (o *Application) SetProfile(v map[string]interface{}) {
	o.Profile = v
}

// GetSignOnMode returns the SignOnMode field value
func (o *Application) GetSignOnMode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SignOnMode
}

// GetSignOnModeOk returns a tuple with the SignOnMode field value
// and a boolean to check if the value has been set.
func (o *Application) GetSignOnModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SignOnMode, true
}

// SetSignOnMode sets field value
func (o *Application) SetSignOnMode(v string) {
	o.SignOnMode = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Application) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Application) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Application) SetStatus(v string) {
	o.Status = &v
}

// GetUniversalLogout returns the UniversalLogout field value if set, zero value otherwise.
func (o *Application) GetUniversalLogout() ApplicationUniversalLogout {
	if o == nil || IsNil(o.UniversalLogout) {
		var ret ApplicationUniversalLogout
		return ret
	}
	return *o.UniversalLogout
}

// GetUniversalLogoutOk returns a tuple with the UniversalLogout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetUniversalLogoutOk() (*ApplicationUniversalLogout, bool) {
	if o == nil || IsNil(o.UniversalLogout) {
		return nil, false
	}
	return o.UniversalLogout, true
}

// HasUniversalLogout returns a boolean if a field has been set.
func (o *Application) HasUniversalLogout() bool {
	if o != nil && !IsNil(o.UniversalLogout) {
		return true
	}

	return false
}

// SetUniversalLogout gets a reference to the given ApplicationUniversalLogout and assigns it to the UniversalLogout field.
func (o *Application) SetUniversalLogout(v ApplicationUniversalLogout) {
	o.UniversalLogout = &v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *Application) GetVisibility() ApplicationVisibility {
	if o == nil || IsNil(o.Visibility) {
		var ret ApplicationVisibility
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetVisibilityOk() (*ApplicationVisibility, bool) {
	if o == nil || IsNil(o.Visibility) {
		return nil, false
	}
	return o.Visibility, true
}

// HasVisibility returns a boolean if a field has been set.
func (o *Application) HasVisibility() bool {
	if o != nil && !IsNil(o.Visibility) {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given ApplicationVisibility and assigns it to the Visibility field.
func (o *Application) SetVisibility(v ApplicationVisibility) {
	o.Visibility = &v
}

// GetEmbedded returns the Embedded field value if set, zero value otherwise.
func (o *Application) GetEmbedded() ApplicationEmbedded {
	if o == nil || IsNil(o.Embedded) {
		var ret ApplicationEmbedded
		return ret
	}
	return *o.Embedded
}

// GetEmbeddedOk returns a tuple with the Embedded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetEmbeddedOk() (*ApplicationEmbedded, bool) {
	if o == nil || IsNil(o.Embedded) {
		return nil, false
	}
	return o.Embedded, true
}

// HasEmbedded returns a boolean if a field has been set.
func (o *Application) HasEmbedded() bool {
	if o != nil && !IsNil(o.Embedded) {
		return true
	}

	return false
}

// SetEmbedded gets a reference to the given ApplicationEmbedded and assigns it to the Embedded field.
func (o *Application) SetEmbedded(v ApplicationEmbedded) {
	o.Embedded = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *Application) GetLinks() ApplicationLinks {
	if o == nil || IsNil(o.Links) {
		var ret ApplicationLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetLinksOk() (*ApplicationLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *Application) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ApplicationLinks and assigns it to the Links field.
func (o *Application) SetLinks(v ApplicationLinks) {
	o.Links = &v
}

func (o Application) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Application) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Accessibility) {
		toSerialize["accessibility"] = o.Accessibility
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.Features) {
		toSerialize["features"] = o.Features
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["label"] = o.Label
	if !IsNil(o.LastUpdated) {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if !IsNil(o.Licensing) {
		toSerialize["licensing"] = o.Licensing
	}
	if !IsNil(o.Orn) {
		toSerialize["orn"] = o.Orn
	}
	if !IsNil(o.Profile) {
		toSerialize["profile"] = o.Profile
	}
	toSerialize["signOnMode"] = o.SignOnMode
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.UniversalLogout) {
		toSerialize["universalLogout"] = o.UniversalLogout
	}
	if !IsNil(o.Visibility) {
		toSerialize["visibility"] = o.Visibility
	}
	if !IsNil(o.Embedded) {
		toSerialize["_embedded"] = o.Embedded
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Application) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"label",
		"signOnMode",
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

	varApplication := _Application{}

	err = json.Unmarshal(data, &varApplication)

	if err != nil {
		return err
	}

	*o = Application(varApplication)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "accessibility")
		delete(additionalProperties, "created")
		delete(additionalProperties, "features")
		delete(additionalProperties, "id")
		delete(additionalProperties, "label")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "licensing")
		delete(additionalProperties, "orn")
		delete(additionalProperties, "profile")
		delete(additionalProperties, "signOnMode")
		delete(additionalProperties, "status")
		delete(additionalProperties, "universalLogout")
		delete(additionalProperties, "visibility")
		delete(additionalProperties, "_embedded")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApplication struct {
	value *Application
	isSet bool
}

func (v NullableApplication) Get() *Application {
	return v.value
}

func (v *NullableApplication) Set(val *Application) {
	v.value = val
	v.isSet = true
}

func (v NullableApplication) IsSet() bool {
	return v.isSet
}

func (v *NullableApplication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplication(val *Application) *NullableApplication {
	return &NullableApplication{value: val, isSet: true}
}

func (v NullableApplication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
