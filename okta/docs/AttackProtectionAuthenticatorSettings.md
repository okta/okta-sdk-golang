# AttackProtectionAuthenticatorSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VerifyKnowledgeSecondWhen2faRequired** | Pointer to **bool** | If true, requires users to verify a possession factor before verifying a knowledge factor when the assurance requires two-factor authentication (2FA). | [optional] [default to false]

## Methods

### NewAttackProtectionAuthenticatorSettings

`func NewAttackProtectionAuthenticatorSettings() *AttackProtectionAuthenticatorSettings`

NewAttackProtectionAuthenticatorSettings instantiates a new AttackProtectionAuthenticatorSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttackProtectionAuthenticatorSettingsWithDefaults

`func NewAttackProtectionAuthenticatorSettingsWithDefaults() *AttackProtectionAuthenticatorSettings`

NewAttackProtectionAuthenticatorSettingsWithDefaults instantiates a new AttackProtectionAuthenticatorSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVerifyKnowledgeSecondWhen2faRequired

`func (o *AttackProtectionAuthenticatorSettings) GetVerifyKnowledgeSecondWhen2faRequired() bool`

GetVerifyKnowledgeSecondWhen2faRequired returns the VerifyKnowledgeSecondWhen2faRequired field if non-nil, zero value otherwise.

### GetVerifyKnowledgeSecondWhen2faRequiredOk

`func (o *AttackProtectionAuthenticatorSettings) GetVerifyKnowledgeSecondWhen2faRequiredOk() (*bool, bool)`

GetVerifyKnowledgeSecondWhen2faRequiredOk returns a tuple with the VerifyKnowledgeSecondWhen2faRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyKnowledgeSecondWhen2faRequired

`func (o *AttackProtectionAuthenticatorSettings) SetVerifyKnowledgeSecondWhen2faRequired(v bool)`

SetVerifyKnowledgeSecondWhen2faRequired sets VerifyKnowledgeSecondWhen2faRequired field to given value.

### HasVerifyKnowledgeSecondWhen2faRequired

`func (o *AttackProtectionAuthenticatorSettings) HasVerifyKnowledgeSecondWhen2faRequired() bool`

HasVerifyKnowledgeSecondWhen2faRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


