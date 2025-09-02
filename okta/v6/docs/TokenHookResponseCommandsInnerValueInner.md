# TokenHookResponseCommandsInnerValueInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Op** | Pointer to **string** | The name of one of the supported ops: &#x60;add&#x60;: Add a claim. &#x60;replace&#x60;: Modify an existing claim and update the token lifetime. &#x60;remove&#x60;: Remove an existing claim. #### &#x60;op: add&#x60; notes  &lt;details&gt; &lt;summary&gt;Add a claim&lt;/summary&gt;    Add a claim    **Existing JSON**    &#x60;&#x60;&#x60;   {     \&quot;employeeId\&quot;: \&quot;00u12345678\&quot;   }   &#x60;&#x60;&#x60;    **Operation**    &#x60;&#x60;&#x60;   {     \&quot;commands\&quot;: [       {         \&quot;type\&quot;: \&quot;com.okta.assertion.patch\&quot;,         \&quot;value\&quot;: [           {             \&quot;op\&quot;: \&quot;add\&quot;,             \&quot;path\&quot;: \&quot;/claims/extPatientId\&quot;,             \&quot;value\&quot;: \&quot;1234\&quot;           }         ]       },       {         \&quot;type\&quot;: \&quot;com.okta.assertion.patch\&quot;,         \&quot;value\&quot;: [           {             \&quot;op\&quot;: \&quot;add\&quot;,             \&quot;path\&quot;: \&quot;/claims/external_guid\&quot;,             \&quot;value\&quot;: \&quot;F0384685-F87D-474B-848D-2058AC5655A7\&quot;           }         ]       }     ]   }   &#x60;&#x60;&#x60;    **Updated JSON**    &#x60;&#x60;&#x60;   {     \&quot;employeeId\&quot;: \&quot;00u12345678\&quot;,     \&quot;extPatientId\&quot;: 1234,     \&quot;external_guid\&quot;: \&quot;F0384685-F87D-474B-848D-2058AC5655A7\&quot;   }   &#x60;&#x60;&#x60;    &gt; **Note:** If you use the &#x60;add&#x60; operation and include an existing claim in your response with a different value, that value is replaced. Use the &#x60;replace&#x60; operation instead. If you attempt to remove a system-specific claim or use an invalid operation, the entire PATCH fails and errors are logged in the token hooks events. See &#x60;op: replace&#x60; notes. &lt;/details&gt;  &lt;details&gt; &lt;summary&gt;Add new members to existing JSON objects&lt;/summary&gt;    If you have a JSON object in a claim called &#x60;employee_profile&#x60;, and you want to add the &#x60;department_id&#x60; member to the claim, the existing JSON is updated by specifying the claim in the path, followed by the name of the object member.    **Existing JSON**    &#x60;&#x60;&#x60;   {     \&quot;employee_profile\&quot;: {       \&quot;employee_id\&quot;: \&quot;1234\&quot;,       \&quot;name\&quot;: \&quot;Anna\&quot;     }   }   &#x60;&#x60;&#x60;    **Operation**    &#x60;&#x60;&#x60;   {     \&quot;commands\&quot;: [       {         \&quot;type\&quot;: \&quot;com.okta.identity.patch\&quot;,         \&quot;value\&quot;: [           {             \&quot;op\&quot;: \&quot;add\&quot;,             \&quot;path\&quot;: \&quot;/claims/employee_profile/department_id\&quot;,             \&quot;value\&quot;: \&quot;4947\&quot;           }         ]       }     ]   }   &#x60;&#x60;&#x60;    **Updated JSON**    &#x60;&#x60;&#x60;   {     \&quot;employee_profile\&quot;: {       \&quot;employee_id\&quot;: \&quot;1234\&quot;,       \&quot;name\&quot;: \&quot;Anna\&quot;,       \&quot;department_id\&quot;: \&quot;4947\&quot;     }   }   &#x60;&#x60;&#x60;    &gt; **Note:** If you attempt to add a member within a JSON object that doesn&#39;t exist or using an invalid operation, the entire PATCH fails and errors are logged in the token hooks events. &lt;/details&gt;  &lt;details&gt; &lt;summary&gt;Add new elements to existing arrays&lt;/summary&gt;    Append an element to an array by specifying the name of the array, followed by the index where you want to insert the element in the path. Alternatively, you can specify the array name followed by a hyphen (-) in the path to append an element at the end of the array. For example, you have an array that contains the user&#39;s preferred airports, and you want to add a new airport to the array. The existing target JSON object is updated by specifying the claim in the path, followed by the index of where to insert the claim.    **Existing JSON**    &#x60;&#x60;&#x60;   {     \&quot;preferred_airports\&quot;:[       \&quot;sjc\&quot;,       \&quot;sfo\&quot;,       \&quot;oak\&quot;     ]   }   &#x60;&#x60;&#x60;    **Operation**    &#x60;&#x60;&#x60;   {     \&quot;commands\&quot;: [       {         \&quot;type\&quot;: \&quot;com.okta.identity.patch\&quot;,         \&quot;value\&quot;: [           {             \&quot;op\&quot;: \&quot;add\&quot;,             \&quot;path\&quot;: \&quot;/claims/preferred_airports/3\&quot;,             \&quot;value\&quot;: \&quot;lax\&quot;           }         ]       }     ]   }   &#x60;&#x60;&#x60;    **Updated JSON**    &#x60;&#x60;&#x60;   {     \&quot;preferred_airports\&quot;:[       \&quot;sjc\&quot;,       \&quot;sfo\&quot;,       \&quot;oak\&quot;,       \&quot;lax\&quot;     ]   }   &#x60;&#x60;&#x60;    &gt; **Note:** If you attempt to add an element within an array that doesn&#39;t exist or specify an invalid index, the entire PATCH fails and errors are logged in the token hooks events. &lt;/details&gt;  #### &#x60;op: replace&#x60; notes  &lt;details&gt; &lt;summary&gt;Modify an existing claim&lt;/summary&gt;    You can modify (&#x60;replace&#x60;) existing custom claims or OIDC standard profile claims, such as &#x60;birthdate&#x60; and &#x60;locale&#x60;. You can&#39;t, however, modify any system-specific claims, such as &#x60;iss&#x60; or &#x60;ver&#x60;. Also, you can&#39;t modify a claim that isn&#39;t currently part of the token in the request payload. Attempting to modify a system-specific claim or using an invalid operation results in the entire PATCH failing and errors logged in the token hooks events.    See [Access Tokens Scopes and Claims](/openapi/okta-oauth/guides/overview/#access-token-scopes-and-claims) for the list of access token-reserved claims that you can&#39;t modify.    &gt; **Note:** Although the &#x60;aud&#x60; and &#x60;sub&#x60; claims are listed as reserved claims, you can modify those claims in access tokens. You can&#39;t modify these claims in ID tokens.    See [ID Token Claims](/openapi/okta-oauth/guides/overview/#id-token-claims) for a list of ID token-reserved claims that you can&#39;t modify.    **Existing target JSON object**    &#x60;&#x60;&#x60;   {     \&quot;employeeId\&quot;: \&quot;00u12345678\&quot;,     \&quot;extPatientId\&quot;: 1234,     \&quot;external_guid\&quot;: \&quot;F0384685-F87D-474B-848D-2058AC5655A7\&quot;   }   &#x60;&#x60;&#x60;    **Operation**    &#x60;&#x60;&#x60;   {     \&quot;commands\&quot;: [       {         \&quot;type\&quot;: \&quot;com.okta.identity.patch\&quot;,         \&quot;value\&quot;: [           {             \&quot;op\&quot;: \&quot;replace\&quot;,             \&quot;path\&quot;: \&quot;/claims/extPatientId\&quot;,             \&quot;value\&quot;: \&quot;12345\&quot;           },           {             \&quot;op\&quot;: \&quot;replace\&quot;,             \&quot;path\&quot;: \&quot;/claims/external_guid\&quot;,             \&quot;value\&quot;: \&quot;D1495796-G98E-585C-959E-1269CD6766B8\&quot;           }         ]       }     ]   }   &#x60;&#x60;&#x60;    **Updated JSON***    &#x60;&#x60;&#x60;   {     \&quot;employeeId\&quot;: \&quot;00u12345678\&quot;,     \&quot;extPatientId\&quot;: 12345,     \&quot;external_guid\&quot;: \&quot;D1495796-G98E-585C-959E-1269CD6766B8\&quot;   }   &#x60;&#x60;&#x60;  &lt;/details&gt;  &lt;details&gt; &lt;summary&gt;Modify members within existing JSON objects and arrays&lt;/summary&gt;    Use the &#x60;replace&#x60; operation to modify members within JSON objects and elements within arrays. For example, you have a JSON object in a claim called &#x60;employee_profile&#x60;, and you want to update the email address of the employee. The existing target JSON object is updated by specifying the claim in the path, followed by the name of the object member that you want to modify.    **Existing target JSON object**    &#x60;&#x60;&#x60;   {     \&quot;employee_profile\&quot;: {       \&quot;employee_id\&quot;:\&quot;1234\&quot;,       \&quot;name\&quot;:\&quot;Anna\&quot;,       \&quot;email\&quot;:\&quot;anna.v@company.com\&quot;       }   }   &#x60;&#x60;&#x60;    **Operation**    &#x60;&#x60;&#x60;   {     \&quot;commands\&quot;: [       {         \&quot;type\&quot;: \&quot;com.okta.identity.patch\&quot;,         \&quot;value\&quot;: [           {             \&quot;op\&quot;: \&quot;replace\&quot;,             \&quot;path\&quot;: \&quot;/claims/employee_profile/email\&quot;,             \&quot;value\&quot;: \&quot;anna@company.com\&quot;           }         ]       }     ]   }   &#x60;&#x60;&#x60;    **Updated JSON**    &#x60;&#x60;&#x60;   {     \&quot;employee_profile\&quot;: {       \&quot;employee_id\&quot;:\&quot;1234\&quot;,       \&quot;name\&quot;:\&quot;Anna\&quot;,       \&quot;email\&quot;:\&quot;anna@company.com\&quot;       }   }   &#x60;&#x60;&#x60;    &gt; **Note:** If you attempt to modify a member within a JSON object that doesn&#39;t exist or use an invalid operation, the entire PATCH fails and errors are logged in the token hooks events.    Similarly, you can replace elements in an array by specifying the array name and the valid index of the element that you want to replace in the path. &lt;/details&gt;  &lt;details&gt; &lt;summary&gt;Modify token lifetimes&lt;/summary&gt;   You can modify how long the access and ID tokens are valid by specifying the &#x60;lifetime&#x60; in seconds. The &#x60;lifetime&#x60; value must be a minimum of five minutes (300 seconds) and a maximum of 24 hours (86,400 seconds).    **Operation**    &#x60;&#x60;&#x60;   {     \&quot;commands\&quot;: [       {         \&quot;type\&quot;: \&quot;com.okta.identity.patch\&quot;,         \&quot;value\&quot;: [           {             \&quot;op\&quot;: \&quot;replace\&quot;,             \&quot;path\&quot;: \&quot;/token/lifetime/expiration\&quot;,             \&quot;value\&quot;: 36000           }         ]       },       {         \&quot;type\&quot;: \&quot;com.okta.access.patch\&quot;,         \&quot;value\&quot;: [           {             \&quot;op\&quot;: \&quot;replace\&quot;,             \&quot;path\&quot;: \&quot;/token/lifetime/expiration\&quot;,             \&quot;value\&quot;: 36000           }         ]       }     ]   }   &#x60;&#x60;&#x60;  &lt;/details&gt;  #### &#x60;op: remove&#x60; notes  &lt;details&gt; &lt;summary&gt;Remove a claim&lt;/summary&gt;    You can remove existing custom claims or OIDC standard profile claims, such as &#x60;birthdate&#x60; or &#x60;locale&#x60;. You can&#39;t, however, remove any system-specific claims, such as &#x60;iss&#x60; or &#x60;ver&#x60;. You also can&#39;t remove a claim that isn&#39;t currently part of the token in the request payload. If you attempt to remove a system-specific claim or use an invalid operation, the entire PATCH fails and errors are logged in the token hooks events.    See [Access Tokens Scopes and Claims](/openapi/okta-oauth/guides/overview/#access-token-scopes-and-claims) for the list of access token-reserved claims that you can&#39;t modify.    See [ID Token Claims](/openapi/okta-oauth/guides/overview/#id-token-claims) for a list of ID token-reserved claims that you can&#39;t modify.    **Operation**    &#x60;&#x60;&#x60;   {     \&quot;commands\&quot;: [       {         \&quot;type\&quot;: \&quot;com.okta.identity.patch\&quot;,         \&quot;value\&quot;: [           {             \&quot;op\&quot;: \&quot;remove\&quot;,             \&quot;path\&quot;: \&quot;/claims/birthdate\&quot;,             \&quot;value\&quot;: null           }         ]       },       {         \&quot;type\&quot;: \&quot;com.okta.access.patch\&quot;,         \&quot;value\&quot;: [           {             \&quot;op\&quot;: \&quot;remove\&quot;,             \&quot;path\&quot;: \&quot;/claims/external_guid\&quot;           }         ]       }     ]   }   &#x60;&#x60;&#x60;    &gt; **Note:** The &#x60;value&#x60; property for the &#x60;remove&#x60; operation isn&#39;t required. If you provide it in the response, it should be set to &#x60;null&#x60;. Providing any other value fails the entire PATCH response.  &lt;/details&gt;  &lt;details&gt; &lt;summary&gt;Remove members from existing arrays&lt;/summary&gt;    Use the &#x60;remove&#x60; operation to remove members from existing arrays. For example, you have an array that contains the user&#39;s preferred airports, and you want to remove an airport from the array. The existing target JSON object is updated by specifying the array name followed by the index of the element that you want to remove. You don&#39;t need to specify a value for the remove operation, but you can specify &#x60;null&#x60; as the value if you want.    **Existing target JSON object**    &#x60;&#x60;&#x60;   {     \&quot;preferred_airports\&quot;: [         \&quot;sjc\&quot;,         \&quot;lax\&quot;,         \&quot;sfo\&quot;,         \&quot;oak\&quot;       ]   }   &#x60;&#x60;&#x60;    **Operation**    &#x60;&#x60;&#x60;   {     \&quot;commands\&quot;: [       {         \&quot;type\&quot;: \&quot;com.okta.identity.patch\&quot;,         \&quot;value\&quot;: [           {             \&quot;op\&quot;: \&quot;remove\&quot;,             \&quot;path\&quot;: \&quot;/claims/preferred_airports/1\&quot;           }         ]       }     ]   }   &#x60;&#x60;&#x60;    **Updated JSON**    &#x60;&#x60;&#x60;   {     \&quot;preferred_airports\&quot;: [         \&quot;sjc\&quot;,         \&quot;sfo\&quot;,         \&quot;oak\&quot;       ]   }   &#x60;&#x60;&#x60;  &lt;/details&gt;  &lt;details&gt; &lt;summary&gt;Remove members from existing JSON objects&lt;/summary&gt;   Use the &#x60;remove&#x60; operation to remove members from existing JSON objects. Do this by specifying the JSON object in the path, followed by the claim member that you would like to remove. For example, you have an &#x60;employee_profile&#x60; claim, and you want to remove &#x60;email&#x60; from it.  **Existing target JSON object**  &#x60;&#x60;&#x60; {   \&quot;employee_profile\&quot;: {     \&quot;employee_id\&quot;:\&quot;1234\&quot;,     \&quot;name\&quot;:\&quot;Anna\&quot;,     \&quot;email\&quot;:\&quot;anna.v@company.com\&quot;     } } &#x60;&#x60;&#x60;  **Operation**  &#x60;&#x60;&#x60; {   \&quot;commands\&quot;: [     {       \&quot;type\&quot;: \&quot;com.okta.identity.patch\&quot;,       \&quot;value\&quot;: [         {           \&quot;op\&quot;: \&quot;remove\&quot;,           \&quot;path\&quot;: \&quot;/claims/employee_profile/email\&quot;         }       ]     }   ] } &#x60;&#x60;&#x60;  **Updated JSON** &#x60;&#x60;&#x60; {   \&quot;employee_profile\&quot;: {     \&quot;employee_id\&quot;:\&quot;1234\&quot;,     \&quot;name\&quot;:\&quot;Anna\&quot;,     } } &#x60;&#x60;&#x60;  &lt;/details&gt; | [optional] 
**Path** | Pointer to **string** | Location within the token to apply the operation, specified as a slash-delimited path. When you add, replace, or remove a claim, this path always begins with &#x60;/claims/&#x60; and is followed by the name of the new claim that you&#39;re adding. When you replace a token lifetime, the path should always be &#x60;/token/lifetime/expiration&#x60;. | [optional] 
**Value** | Pointer to [**TokenHookResponseCommandsInnerValueInnerValue**](TokenHookResponseCommandsInnerValueInnerValue.md) |  | [optional] 

