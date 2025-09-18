# AuthenticatorEnrollmentLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Self** | Pointer to [**HrefObjectSelfLink**](HrefObjectSelfLink.md) |  | [optional] 
**User** | Pointer to [**HrefObject**](HrefObject.md) | Returns information on the specified user | [optional] 
**Authenticator** | Pointer to [**HrefObject**](HrefObject.md) | Returns information about a specific authenticator. See [Retrieve an authenticator](/openapi/okta-management/management/tag/Authenticator/#tag/Authenticator/operation/getAuthenticator). | [optional] 

## Methods

### NewAuthenticatorEnrollmentLinks

`func NewAuthenticatorEnrollmentLinks() *AuthenticatorEnrollmentLinks`

NewAuthenticatorEnrollmentLinks instantiates a new AuthenticatorEnrollmentLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticatorEnrollmentLinksWithDefaults

`func NewAuthenticatorEnrollmentLinksWithDefaults() *AuthenticatorEnrollmentLinks`

NewAuthenticatorEnrollmentLinksWithDefaults instantiates a new AuthenticatorEnrollmentLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelf

`func (o *AuthenticatorEnrollmentLinks) GetSelf() HrefObjectSelfLink`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *AuthenticatorEnrollmentLinks) GetSelfOk() (*HrefObjectSelfLink, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *AuthenticatorEnrollmentLinks) SetSelf(v HrefObjectSelfLink)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *AuthenticatorEnrollmentLinks) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetUser

`func (o *AuthenticatorEnrollmentLinks) GetUser() HrefObject`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *AuthenticatorEnrollmentLinks) GetUserOk() (*HrefObject, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *AuthenticatorEnrollmentLinks) SetUser(v HrefObject)`

SetUser sets User field to given value.

### HasUser

`func (o *AuthenticatorEnrollmentLinks) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetAuthenticator

`func (o *AuthenticatorEnrollmentLinks) GetAuthenticator() HrefObject`

GetAuthenticator returns the Authenticator field if non-nil, zero value otherwise.

### GetAuthenticatorOk

`func (o *AuthenticatorEnrollmentLinks) GetAuthenticatorOk() (*HrefObject, bool)`

GetAuthenticatorOk returns a tuple with the Authenticator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticator

`func (o *AuthenticatorEnrollmentLinks) SetAuthenticator(v HrefObject)`

SetAuthenticator sets Authenticator field to given value.

### HasAuthenticator

`func (o *AuthenticatorEnrollmentLinks) HasAuthenticator() bool`

HasAuthenticator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


