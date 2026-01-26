# OAuth2ClientJsonWebKeyRequestBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alg** | Pointer to **NullableString** | Algorithm used in the key | [optional] 
**Kid** | Pointer to **NullableString** | Unique identifier of the JSON Web Key in the OAUth 2.0 client&#39;s JWKS | [optional] 
**Status** | Pointer to **string** | Status of the OAuth 2.0 client JSON Web Key | [optional] [default to "ACTIVE"]
**Use** | Pointer to **NullableString** | Acceptable use of the JSON Web Key | [optional] 

## Methods

### NewOAuth2ClientJsonWebKeyRequestBase

`func NewOAuth2ClientJsonWebKeyRequestBase() *OAuth2ClientJsonWebKeyRequestBase`

NewOAuth2ClientJsonWebKeyRequestBase instantiates a new OAuth2ClientJsonWebKeyRequestBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuth2ClientJsonWebKeyRequestBaseWithDefaults

`func NewOAuth2ClientJsonWebKeyRequestBaseWithDefaults() *OAuth2ClientJsonWebKeyRequestBase`

NewOAuth2ClientJsonWebKeyRequestBaseWithDefaults instantiates a new OAuth2ClientJsonWebKeyRequestBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlg

`func (o *OAuth2ClientJsonWebKeyRequestBase) GetAlg() string`

GetAlg returns the Alg field if non-nil, zero value otherwise.

### GetAlgOk

`func (o *OAuth2ClientJsonWebKeyRequestBase) GetAlgOk() (*string, bool)`

GetAlgOk returns a tuple with the Alg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlg

`func (o *OAuth2ClientJsonWebKeyRequestBase) SetAlg(v string)`

SetAlg sets Alg field to given value.

### HasAlg

`func (o *OAuth2ClientJsonWebKeyRequestBase) HasAlg() bool`

HasAlg returns a boolean if a field has been set.

### SetAlgNil

`func (o *OAuth2ClientJsonWebKeyRequestBase) SetAlgNil(b bool)`

 SetAlgNil sets the value for Alg to be an explicit nil

### UnsetAlg
`func (o *OAuth2ClientJsonWebKeyRequestBase) UnsetAlg()`

UnsetAlg ensures that no value is present for Alg, not even an explicit nil
### GetKid

`func (o *OAuth2ClientJsonWebKeyRequestBase) GetKid() string`

GetKid returns the Kid field if non-nil, zero value otherwise.

### GetKidOk

`func (o *OAuth2ClientJsonWebKeyRequestBase) GetKidOk() (*string, bool)`

GetKidOk returns a tuple with the Kid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKid

`func (o *OAuth2ClientJsonWebKeyRequestBase) SetKid(v string)`

SetKid sets Kid field to given value.

### HasKid

`func (o *OAuth2ClientJsonWebKeyRequestBase) HasKid() bool`

HasKid returns a boolean if a field has been set.

### SetKidNil

`func (o *OAuth2ClientJsonWebKeyRequestBase) SetKidNil(b bool)`

 SetKidNil sets the value for Kid to be an explicit nil

### UnsetKid
`func (o *OAuth2ClientJsonWebKeyRequestBase) UnsetKid()`

UnsetKid ensures that no value is present for Kid, not even an explicit nil
### GetStatus

`func (o *OAuth2ClientJsonWebKeyRequestBase) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OAuth2ClientJsonWebKeyRequestBase) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OAuth2ClientJsonWebKeyRequestBase) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OAuth2ClientJsonWebKeyRequestBase) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUse

`func (o *OAuth2ClientJsonWebKeyRequestBase) GetUse() string`

GetUse returns the Use field if non-nil, zero value otherwise.

### GetUseOk

`func (o *OAuth2ClientJsonWebKeyRequestBase) GetUseOk() (*string, bool)`

GetUseOk returns a tuple with the Use field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUse

`func (o *OAuth2ClientJsonWebKeyRequestBase) SetUse(v string)`

SetUse sets Use field to given value.

### HasUse

`func (o *OAuth2ClientJsonWebKeyRequestBase) HasUse() bool`

HasUse returns a boolean if a field has been set.

### SetUseNil

`func (o *OAuth2ClientJsonWebKeyRequestBase) SetUseNil(b bool)`

 SetUseNil sets the value for Use to be an explicit nil

### UnsetUse
`func (o *OAuth2ClientJsonWebKeyRequestBase) UnsetUse()`

UnsetUse ensures that no value is present for Use, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