## Methods

### NewTokenHookResponseCommandsInnerValueInner

`func NewTokenHookResponseCommandsInnerValueInner() *TokenHookResponseCommandsInnerValueInner`

NewTokenHookResponseCommandsInnerValueInner instantiates a new TokenHookResponseCommandsInnerValueInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenHookResponseCommandsInnerValueInnerWithDefaults

`func NewTokenHookResponseCommandsInnerValueInnerWithDefaults() *TokenHookResponseCommandsInnerValueInner`

NewTokenHookResponseCommandsInnerValueInnerWithDefaults instantiates a new TokenHookResponseCommandsInnerValueInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOp

`func (o *TokenHookResponseCommandsInnerValueInner) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *TokenHookResponseCommandsInnerValueInner) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *TokenHookResponseCommandsInnerValueInner) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *TokenHookResponseCommandsInnerValueInner) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetPath

`func (o *TokenHookResponseCommandsInnerValueInner) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *TokenHookResponseCommandsInnerValueInner) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *TokenHookResponseCommandsInnerValueInner) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *TokenHookResponseCommandsInnerValueInner) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetValue

`func (o *TokenHookResponseCommandsInnerValueInner) GetValue() TokenHookResponseCommandsInnerValueInnerValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TokenHookResponseCommandsInnerValueInner) GetValueOk() (*TokenHookResponseCommandsInnerValueInnerValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TokenHookResponseCommandsInnerValueInner) SetValue(v TokenHookResponseCommandsInnerValueInnerValue)`

SetValue sets Value field to given value.

### HasValue

`func (o *TokenHookResponseCommandsInnerValueInner) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


