# IdentityProviderApplicationUserLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Self** | Pointer to [**HrefObjectSelfLink**](HrefObjectSelfLink.md) |  | [optional] 
**Next** | Pointer to [**HrefObject**](HrefObject.md) |  | [optional] 
**Idp** | Pointer to [**HrefObject**](HrefObject.md) | The IdP instance | [optional] 
**User** | Pointer to [**HrefObject**](HrefObject.md) | The linked Okta user | [optional] 

## Methods

### NewIdentityProviderApplicationUserLinks

`func NewIdentityProviderApplicationUserLinks() *IdentityProviderApplicationUserLinks`

NewIdentityProviderApplicationUserLinks instantiates a new IdentityProviderApplicationUserLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityProviderApplicationUserLinksWithDefaults

`func NewIdentityProviderApplicationUserLinksWithDefaults() *IdentityProviderApplicationUserLinks`

NewIdentityProviderApplicationUserLinksWithDefaults instantiates a new IdentityProviderApplicationUserLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelf

`func (o *IdentityProviderApplicationUserLinks) GetSelf() HrefObjectSelfLink`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *IdentityProviderApplicationUserLinks) GetSelfOk() (*HrefObjectSelfLink, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *IdentityProviderApplicationUserLinks) SetSelf(v HrefObjectSelfLink)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *IdentityProviderApplicationUserLinks) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetNext

`func (o *IdentityProviderApplicationUserLinks) GetNext() HrefObject`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *IdentityProviderApplicationUserLinks) GetNextOk() (*HrefObject, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *IdentityProviderApplicationUserLinks) SetNext(v HrefObject)`

SetNext sets Next field to given value.

### HasNext

`func (o *IdentityProviderApplicationUserLinks) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetIdp

`func (o *IdentityProviderApplicationUserLinks) GetIdp() HrefObject`

GetIdp returns the Idp field if non-nil, zero value otherwise.

### GetIdpOk

`func (o *IdentityProviderApplicationUserLinks) GetIdpOk() (*HrefObject, bool)`

GetIdpOk returns a tuple with the Idp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdp

`func (o *IdentityProviderApplicationUserLinks) SetIdp(v HrefObject)`

SetIdp sets Idp field to given value.

### HasIdp

`func (o *IdentityProviderApplicationUserLinks) HasIdp() bool`

HasIdp returns a boolean if a field has been set.

### GetUser

`func (o *IdentityProviderApplicationUserLinks) GetUser() HrefObject`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *IdentityProviderApplicationUserLinks) GetUserOk() (*HrefObject, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *IdentityProviderApplicationUserLinks) SetUser(v HrefObject)`

SetUser sets User field to given value.

### HasUser

`func (o *IdentityProviderApplicationUserLinks) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


