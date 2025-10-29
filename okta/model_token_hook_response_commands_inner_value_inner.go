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
)

// checks if the TokenHookResponseCommandsInnerValueInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenHookResponseCommandsInnerValueInner{}

// TokenHookResponseCommandsInnerValueInner struct for TokenHookResponseCommandsInnerValueInner
type TokenHookResponseCommandsInnerValueInner struct {
	// The name of one of the supported ops: `add`: Add a claim. `replace`: Modify an existing claim and update the token lifetime. `remove`: Remove an existing claim. #### `op: add` notes  <details> <summary>Add a claim</summary>    Add a claim    **Existing JSON**    ```   {     \"employeeId\": \"00u12345678\"   }   ```    **Operation**    ```   {     \"commands\": [       {         \"type\": \"com.okta.assertion.patch\",         \"value\": [           {             \"op\": \"add\",             \"path\": \"/claims/extPatientId\",             \"value\": \"1234\"           }         ]       },       {         \"type\": \"com.okta.assertion.patch\",         \"value\": [           {             \"op\": \"add\",             \"path\": \"/claims/external_guid\",             \"value\": \"F0384685-F87D-474B-848D-2058AC5655A7\"           }         ]       }     ]   }   ```    **Updated JSON**    ```   {     \"employeeId\": \"00u12345678\",     \"extPatientId\": 1234,     \"external_guid\": \"F0384685-F87D-474B-848D-2058AC5655A7\"   }   ```    > **Note:** If you use the `add` operation and include an existing claim in your response with a different value, that value is replaced. Use the `replace` operation instead. If you attempt to remove a system-specific claim or use an invalid operation, the entire PATCH fails and errors are logged in the token hooks events. See `op: replace` notes. </details>  <details> <summary>Add new members to existing JSON objects</summary>    If you have a JSON object in a claim called `employee_profile`, and you want to add the `department_id` member to the claim, the existing JSON is updated by specifying the claim in the path, followed by the name of the object member.    **Existing JSON**    ```   {     \"employee_profile\": {       \"employee_id\": \"1234\",       \"name\": \"Anna\"     }   }   ```    **Operation**    ```   {     \"commands\": [       {         \"type\": \"com.okta.identity.patch\",         \"value\": [           {             \"op\": \"add\",             \"path\": \"/claims/employee_profile/department_id\",             \"value\": \"4947\"           }         ]       }     ]   }   ```    **Updated JSON**    ```   {     \"employee_profile\": {       \"employee_id\": \"1234\",       \"name\": \"Anna\",       \"department_id\": \"4947\"     }   }   ```    > **Note:** If you attempt to add a member within a JSON object that doesn't exist or using an invalid operation, the entire PATCH fails and errors are logged in the token hooks events. </details>  <details> <summary>Add new elements to existing arrays</summary>    Append an element to an array by specifying the name of the array, followed by the index where you want to insert the element in the path. Alternatively, you can specify the array name followed by a hyphen (-) in the path to append an element at the end of the array. For example, you have an array that contains the user's preferred airports, and you want to add a new airport to the array. The existing target JSON object is updated by specifying the claim in the path, followed by the index of where to insert the claim.    **Existing JSON**    ```   {     \"preferred_airports\":[       \"sjc\",       \"sfo\",       \"oak\"     ]   }   ```    **Operation**    ```   {     \"commands\": [       {         \"type\": \"com.okta.identity.patch\",         \"value\": [           {             \"op\": \"add\",             \"path\": \"/claims/preferred_airports/3\",             \"value\": \"lax\"           }         ]       }     ]   }   ```    **Updated JSON**    ```   {     \"preferred_airports\":[       \"sjc\",       \"sfo\",       \"oak\",       \"lax\"     ]   }   ```    > **Note:** If you attempt to add an element within an array that doesn't exist or specify an invalid index, the entire PATCH fails and errors are logged in the token hooks events. </details>  #### `op: replace` notes  <details> <summary>Modify an existing claim</summary>    You can modify (`replace`) existing custom claims or OIDC standard profile claims, such as `birthdate` and `locale`. You can't, however, modify any system-specific claims, such as `iss` or `ver`. Also, you can't modify a claim that isn't currently part of the token in the request payload. Attempting to modify a system-specific claim or using an invalid operation results in the entire PATCH failing and errors logged in the token hooks events.    See [Access Tokens Scopes and Claims](/openapi/okta-oauth/guides/overview/#access-token-scopes-and-claims) for the list of access token-reserved claims that you can't modify.    > **Note:** Although the `aud` and `sub` claims are listed as reserved claims, you can modify those claims in access tokens. You can't modify these claims in ID tokens.    See [ID Token Claims](/openapi/okta-oauth/guides/overview/#id-token-claims) for a list of ID token-reserved claims that you can't modify.    **Existing target JSON object**    ```   {     \"employeeId\": \"00u12345678\",     \"extPatientId\": 1234,     \"external_guid\": \"F0384685-F87D-474B-848D-2058AC5655A7\"   }   ```    **Operation**    ```   {     \"commands\": [       {         \"type\": \"com.okta.identity.patch\",         \"value\": [           {             \"op\": \"replace\",             \"path\": \"/claims/extPatientId\",             \"value\": \"12345\"           },           {             \"op\": \"replace\",             \"path\": \"/claims/external_guid\",             \"value\": \"D1495796-G98E-585C-959E-1269CD6766B8\"           }         ]       }     ]   }   ```    **Updated JSON***    ```   {     \"employeeId\": \"00u12345678\",     \"extPatientId\": 12345,     \"external_guid\": \"D1495796-G98E-585C-959E-1269CD6766B8\"   }   ```  </details>  <details> <summary>Modify members within existing JSON objects and arrays</summary>    Use the `replace` operation to modify members within JSON objects and elements within arrays. For example, you have a JSON object in a claim called `employee_profile`, and you want to update the email address of the employee. The existing target JSON object is updated by specifying the claim in the path, followed by the name of the object member that you want to modify.    **Existing target JSON object**    ```   {     \"employee_profile\": {       \"employee_id\":\"1234\",       \"name\":\"Anna\",       \"email\":\"anna.v@company.com\"       }   }   ```    **Operation**    ```   {     \"commands\": [       {         \"type\": \"com.okta.identity.patch\",         \"value\": [           {             \"op\": \"replace\",             \"path\": \"/claims/employee_profile/email\",             \"value\": \"anna@company.com\"           }         ]       }     ]   }   ```    **Updated JSON**    ```   {     \"employee_profile\": {       \"employee_id\":\"1234\",       \"name\":\"Anna\",       \"email\":\"anna@company.com\"       }   }   ```    > **Note:** If you attempt to modify a member within a JSON object that doesn't exist or use an invalid operation, the entire PATCH fails and errors are logged in the token hooks events.    Similarly, you can replace elements in an array by specifying the array name and the valid index of the element that you want to replace in the path. </details>  <details> <summary>Modify token lifetimes</summary>   You can modify how long the access and ID tokens are valid by specifying the `lifetime` in seconds. The `lifetime` value must be a minimum of five minutes (300 seconds) and a maximum of 24 hours (86,400 seconds).    **Operation**    ```   {     \"commands\": [       {         \"type\": \"com.okta.identity.patch\",         \"value\": [           {             \"op\": \"replace\",             \"path\": \"/token/lifetime/expiration\",             \"value\": 36000           }         ]       },       {         \"type\": \"com.okta.access.patch\",         \"value\": [           {             \"op\": \"replace\",             \"path\": \"/token/lifetime/expiration\",             \"value\": 36000           }         ]       }     ]   }   ```  </details>  #### `op: remove` notes  <details> <summary>Remove a claim</summary>    You can remove existing custom claims or OIDC standard profile claims, such as `birthdate` or `locale`. You can't, however, remove any system-specific claims, such as `iss` or `ver`. You also can't remove a claim that isn't currently part of the token in the request payload. If you attempt to remove a system-specific claim or use an invalid operation, the entire PATCH fails and errors are logged in the token hooks events.    See [Access Tokens Scopes and Claims](/openapi/okta-oauth/guides/overview/#access-token-scopes-and-claims) for the list of access token-reserved claims that you can't modify.    See [ID Token Claims](/openapi/okta-oauth/guides/overview/#id-token-claims) for a list of ID token-reserved claims that you can't modify.    **Operation**    ```   {     \"commands\": [       {         \"type\": \"com.okta.identity.patch\",         \"value\": [           {             \"op\": \"remove\",             \"path\": \"/claims/birthdate\",             \"value\": null           }         ]       },       {         \"type\": \"com.okta.access.patch\",         \"value\": [           {             \"op\": \"remove\",             \"path\": \"/claims/external_guid\"           }         ]       }     ]   }   ```    > **Note:** The `value` property for the `remove` operation isn't required. If you provide it in the response, it should be set to `null`. Providing any other value fails the entire PATCH response.  </details>  <details> <summary>Remove members from existing arrays</summary>    Use the `remove` operation to remove members from existing arrays. For example, you have an array that contains the user's preferred airports, and you want to remove an airport from the array. The existing target JSON object is updated by specifying the array name followed by the index of the element that you want to remove. You don't need to specify a value for the remove operation, but you can specify `null` as the value if you want.    **Existing target JSON object**    ```   {     \"preferred_airports\": [         \"sjc\",         \"lax\",         \"sfo\",         \"oak\"       ]   }   ```    **Operation**    ```   {     \"commands\": [       {         \"type\": \"com.okta.identity.patch\",         \"value\": [           {             \"op\": \"remove\",             \"path\": \"/claims/preferred_airports/1\"           }         ]       }     ]   }   ```    **Updated JSON**    ```   {     \"preferred_airports\": [         \"sjc\",         \"sfo\",         \"oak\"       ]   }   ```  </details>  <details> <summary>Remove members from existing JSON objects</summary>   Use the `remove` operation to remove members from existing JSON objects. Do this by specifying the JSON object in the path, followed by the claim member that you would like to remove. For example, you have an `employee_profile` claim, and you want to remove `email` from it.  **Existing target JSON object**  ``` {   \"employee_profile\": {     \"employee_id\":\"1234\",     \"name\":\"Anna\",     \"email\":\"anna.v@company.com\"     } } ```  **Operation**  ``` {   \"commands\": [     {       \"type\": \"com.okta.identity.patch\",       \"value\": [         {           \"op\": \"remove\",           \"path\": \"/claims/employee_profile/email\"         }       ]     }   ] } ```  **Updated JSON** ``` {   \"employee_profile\": {     \"employee_id\":\"1234\",     \"name\":\"Anna\",     } } ```  </details>
	Op *string `json:"op,omitempty"`
	// Location within the token to apply the operation, specified as a slash-delimited path. When you add, replace, or remove a claim, this path always begins with `/claims/` and is followed by the name of the new claim that you're adding. When you replace a token lifetime, the path should always be `/token/lifetime/expiration`.
	Path                 *string                                        `json:"path,omitempty"`
	Value                *TokenHookResponseCommandsInnerValueInnerValue `json:"value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TokenHookResponseCommandsInnerValueInner TokenHookResponseCommandsInnerValueInner

// NewTokenHookResponseCommandsInnerValueInner instantiates a new TokenHookResponseCommandsInnerValueInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenHookResponseCommandsInnerValueInner() *TokenHookResponseCommandsInnerValueInner {
	this := TokenHookResponseCommandsInnerValueInner{}
	return &this
}

// NewTokenHookResponseCommandsInnerValueInnerWithDefaults instantiates a new TokenHookResponseCommandsInnerValueInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenHookResponseCommandsInnerValueInnerWithDefaults() *TokenHookResponseCommandsInnerValueInner {
	this := TokenHookResponseCommandsInnerValueInner{}
	return &this
}

// GetOp returns the Op field value if set, zero value otherwise.
func (o *TokenHookResponseCommandsInnerValueInner) GetOp() string {
	if o == nil || IsNil(o.Op) {
		var ret string
		return ret
	}
	return *o.Op
}

// GetOpOk returns a tuple with the Op field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenHookResponseCommandsInnerValueInner) GetOpOk() (*string, bool) {
	if o == nil || IsNil(o.Op) {
		return nil, false
	}
	return o.Op, true
}

// HasOp returns a boolean if a field has been set.
func (o *TokenHookResponseCommandsInnerValueInner) HasOp() bool {
	if o != nil && !IsNil(o.Op) {
		return true
	}

	return false
}

// SetOp gets a reference to the given string and assigns it to the Op field.
func (o *TokenHookResponseCommandsInnerValueInner) SetOp(v string) {
	o.Op = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *TokenHookResponseCommandsInnerValueInner) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenHookResponseCommandsInnerValueInner) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *TokenHookResponseCommandsInnerValueInner) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *TokenHookResponseCommandsInnerValueInner) SetPath(v string) {
	o.Path = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *TokenHookResponseCommandsInnerValueInner) GetValue() TokenHookResponseCommandsInnerValueInnerValue {
	if o == nil || IsNil(o.Value) {
		var ret TokenHookResponseCommandsInnerValueInnerValue
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenHookResponseCommandsInnerValueInner) GetValueOk() (*TokenHookResponseCommandsInnerValueInnerValue, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *TokenHookResponseCommandsInnerValueInner) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given TokenHookResponseCommandsInnerValueInnerValue and assigns it to the Value field.
func (o *TokenHookResponseCommandsInnerValueInner) SetValue(v TokenHookResponseCommandsInnerValueInnerValue) {
	o.Value = &v
}

func (o TokenHookResponseCommandsInnerValueInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenHookResponseCommandsInnerValueInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Op) {
		toSerialize["op"] = o.Op
	}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TokenHookResponseCommandsInnerValueInner) UnmarshalJSON(data []byte) (err error) {
	varTokenHookResponseCommandsInnerValueInner := _TokenHookResponseCommandsInnerValueInner{}

	err = json.Unmarshal(data, &varTokenHookResponseCommandsInnerValueInner)

	if err != nil {
		return err
	}

	*o = TokenHookResponseCommandsInnerValueInner(varTokenHookResponseCommandsInnerValueInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "op")
		delete(additionalProperties, "path")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTokenHookResponseCommandsInnerValueInner struct {
	value *TokenHookResponseCommandsInnerValueInner
	isSet bool
}

func (v NullableTokenHookResponseCommandsInnerValueInner) Get() *TokenHookResponseCommandsInnerValueInner {
	return v.value
}

func (v *NullableTokenHookResponseCommandsInnerValueInner) Set(val *TokenHookResponseCommandsInnerValueInner) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenHookResponseCommandsInnerValueInner) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenHookResponseCommandsInnerValueInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenHookResponseCommandsInnerValueInner(val *TokenHookResponseCommandsInnerValueInner) *NullableTokenHookResponseCommandsInnerValueInner {
	return &NullableTokenHookResponseCommandsInnerValueInner{value: val, isSet: true}
}

func (v NullableTokenHookResponseCommandsInnerValueInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenHookResponseCommandsInnerValueInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
