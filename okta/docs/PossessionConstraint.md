# PossessionConstraint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationMethods** | Pointer to [**[]AuthenticationMethodObject**](AuthenticationMethodObject.md) | This property specifies the precise authenticator and method for authentication. &lt;x-lifecycle class&#x3D;\&quot;oie\&quot;&gt;&lt;/x-lifecycle&gt; | [optional] 
**ExcludedAuthenticationMethods** | Pointer to [**[]AuthenticationMethodObject**](AuthenticationMethodObject.md) | This property specifies the precise authenticator and method to exclude from authentication. &lt;x-lifecycle class&#x3D;\&quot;oie\&quot;&gt;&lt;/x-lifecycle&gt; | [optional] 
**Methods** | Pointer to **[]string** | The Authenticator methods that are permitted | [optional] 
**ReauthenticateIn** | Pointer to **string** | The duration after which the user must re-authenticate regardless of user activity. This re-authentication interval overrides the Verification Method object&#39;s &#x60;reauthenticateIn&#x60; interval. The supported values use ISO 8601 period format for recurring time intervals (for example, &#x60;PT1H&#x60;). | [optional] 
**Required** | Pointer to **bool** | This property indicates whether the knowledge or possession factor is required by the assurance. It&#39;s optional in the request, but is always returned in the response. By default, this field is &#x60;true&#x60;. If the knowledge or possession constraint has values for &#x60;excludedAuthenticationMethods&#x60; the &#x60;required&#x60; value is false. &lt;x-lifecycle class&#x3D;\&quot;oie\&quot;&gt;&lt;/x-lifecycle&gt; | [optional] 
**Types** | Pointer to **[]string** | The Authenticator types that are permitted | [optional] 
**DeviceBound** | Pointer to **string** | Indicates if device-bound Factors are required. This property is only set for &#x60;POSSESSION&#x60; constraints. | [optional] [default to "OPTIONAL"]
**HardwareProtection** | Pointer to **string** | Indicates if any secrets or private keys used during authentication must be hardware protected and not exportable. This property is only set for &#x60;POSSESSION&#x60; constraints. | [optional] [default to "OPTIONAL"]
**PhishingResistant** | Pointer to **string** | Indicates if phishing-resistant Factors are required. This property is only set for &#x60;POSSESSION&#x60; constraints. | [optional] [default to "OPTIONAL"]
**UserPresence** | Pointer to **string** | Indicates if the user needs to approve an Okta Verify prompt or provide biometrics (meets NIST AAL2 requirements). This property is only set for &#x60;POSSESSION&#x60; constraints. | [optional] [default to "REQUIRED"]
**UserVerification** | Pointer to **string** | Indicates the user interaction requirement (PIN or biometrics) to ensure verification of a possession factor | [optional] [default to "OPTIONAL"]

## Methods

### NewPossessionConstraint

`func NewPossessionConstraint() *PossessionConstraint`

NewPossessionConstraint instantiates a new PossessionConstraint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPossessionConstraintWithDefaults

`func NewPossessionConstraintWithDefaults() *PossessionConstraint`

NewPossessionConstraintWithDefaults instantiates a new PossessionConstraint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationMethods

`func (o *PossessionConstraint) GetAuthenticationMethods() []AuthenticationMethodObject`

GetAuthenticationMethods returns the AuthenticationMethods field if non-nil, zero value otherwise.

### GetAuthenticationMethodsOk

`func (o *PossessionConstraint) GetAuthenticationMethodsOk() (*[]AuthenticationMethodObject, bool)`

GetAuthenticationMethodsOk returns a tuple with the AuthenticationMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationMethods

`func (o *PossessionConstraint) SetAuthenticationMethods(v []AuthenticationMethodObject)`

SetAuthenticationMethods sets AuthenticationMethods field to given value.

### HasAuthenticationMethods

`func (o *PossessionConstraint) HasAuthenticationMethods() bool`

HasAuthenticationMethods returns a boolean if a field has been set.

### GetExcludedAuthenticationMethods

`func (o *PossessionConstraint) GetExcludedAuthenticationMethods() []AuthenticationMethodObject`

GetExcludedAuthenticationMethods returns the ExcludedAuthenticationMethods field if non-nil, zero value otherwise.

### GetExcludedAuthenticationMethodsOk

`func (o *PossessionConstraint) GetExcludedAuthenticationMethodsOk() (*[]AuthenticationMethodObject, bool)`

GetExcludedAuthenticationMethodsOk returns a tuple with the ExcludedAuthenticationMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedAuthenticationMethods

`func (o *PossessionConstraint) SetExcludedAuthenticationMethods(v []AuthenticationMethodObject)`

SetExcludedAuthenticationMethods sets ExcludedAuthenticationMethods field to given value.

### HasExcludedAuthenticationMethods

`func (o *PossessionConstraint) HasExcludedAuthenticationMethods() bool`

