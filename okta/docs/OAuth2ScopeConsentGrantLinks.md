# OAuth2ScopeConsentGrantLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Self** | Pointer to [**HrefObjectSelfLink**](HrefObjectSelfLink.md) |  | [optional] 
**App** | Pointer to [**AppResourceHrefObject**](AppResourceHrefObject.md) | Link to the app resource | [optional] 
**Client** | Pointer to [**AppResourceHrefObject**](AppResourceHrefObject.md) | Link to the client resource | [optional] 
**Scope** | Pointer to [**ScopeResourceHrefObject**](ScopeResourceHrefObject.md) | Link to the scope resource | [optional] 
**User** | Pointer to [**UserResourceHrefObject**](UserResourceHrefObject.md) | Link to the user resource | [optional] 
**AuthorizationServer** | Pointer to [**AuthorizationServerResourceHrefObject**](AuthorizationServerResourceHrefObject.md) | Link to the authorization server resource | [optional] 

## Methods

### NewOAuth2ScopeConsentGrantLinks

`func NewOAuth2ScopeConsentGrantLinks() *OAuth2ScopeConsentGrantLinks`

NewOAuth2ScopeConsentGrantLinks instantiates a new OAuth2ScopeConsentGrantLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuth2ScopeConsentGrantLinksWithDefaults

`func NewOAuth2ScopeConsentGrantLinksWithDefaults() *OAuth2ScopeConsentGrantLinks`

NewOAuth2ScopeConsentGrantLinksWithDefaults instantiates a new OAuth2ScopeConsentGrantLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelf

`func (o *OAuth2ScopeConsentGrantLinks) GetSelf() HrefObjectSelfLink`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *OAuth2ScopeConsentGrantLinks) GetSelfOk() (*HrefObjectSelfLink, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *OAuth2ScopeConsentGrantLinks) SetSelf(v HrefObjectSelfLink)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *OAuth2ScopeConsentGrantLinks) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetApp

`func (o *OAuth2ScopeConsentGrantLinks) GetApp() AppResourceHrefObject`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *OAuth2ScopeConsentGrantLinks) GetAppOk() (*AppResourceHrefObject, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *OAuth2ScopeConsentGrantLinks) SetApp(v AppResourceHrefObject)`

SetApp sets App field to given value.

### HasApp

`func (o *OAuth2ScopeConsentGrantLinks) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetClient

`func (o *OAuth2ScopeConsentGrantLinks) GetClient() AppResourceHrefObject`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *OAuth2ScopeConsentGrantLinks) GetClientOk() (*AppResourceHrefObject, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *OAuth2ScopeConsentGrantLinks) SetClient(v AppResourceHrefObject)`

SetClient sets Client field to given value.

### HasClient

`func (o *OAuth2ScopeConsentGrantLinks) HasClient() bool`

HasClient returns a boolean if a field has been set.

### GetScope

`func (o *OAuth2ScopeConsentGrantLinks) GetScope() ScopeResourceHrefObject`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *OAuth2ScopeConsentGrantLinks) GetScopeOk() (*ScopeResourceHrefObject, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *OAuth2ScopeConsentGrantLinks) SetScope(v ScopeResourceHrefObject)`

SetScope sets Scope field to given value.

### HasScope

`func (o *OAuth2ScopeConsentGrantLinks) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetUser

`func (o *OAuth2ScopeConsentGrantLinks) GetUser() UserResourceHrefObject`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *OAuth2ScopeConsentGrantLinks) GetUserOk() (*UserResourceHrefObject, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *OAuth2ScopeConsentGrantLinks) SetUser(v UserResourceHrefObject)`

SetUser sets User field to given value.

### HasUser

`func (o *OAuth2ScopeConsentGrantLinks) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetAuthorizationServer

`func (o *OAuth2ScopeConsentGrantLinks) GetAuthorizationServer() AuthorizationServerResourceHrefObject`

GetAuthorizationServer returns the AuthorizationServer field if non-nil, zero value otherwise.

### GetAuthorizationServerOk

`func (o *OAuth2ScopeConsentGrantLinks) GetAuthorizationServerOk() (*AuthorizationServerResourceHrefObject, bool)`

GetAuthorizationServerOk returns a tuple with the AuthorizationServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationServer

`func (o *OAuth2ScopeConsentGrantLinks) SetAuthorizationServer(v AuthorizationServerResourceHrefObject)`

SetAuthorizationServer sets AuthorizationServer field to given value.

### HasAuthorizationServer

`func (o *OAuth2ScopeConsentGrantLinks) HasAuthorizationServer() bool`

HasAuthorizationServer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


