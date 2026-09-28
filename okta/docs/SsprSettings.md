# SsprSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowRecoveryEmailWithoutEnrollment** | Pointer to **NullableBool** | &lt;x-lifecycle class&#x3D;\&quot;oie\&quot;&gt;&lt;/x-lifecycle&gt; Controls whether a user can receive a password recovery email at their profile email address even if they haven&#39;t enrolled email as an authenticator.  * &#x60;true&#x60;: Email appears as a recovery option during password reset and account unlock for users without an enrolled email authenticator. The recovery OTP is sent to the email address on the user&#39;s profile. * &#x60;false&#x60; or &#x60;null&#x60;: Email recovery requires an enrolled email authenticator. This is the default behavior for Identity Engine orgs.  For orgs that have migrated from Classic Engine to Identity Engine, &#x60;allowRecoveryEmailWithoutEnrollment&#x60; is set to &#x60;true&#x60;. It&#39;s set to &#x60;true&#x60; by the migration action to preserve Classic Engine behavior where recovery emails are always sent to the profile email address regardless of the user&#39;s enrollment state.  &gt; **Note:** This property is only available if the Email auto-enrollment and recovery control [self-service EA](/openapi/okta-management/guides/release-lifecycle/#early-access-ea) feature is enabled. | [optional] 

## Methods

### NewSsprSettings

`func NewSsprSettings() *SsprSettings`

NewSsprSettings instantiates a new SsprSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSsprSettingsWithDefaults

`func NewSsprSettingsWithDefaults() *SsprSettings`

NewSsprSettingsWithDefaults instantiates a new SsprSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowRecoveryEmailWithoutEnrollment

`func (o *SsprSettings) GetAllowRecoveryEmailWithoutEnrollment() bool`

GetAllowRecoveryEmailWithoutEnrollment returns the AllowRecoveryEmailWithoutEnrollment field if non-nil, zero value otherwise.

### GetAllowRecoveryEmailWithoutEnrollmentOk

`func (o *SsprSettings) GetAllowRecoveryEmailWithoutEnrollmentOk() (*bool, bool)`

GetAllowRecoveryEmailWithoutEnrollmentOk returns a tuple with the AllowRecoveryEmailWithoutEnrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRecoveryEmailWithoutEnrollment

`func (o *SsprSettings) SetAllowRecoveryEmailWithoutEnrollment(v bool)`

SetAllowRecoveryEmailWithoutEnrollment sets AllowRecoveryEmailWithoutEnrollment field to given value.

### HasAllowRecoveryEmailWithoutEnrollment

`func (o *SsprSettings) HasAllowRecoveryEmailWithoutEnrollment() bool`

HasAllowRecoveryEmailWithoutEnrollment returns a boolean if a field has been set.

### SetAllowRecoveryEmailWithoutEnrollmentNil

`func (o *SsprSettings) SetAllowRecoveryEmailWithoutEnrollmentNil(b bool)`

 SetAllowRecoveryEmailWithoutEnrollmentNil sets the value for AllowRecoveryEmailWithoutEnrollment to be an explicit nil

### UnsetAllowRecoveryEmailWithoutEnrollment
`func (o *SsprSettings) UnsetAllowRecoveryEmailWithoutEnrollment()`

UnsetAllowRecoveryEmailWithoutEnrollment ensures that no value is present for AllowRecoveryEmailWithoutEnrollment, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


