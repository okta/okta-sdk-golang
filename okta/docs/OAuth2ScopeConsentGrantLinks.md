# OAuth2ScopeConsentGrantLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Self** | Pointer to [**HrefObjectSelfLink**](HrefObjectSelfLink.md) |  | [optional] 
**App** | Pointer to [**AppCustomHrefObject**](AppCustomHrefObject.md) |  | [optional] 
**Client** | Pointer to [**AppCustomHrefObject**](AppCustomHrefObject.md) |  | [optional] 

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

`func (o *OAuth2ScopeConsentGrantLinks) GetApp() AppCustomHrefObject`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *OAuth2ScopeConsentGrantLinks) GetAppOk() (*AppCustomHrefObject, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *OAuth2ScopeConsentGrantLinks) SetApp(v AppCustomHrefObject)`

SetApp sets App field to given value.

### HasApp

`func (o *OAuth2ScopeConsentGrantLinks) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetClient

`func (o *OAuth2ScopeConsentGrantLinks) GetClient() AppCustomHrefObject`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *OAuth2ScopeConsentGrantLinks) GetClientOk() (*AppCustomHrefObject, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *OAuth2ScopeConsentGrantLinks) SetClient(v AppCustomHrefObject)`

SetClient sets Client field to given value.

### HasClient

`func (o *OAuth2ScopeConsentGrantLinks) HasClient() bool`

HasClient returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


