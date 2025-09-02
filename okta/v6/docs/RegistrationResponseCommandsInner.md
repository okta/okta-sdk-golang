# RegistrationResponseCommandsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The location where you specify the command. To set attributes in the user&#39;s Okta profile, supply a &#x60;type&#x60; property set to &#x60;com.okta.user.profile.update&#x60;, together with a &#x60;value&#x60; property set to a list of key-value pairs corresponding to the Okta user profile attributes you want to set. The attributes must already exist in your user profile schema.  To explicitly allow or deny registration to the user, supply a type property set to &#x60;com.okta.action.update&#x60;, together with a value property set to &#x60;{\&quot;registration\&quot;: \&quot;ALLOW\&quot;}&#x60; or &#x60;{\&quot;registration\&quot;: \&quot;DENY\&quot;}&#x60;. The default is to allow registration.  In Okta Identity Engine, to set attributes in the user&#39;s profile, supply a &#x60;type&#x60; property set to &#x60;com.okta.user.progressive.profile.update&#x60;, together with a &#x60;value&#x60; property set to a list of key-value pairs corresponding to the Progressive Enrollment attributes that you want to set. See [Registration inline hook - Send response](https://developer.okta.com/docs/guides/registration-inline-hook/nodejs/main/#send-response).  Commands are applied in the order that they appear in the array. Within a single &#x60;com.okta.user.profile.update&#x60; or &#x60;com.okta.user.progressive.profile.update command&#x60;, attributes are updated in the order that they appear in the &#x60;value&#x60; object.  You can never use a command to update the user&#39;s password, but you are allowed to set the values of attributes other than password that are designated sensitive in your Okta user schema. However, the values of those sensitive attributes, if included as fields in the Profile Enrollment form, aren&#39;t included in the &#x60;data.userProfile&#x60; object sent to your external service by Okta. See [data.userProfile](/openapi/okta-management/management/tag/InlineHook/#tag/InlineHook/operation/create-registration-hook!path&#x3D;0/data/userProfile&amp;t&#x3D;request). | [optional] 
**Value** | Pointer to **map[string]interface{}** | The &#x60;value&#x60; object is the parameter to pass to the command.  For &#x60;com.okta.user.profile.update&#x60; commands, &#x60;value&#x60; should be an object containing one or more name-value pairs for the attributes you wish to update.  For &#x60;com.okta.action.update&#x60; commands, the value should be an object containing the attribute &#x60;action&#x60; set to a value of either &#x60;ALLOW&#x60; or &#x60;DENY&#x60;, indicating whether the registration should be permitted or not.  Registrations are allowed by default, so setting a value of &#x60;ALLOW&#x60; for the action field is valid but superfluous. | [optional] 

## Methods

### NewRegistrationResponseCommandsInner

`func NewRegistrationResponseCommandsInner() *RegistrationResponseCommandsInner`

NewRegistrationResponseCommandsInner instantiates a new RegistrationResponseCommandsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistrationResponseCommandsInnerWithDefaults

`func NewRegistrationResponseCommandsInnerWithDefaults() *RegistrationResponseCommandsInner`

NewRegistrationResponseCommandsInnerWithDefaults instantiates a new RegistrationResponseCommandsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RegistrationResponseCommandsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RegistrationResponseCommandsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RegistrationResponseCommandsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RegistrationResponseCommandsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *RegistrationResponseCommandsInner) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RegistrationResponseCommandsInner) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *RegistrationResponseCommandsInner) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *RegistrationResponseCommandsInner) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


