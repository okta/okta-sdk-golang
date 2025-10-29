# OAuth2ClientLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Self** | Pointer to [**HrefObjectSelfLink**](HrefObjectSelfLink.md) |  | [optional] 
**Grants** | Pointer to [**GrantResourcesHrefObject**](GrantResourcesHrefObject.md) | Link to the grant resources | [optional] 
**Tokens** | Pointer to [**TokenResourcesHrefObject**](TokenResourcesHrefObject.md) | Link to the token resources | [optional] 

## Methods

### NewOAuth2ClientLinks

`func NewOAuth2ClientLinks() *OAuth2ClientLinks`

NewOAuth2ClientLinks instantiates a new OAuth2ClientLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuth2ClientLinksWithDefaults

`func NewOAuth2ClientLinksWithDefaults() *OAuth2ClientLinks`

NewOAuth2ClientLinksWithDefaults instantiates a new OAuth2ClientLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelf

`func (o *OAuth2ClientLinks) GetSelf() HrefObjectSelfLink`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *OAuth2ClientLinks) GetSelfOk() (*HrefObjectSelfLink, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *OAuth2ClientLinks) SetSelf(v HrefObjectSelfLink)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *OAuth2ClientLinks) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetGrants

`func (o *OAuth2ClientLinks) GetGrants() GrantResourcesHrefObject`

GetGrants returns the Grants field if non-nil, zero value otherwise.

### GetGrantsOk

`func (o *OAuth2ClientLinks) GetGrantsOk() (*GrantResourcesHrefObject, bool)`

GetGrantsOk returns a tuple with the Grants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrants

`func (o *OAuth2ClientLinks) SetGrants(v GrantResourcesHrefObject)`

SetGrants sets Grants field to given value.

### HasGrants

`func (o *OAuth2ClientLinks) HasGrants() bool`

HasGrants returns a boolean if a field has been set.

### GetTokens

`func (o *OAuth2ClientLinks) GetTokens() TokenResourcesHrefObject`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *OAuth2ClientLinks) GetTokensOk() (*TokenResourcesHrefObject, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *OAuth2ClientLinks) SetTokens(v TokenResourcesHrefObject)`

SetTokens sets Tokens field to given value.

### HasTokens

`func (o *OAuth2ClientLinks) HasTokens() bool`

HasTokens returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


