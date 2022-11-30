# PasswordPolicyConditionsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthProvider** | Pointer to [**PasswordPolicyAuthenticationProviderCondition**](PasswordPolicyAuthenticationProviderCondition.md) |  | [optional] 
**People** | Pointer to [**PolicyPeopleCondition**](PolicyPeopleCondition.md) |  | [optional] 

## Methods

### NewPasswordPolicyConditionsAllOf

`func NewPasswordPolicyConditionsAllOf() *PasswordPolicyConditionsAllOf`

NewPasswordPolicyConditionsAllOf instantiates a new PasswordPolicyConditionsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPasswordPolicyConditionsAllOfWithDefaults

`func NewPasswordPolicyConditionsAllOfWithDefaults() *PasswordPolicyConditionsAllOf`

NewPasswordPolicyConditionsAllOfWithDefaults instantiates a new PasswordPolicyConditionsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthProvider

`func (o *PasswordPolicyConditionsAllOf) GetAuthProvider() PasswordPolicyAuthenticationProviderCondition`

GetAuthProvider returns the AuthProvider field if non-nil, zero value otherwise.

### GetAuthProviderOk

`func (o *PasswordPolicyConditionsAllOf) GetAuthProviderOk() (*PasswordPolicyAuthenticationProviderCondition, bool)`

GetAuthProviderOk returns a tuple with the AuthProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthProvider

`func (o *PasswordPolicyConditionsAllOf) SetAuthProvider(v PasswordPolicyAuthenticationProviderCondition)`

SetAuthProvider sets AuthProvider field to given value.

### HasAuthProvider

`func (o *PasswordPolicyConditionsAllOf) HasAuthProvider() bool`

HasAuthProvider returns a boolean if a field has been set.

### GetPeople

`func (o *PasswordPolicyConditionsAllOf) GetPeople() PolicyPeopleCondition`

GetPeople returns the People field if non-nil, zero value otherwise.

### GetPeopleOk

`func (o *PasswordPolicyConditionsAllOf) GetPeopleOk() (*PolicyPeopleCondition, bool)`

GetPeopleOk returns a tuple with the People field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeople

`func (o *PasswordPolicyConditionsAllOf) SetPeople(v PolicyPeopleCondition)`

SetPeople sets People field to given value.

### HasPeople

`func (o *PasswordPolicyConditionsAllOf) HasPeople() bool`

HasPeople returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


