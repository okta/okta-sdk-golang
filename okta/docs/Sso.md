# Sso

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Oidc** | Pointer to [**Oidc**](Oidc.md) |  | [optional] 
**Saml** | Pointer to [**Saml**](Saml.md) |  | [optional] 

## Methods

### NewSso

`func NewSso() *Sso`

NewSso instantiates a new Sso object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSsoWithDefaults

`func NewSsoWithDefaults() *Sso`

NewSsoWithDefaults instantiates a new Sso object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOidc

`func (o *Sso) GetOidc() Oidc`

GetOidc returns the Oidc field if non-nil, zero value otherwise.

### GetOidcOk

`func (o *Sso) GetOidcOk() (*Oidc, bool)`

GetOidcOk returns a tuple with the Oidc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidc

`func (o *Sso) SetOidc(v Oidc)`

SetOidc sets Oidc field to given value.

### HasOidc

`func (o *Sso) HasOidc() bool`

HasOidc returns a boolean if a field has been set.

### GetSaml

`func (o *Sso) GetSaml() Saml`

GetSaml returns the Saml field if non-nil, zero value otherwise.

### GetSamlOk

`func (o *Sso) GetSamlOk() (*Saml, bool)`

GetSamlOk returns a tuple with the Saml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaml

`func (o *Sso) SetSaml(v Saml)`

SetSaml sets Saml field to given value.

### HasSaml

`func (o *Sso) HasSaml() bool`

HasSaml returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


