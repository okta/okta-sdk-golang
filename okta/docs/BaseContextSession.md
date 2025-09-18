# BaseContextSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique identifier for the user&#39;s session | [optional] 
**UserId** | Pointer to **string** | The unique identifier for the user | [optional] 
**Login** | Pointer to **string** | The username used to identify the user. This is often the user&#39;s email address. | [optional] 
**CreatedAt** | Pointer to **time.Time** | Timestamp of when the session was created | [optional] 
**ExpiresAt** | Pointer to **time.Time** | Timestamp of when the session expires | [optional] 
**Status** | Pointer to **string** | Represents the current status of the user&#39;s session | [optional] 
**LastPasswordVerification** | Pointer to **time.Time** | Timestamp of when the user was last authenticated | [optional] 
**Amr** | Pointer to **[]string** | The authentication method reference | [optional] 
**Idp** | Pointer to [**SessionIdentityProvider**](SessionIdentityProvider.md) |  | [optional] 
**MfaActive** | Pointer to **bool** | Describes whether multifactor authentication was enabled | [optional] 

## Methods

### NewBaseContextSession

`func NewBaseContextSession() *BaseContextSession`

NewBaseContextSession instantiates a new BaseContextSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseContextSessionWithDefaults

`func NewBaseContextSessionWithDefaults() *BaseContextSession`

NewBaseContextSessionWithDefaults instantiates a new BaseContextSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BaseContextSession) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BaseContextSession) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BaseContextSession) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BaseContextSession) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUserId

`func (o *BaseContextSession) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *BaseContextSession) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *BaseContextSession) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *BaseContextSession) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetLogin

`func (o *BaseContextSession) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *BaseContextSession) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *BaseContextSession) SetLogin(v string)`

SetLogin sets Login field to given value.

### HasLogin

`func (o *BaseContextSession) HasLogin() bool`

HasLogin returns a boolean if a field has been set.

### GetCreatedAt

`func (o *BaseContextSession) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BaseContextSession) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BaseContextSession) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BaseContextSession) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetExpiresAt

`func (o *BaseContextSession) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *BaseContextSession) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *BaseContextSession) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *BaseContextSession) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetStatus

`func (o *BaseContextSession) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BaseContextSession) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BaseContextSession) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BaseContextSession) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLastPasswordVerification

`func (o *BaseContextSession) GetLastPasswordVerification() time.Time`

GetLastPasswordVerification returns the LastPasswordVerification field if non-nil, zero value otherwise.

### GetLastPasswordVerificationOk

`func (o *BaseContextSession) GetLastPasswordVerificationOk() (*time.Time, bool)`

GetLastPasswordVerificationOk returns a tuple with the LastPasswordVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPasswordVerification

`func (o *BaseContextSession) SetLastPasswordVerification(v time.Time)`

SetLastPasswordVerification sets LastPasswordVerification field to given value.

### HasLastPasswordVerification

`func (o *BaseContextSession) HasLastPasswordVerification() bool`

HasLastPasswordVerification returns a boolean if a field has been set.

### GetAmr

`func (o *BaseContextSession) GetAmr() []string`

GetAmr returns the Amr field if non-nil, zero value otherwise.

### GetAmrOk

`func (o *BaseContextSession) GetAmrOk() (*[]string, bool)`

GetAmrOk returns a tuple with the Amr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmr

`func (o *BaseContextSession) SetAmr(v []string)`

SetAmr sets Amr field to given value.

### HasAmr

`func (o *BaseContextSession) HasAmr() bool`

HasAmr returns a boolean if a field has been set.

### GetIdp

`func (o *BaseContextSession) GetIdp() SessionIdentityProvider`

GetIdp returns the Idp field if non-nil, zero value otherwise.

### GetIdpOk

`func (o *BaseContextSession) GetIdpOk() (*SessionIdentityProvider, bool)`

GetIdpOk returns a tuple with the Idp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdp

`func (o *BaseContextSession) SetIdp(v SessionIdentityProvider)`

SetIdp sets Idp field to given value.

### HasIdp

`func (o *BaseContextSession) HasIdp() bool`

HasIdp returns a boolean if a field has been set.

### GetMfaActive

`func (o *BaseContextSession) GetMfaActive() bool`

GetMfaActive returns the MfaActive field if non-nil, zero value otherwise.

### GetMfaActiveOk

`func (o *BaseContextSession) GetMfaActiveOk() (*bool, bool)`

GetMfaActiveOk returns a tuple with the MfaActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaActive

`func (o *BaseContextSession) SetMfaActive(v bool)`

SetMfaActive sets MfaActive field to given value.

### HasMfaActive

`func (o *BaseContextSession) HasMfaActive() bool`

HasMfaActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


