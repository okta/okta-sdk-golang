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
	"fmt"
)


//model_oneof.mustache
// ListApplications200ResponseInner - struct for ListApplications200ResponseInner
type ListApplications200ResponseInner struct {
	AutoLoginApplication *AutoLoginApplication
	BasicAuthApplication *BasicAuthApplication
	BookmarkApplication *BookmarkApplication
	BrowserPluginApplication *BrowserPluginApplication
	OpenIdConnectApplication *OpenIdConnectApplication
	Saml11Application *Saml11Application
	SamlApplication *SamlApplication
	SecurePasswordStoreApplication *SecurePasswordStoreApplication
	WsFederationApplication *WsFederationApplication
}

// AutoLoginApplicationAsListApplications200ResponseInner is a convenience function that returns AutoLoginApplication wrapped in ListApplications200ResponseInner
func AutoLoginApplicationAsListApplications200ResponseInner(v *AutoLoginApplication) ListApplications200ResponseInner {
	return ListApplications200ResponseInner{
		AutoLoginApplication: v,
	}
}

// BasicAuthApplicationAsListApplications200ResponseInner is a convenience function that returns BasicAuthApplication wrapped in ListApplications200ResponseInner
func BasicAuthApplicationAsListApplications200ResponseInner(v *BasicAuthApplication) ListApplications200ResponseInner {
	return ListApplications200ResponseInner{
		BasicAuthApplication: v,
	}
}

// BookmarkApplicationAsListApplications200ResponseInner is a convenience function that returns BookmarkApplication wrapped in ListApplications200ResponseInner
func BookmarkApplicationAsListApplications200ResponseInner(v *BookmarkApplication) ListApplications200ResponseInner {
	return ListApplications200ResponseInner{
		BookmarkApplication: v,
	}
}

// BrowserPluginApplicationAsListApplications200ResponseInner is a convenience function that returns BrowserPluginApplication wrapped in ListApplications200ResponseInner
func BrowserPluginApplicationAsListApplications200ResponseInner(v *BrowserPluginApplication) ListApplications200ResponseInner {
	return ListApplications200ResponseInner{
		BrowserPluginApplication: v,
	}
}

// OpenIdConnectApplicationAsListApplications200ResponseInner is a convenience function that returns OpenIdConnectApplication wrapped in ListApplications200ResponseInner
func OpenIdConnectApplicationAsListApplications200ResponseInner(v *OpenIdConnectApplication) ListApplications200ResponseInner {
	return ListApplications200ResponseInner{
		OpenIdConnectApplication: v,
	}
}

// Saml11ApplicationAsListApplications200ResponseInner is a convenience function that returns Saml11Application wrapped in ListApplications200ResponseInner
func Saml11ApplicationAsListApplications200ResponseInner(v *Saml11Application) ListApplications200ResponseInner {
	return ListApplications200ResponseInner{
		Saml11Application: v,
	}
}

// SamlApplicationAsListApplications200ResponseInner is a convenience function that returns SamlApplication wrapped in ListApplications200ResponseInner
func SamlApplicationAsListApplications200ResponseInner(v *SamlApplication) ListApplications200ResponseInner {
	return ListApplications200ResponseInner{
		SamlApplication: v,
	}
}

// SecurePasswordStoreApplicationAsListApplications200ResponseInner is a convenience function that returns SecurePasswordStoreApplication wrapped in ListApplications200ResponseInner
func SecurePasswordStoreApplicationAsListApplications200ResponseInner(v *SecurePasswordStoreApplication) ListApplications200ResponseInner {
	return ListApplications200ResponseInner{
		SecurePasswordStoreApplication: v,
	}
}

