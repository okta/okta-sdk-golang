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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
	"fmt"
)


//model_oneof.mustache
// GracePeriodExpiry - struct for GracePeriodExpiry
type GracePeriodExpiry struct {
	ByDateTimeExpiry *ByDateTimeExpiry
	ByDurationExpiry *ByDurationExpiry
}

// ByDateTimeExpiryAsGracePeriodExpiry is a convenience function that returns ByDateTimeExpiry wrapped in GracePeriodExpiry
func ByDateTimeExpiryAsGracePeriodExpiry(v *ByDateTimeExpiry) GracePeriodExpiry {
	return GracePeriodExpiry{
		ByDateTimeExpiry: v,
	}
}

// ByDurationExpiryAsGracePeriodExpiry is a convenience function that returns ByDurationExpiry wrapped in GracePeriodExpiry
func ByDurationExpiryAsGracePeriodExpiry(v *ByDurationExpiry) GracePeriodExpiry {
	return GracePeriodExpiry{
		ByDurationExpiry: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct  CUSTOM
func (dst *GracePeriodExpiry) UnmarshalJSON(data []byte) error {
	var err error
        match := 0
        // try to unmarshal data into ByDateTimeExpiry
        err = json.Unmarshal(data, &dst.ByDateTimeExpiry)
        if err == nil {
                jsonByDateTimeExpiry, _ := json.Marshal(dst.ByDateTimeExpiry)
                if string(jsonByDateTimeExpiry) == "{}" { // empty struct
                        dst.ByDateTimeExpiry = nil
                } else {
                        match++
                }
        } else {
                dst.ByDateTimeExpiry = nil
        }

        // try to unmarshal data into ByDurationExpiry
        err = json.Unmarshal(data, &dst.ByDurationExpiry)
        if err == nil {
                jsonByDurationExpiry, _ := json.Marshal(dst.ByDurationExpiry)
                if string(jsonByDurationExpiry) == "{}" { // empty struct
                        dst.ByDurationExpiry = nil
                } else {
                        match++
                }
        } else {
                dst.ByDurationExpiry = nil
        }

        if match > 1 { // more than 1 match
                // reset to nil
                dst.ByDateTimeExpiry = nil
                dst.ByDurationExpiry = nil

                return fmt.Errorf("Data matches more than one schema in oneOf(GracePeriodExpiry)")
        } else if match == 1 {
                return nil // exactly one match
        } else { // no match
                return fmt.Errorf("Data failed to match schemas in oneOf(GracePeriodExpiry)")
        }
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GracePeriodExpiry) MarshalJSON() ([]byte, error) {
	if src.ByDateTimeExpiry != nil {
		return json.Marshal(&src.ByDateTimeExpiry)
	}

	if src.ByDurationExpiry != nil {
		return json.Marshal(&src.ByDurationExpiry)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GracePeriodExpiry) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ByDateTimeExpiry != nil {
		return obj.ByDateTimeExpiry
	}

	if obj.ByDurationExpiry != nil {
		return obj.ByDurationExpiry
	}

	// all schemas are nil
	return nil
}

type NullableGracePeriodExpiry struct {
	value *GracePeriodExpiry
	isSet bool
}

func (v NullableGracePeriodExpiry) Get() *GracePeriodExpiry {
	return v.value
}

func (v *NullableGracePeriodExpiry) Set(val *GracePeriodExpiry) {
	v.value = val
	v.isSet = true
}

func (v NullableGracePeriodExpiry) IsSet() bool {
	return v.isSet
}

func (v *NullableGracePeriodExpiry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGracePeriodExpiry(val *GracePeriodExpiry) *NullableGracePeriodExpiry {
	return &NullableGracePeriodExpiry{value: val, isSet: true}
}

func (v NullableGracePeriodExpiry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGracePeriodExpiry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


