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

// checks if the CimdClientEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CimdClientEntity{}

// CimdClientEntity Represents a CIMD client entity, which defines rules for matching incoming CIMD clients by client ID pattern and optionally enforcing metadata requirements
type CimdClientEntity struct {
	// Optional list of bindings to create alongside the CIMD client entity
	Bindings []CimdClientEntityBindingRequest `json:"bindings,omitempty"`
	// The client ID match pattern. * For the `EXACT` strategy, this is a single CIMD client ID URL that's stored in normalized form (lowercase and with the trailing slash removed). * For the `REGEX` strategy, this is a regular expression (regex) evaluated against the client's `client_id` URL.
	ClientIdMatchPattern string `json:"clientIdMatchPattern"`
	// The strategy that Okta uses to match a `client_id` URL to a CIMD client entity
	ClientIdMatchStrategy string `json:"clientIdMatchStrategy"`
	// Timestamp when the object was created
	Created *time.Time `json:"created,omitempty"`
	// Description of the CIMD client entity
	Description *string `json:"description,omitempty"`
	// Unique identifier of the CIMD client entity
	Id *string `json:"id,omitempty"`
	// Timestamp when the object was last updated
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
	// The ID of a [CIMD requirement](https://developer.okta.com/docs/api/openapi/okta-management/management/tag/CimdDocumentRequirements/) that Okta evaluates against the metadata document of any CIMD client this entity matches. If any requirement fails, Okta rejects the client's authorization request.
	MetadataRequirementsId NullableString `json:"metadataRequirementsId,omitempty"`
	// Human-readable name of the CIMD client entity
	Name string `json:"name"`
	// Determines the evaluation order when multiple `REGEX`-strategy client entities can match the same `client_id` URL. You must have a value for `priority` if you use a `REGEX` matching strategy and the value must be unique across all `REGEX`-strategy client entities in the org. For example, only one `REGEX`-strategy client entity can have a priority of `0`, only one can have a priority of `1`, and so on.  If you use the `EXACT` strategy, `priority` must be null.  When a `client_id` URL matches more than one `REGEX`-strategy client entity, Okta evaluates client entities from lowest to highest priority value and stops at the first match. Okta applies the matching client entity's configuration, including any associated document requirements, to the authorization request.
	Priority             NullableInt32             `json:"priority,omitempty"`
	Embedded             *CimdClientEntityEmbedded `json:"_embedded,omitempty"`
	Links                *CimdClientEntityLinks    `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CimdClientEntity CimdClientEntity

// NewCimdClientEntity instantiates a new CimdClientEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCimdClientEntity(clientIdMatchPattern string, clientIdMatchStrategy string, name string) *CimdClientEntity {
	this := CimdClientEntity{}
	this.ClientIdMatchPattern = clientIdMatchPattern
	this.ClientIdMatchStrategy = clientIdMatchStrategy
	this.Name = name
	return &this
}

// NewCimdClientEntityWithDefaults instantiates a new CimdClientEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCimdClientEntityWithDefaults() *CimdClientEntity {
	this := CimdClientEntity{}
	return &this
}

// GetBindings returns the Bindings field value if set, zero value otherwise.
func (o *CimdClientEntity) GetBindings() []CimdClientEntityBindingRequest {
	if o == nil || IsNil(o.Bindings) {
		var ret []CimdClientEntityBindingRequest
		return ret
	}
	return o.Bindings
}

// GetBindingsOk returns a tuple with the Bindings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CimdClientEntity) GetBindingsOk() ([]CimdClientEntityBindingRequest, bool) {
	if o == nil || IsNil(o.Bindings) {
		return nil, false
	}
	return o.Bindings, true
}

// HasBindings returns a boolean if a field has been set.
func (o *CimdClientEntity) HasBindings() bool {
	if o != nil && !IsNil(o.Bindings) {
		return true
	}

	return false
}

// SetBindings gets a reference to the given []CimdClientEntityBindingRequest and assigns it to the Bindings field.
func (o *CimdClientEntity) SetBindings(v []CimdClientEntityBindingRequest) {
	o.Bindings = v
}

// GetClientIdMatchPattern returns the ClientIdMatchPattern field value
func (o *CimdClientEntity) GetClientIdMatchPattern() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientIdMatchPattern
}

// GetClientIdMatchPatternOk returns a tuple with the ClientIdMatchPattern field value
// and a boolean to check if the value has been set.
func (o *CimdClientEntity) GetClientIdMatchPatternOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientIdMatchPattern, true
}

// SetClientIdMatchPattern sets field value
func (o *CimdClientEntity) SetClientIdMatchPattern(v string) {
	o.ClientIdMatchPattern = v
}

// GetClientIdMatchStrategy returns the ClientIdMatchStrategy field value
func (o *CimdClientEntity) GetClientIdMatchStrategy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientIdMatchStrategy
}

// GetClientIdMatchStrategyOk returns a tuple with the ClientIdMatchStrategy field value
// and a boolean to check if the value has been set.
func (o *CimdClientEntity) GetClientIdMatchStrategyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientIdMatchStrategy, true
}

// SetClientIdMatchStrategy sets field value
func (o *CimdClientEntity) SetClientIdMatchStrategy(v string) {
	o.ClientIdMatchStrategy = v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *CimdClientEntity) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CimdClientEntity) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *CimdClientEntity) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *CimdClientEntity) SetCreated(v time.Time) {
	o.Created = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CimdClientEntity) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CimdClientEntity) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CimdClientEntity) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CimdClientEntity) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CimdClientEntity) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CimdClientEntity) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CimdClientEntity) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CimdClientEntity) SetId(v string) {
	o.Id = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *CimdClientEntity) GetLastUpdated() time.Time {
	if o == nil || IsNil(o.LastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CimdClientEntity) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *CimdClientEntity) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *CimdClientEntity) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetMetadataRequirementsId returns the MetadataRequirementsId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CimdClientEntity) GetMetadataRequirementsId() string {
	if o == nil || IsNil(o.MetadataRequirementsId.Get()) {
		var ret string
		return ret
	}
	return *o.MetadataRequirementsId.Get()
}

// GetMetadataRequirementsIdOk returns a tuple with the MetadataRequirementsId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CimdClientEntity) GetMetadataRequirementsIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MetadataRequirementsId.Get(), o.MetadataRequirementsId.IsSet()
}

// HasMetadataRequirementsId returns a boolean if a field has been set.
func (o *CimdClientEntity) HasMetadataRequirementsId() bool {
	if o != nil && o.MetadataRequirementsId.IsSet() {
		return true
	}

	return false
}

// SetMetadataRequirementsId gets a reference to the given NullableString and assigns it to the MetadataRequirementsId field.
func (o *CimdClientEntity) SetMetadataRequirementsId(v string) {
	o.MetadataRequirementsId.Set(&v)
}

// SetMetadataRequirementsIdNil sets the value for MetadataRequirementsId to be an explicit nil
func (o *CimdClientEntity) SetMetadataRequirementsIdNil() {
	o.MetadataRequirementsId.Set(nil)
}

// UnsetMetadataRequirementsId ensures that no value is present for MetadataRequirementsId, not even an explicit nil
func (o *CimdClientEntity) UnsetMetadataRequirementsId() {
	o.MetadataRequirementsId.Unset()
}

// GetName returns the Name field value
func (o *CimdClientEntity) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CimdClientEntity) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CimdClientEntity) SetName(v string) {
	o.Name = v
}

// GetPriority returns the Priority field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CimdClientEntity) GetPriority() int32 {
	if o == nil || IsNil(o.Priority.Get()) {
		var ret int32
		return ret
	}
	return *o.Priority.Get()
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CimdClientEntity) GetPriorityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Priority.Get(), o.Priority.IsSet()
}

// HasPriority returns a boolean if a field has been set.
func (o *CimdClientEntity) HasPriority() bool {
	if o != nil && o.Priority.IsSet() {
		return true
	}

	return false
}

// SetPriority gets a reference to the given NullableInt32 and assigns it to the Priority field.
func (o *CimdClientEntity) SetPriority(v int32) {
	o.Priority.Set(&v)
}

// SetPriorityNil sets the value for Priority to be an explicit nil
func (o *CimdClientEntity) SetPriorityNil() {
	o.Priority.Set(nil)
}

// UnsetPriority ensures that no value is present for Priority, not even an explicit nil
func (o *CimdClientEntity) UnsetPriority() {
	o.Priority.Unset()
}

// GetEmbedded returns the Embedded field value if set, zero value otherwise.
func (o *CimdClientEntity) GetEmbedded() CimdClientEntityEmbedded {
	if o == nil || IsNil(o.Embedded) {
		var ret CimdClientEntityEmbedded
		return ret
	}
	return *o.Embedded
}

// GetEmbeddedOk returns a tuple with the Embedded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CimdClientEntity) GetEmbeddedOk() (*CimdClientEntityEmbedded, bool) {
	if o == nil || IsNil(o.Embedded) {
		return nil, false
	}
	return o.Embedded, true
}

// HasEmbedded returns a boolean if a field has been set.
func (o *CimdClientEntity) HasEmbedded() bool {
	if o != nil && !IsNil(o.Embedded) {
		return true
	}

	return false
}

// SetEmbedded gets a reference to the given CimdClientEntityEmbedded and assigns it to the Embedded field.
func (o *CimdClientEntity) SetEmbedded(v CimdClientEntityEmbedded) {
	o.Embedded = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *CimdClientEntity) GetLinks() CimdClientEntityLinks {
	if o == nil || IsNil(o.Links) {
		var ret CimdClientEntityLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CimdClientEntity) GetLinksOk() (*CimdClientEntityLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *CimdClientEntity) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given CimdClientEntityLinks and assigns it to the Links field.
func (o *CimdClientEntity) SetLinks(v CimdClientEntityLinks) {
	o.Links = &v
}

func (o CimdClientEntity) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CimdClientEntity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Bindings) {
		toSerialize["bindings"] = o.Bindings
	}
	toSerialize["clientIdMatchPattern"] = o.ClientIdMatchPattern
	toSerialize["clientIdMatchStrategy"] = o.ClientIdMatchStrategy
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.LastUpdated) {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if o.MetadataRequirementsId.IsSet() {
		toSerialize["metadataRequirementsId"] = o.MetadataRequirementsId.Get()
	}
	toSerialize["name"] = o.Name
	if o.Priority.IsSet() {
		toSerialize["priority"] = o.Priority.Get()
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

func (o *CimdClientEntity) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"clientIdMatchPattern",
		"clientIdMatchStrategy",
		"name",
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

	varCimdClientEntity := _CimdClientEntity{}

	err = json.Unmarshal(data, &varCimdClientEntity)

	if err != nil {
		return err
	}

	*o = CimdClientEntity(varCimdClientEntity)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "bindings")
		delete(additionalProperties, "clientIdMatchPattern")
		delete(additionalProperties, "clientIdMatchStrategy")
		delete(additionalProperties, "created")
		delete(additionalProperties, "description")
		delete(additionalProperties, "id")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "metadataRequirementsId")
		delete(additionalProperties, "name")
		delete(additionalProperties, "priority")
		delete(additionalProperties, "_embedded")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCimdClientEntity struct {
	value *CimdClientEntity
	isSet bool
}

func (v NullableCimdClientEntity) Get() *CimdClientEntity {
	return v.value
}

func (v *NullableCimdClientEntity) Set(val *CimdClientEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableCimdClientEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableCimdClientEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCimdClientEntity(val *CimdClientEntity) *NullableCimdClientEntity {
	return &NullableCimdClientEntity{value: val, isSet: true}
}

func (v NullableCimdClientEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCimdClientEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
