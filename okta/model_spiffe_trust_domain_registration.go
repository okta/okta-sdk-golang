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
	"fmt"
	"time"
)

// checks if the SpiffeTrustDomainRegistration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpiffeTrustDomainRegistration{}

// SpiffeTrustDomainRegistration Represents a SPIFFE trust domain registration, which binds a SPIFFE trust domain to the HTTPS endpoint that serves its JWK bundle. Workloads whose SPIFFE IDs belong to the registered trust domain can authenticate against your Okta org using SPIFFE JWT-SVIDs, with signing keys fetched from the registered `bundleUrl`.
type SpiffeTrustDomainRegistration struct {
	// HTTPS URL of the SPIFFE bundle endpoint that serves the trust domain's JWK Set. Okta fetches this URL on demand to resolve the signing keys for incoming SPIFFE JWT-SVID client assertions.
	BundleUrl string `json:"bundleUrl" validate:"regexp=^https:\\/\\/"`
	// Timestamp when the registration was created
	Created *time.Time `json:"created,omitempty"`
	// Unique identifier of the SPIFFE trust domain registration
	Id *string `json:"id,omitempty"`
	// Timestamp when the registration was last updated
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
	// The trust domain of a workload's SPIFFE ID. For example, in the SPIFFE ID `spiffe://example.org/payments-service`, the trust domain is `example.org`. The trust domain is immutable after creation; use the DELETE endpoint and create a new registration if the trust domain changes.
	TrustDomain          string `json:"trustDomain"`
	AdditionalProperties map[string]interface{}
}

type _SpiffeTrustDomainRegistration SpiffeTrustDomainRegistration

// NewSpiffeTrustDomainRegistration instantiates a new SpiffeTrustDomainRegistration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpiffeTrustDomainRegistration(bundleUrl string, trustDomain string) *SpiffeTrustDomainRegistration {
	this := SpiffeTrustDomainRegistration{}
	this.BundleUrl = bundleUrl
	this.TrustDomain = trustDomain
	return &this
}

// NewSpiffeTrustDomainRegistrationWithDefaults instantiates a new SpiffeTrustDomainRegistration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpiffeTrustDomainRegistrationWithDefaults() *SpiffeTrustDomainRegistration {
	this := SpiffeTrustDomainRegistration{}
	return &this
}

// GetBundleUrl returns the BundleUrl field value
func (o *SpiffeTrustDomainRegistration) GetBundleUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BundleUrl
}

// GetBundleUrlOk returns a tuple with the BundleUrl field value
// and a boolean to check if the value has been set.
func (o *SpiffeTrustDomainRegistration) GetBundleUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BundleUrl, true
}

// SetBundleUrl sets field value
func (o *SpiffeTrustDomainRegistration) SetBundleUrl(v string) {
	o.BundleUrl = v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *SpiffeTrustDomainRegistration) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpiffeTrustDomainRegistration) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *SpiffeTrustDomainRegistration) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *SpiffeTrustDomainRegistration) SetCreated(v time.Time) {
	o.Created = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SpiffeTrustDomainRegistration) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpiffeTrustDomainRegistration) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SpiffeTrustDomainRegistration) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SpiffeTrustDomainRegistration) SetId(v string) {
	o.Id = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *SpiffeTrustDomainRegistration) GetLastUpdated() time.Time {
	if o == nil || IsNil(o.LastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpiffeTrustDomainRegistration) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *SpiffeTrustDomainRegistration) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *SpiffeTrustDomainRegistration) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetTrustDomain returns the TrustDomain field value
func (o *SpiffeTrustDomainRegistration) GetTrustDomain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TrustDomain
}

// GetTrustDomainOk returns a tuple with the TrustDomain field value
// and a boolean to check if the value has been set.
func (o *SpiffeTrustDomainRegistration) GetTrustDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TrustDomain, true
}

// SetTrustDomain sets field value
func (o *SpiffeTrustDomainRegistration) SetTrustDomain(v string) {
	o.TrustDomain = v
}

func (o SpiffeTrustDomainRegistration) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpiffeTrustDomainRegistration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["bundleUrl"] = o.BundleUrl
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.LastUpdated) {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	toSerialize["trustDomain"] = o.TrustDomain

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SpiffeTrustDomainRegistration) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"bundleUrl",
		"trustDomain",
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

	varSpiffeTrustDomainRegistration := _SpiffeTrustDomainRegistration{}

	err = json.Unmarshal(data, &varSpiffeTrustDomainRegistration)

	if err != nil {
		return err
	}

	*o = SpiffeTrustDomainRegistration(varSpiffeTrustDomainRegistration)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "bundleUrl")
		delete(additionalProperties, "created")
		delete(additionalProperties, "id")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "trustDomain")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSpiffeTrustDomainRegistration struct {
	value *SpiffeTrustDomainRegistration
	isSet bool
}

func (v NullableSpiffeTrustDomainRegistration) Get() *SpiffeTrustDomainRegistration {
	return v.value
}

func (v *NullableSpiffeTrustDomainRegistration) Set(val *SpiffeTrustDomainRegistration) {
	v.value = val
	v.isSet = true
}

func (v NullableSpiffeTrustDomainRegistration) IsSet() bool {
	return v.isSet
}

func (v *NullableSpiffeTrustDomainRegistration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpiffeTrustDomainRegistration(val *SpiffeTrustDomainRegistration) *NullableSpiffeTrustDomainRegistration {
	return &NullableSpiffeTrustDomainRegistration{value: val, isSet: true}
}

func (v NullableSpiffeTrustDomainRegistration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpiffeTrustDomainRegistration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