// WsFederationApplicationAsListApplications200ResponseInner is a convenience function that returns WsFederationApplication wrapped in ListApplications200ResponseInner
func WsFederationApplicationAsListApplications200ResponseInner(v *WsFederationApplication) ListApplications200ResponseInner {
	return ListApplications200ResponseInner{
		WsFederationApplication: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct  CUSTOM
func (dst *ListApplications200ResponseInner) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'AUTO_LOGIN'
	if jsonDict["signOnMode"] == "AUTO_LOGIN" {
		// try to unmarshal JSON data into AutoLoginApplication
		err = json.Unmarshal(data, &dst.AutoLoginApplication)
		if err == nil {
			return nil // data stored in dst.AutoLoginApplication, return on the first match
		} else {
			dst.AutoLoginApplication = nil
			return fmt.Errorf("Failed to unmarshal ListApplications200ResponseInner as AutoLoginApplication: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AutoLoginApplication'
	if jsonDict["signOnMode"] == "AutoLoginApplication" {
		// try to unmarshal JSON data into AutoLoginApplication
		err = json.Unmarshal(data, &dst.AutoLoginApplication)
		if err == nil {
			return nil // data stored in dst.AutoLoginApplication, return on the first match
		} else {
			dst.AutoLoginApplication = nil
			return fmt.Errorf("Failed to unmarshal ListApplications200ResponseInner as AutoLoginApplication: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BASIC_AUTH'
	if jsonDict["signOnMode"] == "BASIC_AUTH" {
		// try to unmarshal JSON data into BasicAuthApplication
		err = json.Unmarshal(data, &dst.BasicAuthApplication)
		if err == nil {
			return nil // data stored in dst.BasicAuthApplication, return on the first match
		} else {
			dst.BasicAuthApplication = nil
			return fmt.Errorf("Failed to unmarshal ListApplications200ResponseInner as BasicAuthApplication: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BOOKMARK'
	if jsonDict["signOnMode"] == "BOOKMARK" {
		// try to unmarshal JSON data into BookmarkApplication
		err = json.Unmarshal(data, &dst.BookmarkApplication)
		if err == nil {
			return nil // data stored in dst.BookmarkApplication, return on the first match
		} else {
			dst.BookmarkApplication = nil
			return fmt.Errorf("Failed to unmarshal ListApplications200ResponseInner as BookmarkApplication: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BROWSER_PLUGIN'
	if jsonDict["signOnMode"] == "BROWSER_PLUGIN" {
		// try to unmarshal JSON data into BrowserPluginApplication
		err = json.Unmarshal(data, &dst.BrowserPluginApplication)
		if err == nil {
			return nil // data stored in dst.BrowserPluginApplication, return on the first match
		} else {
			dst.BrowserPluginApplication = nil
			return fmt.Errorf("Failed to unmarshal ListApplications200ResponseInner as BrowserPluginApplication: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BasicAuthApplication'
	if jsonDict["signOnMode"] == "BasicAuthApplication" {
		// try to unmarshal JSON data into BasicAuthApplication
		err = json.Unmarshal(data, &dst.BasicAuthApplication)
		if err == nil {
			return nil // data stored in dst.BasicAuthApplication, return on the first match
		} else {
			dst.BasicAuthApplication = nil
			return fmt.Errorf("Failed to unmarshal ListApplications200ResponseInner as BasicAuthApplication: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BookmarkApplication'
	if jsonDict["signOnMode"] == "BookmarkApplication" {
		// try to unmarshal JSON data into BookmarkApplication
		err = json.Unmarshal(data, &dst.BookmarkApplication)
		if err == nil {
			return nil // data stored in dst.BookmarkApplication, return on the first match
		} else {
			dst.BookmarkApplication = nil
			return fmt.Errorf("Failed to unmarshal ListApplications200ResponseInner as BookmarkApplication: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BrowserPluginApplication'
	if jsonDict["signOnMode"] == "BrowserPluginApplication" {
		// try to unmarshal JSON data into BrowserPluginApplication
		err = json.Unmarshal(data, &dst.BrowserPluginApplication)
		if err == nil {
			return nil // data stored in dst.BrowserPluginApplication, return on the first match
		} else {
			dst.BrowserPluginApplication = nil
			return fmt.Errorf("Failed to unmarshal ListApplications200ResponseInner as BrowserPluginApplication: %s", err.Error())
		}
	}

	// check if the discriminator value is 'OPENID_CONNECT'
	if jsonDict["signOnMode"] == "OPENID_CONNECT" {
		// try to unmarshal JSON data into OpenIdConnectApplication
		err = json.Unmarshal(data, &dst.OpenIdConnectApplication)
		if err == nil {
			return nil // data stored in dst.OpenIdConnectApplication, return on the first match
		} else {
			dst.OpenIdConnectApplication = nil
			return fmt.Errorf("Failed to unmarshal ListApplications200ResponseInner as OpenIdConnectApplication: %s", err.Error())
		}
	}

	// check if the discriminator value is 'OpenIdConnectApplication'
	if jsonDict["signOnMode"] == "OpenIdConnectApplication" {
		// try to unmarshal JSON data into OpenIdConnectApplication
		err = json.Unmarshal(data, &dst.OpenIdConnectApplication)
		if err == nil {
			return nil // data stored in dst.OpenIdConnectApplication, return on the first match
		} else {
			dst.OpenIdConnectApplication = nil
			return fmt.Errorf("Failed to unmarshal ListApplications200ResponseInner as OpenIdConnectApplication: %s", err.Error())
		}
	}

	// check if the discriminator value is 'SAML_1_1'
	if jsonDict["signOnMode"] == "SAML_1_1" {
		// try to unmarshal JSON data into Saml11Application
		err = json.Unmarshal(data, &dst.Saml11Application)
		if err == nil {
			return nil // data stored in dst.Saml11Application, return on the first match
		} else {
			dst.Saml11Application = nil
			return fmt.Errorf("Failed to unmarshal ListApplications200ResponseInner as Saml11Application: %s", err.Error())
		}
	}

	// check if the discriminator value is 'SAML_2_0'
	if jsonDict["signOnMode"] == "SAML_2_0" {
		// try to unmarshal JSON data into SamlApplication
		err = json.Unmarshal(data, &dst.SamlApplication)
		if err == nil {
			return nil // data stored in dst.SamlApplication, return on the first match
		} else {
			dst.SamlApplication = nil
			return fmt.Errorf("Failed to unmarshal ListApplications200ResponseInner as SamlApplication: %s", err.Error())
		}
	}

	// check if the discriminator value is 'SECURE_PASSWORD_STORE'
	if jsonDict["signOnMode"] == "SECURE_PASSWORD_STORE" {
		// try to unmarshal JSON data into SecurePasswordStoreApplication
		err = json.Unmarshal(data, &dst.SecurePasswordStoreApplication)
		if err == nil {
			return nil // data stored in dst.SecurePasswordStoreApplication, return on the first match
		} else {
			dst.SecurePasswordStoreApplication = nil
			return fmt.Errorf("Failed to unmarshal ListApplications200ResponseInner as SecurePasswordStoreApplication: %s", err.Error())
		}
	}

	// check if the discriminator value is 'Saml11Application'
	if jsonDict["signOnMode"] == "Saml11Application" {
		// try to unmarshal JSON data into Saml11Application
		err = json.Unmarshal(data, &dst.Saml11Application)
		if err == nil {
			return nil // data stored in dst.Saml11Application, return on the first match
		} else {
			dst.Saml11Application = nil
			return fmt.Errorf("Failed to unmarshal ListApplications200ResponseInner as Saml11Application: %s", err.Error())
		}
	}

	// check if the discriminator value is 'SamlApplication'
	if jsonDict["signOnMode"] == "SamlApplication" {
		// try to unmarshal JSON data into SamlApplication
		err = json.Unmarshal(data, &dst.SamlApplication)
		if err == nil {
			return nil // data stored in dst.SamlApplication, return on the first match
		} else {
			dst.SamlApplication = nil
			return fmt.Errorf("Failed to unmarshal ListApplications200ResponseInner as SamlApplication: %s", err.Error())
		}
	}

	// check if the discriminator value is 'SecurePasswordStoreApplication'
	if jsonDict["signOnMode"] == "SecurePasswordStoreApplication" {
		// try to unmarshal JSON data into SecurePasswordStoreApplication
		err = json.Unmarshal(data, &dst.SecurePasswordStoreApplication)
		if err == nil {
			return nil // data stored in dst.SecurePasswordStoreApplication, return on the first match
		} else {
			dst.SecurePasswordStoreApplication = nil
			return fmt.Errorf("Failed to unmarshal ListApplications200ResponseInner as SecurePasswordStoreApplication: %s", err.Error())
		}
	}

	// check if the discriminator value is 'WS_FEDERATION'
	if jsonDict["signOnMode"] == "WS_FEDERATION" {
		// try to unmarshal JSON data into WsFederationApplication
		err = json.Unmarshal(data, &dst.WsFederationApplication)
		if err == nil {
			return nil // data stored in dst.WsFederationApplication, return on the first match
		} else {
			dst.WsFederationApplication = nil
			return fmt.Errorf("Failed to unmarshal ListApplications200ResponseInner as WsFederationApplication: %s", err.Error())
		}
	}

	// check if the discriminator value is 'WsFederationApplication'
	if jsonDict["signOnMode"] == "WsFederationApplication" {
		// try to unmarshal JSON data into WsFederationApplication
		err = json.Unmarshal(data, &dst.WsFederationApplication)
		if err == nil {
			return nil // data stored in dst.WsFederationApplication, return on the first match
		} else {
			dst.WsFederationApplication = nil
			return fmt.Errorf("Failed to unmarshal ListApplications200ResponseInner as WsFederationApplication: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ListApplications200ResponseInner) MarshalJSON() ([]byte, error) {
	if src.AutoLoginApplication != nil {
		return json.Marshal(&src.AutoLoginApplication)
	}

	if src.BasicAuthApplication != nil {
		return json.Marshal(&src.BasicAuthApplication)
	}

	if src.BookmarkApplication != nil {
		return json.Marshal(&src.BookmarkApplication)
	}

	if src.BrowserPluginApplication != nil {
		return json.Marshal(&src.BrowserPluginApplication)
	}

	if src.OpenIdConnectApplication != nil {
		return json.Marshal(&src.OpenIdConnectApplication)
	}

	if src.Saml11Application != nil {
		return json.Marshal(&src.Saml11Application)
	}

	if src.SamlApplication != nil {
		return json.Marshal(&src.SamlApplication)
	}

	if src.SecurePasswordStoreApplication != nil {
		return json.Marshal(&src.SecurePasswordStoreApplication)
	}

	if src.WsFederationApplication != nil {
		return json.Marshal(&src.WsFederationApplication)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ListApplications200ResponseInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.AutoLoginApplication != nil {
		return obj.AutoLoginApplication
	}

	if obj.BasicAuthApplication != nil {
		return obj.BasicAuthApplication
	}

	if obj.BookmarkApplication != nil {
		return obj.BookmarkApplication
	}

	if obj.BrowserPluginApplication != nil {
		return obj.BrowserPluginApplication
	}

	if obj.OpenIdConnectApplication != nil {
		return obj.OpenIdConnectApplication
	}

	if obj.Saml11Application != nil {
		return obj.Saml11Application
	}

	if obj.SamlApplication != nil {
		return obj.SamlApplication
	}

	if obj.SecurePasswordStoreApplication != nil {
		return obj.SecurePasswordStoreApplication
	}

	if obj.WsFederationApplication != nil {
		return obj.WsFederationApplication
	}

	// all schemas are nil
	return nil
}

type NullableListApplications200ResponseInner struct {
	value *ListApplications200ResponseInner
	isSet bool
}

func (v NullableListApplications200ResponseInner) Get() *ListApplications200ResponseInner {
	return v.value
}

func (v *NullableListApplications200ResponseInner) Set(val *ListApplications200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListApplications200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListApplications200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListApplications200ResponseInner(val *ListApplications200ResponseInner) *NullableListApplications200ResponseInner {
	return &NullableListApplications200ResponseInner{value: val, isSet: true}
}

func (v NullableListApplications200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListApplications200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