HasExcludedAuthenticationMethods returns a boolean if a field has been set.

### GetMethods

`func (o *PossessionConstraint) GetMethods() []string`

GetMethods returns the Methods field if non-nil, zero value otherwise.

### GetMethodsOk

`func (o *PossessionConstraint) GetMethodsOk() (*[]string, bool)`

GetMethodsOk returns a tuple with the Methods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethods

`func (o *PossessionConstraint) SetMethods(v []string)`

SetMethods sets Methods field to given value.

### HasMethods

`func (o *PossessionConstraint) HasMethods() bool`

HasMethods returns a boolean if a field has been set.

### GetReauthenticateIn

`func (o *PossessionConstraint) GetReauthenticateIn() string`

GetReauthenticateIn returns the ReauthenticateIn field if non-nil, zero value otherwise.

### GetReauthenticateInOk

`func (o *PossessionConstraint) GetReauthenticateInOk() (*string, bool)`

GetReauthenticateInOk returns a tuple with the ReauthenticateIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReauthenticateIn

`func (o *PossessionConstraint) SetReauthenticateIn(v string)`

SetReauthenticateIn sets ReauthenticateIn field to given value.

### HasReauthenticateIn

`func (o *PossessionConstraint) HasReauthenticateIn() bool`

HasReauthenticateIn returns a boolean if a field has been set.

### GetRequired

`func (o *PossessionConstraint) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *PossessionConstraint) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *PossessionConstraint) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *PossessionConstraint) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetTypes

`func (o *PossessionConstraint) GetTypes() []string`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *PossessionConstraint) GetTypesOk() (*[]string, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *PossessionConstraint) SetTypes(v []string)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *PossessionConstraint) HasTypes() bool`

HasTypes returns a boolean if a field has been set.

### GetDeviceBound

`func (o *PossessionConstraint) GetDeviceBound() string`

GetDeviceBound returns the DeviceBound field if non-nil, zero value otherwise.

### GetDeviceBoundOk

`func (o *PossessionConstraint) GetDeviceBoundOk() (*string, bool)`

GetDeviceBoundOk returns a tuple with the DeviceBound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceBound

`func (o *PossessionConstraint) SetDeviceBound(v string)`

SetDeviceBound sets DeviceBound field to given value.

### HasDeviceBound

`func (o *PossessionConstraint) HasDeviceBound() bool`

HasDeviceBound returns a boolean if a field has been set.

### GetHardwareProtection

`func (o *PossessionConstraint) GetHardwareProtection() string`

GetHardwareProtection returns the HardwareProtection field if non-nil, zero value otherwise.

### GetHardwareProtectionOk

`func (o *PossessionConstraint) GetHardwareProtectionOk() (*string, bool)`

GetHardwareProtectionOk returns a tuple with the HardwareProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareProtection

`func (o *PossessionConstraint) SetHardwareProtection(v string)`

SetHardwareProtection sets HardwareProtection field to given value.

### HasHardwareProtection

`func (o *PossessionConstraint) HasHardwareProtection() bool`

HasHardwareProtection returns a boolean if a field has been set.

### GetPhishingResistant

`func (o *PossessionConstraint) GetPhishingResistant() string`

GetPhishingResistant returns the PhishingResistant field if non-nil, zero value otherwise.

### GetPhishingResistantOk

`func (o *PossessionConstraint) GetPhishingResistantOk() (*string, bool)`

GetPhishingResistantOk returns a tuple with the PhishingResistant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhishingResistant

`func (o *PossessionConstraint) SetPhishingResistant(v string)`

SetPhishingResistant sets PhishingResistant field to given value.

### HasPhishingResistant

`func (o *PossessionConstraint) HasPhishingResistant() bool`

HasPhishingResistant returns a boolean if a field has been set.

### GetUserPresence

`func (o *PossessionConstraint) GetUserPresence() string`

GetUserPresence returns the UserPresence field if non-nil, zero value otherwise.

### GetUserPresenceOk

`func (o *PossessionConstraint) GetUserPresenceOk() (*string, bool)`

GetUserPresenceOk returns a tuple with the UserPresence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPresence

`func (o *PossessionConstraint) SetUserPresence(v string)`

SetUserPresence sets UserPresence field to given value.

### HasUserPresence

`func (o *PossessionConstraint) HasUserPresence() bool`

HasUserPresence returns a boolean if a field has been set.

### GetUserVerification

`func (o *PossessionConstraint) GetUserVerification() string`

GetUserVerification returns the UserVerification field if non-nil, zero value otherwise.

### GetUserVerificationOk

`func (o *PossessionConstraint) GetUserVerificationOk() (*string, bool)`

GetUserVerificationOk returns a tuple with the UserVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserVerification

`func (o *PossessionConstraint) SetUserVerification(v string)`

SetUserVerification sets UserVerification field to given value.

### HasUserVerification

`func (o *PossessionConstraint) HasUserVerification() bool`

HasUserVerification returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


