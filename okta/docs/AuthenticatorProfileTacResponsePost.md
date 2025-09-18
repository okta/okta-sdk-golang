# AuthenticatorProfileTacResponsePost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiresAt** | Pointer to **time.Time** | The time when the TAC enrollment expires in the UTC timezone | [optional] 
**MultiUse** | Pointer to **bool** | Determines whether an enrollment can be used more than once | [optional] 
**Tac** | Pointer to **string** | A temporary access code used for authentication. It can be used one or more times and is valid for a defined period specified by the &#x60;ttl&#x60; property. The &#x60;tac&#x60; is returned in the response when the enrollment is created. It is not returned when the enrollment is retrieved. Issuing a new TAC invalidates any existing TAC for this user. | [optional] 

## Methods

### NewAuthenticatorProfileTacResponsePost

`func NewAuthenticatorProfileTacResponsePost() *AuthenticatorProfileTacResponsePost`

NewAuthenticatorProfileTacResponsePost instantiates a new AuthenticatorProfileTacResponsePost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticatorProfileTacResponsePostWithDefaults

`func NewAuthenticatorProfileTacResponsePostWithDefaults() *AuthenticatorProfileTacResponsePost`

NewAuthenticatorProfileTacResponsePostWithDefaults instantiates a new AuthenticatorProfileTacResponsePost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiresAt

`func (o *AuthenticatorProfileTacResponsePost) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *AuthenticatorProfileTacResponsePost) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *AuthenticatorProfileTacResponsePost) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *AuthenticatorProfileTacResponsePost) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetMultiUse

`func (o *AuthenticatorProfileTacResponsePost) GetMultiUse() bool`

GetMultiUse returns the MultiUse field if non-nil, zero value otherwise.

### GetMultiUseOk

`func (o *AuthenticatorProfileTacResponsePost) GetMultiUseOk() (*bool, bool)`

GetMultiUseOk returns a tuple with the MultiUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiUse

`func (o *AuthenticatorProfileTacResponsePost) SetMultiUse(v bool)`

SetMultiUse sets MultiUse field to given value.

### HasMultiUse

`func (o *AuthenticatorProfileTacResponsePost) HasMultiUse() bool`

HasMultiUse returns a boolean if a field has been set.

### GetTac

`func (o *AuthenticatorProfileTacResponsePost) GetTac() string`

GetTac returns the Tac field if non-nil, zero value otherwise.

### GetTacOk

`func (o *AuthenticatorProfileTacResponsePost) GetTacOk() (*string, bool)`

GetTacOk returns a tuple with the Tac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTac

`func (o *AuthenticatorProfileTacResponsePost) SetTac(v string)`

SetTac sets Tac field to given value.

### HasTac

`func (o *AuthenticatorProfileTacResponsePost) HasTac() bool`

HasTac returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


