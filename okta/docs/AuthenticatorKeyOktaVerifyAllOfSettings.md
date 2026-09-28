# AuthenticatorKeyOktaVerifyAllOfSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelBinding** | Pointer to [**ChannelBinding**](ChannelBinding.md) |  | [optional] 
**Compliance** | Pointer to [**Compliance**](Compliance.md) |  | [optional] 
**UserVerification** | Pointer to **string** | User verification settings. Possible values &#x60;DISCOURAGED&#x60; (the authenticator isn&#39;t asked to perform user verification, but may do so at its discretion), &#x60;PREFERRED&#x60; (the client uses an authenticator capable of user verification if possible), or &#x60;REQUIRED&#x60; (the client uses only an authenticator capable of user verification) | [optional] 
**EnrollmentSecurityLevel** | Pointer to **string** | The security level required for device enrollment. &#x60;HIGH&#x60; restricts enrollment to hardware-bound keys only. &#x60;ANY&#x60; allows all enrollment channels.  &gt; &lt;x-lifecycle class&#x3D;\&quot;ea\&quot;&gt;&lt;/x-lifecycle&gt; **Note:** When the Flexible Okta Verify authenticator configuration feature is enabled, this is a common setting shared across all per-method authenticators (&#x60;okta_verify_totp&#x60;, &#x60;okta_verify_push&#x60;, &#x60;okta_verify_fastpass&#x60;). Updating this value on any one authenticator applies the change to all three. | [optional] 
**UserVerificationMethods** | Pointer to **[]string** |  | [optional] 
**AppInstanceId** | Pointer to **string** | The application instance ID | [optional] 

## Methods

### NewAuthenticatorKeyOktaVerifyAllOfSettings

`func NewAuthenticatorKeyOktaVerifyAllOfSettings() *AuthenticatorKeyOktaVerifyAllOfSettings`

NewAuthenticatorKeyOktaVerifyAllOfSettings instantiates a new AuthenticatorKeyOktaVerifyAllOfSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticatorKeyOktaVerifyAllOfSettingsWithDefaults

`func NewAuthenticatorKeyOktaVerifyAllOfSettingsWithDefaults() *AuthenticatorKeyOktaVerifyAllOfSettings`

NewAuthenticatorKeyOktaVerifyAllOfSettingsWithDefaults instantiates a new AuthenticatorKeyOktaVerifyAllOfSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelBinding

`func (o *AuthenticatorKeyOktaVerifyAllOfSettings) GetChannelBinding() ChannelBinding`

GetChannelBinding returns the ChannelBinding field if non-nil, zero value otherwise.

### GetChannelBindingOk

`func (o *AuthenticatorKeyOktaVerifyAllOfSettings) GetChannelBindingOk() (*ChannelBinding, bool)`

GetChannelBindingOk returns a tuple with the ChannelBinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelBinding

`func (o *AuthenticatorKeyOktaVerifyAllOfSettings) SetChannelBinding(v ChannelBinding)`

SetChannelBinding sets ChannelBinding field to given value.

### HasChannelBinding

`func (o *AuthenticatorKeyOktaVerifyAllOfSettings) HasChannelBinding() bool`

HasChannelBinding returns a boolean if a field has been set.

### GetCompliance

`func (o *AuthenticatorKeyOktaVerifyAllOfSettings) GetCompliance() Compliance`

GetCompliance returns the Compliance field if non-nil, zero value otherwise.

### GetComplianceOk

`func (o *AuthenticatorKeyOktaVerifyAllOfSettings) GetComplianceOk() (*Compliance, bool)`

GetComplianceOk returns a tuple with the Compliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompliance

`func (o *AuthenticatorKeyOktaVerifyAllOfSettings) SetCompliance(v Compliance)`

SetCompliance sets Compliance field to given value.

### HasCompliance

`func (o *AuthenticatorKeyOktaVerifyAllOfSettings) HasCompliance() bool`

HasCompliance returns a boolean if a field has been set.

### GetUserVerification

`func (o *AuthenticatorKeyOktaVerifyAllOfSettings) GetUserVerification() string`

GetUserVerification returns the UserVerification field if non-nil, zero value otherwise.

### GetUserVerificationOk

`func (o *AuthenticatorKeyOktaVerifyAllOfSettings) GetUserVerificationOk() (*string, bool)`

GetUserVerificationOk returns a tuple with the UserVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserVerification

`func (o *AuthenticatorKeyOktaVerifyAllOfSettings) SetUserVerification(v string)`

SetUserVerification sets UserVerification field to given value.

### HasUserVerification

`func (o *AuthenticatorKeyOktaVerifyAllOfSettings) HasUserVerification() bool`

HasUserVerification returns a boolean if a field has been set.

### GetEnrollmentSecurityLevel

`func (o *AuthenticatorKeyOktaVerifyAllOfSettings) GetEnrollmentSecurityLevel() string`

GetEnrollmentSecurityLevel returns the EnrollmentSecurityLevel field if non-nil, zero value otherwise.

### GetEnrollmentSecurityLevelOk

`func (o *AuthenticatorKeyOktaVerifyAllOfSettings) GetEnrollmentSecurityLevelOk() (*string, bool)`

GetEnrollmentSecurityLevelOk returns a tuple with the EnrollmentSecurityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentSecurityLevel

`func (o *AuthenticatorKeyOktaVerifyAllOfSettings) SetEnrollmentSecurityLevel(v string)`

SetEnrollmentSecurityLevel sets EnrollmentSecurityLevel field to given value.

### HasEnrollmentSecurityLevel

`func (o *AuthenticatorKeyOktaVerifyAllOfSettings) HasEnrollmentSecurityLevel() bool`

HasEnrollmentSecurityLevel returns a boolean if a field has been set.

### GetUserVerificationMethods

`func (o *AuthenticatorKeyOktaVerifyAllOfSettings) GetUserVerificationMethods() []string`

GetUserVerificationMethods returns the UserVerificationMethods field if non-nil, zero value otherwise.

### GetUserVerificationMethodsOk

`func (o *AuthenticatorKeyOktaVerifyAllOfSettings) GetUserVerificationMethodsOk() (*[]string, bool)`

GetUserVerificationMethodsOk returns a tuple with the UserVerificationMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserVerificationMethods

`func (o *AuthenticatorKeyOktaVerifyAllOfSettings) SetUserVerificationMethods(v []string)`

SetUserVerificationMethods sets UserVerificationMethods field to given value.

### HasUserVerificationMethods

`func (o *AuthenticatorKeyOktaVerifyAllOfSettings) HasUserVerificationMethods() bool`

HasUserVerificationMethods returns a boolean if a field has been set.

### GetAppInstanceId

`func (o *AuthenticatorKeyOktaVerifyAllOfSettings) GetAppInstanceId() string`

GetAppInstanceId returns the AppInstanceId field if non-nil, zero value otherwise.

### GetAppInstanceIdOk

`func (o *AuthenticatorKeyOktaVerifyAllOfSettings) GetAppInstanceIdOk() (*string, bool)`

GetAppInstanceIdOk returns a tuple with the AppInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppInstanceId

`func (o *AuthenticatorKeyOktaVerifyAllOfSettings) SetAppInstanceId(v string)`

SetAppInstanceId sets AppInstanceId field to given value.

### HasAppInstanceId

`func (o *AuthenticatorKeyOktaVerifyAllOfSettings) HasAppInstanceId() bool`

HasAppInstanceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


