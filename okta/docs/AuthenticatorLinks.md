# AuthenticatorLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Self** | Pointer to [**HrefObjectSelfLink**](HrefObjectSelfLink.md) |  | [optional] 
**Activate** | Pointer to [**HrefObjectActivateLink**](HrefObjectActivateLink.md) |  | [optional] 
**Deactivate** | Pointer to [**HrefObjectDeactivateLink**](HrefObjectDeactivateLink.md) |  | [optional] 
**Methods** | Pointer to [**HrefObject**](HrefObject.md) |  | [optional] 

## Methods

### NewAuthenticatorLinks

`func NewAuthenticatorLinks() *AuthenticatorLinks`

NewAuthenticatorLinks instantiates a new AuthenticatorLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticatorLinksWithDefaults

`func NewAuthenticatorLinksWithDefaults() *AuthenticatorLinks`

NewAuthenticatorLinksWithDefaults instantiates a new AuthenticatorLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelf

`func (o *AuthenticatorLinks) GetSelf() HrefObjectSelfLink`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *AuthenticatorLinks) GetSelfOk() (*HrefObjectSelfLink, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *AuthenticatorLinks) SetSelf(v HrefObjectSelfLink)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *AuthenticatorLinks) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetActivate

`func (o *AuthenticatorLinks) GetActivate() HrefObjectActivateLink`

GetActivate returns the Activate field if non-nil, zero value otherwise.

### GetActivateOk

`func (o *AuthenticatorLinks) GetActivateOk() (*HrefObjectActivateLink, bool)`

GetActivateOk returns a tuple with the Activate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivate

`func (o *AuthenticatorLinks) SetActivate(v HrefObjectActivateLink)`

SetActivate sets Activate field to given value.

### HasActivate

`func (o *AuthenticatorLinks) HasActivate() bool`

HasActivate returns a boolean if a field has been set.

### GetDeactivate

`func (o *AuthenticatorLinks) GetDeactivate() HrefObjectDeactivateLink`

GetDeactivate returns the Deactivate field if non-nil, zero value otherwise.

### GetDeactivateOk

`func (o *AuthenticatorLinks) GetDeactivateOk() (*HrefObjectDeactivateLink, bool)`

GetDeactivateOk returns a tuple with the Deactivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeactivate

`func (o *AuthenticatorLinks) SetDeactivate(v HrefObjectDeactivateLink)`

SetDeactivate sets Deactivate field to given value.

### HasDeactivate

`func (o *AuthenticatorLinks) HasDeactivate() bool`

HasDeactivate returns a boolean if a field has been set.

### GetMethods

`func (o *AuthenticatorLinks) GetMethods() HrefObject`

GetMethods returns the Methods field if non-nil, zero value otherwise.

### GetMethodsOk

`func (o *AuthenticatorLinks) GetMethodsOk() (*HrefObject, bool)`

GetMethodsOk returns a tuple with the Methods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethods

`func (o *AuthenticatorLinks) SetMethods(v HrefObject)`

SetMethods sets Methods field to given value.

### HasMethods

`func (o *AuthenticatorLinks) HasMethods() bool`

HasMethods returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


