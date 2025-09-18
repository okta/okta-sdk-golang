# AuthenticatorEnrollmentPolicySettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authenticators** | Pointer to [**[]AuthenticatorEnrollmentPolicyAuthenticatorSettings**](AuthenticatorEnrollmentPolicyAuthenticatorSettings.md) | List of authenticator policy settings  &lt;x-lifecycle class&#x3D;\&quot;oie\&quot;&gt;&lt;/x-lifecycle&gt; For orgs with the Authenticator enrollment policy feature enabled, the new default authenticator enrollment policy created by Okta contains the &#x60;authenticators&#x60; property in the policy settings. Existing default authenticator enrollment policies from a migrated Classic Engine org remain unchanged. The policies still use the &#x60;factors&#x60; property in their settings. The &#x60;authenticators&#x60; parameter allows you to configure all available authenticators, including authentication and recovery. The &#x60;factors&#x60; parameter only allows you to configure multifactor authentication.  | [optional] 
**Type** | Pointer to **string** | Type of policy configuration object  &lt;x-lifecycle class&#x3D;\&quot;oie\&quot;&gt;&lt;/x-lifecycle&gt; The &#x60;type&#x60; property in the policy &#x60;settings&#x60; is only applicable to the authenticator enrollment policy available in Identity Engine.  | [optional] [default to "FACTORS"]

## Methods

### NewAuthenticatorEnrollmentPolicySettings

`func NewAuthenticatorEnrollmentPolicySettings() *AuthenticatorEnrollmentPolicySettings`

NewAuthenticatorEnrollmentPolicySettings instantiates a new AuthenticatorEnrollmentPolicySettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticatorEnrollmentPolicySettingsWithDefaults

`func NewAuthenticatorEnrollmentPolicySettingsWithDefaults() *AuthenticatorEnrollmentPolicySettings`

NewAuthenticatorEnrollmentPolicySettingsWithDefaults instantiates a new AuthenticatorEnrollmentPolicySettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticators

`func (o *AuthenticatorEnrollmentPolicySettings) GetAuthenticators() []AuthenticatorEnrollmentPolicyAuthenticatorSettings`

GetAuthenticators returns the Authenticators field if non-nil, zero value otherwise.

### GetAuthenticatorsOk

`func (o *AuthenticatorEnrollmentPolicySettings) GetAuthenticatorsOk() (*[]AuthenticatorEnrollmentPolicyAuthenticatorSettings, bool)`

GetAuthenticatorsOk returns a tuple with the Authenticators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticators

`func (o *AuthenticatorEnrollmentPolicySettings) SetAuthenticators(v []AuthenticatorEnrollmentPolicyAuthenticatorSettings)`

SetAuthenticators sets Authenticators field to given value.

### HasAuthenticators

`func (o *AuthenticatorEnrollmentPolicySettings) HasAuthenticators() bool`

HasAuthenticators returns a boolean if a field has been set.

### GetType

`func (o *AuthenticatorEnrollmentPolicySettings) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuthenticatorEnrollmentPolicySettings) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuthenticatorEnrollmentPolicySettings) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AuthenticatorEnrollmentPolicySettings) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


