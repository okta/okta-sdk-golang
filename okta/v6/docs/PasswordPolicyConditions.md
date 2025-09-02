# PasswordPolicyConditions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthProvider** | Pointer to [**PasswordPolicyAuthenticationProviderCondition**](PasswordPolicyAuthenticationProviderCondition.md) |  | [optional] 
**People** | Pointer to [**AuthenticatorEnrollmentPolicyConditionsAllOfPeople**](AuthenticatorEnrollmentPolicyConditionsAllOfPeople.md) |  | [optional] 

## Methods

### NewPasswordPolicyConditions

`func NewPasswordPolicyConditions() *PasswordPolicyConditions`

NewPasswordPolicyConditions instantiates a new PasswordPolicyConditions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPasswordPolicyConditionsWithDefaults

`func NewPasswordPolicyConditionsWithDefaults() *PasswordPolicyConditions`

NewPasswordPolicyConditionsWithDefaults instantiates a new PasswordPolicyConditions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthProvider

`func (o *PasswordPolicyConditions) GetAuthProvider() PasswordPolicyAuthenticationProviderCondition`

GetAuthProvider returns the AuthProvider field if non-nil, zero value otherwise.

### GetAuthProviderOk

`func (o *PasswordPolicyConditions) GetAuthProviderOk() (*PasswordPolicyAuthenticationProviderCondition, bool)`

GetAuthProviderOk returns a tuple with the AuthProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthProvider

`func (o *PasswordPolicyConditions) SetAuthProvider(v PasswordPolicyAuthenticationProviderCondition)`

SetAuthProvider sets AuthProvider field to given value.

### HasAuthProvider

`func (o *PasswordPolicyConditions) HasAuthProvider() bool`

HasAuthProvider returns a boolean if a field has been set.

### GetPeople

`func (o *PasswordPolicyConditions) GetPeople() AuthenticatorEnrollmentPolicyConditionsAllOfPeople`

GetPeople returns the People field if non-nil, zero value otherwise.

### GetPeopleOk

`func (o *PasswordPolicyConditions) GetPeopleOk() (*AuthenticatorEnrollmentPolicyConditionsAllOfPeople, bool)`

GetPeopleOk returns a tuple with the People field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeople

`func (o *PasswordPolicyConditions) SetPeople(v AuthenticatorEnrollmentPolicyConditionsAllOfPeople)`

SetPeople sets People field to given value.

### HasPeople

`func (o *PasswordPolicyConditions) HasPeople() bool`

HasPeople returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


