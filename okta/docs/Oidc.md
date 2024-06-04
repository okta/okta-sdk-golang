# Oidc

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Doc** | **string** | The URL to your customer-facing instructions for configuring your OIDC integration. See [Customer configuration document guidelines](https://developer.okta.com/docs/guides/submit-app-prereq/main/#customer-configuration-document-guidelines). | 
**InitiateLoginUri** | Pointer to **string** | The URL to redirect users when they click on your app from their Okta End-User Dashboard | [optional] 
**PostLogoutUris** | Pointer to **[]string** | The sign-out redirect URIs for your app. You can send a request to &#x60;/v1/logout&#x60; to sign the user out and redirect them to one of these URIs. | [optional] 
**RedirectUris** | **[]string** | List of sign-in redirect URIs | 

## Methods

### NewOidc

`func NewOidc(doc string, redirectUris []string, ) *Oidc`

NewOidc instantiates a new Oidc object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOidcWithDefaults

`func NewOidcWithDefaults() *Oidc`

NewOidcWithDefaults instantiates a new Oidc object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDoc

`func (o *Oidc) GetDoc() string`

GetDoc returns the Doc field if non-nil, zero value otherwise.

### GetDocOk

`func (o *Oidc) GetDocOk() (*string, bool)`

GetDocOk returns a tuple with the Doc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoc

`func (o *Oidc) SetDoc(v string)`

SetDoc sets Doc field to given value.


### GetInitiateLoginUri

`func (o *Oidc) GetInitiateLoginUri() string`

GetInitiateLoginUri returns the InitiateLoginUri field if non-nil, zero value otherwise.

### GetInitiateLoginUriOk

`func (o *Oidc) GetInitiateLoginUriOk() (*string, bool)`

GetInitiateLoginUriOk returns a tuple with the InitiateLoginUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiateLoginUri

`func (o *Oidc) SetInitiateLoginUri(v string)`

SetInitiateLoginUri sets InitiateLoginUri field to given value.

### HasInitiateLoginUri

`func (o *Oidc) HasInitiateLoginUri() bool`

HasInitiateLoginUri returns a boolean if a field has been set.

### GetPostLogoutUris

`func (o *Oidc) GetPostLogoutUris() []string`

GetPostLogoutUris returns the PostLogoutUris field if non-nil, zero value otherwise.

### GetPostLogoutUrisOk

`func (o *Oidc) GetPostLogoutUrisOk() (*[]string, bool)`

GetPostLogoutUrisOk returns a tuple with the PostLogoutUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostLogoutUris

`func (o *Oidc) SetPostLogoutUris(v []string)`

SetPostLogoutUris sets PostLogoutUris field to given value.

### HasPostLogoutUris

`func (o *Oidc) HasPostLogoutUris() bool`

HasPostLogoutUris returns a boolean if a field has been set.

### GetRedirectUris

`func (o *Oidc) GetRedirectUris() []string`

GetRedirectUris returns the RedirectUris field if non-nil, zero value otherwise.

### GetRedirectUrisOk

`func (o *Oidc) GetRedirectUrisOk() (*[]string, bool)`

GetRedirectUrisOk returns a tuple with the RedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUris

`func (o *Oidc) SetRedirectUris(v []string)`

SetRedirectUris sets RedirectUris field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


