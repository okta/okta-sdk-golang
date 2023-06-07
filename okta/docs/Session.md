# Session

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amr** | Pointer to [**[]SessionAuthenticationMethod**](SessionAuthenticationMethod.md) |  | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**ExpiresAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Idp** | Pointer to [**SessionIdentityProvider**](SessionIdentityProvider.md) |  | [optional] 
**LastFactorVerification** | Pointer to **time.Time** |  | [optional] [readonly] 
**LastPasswordVerification** | Pointer to **time.Time** |  | [optional] [readonly] 
**Login** | Pointer to **string** |  | [optional] [readonly] 
**Status** | Pointer to [**SessionStatus**](SessionStatus.md) |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] [readonly] 
**Links** | Pointer to **map[string]map[string]interface{}** |  | [optional] [readonly] 

## Methods

### NewSession

`func NewSession() *Session`

NewSession instantiates a new Session object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionWithDefaults

`func NewSessionWithDefaults() *Session`

NewSessionWithDefaults instantiates a new Session object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmr

`func (o *Session) GetAmr() []SessionAuthenticationMethod`

GetAmr returns the Amr field if non-nil, zero value otherwise.

### GetAmrOk

`func (o *Session) GetAmrOk() (*[]SessionAuthenticationMethod, bool)`

GetAmrOk returns a tuple with the Amr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmr

`func (o *Session) SetAmr(v []SessionAuthenticationMethod)`

SetAmr sets Amr field to given value.

### HasAmr

`func (o *Session) HasAmr() bool`

HasAmr returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Session) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Session) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Session) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Session) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetExpiresAt

`func (o *Session) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *Session) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *Session) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *Session) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetId

`func (o *Session) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Session) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Session) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Session) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdp

`func (o *Session) GetIdp() SessionIdentityProvider`

GetIdp returns the Idp field if non-nil, zero value otherwise.

### GetIdpOk

`func (o *Session) GetIdpOk() (*SessionIdentityProvider, bool)`

GetIdpOk returns a tuple with the Idp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdp

`func (o *Session) SetIdp(v SessionIdentityProvider)`

SetIdp sets Idp field to given value.

### HasIdp

`func (o *Session) HasIdp() bool`

HasIdp returns a boolean if a field has been set.

### GetLastFactorVerification

`func (o *Session) GetLastFactorVerification() time.Time`

GetLastFactorVerification returns the LastFactorVerification field if non-nil, zero value otherwise.

### GetLastFactorVerificationOk

`func (o *Session) GetLastFactorVerificationOk() (*time.Time, bool)`

GetLastFactorVerificationOk returns a tuple with the LastFactorVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFactorVerification

`func (o *Session) SetLastFactorVerification(v time.Time)`

SetLastFactorVerification sets LastFactorVerification field to given value.

### HasLastFactorVerification

`func (o *Session) HasLastFactorVerification() bool`

HasLastFactorVerification returns a boolean if a field has been set.

### GetLastPasswordVerification

`func (o *Session) GetLastPasswordVerification() time.Time`

GetLastPasswordVerification returns the LastPasswordVerification field if non-nil, zero value otherwise.

### GetLastPasswordVerificationOk

`func (o *Session) GetLastPasswordVerificationOk() (*time.Time, bool)`

GetLastPasswordVerificationOk returns a tuple with the LastPasswordVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPasswordVerification

`func (o *Session) SetLastPasswordVerification(v time.Time)`

SetLastPasswordVerification sets LastPasswordVerification field to given value.

### HasLastPasswordVerification

`func (o *Session) HasLastPasswordVerification() bool`

HasLastPasswordVerification returns a boolean if a field has been set.

### GetLogin

`func (o *Session) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *Session) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *Session) SetLogin(v string)`

SetLogin sets Login field to given value.

### HasLogin

`func (o *Session) HasLogin() bool`

HasLogin returns a boolean if a field has been set.

### GetStatus

`func (o *Session) GetStatus() SessionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Session) GetStatusOk() (*SessionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Session) SetStatus(v SessionStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Session) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUserId

`func (o *Session) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Session) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Session) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *Session) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetLinks

`func (o *Session) GetLinks() map[string]map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Session) GetLinksOk() (*map[string]map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Session) SetLinks(v map[string]map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Session) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


