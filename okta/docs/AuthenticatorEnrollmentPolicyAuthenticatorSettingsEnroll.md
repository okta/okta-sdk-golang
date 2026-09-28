# AuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Self** | Pointer to **string** | Requirements for the user-initiated enrollment | [optional] [default to "NOT_ALLOWED"]
**AutoEnroll** | Pointer to **NullableBool** | &lt;x-lifecycle class&#x3D;\&quot;oie\&quot;&gt;&lt;/x-lifecycle&gt; Controls whether the email authenticator is automatically enrolled for users during sign-up, user account creation, or if the user&#39;s profile email address is changed or updated.    * &#x60;true&#x60; or &#x60;null&#x60;: The email authenticator is auto-enrolled. This is the default behavior for the email authenticator in Identity Engine.   * &#x60;false&#x60;: The email authenticator auto-enrollment is skipped. Users see email as an optional authenticator during enrollment and can choose to skip it.    &gt; **Note:** &#x60;autoEnroll&#x60; can only be set to &#x60;false&#x60; when [&#x60;self&#x60;](/openapi/okta-management/management/tags/policy/other/createpolicy#other/createpolicy/t&#x3D;request&amp;path&#x3D;&amp;d&#x3D;1/settings/authenticators/enroll/self) is &#x60;OPTIONAL&#x60; or &#x60;NOT_ALLOWED&#x60;. This property is only available if the Email auto-enrollment and recovery control [self-service EA](/openapi/okta-management/guides/release-lifecycle/#early-access-ea) feature is enabled.   | [optional] 
**GracePeriod** | Pointer to [**EnrollmentPolicyAuthenticatorGracePeriod**](EnrollmentPolicyAuthenticatorGracePeriod.md) |  | [optional] 
**Promotion** | Pointer to [**EnrollmentPolicyAuthenticatorPromotion**](EnrollmentPolicyAuthenticatorPromotion.md) |  | [optional] 

## Methods

### NewAuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll

`func NewAuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll() *AuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll`

NewAuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll instantiates a new AuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticatorEnrollmentPolicyAuthenticatorSettingsEnrollWithDefaults

`func NewAuthenticatorEnrollmentPolicyAuthenticatorSettingsEnrollWithDefaults() *AuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll`

NewAuthenticatorEnrollmentPolicyAuthenticatorSettingsEnrollWithDefaults instantiates a new AuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelf

`func (o *AuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *AuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *AuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *AuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetAutoEnroll

`func (o *AuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll) GetAutoEnroll() bool`

GetAutoEnroll returns the AutoEnroll field if non-nil, zero value otherwise.

### GetAutoEnrollOk

`func (o *AuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll) GetAutoEnrollOk() (*bool, bool)`

GetAutoEnrollOk returns a tuple with the AutoEnroll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoEnroll

`func (o *AuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll) SetAutoEnroll(v bool)`

SetAutoEnroll sets AutoEnroll field to given value.

### HasAutoEnroll

`func (o *AuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll) HasAutoEnroll() bool`

HasAutoEnroll returns a boolean if a field has been set.

### SetAutoEnrollNil

`func (o *AuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll) SetAutoEnrollNil(b bool)`

 SetAutoEnrollNil sets the value for AutoEnroll to be an explicit nil

### UnsetAutoEnroll
`func (o *AuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll) UnsetAutoEnroll()`

UnsetAutoEnroll ensures that no value is present for AutoEnroll, not even an explicit nil
### GetGracePeriod

`func (o *AuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll) GetGracePeriod() EnrollmentPolicyAuthenticatorGracePeriod`

GetGracePeriod returns the GracePeriod field if non-nil, zero value otherwise.

### GetGracePeriodOk

`func (o *AuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll) GetGracePeriodOk() (*EnrollmentPolicyAuthenticatorGracePeriod, bool)`

GetGracePeriodOk returns a tuple with the GracePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGracePeriod

`func (o *AuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll) SetGracePeriod(v EnrollmentPolicyAuthenticatorGracePeriod)`

SetGracePeriod sets GracePeriod field to given value.

### HasGracePeriod

`func (o *AuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll) HasGracePeriod() bool`

HasGracePeriod returns a boolean if a field has been set.

### GetPromotion

`func (o *AuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll) GetPromotion() EnrollmentPolicyAuthenticatorPromotion`

GetPromotion returns the Promotion field if non-nil, zero value otherwise.

### GetPromotionOk

`func (o *AuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll) GetPromotionOk() (*EnrollmentPolicyAuthenticatorPromotion, bool)`

GetPromotionOk returns a tuple with the Promotion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromotion

`func (o *AuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll) SetPromotion(v EnrollmentPolicyAuthenticatorPromotion)`

SetPromotion sets Promotion field to given value.

### HasPromotion

`func (o *AuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll) HasPromotion() bool`

HasPromotion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


