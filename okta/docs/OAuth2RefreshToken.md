# OAuth2RefreshToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | Client ID | [optional] 
**Created** | Pointer to **time.Time** | Timestamp when the object was created | [optional] [readonly] 
**ExpiresAt** | Pointer to **time.Time** | Expiration time of the OAuth 2.0 Token | [optional] [readonly] 
**Id** | Pointer to **string** | ID of the Token object | [optional] [readonly] 
**Issuer** | Pointer to **string** | The complete URL of the authorization server that issued the Token | [optional] 
**LastUpdated** | Pointer to **time.Time** | Timestamp when the object was last updated | [optional] [readonly] 
**Scopes** | Pointer to **[]string** | The scope names attached to the Token | [optional] 
**Status** | Pointer to **string** | Status | [optional] [readonly] 
**UserId** | Pointer to **string** | The ID of the user associated with the Token | [optional] 
**Embedded** | Pointer to [**OAuth2RefreshTokenEmbedded**](OAuth2RefreshTokenEmbedded.md) |  | [optional] 
**Links** | Pointer to [**OAuth2RefreshTokenLinks**](OAuth2RefreshTokenLinks.md) |  | [optional] 

## Methods

### NewOAuth2RefreshToken

`func NewOAuth2RefreshToken() *OAuth2RefreshToken`

NewOAuth2RefreshToken instantiates a new OAuth2RefreshToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuth2RefreshTokenWithDefaults

`func NewOAuth2RefreshTokenWithDefaults() *OAuth2RefreshToken`

NewOAuth2RefreshTokenWithDefaults instantiates a new OAuth2RefreshToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *OAuth2RefreshToken) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *OAuth2RefreshToken) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *OAuth2RefreshToken) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *OAuth2RefreshToken) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetCreated

`func (o *OAuth2RefreshToken) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *OAuth2RefreshToken) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *OAuth2RefreshToken) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *OAuth2RefreshToken) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetExpiresAt

`func (o *OAuth2RefreshToken) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *OAuth2RefreshToken) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *OAuth2RefreshToken) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *OAuth2RefreshToken) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetId

`func (o *OAuth2RefreshToken) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OAuth2RefreshToken) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OAuth2RefreshToken) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OAuth2RefreshToken) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIssuer

`func (o *OAuth2RefreshToken) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *OAuth2RefreshToken) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *OAuth2RefreshToken) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *OAuth2RefreshToken) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetLastUpdated

`func (o *OAuth2RefreshToken) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *OAuth2RefreshToken) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *OAuth2RefreshToken) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *OAuth2RefreshToken) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetScopes

`func (o *OAuth2RefreshToken) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *OAuth2RefreshToken) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *OAuth2RefreshToken) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *OAuth2RefreshToken) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetStatus

`func (o *OAuth2RefreshToken) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OAuth2RefreshToken) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OAuth2RefreshToken) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OAuth2RefreshToken) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUserId

`func (o *OAuth2RefreshToken) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *OAuth2RefreshToken) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *OAuth2RefreshToken) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *OAuth2RefreshToken) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetEmbedded

`func (o *OAuth2RefreshToken) GetEmbedded() OAuth2RefreshTokenEmbedded`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *OAuth2RefreshToken) GetEmbeddedOk() (*OAuth2RefreshTokenEmbedded, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *OAuth2RefreshToken) SetEmbedded(v OAuth2RefreshTokenEmbedded)`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *OAuth2RefreshToken) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetLinks

`func (o *OAuth2RefreshToken) GetLinks() OAuth2RefreshTokenLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *OAuth2RefreshToken) GetLinksOk() (*OAuth2RefreshTokenLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *OAuth2RefreshToken) SetLinks(v OAuth2RefreshTokenLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *OAuth2RefreshToken) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


