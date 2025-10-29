# AuthenticatorEnrollmentPolicyAuthenticatorSettingsConstraints

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AaguidGroups** | Pointer to **[]string** | The list of FIDO2 WebAuthn authenticator groups allowed for enrollment. The authenticators in the group are based on FIDO Alliance Metadata Service that&#39;s identified by name or the Authenticator Attestation Global Unique Identifier ([AAGUID](https://support.yubico.com/hc/en-us/articles/360016648959-YubiKey-Hardware-FIDO2-AAGUIDs)) number. These groups are defined in the [WebAuthn authenticator method settings](https://developer.okta.com/docs/api/openapi/okta-management/management/tag/Authenticator/#tag/Authenticator/operation/listAuthenticatorMethods). | [optional] 

## Methods

### NewAuthenticatorEnrollmentPolicyAuthenticatorSettingsConstraints

`func NewAuthenticatorEnrollmentPolicyAuthenticatorSettingsConstraints() *AuthenticatorEnrollmentPolicyAuthenticatorSettingsConstraints`

NewAuthenticatorEnrollmentPolicyAuthenticatorSettingsConstraints instantiates a new AuthenticatorEnrollmentPolicyAuthenticatorSettingsConstraints object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticatorEnrollmentPolicyAuthenticatorSettingsConstraintsWithDefaults

`func NewAuthenticatorEnrollmentPolicyAuthenticatorSettingsConstraintsWithDefaults() *AuthenticatorEnrollmentPolicyAuthenticatorSettingsConstraints`

NewAuthenticatorEnrollmentPolicyAuthenticatorSettingsConstraintsWithDefaults instantiates a new AuthenticatorEnrollmentPolicyAuthenticatorSettingsConstraints object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAaguidGroups

`func (o *AuthenticatorEnrollmentPolicyAuthenticatorSettingsConstraints) GetAaguidGroups() []string`

GetAaguidGroups returns the AaguidGroups field if non-nil, zero value otherwise.

### GetAaguidGroupsOk

`func (o *AuthenticatorEnrollmentPolicyAuthenticatorSettingsConstraints) GetAaguidGroupsOk() (*[]string, bool)`

GetAaguidGroupsOk returns a tuple with the AaguidGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAaguidGroups

`func (o *AuthenticatorEnrollmentPolicyAuthenticatorSettingsConstraints) SetAaguidGroups(v []string)`

SetAaguidGroups sets AaguidGroups field to given value.

### HasAaguidGroups

`func (o *AuthenticatorEnrollmentPolicyAuthenticatorSettingsConstraints) HasAaguidGroups() bool`

HasAaguidGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


