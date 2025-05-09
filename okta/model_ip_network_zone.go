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
	"reflect"
	"strings"
)

// IPNetworkZone struct for IPNetworkZone
type IPNetworkZone struct {
	NetworkZone
	// The IP addresses (range or CIDR form) for an IP Network Zone. The maximum array length is 150 entries for admin-created IP zones, 1000 entries for IP blocklist zones, and 5000 entries for the default system IP Zone.
	Gateways []NetworkZoneAddress `json:"gateways,omitempty"`
	// The IP addresses (range or CIDR form) that are allowed to forward a request from gateway addresses for an IP Network Zone. These proxies are automatically trusted by Threat Insights and used to identify the client IP of a request. The maximum array length is 150 entries for admin-created zones and 5000 entries for the default system IP Zone.
	Proxies []NetworkZoneAddress `json:"proxies,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IPNetworkZone IPNetworkZone

// NewIPNetworkZone instantiates a new IPNetworkZone object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIPNetworkZone(name string, type_ string) *IPNetworkZone {
	this := IPNetworkZone{}
	this.Name = name
	this.Type = type_
	return &this
}

// NewIPNetworkZoneWithDefaults instantiates a new IPNetworkZone object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIPNetworkZoneWithDefaults() *IPNetworkZone {
	this := IPNetworkZone{}
	return &this
}

// GetGateways returns the Gateways field value if set, zero value otherwise.
func (o *IPNetworkZone) GetGateways() []NetworkZoneAddress {
	if o == nil || o.Gateways == nil {
		var ret []NetworkZoneAddress
		return ret
	}
	return o.Gateways
}

// GetGatewaysOk returns a tuple with the Gateways field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPNetworkZone) GetGatewaysOk() ([]NetworkZoneAddress, bool) {
	if o == nil || o.Gateways == nil {
		return nil, false
	}
	return o.Gateways, true
}

// HasGateways returns a boolean if a field has been set.
func (o *IPNetworkZone) HasGateways() bool {
	if o != nil && o.Gateways != nil {
		return true
	}

	return false
}

// SetGateways gets a reference to the given []NetworkZoneAddress and assigns it to the Gateways field.
func (o *IPNetworkZone) SetGateways(v []NetworkZoneAddress) {
	o.Gateways = v
}

// GetProxies returns the Proxies field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IPNetworkZone) GetProxies() []NetworkZoneAddress {
	if o == nil {
		var ret []NetworkZoneAddress
		return ret
	}
	return o.Proxies
}

// GetProxiesOk returns a tuple with the Proxies field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IPNetworkZone) GetProxiesOk() ([]NetworkZoneAddress, bool) {
	if o == nil || o.Proxies == nil {
		return nil, false
	}
	return o.Proxies, true
}

// HasProxies returns a boolean if a field has been set.
func (o *IPNetworkZone) HasProxies() bool {
	if o != nil && o.Proxies != nil {
		return true
	}

	return false
}

// SetProxies gets a reference to the given []NetworkZoneAddress and assigns it to the Proxies field.
func (o *IPNetworkZone) SetProxies(v []NetworkZoneAddress) {
	o.Proxies = v
}

func (o IPNetworkZone) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedNetworkZone, errNetworkZone := json.Marshal(o.NetworkZone)
	if errNetworkZone != nil {
		return []byte{}, errNetworkZone
	}
	errNetworkZone = json.Unmarshal([]byte(serializedNetworkZone), &toSerialize)
	if errNetworkZone != nil {
		return []byte{}, errNetworkZone
	}
	if o.Gateways != nil {
		toSerialize["gateways"] = o.Gateways
	}
	if o.Proxies != nil {
		toSerialize["proxies"] = o.Proxies
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IPNetworkZone) UnmarshalJSON(bytes []byte) (err error) {
	type IPNetworkZoneWithoutEmbeddedStruct struct {
		// The IP addresses (range or CIDR form) for an IP Network Zone. The maximum array length is 150 entries for admin-created IP zones, 1000 entries for IP blocklist zones, and 5000 entries for the default system IP Zone.
		Gateways []NetworkZoneAddress `json:"gateways,omitempty"`
		// The IP addresses (range or CIDR form) that are allowed to forward a request from gateway addresses for an IP Network Zone. These proxies are automatically trusted by Threat Insights and used to identify the client IP of a request. The maximum array length is 150 entries for admin-created zones and 5000 entries for the default system IP Zone.
		Proxies []NetworkZoneAddress `json:"proxies,omitempty"`
	}

	varIPNetworkZoneWithoutEmbeddedStruct := IPNetworkZoneWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varIPNetworkZoneWithoutEmbeddedStruct)
	if err == nil {
		varIPNetworkZone := _IPNetworkZone{}
		varIPNetworkZone.Gateways = varIPNetworkZoneWithoutEmbeddedStruct.Gateways
		varIPNetworkZone.Proxies = varIPNetworkZoneWithoutEmbeddedStruct.Proxies
		*o = IPNetworkZone(varIPNetworkZone)
	} else {
		return err
	}

	varIPNetworkZone := _IPNetworkZone{}

	err = json.Unmarshal(bytes, &varIPNetworkZone)
	if err == nil {
		o.NetworkZone = varIPNetworkZone.NetworkZone
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "gateways")
		delete(additionalProperties, "proxies")

		// remove fields from embedded structs
		reflectNetworkZone := reflect.ValueOf(o.NetworkZone)
		for i := 0; i < reflectNetworkZone.Type().NumField(); i++ {
			t := reflectNetworkZone.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableIPNetworkZone struct {
	value *IPNetworkZone
	isSet bool
}

func (v NullableIPNetworkZone) Get() *IPNetworkZone {
	return v.value
}

func (v *NullableIPNetworkZone) Set(val *IPNetworkZone) {
	v.value = val
	v.isSet = true
}

func (v NullableIPNetworkZone) IsSet() bool {
	return v.isSet
}

func (v *NullableIPNetworkZone) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIPNetworkZone(val *IPNetworkZone) *NullableIPNetworkZone {
	return &NullableIPNetworkZone{value: val, isSet: true}
}

func (v NullableIPNetworkZone) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIPNetworkZone) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

