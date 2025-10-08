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

// GracePeriodExpiry - struct for GracePeriodExpiry
type GracePeriodExpiry struct {
	String   *string
	TimeTime *time.Time
}

// stringAsGracePeriodExpiry is a convenience function that returns string wrapped in GracePeriodExpiry
func StringAsGracePeriodExpiry(v *string) GracePeriodExpiry {
	return GracePeriodExpiry{
		String: v,
	}
}

// time.TimeAsGracePeriodExpiry is a convenience function that returns time.Time wrapped in GracePeriodExpiry
func TimeTimeAsGracePeriodExpiry(v *time.Time) GracePeriodExpiry {
	return GracePeriodExpiry{
		TimeTime: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *GracePeriodExpiry) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into String
	err = json.Unmarshal(data, &dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			match++
		}
	} else {
		dst.String = nil
	}

	// try to unmarshal data into TimeTime
	err = json.Unmarshal(data, &dst.TimeTime)
	if err == nil {
		jsonTimeTime, _ := json.Marshal(dst.TimeTime)
		if string(jsonTimeTime) == "{}" { // empty struct
			dst.TimeTime = nil
		} else {
			match++
		}
	} else {
		dst.TimeTime = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.String = nil
		dst.TimeTime = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GracePeriodExpiry)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GracePeriodExpiry)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GracePeriodExpiry) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	if src.TimeTime != nil {
		return json.Marshal(&src.TimeTime)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GracePeriodExpiry) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.String != nil {
		return obj.String
	}

	if obj.TimeTime != nil {
		return obj.TimeTime
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj GracePeriodExpiry) GetActualInstanceValue() interface{} {
	if obj.String != nil {
		return *obj.String
	}

	if obj.TimeTime != nil {
		return *obj.TimeTime
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
